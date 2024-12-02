package searchEngine

import (
	"BetterPC_2.0/pkg/data/models/products"
	"go.mongodb.org/mongo-driver/bson"
	"strings"
)

func GetSearchFilter(searchQuery string) ([]products.ProductType, bson.M) {
	if searchQuery == "" {
		return products.ProductTypes.GetAll(), bson.M{}
	}

	var productTypeFilter []products.ProductType

	productKeywords := map[products.ProductType][]string{
		products.ProductTypes.Cpu:         {"cpu", "processor", "cores", "threads"},
		products.ProductTypes.Motherboard: {"motherboard", "chipset", "socket", "form factor"},
		products.ProductTypes.Ram:         {"ram", "memory", "ddr", "capacity"},
		products.ProductTypes.Gpu:         {"gpu", "graphics", "card", "video card", "vram"},
		products.ProductTypes.Ssd:         {"ssd", "solid", "state", "drive", "nvme", "sata"},
		products.ProductTypes.Hdd:         {"hdd", "hard drive", "storage", "m2"},
		products.ProductTypes.Cooling:     {"cooling", "cooler", "fan", "tdp"},
		products.ProductTypes.PowerSupply: {"power", "supply", "psu", "wattage", "efficiency"},
		products.ProductTypes.Housing:     {"case", "housing", "chassis"},
	}

	searchWords := strings.Split(searchQuery, " ")
	var searchFilters []bson.M

	for i, word := range searchWords {
		word = strings.ToLower(word)

		searchFilters[i] = bson.M{
			"$or": []bson.M{
				{"general.model": bson.M{"$regex": word, "$options": "i"}},
				{"general.manufacturer": bson.M{"$regex": word, "$options": "i"}},
				{"main.category": bson.M{"$regex": word, "$options": "i"}},
			},
		}

		for productType, keywords := range productKeywords {
			for _, keyword := range keywords {
				if strings.Contains(word, keyword) {
					productTypeFilter = append(productTypeFilter, productType)
					continue
				}
			}
		}
	}

	if len(productTypeFilter) == 0 {
		productTypeFilter = products.ProductTypes.GetAll()
	}

	return productTypeFilter, bson.M{"$and": searchFilters}
}
