package patch

import (
	"BetterPC_2.0/pkg/data/helpers/decomposers"
	"BetterPC_2.0/pkg/data/helpers/validators"
	"time"
)

type UpdateUserInfoRequest struct {
	Name    string     `bson:"user_info.name" json:"name"`
	Surname string     `bson:"user_info.surname" json:"surname"`
	Dob     *time.Time `bson:"user_info.dob" json:"dob"`
}

func (u *UpdateUserInfoRequest) Validate() error {
	return validators.ValidateStruct(&u)
}

func (u *UpdateUserInfoRequest) Decompose() (map[string]interface{}, error) {
	return decomposers.DecomposeWithTag(u, "bson")
}
