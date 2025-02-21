package migrate

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/migrate"
	"log"
)

func RunMigrations(dsn string) {
	sqldb := pgdriver.NewConnector(pgdriver.WithDSN(dsn))
	db := bun.NewDB(sqldb, pgdriver.New())

	migrator := migrate.NewMigrator(db, migrate.NewMigrations())

	if err := migrator.Init(context.Background()); err != nil {
		log.Fatal(err)
	}
	if err := migrator.Up(context.Background()); err != nil {
		log.Fatal(err)
	}

	log.Println("✅ Migrations applied successfully!")
}
