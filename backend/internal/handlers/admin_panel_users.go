package handlers

import (
	"BetterPC_2.0/internal/handlers/helpers/responseManager"
	userFilters "BetterPC_2.0/pkg/data/models/users/filters"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetUserList(c *gin.Context) {
	var filters userFilters.AdminUserFilters

	if err := c.ShouldBindJSON(&filters); err != nil {
		responseManager.ErrorResponse(c, http.StatusBadRequest, "invalid filter parameters")
		return
	}

}

func (h *Handler) GetUser(c *gin.Context) {
	userId := c.Param("id")
	if userId == "" {
		responseManager.ErrorResponse(c, http.StatusBadRequest, "empty user id")
		return
	}

}

func (h *Handler) CreateUser(c *gin.Context) {

}

func (h *Handler) UpdateUser(c *gin.Context) {
	userId := c.Param("id")
	if userId == "" {
		responseManager.ErrorResponse(c, http.StatusBadRequest, "empty user id")
		return
	}

}

func (h *Handler) DeleteUser(c *gin.Context) {
	userId := c.Param("id")
	if userId == "" {
		responseManager.ErrorResponse(c, http.StatusBadRequest, "empty user id")
		return
	}

}
