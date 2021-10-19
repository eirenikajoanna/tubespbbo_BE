package repository

import (
	"errors"
	"tubespbbo/db"
	"tubespbbo/domains/expense/model"
)

func FindExpense() (*[]model.Expense, error) {
	var expenses []model.Expense
	result := db.Orm.Find(&pexpenses)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("no expenses found")
	}

	return &expenses, nil
}

func FindOneExpense(id int64) (*model.Expense, error) {
	var expenses model.Expense
	result := db.Orm.First(&expenses, id)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("no Expenses found")
	}

	return &expenses, nil
}

func CreateExpense(expenses *model.Expense) error {
	result := db.Orm.Create(expenses)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func UpdateExpense(expenses *model.Expense) (*model.Expense, error) {
	result := db.Orm.Save(&expenses)
	if result.Error != nil {
		return nil, result.Error
	}
	return expenses, nil
}

func DeleteExpense(expenses *model.PExpense) (*model.Expense, error) {
	if expenses.Payment != nil {
		return nil, errors.New("can't delete cause of relational")
	}
	result := db.Orm.Delete(&expenses)
	if result.Error != nil {
		return nil, result.Error
	}

	return expenses, nil
}
