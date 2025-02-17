package commands

import (
	"pasour/internal/domain/types"
)

type AuthCmd struct {
	Username string         `json:"username" validate:"required,min=6"`
	Password string         `json:"password" validate:"required,min=8"`
	Role     types.UserRole `json:"role" validate:"required"`
}

type UserSignUpCmd struct {
	Username string `json:"username" validate:"required,min=6"`
	Password string `json:"password" validate:"required,min=8"`
}

type TokenEncodeCmd struct {
	Sub string
}
