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
		product, err := repository.FindOneProduct(createDto.TransactionDetails[i].ProductId)
		if err != nil {
			return nil, err
		}
		product.Quantity = product.Quantity - createDto.TransactionDetails[i].Quantity
		repository.UpdateProduct(product)
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
	for i := 0; i < len(transaction.TransactionDetails); i++ {
		transactionDetail := transaction.TransactionDetails[i]
		product, err := repository.FindOneProduct(transactionDetail.ProductId)
		if err != nil {
			return nil, err
		}
		product.Quantity = product.Quantity + transactionDetail.Quantity
		repository.UpdateProduct(product)
		repository.DeleteTransactionDetail(&transactionDetail)
	}
	return repository.DeleteTransaction(transaction)
}
