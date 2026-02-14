#!/bin/bash

# Build script for Linux binary

set -e

echo "Building mergit for Linux..."

cd "$(dirname "$0")/.."

go build -ldflags="-s -w" -o build/mergit

if [ $? -eq 0 ]; then
    echo "✓ Linux binary created: build/mergit"
    ls -lh build/mergit
    echo ""
    echo "To run: ./build/mergit"
else
    echo "✗ Build failed"
    exit 1
fi
