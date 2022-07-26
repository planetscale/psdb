package client

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"os"
	"sync"

	"github.com/bufbuild/connect-go"
)

type config struct {
	tlsConfig          *tls.Config
	extraClientOptions []connect.ClientOption
}

type Option func(*config)

func WithTLSConfig(cfg *tls.Config) Option {
	return func(c *config) {
		c.tlsConfig = cfg
	}
}

func WithExtraClientOptions(opts ...connect.ClientOption) Option {
	return func(c *config) {
		c.extraClientOptions = opts
	}
}

var (
	defaultTLSConfig     *tls.Config
	defaultTLSConfigOnce sync.Once
)

func DefaultTLSConfig() *tls.Config {
	defaultTLSConfigOnce.Do(initDefaultTLSConfig)
	return defaultTLSConfig
}

func initDefaultTLSConfig() {
	certPool, _ := x509.SystemCertPool()
	defaultTLSConfig = TLSConfigWithCertPool(certPool)
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
