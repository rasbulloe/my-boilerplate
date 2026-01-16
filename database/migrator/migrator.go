package migrator

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

// Mesin eksekusi migrate yang ada di cmd/migrate.go

func RunMigrations(db *sql.DB, dir string) error {
	if err := EnsureSchemaTable(db); err != nil {
		return err
	}

	files, err := filepath.Glob(filepath.Join(dir, "*.up.sql"))
	if err != nil {
		return err
	}
	sort.Strings(files)

	for _, file := range files {
		filename := filepath.Base(file)

		var exists int
		err := db.QueryRow(
			"SELECT COUNT(*) FROM schema_migrations WHERE filename = ?",
			filename,
		).Scan(&exists)
		if err != nil {
			return err
		}

		if exists > 0 {
			continue // sudah dijalankan
		}

		sqlBytes, err := os.ReadFile(file)
		if err != nil {
			return err
		}

		fmt.Println("Running migration:", filename)

		if _, err := db.Exec(string(sqlBytes)); err != nil {
			return fmt.Errorf("migration %s failed: %w", filename, err)
		}

		_, err = db.Exec(
			"INSERT INTO schema_migrations (filename) VALUES (?)",
			filename,
		)
		if err != nil {
			return err
		}

		fmt.Println("✓ Done:", filename)
	}

	return nil
}
