#!/bin/sh

set -e  # Berhenti jika ada error

echo "🚀 Starting Go application..."

# Run migrations automatically (only 'up' in production)
if [ -f ./migrate ]; then
  echo "Running database migrations..."
  ./migrate up
  if [ $? -ne 0 ]; then
    echo "Migration failed!"
    exit 1
  fi
fi


# Jalankan aplikasi
echo "✅ Application is running!"
exec ./main 