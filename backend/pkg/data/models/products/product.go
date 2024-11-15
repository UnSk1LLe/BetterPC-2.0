package products

import (
	generalResponses "BetterPC_2.0/pkg/data/models/products/general/responses"
)

var ProductTypeFactory = map[string]func() Product{
	"cpu":         func() Product { return &Cpu{} },
	"motherboard": func() Product { return &Motherboard{} },
	"ram":         func() Product { return &Ram{} },
	"gpu":         func() Product { return &Gpu{} },
	"ssd":         func() Product { return &Ssd{} },
	"hdd":         func() Product { return &Hdd{} },
	"cooling":     func() Product { return &Cooling{} },
	"housing":     func() Product { return &Housing{} },
	"powersupply": func() Product { return &PowerSupply{} },
}

type Product interface {
	GetProductModel() string
	ProductFinalPrice() int
	Standardize() generalResponses.StandardizedProductData
}
