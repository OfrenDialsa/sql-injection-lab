# SQL Injection Lab

Repositori ini adalah lingkungan laboratorium sederhana untuk mempelajari dan menguji berbagai teknik **SQL Injection (SQLi)**. Lab ini terdiri dari dua komponen utama: server API yang sengaja dibuat rentan dan alat penguji (*tester*) otomatis.

## Prasyarat

Sebelum memulai, pastikan Anda telah menginstal:
* [Go](https://go.dev/doc/install) (versi 1.26.1 atau yang lebih baru)

## Instalasi

1. Masuk ke direktori proyek:
```bash
cd sql-injection-lab
```

2. Unduh dependencies:
```bash
go get modernc.org/sqlite
go mod tidy
```

---

## Cara Penggunaan

Aplikasi ini menggunakan sistem *command-line interface* (CLI) berbasis Cobra.

### 1. Menjalankan Server API (Vulnerable Target)
Jalankan server API yang akan menjadi target simulasi:
 ```bash
go run main.go api
```
*Server secara default berjalan di http://localhost:8080.*

### 2. Menjalankan SQL Injection Tester
Buka terminal baru, lalu jalankan perintah \`inject\` untuk mulai melakukan pengujian.

**Contoh Perintah:**
 ```bash
# Mode default (basic)
go run main.go inject --url http://localhost:8080 --mode basic

# Mode Auth Bypass SQLi
go run main.go inject --url http://localhost:8080 --mode login

# Mode Time-based Blind SQLi
go run main.go inject --url http://localhost:8080 --mode time

# Mode Boolean Blind SQLi
go run main.go inject --url http://localhost:8080 --mode boolean
```

---

## Flag & Mode

| Flag | Deskripsi | Default |
| :--- | :--- | :--- |
| `--url` | URL target API yang akan diuji | `http://localhost:8080\` |
| `--mode` | Jenis teknik: `basic`, `boolean`, `time`, `login` | `basic` |

---

> [!WARNING]  
> **Peringatan:** Kode ini dibuat hanya untuk tujuan edukasi dan pengujian lokal. Jangan digunakan pada sistem tanpa izin resmi.