package requests

import (
	"BetterPC_2.0/pkg/data/helpers/decomposers"
	"BetterPC_2.0/pkg/data/helpers/validators"
	generalRequests "BetterPC_2.0/pkg/data/models/products/general/requests"
)

type UpdateHddRequest struct {
	General      *generalRequests.UpdateGeneralRequest `bson:"general"`
	Type         *string                               `bson:"type"`
	Capacity     *int                                  `bson:"capacity"`
	Interface    *string                               `bson:"interface"`
	WriteMethod  *string                               `bson:"write_method"`
	TransferRate *int                                  `bson:"transfer_rate"`
	SpindleSpeed *int                                  `bson:"spindle_speed"`
	FormFactor   *string                               `bson:"form_factor"`
	Mftb         *int                                  `bson:"mftb"`
	Size         *[]float64                            `bson:"size"`
	Weight       *int                                  `bson:"weight"`
}

func (hddRequest *UpdateHddRequest) Validate() error {
	return validators.ValidateStruct(hddRequest)
}

func (hddRequest *UpdateHddRequest) Decompose() (map[string]interface{}, error) {
	return decomposers.DecomposeWithTag(hddRequest, "bson")
}

func (hddRequest *UpdateHddRequest) SetImage(imageName *string) {
	hddRequest.General.Image = imageName
}
