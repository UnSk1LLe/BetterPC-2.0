package responses

import "time"

type UserResponse struct {
	ID      string    `json:"id"`
	Name    string    `json:"name"`
	Surname string    `json:"surname"`
	Dob     time.Time `json:"dob"`
	Role    string    `json:"role"`
	Email   string    `json:"email"`
	Image   string    `json:"image"`
}
