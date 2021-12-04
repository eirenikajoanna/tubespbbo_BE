package dto

import (
	"tubespbbo/base"
)

type CreateShoppingCart struct {
	UserId    int64 `json:"userId"`
	ProductId int64 `json:"productId"`
	Quantity  int64 `json:"quantity"`
}

type UpdateShoppingCart struct {
	Quantity int64 `json:"quantity"`
}

type ShoppingCart struct {
	base.DTO
	UserId    int64       `json:"userId"`
	ProductId int64       `json:"productId"`
	Product   *ProductDTO `json:"product"`
	Quantity  int64       `json:"quantity"`
}
