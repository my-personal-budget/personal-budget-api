package store

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"log"
	"time"
)

func CreateMigration(name string) {
	migrationsDir := "migrations"

	if _, err := os.Stat(migrationsDir); os.IsNotExist(err) {
		if err := os.Mkdir(migrationsDir, 0755); err != nil {
			log.Fatalf("failed to create migrations directory: %w", err)
			return
		}
	}

	timestamp := time.Now().Format("20060102150405")
	fileNameWithoutExtension := fmt.Sprintf("%s_%s", timestamp, name)

	migrationFileName := filepath.Join(migrationsDir, fileNameWithoutExtension+".sql")

	if err := os.WriteFile(migrationFileName, []byte("-- Write your migration here\n"), 0644); err != nil {
		log.Fatalf("failed to create up migration: %w", err)
		return
	}

	log.Println("Created migration files: " + migrationFileName)
	return 
}

// Migrate applies all pending SQL migrations in ./migrations
func Migrate(s *Store) error {
	// Ensure schema_migrations table exists
	_, err := s.DB.Exec(`
		CREATE TABLE IF NOT EXISTS schema_migrations (
			version TEXT PRIMARY KEY
		)
	`)
	if err != nil {
		return fmt.Errorf("failed to ensure schema_migrations table: %w", err)
	}

	// Get applied migrations
	applied := make(map[string]bool)
	rows, err := s.DB.Query("SELECT version FROM schema_migrations")
	if err != nil {
		return fmt.Errorf("failed to query schema_migrations: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var version string
		if err := rows.Scan(&version); err != nil {
			return fmt.Errorf("failed to scan schema_migrations: %w", err)
		}
		applied[version] = true
	}

	// Load migration files
	dir := "migrations"
	entries, err := os.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("No migrations directory found, skipping.")
			return nil
		}
		return fmt.Errorf("failed to read migrations dir: %w", err)
	}

	var files []string
	for _, e := range entries {
		if !e.IsDir() && strings.HasSuffix(e.Name(), ".sql") {
			files = append(files, e.Name())
		}
	}

	// Sort migrations by filename (timestamp prefix)
	sort.Strings(files)

	// Apply new migrations
	for _, f := range files {
		if applied[f] {
			continue
		}

		path := filepath.Join(dir, f)
		sqlBytes, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("failed to read migration file %s: %w", f, err)
		}

		tx, err := s.DB.Begin()
		if err != nil {
			return fmt.Errorf("failed to start tx for migration %s: %w", f, err)
		}

		// Execute migration SQL
		if _, err := tx.Exec(string(sqlBytes)); err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to execute migration %s: %w", f, err)
		}

		// Record migration version
		if _, err := tx.Exec("INSERT INTO schema_migrations (version) VALUES (?)", f); err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to record migration %s: %w", f, err)
		}

		if err := tx.Commit(); err != nil {
			return fmt.Errorf("failed to commit migration %s: %w", f, err)
		}

		fmt.Printf("Applied migration: %s\n", f)
	}

	fmt.Println("All migrations applied successfully.")
	return nil
}
