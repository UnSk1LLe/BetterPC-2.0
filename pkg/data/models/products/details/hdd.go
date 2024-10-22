package details

import (
	"BetterPC_2.0/pkg/data/models/products/general"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Hdd struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	General      general.General    `bson:"general"`
	Type         string             `bson:"type"`
	Capacity     int                `bson:"capacity"`
	Interface    string             `bson:"interface"`
	WriteMethod  string             `bson:"write_method"`
	TransferRate int                `bson:"transfer_rate"`
	SpindleSpeed int                `bson:"spindle_speed"`
	FormFactor   string             `bson:"form_factor"`
	Mftb         int                `bson:"mftb"`
	Size         []float64          `bson:"size"`
	Weight       int                `bson:"weight"`
}

type UpdateHddInput struct {
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

func (h UpdateHddInput) Validate() error {
	return general.ValidateStruct(&h)
}
