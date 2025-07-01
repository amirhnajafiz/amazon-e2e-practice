#!/bin/bash

# Create docker network
docker network create --driver bridge e2e-close-network

# Start containers one by one
docker run  \
    --network e2e-close-network \
    -d \
    --name postgres \
    -e POSTGRES_PASSWORD=password \
    -p 5432:5432 \
    postgres
docker run  \
    --network e2e-close-network \
    -d \
    --name e2e-practice-api \
    -e AEP_revision=1 \
    -p 8080:8080 \
    amazon-e2e-practice-api:latest
docker run  \
    --network e2e-close-network \
    -d \
    --name e2e-practice-frontend \
    -p 3000:4173 \
    amazon-e2e-practice-frontend:latest
