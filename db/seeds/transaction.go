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

	Transactions := make([]model.Transaction, 5)
	Transactions[0] = model.Transaction{
		UserId:1,
		ReceiptNumber:"111222333",
		Status:"Selesai",
	}
	Transactions[1] = model.Transaction{
		UserId:6,
		ReceiptNumber:"-",
		Status:"Dikemas",
	}
	Transactions[2] = model.Transaction{
		UserId:7,
		ReceiptNumber:"-",
		Status:"Belum Proses",
	}
	Transactions[3] = model.Transaction{
		UserId:7,
		ReceiptNumber:"-",
		Status:"Belum Terverifikasi",
	}
	Transactions[4] = model.Transaction{
		UserId:7,
		ReceiptNumber:"111222444",
		Status:"Dikirim",
	}

	_ = db.Orm.Create(&Transactions)
}
