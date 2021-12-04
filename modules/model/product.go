package model

import (
	"tubespbbo/base"
)

type Product struct {
	base.Model         `gorm:"extends"`
	Id                 int64
	Name               string  `gorm:"varchar(100)"`
	Code               string  `gorm:"varchar(100)"`
	Description        string  `gorm:"varchar(255)"`
	Price              float32 `gorm:"double"`
	Quantity           int64
	ImageUrl           string              `gorm:"varchar(255)"`
	TransactionDetails []TransactionDetail `gorm:"ForeignKey:ProductId;references:Id"`
	ShoppingCart       []ShoppingCart      `gorm:"ForeignKey:ProductId;references:Id"`
}

func init() {
}
