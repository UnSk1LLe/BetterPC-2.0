package service

import (
	"BetterPC_2.0/internal/repository"
	"BetterPC_2.0/pkg/data/models/orders"
	"BetterPC_2.0/pkg/data/models/products"
	"BetterPC_2.0/pkg/data/models/products/general"
	"BetterPC_2.0/pkg/data/models/users"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product interface {
	Create(product products.Product, productType string) (primitive.ObjectID, error)
	GetById(id primitive.ObjectID, productType string) (products.Product, error)
	GetList(filter bson.M, productType string) ([]products.Product, error)
	UpdateById(id primitive.ObjectID, input products.ProductInput, productType string) error
	UpdateGeneralInfoById(productId primitive.ObjectID, input general.UpdateGeneralInput, productType string) error
	DeleteById(productId primitive.ObjectID, productType string) error
}

type Authorization interface {
	CreateUser(user users.User) (primitive.ObjectID, error)
	GenerateTokenPair(email, password string) (TokenPair, error)
	ParseAccessToken(accessToken string) (string, error)
	RefreshTokens(refreshToken string) (TokenPair, error)
}

type User interface {
	Create(user users.User) (primitive.ObjectID, error)
	Update(userId primitive.ObjectID, input users.UpdateUserInput) error
	Delete(userId primitive.ObjectID) error
	GetAll(filter bson.M)
	GetById(userId primitive.ObjectID)
}

type Order interface {
	Create(order orders.Order) (primitive.ObjectID, error)
	Update(orderId primitive.ObjectID, input orders.UpdateOrderInput) error
	SetStatus(orderId primitive.ObjectID, status string) error
	Delete(orderId primitive.ObjectID) error
}

type Service struct {
	Authorization
	User
	Product
	Order
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Product:       NewProductService(repos.Product),
	}
}
