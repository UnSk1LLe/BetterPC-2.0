package general

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reflect"
)

type Product interface {
	General
	ProductFinalPrice() int
	Standardize()
}

type ProductHeader struct {
	ID          primitive.ObjectID
	ProductType string
}

type StandardizedProductData struct {
	ProductHeader ProductHeader
	General       General
	Name          string
	Description   string
}

type General struct {
	Manufacturer string `bson:"manufacturer" json:"manufacturer"`
	Model        string `bson:"model" json:"model"`
	Price        int    `bson:"price" json:"price"`
	Discount     int    `bson:"discount" json:"discount"`
	Amount       int    `bson:"amount" json:"amount"`
}

type UpdateGeneralInput struct {
	Manufacturer *string `bson:"manufacturer" json:"manufacturer"`
	Model        *string `bson:"model" json:"model"`
	Price        *int    `bson:"price" json:"price"`
	Discount     *int    `bson:"discount" json:"discount"`
	Amount       *int    `bson:"amount" json:"amount"`
}

func (g *General) ProductFinalPrice() int {
	return g.Price - (g.Price * g.Discount / 100)
}

func ValidateStruct(input interface{}) error {
	v := reflect.ValueOf(input)

	// Ensure that we are working with a pointer to a struct
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return errors.New("input must be a pointer to a struct")
	}

	v = v.Elem()

	// Iterate over the struct fields
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)

		// If any field is not nil, return nil (validation passed)
		if !field.IsNil() {
			return nil
		}
	}

	// If all fields are nil, return an error
	return errors.New("all item fields are empty")
}
