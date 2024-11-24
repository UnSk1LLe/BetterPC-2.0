package requests

import (
	"BetterPC_2.0/pkg/data/helpers/decomposers"
	"BetterPC_2.0/pkg/data/helpers/validators"
	generalRequests "BetterPC_2.0/pkg/data/models/products/general/requests"
)

type UpdatePowerSupplyRequest struct {
	General     *generalRequests.UpdateGeneralRequest `bson:"general" json:"general"`
	FormFactor  string                                `bson:"form_factor"`
	OutputPower *int                                  `bson:"output_power"`
	Connectors  *UpdateConnectors                     `bson:"connectors"`
	Modules     *bool                                 `bson:"modules"`
	MbPower     *int                                  `bson:"mb_power"`
	CpuPower    *UpdateCpuPower                       `bson:"cpu_power"`
}

type UpdateConnectors struct {
	Sata  *int   `bson:"SATA"`
	Molex *int   `bson:"MOLEX"`
	PciE  *[]int `bson:"PCI_E"`
}

type UpdateCpuPower struct {
	Amount *int   `bson:"amount"`
	Type   *[]int `bson:"type"`
}

func (powerSupplyRequest *UpdatePowerSupplyRequest) Validate() error {
	return validators.ValidateStruct(powerSupplyRequest)
}

func (powerSupplyRequest *UpdatePowerSupplyRequest) Decompose() (map[string]interface{}, error) {
	return decomposers.DecomposeWithTag(powerSupplyRequest, "bson")
}

func (powerSupplyRequest *UpdatePowerSupplyRequest) SetImage(imageName *string) {
	powerSupplyRequest.General.Image = imageName
}
