package details

import (
	"BetterPC_2.0/pkg/data/models/products/general"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Housing struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	General         general.General    `bson:"general"`
	FormFactor      string             `bson:"form_factor"`
	DriveBays       driveBays          `bson:"drive_bays"`
	MbFormFactor    string             `bson:"mb_form_factor"`
	PsFormFactor    string             `bson:"ps_form_factor"`
	ExpansionSlots  int                `bson:"expansion_slots"`
	GraphicCardSize int                `bson:"graphic_card_size"`
	CoolerHeight    int                `bson:"cooler_height"`
	Size            []int              `bson:"size"`
	Weight          float64            `bson:"weight"`
}

type driveBays struct {
	D35 int `bson:"3_5"`
	D25 int `bson:"2_5"`
}

type UpdateHousingInput struct {
	General         *general.General `bson:"general"`
	FormFactor      *string          `bson:"form_factor"`
	DriveBays       *driveBays       `bson:"drive_bays"`
	MbFormFactor    *string          `bson:"mb_form_factor"`
	PsFormFactor    *string          `bson:"ps_form_factor"`
	ExpansionSlots  *int             `bson:"expansion_slots"`
	GraphicCardSize *int             `bson:"graphic_card_size"`
	CoolerHeight    *int             `bson:"cooler_height"`
	Size            *[]int           `bson:"size"`
	Weight          *float64         `bson:"weight"`
}

func (h UpdateHousingInput) Validate() error {
	return general.ValidateStruct(&h)
}
