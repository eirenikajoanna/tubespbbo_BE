package dto

import (
	"tubespbbo/base"
)

type ExpenseDTO struct {
	base.DTO
	Name        string   `json:"name"`
	ReleaseDate string   `json:"releaseDate"`
	Quantity    float32  `json:"quantity"`
	Category    int64    `json:"category"`
	Description string   `json:"description"`
	TotalAmount float32  `json:"totalAmount"`
	UserId      int64    `json:"userId"`
	User        *UserDTO `json:"user"`
}
