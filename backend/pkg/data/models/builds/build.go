package builds

import (
	"BetterPC_2.0/pkg/data/models/products"
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
	CPU         products.Cpu
	Motherboard products.Motherboard
	RAM         products.Ram
	GPU         products.Gpu
	SSD         products.Ssd
	HDD         products.Hdd
	Cooling     products.Cooling
	PowerSupply products.PowerSupply
	Housing     products.Housing
}
