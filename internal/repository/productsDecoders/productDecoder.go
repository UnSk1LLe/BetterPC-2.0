package productsDecoders

import (
	"BetterPC_2.0/internal/repository/productsDecoders/Products"
	"BetterPC_2.0/pkg/data/models/products"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
)

func DecodeProduct(productType string, result *mongo.SingleResult) (products.Product, error) {
	switch productType {
	case "cpu":
		return Products.DecodeCpu(result)
	// case "ram": return DecodeRam(result)
	// case "motherboard": return DecodeMotherboard(result)
	// case "ssd": return DecodeSsd(result)
	// case "hdd": return DecodeHdd(result)
	// case "powersupply": return DecodePowerSupply(result)
	// case "cooling": return DecodeCooling(result)
	// case "housing": return DecodeHousing(result)
	// case "gpu": return DecodeGpu(result)

	default:
		return nil, errors.New("unsupported product type")
	}

}

func DecodeProductsList(productType string, cur *mongo.Cursor) (*[]products.Product, error) {
	switch productType {
	case "cpu":
		return Products.DecodeCpuList(cur)
	// case "ram": return DecodeRamList(cur)
	// case "motherboard": return DecodeMotherboardList(cur)
	// case "ssd": return DecodeSsdList(cur)
	// case "hdd": return DecodeHddList(cur)
	// case "powersupply": return DecodePowerSupplyList(cur)
	// case "cooling": return DecodeCoolingList(cur)
	// case "housing": return DecodeHousingList(cur)
	// case "gpu": return DecodeGpuList(cur)

	default:
		return nil, errors.New("unsupported product type")
	}
}
