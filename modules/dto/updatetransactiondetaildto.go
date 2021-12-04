package dto

type UpdateTransactionDetailDTO struct {
	Id        int64 `json:"id"`
	ProductId int64 `json:"productId"`
	Quantity  int64 `json:"quantity"`
}
