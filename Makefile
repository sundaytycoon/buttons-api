
# Install protoc compiler
PROTOC_REPO=https://github.com/protocolbuffers/protobuf
PROTOC_VER=3.14.0

PROTOC_GEN_GO_VER=v1.25.0
PROTOC_GEN_GO_GRPC_VER=v1.1
PROTOC_GEN_DOC=v1.5.0
PROTOC_GEN_GATEWAY=v2.6.0
PROTOC_GEN_OPENAPI2=v2.6.0
PROTOC_ZIP=protoc-${PROTOC_VER}-osx-x86_64.zip
USER=`whoami`



#		--ignoreprefixes \
#			ent \
#		-maxlevel 10\


.PHONY: generate-docs
generate-docs:
	# source code dependency graph // source code visualization utility
	godepgraph -s -o github.com/sundaytycoon/buttons-api \
		-novendor \
		-ignoreprefixes github.com/sundaytycoon/buttons-api/ent \
		-ignorepackages github.com/sundaytycoon/buttons-api/pkg/er,github.com/sundaytycoon/buttons-api/pkg/testdockercontainer \
		github.com/sundaytycoon/buttons-api/internal/handler/user \
		| dot -Tpng -o doc/_images/godepgraph.png

.PHONY: swagger-ui-gen
swagger-ui-gen:
	mkdir -p ./doc/OpenAPI
	curl -o ./doc/OpenAPI/swaagerui.tar.gz -L https://github.com/swagger-api/swagger-ui/archive/refs/tags/v4.1.3.tar.gz
	tar -xf ./doc/OpenAPI/swaagerui.tar.gz  -C ./doc/OpenAPI
	mv ./doc/OpenAPI/swagger-ui-4.1.3/dist/* ./doc/OpenAPI
	rm -rf ./doc/OpenAPI/swaagerui.tar.gz
	rm -rf ./doc/OpenAPI/swagger-ui-4.1.3
	echo "You have to modify 'doc/OpenAPI/index.html' for looking generated protobuf OAS(Open Api Swagger)"
	echo "You have to modify 'doc/OpenAPI/index.html' for looking generated protobuf OAS(Open Api Swagger)"
	echo "You have to modify 'doc/OpenAPI/index.html' for looking generated protobuf OAS(Open Api Swagger)"


.PHONY: protogen
protogen:
	rm -rf gen
	mkdir -p ./gen/go
	buf build -o -#format=json | jq '.file[] | .package' | sort | uniq | head
	buf generate

.PHONY: protolint
protolint:
	buf breaking --against 'https://github.com/sundaytycoon/buttons-api.git#branch=main'
	buf lint

.PHONY: protosetup
protosetup:
#	curl -OL ${PROTOC_REPO}/releases/download/v${PROTOC_VER}/${PROTOC_ZIP}
#	sudo unzip -o ${PROTOC_ZIP} -d /usr/local bin/protoc
#	sudo unzip -o ${PROTOC_ZIP} -d /usr/local 'include/*'
#	sudo chown -R ${USER}:staff /usr/local/include
#	rm -f ${PROTOC_ZIP}
#
#	# Install protoc-gen-go plugin
#	go install google.golang.org/protobuf/cmd/protoc-gen-go@${PROTOC_GEN_GO_VER}
#	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@${PROTOC_GEN_GO_GRPC_VER}
#	go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@${PROTOC_GEN_DOC}
#
#	# for gateway
#	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@${PROTOC_GEN_GATEWAY}
#	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@${PROTOC_GEN_OPENAPI2}

	# buf install https://docs.buf.build/installation#github-releases
	brew tap bufbuild/buf
	brew install buf

	brew install jq
