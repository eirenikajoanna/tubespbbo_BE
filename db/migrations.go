package db

import (
	"tubespbbo/modules/model"
)

func migrate() {
	_ = Orm.AutoMigrate(new(model.User))
	_ = Orm.AutoMigrate(new(model.PaymentMethod))
	_ = Orm.AutoMigrate(new(model.Payment))
	_ = Orm.AutoMigrate(new(model.Product))
	_ = Orm.AutoMigrate(new(model.Expense))
	_ = Orm.AutoMigrate(new(model.Transaction))
	_ = Orm.AutoMigrate(new(model.TransactionDetail))
	_ = Orm.AutoMigrate(new(model.ShoppingCart))
}
