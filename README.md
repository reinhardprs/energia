# Energia API Mini Project

Energia API adalah API sederhana yang dirancang untuk mengelola penggunaan perangkat listrik di rumah. Selain itu, API ini juga memberikan saran penggunaan perangkat berdasarkan informasi cuaca terkini.

---

## Entity Relationship Diagram (ERD)

Berikut adalah ERD yang digunakan dalam project ini:

![ERD](images/Energia_ERD.png)

---

## High-Level Architecture (HLA)

Berikut adalah High-Level Architecture (HLA) dari Energia API:

![HLA](images/Energia_HLA.png)

---

## Host untuk Penggunaan API

Energia API telah dihosting di cloud sehingga Anda dapat menggunakannya tanpa instalasi tambahan. Anda dapat mengakses API melalui host berikut:

http://52.65.161.24

---

## Dokumentasi API

Untuk mengakses dokumentasi API, ikuti langkah-langkah berikut:

### 1. Clone repositori ini

Clone repositori ke dalam direktori lokal Anda menggunakan perintah berikut:

```bash  
git clone <repository_url>
```

### 2. Buat file .env
Buat file .env di direktori root proyek dan isi dengan konfigurasi berikut:

```plaintext
DATABASE_HOST=""
DATABASE_PORT=""
DATABASE_USER=""
DATABASE_PASSWORD=""
DATABASE_NAME=""
JWT_SECRET_KEY=""
OPENWEATHER_API_KEY=""
OPENAI_API_KEY=""
MAIL_USER=""
MAIL_PASSWORD=""
MAIL_HOST=""
MAIL_PORT=""
```

### 3. Jalankan aplikasi
Jalankan perintah berikut untuk memulai aplikasi:

```bash
go run main.go
```

### 4. Akses Dokumentasi API
Buka browser Anda dan akses dokumentasi API melalui URL berikut:

```plaintext
http://{{host}}/swagger/index.html#
```

Catatan: Gantilah {{host}} dengan alamat host yang sesuai, misalnya localhost:8080 jika aplikasi dijalankan secara lokal.