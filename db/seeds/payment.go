package seeds

import (
	"tubespbbo/db"
	"tubespbbo/modules/model"
)

func (s *Seed) SeedPayment() {
	var pm []model.Payment
	var count int64
	db.Orm.Model(&pm).Count(&count)

	if count > 0 {
		return
	}

	payments := make([]model.Payment, 4)
	payments[0] = model.Payment{
		TransactionId:     1,
		PaymentMethodId: 1,
		Status:          "Lunas",
		AccountNumber:      "987654321123",
		Amount:          50000,
	}
	payments[1] = model.Payment{
		TransactionId:     2,
		PaymentMethodId: 2,
		Status:          "Lunas",
		AccountNumber:      "085278633421",
		Amount:          40000,
	}
	payments[2] = model.Payment{
		TransactionId:     3,
		PaymentMethodId: 3,
		Status:          "Lunas",
		AccountNumber:      "27852684",
		Amount:          2100000,
	}
	payments[3] = model.Payment{
		TransactionId:     5,
		PaymentMethodId: 3,
		Status:          "Lunas",
		AccountNumber:      "2785268443",
		Amount:          300000,
	}

	_ = db.Orm.Create(&payments)
}
