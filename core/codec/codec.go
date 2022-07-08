// This codec is meant to iterop between 3 paths and replace
// the built in default "proto" codec in grpc-go.
// This codec handles messages that are formatted as "raw", which is
// used by the transparent proxy, vtprotobuf, which is our protos
// that are optimized by the vtprotobuf tool, and lastly falling
// back to normal protobuf to interop with things like etcd.
// This codec can and should be used in all cases within our codebase.
package codec

import (
	"fmt"

	// use the original golang/protobuf package so we can continue serializing
	// messages from our dependencies, particularly from etcd
	//lint:ignore SA1019 use the original golang/protobuf package
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc/encoding"
	_ "google.golang.org/grpc/encoding/proto"
)

const Name = "proto"

type codec struct{}

var DefaultCodec = &codec{}

// used by the transparent proxy
type rawMessage interface {
	MarshalRaw() ([]byte, error)
	UnmarshalRaw([]byte) error
}

// used by vtprotobuf
type vtprotoMessage interface {
	MarshalVT() ([]byte, error)
	UnmarshalVT([]byte) error
}

func (*codec) Marshal(v any) ([]byte, error) {
	switch vv := v.(type) {
	case rawMessage:
		return vv.MarshalRaw()
	case vtprotoMessage:
		return vv.MarshalVT()
	case proto.Message:
		return proto.Marshal(vv)
	}
	return nil, fmt.Errorf("failed to marshal, message is %T, want proto.Message", v)
}

func (*codec) Unmarshal(data []byte, v any) error {
	switch vv := v.(type) {
	case rawMessage:
		return vv.UnmarshalRaw(data)
	case vtprotoMessage:
		return vv.UnmarshalVT(data)
	case proto.Message:
		return proto.Unmarshal(data, vv)
	}
	return fmt.Errorf("failed to unmarshal, message is %T, want proto.Message", v)
}

func (*codec) Name() string {
	return Name
}

func (*codec) String() string {
	return Name
}

func init() {
	encoding.RegisterCodec(DefaultCodec)
}
