package dto

import "time"

type ReportDTO struct {
	Id        int64     `json:"id"`
	Date      string    `json:"date"`
	Type      string    `json:"type"`
	Amount    float32   `json:"amount"`
	CreatedAt time.Time `json:"createdAt"`
}
