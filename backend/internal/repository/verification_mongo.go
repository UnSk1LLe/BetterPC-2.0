package repository

import (
	"BetterPC_2.0/internal/repository/database/mongoDb"
	"BetterPC_2.0/pkg/data/models/users"
	userErrors "BetterPC_2.0/pkg/data/models/users/errors"
	userUpdateRequests "BetterPC_2.0/pkg/data/models/users/requests/patch"
	"context"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type VerificationMongo struct {
	db mongoDb.Database
}

func NewVerificationMongo(conn mongoDb.Database) *VerificationMongo {
	return &VerificationMongo{db: conn}
}

func (v *VerificationMongo) SetTokenByEmail(email, token string, expTime primitive.DateTime) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"user_info.email": email}
	update := bson.M{
		"$set": bson.M{
			"verification.token":      token,
			"verification.created_at": time.Now(),
			"verification.expires_at": expTime,
		},
	}

	res, err := v.db.GetUsersCollection().UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if res.MatchedCount == 0 {
		return userErrors.ErrUserNotFound
	}

	return nil
}

func (v *VerificationMongo) CompareUserToken(token string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{
		"verification.token":      token,
		"verification.expires_at": bson.M{"$lt": primitive.NewDateTimeFromTime(time.Now())},
	}

	res := v.db.GetUsersCollection().FindOne(ctx, filter)
	if errors.Is(res.Err(), mongo.ErrNoDocuments) {
		return false, userErrors.ErrUserNotFound
	}
	if res.Err() != nil {
		return false, res.Err()
	}

	return false, nil
}

func (v *VerificationMongo) IsVerifiedUser(email string) (bool, error) {
	var user users.User

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"user_info.email": email}

	res := v.db.GetUsersCollection().FindOne(ctx, filter)
	if errors.Is(res.Err(), mongo.ErrNoDocuments) {
		return false, userErrors.ErrUserNotFound
	}
	if res.Err() != nil {
		return false, res.Err()
	}

	err := res.Decode(&user)
	if err != nil {
		return false, errors.New("error decoding verification data: " + err.Error())
	}

	return user.Verification.IsVerified, nil
}

func (v *VerificationMongo) UpdateVerificationDataById(userId primitive.ObjectID, input userUpdateRequests.UpdateUserVerificationDataRequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fieldsValues, err := input.Decompose()
	if err != nil {
		return err
	}

	logrus.Info(fieldsValues)

	update := bson.M{"$set": fieldsValues}

	updRes, err := v.db.GetUsersCollection().UpdateByID(ctx, userId, update)
	if err != nil {
		return errors.New("error updating verification data: " + err.Error())
	}
	if updRes.MatchedCount == 0 {
		return userErrors.ErrUserNotFound
	}

	return nil
}

func (v *VerificationMongo) GetUserByVerificationToken(token string) (users.User, error) {
	var user users.User

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"verification.token": token}

	res := v.db.GetUsersCollection().FindOne(ctx, filter)

	if errors.Is(res.Err(), mongo.ErrNoDocuments) {
		return user, userErrors.ErrUserNotFound
	}
	if res.Err() != nil {
		return user, res.Err()
	}

	err := res.Decode(&user)
	if err != nil {
		return user, errors.New("error decoding verification data: " + err.Error())
	}

	if user.Verification.ExpiresAt < primitive.NewDateTimeFromTime(time.Now()) {
		return user, userErrors.ErrTokenExpired
	}

	return user, nil
}

func (v *VerificationMongo) UpdateUserPasswordById(userId primitive.ObjectID, password string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	update := bson.M{"$set": bson.M{
		"user_info.password": password,
		"verification.token": "",
	}}

	updRes, err := v.db.GetUsersCollection().UpdateByID(ctx, userId, update)
	if err != nil {
		return errors.New("error updating verification data: " + err.Error())
	}
	if updRes.MatchedCount == 0 {
		return userErrors.ErrUserNotFound
	}

	return nil
}
