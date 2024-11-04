package repository

import (
	"BetterPC_2.0/internal/repository/productTypes"
	"BetterPC_2.0/internal/repository/productsDecoders"
	"BetterPC_2.0/pkg/data/models/products"
	"BetterPC_2.0/pkg/data/models/products/details"
	"BetterPC_2.0/pkg/data/models/products/general"
	"BetterPC_2.0/pkg/database/mongoDb"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type ProductMongo struct {
	db *mongoDb.MongoConnection
}

var productTypeFactory = map[string]func() products.Product{
	"cpu":         func() products.Product { return &details.Cpu{} },
	"motherboard": func() products.Product { return &details.Motherboard{} },
	"ram":         func() products.Product { return &details.Ram{} },
	"gpu":         func() products.Product { return &details.Gpu{} },
	"ssd":         func() products.Product { return &details.Ssd{} },
	"hdd":         func() products.Product { return &details.Hdd{} },
	"cooling":     func() products.Product { return &details.Cooling{} },
	"housing":     func() products.Product { return &details.Housing{} },
	"powersupply": func() products.Product { return &details.PowerSupply{} },
}

func NewProductMongo(mongoConn *mongoDb.MongoConnection) *ProductMongo {
	return &ProductMongo{db: mongoConn}
}

func (p *ProductMongo) Create(product products.Product, productType string) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res := p.db.Collections[productType].FindOne(ctx, bson.M{"general.model": product.GetProductModel()})
	if !errors.Is(res.Err(), mongo.ErrNoDocuments) {
		return primitive.NilObjectID, errors.New(fmt.Sprintf("cpu model <%s> already exists", product.GetProductModel()))
	}

	newProduct, err := p.db.Collections[productType].InsertOne(ctx, product)
	if err != nil {
		return primitive.NilObjectID, errors.New(fmt.Sprintf("error inserting %s: %s", productType, err.Error()))
	}

	return newProduct.InsertedID.(primitive.ObjectID), nil
}

func (p *ProductMongo) GetById(id primitive.ObjectID, productType string) (products.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res := p.db.Collections[productType].FindOne(ctx, bson.M{"_id": id})
	if errors.Is(res.Err(), mongo.ErrNoDocuments) {
		return nil, errors.New(fmt.Sprintf("product of type <%s> with id <%s> not found: %s", productType, id.Hex(), res.Err().Error()))
	} else if res.Err() != nil {
		return nil, res.Err()
	}

	factory, ok := productTypeFactory[productType]
	if !ok {
		return nil, errors.New(fmt.Sprintf("unsupported product type: %s", productType))
	}

	product, err := productsDecoders.DecodeProduct(res, factory)
	if err != nil {
		return nil, err
	}
	return *product, nil
}

func (p *ProductMongo) GetList(filter bson.M, productType string) ([]products.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cur, err := p.db.Collections[productType].Find(ctx, filter)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, errors.New(fmt.Sprintf("no productTypes found using filter: %v", filter))
	} else if err != nil {
		return nil, errors.New(fmt.Sprintf("error finding productTypes: %s", err.Error()))
	}

	factory, ok := productTypeFactory[productType]
	if !ok {
		return nil, errors.New(fmt.Sprintf("unsupported product type: %s", productType))
	}

	productsList, err := productsDecoders.DecodeProductsList(cur, factory)
	if err != nil {
		return nil, err
	}

	return *productsList, nil
}

func (p *ProductMongo) UpdateById(productId primitive.ObjectID, input products.ProductInput, productType string) error {
	switch productType {
	case "cpu":
		return productTypes.UpdateCpuById(productId, input)
	/*case "motherboard":
		return UpdateMotherboardById(productId, input)
	case "ram":
		return UpdateRamById(productId, input)
	case "gpu":
		return UpdateGpuById(productId, input)
	case "ssd":
		return UpdateSsdById(productId, input)
	case "hdd":
		return UpdateHddById(productId, input)
	case "cooling":
		return UpdateCoolingById(productId, input)
	case "powersupply":
		return UpdatePowerSupplyById(productId, input)
	case "housing":
		return UpdateHosuingById(productId, input)*/
	default:
		return errors.New(fmt.Sprintf("unknown product type: %s", productType))
	}
}

func (p *ProductMongo) UpdateGeneralInfoById(productId primitive.ObjectID, input general.UpdateGeneralInput, collectionName string) error {
	/*ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()


	_, err := p.db.Collections[collectionName].UpdateOne(ctx, bson.M{"product_id": productId}, bson.M{"$set": input})*/

	return nil
}

func (p *ProductMongo) DeleteById(productId primitive.ObjectID, productType string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := p.db.Collections[productType].DeleteOne(ctx, bson.M{"productId": productId})
	if err != nil {
		return errors.New(fmt.Sprintf("error deleting product with id <%s> in collection <%s>: %s", productId, productType, err.Error()))
	}

	return nil
}
