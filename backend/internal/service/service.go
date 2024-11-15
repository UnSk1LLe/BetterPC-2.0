package service

import (
	"BetterPC_2.0/internal/repository"
	"BetterPC_2.0/pkg/data/models/categories"
	"BetterPC_2.0/pkg/data/models/orders"
	"BetterPC_2.0/pkg/data/models/products"
	generalRequests "BetterPC_2.0/pkg/data/models/products/general/requests"
	generalResponses "BetterPC_2.0/pkg/data/models/products/general/responses"
	productRequests "BetterPC_2.0/pkg/data/models/products/requests"
	"BetterPC_2.0/pkg/data/models/users"
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
	HasRole(userId primitive.ObjectID, roles ...string) (bool, error)
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
	Create(user users.User) (primitive.ObjectID, error)
	Update(userId primitive.ObjectID, input userUpdateRequests.UpdateUserInfoRequest) error
	Delete(userId primitive.ObjectID) error
	GetList(filter bson.M) ([]users.User, error)
	GetById(userId primitive.ObjectID) (users.User, error)
}

type Categories interface {
	GetList(filter bson.M) ([]categories.Category, error)
	GetById(id primitive.ObjectID) (categories.Category, error)
	UpdateById(id primitive.ObjectID, input categories.UpdateCategoryInput) error
}

type Product interface {
	Create(product products.Product, productType string) (primitive.ObjectID, error)
	GetById(id primitive.ObjectID, productType string) (products.Product, error)
	GetList(filter bson.M, productType string) ([]products.Product, error)
	GetStandardizedList(filter bson.M, productType string) ([]generalResponses.StandardizedProductData, error)
	UpdateById(id primitive.ObjectID, input productRequests.ProductUpdateRequest, productType string) error
	UpdateGeneralInfoById(productId primitive.ObjectID, input generalRequests.UpdateGeneralRequest, productType string) error
	DeleteById(productId primitive.ObjectID, productType string) error
}

type Filters interface {
	SetSearchFilter()
	GetSearchFilter() bson.M
	SetBuildFilter()
	GetBuildFilter() bson.M
}

type Order interface {
	Create(order orders.Order) (primitive.ObjectID, error)
	Update(orderId primitive.ObjectID, input orders.UpdateOrderInput) error
	SetStatus(orderId primitive.ObjectID, status string) error
	Delete(orderId primitive.ObjectID) error
	GetById(id primitive.ObjectID) (orders.Order, error)
	GetList(filter bson.M) ([]orders.Order, error)
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
		Order:         NewOrderService(repos.Order),
		User:          NewUserService(repos.User),
	}
}
