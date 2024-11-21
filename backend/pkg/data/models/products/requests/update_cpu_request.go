package requests

import (
	"BetterPC_2.0/pkg/data/helpers/decomposers"
	"BetterPC_2.0/pkg/data/helpers/validators"
	generalRequests "BetterPC_2.0/pkg/data/models/products/general/requests"
)

type UpdateCpuRequest struct {
	General        *generalRequests.UpdateGeneralRequest `bson:"general,omitempty" json:"general"`
	Main           *UpdateMainCpu                        `bson:"main,omitempty" json:"main,omitempty"`
	Cores          *UpdateCoresCpu                       `bson:"cores,omitempty" json:"cores"`
	ClockFrequency *UpdateClockFrequencyCpu              `bson:"clock_frequency,omitempty" json:"clock_frequency,omitempty"`
	Ram            *UpdateRamCpu                         `bson:"ram,omitempty" json:"ram"`
	Tdp            *int                                  `bson:"tdp,omitempty" json:"tdp,omitempty"`
	Graphics       *string                               `bson:"graphics,omitempty" json:"graphics,omitempty"`
	PciE           *int                                  `bson:"pci_e,omitempty" json:"pci_e,omitempty"`
	MaxTemperature *int                                  `bson:"max_temperature,omitempty" json:"max_temperature"`
}

type UpdateMainCpu struct {
	Category   string `bson:"category" json:"category,omitempty"`
	Generation string `bson:"generation" json:"generation,omitempty"`
	Socket     string `bson:"socket" json:"socket,omitempty"`
	Year       *int   `bson:"year" json:"year,omitempty"`
}

type UpdateCoresCpu struct {
	Pcores           *int `bson:"p-cores" json:"p-cores,omitempty"`
	Ecores           *int `bson:"e-cores" json:"e-cores,omitempty"`
	Threads          *int `bson:"threads" json:"threads,omitempty"`
	TechnicalProcess *int `bson:"technical_process" json:"technical_process,omitempty"`
}

type UpdateClockFrequencyCpu struct {
	Pcores         *[]float64 `bson:"p-cores" json:"p-cores,omitempty"`
	Ecores         *[]float64 `bson:"e-cores" json:"e-cores,omitempty"`
	FreeMultiplier *bool      `bson:"free_multiplier" json:"free_multiplier,omitempty"`
}

type UpdateRamCpu struct {
	Channels     *int   `bson:"channels" json:"channels,omitempty"`
	MaxFrequency *[]int `bson:"max_frequency" json:"max_frequency,omitempty"`
	MaxCapacity  *int   `bson:"max_capacity" json:"max_capacity,omitempty"`
}

func (c *UpdateCpuRequest) Validate() error {
	return validators.ValidateStruct(c)
}

func (c *UpdateCpuRequest) Decompose() (map[string]interface{}, error) {
	return decomposers.DecomposeWithTag(c, "bson")
}
