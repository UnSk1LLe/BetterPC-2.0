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
	return primitive.NilObjectID, nil
}

func (p *ProductService) GetById(id primitive.ObjectID, productType string) (products.Product, error) {
	return nil, nil
}

func (p *ProductService) GetList(filter bson.M, productType string) ([]products.Product, error) {
	return nil, nil
}

func (p *ProductService) DeleteById(id primitive.ObjectID, productType string) error {
	return nil
}

func (p *ProductService) UpdateById(id primitive.ObjectID, input products.ProductInput, productType string) error {
	return nil
}

func (p *ProductService) UpdateGeneralInfoById(id primitive.ObjectID, input general.UpdateGeneralInput, productType string) error {
	return nil
}
