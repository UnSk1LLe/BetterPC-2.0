package users

import (
	"BetterPC_2.0/pkg/data/models/products/general"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID                primitive.ObjectID `bson:"_id"`
	UserInfo          UserInfo           `bson:"user_info"`
	VerificationToken string             `bson:"verification_token"`
	Verified          bool               `bson:"verified"`
}

type UserInfo struct {
	Name     string    `bson:"name"`
	Surname  string    `bson:"surname"`
	Dob      time.Time `bson:"dob"`
	Email    string    `bson:"email"`
	Password string    `bson:"password"`
	Role     string    `bson:"role"`
	Image    string    `bson:"image"`
}

type UpdateUserInput struct {
	Name    *string    `bson:"name"`
	Surname *string    `bson:"surname"`
	Dob     *time.Time `bson:"dob"`
}

func (u UpdateUserInput) Validate() error {
	return general.ValidateStruct(&u)
}
