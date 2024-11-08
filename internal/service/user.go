package service

import (
	"BetterPC_2.0/internal/repository"
	"BetterPC_2.0/pkg/data/models/users"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (u *UserService) Create(user users.User) (primitive.ObjectID, error) {
	return u.repo.Create(user)
}

func (u *UserService) Update(userId primitive.ObjectID, input users.UpdateUserInput) error {
	return u.repo.Update(userId, input)
}

func (u *UserService) Delete(userId primitive.ObjectID) error {
	return u.repo.Delete(userId)
}

func (u *UserService) GetList(filter bson.M) ([]users.User, error) {
	return u.repo.GetList(filter)
}

func (u *UserService) GetById(userId primitive.ObjectID) (users.User, error) {
	return u.repo.GetById(userId)
}
