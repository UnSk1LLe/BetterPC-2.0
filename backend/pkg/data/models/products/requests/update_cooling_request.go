package requests

import (
	"BetterPC_2.0/pkg/data/helpers/decomposers"
	"BetterPC_2.0/pkg/data/helpers/validators"
	generalRequests "BetterPC_2.0/pkg/data/models/products/general/requests"
)

type UpdateCoolingRequest struct {
	General    *generalRequests.UpdateGeneralRequest `bson:"general" json:"general"`
	Type       *string                               `bson:"type" json:"type"`
	Sockets    *[]string                             `bson:"sockets" json:"sockets"`
	Fans       *[]int                                `bson:"fans" json:"fans"`
	Rpm        *[]int                                `bson:"rpm" json:"rpm"`
	Tdp        *int                                  `bson:"tdp" json:"tdp"`
	NoiseLevel *int                                  `bson:"noise_level" json:"noiseLevel"`
	MountType  *string                               `bson:"mount_type" json:"mountType"`
	Power      *int                                  `bson:"power" json:"power"`
	Height     *int                                  `bson:"height" json:"height"`
}

func (m *UpdateCoolingRequest) Validate() error {
	return validators.ValidateStruct(m)
}

func (m *UpdateCoolingRequest) Decompose() (map[string]interface{}, error) {
	return decomposers.DecomposeWithTag(m, "bson")
}
