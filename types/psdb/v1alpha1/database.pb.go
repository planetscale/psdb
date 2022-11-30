// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.3
// source: psdb/v1alpha1/database.proto

package psdbv1alpha1

import (
	v1 "github.com/planetscale/psdb/types/psdb/data/v1"
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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string  `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Psid     string  `protobuf:"bytes,2,opt,name=psid,proto3" json:"psid,omitempty"`
	Role     v1.Role `protobuf:"varint,3,opt,name=role,proto3,enum=psdb.data.v1.Role" json:"role,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_psdb_v1alpha1_database_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_psdb_v1alpha1_database_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_psdb_v1alpha1_database_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetPsid() string {
	if x != nil {
		return x.Psid
	}
	return ""
}

func (x *User) GetRole() v1.Role {
	if x != nil {
		return x.Role
	}
	return v1.Role(0)
}

type CreateSessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateSessionRequest) Reset() {
	*x = CreateSessionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_psdb_v1alpha1_database_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSessionRequest) ProtoMessage() {}

func (x *CreateSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_psdb_v1alpha1_database_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSessionRequest.ProtoReflect.Descriptor instead.
func (*CreateSessionRequest) Descriptor() ([]byte, []int) {
	return file_psdb_v1alpha1_database_proto_rawDescGZIP(), []int{1}
}

type CreateSessionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateSessionResponse) Reset() {
	*x = CreateSessionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_psdb_v1alpha1_database_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSessionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSessionResponse) ProtoMessage() {}

func (x *CreateSessionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_psdb_v1alpha1_database_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSessionResponse.ProtoReflect.Descriptor instead.
func (*CreateSessionResponse) Descriptor() ([]byte, []int) {
	return file_psdb_v1alpha1_database_proto_rawDescGZIP(), []int{2}
}

var File_psdb_v1alpha1_database_proto protoreflect.FileDescriptor

var file_psdb_v1alpha1_database_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x73, 0x64, 0x62, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d,
	0x70, 0x73, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x17, 0x70,
	0x73, 0x64, 0x62, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5e, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x73,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x73, 0x69, 0x64, 0x12, 0x26,
	0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x70,
	0x73, 0x64, 0x62, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6c, 0x65,
	0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x16, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x17,
	0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x68, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x12, 0x5c, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x2e, 0x70, 0x73, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x70, 0x73, 0x64, 0x62,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2f, 0x70, 0x73, 0x64, 0x62,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x70, 0x73, 0x64, 0x62, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x3b, 0x70, 0x73, 0x64, 0x62, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_psdb_v1alpha1_database_proto_rawDescOnce sync.Once
	file_psdb_v1alpha1_database_proto_rawDescData = file_psdb_v1alpha1_database_proto_rawDesc
)

func file_psdb_v1alpha1_database_proto_rawDescGZIP() []byte {
	file_psdb_v1alpha1_database_proto_rawDescOnce.Do(func() {
		file_psdb_v1alpha1_database_proto_rawDescData = protoimpl.X.CompressGZIP(file_psdb_v1alpha1_database_proto_rawDescData)
	})
	return file_psdb_v1alpha1_database_proto_rawDescData
}

var file_psdb_v1alpha1_database_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_psdb_v1alpha1_database_proto_goTypes = []interface{}{
	(*User)(nil),                  // 0: psdb.v1alpha1.User
	(*CreateSessionRequest)(nil),  // 1: psdb.v1alpha1.CreateSessionRequest
	(*CreateSessionResponse)(nil), // 2: psdb.v1alpha1.CreateSessionResponse
	(v1.Role)(0),                  // 3: psdb.data.v1.Role
}
var file_psdb_v1alpha1_database_proto_depIdxs = []int32{
	3, // 0: psdb.v1alpha1.User.role:type_name -> psdb.data.v1.Role
	1, // 1: psdb.v1alpha1.Database.CreateSession:input_type -> psdb.v1alpha1.CreateSessionRequest
	2, // 2: psdb.v1alpha1.Database.CreateSession:output_type -> psdb.v1alpha1.CreateSessionResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_psdb_v1alpha1_database_proto_init() }
func file_psdb_v1alpha1_database_proto_init() {
	if File_psdb_v1alpha1_database_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_psdb_v1alpha1_database_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_psdb_v1alpha1_database_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSessionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_psdb_v1alpha1_database_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSessionResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_psdb_v1alpha1_database_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_psdb_v1alpha1_database_proto_goTypes,
		DependencyIndexes: file_psdb_v1alpha1_database_proto_depIdxs,
		MessageInfos:      file_psdb_v1alpha1_database_proto_msgTypes,
	}.Build()
	File_psdb_v1alpha1_database_proto = out.File
	file_psdb_v1alpha1_database_proto_rawDesc = nil
	file_psdb_v1alpha1_database_proto_goTypes = nil
	file_psdb_v1alpha1_database_proto_depIdxs = nil
}
