#!/bin/bash

echo "🚀 Setting up the project..."

# Load environment variables jika menggunakan .env
if [ -f .env ]; then
    echo "📦 Loading environment variables..."
    export $(grep -v '^#' .env | xargs)
fi

echo "🔄 Running database migrations..."
make migrate

echo "✅ Setup completed!"