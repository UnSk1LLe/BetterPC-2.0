package patch

import (
	"BetterPC_2.0/pkg/data/helpers/decomposers"
	"BetterPC_2.0/pkg/data/helpers/validators"
	"time"
)

type UpdateUserVerificationDataRequest struct {
	VerificationToken *string    `bson:"verification.token,omitempty"`
	CreatedAt         *time.Time `bson:"verification.created_at,omitempty"`
	UpdatedAt         *time.Time `bson:"verification.updated_at,omitempty"`
	ExpiresAt         *time.Time `bson:"verification.expires_at,omitempty"`
	IsVerified        *bool      `bson:"verification.is_verified,omitempty"`
}

func (u *UpdateUserVerificationDataRequest) Validate() error {
	return validators.ValidateStruct(&u)
}

func (u *UpdateUserVerificationDataRequest) Decompose() (map[string]interface{}, error) {
	return decomposers.DecomposeWithTag(u, "bson")
}
