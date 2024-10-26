package productTypes

import (
	"BetterPC_2.0/pkg/data/models/products"
	"BetterPC_2.0/pkg/data/models/products/details"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateCpuById(cpuId primitive.ObjectID, input products.ProductInput) error {
	var cpuInput details.UpdateCpuInput

	err := cpuInput.ConvertInput(input)
	if err != nil {
		return err
	}

	return nil
}
