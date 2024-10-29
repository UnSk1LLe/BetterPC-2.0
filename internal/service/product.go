package service

import (
	"BetterPC_2.0/internal/repository"
	"BetterPC_2.0/pkg/data/models/products"
	"BetterPC_2.0/pkg/data/models/products/general"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductService struct {
	repo repository.Product
}

func NewProductService(repo repository.Product) *ProductService {
	return &ProductService{repo: repo}
}

func (p *ProductService) Create(product products.Product, productType string) (primitive.ObjectID, error) {
	return p.repo.Create(product, productType)
}

func (p *ProductService) GetById(id primitive.ObjectID, productType string) (products.Product, error) {
	return p.repo.GetById(id, productType)
}

func (p *ProductService) GetStandardizedList(filter bson.M, productType string) ([]general.StandardizedProductData, error) {
	productsList, err := p.repo.GetList(filter, productType)
	if err != nil {
		return nil, err
	}

	standardizedProductsList := make([]general.StandardizedProductData, len(productsList))
	for i, product := range productsList {
		standardizedProductsList[i] = product.Standardize()
	}

	return standardizedProductsList, nil
}

func (p *ProductService) GetList(filter bson.M, productType string) ([]products.Product, error) {
	return p.repo.GetList(filter, productType)
}

func (p *ProductService) DeleteById(id primitive.ObjectID, productType string) error {
	return p.repo.DeleteById(id, productType)
}

func (p *ProductService) UpdateById(id primitive.ObjectID, input products.ProductInput, productType string) error {
	return p.repo.UpdateById(id, input, productType)
}

func (p *ProductService) UpdateGeneralInfoById(id primitive.ObjectID, input general.UpdateGeneralInput, productType string) error {
	return p.repo.UpdateGeneralInfoById(id, input, productType)
}
