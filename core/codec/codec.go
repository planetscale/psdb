// This codec is meant to iterop between both normal protobuf messages
// as well as vtprotbuf optimized messages.
package codec

import (
	"fmt"

	"google.golang.org/protobuf/proto"
)

const Name = "proto"

type codec struct{}

var DefaultCodec = &codec{}

// used by vtprotobuf
type vtprotoMessage interface {
	MarshalVT() ([]byte, error)
	UnmarshalVT([]byte) error
}

func (*codec) Marshal(v any) ([]byte, error) {
	switch vv := v.(type) {
	case vtprotoMessage:
		return vv.MarshalVT()
	case proto.Message:
		return proto.Marshal(vv)
	}
	return nil, fmt.Errorf("failed to marshal, message is %T, want proto.Message", v)
}

func (*codec) Unmarshal(data []byte, v any) error {
	switch vv := v.(type) {
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
