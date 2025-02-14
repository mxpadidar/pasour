package sqlrepo

import (
	"database/sql"
	"pasour/internal/domain/entities"
	"pasour/internal/domain/errors"
)

type SqlUserRepo struct {
	db *sql.DB
}

func NewSqlUserRepo(db *sql.DB) *SqlUserRepo {
	return &SqlUserRepo{db}
}

func (repo SqlUserRepo) FindByUsername(username string) (*entities.User, *errors.DomainErr) {
	query := `SELECT id, username, hashed_password, is_admin, created_at FROM users WHERE username = $1`
	row := repo.db.QueryRow(query, username)

	user, err := scan(row)
	if err == sql.ErrNoRows {
		return nil, errors.NewNotFoundErr("user not found")
	}
	return user, err
}

func (store SqlUserRepo) Save(user *entities.User) *errors.DomainErr {
	query := `INSERT INTO users (username, hashed_password, is_admin, created_at)
		VAlUES ($1, $2, $3, $4) RETURNING id`
	if err := store.db.QueryRow(
		query,
		user.Username,
		user.HashedPassword,
		user.IsAdmin,
		user.CreatedAt,
	).Scan(&user.ID); err != nil {
		return errors.NewValidationErr(err)
	} else {
		return nil
	}
}

func scan(row *sql.Row) (*entities.User, *errors.DomainErr) {
	var user entities.User
	err := row.Scan(&user.ID, &user.Username, &user.HashedPassword, &user.IsAdmin, &user.CreatedAt)

	if err == sql.ErrNoRows {
		return nil, errors.NewNotFoundErr("user not found")
	} else if err != nil {
		return nil, errors.NewInternalErr(err)
	}
	return &user, nil
}
