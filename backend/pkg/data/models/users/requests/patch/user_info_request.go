package patch

import (
	"BetterPC_2.0/pkg/data/helpers/decomposers"
	"BetterPC_2.0/pkg/data/helpers/validators"
	"time"
)

type UpdateUserInfoRequest struct {
	Name    *string    `bson:"name"`
	Surname *string    `bson:"surname"`
	Dob     *time.Time `bson:"dob"`
}

func (u *UpdateUserInfoRequest) Validate() error {
	return validators.ValidateStruct(&u)
}

func (u *UpdateUserInfoRequest) Decompose() (map[string]interface{}, error) {
	return decomposers.DecomposeWithTag(u, "bson")
}
