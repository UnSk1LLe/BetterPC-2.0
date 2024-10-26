package details

import (
	"BetterPC_2.0/pkg/data/models/products/general"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (r Ram) GetProductModel() string {
	return r.General.Model
}

// TODO add standardizers for each product type
func (r Ram) Standardize() general.StandardizedProductData {
	return general.StandardizedProductData{}
}

func (r Ram) ProductFinalPrice() int {
	return r.General.CalculateFinalPrice()
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
