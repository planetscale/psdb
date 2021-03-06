package auth

import (
	"bytes"
	"context"
	"errors"
	"strings"

	"github.com/segmentio/asm/base64"
	"google.golang.org/grpc"
)

const (
	sha1Size   = 20
	sha256Size = 32
)

var (
	ErrMalformedAuthorization = errors.New("malformed Authorization")
	ErrInvalidMySQLAlgorithm  = errors.New("unknown MySQL password hashing algorithm")
)

//enumcheck:relaxed
type AuthType string

const (
	BasicAuthType       = AuthType("Basic")
	MysqlSha1AuthType   = AuthType("mysql-sha1")
	MysqlSha256AuthType = AuthType("mysql-sha256")
)

func (t AuthType) String() string {
	return string(t)
}

type Authorization struct {
	authType    AuthType
	username    string
	headerValue string
	secretBytes []byte
}

func (a *Authorization) Type() AuthType       { return a.authType }
func (a *Authorization) Username() string     { return a.username }
func (a *Authorization) HeaderValue() string  { return a.headerValue }
func (a *Authorization) HasSecretBytes() bool { return a.secretBytes != nil }
func (a *Authorization) SecretBytes() []byte  { return a.secretBytes }
func (a *Authorization) PasswordLength() int {
	switch a.authType {
	case BasicAuthType:
		return len(a.secretBytes)
	case MysqlSha1AuthType:
		return sha1Size
	case MysqlSha256AuthType:
		return sha256Size
	}
	panic("unknown AuthType")
}

func (a *Authorization) CallOption() grpc.CallOption {
	return grpc.PerRPCCredentials(a)
}

func (a *Authorization) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": a.authType.String() + " " + a.headerValue,
	}, nil
}

// Needs to always be false becuase grpc doesn't actually know we're using TLS
// since we wrap the socket ourselves. Maybe rethink this to use TransportCredentials?
func (a *Authorization) RequireTransportSecurity() bool { return false }

func NewBasicAuth(username, password string) *Authorization {
	return &Authorization{
		authType:    BasicAuthType,
		username:    username,
		headerValue: b64encode([]byte(username + ":" + password)),
	}
}

func NewMySQLAuth(username string, salt, pwHash []byte) (*Authorization, error) {
	var authType AuthType
	switch len(pwHash) {
	case sha1Size:
		authType = MysqlSha1AuthType
	case sha256Size:
		authType = MysqlSha256AuthType
	default:
		return nil, ErrInvalidMySQLAlgorithm
	}

	value := make([]byte, len(username)+1+len(pwHash)+len(salt))
	n := 0
	n += copy(value[n:], username)
	value[n] = ':'
	n++
	n += copy(value[n:], pwHash)
	copy(value[n:], salt)

	return &Authorization{
		authType:    authType,
		username:    username,
		headerValue: b64encode(value),
	}, nil
}

func Parse(value string) (*Authorization, error) {
	return parse(value, false)
}

func ParseWithSecret(value string) (*Authorization, error) {
	return parse(value, true)
}

func parse(value string, keepSecret bool) (*Authorization, error) {
	spacePos := strings.IndexByte(value, ' ')
	if spacePos < 0 {
		return nil, ErrMalformedAuthorization
	}

	authTypeS := value[:spacePos]
	authType, err := makeAuthType(authTypeS)
	if err != nil {
		return nil, err
	}

	value = value[spacePos+1:]

	username, secret, err := splitUsernameSecret(value)
	if err != nil {
		return nil, err
	}

	if keepSecret && len(secret) == 0 {
		return nil, ErrMalformedAuthorization
	}

	if !keepSecret {
		secret = nil
	}

	return &Authorization{
		authType:    AuthType(authType),
		username:    string(username),
		headerValue: value,
		secretBytes: secret,
	}, nil
}

func splitUsernameSecret(value string) ([]byte, []byte, error) {
	b, err := b64decode(value)
	if err != nil {
		return nil, nil, ErrMalformedAuthorization
	}
	colonPos := bytes.IndexByte(b, ':')
	if colonPos < 1 {
		return nil, nil, ErrMalformedAuthorization
	}
	return b[:colonPos], b[colonPos+1:], nil
}

func b64encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func b64decode(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}

func makeAuthType(s string) (AuthType, error) {
	switch s {
	case BasicAuthType.String():
		return BasicAuthType, nil
	case MysqlSha1AuthType.String():
		return MysqlSha1AuthType, nil
	case MysqlSha256AuthType.String():
		return MysqlSha256AuthType, nil
	}
	return AuthType(""), ErrMalformedAuthorization
}
