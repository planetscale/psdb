gomod := github.com/planetscale/psdb

PROTO_OUT := types
PSDB_OUT := $(PROTO_OUT)/psdb
PSDBDATA_OUT := $(PSDB_OUT)/data
PSDBCONNECT_OUT := $(PROTO_OUT)/psdbconnect

PSDBDATAV1_OUT := $(PSDBDATA_OUT)/v1
PSDBV1ALPHA1_OUT := $(PSDB_OUT)/v1alpha1
PSDBCONNECTV1_OUT := $(PSDBCONNECT_OUT)/v1

PROTO_SRC := proto-src
PSDB_SRC := $(PROTO_SRC)/psdb
PSDBDATA_SRC := $(PSDB_SRC)/data
PSDBCONNECT_SRC := $(PROTO_SRC)/psdbconnect

PSDBDATAV1_SRC := $(PSDBDATA_SRC)/v1
PSDBV1ALPHA1_SRC := $(PSDB_SRC)/v1alpha1
PSDBCONNECTV1_SRC := $(PSDBCONNECT_SRC)/v1

BIN := bin

clean: clean-proto clean-bin

clean-proto:
	rm -rf $(PROTO_OUT)

clean-bin:
	rm -rf $(BIN)

$(BIN):
	mkdir -p $(BIN)

$(PROTO_OUT):
	mkdir -p $(PROTO_OUT)

TOOL_INSTALL := env GOBIN=$(PWD)/$(BIN) go install

$(BIN)/protoc-gen-go: Makefile | $(BIN)
	$(TOOL_INSTALL) google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1

$(BIN)/protoc-gen-go-vtproto: Makefile | $(BIN)
	$(TOOL_INSTALL) github.com/planetscale/vtprotobuf/cmd/protoc-gen-go-vtproto@v0.4.0

$(BIN)/protoc-gen-connect-go: Makefile | $(BIN)
	$(TOOL_INSTALL) github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go@v1.5.1

$(BIN)/gofumpt: Makefile | $(BIN)
	$(TOOL_INSTALL) mvdan.cc/gofumpt@v0.4.0

$(BIN)/staticcheck: Makefile | $(BIN)
	$(TOOL_INSTALL) honnef.co/go/tools/cmd/staticcheck@v0.3.3

$(BIN)/enumcheck: Makefile | $(BIN)
	$(TOOL_INSTALL) loov.dev/enumcheck@v0.0.0-20220314183541-8aa7b787306e

$(BIN)/govulncheck: Makefile | $(BIN)
	$(TOOL_INSTALL) golang.org/x/vuln/cmd/govulncheck@v0.0.0-20230110180137-6ad3e3d07815

$(BIN)/buf: Makefile | $(BIN)
	$(TOOL_INSTALL) github.com/bufbuild/buf/cmd/buf@v1.14.0

$(BIN)/yq: Makefile | $(BIN)
	$(TOOL_INSTALL) github.com/mikefarah/yq/v4@v4.30.8

PROTO_TOOLS := $(BIN)/protoc-gen-go $(BIN)/protoc-gen-connect-go $(BIN)/protoc-gen-go-vtproto $(BIN)/buf
tools: $(PROTO_TOOLS) $(BIN)/gofumpt $(BIN)/staticcheck $(BIN)/enumcheck $(BIN)/govulncheck $(BIN)/yq

proto: \
	$(PSDBCONNECTV1_OUT)/connect.pb.go \
	$(PSDBDATAV1_OUT)/data.pb.go \
	$(PSDBV1ALPHA1_OUT)/database.pb.go

$(PSDBCONNECTV1_OUT)/connect.pb.go: $(PROTO_TOOLS) $(PSDBCONNECTV1_SRC)/connect.proto | $(PROTO_OUT)
	$(BIN)/buf generate --path $(PSDBCONNECTV1_SRC)/connect.proto

$(PSDBDATAV1_OUT)/data.pb.go: $(PROTO_TOOLS) $(PSDBDATAV1_SRC)/data.proto | $(PROTO_OUT)
	$(BIN)/buf generate --path $(PSDBDATAV1_SRC)/data.proto

$(PSDBV1ALPHA1_OUT)/database.pb.go: $(PROTO_TOOLS) $(PSDBV1ALPHA1_SRC)/database.proto | $(PROTO_OUT)
	$(BIN)/buf generate --path $(PSDBV1ALPHA1_SRC)/database.proto

fmt: fmt-go fmt-proto

fmt-go: $(BIN)/gofumpt
	$(BIN)/gofumpt -l -w .

fmt-proto: $(BIN)/buf
	$(BIN)/buf format -w proto-src

fmt-yaml: $(BIN)/yq
ifeq (, $(shell command -v fd 2>/dev/null))
	@echo "!! Maybe install 'fd', it's a lot faster (https://github.com/sharkdp/fd)"
	find . -type f \( -name '*.yaml' -o -name '*.yml' \) -exec $(BIN)/yq -iP eval-all . {} \;
else
	fd . -t f -e yaml -e yml -x $(BIN)/yq -iP eval-all . {} \;
endif

lint: lint-vet lint-staticcheck lint-enumcheck lint-govulncheck lint-proto

lint-vet:
	go vet ./...

lint-staticcheck: $(BIN)/staticcheck
	$(BIN)/staticcheck -f=stylish ./...

lint-enumcheck: $(BIN)/enumcheck
	$(BIN)/enumcheck ./...

lint-govulncheck: $(BIN)/govulncheck
	$(BIN)/govulncheck ./...

lint-proto: $(BIN)/buf
	$(BIN)/buf lint -v proto-src

tests:
	go test -v ./...

update:
	go get -v -u ./...
	go mod tidy
	$(MAKE) clean proto

.PHONY: proto tools update \
		clean clean-proto clean-bin \
		fmt fmt-go fmt-proto fmt-yaml \
		lint lint-vet lint-staticcheck lint-enumcheck lint-govulncheck lint-proto \
		tests
