package handlers

import (
	"BetterPC_2.0/internal/handlers/helpers/responseManager"
	"BetterPC_2.0/internal/service"
	userErrors "BetterPC_2.0/pkg/data/models/users/errors"
	"BetterPC_2.0/pkg/email/verification"
	emailVerificationErrors "BetterPC_2.0/pkg/email/verification/errors"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type EmailInput struct {
	Email string `json:"email" binding:"required"`
}

func (h *Handler) SendRecoveryLink(c *gin.Context) {
	var input EmailInput

	//TODO sendMessage despite the email was correct or not("If account with the specified email exists, we will send you a recovery link.")

	deferResponseStatus := http.StatusOK
	deferMessage := "If your email is correct, we will send a recovery link to it."

	//initially response supposed to be a messageResponse, but if critical error occurs it will be changed to errorResponse
	var responseFunc func(c *gin.Context, status int, message ...string) = responseManager.MessageResponse
	defer func() {
		responseFunc(c, deferResponseStatus, deferMessage)
	}() //using anonymous function so that defer would consider the changed parameters

	err := c.ShouldBindJSON(&input)
	if err != nil {
		deferMessage = err.Error()
		deferResponseStatus = http.StatusBadRequest
		responseFunc = responseManager.ErrorResponse
		return
	}

	isReal, err := verification.IsRealEmail(input.Email)
	if !isReal {
		message := fmt.Sprintf("%s is not a real email: %s", input.Email, err.Error())
		h.logger.Error(message)
		deferMessage = message
		deferResponseStatus = http.StatusNotAcceptable

		if err != nil && !errors.As(err, &emailVerificationErrors.EmailVerificationError{}) {
			deferMessage = err.Error()
			deferResponseStatus = http.StatusInternalServerError
		}

		responseFunc = responseManager.ErrorResponse
		return
	}

	isVerified, err := h.services.Verification.IsVerifiedUser(input.Email)
	if err != nil {

		if !errors.Is(err, userErrors.ErrUserNotFound) {
			deferMessage = err.Error()
			deferResponseStatus = http.StatusInternalServerError
			responseFunc = responseManager.ErrorResponse
			return
		}

		message := err.Error()
		h.logger.Error(message)
		return
	}
	if !isVerified {
		message := fmt.Sprintf("email %s is not verified", input.Email)
		h.logger.Error(message)
		return
	}

	token, err := h.services.Verification.SetNewToken(input.Email, service.DefaultVerificationTokenTTL)
	if err != nil {
		deferMessage = err.Error()
		deferResponseStatus = http.StatusInternalServerError
		responseFunc = responseManager.ErrorResponse
		return
	}

	serverUrl := h.cfg.Server.Url
	endpointPath := c.Request.URL.Path
	recoveryUrl := fmt.Sprintf("%s%s/%s", serverUrl, endpointPath, token)

	emailSubject := "Password recovery"
	emailBody := fmt.Sprintf(`Use the following link to create a new password for your account:
					%s`, recoveryUrl)

	err = h.services.Notification.SendEmailToUser(input.Email, emailSubject, emailBody)
	if err != nil {
		message := err.Error()
		responseManager.ErrorResponseWithLog(c, http.StatusInternalServerError, message)
		return
	}

	h.logger.Info("recovery link successfully created and send")
}

type PasswordRecoveryInput struct {
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

func (h *Handler) RecoverPassword(c *gin.Context) {
	var input PasswordRecoveryInput

	token := c.Param("token")
	if len(token) == service.VerificationTokenByteLength {
		responseManager.ErrorResponseWithLog(c, http.StatusBadRequest, "Token required")
		return
	}

	err := c.ShouldBindJSON(&input)
	if err != nil {
		responseManager.ErrorResponseWithLog(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Verification.UpdatePasswordByToken(token, input.Password)
	if err != nil {
		var status int
		message := err.Error()

		switch {
		case errors.Is(err, userErrors.ErrUserNotFound):
			status = http.StatusNotFound
			message = "token is invalid"
		case errors.Is(err, userErrors.ErrUserNotVerified):
			status = http.StatusNotModified
		default:
			status = http.StatusInternalServerError
		}
		responseManager.ErrorResponseWithLog(c, status, message)
		return
	}
	//TODO go func optimization for recovery and sending link

	responseManager.MessageResponse(c, http.StatusOK, "Password recovered successfully")
}
