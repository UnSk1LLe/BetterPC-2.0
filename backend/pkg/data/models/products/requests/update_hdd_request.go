package requests

import (
	"BetterPC_2.0/pkg/data/helpers/decomposers"
	"BetterPC_2.0/pkg/data/helpers/validators"
	"BetterPC_2.0/pkg/data/models/products/general"
)

type UpdateHddRequest struct {
	General      *general.General `bson:"general"`
	Type         *string          `bson:"type"`
	Capacity     *int             `bson:"capacity"`
	Interface    *string          `bson:"interface"`
	WriteMethod  *string          `bson:"write_method"`
	TransferRate *int             `bson:"transfer_rate"`
	SpindleSpeed *int             `bson:"spindle_speed"`
	FormFactor   *string          `bson:"form_factor"`
	Mftb         *int             `bson:"mftb"`
	Size         *[]float64       `bson:"size"`
	Weight       *int             `bson:"weight"`
}

func (m *UpdateHddRequest) Validate() error {
	return validators.ValidateStruct(m)
}

func (m *UpdateHddRequest) Decompose() (map[string]interface{}, error) {
	return decomposers.DecomposeWithTag(m, "bson")
}
