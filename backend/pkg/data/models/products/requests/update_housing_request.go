package requests

import (
	"BetterPC_2.0/pkg/data/helpers/decomposers"
	"BetterPC_2.0/pkg/data/helpers/validators"
	"BetterPC_2.0/pkg/data/models/products"
	"BetterPC_2.0/pkg/data/models/products/general"
)

type UpdateHousingRequest struct {
	General         *general.General    `bson:"general"`
	FormFactor      *string             `bson:"form_factor"`
	DriveBays       *products.DriveBays `bson:"drive_bays"`
	MbFormFactor    *string             `bson:"mb_form_factor"`
	PsFormFactor    *string             `bson:"ps_form_factor"`
	ExpansionSlots  *int                `bson:"expansion_slots"`
	GraphicCardSize *int                `bson:"graphic_card_size"`
	CoolerHeight    *int                `bson:"cooler_height"`
	Size            *[]int              `bson:"size"`
	Weight          *float64            `bson:"weight"`
}

func (m *UpdateHousingRequest) Validate() error {
	return validators.ValidateStruct(m)
}

func (m *UpdateHousingRequest) Decompose() (map[string]interface{}, error) {
	return decomposers.DecomposeWithTag(m, "bson")
}
