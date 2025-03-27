package migrate

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/kalyan-velu/weetrival-localize/cmd/migrate/migrations"
	"github.com/kalyan-velu/weetrival-localize/internal/models"

	_ "github.com/lib/pq"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/migrate"
)

// RunMigrations applies pending database migrations.
func RunMigrations(dsn string) {
	// Create a database connection
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	// Ensure the database connection is working
	if err := sqldb.Ping(); err != nil {
		log.Fatalf("❌ Database connection failed: %v", err)
	}

	// Initialize Bun ORM
	db := bun.NewDB(sqldb, pgdialect.New())
	defer func(db *bun.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("<UNK> Database connection failed: %v", err)
		}
	}(db) // Ensure cleanup

	modelsToMigrate := []interface{}{
		(*models.User)(nil),
		(*models.AuthProvider)(nil),
		(*models.UserResponse)(nil),
		(*models.RefreshToken)(nil),
		(*models.RolePermission)(nil),
	}

	for _, model := range modelsToMigrate {
		_, err := db.NewCreateTable().Model(model).IfNotExists().Exec(context.Background())
		if err != nil {
			log.Fatalf("❌ Failed to migrate %T: %v", model, err)
		}
	}

	// Initialize the migrator
	migrator := migrate.NewMigrator(db, migrations.Migrations)

	// Initialize migration table
	if err := migrator.Init(context.Background()); err != nil {
		log.Fatalf("❌ Failed to initialize migrations: %v", err)
	}

	// Run migrations
	ctx := context.Background()
	group, err := migrator.Migrate(ctx)
	if err != nil {
		log.Fatalf("❌ Migration failed: %v", err)
	}

	if group.ID == 0 {
		fmt.Printf("there are no new migrations to run\n")
		return
	}

	log.Printf("✅ %s migrations applied successfully!", group)
}
