#!/bin/bash
echo "📦 Checking and installing dependencies..."
go mod tidy

echo "🚀 Starting Pismo Tech Case API (without docker)..."
go run cmd/api/main.go