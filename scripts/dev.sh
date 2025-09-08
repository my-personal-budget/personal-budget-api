#!/usr/bin/env bash
set -e

# Ensure Go is installed
if ! command -v go &> /dev/null; then
    echo "Go is not installed. Please install Go first."
    exit 1
fi

# Ensure GOPATH/bin is in PATH
if [[ ":$PATH:" != *":$(go env GOPATH)/bin:"* ]]; then
    export PATH="$PATH:$(go env GOPATH)/bin"
fi

# Check if air is installed
if ! command -v air &> /dev/null; then
    echo "Air not found. Installing..."
    go install github.com/air-verse/air@latest
    echo "Air installed."
fi

# Run air
echo "Starting air..."
air
