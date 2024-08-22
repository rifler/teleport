// Copyright 2024 Gravitational, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: teleport/autoupdate/v1/autoupdate.proto

package autoupdate

import (
	v1 "github.com/gravitational/teleport/api/gen/proto/go/teleport/header/v1"
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

// AutoupdateConfig is a config singleton used to configure cluster
// autoupdate settings.
type AutoupdateConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind     string                `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	SubKind  string                `protobuf:"bytes,2,opt,name=sub_kind,json=subKind,proto3" json:"sub_kind,omitempty"`
	Version  string                `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Metadata *v1.Metadata          `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Spec     *AutoupdateConfigSpec `protobuf:"bytes,5,opt,name=spec,proto3" json:"spec,omitempty"`
}

func (x *AutoupdateConfig) Reset() {
	*x = AutoupdateConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_autoupdate_v1_autoupdate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutoupdateConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutoupdateConfig) ProtoMessage() {}

func (x *AutoupdateConfig) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_autoupdate_v1_autoupdate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutoupdateConfig.ProtoReflect.Descriptor instead.
func (*AutoupdateConfig) Descriptor() ([]byte, []int) {
	return file_teleport_autoupdate_v1_autoupdate_proto_rawDescGZIP(), []int{0}
}

func (x *AutoupdateConfig) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *AutoupdateConfig) GetSubKind() string {
	if x != nil {
		return x.SubKind
	}
	return ""
}

func (x *AutoupdateConfig) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *AutoupdateConfig) GetMetadata() *v1.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *AutoupdateConfig) GetSpec() *AutoupdateConfigSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

// AutoupdateConfigSpec encodes the parameters of the autoupdate config object.
type AutoupdateConfigSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ToolsAutoupdate encodes the feature flag to enable/disable tools autoupdates.
	ToolsAutoupdate bool `protobuf:"varint,1,opt,name=tools_autoupdate,json=toolsAutoupdate,proto3" json:"tools_autoupdate,omitempty"`
}

func (x *AutoupdateConfigSpec) Reset() {
	*x = AutoupdateConfigSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_autoupdate_v1_autoupdate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutoupdateConfigSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutoupdateConfigSpec) ProtoMessage() {}

func (x *AutoupdateConfigSpec) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_autoupdate_v1_autoupdate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutoupdateConfigSpec.ProtoReflect.Descriptor instead.
func (*AutoupdateConfigSpec) Descriptor() ([]byte, []int) {
	return file_teleport_autoupdate_v1_autoupdate_proto_rawDescGZIP(), []int{1}
}

func (x *AutoupdateConfigSpec) GetToolsAutoupdate() bool {
	if x != nil {
		return x.ToolsAutoupdate
	}
	return false
}

// AutoupdateVersion is a resource singleton with version required for
// tools autoupdate.
type AutoupdateVersion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind     string                 `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	SubKind  string                 `protobuf:"bytes,2,opt,name=sub_kind,json=subKind,proto3" json:"sub_kind,omitempty"`
	Version  string                 `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Metadata *v1.Metadata           `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Spec     *AutoupdateVersionSpec `protobuf:"bytes,5,opt,name=spec,proto3" json:"spec,omitempty"`
}

func (x *AutoupdateVersion) Reset() {
	*x = AutoupdateVersion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_autoupdate_v1_autoupdate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutoupdateVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutoupdateVersion) ProtoMessage() {}

func (x *AutoupdateVersion) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_autoupdate_v1_autoupdate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutoupdateVersion.ProtoReflect.Descriptor instead.
func (*AutoupdateVersion) Descriptor() ([]byte, []int) {
	return file_teleport_autoupdate_v1_autoupdate_proto_rawDescGZIP(), []int{2}
}

func (x *AutoupdateVersion) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *AutoupdateVersion) GetSubKind() string {
	if x != nil {
		return x.SubKind
	}
	return ""
}

func (x *AutoupdateVersion) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *AutoupdateVersion) GetMetadata() *v1.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *AutoupdateVersion) GetSpec() *AutoupdateVersionSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

// AutoupdateVersionSpec encodes the parameters of the autoupdate versions.
type AutoupdateVersionSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ToolsVersion is the semantic version required for tools autoupdates.
	ToolsVersion string `protobuf:"bytes,1,opt,name=tools_version,json=toolsVersion,proto3" json:"tools_version,omitempty"`
}

func (x *AutoupdateVersionSpec) Reset() {
	*x = AutoupdateVersionSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_autoupdate_v1_autoupdate_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutoupdateVersionSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutoupdateVersionSpec) ProtoMessage() {}

func (x *AutoupdateVersionSpec) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_autoupdate_v1_autoupdate_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutoupdateVersionSpec.ProtoReflect.Descriptor instead.
func (*AutoupdateVersionSpec) Descriptor() ([]byte, []int) {
	return file_teleport_autoupdate_v1_autoupdate_proto_rawDescGZIP(), []int{3}
}

func (x *AutoupdateVersionSpec) GetToolsVersion() string {
	if x != nil {
		return x.ToolsVersion
	}
	return ""
}

var File_teleport_autoupdate_v1_autoupdate_proto protoreflect.FileDescriptor

var file_teleport_autoupdate_v1_autoupdate_proto_rawDesc = []byte{
	0x0a, 0x27, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x74, 0x65, 0x6c, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x76,
	0x31, 0x1a, 0x21, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd7, 0x01, 0x0a, 0x10, 0x41, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x19, 0x0a,
	0x08, 0x73, 0x75, 0x62, 0x5f, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x73, 0x75, 0x62, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x40, 0x0a, 0x04,
	0x73, 0x70, 0x65, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x74, 0x65, 0x6c,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x22, 0x41,
	0x0a, 0x14, 0x41, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x53, 0x70, 0x65, 0x63, 0x12, 0x29, 0x0a, 0x10, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x5f,
	0x61, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0f, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x41, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x22, 0xd9, 0x01, 0x0a, 0x11, 0x41, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73,
	0x75, 0x62, 0x5f, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73,
	0x75, 0x62, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x38, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x41, 0x0a, 0x04, 0x73, 0x70,
	0x65, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x22, 0x3c, 0x0a,
	0x15, 0x41, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x5f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74,
	0x6f, 0x6f, 0x6c, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x56, 0x5a, 0x54, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x61, 0x76, 0x69, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_teleport_autoupdate_v1_autoupdate_proto_rawDescOnce sync.Once
	file_teleport_autoupdate_v1_autoupdate_proto_rawDescData = file_teleport_autoupdate_v1_autoupdate_proto_rawDesc
)

func file_teleport_autoupdate_v1_autoupdate_proto_rawDescGZIP() []byte {
	file_teleport_autoupdate_v1_autoupdate_proto_rawDescOnce.Do(func() {
		file_teleport_autoupdate_v1_autoupdate_proto_rawDescData = protoimpl.X.CompressGZIP(file_teleport_autoupdate_v1_autoupdate_proto_rawDescData)
	})
	return file_teleport_autoupdate_v1_autoupdate_proto_rawDescData
}

var file_teleport_autoupdate_v1_autoupdate_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_teleport_autoupdate_v1_autoupdate_proto_goTypes = []any{
	(*AutoupdateConfig)(nil),      // 0: teleport.autoupdate.v1.AutoupdateConfig
	(*AutoupdateConfigSpec)(nil),  // 1: teleport.autoupdate.v1.AutoupdateConfigSpec
	(*AutoupdateVersion)(nil),     // 2: teleport.autoupdate.v1.AutoupdateVersion
	(*AutoupdateVersionSpec)(nil), // 3: teleport.autoupdate.v1.AutoupdateVersionSpec
	(*v1.Metadata)(nil),           // 4: teleport.header.v1.Metadata
}
var file_teleport_autoupdate_v1_autoupdate_proto_depIdxs = []int32{
	4, // 0: teleport.autoupdate.v1.AutoupdateConfig.metadata:type_name -> teleport.header.v1.Metadata
	1, // 1: teleport.autoupdate.v1.AutoupdateConfig.spec:type_name -> teleport.autoupdate.v1.AutoupdateConfigSpec
	4, // 2: teleport.autoupdate.v1.AutoupdateVersion.metadata:type_name -> teleport.header.v1.Metadata
	3, // 3: teleport.autoupdate.v1.AutoupdateVersion.spec:type_name -> teleport.autoupdate.v1.AutoupdateVersionSpec
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_teleport_autoupdate_v1_autoupdate_proto_init() }
func file_teleport_autoupdate_v1_autoupdate_proto_init() {
	if File_teleport_autoupdate_v1_autoupdate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_teleport_autoupdate_v1_autoupdate_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*AutoupdateConfig); i {
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
		file_teleport_autoupdate_v1_autoupdate_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AutoupdateConfigSpec); i {
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
		file_teleport_autoupdate_v1_autoupdate_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*AutoupdateVersion); i {
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
		file_teleport_autoupdate_v1_autoupdate_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*AutoupdateVersionSpec); i {
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
			RawDescriptor: file_teleport_autoupdate_v1_autoupdate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_teleport_autoupdate_v1_autoupdate_proto_goTypes,
		DependencyIndexes: file_teleport_autoupdate_v1_autoupdate_proto_depIdxs,
		MessageInfos:      file_teleport_autoupdate_v1_autoupdate_proto_msgTypes,
	}.Build()
	File_teleport_autoupdate_v1_autoupdate_proto = out.File
	file_teleport_autoupdate_v1_autoupdate_proto_rawDesc = nil
	file_teleport_autoupdate_v1_autoupdate_proto_goTypes = nil
	file_teleport_autoupdate_v1_autoupdate_proto_depIdxs = nil
}