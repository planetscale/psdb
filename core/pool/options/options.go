package options

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"os"
	"time"

	"google.golang.org/grpc"
)

type ClientOption interface {
	Apply(*ClientOptions)
}

type ClientOptions struct {
	ConnPoolSize        int
	AddrRefreshInterval time.Duration
	AddrPrefix          string
	KeepaliveTime       time.Duration
	KeepaliveTimeout    time.Duration
	UseCompression      bool
	UseRoundRobin       bool
	ExtraDialOptions    []grpc.DialOption
	ExtraCallOptions    []grpc.CallOption
	TLSConfig           *tls.Config
	ResolverAddress     string
}

func WithConnectionPool(size int) ClientOption {
	return withConnectionPool(size)
}

type withConnectionPool int

func (w withConnectionPool) Apply(o *ClientOptions) {
	o.ConnPoolSize = int(w)
}

func WithAddressRefreshInterval(t time.Duration) ClientOption {
	return withAddressRefreshInterval(t)
}

type withAddressRefreshInterval time.Duration

func (w withAddressRefreshInterval) Apply(o *ClientOptions) {
	o.AddrRefreshInterval = time.Duration(w)
}

func WithAddressPrefix(s string) ClientOption {
	return withAddressPrefix(s)
}

type withAddressPrefix string

func (w withAddressPrefix) Apply(o *ClientOptions) {
	o.AddrPrefix = string(w)
}

func WithKeepaliveTime(t time.Duration) ClientOption {
	return withKeepaliveTime(t)
}

type withKeepaliveTime time.Duration

func (w withKeepaliveTime) Apply(o *ClientOptions) {
	o.KeepaliveTime = time.Duration(w)
}

func WithKeepaliveTimeout(t time.Duration) ClientOption {
	return withKeepaliveTimeout(t)
}

type withKeepaliveTimeout time.Duration

func (w withKeepaliveTimeout) Apply(o *ClientOptions) {
	o.KeepaliveTimeout = time.Duration(w)
}

func WithCompression(b bool) ClientOption {
	return withCompression(b)
}

type withCompression bool

func (w withCompression) Apply(o *ClientOptions) {
	o.UseCompression = bool(w)
}

func WithRoundRobin(b bool) ClientOption {
	return withRoundRobin(b)
}

type withRoundRobin bool

func (w withRoundRobin) Apply(o *ClientOptions) {
	o.UseRoundRobin = bool(w)
}

func WithExtraDialOptions(do []grpc.DialOption) ClientOption {
	return withExtraDialOptions(do)
}

type withExtraDialOptions []grpc.DialOption

func (w withExtraDialOptions) Apply(o *ClientOptions) {
	o.ExtraDialOptions = []grpc.DialOption(w)
}

func WithExtraCallOption(co grpc.CallOption) ClientOption {
	return withExtraCallOption{co}
}

type withExtraCallOption struct {
	grpc.CallOption
}

func (w withExtraCallOption) Apply(o *ClientOptions) {
	o.ExtraCallOptions = append(o.ExtraCallOptions, w.CallOption)
}

func WithTLSConfig(cfg *tls.Config) ClientOption {
	return (*withTLSConfig)(cfg)
}

type withTLSConfig tls.Config

func (w *withTLSConfig) Apply(o *ClientOptions) {
	o.TLSConfig = (*tls.Config)(w)
}

func WithDefaultTLSConfig() ClientOption {
	return WithTLSConfig(DefaultTLSConfig())
}

func WithResolverAddress(addr string) withResolverAddress {
	return withResolverAddress(addr)
}

type withResolverAddress string

func (w withResolverAddress) Apply(o *ClientOptions) {
	o.ResolverAddress = string(w)
}

func DefaultClientOptions() *ClientOptions {
	return &ClientOptions{
		ConnPoolSize:        1,
		AddrRefreshInterval: 0,
		KeepaliveTime:       10 * time.Second,
		KeepaliveTimeout:    10 * time.Second,
		UseCompression:      false,
		UseRoundRobin:       true,
	}
}

func DefaultTLSConfig() *tls.Config {
	certPool, _ := x509.SystemCertPool()
	return TLSConfigWithCertPool(certPool)
}

func TLSConfigWithCertPool(roots *x509.CertPool) *tls.Config {
	return &tls.Config{
		MinVersion: tls.VersionTLS12,
		RootCAs:    roots,
	}
}

func TLSConfigWithRoot(cert string) (*tls.Config, error) {
	b, err := os.ReadFile(cert)
	if err != nil {
		return nil, err
	}
	pool := x509.NewCertPool()
	if !pool.AppendCertsFromPEM(b) {
		return nil, errors.New("no certificates found")
	}
	return TLSConfigWithCertPool(pool), nil
}
