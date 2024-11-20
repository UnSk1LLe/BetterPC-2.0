package cookies

import (
	"BetterPC_2.0/pkg/data/models/orders"
	"BetterPC_2.0/pkg/data/models/products"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/url"
)

func GetCart(c *gin.Context, userId string) (map[products.ProductType][]orders.ItemHeader, error) {
	var cart map[products.ProductType][]orders.ItemHeader

	cartCookieName := fmt.Sprintf("cart-%s", userId)
	cartCookie, err := c.Cookie(cartCookieName)
	if err != nil {
		return cart, err
	}

	decodedValue, err := url.QueryUnescape(cartCookie)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(decodedValue), &cart)
	if err != nil {
		return cart, err
	}

	return cart, nil
}
