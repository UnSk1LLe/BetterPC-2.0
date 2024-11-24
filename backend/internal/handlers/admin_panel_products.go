package handlers

import (
	"BetterPC_2.0/internal/handlers/helpers/responseManager"
	validatorErrors "BetterPC_2.0/pkg/data/helpers/validators/errors"
	"BetterPC_2.0/pkg/data/models/products"
	productErrors "BetterPC_2.0/pkg/data/models/products/errors"
	generalRequests "BetterPC_2.0/pkg/data/models/products/general/requests"
	productRequests "BetterPC_2.0/pkg/data/models/products/requests"
	"BetterPC_2.0/pkg/html"
	"BetterPC_2.0/pkg/messages"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mime/multipart"
	"net/http"
)

func (h *Handler) CreateProductPage(c *gin.Context) {

	html.Render(c, http.StatusOK, "templates/pages/addProduct", gin.H{})
}

func (h *Handler) CreateProduct(c *gin.Context) {
	productType, err := products.ProductTypeFromString(c.Param("product_type"))
	if err != nil {
		responseManager.ErrorResponseWithLog(c, http.StatusBadRequest, err.Error())
		return
	}

	productFactory, ok := products.ProductTypeFactory[productType]
	if !ok {
		message := fmt.Sprintf("unsupported product type %s", productType)
		responseManager.ErrorResponseWithLog(c, http.StatusBadRequest, message)
		return
	}

	product := productFactory()

	/*if err := c.BindJSON(&product); err != nil {
		message := "Could not parse product model"
		responseManager.ErrorResponseWithLog(c, http.StatusInternalServerError, message)
		//errors.RenderError(c, http.StatusInternalServerError, "adminPanel/categories", "get", err, message)
		return
	}*/

	productData := c.PostForm("product_data")
	if productData == "" {
		err := json.Unmarshal([]byte(productData), &product)
		if err != nil {
			responseManager.ErrorResponseWithLog(c, http.StatusBadRequest, err.Error())
			return
		}
	}

	logrus.Info(product)

	var fileHeader *multipart.FileHeader
	fileHeader, err = c.FormFile("image")
	if err != nil {
		if !errors.Is(err, http.ErrMissingFile) {
			responseManager.ErrorResponseWithLog(c, http.StatusBadRequest, err.Error())
			return
		}
	}

	productId, err := h.services.Product.Create(product, productType, fileHeader)
	if err != nil {

		if errors.Is(err, productErrors.ErrProductModelAlreadyExists) {
			message := fmt.Sprintf(
				"could not create %s with the model %s: %s",
				productType, product.GetModel(), err.Error(),
			)

			responseManager.ErrorResponseWithLog(c, http.StatusConflict, message)
			return
		}

		message := "Server error while creating the product"

		responseManager.ErrorResponseWithLog(c, http.StatusInternalServerError, message+": "+err.Error(), message)
		//errors.RenderError(c, http.StatusInternalServerError, "adminPanel/categories", "get", err, message)
		return
	}

	message := fmt.Sprintf("Product successfully CREATED with ID %s!", productId.Hex())
	responseManager.MessageResponseWithLog(c, http.StatusCreated, message)
	//message := fmt.Sprintf("Product successfully CREATED. ID = %s!", productId.Hex())
	//messages.RenderMessage(c, http.StatusOK, "/adminPanel/categories", "get", message)
}

func (h *Handler) DeleteProduct(c *gin.Context) {
	productType, err := products.ProductTypeFromString(c.Param("product_type"))
	if err != nil {
		responseManager.ErrorResponseWithLog(c, http.StatusBadRequest, err.Error())
		return
	}

	productId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		message := "invalid product id"
		responseManager.ErrorResponseWithLog(c, http.StatusBadRequest, message)
		//errors.RenderError(c, http.StatusInternalServerError, "adminPanel/categories", "get", err, message)
		return
	}

	err = h.services.Product.DeleteById(productId, productType)
	if err != nil {
		message := "Server error while deleting the product: " + err.Error()
		status := http.StatusInternalServerError
		if errors.Is(err, productErrors.ErrNoProductsFound) {
			status = http.StatusNotFound
		}

		responseManager.ErrorResponseWithLog(c, status, message)
		//errors.RenderError(c, http.StatusInternalServerError, "adminPanel/categories", "get", err, message)
		return
	}

	message := fmt.Sprintf("Product with ID %s DELETED successfully!", productId.Hex())
	responseManager.MessageResponseWithLog(c, http.StatusOK, message)
	//message := fmt.Sprintf("Product with ID = %s DELETED successfully!", productId.Hex())
	//messages.RenderMessage(c, http.StatusOK, "/adminPanel/categories", "get", message)
}

func (h *Handler) UpdateProduct(c *gin.Context) {
	productType, err := products.ProductTypeFromString(c.Param("product_type"))
	if err != nil {
		responseManager.ErrorResponseWithLog(c, http.StatusBadRequest, err.Error())
		return
	}

	productId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		message := fmt.Sprintf("invalid product id: %s", err.Error())
		responseManager.ErrorResponseWithLog(c, http.StatusBadRequest, message)
		return
	}

	factory, ok := productRequests.ProductUpdateRequestFactory[productType]
	if !ok {
		message := fmt.Sprintf("%s: %s", productErrors.ErrUnsupportedProductType.Error(), productType)
		responseManager.ErrorResponseWithLog(c, http.StatusBadRequest, message)
		return
	}
	input := factory()

	/*err = c.ShouldBindJSON(&input)
	if err != nil {
		message := fmt.Sprintf("could not bind product: %s", err.Error())
		responseManager.ErrorResponseWithLog(c, http.StatusBadRequest, message)
		return
	}*/

	productData := c.PostForm("product_data")
	if productData == "" {
		err := json.Unmarshal([]byte(productData), &input)
		if err != nil {
			responseManager.ErrorResponseWithLog(c, http.StatusBadRequest, err.Error())
			return
		}
	}

	var fileHeader *multipart.FileHeader
	fileHeader, err = c.FormFile("image")
	if err != nil {
		if !errors.Is(err, http.ErrMissingFile) {
			responseManager.ErrorResponseWithLog(c, http.StatusBadRequest, err.Error())
			return
		}
	}

	err = h.services.Product.UpdateById(productId, input, productType, fileHeader)
	if err != nil {
		message := fmt.Sprintf("server error while updating the product: %s", err.Error())
		status := http.StatusInternalServerError

		if errors.As(err, validatorErrors.ValidatorError{}) {
			status = http.StatusBadRequest
		}

		responseManager.ErrorResponseWithLog(c, status, message)
		return
	}

	message := "Successfully updated the product"
	logMessage := fmt.Sprintf("Product with ID %s UPDATED successfully!", productId.Hex())
	responseManager.MessageResponseWithLog(c, http.StatusOK, logMessage, message)
	//message := fmt.Sprintf("Product with ID = %s DELETED successfully!", productId.Hex())
	//messages.RenderMessage(c, http.StatusOK, "/adminPanel/categories", "get", message)
}

func (h *Handler) UpdateProductGeneral(c *gin.Context) {
	productType, err := products.ProductTypeFromString(c.Param("product_type"))
	if err != nil {
		responseManager.ErrorResponseWithLog(c, http.StatusBadRequest, err.Error())
		return
	}

	productId, err := primitive.ObjectIDFromHex(c.Param("id")) //TODO move the converting logic to the service layer
	if err != nil {
		message := "invalid product id"
		responseManager.ErrorResponseWithLog(c, http.StatusBadRequest, message)
		return
	}

	var input generalRequests.UpdateGeneralRequest

	productData := c.PostForm("product_data")
	if productData == "" {
		err := json.Unmarshal([]byte(productData), &input)
		if err != nil {
			responseManager.ErrorResponseWithLog(c, http.StatusBadRequest, err.Error())
			return
		}
	}

	logrus.Info(input)

	var fileHeader *multipart.FileHeader
	fileHeader, err = c.FormFile("image")
	if err != nil {
		if !errors.Is(err, http.ErrMissingFile) {
			responseManager.ErrorResponseWithLog(c, http.StatusBadRequest, err.Error())
			return
		}
	}

	err = h.services.Product.UpdateGeneralInfoById(productId, input, productType, fileHeader)
	if err != nil {
		message := fmt.Sprintf("server error while updating the product: %s", err.Error())
		responseManager.ErrorResponse(c, http.StatusInternalServerError, message)
	}
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
