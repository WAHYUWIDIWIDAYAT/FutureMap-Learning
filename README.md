# Final Project Ruangguru

## Table of Contents

- ⚙️ [Backend](#backend)
  - [Teknologi Backend](#teknologi-backend)
  - [Development Backend](#development-backend)
  - [Contoh Response](#contoh-response)

---

## Backend

Buka terminal favorit kalian dan jalankan perintah-perintah berikut ini. Selain itu juga disediakan _build version_ berupa file exe.

Buat key cloud google and replace futurego.json di controllers/learn.go
### 📚 Teknologi Backend

- Golang
- Gin Gonic
- JWT-GO
- Cloud Firestore (Firebase)

### 🛠 Development Backend

Untuk menjalankan Project Backend:

```bash
go run main.go
```

Untuk menambahkan Admin:

```bash
localhost:8080/admin/register
```

Untuk login Admin:

```bash
localhost:8080/login
```

### 📲 Contoh Response

JSON data Admin register:

```json
{
  "username": "admin",
  "phone": "12345",
  "email": "admin123",
  "password": "admin12345"
}
```

JSON data Admin login:

```json
{
  "email": "admin123",
  "password": "admin12345"
}
```

Contoh response register dan login untuk User, sama dengan Admin.

---

