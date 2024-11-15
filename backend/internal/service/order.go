package service

import (
	"BetterPC_2.0/internal/repository"
	"BetterPC_2.0/pkg/data/models/orders"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderService struct {
	repo repository.Order
}

func NewOrderService(repo repository.Order) *OrderService {
	return &OrderService{repo}
}

func (o *OrderService) Create(order orders.Order) (primitive.ObjectID, error) {
	return o.repo.Create(order)
}

func (o *OrderService) Update(orderId primitive.ObjectID, input orders.UpdateOrderInput) error {
	return o.repo.Update(orderId, input)
}

func (o *OrderService) SetStatus(orderId primitive.ObjectID, status string) error {
	return o.repo.SetStatus(orderId, status)
}

func (o *OrderService) Delete(orderId primitive.ObjectID) error {
	return o.repo.Delete(orderId)
}

func (o *OrderService) GetById(id primitive.ObjectID) (orders.Order, error) {
	return o.repo.GetById(id)
}

func (o *OrderService) GetList(filter bson.M) ([]orders.Order, error) {
	return o.repo.GetList(filter)
}
