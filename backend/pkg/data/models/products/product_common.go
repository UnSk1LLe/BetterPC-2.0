package products

import (
	productErrors "BetterPC_2.0/pkg/data/models/products/errors"
	generalResponses "BetterPC_2.0/pkg/data/models/products/general/responses"
	"github.com/pkg/errors"
	"strings"
)

const MaxProductTypeLength = 20

type ProductType string

var ProductTypes = struct {
	Cpu         ProductType
	Motherboard ProductType
	Ram         ProductType
	Gpu         ProductType
	Ssd         ProductType
	Hdd         ProductType
	Cooling     ProductType
	PowerSupply ProductType
	Housing     ProductType
}{
	Cpu:         "cpu",
	Motherboard: "motherboard",
	Ram:         "ram",
	Gpu:         "gpu",
	Ssd:         "ssd",
	Hdd:         "hdd",
	Cooling:     "cooling",
	PowerSupply: "powersupply",
	Housing:     "housing",
}

func (pt ProductType) String() string {
	return string(pt)
}

func ProductTypeFromString(input string) (ProductType, error) {
	if len(input) > MaxProductTypeLength {
		return "", errors.Wrapf(productErrors.ErrUnsupportedProductType, "product type must be shorter than %d characters", MaxProductTypeLength)
	}

	normalizedInput := strings.ToLower(strings.TrimSpace(input))
	if len(normalizedInput) == 0 {
		return "", errors.Wrap(productErrors.ErrUnsupportedProductType, "product type cannot not be empty")
	}

	switch ProductType(normalizedInput) {
	case ProductTypes.Cpu:
		return ProductTypes.Cpu, nil
	case ProductTypes.Motherboard:
		return ProductTypes.Motherboard, nil
	case ProductTypes.Ram:
		return ProductTypes.Ram, nil
	case ProductTypes.Gpu:
		return ProductTypes.Gpu, nil
	case ProductTypes.Ssd:
		return ProductTypes.Ssd, nil
	case ProductTypes.Hdd:
		return ProductTypes.Hdd, nil
	case ProductTypes.Cooling:
		return ProductTypes.Cooling, nil
	case ProductTypes.PowerSupply:
		return ProductTypes.PowerSupply, nil
	case ProductTypes.Housing:
		return ProductTypes.Housing, nil
	}

	/*mapping := map[string]ProductType{
		"cpu":          ProductTypes.Cpu,
		"motherboard":  ProductTypes.Motherboard,
		"ram":          ProductTypes.Ram,
		"gpu":          ProductTypes.Gpu,
		"ssd":          ProductTypes.Ssd,
		"hdd":          ProductTypes.Hdd,
		"cooling":      ProductTypes.Cooling,
		"powersupply":  ProductTypes.PowerSupply,
		"power supply": ProductTypes.PowerSupply,
		"housing":      ProductTypes.Housing,
	}

	if productType, exists := mapping[normalizedInput]; exists {
		return productType, nil
	}*/

	return "", productErrors.ErrUnsupportedProductType
}

var ProductTypeFactory = map[ProductType]func() Product{
	ProductTypes.Cpu:         func() Product { return &Cpu{} },
	ProductTypes.Motherboard: func() Product { return &Motherboard{} },
	ProductTypes.Ram:         func() Product { return &Ram{} },
	ProductTypes.Gpu:         func() Product { return &Gpu{} },
	ProductTypes.Ssd:         func() Product { return &Ssd{} },
	ProductTypes.Hdd:         func() Product { return &Hdd{} },
	ProductTypes.Cooling:     func() Product { return &Cooling{} },
	ProductTypes.Housing:     func() Product { return &Housing{} },
	ProductTypes.PowerSupply: func() Product { return &PowerSupply{} },
}

type Product interface {
	GetModel() string
	GetStock() int
	GetImage() string
	SetImage(imageName string)
	CalculateFinalPrice() int
	Standardize() generalResponses.StandardizedProductData
}
