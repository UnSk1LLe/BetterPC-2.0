package responses

import "BetterPC_2.0/pkg/data/models/products/general"

type ProductHeader struct {
	ID          string
	ProductType string
}

type StandardizedProductData struct {
	ProductHeader ProductHeader
	General       general.General
	Name          string
	Description   string
}
