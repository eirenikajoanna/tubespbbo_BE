package model

import (
	"tubespbbo/base"
)

type ShoppingCart struct {
	base.Model `gorm:"extends"`
	Id         int64
	UserId     int64
	ProductId  int64
	Quantity   int64 
	Product    Product
}
