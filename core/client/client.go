package client

import (
	"crypto/tls"
	"net"
	"net/http"
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
	cfg := &config{}
	for _, o := range opts {
		o(cfg)
	}

	cOpts := []connect.ClientOption{
		clientCompressionOpt(defaultCompressionName, defaultCompressionLevel),
		connect.WithSendCompression(defaultCompressionName),
		connect.WithCodec(codec.DefaultCodec),
		connect.WithReadMaxBytes(maxMessageSize),
		connect.WithInterceptors(
			&setHeadersInterceptor{
				"Authorization",
				auth.Type().String() + " " + auth.HeaderValue(),
			},
		),
	}

	if len(cfg.extraClientOptions) > 0 {
		cOpts = append(cOpts, cfg.extraClientOptions...)
	}

	return fn(newClient(cfg.tlsConfig), "https://"+addr, cOpts...)
}

func clientCompressionOpt(name string, level compress.Level) connect.ClientOption {
	opt, _ := compress.Select(name, level)
	return opt
}

func newClient(tlsConfig *tls.Config) connect.HTTPClient {
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
