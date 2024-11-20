package orders

import (
	generalResponses "BetterPC_2.0/pkg/data/models/products/general/responses"
)

//TODO try to use {"product_type": []ItemAmounts}
//like {"cpu": [{"id1": 3}, {"id2": 1}], "motherboard"[{"id3": 1}]}

type ItemHeader struct {
	ID             string `json:"id" binding:"required"`
	SelectedAmount int    `json:"selected_amount" binding:"required,gte=1,lte=20"`
}

type Item struct {
	generalResponses.StandardizedProductData
	Price          int `json:"price"`
	SelectedAmount int `json:"selected_amount"`
	MaxAmount      int `json:"max_amount"`
}

func (i Item) ItemFinalPrice() int {
	finalPrice := i.SelectedAmount * i.General.GetFinalPrice()
	return finalPrice
}
