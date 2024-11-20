package builds

import (
	"BetterPC_2.0/pkg/data/models/products"
	"go.mongodb.org/mongo-driver/bson"
	"strings"
)

type Build struct {
	CPU         *products.Cpu
	Motherboard *products.Motherboard
	RAM         *products.Ram
	GPU         *products.Gpu
	SSD         []*products.Ssd
	HDD         []*products.Hdd
	Cooling     *products.Cooling
	PowerSupply *products.PowerSupply
	Housing     *products.Housing
}

func (b *Build) GetCpuFilter() bson.M {
	var filter bson.M
	var conditions []bson.M

	if b.Motherboard != nil {
		conditions = append(conditions, bson.M{"main.socket": b.Motherboard.Socket})
	}
	if b.RAM != nil {
		switch b.RAM.Type {
		case "DDR4":
			conditions = append(conditions, bson.M{"ram.max_frequency.0": bson.M{"$gte": b.RAM.Frequency}})
		case "DDR5":
			conditions = append(conditions, bson.M{"ram.max_frequency.1": bson.M{"$gte": b.RAM.Frequency}})
		}

		conditions = append(conditions, bson.M{"ram.max_capacity": bson.M{"$gte": b.RAM.Capacity * b.RAM.Number}})
	}
	if b.Cooling != nil {
		conditions = append(conditions, bson.M{"main.socket": bson.M{"$in": b.Cooling.Sockets}, "tdp": bson.M{"$lte": b.Cooling.Tdp}})
	}

	if len(conditions) > 0 {
		filter["$and"] = conditions
	}

	return filter
}

func (b *Build) GetMotherboardFilter() bson.M {
	var filter bson.M
	var conditions []bson.M

	if b.CPU != nil {
		conditions = append(conditions, bson.M{"socket": b.CPU.Main.Socket})
	}
	if b.Housing != nil {
		conditions = append(conditions, bson.M{"form_factor": b.Housing.MbFormFactor})
	}
	if b.RAM != nil {
		conditions = append(conditions, bson.M{"ram.type": b.RAM.Type, "ram.max_frequency": bson.M{"$gte": b.RAM.Frequency},
			"ram.max_capacity": bson.M{"$gte": b.RAM.Capacity * b.RAM.Number}, "ram.slots": bson.M{"$gte": b.RAM.Number}})
	}
	if b.SSD != nil {
		m2number := 0
		sata3number := 0
		for _, ssd := range b.SSD {
			switch {
			case strings.Contains(ssd.FormFactor, "M2"):
				m2number++
				conditions = append(conditions, bson.M{"interfaces.M2": bson.M{"$gte": m2number}})
			case strings.Contains(ssd.FormFactor, "2.5") && ssd.Interface == "SATA 3":
				sata3number++
				conditions = append(conditions, bson.M{"interfaces.SATA3": bson.M{"$gte": sata3number}})
			}
		}
	}

	if len(conditions) > 0 {
		filter["$and"] = conditions
	}

	return filter
}

func (b *Build) GetRamFilter() bson.M {
	var filter bson.M
	var conditions []bson.M

	if b.Motherboard != nil {
		conditions = append(conditions, bson.M{"type": b.Motherboard.Ram.Type, "frequency": bson.M{"$lte": b.Motherboard.Ram.MaxFrequency}, "form_factor": "DIMM"})
	}
	if b.CPU != nil {
		if b.CPU.Ram.MaxFrequency[0] != 0 {
			conditions = append(conditions, bson.M{"frequency": bson.M{"$lte": b.CPU.Ram.MaxFrequency[0]}})
		}
		if b.CPU.Ram.MaxFrequency[1] != 0 {
			conditions = append(conditions, bson.M{"frequency": bson.M{"$lte": b.CPU.Ram.MaxFrequency[1]}})
		}
		conditions = append(conditions, bson.M{"$expr": bson.M{"$lte": bson.A{bson.M{"$multiply": bson.A{"$capacity", "$number"}}, b.CPU.Ram.MaxCapacity}}})
	}

	if len(conditions) > 0 {
		filter["$and"] = conditions
	}
	return filter
}

func (b *Build) GetGPUFilter() bson.M {
	var filter bson.M
	var conditions []bson.M

	if b.PowerSupply != nil {
		conditions = append(conditions, bson.M{"tdp_r": bson.M{"$lte": b.PowerSupply.OutputPower}})
	}
	if b.Housing != nil {
		conditions = append(conditions, bson.M{"size.0": bson.M{"$lte": b.Housing.GraphicCardSize}})
	}

	if len(conditions) > 0 {
		filter["$and"] = conditions
	}
	return filter
}

func (b *Build) GetSSDFilter() bson.M {
	var filter bson.M
	var conditions []bson.M

	if b.Motherboard != nil {
		if b.Motherboard.Interfaces.M2 > 0 && b.Motherboard.Interfaces.Sata3 > 0 {
			filter = bson.M{"$or": []bson.M{{"form_factor": bson.M{"$regex": "M.2"}}, {"interface": "SATA 3"}}}
		} else if b.Motherboard.Interfaces.M2 > 0 {
			conditions = append(conditions, bson.M{"form_factor": bson.M{"$regex": "M.2"}})
		} else if b.Motherboard.Interfaces.Sata3 > 0 {
			conditions = append(conditions, bson.M{"interface": "SATA 3"})
		}
	}

	if len(conditions) > 0 {
		filter["$and"] = conditions
	}
	return filter
}

func (b *Build) GetHDDFilter() bson.M {
	var filter bson.M
	var conditions []bson.M

	if b.Motherboard != nil {
		if b.Motherboard.Interfaces.Sata3 > 0 {
			filter = bson.M{"interface": "SATA 3"}
		}
	}

	if len(conditions) > 0 {
		filter["$and"] = conditions
	}
	return filter
}

func (b *Build) GetCoolingFilter() bson.M {
	var filter bson.M
	var conditions []bson.M

	if b.CPU != nil {
		conditions = append(conditions, bson.M{"sockets": b.CPU.Main.Socket})
		conditions = append(conditions, bson.M{"tdp": bson.M{"$gte": b.CPU.Tdp}})
	}
	if b.Motherboard != nil {
		conditions = append(conditions, bson.M{"sockets": b.Motherboard.Socket})
	}
	if b.Housing != nil {
		conditions = append(conditions, bson.M{"height": bson.M{"$lte": b.Housing.CoolerHeight}})
	}

	if len(conditions) > 0 {
		filter["$and"] = conditions
	}
	return filter
}

func (b *Build) GetPowerSupplyFilter() bson.M {
	var filter bson.M
	var conditions []bson.M

	if b.GPU != nil {
		conditions = append(conditions, bson.M{"output_power": bson.M{"$gte": b.GPU.TdpR}})
	}
	if b.Housing != nil {
		conditions = append(conditions, bson.M{"form_factor": b.Housing.PsFormFactor})
	}

	if len(conditions) > 0 {
		filter["$and"] = conditions
	}
	return filter
}

func (b *Build) GetHousingFilter() bson.M {
	var filter bson.M
	var conditions []bson.M

	if b.Motherboard != nil {
		conditions = append(conditions, bson.M{"mb_form_factor": b.Motherboard.FormFactor})
	}
	if b.PowerSupply != nil {
		conditions = append(conditions, bson.M{"ps_form_factor": b.PowerSupply.FormFactor})
	}
	if b.SSD != nil {
		driveBaysNumber := 0
		for _, ssd := range b.SSD {
			if ssd.FormFactor == "2.5" {
				driveBaysNumber++
			}
		}

		if driveBaysNumber > 0 {
			conditions = append(conditions, bson.M{"drive_bays.2_5": bson.M{"$gte": driveBaysNumber}})
		}
	}
	if b.HDD != nil {
		driveBaysNumber := 0
		for _, hdd := range b.HDD {
			if hdd.FormFactor == "3.5" {
				driveBaysNumber++
			}
		}

		if driveBaysNumber > 0 {
			conditions = append(conditions, bson.M{"drive_bays.3_5": bson.M{"$gte": driveBaysNumber}})
		}
	}

	if len(conditions) > 0 {
		filter["$and"] = conditions
	}
	return filter
}
