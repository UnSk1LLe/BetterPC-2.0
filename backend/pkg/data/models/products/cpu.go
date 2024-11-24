package products

import (
	"BetterPC_2.0/pkg/data/models/products/general"
	generalResponses "BetterPC_2.0/pkg/data/models/products/general/responses"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
)

type Cpu struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	General        general.General    `bson:"general" json:"general"`
	Main           MainCpu            `bson:"main"  json:"main"`
	Cores          CoresCpu           `bson:"cores" json:"cores,omitempty"`
	ClockFrequency ClockFrequencyCpu  `bson:"clock_frequency" json:"clock_frequency,omitempty"`
	Ram            RamCpu             `bson:"ram" json:"ram,omitempty"`
	Tdp            int                `bson:"tdp" json:"tdp,omitempty"`
	Graphics       string             `bson:"graphics" json:"graphics,omitempty"`
	PciE           int                `bson:"pci_e" json:"pci_e,omitempty"`
	MaxTemperature int                `bson:"max_temperature" json:"max_temperature,omitempty"`
}

type MainCpu struct {
	Category   string `bson:"category" json:"category,omitempty"`
	Generation string `bson:"generation" json:"generation,omitempty"`
	Socket     string `bson:"socket" json:"socket,omitempty"`
	Year       int    `bson:"year" json:"year,omitempty"`
}

type CoresCpu struct {
	Pcores           int `bson:"p-cores" json:"p-cores,omitempty"`
	Ecores           int `bson:"e-cores" json:"e-cores,omitempty"`
	Threads          int `bson:"threads" json:"threads,omitempty"`
	TechnicalProcess int `bson:"technical_process" json:"technical_process,omitempty"`
}

type ClockFrequencyCpu struct {
	Pcores         []float64 `bson:"p-cores" json:"p-cores,omitempty"`
	Ecores         []float64 `bson:"e-cores" json:"e-cores,omitempty"`
	FreeMultiplier bool      `bson:"free_multiplier" json:"free_multiplier,omitempty"`
}

type RamCpu struct {
	Channels     int   `bson:"channels" json:"channels,omitempty"`
	MaxFrequency []int `bson:"max_frequency" json:"max_frequency,omitempty"`
	MaxCapacity  int   `bson:"max_capacity" json:"max_capacity,omitempty"`
}

func (cpu Cpu) GetModel() string {
	return cpu.General.Model
}

func (cpu Cpu) GetStock() int {
	return cpu.General.Amount
}

func (cpu Cpu) GetImage() string {
	return cpu.General.Image
}

func (cpu Cpu) SetImage(imageName string) {
	cpu.General.Image = imageName
}

func (cpu Cpu) Standardize() generalResponses.StandardizedProductData {
	var product generalResponses.StandardizedProductData
	product.ProductHeader.ID = cpu.ID.Hex()
	product.ProductHeader.ProductType = "cpu"
	product.General = cpu.General
	product.Name = cpu.Main.Category + " " + cpu.General.Model
	var cores string
	if cpu.Cores.Ecores > 0 {
		cores = "P-cores: " + strconv.Itoa(cpu.Cores.Pcores) + " E-cores: " + strconv.Itoa(cpu.Cores.Ecores) + ","
	} else {
		cores = strconv.Itoa(cpu.Cores.Pcores) + ","
	}
	product.Description = cpu.Main.Category + ", " + cpu.Main.Generation + " Generation, " +
		cpu.Main.Socket + " Socket, " + "Cores: " + cores + " Threads: " + strconv.Itoa(cpu.Cores.Threads) +
		", Technical process " + strconv.Itoa(cpu.Cores.TechnicalProcess) + " nm, "
	return product
}

func (cpu Cpu) CalculateFinalPrice() int {
	return cpu.General.GetFinalPrice()
}
