# Golang Crud
Ini adalah project simple CRUD (Create, Read, Update, Delete) menggunakan bahasa Go. Yang dimana pengguna bisa memanajemen Product serta Category. Disini pembuat menggunakan database PostgreSQL untuk menyimpan data, menggunakan environment variabel untuk memudahkan dalam mengatur konfigurasi serta mengamankan data kredensial, seperti Password database.

## Cara Penggunaan
- Buat file `.env` di working directory, lalu isi key and value samakan dengan yang ada di [.env.examples](.env.examples) (key nya)
- Copy kan query yang ada di [table.sql](table.sql) untuk membuat struktur tabel. Pastikan sudah membuat database nya dan didaftarkan ke env **DB_NAME**. Jika menggunakan database selain PostgreSQL, ubah konfigurasi yang ada di `config/database.go` lalu ubah variabel `dbInfo` menjadi statement untuk mengakses ke database yang digunakan
- Jalankan aplikasi menggunakan perintah `go run .` atau `go run main.go`
- Lalu buka browser dan ketikkan di pencarian `http://localhost:8080`

Dan aplikasi sudah bisa untuk digunakan
