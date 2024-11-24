package requests

import (
	"BetterPC_2.0/pkg/data/helpers/decomposers"
	"BetterPC_2.0/pkg/data/helpers/validators"
	generalRequests "BetterPC_2.0/pkg/data/models/products/general/requests"
)

type UpdateMotherboardRequest struct {
	General     *generalRequests.UpdateGeneralRequest `bson:"general"`
	Socket      string                                `bson:"socket"`
	Chipset     string                                `bson:"chipset"`
	FormFactor  string                                `bson:"form_factor"`
	Ram         *UpdateRamMb                          `bson:"ram"`
	Interfaces  *UpdateInterfaces                     `bson:"interfaces"`
	PciStandard *int                                  `bson:"pci_standard"`
	MbPower     *int                                  `bson:"mb_power"`
	CpuPower    *int                                  `bson:"cpu_power"`
}

type UpdateRamMb struct {
	Slots        *int   `bson:"slots"`
	Type         string `bson:"type"`
	MaxFrequency *int   `bson:"max_frequency"`
	MaxCapacity  *int   `bson:"max_capacity"`
}

type UpdateInterfaces struct {
	Sata3   *int `bson:"SATA3"`
	M2      *int `bson:"M2"`
	PciE1x  *int `bson:"PCI_E_1x"`
	PciE16x *int `bson:"PCI_E_16x"`
}

func (motherboardRequest *UpdateMotherboardRequest) Validate() error {
	return validators.ValidateStruct(motherboardRequest)
}

func (motherboardRequest *UpdateMotherboardRequest) Decompose() (map[string]interface{}, error) {
	return decomposers.DecomposeWithTag(motherboardRequest, "bson")
}

func (motherboardRequest *UpdateMotherboardRequest) SetImage(imageName *string) {
	motherboardRequest.General.Image = imageName
}
