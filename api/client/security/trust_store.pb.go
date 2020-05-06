// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.2
// source: security/trust_store.proto

package security

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Configuration for a custom trust store.
type TrustStore struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether this custom trust store is enabled.
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// The path to a key store in JKS format containing certification authorities
	// that should be trusted.
	TrustStore string `protobuf:"bytes,2,opt,name=trustStore,proto3" json:"trustStore,omitempty"`
	// The password for the supplied trustStore.
	TrustStorePassword string `protobuf:"bytes,3,opt,name=trustStorePassword,proto3" json:"trustStorePassword,omitempty"`
}

func (x *TrustStore) Reset() {
	*x = TrustStore{}
	if protoimpl.UnsafeEnabled {
		mi := &file_security_trust_store_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrustStore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrustStore) ProtoMessage() {}

func (x *TrustStore) ProtoReflect() protoreflect.Message {
	mi := &file_security_trust_store_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrustStore.ProtoReflect.Descriptor instead.
func (*TrustStore) Descriptor() ([]byte, []int) {
	return file_security_trust_store_proto_rawDescGZIP(), []int{0}
}

func (x *TrustStore) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *TrustStore) GetTrustStore() string {
	if x != nil {
		return x.TrustStore
	}
	return ""
}

func (x *TrustStore) GetTrustStorePassword() string {
	if x != nil {
		return x.TrustStorePassword
	}
	return ""
}

// Configuration for webhooks.
type WebhookConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A custom trust store to use for outgoing webhook connections.
	Trust *TrustStore `protobuf:"bytes,1,opt,name=trust,proto3" json:"trust,omitempty"`
}

func (x *WebhookConfig) Reset() {
	*x = WebhookConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_security_trust_store_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebhookConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebhookConfig) ProtoMessage() {}

func (x *WebhookConfig) ProtoReflect() protoreflect.Message {
	mi := &file_security_trust_store_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebhookConfig.ProtoReflect.Descriptor instead.
func (*WebhookConfig) Descriptor() ([]byte, []int) {
	return file_security_trust_store_proto_rawDescGZIP(), []int{1}
}

func (x *WebhookConfig) GetTrust() *TrustStore {
	if x != nil {
		return x.Trust
	}
	return nil
}

var File_security_trust_store_proto protoreflect.FileDescriptor

var file_security_trust_store_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2f, 0x74, 0x72, 0x75, 0x73, 0x74,
	0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x22, 0x76, 0x0a, 0x0a,
	0x54, 0x72, 0x75, 0x73, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x72, 0x75, 0x73, 0x74, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x72, 0x75, 0x73, 0x74, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x74, 0x72, 0x75, 0x73, 0x74, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x12, 0x74, 0x72, 0x75, 0x73, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x22, 0x41, 0x0a, 0x0d, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x30, 0x0a, 0x05, 0x74, 0x72, 0x75, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x54, 0x72, 0x75, 0x73, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x52, 0x05, 0x74, 0x72, 0x75, 0x73, 0x74, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72, 0x2f,
	0x6b, 0x6c, 0x65, 0x61, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_security_trust_store_proto_rawDescOnce sync.Once
	file_security_trust_store_proto_rawDescData = file_security_trust_store_proto_rawDesc
)

func file_security_trust_store_proto_rawDescGZIP() []byte {
	file_security_trust_store_proto_rawDescOnce.Do(func() {
		file_security_trust_store_proto_rawDescData = protoimpl.X.CompressGZIP(file_security_trust_store_proto_rawDescData)
	})
	return file_security_trust_store_proto_rawDescData
}

var file_security_trust_store_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_security_trust_store_proto_goTypes = []interface{}{
	(*TrustStore)(nil),    // 0: proto.security.TrustStore
	(*WebhookConfig)(nil), // 1: proto.security.WebhookConfig
}
var file_security_trust_store_proto_depIdxs = []int32{
	0, // 0: proto.security.WebhookConfig.trust:type_name -> proto.security.TrustStore
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_security_trust_store_proto_init() }
func file_security_trust_store_proto_init() {
	if File_security_trust_store_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_security_trust_store_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrustStore); i {
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
		file_security_trust_store_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebhookConfig); i {
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
			RawDescriptor: file_security_trust_store_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_security_trust_store_proto_goTypes,
		DependencyIndexes: file_security_trust_store_proto_depIdxs,
		MessageInfos:      file_security_trust_store_proto_msgTypes,
	}.Build()
	File_security_trust_store_proto = out.File
	file_security_trust_store_proto_rawDesc = nil
	file_security_trust_store_proto_goTypes = nil
	file_security_trust_store_proto_depIdxs = nil
}
