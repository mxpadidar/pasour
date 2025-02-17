package entities

import (
	"errors"
	"time"
)

type User struct {
	ID             int
	Username       string
	HashedPassword string
	IsAdmin        bool
	CreatedAt      time.Time
}

func NewUser(username, password string, isAdmin bool) (*User, error) {
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

func (u *User) SetPassword(rawPassword string) error {
	// validate password length
	if len(rawPassword) < 8 {
		return errors.New("invalid password")
	}

	// hash password
	hashPassword := rawPassword + "hashed"
	u.HashedPassword = hashPassword
	return nil
}

func (u *User) SetUsername(username string) error {
	// validate username length
	if len(username) < 4 {
		return errors.New("invalid username")
	}

	u.Username = username
	return nil
}

func (u *User) CheckPassword(password string) bool {
	// hash password
	hashPassword := password + "hashed"
	return u.HashedPassword == hashPassword
}
