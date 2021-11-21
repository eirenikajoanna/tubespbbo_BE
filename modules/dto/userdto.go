package dto

import (
	"tubespbbo/base"
)

type UserDTO struct {
	base.DTO
	Username string `json:"username"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Role     int    `json:"role"`
}
