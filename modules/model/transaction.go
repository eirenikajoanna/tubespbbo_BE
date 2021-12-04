package model

import (
	"tubespbbo/base"
)

type Transaction struct {
	base.Model         `gorm:"extends"`
	Id                 int64
	UserId             int64
	ReceiptNumber      string              `gorm:"varchar(100)"`
	Status             string              `gorm:"varchar(100)"`
	TransactionDetails []TransactionDetail `gorm:"ForeignKey:TransactionId;references:Id"`
	Payment            Payment             `gorm:"ForeignKey:TransactionId"`
}

type TransactionDetail struct {
	base.Model    `gorm:"extends"`
	Id            int64
	TransactionId int64
	ProductId     int64
	Quantity      int64
	Product       Product
}

func init() {
}
