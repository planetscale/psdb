package psdb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigFromURL(t *testing.T) {
	cases := []struct {
		in  string
		cfg Config
	}{
		{"mysql://foo:bar@example.com/", Config{"example.com", "foo", "bar"}},
		{"mysql://foo@example.com/", Config{"example.com", "foo", ""}},
		{"mysql://foo:bar@example.com:9999", Config{"example.com:9999", "foo", "bar"}},
		{"", Config{}},
	}

	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			cfg := ConfigFromURL(c.in)
			assert.Equal(t, cfg, c.cfg)
		})
	}
}
