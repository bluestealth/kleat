// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: spinnaker/extensibility.proto

package spinnaker

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
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

// Spinnaker flags
type Spinnaker struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Extensibility *Extensibility `protobuf:"bytes,1,opt,name=extensibility,proto3" json:"extensibility,omitempty"`
}

func (x *Spinnaker) Reset() {
	*x = Spinnaker{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spinnaker_extensibility_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Spinnaker) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Spinnaker) ProtoMessage() {}

func (x *Spinnaker) ProtoReflect() protoreflect.Message {
	mi := &file_spinnaker_extensibility_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Spinnaker.ProtoReflect.Descriptor instead.
func (*Spinnaker) Descriptor() ([]byte, []int) {
	return file_spinnaker_extensibility_proto_rawDescGZIP(), []int{0}
}

func (x *Spinnaker) GetExtensibility() *Extensibility {
	if x != nil {
		return x.Extensibility
	}
	return nil
}

// Extensibility flags
type Extensibility struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Map with the spinnaker plugins.
	Plugins map[string]*Plugin `protobuf:"bytes,1,rep,name=plugins,proto3" json:"plugins,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Map with the plugin repositories.
	Repositories map[string]*Repository `protobuf:"bytes,2,rep,name=repositories,proto3" json:"repositories,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Deck Proxy for gate
	DeckProxy *DeckProxy `protobuf:"bytes,3,opt,name=deckProxy,proto3" json:"deckProxy,omitempty"`
}

func (x *Extensibility) Reset() {
	*x = Extensibility{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spinnaker_extensibility_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Extensibility) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Extensibility) ProtoMessage() {}

func (x *Extensibility) ProtoReflect() protoreflect.Message {
	mi := &file_spinnaker_extensibility_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Extensibility.ProtoReflect.Descriptor instead.
func (*Extensibility) Descriptor() ([]byte, []int) {
	return file_spinnaker_extensibility_proto_rawDescGZIP(), []int{1}
}

func (x *Extensibility) GetPlugins() map[string]*Plugin {
	if x != nil {
		return x.Plugins
	}
	return nil
}

func (x *Extensibility) GetRepositories() map[string]*Repository {
	if x != nil {
		return x.Repositories
	}
	return nil
}

func (x *Extensibility) GetDeckProxy() *DeckProxy {
	if x != nil {
		return x.DeckProxy
	}
	return nil
}

// DeckProxy flags
type DeckProxy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether deckproxy is enabled
	Enabled *wrapperspb.BoolValue `protobuf:"bytes,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// Map with the spinnaker plugins.
	Plugins map[string]*Plugin `protobuf:"bytes,2,rep,name=plugins,proto3" json:"plugins,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *DeckProxy) Reset() {
	*x = DeckProxy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spinnaker_extensibility_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeckProxy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeckProxy) ProtoMessage() {}

func (x *DeckProxy) ProtoReflect() protoreflect.Message {
	mi := &file_spinnaker_extensibility_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeckProxy.ProtoReflect.Descriptor instead.
func (*DeckProxy) Descriptor() ([]byte, []int) {
	return file_spinnaker_extensibility_proto_rawDescGZIP(), []int{2}
}

func (x *DeckProxy) GetEnabled() *wrapperspb.BoolValue {
	if x != nil {
		return x.Enabled
	}
	return nil
}

func (x *DeckProxy) GetPlugins() map[string]*Plugin {
	if x != nil {
		return x.Plugins
	}
	return nil
}

// Spinnaker plugin flags
type Plugin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether plugin is enabled.
	Enabled *wrapperspb.BoolValue `protobuf:"bytes,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// Plugin version to use.
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// Plugin configuration
	Config map[string]string `protobuf:"bytes,3,rep,name=config,proto3" json:"config,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Plugin) Reset() {
	*x = Plugin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spinnaker_extensibility_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Plugin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Plugin) ProtoMessage() {}

func (x *Plugin) ProtoReflect() protoreflect.Message {
	mi := &file_spinnaker_extensibility_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Plugin.ProtoReflect.Descriptor instead.
func (*Plugin) Descriptor() ([]byte, []int) {
	return file_spinnaker_extensibility_proto_rawDescGZIP(), []int{3}
}

func (x *Plugin) GetEnabled() *wrapperspb.BoolValue {
	if x != nil {
		return x.Enabled
	}
	return nil
}

func (x *Plugin) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Plugin) GetConfig() map[string]string {
	if x != nil {
		return x.Config
	}
	return nil
}

// Spinnaker plugin repository flags
type Repository struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// unique repository name.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// URL of plugins.json file.
	Url string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *Repository) Reset() {
	*x = Repository{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spinnaker_extensibility_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Repository) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Repository) ProtoMessage() {}

func (x *Repository) ProtoReflect() protoreflect.Message {
	mi := &file_spinnaker_extensibility_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Repository.ProtoReflect.Descriptor instead.
func (*Repository) Descriptor() ([]byte, []int) {
	return file_spinnaker_extensibility_proto_rawDescGZIP(), []int{4}
}

func (x *Repository) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Repository) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_spinnaker_extensibility_proto protoreflect.FileDescriptor

var file_spinnaker_extensibility_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x73, 0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x51, 0x0a, 0x09, 0x53, 0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72, 0x12, 0x44, 0x0a,
	0x0d, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x70, 0x69,
	0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x52, 0x0d, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x22, 0x99, 0x03, 0x0a, 0x0d, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x45, 0x0a, 0x07, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73,
	0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x07, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x12, 0x54, 0x0a, 0x0c,
	0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x30, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x70, 0x69, 0x6e, 0x6e,
	0x61, 0x6b, 0x65, 0x72, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x69,
	0x65, 0x73, 0x12, 0x38, 0x0a, 0x09, 0x64, 0x65, 0x63, 0x6b, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x70,
	0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x63, 0x6b, 0x50, 0x72, 0x6f, 0x78,
	0x79, 0x52, 0x09, 0x64, 0x65, 0x63, 0x6b, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x1a, 0x53, 0x0a, 0x0c,
	0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2d,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72, 0x2e,
	0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x1a, 0x5c, 0x0a, 0x11, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x65,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x31, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x73, 0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x6f, 0x72, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0xd9, 0x01, 0x0a, 0x09, 0x44, 0x65, 0x63, 0x6b, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x12, 0x34, 0x0a,
	0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x12, 0x41, 0x0a, 0x07, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x70, 0x69,
	0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x63, 0x6b, 0x50, 0x72, 0x6f, 0x78, 0x79,
	0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x1a, 0x53, 0x0a, 0x0c, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x73, 0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xd0, 0x01, 0x0a, 0x06,
	0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x34, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73,
	0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x1a, 0x39, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x2e,
	0x0a, 0x0a, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x42, 0x31,
	0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x69,
	0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72, 0x2f, 0x6b, 0x6c, 0x65, 0x61, 0x74, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x73, 0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spinnaker_extensibility_proto_rawDescOnce sync.Once
	file_spinnaker_extensibility_proto_rawDescData = file_spinnaker_extensibility_proto_rawDesc
)

func file_spinnaker_extensibility_proto_rawDescGZIP() []byte {
	file_spinnaker_extensibility_proto_rawDescOnce.Do(func() {
		file_spinnaker_extensibility_proto_rawDescData = protoimpl.X.CompressGZIP(file_spinnaker_extensibility_proto_rawDescData)
	})
	return file_spinnaker_extensibility_proto_rawDescData
}

var file_spinnaker_extensibility_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_spinnaker_extensibility_proto_goTypes = []interface{}{
	(*Spinnaker)(nil),            // 0: proto.spinnaker.Spinnaker
	(*Extensibility)(nil),        // 1: proto.spinnaker.Extensibility
	(*DeckProxy)(nil),            // 2: proto.spinnaker.DeckProxy
	(*Plugin)(nil),               // 3: proto.spinnaker.Plugin
	(*Repository)(nil),           // 4: proto.spinnaker.Repository
	nil,                          // 5: proto.spinnaker.Extensibility.PluginsEntry
	nil,                          // 6: proto.spinnaker.Extensibility.RepositoriesEntry
	nil,                          // 7: proto.spinnaker.DeckProxy.PluginsEntry
	nil,                          // 8: proto.spinnaker.Plugin.ConfigEntry
	(*wrapperspb.BoolValue)(nil), // 9: google.protobuf.BoolValue
}
var file_spinnaker_extensibility_proto_depIdxs = []int32{
	1,  // 0: proto.spinnaker.Spinnaker.extensibility:type_name -> proto.spinnaker.Extensibility
	5,  // 1: proto.spinnaker.Extensibility.plugins:type_name -> proto.spinnaker.Extensibility.PluginsEntry
	6,  // 2: proto.spinnaker.Extensibility.repositories:type_name -> proto.spinnaker.Extensibility.RepositoriesEntry
	2,  // 3: proto.spinnaker.Extensibility.deckProxy:type_name -> proto.spinnaker.DeckProxy
	9,  // 4: proto.spinnaker.DeckProxy.enabled:type_name -> google.protobuf.BoolValue
	7,  // 5: proto.spinnaker.DeckProxy.plugins:type_name -> proto.spinnaker.DeckProxy.PluginsEntry
	9,  // 6: proto.spinnaker.Plugin.enabled:type_name -> google.protobuf.BoolValue
	8,  // 7: proto.spinnaker.Plugin.config:type_name -> proto.spinnaker.Plugin.ConfigEntry
	3,  // 8: proto.spinnaker.Extensibility.PluginsEntry.value:type_name -> proto.spinnaker.Plugin
	4,  // 9: proto.spinnaker.Extensibility.RepositoriesEntry.value:type_name -> proto.spinnaker.Repository
	3,  // 10: proto.spinnaker.DeckProxy.PluginsEntry.value:type_name -> proto.spinnaker.Plugin
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_spinnaker_extensibility_proto_init() }
func file_spinnaker_extensibility_proto_init() {
	if File_spinnaker_extensibility_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spinnaker_extensibility_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Spinnaker); i {
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
		file_spinnaker_extensibility_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Extensibility); i {
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
		file_spinnaker_extensibility_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeckProxy); i {
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
		file_spinnaker_extensibility_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Plugin); i {
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
		file_spinnaker_extensibility_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Repository); i {
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
			RawDescriptor: file_spinnaker_extensibility_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_spinnaker_extensibility_proto_goTypes,
		DependencyIndexes: file_spinnaker_extensibility_proto_depIdxs,
		MessageInfos:      file_spinnaker_extensibility_proto_msgTypes,
	}.Build()
	File_spinnaker_extensibility_proto = out.File
	file_spinnaker_extensibility_proto_rawDesc = nil
	file_spinnaker_extensibility_proto_goTypes = nil
	file_spinnaker_extensibility_proto_depIdxs = nil
}
