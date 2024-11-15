package requests

import (
	"BetterPC_2.0/pkg/data/helpers/decomposers"
	"BetterPC_2.0/pkg/data/helpers/validators"
	"BetterPC_2.0/pkg/data/models/products"
	"BetterPC_2.0/pkg/data/models/products/general"
)

type UpdateMotherboardRequest struct {
	General     *general.General     `bson:"general"`
	Socket      *string              `bson:"socket"`
	Chipset     *string              `bson:"chipset"`
	FormFactor  *string              `bson:"form_factor"`
	Ram         *products.RamMb      `bson:"ram"`
	Interfaces  *products.Interfaces `bson:"interfaces"`
	PciStandard *int                 `bson:"pci_standard"`
	MbPower     *int                 `bson:"mb_power"`
	CpuPower    *int                 `bson:"cpu_power"`
}

func (m *UpdateMotherboardRequest) Validate() error {
	return validators.ValidateStruct(m)
}

func (m *UpdateMotherboardRequest) Decompose() (map[string]interface{}, error) {
	return decomposers.DecomposeWithTag(m, "bson")
}
