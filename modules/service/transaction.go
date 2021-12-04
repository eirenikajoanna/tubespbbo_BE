package service

import (
	"tubespbbo/mapper"
	"tubespbbo/modules/dto"
	"tubespbbo/modules/model"
	"tubespbbo/modules/repository"
)

func FindTransaction() (*[]model.Transaction, error) {
	return repository.FindTransaction()
}

func FindTransactionByUser(userId int64) (*[]model.Transaction, error) {
	return repository.FindTransactionByUser(userId)
}

func FindOneTransaction(id int64) (*model.Transaction, error) {
	return repository.FindOneTransaction(id)
}

func CreateTransaction(createDto *dto.CreateTransactionDTO) (*model.Transaction, error) {
	var transaction model.Transaction
	mapper.Map(createDto, &transaction)
	transaction.Status = "Belum Terverifikasi"
	repository.CreateTransaction(&transaction)
	for i := 0; i < len(createDto.TransactionDetails); i++ {
		var transactionDetail model.TransactionDetail
		mapper.Map(createDto.TransactionDetails[i], &transactionDetail)
		transactionDetail.TransactionId = transaction.Id
		repository.CreateTransactionDetail(&transactionDetail)
		transaction.TransactionDetails = append(transaction.TransactionDetails, transactionDetail)
	}
	return &transaction, nil
}

func UpdateTransaction(updateDto *dto.UpdateTransactionDTO, id int64) (*model.Transaction, error) {
	pm, err := repository.FindOneTransaction(id)
	if err != nil {
		return nil, err
	}
	if updateDto.UserId != 0 {
		pm.UserId = updateDto.UserId
	}
	if updateDto.ReceiptNumber != "" {
		pm.ReceiptNumber = updateDto.ReceiptNumber
	}
	if updateDto.Status != "" {
		pm.Status = updateDto.Status
	}
	pm, err = repository.UpdateTransaction(pm)

	var transactionDetails []model.TransactionDetail
	for i := 0; i < len(updateDto.TransactionDetails); i++ {
		if updateDto.TransactionDetails[i].Id != 0 {
			pm, err := UpdateTransactionDetail(&updateDto.TransactionDetails[i], updateDto.TransactionDetails[i].Id)
			if err != nil {
				return nil, err
			}
			transactionDetails = append(transactionDetails, *pm)
		}
		if updateDto.TransactionDetails[i].Id == 0 {
			var transactionDetail model.TransactionDetail
			mapper.Map(updateDto.TransactionDetails[i], &transactionDetail)
			transactionDetail.TransactionId = pm.Id
			repository.CreateTransactionDetail(&transactionDetail)
			transactionDetails = append(transactionDetails, transactionDetail)
		}
	}

	for i := 0; i < len(pm.TransactionDetails); i++ {
		if !contains(updateDto.TransactionDetails, pm.TransactionDetails[i].Id) {
			DeleteTransactionDetail(pm.TransactionDetails[i].Id)
		}
	}

	pm.TransactionDetails = transactionDetails

	if err != nil {
		return nil, err
	}
	return pm, nil
}

func DeleteTransaction(id int64) (*model.Transaction, error) {
	transaction, err := repository.FindOneTransaction(id)
	if err != nil {
		return nil, err
	}
	return repository.DeleteTransaction(transaction)
}

func contains(s []dto.UpdateTransactionDetailDTO, str int64) bool {
	for _, v := range s {
		if v.Id == str {
			return true
		}
	}

	return false
}
