package orders

import (
	"BetterPC_2.0/pkg/data/helpers/validators"
	orderErrors "BetterPC_2.0/pkg/data/models/orders/errors"
	"BetterPC_2.0/pkg/data/models/products"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
	"time"
)

const MaxStatusLength = 20

type OrderStatus string

var OrderStatuses = struct {
	Created   OrderStatus
	Pending   OrderStatus
	Delivered OrderStatus
	Cancelled OrderStatus
	Closed    OrderStatus
}{
	Created:   "CREATED",
	Pending:   "PENDING",
	Delivered: "DELIVERED",
	Cancelled: "CANCELLED",
	Closed:    "CLOSED",
}

func ParseOrderStatus(input string) (OrderStatus, error) {

	if len(input) > MaxStatusLength {
		return "", errors.Wrapf(orderErrors.ErrInvalidInput, "product type must be shorter than %d characters", MaxStatusLength)
	}

	normalizedInput := strings.ToUpper(strings.TrimSpace(input))
	if len(normalizedInput) == 0 {
		return "", errors.Wrap(orderErrors.ErrInvalidInput, "product type cannot not be empty")
	}

	switch OrderStatus(normalizedInput) {
	case OrderStatuses.Created:
		return OrderStatuses.Created, nil
	case OrderStatuses.Pending:
		return OrderStatuses.Pending, nil
	case OrderStatuses.Delivered:
		return OrderStatuses.Delivered, nil
	case OrderStatuses.Cancelled:
		return OrderStatuses.Cancelled, nil
	}

	return "", orderErrors.ErrUnsupportedOrderStatus
}

type Order struct {
	ID          primitive.ObjectID                       `bson:"_id,omitempty" json:"id"`
	ProductList map[products.ProductType][]ProductHeader `bson:"product_list" json:"product_list"`
	UserID      primitive.ObjectID                       `bson:"user_id" json:"user_id"`
	Price       int                                      `bson:"price" json:"price"`
	Status      OrderStatus                              `bson:"status" json:"status"`
	Payment     PaymentDetails                           `bson:"payment" json:"payment"`
	Refunds     []RefundDetails                          `bson:"refunds,omitempty" json:"refunds,omitempty"`
	CreatedAt   primitive.DateTime                       `bson:"created_at" json:"created_at"`
	UpdatedAt   primitive.DateTime                       `bson:"updated_at" json:"updated_at"`
}

type PaymentDetails struct {
	PaymentIntentId string `bson:"intent_id,omitempty" json:"intent_id,omitempty"`
	IsPaid          bool   `bson:"is_paid" json:"is_paid"`
}

type RefundDetails struct {
	RefundID   string             `bson:"refund_id" json:"refund_id"`
	Amount     int64              `bson:"amount" json:"amount"`
	Currency   string             `bson:"currency" json:"currency"`
	RefundedAt primitive.DateTime `bson:"refunded_at" json:"refunded_at"`
}

type ProductHeader struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Price          int                `bson:"price" json:"price"`
	SelectedAmount int                `bson:"selected_amount" json:"selected_amount"`
}

func (o *Order) IsActive() bool {
	return !(o.Status == OrderStatuses.Closed || o.Status == OrderStatuses.Cancelled)
}

func NewOrder(userId primitive.ObjectID, productHeaders map[products.ProductType][]ProductHeader) Order {
	return Order{
		ID:          primitive.NewObjectID(),
		ProductList: productHeaders,
		UserID:      userId,
		Price:       CalculateOrderPrice(productHeaders),
		Status:      OrderStatuses.Created,
		Payment: PaymentDetails{
			PaymentIntentId: "",
			IsPaid:          false,
		},
		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
		UpdatedAt: primitive.NewDateTimeFromTime(time.Now()),
	}
}

func CalculateOrderPrice(products map[products.ProductType][]ProductHeader) int {
	totalPrice := 0
	for _, headers := range products {
		for _, header := range headers {
			totalPrice += header.Price * header.SelectedAmount
		}
	}
	return totalPrice
}

type UpdateOrderInput struct {
	RemovedItems map[products.ProductType][]ItemHeader `json:"removed_items"`
	AddedItems   map[products.ProductType][]ItemHeader `json:"added_items"`
}

func (o *UpdateOrderInput) Validate() error {
	return validators.ValidateStruct(&o)
}
