CREATE INDEX idx_users_is_verified ON users(is_verified);

-- Index dibuat supaya database bisa mencari data jauh lebih cepat, tanpa harus membaca seluruh table.
-- Misal kita ingin mencari user yang sudah terverifikasi (is_verified = true), maka database bisa langsung menggunakan index ini untuk menemukan data tersebut dengan cepat.
-- Index ini akan sangat berguna jika table users memiliki jutaan data, sehingga proses pencarian menjadi lebih efisien.
-- Namun perlu diingat, pembuatan index juga memiliki trade-off, yaitu menambah beban saat proses insert, update, atau delete data, karena index juga harus diperbarui.
-- Oleh karena itu, index sebaiknya dibuat pada kolom yang sering digunakan dalam kondisi pencarian (WHERE clause) atau pengurutan (ORDER BY clause).