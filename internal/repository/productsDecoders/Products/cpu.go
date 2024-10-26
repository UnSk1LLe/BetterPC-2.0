package Products

import (
	"BetterPC_2.0/pkg/data/models/products"
	"BetterPC_2.0/pkg/data/models/products/details"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func DecodeCpu(result *mongo.SingleResult) (*details.Cpu, error) {
	var cpu details.Cpu
	err := result.Decode(&cpu)
	if err != nil {
		return nil, err
	}
	return &cpu, nil
}

func DecodeCpuList(cur *mongo.Cursor) (*[]products.Product, error) {
	var cpuList []products.Product
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	for cur.Next(ctx) {
		var cpu details.Cpu
		if err := cur.Decode(&cpu); err != nil {
			return nil, err
		}
		cpuList = append(cpuList, cpu)
	}
	return &cpuList, nil
}
