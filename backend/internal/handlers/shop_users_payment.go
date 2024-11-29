package handlers

import (
	"BetterPC_2.0/internal/handlers/helpers/responseManager"
	"BetterPC_2.0/internal/middlewares/helpers/userContext"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetUserPaymentMethods(c *gin.Context) {

	user, err := userContext.GetUserCtx(c)
	if err != nil {
		responseManager.ErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	paymentMethods, err := h.services.ListPaymentMethodsByUser(user.ID)
	if err != nil {
		responseManager.ErrorResponseWithLog(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, paymentMethods)
	return
}

func (h *Handler) AttachNewPaymentMethod(c *gin.Context) {
	paymentMethod := c.Param("id")
	if paymentMethod == "" {
		responseManager.ErrorResponse(c, http.StatusBadRequest, "empty payment id")
		return
	}

	user, err := userContext.GetUserCtx(c)
	if err != nil {
		responseManager.ErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	err = h.services.AttachPaymentMethodToUser(user.ID, paymentMethod)
	if err != nil {
		responseManager.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	message := "Payment method attached successfully!"
	responseManager.MessageResponseWithLog(c, http.StatusOK, message)
	return
}

func (h *Handler) RemovePaymentMethod(c *gin.Context) {
	paymentMethodId := c.Param("id")
	if paymentMethodId == "" {
		responseManager.ErrorResponseWithLog(c, http.StatusBadRequest, "empty payment method id")
		return
	}

	user, err := userContext.GetUserCtx(c)
	if err != nil {
		responseManager.ErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	err = h.services.RemovePaymentMethod(user.ID, paymentMethodId)
	if err != nil {
		responseManager.ErrorResponseWithLog(c, http.StatusInternalServerError, err.Error())
		return
	}

	message := "Payment method removed successfully!"
	responseManager.MessageResponse(c, http.StatusOK, message)
	return
}
