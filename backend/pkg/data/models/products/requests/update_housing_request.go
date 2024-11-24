package requests

import (
	"BetterPC_2.0/pkg/data/helpers/decomposers"
	"BetterPC_2.0/pkg/data/helpers/validators"
	"BetterPC_2.0/pkg/data/models/products"
	generalRequests "BetterPC_2.0/pkg/data/models/products/general/requests"
)

type UpdateHousingRequest struct {
	General         *generalRequests.UpdateGeneralRequest `bson:"general"`
	FormFactor      string                                `bson:"form_factor"`
	DriveBays       *products.DriveBays                   `bson:"drive_bays"`
	MbFormFactor    string                                `bson:"mb_form_factor"`
	PsFormFactor    string                                `bson:"ps_form_factor"`
	ExpansionSlots  *int                                  `bson:"expansion_slots"`
	GraphicCardSize *int                                  `bson:"graphic_card_size"`
	CoolerHeight    *int                                  `bson:"cooler_height"`
	Size            *[]int                                `bson:"size"`
	Weight          *float64                              `bson:"weight"`
}

type DriveBays struct {
	D35 *int `bson:"3_5"`
	D25 *int `bson:"2_5"`
}

func (housingRequest *UpdateHousingRequest) Validate() error {
	return validators.ValidateStruct(housingRequest)
}

func (housingRequest *UpdateHousingRequest) Decompose() (map[string]interface{}, error) {
	return decomposers.DecomposeWithTag(housingRequest, "bson")
}

func (housingRequest *UpdateHousingRequest) SetImage(imageName *string) {
	housingRequest.General.Image = imageName
}
