package users

import (
	"BetterPC_2.0/configs"
	"BetterPC_2.0/pkg/data/models/products/general"
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
	VerificationToken string             `bson:"verification_token"`
	CreatedAt         primitive.DateTime `bson:"created_at"`
	UpdatedAt         primitive.DateTime `bson:"updated_at"`
	ExpiresAt         primitive.DateTime `bson:"expires_at"`
	IsVerified        bool               `bson:"is_verified"`
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
	Role    string    `json:"role"`
	Email   string    `json:"email"`
	Image   string    `json:"image"`
}

func (u User) ConvertToUserResponse() UserResponse {
	return UserResponse{
		ID:      u.ID.Hex(),
		Name:    u.UserInfo.Name,
		Surname: u.UserInfo.Surname,
		Dob:     u.UserInfo.Dob.Time(),
		Role:    u.UserInfo.Role,
		Email:   helpers.HideEmail(u.UserInfo.Email),
		Image:   u.UserInfo.Image,
	}
}

func NewUserDefault(cfg *configs.Config) *User {
	user := &User{
		ID: primitive.NewObjectID(),
		UserInfo: UserInfo{
			Image: cfg.User.Image,
			Role:  cfg.User.Roles.CustomerRole,
		},
		Verification: Verification{
			VerificationToken: "", //TODO verification token for email verification and password recovery
			CreatedAt:         primitive.NewDateTimeFromTime(time.Now()),
			UpdatedAt:         primitive.NewDateTimeFromTime(time.Now()),
			ExpiresAt:         primitive.NewDateTimeFromTime(time.Now().Add(cfg.User.VerificationTokenTTL)),
			IsVerified:        false,
		},
	}
	return user
}

func (u UpdateUserInput) Validate() error {
	return general.ValidateStruct(&u)
}
