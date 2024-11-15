package repository

import (
	"BetterPC_2.0/pkg/data/models/orders"
	"BetterPC_2.0/pkg/database/mongoDb"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type OrderMongo struct {
	db *mongoDb.MongoConnection
}

func NewOrderMongo(db *mongoDb.MongoConnection) *OrderMongo {
	return &OrderMongo{db: db}
}

func (o *OrderMongo) Create(order orders.Order) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	newOrder, err := o.db.Collections["orders"].InsertOne(ctx, order)
	if err != nil {
		return primitive.NilObjectID, errors.New(fmt.Sprintf("error creating new order: %s", err.Error()))
	}

	return newOrder.InsertedID.(primitive.ObjectID), nil
}

func (o *OrderMongo) Update(orderId primitive.ObjectID, input orders.UpdateOrderInput) error {
	return nil
}

func (o *OrderMongo) SetStatus(orderId primitive.ObjectID, status string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := o.db.Collections["orders"].UpdateOne(ctx, bson.M{"_id": orderId}, bson.M{"$set": bson.M{"status": status}})
	if err != nil {
		return errors.New(fmt.Sprintf("error updating order status: %s", err.Error()))
	}

	return nil
}

func (o *OrderMongo) Delete(orderId primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := o.db.Collections["orders"].DeleteOne(ctx, bson.M{"_id": orderId})
	if err != nil {
		return errors.New(fmt.Sprintf("error deleting order: %s", err.Error()))
	}

	return nil
}

func (o *OrderMongo) GetById(id primitive.ObjectID) (orders.Order, error) {
	var order orders.Order

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res := o.db.Collections["orders"].FindOne(ctx, bson.M{"_id": id})
	if res.Err() != nil {
		return order, errors.New(fmt.Sprintf("error getting order: %s", res.Err()))
	}

	err := res.Decode(&order)
	if err != nil {
		return order, errors.New(fmt.Sprintf("error decoding order: %s", res.Err()))
	}

	return order, nil
}

func (o *OrderMongo) GetList(filter bson.M) ([]orders.Order, error) {
	var orderList []orders.Order

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cur, err := o.db.Collections["orders"].Find(ctx, filter)
	if err != nil {
		return orderList, errors.New(fmt.Sprintf("error getting orderList: %s", err.Error()))
	}

	err = cur.All(ctx, &orderList)
	if err != nil {
		return orderList, errors.New(fmt.Sprintf("error decoding orderList: %s", err.Error()))
	}

	return orderList, nil
}
