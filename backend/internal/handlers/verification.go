package handlers

import (
	"BetterPC_2.0/internal/handlers/helpers/responseManager"
	"BetterPC_2.0/internal/service"
	userErrors "BetterPC_2.0/pkg/data/models/users/errors"
	"BetterPC_2.0/pkg/email/verification"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

func (h *Handler) SendVerificationLink(c *gin.Context) {
	var input EmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		responseManager.ErrorResponseWithLog(c, http.StatusBadRequest, err.Error())
		return
	}

	ok, err := verification.IsRealEmail(input.Email)
	if err != nil {
		logMessage := fmt.Sprintf("invalid email address %s: %s", input.Email, err.Error())
		responseManager.ErrorResponseWithLog(c, http.StatusNotAcceptable, logMessage)
		return
	}
	if !ok {
		logMessage := fmt.Sprintf("invalid email address %s", input.Email)
		responseManager.ErrorResponseWithLog(c, http.StatusNotAcceptable, logMessage)
		return
	}

	isVerified, err := h.services.Verification.IsVerifiedUser(input.Email)
	if err != nil {
		responseManager.ErrorResponseWithLog(c, http.StatusInternalServerError, err.Error())
		return
	}
	if isVerified {
		message := "user is already verified"
		responseManager.ErrorResponseWithLog(c, http.StatusConflict, message)
		return
	}

	token, err := h.services.Verification.SetNewToken(input.Email, service.DefaultVerificationTokenTTL)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, userErrors.ErrUserNotFound) {
			status = http.StatusNotFound
		}
		responseManager.ErrorResponseWithLog(c, status, err.Error())
		return
	}

	serverUrl := h.cfg.Server.Url
	endpointPath := c.Request.URL.Path
	verificationUrl := fmt.Sprintf("%s%s%s", serverUrl, endpointPath, token)

	emailSubject := "Email verification"
	emailBody := fmt.Sprintf(`Please click the following link to verify your email address: 
								%s`, verificationUrl)

	err = h.services.Notification.SendEmailToUser(input.Email, emailSubject, emailBody)
	if err != nil {
		responseManager.ErrorResponseWithLog(c, http.StatusInternalServerError, err.Error())
	}

	message := "Verification link has been sent to your email"
	logMessage := fmt.Sprintf("Verification link has been sent to email: %s", input.Email)
	responseManager.MessageResponseWithLog(c, http.StatusOK, logMessage, message)
}

func (h *Handler) SendNewVerificationLink(c *gin.Context) {

}

func (h *Handler) VerifyUser(c *gin.Context) {
	token := c.Param("token")
	if token == "" {
		responseManager.ErrorResponseWithLog(c, http.StatusBadRequest, "empty token")
		return
	}

	err := h.services.Verification.VerifyUser(token) //TODO handle error properly
	if err != nil {
		message := fmt.Sprintf("verification error: %s", err.Error())
		responseManager.ErrorResponseWithLog(c, http.StatusNotFound, message)
		return
	}

	message := fmt.Sprintf("user successfully verified")
	logMessage := fmt.Sprintf("user %s successfully verified", token)
	responseManager.ErrorResponseWithLog(c, http.StatusOK, logMessage, message)
}
