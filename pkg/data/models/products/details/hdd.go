package details

import (
	"BetterPC_2.0/pkg/data/models/products/general"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
)

type Hdd struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	General      general.General    `bson:"general"`
	Type         string             `bson:"type"`
	Capacity     int                `bson:"capacity"`
	Interface    string             `bson:"interface"`
	WriteMethod  string             `bson:"write_method"`
	TransferRate int                `bson:"transfer_rate"`
	SpindleSpeed int                `bson:"spindle_speed"`
	FormFactor   string             `bson:"form_factor"`
	Mftb         int                `bson:"mftb"`
	Size         []float64          `bson:"size"`
	Weight       int                `bson:"weight"`
}

func (hdd Hdd) GetProductModel() string {
	return hdd.General.Model
}

func (hdd Hdd) Standardize() general.StandardizedProductData {
	var product general.StandardizedProductData
	product.ProductHeader.ID = hdd.ID.Hex()
	product.ProductHeader.ProductType = "hdd"
	product.Name = hdd.General.Model
	product.General = hdd.General
	product.Description = "Type: " + hdd.Type + ", Capacity: " + strconv.Itoa(hdd.Capacity) + "GB, " +
		"Interface: " + hdd.Interface + ", Spindle Speed: " + strconv.Itoa(hdd.SpindleSpeed) + "RPM"
	return product
}

func (hdd Hdd) ProductFinalPrice() int {
	return hdd.General.CalculateFinalPrice()
}

type UpdateHddInput struct {
	General      *general.General `bson:"general"`
	Type         *string          `bson:"type"`
	Capacity     *int             `bson:"capacity"`
	Interface    *string          `bson:"interface"`
	WriteMethod  *string          `bson:"write_method"`
	TransferRate *int             `bson:"transfer_rate"`
	SpindleSpeed *int             `bson:"spindle_speed"`
	FormFactor   *string          `bson:"form_factor"`
	Mftb         *int             `bson:"mftb"`
	Size         *[]float64       `bson:"size"`
	Weight       *int             `bson:"weight"`
}

func (h UpdateHddInput) Validate() error {
	return general.ValidateStruct(&h)
}
