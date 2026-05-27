# Shop API

Educational Go backend for the SQL workshop.

## Prerequisites

- Go 1.25+
- Docker (for PostgreSQL)
- [swag](https://github.com/swaggo/swag) CLI (`go install github.com/swaggo/swag/cmd/swag@latest`)

## Quick start

```bash
# Start PostgreSQL
docker run -d --name shop-db -e POSTGRES_PASSWORD=password -p 55000:5432 postgres:18

# Create database and tables (connect and run scripts/ manually)

# Install dependencies
go mod tidy

# Run with hot-reload
air

# Or run directly
go run ./cmd/api/main.go
```

The server starts on `http://localhost:8080`.

## Environment variables

| Variable       | Default                                                          |
|----------------|------------------------------------------------------------------|
| `PORT`         | `8080`                                                           |
| `DATABASE_URL` | `postgres://postgres:password@localhost:55000/shop?sslmode=disable` |

Set via `.env` file or system env.

## Endpoints

| Method | Path                   | Description            |
|--------|------------------------|------------------------|
| GET    | `/api/products`        | List products (offset) |
| GET    | `/api/products/cursor` | List products (cursor) |
| GET    | `/api/swagger/index.html` | Swagger UI          |

## Regenerate Swagger docs

```bash
swag init -g ./cmd/api/main.go
```

Auto-regenerated on `air` start (see `.air.toml`).
