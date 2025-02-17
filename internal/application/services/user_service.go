package services

import (
	"pasour/internal/domain/commands"
	"pasour/internal/domain/dtos"
	"pasour/internal/domain/entities"
	"pasour/internal/domain/errors"
	"pasour/internal/domain/repos"
	"pasour/internal/domain/types"
)

type UserService struct {
	repo repos.UserRepo
}

func NewUserService(repo repos.UserRepo) *UserService {
	return &UserService{repo: repo}
}

func (userService *UserService) FindByUsername(username string) (*dtos.UserDTO, *errors.DomainErr) {
	user, err := userService.repo.FindByUsername(username)
	if err != nil {
		return nil, err
	}

	return dtos.NewUserDTO(user), nil
}

func (userService *UserService) Authenticate(cmd *commands.AuthCmd) (*dtos.UserDTO, *errors.DomainErr) {
	user, err := userService.repo.FindByUsername(cmd.Username)
	if err != nil {
		return nil, err
	}

	if !user.CheckPassword(cmd.Password) {
		return nil, errors.NewUnAuthorizedErr("authentication failed")
	}

	if cmd.Role == types.RoleAdmin && !user.IsAdmin {
		return nil, errors.NewUnAuthorizedErr("authentication failed")
	}

	return dtos.NewUserDTO(user), nil
}

func (userService *UserService) SignUp(cmd *commands.UserSignUpCmd) (*dtos.UserDTO, *errors.DomainErr) {
	existingUser, err := userService.repo.FindByUsername(cmd.Username)

	if err != nil && err.Type != types.NotFoundErr {
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
