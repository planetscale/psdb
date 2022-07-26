package client

import (
	"context"

	"github.com/bufbuild/connect-go"
)

type setHeadersInterceptor struct {
	key, value string
}

func (i *setHeadersInterceptor) WrapUnary(next connect.UnaryFunc) connect.UnaryFunc {
	return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
		if !req.Spec().IsClient {
			return next(ctx, req)
		}
		req.Header()[i.key] = []string{i.value}
		return next(ctx, req)
	}
}

func (i *setHeadersInterceptor) WrapStreamingClient(next connect.StreamingClientFunc) connect.StreamingClientFunc {
	return func(ctx context.Context, spec connect.Spec) connect.StreamingClientConn {
		conn := next(ctx, spec)
		conn.RequestHeader()[i.key] = []string{i.value}
		return conn
	}
}

func (*setHeadersInterceptor) WrapStreamingHandler(next connect.StreamingHandlerFunc) connect.StreamingHandlerFunc {
	return next
}
