package details

import (
	"BetterPC_2.0/pkg/data/models/products/general"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
)

type PowerSupply struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	General     general.General    `bson:"general"`
	FormFactor  string             `bson:"form_factor"`
	OutputPower int                `bson:"output_power"`
	Connectors  connectors         `bson:"connectors"`
	Modules     bool               `bson:"modules"`
	MbPower     int                `bson:"mb_power"`
	CpuPower    bsoncore.Array     `bson:"cpu_power"`
}

type connectors struct {
	Sata  int   `bson:"SATA"`
	Molex int   `bson:"MOLEX"`
	PciE  []int `bson:"PCI_E"`
}

type UpdatePowerSupplyInput struct {
	General     *general.General `bson:"general"`
	FormFactor  *string          `bson:"form_factor"`
	OutputPower *int             `bson:"output_power"`
	Connectors  *connectors      `bson:"connectors"`
	Modules     *bool            `bson:"modules"`
	MbPower     *int             `bson:"mb_power"`
	CpuPower    *bsoncore.Array  `bson:"cpu_power"`
}

func (p UpdatePowerSupplyInput) Validate() error {
	return general.ValidateStruct(&p)
}
