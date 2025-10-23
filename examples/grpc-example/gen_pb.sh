#!/bin/bash

GO_BIN="$(go env GOPATH)/bin"

if [ -f "$GO_BIN/protoc-gen-go" ]
then
  echo "protoc-gen-go ok"
else
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
fi

if [ -f "$GO_BIN/protoc-gen-go-grpc" ]
then
  echo "protoc-gen-go-grpc ok"
else
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
fi

#if [ -f "/usr/local/bin/protoc"]
#then
#  echo "protoc ok"
#else
#  curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v33.0/protoc-33.0-linux-aarch_64.zip
#fi

cd ./proto

for file in *.proto
do
  protoc --go_out=. --go_opt=paths=source_relative \
      --go-grpc_out=. --go-grpc_opt=paths=source_relative \
      $file
done



