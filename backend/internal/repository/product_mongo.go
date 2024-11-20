package repository

import (
	"BetterPC_2.0/internal/repository/database/mongoDb"
	"BetterPC_2.0/internal/repository/helpers/productsDecoders"
	"BetterPC_2.0/internal/repository/helpers/typeValidators"
	"BetterPC_2.0/pkg/data/models/products"
	productErrors "BetterPC_2.0/pkg/data/models/products/errors"
	generalProductRequests "BetterPC_2.0/pkg/data/models/products/general/requests"
	productRequests "BetterPC_2.0/pkg/data/models/products/requests"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
	"time"
)

type productStructsTypes struct {
	productType       reflect.Type
	updateProductType reflect.Type
}

var ProductTypesMap = map[products.ProductType]productStructsTypes{
	products.ProductTypes.Cpu: {
		productType:       reflect.TypeOf(products.Cpu{}),
		updateProductType: reflect.TypeOf(productRequests.UpdateCpuRequest{}),
	},
	products.ProductTypes.Ram: {
		productType:       reflect.TypeOf(products.Ram{}),
		updateProductType: reflect.TypeOf(productRequests.UpdateRamRequest{}),
	},
	products.ProductTypes.Motherboard: {
		productType:       reflect.TypeOf(products.Motherboard{}),
		updateProductType: reflect.TypeOf(productRequests.UpdateMotherboardRequest{}),
	},
	products.ProductTypes.Gpu: {
		productType:       reflect.TypeOf(products.Gpu{}),
		updateProductType: reflect.TypeOf(productRequests.UpdateGpuRequest{}),
	},
	products.ProductTypes.Ssd: {
		productType:       reflect.TypeOf(products.Ssd{}),
		updateProductType: reflect.TypeOf(productRequests.UpdateSsdRequest{}),
	},
	products.ProductTypes.Hdd: {
		productType:       reflect.TypeOf(products.Hdd{}),
		updateProductType: reflect.TypeOf(productRequests.UpdateHddRequest{}),
	},
	products.ProductTypes.PowerSupply: {
		productType:       reflect.TypeOf(products.PowerSupply{}),
		updateProductType: reflect.TypeOf(productRequests.UpdatePowerSupplyRequest{}),
	},
	products.ProductTypes.Cooling: {
		productType:       reflect.TypeOf(products.Cooling{}),
		updateProductType: reflect.TypeOf(productRequests.UpdateCoolingRequest{}),
	},
	products.ProductTypes.Housing: {
		productType:       reflect.TypeOf(products.Housing{}),
		updateProductType: reflect.TypeOf(productRequests.UpdateHousingRequest{}),
	},
}

type ProductMongo struct {
	db *mongoDb.MongoConnection
}

func NewProductMongo(mongoConn *mongoDb.MongoConnection) *ProductMongo {
	return &ProductMongo{db: mongoConn}
}

func (p *ProductMongo) Create(product products.Product, productType products.ProductType) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res := p.db.Collections[productType.String()].FindOne(ctx, bson.M{"general.model": product.GetModel()})
	if !errors.Is(res.Err(), mongo.ErrNoDocuments) {
		return primitive.NilObjectID, productErrors.ErrProductModelAlreadyExists
	}

	newProduct, err := p.db.Collections[productType.String()].InsertOne(ctx, product)
	if err != nil {
		return primitive.NilObjectID, errors.New(fmt.Sprintf("error inserting %s: %s", productType, err.Error()))
	}

	return newProduct.InsertedID.(primitive.ObjectID), nil
}

func (p *ProductMongo) GetById(id primitive.ObjectID, productType products.ProductType) (products.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res := p.db.Collections[productType.String()].FindOne(ctx, bson.M{"_id": id})
	if errors.Is(res.Err(), mongo.ErrNoDocuments) {
		return nil, productErrors.ErrNoProductsFound
	} else if res.Err() != nil {
		return nil, res.Err()
	}

	factory, ok := products.ProductTypeFactory[productType]
	if !ok {
		return nil, productErrors.ErrUnsupportedProductType
	}
	product, err := productsDecoders.DecodeProduct(res, factory)
	if err != nil {
		return nil, err
	}
	return *product, nil
}

func (p *ProductMongo) GetList(filter bson.M, productType products.ProductType) ([]products.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cur, err := p.db.Collections[productType.String()].Find(ctx, filter)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, productErrors.ErrNoProductsFound
	} else if err != nil {
		return nil, errors.New(fmt.Sprintf("error finding productTypes: %s", err.Error()))
	}

	factory, ok := products.ProductTypeFactory[productType]
	if !ok {
		return nil, productErrors.ErrUnsupportedProductType
	}

	productsList, err := productsDecoders.DecodeProductsList(cur, factory)
	if err != nil {
		return nil, err
	}

	return *productsList, nil
}

func (p *ProductMongo) UpdateById(productId primitive.ObjectID, input productRequests.ProductUpdateRequest, productType products.ProductType) error {
	err := typeValidators.ValidateType(input, ProductTypesMap[productType].updateProductType)
	if err != nil {
		return productErrors.ErrProductTypesMismatch
	}

	collection, ok := p.db.Collections[productType.String()]
	if !ok {
		return productErrors.ErrProductTypesMismatch
	}
	if collection == nil {
		return errors.New(fmt.Sprintf("mongo collection reference error: collection %s is nil", productType))
	}

	fieldsValues, err := input.Decompose()
	if err != nil {
		return err
	}

	update := bson.M{"$set": fieldsValues}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	updRes, err := collection.UpdateByID(ctx, productId, update)
	if err != nil {
		return err
	}
	if updRes.MatchedCount == 0 {
		return productErrors.ErrNoProductsFound
	}
	return nil
}

func (p *ProductMongo) UpdateGeneralInfoById(productId primitive.ObjectID, input generalProductRequests.UpdateGeneralRequest, productType products.ProductType) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	generalUpdateData, err := input.Decompose()
	if err != nil {
		return err
	}

	update := bson.M{"$set": generalUpdateData}

	res, err := p.db.Collections[productType.String()].UpdateByID(ctx, productId, update)
	if err != nil {
		return err
	}
	if res.MatchedCount == 0 {
		return productErrors.ErrNoProductsFound
	}
	if res.ModifiedCount == 0 {
		return productErrors.ErrProductNotModified
	}

	return nil
}

func (p *ProductMongo) DeleteById(productId primitive.ObjectID, productType products.ProductType) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	delRes, err := p.db.Collections[productType.String()].DeleteOne(ctx, bson.M{"productId": productId})
	if err != nil {
		return errors.New(fmt.Sprintf("error deleting product with id <%s> in collection <%s>: %s", productId, productType, err.Error()))
	} else if delRes.DeletedCount == 0 {
		return productErrors.ErrNoProductsFound
	}

	return nil
}
