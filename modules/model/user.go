package model

import (
	"tubespbbo/base"
)

type User struct {
	base.Model   `gorm:"extends"`
	Id           int64
	Username     string `gorm:"varchar(100)"`
	Password     string `gorm:"varchar(255)"`
	Name         string `gorm:"varchar(40)"`
	Dob          string `gorm:"varchar(10)"`
	Role         Role   `gorm:"smallint"`
	Type         int
	Address      string        `gorm:"varchar(50)"`
	Phone        string        `gorm:"varchar(15)"`
	Email        string        `gorm:"varchar(100) unique=true"`
	Expenses     []Expense     `gorm:"ForeignKey:UserId;references:Id"`
	Transactions []Transaction `gorm:"ForeignKey:UserId;references:Id"`
}

func init() {
}
