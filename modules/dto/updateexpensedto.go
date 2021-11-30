package dto

type UpdateExpenseDTO struct {
	UserId      int64   `json:"userId"`
	Name        string  `json:"name"`
	ReleaseDate string  `json:"releaseDate"`
	Quantity    float32 `json:"quantity"`
	Category    int64   `json:"category"`
	Description string  `json:"description"`
	TotalAmount	float32 `json:"totalAmount"`
}
