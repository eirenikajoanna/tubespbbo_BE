package seeds

import (
	"tubespbbo/db"
	"tubespbbo/domains/transaction/model"
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
		PembayaranId:     1,
		TanggalTransaksi: "2021-10-15",
		NoResi:           "111222333",
		Status:           "Selesai",
	}
	Transactions[1] = model.Transaction{
		UserId:           2,
		PembayaranId:     2,
		TanggalTransaksi: "2021-10-16",
		NoResi:           "111222444",
		Status:           "Proses",
	}
	Transactions[2] = model.Transaction{
		UserId:           3,
		PembayaranId:     3,
		TanggalTransaksi: "2021-10-17",
		NoResi:           "111222555",
		Status:           "Belum Proses",
	}

	_ = db.Orm.Create(&Transactions)
}