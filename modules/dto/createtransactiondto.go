package dto

type CreateTransactionDTO struct {
	UserId             int64                        `json:"userId"`
	TransactionDetails []CreateTransactionDetailDTO `json:"details"`
}
