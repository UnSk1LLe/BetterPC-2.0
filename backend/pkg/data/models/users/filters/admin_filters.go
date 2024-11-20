package filters

import "time"

type AdminUserFilters struct {
	DateFrom   *time.Time `json:"date_from"`
	DateTo     *time.Time `json:"date_to"`
	IsVerified *bool      `json:"is_verified"`
	Roles      []string   `json:"roles"`
}
