#!/bin/sh

# Executar migrações
echo "🔄 Running migrations..."
go run ./cmd/api/main.go --migrate

# Executar seeders
echo "🌱 Running seeders..."
go run ./cmd/api/main.go --seed

# Iniciar Air para hot reload
echo "🔥 Starting Air for hot reload..."
exec air
