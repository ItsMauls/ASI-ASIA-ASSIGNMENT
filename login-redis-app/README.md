# Sistem Login Sederhana dengan Go Fiber dan Redis (JSON API)

## Deskripsi
API sederhana untuk sistem login menggunakan:
- **Go Fiber** sebagai web framework
- **Redis** sebagai database untuk menyimpan data user
- **SHA1** untuk hashing password
- **JSON** untuk request dan response

## Struktur Data di Redis
```
Key: login_<username>
Value: {
  "realname": "Nama Lengkap",
  "email": "email@example.com", 
  "password": "hash_sha1_password"
}
```

## Cara Menjalankan

### 1. Install Dependencies
```bash
go mod tidy
```

### 2. Pastikan Redis Running
Pastikan Redis server berjalan di `localhost:6379`

### 3. Jalankan Aplikasi
```bash
go run main.go
```

## API Endpoints

### 1. POST /create-user
Membuat user baru

**Request:**
```json
{
  "username": "johndoe",
  "realname": "John Doe",
  "email": "john@example.com",
  "password": "mypassword"
}
```

**Response Success:**
```json
{
  "success": true,
  "message": "User berhasil dibuat",
  "data": {
    "username": "johndoe",
    "realname": "John Doe",
    "email": "john@example.com",
    "password_hash": "5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8"
  }
}
```

### 2. POST /login
Login user

**Request:**
```json
{
  "username": "johndoe",
  "password": "mypassword"
}
```

**Response Success:**
```json
{
  "success": true,
  "message": "Login berhasil",
  "data": {
    "username": "johndoe",
    "realname": "John Doe",
    "email": "john@example.com"
  }
}
```

**Response Error:**
```json
{
  "success": false,
  "message": "Username tidak ditemukan"
}
```

## Testing dengan curl

### Buat User Baru:
```bash
curl -X POST http://localhost:3000/create-user \
  -H "Content-Type: application/json" \
  -d '{
    "username": "johndoe",
    "realname": "John Doe", 
    "email": "john@example.com",
    "password": "mypassword"
  }'
```

### Login:
```bash
curl -X POST http://localhost:3000/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "johndoe",
    "password": "mypassword"
  }'
```