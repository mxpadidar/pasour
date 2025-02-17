package services

import (
	"pasour/internal/domain/commands"
	"pasour/internal/domain/dtos"
	"pasour/internal/domain/errors"
)

type UserService interface {
	FindByUsername(string) (*dtos.UserDTO, *errors.DomainErr)
	Authenticate(*commands.AuthCmd) (*dtos.UserDTO, *errors.DomainErr)
	SignUp(*commands.UserSignUpCmd) (*dtos.UserDTO, *errors.DomainErr)
}

type TokenService interface {
	Encode(*commands.TokenEncodeCmd) (*dtos.TokenDTO, *errors.DomainErr)
	Decode(string) (string, *errors.DomainErr)
	GetTokenFromHeader(string) (string, *errors.DomainErr)
}
