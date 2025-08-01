# 🧪 Go Minimal REST API Server

This is a minimal HTTP server written purely using Go's **standard library**, showcasing how to build a basic RESTful API with no external dependencies.

---

## ✨ Features

- Built using only `net/http` (no frameworks)
- Thread-safe in-memory user storage with `sync.RWMutex`
- RESTful endpoints:
  - `GET /user/{id}` – Get user by ID
  - `POST /user` – Create new user
  - `DELETE /user/{id}` – Delete user by ID
- Go 1.22+ route pattern matching (e.g. `/user/{id}`)

---

## 🧱 Tech Stack

- **Language:** Go (>= 1.22)
- **Storage:** In-memory map with mutex locking
- **Frameworks:** None

---

## 🚀 Getting Started

### Prerequisites

- Go 1.22 or higher

### Running the Server

```bash
go run main.go
```

Server will start at:

```
http://localhost:8080
```

---

## 📡 API Endpoints

### `GET /`
Returns a simple welcome message.

**Response:**
```
Hello Client!
```

---

### `POST /user`

Creates a new user.

**Request Header:**
```
Content-Type: application/json
```

**Request Body:**
```json
{
  "name": "Ameer"
}
```

**Responses:**
- `202 Accepted` – User successfully created
- `400 Bad Request` – Invalid or empty name

---

### `GET /user/{id}`

Fetches a user by ID.

**Example:**
```
GET /user/1
```

**Response:**
```json
{
  "name": "Ameer"
}
```

**Responses:**
- `200 OK` – User found
- `400 Bad Request` – Invalid ID
- `404 Not Found` – User not found

---

### `DELETE /user/{id}`

Deletes a user by ID.

**Example:**
```
DELETE /user/1
```

**Responses:**
- `204 No Content` – User deleted
- `400 Bad Request` – Invalid ID
- `404 Not Found` – User not found

---

## ⚠️ Limitations

- Data is stored in-memory (non-persistent)
- No input sanitization beyond basic validation
- Meant for experimentation, not production

---

## 🧠 Why This Project?

This example demonstrates:

- How to build an HTTP server with Go’s standard library
- RESTful route handling in Go 1.22+
- Basic concurrency-safe in-memory data operations

---

## 📜 License

MIT License — feel free to use and modify.

---

> Built with ❤️ in Go by [Ameer Heiba](https://github.com/AmeerHeiba)
