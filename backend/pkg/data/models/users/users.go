package users

import (
	"BetterPC_2.0/configs"
	userErrors "BetterPC_2.0/pkg/data/models/users/errors"
	userResponses "BetterPC_2.0/pkg/data/models/users/responses"
	"BetterPC_2.0/pkg/email/helpers"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
	"time"
)

const MaxUserRoleLength = 20

type UserRole string

var UserRoles = struct {
	Customer      UserRole
	ShopAssistant UserRole
	Admin         UserRole
}{
	Customer:      "CUSTOMER",
	ShopAssistant: "SHOP_ASSISTANT",
	Admin:         "ADMIN",
} // TODO change all user roles dependencies to this struct

func (u UserRole) String() string {
	return string(u)
}

func UserRoleFromString(input string) (UserRole, error) {
	if len(input) > MaxUserRoleLength {
		return "", errors.Wrapf(userErrors.ErrInvalidUserRole, "user role must be shorter than %d characters", MaxUserRoleLength)
	}

	normalizedInput := strings.ToLower(strings.TrimSpace(input))
	if len(normalizedInput) == 0 {
		return "", errors.Wrap(userErrors.ErrInvalidUserRole, "user role cannot not be empty")
	}

	switch UserRole(normalizedInput) {
	case UserRoles.Customer:
		return UserRoles.Customer, nil
	case UserRoles.ShopAssistant:
		return UserRoles.ShopAssistant, nil
	case UserRoles.Admin:
		return UserRoles.Admin, nil
	}

	return "", userErrors.ErrInvalidUserRole
}

type User struct {
	ID           primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	UserInfo     UserInfo           `bson:"user_info" json:"user_info"`
	Verification Verification       `bson:"verification" json:"verification"`
	StripeId     string             `bson:"stripe_id,omitempty" json:"stripe_id,omitempty"`
	CreatedAt    primitive.DateTime `bson:"created_at" json:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt    primitive.DateTime `bson:"updated_at" json:"updated_at,omitempty" json:"updated_at,omitempty"`
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
	currentTime := time.Now()
	currentDateTime := primitive.NewDateTimeFromTime(currentTime)
	user := &User{
		ID: primitive.NewObjectID(),
		UserInfo: UserInfo{
			Role: cfg.User.Roles.CustomerRole,
		},
		Verification: Verification{
			Token:      token,
			CreatedAt:  currentDateTime,
			UpdatedAt:  currentDateTime,
			ExpiresAt:  primitive.NewDateTimeFromTime(currentTime.Add(cfg.Tokens.VerificationTokenTTL)),
			IsVerified: false,
		},
		CreatedAt: currentDateTime,
		UpdatedAt: currentDateTime,
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
