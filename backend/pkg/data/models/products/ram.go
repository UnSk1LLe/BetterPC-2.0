package products

import (
	"BetterPC_2.0/pkg/data/models/products/general"
	generalResponses "BetterPC_2.0/pkg/data/models/products/general/responses"
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

func (ram Ram) GetModel() string {
	return ram.General.Model
}

func (ram Ram) GetStock() int {
	return ram.General.Amount
}

func (ram Ram) GetImage() string {
	return ram.General.Image
}

func (ram Ram) SetImage(imageName string) {
	ram.General.Image = imageName
}

func (ram Ram) Standardize() generalResponses.StandardizedProductData {
	var product generalResponses.StandardizedProductData
	product.ProductHeader.ID = ram.ID.Hex()
	product.ProductHeader.ProductType = "ram"
	product.Name = ram.General.Model
	product.General = ram.General
	product.Description = "Capacity: " + strconv.Itoa(ram.Capacity) + "GB, Type: " + ram.Type +
		", Frequency: " + strconv.Itoa(ram.Frequency) + "MHz, CAS Latency: " + ram.CasLatency
	return product
}
func (ram Ram) CalculateFinalPrice() int {
	return ram.General.GetFinalPrice()
}
