package dto

type CreateUserDTO struct {
	Username string `json:"username" validate:"empty=false"`
	Password string `json:"password" validate:"empty=false"`
	Name     string `json:"name"`
	Dob      string `json:"dob"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	Email    string `json:"email" validate:"empty=false"`
}
