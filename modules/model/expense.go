package model

import (
	"tubespbbo/base"
)

type Expense struct {
	base.Model  `gorm:"extends"`
	Id          int64
	UserId      int64
	Name        string  `gorm:"varchar(100)"`
	ReleaseDate string  `gorm:"varchar(10)"`
	Quantity    float32 `gorm:"double"`
	Category    int64
	Description string `gorm:"varchar(255)"`
	TotalAmount	float32 `gorm:"double"`
}

func init() {
}
