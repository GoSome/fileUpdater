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
GOOS=linux GOARCH=amd64 go build -o ../bin/fileupdater-amd64-linux
GOOS=linux GOARCH=arm64 go build -o ../bin/fileupdater-arm64-linux
echo "clean statics"

rm -rf build
