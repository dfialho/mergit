#!/bin/bash

# Build script for Windows binary (cross-compile from Linux)

set -e

echo "Building mergit for Windows..."

cd "$(dirname "$0")/.."

# Check if mingw-w64 is installed
if ! command -v x86_64-w64-mingw32-gcc &> /dev/null; then
    echo "Error: mingw-w64 is not installed"
    echo "Install it with: sudo apt-get install gcc-mingw-w64"
    exit 1
fi

GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc \
  go build -ldflags="-s -w -H windowsgui" -o build/mergit.exe

if [ $? -eq 0 ]; then
    echo "✓ Windows binary created: build/mergit.exe"
    ls -lh build/mergit.exe
    echo ""
    echo "To run on Windows: Double-click mergit.exe"
else
    echo "✗ Build failed"
    echo "Make sure mingw-w64 is installed: sudo apt-get install gcc-mingw-w64"
    exit 1
fi
