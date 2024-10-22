package repository

import (
	"BetterPC_2.0/configs"
	"BetterPC_2.0/pkg/logging"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MongoDbConnection struct {
	Client      *mongo.Client
	Collections map[string]*mongo.Collection
}

func Init(cfg *configs.Config, logger *logging.Logger) (*MongoDbConnection, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().
		ApplyURI(cfg.MongoDB.Url))
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			return
		}
	}()

	mongoDb := &MongoDbConnection{
		Client:      client,
		Collections: make(map[string]*mongo.Collection),
	}

	for _, collectionName := range cfg.MongoDB.UsersCollectionsNames {
		mongoDb.Collections[collectionName], err = InitCollection(client, cfg.MongoDB.UsersDbName, collectionName, logger)
		if err != nil {
			logger.Fatalf("MongoDb initializing error: %s", err.Error())
		}
	}
	for _, collectionName := range cfg.MongoDB.ShopCollectionsNames {
		mongoDb.Collections[collectionName], err = InitCollection(client, cfg.MongoDB.ShopDbName, collectionName, logger)
		if err != nil {
			logger.Fatalf("MongoDb initializing error: %s", err.Error())
		}
	}

	return mongoDb, nil
}

func InitCollection(client *mongo.Client, dbName, collectionName string, logger *logging.Logger) (*mongo.Collection, error) {
	collection := client.Database(dbName).Collection(collectionName)
	var result bson.M
	err := collection.FindOne(context.TODO(), bson.D{}).Decode(&result)
	if errors.Is(err, mongo.ErrNoDocuments) {
		logger.Infof("Collection \"%s\" has no documents", collectionName)
		return collection, nil
	}
	return nil, err
}
