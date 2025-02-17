package services

import (
	"pasour/internal/domain/commands"
	"pasour/internal/domain/dtos"
)

type UserService interface {
	FindByUsername(string) (*dtos.UserDTO, error)
	Authenticate(*commands.AuthCmd) (*dtos.UserDTO, error)
	SignUp(*commands.UserSignUpCmd) (*dtos.UserDTO, error)
}

type TokenService interface {
	Encode(*commands.TokenEncodeCmd) (*dtos.TokenDTO, error)
	Decode(string) (string, error)
	GetTokenFromHeader(string) (string, error)
}
