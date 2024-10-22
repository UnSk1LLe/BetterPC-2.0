package details

import (
	"BetterPC_2.0/pkg/data/models/products/general"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Gpu struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	General       general.General    `bson:"general"`
	Architecture  string             `bson:"architecture"`
	Memory        memoryGpu          `bson:"memory"`
	GpuFrequency  int                `bson:"gpu_frequency"`
	ProcessSize   int                `bson:"process_size"`
	MaxResolution string             `bson:"max_resolution"`
	Interfaces    []interfacesGpu    `bson:"interfaces"`
	MaxMonitors   int                `bson:"max_monitors"`
	Cooling       coolingGpu         `bson:"cooling"`
	Tdp           int                `bson:"tdp"`
	TdpR          int                `bson:"tdp_r"`
	PowerSupply   []int              `bson:"power_supply"`
	Slots         float64            `bson:"slots"`
	Size          []int              `bson:"size"`
}

type memoryGpu struct {
	Capacity       int    `bson:"capacity"`
	Type           string `bson:"type"`
	InterfaceWidth int    `bson:"interface_width"`
	Frequency      int    `bson:"frequency"`
}

type interfacesGpu struct {
	Type   string `bson:"type"`
	Number int    `bson:"number"`
}

type coolingGpu struct {
	Type      string `bson:"type"`
	FanNumber int    `bson:"fan_number"`
}

type UpdateGpuInput struct {
	General       *general.General `bson:"general"`
	Architecture  *string          `bson:"architecture"`
	Memory        *memoryGpu       `bson:"memory"`
	GpuFrequency  *int             `bson:"gpu_frequency"`
	ProcessSize   *int             `bson:"process_size"`
	MaxResolution *string          `bson:"max_resolution"`
	Interfaces    *[]interfacesGpu `bson:"interfaces"`
	MaxMonitors   *int             `bson:"max_monitors"`
	Cooling       *coolingGpu      `bson:"cooling"`
	Tdp           *int             `bson:"tdp"`
	TdpR          *int             `bson:"tdp_r"`
	PowerSupply   *[]int           `bson:"power_supply"`
	Slots         *float64         `bson:"slots"`
	Size          *[]int           `bson:"size"`
}

func (g UpdateGpuInput) Validate() error {
	return general.ValidateStruct(&g)
}
