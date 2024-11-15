package requests

import (
	"BetterPC_2.0/pkg/data/helpers/decomposers"
	"BetterPC_2.0/pkg/data/helpers/validators"
	"BetterPC_2.0/pkg/data/models/products/general"
)

type UpdateRamRequest struct {
	General      *general.General `bson:"general"`
	Capacity     *int             `bson:"capacity"`
	Number       *int             `bson:"number"`
	FormFactor   *string          `bson:"form_factor"`
	Rank         *int             `bson:"rank"`
	Type         *string          `bson:"type"`
	Frequency    *int             `bson:"frequency"`
	Bandwidth    *int             `bson:"bandwidth"`
	CasLatency   *string          `bson:"cas_latency"`
	TimingScheme *[]int           `bson:"timing_scheme"`
	Voltage      *float64         `bson:"voltage"`
	Cooling      *string          `bson:"cooling"`
	Height       *int             `bson:"height"`
}

func (m *UpdateRamRequest) Validate() error {
	return validators.ValidateStruct(m)
}

func (m *UpdateRamRequest) Decompose() (map[string]interface{}, error) {
	return decomposers.DecomposeWithTag(m, "bson")
}
