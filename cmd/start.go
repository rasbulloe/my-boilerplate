package cmd

import (
	"go-boilerplate/internal/apps"

	"github.com/spf13/cobra"
)

// startCmd adalah subcommand dari aplikasi CLI kamu.
// "go-boilerplate start" di CLI akan menjalankan command ini.

// Mendefinisikan command baru
// Tipe data: *cobra.Command
// Disimpan sebagai variabel global agar bisa di-register
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start",
	Long:  "start",

	// Ketika user menjalankan go-boilerplate start, Cobra akan memanggil fungsi ini.
	Run: func(cmd *cobra.Command, args []string) {
		// apps berada di folder internal/apps
		apps.RunServer()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
