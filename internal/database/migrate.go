package database

import (
	"errors"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"zuqui-core/internal"
)

// RunMigrations applies database migrations to the specified database
func (db *Database) RunMigrations() error {
	m, err := migrate.New("file://internal/database/migrations", internal.Env.DATABASE_URL)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}
	srcErr, dbErr := m.Close()
	if srcErr != nil {
		return srcErr
	}
	if dbErr != nil {
		return dbErr
	}
	log.Println("Migrations applied successfully")
	return nil
}
