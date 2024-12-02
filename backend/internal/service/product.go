package service

import (
	"BetterPC_2.0/internal/repository"
	"BetterPC_2.0/internal/service/helpers/searchEngine"
	"BetterPC_2.0/pkg/data/models/products"
	generalRequests "BetterPC_2.0/pkg/data/models/products/general/requests"
	generalResponses "BetterPC_2.0/pkg/data/models/products/general/responses"
	productRequests "BetterPC_2.0/pkg/data/models/products/requests"
	"BetterPC_2.0/pkg/logging"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mime/multipart"
)

type ProductService struct {
	repo        repository.Product
	fileService *FileService
	logger      *logging.Logger
}

func NewProductService(repo repository.Product, fileService *FileService, logger *logging.Logger) *ProductService {
	return &ProductService{
		repo:        repo,
		fileService: fileService,
		logger:      logger,
	}
}

func (p *ProductService) Create(product products.Product, productType products.ProductType, image *multipart.FileHeader) (primitive.ObjectID, error) {
	imageName, err := p.fileService.AddProductImage(image)
	if err != nil {
		p.logger.Error(err.Error())
		return primitive.NilObjectID, err
	}

	product.SetImage(imageName)

	return p.repo.Create(product, productType)
}

func (p *ProductService) GetById(id primitive.ObjectID, productType products.ProductType) (products.Product, error) {
	return p.repo.GetById(id, productType)
}

func (p *ProductService) GetStandardizedList(searchQuery string, propertyFilters map[string]interface{}, productType products.ProductType) ([]generalResponses.StandardizedProductData, error) {
	_, filter := searchEngine.GetSearchFilter(searchQuery)

	productsList, err := p.repo.GetList(filter, productType)
	if err != nil {
		return nil, err
	}

	standardizedProductsList := make([]generalResponses.StandardizedProductData, len(productsList))
	for i, product := range productsList {
		standardizedProductsList[i] = product.Standardize()
	}

	return standardizedProductsList, nil
}

func (p *ProductService) GetList(searchQuery string, propertyFilter map[string]interface{}, productType products.ProductType) ([]products.Product, error) {
	_, searchFilter := searchEngine.GetSearchFilter(searchQuery)
	return p.repo.GetList(searchFilter, productType)
}

func (p *ProductService) CountProductsForEachCategory(searchQuery string) (map[products.ProductType]int, error) {
	productTypeCount := make(map[products.ProductType]int)

	productTypes, filter := searchEngine.GetSearchFilter(searchQuery)
	if len(productTypes) != 0 {
		for _, productType := range productTypes {
			count, err := p.repo.CountCategoryProducts(filter, productType)
			if err != nil {
				return productTypeCount, err
			}
			productTypeCount[productType] = count
		}
	}
	return productTypeCount, nil
}

func (p *ProductService) DeleteById(id primitive.ObjectID, productType products.ProductType) error {
	product, err := p.repo.DeleteById(id, productType)
	if err != nil {
		p.logger.Error(err.Error())
		return err
	}

	if err := p.fileService.DeleteProductImage(product.GetImage()); err != nil {
		p.logger.Error(err.Error())
	}

	return nil
}

func (p *ProductService) UpdateById(id primitive.ObjectID,
	input productRequests.ProductUpdateRequest,
	productType products.ProductType,
	image *multipart.FileHeader,
) error {
	err := input.Validate()
	if err != nil {
		return err
	}

	if image != nil {
		imageName, err := p.fileService.AddProductImage(image)
		if err != nil {
			p.logger.Error(err.Error())
			return err
		}

		input.SetImage(&imageName)
	}

	return p.repo.UpdateById(id, input, productType)
}

func (p *ProductService) UpdateGeneralInfoById(
	id primitive.ObjectID,
	input generalRequests.UpdateGeneralRequest,
	productType products.ProductType,
	image *multipart.FileHeader,
) error {
	var imageName string

	err := input.Validate()
	if err != nil {
		return err
	}

	if image != nil {
		imageName, err = p.fileService.AddProductImage(image)
		if err != nil {
			p.logger.Error(err.Error())
			return err
		}

		input.Image = &imageName
	}

	if err := p.repo.UpdateGeneralInfoById(id, input, productType); err != nil {
		return err
	}

	if imageName != "" {
		if err := p.fileService.DeleteProductImage(imageName); err != nil {
			p.logger.Error(err.Error())
		}
	}

	return nil
}
