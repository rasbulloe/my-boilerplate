DROP INDEX idx_users_is_verified ON users;
-- Menghapus index idx_users_is_verified yang dibuat pada migration 002_add_users_indexes.up.sql
-- Ini dilakukan untuk mengembalikan struktur database ke kondisi sebelum migration tersebut dijalankan
-- Sehingga jika migration ini di-rollback, index tersebut tidak akan ada lagi di table users
-- Hal ini penting untuk menjaga konsistensi struktur database sesuai dengan versi migration yang sedang digunakan
-- Selain itu, menghapus index juga dapat membantu mengurangi beban pada operasi insert, update, atau delete data pada table users
-- karena index tidak perlu diperbarui lagi setelah dihapus
-- Pastikan untuk selalu menyesuaikan migration up dan down agar perubahan pada database dapat dikelola dengan baik