package requests

import (
	"BetterPC_2.0/pkg/data/helpers/decomposers"
	"BetterPC_2.0/pkg/data/helpers/validators"
	"BetterPC_2.0/pkg/data/models/products/general"
)

type UpdatePowerSupplyRequest struct {
	General    *general.General `bson:"general" json:"general"`
	Type       *string          `bson:"type" json:"type"`
	Capacity   *int             `bson:"capacity" json:"capacity"`
	Interface  *string          `bson:"interface" json:"interface"`
	MemoryType *string          `bson:"memory_type" json:"memory_type"`
	Read       *int             `bson:"read" json:"read"`
	Write      *int             `bson:"write" json:"write"`
	FormFactor *string          `bson:"form_factor" json:"form_factor"`
	Mftb       *float64         `bson:"mftb" json:"mftb"`
	Size       *[]float64       `bson:"size" json:"size"`
	Weight     *int             `bson:"weight" json:"weight"`
}

func (m *UpdatePowerSupplyRequest) Validate() error {
	return validators.ValidateStruct(m)
}

func (m *UpdatePowerSupplyRequest) Decompose() (map[string]interface{}, error) {
	return decomposers.DecomposeWithTag(m, "bson")
}
