package products

import (
	"BetterPC_2.0/pkg/data/models/products/general"
	generalResponses "BetterPC_2.0/pkg/data/models/products/general/responses"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
)

type Ssd struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	General    general.General    `bson:"general"`
	Type       string             `bson:"type"`
	Capacity   int                `bson:"capacity"`
	Interface  string             `bson:"interface"`
	MemoryType string             `bson:"memory_type"`
	Read       int                `bson:"read"`
	Write      int                `bson:"write"`
	FormFactor string             `bson:"form_factor"`
	Mftb       float64            `bson:"mftb"`
	Size       []float64          `bson:"size"`
	Weight     int                `bson:"weight"`
}

func (ssd Ssd) GetModel() string {
	return ssd.General.Model
}

func (ssd Ssd) GetStock() int {
	return ssd.General.Amount
}

func (ssd Ssd) GetImage() string {
	return ssd.General.Image
}

func (ssd Ssd) SetImage(imageName string) {
	ssd.General.Image = imageName
}

func (ssd Ssd) Standardize() generalResponses.StandardizedProductData {
	var product generalResponses.StandardizedProductData
	product.ProductHeader.ID = ssd.ID.Hex()
	product.ProductHeader.ProductType = "ssd"
	product.Name = ssd.General.Model
	product.General = ssd.General
	product.Description = "Type: " + ssd.Type + ", Capacity: " + strconv.Itoa(ssd.Capacity) + "GB, " +
		"Interface: " + ssd.Interface + ", Read Speed: " + strconv.Itoa(ssd.Read) + "MB/s, " +
		"Write Speed: " + strconv.Itoa(ssd.Write) + "MB/s"
	return product
}

func (ssd Ssd) CalculateFinalPrice() int {
	return ssd.General.GetFinalPrice()
}
