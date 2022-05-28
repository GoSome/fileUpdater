#!/usr/bin/env bash
# build vue for further use
go get github.com/GeertJohan/go.rice/rice
go install github.com/GeertJohan/go.rice/rice
cd vue
npm install
npm run build
echo "move file for go binding data"
mkdir -pv ../pkg/binding/dist/
mv -f dist/statics/{css,fonts,img,js} ../pkg/binding/dist/
mv dist/{index.html,favicon.ico} ../pkg/binding/dist/
echo "binding data"
cd ../pkg/binding
rice embed-go