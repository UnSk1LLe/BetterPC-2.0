package service

import (
	"BetterPC_2.0/configs"
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
	adminUserRequests "BetterPC_2.0/pkg/data/models/users/requests/admin"
	userAuthRequests "BetterPC_2.0/pkg/data/models/users/requests/auth"
	userUpdateRequests "BetterPC_2.0/pkg/data/models/users/requests/patch"
	userResponses "BetterPC_2.0/pkg/data/models/users/responses"
	"BetterPC_2.0/pkg/email/smtpServer"
	"BetterPC_2.0/pkg/logging"
	"github.com/stripe/stripe-go/v81"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mime/multipart"
	"time"
)

const StaticFilesPath = "./static"

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
	Create(user adminUserRequests.CreateUserRequest) (string, error)
	UpdateUserInfo(userId string, input userUpdateRequests.UpdateUserInfoRequest) error
	SetRole(userId string, role string) error
	Delete(userId string) error
	GetList(filters userFilters.AdminUserFilters) ([]users.User, error)
	GetById(userId string) (users.User, error)
	UpdateUserImage(userId string, image *multipart.FileHeader) error

	AttachPaymentMethodToUser(userId, paymentMethodId string) error
	RemovePaymentMethod(userId, paymentMethodId string) error
	ListPaymentMethodsByUser(userId string) ([]*stripe.PaymentMethod, error)
}

type Categories interface {
	GetList(filter bson.M) ([]categories.Category, error)
	GetById(id primitive.ObjectID) (categories.Category, error)
	UpdateById(id primitive.ObjectID, input categories.UpdateCategoryInput) error
}

type Product interface {
	Create(product products.Product, productType products.ProductType, image *multipart.FileHeader) (primitive.ObjectID, error)
	GetById(id primitive.ObjectID, productType products.ProductType) (products.Product, error)
	GetList(filter bson.M, productType products.ProductType) ([]products.Product, error)
	GetStandardizedList(filter bson.M, productType products.ProductType) ([]generalResponses.StandardizedProductData, error)
	UpdateById(id primitive.ObjectID, input productRequests.ProductUpdateRequest,
		productType products.ProductType, image *multipart.FileHeader) error
	UpdateGeneralInfoById(productId primitive.ObjectID, input generalRequests.UpdateGeneralRequest,
		productType products.ProductType, image *multipart.FileHeader) error
	DeleteById(productId primitive.ObjectID, productType products.ProductType) error
}

type Files interface {
	AddUserImage(file *multipart.FileHeader) (string, error)
	AddProductImage(file *multipart.FileHeader) (string, error)
	AddImage(file *multipart.FileHeader, subDirectory string) (string, error)
	DeleteUserImage(imageName string) error
	DeleteProductImage(imageName string) error
	DeleteImage(imageName, subDirectory string) error
}

type Order interface {
	CreateWithItemHeaders(userId string, itemHeaders orderRequests.CreateOrderRequest) (primitive.ObjectID, error)
	//Update(orderId string, input orders.UpdateOrderInput) error
	PayForOrder(userId, orderId string, amount int64, currency, paymentMethodId, returnUrl string) (string, error)
	CancelOrder(userId, orderId string) error
	SetStatus(orderId string, status string) error
	Delete(orderId string) error
	GetById(orderId string) (orders.Order, error)
	GetUserOrders(userId string) ([]orders.Order, error)
	GetUserOrder(userId, orderId string) (orders.Order, error)
	GetList(filter orderFilters.AdminOrderFilters) ([]orders.Order, error)
}

type Stripe interface {
	CreateCustomer(email string, metadata map[string]string) (string, error)
	AttachPaymentMethodToCustomer(customerId, paymentMethodId string) error
	RemovePaymentMethod(paymentMethodId string) error
	ListPaymentMethodsByCustomer(customerId string) ([]*stripe.PaymentMethod, error)
	CreatePaymentIntent(amount int64, currency, paymentMethodID, returnUrl string, metadata map[string]string) (*stripe.PaymentIntent, error)
	GetPaymentIntent(paymentIntentId string, params *stripe.PaymentIntentParams) (*stripe.PaymentIntent, error)
	RefundPayment(chargeId string, amount int64) (*stripe.Refund, error)
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

func NewService(repos *repository.Repository, logger *logging.Logger) *Service {

	fileService := NewFileService(StaticFilesPath)
	notificationService := NewNotificationService(smtpServer.MustGet(), logger)
	stripeService := NewStripeService(configs.GetConfig().Stripe.PrivateKey)

	return &Service{
		Categories:    NewCategoryService(repos.Categories, logger),
		Authorization: NewAuthService(repos.Authorization, repos.User, logger),
		Verification:  NewVerificationService(repos.Verification, notificationService, logger),
		Notification:  NewNotificationService(smtpServer.MustGet(), logger),
		Product:       NewProductService(repos.Product, fileService, logger),
		Order:         NewOrderService(repos.Order, repos.Product, stripeService, logger),
		User:          NewUserService(repos.User, fileService, stripeService, logger),
	}
}
