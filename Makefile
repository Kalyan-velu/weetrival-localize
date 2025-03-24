# Load environment variables from .env
include .env
export $(shell sed 's/=.*//' .env)

# Example target that prints a variable
print-env:
	@echo "Database URL: $(DATABASE_URL)"
# Paths
MIGRATE_DIR = ./cmd/migrate
BUN_CMD = $(shell go env GOPATH)/bin/bun

# Install Bun CLI
install:
	@echo "📦 Installing Bun CLI..."
	go install github.com/uptrace/bun/cmd/bun@latest

# Run migrations
migrate:
	@echo "🚀 Running migrations..."
	DATABASE_URL=$(DATABASE_URL) go run $(MIGRATE_DIR)/main.go

# Rollback last migration
rollback:
	@echo "⏪ Rolling back last migration..."
	$(BUN_CMD) migrate rollback

# Create a new migration (Usage: make new-migration name=migration_name)
new-migration:
ifndef name
	$(error "❌ Migration name is required. Usage: make new-migration name=migration_name")
endif
	@echo "📜 Creating new migration: $(name)"
	$(BUN_CMD) migrate new $(name)

# Show migration status
status:
	@echo "📊 Checking migration status..."
	$(BUN_CMD) migrate status

# Reset all migrations (WARNING: This will delete data!)
reset:
	@echo "⚠️  Resetting database (Dropping all tables)..."
	$(BUN_CMD) migrate reset

# Help menu
help:
	@echo "📖 Available commands:"
	@echo "  make install           - Install Bun CLI"
	@echo "  make migrate           - Run all migrations"
	@echo "  make rollback          - Rollback last migration"
	@echo "  make new-migration name=migration_name - Create a new migration"
	@echo "  make status            - Show migration status"
	@echo "  make reset             - Drop all tables and reset"

.PHONY: install migrate rollback new-migration status reset help
