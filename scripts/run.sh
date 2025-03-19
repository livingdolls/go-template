#!/bin/bash

echo "🚀 Starting the server..."

# Load environment variables jika menggunakan .env
if [ -f .env ]; then
    echo "📦 Loading environment variables..."
    export $(grep -v '^#' .env | xargs)
fi

echo "🔄 Running the application..."
go run cmd/api/main.go
