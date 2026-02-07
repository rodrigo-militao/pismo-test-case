#!/bin/bash
echo "🚀 Building and starting Pismo Tech Case API..."

if ! docker info > /dev/null 2>&1; then
  echo "Erro: Docker não está rodando. Inicie o Docker Desktop."
  exit 1
fi

docker build -t pismo-api . && docker run --rm -p 8080:8080 pismo-api