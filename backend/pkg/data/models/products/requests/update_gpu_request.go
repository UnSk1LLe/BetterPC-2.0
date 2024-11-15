package requests

import (
	"BetterPC_2.0/pkg/data/helpers/decomposers"
	"BetterPC_2.0/pkg/data/helpers/validators"
	"BetterPC_2.0/pkg/data/models/products"
	"BetterPC_2.0/pkg/data/models/products/general"
)

type UpdateGpuRequest struct {
	General       *general.General          `bson:"general"`
	Architecture  *string                   `bson:"architecture"`
	Memory        *products.MemoryGpu       `bson:"memory"`
	GpuFrequency  *int                      `bson:"gpu_frequency"`
	ProcessSize   *int                      `bson:"process_size"`
	MaxResolution *string                   `bson:"max_resolution"`
	Interfaces    *[]products.InterfacesGpu `bson:"Interfaces"`
	MaxMonitors   *int                      `bson:"max_monitors"`
	Cooling       *products.CoolingGpu      `bson:"cooling"`
	Tdp           *int                      `bson:"tdp"`
	TdpR          *int                      `bson:"tdp_r"`
	PowerSupply   *[]int                    `bson:"power_supply"`
	Slots         *float64                  `bson:"slots"`
	Size          *[]int                    `bson:"size"`
}

func (m *UpdateGpuRequest) Validate() error {
	return validators.ValidateStruct(m)
}

func (m *UpdateGpuRequest) Decompose() (map[string]interface{}, error) {
	return decomposers.DecomposeWithTag(m, "bson")
}
