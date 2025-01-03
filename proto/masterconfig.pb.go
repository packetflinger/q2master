// compile with:
// protoc --go_out=. --go_opt=paths=source_relative masterconfig.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.2
// source: masterconfig.proto

package proto

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

type MasterConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr       string `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	Port       uint32 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	Foreground bool   `protobuf:"varint,3,opt,name=foreground,proto3" json:"foreground,omitempty"`
	Logfile    string `protobuf:"bytes,4,opt,name=logfile,proto3" json:"logfile,omitempty"`
	Api        bool   `protobuf:"varint,5,opt,name=api,proto3" json:"api,omitempty"`
	ApiAddr    string `protobuf:"bytes,6,opt,name=api_addr,json=apiAddr,proto3" json:"api_addr,omitempty"`
	ApiPort    uint32 `protobuf:"varint,7,opt,name=api_port,json=apiPort,proto3" json:"api_port,omitempty"`
}

func (x *MasterConfig) Reset() {
	*x = MasterConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_masterconfig_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MasterConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MasterConfig) ProtoMessage() {}

func (x *MasterConfig) ProtoReflect() protoreflect.Message {
	mi := &file_masterconfig_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MasterConfig.ProtoReflect.Descriptor instead.
func (*MasterConfig) Descriptor() ([]byte, []int) {
	return file_masterconfig_proto_rawDescGZIP(), []int{0}
}

func (x *MasterConfig) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *MasterConfig) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *MasterConfig) GetForeground() bool {
	if x != nil {
		return x.Foreground
	}
	return false
}

func (x *MasterConfig) GetLogfile() string {
	if x != nil {
		return x.Logfile
	}
	return ""
}

func (x *MasterConfig) GetApi() bool {
	if x != nil {
		return x.Api
	}
	return false
}

func (x *MasterConfig) GetApiAddr() string {
	if x != nil {
		return x.ApiAddr
	}
	return ""
}

func (x *MasterConfig) GetApiPort() uint32 {
	if x != nil {
		return x.ApiPort
	}
	return 0
}

var File_masterconfig_proto protoreflect.FileDescriptor

var file_masterconfig_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb8, 0x01, 0x0a, 0x0c,
	0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04,
	0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04,
	0x70, 0x6f, 0x72, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x6f, 0x72, 0x65, 0x67, 0x72, 0x6f, 0x75,
	0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x66, 0x6f, 0x72, 0x65, 0x67, 0x72,
	0x6f, 0x75, 0x6e, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x6f, 0x67, 0x66, 0x69, 0x6c, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f, 0x67, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x61, 0x70, 0x69, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x61, 0x70, 0x69,
	0x12, 0x19, 0x0a, 0x08, 0x61, 0x70, 0x69, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x70, 0x69, 0x41, 0x64, 0x64, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x61,
	0x70, 0x69, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x61,
	0x70, 0x69, 0x50, 0x6f, 0x72, 0x74, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x66, 0x6c, 0x69, 0x6e, 0x67,
	0x65, 0x72, 0x2f, 0x71, 0x32, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_masterconfig_proto_rawDescOnce sync.Once
	file_masterconfig_proto_rawDescData = file_masterconfig_proto_rawDesc
)

func file_masterconfig_proto_rawDescGZIP() []byte {
	file_masterconfig_proto_rawDescOnce.Do(func() {
		file_masterconfig_proto_rawDescData = protoimpl.X.CompressGZIP(file_masterconfig_proto_rawDescData)
	})
	return file_masterconfig_proto_rawDescData
}

var file_masterconfig_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_masterconfig_proto_goTypes = []interface{}{
	(*MasterConfig)(nil), // 0: proto.MasterConfig
}
var file_masterconfig_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_masterconfig_proto_init() }
func file_masterconfig_proto_init() {
	if File_masterconfig_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_masterconfig_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MasterConfig); i {
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
			RawDescriptor: file_masterconfig_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_masterconfig_proto_goTypes,
		DependencyIndexes: file_masterconfig_proto_depIdxs,
		MessageInfos:      file_masterconfig_proto_msgTypes,
	}.Build()
	File_masterconfig_proto = out.File
	file_masterconfig_proto_rawDesc = nil
	file_masterconfig_proto_goTypes = nil
	file_masterconfig_proto_depIdxs = nil
}
