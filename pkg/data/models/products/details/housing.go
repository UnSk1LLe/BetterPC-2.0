package details

import (
	"BetterPC_2.0/pkg/data/models/products/general"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
)

type Housing struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	General         general.General    `bson:"general"`
	FormFactor      string             `bson:"form_factor"`
	DriveBays       driveBays          `bson:"drive_bays"`
	MbFormFactor    string             `bson:"mb_form_factor"`
	PsFormFactor    string             `bson:"ps_form_factor"`
	ExpansionSlots  int                `bson:"expansion_slots"`
	GraphicCardSize int                `bson:"graphic_card_size"`
	CoolerHeight    int                `bson:"cooler_height"`
	Size            []int              `bson:"size"`
	Weight          float64            `bson:"weight"`
}

type driveBays struct {
	D35 int `bson:"3_5"`
	D25 int `bson:"2_5"`
}

type UpdateHousingInput struct {
	General         *general.General `bson:"general"`
	FormFactor      *string          `bson:"form_factor"`
	DriveBays       *driveBays       `bson:"drive_bays"`
	MbFormFactor    *string          `bson:"mb_form_factor"`
	PsFormFactor    *string          `bson:"ps_form_factor"`
	ExpansionSlots  *int             `bson:"expansion_slots"`
	GraphicCardSize *int             `bson:"graphic_card_size"`
	CoolerHeight    *int             `bson:"cooler_height"`
	Size            *[]int           `bson:"size"`
	Weight          *float64         `bson:"weight"`
}

func (housing Housing) GetProductModel() string {
	return housing.General.Model
}

func (housing Housing) Standardize() general.StandardizedProductData {
	var product general.StandardizedProductData
	product.ProductHeader.ID = housing.ID.Hex()
	product.ProductHeader.ProductType = "housing"
	product.Name = housing.General.Model
	product.General = housing.General
	product.Description = "Form Factor: " + housing.FormFactor + ", Motherboard Form Factor: " +
		housing.MbFormFactor + ", Expansion Slots: " + strconv.Itoa(housing.ExpansionSlots)
	return product
}

func (housing Housing) ProductFinalPrice() int {
	return housing.General.CalculateFinalPrice()
}

func (h UpdateHousingInput) Validate() error {
	return general.ValidateStruct(&h)
}
