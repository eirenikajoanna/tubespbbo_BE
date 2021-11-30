package service

import (
	"tubespbbo/modules/dto"
	"tubespbbo/modules/model"
	"tubespbbo/modules/repository"
)

func FindExpense() (*[]model.Expense, error) {
	return repository.FindExpense()
}

func FindOneExpense(id int64) (*model.Expense, error) {
	return repository.FindOneExpense(id)
}

func CreateExpense(expense *model.Expense) error {
	return repository.CreateExpense(expense)
}

func UpdateExpense(updateDto *dto.UpdateExpenseDTO, id int64) (*model.Expense, error) {
	pm, err := repository.FindOneExpense(id)
	if err != nil {
		return nil, err
	}
	if updateDto.UserId != 0 {
		pm.UserId = updateDto.UserId
	}
	if updateDto.Name != "" {
		pm.Name = updateDto.Name
	}
	if updateDto.ReleaseDate != "" {
		pm.ReleaseDate = updateDto.ReleaseDate
	}
	if updateDto.Quantity != 0 {
		pm.Quantity = updateDto.Quantity
	}
	if updateDto.Category != 0 {
		pm.Category = updateDto.Category
	}
	if updateDto.Description != "" {
		pm.Description = updateDto.Description
	}
	if updateDto.TotalAmount != 0 {
		pm.TotalAmount = updateDto.TotalAmount
	}
	pm, err = repository.UpdateExpense(pm)
	if err != nil {
		return nil, err
	}
	return pm, nil
}

func DeleteExpense(id int64) (*model.Expense, error) {
	expense, err := repository.FindOneExpense(id)
	if err != nil {
		return nil, err
	}
	return repository.DeleteExpense(expense)
}
