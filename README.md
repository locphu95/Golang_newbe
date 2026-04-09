# Golang_newbe

Repo này dùng để học và thực hành Golang.

## Cấu trúc thư mục
- `api/` : chứa các định nghĩa API
- `internal/` : code nội bộ, không export ra ngoài
- `pkg/` : các package có thể tái sử dụng
- `deployments/` : cấu hình triển khai

## Cách chạy
```bash
go run main.go

## Chạy test
go test


1. STRUCTURE

go-backend-template
│
├ cmd
│   └ api
│       └ main.go
│
├ internal
│   ├ domain
│   │   └ user.go
│   │
│   ├ repository
│   │   ├ user_repository.go        (interface)
│   │   └ user_repository_pg.go     (implementation)
│   │
│   ├ service
│   │   └ user_service.go
│   │
│   ├ handler
│   │   └ user_handler.go
│   │
│   ├ transport
│   │   └ http
│   │       ├ router.go
│   │       └ middleware
│   │           ├ logger.go
│   │           ├ recovery.go
│   │           └ auth.go
│
├ pkg
│   ├ config
│   │   └ config.go
│   │
│   ├ database
│   │   └ postgres.go
│   │
│   ├ redis
│   │   └ client.go
│   │
│   ├ logger
│   │   └ zap.go
│   │
│   ├ server
│   │   └ http_server.go
│
├ configs
│   └ config.yaml
│
├ migrations
│
├ docker-compose.yml
├ Dockerfile
├ go.mod


2. FLOW CHUẨN
HTTP
 ↓
middleware
 ↓
handler
 ↓
service
 ↓
repository
 ↓
postgres / redis