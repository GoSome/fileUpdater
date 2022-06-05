#!/usr/bin/env bash

set -e
echo "go get"
go get -v ./...

# ui
cd ui
npm install
npm run build

echo "move file for go binding data"
rm -rf ../cmd/build
mv build ../cmd

# cmd
echo "build server"
cd ../cmd
GOOS=darwin GOARCH=amd64 go build -o ../bin/fileupdater-amd64-darwin
GOOS=darwin GOARCH=arm64 go build -o ../bin/fileupdater-arm64-darwin
echo "clean statics"

rm -rf build
