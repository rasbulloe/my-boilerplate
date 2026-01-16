package migrator

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func RollbackLast(db *sql.DB, dir string) error {
	var filename string

	err := db.QueryRow(`
		SELECT filename
		FROM schema_migrations
		ORDER BY id DESC
		LIMIT 1
	`).Scan(&filename)

	if err == sql.ErrNoRows {
		return fmt.Errorf("no migration to rollback")
	}
	if err != nil {
		return err
	}

	downFile := strings.Replace(filename, ".up.sql", ".down.sql", 1)
	fullPath := filepath.Join(dir, downFile)

	sqlBytes, err := os.ReadFile(fullPath)
	if err != nil {
		return fmt.Errorf("rollback file not found: %s", downFile)
	}

	fmt.Println("Rolling back:", downFile)

	if _, err := db.Exec(string(sqlBytes)); err != nil {
		return err
	}

	_, err = db.Exec(
		"DELETE FROM schema_migrations WHERE filename = ?",
		filename,
	)
	return err
}
