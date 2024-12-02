package service

import (
	"BetterPC_2.0/internal/repository"
	"BetterPC_2.0/pkg/data/models/orders"
	orderErrors "BetterPC_2.0/pkg/data/models/orders/errors"
	orderFilters "BetterPC_2.0/pkg/data/models/orders/filters"
	orderRequests "BetterPC_2.0/pkg/data/models/orders/requests"
	"BetterPC_2.0/pkg/data/models/products"
	productErrors "BetterPC_2.0/pkg/data/models/products/errors"
	"BetterPC_2.0/pkg/logging"
	"fmt"
	"github.com/pkg/errors"
	"github.com/stripe/stripe-go/v81"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderService struct {
	repo          repository.Order
	productRepo   repository.Product
	stripeService *StripeService
	logger        *logging.Logger
}

func NewOrderService(repo repository.Order, productRepo repository.Product, stripeService *StripeService, logger *logging.Logger) *OrderService {
	return &OrderService{
		repo:          repo,
		productRepo:   productRepo,
		stripeService: stripeService,
		logger:        logger,
	}
}

func (o *OrderService) CreateWithItemHeaders(userId string, input orderRequests.CreateOrderRequest) (primitive.ObjectID, error) {
	var order orders.Order

	userObjId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return primitive.NilObjectID, err
	}

	productTypeItemHeaders := input.ProductTypeItemHeaders

	var productList = make(map[products.ProductType][]orders.ProductHeader)

	for productType, items := range productTypeItemHeaders {
		var productTypeHeaders []orders.ProductHeader
		//Set information for each product of a certain type
		for _, item := range items { //TODO optimization using GetById for single item
			var productHeader orders.ProductHeader
			var err error

			//Getting ID from string
			productHeader.ID, err = primitive.ObjectIDFromHex(item.ID)
			if err != nil {
				return primitive.NilObjectID, err
			}

			//Getting amount
			productHeader.SelectedAmount = item.SelectedAmount

			//Getting price from DB
			product, err := o.productRepo.GetById(productHeader.ID, productType)
			if err != nil {
				return primitive.NilObjectID, err
			}

			//Checking the availability
			if product.GetStock() < productHeader.SelectedAmount {
				message := fmt.Sprintf("not enough %s with model %s", productType, product.GetModel())
				return primitive.NilObjectID, errors.Wrap(productErrors.ErrInsufficientStock, message)
			}

			productHeader.Price = product.CalculateFinalPrice()

			//Adding product info
			productTypeHeaders = append(productTypeHeaders, productHeader)
		}
		productList[productType] = append(productList[productType], productTypeHeaders...)
	}

	order = orders.NewOrder(userObjId, productList)

	return o.repo.Create(order)
}

func (o *OrderService) PayForOrder(userId, orderId string, amount int64, currency, paymentMethodId, returnUrl string) (string, error) {
	userObjId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return "", err
	}

	orderObjId, err := primitive.ObjectIDFromHex(orderId)
	if err != nil {
		return "", err
	}

	order, err := o.repo.GetById(orderObjId)
	if err != nil {
		return "", err
	}

	if order.UserID != userObjId {
		return "", orderErrors.ErrOrderOwnerMismatch
	}

	metadata := map[string]string{
		"order_id": orderId,
		"user_id":  userId,
	}

	if currency != "USD" {
		return "", errors.New("payment must be in USD")
	}

	paymentIntent, err := o.stripeService.CreatePaymentIntent(amount, currency, paymentMethodId, returnUrl, metadata)
	if err != nil {
		return "", fmt.Errorf("failed to create payment intent: %w", err)
	}

	if paymentIntent.Status == stripe.PaymentIntentStatusSucceeded {
		input := orders.PaymentDetails{
			PaymentIntentId: paymentIntent.ID,
			IsPaid:          true,
		}
		err := o.repo.UpdatePaymentDetails(orderObjId, input)
		if err != nil {
			return paymentIntent.ID, fmt.Errorf("failed to update order status: %w", err)
		}
	} else {
		return paymentIntent.ID, fmt.Errorf("payment was not successful, status: %s", paymentIntent.Status)
	}

	return paymentIntent.ID, nil
}

func (o *OrderService) CancelOrder(userId, orderId string) error {
	orderObjId, err := primitive.ObjectIDFromHex(orderId)
	if err != nil {
		return err
	}

	userObjId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}

	order, err := o.repo.GetById(orderObjId)
	if err != nil {
		return err
	}

	//only proceed if the order is active
	if !order.IsActive() {
		return orderErrors.ErrNotActiveOrder
	}

	err = o.repo.Cancel(userObjId, orderObjId)
	if err != nil {
		return err
	}

	//TODO add logic if order is paid then give money back to the client, else cancel order
	if order.Payment.IsPaid && order.IsActive() {
		_, err := o.stripeService.RefundPayment(order.Payment.PaymentIntentId, int64(order.Price))
		if err != nil {
			return err
		}
	}

	return nil
}

/*func (o *OrderService) Update(orderId string, input orders.UpdateOrderInput) error {
	orderObjId, err := primitive.ObjectIDFromHex(orderId)
	if err != nil {
		return err
	}

	return o.repo.Update(orderObjId, input)
}*/

func (o *OrderService) SetStatus(orderId string, status string) error {
	orderObjId, err := primitive.ObjectIDFromHex(orderId)
	if err != nil {
		return err
	}

	orderStatus, err := orders.ParseOrderStatus(status)
	if err != nil {
		return err
	}
	if orderStatus == orders.OrderStatuses.Cancelled || orderStatus == orders.OrderStatuses.Closed {
		return orderErrors.ErrOrderCancelled
	}

	return o.repo.SetStatus(orderObjId, orderStatus)
}

func (o *OrderService) Delete(orderId string) error {
	orderObjId, err := primitive.ObjectIDFromHex(orderId)
	if err != nil {
		return err
	}

	return o.repo.Delete(orderObjId)
}

func (o *OrderService) GetById(orderId string) (orders.Order, error) {
	var order orders.Order

	orderObjId, err := primitive.ObjectIDFromHex(orderId)
	if err != nil {
		return order, err
	}

	order, err = o.repo.GetById(orderObjId)
	if err != nil {
		return order, err
	}

	return order, nil
}

func (o *OrderService) GetUserOrders(userId string) ([]orders.Order, error) {
	userObjId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, errors.Wrap(err, "invalid user id: %v")
	}

	bsonFilter := bson.M{"user_id": userObjId}

	return o.repo.GetList(bsonFilter)
}

func (o *OrderService) GetUserOrder(userId, orderId string) (orders.Order, error) {
	var order orders.Order

	userObjId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return order, errors.Wrap(err, "invalid user id: %v")
	}

	orderObjId, err := primitive.ObjectIDFromHex(orderId)
	if err != nil {
		return order, errors.Wrap(err, "invalid order id: %v")
	}

	bsonFilter := bson.M{"_id": orderObjId, "user_id": userObjId}

	return o.repo.GetOneByFilter(bsonFilter)
}

func (o *OrderService) GetList(filters orderFilters.AdminOrderFilters) ([]orders.Order, error) {
	var bsonFilter bson.M

	if !filters.DateFrom.IsZero() {
		bsonFilter["created_at"] = bson.M{"$gte": filters.DateFrom}
	}
	if !filters.DateTo.IsZero() {
		if existingFilter, ok := bsonFilter["created_at"].(bson.M); ok {
			existingFilter["$lte"] = filters.DateTo
		} else {
			bsonFilter["created_at"] = bson.M{"$lte": filters.DateTo}
		}
	}

	if filters.UserId != "" {
		userId, err := primitive.ObjectIDFromHex(filters.UserId)
		if err != nil {
			return nil, errors.Wrap(err, "invalid user_id: %v")
		}
		bsonFilter["user_id"] = userId
	}

	if len(filters.ProductTypes) > 0 {
		productTypeQuery := bson.A{}
		for _, productType := range filters.ProductTypes {
			productTypeQuery = append(productTypeQuery, bson.M{
				"product_list." + productType: bson.M{"$exists": true},
			})
		}
		bsonFilter["$or"] = productTypeQuery
	}

	if filters.PriceFrom > 0 || filters.PriceTo > 0 {
		priceFilter := bson.M{}
		if filters.PriceFrom > 0 {
			priceFilter["$gte"] = filters.PriceFrom
		}
		if filters.PriceTo > 0 {
			priceFilter["$lte"] = filters.PriceTo
		}
		bsonFilter["price"] = priceFilter
	}

	if len(filters.Statuses) > 0 {
		bsonFilter["status"] = bson.M{"$in": filters.Statuses}
	}

	return o.repo.GetList(bsonFilter)
}
