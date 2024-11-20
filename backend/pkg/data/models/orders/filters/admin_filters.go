package orderFilters

import "time"

type AdminOrderFilters struct {
	DateFrom     time.Time `json:"date_from,omitempty"`
	DateTo       time.Time `json:"date_to,omitempty"`
	UserId       string    `json:"user_id,omitempty"`
	ProductTypes []string  `json:"product_types,omitempty"`
	PriceFrom    int       `json:"price_from,omitempty"`
	PriceTo      int       `json:"price_to,omitempty"`
	Statuses     []string  `json:"statuses,omitempty"`
}
