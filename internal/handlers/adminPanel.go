package handlers

import (
	"BetterPC_2.0/pkg/data/models/products"
	"BetterPC_2.0/pkg/errors"
	"BetterPC_2.0/pkg/messages"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func (h *Handler) CreateProduct(c *gin.Context) {
	var product products.Product
	productType := c.Param("productType")

	if err := c.BindJSON(&product); err != nil {
		message := "Could not parse product model"
		errors.RenderError(c, http.StatusInternalServerError, "adminPanel/categories", "get", err, message)
		return
	}

	productId, err := h.services.Product.Create(product, productType)
	if err != nil {
		message := "Server error while creating the product"
		errors.RenderError(c, http.StatusInternalServerError, "adminPanel/categories", "get", err, message)
		return
	}

	message := fmt.Sprintf("Product successfully CREATED. ID = %s!", productId.Hex())
	messages.RenderMessage(c, http.StatusOK, "/adminPanel/categories", "get", message)
}

func (h *Handler) DeleteProduct(c *gin.Context) {
	productType := c.Param("category_name")
	productId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		message := "invalid product id"
		errors.RenderError(c, http.StatusInternalServerError, "adminPanel/categories", "get", err, message)
		return
	}

	err = h.services.Product.DeleteById(productId, productType)
	if err != nil {
		message := "Server error while deleting the product"
		errors.RenderError(c, http.StatusInternalServerError, "adminPanel/categories", "get", err, message)
		return
	}

	message := fmt.Sprintf("Product with ID = %s DELETED successfully!", productId.Hex())
	messages.RenderMessage(c, http.StatusOK, "/adminPanel/categories", "get", message)
}

func (h *Handler) GetCategoriesAdmin(c *gin.Context) {

	/*categoriesList, err := h.services.Categories.GetList(bson.M{})
	if err != nil {
		message := "Error getting categories"
		errors.RenderError(c, http.StatusInternalServerError, "adminPanel/categories", "get", err, message)
		return
	}*/

	message := fmt.Sprintf("ok")
	messages.RenderMessage(c, http.StatusOK, "/adminPanel/categories", "get", message)
}
