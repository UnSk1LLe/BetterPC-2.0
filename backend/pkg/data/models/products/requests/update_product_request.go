package requests

var ProductUpdateRequestFactory = map[string]func() ProductUpdateRequest{
	"cpu":         func() ProductUpdateRequest { return &UpdateCpuRequest{} },
	"motherboard": func() ProductUpdateRequest { return &UpdateMotherboardRequest{} },
	"ram":         func() ProductUpdateRequest { return &UpdateRamRequest{} },
	"gpu":         func() ProductUpdateRequest { return &UpdateGpuRequest{} },
	"ssd":         func() ProductUpdateRequest { return &UpdateSsdRequest{} },
	"hdd":         func() ProductUpdateRequest { return &UpdateHddRequest{} },
	"powersupply": func() ProductUpdateRequest { return &UpdatePowerSupplyRequest{} },
	"cooling":     func() ProductUpdateRequest { return &UpdateCoolingRequest{} },
	"housing":     func() ProductUpdateRequest { return &UpdateHousingRequest{} },
}

type ProductUpdateRequest interface {
	Validate() error
	Decompose() (map[string]interface{}, error)
}
