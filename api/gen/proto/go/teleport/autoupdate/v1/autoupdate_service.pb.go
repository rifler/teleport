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
// source: teleport/autoupdate/v1/autoupdate_service.proto

package autoupdate

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request for GetAutoupdateConfig.
type GetAutoupdateConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAutoupdateConfigRequest) Reset() {
	*x = GetAutoupdateConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_autoupdate_v1_autoupdate_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAutoupdateConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAutoupdateConfigRequest) ProtoMessage() {}

func (x *GetAutoupdateConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_autoupdate_v1_autoupdate_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAutoupdateConfigRequest.ProtoReflect.Descriptor instead.
func (*GetAutoupdateConfigRequest) Descriptor() ([]byte, []int) {
	return file_teleport_autoupdate_v1_autoupdate_service_proto_rawDescGZIP(), []int{0}
}

// Request for UpsertAutoupdateConfig.
type UpsertAutoupdateConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config *AutoupdateConfig `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *UpsertAutoupdateConfigRequest) Reset() {
	*x = UpsertAutoupdateConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_autoupdate_v1_autoupdate_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsertAutoupdateConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertAutoupdateConfigRequest) ProtoMessage() {}

func (x *UpsertAutoupdateConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_autoupdate_v1_autoupdate_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertAutoupdateConfigRequest.ProtoReflect.Descriptor instead.
func (*UpsertAutoupdateConfigRequest) Descriptor() ([]byte, []int) {
	return file_teleport_autoupdate_v1_autoupdate_service_proto_rawDescGZIP(), []int{1}
}

func (x *UpsertAutoupdateConfigRequest) GetConfig() *AutoupdateConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

// Request for DeleteAutoupdateConfig.
type DeleteAutoupdateConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteAutoupdateConfigRequest) Reset() {
	*x = DeleteAutoupdateConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_autoupdate_v1_autoupdate_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAutoupdateConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAutoupdateConfigRequest) ProtoMessage() {}

func (x *DeleteAutoupdateConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_autoupdate_v1_autoupdate_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAutoupdateConfigRequest.ProtoReflect.Descriptor instead.
func (*DeleteAutoupdateConfigRequest) Descriptor() ([]byte, []int) {
	return file_teleport_autoupdate_v1_autoupdate_service_proto_rawDescGZIP(), []int{2}
}

// Request for GetAutoupdateVersion.
type GetAutoupdateVersionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAutoupdateVersionRequest) Reset() {
	*x = GetAutoupdateVersionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_autoupdate_v1_autoupdate_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAutoupdateVersionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAutoupdateVersionRequest) ProtoMessage() {}

func (x *GetAutoupdateVersionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_autoupdate_v1_autoupdate_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAutoupdateVersionRequest.ProtoReflect.Descriptor instead.
func (*GetAutoupdateVersionRequest) Descriptor() ([]byte, []int) {
	return file_teleport_autoupdate_v1_autoupdate_service_proto_rawDescGZIP(), []int{3}
}

// Request for UpsertAutoupdateVersion.
type UpsertAutoupdateVersionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version *AutoupdateVersion `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *UpsertAutoupdateVersionRequest) Reset() {
	*x = UpsertAutoupdateVersionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_autoupdate_v1_autoupdate_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsertAutoupdateVersionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertAutoupdateVersionRequest) ProtoMessage() {}

func (x *UpsertAutoupdateVersionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_autoupdate_v1_autoupdate_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertAutoupdateVersionRequest.ProtoReflect.Descriptor instead.
func (*UpsertAutoupdateVersionRequest) Descriptor() ([]byte, []int) {
	return file_teleport_autoupdate_v1_autoupdate_service_proto_rawDescGZIP(), []int{4}
}

func (x *UpsertAutoupdateVersionRequest) GetVersion() *AutoupdateVersion {
	if x != nil {
		return x.Version
	}
	return nil
}

// Request for DeleteAutoupdateVersion.
type DeleteAutoupdateVersionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteAutoupdateVersionRequest) Reset() {
	*x = DeleteAutoupdateVersionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_autoupdate_v1_autoupdate_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAutoupdateVersionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAutoupdateVersionRequest) ProtoMessage() {}

func (x *DeleteAutoupdateVersionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_autoupdate_v1_autoupdate_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAutoupdateVersionRequest.ProtoReflect.Descriptor instead.
func (*DeleteAutoupdateVersionRequest) Descriptor() ([]byte, []int) {
	return file_teleport_autoupdate_v1_autoupdate_service_proto_rawDescGZIP(), []int{5}
}

var File_teleport_autoupdate_v1_autoupdate_service_proto protoreflect.FileDescriptor

var file_teleport_autoupdate_v1_autoupdate_service_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x16, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x61, 0x75, 0x74, 0x6f,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x2f, 0x61, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x1c, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x61, 0x0a,
	0x1d, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x41, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x40,
	0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28,
	0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x22, 0x1f, 0x0a, 0x1d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x75, 0x74, 0x6f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x1d, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x65, 0x0a, 0x1e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x41, 0x75, 0x74, 0x6f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x43, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x61,
	0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74,
	0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x20, 0x0a, 0x1e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x41, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0xcd, 0x05, 0x0a, 0x11, 0x41, 0x75,
	0x74, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x73, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x32, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x74, 0x65, 0x6c,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x79, 0x0a, 0x16, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x41, 0x75,
	0x74, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x35,
	0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x41, 0x75,
	0x74, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x61, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x67, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x35, 0x2e, 0x74, 0x65, 0x6c, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x76, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41,
	0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x33, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x61, 0x75, 0x74, 0x6f,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74,
	0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x61, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x7c, 0x0a, 0x17, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x41, 0x75, 0x74, 0x6f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x2e, 0x74, 0x65,
	0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x41, 0x75, 0x74, 0x6f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x61,
	0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74,
	0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x69,
	0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x2e, 0x74, 0x65, 0x6c, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x56, 0x5a, 0x54, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x61, 0x76, 0x69, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f,
	0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_teleport_autoupdate_v1_autoupdate_service_proto_rawDescOnce sync.Once
	file_teleport_autoupdate_v1_autoupdate_service_proto_rawDescData = file_teleport_autoupdate_v1_autoupdate_service_proto_rawDesc
)

func file_teleport_autoupdate_v1_autoupdate_service_proto_rawDescGZIP() []byte {
	file_teleport_autoupdate_v1_autoupdate_service_proto_rawDescOnce.Do(func() {
		file_teleport_autoupdate_v1_autoupdate_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_teleport_autoupdate_v1_autoupdate_service_proto_rawDescData)
	})
	return file_teleport_autoupdate_v1_autoupdate_service_proto_rawDescData
}

var file_teleport_autoupdate_v1_autoupdate_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_teleport_autoupdate_v1_autoupdate_service_proto_goTypes = []any{
	(*GetAutoupdateConfigRequest)(nil),     // 0: teleport.autoupdate.v1.GetAutoupdateConfigRequest
	(*UpsertAutoupdateConfigRequest)(nil),  // 1: teleport.autoupdate.v1.UpsertAutoupdateConfigRequest
	(*DeleteAutoupdateConfigRequest)(nil),  // 2: teleport.autoupdate.v1.DeleteAutoupdateConfigRequest
	(*GetAutoupdateVersionRequest)(nil),    // 3: teleport.autoupdate.v1.GetAutoupdateVersionRequest
	(*UpsertAutoupdateVersionRequest)(nil), // 4: teleport.autoupdate.v1.UpsertAutoupdateVersionRequest
	(*DeleteAutoupdateVersionRequest)(nil), // 5: teleport.autoupdate.v1.DeleteAutoupdateVersionRequest
	(*AutoupdateConfig)(nil),               // 6: teleport.autoupdate.v1.AutoupdateConfig
	(*AutoupdateVersion)(nil),              // 7: teleport.autoupdate.v1.AutoupdateVersion
	(*emptypb.Empty)(nil),                  // 8: google.protobuf.Empty
}
var file_teleport_autoupdate_v1_autoupdate_service_proto_depIdxs = []int32{
	6, // 0: teleport.autoupdate.v1.UpsertAutoupdateConfigRequest.config:type_name -> teleport.autoupdate.v1.AutoupdateConfig
	7, // 1: teleport.autoupdate.v1.UpsertAutoupdateVersionRequest.version:type_name -> teleport.autoupdate.v1.AutoupdateVersion
	0, // 2: teleport.autoupdate.v1.AutoupdateService.GetAutoupdateConfig:input_type -> teleport.autoupdate.v1.GetAutoupdateConfigRequest
	1, // 3: teleport.autoupdate.v1.AutoupdateService.UpsertAutoupdateConfig:input_type -> teleport.autoupdate.v1.UpsertAutoupdateConfigRequest
	2, // 4: teleport.autoupdate.v1.AutoupdateService.DeleteAutoupdateConfig:input_type -> teleport.autoupdate.v1.DeleteAutoupdateConfigRequest
	3, // 5: teleport.autoupdate.v1.AutoupdateService.GetAutoupdateVersion:input_type -> teleport.autoupdate.v1.GetAutoupdateVersionRequest
	4, // 6: teleport.autoupdate.v1.AutoupdateService.UpsertAutoupdateVersion:input_type -> teleport.autoupdate.v1.UpsertAutoupdateVersionRequest
	5, // 7: teleport.autoupdate.v1.AutoupdateService.DeleteAutoupdateVersion:input_type -> teleport.autoupdate.v1.DeleteAutoupdateVersionRequest
	6, // 8: teleport.autoupdate.v1.AutoupdateService.GetAutoupdateConfig:output_type -> teleport.autoupdate.v1.AutoupdateConfig
	6, // 9: teleport.autoupdate.v1.AutoupdateService.UpsertAutoupdateConfig:output_type -> teleport.autoupdate.v1.AutoupdateConfig
	8, // 10: teleport.autoupdate.v1.AutoupdateService.DeleteAutoupdateConfig:output_type -> google.protobuf.Empty
	7, // 11: teleport.autoupdate.v1.AutoupdateService.GetAutoupdateVersion:output_type -> teleport.autoupdate.v1.AutoupdateVersion
	7, // 12: teleport.autoupdate.v1.AutoupdateService.UpsertAutoupdateVersion:output_type -> teleport.autoupdate.v1.AutoupdateVersion
	8, // 13: teleport.autoupdate.v1.AutoupdateService.DeleteAutoupdateVersion:output_type -> google.protobuf.Empty
	8, // [8:14] is the sub-list for method output_type
	2, // [2:8] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_teleport_autoupdate_v1_autoupdate_service_proto_init() }
func file_teleport_autoupdate_v1_autoupdate_service_proto_init() {
	if File_teleport_autoupdate_v1_autoupdate_service_proto != nil {
		return
	}
	file_teleport_autoupdate_v1_autoupdate_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_teleport_autoupdate_v1_autoupdate_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetAutoupdateConfigRequest); i {
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
		file_teleport_autoupdate_v1_autoupdate_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*UpsertAutoupdateConfigRequest); i {
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
		file_teleport_autoupdate_v1_autoupdate_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteAutoupdateConfigRequest); i {
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
		file_teleport_autoupdate_v1_autoupdate_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetAutoupdateVersionRequest); i {
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
		file_teleport_autoupdate_v1_autoupdate_service_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*UpsertAutoupdateVersionRequest); i {
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
		file_teleport_autoupdate_v1_autoupdate_service_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteAutoupdateVersionRequest); i {
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
			RawDescriptor: file_teleport_autoupdate_v1_autoupdate_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_teleport_autoupdate_v1_autoupdate_service_proto_goTypes,
		DependencyIndexes: file_teleport_autoupdate_v1_autoupdate_service_proto_depIdxs,
		MessageInfos:      file_teleport_autoupdate_v1_autoupdate_service_proto_msgTypes,
	}.Build()
	File_teleport_autoupdate_v1_autoupdate_service_proto = out.File
	file_teleport_autoupdate_v1_autoupdate_service_proto_rawDesc = nil
	file_teleport_autoupdate_v1_autoupdate_service_proto_goTypes = nil
	file_teleport_autoupdate_v1_autoupdate_service_proto_depIdxs = nil
}