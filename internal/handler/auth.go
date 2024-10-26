package handler

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
	var user users.User
	err := c.BindJSON(&user)

	if err != nil {
		message := "Your credentials are not valid"
		errors.RenderError(c, http.StatusBadRequest, "/registerForm", "get", err, message)
		return
	}

	ObjId, err := h.services.Authorization.CreateUser(user)
	if err != nil {
		message := fmt.Sprintf("Error creating user: %s", err.Error())
		h.logger.Error(message)
		errors.RenderError(c, http.StatusInternalServerError, "/registerForm", "get", err, message)
		return
	}

	h.logger.Infof("User with ID <%s> was CREATED!", ObjId.Hex())
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

	err := c.BindJSON(&input)
	if err != nil {
		message := "Your credentials are not valid"
		errors.RenderError(c, http.StatusBadRequest, "/auth/login", "get", err, message)
		return
	}

	tokens, err := h.services.Authorization.GenerateTokenPair(input.Email, input.Password)
	if err != nil {
		message := err.Error()
		errors.RenderError(c, http.StatusInternalServerError, "/auth/login", "get", err, message)
		return
	}

	//Setting tokens to cookies with httpOnly flag
	c.SetCookie("access_token", tokens.AccessToken, 3600, "/", "", true, true)

	c.SetCookie("refresh_token", tokens.RefreshToken, 7*24*3600, "/", "", true, true)

	c.JSON(http.StatusOK, gin.H{"message": "Login successful!"})
	return
}
