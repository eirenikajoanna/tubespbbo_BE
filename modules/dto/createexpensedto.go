package dto

type CreateExpenseDTO struct {
	UserId		int64	`json:"userId"`
	Name        string  `json:"name" validate:"empty=false"`
	ReleaseDate string  `json:"releaseDate" validate:"empty=false"`
	Quantity    float32 `json:"quantity"`
	Category    int64   `json:"category"`
	Description string  `json:"description"`
	TotalAmount	float32 `json:"totalAmount"`
}
