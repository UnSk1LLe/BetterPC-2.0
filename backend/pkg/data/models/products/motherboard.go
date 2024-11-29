package products

import (
	"BetterPC_2.0/pkg/data/models/products/general"
	generalResponses "BetterPC_2.0/pkg/data/models/products/general/responses"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
)

type Motherboard struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	General     general.General    `bson:"general" json:"general"`
	Socket      string             `bson:"socket" json:"socket,omitempty"`
	Chipset     string             `bson:"chipset" json:"chipset,omitempty"`
	FormFactor  string             `bson:"form_factor" json:"form_factor,omitempty"`
	Ram         RamMb              `bson:"ram" json:"ram"`
	Interfaces  Interfaces         `bson:"interfaces" json:"interfaces"`
	PciStandard int                `bson:"pci_standard" json:"pci_standard,omitempty"`
	MbPower     int                `bson:"mb_power" json:"mb_power,omitempty"`
	CpuPower    int                `bson:"cpu_power" json:"cpu_power,omitempty"`
}

type RamMb struct {
	Slots        int    `bson:"slots" json:"slots,omitempty"`
	Type         string `bson:"type" json:"type,omitempty"`
	MaxFrequency int    `bson:"max_frequency" json:"max_frequency,omitempty"`
	MaxCapacity  int    `bson:"max_capacity" json:"max_capacity,omitempty"`
}

type Interfaces struct {
	Sata3   int `bson:"sata3" json:"sata3,omitempty"`
	M2      int `bson:"m2" json:"m2,omitempty"`
	PciE1x  int `bson:"pci_e_1x" json:"pci_e_1_x,omitempty"`
	PciE16x int `bson:"pci_e_16x" json:"pci_e_16_x,omitempty"`
}

func (motherboard *Motherboard) GetModel() string {
	return motherboard.General.Model
}

func (motherboard *Motherboard) GetStock() int {
	return motherboard.General.Amount
}

func (motherboard *Motherboard) GetImage() string {
	return motherboard.General.Image
}

func (motherboard *Motherboard) SetImage(imageName string) {
	motherboard.General.Image = imageName
}

func (motherboard *Motherboard) Standardize() generalResponses.StandardizedProductData {
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

func (motherboard *Motherboard) CalculateFinalPrice() int {
	return motherboard.General.GetFinalPrice()
}
