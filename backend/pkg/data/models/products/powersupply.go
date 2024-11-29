package products

import (
	"BetterPC_2.0/pkg/data/models/products/general"
	generalResponses "BetterPC_2.0/pkg/data/models/products/general/responses"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
)

type PowerSupply struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	General     general.General    `bson:"general" json:"general"`
	FormFactor  string             `bson:"form_factor" json:"form_factor,omitempty"`
	OutputPower int                `bson:"output_power" json:"output_power,omitempty"`
	Connectors  Connectors         `bson:"connectors" json:"connectors"`
	Modules     bool               `bson:"modules" json:"modules,omitempty"`
	MbPower     int                `bson:"mb_power" json:"mb_power,omitempty"`
	CpuPower    CpuPower           `bson:"cpu_power" json:"cpu_power,omitempty"`
}

type Connectors struct {
	Sata  int   `bson:"sata,omitempty" json:"sata,omitempty"`
	Molex int   `bson:"molex,omitempty" json:"molex,omitempty"`
	PciE  []int `bson:"pci_e,omitempty" json:"pci_e,omitempty"`
}

type CpuPower struct {
	Amount int   `bson:"amount" json:"amount,omitempty"`
	Type   []int `bson:"type" json:"type,omitempty"`
}

func (powerSupply *PowerSupply) GetModel() string {
	return powerSupply.General.Model
}

func (powerSupply *PowerSupply) GetStock() int {
	return powerSupply.General.Amount
}

func (powerSupply *PowerSupply) GetImage() string {
	return powerSupply.General.Image
}

func (powerSupply *PowerSupply) SetImage(imageName string) {
	powerSupply.General.Image = imageName
}

func (powerSupply *PowerSupply) Standardize() generalResponses.StandardizedProductData {
	var product generalResponses.StandardizedProductData
	product.ProductHeader.ID = powerSupply.ID.Hex()
	product.ProductHeader.ProductType = "powersupply"
	product.Name = powerSupply.General.Model
	product.General = powerSupply.General
	product.Description = "Form Factor: " + powerSupply.FormFactor + ", Output Power: " +
		strconv.Itoa(powerSupply.OutputPower) + "W, Modular: " + strconv.FormatBool(powerSupply.Modules)
	return product
}

func (powerSupply *PowerSupply) CalculateFinalPrice() int {
	return powerSupply.General.GetFinalPrice()
}
