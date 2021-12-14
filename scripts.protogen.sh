#!/bin/bash

PROTO_IN=./proto
PROTO_OUT=gen/go

mkdir -p gen/go

for f in ${PROTO_IN}/**/**/*.proto
do
    echo ${f}
    protoc \
        -I=. \
        -I=.. \
        --go_opt paths=source_relative \
        --go-grpc_opt paths=source_relative \
        \
        --go_out ${PROTO_OUT} \
        --go-grpc_out ${PROTO_OUT} \
        --grpc-gateway_out ${PROTO_OUT} \
        \
        --grpc-gateway_opt logtostderr=true \
        --grpc-gateway_opt paths=source_relative \
        --grpc-gateway_opt generate_unbound_methods=true \
        \
        --doc_out . \
        --doc_opt html,protobuf.html \
        ${f}
done

