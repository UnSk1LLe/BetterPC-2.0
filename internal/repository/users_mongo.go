package repository

import (
	"BetterPC_2.0/pkg/data/models/users"
	"BetterPC_2.0/pkg/database/mongoDb"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type UsersMongo struct {
	db *mongoDb.MongoConnection
}

func NewUsersMongo(conn *mongoDb.MongoConnection) *UsersMongo {
	return &UsersMongo{db: conn}
}

func (u *UsersMongo) Create(user users.User) (primitive.ObjectID, error) {
	//checking if user with the given Email already exists
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	res := u.db.Collections["users"].FindOne(ctx, bson.M{"user_info.email": user.UserInfo.Email})

	if !errors.Is(res.Err(), mongo.ErrNoDocuments) && res.Err() != nil {
		return primitive.NilObjectID, res.Err()
	}
	if res.Err() == nil {
		return primitive.NilObjectID, errors.New("user with this email already exists")
	}

	//Inserting the user into users collection
	newUser, err := u.db.Collections["users"].InsertOne(ctx, user)
	if err != nil {
		return primitive.NilObjectID, errors.New(fmt.Sprintf("error creating new user: %s", err.Error()))
	}

	return newUser.InsertedID.(primitive.ObjectID), nil
}

func (u *UsersMongo) Update(userId primitive.ObjectID, input users.UpdateUserInput) error {
	return nil
}

func (u *UsersMongo) Delete(userId primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := u.db.Collections["users"].DeleteOne(ctx, bson.M{"_id": userId})
	if err != nil {
		return errors.New(fmt.Sprintf("error deleting user: %s", err.Error()))
	}

	return nil
}

func (u *UsersMongo) GetList(filter bson.M) ([]users.User, error) {
	var usersList []users.User

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cur, err := u.db.Collections["users"].Find(ctx, filter)
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

	result := u.db.Collections["users"].FindOne(ctx, bson.M{"_id": userId})
	if result.Err() != nil {
		return users.User{}, result.Err()
	}

	err := result.Decode(&user)
	if err != nil {
		return users.User{}, err
	}

	return user, nil
}
