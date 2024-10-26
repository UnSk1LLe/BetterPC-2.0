package productsDecoders

import (
	"BetterPC_2.0/internal/repository/productsDecoders/Products"
	"BetterPC_2.0/pkg/data/models/products"
	"BetterPC_2.0/pkg/data/models/products/details"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
)

func DecodeProduct(productType string, result *mongo.SingleResult) (products.Product, error) {
	switch productType {
	case "cpu":
		return Products.DecodeCpu(result)
	case "ram":
		return details.Ram{}, nil
	// добавьте другие типы продуктов здесь
	default:
		return nil, errors.New("unsupported product type")
	}

}

func DecodeProductsList(productType string, cur *mongo.Cursor) (*[]products.Product, error) {
	switch productType {
	case "cpu":
		return Products.DecodeCpuList(cur)
	// case "ram": return DecodeRamList(cur)
	// добавьте другие типы продуктов здесь
	default:
		return nil, errors.New("unsupported product type")
	}
}
