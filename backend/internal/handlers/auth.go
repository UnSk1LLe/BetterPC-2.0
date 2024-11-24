package handlers

import (
	"BetterPC_2.0/internal/handlers/helpers/responseManager"
	userErrors "BetterPC_2.0/pkg/data/models/users/errors"
	userRequests "BetterPC_2.0/pkg/data/models/users/requests/auth"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

/*func (h *Handler) RegisterForm(c *gin.Context) {

	html.Render(c, http.StatusOK, "templates/pages/register", gin.H{
		"title": "Registration form",
	})

}*/

func (h *Handler) Register(c *gin.Context) {
	var input userRequests.RegisterRequest

	err := c.ShouldBindJSON(&input)
	if err != nil {
		message := "Your credentials are not valid"
		responseManager.ErrorResponseWithLog(c, http.StatusBadRequest, message)
		//errors.RenderError(c, http.StatusBadRequest, "/auth/register", "get", err, message)
		return
	}

	ObjId, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		message := fmt.Sprintf("Error creating user: %s", err.Error())
		h.logger.Error(message)

		if errors.Is(err, userErrors.ErrUserAlreadyExists) {
			responseManager.ErrorResponseWithLog(c, http.StatusConflict, message)
			return
		}

		responseManager.ErrorResponseWithLog(c, http.StatusInternalServerError, message)
		//errors.RenderError(c, http.StatusInternalServerError, "/auth/register", "get", err, message)
		return
	}

	h.logger.Infof("User with ID <%s> was REGISTERED!", ObjId.Hex())
	responseManager.MessageResponseWithLog(c, http.StatusCreated, "User created successfully")
}

/*func (h *Handler) LoginForm(c *gin.Context) {

	html.Render(c, http.StatusOK, "templates/pages/login", gin.H{
		"title": "Login form",
	})

}*/

func (h *Handler) Login(c *gin.Context) {
	var input userRequests.LoginRequest

	err := c.ShouldBindJSON(&input)
	if err != nil {
		message := "Your credentials are not valid"
		responseManager.ErrorResponseWithLog(c, http.StatusBadRequest, err.Error(), message)
		//errors.RenderError(c, http.StatusBadRequest, "/auth/login", "get", err, message)
		return
	}

	tokens, err := h.services.Authorization.GenerateTokenPair(input.Email, input.Password)
	if err != nil {

		if errors.Is(err, userErrors.ErrUserNotFound) {
			message := fmt.Sprintf("Error generating token pair: %s", err.Error())
			responseManager.ErrorResponseWithLog(c, http.StatusConflict, message)
			return
		}

		message := "Could not authorize user"
		responseManager.ErrorResponseWithLog(c, http.StatusInternalServerError, message+err.Error())
		//errors.RenderError(c, http.StatusInternalServerError, "/auth/login", "get", err, message)
		return
	}

	//Setting tokens to cookies with httpOnly flag
	c.SetCookie("accessToken", tokens.AccessToken, 15*60, "/", "", true, true)
	c.SetCookie("refreshToken", tokens.RefreshToken, 7*24*60*60, "/", "", true, true)

	message := "User has logged in successfully"
	logMessage := fmt.Sprintf("user with email %s has logged in!", input.Email)

	responseManager.MessageResponseWithLog(c, http.StatusOK, logMessage, message)
	return
}

func (h *Handler) Logout(c *gin.Context) {
	c.SetCookie("accessToken", "", -1, "/", "", true, true)
	c.SetCookie("refreshToken", "", -1, "/", "", true, true)

	logMessage := "access & refresh token cookies were removed due to user log out"
	message := "User has logged out"

	responseManager.MessageResponseWithLog(c, http.StatusOK, logMessage, message)
	return
}

func (h *Handler) Refresh(c *gin.Context) {

	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		message := fmt.Sprintf("Error getting refresh token: %s", err.Error())
		responseManager.ErrorResponseWithLog(c, http.StatusBadRequest, message)
		//errorHandler.RenderError(c, http.StatusBadRequest, "/auth/login", "get", err, message)
		return
	}

	_, tokens, err := h.services.Authorization.RefreshTokens(refreshToken)
	if err != nil {
		message := err.Error()
		responseManager.ErrorResponseWithLog(c, http.StatusInternalServerError, message)
		//errorHandler.RenderError(c, http.StatusInternalServerError, "/auth/login", "get", err, message)
		return
	}

	//Setting tokens to cookies with httpOnly flag
	c.SetCookie("accessToken", tokens.AccessToken, 3600, "/", "", true, true)
	c.SetCookie("refreshToken", tokens.RefreshToken, 7*24*3600, "/", "", true, true)

	c.Redirect(http.StatusFound, "/shop/categories")
	return
}
