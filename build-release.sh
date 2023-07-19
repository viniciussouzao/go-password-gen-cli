#!/bin/bash

VERSION=$1
NOTE=$2

if [ -z "$VERSION" ]; then
  echo "Version is required"
  exit 1
fi

# Build binaries
echo "Building binaries..."
go build -o bin/password-gen-linux-amd64 main.go
GOOS=windows GOARCH=amd64 go build -o bin/password-gen-windows-amd64.exe main.go
GOOS=darwin GOARCH=amd64 go build -o bin/password-gen-darwin-amd64 main.go

# Create release
echo "Creating release..."
gh release create $VERSION \
  --title "$VERSION" \
  --notes "$NOTE" \
  bin/password-gen-linux-amd64 \
  bin/password-gen-windows-amd64.exe \
  bin/password-gen-darwin-amd64