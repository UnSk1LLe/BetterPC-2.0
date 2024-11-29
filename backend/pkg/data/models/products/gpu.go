package products

import (
	"BetterPC_2.0/pkg/data/models/products/general"
	generalResponses "BetterPC_2.0/pkg/data/models/products/general/responses"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
)

type Gpu struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	General       general.General    `bson:"general"`
	Architecture  string             `bson:"architecture"`
	Memory        MemoryGpu          `bson:"memory"`
	GpuFrequency  int                `bson:"gpu_frequency"`
	ProcessSize   int                `bson:"process_size"`
	MaxResolution string             `bson:"max_resolution"`
	Interfaces    []InterfacesGpu    `bson:"interfaces"`
	MaxMonitors   int                `bson:"max_monitors"`
	Cooling       CoolingGpu         `bson:"cooling"`
	Tdp           int                `bson:"tdp"`
	TdpR          int                `bson:"tdp_r"`
	PowerSupply   []int              `bson:"power_supply"`
	Slots         float64            `bson:"slots"`
	Size          []int              `bson:"size"`
}

type MemoryGpu struct {
	Capacity       int    `bson:"capacity"`
	Type           string `bson:"type"`
	InterfaceWidth int    `bson:"interface_width"`
	Frequency      int    `bson:"frequency"`
}

type InterfacesGpu struct {
	Type   string `bson:"type"`
	Number int    `bson:"number"`
}

type CoolingGpu struct {
	Type      string `bson:"type"`
	FanNumber int    `bson:"fan_number"`
}

func (gpu *Gpu) GetModel() string {
	return gpu.General.Model
}

func (gpu *Gpu) GetStock() int {
	return gpu.General.Amount
}

func (gpu *Gpu) GetImage() string {
	return gpu.General.Image
}

func (gpu *Gpu) SetImage(imageName string) {
	gpu.General.Image = imageName
}

func (gpu *Gpu) Standardize() generalResponses.StandardizedProductData {
	var product generalResponses.StandardizedProductData
	product.ProductHeader.ID = gpu.ID.Hex()
	product.ProductHeader.ProductType = "gpu"
	product.Name = gpu.General.Model
	product.General = gpu.General
	product.Description = "Architecture: " + gpu.Architecture + ", Memory: " + strconv.Itoa(gpu.Memory.Capacity) +
		"GB " + gpu.Memory.Type + ", Frequency: " + strconv.Itoa(gpu.GpuFrequency) + "MHz, " +
		"Max Resolution: " + gpu.MaxResolution
	return product
}

func (gpu *Gpu) CalculateFinalPrice() int {
	return gpu.General.GetFinalPrice()
}
