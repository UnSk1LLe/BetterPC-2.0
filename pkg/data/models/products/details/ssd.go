package details

import (
	"BetterPC_2.0/pkg/data/models/products/general"
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

func (ssd Ssd) GetProductModel() string {
	return ssd.General.Model
}

func (ssd Ssd) Standardize() general.StandardizedProductData {
	var product general.StandardizedProductData
	product.ProductHeader.ID = ssd.ID.Hex()
	product.ProductHeader.ProductType = "ssd"
	product.Name = ssd.General.Model
	product.General = ssd.General
	product.Description = "Type: " + ssd.Type + ", Capacity: " + strconv.Itoa(ssd.Capacity) + "GB, " +
		"Interface: " + ssd.Interface + ", Read Speed: " + strconv.Itoa(ssd.Read) + "MB/s, " +
		"Write Speed: " + strconv.Itoa(ssd.Write) + "MB/s"
	return product
}

func (ssd Ssd) ProductFinalPrice() int {
	return ssd.General.CalculateFinalPrice()
}

type UpdateSsdInput struct {
	General    general.General `bson:"general" json:"general"`
	Type       string          `bson:"type" json:"type"`
	Capacity   int             `bson:"capacity" json:"capacity"`
	Interface  string          `bson:"interface" json:"interface"`
	MemoryType string          `bson:"memory_type" json:"memory_type"`
	Read       int             `bson:"read" json:"read"`
	Write      int             `bson:"write" json:"write"`
	FormFactor string          `bson:"form_factor" json:"form_factor"`
	Mftb       float64         `bson:"mftb" json:"mftb"`
	Size       []float64       `bson:"size" json:"size"`
	Weight     int             `bson:"weight" json:"weight"`
}

func (s UpdateSsdInput) Validate() error {
	return general.ValidateStruct(&s)
}
