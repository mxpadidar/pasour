package repos

import (
	"pasour/internal/domain/entities"
)

type UserRepo interface {
	FindByUsername(username string) (*entities.User, error)
	Save(user *entities.User) error
}
