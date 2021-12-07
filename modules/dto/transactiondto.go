package dto

import (
	"tubespbbo/base"
)

type TransactionDTO struct {
	base.DTO
	UserId             int64                   `json:"userId"`
	User               *UserDTO                `json:"user"`
	PaymentId          int64                   `json:"paymentId"`
	Payment            *PaymentDTO             `json:"payment"`
	ReceiptNumber      string                  `json:"receiptNumber"`
	Status             string                  `json:"status"`
	TransactionDetails []*TransactionDetailDTO `json:"transactionDetails"`
}
