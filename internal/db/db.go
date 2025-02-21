package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

// DB is the global database instance
var DB *bun.DB

func ConnectDB() {
	dsn := "postgres://postgres:kalyan%40postgre@localhost:5432/weetrival?sslmode=disable"
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	DB = bun.NewDB(sqldb, pgdialect.New())

	// Test the connection
	if err := DB.Ping(); err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	log.Println("Database connected successfully")
}

func CloseDB() {
	if err := DB.Close(); err != nil {
		log.Printf("Error closing database: %v", err)
	}
}
