package client

import (
	"crypto/tls"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/bufbuild/connect-go"
	compress "github.com/klauspost/connect-compress"
	"github.com/planetscale/psdb/auth"
	"github.com/planetscale/psdb/core/codec"
)

// simple connect.HTTPClient implementation
// that doesn't support cookies, or redirects, or any baggage that comes
// from http.Client
type simpleClient struct {
	http.RoundTripper
}

func (c *simpleClient) Do(req *http.Request) (*http.Response, error) {
	return c.RoundTripper.RoundTrip(req)
}

const (
	defaultCompressionName  = compress.S2
	defaultCompressionLevel = compress.LevelFastest
	// Slightly larger than the max message size allowed by edge gateway.
	maxMessageSize = 100 * 1024 * 1024
)

func New[T any](
	addr string,
	fn func(connect.HTTPClient, string, ...connect.ClientOption) T,
	auth *auth.Authorization,
	opts ...Option,
) T {
	cfg := configFromOptions(opts...)

	cOpts := append(defaultClientOptions(), connect.WithInterceptors(
		&setHeadersInterceptor{
			"Authorization",
			auth.Type().String() + " " + auth.HeaderValue(),
		},
	))

	if len(cfg.extraClientOptions) > 0 {
		cOpts = append(cOpts, cfg.extraClientOptions...)
	}

	return fn(cfg.httpClient, "https://"+addr, cOpts...)
}

type ClientPool[T any] struct {
	fn            func(connect.HTTPClient, string, ...connect.ClientOption) T
	pool          map[string]T
	poolMu        sync.RWMutex
	cfg           *config
	clientOptions []connect.ClientOption
}

func (p *ClientPool[T]) Get(addr string) T {
	// First check the fast path, if there's already
	// an initialized client for this address
	p.poolMu.RLock()
	client, ok := p.pool[addr]
	p.poolMu.RUnlock()
	if ok {
		return client
	}

	// create a new one outside a lock, worst case, we create multiple, last one wins
	// and the others are GC'd
	client = p.fn(p.cfg.httpClient, "https://"+addr, p.clientOptions...)

	// lock the map to store the client
	p.poolMu.Lock()
	p.pool[addr] = client
	p.poolMu.Unlock()

	// finally, read from the map again to make sure we use the instance that
	// was actually written. This is because of the race condition above where multiple
	// instances were created but only one can be written into the map
	p.poolMu.RLock()
	client = p.pool[addr]
	p.poolMu.RUnlock()
	return client
}

func (p *ClientPool[T]) Release(addr string) {
	p.poolMu.Lock()
	delete(p.pool, addr)
	p.poolMu.Unlock()
}

func (p *ClientPool[T]) Len() int {
	p.poolMu.RLock()
	defer p.poolMu.RUnlock()
	return len(p.pool)
}

func NewUnauthenticatedPool[T any](
	fn func(connect.HTTPClient, string, ...connect.ClientOption) T,
	opts ...Option,
) *ClientPool[T] {
	cfg := configFromOptions(opts...)

	cOpts := defaultClientOptions()
	if len(cfg.extraClientOptions) > 0 {
		cOpts = append(cOpts, cfg.extraClientOptions...)
	}

	return &ClientPool[T]{
		fn:            fn,
		pool:          make(map[string]T),
		cfg:           cfg,
		clientOptions: cOpts,
	}
}

func clientCompressionOpt(name string, level compress.Level) connect.ClientOption {
	opt, _ := compress.Select(name, level)
	return opt
}

func defaultClientOptions() []connect.ClientOption {
	return []connect.ClientOption{
		clientCompressionOpt(defaultCompressionName, defaultCompressionLevel),
		connect.WithSendCompression(defaultCompressionName),
		connect.WithCodec(codec.DefaultCodec),
		connect.WithReadMaxBytes(maxMessageSize),
	}
}

func defaultHTTPClient(tlsConfig *tls.Config) connect.HTTPClient {
	if tlsConfig == nil {
		tlsConfig = DefaultTLSConfig()
	}
	return &simpleClient{
		&http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			ForceAttemptHTTP2:     true,
			MaxIdleConnsPerHost:   20,
			IdleConnTimeout:       30 * time.Minute,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			DisableCompression:    true,
			TLSClientConfig:       tlsConfig,
		},
	}
}
