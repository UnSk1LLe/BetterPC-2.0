package users

import (
	"BetterPC_2.0/configs"
	userResponses "BetterPC_2.0/pkg/data/models/users/responses"
	"BetterPC_2.0/pkg/email/helpers"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id"`
	UserInfo     UserInfo           `bson:"user_info"`
	Verification Verification       `bson:"verification"`
}

type UserInfo struct {
	Name     string             `bson:"name" json:"name,omitempty"`
	Surname  string             `bson:"surname" json:"surname,omitempty"`
	Dob      primitive.DateTime `bson:"dob" json:"dob,omitempty"`
	Email    string             `bson:"email" json:"email,omitempty"`
	Password string             `bson:"password" json:"password,omitempty"`
	Role     string             `bson:"role" json:"role,omitempty"`
	Image    string             `bson:"image" json:"image,omitempty"`
}

type Verification struct {
	Token      string             `bson:"token"`
	CreatedAt  primitive.DateTime `bson:"created_at"`
	UpdatedAt  primitive.DateTime `bson:"updated_at"`
	ExpiresAt  primitive.DateTime `bson:"expires_at"`
	IsVerified bool               `bson:"is_verified"`
}

func NewUserDefault(token string, cfg *configs.Config) *User {
	user := &User{
		ID: primitive.NewObjectID(),
		UserInfo: UserInfo{
			Image: cfg.User.Image,
			Role:  cfg.User.Roles.CustomerRole,
		},
		Verification: Verification{
			Token:      token,
			CreatedAt:  primitive.NewDateTimeFromTime(time.Now()),
			UpdatedAt:  primitive.NewDateTimeFromTime(time.Now()),
			ExpiresAt:  primitive.NewDateTimeFromTime(time.Now().Add(cfg.Tokens.VerificationTokenTTL)),
			IsVerified: false,
		},
	}
	return user
}

func (u User) ConvertToUserResponse() userResponses.UserResponse {
	return userResponses.UserResponse{
		ID:      u.ID.Hex(),
		Name:    u.UserInfo.Name,
		Surname: u.UserInfo.Surname,
		Dob:     u.UserInfo.Dob.Time(),
		Role:    u.UserInfo.Role,
		Email:   helpers.HideEmail(u.UserInfo.Email),
		Image:   u.UserInfo.Image,
	}
}
