package service

import (
	"tubespbbo/modules/dto"
	"tubespbbo/modules/model"
	"tubespbbo/modules/repository"
)

func FindPayment() (*[]model.Payment, error) {
	return repository.FindPayment()
}

func FindOnePayment(id int64) (*model.Payment, error) {
	return repository.FindOnePayment(id)
}

func CreatePayment(payment *model.Payment) error {
	transaction, _ := repository.FindOneTransaction(payment.TransactionId)
	transaction.Status = "Belum Proses"
	repository.UpdateTransaction(transaction)
	return repository.CreatePayment(payment)
}

func UpdatePayment(updateDto *dto.UpdatePaymentDTO, id int64) (*model.Payment, error) {
	pm, err := repository.FindOnePayment(id)
	if err != nil {
		return nil, err
	}
	if updateDto.AccountNumber != "" {
		pm.AccountNumber = updateDto.AccountNumber
	}
	if updateDto.Amount != 0 {
		pm.Amount = updateDto.Amount
	}
	if updateDto.PaymentMethodId != 0 {
		pm.PaymentMethodId = updateDto.PaymentMethodId
	}
	pm, err = repository.UpdatePayment(pm)
	if err != nil {
		return nil, err
	}
	return pm, nil
}

func DeletePayment(id int64) (*model.Payment, error) {
	payment, err := repository.FindOnePayment(id)
	if err != nil {
		return nil, err
	}
	return repository.DeletePayment(payment)
}
