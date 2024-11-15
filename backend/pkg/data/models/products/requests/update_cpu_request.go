package requests

import (
	"BetterPC_2.0/pkg/data/helpers/decomposers"
	"BetterPC_2.0/pkg/data/helpers/validators"
	"BetterPC_2.0/pkg/data/models/products"
	"BetterPC_2.0/pkg/data/models/products/general"
)

type UpdateCpuRequest struct {
	General        *general.General            `bson:"general,omitempty" json:"general"`
	Main           *products.MainCpu           `bson:"main,omitempty" json:"main,omitempty"`
	Cores          *products.CoresCpu          `bson:"cores,omitempty" json:"cores"`
	ClockFrequency *products.ClockFrequencyCpu `bson:"clock_frequency,omitempty" json:"clock_frequency,omitempty"`
	Ram            *products.RamCpu            `bson:"ram,omitempty" json:"ram"`
	Tdp            *int                        `bson:"tdp,omitempty" json:"tdp,omitempty"`
	Graphics       *string                     `bson:"graphics,omitempty" json:"graphics,omitempty"`
	PciE           *int                        `bson:"pci_e,omitempty" json:"pci_e,omitempty"`
	MaxTemperature *int                        `bson:"max_temperature,omitempty" json:"max_temperature"`
}

func (c *UpdateCpuRequest) Validate() error {
	return validators.ValidateStruct(c)
}

func (c *UpdateCpuRequest) Decompose() (map[string]interface{}, error) {
	return decomposers.DecomposeWithTag(c, "bson")
}
