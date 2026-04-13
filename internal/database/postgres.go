package database

import (
	"database/sql"
	"embed"
	"fmt"
	"log/slog"
	"time"

	_ "github.com/lib/pq"
)

// Connect opens a PostgreSQL connection pool and verifies connectivity.
func Connect(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %w", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("db.Ping: %w", err)
	}

	slog.Info("Database connected")
	return db, nil
}

//go:embed ../../migrations/*.sql
var migrationsFS embed.FS

// RunMigrations executes all SQL migration files in order.
func RunMigrations(db *sql.DB) error {
	entries, err := migrationsFS.ReadDir("migrations")
	if err != nil {
		return err
	}
	_, _ = db.Exec(`CREATE TABLE IF NOT EXISTS schema_migrations (version TEXT PRIMARY KEY)`)
	for _, entry := range entries {
		var exists bool
		_ = db.QueryRow("SELECT EXISTS(SELECT 1 FROM schema_migrations WHERE version=$1)", entry.Name()).Scan(&exists)
		if exists {
			continue
		}
		data, _ := migrationsFS.ReadFile("migrations/" + entry.Name())
		if _, err := db.Exec(string(data)); err != nil {
			return fmt.Errorf("migration %s: %w", entry.Name(), err)
		}
		_, _ = db.Exec("INSERT INTO schema_migrations (version) VALUES ($1)", entry.Name())
		slog.Info("Migration applied", "file", entry.Name())
	}
	return nil
}
