package seeds

import (
	"tubespbbo/db"
	"tubespbbo/modules/model"
)

func (s *Seed) SeedTransaction() {
	var Transaction []model.Transaction
	var count int64
	db.Orm.Model(&Transaction).Count(&count)

	if count > 0 {
		return
	}

	Transactions := make([]model.Transaction, 3)
	Transactions[0] = model.Transaction{
		UserId:           1,
		ReceiptNumber:           "111222333",
		Status:           "Selesai",
	}
	Transactions[1] = model.Transaction{
		UserId:           2,
		ReceiptNumber:           "111222444",
		Status:           "Dikemas",
	}
	Transactions[2] = model.Transaction{
		UserId:           3,
		ReceiptNumber:           "111222555",
		Status:           "Belum Proses",
	}

	_ = db.Orm.Create(&Transactions)
}
