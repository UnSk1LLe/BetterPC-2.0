package handlers

import (
	"BetterPC_2.0/pkg/html"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) ListCategories(c *gin.Context) {

	html.Render(c, http.StatusOK, "templates/pages/index", gin.H{
		"title": "Categories",
	})

}
