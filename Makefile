LWDP_SRC_MAIN_PROTO_DIR=grpc/src/main/proto
COMMON_SRC_MAIN_PROTO_DIR=src/main/proto
GITHUB_TH2=github.com/th2-net

TH2_LW_DATA_PROVIDER=th2-lw-data-provider
TH2_LW_DATA_PROVIDER_URL=$(GITHUB_TH2)/$(TH2_LW_DATA_PROVIDER)@go_package # TODO: replace to a tag after submit PR https://github.com/th2-net/th2-lw-data-provider/pull/91
TH2_GRPC_COMMON=th2-grpc-common
TH2_GRPC_COMMON_URL=$(GITHUB_TH2)/$(TH2_GRPC_COMMON)@go_package # TODO: replace to a tag after submit PR https://github.com/th2-net/th2-lw-data-provider/pull/91

MODULE_DIR=th2-grpc

PROTOC_VERSION=21.12

init-work-space: clean-grpc-module prepare-grpc-module configure-grpc-generator generate-grpc-files tidy

configure-grpc-generator:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.5
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1

clean-grpc-module:
	-rm -rf $(MODULE_DIR)

prepare-grpc-module:
	- mkdir $(MODULE_DIR)
	cd $(MODULE_DIR)

	cd $(MODULE_DIR) \
		&& go get $(TH2_LW_DATA_PROVIDER_URL) \
		&& go get $(TH2_GRPC_COMMON_URL) \
		&& go get google.golang.org/protobuf@v1.36.5 \
		&& go get github.com/google/go-cmp@v0.6.0

generate-grpc-files:
	$(eval $@_LWDP_PROTO_DIR := $(shell go list -m -f '{{.Dir}}' $(TH2_LW_DATA_PROVIDER_URL))/$(LWDP_SRC_MAIN_PROTO_DIR))
	$(eval $@_COMMON_PROTO_DIR := $(shell go list -m -f '{{.Dir}}' $(TH2_GRPC_COMMON_URL))/$(COMMON_SRC_MAIN_PROTO_DIR))
	protoc \
		--go_out=$(MODULE_DIR) \
		--go-grpc_out=$(MODULE_DIR) \
		--go_opt=module=github.com/th2-net/th2-lw-data-provider/grpc/src/main/proto \
		--go-grpc_opt=module=github.com/th2-net/th2-lw-data-provider/grpc/src/main/proto \
		--proto_path=$($@_COMMON_PROTO_DIR) \
		--proto_path=$($@_LWDP_PROTO_DIR) \
		$(shell find $($@_LWDP_PROTO_DIR) -name '*.proto' )

tidy:
	go mod tidy -v

build:
	go vet ./...
	go build -v ./...

run-test:
	go test -v -race ./...