package details

import (
	"BetterPC_2.0/pkg/data/models/products/general"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Cooling struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	General    general.General    `bson:"general"`
	Type       string             `bson:"type"`
	Sockets    []string           `bson:"sockets"`
	Fans       []int              `bson:"fans"`
	Rpm        []int              `bson:"rpm"`
	Tdp        int                `bson:"tdp"`
	NoiseLevel int                `bson:"noise_level"`
	MountType  string             `bson:"mount_type"`
	Power      int                `bson:"power"`
	Height     int                `bson:"height"`
}

type UpdateCoolingInput struct {
	General    *general.General `bson:"general"`
	Type       *string          `bson:"type"`
	Sockets    *[]string        `bson:"sockets"`
	Fans       *[]int           `bson:"fans"`
	Rpm        *[]int           `bson:"rpm"`
	Tdp        *int             `bson:"tdp"`
	NoiseLevel *int             `bson:"noise_level"`
	MountType  *string          `bson:"mount_type"`
	Power      *int             `bson:"power"`
	Height     *int             `bson:"height"`
}

func (c UpdateCoolingInput) Validate() error {
	return general.ValidateStruct(&c)
}
