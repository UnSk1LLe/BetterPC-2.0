package service

/*type CpuService struct {
	repo        repository.Cpu
	productRepo repository.Product
}

func NewCpuService(repo repository.Cpu, productRepo repository.Product) *CpuService {
	return &CpuService{repo: repo, productRepo: productRepo}
}

func (c *CpuService) CreateCpu(cpu details.Cpu) (primitive.ObjectID, error) {
	return c.repo.Create(cpu)
}

func (c *CpuService) UpdateGeneralInfo(cpuId primitive.ObjectID, input general.UpdateGeneralInput) error {
	cpuCollection := "cpu"
	return c.productRepo.UpdateGeneralInfo(cpuId, input, cpuCollection)
}

func (c *CpuService) UpdateCpu(cpuId primitive.ObjectID, updateCpuInput details.UpdateCpuInput) (primitive.ObjectID, error) {
	return c.repo.Update(cpuId, updateCpuInput)
}

func (c *CpuService) DeleteProduct(cpuId primitive.ObjectID, productType string) error {
	return c.productRepo.DeleteProduct(cpuId, productType)
}

func (c *CpuService) GetCpuList(filter bson.M) ([]details.Cpu, error) {
	return c.repo.GetList(filter)
}

func (c *CpuService) GetCpuById(cpuId primitive.ObjectID) (details.Cpu, error) {
	return c.repo.GetById(cpuId)
}
*/
