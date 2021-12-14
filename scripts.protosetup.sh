#!/bin/bash


set -x

# Install protoc compiler
PROTOC_REPO=https://github.com/protocolbuffers/protobuf
PROTOC_VER=3.14.0

PROTOC_GEN_GO_VER=v1.25.0
PROTOC_GEN_GO_GRPC_VER=v1.1
PROTOC_GEN_DOC=v1.5.0
PROTOC_GEN_GATEWAY=v2.6.0
PROTOC_GEN_OPENAPI2=v2.6.0

PROTOC_ZIP=protoc-$PROTOC_VER-osx-x86_64.zip

USER=`whoami`
curl -OL $PROTOC_REPO/releases/download/v$PROTOC_VER/$PROTOC_ZIP
sudo unzip -o $PROTOC_ZIP -d /usr/local bin/protoc
sudo unzip -o $PROTOC_ZIP -d /usr/local 'include/*'
sudo chown -R ${USER}:staff /usr/local/include

rm -f $PROTOC_ZIP

# Install protoc-gen-go plugin
go install google.golang.org/protobuf/cmd/protoc-gen-go@${PROTOC_GEN_GO_VER}
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@${PROTOC_GEN_GO_GRPC_VER}
go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@${PROTOC_GEN_DOC}

# for gateway
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@${PROTOC_GEN_GATEWAY}
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@${PROTOC_GEN_OPENAPI2}
