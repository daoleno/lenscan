#!/bin/bash

# Set environment variables
REGISTRY="ghcr.io"
IMAGE_NAME="daoleno/lenscan"

# Ask for the tag
echo "Please enter the tag for the Docker image:"
read TAG

# Log in to the container registry
# echo "Logging in to the container registry..."
# echo "Please enter your GitHub Personal Access Token:"
# read -s GITHUB_TOKEN # Read the GitHub Token securely
# echo $GITHUB_TOKEN | docker login $REGISTRY -u daoleno --password-stdin

# Ask for the DATABASE_URL
echo "Please enter the DATABASE_URL:"
read DATABASE_URL

# Build the Docker image, passing DATABASE_URL as a build argument
echo "Building the Docker image..."
docker build --build-arg DATABASE_URL="$DATABASE_URL" -t $REGISTRY/$IMAGE_NAME:$TAG .

# Push the Docker image to the registry
echo "Pushing the Docker image..."
docker push $REGISTRY/$IMAGE_NAME:$TAG
# Log out from the container registry
# echo "Logging out from the container registry..."
# docker logout $REGISTRY
echo "Docker image build and push completed."
