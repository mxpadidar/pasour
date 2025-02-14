package services

import (
	"pasour/internal/domain/commands"
	"pasour/internal/domain/dtos"
	"pasour/internal/domain/entities"
	"pasour/internal/domain/errors"
	"pasour/internal/domain/repos"
)

type UserService struct {
	repo repos.UserRepo
}

func NewUserService(repo repos.UserRepo) *UserService {
	return &UserService{repo: repo}
}

func (userService *UserService) SignUp(cmd *commands.UserSignUpCommand) (*dtos.UserDTO, *errors.DomainErr) {
	existingUser, err := userService.repo.FindByUsername(cmd.Username)

	if err != nil && err.Type != errors.NotFoundErr {
		return nil, err
	}

	if existingUser != nil {
		return nil, errors.NewConflictErr("username already exists")
	}

	user, err := entities.NewUser(cmd.Username, cmd.Password, false)
	if err != nil {
		return nil, err
	}
	println("saving user")
	if err := userService.repo.Save(user); err != nil {
		return nil, err
	}

	return dtos.NewUserDTO(user), nil
}
