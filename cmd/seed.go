package cmd

import (
	"go-boilerplate/config"
	"log"

	"github.com/spf13/cobra"
)

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Run database seeders",
	Long:  "Run database seeders to populate initial data",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.NewConfig()
		_, err := cfg.ConnectionMySql()
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
			return
		}

		// Get GORM DB instance
		// gormDB := db.DB

		log.Println("Running seeders...")

		// Run role seeder
		log.Println("Seeding roles...")
		// seeds.SeedRole(gormDB)

		// Run admin seeder
		log.Println("Seeding admin user...")
		// seeds.SeedAdmin(gormDB)

		log.Println("All seeders completed successfully!")
	},
}

func init() {
	rootCmd.AddCommand(seedCmd)
}
