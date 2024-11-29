package decoders

import (
	"BetterPC_2.0/pkg/data/models/products"
	"encoding/json"
)

func DecodeCpuList(data []byte) ([]*products.Cpu, error) {
	var cpuList []*products.Cpu
	if err := json.Unmarshal(data, &cpuList); err != nil {
		return nil, err
	}
	return cpuList, nil
}

func DecodeMotherboardList(data []byte) ([]*products.Motherboard, error) {
	var cpuList []*products.Motherboard
	if err := json.Unmarshal(data, &cpuList); err != nil {
		return nil, err
	}
	return cpuList, nil
}

func DecodeRamList(data []byte) ([]*products.Ram, error) {
	var cpuList []*products.Ram
	if err := json.Unmarshal(data, &cpuList); err != nil {
		return nil, err
	}
	return cpuList, nil
}

func DecodeGpuList(data []byte) ([]*products.Gpu, error) {
	var cpuList []*products.Gpu
	if err := json.Unmarshal(data, &cpuList); err != nil {
		return nil, err
	}
	return cpuList, nil
}

func DecodeSsdList(data []byte) ([]*products.Ssd, error) {
	var cpuList []*products.Ssd
	if err := json.Unmarshal(data, &cpuList); err != nil {
		return nil, err
	}
	return cpuList, nil
}

func DecodeHddList(data []byte) ([]*products.Hdd, error) {
	var cpuList []*products.Hdd
	if err := json.Unmarshal(data, &cpuList); err != nil {
		return nil, err
	}
	return cpuList, nil
}

func DecodeCoolingList(data []byte) ([]*products.Cooling, error) {
	var cpuList []*products.Cooling
	if err := json.Unmarshal(data, &cpuList); err != nil {
		return nil, err
	}
	return cpuList, nil
}

func DecodePowerSupplyList(data []byte) ([]*products.PowerSupply, error) {
	var cpuList []*products.PowerSupply
	if err := json.Unmarshal(data, &cpuList); err != nil {
		return nil, err
	}
	return cpuList, nil
}

func DecodeHousingList(data []byte) ([]*products.Housing, error) {
	var cpuList []*products.Housing
	if err := json.Unmarshal(data, &cpuList); err != nil {
		return nil, err
	}
	return cpuList, nil
}
