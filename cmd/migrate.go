package cmd

import (
	"go-boilerplate/config"
	"go-boilerplate/database/migrator"
	"log"

	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migrations",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.NewConfig()

		dbConn, err := cfg.ConnectionMySql()
		if err != nil {
			log.Fatal(err)
		}

		sqlDB, err := dbConn.DB.DB()
		if err != nil {
			log.Fatal(err)
		}
		defer sqlDB.Close()

		if err := migrator.RunMigrations(sqlDB, "database/migrations"); err != nil {
			log.Fatal(err)
		}

		log.Println("All migrations applied successfully")
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}

var rollbackCmd = &cobra.Command{
	Use:   "rollback",
	Short: "Rollback last migration",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.NewConfig()
		dbConn, err := cfg.ConnectionMySql()
		if err != nil {
			log.Fatal(err)
		}

		sqlDB, err := dbConn.DB.DB()
		if err != nil {
			log.Fatal(err)
		}
		defer sqlDB.Close()

		if err := migrator.RollbackLast(sqlDB, "database/migrations"); err != nil {
			log.Fatal(err)
		}

		log.Println("Rollback completed successfully")
	},
}

func init() {
	migrateCmd.AddCommand(rollbackCmd)
}
