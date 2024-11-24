package requests

import (
	"BetterPC_2.0/pkg/data/helpers/decomposers"
	"BetterPC_2.0/pkg/data/helpers/validators"
	generalRequests "BetterPC_2.0/pkg/data/models/products/general/requests"
)

type UpdateGpuRequest struct {
	General       *generalRequests.UpdateGeneralRequest `bson:"general"`
	Architecture  *string                               `bson:"architecture"`
	Memory        *UpdateMemoryGpu                      `bson:"memory"`
	GpuFrequency  *int                                  `bson:"gpu_frequency"`
	ProcessSize   *int                                  `bson:"process_size"`
	MaxResolution *string                               `bson:"max_resolution"`
	Interfaces    *[]UpdateInterfacesGpu                `bson:"Interfaces"`
	MaxMonitors   *int                                  `bson:"max_monitors"`
	Cooling       *UpdateCoolingGpu                     `bson:"cooling"`
	Tdp           *int                                  `bson:"tdp"`
	TdpR          *int                                  `bson:"tdp_r"`
	PowerSupply   *[]int                                `bson:"power_supply"`
	Slots         *float64                              `bson:"slots"`
	Size          *[]int                                `bson:"size"`
}

type UpdateMemoryGpu struct {
	Capacity       *int   `bson:"capacity"`
	Type           string `bson:"type"`
	InterfaceWidth *int   `bson:"interface_width"`
	Frequency      *int   `bson:"frequency"`
}

type UpdateInterfacesGpu struct {
	Type   string `bson:"type"`
	Number *int   `bson:"number"`
}

type UpdateCoolingGpu struct {
	Type      string `bson:"type"`
	FanNumber *int   `bson:"fan_number"`
}

func (gpuRequest *UpdateGpuRequest) Validate() error {
	return validators.ValidateStruct(gpuRequest)
}

func (gpuRequest *UpdateGpuRequest) Decompose() (map[string]interface{}, error) {
	return decomposers.DecomposeWithTag(gpuRequest, "bson")
}

func (gpuRequest *UpdateGpuRequest) SetImage(imageName *string) {
	gpuRequest.General.Image = imageName
}
