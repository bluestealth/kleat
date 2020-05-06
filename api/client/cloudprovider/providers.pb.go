// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.2
// source: cloudprovider/providers.proto

package cloudprovider

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

// Configuration for cloud provider integrations.
type Providers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kubernetes     *Kubernetes          `protobuf:"bytes,1,opt,name=kubernetes,proto3" json:"kubernetes,omitempty"`
	Google         *GoogleComputeEngine `protobuf:"bytes,2,opt,name=google,proto3" json:"google,omitempty"`
	Appengine      *Appengine           `protobuf:"bytes,3,opt,name=appengine,proto3" json:"appengine,omitempty"`
	Aws            *Aws                 `protobuf:"bytes,4,opt,name=aws,proto3" json:"aws,omitempty"`
	Azure          *Azure               `protobuf:"bytes,5,opt,name=azure,proto3" json:"azure,omitempty"`
	Cloudfoundry   *CloudFoundry        `protobuf:"bytes,6,opt,name=cloudfoundry,proto3" json:"cloudfoundry,omitempty"`
	Dcos           *Dcos                `protobuf:"bytes,7,opt,name=dcos,proto3" json:"dcos,omitempty"`
	DockerRegistry *DockerRegistry      `protobuf:"bytes,8,opt,name=dockerRegistry,proto3" json:"dockerRegistry,omitempty"`
	Ecs            *Ecs                 `protobuf:"bytes,9,opt,name=ecs,proto3" json:"ecs,omitempty"`
	Huaweicloud    *HuaweiCloud         `protobuf:"bytes,10,opt,name=huaweicloud,proto3" json:"huaweicloud,omitempty"`
	Oracle         *Oracle              `protobuf:"bytes,11,opt,name=oracle,proto3" json:"oracle,omitempty"`
}

func (x *Providers) Reset() {
	*x = Providers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_providers_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Providers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Providers) ProtoMessage() {}

func (x *Providers) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_providers_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Providers.ProtoReflect.Descriptor instead.
func (*Providers) Descriptor() ([]byte, []int) {
	return file_cloudprovider_providers_proto_rawDescGZIP(), []int{0}
}

func (x *Providers) GetKubernetes() *Kubernetes {
	if x != nil {
		return x.Kubernetes
	}
	return nil
}

func (x *Providers) GetGoogle() *GoogleComputeEngine {
	if x != nil {
		return x.Google
	}
	return nil
}

func (x *Providers) GetAppengine() *Appengine {
	if x != nil {
		return x.Appengine
	}
	return nil
}

func (x *Providers) GetAws() *Aws {
	if x != nil {
		return x.Aws
	}
	return nil
}

func (x *Providers) GetAzure() *Azure {
	if x != nil {
		return x.Azure
	}
	return nil
}

func (x *Providers) GetCloudfoundry() *CloudFoundry {
	if x != nil {
		return x.Cloudfoundry
	}
	return nil
}

func (x *Providers) GetDcos() *Dcos {
	if x != nil {
		return x.Dcos
	}
	return nil
}

func (x *Providers) GetDockerRegistry() *DockerRegistry {
	if x != nil {
		return x.DockerRegistry
	}
	return nil
}

func (x *Providers) GetEcs() *Ecs {
	if x != nil {
		return x.Ecs
	}
	return nil
}

func (x *Providers) GetHuaweicloud() *HuaweiCloud {
	if x != nil {
		return x.Huaweicloud
	}
	return nil
}

func (x *Providers) GetOracle() *Oracle {
	if x != nil {
		return x.Oracle
	}
	return nil
}

var File_cloudprovider_providers_proto protoreflect.FileDescriptor

var file_cloudprovider_providers_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x1a, 0x1d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x2f, 0x61, 0x77, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x61, 0x7a, 0x75, 0x72,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66, 0x6f, 0x75, 0x6e,
	0x64, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x64, 0x63, 0x6f, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x2f, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x65, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1a, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x68, 0x75, 0x61,
	0x77, 0x65, 0x69, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x6b, 0x75,
	0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x6f, 0x72,
	0x61, 0x63, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x92, 0x05, 0x0a, 0x09, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x12, 0x3f, 0x0a, 0x0a, 0x6b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x52, 0x0a, 0x6b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x12, 0x40, 0x0a, 0x06, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x45, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x52, 0x06, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x12, 0x3c, 0x0a, 0x09, 0x61,
	0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x2e, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x52, 0x09,
	0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x12, 0x2a, 0x0a, 0x03, 0x61, 0x77, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x41, 0x77, 0x73,
	0x52, 0x03, 0x61, 0x77, 0x73, 0x12, 0x30, 0x0a, 0x05, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x41, 0x7a, 0x75, 0x72, 0x65,
	0x52, 0x05, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x12, 0x45, 0x0a, 0x0c, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x66, 0x6f, 0x75, 0x6e, 0x64, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x72, 0x79,
	0x52, 0x0c, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x72, 0x79, 0x12, 0x2d,
	0x0a, 0x04, 0x64, 0x63, 0x6f, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x2e, 0x44, 0x63, 0x6f, 0x73, 0x52, 0x04, 0x64, 0x63, 0x6f, 0x73, 0x12, 0x4b, 0x0a,
	0x0e, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x44, 0x6f, 0x63, 0x6b,
	0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x52, 0x0e, 0x64, 0x6f, 0x63, 0x6b,
	0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12, 0x2a, 0x0a, 0x03, 0x65, 0x63,
	0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x45, 0x63,
	0x73, 0x52, 0x03, 0x65, 0x63, 0x73, 0x12, 0x42, 0x0a, 0x0b, 0x68, 0x75, 0x61, 0x77, 0x65, 0x69,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x2e, 0x48, 0x75, 0x61, 0x77, 0x65, 0x69, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x52, 0x0b, 0x68,
	0x75, 0x61, 0x77, 0x65, 0x69, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x12, 0x33, 0x0a, 0x06, 0x6f, 0x72,
	0x61, 0x63, 0x6c, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x2e, 0x4f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x52, 0x06, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x42,
	0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70,
	0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72, 0x2f, 0x6b, 0x6c, 0x65, 0x61, 0x74, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloudprovider_providers_proto_rawDescOnce sync.Once
	file_cloudprovider_providers_proto_rawDescData = file_cloudprovider_providers_proto_rawDesc
)

func file_cloudprovider_providers_proto_rawDescGZIP() []byte {
	file_cloudprovider_providers_proto_rawDescOnce.Do(func() {
		file_cloudprovider_providers_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloudprovider_providers_proto_rawDescData)
	})
	return file_cloudprovider_providers_proto_rawDescData
}

var file_cloudprovider_providers_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_cloudprovider_providers_proto_goTypes = []interface{}{
	(*Providers)(nil),           // 0: proto.cloudprovider.Providers
	(*Kubernetes)(nil),          // 1: proto.cloudprovider.Kubernetes
	(*GoogleComputeEngine)(nil), // 2: proto.cloudprovider.GoogleComputeEngine
	(*Appengine)(nil),           // 3: proto.cloudprovider.Appengine
	(*Aws)(nil),                 // 4: proto.cloudprovider.Aws
	(*Azure)(nil),               // 5: proto.cloudprovider.Azure
	(*CloudFoundry)(nil),        // 6: proto.cloudprovider.CloudFoundry
	(*Dcos)(nil),                // 7: proto.cloudprovider.Dcos
	(*DockerRegistry)(nil),      // 8: proto.cloudprovider.DockerRegistry
	(*Ecs)(nil),                 // 9: proto.cloudprovider.Ecs
	(*HuaweiCloud)(nil),         // 10: proto.cloudprovider.HuaweiCloud
	(*Oracle)(nil),              // 11: proto.cloudprovider.Oracle
}
var file_cloudprovider_providers_proto_depIdxs = []int32{
	1,  // 0: proto.cloudprovider.Providers.kubernetes:type_name -> proto.cloudprovider.Kubernetes
	2,  // 1: proto.cloudprovider.Providers.google:type_name -> proto.cloudprovider.GoogleComputeEngine
	3,  // 2: proto.cloudprovider.Providers.appengine:type_name -> proto.cloudprovider.Appengine
	4,  // 3: proto.cloudprovider.Providers.aws:type_name -> proto.cloudprovider.Aws
	5,  // 4: proto.cloudprovider.Providers.azure:type_name -> proto.cloudprovider.Azure
	6,  // 5: proto.cloudprovider.Providers.cloudfoundry:type_name -> proto.cloudprovider.CloudFoundry
	7,  // 6: proto.cloudprovider.Providers.dcos:type_name -> proto.cloudprovider.Dcos
	8,  // 7: proto.cloudprovider.Providers.dockerRegistry:type_name -> proto.cloudprovider.DockerRegistry
	9,  // 8: proto.cloudprovider.Providers.ecs:type_name -> proto.cloudprovider.Ecs
	10, // 9: proto.cloudprovider.Providers.huaweicloud:type_name -> proto.cloudprovider.HuaweiCloud
	11, // 10: proto.cloudprovider.Providers.oracle:type_name -> proto.cloudprovider.Oracle
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_cloudprovider_providers_proto_init() }
func file_cloudprovider_providers_proto_init() {
	if File_cloudprovider_providers_proto != nil {
		return
	}
	file_cloudprovider_appengine_proto_init()
	file_cloudprovider_aws_proto_init()
	file_cloudprovider_azure_proto_init()
	file_cloudprovider_cloudfoundry_proto_init()
	file_cloudprovider_dcos_proto_init()
	file_cloudprovider_docker_registry_proto_init()
	file_cloudprovider_ecs_proto_init()
	file_cloudprovider_google_proto_init()
	file_cloudprovider_huaweicloud_proto_init()
	file_cloudprovider_kubernetes_proto_init()
	file_cloudprovider_oracle_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cloudprovider_providers_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Providers); i {
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
			RawDescriptor: file_cloudprovider_providers_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloudprovider_providers_proto_goTypes,
		DependencyIndexes: file_cloudprovider_providers_proto_depIdxs,
		MessageInfos:      file_cloudprovider_providers_proto_msgTypes,
	}.Build()
	File_cloudprovider_providers_proto = out.File
	file_cloudprovider_providers_proto_rawDesc = nil
	file_cloudprovider_providers_proto_goTypes = nil
	file_cloudprovider_providers_proto_depIdxs = nil
}
