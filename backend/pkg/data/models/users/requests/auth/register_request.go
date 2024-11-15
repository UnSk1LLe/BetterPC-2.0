package auth

type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Surname  string `json:"surname" binding:"required"`
	Dob      string `json:"dob" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
