# Asaba Back End

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

   Database akan auto migrate.

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

4. Rute untuk melihat perubahan stok.

   ```http
   GET /api/stok
   ```

# Note

Pastikan untuk menjalankan server backend terlebih dahulu
