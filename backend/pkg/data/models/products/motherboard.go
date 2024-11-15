package products

import (
	"BetterPC_2.0/pkg/data/models/products/general"
	generalResponses "BetterPC_2.0/pkg/data/models/products/general/responses"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
)

type Motherboard struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	General     general.General    `bson:"general"`
	Socket      string             `bson:"socket"`
	Chipset     string             `bson:"chipset"`
	FormFactor  string             `bson:"form_factor"`
	Ram         RamMb              `bson:"ram"`
	Interfaces  Interfaces         `bson:"interfaces"`
	PciStandard int                `bson:"pci_standard"`
	MbPower     int                `bson:"mb_power"`
	CpuPower    int                `bson:"cpu_power"`
}

type RamMb struct {
	Slots        int    `bson:"slots"`
	Type         string `bson:"type"`
	MaxFrequency int    `bson:"max_frequency"`
	MaxCapacity  int    `bson:"max_capacity"`
}

type Interfaces struct {
	Sata3   int `bson:"SATA3"`
	M2      int `bson:"M2"`
	PciE1x  int `bson:"PCI_E_1x"`
	PciE16x int `bson:"PCI_E_16x"`
}

func (motherboard Motherboard) GetProductModel() string {
	return motherboard.General.Model
}

func (motherboard Motherboard) Standardize() generalResponses.StandardizedProductData {
	var product generalResponses.StandardizedProductData
	product.ProductHeader.ID = motherboard.ID.Hex()
	product.ProductHeader.ProductType = "motherboard"
	product.Name = motherboard.General.Model
	product.General = motherboard.General
	product.Description = "Socket: " + motherboard.Socket + ", Chipset: " + motherboard.Chipset +
		", Form Factor: " + motherboard.FormFactor + ", RAM: " + motherboard.Ram.Type + " " +
		strconv.Itoa(motherboard.Ram.MaxCapacity) + "GB"
	return product
}

func (motherboard Motherboard) ProductFinalPrice() int {
	return motherboard.General.CalculateFinalPrice()
}
