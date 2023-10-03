# Go Fast Print

Selamat datang di Proyek Go Fast Print! Proyek ini berisi implementasi backend untuk layanan Anda. Di bawah ini, Anda akan menemukan informasi tentang cara menjalankan proyek, endpoint yang tersedia, dan bagaimana cara menggunakannya.

## Menjalankan Proyek

1. Pastikan Anda memiliki Go (Golang) diinstal pada mesin Anda.
2. Clone repositori ini ke direktori lokal Anda.
3. Buka terminal dan navigasikan ke direktori repositori yang baru saja Anda klon.
4. Jalankan perintah berikut untuk menginstal dependensi:

   ```sh

   go mod download

   ```

5. Konfigurasi PostgreSQL (Golang)

   Untuk menyesuaikan konfigurasi PostgreSQL pada proyek Golang menggunakan GORM:

   Gantilah variabel dsn pada file setup.go dengan konfigurasi yang sesuai:

   ```sh
   dsn := "host=your_host user=your_user password=your_password dbname=your_db port=5432 sslmode=disable TimeZone=Asia/Jakarta"
   ```

6. Setelah dependensi terinstal, database telah di konfigurasi lalu jalankan perintah berikut untuk memulai server:

   ```sh

   go run main.go
   ```

   Server akan berjalan pada alamat http://localhost:8080.
   Database akan auto migrate

# Endpoint

1. Rute untuk mengambil data barang.

   ```http
   GET /api/barang
   ```

2. Rute untuk menambah data barang.

   ```http
   POST /api/barang
   ```

   Body (Json)

   ```sh
    [
        {
            code: "",
            nama: "",
            jumlah: "",
            deskripsi: "",
            isActive: "",
        }
    ]
   ```

3. Rute untuk mengubah stok.

   ```http
   POST /api/stok
   ```

   Body (Json)

   ```sh
   [
        {
            code: "",
            tipe: "",
            jumlah: "",
        }
   ]
   ```

# Note

Untuk menambahkan kategori dan status dilakukan dengan melakukan POST data melalui postman. Hal tersebut diperlukan karena dalam view Vue JS saya tidak membuat fitur untuk mengelola kategori dan status

Pastikan untuk menjalankan server backend terlebih dahulu
