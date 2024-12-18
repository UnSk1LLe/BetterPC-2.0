package repository

import (
	"BetterPC_2.0/internal/repository/database/mongoDb"
	"BetterPC_2.0/pkg/data/models/categories"
	"BetterPC_2.0/pkg/data/models/orders"
	"BetterPC_2.0/pkg/data/models/products"
	generalRequests "BetterPC_2.0/pkg/data/models/products/general/requests"
	productRequests "BetterPC_2.0/pkg/data/models/products/requests"
	"BetterPC_2.0/pkg/data/models/users"
	userUpdateRequests "BetterPC_2.0/pkg/data/models/users/requests/patch"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Authorization interface {
	CreateUser(user users.User) (primitive.ObjectID, error)
	GetUserByEmail(email string) (users.User, error)
	CheckUserExists(userId primitive.ObjectID) (bool, error)
	HasRole(userId primitive.ObjectID, roles []string) (bool, error)
}

type Verification interface {
	SetTokenByEmail(email, token string, expTime primitive.DateTime) error
	CompareUserToken(token string) (bool, error)
	IsVerifiedUser(email string) (bool, error)
	UpdateVerificationDataById(userId primitive.ObjectID, input userUpdateRequests.UpdateUserVerificationDataRequest) error
	GetUserByVerificationToken(token string) (users.User, error)
	UpdateUserPasswordById(userId primitive.ObjectID, password string) error
}

type User interface {
	Create(user users.User) (primitive.ObjectID, error)
	UpdateUserImageById(userId primitive.ObjectID, imageUrl string) (string, error)
	UpdateUserInfoById(userId primitive.ObjectID, input userUpdateRequests.UpdateUserInfoRequest) error
	SetRole(userId primitive.ObjectID, role users.UserRole) error
	DeleteById(userId primitive.ObjectID) error
	GetList(filter bson.M) ([]users.User, error)
	GetById(userId primitive.ObjectID) (users.User, error)
}

type Categories interface {
	GetList(filter bson.M) ([]categories.Category, error)
	GetById(id primitive.ObjectID) (categories.Category, error)
	UpdateById(id primitive.ObjectID, input categories.UpdateCategoryInput) error
}

type Product interface {
	Create(product products.Product, productType products.ProductType) (primitive.ObjectID, error)
	GetById(id primitive.ObjectID, productType products.ProductType) (products.Product, error)
	GetList(filter bson.M, productType products.ProductType) ([]products.Product, error)
	CountCategoryProducts(filter bson.M, productType products.ProductType) (int, error)
	UpdateById(id primitive.ObjectID, input productRequests.ProductUpdateRequest, productType products.ProductType) error
	UpdateGeneralInfoById(productId primitive.ObjectID, input generalRequests.UpdateGeneralRequest, productType products.ProductType) error
	DeleteById(productId primitive.ObjectID, productType products.ProductType) (products.Product, error)
}

type Order interface {
	Create(order orders.Order) (primitive.ObjectID, error)
	Cancel(userId, orderId primitive.ObjectID) error
	UpdatePaymentDetails(orderId primitive.ObjectID, input orders.PaymentDetails) error
	//Update(orderId primitive.ObjectID, input orders.UpdateOrderInput) error
	SetStatus(orderId primitive.ObjectID, status orders.OrderStatus) error
	Delete(orderId primitive.ObjectID) error
	GetOneByFilter(filter bson.M) (orders.Order, error)
	GetById(id primitive.ObjectID) (orders.Order, error)
	GetList(filter bson.M) ([]orders.Order, error)
}

type Repository struct {
	Categories
	Authorization
	Verification
	User
	Product
	Order
}

func NewRepository(MongoConnection mongoDb.Database) *Repository {
	return &Repository{
		User:          NewUsersMongo(MongoConnection),
		Authorization: NewAuthMongo(MongoConnection),
		Verification:  NewVerificationMongo(MongoConnection),
		Product:       NewProductMongo(MongoConnection),
		Order:         NewOrderMongo(MongoConnection),
	}
}
