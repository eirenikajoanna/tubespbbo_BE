package repository

import (
	"errors"
	"fmt"
	"time"
	"tubespbbo/db"
	"tubespbbo/modules/model"
)

func FindTransaction() (*[]model.Transaction, error) {
	var transactions []model.Transaction
	result := db.Orm.Find(&transactions).Preload("Payment").Preload("Payment.PaymentMethod").Preload("TransactionDetails").Preload("TransactionDetails.Product").Find(&transactions).Not("deleted_at = ?", nil)
	if result.Error != nil {
		return nil, result.Error
	}
	return &transactions, nil
}

func FindTransactionByUser(userId int64) (*[]model.Transaction, error) {
	var transactions []model.Transaction
	result := db.Orm.Find(&transactions).Where("user_id = ?", userId).Preload("Payment").Preload("Payment.PaymentMethod").Preload("TransactionDetails").Preload("TransactionDetails.Product").Find(&transactions).Not("deleted_at = ?", nil)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("no transaction found")
	}

	return &transactions, nil
}

func FindTransactionOneMonth(month int, year int) (*[]model.Transaction, error) {
	now := time.Now()

	lastMonth := month
	lastYear := year
	if month+1 == 13 {
		lastYear += 1
		lastMonth = 1
	}
	endOfLastMonth := time.Date(lastYear, time.Month(lastMonth), 1, 0, 0, 0, 0, now.Location())
	fmt.Printf("End of the last month: %s\n", endOfLastMonth)
	firstDayOfThisMonth := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, now.Location())
	fmt.Printf("The fist day of the actual month: %s\n", firstDayOfThisMonth)
	var transactions []model.Transaction
	result := db.Orm.Find(&transactions).Where("created_at BETWEEN ? AND ?", firstDayOfThisMonth, endOfLastMonth).Preload("Payment").Preload("Payment.PaymentMethod").Preload("TransactionDetails").Preload("TransactionDetails.Product").Find(&transactions).Not("deleted_at = ?", nil)
	if result.Error != nil {
		return nil, result.Error
	}

	return &transactions, nil
}

func FindOneTransaction(id int64) (*model.Transaction, error) {
	var transaction model.Transaction
	result := db.Orm.First(&transaction, id).Preload("Payment").Preload("Payment.PaymentMethod").Preload("TransactionDetails").Preload("TransactionDetails.Product").First(&transaction, id)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("no transaction found")
	}

	return &transaction, nil
}

func CreateTransaction(pm *model.Transaction) error {
	result := db.Orm.Create(&pm)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func UpdateTransaction(pm *model.Transaction) (*model.Transaction, error) {
	result := db.Orm.Save(&pm)
	if result.Error != nil {
		return nil, result.Error
	}
	pm, _ = FindOneTransaction(pm.Id)
	return pm, nil
}

func DeleteTransaction(pm *model.Transaction) (*model.Transaction, error) {
	result := db.Orm.Delete(&pm)
	if result.Error != nil {
		return nil, result.Error
	}
	return pm, nil
}
