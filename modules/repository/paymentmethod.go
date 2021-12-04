package repository

import (
	"errors"
	"tubespbbo/db"
	"tubespbbo/modules/model"
)

func FindPaymentMethod() (*[]model.PaymentMethod, error) {
	var paymentMethods []model.PaymentMethod
	result := db.Orm.Find(&paymentMethods)
	if result.Error != nil {
		return nil, result.Error
	}

	return &paymentMethods, nil
}

func FindOnePaymentMethod(id int64) (*model.PaymentMethod, error) {
	var pm model.PaymentMethod
	result := db.Orm.First(&pm, id)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("no payment methods found")
	}

	return &pm, nil
}

func CreatePaymentMethod(pm *model.PaymentMethod) error {
	result := db.Orm.Create(pm)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func UpdatePaymentMethod(pm *model.PaymentMethod) (*model.PaymentMethod, error) {
	result := db.Orm.Save(&pm)
	if result.Error != nil {
		return nil, result.Error
	}
	return pm, nil
}

func DeletePaymentMethod(pm *model.PaymentMethod) (*model.PaymentMethod, error) {
	if pm.Payment != nil {
		return nil, errors.New("can't delete cause of relational")
	}
	result := db.Orm.Delete(&pm)
	if result.Error != nil {
		return nil, result.Error
	}

	return pm, nil
}
