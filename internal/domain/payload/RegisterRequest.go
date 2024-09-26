package payload

import _ "car-rental-api/internal/domain"

type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Address  string `json:"address"`
	Dob      string `json:"birthDate"`
}
