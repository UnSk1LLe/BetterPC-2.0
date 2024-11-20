package handlers

import (
	"BetterPC_2.0/internal/handlers/helpers/responseManager"
	userFilters "BetterPC_2.0/pkg/data/models/users/filters"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) ListUsers(c *gin.Context) {
	var filters userFilters.AdminUserFilters

	if err := c.ShouldBindJSON(&filters); err != nil {
		responseManager.ErrorResponse(c, http.StatusBadRequest, "invalid filter parameters")
		return
	}

}
