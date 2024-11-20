package service

import (
	"BetterPC_2.0/internal/repository"
	"BetterPC_2.0/pkg/data/models/users"
	userFilters "BetterPC_2.0/pkg/data/models/users/filters"
	userUpdateRequests "BetterPC_2.0/pkg/data/models/users/requests/patch"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (u *UserService) Create(user users.User) (string, error) {
	userObjId, err := u.repo.Create(user)
	if err != nil {
		return "", err
	}

	return userObjId.String(), nil
}

func (u *UserService) Update(userId string, input userUpdateRequests.UpdateUserInfoRequest) error {
	userObjId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}

	return u.repo.UpdateUserInfoById(userObjId, input)
}

func (u *UserService) Delete(userId string) error {
	userObjId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}

	return u.repo.DeleteById(userObjId)
}

func (u *UserService) GetList(filters userFilters.AdminUserFilters) ([]users.User, error) {
	var bsonFilter bson.M

	if filters.DateFrom != nil && filters.DateTo != nil {
		bsonFilter["created_at"] = bson.M{
			"$gte": *filters.DateFrom,
			"$lte": *filters.DateTo,
		}
	} else if filters.DateFrom != nil {
		bsonFilter["created_at"] = bson.M{"$gte": *filters.DateFrom}
	} else if filters.DateTo != nil {
		bsonFilter["created_at"] = bson.M{"$lte": *filters.DateTo}
	}

	if filters.IsVerified != nil {
		bsonFilter["verification.is_verified"] = *filters.IsVerified
	}

	if len(filters.Roles) > 0 {
		bsonFilter["user_info.roles"] = bson.M{"$in": filters.Roles}
	}

	return u.repo.GetList(bsonFilter)
}

func (u *UserService) GetById(userId string) (users.User, error) {
	userObjId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return users.User{}, err
	}

	return u.repo.GetById(userObjId)
}
