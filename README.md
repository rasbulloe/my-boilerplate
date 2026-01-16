# Boilerplate Go

Repositori ini berisi proyek boilerplate untuk memulai aplikasi Golang baru dari awal. Ini menyediakan struktur dasar, konfigurasi, pengaturan database, dan komponen penting untuk mempercepat pengembangan proyek Go Anda.

## Fitur

- **Perintah CLI**: Termasuk perintah untuk migrasi database, seeding, dan memulai aplikasi.
- **Integrasi Database**: Migrasi dan skrip seeding database yang sudah dikonfigurasi sebelumnya.
- **Manajemen Konfigurasi**: Penanganan konfigurasi berbasis environment.
- **Pengaturan Framework Web**: Struktur routing dan aplikasi dasar menggunakan paket internal.
- **Struktur Proyek**: Direktori yang terorganisir untuk adapter, aplikasi, konfigurasi, dan utilitas.

## Struktur Proyek

```
.
├── cmd/                    # Perintah CLI (migrate, seed, start, root)
├── config/                 # File konfigurasi (config.go, database.go)
├── database/               # File terkait database
│   ├── migrations/         # File migrasi SQL
│   └── migrator/           # Utilitas migrasi
├── internal/               # Kode aplikasi internal
│   ├── adapter/            # Adapter (misalnya, adapter database)
│   └── apps/               # Logika aplikasi (app.go, router.go)
├── utils/                  # Fungsi utilitas
├── .env-example            # Contoh variabel environment
├── .gitignore              # Aturan git ignore
├── go.mod                  # File modul Go
├── go.sum                  # Checksum dependensi Go
└── main.go                 # Titik masuk aplikasi
```

## Memulai

1. **Clone Repositori**:

   ```
   git clone https://github.com/your-username/go-boilerplate.git
   cd go-boilerplate
   ```

2. **Instal Dependensi**:

   ```
   go mod download
   ```

3. **Pengaturan Environment**:

   - Salin `.env-example` ke `.env`
   - Perbarui variabel environment sesuai kebutuhan (misalnya, string koneksi database)

4. **Jalankan Migrasi Database**:

   ```
   go run main.go migrate
   ```

5. **Seed Database** (jika berlaku):

   ```
   go run main.go seed
   ```

6. **Mulai Aplikasi**:
   ```
   go run main.go start
   ```

## Penggunaan

Boilerplate ini dirancang sebagai titik awal untuk proyek Go baru. Anda dapat menyesuaikan kode di direktori `internal/`, menambahkan perintah CLI baru di `cmd/`, dan memperluas skema database sesuai kebutuhan.

## Kontribusi

Silakan berkontribusi dengan mengirimkan masalah atau pull request untuk meningkatkan boilerplate ini.

## Lisensi

Proyek ini dilisensikan di bawah Lisensi MIT.
