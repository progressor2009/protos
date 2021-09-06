// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: ras/protocol/v1/param.proto

package protocolv1

import (
	_ "github.com/v8platform/protoc-gen-go-ras/plugin/ras/encoding"
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

type ParamType int32

const (
	ParamType_PARAM_UNKNOWN_TYPE  ParamType = 0
	ParamType_PARAM_BOOLEAN       ParamType = 1
	ParamType_PARAM_BYTE          ParamType = 2
	ParamType_PARAM_SHORT         ParamType = 3
	ParamType_PARAM_INT           ParamType = 4
	ParamType_PARAM_LONG          ParamType = 5
	ParamType_PARAM_FLOAT         ParamType = 6
	ParamType_PARAM_DOUBLE        ParamType = 7
	ParamType_PARAM_SIZE          ParamType = 8
	ParamType_PARAM_NULLABLE_SIZE ParamType = 9
	ParamType_PARAM_STRING        ParamType = 10
	ParamType_PARAM_UUID          ParamType = 11
	ParamType_PARAM_TYPE          ParamType = 12
	ParamType_PARAM_ENDPOINT_ID   ParamType = 13
)

// Enum value maps for ParamType.
var (
	ParamType_name = map[int32]string{
		0:  "PARAM_UNKNOWN_TYPE",
		1:  "PARAM_BOOLEAN",
		2:  "PARAM_BYTE",
		3:  "PARAM_SHORT",
		4:  "PARAM_INT",
		5:  "PARAM_LONG",
		6:  "PARAM_FLOAT",
		7:  "PARAM_DOUBLE",
		8:  "PARAM_SIZE",
		9:  "PARAM_NULLABLE_SIZE",
		10: "PARAM_STRING",
		11: "PARAM_UUID",
		12: "PARAM_TYPE",
		13: "PARAM_ENDPOINT_ID",
	}
	ParamType_value = map[string]int32{
		"PARAM_UNKNOWN_TYPE":  0,
		"PARAM_BOOLEAN":       1,
		"PARAM_BYTE":          2,
		"PARAM_SHORT":         3,
		"PARAM_INT":           4,
		"PARAM_LONG":          5,
		"PARAM_FLOAT":         6,
		"PARAM_DOUBLE":        7,
		"PARAM_SIZE":          8,
		"PARAM_NULLABLE_SIZE": 9,
		"PARAM_STRING":        10,
		"PARAM_UUID":          11,
		"PARAM_TYPE":          12,
		"PARAM_ENDPOINT_ID":   13,
	}
)

func (x ParamType) Enum() *ParamType {
	p := new(ParamType)
	*p = x
	return p
}

func (x ParamType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ParamType) Descriptor() protoreflect.EnumDescriptor {
	return file_ras_protocol_v1_param_proto_enumTypes[0].Descriptor()
}

func (ParamType) Type() protoreflect.EnumType {
	return &file_ras_protocol_v1_param_proto_enumTypes[0]
}

func (x ParamType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ParamType.Descriptor instead.
func (ParamType) EnumDescriptor() ([]byte, []int) {
	return file_ras_protocol_v1_param_proto_rawDescGZIP(), []int{0}
}

type Param struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  ParamType `protobuf:"varint,1,opt,name=type,proto3,enum=ras.protocol.v1.ParamType" json:"type,omitempty"`
	Value []byte    `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Param) Reset() {
	*x = Param{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ras_protocol_v1_param_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Param) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Param) ProtoMessage() {}

func (x *Param) ProtoReflect() protoreflect.Message {
	mi := &file_ras_protocol_v1_param_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Param.ProtoReflect.Descriptor instead.
func (*Param) Descriptor() ([]byte, []int) {
	return file_ras_protocol_v1_param_proto_rawDescGZIP(), []int{0}
}

func (x *Param) GetType() ParamType {
	if x != nil {
		return x.Type
	}
	return ParamType_PARAM_UNKNOWN_TYPE
}

func (x *Param) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_ras_protocol_v1_param_proto protoreflect.FileDescriptor

var file_ras_protocol_v1_param_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x72, 0x61, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76,
	0x31, 0x2f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x72,
	0x61, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x1a, 0x16,
	0x72, 0x61, 0x73, 0x2f, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2f, 0x72, 0x61, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6d, 0x0a, 0x05, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12,
	0x3e, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e,
	0x72, 0x61, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x42, 0x0e, 0x82, 0xf5, 0xea, 0x94, 0x0e,
	0x08, 0x0a, 0x04, 0x62, 0x79, 0x74, 0x65, 0x10, 0x01, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x24, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x0e,
	0x82, 0xf5, 0xea, 0x94, 0x0e, 0x08, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x10, 0x02, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x2a, 0x8b, 0x02, 0x0a, 0x09, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x5f, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x50,
	0x41, 0x52, 0x41, 0x4d, 0x5f, 0x42, 0x4f, 0x4f, 0x4c, 0x45, 0x41, 0x4e, 0x10, 0x01, 0x12, 0x0e,
	0x0a, 0x0a, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x5f, 0x42, 0x59, 0x54, 0x45, 0x10, 0x02, 0x12, 0x0f,
	0x0a, 0x0b, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x5f, 0x53, 0x48, 0x4f, 0x52, 0x54, 0x10, 0x03, 0x12,
	0x0d, 0x0a, 0x09, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x5f, 0x49, 0x4e, 0x54, 0x10, 0x04, 0x12, 0x0e,
	0x0a, 0x0a, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x5f, 0x4c, 0x4f, 0x4e, 0x47, 0x10, 0x05, 0x12, 0x0f,
	0x0a, 0x0b, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x5f, 0x46, 0x4c, 0x4f, 0x41, 0x54, 0x10, 0x06, 0x12,
	0x10, 0x0a, 0x0c, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x5f, 0x44, 0x4f, 0x55, 0x42, 0x4c, 0x45, 0x10,
	0x07, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x5f, 0x53, 0x49, 0x5a, 0x45, 0x10,
	0x08, 0x12, 0x17, 0x0a, 0x13, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x5f, 0x4e, 0x55, 0x4c, 0x4c, 0x41,
	0x42, 0x4c, 0x45, 0x5f, 0x53, 0x49, 0x5a, 0x45, 0x10, 0x09, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x41,
	0x52, 0x41, 0x4d, 0x5f, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x0a, 0x12, 0x0e, 0x0a, 0x0a,
	0x50, 0x41, 0x52, 0x41, 0x4d, 0x5f, 0x55, 0x55, 0x49, 0x44, 0x10, 0x0b, 0x12, 0x0e, 0x0a, 0x0a,
	0x50, 0x41, 0x52, 0x41, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x0c, 0x12, 0x15, 0x0a, 0x11,
	0x50, 0x41, 0x52, 0x41, 0x4d, 0x5f, 0x45, 0x4e, 0x44, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x5f, 0x49,
	0x44, 0x10, 0x0d, 0x42, 0xbc, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x72, 0x61, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x38, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x72, 0x61, 0x73, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x52, 0x50, 0x58, 0xaa, 0x02, 0x0f, 0x52,
	0x61, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x56, 0x31, 0xca, 0x02,
	0x0f, 0x52, 0x61, 0x73, 0x5c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5c, 0x56, 0x31,
	0xe2, 0x02, 0x1b, 0x52, 0x61, 0x73, 0x5c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x11, 0x52, 0x61, 0x73, 0x3a, 0x3a, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ras_protocol_v1_param_proto_rawDescOnce sync.Once
	file_ras_protocol_v1_param_proto_rawDescData = file_ras_protocol_v1_param_proto_rawDesc
)

func file_ras_protocol_v1_param_proto_rawDescGZIP() []byte {
	file_ras_protocol_v1_param_proto_rawDescOnce.Do(func() {
		file_ras_protocol_v1_param_proto_rawDescData = protoimpl.X.CompressGZIP(file_ras_protocol_v1_param_proto_rawDescData)
	})
	return file_ras_protocol_v1_param_proto_rawDescData
}

var file_ras_protocol_v1_param_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ras_protocol_v1_param_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ras_protocol_v1_param_proto_goTypes = []interface{}{
	(ParamType)(0), // 0: ras.protocol.v1.ParamType
	(*Param)(nil),  // 1: ras.protocol.v1.Param
}
var file_ras_protocol_v1_param_proto_depIdxs = []int32{
	0, // 0: ras.protocol.v1.Param.type:type_name -> ras.protocol.v1.ParamType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ras_protocol_v1_param_proto_init() }
func file_ras_protocol_v1_param_proto_init() {
	if File_ras_protocol_v1_param_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ras_protocol_v1_param_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Param); i {
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
			RawDescriptor: file_ras_protocol_v1_param_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ras_protocol_v1_param_proto_goTypes,
		DependencyIndexes: file_ras_protocol_v1_param_proto_depIdxs,
		EnumInfos:         file_ras_protocol_v1_param_proto_enumTypes,
		MessageInfos:      file_ras_protocol_v1_param_proto_msgTypes,
	}.Build()
	File_ras_protocol_v1_param_proto = out.File
	file_ras_protocol_v1_param_proto_rawDesc = nil
	file_ras_protocol_v1_param_proto_goTypes = nil
	file_ras_protocol_v1_param_proto_depIdxs = nil
}
