package mongoDb

import (
	"BetterPC_2.0/configs"
	"BetterPC_2.0/pkg/logging"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var MongoConn MongoConnection

type MongoConnection struct {
	Client      *mongo.Client
	Collections map[string]*mongo.Collection
}

func Init(cfg *configs.Config, logger *logging.Logger) error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().
		ApplyURI(cfg.MongoDB.Url))
	if err != nil {
		return err
	}

	MongoConn = MongoConnection{
		Client:      client,
		Collections: make(map[string]*mongo.Collection),
	}

	for _, collectionName := range cfg.MongoDB.UsersCollectionsNames {
		MongoConn.Collections[collectionName], err = InitCollection(client, cfg.MongoDB.UsersDbName, collectionName, logger)
		if err != nil {
			logger.Fatalf("MongoDb initializing error: %s", err.Error())
		}
	}
	for _, collectionName := range cfg.MongoDB.ShopCollectionsNames {
		MongoConn.Collections[collectionName], err = InitCollection(client, cfg.MongoDB.ShopDbName, collectionName, logger)
		if err != nil {
			logger.Fatalf("MongoDb initializing error: %s", err.Error())
		}
	}

	fmt.Println(MongoConn.Collections)

	return nil
}

func InitCollection(client *mongo.Client, dbName, collectionName string, logger *logging.Logger) (*mongo.Collection, error) {
	collection := client.Database(dbName).Collection(collectionName)
	if collection == nil {
		return nil, errors.New(fmt.Sprintf("nil pointer for collection <%s> in db <%s>. collection does not exist", collectionName, dbName))
	}
	logger.Infof("Collection \"%s\" found", collectionName)

	res := collection.FindOne(context.TODO(), bson.D{})
	if errors.Is(res.Err(), mongo.ErrNoDocuments) {
		logger.Warnf("Collection \"%s\" has no documents", collectionName)
	}

	return collection, nil
}

func GetConnection() (*MongoConnection, error) {
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	err := MongoConn.Client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	return &MongoConn, nil
}

func CloseConnection() error {
	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)

	MongoConn.Collections = map[string]*mongo.Collection{}
	if err := MongoConn.Client.Disconnect(ctx); err != nil {
		return err
	}

	return nil
}
