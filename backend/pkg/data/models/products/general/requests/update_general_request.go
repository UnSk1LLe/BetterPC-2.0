package requests

import (
	"BetterPC_2.0/pkg/data/helpers/decomposers"
	"BetterPC_2.0/pkg/data/helpers/validators"
)

type UpdateGeneralRequest struct {
	Manufacturer *string `bson:"manufacturer" json:"manufacturer"`
	Model        *string `bson:"model" json:"model"`
	Price        *int    `bson:"price" json:"price"`
	Discount     *int    `bson:"discount" json:"discount"`
	Amount       *int    `bson:"amount" json:"amount"`
	Image        *string `bson:"image" json:"image"`
}

func (u *UpdateGeneralRequest) Validate() error {
	return validators.ValidateStruct(&u)
}

func (u *UpdateGeneralRequest) Decompose() (map[string]interface{}, error) {
	return decomposers.DecomposeWithTag(u, "bson")
}
