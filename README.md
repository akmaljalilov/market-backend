# Market — Backend Service (Golang)

## 📌 Overview

**Market** is a backend service written in **Golang**, designed to power a marketplace platform.
It provides APIs for managing users, products, orders, and payments, and is built with scalability and microservice-friendly architecture in mind.

---

## 🚀 Features

* RESTful API built with Golang
* Modular project structure for easy scaling
* Database integration (PostgreSQL/MySQL supported)
* Authentication & authorization ready
* Environment-based configuration
* Docker support for containerized deployment
* CI/CD friendly

---

## 🏗️ Project Structure

```
market/
│
├── cmd/                # Application entrypoints
├── internal/           # Core business logic
│   ├── handlers/       # HTTP handlers
│   ├── services/       # Business services
│   ├── repositories/   # Database layer
│   └── models/         # Data models
│
├── pkg/                # Shared utilities
├── configs/            # Config files
├── migrations/         # DB migrations
├── scripts/            # Helper scripts
├── docker/             # Docker files
└── README.md
```

---

## ⚙️ Requirements

* Go 1.20+
* PostgreSQL / MySQL
* Docker (optional but recommended)

---

## 🔧 Installation

### 1. Clone repository

```bash
git clone https://github.com/yourusername/market.git
cd market
```

### 2. Install dependencies

```bash
go mod download
```

### 3. Configure environment

Create `.env` file:

```
APP_PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=market
```

### 4. Run migrations

```bash
make migrate-up
```

### 5. Start server

```bash
go run ./cmd/app/main.go
```

---

## 🐳 Run with Docker

```bash
docker-compose up --build
```

---

## 📡 API Example

```
GET /health        -> service status
POST /users        -> create user
POST /products     -> create product
POST /orders       -> create order
```

---

## 🧪 Testing

```bash
go test ./...
```

---

## 📦 Build

```bash
go build -o market ./cmd/app
```

---

## 🔐 Environment Variables

| Variable    | Description   |
| ----------- | ------------- |
| APP_PORT    | Server port   |
| DB_HOST     | Database host |
| DB_PORT     | Database port |
| DB_USER     | DB user       |
| DB_PASSWORD | DB password   |
| DB_NAME     | Database name |

---

## 🤝 Contributing

1. Fork repository
2. Create feature branch
3. Commit changes
4. Open Pull Request

---

## 📄 License

This project is licensed under the MIT License.

---

## 👨‍💻 Author

Backend powered by Golang for scalable marketplace systems.
