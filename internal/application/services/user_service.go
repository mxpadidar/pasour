package services

import (
	"fmt"
	"pasour/internal/domain/commands"
	"pasour/internal/domain/dtos"
	"pasour/internal/domain/entities"
	"pasour/internal/domain/repos"
	"pasour/internal/domain/types"
	"strings"
)

type UserService struct {
	repo repos.UserRepo
}

func NewUserService(repo repos.UserRepo) *UserService {
	return &UserService{repo: repo}
}

func (userService *UserService) FindByUsername(username string) (*dtos.UserDTO, error) {
	user, err := userService.repo.FindByUsername(username)
	if err != nil {
		return nil, err
	}

	return dtos.NewUserDTO(user), nil
}

func (userService *UserService) Authenticate(cmd *commands.AuthCmd) (*dtos.UserDTO, error) {
	user, err := userService.repo.FindByUsername(cmd.Username)
	if err != nil {
		return nil, err
	}

	if !user.CheckPassword(cmd.Password) {
		return nil, fmt.Errorf("authentication failed")
	}

	if cmd.Role == types.RoleAdmin && !user.IsAdmin {
		return nil, fmt.Errorf("authentication failed")
	}

	return dtos.NewUserDTO(user), nil
}

func (userService *UserService) SignUp(cmd *commands.UserSignUpCmd) (*dtos.UserDTO, error) {
	existingUser, err := userService.repo.FindByUsername(cmd.Username)

	if err != nil {
		// Skip not found errors
		if strings.Contains(err.Error(), "not found") {
			goto ContinueSignup
		}
		return nil, err
	}

	if existingUser != nil {
		return nil, fmt.Errorf("username already exists")
	}

ContinueSignup:
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
