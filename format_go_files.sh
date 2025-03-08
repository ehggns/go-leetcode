#!/bin/bash

# Script to recursively find all main.go files and apply gofmt to them

echo "Starting to format main.go files..."

# Find all main.go files recursively and apply gofmt to each
find . -type f -name "main.go" -exec sh -c 'echo "Formatting $1"; gofmt -w "$1"' _ {} \;

echo "Formatting complete!"