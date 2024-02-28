# Proyek Golang dengan Database dan Konfigurasi .env

Proyek ini adalah contoh aplikasi Golang sederhana yang menggunakan database, dan konfigurasi database didefinisikan dalam file `.env`.

## Persiapan

1. Pastikan Golang sudah terinstal di komputer Anda. Jika belum, unduh dan instal dari [https://golang.org/dl/](https://golang.org/dl/).

2. Pastikan [Git](https://git-scm.com/) juga terinstal di komputer Anda.

3. Unduh proyek ini ke dalam folder lokal:

    ```bash
    git clone https://github.com/saahalla/online-store.git
    ```

4. Pindah ke direktori proyek:

    ```bash
    cd online-store
    ```

5. Instal dependensi proyek:

    ```bash
    go mod tidy
    ```

## Konfigurasi Database

1. Salin contoh file `.env.example` ke `.env`:

    ```bash
    cp .env.example .env
    ```

2. Edit file `.env` dan sesuaikan informasi konfigurasi database Anda:

    ```env
    MYSQL_HOST=localhost
    MYSQL_PORT=3306
    MYSQL_USER=your_username
    MYSQL_PASSWORD=your_password
    MYSQL_DBNAME=your_database_name
    ```

## Menjalankan Aplikasi

1. Pastikan server database Anda sudah berjalan.

2. Jalankan aplikasi Golang:

    ```bash
    go run main.go
    ```

    Aplikasi akan berjalan pada [http://localhost:3030](http://localhost:3030).

## Menjalankan Aplikasi dengan docker

1. Build Image
    
    ```bash
    docker build . -t saahalla/online-store-be-golang
    ```

2. Run Image beserta env database

    ```bash
    docker run -e MYSQL_HOST=127.0.0.1 -e MYSQL_USER=user -e MYSQL_PASSWORD=password -e MYSQL_DBNAME=online_store -e MYSQL_PORT=3306 -p 3030:3030 saahalla/online-store-be-golang
    ```

## Endpoints

- `/category` [GET] : mendapatkan list category
- `/category/:id` [GET] : mendapatkan detail category
- `/category/:id` [PUT] : untuk update category
- `/category/` [ADD] : untuk menambah category
- `/category/:id` [DELETE] : untuk menghapus category

- `/product` [GET] : mendapatkan list product
- `/product/:id` [GET] : mendapatkan detail product
- `/product/:id` [PUT] : untuk update product
- `/product/` [ADD] : untuk menambah product
- `/product/:id` [DELETE] : untuk menghapus product

- `/auth/register` [POST]: untuk mendaftar akun
- `/auth/login` [POST]: untuk login akun

## Catatan

- Pastikan untuk menjaga file `.env` agar tidak dibagikan bersamaan dengan kode sumber Anda, karena berisi informasi rahasia.

- Sesuaikan database dan tabel sesuai dengan kebutuhan proyek Anda.

- Jangan lupa memastikan dependensi proyek terinstal dengan benar menggunakan perintah `go mod tidy`.

Selamat menjelajahi proyek Golang Anda dengan database dan konfigurasi .env!

## ERD Database

![alt text](https://github.com/saahalla/online-store/blob/master/erd.png?raw=true)