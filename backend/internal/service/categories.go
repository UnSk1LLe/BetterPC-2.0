package service

import (
	"BetterPC_2.0/internal/repository"
	"BetterPC_2.0/pkg/data/models/categories"
	"BetterPC_2.0/pkg/logging"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CategoryService struct {
	repo repository.Categories
}

func NewCategoryService(repo repository.Categories, logger *logging.Logger) *CategoryService {
	return &CategoryService{repo: repo}
}

func (c *CategoryService) GetList(filter bson.M) ([]categories.Category, error) {
	return c.repo.GetList(filter)
}

func (c *CategoryService) GetById(id primitive.ObjectID) (categories.Category, error) {
	return c.repo.GetById(id)
}

func (c *CategoryService) UpdateById(id primitive.ObjectID, input categories.UpdateCategoryInput) error {
	return c.repo.UpdateById(id, input)
}
