package dto

type CreatePaymentMethodDTO struct {
	Name string `json:"name" validate:"empty=false"`
	Code string `json:"code" validate:"empty=false"`
}
