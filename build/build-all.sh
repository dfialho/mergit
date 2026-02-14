#!/bin/bash

# Build script for both Linux and Windows binaries

set -e

echo "==================================="
echo "Building mergit for all platforms"
echo "==================================="
echo ""

cd "$(dirname "$0")"

./build-linux.sh
echo ""
./build-windows.sh

echo ""
echo "==================================="
echo "All builds completed successfully!"
echo "==================================="
