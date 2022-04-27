package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	mysqlSalt       = []byte{1, 1, 1, 1}
	mysqlSha1PwHash = []byte{
		// 20 bytes long
		0, 0, 0, 0, 0,
		0, 0, 0, 0, 0,
		0, 0, 0, 0, 0,
		0, 0, 0, 0, 0,
	}
	mysqlSha256PwHash = []byte{
		// 32 bytes long
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
	}
)

func TestNewBasicAuth(t *testing.T) {
	cases := []struct {
		user, pw string
		auth     *Authorization
	}{
		{"xyzxyzxyz123", "bar", &Authorization{
			authType:    BasicAuthType,
			username:    "xyzxyzxyz123",
			headerValue: "eHl6eHl6eHl6MTIzOmJhcg==",
			secretBytes: nil,
		}},
	}
	for _, c := range cases {
		auth := NewBasicAuth(c.user, c.pw)
		assert.Equal(t, auth, c.auth)
	}
}

func TestNewMySQLAuth(t *testing.T) {
	cases := []struct {
		user         string
		salt, pwHash []byte
		auth         *Authorization
		err          error
	}{
		{"xxxxxxxxxxroute-hint", mysqlSalt, mysqlSha1PwHash, &Authorization{
			authType:    MysqlSha1AuthType,
			username:    "xxxxxxxxxxroute-hint",
			headerValue: "eHh4eHh4eHh4eHJvdXRlLWhpbnQ6AAAAAAAAAAAAAAAAAAAAAAAAAAABAQEB",
			secretBytes: nil,
		}, nil},
		{"xxxxxxxxxxroute-hint", mysqlSalt, mysqlSha256PwHash, &Authorization{
			authType:    MysqlSha256AuthType,
			username:    "xxxxxxxxxxroute-hint",
			headerValue: "eHh4eHh4eHh4eHJvdXRlLWhpbnQ6AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABAQEB",
			secretBytes: nil,
		}, nil},
		{"xxxxxxxxxxroute-hint", mysqlSalt, mysqlSha256PwHash[:5], nil, ErrInvalidMySQLAlgorithm},
		{"xxxxxxxxxxroute-hint", []byte{}, []byte{}, nil, ErrInvalidMySQLAlgorithm},
	}
	for _, c := range cases {
		auth, err := NewMySQLAuth(c.user, c.salt, c.pwHash)
		if c.err == nil {
			assert.Nil(t, err)
			assert.Equal(t, auth, c.auth)
		} else {
			assert.Equal(t, err, c.err)
		}
	}
}

func TestParse(t *testing.T) {
	cases := []struct {
		in   string
		auth *Authorization
		err  error
	}{
		{"Basic eHh4eHh4eHh4eHJvdXRlLWhpbnQ6YmFy", &Authorization{
			authType:    BasicAuthType,
			username:    "xxxxxxxxxxroute-hint",
			headerValue: "eHh4eHh4eHh4eHJvdXRlLWhpbnQ6YmFy",
			secretBytes: nil,
		}, nil},
		{"mysql-sha1 eHh4eHh4eHh4eHJvdXRlLWhpbnQ6AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA", &Authorization{
			authType:    MysqlSha1AuthType,
			username:    "xxxxxxxxxxroute-hint",
			headerValue: "eHh4eHh4eHh4eHJvdXRlLWhpbnQ6AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
			secretBytes: nil,
		}, nil},
		{"mysql-sha256 eHh4eHh4eHh4eHJvdXRlLWhpbnQ6AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA", &Authorization{
			authType:    MysqlSha256AuthType,
			username:    "xxxxxxxxxxroute-hint",
			headerValue: "eHh4eHh4eHh4eHJvdXRlLWhpbnQ6AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
			secretBytes: nil,
		}, nil},
		{"mysql-sha256 x", nil, ErrMalformedAuthorization},
		{"mysql-sha256 ", nil, ErrMalformedAuthorization},
		{"xxx eHh4eHh4eHh4eHJvdXRlLWhpbnQ6YmFy", nil, ErrMalformedAuthorization},
		{"x", nil, ErrMalformedAuthorization},
		{"", nil, ErrMalformedAuthorization},
	}
	for _, c := range cases {
		auth, err := Parse(c.in)
		if c.err == nil {
			assert.Nil(t, err)
			assert.Equal(t, auth, c.auth)
		} else {
			assert.Equal(t, err, c.err)
		}
	}
}

func TestParseWithSecret(t *testing.T) {
	cases := []struct {
		in   string
		auth *Authorization
		err  error
	}{
		{"Basic eHh4eHh4eHh4eHJvdXRlLWhpbnQ6YmFy", &Authorization{
			authType:    BasicAuthType,
			username:    "xxxxxxxxxxroute-hint",
			headerValue: "eHh4eHh4eHh4eHJvdXRlLWhpbnQ6YmFy",
			secretBytes: []byte("bar"),
		}, nil},
		{"mysql-sha1 eHh4eHh4eHh4eHJvdXRlLWhpbnQ6AAAAAAAAAAAAAAAAAAAAAAAAAAABAQEB", &Authorization{
			authType:    MysqlSha1AuthType,
			username:    "xxxxxxxxxxroute-hint",
			headerValue: "eHh4eHh4eHh4eHJvdXRlLWhpbnQ6AAAAAAAAAAAAAAAAAAAAAAAAAAABAQEB",
			secretBytes: joinBytes(mysqlSha1PwHash, mysqlSalt),
		}, nil},
	}
	for _, c := range cases {
		auth, err := ParseWithSecret(c.in)
		if c.err == nil {
			assert.Nil(t, err)
			assert.Equal(t, auth, c.auth)
		} else {
			assert.Equal(t, err, c.err)
		}
	}
}

func joinBytes(a []byte, rest ...[]byte) []byte {
	ll := len(a)
	for _, b := range rest {
		ll += len(b)
	}
	out := make([]byte, ll)
	n := copy(out[:len(a)], a)
	for _, b := range rest {
		n += copy(out[n:n+len(b)], b)
	}
	return out
}
