package handlers

import (
	"BetterPC_2.0/internal/handlers/helpers/responseManager"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) ListCategories(c *gin.Context) {

	searchQuery := c.Query("search")

	productTypeCount, err := h.services.CountProductsForEachCategory(searchQuery)
	if err != nil {
		responseManager.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, productTypeCount)
}
