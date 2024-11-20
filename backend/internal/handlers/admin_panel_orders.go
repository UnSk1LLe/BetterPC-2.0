package handlers

import (
	"BetterPC_2.0/internal/handlers/helpers/responseManager"
	"BetterPC_2.0/internal/middlewares"
	"BetterPC_2.0/pkg/data/models/orders"
	orderErrors "BetterPC_2.0/pkg/data/models/orders/errors"
	orderFilters "BetterPC_2.0/pkg/data/models/orders/filters"
	"BetterPC_2.0/pkg/data/models/orders/requests"
	orderRequests "BetterPC_2.0/pkg/data/models/orders/requests"
	productErrors "BetterPC_2.0/pkg/data/models/products/errors"
	userResponses "BetterPC_2.0/pkg/data/models/users/responses"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"sync"
)

type VerificationCheckResult struct {
	IsVerified bool
	Err        error
}

func (h *Handler) CreateOrderWithItemHeaders(c *gin.Context) {
	//get userId
	user, ok := c.Get(middlewares.UserCtx)
	if !ok {
		logMessage := "empty user context"
		responseManager.ErrorResponseWithLog(c, http.StatusUnauthorized, logMessage)
		return
	}

	userData, ok := user.(userResponses.UserResponse)
	if !ok {
		logMessage := "invalid user response"
		responseManager.ErrorResponseWithLog(c, http.StatusBadRequest, logMessage)
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
		isVerified, err := h.services.Verification.IsVerifiedUser(userData.Email)
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

	var request requests.CreateOrderRequest

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
	orderId, err := h.services.Order.CreateWithItemHeaders(userData.ID, request)
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

func (h *Handler) SetStatus(c *gin.Context) {
	status := c.Param("status")
	if status == "" {
		responseManager.ErrorResponse(c, http.StatusBadRequest, "empty status parameter")
		return
	}

	orderId := c.Param("id")
	if orderId == "" {
		responseManager.ErrorResponse(c, http.StatusBadRequest, "empty order id")
		return
	}

	err := h.services.Order.SetStatus(orderId, status)
	if err != nil {
		switch {
		case errors.As(err, &orderErrors.OrderError{}):
			responseManager.ErrorResponseWithLog(c, http.StatusConflict, err.Error(), "An error occurred: "+err.Error())
			return
		}
		responseManager.ErrorResponseWithLog(c, http.StatusInternalServerError, err.Error(), "An error occurred: "+err.Error())
		return
	}

	message := fmt.Sprintf("Order status changed to '%s'", status)
	responseManager.MessageResponse(c, http.StatusOK, message)
	return
}

func (h *Handler) CancelOrder(c *gin.Context) {
	orderId := c.Param("id")
	if orderId == "" {
		responseManager.ErrorResponseWithLog(c, http.StatusBadRequest, "empty order id")
		return
	}

	err := h.services.CancelOrder(orderId)
	if err != nil {
		switch {
		case errors.As(err, &orderErrors.OrderError{}):
			responseManager.ErrorResponseWithLog(c, http.StatusConflict, err.Error())
			return
		}
		responseManager.ErrorResponseWithLog(c, http.StatusInternalServerError, err.Error())
		return
	}

	logMessage := fmt.Sprintf("order with ID %s was CANCELED successfully", orderId)
	message := "Order canceled successfully!"
	responseManager.MessageResponseWithLog(c, http.StatusOK, logMessage, message)
	return
}

func (h *Handler) DeleteOrder(c *gin.Context) {
	orderId := c.Param("id")
	if orderId == "" {
		responseManager.ErrorResponseWithLog(c, http.StatusBadRequest, "empty order id")
		return
	}

	err := h.services.Order.Delete(orderId)
	if err != nil {
		switch {
		case errors.As(err, &orderErrors.OrderError{}):
			responseManager.ErrorResponseWithLog(c, http.StatusConflict, err.Error())
			return
		}
		responseManager.ErrorResponseWithLog(c, http.StatusInternalServerError, err.Error())
		return
	}
	logMessage := fmt.Sprintf("order with ID %s was DELETED successfully", orderId)
	message := "Order deleted successfully!"
	responseManager.MessageResponseWithLog(c, http.StatusOK, logMessage, message)
}

func (h *Handler) ListOrders(c *gin.Context) {
	//getfilters
	var filters orderFilters.AdminOrderFilters
	if err := c.ShouldBindJSON(&filters); err != nil {
		logMessage := fmt.Sprintf("invalid request body: %s", err.Error())
		responseManager.ErrorResponseWithLog(c, http.StatusBadRequest, logMessage)
		return
	}

	var orderList []orders.Order

	orderList, err := h.services.Order.GetList(filters)
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

func (h *Handler) OrderDetails(c *gin.Context) {
	orderId := c.Param("id")
	if orderId == "" {
		responseManager.ErrorResponse(c, http.StatusBadRequest, "empty order id")
		return
	}

	var order orders.Order

	order, err := h.services.Order.GetById(orderId)
	if err != nil {
		switch {
		case errors.As(err, &orderErrors.OrderError{}):
			responseManager.ErrorResponse(c, http.StatusConflict, err.Error())
			return
		}
		responseManager.ErrorResponseWithLog(c, http.StatusInternalServerError, err.Error(), "An error occurred: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, order)
	return
}
