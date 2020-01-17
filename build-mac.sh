#!/usr/bin/env bash

echo "go get"
go get github.com/GeertJohan/go.rice/rice
go get -v ./...
set -e
echo "clean old files"
rm -rf pkg/binding/{css,fonts,img,js,index.html,favicon.ico}
echo "build vue"

# vue
cd vue
npm install
npm run build

echo "move file for go binding data"
mkdir -pv ../pkg/binding/dist/
mv -f dist/statics/{css,fonts,img,js} ../pkg/binding/dist/
mv dist/{index.html,favicon.ico} ../pkg/binding/dist/

# pkg/binding

echo "binding data"
cd ../pkg/binding
rice embed-go

# cmd
echo "build server"
cd ../../cmd
GOOS=darwin GOARCH=amd64 go build -o ../bin/fileupdater-amd64-darwin
echo "clean statics"

cd ..
rm -rf pkg/binding/dist
rm -rf vue/dist