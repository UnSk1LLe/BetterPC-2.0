package orders

import (
	"BetterPC_2.0/pkg/data/helpers/validators"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Order struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Items  []Item             `bson:"items" json:"items"`
	UserID primitive.ObjectID `bson:"user_id" json:"user_id"`
	Date   time.Time          `bson:"date" json:"date"`
	Price  int                `bson:"price" json:"price"`
	Status string             `bson:"status" json:"status"`
}

type Item struct {
	ItemHeader   ItemHeader `bson:"item_header" json:"itemHeader"`
	Manufacturer string     `bson:"manufacturer" json:"manufacturer"`
	Model        string     `bson:"model,omitempty" json:"model"`
	Price        int        `bson:"price,omitempty" json:"price"`
	MaxAmount    int        `bson:"max_amount,omitempty" json:"maxAmount"`
}

type ItemHeader struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	ProductType string             `bson:"product_type"`
	Amount      int                `bson:"amount"`
}

func (i Item) ItemFinalPrice() int {
	finalPrice := i.ItemHeader.Amount * i.Price
	return finalPrice
}

func (o *Order) CalculateOrderPrice() {
	totalPrice := 0
	for _, item := range o.Items {
		totalPrice += item.ItemFinalPrice()
	}
	o.Price = totalPrice
}

type UpdateOrderInput struct {
	Items  *[]Item             `bson:"items" json:"items"`
	UserID *primitive.ObjectID `bson:"user_id" json:"user_id"`
	Date   *time.Time          `bson:"date" json:"date"`
	Price  *int                `bson:"price" json:"price"`
	Status *string             `bson:"status" json:"status"`
}

func (o *UpdateOrderInput) Validate() error {
	return validators.ValidateStruct(&o)
}
