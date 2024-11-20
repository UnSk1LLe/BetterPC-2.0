package general

type General struct {
	Manufacturer string `bson:"manufacturer" json:"manufacturer"`
	Model        string `bson:"model" json:"model"`
	Price        int    `bson:"price" json:"price"`
	Discount     int    `bson:"discount" json:"discount"`
	Amount       int    `bson:"amount" json:"amount"`
	Image        string `bson:"image" json:"image"`
}

func (g *General) GetFinalPrice() int {
	return g.Price - (g.Price * g.Discount / 100)
}
