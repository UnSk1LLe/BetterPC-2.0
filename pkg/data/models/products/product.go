package products

import "BetterPC_2.0/pkg/data/models/products/general"

type Product interface {
	GetProductModel() string
	ProductFinalPrice() int
	Standardize() general.StandardizedProductData
}

type ProductInput interface {
	Validate() error
	ConvertInput(input ProductInput) error
}
