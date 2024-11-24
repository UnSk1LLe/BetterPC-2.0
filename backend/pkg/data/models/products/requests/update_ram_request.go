package requests

import (
	"BetterPC_2.0/pkg/data/helpers/decomposers"
	"BetterPC_2.0/pkg/data/helpers/validators"
	generalRequests "BetterPC_2.0/pkg/data/models/products/general/requests"
)

type UpdateRamRequest struct {
	General      *generalRequests.UpdateGeneralRequest `bson:"general"`
	Capacity     *int                                  `bson:"capacity"`
	Number       *int                                  `bson:"number"`
	FormFactor   string                                `bson:"form_factor"`
	Rank         *int                                  `bson:"rank"`
	Type         string                                `bson:"type"`
	Frequency    *int                                  `bson:"frequency"`
	Bandwidth    *int                                  `bson:"bandwidth"`
	CasLatency   string                                `bson:"cas_latency"`
	TimingScheme *[]int                                `bson:"timing_scheme"`
	Voltage      *float64                              `bson:"voltage"`
	Cooling      string                                `bson:"cooling"`
	Height       *int                                  `bson:"height"`
}

func (ramRequest *UpdateRamRequest) Validate() error {
	return validators.ValidateStruct(ramRequest)
}

func (ramRequest *UpdateRamRequest) Decompose() (map[string]interface{}, error) {
	return decomposers.DecomposeWithTag(ramRequest, "bson")
}

func (ramRequest *UpdateRamRequest) SetImage(imageName *string) {
	ramRequest.General.Image = imageName
}
