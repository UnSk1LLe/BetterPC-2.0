package repository

import (
	"BetterPC_2.0/internal/repository/database/mongoDb"
	"BetterPC_2.0/pkg/data/models/users"
	userErrors "BetterPC_2.0/pkg/data/models/users/errors"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type AuthMongo struct {
	db mongoDb.Database
}

func NewAuthMongo(conn mongoDb.Database) *AuthMongo {
	return &AuthMongo{db: conn}
}

func (a *AuthMongo) CreateUser(user users.User) (primitive.ObjectID, error) {
	//checking if user with the given Email already exists
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	res := a.db.GetUsersCollection().FindOne(ctx, bson.M{"user_info.email": user.UserInfo.Email})

	if !errors.Is(res.Err(), mongo.ErrNoDocuments) && res.Err() != nil {
		return primitive.NilObjectID, res.Err()
	}
	if res.Err() == nil {
		return primitive.NilObjectID, userErrors.ErrUserAlreadyExists
	}

	//Inserting the user into users collection
	newUser, err := a.db.GetUsersCollection().InsertOne(ctx, user)
	if err != nil {
		return primitive.NilObjectID, errors.New(fmt.Sprintf("error creating new user: %s", err.Error()))
	}

	return newUser.InsertedID.(primitive.ObjectID), nil
}

func (a *AuthMongo) GetUserByEmail(email string) (users.User, error) {
	var user users.User

	res := a.db.GetUsersCollection().FindOne(context.TODO(), bson.M{"user_info.email": email})
	if res.Err() != nil {
		return users.User{}, res.Err()
	}

	err := res.Decode(&user)
	if err != nil {
		return users.User{}, errors.New("error decoding the user")
	}
	return user, nil
}

func (a *AuthMongo) CheckUserExists(userId primitive.ObjectID) (bool, error) {
	//checking if user with the given id already exists
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res := a.db.GetUsersCollection().FindOne(ctx, bson.M{"_id": userId})

	if res.Err() != nil {
		return false, res.Err()
	}

	return true, nil
}

func (a *AuthMongo) HasRole(userId primitive.ObjectID, roles []string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{
		"_id":            userId,
		"user_info.role": bson.M{"$in": roles},
	}

	res := a.db.GetUsersCollection().FindOne(ctx, filter)

	if res.Err() != nil {
		return false, res.Err()
	}

	return true, nil
}
