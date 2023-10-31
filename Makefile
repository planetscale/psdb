gomod := github.com/planetscale/psdb

PROTO_SRC := proto-src
PROTO_OUT := types

proto_pkgs := psdb/data/v1 \
			  psdb/v1alpha1 \
			  psdbconnect/v1alpha1

proto_files := $(wildcard $(addsuffix /*.proto,$(addprefix $(PROTO_SRC)/,$(proto_pkgs))))

BIN := bin

clean: clean-proto clean-bin

clean-proto:
	rm -rf $(PROTO_OUT)

clean-bin:
	rm -rf $(BIN)

$(BIN):
	mkdir -p $(BIN)

TOOL_INSTALL := env GOBIN=$(PWD)/$(BIN) go install

$(BIN)/protoc-gen-go: go.mod | $(BIN)
	$(TOOL_INSTALL) google.golang.org/protobuf/cmd/protoc-gen-go

$(BIN)/protoc-gen-go-vtproto: Makefile | $(BIN)
	$(TOOL_INSTALL) github.com/planetscale/vtprotobuf/cmd/protoc-gen-go-vtproto@v0.5.0

$(BIN)/protoc-gen-connect-go: go.mod | $(BIN)
	$(TOOL_INSTALL) connectrpc.com/connect/cmd/protoc-gen-connect-go

$(BIN)/gofumpt: Makefile | $(BIN)
	$(TOOL_INSTALL) mvdan.cc/gofumpt@v0.5.0

$(BIN)/staticcheck: Makefile | $(BIN)
	$(TOOL_INSTALL) honnef.co/go/tools/cmd/staticcheck@v0.4.6

$(BIN)/enumcheck: Makefile | $(BIN)
	$(TOOL_INSTALL) loov.dev/enumcheck@v0.0.0-20220314183541-8aa7b787306e

$(BIN)/govulncheck: Makefile | $(BIN)
	$(TOOL_INSTALL) golang.org/x/vuln/cmd/govulncheck@v1.0.1

$(BIN)/buf: Makefile | $(BIN)
	$(TOOL_INSTALL) github.com/bufbuild/buf/cmd/buf@v1.28.1

$(BIN)/yq: Makefile | $(BIN)
	$(TOOL_INSTALL) github.com/mikefarah/yq/v4@v4.30.8

PROTO_TOOLS := \
			   $(BIN)/protoc-gen-go \
			   $(BIN)/protoc-gen-connect-go \
			   $(BIN)/protoc-gen-go-vtproto \
			   $(BIN)/buf

tools: \
	$(PROTO_TOOLS) \
	$(BIN)/gofumpt \
	$(BIN)/staticcheck \
	$(BIN)/enumcheck \
	$(BIN)/govulncheck \
	$(BIN)/yq

proto: $(PROTO_OUT)/.done
$(PROTO_OUT)/.done: $(BIN)/buf $(PROTO_TOOLS) $(proto_files)
	$< generate -v $(PROTO_SRC)
	@touch $@

fmt: fmt-go fmt-proto

fmt-go: $(BIN)/gofumpt
	$< -l -w .

fmt-proto: $(BIN)/buf
	$< format -w $(PROTO_SRC)

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
	$< lint -v $(PROTO_SRC)

tests:
	go test -v ./...

update-proto: $(BIN)/buf
	$< mod update $(PROTO_SRC)

push-proto: $(BIN)/buf
	$< push $(PROTO_SRC)

update: update-proto
	go get -v -u ./...
	go mod tidy
	$(MAKE) clean proto

.PHONY: proto tools update update-proto push-proto \
		clean clean-proto clean-bin \
		fmt fmt-go fmt-proto fmt-yaml \
		lint lint-vet lint-staticcheck lint-enumcheck lint-govulncheck lint-proto \
		tests
