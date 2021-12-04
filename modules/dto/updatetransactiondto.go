package dto

type UpdateTransactionDTO struct {
	UserId             int64                        `json:"userId"`
	ReceiptNumber      string                       `json:"receiptNumber"`
	Status             string                       `json:"status"`
	TransactionDetails []UpdateTransactionDetailDTO `json:"details"`
}
