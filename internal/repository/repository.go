package repository

import (
	"BetterPC_2.0/pkg/data/models/categories"
	"BetterPC_2.0/pkg/data/models/orders"
	"BetterPC_2.0/pkg/data/models/products"
	"BetterPC_2.0/pkg/data/models/products/general"
	"BetterPC_2.0/pkg/data/models/users"
	"BetterPC_2.0/pkg/database/mongoDb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Categories interface {
	GetList(filter bson.M) ([]categories.Category, error)
	GetById(id primitive.ObjectID) (categories.Category, error)
	UpdateById(id primitive.ObjectID, input categories.UpdateCategoryInput) error
}

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
	GetUser(email string) (users.User, error)
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

type Repository struct {
	Categories
	Authorization
	User
	Product
	Order
}

func NewRepository(MongoConnection *mongoDb.MongoConnection) *Repository {
	return &Repository{
		Authorization: NewAuthMongo(MongoConnection),
		Product:       NewProductMongo(MongoConnection),
	}
}
