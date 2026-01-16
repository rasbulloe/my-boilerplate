package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var rootCmd = &cobra.Command{

	// Artinya:
	// Ini command utama
	// Nama aplikasinya: go-boilerplate
	// build dulu aplikasinya:
	// go build -o go-boilerplate main.go
	// Jalankan di terminal:
	// ./go-boilerplate
	// Maka Cobra akan menjalankan command ini

	Use:   "go-boilerplate",
	Short: "A brief description of your application",

	// Maknanya:
	// Jika user hanya mengetik:
	// "go-boilerplate" di command line
	// Maka Cobra akan menjalankan startCmd
	// startCmd didefinisikan di file lain (start.go)

	Run: func(cmd *cobra.Command, args []string) {
		// startCmd ada di file start.go
		cmd.Run(startCmd, nil)
	},
}

func Execute() {
	// Fungsi ini:
	// Dipanggil dari main.go
	// Memulai seluruh proses CLI:
	// parsing command
	// parsing flag
	// menjalankan subcommand
	// 📌 Tanpa Execute(), CLI tidak akan berjalan
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	// Disini kita bisa menambahkan flag atau konfigurasi lain jika diperlukan
	// Kapan init() dijalankan?
	// SEBELUM main()
	// Otomatis oleh Go
	// Untuk semua file & package

	// Inisialisasi konfigurasi sebelum menjalankan command apapun
	// Saat Cobra mulai berjalan, jalankan initConfig()
	cobra.OnInitialize(initConfig)

	// Definisi flag global untuk semua command
	// Flag: --config
	// Digunakan untuk menentukan file konfigurasi yang akan dibaca oleh Viper
	// karna kita set nama buildernya "go-boilerplate"
	// Maka user bisa menjalankan:
	// ./go-boilerplate --config=.env.prod
	// Maka Viper akan membaca dan menggunakan file .env.prod sebagai konfigurasi
	rootCmd.PersistentFlags().StringVar(
		&cfgFile,
		"config",
		"",
		"config file (default is .env)",
	)

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigFile(`.env`)
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

// Tujuannya:

// Membuat aplikasi Go kamu bisa dijalankan lewat terminal sebagai command

// Contoh:

// go-boilerplate start
// go-boilerplate --config=.env.prod

// Cobra → membuat CLI command (go-boilerplate start)
// Viper → membaca konfigurasi (.env, env variable, config file)

// Dengan struktur file seperti diatas, biasanya digunakan untuk arsitektur Microservices, karna mungkin banyak environment yang berbeda untuk setiap service nya.

// =====================================================================//

// Bagaimana cara kalo aplikasinya cuman Monolith?
// Apakah perlu dipisah-pisah file seperti ini?
// Jawabannya: Tergantung kebutuhan dan preferensi masing-masing developer
// Kalau aplikasinya kecil dan sederhana, boleh saja ditaruh di satu file main.go
// Ada 3 Model umum:

// 1. Semua di main.go
// Cocok untuk aplikasi kecil dan sederhana
// Contoh: Aplikasi CLI sederhana, script otomatisasi kecil
// 2. Dipisah file cmd/ seperti ini
// Cocok untuk aplikasi yang mulai kompleks, ada banyak command
// Contoh: Aplikasi CLI dengan banyak fitur, Microservices sederhana
// 3. Struktur proyek lengkap dengan internal/, pkg/, cmd/, dll
// Cocok untuk aplikasi besar dan kompleks
// Contoh: Aplikasi Microservices besar, aplikasi web lengkap dengan backend dan frontend

// Contoh model 1: Semua di main.go
// func main() {
// 	_ = godotenv.Load()

// 	port := os.Getenv("APP_PORT")
// 	if port == "" {
// 		port = "8080"
// 	}

// 	log.Println("Server running on port", port)
// }
