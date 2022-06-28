// Copyright 2015 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This is highly based on the connection pooler found within
// Google's SDK. It's highly modified to stand in as our SDK and
// has nothing Google specific, all PlanetScale specific.
package pool

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/encoding"
	"google.golang.org/grpc/keepalive"

	"github.com/planetscale/psdb/core/pool/options"
	"github.com/planetscale/psdb/core/resolver"
)

// Slightly larger than the max message size allowed by vtgate.
const maxMessageSize = 100 * 1024 * 1024

func Dial(ctx context.Context, addr string, opts ...options.ClientOption) (ConnPool, error) {
	return dialPool(ctx, addr, processOpts(opts))
}

func dialPool(ctx context.Context, addr string, o *options.ClientOptions) (ConnPool, error) {
	if o.ConnPoolSize == 0 || o.ConnPoolSize == 1 {
		// Fast path for common case for a connection pool with a single connection.
		conn, err := dial(ctx, addr, o)
		if err != nil {
			return nil, err
		}
		return &singleConnPool{conn}, nil
	}

	pool := &roundRobinConnPool{}
	for i := 0; i < o.ConnPoolSize; i++ {
		conn, err := dial(ctx, addr, o)
		if err != nil {
			defer pool.Close() // NOTE: error from Close is ignored.
			return nil, err
		}
		pool.conns = append(pool.conns, conn)
	}
	return pool, nil
}

func dial(ctx context.Context, addr string, o *options.ClientOptions) (*grpc.ClientConn, error) {
	if colonPos := strings.LastIndex(addr, ":"); colonPos == -1 {
		if o.TLSConfig != nil {
			addr += ":443"
		} else {
			addr += ":80"
		}
	}

	dialOpts := []grpc.DialOption{
		// always insecure, because we handle the dialer and TLS wrapping
		// ourselves
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithContextDialer(newDialer(addr, o.TLSConfig)),
	}
	callOpts := []grpc.CallOption{
		grpc.MaxCallRecvMsgSize(maxMessageSize),
		grpc.MaxCallSendMsgSize(maxMessageSize),
	}

	if len(o.ExtraCallOptions) > 0 {
		callOpts = append(callOpts, o.ExtraCallOptions...)
	}

	if o.UseCompression {
		if encoding.GetCompressor("snappy") == nil {
			return nil, errors.New("snappy compressor not installed")
		}
		callOpts = append(callOpts, grpc.UseCompressor("snappy"))
	}

	if o.UseRoundRobin {
		dialOpts = append(dialOpts, grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin":{}}]}`))
	}

	if o.KeepaliveTime != 0 || o.KeepaliveTimeout != 0 {
		dialOpts = append(dialOpts, grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                o.KeepaliveTime,
			Timeout:             o.KeepaliveTimeout,
			PermitWithoutStream: true,
		}))
	}

	if o.AddrRefreshInterval > 0 || o.AddrPrefix != "" {
		addr = fmt.Sprintf("%s:///%s?refresh=%s&prefix=%s", resolver.SchemeName, addr, o.AddrRefreshInterval, o.AddrPrefix)
	}

	dialOpts = append(dialOpts, grpc.WithDefaultCallOptions(callOpts...))

	if len(o.ExtraDialOptions) > 0 {
		dialOpts = append(dialOpts, o.ExtraDialOptions...)
	}

	return grpc.DialContext(ctx, addr, dialOpts...)
}

func newDialer(addr string, config *tls.Config) func(context.Context, string) (net.Conn, error) {
	if config != nil {
		if !config.InsecureSkipVerify && config.ServerName == "" {
			// Clone() before we mutate
			config = config.Clone()

			// we assure that there is a colon before we pass `addr` in here.
			hostname := addr[:strings.LastIndex(addr, ":")]
			config.ServerName = hostname
		}
	}
	return func(ctx context.Context, resolvedAddr string) (net.Conn, error) {
		var d net.Dialer
		rawConn, err := d.DialContext(ctx, "tcp", resolvedAddr)
		if err != nil {
			return nil, err
		}
		if config == nil {
			return rawConn, nil
		}

		conn := tls.Client(rawConn, config)
		if err := conn.HandshakeContext(ctx); err != nil {
			return nil, err
		}
		return conn, nil
	}
}

func processOpts(opts []options.ClientOption) *options.ClientOptions {
	o := options.DefaultClientOptions()
	for _, opt := range opts {
		opt.Apply(o)
	}

	return o
}
