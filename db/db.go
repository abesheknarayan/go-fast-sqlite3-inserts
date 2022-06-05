package db

import (
	"database/sql"
	"embed"
	"fmt"
	"net/http"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/httpfs"
)

//go:embed migrations
var migrations embed.FS

const schemaVersion = 1

func NewDB(dbPath string) (*sql.DB, error) {
	sqliteDB, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	return sqliteDB, nil
}

func RunMigrationScripts(db *sql.DB) error {

	// reads migrations from vfs by making a file server with go:embed (TODO: know y it is this way or any alternatives present)
	sourceInstance, err := httpfs.New(http.FS(migrations), "migrations")

	if err != nil {
		return fmt.Errorf("error while creating migrations instance %s", err)
	}

	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})

	if err != nil {
		return fmt.Errorf("creating sqlite3 driver failed %s", err)
	}
	m, err := migrate.NewWithInstance("https", sourceInstance, "sqlite3", driver)
	if err != nil {
		return fmt.Errorf("failed to initialize migrate instance, %w", err)
	}

	// up migrations
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return sourceInstance.Close()

}
