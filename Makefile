# Makefile.common
BASE_DIR := $(shell pwd)

.PHONY: check_requirements
check_requirements:
	./scripts/check_requirements.sh

.PHONY: gen
gen: gen-codes gen-docs go-codetest


.PHONY: gen-codes
gen-codes:
	oapi-codegen -package v1 -generate types,chi-server,spec api/oapi/v1/v1.yaml > api/oapi/v1/v1.oapi.go
	buf build -o -#format=json | jq '.file[] | .package' | sort | uniq | head
	buf generate
	go generate ./internal/storage/servicedb/ent
	go generate ./internal/domain/repository/...
	go generate ./internal/adapter/...


.PHONY: gen-docs
gen-docs:
	# source code dependency graph // source code visualization utility
	godepgraph -s -o github.com/sundaytycoon/buttons-api \
		-novendor \
		-ignoreprefixes github.com/sundaytycoon/buttons-api/internal/storage/servicedb/ent \
		-ignorepackages \
github.com/sundaytycoon/buttons-api\
,github.com/sundaytycoon/buttons-api/internal/utils/er\
,github.com/sundaytycoon/buttons-api/internal/utils/retry\
,github.com/sundaytycoon/buttons-api/internal/utils/unittestdocker\
,github.com/sundaytycoon/buttons-api/internal/utils/recovery\
		github.com/sundaytycoon/buttons-api/cmd/http_server \
		| dot -Tpng -o doc/_images/godepgraph.png

.PHONY: test
test: go-codetest

.PHONY: go-codetest
go-codetest:
	# buf breaking --against 'https://github.com/sundaytycoon/buttons-api.git#branch=main'
	# buf lint
	go fmt ./...
	go vet ./...
	go test ./...

.PHONY: docker-up
docker-up:
	docker-compose -f docker-compose.yml up --build --remove-orphans -d
