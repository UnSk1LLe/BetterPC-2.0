package products

import (
	"BetterPC_2.0/pkg/data/models/products/general"
	generalResponses "BetterPC_2.0/pkg/data/models/products/general/responses"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
	"strings"
)

type Cooling struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	General    general.General    `bson:"general"`
	Type       string             `bson:"type"`
	Sockets    []string           `bson:"sockets"`
	Fans       []int              `bson:"fans"`
	Rpm        []int              `bson:"rpm"`
	Tdp        int                `bson:"tdp"`
	NoiseLevel int                `bson:"noise_level"`
	MountType  string             `bson:"mount_type"`
	Power      int                `bson:"power"`
	Height     int                `bson:"height"`
}

func (cooling Cooling) GetModel() string {
	return cooling.General.Model
}

func (cooling Cooling) GetStock() int {
	return cooling.General.Amount
}

func (cooling Cooling) GetImage() string {
	return cooling.General.Image
}

func (cooling Cooling) SetImage(imageName string) {
	cooling.General.Image = imageName
}

func (cooling Cooling) Standardize() generalResponses.StandardizedProductData {
	var product generalResponses.StandardizedProductData
	product.ProductHeader.ID = cooling.ID.Hex()
	product.ProductHeader.ProductType = "cooling"
	product.Name = cooling.General.Model
	product.General = cooling.General
	product.Description = "Type: " + cooling.Type + ", Sockets: " + strings.Join(cooling.Sockets, ", ") +
		", Fans: " + strconv.Itoa(len(cooling.Fans)) + ", TDP: " + strconv.Itoa(cooling.Tdp) + "W"
	return product
}

func (cooling Cooling) CalculateFinalPrice() int {
	return cooling.General.GetFinalPrice()
}
