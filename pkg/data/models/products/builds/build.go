package builds

import (
	"BetterPC_2.0/pkg/data/models/products/details"
	"BetterPC_2.0/pkg/data/models/products/general"
)

type Build struct {
	CPU         general.StandardizedProductData
	Motherboard general.StandardizedProductData
	RAM         general.StandardizedProductData
	GPU         general.StandardizedProductData
	SSD         general.StandardizedProductData
	HDD         general.StandardizedProductData
	Cooling     general.StandardizedProductData
	PowerSupply general.StandardizedProductData
	Housing     general.StandardizedProductData
}

type FullBuild struct {
	CPU         details.Cpu
	Motherboard details.Motherboard
	RAM         details.Ram
	GPU         details.Gpu
	SSD         details.Ssd
	HDD         details.Hdd
	Cooling     details.Cooling
	PowerSupply details.PowerSupply
	Housing     details.Housing
}
