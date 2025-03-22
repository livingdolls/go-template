# Stage 1: Build aplikasi Go
FROM golang:1.23.4 AS builder

WORKDIR /app

# Copy go.mod dan go.sum, lalu unduh dependensi
COPY go.mod go.sum ./
RUN go mod download

# Copy seluruh kode sumber
COPY . .

# Build aplikasi dengan static binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/bin/main ./cmd/api/main.go

# Stage 2: Run aplikasi dengan Alpine
FROM alpine:latest

WORKDIR /app

# Install dependencies jika diperlukan (contoh: SSL)
RUN apk --no-cache add ca-certificates

# Copy binary hasil build dari stage builder
COPY --from=builder /app/bin/main /app/bin/main

# Copy file konfigurasi
COPY config/config.yaml /app/config/config.yaml
COPY start.sh /app
# Pastikan file bisa dieksekusi
RUN chmod +x /app/bin/main
RUN chmod +x /app/start.sh
# Expose port aplikasi
EXPOSE 8080

CMD ["/app/start.sh"]