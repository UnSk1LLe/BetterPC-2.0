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

type AuthMongo struct {
	db *mongoDb.MongoConnection
}

func NewAuthMongo(conn *mongoDb.MongoConnection) *AuthMongo {
	return &AuthMongo{db: conn}
}

func (a *AuthMongo) CreateUser(user users.User) (primitive.ObjectID, error) {
	//checking if user with the given Email already exists
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	res := a.db.Collections["users"].FindOne(ctx, bson.M{"user_info.email": user.UserInfo.Email})

	if !errors.Is(res.Err(), mongo.ErrNoDocuments) && res.Err() != nil {
		return primitive.NilObjectID, res.Err()
	}
	if res.Err() == nil {
		return primitive.NilObjectID, errors.New("user with this email already exists")
	}

	//Inserting the user into users collection
	newUser, err := a.db.Collections["users"].InsertOne(ctx, user)
	if err != nil {
		return primitive.NilObjectID, errors.New(fmt.Sprintf("error creating new user: %s", err.Error()))
	}

	return newUser.InsertedID.(primitive.ObjectID), nil
}

func (a *AuthMongo) GetUser(email string) (users.User, error) {
	var user users.User

	res := a.db.Collections["users"].FindOne(context.TODO(), bson.M{"user_info.email": email})
	if res.Err() != nil {
		return users.User{}, res.Err()
	}

	err := res.Decode(&user)
	if err != nil {
		return users.User{}, errors.New("error decoding the user")
	}
	return user, nil
}
