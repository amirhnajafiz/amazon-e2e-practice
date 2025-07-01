#!/bin/bash

# Build the Docker images
docker buildx bake -f build/docker-bake.hcl
