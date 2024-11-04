package handlers

import (
	"BetterPC_2.0/pkg/data/models/users"
	"BetterPC_2.0/pkg/errors"
	"BetterPC_2.0/pkg/html"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) RegisterForm(c *gin.Context) {

	html.Render(c, http.StatusOK, "templates/pages/register", gin.H{
		"title": "Registration form",
	})

}

func (h *Handler) Register(c *gin.Context) {
	var input users.RegisterInput
	err := c.ShouldBind(&input)

	if err != nil {
		message := "Your credentials are not valid"
		errors.RenderError(c, http.StatusBadRequest, "/auth/register", "get", err, message)
		return
	}

	ObjId, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		message := fmt.Sprintf("Error creating user: %s", err.Error())
		h.logger.Error(message)
		errors.RenderError(c, http.StatusInternalServerError, "/auth/register", "get", err, message)
		return
	}

	h.logger.Infof("User with ID <%s> was REGISTERED!", ObjId.Hex())
	c.Redirect(http.StatusCreated, "/auth/login")
}

func (h *Handler) LoginForm(c *gin.Context) {

	html.Render(c, http.StatusOK, "templates/pages/login", gin.H{
		"title": "Login form",
	})

}

type loginInput struct {
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func (h *Handler) Login(c *gin.Context) {
	var input loginInput

	err := c.ShouldBind(&input)
	if err != nil {
		message := "Your credentials are not valid"
		errors.RenderError(c, http.StatusBadRequest, "/auth/login", "get", err, message)
		return
	}

	tokens, err := h.services.Authorization.GenerateTokenPair(input.Email, input.Password)
	if err != nil {
		h.logger.Error(err)
		message := "Could not authorize user"
		errors.RenderError(c, http.StatusInternalServerError, "/auth/login", "get", err, message)
		return
	}

	//Setting tokens to cookies with httpOnly flag
	c.SetCookie("accessToken", tokens.AccessToken, 3600, "/", "", true, true)

	c.SetCookie("refreshToken", tokens.RefreshToken, 7*24*3600, "/", "", true, true)

	c.Redirect(http.StatusFound, "/shop/categories")
	return
}

func (h *Handler) Refresh(c *gin.Context) {

	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		message := err.Error()
		errors.RenderError(c, http.StatusBadRequest, "/auth/login", "get", err, message)
		return
	}

	tokens, _, err := h.services.Authorization.RefreshTokens(refreshToken)
	if err != nil {
		message := err.Error()
		errors.RenderError(c, http.StatusInternalServerError, "/auth/login", "get", err, message)
		return
	}

	//Setting tokens to cookies with httpOnly flag
	c.SetCookie("accessToken", tokens.AccessToken, 3600, "/", "", true, true)

	c.SetCookie("refreshToken", tokens.RefreshToken, 7*24*3600, "/", "", true, true)

	c.Redirect(http.StatusFound, "/shop/categories")
	return
}
