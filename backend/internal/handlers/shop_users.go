package handlers

import (
	"BetterPC_2.0/internal/handlers/helpers/responseManager"
	"BetterPC_2.0/internal/middlewares/helpers/userContext"
	orderErrors "BetterPC_2.0/pkg/data/models/orders/errors"
	userErrors "BetterPC_2.0/pkg/data/models/users/errors"
	userUpdateRequests "BetterPC_2.0/pkg/data/models/users/requests/patch"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

func (h *Handler) UploadUserImage(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		responseManager.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := userContext.GetUserCtx(c)
	if err != nil {
		responseManager.ErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	// Save the file
	err = h.services.User.UpdateUserImage(user.ID, file)
	if err != nil {
		responseManager.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	responseManager.MessageResponse(c, http.StatusOK, "Image updated successfully!")
}

func (h *Handler) UpdateUserInfo(c *gin.Context) {
	user, err := userContext.GetUserCtx(c)
	if err != nil {
		responseManager.ErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	var input userUpdateRequests.UpdateUserInfoRequest

	err = c.ShouldBindJSON(&input)
	if err != nil {
		responseManager.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.User.UpdateUserInfo(user.ID, input)
	if err != nil {
		switch {
		case errors.As(err, &userErrors.UserError{}):
			responseManager.ErrorResponse(c, http.StatusConflict, err.Error())
			return
		}
		responseManager.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	responseManager.MessageResponse(c, http.StatusOK, "User updated successfully!")
	return
}

func (h *Handler) GetUserInfo(c *gin.Context) {
	user, err := userContext.GetUserCtx(c)
	if err != nil {
		responseManager.ErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
	return
}

func (h *Handler) CancelUserOrder(c *gin.Context) {
	orderId := c.Param("id")
	if orderId == "" {
		responseManager.ErrorResponseWithLog(c, http.StatusBadRequest, "empty order id")
		return
	}

	user, err := userContext.GetUserCtx(c)
	if err != nil {
		responseManager.ErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	err = h.services.CancelOrder(user.ID, orderId)
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
