# Market — Golang Backend Service

## 📌 Overview

**Market** is a backend service written in **Go**, designed for a marketplace system.
The project follows a clean, modular structure using `internal` packages, SQLC for database access, migrations, and Docker for local development.

---

## 🏗️ Project Structure

```
.
├── cmd/server/main.go        # Application entrypoint
├── internal/                 # Private application code
│   ├── bootstrap/            # App initialization & wiring
│   ├── config/               # Config loader & models
│   ├── app/                  # Business entities & logic
│   ├── repository/           # Database access layer
│   ├── transport/
│   │   ├── http/             # HTTP server
│   │   │   ├── handlers/     # HTTP handlers
│   │   │   ├── middleware/   # HTTP middleware
│   │   │   └── router.go     # Route setup
│   │   └── websocket/        # WebSocket transport
│   └── utils/                # Shared utilities
│
├── db/
│   ├── migrations/           # SQL migrations
│   └── query/                # SQLC generated queries
│
├── config.yml                # App configuration
├── sqlc.yaml                 # SQLC configuration
├── docker-compose.yml        # Local services setup
├── Makefile                  # Project commands
└── README.md
```

---

## 🚀 Features

* Clean internal package architecture
* SQLC type-safe queries
* Migration-ready PostgreSQL setup
* HTTP API with middleware support
* WebSocket support
* YAML-based configuration
* Docker environment for development

---

## ⚙️ Requirements

* Go 1.20+
* PostgreSQL
* sqlc
* Docker (optional)

---

## 🔧 Setup

### 1. Install dependencies

```bash
go mod download
```

### 2. Configure application

Edit:

```
config.yml
```

Set database credentials, ports, and service settings.

---

### 3. Run migrations

```bash
make migrate-up
```

---

### 4. Generate SQLC code

```bash
make sqlc
```

---

### 5. Start server

```bash
go run ./cmd/server
```

---

## 🐳 Run with Docker

```bash
docker-compose up --build
```

---

## 🧪 Testing

```bash
go test ./...
```

---

## 📦 Common Make Commands

```bash
make run           # Start application
make migrate-up    # Apply DB migrations
make migrate-down  # Rollback migrations
make sqlc          # Generate SQLC queries
```

---

## 🔐 Configuration

All runtime configuration is loaded from:

```
config.yml
```

Handled by the `internal/config` package.

---

## 📄 License

MIT

---

## 👨‍💻 Notes

This project keeps all business logic inside `internal` to prevent unintended external imports and enforce clear service boundaries.

---

If you want, I can also generate:

* Production-ready **Dockerfile**
* Example **config.yml**
* CI pipeline for **GitHub Actions**
* Base **
