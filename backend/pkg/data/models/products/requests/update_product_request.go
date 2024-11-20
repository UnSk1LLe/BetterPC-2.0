package requests

import "BetterPC_2.0/pkg/data/models/products"

var ProductUpdateRequestFactory = map[products.ProductType]func() ProductUpdateRequest{
	products.ProductTypes.Cpu:         func() ProductUpdateRequest { return &UpdateCpuRequest{} },
	products.ProductTypes.Motherboard: func() ProductUpdateRequest { return &UpdateMotherboardRequest{} },
	products.ProductTypes.Ram:         func() ProductUpdateRequest { return &UpdateRamRequest{} },
	products.ProductTypes.Gpu:         func() ProductUpdateRequest { return &UpdateGpuRequest{} },
	products.ProductTypes.Ssd:         func() ProductUpdateRequest { return &UpdateSsdRequest{} },
	products.ProductTypes.Hdd:         func() ProductUpdateRequest { return &UpdateHddRequest{} },
	products.ProductTypes.PowerSupply: func() ProductUpdateRequest { return &UpdatePowerSupplyRequest{} },
	products.ProductTypes.Cooling:     func() ProductUpdateRequest { return &UpdateCoolingRequest{} },
	products.ProductTypes.Housing:     func() ProductUpdateRequest { return &UpdateHousingRequest{} },
}

type ProductUpdateRequest interface {
	Validate() error
	Decompose() (map[string]interface{}, error)
}
