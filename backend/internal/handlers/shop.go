package handlers

import (
	"BetterPC_2.0/internal/handlers/helpers/responseManager"
	"BetterPC_2.0/pkg/data/models/products"
	productErrors "BetterPC_2.0/pkg/data/models/products/errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func (h *Handler) Shop(c *gin.Context) {

	c.Redirect(http.StatusOK, "shop/categories")

	return
}

func (h *Handler) ListStandardizedProducts(c *gin.Context) {

	productType, err := products.ProductTypeFromString(c.Param("product_type"))
	if err != nil {
		responseManager.ErrorResponseWithLog(c, http.StatusBadRequest, err.Error())
		return
	}

	searchQuery := c.Query("search")

	//TODO fetch filters
	standardizedProductsList, err := h.services.Product.GetStandardizedList(searchQuery, nil, productType)
	if err != nil {
		//errors.RenderError(c, http.StatusInternalServerError, "shop/categories", "get", err)
		responseManager.ErrorResponseWithLog(c, http.StatusInternalServerError, err.Error())
		return
	}

	/*html.Render(c, http.StatusOK, "templates/pages/listProducts", gin.H{
		"title":        "List Products",
		"ProductsList": standardizedProductsList,
	})*/

	c.JSON(http.StatusOK, gin.H{
		"productsList": standardizedProductsList,
	})

	return
}

func (h *Handler) ShowProductInfo(c *gin.Context) {
	productId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		logMessage := fmt.Sprintf("error getting product id: %s", err.Error())
		responseManager.ErrorResponseWithLog(c, http.StatusInternalServerError, logMessage)
		//errors.RenderError(c, http.StatusInternalServerError, "shop/categories", "get", err)
		return
	}

	productType, err := products.ProductTypeFromString(c.Param("product_type"))
	if err != nil {
		responseManager.ErrorResponseWithLog(c, http.StatusBadRequest, err.Error())
		return
	}

	product, err := h.services.Product.GetById(productId, productType)

	if err != nil {
		switch {
		case errors.Is(err, productErrors.ErrNoProductsFound):
			responseManager.ErrorResponse(c, http.StatusNotFound, err.Error())
			return
		}
		message := fmt.Sprintf("error getting product by id: %s", err.Error())
		h.logger.Error(message)
		responseManager.ErrorResponse(c, http.StatusInternalServerError, message)
		//errors.RenderError(c, http.StatusInternalServerError, "shop/categories", "get", err)
		return
	}

	/*html.Render(c, http.StatusOK, "templates/pages/productInfo", gin.H{
		"title":       fmt.Sprintf("%s info", productType),
		"Product":     product,
		"ProductType": productType,
		"Build":       nil,
	})*/

	c.JSON(http.StatusOK, gin.H{
		"product":     product,
		"productType": productType,
		"build":       nil,
	})

}
