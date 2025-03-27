package migrations

import (
	"context"
	"fmt"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
)

var Migrations = migrate.NewMigrations()

// Register migration
func init() {
	Migrations.MustRegister(
		// UP Migration
		func(ctx context.Context, db *bun.DB) error {
			fmt.Println("Applying migration: Adding a new table")
			_, err := db.Exec(`CREATE TABLE users (id SERIAL PRIMARY KEY, name TEXT)`)
			return err
		},
		// DOWN Migration
		func(ctx context.Context, db *bun.DB) error {
			fmt.Println("Rolling back migration: Removing users table")
			_, err := db.Exec(`DROP TABLE users`)
			return err
		},
	)
}
