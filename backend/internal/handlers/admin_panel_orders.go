package handlers

import (
	"BetterPC_2.0/internal/handlers/helpers/responseManager"
	"BetterPC_2.0/pkg/data/models/orders"
	orderErrors "BetterPC_2.0/pkg/data/models/orders/errors"
	orderFilters "BetterPC_2.0/pkg/data/models/orders/filters"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

type VerificationCheckResult struct {
	IsVerified bool
	Err        error
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

	userId := c.Param("user_id")
	if userId == "" {
		responseManager.ErrorResponseWithLog(c, http.StatusBadRequest, "empty order id")
		return
	}

	err := h.services.CancelOrder(userId, orderId)
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
