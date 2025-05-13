# 🧾 Asset Management

**Asset Management** adalah sebuah sistem backend yang dibangun untuk mengelola aset, pemeliharaan, dan pencatatan logistik dengan kontrol akses berbasis peran (role-based access control). Proyek ini dibuat menggunakan **Golang**, dengan framework **Gin**, **MySQL** sebagai database, dan **Docker** untuk manajemen kontainerisasi.

---

## 🚀 Teknologi yang Digunakan

- **Go** (Golang)
- **Gin** (HTTP Web Framework)
- **MySQL** (Relational Database)
- **Docker** (Containerization)
- **Postman** (API Testing)

---

## 📁 Struktur Proyek

├── config # Konfigurasi database, env

├── controllers # Handler untuk permintaan HTTP

├── dto # Data Transfer Objects

├── middleware # Middleware untuk autentikasi & otorisasi

├── models # Model untuk representasi tabel DB

├── repositories # Query dan operasi DB

├── routes # Routing endpoint API

├── service # Logika bisnis aplikasi

├── Dockerfile # Dockerfile untuk build container


---

## 🔐 Role dan Fitur

| Role      | Fitur                                                                 |
|-----------|------------------------------------------------------------------------|
| **Auth**     | Register, Login                                                      |
| **Logistik** | CRUD Asset, Get Maintenance, Get Asset Log                          |
| **Engineer** | Create Maintenance, Get Maintenance, Get Asset                      |
| **Manajer**  | Get Asset, Get Maintenance, Get Asset Log                           |

---

## ⚙️ Cara Menjalankan (Dengan Docker)

### 1. Clone Repositori

bash
git clone https://github.com/anggajayadii/final-project-dibimbing.git
cd final-project-dibimbing

### 2. Build dan Jalankan dengan Docker
docker build -t asset-management .
docker run -p 8080:8080 asset-management

