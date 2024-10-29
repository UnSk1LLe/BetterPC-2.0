package repository

import (
	"BetterPC_2.0/pkg/database/mongoDb"
	"go.mongodb.org/mongo-driver/bson"
)

type CategoriesMongo struct {
	db *mongoDb.MongoConnection
}

func NewCategoriesMongo(db *mongoDb.MongoConnection) *CategoriesMongo {
	return &CategoriesMongo{db: db}
}

func (c *CategoriesMongo) GetList(filter bson.M) {

}
