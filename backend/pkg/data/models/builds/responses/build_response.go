package responses

import generalResponses "BetterPC_2.0/pkg/data/models/products/general/responses"

type BuildResponse struct {
	CPU         *generalResponses.StandardizedProductData
	Motherboard *generalResponses.StandardizedProductData
	RAM         *generalResponses.StandardizedProductData
	GPU         *generalResponses.StandardizedProductData
	SSD         []*generalResponses.StandardizedProductData
	HDD         []*generalResponses.StandardizedProductData
	Cooling     *generalResponses.StandardizedProductData
	PowerSupply *generalResponses.StandardizedProductData
	Housing     *generalResponses.StandardizedProductData
}
