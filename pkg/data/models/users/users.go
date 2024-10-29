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
	Name     string    `bson:"name" json:"name,omitempty"`
	Surname  string    `bson:"surname" json:"surname,omitempty"`
	Dob      time.Time `bson:"dob" json:"dob,omitempty"`
	Email    string    `bson:"email" json:"email,omitempty"`
	Password string    `bson:"password" json:"password,omitempty"`
	Role     string    `bson:"role" json:"role,omitempty"`
	Image    string    `bson:"image" json:"image,omitempty"`
}

type UpdateUserInput struct {
	Name    *string    `bson:"name"`
	Surname *string    `bson:"surname"`
	Dob     *time.Time `bson:"dob"`
}

type RegisterInput struct {
	Name     string `form:"name" binding:"required"`
	Surname  string `form:"surname" binding:"required"`
	Dob      string `form:"dob" binding:"required"`
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

type UserResponse struct {
	ID      string    `json:"id"`
	Name    string    `json:"name"`
	Surname string    `json:"surname"`
	Dob     time.Time `json:"dob"`
	Email   string    `json:"email"`
}

func (u User) ConvertToUserResponse() *UserResponse {
	return &UserResponse{
		ID:      u.ID.Hex(),
		Name:    u.UserInfo.Name,
		Surname: u.UserInfo.Surname,
		Dob:     u.UserInfo.Dob,
		Email:   u.UserInfo.Email,
	}
}

func NewUserDefault() *User {
	user := &User{
		ID: primitive.NewObjectID(),
		UserInfo: UserInfo{
			Role:  "CUSTOMER",
			Image: "/assets/icons/userdefault.png",
		},
		VerificationToken: "", //TODO verification token for email verification and password recovery
		Verified:          false,
	}
	return user
}

func (u UpdateUserInput) Validate() error {
	return general.ValidateStruct(&u)
}
