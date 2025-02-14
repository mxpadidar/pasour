package entities

import (
	"pasour/internal/domain/errors"
	"time"
)

type User struct {
	ID             int
	Username       string
	HashedPassword string
	IsAdmin        bool
	CreatedAt      time.Time
}

func NewUser(username, password string, isAdmin bool) (*User, *errors.DomainErr) {
	// create user
	user := &User{}
	if err := user.SetUsername(username); err != nil {
		return nil, err
	}
	if err := user.SetPassword(password); err != nil {
		return nil, err
	}
	user.IsAdmin = isAdmin
	user.CreatedAt = time.Now()
	return user, nil
}

func (u *User) SetPassword(rawPassword string) *errors.DomainErr {
	// validate password length
	if len(rawPassword) < 8 {
		return errors.NewValidationErr("password must be at least 8 characters")
	}

	// hash password
	hashPassword := rawPassword + "hashed"
	u.HashedPassword = hashPassword
	return nil
}

func (u *User) SetUsername(username string) *errors.DomainErr {
	// validate username length
	if len(username) < 4 {
		return errors.NewValidationErr("username must be at least 4 characters")
	}

	u.Username = username
	return nil
}

func (u *User) CheckPassword(password string) bool {
	// hash password
	hashPassword := password + "hashed"
	return u.HashedPassword == hashPassword
}
