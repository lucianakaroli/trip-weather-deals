#!/bin/bash

set -e

IMAGE="ghcr.io/lucianakaroli/trip-weather-deals:latest"

echo "Baixando a imagem mais recente..."
docker pull $IMAGE

echo "Subindo a aplicação com Docker Compose..."
docker-compose up -d

echo "Deploy concluído com sucesso."
echo "Aplicação disponível em: http://localhost:8080/recommendations"