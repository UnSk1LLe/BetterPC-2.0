package products

import (
	"BetterPC_2.0/pkg/data/models/products/general"
	generalResponses "BetterPC_2.0/pkg/data/models/products/general/responses"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
)

type Housing struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	General         general.General    `bson:"general"`
	FormFactor      string             `bson:"form_factor"`
	DriveBays       DriveBays          `bson:"drive_bays"`
	MbFormFactor    string             `bson:"mb_form_factor"`
	PsFormFactor    string             `bson:"ps_form_factor"`
	ExpansionSlots  int                `bson:"expansion_slots"`
	GraphicCardSize int                `bson:"graphic_card_size"`
	CoolerHeight    int                `bson:"cooler_height"`
	Size            []int              `bson:"size"`
	Weight          float64            `bson:"weight"`
}

type DriveBays struct {
	D35 int `bson:"3_5"`
	D25 int `bson:"2_5"`
}

func (housing Housing) GetModel() string {
	return housing.General.Model
}

func (housing Housing) GetStock() int {
	return housing.General.Amount
}

func (housing Housing) GetImage() string {
	return housing.General.Image
}

func (housing Housing) SetImage(imageName string) {
	housing.General.Image = imageName
}

func (housing Housing) Standardize() generalResponses.StandardizedProductData {
	var product generalResponses.StandardizedProductData
	product.ProductHeader.ID = housing.ID.Hex()
	product.ProductHeader.ProductType = "housing"
	product.Name = housing.General.Model
	product.General = housing.General
	product.Description = "Form Factor: " + housing.FormFactor + ", Motherboard Form Factor: " +
		housing.MbFormFactor + ", Expansion Slots: " + strconv.Itoa(housing.ExpansionSlots)
	return product
}

func (housing Housing) CalculateFinalPrice() int {
	return housing.General.GetFinalPrice()
}
