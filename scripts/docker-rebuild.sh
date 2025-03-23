#!/bin/bash

echo "🔄 Rebuilding Docker containers..."
docker-compose down
docker-compose build --no-cache
docker-compose up -d
echo "✅ Rebuild complete!"
