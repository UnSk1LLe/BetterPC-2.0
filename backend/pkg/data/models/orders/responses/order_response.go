package responses

import (
	"BetterPC_2.0/pkg/data/models/orders"
	"BetterPC_2.0/pkg/data/models/products"
)

type ItemsResponse struct {
	Items map[products.ProductType][]orders.Item `json:"items"`
}
