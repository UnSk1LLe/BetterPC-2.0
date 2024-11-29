package handlers

import (
	"BetterPC_2.0/internal/handlers/helpers/responseManager"
	"BetterPC_2.0/internal/middlewares/helpers/userContext"
	"BetterPC_2.0/pkg/data/models/orders"
	orderErrors "BetterPC_2.0/pkg/data/models/orders/errors"
	orderRequests "BetterPC_2.0/pkg/data/models/orders/requests"
	productErrors "BetterPC_2.0/pkg/data/models/products/errors"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
	"sync"
)

func (h *Handler) CreateOrderWithItemHeaders(c *gin.Context) {
	//get userId
	user, err := userContext.GetUserCtx(c)
	if err != nil {
		responseManager.ErrorResponseWithLog(c, http.StatusUnauthorized, err.Error())
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg := sync.WaitGroup{}
	wg.Add(2)

	cartChan := make(chan orderRequests.CreateOrderRequest, 1)
	verResChan := make(chan VerificationCheckResult, 1)
	errChan := make(chan error, 1) // Channel to communicate errors

	//get user cart
	go func() {
		defer wg.Done()
		var cart orderRequests.CreateOrderRequest

		//cart, err := cookies.GetCart(c, userData.ID)
		err := c.ShouldBindJSON(&cart)
		if err != nil {
			errChan <- err
			cancel()
			return
		}

		err = cart.Validate()
		if err != nil {
			errChan <- err
			cancel()
			return
		}

		select {
		case cartChan <- cart:
		case <-ctx.Done():
		}

	}()

	go func() {
		defer wg.Done()
		isVerified, err := h.services.Verification.IsVerifiedUser(user.Email)
		if err != nil {
			errChan <- err
			cancel()
			return
		}
		select {
		case verResChan <- VerificationCheckResult{isVerified, err}:
		case <-ctx.Done():
		}

	}()

	go func() {
		wg.Wait()
		close(cartChan)
		close(verResChan)
		close(errChan)
	}()

	var request orderRequests.CreateOrderRequest

	select {
	case cartResult := <-cartChan:
		request = cartResult

	case verResult := <-verResChan:
		if verResult.Err != nil {
			responseManager.ErrorResponseWithLog(c, http.StatusInternalServerError, verResult.Err.Error())
			return
		}

		if !verResult.IsVerified {
			responseManager.ErrorResponseWithLog(c, http.StatusForbidden, verResult.Err.Error())
			return
		}

	case err := <-errChan:
		if errors.As(err, &orderErrors.OrderError{}) {
			message := fmt.Sprintf("Cannot create order: %s", err.Error())
			responseManager.ErrorResponseWithLog(c, http.StatusBadRequest, err.Error(), message)
			return
		}
		responseManager.ErrorResponseWithLog(c, http.StatusInternalServerError, err.Error(), "An error occurred: "+err.Error())
		return

	case <-ctx.Done():
		// Handle cancellation if needed
		fmt.Println("Operation canceled")
		return

	}

	//creating order
	orderId, err := h.services.Order.CreateWithItemHeaders(user.ID, request)
	if err != nil {
		switch {
		case errors.As(err, &productErrors.ProductError{}) || errors.As(err, &orderErrors.OrderError{}):
			responseManager.ErrorResponseWithLog(c, http.StatusConflict, err.Error())
			return
		}
		responseManager.ErrorResponseWithLog(c, http.StatusInternalServerError, err.Error())
		return
	}

	//server response
	logMessage := fmt.Sprintf("order with ID %s was CREATED successfully", orderId.Hex())
	message := "Order created successfully!"
	responseManager.MessageResponseWithLog(c, http.StatusCreated, logMessage, message)
}

func (h *Handler) ListUserOrders(c *gin.Context) {
	var orderList []orders.Order

	user, err := userContext.GetUserCtx(c)
	if err != nil {
		responseManager.ErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	orderList, err = h.services.Order.GetUserOrders(user.ID)
	if err != nil {
		switch {
		case errors.As(err, &orderErrors.OrderError{}):
			responseManager.ErrorResponse(c, http.StatusConflict, err.Error())
			return
		}
		responseManager.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, orderList)
	return
}

func (h *Handler) GetUserOrder(c *gin.Context) {
	var orderList orders.Order

	orderId := c.Param(":id")
	if orderId == "" {
		responseManager.ErrorResponse(c, http.StatusBadRequest, "order id required")
		return
	}

	user, err := userContext.GetUserCtx(c)
	if err != nil {
		responseManager.ErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	orderList, err = h.services.Order.GetUserOrder(user.ID, orderId)
	if err != nil {
		switch {
		case errors.As(err, &orderErrors.OrderError{}):
			responseManager.ErrorResponse(c, http.StatusConflict, err.Error())
			return
		}
		responseManager.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, orderList)
	return
}

func (h *Handler) ProcessOrderPayment(c *gin.Context) {
	type PaymentRequest struct {
		Amount          int64  `json:"amount"`
		Currency        string `json:"currency"`
		PaymentMethodId string `json:"payment_method_id"`
	}

	user, err := userContext.GetUserCtx(c)
	if err != nil {
		responseManager.ErrorResponse(c, http.StatusUnauthorized, "unauthorized")
		return
	}

	orderId := c.Param("id")
	if orderId == "" {
		responseManager.ErrorResponse(c, http.StatusBadRequest, "no order id provided")
		return
	}

	var req PaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		responseManager.ErrorResponse(c, http.StatusBadRequest, "invalid request")
		return
	}

	host := c.Request.URL.Host
	endpoint := "shop/orders/"
	returnUrl := "http://localhost:8080/shop/orders/"
	logrus.Info(host, " ", endpoint)
	paymentIntentId, err := h.services.Order.PayForOrder(user.ID, orderId, req.Amount, req.Currency, req.PaymentMethodId, returnUrl)
	if err != nil {
		switch {
		case errors.As(err, &orderErrors.OrderError{}):
			responseManager.ErrorResponse(c, http.StatusConflict, err.Error())
			return
		default:
			responseManager.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":           true,
		"payment_intent_id": paymentIntentId,
	})
}
