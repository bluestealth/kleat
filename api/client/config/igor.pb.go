// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.1
// source: config/igor.proto

package config

import (
	proto "github.com/golang/protobuf/proto"
	artifact "github.com/spinnaker/kleat/api/client/artifact"
	ci "github.com/spinnaker/kleat/api/client/ci"
	repository "github.com/spinnaker/kleat/api/client/repository"
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

// Configuration for the Igor microservice.
type Igor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DockerRegistry *Igor_DockerRegistry    `protobuf:"bytes,1,opt,name=dockerRegistry,proto3" json:"dockerRegistry,omitempty"`
	Artifacts      *Igor_Artifacts         `protobuf:"bytes,2,opt,name=artifacts,proto3" json:"artifacts,omitempty"`
	Artifactory    *repository.Artifactory `protobuf:"bytes,3,opt,name=artifactory,proto3" json:"artifactory,omitempty"`
	Gcb            *ci.GoogleCloudBuild    `protobuf:"bytes,4,opt,name=gcb,proto3" json:"gcb,omitempty"`
	Codebuild      *ci.CodeBuild           `protobuf:"bytes,5,opt,name=codebuild,proto3" json:"codebuild,omitempty"`
	Concourse      *ci.Concourse           `protobuf:"bytes,6,opt,name=concourse,proto3" json:"concourse,omitempty"`
	Jenkins        *ci.Jenkins             `protobuf:"bytes,7,opt,name=jenkins,proto3" json:"jenkins,omitempty"`
	Travis         *ci.Travis              `protobuf:"bytes,8,opt,name=travis,proto3" json:"travis,omitempty"`
	Wercker        *ci.Wercker             `protobuf:"bytes,9,opt,name=wercker,proto3" json:"wercker,omitempty"`
}

func (x *Igor) Reset() {
	*x = Igor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_igor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Igor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Igor) ProtoMessage() {}

func (x *Igor) ProtoReflect() protoreflect.Message {
	mi := &file_config_igor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Igor.ProtoReflect.Descriptor instead.
func (*Igor) Descriptor() ([]byte, []int) {
	return file_config_igor_proto_rawDescGZIP(), []int{0}
}

func (x *Igor) GetDockerRegistry() *Igor_DockerRegistry {
	if x != nil {
		return x.DockerRegistry
	}
	return nil
}

func (x *Igor) GetArtifacts() *Igor_Artifacts {
	if x != nil {
		return x.Artifacts
	}
	return nil
}

func (x *Igor) GetArtifactory() *repository.Artifactory {
	if x != nil {
		return x.Artifactory
	}
	return nil
}

func (x *Igor) GetGcb() *ci.GoogleCloudBuild {
	if x != nil {
		return x.Gcb
	}
	return nil
}

func (x *Igor) GetCodebuild() *ci.CodeBuild {
	if x != nil {
		return x.Codebuild
	}
	return nil
}

func (x *Igor) GetConcourse() *ci.Concourse {
	if x != nil {
		return x.Concourse
	}
	return nil
}

func (x *Igor) GetJenkins() *ci.Jenkins {
	if x != nil {
		return x.Jenkins
	}
	return nil
}

func (x *Igor) GetTravis() *ci.Travis {
	if x != nil {
		return x.Travis
	}
	return nil
}

func (x *Igor) GetWercker() *ci.Wercker {
	if x != nil {
		return x.Wercker
	}
	return nil
}

type Igor_DockerRegistry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
}

func (x *Igor_DockerRegistry) Reset() {
	*x = Igor_DockerRegistry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_igor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Igor_DockerRegistry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Igor_DockerRegistry) ProtoMessage() {}

func (x *Igor_DockerRegistry) ProtoReflect() protoreflect.Message {
	mi := &file_config_igor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Igor_DockerRegistry.ProtoReflect.Descriptor instead.
func (*Igor_DockerRegistry) Descriptor() ([]byte, []int) {
	return file_config_igor_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Igor_DockerRegistry) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

type Igor_Artifacts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Templates []*artifact.Template `protobuf:"bytes,1,rep,name=templates,proto3" json:"templates,omitempty"`
}

func (x *Igor_Artifacts) Reset() {
	*x = Igor_Artifacts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_igor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Igor_Artifacts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Igor_Artifacts) ProtoMessage() {}

func (x *Igor_Artifacts) ProtoReflect() protoreflect.Message {
	mi := &file_config_igor_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Igor_Artifacts.ProtoReflect.Descriptor instead.
func (*Igor_Artifacts) Descriptor() ([]byte, []int) {
	return file_config_igor_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Igor_Artifacts) GetTemplates() []*artifact.Template {
	if x != nil {
		return x.Templates
	}
	return nil
}

var File_config_igor_proto protoreflect.FileDescriptor

var file_config_igor_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x69, 0x67, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x1a, 0x17, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2f, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x72, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x63, 0x69, 0x2f, 0x63, 0x6f, 0x64,
	0x65, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x63, 0x69,
	0x2f, 0x63, 0x6f, 0x6e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0c, 0x63, 0x69, 0x2f, 0x67, 0x63, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10,
	0x63, 0x69, 0x2f, 0x6a, 0x65, 0x6e, 0x6b, 0x69, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0f, 0x63, 0x69, 0x2f, 0x74, 0x72, 0x61, 0x76, 0x69, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x10, 0x63, 0x69, 0x2f, 0x77, 0x65, 0x72, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xd7, 0x04, 0x0a, 0x04, 0x49, 0x67, 0x6f, 0x72, 0x12, 0x49, 0x0a, 0x0e,
	0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x49, 0x67, 0x6f, 0x72, 0x2e, 0x44, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x52, 0x0e, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12, 0x3a, 0x0a, 0x09, 0x61, 0x72, 0x74, 0x69, 0x66,
	0x61, 0x63, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x49, 0x67, 0x6f, 0x72, 0x2e, 0x41,
	0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x52, 0x09, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x73, 0x12, 0x3f, 0x0a, 0x0b, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x41, 0x72, 0x74, 0x69,
	0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x0b, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63,
	0x74, 0x6f, 0x72, 0x79, 0x12, 0x2c, 0x0a, 0x03, 0x67, 0x63, 0x62, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x69, 0x2e, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x03, 0x67,
	0x63, 0x62, 0x12, 0x31, 0x0a, 0x09, 0x63, 0x6f, 0x64, 0x65, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x69,
	0x2e, 0x43, 0x6f, 0x64, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x09, 0x63, 0x6f, 0x64, 0x65,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x12, 0x31, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x63, 0x69, 0x2e, 0x43, 0x6f, 0x6e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x09, 0x63,
	0x6f, 0x6e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x07, 0x6a, 0x65, 0x6e, 0x6b,
	0x69, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x63, 0x69, 0x2e, 0x4a, 0x65, 0x6e, 0x6b, 0x69, 0x6e, 0x73, 0x52, 0x07, 0x6a, 0x65,
	0x6e, 0x6b, 0x69, 0x6e, 0x73, 0x12, 0x28, 0x0a, 0x06, 0x74, 0x72, 0x61, 0x76, 0x69, 0x73, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x69,
	0x2e, 0x54, 0x72, 0x61, 0x76, 0x69, 0x73, 0x52, 0x06, 0x74, 0x72, 0x61, 0x76, 0x69, 0x73, 0x12,
	0x2b, 0x0a, 0x07, 0x77, 0x65, 0x72, 0x63, 0x6b, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x69, 0x2e, 0x57, 0x65, 0x72, 0x63,
	0x6b, 0x65, 0x72, 0x52, 0x07, 0x77, 0x65, 0x72, 0x63, 0x6b, 0x65, 0x72, 0x1a, 0x2a, 0x0a, 0x0e,
	0x44, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12, 0x18,
	0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x1a, 0x43, 0x0a, 0x09, 0x41, 0x72, 0x74, 0x69,
	0x66, 0x61, 0x63, 0x74, 0x73, 0x12, 0x36, 0x0a, 0x09, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x52, 0x09, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x42, 0x2e, 0x5a,
	0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x69, 0x6e,
	0x6e, 0x61, 0x6b, 0x65, 0x72, 0x2f, 0x6b, 0x6c, 0x65, 0x61, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_igor_proto_rawDescOnce sync.Once
	file_config_igor_proto_rawDescData = file_config_igor_proto_rawDesc
)

func file_config_igor_proto_rawDescGZIP() []byte {
	file_config_igor_proto_rawDescOnce.Do(func() {
		file_config_igor_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_igor_proto_rawDescData)
	})
	return file_config_igor_proto_rawDescData
}

var file_config_igor_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_config_igor_proto_goTypes = []interface{}{
	(*Igor)(nil),                   // 0: proto.config.Igor
	(*Igor_DockerRegistry)(nil),    // 1: proto.config.Igor.DockerRegistry
	(*Igor_Artifacts)(nil),         // 2: proto.config.Igor.Artifacts
	(*repository.Artifactory)(nil), // 3: proto.repository.Artifactory
	(*ci.GoogleCloudBuild)(nil),    // 4: proto.ci.GoogleCloudBuild
	(*ci.CodeBuild)(nil),           // 5: proto.ci.CodeBuild
	(*ci.Concourse)(nil),           // 6: proto.ci.Concourse
	(*ci.Jenkins)(nil),             // 7: proto.ci.Jenkins
	(*ci.Travis)(nil),              // 8: proto.ci.Travis
	(*ci.Wercker)(nil),             // 9: proto.ci.Wercker
	(*artifact.Template)(nil),      // 10: proto.artifact.Template
}
var file_config_igor_proto_depIdxs = []int32{
	1,  // 0: proto.config.Igor.dockerRegistry:type_name -> proto.config.Igor.DockerRegistry
	2,  // 1: proto.config.Igor.artifacts:type_name -> proto.config.Igor.Artifacts
	3,  // 2: proto.config.Igor.artifactory:type_name -> proto.repository.Artifactory
	4,  // 3: proto.config.Igor.gcb:type_name -> proto.ci.GoogleCloudBuild
	5,  // 4: proto.config.Igor.codebuild:type_name -> proto.ci.CodeBuild
	6,  // 5: proto.config.Igor.concourse:type_name -> proto.ci.Concourse
	7,  // 6: proto.config.Igor.jenkins:type_name -> proto.ci.Jenkins
	8,  // 7: proto.config.Igor.travis:type_name -> proto.ci.Travis
	9,  // 8: proto.config.Igor.wercker:type_name -> proto.ci.Wercker
	10, // 9: proto.config.Igor.Artifacts.templates:type_name -> proto.artifact.Template
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_config_igor_proto_init() }
func file_config_igor_proto_init() {
	if File_config_igor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_igor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Igor); i {
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
		file_config_igor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Igor_DockerRegistry); i {
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
		file_config_igor_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Igor_Artifacts); i {
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
			RawDescriptor: file_config_igor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_igor_proto_goTypes,
		DependencyIndexes: file_config_igor_proto_depIdxs,
		MessageInfos:      file_config_igor_proto_msgTypes,
	}.Build()
	File_config_igor_proto = out.File
	file_config_igor_proto_rawDesc = nil
	file_config_igor_proto_goTypes = nil
	file_config_igor_proto_depIdxs = nil
}
