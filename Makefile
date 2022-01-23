
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

	# buf install https://docs.buf.build/installation#github-releases
	brew tap bufbuild/buf
	brew install buf

	brew install jq

.PHONY: test
test:
	docker-compose up --build -d
