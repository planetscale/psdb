//go:build tools
// +build tools

package tools

import (
	_ "github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go"
	_ "github.com/planetscale/vtprotobuf/cmd/protoc-gen-go-vtproto"
	_ "honnef.co/go/tools/cmd/staticcheck"
	_ "loov.dev/enumcheck"
	_ "mvdan.cc/gofumpt"
)
