// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: ras/protocol/v1/connect.proto

package protocolv1

import (
	_ "github.com/v8platform/protoc-gen-go-ras/plugin/ras/encoding"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
	_ "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type NegotiateMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Magic    int32 `protobuf:"varint,1,opt,name=magic,proto3" json:"magic,omitempty"` // Use only one value `475223888`
	Protocol int32 `protobuf:"varint,2,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Version  int32 `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *NegotiateMessage) Reset() {
	*x = NegotiateMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ras_protocol_v1_connect_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NegotiateMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NegotiateMessage) ProtoMessage() {}

func (x *NegotiateMessage) ProtoReflect() protoreflect.Message {
	mi := &file_ras_protocol_v1_connect_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NegotiateMessage.ProtoReflect.Descriptor instead.
func (*NegotiateMessage) Descriptor() ([]byte, []int) {
	return file_ras_protocol_v1_connect_proto_rawDescGZIP(), []int{0}
}

func (x *NegotiateMessage) GetMagic() int32 {
	if x != nil {
		return x.Magic
	}
	return 0
}

func (x *NegotiateMessage) GetProtocol() int32 {
	if x != nil {
		return x.Protocol
	}
	return 0
}

func (x *NegotiateMessage) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

type ConnectMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Известные параметры
	// "connect.timeout" = 2000
	Params map[string]*Param `protobuf:"bytes,1,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ConnectMessage) Reset() {
	*x = ConnectMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ras_protocol_v1_connect_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectMessage) ProtoMessage() {}

func (x *ConnectMessage) ProtoReflect() protoreflect.Message {
	mi := &file_ras_protocol_v1_connect_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectMessage.ProtoReflect.Descriptor instead.
func (*ConnectMessage) Descriptor() ([]byte, []int) {
	return file_ras_protocol_v1_connect_proto_rawDescGZIP(), []int{1}
}

func (x *ConnectMessage) GetParams() map[string]*Param {
	if x != nil {
		return x.Params
	}
	return nil
}

type DisconnectMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Params map[string]*Param `protobuf:"bytes,1,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *DisconnectMessage) Reset() {
	*x = DisconnectMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ras_protocol_v1_connect_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisconnectMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisconnectMessage) ProtoMessage() {}

func (x *DisconnectMessage) ProtoReflect() protoreflect.Message {
	mi := &file_ras_protocol_v1_connect_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisconnectMessage.ProtoReflect.Descriptor instead.
func (*DisconnectMessage) Descriptor() ([]byte, []int) {
	return file_ras_protocol_v1_connect_proto_rawDescGZIP(), []int{2}
}

func (x *DisconnectMessage) GetParams() map[string]*Param {
	if x != nil {
		return x.Params
	}
	return nil
}

type ConnectMessageAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConnectMessageAck) Reset() {
	*x = ConnectMessageAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ras_protocol_v1_connect_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectMessageAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectMessageAck) ProtoMessage() {}

func (x *ConnectMessageAck) ProtoReflect() protoreflect.Message {
	mi := &file_ras_protocol_v1_connect_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectMessageAck.ProtoReflect.Descriptor instead.
func (*ConnectMessageAck) Descriptor() ([]byte, []int) {
	return file_ras_protocol_v1_connect_proto_rawDescGZIP(), []int{3}
}

var File_ras_protocol_v1_connect_proto protoreflect.FileDescriptor

var file_ras_protocol_v1_connect_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x72, 0x61, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0f, 0x72, 0x61, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x31,
	0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x72,
	0x61, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x72, 0x61, 0x73, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x72, 0x61, 0x73, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x72, 0x61, 0x73, 0x2f, 0x65, 0x6e, 0x63, 0x6f, 0x64,
	0x69, 0x6e, 0x67, 0x2f, 0x72, 0x61, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb2, 0x01,
	0x0a, 0x10, 0x4e, 0x65, 0x67, 0x6f, 0x74, 0x69, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x25, 0x0a, 0x05, 0x6d, 0x61, 0x67, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x42, 0x0f, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x09, 0x0a, 0x05, 0x69, 0x6e, 0x74, 0x33, 0x32,
	0x10, 0x01, 0x52, 0x05, 0x6d, 0x61, 0x67, 0x69, 0x63, 0x12, 0x2b, 0x0a, 0x08, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0f, 0x82, 0xf5, 0xea,
	0x94, 0x0e, 0x09, 0x0a, 0x05, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x10, 0x02, 0x52, 0x08, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x29, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0f, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x09, 0x0a,
	0x05, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x10, 0x03, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x3a, 0x1f, 0x8a, 0xf5, 0xea, 0x94, 0x0e, 0x19, 0x2a, 0x15, 0x50, 0x41, 0x43, 0x4b, 0x45,
	0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x45, 0x47, 0x4f, 0x54, 0x49, 0x41, 0x54, 0x45,
	0x58, 0x01, 0x22, 0xcf, 0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x4d, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x72, 0x61, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x42, 0x08, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x02, 0x10, 0x01, 0x52, 0x06, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x1a, 0x51, 0x0a, 0x0b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x72, 0x61, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x3a, 0x1b, 0x8a, 0xf5, 0xea, 0x94, 0x0e, 0x15, 0x2a,
	0x13, 0x50, 0x41, 0x43, 0x4b, 0x45, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x4e,
	0x4e, 0x45, 0x43, 0x54, 0x22, 0xd8, 0x01, 0x0a, 0x11, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x50, 0x0a, 0x06, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x72, 0x61, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x73,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x08, 0x82, 0xf5, 0xea, 0x94,
	0x0e, 0x02, 0x10, 0x01, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x51, 0x0a, 0x0b,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2c, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x72,
	0x61, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x3a,
	0x1e, 0x8a, 0xf5, 0xea, 0x94, 0x0e, 0x18, 0x2a, 0x16, 0x50, 0x41, 0x43, 0x4b, 0x45, 0x54, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x49, 0x53, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x22,
	0x36, 0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x41, 0x63, 0x6b, 0x3a, 0x21, 0x8a, 0xf5, 0xea, 0x94, 0x0e, 0x1b, 0x08, 0x01, 0x2a, 0x17,
	0x50, 0x41, 0x43, 0x4b, 0x45, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x4e,
	0x45, 0x43, 0x54, 0x5f, 0x41, 0x43, 0x4b, 0x42, 0xbe, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e,
	0x72, 0x61, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x42,
	0x0c, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x38, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x72, 0x61, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76,
	0x31, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x52,
	0x50, 0x58, 0xaa, 0x02, 0x0f, 0x52, 0x61, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0f, 0x52, 0x61, 0x73, 0x5c, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1b, 0x52, 0x61, 0x73, 0x5c, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x52, 0x61, 0x73, 0x3a, 0x3a, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ras_protocol_v1_connect_proto_rawDescOnce sync.Once
	file_ras_protocol_v1_connect_proto_rawDescData = file_ras_protocol_v1_connect_proto_rawDesc
)

func file_ras_protocol_v1_connect_proto_rawDescGZIP() []byte {
	file_ras_protocol_v1_connect_proto_rawDescOnce.Do(func() {
		file_ras_protocol_v1_connect_proto_rawDescData = protoimpl.X.CompressGZIP(file_ras_protocol_v1_connect_proto_rawDescData)
	})
	return file_ras_protocol_v1_connect_proto_rawDescData
}

var file_ras_protocol_v1_connect_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_ras_protocol_v1_connect_proto_goTypes = []interface{}{
	(*NegotiateMessage)(nil),  // 0: ras.protocol.v1.NegotiateMessage
	(*ConnectMessage)(nil),    // 1: ras.protocol.v1.ConnectMessage
	(*DisconnectMessage)(nil), // 2: ras.protocol.v1.DisconnectMessage
	(*ConnectMessageAck)(nil), // 3: ras.protocol.v1.ConnectMessageAck
	nil,                       // 4: ras.protocol.v1.ConnectMessage.ParamsEntry
	nil,                       // 5: ras.protocol.v1.DisconnectMessage.ParamsEntry
	(*Param)(nil),             // 6: ras.protocol.v1.Param
}
var file_ras_protocol_v1_connect_proto_depIdxs = []int32{
	4, // 0: ras.protocol.v1.ConnectMessage.params:type_name -> ras.protocol.v1.ConnectMessage.ParamsEntry
	5, // 1: ras.protocol.v1.DisconnectMessage.params:type_name -> ras.protocol.v1.DisconnectMessage.ParamsEntry
	6, // 2: ras.protocol.v1.ConnectMessage.ParamsEntry.value:type_name -> ras.protocol.v1.Param
	6, // 3: ras.protocol.v1.DisconnectMessage.ParamsEntry.value:type_name -> ras.protocol.v1.Param
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_ras_protocol_v1_connect_proto_init() }
func file_ras_protocol_v1_connect_proto_init() {
	if File_ras_protocol_v1_connect_proto != nil {
		return
	}
	file_ras_protocol_v1_param_proto_init()
	file_ras_protocol_v1_types_proto_init()
	file_ras_protocol_v1_packet_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ras_protocol_v1_connect_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NegotiateMessage); i {
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
		file_ras_protocol_v1_connect_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectMessage); i {
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
		file_ras_protocol_v1_connect_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisconnectMessage); i {
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
		file_ras_protocol_v1_connect_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectMessageAck); i {
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
			RawDescriptor: file_ras_protocol_v1_connect_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ras_protocol_v1_connect_proto_goTypes,
		DependencyIndexes: file_ras_protocol_v1_connect_proto_depIdxs,
		MessageInfos:      file_ras_protocol_v1_connect_proto_msgTypes,
	}.Build()
	File_ras_protocol_v1_connect_proto = out.File
	file_ras_protocol_v1_connect_proto_rawDesc = nil
	file_ras_protocol_v1_connect_proto_goTypes = nil
	file_ras_protocol_v1_connect_proto_depIdxs = nil
}
