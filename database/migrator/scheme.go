package migrator

import "database/sql"

func EnsureSchemaTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS schema_migrations (
		id INT AUTO_INCREMENT PRIMARY KEY,
		filename VARCHAR(255) UNIQUE,
		applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`
	_, err := db.Exec(query)
	return err
}
