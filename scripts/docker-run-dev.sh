#!/bin/bash

echo "🐳 Starting Docker containers..."
docker-compose -f docker-compose.dev.yaml up
echo "✅ Docker containers are running!"