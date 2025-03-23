#!/bin/sh

set -e  # Berhenti jika ada error

echo "🚀 Starting Go application..."

# Load environment variables
if [ -f .env ]; then
    echo "📦 Loading environment variables..."
    set -o allexport
    source .env
    set +o allexport
fi

# Jalankan migrasi database sebelum memulai aplikasi
echo "🔄 Running database migrations..."
make migrate

# Jalankan aplikasi
echo "✅ Application is running!"
./app/bin/main
