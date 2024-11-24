package errors

import "github.com/pkg/errors"

type ProductError struct {
	err error
}

func (p ProductError) Error() string {
	return p.err.Error()
}

var (
	ErrNoProductsFound           = ProductError{errors.New("no products found")}
	ErrProductModelAlreadyExists = ProductError{errors.New("product model already exists")}
	ErrUnsupportedProductType    = ProductError{errors.New("unsupported product type")}
	ErrProductTypesMismatch      = ProductError{errors.New("product types mismatch")}
	ErrInsufficientStock         = ProductError{errors.New("not enough stock available")}
	ErrProductNotModified        = ProductError{errors.New("product not modified, no new changes for product")}
)
