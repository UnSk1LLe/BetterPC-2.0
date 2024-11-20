package products

import (
	"BetterPC_2.0/pkg/data/models/products/general"
	generalResponses "BetterPC_2.0/pkg/data/models/products/general/responses"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
	"strconv"
)

type PowerSupply struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	General     general.General    `bson:"general"`
	FormFactor  string             `bson:"form_factor"`
	OutputPower int                `bson:"output_power"`
	Connectors  Connectors         `bson:"connectors"`
	Modules     bool               `bson:"modules"`
	MbPower     int                `bson:"mb_power"`
	CpuPower    bsoncore.Array     `bson:"cpu_power"`
}

type Connectors struct {
	Sata  int   `bson:"SATA"`
	Molex int   `bson:"MOLEX"`
	PciE  []int `bson:"PCI_E"`
}

func (powerSupply PowerSupply) GetModel() string {
	return powerSupply.General.Model
}

func (powerSupply PowerSupply) GetStock() int {
	return powerSupply.General.Amount
}

func (powerSupply PowerSupply) Standardize() generalResponses.StandardizedProductData {
	var product generalResponses.StandardizedProductData
	product.ProductHeader.ID = powerSupply.ID.Hex()
	product.ProductHeader.ProductType = "powersupply"
	product.Name = powerSupply.General.Model
	product.General = powerSupply.General
	product.Description = "Form Factor: " + powerSupply.FormFactor + ", Output Power: " +
		strconv.Itoa(powerSupply.OutputPower) + "W, Modular: " + strconv.FormatBool(powerSupply.Modules)
	return product
}

func (powerSupply PowerSupply) CalculateFinalPrice() int {
	return powerSupply.General.GetFinalPrice()
}
