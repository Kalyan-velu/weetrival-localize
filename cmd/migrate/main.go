package main

import (
	"log"
	"os"

	"github.com/kalyan-velu/weetrival-localize/internal/migrate"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("❌ DATABASE_URL is not set!")
	}

	migrate.RunMigrations(dsn)
}
