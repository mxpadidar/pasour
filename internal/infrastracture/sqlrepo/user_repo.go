package sqlrepo

import (
	"database/sql"
	"fmt"
	"pasour/internal/domain/entities"
)

type SqlUserRepo struct {
	db *sql.DB
}

func NewSqlUserRepo(db *sql.DB) *SqlUserRepo {
	return &SqlUserRepo{db}
}

func (repo SqlUserRepo) FindByUsername(username string) (*entities.User, error) {
	query := `SELECT id, username, hashed_password, is_admin, created_at FROM users WHERE username = $1`
	row := repo.db.QueryRow(query, username)

	user, err := scan(row)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user not found")
	}
	return user, err
}

func (store SqlUserRepo) Save(user *entities.User) error {
	query := `INSERT INTO users (username, hashed_password, is_admin, created_at)
		VAlUES ($1, $2, $3, $4) RETURNING id`
	if err := store.db.QueryRow(
		query,
		user.Username,
		user.HashedPassword,
		user.IsAdmin,
		user.CreatedAt,
	).Scan(&user.ID); err != nil {
		return fmt.Errorf("validation error: %w", err)
	} else {
		return nil
	}
}

func scan(row *sql.Row) (*entities.User, error) {
	var user entities.User
	err := row.Scan(&user.ID, &user.Username, &user.HashedPassword, &user.IsAdmin, &user.CreatedAt)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user not found")
	} else if err != nil {
		return nil, fmt.Errorf("internal error: %w", err)
	}
	return &user, nil
}
