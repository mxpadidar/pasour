package main

import (
	"log"
	"os"
	"pasour/internal/infrastracture"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	baseDir := infrastracture.Configs.RootDir
	db, domainErr := infrastracture.NewDB()

	if domainErr != nil {
		log.Fatal(domainErr)
	}

	defer db.Close()

	// Create SQLite driver instance
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		log.Fatalf("failed to create driver instance: %v", err)
	}

	// Create migration instance
	mgrtDir := "file://" + baseDir + "/migrations/sql"
	mgrt, err := migrate.NewWithDatabaseInstance(mgrtDir, "sqlite3", driver)
	if err != nil {
		log.Fatalf("failed to create migration instance: %v", err)
	}

	// get the command from the arguments
	cmd := os.Args[len(os.Args)-1]
	switch cmd {
	case "up":
		err := mgrt.Up()
		if err != nil && err != migrate.ErrNoChange {
			log.Fatalf("failed to apply migration: %v", err)
		} else if err != nil && err == migrate.ErrNoChange {
			log.Println("db is up to date")
		} else {
			log.Println("migration applied successfully")
		}

	case "down":
		if err := mgrt.Down(); err != nil {
			log.Fatalf("failed to rollback migration: %v", err)
		} else {
			log.Println("migration rolled back successfully")
		}

	default:
		log.Fatalf("unknown command: %s", cmd)
	}

}
