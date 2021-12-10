package seeds

import (
	"tubespbbo/db"
	"tubespbbo/modules/model"
)

func (s *Seed) SeedTD() {
	var pm []model.TransactionDetail
	var count int64
	db.Orm.Model(&pm).Count(&count)

	if count > 0 {
		return
	}

	TransactionDetails := make([]model.TransactionDetail, 5)
	TransactionDetails[0] = model.TransactionDetail{
		TransactionId:1,
		ProductId:1,
		Quantity:2,
	}
	TransactionDetails[1] = model.TransactionDetail{
		TransactionId:2,
		ProductId:2,
		Quantity:2,
	}
	TransactionDetails[2] = model.TransactionDetail{
		TransactionId:3,
		ProductId:3,
		Quantity:2,
	}
	TransactionDetails[3] = model.TransactionDetail{
		TransactionId:3,
		ProductId:4,
		Quantity:4,
	}
	TransactionDetails[3] = model.TransactionDetail{
		TransactionId:4,
		ProductId:5,
		Quantity:1,
	}
	TransactionDetails[4] = model.TransactionDetail{
		TransactionId:5,
		ProductId:3,
		Quantity:2,
	}

	_ = db.Orm.Create(&TransactionDetails)
}
