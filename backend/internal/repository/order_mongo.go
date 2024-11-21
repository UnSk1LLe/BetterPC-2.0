package repository

import (
	"BetterPC_2.0/internal/repository/database/mongoDb"
	"BetterPC_2.0/pkg/data/models/orders"
	orderErrors "BetterPC_2.0/pkg/data/models/orders/errors"
	"BetterPC_2.0/pkg/data/models/products"
	productErrors "BetterPC_2.0/pkg/data/models/products/errors"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type OrderMongo struct {
	db mongoDb.Database
}

func NewOrderMongo(db mongoDb.Database) *OrderMongo {
	return &OrderMongo{db: db}
}

func (o *OrderMongo) Create(order orders.Order) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	productList := order.ProductList

	callback := func(sesCtx mongo.SessionContext) (interface{}, error) {
		err := o.updateProductStock(sesCtx, productList, false)
		if err != nil {
			return nil, err
		}
		res, err := o.db.GetOrdersCollection().InsertOne(ctx, order)
		if err != nil {
			return nil, err
		}

		return res.InsertedID, nil
	}

	session, err := o.db.GetClient().StartSession()
	if err != nil {
		return primitive.ObjectID{}, err
	}
	defer session.EndSession(ctx)

	result, err := session.WithTransaction(ctx, callback)
	if err != nil {
		message := "error creating new order"
		return primitive.NilObjectID, errors.Wrap(err, message)
	}
	logrus.Info(result)

	return result.(primitive.ObjectID), nil
}

func (o *OrderMongo) Cancel(orderId primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	callback := func(sesCtx mongo.SessionContext) (interface{}, error) {
		var order orders.Order

		filter := bson.M{"_id": orderId}
		update := bson.M{"$set": bson.M{
			"status":     orders.OrderStatuses.Cancelled,
			"updated_at": time.Now(),
		}}
		res := o.db.GetOrdersCollection().FindOneAndUpdate(sesCtx, filter, update)
		if res.Err() != nil {
			switch {
			case errors.Is(res.Err(), mongo.ErrNoDocuments):
				return nil, orderErrors.ErrOrderNotFound
			}
			return nil, errors.Wrapf(res.Err(), "error updating order %s", orderId.Hex())
		}

		err := res.Decode(&order)
		if err != nil {
			return nil, err
		}

		switch order.Status {
		case orders.OrderStatuses.Cancelled:
			return nil, orderErrors.ErrOrderCancelled
		case orders.OrderStatuses.Closed:
			return nil, orderErrors.ErrOrderClosed
		}

		err = o.updateProductStock(sesCtx, order.ProductList, true)
		if err != nil {
			return nil, err
		}

		return order.ID, nil
	}

	session, err := o.db.GetClient().StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(ctx)

	_, err = session.WithTransaction(ctx, callback)
	if err != nil {
		return err
	}

	return nil
}

func (o *OrderMongo) updateProductStock(ctx mongo.SessionContext, productList map[products.ProductType][]orders.ProductHeader, increment bool) error {
	for productType, productHeaders := range productList {
		collection := o.db.GetProductCollection(productType)
		if collection == nil {
			return productErrors.ErrUnsupportedProductType
		}

		var operations []mongo.WriteModel

		if increment {
			for _, header := range productHeaders {
				operations = append(operations, mongo.NewUpdateOneModel().
					SetFilter(bson.M{
						"_id": header.ID,
					}).
					SetUpdate(bson.M{"$inc": bson.M{"general.amount": header.SelectedAmount}}))
			}
		} else {
			for _, header := range productHeaders {
				operations = append(operations, mongo.NewUpdateOneModel().
					SetFilter(bson.M{
						"_id":            header.ID,
						"general.amount": bson.M{"$gte": header.SelectedAmount},
					}).
					SetUpdate(bson.M{"$inc": bson.M{"general.amount": -header.SelectedAmount}}))
			}
		}

		bwRes, err := collection.BulkWrite(ctx, operations)

		if err != nil {
			return err
		}
		if int(bwRes.MatchedCount) != len(productHeaders) {
			return productErrors.ErrInsufficientStock
		}
		if int(bwRes.MatchedCount) != int(bwRes.ModifiedCount) {
			return errors.Wrap(productErrors.ErrProductNotModified, "amount of products in stock is not changed")
		}

	}

	return nil
}

func (o *OrderMongo) Update(orderId primitive.ObjectID, input orders.UpdateOrderInput) error {

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	callback := func(sesCtx mongo.SessionContext) (interface{}, error) {
		var order orders.Order

		filter := bson.M{"_id": orderId}
		update := bson.M{"$set": bson.M{}}

		res := o.db.GetOrdersCollection().FindOneAndUpdate(sesCtx, filter, update)
		if res.Err() != nil {
			switch {
			case errors.Is(res.Err(), mongo.ErrNoDocuments):
				return nil, orderErrors.ErrOrderNotFound
			}
			return nil, res.Err()
		}

		err := res.Decode(&order)
		if err != nil {
			return nil, err
		}

		//only proceed if the order is active
		if err := order.IsActive(); err != nil {
			return nil, err
		}

		return nil, nil
	}

	session, err := o.db.GetClient().StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(ctx)

	if _, err := session.WithTransaction(ctx, callback); err != nil {
		return err
	}

	return nil
}

func (o *OrderMongo) SetStatus(orderId primitive.ObjectID, status orders.OrderStatus) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := o.db.GetOrdersCollection().UpdateOne(ctx, bson.M{"_id": orderId}, bson.M{"$set": bson.M{"status": status}})
	if err != nil {
		return errors.New(fmt.Sprintf("error updating order status: %s", err.Error()))
	}

	return nil
}

func (o *OrderMongo) Delete(orderId primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	errMsg := "error deleting order"
	defer cancel()

	session, err := o.db.GetClient().StartSession()
	if err != nil {
		return errors.Wrap(err, errMsg)
	}
	defer session.EndSession(ctx)

	callback := func(sesCtx mongo.SessionContext) (interface{}, error) {
		var order orders.Order

		res := o.db.GetOrdersCollection().FindOneAndDelete(sesCtx, bson.M{"_id": orderId})
		if res.Err() != nil {

			switch {
			case errors.Is(res.Err(), mongo.ErrNoDocuments):
				ctx.Done()
				return nil, orderErrors.ErrOrderNotFound
			}

			return nil, res.Err()
		}

		err := res.Decode(&order)
		if err != nil {
			return nil, err
		}
		if order.Status != orders.OrderStatuses.Cancelled && order.Status != orders.OrderStatuses.Closed {
			return nil, orderErrors.ErrActiveOrder
		}

		return nil, nil
	}

	_, err = session.WithTransaction(ctx, callback)
	if err != nil {
		return errors.Wrap(err, errMsg)
	}
	return nil
}

func (o *OrderMongo) GetById(id primitive.ObjectID) (orders.Order, error) {
	var order orders.Order

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res := o.db.GetOrdersCollection().FindOne(ctx, bson.M{"_id": id})
	if res.Err() != nil {
		switch {
		case errors.Is(res.Err(), mongo.ErrNoDocuments):
			return order, errors.Wrapf(orderErrors.ErrOrderNotFound, "error getting order")
		}
		return order, errors.Wrap(res.Err(), "error getting order")
	}

	err := res.Decode(&order)
	if err != nil {
		return order, errors.Wrap(res.Err(), "error decoding order")
	}

	return order, nil
}

func (o *OrderMongo) GetList(filter bson.M) ([]orders.Order, error) {
	var orderList []orders.Order
	errMsg := "error getting order"

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cur, err := o.db.GetOrdersCollection().Find(ctx, filter)
	if err != nil {
		switch {
		case errors.Is(err, mongo.ErrNoDocuments):
			return orderList, errors.Wrapf(orderErrors.ErrOrderNotFound, errMsg)
		}
		return orderList, errors.Wrap(err, errMsg)
	}

	err = cur.All(ctx, &orderList)
	if err != nil {
		return orderList, errors.Wrap(err, "error decoding orderList")
	}

	return orderList, nil
}
