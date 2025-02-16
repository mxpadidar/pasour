package services

import (
	"pasour/internal/domain/commands"
	"pasour/internal/domain/dtos"
	"pasour/internal/domain/errors"
)

type UserService interface {
	FindByUsername(string) (*dtos.UserDTO, *errors.DomainErr)
	SignUp(*commands.UserSignUpCmd) (*dtos.UserDTO, *errors.DomainErr)
}

type TokenService interface {
	// generate a jwt based on the given command
	Encode(*commands.TokenEncodeCmd) (string, *errors.DomainErr)

	// decode a jwt token and return the subject of the token
	Decode(string) (string, *errors.DomainErr)

	// get the token from authorization header
	GetTokenFromHeader(string) (string, *errors.DomainErr)
}
