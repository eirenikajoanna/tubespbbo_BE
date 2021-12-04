package repository

import (
	"errors"
	"fmt"
	"time"
	"tubespbbo/db"
	"tubespbbo/modules/model"
)

func FindExpense() (*[]model.Expense, error) {
	var expenses []model.Expense
	result := db.Orm.Find(&expenses).Preload("User").Find(&expenses).Not("deleted_at = ?", nil)
	if result.Error != nil {
		return nil, result.Error
	}

	return &expenses, nil
}

func FindExpenseOneMonth(month int, year int) (*[]model.Expense, error) {
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
	var expenses []model.Expense
	result := db.Orm.Where("created_at BETWEEN ? AND ?", firstDayOfThisMonth, endOfLastMonth).Find(&expenses).Preload("User").Find(&expenses).Not("deleted_at = ?", nil)
	if result.Error != nil {
		return nil, result.Error
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

func DeleteExpense(expenses *model.Expense) (*model.Expense, error) {
	result := db.Orm.Delete(&expenses)
	if result.Error != nil {
		return nil, result.Error
	}

	return expenses, nil
}
