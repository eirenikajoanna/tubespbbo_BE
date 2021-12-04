package repository

import (
	"errors"
	"tubespbbo/db"
	"tubespbbo/modules/model"
)

func FindPayment() (*[]model.Payment, error) {
	var payments []model.Payment
	result := db.Orm.Find(&payments).Preload("PaymentMethod").Find(&payments).Not("deleted_at = ?", nil)
	if result.Error != nil {
		return nil, result.Error
	}

	return &payments, nil
}

func FindOnePayment(id int64) (*model.Payment, error) {
	var payment model.Payment
	result := db.Orm.First(&payment, id).Preload("PaymentMethod").First(&payment, id)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("no payment found")
	}

	return &payment, nil
}

func CreatePayment(pm *model.Payment) error {
	result := db.Orm.Create(pm)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func UpdatePayment(pm *model.Payment) (*model.Payment, error) {
	result := db.Orm.Save(&pm)
	if result.Error != nil {
		return nil, result.Error
	}
	pm, _ = FindOnePayment(pm.Id)
	return pm, nil
}

func DeletePayment(pm *model.Payment) (*model.Payment, error) {
	result := db.Orm.Delete(&pm)
	if result.Error != nil {
		return nil, result.Error
	}
	return pm, nil
}
