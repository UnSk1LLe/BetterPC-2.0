package requests

import (
	"BetterPC_2.0/pkg/data/models/orders"
	orderErrors "BetterPC_2.0/pkg/data/models/orders/errors"
	"BetterPC_2.0/pkg/data/models/products"
	productErrors "BetterPC_2.0/pkg/data/models/products/errors"
	"github.com/pkg/errors"
)

const (
	MaxSelectedAmount = 20
	MinSelectedAmount = 1
)

type CreateOrderRequest struct {
	ProductTypeItemHeaders map[products.ProductType][]orders.ItemHeader `json:"item_headers" binding:"dive"`
}

func (cor *CreateOrderRequest) Validate() error {

	for productType, items := range cor.ProductTypeItemHeaders {
		_, err := products.ProductTypeFromString(productType.String())
		if err != nil {
			return errors.Wrapf(orderErrors.ErrInvalidInput, "%s: %s", productErrors.ErrUnsupportedProductType.Error(), productType)
		}

		for _, item := range items {
			if item.SelectedAmount < MinSelectedAmount || item.SelectedAmount > MaxSelectedAmount {
				return errors.Wrapf(orderErrors.ErrInvalidInput, "selected amount must be between %d and %d", MinSelectedAmount, MaxSelectedAmount)
			}

			if item.ID == "" {
				return errors.Wrapf(orderErrors.ErrInvalidInput, "product type %s has missing ID", productType)
			}
		}
	}

	return nil
}
