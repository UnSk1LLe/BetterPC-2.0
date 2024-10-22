package repository

import (
	"BetterPC_2.0/pkg/data/models/orders"
	"BetterPC_2.0/pkg/data/models/products/details"
	"BetterPC_2.0/pkg/data/models/products/general"
	"BetterPC_2.0/pkg/data/models/users"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Product interface {
	general.General
	UpdateGeneralInfo(productId primitive.ObjectID, collection *mongo.Collection, input general.UpdateGeneralInput)
}

type Authorization interface {
	CreateUser(user users.User) (primitive.ObjectID, error)
	UpdateUser(userId primitive.ObjectID, input users.UpdateUserInput) error
	DeleteUser(userId primitive.ObjectID) error
	GetAllUsers(filter bson.M)
	GetUserById(userId primitive.ObjectID)
}

type Orders interface {
	CreateOrder(order orders.Order) (primitive.ObjectID, error)
	UpdateOrder(orderId primitive.ObjectID, input orders.UpdateOrderInput) error
	SetOrderStatus(orderId primitive.ObjectID, status string) error
	DeleteOrder(orderId primitive.ObjectID) error
}

type Cpu interface {
	CreateCpu(cpu details.Cpu) (primitive.ObjectID, error)
	UpdateCpu(cpuId primitive.ObjectID, input details.UpdateCpuInput) (primitive.ObjectID, error)
	DeleteCpu(cpuId primitive.ObjectID) error
	GetAllCpu(filter bson.M) ([]details.Cpu, error)
	GetCpuById(cpuId primitive.ObjectID) (details.Cpu, error)
}

type Motherboard interface {
	CreateMotherboard(motherboard details.Motherboard) (primitive.ObjectID, error)
	UpdateMotherboard(motherboardId primitive.ObjectID, input details.UpdateMotherboardInput) (primitive.ObjectID, error)
	DeleteMotherboard(motherboardId primitive.ObjectID) error
	GetAllMotherboard(filter bson.M) ([]details.Motherboard, error)
	GetMotherboardById(motherboardId primitive.ObjectID) (details.Motherboard, error)
}

type Ssd interface {
	CreateSsd(ssd details.Ssd) (primitive.ObjectID, error)
	UpdateSsd(ssdId primitive.ObjectID, input details.UpdateSsdInput) (primitive.ObjectID, error)
	DeleteSsd(ssdId primitive.ObjectID) error
	GetAllSsd(filter bson.M) ([]details.Ssd, error)
	GetSsdById(ssdId primitive.ObjectID) (details.Ssd, error)
}

type Hdd interface {
	CreateHdd(hdd details.Hdd) (primitive.ObjectID, error)
	UpdateHdd(hddId primitive.ObjectID, input details.UpdateHddInput) (primitive.ObjectID, error)
	DeleteHdd(hddId primitive.ObjectID) error
	GetAllHdd(filter bson.M) ([]details.Hdd, error)
	GetHddById(hddId primitive.ObjectID) (details.Hdd, error)
}

type Ram interface {
	CreateRam(ram details.Ram) (primitive.ObjectID, error)
	UpdateRam(ramId primitive.ObjectID, input details.UpdateRamInput) (primitive.ObjectID, error)
	DeleteRam(ramId primitive.ObjectID) error
	GetAllRam(filter bson.M) ([]details.Ram, error)
	GetRamById(ramId primitive.ObjectID) (details.Ram, error)
}

type Gpu interface {
	CreateGpu(gpu details.Gpu) (primitive.ObjectID, error)
	UpdateGpu(gpuId primitive.ObjectID, input details.UpdateGpuInput) (primitive.ObjectID, error)
	DeleteGpu(gpuId primitive.ObjectID) error
	GetAllGpu(filter bson.M) ([]details.Gpu, error)
	GetGpuById(gpuId primitive.ObjectID) (details.Gpu, error)
}

type Cooling interface {
	CreateCooling(cooling details.Cooling) (primitive.ObjectID, error)
	UpdateCooling(coolingId primitive.ObjectID, input details.UpdateCoolingInput) (primitive.ObjectID, error)
	DeleteCooling(coolingId primitive.ObjectID) error
	GetAllCooling(filter bson.M) ([]details.Cooling, error)
	GetCoolingById(coolingId primitive.ObjectID) (details.Cooling, error)
}

type Housing interface {
	CreateHousing(housing details.Housing) (primitive.ObjectID, error)
	UpdateHousing(housingId primitive.ObjectID, input details.UpdateHousingInput) (primitive.ObjectID, error)
	DeleteHousing(housingId primitive.ObjectID) error
	GetAllHousing(filter bson.M) ([]details.Housing, error)
	GetHousingById(housingId primitive.ObjectID) (details.Housing, error)
}

type PowerSupply interface {
	CreatePowerSupply(powerSupply details.PowerSupply) (primitive.ObjectID, error)
	UpdatePowerSupply(powerSupplyId primitive.ObjectID, input details.UpdatePowerSupplyInput) (primitive.ObjectID, error)
	DeletePowerSupply(powerSupplyId primitive.ObjectID) error
	GetAllPowerSupply(filter bson.M) ([]details.PowerSupply, error)
	GetPowerSupplyById(powerSupplyId primitive.ObjectID) (details.PowerSupply, error)
}

type Repository struct {
	Authorization
	Cpu
	Motherboard
	Ram
	Gpu
	Cooling
	PowerSupply
	Hdd
	Ssd
	Housing
	Orders
}

func NewRepository(MongoConnection *MongoDbConnection) *Repository {
	return &Repository{}
}
