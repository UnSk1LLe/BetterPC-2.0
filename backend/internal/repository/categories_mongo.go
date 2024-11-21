package repository

import (
	"BetterPC_2.0/internal/repository/database/mongoDb"
	"go.mongodb.org/mongo-driver/bson"
)

type CategoriesMongo struct {
	db mongoDb.Database
}

func NewCategoriesMongo(db mongoDb.Database) *CategoriesMongo {
	return &CategoriesMongo{db: db}
}

func (c *CategoriesMongo) GetList(filter bson.M) {

}
