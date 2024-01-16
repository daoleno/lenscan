#!/bin/bash

# Set environment variables
REGISTRY="ghcr.io"
IMAGE_NAME="daoleno/lenscan"
DUCKDB_PATH=lens-public-data/v2_polygon.db

# Ask for the tag
echo "Please enter the tag for the Docker image:"
read TAG

# Log in to the container registry
echo "Logging in to the container registry..."
echo "Please enter your GitHub Personal Access Token:"
read -s GITHUB_TOKEN # Read the GitHub Token securely
echo $GITHUB_TOKEN | docker login $REGISTRY -u daoleno --password-stdin

# Build the Docker image, add --build-arg if needed
echo "Building the Docker image..."
docker build -t $REGISTRY/$IMAGE_NAME:$TAG ./dashboard # Ensure the correct build context is used

# Push the Docker image to the registry
echo "Pushing the Docker image..."
docker push $REGISTRY/$IMAGE_NAME:$TAG

# Log out from the container registry
echo "Logging out from the container registry..."
docker logout $REGISTRY

echo "Docker image build and push completed."
