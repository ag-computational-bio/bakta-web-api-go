#!/bin/sh

rm -rf API
git clone git@github.com:ag-computational-bio/bakta-web-api.git

mkdir build

# Builds the go stubs from the protobuf files and places them in their directories
# Temporary solution until ci/cd is enabled 
cd bakta-web-api
git checkout main
protoc --proto_path=. --go-grpc_out=../build --go_out=../build --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative proto/api/*.proto
protoc -I . --grpc-gateway_out ../build --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative proto/api/*

protoc -I . --openapiv2_out ../build --openapiv2_opt logtostderr=true proto/api/*

cd ..

mv build/proto/api/Jobs.swagger.json swagger

go generate swagger/SwaggerGen.go
mv build/proto/api/*.go api

rm -rf bakta-web-api
rm -rf build