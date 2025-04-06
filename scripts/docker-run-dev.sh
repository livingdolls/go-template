#!/bin/bash

echo "🐳 Starting Docker containers..."
docker-compose -f docker-compose.dev.yml up -d
echo "✅ Docker containers are running!"