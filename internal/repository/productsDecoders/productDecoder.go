package productsDecoders

import (
	"BetterPC_2.0/pkg/data/models/products"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func DecodeProduct(result *mongo.SingleResult, factory func() products.Product) (*products.Product, error) {
	product := factory() // Create a new instance of the product
	if err := result.Decode(product); err != nil {
		return nil, err
	}
	return &product, nil
}

// DecodeProductsList decodes documents from a cursor into a list of products using the provided factory function.
func DecodeProductsList(cur *mongo.Cursor, factory func() products.Product) (*[]products.Product, error) {
	var productsList []products.Product
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	for cur.Next(ctx) {
		product := factory() // Create a new instance of the product type
		if err := cur.Decode(product); err != nil {
			return nil, err
		}
		productsList = append(productsList, product)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return &productsList, nil
}
