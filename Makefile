.PHONY: run migrate up down fmt lint test build

# Jalankan server
run:
	go run cmd/api/main.go

# Lakukan migrasi database
migrate:
	go run cmd/migrate/main.go up

# Rollback migrasi database
migrate-down:
	go run cmd/migrate/main.go down

# Jalankan Docker (jika ada docker-compose)
up:
	docker-compose up -d

# Hentikan Docker
down:
	docker-compose down

# Format kode agar rapi
fmt:
	go fmt ./...

# Linting kode untuk memastikan best practices
lint:
	golangci-lint run

# Jalankan unit test
test:
	go test ./...

# Build aplikasi
build:
	go build -o bin/app cmd/api/main.go
