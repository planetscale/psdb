gomod := github.com/planetscale/psdb

PSDB_PROTO_OUT := types
PSDB_PROTO_ROOT := $(PSDB_PROTO_OUT)/psdb
PSDB_DATA_V1 := $(PSDB_PROTO_ROOT)/data/v1
PSDB_V1ALPHA1 := $(PSDB_PROTO_ROOT)/v1alpha1

PROTOC_VERSION=21.3

BIN := bin

UNAME_OS := $(shell uname -s)
UNAME_ARCH := $(shell uname -m)

proto: \
	$(PSDB_DATA_V1)/data.pb.go \
	$(PSDB_V1ALPHA1)/database.pb.go

clean: clean-proto clean-bin

clean-proto:
	rm -rf $(PSDB_PROTO_OUT)

clean-bin:
	rm -rf $(BIN)

$(BIN):
	mkdir -p $(BIN)

$(PSDB_PROTO_OUT):
	mkdir -p $(PSDB_PROTO_OUT)

TOOL_INSTALL := cd tools && env GOBIN=$(PWD)/$(BIN) go install

$(BIN)/protoc-gen-go: | $(BIN)
	$(TOOL_INSTALL) google.golang.org/protobuf/cmd/protoc-gen-go

$(BIN)/protoc-gen-connect-go: | $(BIN)
	$(TOOL_INSTALL) github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go

$(BIN)/protoc-gen-go-vtproto: | $(BIN)
	$(TOOL_INSTALL) github.com/planetscale/vtprotobuf/cmd/protoc-gen-go-vtproto

$(BIN)/gofumpt: | $(BIN)
	$(TOOL_INSTALL) mvdan.cc/gofumpt

$(BIN)/staticcheck: | $(BIN)
	$(TOOL_INSTALL) honnef.co/go/tools/cmd/staticcheck

$(BIN)/enumcheck: | $(BIN)
	$(TOOL_INSTALL) loov.dev/enumcheck

ifeq ($(UNAME_OS),Darwin)
PROTOC_OS := osx
ifeq ($(UNAME_ARCH),arm64)
PROTOC_ARCH := aarch_64
else
PROTOC_ARCH := x86_64
endif
endif
ifeq ($(UNAME_OS),Linux)
PROTOC_OS = linux
ifeq ($(UNAME_ARCH),aarch64)
PROTOC_ARCH := aarch_64
else
PROTOC_ARCH := $(UNAME_ARCH)
endif
endif

$(BIN)/protoc: | $(BIN)
	rm -rf tmp-protoc
	mkdir -p tmp-protoc
	wget -O tmp-protoc/protoc.zip https://github.com/protocolbuffers/protobuf/releases/download/v$(PROTOC_VERSION)/protoc-$(PROTOC_VERSION)-$(PROTOC_OS)-$(PROTOC_ARCH).zip
	unzip -d tmp-protoc tmp-protoc/protoc.zip
	mv tmp-protoc/bin/protoc $(BIN)/
	rm -rf tmp-protoc

PROTO_TOOLS := $(BIN)/protoc $(BIN)/protoc-gen-go $(BIN)/protoc-gen-connect-go $(BIN)/protoc-gen-go-vtproto
tools: $(PROTO_TOOLS) $(BIN)/gofumpt $(BIN)/staticcheck $(BIN)/enumcheck

$(PSDB_DATA_V1)/data.pb.go: $(PROTO_TOOLS) proto-src/psdb/data/v1/data.proto | $(PSDB_PROTO_OUT)
	$(BIN)/protoc \
	  --plugin=protoc-gen-go=$(BIN)/protoc-gen-go \
	  --plugin=protoc-gen-go-vtproto=$(BIN)/protoc-gen-go-vtproto \
	  --go_out=$(PSDB_PROTO_OUT) \
	  --go-vtproto_out=$(PSDB_PROTO_OUT) \
	  --go_opt=paths=source_relative \
	  --go-vtproto_opt=features=marshal+unmarshal+size \
	  --go-vtproto_opt=paths=source_relative \
	  -I proto-src \
	  proto-src/psdb/data/v1/data.proto

$(PSDB_V1ALPHA1)/database.pb.go: $(PROTO_TOOLS) proto-src/psdb/v1alpha1/database.proto | $(PSDB_PROTO_OUT)
	$(BIN)/protoc \
	  --plugin=protoc-gen-go=$(BIN)/protoc-gen-go \
	  --plugin=protoc-gen-go-vtproto=$(BIN)/protoc-gen-go-vtproto \
	  --plugin=protoc-gen-connect-go=$(BIN)/protoc-gen-connect-go \
	  --go_out=$(PSDB_PROTO_OUT) \
	  --go-vtproto_out=$(PSDB_PROTO_OUT) \
	  --connect-go_out=$(PSDB_PROTO_OUT) \
	  --go_opt=paths=source_relative \
	  --go-vtproto_opt=features=marshal+unmarshal+size \
	  --go-vtproto_opt=paths=source_relative \
	  --connect-go_opt=paths=source_relative \
	  -I proto-src \
	  proto-src/psdb/v1alpha1/database.proto

fmt: fmt-go

fmt-go: $(BIN)/gofumpt
	$(BIN)/gofumpt -l -w .

lint: lint-vet lint-staticcheck lint-enumcheck

lint-vet:
	go vet ./...

lint-staticcheck: $(BIN)/staticcheck
	$(BIN)/staticcheck -f=stylish ./...

lint-enumcheck: $(BIN)/enumcheck
	$(BIN)/enumcheck ./...

tests:
	go test -v ./...

update:
	go get -v -u ./...
	go mod tidy
	cd tools && go get -v -u ./internal
	cd tools && go mod tidy
	$(MAKE) clean proto

.PHONY: proto tools update \
		clean clean-proto clean-bin \
		fmt fmt-go \
		lint lint-vet lint-staticcheck lint-enumcheck \
		tests
