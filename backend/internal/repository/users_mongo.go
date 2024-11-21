package repository

import (
	"BetterPC_2.0/internal/repository/database/mongoDb"
	"BetterPC_2.0/pkg/data/models/users"
	userErrors "BetterPC_2.0/pkg/data/models/users/errors"
	userUpdateRequests "BetterPC_2.0/pkg/data/models/users/requests/patch"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type UsersMongo struct {
	db mongoDb.Database
}

func NewUsersMongo(conn mongoDb.Database) *UsersMongo {
	return &UsersMongo{db: conn}
}

func (u *UsersMongo) Create(user users.User) (primitive.ObjectID, error) {
	//checking if user with the given Email already exists
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res := u.db.GetUsersCollection().FindOne(ctx, bson.M{"user_info.email": user.UserInfo.Email})

	if !errors.Is(res.Err(), mongo.ErrNoDocuments) && res.Err() != nil {
		return primitive.NilObjectID, res.Err()
	}
	if res.Err() == nil {
		return primitive.NilObjectID, userErrors.ErrUserAlreadyExists
	}

	//Inserting the user into users collection
	newUser, err := u.db.GetUsersCollection().InsertOne(ctx, user)
	if err != nil {
		return primitive.NilObjectID, errors.New(fmt.Sprintf("error creating new user: %s", err.Error()))
	}

	return newUser.InsertedID.(primitive.ObjectID), nil
}

func (u *UsersMongo) UpdateUserImageById(userId primitive.ObjectID, imageUrl string) error {
	return nil
}

func (u *UsersMongo) UpdateUserInfoById(userId primitive.ObjectID, input userUpdateRequests.UpdateUserInfoRequest) error {
	/*fields, values := input.Decompose()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := u.db.GetUsersCollection().UpdateByID(ctx, userId, input)*/
	return nil
}

func (u *UsersMongo) UpdatePasswordById(userId primitive.ObjectID, password string) error {
	return nil
}

func (u *UsersMongo) UpdateVerificationDataById(userId primitive.ObjectID, input userUpdateRequests.UpdateUserVerificationDataRequest) error {
	return nil
}

func (u *UsersMongo) DeleteById(userId primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	delRes, err := u.db.GetUsersCollection().DeleteOne(ctx, bson.M{"_id": userId})
	if err != nil {
		return errors.New(fmt.Sprintf("error deleting user: %s", err.Error()))
	} else if delRes.DeletedCount == 0 {
		return userErrors.ErrUserNotFound
	}

	return nil
}

func (u *UsersMongo) GetList(filter bson.M) ([]users.User, error) {
	var usersList []users.User

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cur, err := u.db.GetUsersCollection().Find(ctx, filter)
	if err != nil {
		return usersList, errors.New(fmt.Sprintf("error finding users: %s", err.Error()))
	}

	err = cur.All(ctx, &usersList)
	if err != nil {
		return usersList, errors.New(fmt.Sprintf("error decoding users cursor: %s", err.Error()))
	}

	return usersList, nil
}

func (u *UsersMongo) GetById(userId primitive.ObjectID) (users.User, error) {
	var user users.User

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result := u.db.GetUsersCollection().FindOne(ctx, bson.M{"_id": userId})
	if errors.Is(result.Err(), mongo.ErrNoDocuments) {
		return user, userErrors.ErrUserNotFound
	}
	if result.Err() != nil {
		return user, result.Err()
	}

	err := result.Decode(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}
