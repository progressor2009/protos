// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: v8platform/serialize/v1/licanses.proto

package serializev1

import (
	_ "github.com/v8platform/protoc-gen-go-ras/plugin/ras/encoding"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LicenseInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ignore
	// need fill
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	ProcessId string `protobuf:"bytes,2,opt,name=process_id,json=processId,proto3" json:"process_id,omitempty"`
	SessionId string `protobuf:"bytes,3,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	UserName  string `protobuf:"bytes,4,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	Host      string `protobuf:"bytes,6,opt,name=host,proto3" json:"host,omitempty"`
	AppId     string `protobuf:"bytes,7,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	// all version
	FullName          string `protobuf:"bytes,8,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	FullPresentation  string `protobuf:"bytes,9,opt,name=full_presentation,json=fullPresentation,proto3" json:"full_presentation,omitempty"`
	IssuedByServer    bool   `protobuf:"varint,10,opt,name=issued_by_server,json=issuedByServer,proto3" json:"issued_by_server,omitempty"`
	LicenseType       int32  `protobuf:"varint,11,opt,name=license_type,json=licenseType,proto3" json:"license_type,omitempty"`
	MaxUsersAll       int32  `protobuf:"varint,12,opt,name=max_users_all,json=maxUsersAll,proto3" json:"max_users_all,omitempty"`
	MaxUsersCur       int32  `protobuf:"varint,13,opt,name=max_users_cur,json=maxUsersCur,proto3" json:"max_users_cur,omitempty"`
	Net               bool   `protobuf:"varint,14,opt,name=net,proto3" json:"net,omitempty"`
	RmngrAddress      string `protobuf:"bytes,15,opt,name=rmngr_address,json=rmngrAddress,proto3" json:"rmngr_address,omitempty"`
	RmngrPid          string `protobuf:"bytes,16,opt,name=rmngr_pid,json=rmngrPid,proto3" json:"rmngr_pid,omitempty"`
	RmngrPort         int32  `protobuf:"varint,17,opt,name=rmngr_port,json=rmngrPort,proto3" json:"rmngr_port,omitempty"`
	Series            string `protobuf:"bytes,18,opt,name=series,proto3" json:"series,omitempty"`
	ShortPresentation string `protobuf:"bytes,19,opt,name=short_presentation,json=shortPresentation,proto3" json:"short_presentation,omitempty"`
}

func (x *LicenseInfo) Reset() {
	*x = LicenseInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v8platform_serialize_v1_licanses_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LicenseInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LicenseInfo) ProtoMessage() {}

func (x *LicenseInfo) ProtoReflect() protoreflect.Message {
	mi := &file_v8platform_serialize_v1_licanses_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LicenseInfo.ProtoReflect.Descriptor instead.
func (*LicenseInfo) Descriptor() ([]byte, []int) {
	return file_v8platform_serialize_v1_licanses_proto_rawDescGZIP(), []int{0}
}

func (x *LicenseInfo) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *LicenseInfo) GetProcessId() string {
	if x != nil {
		return x.ProcessId
	}
	return ""
}

func (x *LicenseInfo) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *LicenseInfo) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *LicenseInfo) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *LicenseInfo) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *LicenseInfo) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *LicenseInfo) GetFullPresentation() string {
	if x != nil {
		return x.FullPresentation
	}
	return ""
}

func (x *LicenseInfo) GetIssuedByServer() bool {
	if x != nil {
		return x.IssuedByServer
	}
	return false
}

func (x *LicenseInfo) GetLicenseType() int32 {
	if x != nil {
		return x.LicenseType
	}
	return 0
}

func (x *LicenseInfo) GetMaxUsersAll() int32 {
	if x != nil {
		return x.MaxUsersAll
	}
	return 0
}

func (x *LicenseInfo) GetMaxUsersCur() int32 {
	if x != nil {
		return x.MaxUsersCur
	}
	return 0
}

func (x *LicenseInfo) GetNet() bool {
	if x != nil {
		return x.Net
	}
	return false
}

func (x *LicenseInfo) GetRmngrAddress() string {
	if x != nil {
		return x.RmngrAddress
	}
	return ""
}

func (x *LicenseInfo) GetRmngrPid() string {
	if x != nil {
		return x.RmngrPid
	}
	return ""
}

func (x *LicenseInfo) GetRmngrPort() int32 {
	if x != nil {
		return x.RmngrPort
	}
	return 0
}

func (x *LicenseInfo) GetSeries() string {
	if x != nil {
		return x.Series
	}
	return ""
}

func (x *LicenseInfo) GetShortPresentation() string {
	if x != nil {
		return x.ShortPresentation
	}
	return ""
}

var File_v8platform_serialize_v1_licanses_proto protoreflect.FileDescriptor

var file_v8platform_serialize_v1_licanses_proto_rawDesc = []byte{
	0x0a, 0x26, 0x76, 0x38, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x73, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x63, 0x61, 0x6e, 0x73,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x76, 0x38, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x2e, 0x76,
	0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x16, 0x72, 0x61, 0x73, 0x2f, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67,
	0x2f, 0x72, 0x61, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc3, 0x05, 0x0a, 0x0b, 0x4c,
	0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f,
	0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12,
	0x25, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x08, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x02, 0x10, 0x01, 0x52, 0x08, 0x66, 0x75,
	0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x11, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x70,
	0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x08, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x02, 0x10, 0x02, 0x52, 0x10, 0x66, 0x75, 0x6c,
	0x6c, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x0a,
	0x10, 0x69, 0x73, 0x73, 0x75, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x42, 0x08, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x02, 0x10,
	0x03, 0x52, 0x0e, 0x69, 0x73, 0x73, 0x75, 0x65, 0x64, 0x42, 0x79, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x12, 0x2b, 0x0a, 0x0c, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x42, 0x08, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x02, 0x10,
	0x04, 0x52, 0x0b, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2c,
	0x0a, 0x0d, 0x6d, 0x61, 0x78, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x61, 0x6c, 0x6c, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x05, 0x42, 0x08, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x02, 0x10, 0x05, 0x52,
	0x0b, 0x6d, 0x61, 0x78, 0x55, 0x73, 0x65, 0x72, 0x73, 0x41, 0x6c, 0x6c, 0x12, 0x2c, 0x0a, 0x0d,
	0x6d, 0x61, 0x78, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x63, 0x75, 0x72, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x05, 0x42, 0x08, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x02, 0x10, 0x06, 0x52, 0x0b, 0x6d,
	0x61, 0x78, 0x55, 0x73, 0x65, 0x72, 0x73, 0x43, 0x75, 0x72, 0x12, 0x1a, 0x0a, 0x03, 0x6e, 0x65,
	0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x42, 0x08, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x02, 0x10,
	0x07, 0x52, 0x03, 0x6e, 0x65, 0x74, 0x12, 0x2d, 0x0a, 0x0d, 0x72, 0x6d, 0x6e, 0x67, 0x72, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0x82,
	0xf5, 0xea, 0x94, 0x0e, 0x02, 0x10, 0x08, 0x52, 0x0c, 0x72, 0x6d, 0x6e, 0x67, 0x72, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x25, 0x0a, 0x09, 0x72, 0x6d, 0x6e, 0x67, 0x72, 0x5f, 0x70,
	0x69, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x02,
	0x10, 0x09, 0x52, 0x08, 0x72, 0x6d, 0x6e, 0x67, 0x72, 0x50, 0x69, 0x64, 0x12, 0x27, 0x0a, 0x0a,
	0x72, 0x6d, 0x6e, 0x67, 0x72, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28, 0x05,
	0x42, 0x08, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x02, 0x10, 0x0a, 0x52, 0x09, 0x72, 0x6d, 0x6e, 0x67,
	0x72, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x20, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x18,
	0x12, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x02, 0x10, 0x0b, 0x52,
	0x06, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x12, 0x37, 0x0a, 0x12, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x5f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x13, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x08, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x02, 0x10, 0x0c, 0x52, 0x11, 0x73,
	0x68, 0x6f, 0x72, 0x74, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0xf0, 0x01, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x38, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x2e, 0x76, 0x31,
	0x42, 0x0d, 0x4c, 0x69, 0x63, 0x61, 0x6e, 0x73, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x38,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x76, 0x38, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x73,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x69, 0x7a, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x56, 0x53, 0x58, 0xaa, 0x02, 0x17,
	0x56, 0x38, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x53, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x69, 0x7a, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x17, 0x56, 0x38, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x5c, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x5c, 0x56,
	0x31, 0xe2, 0x02, 0x23, 0x56, 0x38, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5c, 0x53,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x19, 0x56, 0x38, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v8platform_serialize_v1_licanses_proto_rawDescOnce sync.Once
	file_v8platform_serialize_v1_licanses_proto_rawDescData = file_v8platform_serialize_v1_licanses_proto_rawDesc
)

func file_v8platform_serialize_v1_licanses_proto_rawDescGZIP() []byte {
	file_v8platform_serialize_v1_licanses_proto_rawDescOnce.Do(func() {
		file_v8platform_serialize_v1_licanses_proto_rawDescData = protoimpl.X.CompressGZIP(file_v8platform_serialize_v1_licanses_proto_rawDescData)
	})
	return file_v8platform_serialize_v1_licanses_proto_rawDescData
}

var file_v8platform_serialize_v1_licanses_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_v8platform_serialize_v1_licanses_proto_goTypes = []interface{}{
	(*LicenseInfo)(nil), // 0: v8platform.serialize.v1.LicenseInfo
}
var file_v8platform_serialize_v1_licanses_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v8platform_serialize_v1_licanses_proto_init() }
func file_v8platform_serialize_v1_licanses_proto_init() {
	if File_v8platform_serialize_v1_licanses_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v8platform_serialize_v1_licanses_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LicenseInfo); i {
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
			RawDescriptor: file_v8platform_serialize_v1_licanses_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v8platform_serialize_v1_licanses_proto_goTypes,
		DependencyIndexes: file_v8platform_serialize_v1_licanses_proto_depIdxs,
		MessageInfos:      file_v8platform_serialize_v1_licanses_proto_msgTypes,
	}.Build()
	File_v8platform_serialize_v1_licanses_proto = out.File
	file_v8platform_serialize_v1_licanses_proto_rawDesc = nil
	file_v8platform_serialize_v1_licanses_proto_goTypes = nil
	file_v8platform_serialize_v1_licanses_proto_depIdxs = nil
}
