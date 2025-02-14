package infrastracture

import (
	"database/sql"
	"pasour/internal/domain/errors"

	_ "github.com/mattn/go-sqlite3"
)

func NewDB() (*sql.DB, *errors.DomainErr) {
	db, err := sql.Open("sqlite3", Configs.RootDir+"/pasour.db")
	if err != nil {
		return nil, errors.NewInternalErr(err)
	}

	if err := db.Ping(); err != nil {
		return nil, errors.NewInternalErr(err)
	}

	return db, nil
}
