# Market — Golang Backend

## 📌 Overview

**Market** is a Golang backend service for a marketplace system.
The project is structured to support modular business domains (users, products, sales), database migrations, SQLC-based queries, WebSocket support, and Dockerized deployment.

---

## 🏗️ Project Structure

```
.
├── main.go                # Application entry point
├── app/                   # Core application modules
│   ├── app.go             # App bootstrap / wiring
│   ├── users/             # User domain logic
│   ├── products/          # Product domain logic
│   └── sales/             # Sales domain logic
│
├── api/                   # HTTP API layer
│   ├── api.go             # Router initialization
│   ├── public/            # Public endpoints
│   └── private/           # Auth-protected endpoints
│
├── socket/                # WebSocket handling
│   └── socket.go
│
├── db/                    # Database layer
│   ├── migrations/        # SQL migrations
│   └── query/             # SQLC generated queries
│
├── services/              # External service integrations
│   └── posgresql/         # PostgreSQL connection/service logic
│
├── pkg/                   # Shared packages
│   └── config/            # Config loader & models
│
├── utils/                 # Helper utilities
│   └── yml.go             # YAML helpers
│
├── config.yml             # Application configuration
├── sqlc.yaml              # SQLC configuration
├── docker-compose.yml     # Local environment setup
├── Makefile               # Project commands
├── go.mod / go.sum        # Dependencies
└── README.md
```

---

## 🚀 Features

* Modular domain-driven structure
* SQLC-based type-safe database queries
* Migration-ready PostgreSQL setup
* REST API with public/private routing
* WebSocket support
* YAML-based configuration
* Docker environment for local development

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

### 2. Configure app

Edit `config.yml` with your database and service settings.

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

### 5. Run application

```bash
go run main.go
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

## 📦 Useful Make Commands

```bash
make run          # start app
make migrate-up   # apply migrations
make migrate-down # rollback migration
make sqlc         # regenerate SQLC code
```

---

## 🔐 Configuration

All runtime configuration is stored in:

```
config.yml
```

It is loaded via `pkg/config`.

---

## 📄 License

MIT

---

## 👨‍💻 Notes

This project follows a modular service-oriented structure rather than strict clean architecture, allowing faster development while keeping domains isolated.

---

If you want, I can also generate for you:

* 🔹 `.env` support instead of YAML
* 🔹 GitHub Actions CI
* 🔹 Production-ready Dockerfile
* 🔹 Example config.yml template

Just tell me which one you want next.
