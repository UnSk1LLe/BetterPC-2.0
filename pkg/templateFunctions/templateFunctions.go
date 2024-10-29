package templateFunctions

import (
	"html/template"
	"math"
)

var TmplFuncs = template.FuncMap{
	"add":                    add,
	"sub":                    sub,
	"until":                  until,
	"mul":                    mul,
	"div":                    div,
	"calculateDiscountPrice": calculateDiscountPrice,
}

func add(a, b int) int { return a + b }
func sub(a, b int) int { return a - b }
func mul(a, b int) int { return a * b }
func div(a, b int) int {
	if b == 0 {
		return 0
	}
	return a / b
}

func until(start, end int) []int {
	result := make([]int, end-start+1)
	for i := range result {
		result[i] = start + i
	}
	return result
}

func calculateDiscountPrice(price, discount int) int {
	if discount == 0 {
		return price
	}
	discountPercentage := float64(discount) / 100.0
	discountedPrice := float64(price) * (1 - discountPercentage)
	return int(math.Ceil(discountedPrice))
}
