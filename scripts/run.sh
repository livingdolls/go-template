#!/bin/bash

set -e  # Exit jika ada error

echo "🚀 Starting the server..."

# Load environment variables dengan cara yang lebih aman
if [ -f .env ]; then
    echo "📦 Loading environment variables..."
    set -o allexport
    source .env
    set +o allexport
fi

# Fungsi untuk memeriksa apakah sebuah perintah tersedia
command_exists() {
    command -v "$1" &> /dev/null
}

# Jalankan dengan Air jika tersedia
if command_exists air; then
    echo "🔥 Running with Air (hot-reload enabled)"
    air
    exit 0
fi

# Jika ingin menjalankan Docker services sebelum aplikasi
if [ "$1" == "docker-db" ]; then
    echo "🐳 Starting Docker services (DB, RabbitMQ, etc.)..."
    docker container start mysql-container
fi

# Pastikan Go terinstal sebelum menjalankan aplikasi
if command_exists go; then
    echo "🔄 Running the application..."
    go run cmd/api/main.go
else
    echo "❌ Error: Go is not installed. Please install Go first."
    exit 1
fi
