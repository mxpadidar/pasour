package dtos

import "pasour/internal/domain/entities"

type UserDTO struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	IsAdmin   bool   `json:"is_admin"`
	CreatedAt string `json:"created_at"`
}

func NewUserDTO(user *entities.User) *UserDTO {
	return &UserDTO{
		ID:        user.ID,
		Username:  user.Username,
		IsAdmin:   user.IsAdmin,
		CreatedAt: user.CreatedAt.String(),
	}
}
