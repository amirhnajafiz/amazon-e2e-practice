#!/bin/bash

# Stop and remove containers
docker stop e2e-practice-frontend e2e-practice-api postgres
docker rm e2e-practice-frontend e2e-practice-api postgres

# Remove docker network
docker network rm e2e-close-network
