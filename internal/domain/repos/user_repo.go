package repos

import (
	"pasour/internal/domain/entities"
	"pasour/internal/domain/errors"
)

type UserRepo interface {
	FindByUsername(username string) (*entities.User, *errors.DomainErr)
	Save(user *entities.User) *errors.DomainErr
}
