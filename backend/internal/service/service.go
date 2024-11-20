package service

import (
	"BetterPC_2.0/internal/repository"
	"BetterPC_2.0/pkg/data/models/categories"
	"BetterPC_2.0/pkg/data/models/orders"
	orderFilters "BetterPC_2.0/pkg/data/models/orders/filters"
	orderRequests "BetterPC_2.0/pkg/data/models/orders/requests"
	"BetterPC_2.0/pkg/data/models/products"
	generalRequests "BetterPC_2.0/pkg/data/models/products/general/requests"
	generalResponses "BetterPC_2.0/pkg/data/models/products/general/responses"
	productRequests "BetterPC_2.0/pkg/data/models/products/requests"
	"BetterPC_2.0/pkg/data/models/users"
	userFilters "BetterPC_2.0/pkg/data/models/users/filters"
	userAuthRequests "BetterPC_2.0/pkg/data/models/users/requests/auth"
	userUpdateRequests "BetterPC_2.0/pkg/data/models/users/requests/patch"
	userResponses "BetterPC_2.0/pkg/data/models/users/responses"
	"BetterPC_2.0/pkg/email/smtpServer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Authorization interface {
	CreateUser(input userAuthRequests.RegisterRequest) (primitive.ObjectID, error)
	GenerateTokenPair(email, password string) (TokenPair, error)
	ParseAccessToken(accessToken string) (userResponses.UserResponse, error)
	RefreshTokens(refreshToken string) (userResponses.UserResponse, TokenPair, error)
	HasRole(userId string, roles ...string) (bool, error)
}

type Verification interface {
	SetNewToken(email string, tokenTTL time.Duration) (string, error)
	VerifyUser(token string) error
	IsVerifiedUser(email string) (bool, error)

	GenerateRecoveryToken(email string) (string, error)
	UpdatePasswordByToken(password, token string) error
}

type Notification interface {
	SendEmailToUser(userEmail, subject, body string) error
	SendEmailToGroup(userEmailList []string, subject, body string) error
}

type User interface {
	Create(user users.User) (string, error)
	Update(userId string, input userUpdateRequests.UpdateUserInfoRequest) error
	Delete(userId string) error
	GetList(filters userFilters.AdminUserFilters) ([]users.User, error)
	GetById(userId string) (users.User, error)
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
	GetStandardizedList(filter bson.M, productType products.ProductType) ([]generalResponses.StandardizedProductData, error)
	UpdateById(id primitive.ObjectID, input productRequests.ProductUpdateRequest, productType products.ProductType) error
	UpdateGeneralInfoById(productId primitive.ObjectID, input generalRequests.UpdateGeneralRequest, productType products.ProductType) error
	DeleteById(productId primitive.ObjectID, productType products.ProductType) error
}

type Filters interface {
	SetSearchFilter()
	GetSearchFilter() bson.M
	SetBuildFilter()
	GetBuildFilter() bson.M
}

type Order interface {
	CreateWithItemHeaders(userId string, itemHeaders orderRequests.CreateOrderRequest) (primitive.ObjectID, error)
	Update(orderId string, input orders.UpdateOrderInput) error
	CancelOrder(orderId string) error
	SetStatus(orderId string, status string) error
	Delete(orderId string) error
	GetById(orderId string) (orders.Order, error)
	GetList(filter orderFilters.AdminOrderFilters) ([]orders.Order, error)
}

type Service struct {
	Categories
	Authorization
	Verification
	Notification
	User
	Product
	Order
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Categories:    NewCategoryService(repos.Categories),
		Authorization: NewAuthService(repos.Authorization, repos.User),
		Verification:  NewVerificationService(repos.Verification),
		Notification:  NewNotificationService(smtpServer.MustGet()),
		Product:       NewProductService(repos.Product),
		Order:         NewOrderService(repos.Order, repos.Product),
		User:          NewUserService(repos.User),
	}
}
