// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: vitess/v14/vitess.proto

package vitessv14

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

// Type defines the various supported data types in bind vars
// and query results.
type Type int32

const (
	// NULL_TYPE specifies a NULL type.
	Type_NULL_TYPE Type = 0
	// INT8 specifies a TINYINT type.
	// Properties: 1, IsNumber.
	Type_INT8 Type = 257
	// UINT8 specifies a TINYINT UNSIGNED type.
	// Properties: 2, IsNumber, IsUnsigned.
	Type_UINT8 Type = 770
	// INT16 specifies a SMALLINT type.
	// Properties: 3, IsNumber.
	Type_INT16 Type = 259
	// UINT16 specifies a SMALLINT UNSIGNED type.
	// Properties: 4, IsNumber, IsUnsigned.
	Type_UINT16 Type = 772
	// INT24 specifies a MEDIUMINT type.
	// Properties: 5, IsNumber.
	Type_INT24 Type = 261
	// UINT24 specifies a MEDIUMINT UNSIGNED type.
	// Properties: 6, IsNumber, IsUnsigned.
	Type_UINT24 Type = 774
	// INT32 specifies a INTEGER type.
	// Properties: 7, IsNumber.
	Type_INT32 Type = 263
	// UINT32 specifies a INTEGER UNSIGNED type.
	// Properties: 8, IsNumber, IsUnsigned.
	Type_UINT32 Type = 776
	// INT64 specifies a BIGINT type.
	// Properties: 9, IsNumber.
	Type_INT64 Type = 265
	// UINT64 specifies a BIGINT UNSIGNED type.
	// Properties: 10, IsNumber, IsUnsigned.
	Type_UINT64 Type = 778
	// FLOAT32 specifies a FLOAT type.
	// Properties: 11, IsFloat.
	Type_FLOAT32 Type = 1035
	// FLOAT64 specifies a DOUBLE or REAL type.
	// Properties: 12, IsFloat.
	Type_FLOAT64 Type = 1036
	// TIMESTAMP specifies a TIMESTAMP type.
	// Properties: 13, IsQuoted.
	Type_TIMESTAMP Type = 2061
	// DATE specifies a DATE type.
	// Properties: 14, IsQuoted.
	Type_DATE Type = 2062
	// TIME specifies a TIME type.
	// Properties: 15, IsQuoted.
	Type_TIME Type = 2063
	// DATETIME specifies a DATETIME type.
	// Properties: 16, IsQuoted.
	Type_DATETIME Type = 2064
	// YEAR specifies a YEAR type.
	// Properties: 17, IsNumber, IsUnsigned.
	Type_YEAR Type = 785
	// DECIMAL specifies a DECIMAL or NUMERIC type.
	// Properties: 18, None.
	Type_DECIMAL Type = 18
	// TEXT specifies a TEXT type.
	// Properties: 19, IsQuoted, IsText.
	Type_TEXT Type = 6163
	// BLOB specifies a BLOB type.
	// Properties: 20, IsQuoted, IsBinary.
	Type_BLOB Type = 10260
	// VARCHAR specifies a VARCHAR type.
	// Properties: 21, IsQuoted, IsText.
	Type_VARCHAR Type = 6165
	// VARBINARY specifies a VARBINARY type.
	// Properties: 22, IsQuoted, IsBinary.
	Type_VARBINARY Type = 10262
	// CHAR specifies a CHAR type.
	// Properties: 23, IsQuoted, IsText.
	Type_CHAR Type = 6167
	// BINARY specifies a BINARY type.
	// Properties: 24, IsQuoted, IsBinary.
	Type_BINARY Type = 10264
	// BIT specifies a BIT type.
	// Properties: 25, IsQuoted.
	Type_BIT Type = 2073
	// ENUM specifies an ENUM type.
	// Properties: 26, IsQuoted.
	Type_ENUM Type = 2074
	// SET specifies a SET type.
	// Properties: 27, IsQuoted.
	Type_SET Type = 2075
	// TUPLE specifies a tuple. This cannot
	// be returned in a QueryResult, but it can
	// be sent as a bind var.
	// Properties: 28, None.
	Type_TUPLE Type = 28
	// GEOMETRY specifies a GEOMETRY type.
	// Properties: 29, IsQuoted.
	Type_GEOMETRY Type = 2077
	// JSON specifies a JSON type.
	// Properties: 30, IsQuoted.
	Type_JSON Type = 2078
	// EXPRESSION specifies a SQL expression.
	// This type is for internal use only.
	// Properties: 31, None.
	Type_EXPRESSION Type = 31
	// HEXNUM specifies a HEXNUM type (unquoted varbinary).
	// Properties: 32, IsText.
	Type_HEXNUM Type = 4128
	// HEXVAL specifies a HEXVAL type (unquoted varbinary).
	// Properties: 33, IsText.
	Type_HEXVAL Type = 4129
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0:     "NULL_TYPE",
		257:   "INT8",
		770:   "UINT8",
		259:   "INT16",
		772:   "UINT16",
		261:   "INT24",
		774:   "UINT24",
		263:   "INT32",
		776:   "UINT32",
		265:   "INT64",
		778:   "UINT64",
		1035:  "FLOAT32",
		1036:  "FLOAT64",
		2061:  "TIMESTAMP",
		2062:  "DATE",
		2063:  "TIME",
		2064:  "DATETIME",
		785:   "YEAR",
		18:    "DECIMAL",
		6163:  "TEXT",
		10260: "BLOB",
		6165:  "VARCHAR",
		10262: "VARBINARY",
		6167:  "CHAR",
		10264: "BINARY",
		2073:  "BIT",
		2074:  "ENUM",
		2075:  "SET",
		28:    "TUPLE",
		2077:  "GEOMETRY",
		2078:  "JSON",
		31:    "EXPRESSION",
		4128:  "HEXNUM",
		4129:  "HEXVAL",
	}
	Type_value = map[string]int32{
		"NULL_TYPE":  0,
		"INT8":       257,
		"UINT8":      770,
		"INT16":      259,
		"UINT16":     772,
		"INT24":      261,
		"UINT24":     774,
		"INT32":      263,
		"UINT32":     776,
		"INT64":      265,
		"UINT64":     778,
		"FLOAT32":    1035,
		"FLOAT64":    1036,
		"TIMESTAMP":  2061,
		"DATE":       2062,
		"TIME":       2063,
		"DATETIME":   2064,
		"YEAR":       785,
		"DECIMAL":    18,
		"TEXT":       6163,
		"BLOB":       10260,
		"VARCHAR":    6165,
		"VARBINARY":  10262,
		"CHAR":       6167,
		"BINARY":     10264,
		"BIT":        2073,
		"ENUM":       2074,
		"SET":        2075,
		"TUPLE":      28,
		"GEOMETRY":   2077,
		"JSON":       2078,
		"EXPRESSION": 31,
		"HEXNUM":     4128,
		"HEXVAL":     4129,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_vitess_v14_vitess_proto_enumTypes[0].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_vitess_v14_vitess_proto_enumTypes[0]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_vitess_v14_vitess_proto_rawDescGZIP(), []int{0}
}

// Field describes a single column returned by a query
type Field struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name of the field as returned by mysql C API
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// vitess-defined type. Conversion function is in sqltypes package.
	Type Type `protobuf:"varint,2,opt,name=type,proto3,enum=vitess.v14.Type" json:"type,omitempty"`
	// Remaining fields from mysql C API.
	// These fields are only populated when ExecuteOptions.included_fields
	// is set to IncludedFields.ALL.
	Table    string `protobuf:"bytes,3,opt,name=table,proto3" json:"table,omitempty"`
	OrgTable string `protobuf:"bytes,4,opt,name=org_table,json=orgTable,proto3" json:"org_table,omitempty"`
	Database string `protobuf:"bytes,5,opt,name=database,proto3" json:"database,omitempty"`
	OrgName  string `protobuf:"bytes,6,opt,name=org_name,json=orgName,proto3" json:"org_name,omitempty"`
	// column_length is really a uint32. All 32 bits can be used.
	ColumnLength uint32 `protobuf:"varint,7,opt,name=column_length,json=columnLength,proto3" json:"column_length,omitempty"`
	// charset is actually a uint16. Only the lower 16 bits are used.
	Charset uint32 `protobuf:"varint,8,opt,name=charset,proto3" json:"charset,omitempty"`
	// decimals is actually a uint8. Only the lower 8 bits are used.
	Decimals uint32 `protobuf:"varint,9,opt,name=decimals,proto3" json:"decimals,omitempty"`
	// flags is actually a uint16. Only the lower 16 bits are used.
	Flags uint32 `protobuf:"varint,10,opt,name=flags,proto3" json:"flags,omitempty"`
	// column_type is optionally populated from information_schema.columns
	ColumnType string `protobuf:"bytes,11,opt,name=column_type,json=columnType,proto3" json:"column_type,omitempty"`
}

func (x *Field) Reset() {
	*x = Field{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vitess_v14_vitess_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Field) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Field) ProtoMessage() {}

func (x *Field) ProtoReflect() protoreflect.Message {
	mi := &file_vitess_v14_vitess_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Field.ProtoReflect.Descriptor instead.
func (*Field) Descriptor() ([]byte, []int) {
	return file_vitess_v14_vitess_proto_rawDescGZIP(), []int{0}
}

func (x *Field) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Field) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_NULL_TYPE
}

func (x *Field) GetTable() string {
	if x != nil {
		return x.Table
	}
	return ""
}

func (x *Field) GetOrgTable() string {
	if x != nil {
		return x.OrgTable
	}
	return ""
}

func (x *Field) GetDatabase() string {
	if x != nil {
		return x.Database
	}
	return ""
}

func (x *Field) GetOrgName() string {
	if x != nil {
		return x.OrgName
	}
	return ""
}

func (x *Field) GetColumnLength() uint32 {
	if x != nil {
		return x.ColumnLength
	}
	return 0
}

func (x *Field) GetCharset() uint32 {
	if x != nil {
		return x.Charset
	}
	return 0
}

func (x *Field) GetDecimals() uint32 {
	if x != nil {
		return x.Decimals
	}
	return 0
}

func (x *Field) GetFlags() uint32 {
	if x != nil {
		return x.Flags
	}
	return 0
}

func (x *Field) GetColumnType() string {
	if x != nil {
		return x.ColumnType
	}
	return ""
}

// Row is a database row.
type Row struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// lengths contains the length of each value in values.
	// A length of -1 means that the field is NULL. While
	// reading values, you have to accummulate the length
	// to know the offset where the next value begins in values.
	Lengths []int64 `protobuf:"zigzag64,1,rep,packed,name=lengths,proto3" json:"lengths,omitempty"`
	// values contains a concatenation of all values in the row.
	Values []byte `protobuf:"bytes,2,opt,name=values,proto3" json:"values,omitempty"`
}

func (x *Row) Reset() {
	*x = Row{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vitess_v14_vitess_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Row) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Row) ProtoMessage() {}

func (x *Row) ProtoReflect() protoreflect.Message {
	mi := &file_vitess_v14_vitess_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Row.ProtoReflect.Descriptor instead.
func (*Row) Descriptor() ([]byte, []int) {
	return file_vitess_v14_vitess_proto_rawDescGZIP(), []int{1}
}

func (x *Row) GetLengths() []int64 {
	if x != nil {
		return x.Lengths
	}
	return nil
}

func (x *Row) GetValues() []byte {
	if x != nil {
		return x.Values
	}
	return nil
}

// QueryResult is returned by Execute and ExecuteStream.
//
// As returned by Execute, len(fields) is always equal to len(row)
// (for each row in rows).
//
// As returned by StreamExecute, the first QueryResult has the fields
// set, and subsequent QueryResult have rows set. And as Execute,
// len(QueryResult[0].fields) is always equal to len(row) (for each
// row in rows for each QueryResult in QueryResult[1:]).
type QueryResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fields       []*Field `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields,omitempty"`
	RowsAffected uint64   `protobuf:"varint,2,opt,name=rows_affected,json=rowsAffected,proto3" json:"rows_affected,omitempty"`
	InsertId     uint64   `protobuf:"varint,3,opt,name=insert_id,json=insertId,proto3" json:"insert_id,omitempty"`
	Rows         []*Row   `protobuf:"bytes,4,rep,name=rows,proto3" json:"rows,omitempty"`
	Info         string   `protobuf:"bytes,6,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *QueryResult) Reset() {
	*x = QueryResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vitess_v14_vitess_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryResult) ProtoMessage() {}

func (x *QueryResult) ProtoReflect() protoreflect.Message {
	mi := &file_vitess_v14_vitess_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryResult.ProtoReflect.Descriptor instead.
func (*QueryResult) Descriptor() ([]byte, []int) {
	return file_vitess_v14_vitess_proto_rawDescGZIP(), []int{2}
}

func (x *QueryResult) GetFields() []*Field {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *QueryResult) GetRowsAffected() uint64 {
	if x != nil {
		return x.RowsAffected
	}
	return 0
}

func (x *QueryResult) GetInsertId() uint64 {
	if x != nil {
		return x.InsertId
	}
	return 0
}

func (x *QueryResult) GetRows() []*Row {
	if x != nil {
		return x.Rows
	}
	return nil
}

func (x *QueryResult) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

var File_vitess_v14_vitess_proto protoreflect.FileDescriptor

var file_vitess_v14_vitess_proto_rawDesc = []byte{
	0x0a, 0x17, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x34, 0x2f, 0x76, 0x69, 0x74,
	0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x76, 0x69, 0x74, 0x65, 0x73,
	0x73, 0x2e, 0x76, 0x31, 0x34, 0x22, 0xbd, 0x02, 0x0a, 0x05, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x10, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x34, 0x2e, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x6f, 0x72, 0x67, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6f, 0x72, 0x67, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x67, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x67, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x5f, 0x6c, 0x65,
	0x6e, 0x67, 0x74, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x63, 0x6f, 0x6c, 0x75,
	0x6d, 0x6e, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x72,
	0x73, 0x65, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x63, 0x68, 0x61, 0x72, 0x73,
	0x65, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x73, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x66,
	0x6c, 0x61, 0x67, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6c, 0x75, 0x6d,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0x37, 0x0a, 0x03, 0x52, 0x6f, 0x77, 0x12, 0x18, 0x0a, 0x07,
	0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x12, 0x52, 0x07, 0x6c,
	0x65, 0x6e, 0x67, 0x74, 0x68, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0xb9,
	0x01, 0x0a, 0x0b, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x29,
	0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x34, 0x2e, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x6f, 0x77,
	0x73, 0x5f, 0x61, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0c, 0x72, 0x6f, 0x77, 0x73, 0x41, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x69, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x08, 0x69, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x04, 0x72,
	0x6f, 0x77, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x76, 0x69, 0x74, 0x65,
	0x73, 0x73, 0x2e, 0x76, 0x31, 0x34, 0x2e, 0x52, 0x6f, 0x77, 0x52, 0x04, 0x72, 0x6f, 0x77, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x69, 0x6e, 0x66, 0x6f, 0x4a, 0x04, 0x08, 0x05, 0x10, 0x06, 0x2a, 0xb3, 0x03, 0x0a, 0x04, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x4e, 0x55, 0x4c, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x10, 0x00, 0x12, 0x09, 0x0a, 0x04, 0x49, 0x4e, 0x54, 0x38, 0x10, 0x81, 0x02, 0x12, 0x0a, 0x0a,
	0x05, 0x55, 0x49, 0x4e, 0x54, 0x38, 0x10, 0x82, 0x06, 0x12, 0x0a, 0x0a, 0x05, 0x49, 0x4e, 0x54,
	0x31, 0x36, 0x10, 0x83, 0x02, 0x12, 0x0b, 0x0a, 0x06, 0x55, 0x49, 0x4e, 0x54, 0x31, 0x36, 0x10,
	0x84, 0x06, 0x12, 0x0a, 0x0a, 0x05, 0x49, 0x4e, 0x54, 0x32, 0x34, 0x10, 0x85, 0x02, 0x12, 0x0b,
	0x0a, 0x06, 0x55, 0x49, 0x4e, 0x54, 0x32, 0x34, 0x10, 0x86, 0x06, 0x12, 0x0a, 0x0a, 0x05, 0x49,
	0x4e, 0x54, 0x33, 0x32, 0x10, 0x87, 0x02, 0x12, 0x0b, 0x0a, 0x06, 0x55, 0x49, 0x4e, 0x54, 0x33,
	0x32, 0x10, 0x88, 0x06, 0x12, 0x0a, 0x0a, 0x05, 0x49, 0x4e, 0x54, 0x36, 0x34, 0x10, 0x89, 0x02,
	0x12, 0x0b, 0x0a, 0x06, 0x55, 0x49, 0x4e, 0x54, 0x36, 0x34, 0x10, 0x8a, 0x06, 0x12, 0x0c, 0x0a,
	0x07, 0x46, 0x4c, 0x4f, 0x41, 0x54, 0x33, 0x32, 0x10, 0x8b, 0x08, 0x12, 0x0c, 0x0a, 0x07, 0x46,
	0x4c, 0x4f, 0x41, 0x54, 0x36, 0x34, 0x10, 0x8c, 0x08, 0x12, 0x0e, 0x0a, 0x09, 0x54, 0x49, 0x4d,
	0x45, 0x53, 0x54, 0x41, 0x4d, 0x50, 0x10, 0x8d, 0x10, 0x12, 0x09, 0x0a, 0x04, 0x44, 0x41, 0x54,
	0x45, 0x10, 0x8e, 0x10, 0x12, 0x09, 0x0a, 0x04, 0x54, 0x49, 0x4d, 0x45, 0x10, 0x8f, 0x10, 0x12,
	0x0d, 0x0a, 0x08, 0x44, 0x41, 0x54, 0x45, 0x54, 0x49, 0x4d, 0x45, 0x10, 0x90, 0x10, 0x12, 0x09,
	0x0a, 0x04, 0x59, 0x45, 0x41, 0x52, 0x10, 0x91, 0x06, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45, 0x43,
	0x49, 0x4d, 0x41, 0x4c, 0x10, 0x12, 0x12, 0x09, 0x0a, 0x04, 0x54, 0x45, 0x58, 0x54, 0x10, 0x93,
	0x30, 0x12, 0x09, 0x0a, 0x04, 0x42, 0x4c, 0x4f, 0x42, 0x10, 0x94, 0x50, 0x12, 0x0c, 0x0a, 0x07,
	0x56, 0x41, 0x52, 0x43, 0x48, 0x41, 0x52, 0x10, 0x95, 0x30, 0x12, 0x0e, 0x0a, 0x09, 0x56, 0x41,
	0x52, 0x42, 0x49, 0x4e, 0x41, 0x52, 0x59, 0x10, 0x96, 0x50, 0x12, 0x09, 0x0a, 0x04, 0x43, 0x48,
	0x41, 0x52, 0x10, 0x97, 0x30, 0x12, 0x0b, 0x0a, 0x06, 0x42, 0x49, 0x4e, 0x41, 0x52, 0x59, 0x10,
	0x98, 0x50, 0x12, 0x08, 0x0a, 0x03, 0x42, 0x49, 0x54, 0x10, 0x99, 0x10, 0x12, 0x09, 0x0a, 0x04,
	0x45, 0x4e, 0x55, 0x4d, 0x10, 0x9a, 0x10, 0x12, 0x08, 0x0a, 0x03, 0x53, 0x45, 0x54, 0x10, 0x9b,
	0x10, 0x12, 0x09, 0x0a, 0x05, 0x54, 0x55, 0x50, 0x4c, 0x45, 0x10, 0x1c, 0x12, 0x0d, 0x0a, 0x08,
	0x47, 0x45, 0x4f, 0x4d, 0x45, 0x54, 0x52, 0x59, 0x10, 0x9d, 0x10, 0x12, 0x09, 0x0a, 0x04, 0x4a,
	0x53, 0x4f, 0x4e, 0x10, 0x9e, 0x10, 0x12, 0x0e, 0x0a, 0x0a, 0x45, 0x58, 0x50, 0x52, 0x45, 0x53,
	0x53, 0x49, 0x4f, 0x4e, 0x10, 0x1f, 0x12, 0x0b, 0x0a, 0x06, 0x48, 0x45, 0x58, 0x4e, 0x55, 0x4d,
	0x10, 0xa0, 0x20, 0x12, 0x0b, 0x0a, 0x06, 0x48, 0x45, 0x58, 0x56, 0x41, 0x4c, 0x10, 0xa1, 0x20,
	0x42, 0x9e, 0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2e,
	0x76, 0x31, 0x34, 0x42, 0x0b, 0x56, 0x69, 0x74, 0x65, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70,
	0x6c, 0x61, 0x6e, 0x65, 0x74, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2f, 0x70, 0x73, 0x64, 0x62, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x34,
	0x3b, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x76, 0x31, 0x34, 0xa2, 0x02, 0x03, 0x56, 0x58, 0x58,
	0xaa, 0x02, 0x0a, 0x56, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2e, 0x56, 0x31, 0x34, 0xca, 0x02, 0x0a,
	0x56, 0x69, 0x74, 0x65, 0x73, 0x73, 0x5c, 0x56, 0x31, 0x34, 0xe2, 0x02, 0x16, 0x56, 0x69, 0x74,
	0x65, 0x73, 0x73, 0x5c, 0x56, 0x31, 0x34, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x0b, 0x56, 0x69, 0x74, 0x65, 0x73, 0x73, 0x3a, 0x3a, 0x56, 0x31,
	0x34, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vitess_v14_vitess_proto_rawDescOnce sync.Once
	file_vitess_v14_vitess_proto_rawDescData = file_vitess_v14_vitess_proto_rawDesc
)

func file_vitess_v14_vitess_proto_rawDescGZIP() []byte {
	file_vitess_v14_vitess_proto_rawDescOnce.Do(func() {
		file_vitess_v14_vitess_proto_rawDescData = protoimpl.X.CompressGZIP(file_vitess_v14_vitess_proto_rawDescData)
	})
	return file_vitess_v14_vitess_proto_rawDescData
}

var file_vitess_v14_vitess_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_vitess_v14_vitess_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_vitess_v14_vitess_proto_goTypes = []interface{}{
	(Type)(0),           // 0: vitess.v14.Type
	(*Field)(nil),       // 1: vitess.v14.Field
	(*Row)(nil),         // 2: vitess.v14.Row
	(*QueryResult)(nil), // 3: vitess.v14.QueryResult
}
var file_vitess_v14_vitess_proto_depIdxs = []int32{
	0, // 0: vitess.v14.Field.type:type_name -> vitess.v14.Type
	1, // 1: vitess.v14.QueryResult.fields:type_name -> vitess.v14.Field
	2, // 2: vitess.v14.QueryResult.rows:type_name -> vitess.v14.Row
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_vitess_v14_vitess_proto_init() }
func file_vitess_v14_vitess_proto_init() {
	if File_vitess_v14_vitess_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vitess_v14_vitess_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Field); i {
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
		file_vitess_v14_vitess_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Row); i {
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
		file_vitess_v14_vitess_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryResult); i {
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
			RawDescriptor: file_vitess_v14_vitess_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_vitess_v14_vitess_proto_goTypes,
		DependencyIndexes: file_vitess_v14_vitess_proto_depIdxs,
		EnumInfos:         file_vitess_v14_vitess_proto_enumTypes,
		MessageInfos:      file_vitess_v14_vitess_proto_msgTypes,
	}.Build()
	File_vitess_v14_vitess_proto = out.File
	file_vitess_v14_vitess_proto_rawDesc = nil
	file_vitess_v14_vitess_proto_goTypes = nil
	file_vitess_v14_vitess_proto_depIdxs = nil
}
