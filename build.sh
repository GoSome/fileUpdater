#!/usr/bin/env bash

set -e
echo "clean old files"
rm -rf pkg/binding/{css,fonts,img,js,index.html,favicon.ico}
echo "build vue"

cd vue
npm run build

echo "move file for go binding data"
mv -f dist/statics/{css,fonts,img,js} ../pkg/binding/
mv dist/{index.html,favicon.ico} ../pkg/binding/

echo "build server"
cd ../cmd
GOOS=linux GOARCH=amd64 go build -o ../bin/fileupdater-amd64-linux
echo "clean statics"

cd ..
rm -rf pkg/binding/{css,fonts,img,js,index.html,favicon.ico}
rm -rf vue/dist