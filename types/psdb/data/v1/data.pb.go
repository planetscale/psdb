// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: psdb/data/v1/data.proto

package psdbdatav1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

//enumcheck:exhaustive
type Role int32

const (
	Role_reader     Role = 0
	Role_writer     Role = 1
	Role_readwriter Role = 2
	Role_admin      Role = 3
)

// Enum value maps for Role.
var (
	Role_name = map[int32]string{
		0: "reader",
		1: "writer",
		2: "readwriter",
		3: "admin",
	}
	Role_value = map[string]int32{
		"reader":     0,
		"writer":     1,
		"readwriter": 2,
		"admin":      3,
	}
)

func (x Role) Enum() *Role {
	p := new(Role)
	*p = x
	return p
}

func (x Role) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Role) Descriptor() protoreflect.EnumDescriptor {
	return file_psdb_data_v1_data_proto_enumTypes[0].Descriptor()
}

func (Role) Type() protoreflect.EnumType {
	return &file_psdb_data_v1_data_proto_enumTypes[0]
}

func (x Role) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Role.Descriptor instead.
func (Role) EnumDescriptor() ([]byte, []int) {
	return file_psdb_data_v1_data_proto_rawDescGZIP(), []int{0}
}

var File_psdb_data_v1_data_proto protoreflect.FileDescriptor

var file_psdb_data_v1_data_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x73, 0x64, 0x62, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70, 0x73, 0x64, 0x62, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2a, 0x39, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12,
	0x0a, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x64, 0x65, 0x72, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x77,
	0x72, 0x69, 0x74, 0x65, 0x72, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x72, 0x65, 0x61, 0x64, 0x77,
	0x72, 0x69, 0x74, 0x65, 0x72, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x10, 0x03, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2f, 0x70, 0x73, 0x64,
	0x62, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x70, 0x73, 0x64, 0x62, 0x2f, 0x64, 0x61, 0x74,
	0x61, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x73, 0x64, 0x62, 0x64, 0x61, 0x74, 0x61, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_psdb_data_v1_data_proto_rawDescOnce sync.Once
	file_psdb_data_v1_data_proto_rawDescData = file_psdb_data_v1_data_proto_rawDesc
)

func file_psdb_data_v1_data_proto_rawDescGZIP() []byte {
	file_psdb_data_v1_data_proto_rawDescOnce.Do(func() {
		file_psdb_data_v1_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_psdb_data_v1_data_proto_rawDescData)
	})
	return file_psdb_data_v1_data_proto_rawDescData
}

var file_psdb_data_v1_data_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_psdb_data_v1_data_proto_goTypes = []interface{}{
	(Role)(0), // 0: psdb.data.v1.Role
}
var file_psdb_data_v1_data_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_psdb_data_v1_data_proto_init() }
func file_psdb_data_v1_data_proto_init() {
	if File_psdb_data_v1_data_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_psdb_data_v1_data_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_psdb_data_v1_data_proto_goTypes,
		DependencyIndexes: file_psdb_data_v1_data_proto_depIdxs,
		EnumInfos:         file_psdb_data_v1_data_proto_enumTypes,
	}.Build()
	File_psdb_data_v1_data_proto = out.File
	file_psdb_data_v1_data_proto_rawDesc = nil
	file_psdb_data_v1_data_proto_goTypes = nil
	file_psdb_data_v1_data_proto_depIdxs = nil
}
