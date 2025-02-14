package services

import (
	"pasour/internal/domain/commands"
	"pasour/internal/domain/dtos"
	"pasour/internal/domain/errors"
)

type UserService interface {
	SignUp(cmd *commands.UserSignUpCommand) (*dtos.UserDTO, *errors.DomainErr)
}
