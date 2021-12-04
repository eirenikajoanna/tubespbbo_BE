package repository

import (
	"errors"
	"tubespbbo/db"
	"tubespbbo/modules/model"
)

func FindTransactionDetail() (*[]model.TransactionDetail, error) {
	var transactionDetails []model.TransactionDetail
	result := db.Orm.Find(&transactionDetails)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("no payment methods found")
	}

	return &transactionDetails, nil
}

func FindOneTransactionDetail(id int64) (*model.TransactionDetail, error) {
	var pm model.TransactionDetail
	result := db.Orm.First(&pm, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &pm, nil
}

func CreateTransactionDetail(pm *model.TransactionDetail) error {
	result := db.Orm.Create(&pm)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func UpdateTransactionDetail(pm *model.TransactionDetail) (*model.TransactionDetail, error) {
	result := db.Orm.Save(&pm)
	if result.Error != nil {
		return nil, result.Error
	}
	return pm, nil
}

func DeleteTransactionDetail(pm *model.TransactionDetail) (*model.TransactionDetail, error) {
	result := db.Orm.Delete(&pm)
	if result.Error != nil {
		return nil, result.Error
	}

	return pm, nil
}
