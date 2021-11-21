package model

import "errors"

type Role int

const (
	CUSTOMER Role = iota
	PENANAM_BUNGA
	KURIR
	ADMIN
	OTHERS
)

func (r Role) String() string {
	return [...]string{"Customer", "Penanam Bunga", "Kurir", "Admin"}[r]
}

func (r Role) EnumIndex() int {
	return int(r)
}

func RoleFromString(s string) (Role, error) {
	switch s {
	case "Customer":
		return CUSTOMER, nil
	case "Penanam Bunga":
		return PENANAM_BUNGA, nil
	case "Kurir":
		return KURIR, nil
	case "Admin":
		return ADMIN, nil
	default:
		return OTHERS, errors.New("Role not supported")
	}
}
