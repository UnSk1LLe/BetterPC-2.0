package details

import (
	"BetterPC_2.0/pkg/data/models/products/general"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
)

type Ram struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	General      general.General    `bson:"general"`
	Capacity     int                `bson:"capacity"`
	Number       int                `bson:"number"`
	FormFactor   string             `bson:"form_factor"`
	Rank         int                `bson:"rank"`
	Type         string             `bson:"type"`
	Frequency    int                `bson:"frequency"`
	Bandwidth    int                `bson:"bandwidth"`
	CasLatency   string             `bson:"cas_latency"`
	TimingScheme []int              `bson:"timing_scheme"`
	Voltage      float64            `bson:"voltage"`
	Cooling      string             `bson:"cooling"`
	Height       int                `bson:"height"`
}

func (ram Ram) GetProductModel() string {
	return ram.General.Model
}

func (ram Ram) Standardize() general.StandardizedProductData {
	var product general.StandardizedProductData
	product.ProductHeader.ID = ram.ID.Hex()
	product.ProductHeader.ProductType = "ram"
	product.Name = ram.General.Model
	product.General = ram.General
	product.Description = "Capacity: " + strconv.Itoa(ram.Capacity) + "GB, Type: " + ram.Type +
		", Frequency: " + strconv.Itoa(ram.Frequency) + "MHz, CAS Latency: " + ram.CasLatency
	return product
}
func (ram Ram) ProductFinalPrice() int {
	return ram.General.CalculateFinalPrice()
}

type UpdateRamInput struct {
	General      *general.General `bson:"general"`
	Capacity     *int             `bson:"capacity"`
	Number       *int             `bson:"number"`
	FormFactor   *string          `bson:"form_factor"`
	Rank         *int             `bson:"rank"`
	Type         *string          `bson:"type"`
	Frequency    *int             `bson:"frequency"`
	Bandwidth    *int             `bson:"bandwidth"`
	CasLatency   *string          `bson:"cas_latency"`
	TimingScheme *[]int           `bson:"timing_scheme"`
	Voltage      *float64         `bson:"voltage"`
	Cooling      *string          `bson:"cooling"`
	Height       *int             `bson:"height"`
}

func (r UpdateRamInput) Validate() error {
	return general.ValidateStruct(&r)
}
