package requests

import (
	"BetterPC_2.0/pkg/data/helpers/decomposers"
	"BetterPC_2.0/pkg/data/helpers/validators"
	"BetterPC_2.0/pkg/data/models/products/general"
)

type UpdateCoolingRequest struct {
	General    *general.General `bson:"general"`
	Type       *string          `bson:"type"`
	Sockets    *[]string        `bson:"sockets"`
	Fans       *[]int           `bson:"fans"`
	Rpm        *[]int           `bson:"rpm"`
	Tdp        *int             `bson:"tdp"`
	NoiseLevel *int             `bson:"noise_level"`
	MountType  *string          `bson:"mount_type"`
	Power      *int             `bson:"power"`
	Height     *int             `bson:"height"`
}

func (m *UpdateCoolingRequest) Validate() error {
	return validators.ValidateStruct(m)
}

func (m *UpdateCoolingRequest) Decompose() (map[string]interface{}, error) {
	return decomposers.DecomposeWithTag(m, "bson")
}
