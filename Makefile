
VERSION ?= $(shell git describe --tags --abbrev=0)
SHA1 ?= $(shell git rev-parse HEAD)
BUILD = $(shell date +%FT%T%z)

LDFLAGS=-ldflags \
	'-X github.com/rudeigerc/dexctl/pkg/cmd/version.gitVersion=$(VERSION) \
	-X github.com/rudeigerc/dexctl/pkg/cmd/version.gitCommit=$(SHA1) \
	-X github.com/rudeigerc/dexctl/pkg/cmd/version.buildDate=$(BUILD)'

.PHONY: all
all: build

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: generate
generate: ## Run protoc to generate gRPC code from proto files.
	protoc --go_out=paths=source_relative:. \
		--go-grpc_out=paths=source_relative:. \
		pkg/protos/dex.proto

.PHONY: fmt
fmt: ## Run go fmt against code.
	go fmt ./...

.PHONY: vet
vet: ## Run go vet against code.
	go vet ./...

.PHONY: build
build: generate fmt vet ## Build manager binary.
	go build ${LDFLAGS} -o bin/dexctl cmd/dexctl/main.go

.PHONY: clean
clean:
	@rm -rf bin/ pkg/protos/*.pb.go
