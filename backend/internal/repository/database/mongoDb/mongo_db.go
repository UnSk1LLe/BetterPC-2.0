package mongoDb

import (
	"BetterPC_2.0/configs"
	"BetterPC_2.0/pkg/data/models/products"
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

type Database interface {
	GetUsersCollection() *mongo.Collection
	GetOrdersCollection() *mongo.Collection
	GetProductCollection(productType products.ProductType) *mongo.Collection
	GetClient() *mongo.Client
}

type MongoConnection struct {
	Client             *mongo.Client
	UsersCollection    *mongo.Collection
	OrdersCollection   *mongo.Collection
	ProductCollections map[string]*mongo.Collection
}

func (db *MongoConnection) GetUsersCollection() *mongo.Collection {
	return db.UsersCollection
}

func (db *MongoConnection) GetOrdersCollection() *mongo.Collection {
	return db.OrdersCollection
}

func (db *MongoConnection) GetProductCollection(productType products.ProductType) *mongo.Collection {
	collection, ok := db.ProductCollections[productType.String()]
	if !ok {
		return nil
	}
	return collection
}

func (db *MongoConnection) GetClient() *mongo.Client {
	return db.Client
}

func MustConnectMongo(cfg *configs.Config, logger *logging.Logger) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	mongoDbUrl := fmt.Sprintf("mongodb+srv://%s:%s@%s/?%s",
		cfg.MongoDB.Username, cfg.MongoDB.Password, cfg.MongoDB.ClusterAddress, cfg.MongoDB.Options)

	opts := options.Client().ApplyURI(mongoDbUrl).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		logger.Fatalf("error connecting to MongoDB client: %v", err)
	}

	usersCollection, err := ConnectCollection(client, cfg.MongoDB.UsersDbName, cfg.MongoDB.UsersCollectionName, logger)
	if err != nil {
		logger.Fatalf("error initializing users collection: %v", err)
	}

	ordersCollection, err := ConnectCollection(client, cfg.MongoDB.UsersDbName, cfg.MongoDB.UsersCollectionName, logger)
	if err != nil {
		logger.Fatalf("error initializing orders collection: %v", err)
	}

	MongoConn = MongoConnection{
		Client:             client,
		UsersCollection:    usersCollection,
		OrdersCollection:   ordersCollection,
		ProductCollections: make(map[string]*mongo.Collection),
	}

	for _, collectionName := range cfg.MongoDB.ProductsCollectionNames { //TODO include dependency from the product types
		MongoConn.ProductCollections[collectionName], err = ConnectCollection(client, cfg.MongoDB.ShopDbName, collectionName, logger)
		if err != nil {
			logger.Fatalf("error initializing %s collection: %s", collectionName, err.Error())
		}
	}
}

func ConnectCollection(client *mongo.Client, dbName, collectionName string, logger *logging.Logger) (*mongo.Collection, error) {
	collection := client.Database(dbName).Collection(collectionName)
	if collection == nil {
		return nil, errors.New(fmt.Sprintf("nil pointer for collection <%s> in db <%s>.", collectionName, dbName))
	}
	logger.Infof("Collection \"%s\" found", collectionName)

	res := collection.FindOne(context.TODO(), bson.D{})
	if errors.Is(res.Err(), mongo.ErrNoDocuments) {
		logger.Warnf("Collection \"%s\" has no documents", collectionName)
	}

	return collection, nil
}

func GetMongoDB() (Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	err := MongoConn.Client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	return &MongoConn, nil
}

func CloseConnection() error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	MongoConn.UsersCollection = nil
	MongoConn.OrdersCollection = nil
	MongoConn.ProductCollections = map[string]*mongo.Collection{}

	if err := MongoConn.Client.Disconnect(ctx); err != nil {
		return err
	}

	return nil
}
