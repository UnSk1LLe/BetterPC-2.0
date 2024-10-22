package details

import (
	"BetterPC_2.0/pkg/data/models/products/general"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Cpu struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	General        general.General    `bson:"general"`
	Main           MainCpu            `bson:"main"`
	Cores          CoresCpu           `bson:"cores"`
	ClockFrequency ClockFrequencyCpu  `bson:"clock_frequency" json:"clock_frequency,omitempty"`
	Ram            RamCpu             `bson:"ram"`
	Tdp            int                `bson:"tdp" json:"tdp,omitempty"`
	Graphics       string             `bson:"graphics" json:"graphics,omitempty"`
	PciE           int                `bson:"pci-e" json:"pci_e,omitempty"`
	MaxTemperature int                `bson:"max_temperature"`
}

type MainCpu struct {
	Category   string `bson:"category" json:"category,omitempty"`
	Generation string `bson:"generation" json:"generation,omitempty"`
	Socket     string `bson:"socket" json:"socket,omitempty"`
	Year       int    `bson:"year" json:"year,omitempty"`
}

type CoresCpu struct {
	Pcores           int `bson:"p-cores" json:"p-cores,omitempty"`
	Ecores           int `bson:"e-cores" json:"e-cores,omitempty"`
	Threads          int `bson:"threads" json:"threads,omitempty"`
	TechnicalProcess int `bson:"technical_process" json:"technical_process,omitempty"`
}

type ClockFrequencyCpu struct {
	Pcores         []float64 `bson:"p-cores" json:"p-cores,omitempty"`
	Ecores         []float64 `bson:"e-cores" json:"e-cores,omitempty"`
	FreeMultiplier bool      `bson:"free_multiplier" json:"free_multiplier,omitempty"`
}

type RamCpu struct {
	Channels     int   `bson:"channels"`
	MaxFrequency []int `bson:"max_frequency"`
	MaxCapacity  int   `bson:"max_capacity"`
}

type UpdateCpuInput struct {
	General        *general.General   `bson:"general"`
	Main           *MainCpu           `bson:"main"`
	Cores          *CoresCpu          `bson:"cores"`
	ClockFrequency *ClockFrequencyCpu `bson:"clock_frequency" json:"clock_frequency,omitempty"`
	Ram            *RamCpu            `bson:"ram"`
	Tdp            *int               `bson:"tdp" json:"tdp,omitempty"`
	Graphics       *string            `bson:"graphics" json:"graphics,omitempty"`
	PciE           *int               `bson:"pci-e" json:"pci_e,omitempty"`
	MaxTemperature *int               `bson:"max_temperature"`
}

func (c UpdateCpuInput) Validate() error {
	return general.ValidateStruct(&c)
}
