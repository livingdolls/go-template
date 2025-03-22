#!/bin/bash

set -e  # Berhenti jika ada error

echo "🚀 Setting up the project..."

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

# Pastikan `make` tersedia sebelum menjalankan migrasi
if command_exists make; then
    echo "🔄 Running database migrations..."
    make migrate
else
    echo "❌ Error: 'make' is not installed. Please install it first."
    exit 1
fi

echo "✅ Setup completed!"
