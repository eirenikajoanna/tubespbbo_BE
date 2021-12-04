package dto

type CreateTransactionDetailDTO struct {
	ProductId int64 `json:"productId"`
	Quantity  int64 `json:"quantity"`
}
