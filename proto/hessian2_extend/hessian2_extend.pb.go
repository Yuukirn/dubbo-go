//
// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.3
// source: hessian2_extend.proto

package hessian2_extend

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Hessian2MessageOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JavaClassName string `protobuf:"bytes,1,opt,name=java_class_name,json=javaClassName,proto3" json:"java_class_name,omitempty"`
	ReferencePath string `protobuf:"bytes,2,opt,name=reference_path,json=referencePath,proto3" json:"reference_path,omitempty"`
	IsInheritance bool   `protobuf:"varint,3,opt,name=is_inheritance,json=isInheritance,proto3" json:"is_inheritance,omitempty"`
	ExtendArgs    bool   `protobuf:"varint,4,opt,name=extend_args,json=extendArgs,proto3" json:"extend_args,omitempty"`
}

func (x *Hessian2MessageOptions) Reset() {
	*x = Hessian2MessageOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hessian2_extend_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hessian2MessageOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hessian2MessageOptions) ProtoMessage() {}

func (x *Hessian2MessageOptions) ProtoReflect() protoreflect.Message {
	mi := &file_hessian2_extend_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hessian2MessageOptions.ProtoReflect.Descriptor instead.
func (*Hessian2MessageOptions) Descriptor() ([]byte, []int) {
	return file_hessian2_extend_proto_rawDescGZIP(), []int{0}
}

func (x *Hessian2MessageOptions) GetJavaClassName() string {
	if x != nil {
		return x.JavaClassName
	}
	return ""
}

func (x *Hessian2MessageOptions) GetReferencePath() string {
	if x != nil {
		return x.ReferencePath
	}
	return ""
}

func (x *Hessian2MessageOptions) GetIsInheritance() bool {
	if x != nil {
		return x.IsInheritance
	}
	return false
}

func (x *Hessian2MessageOptions) GetExtendArgs() bool {
	if x != nil {
		return x.ExtendArgs
	}
	return false
}

type Hessian2MethodOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MethodName string `protobuf:"bytes,1,opt,name=method_name,json=methodName,proto3" json:"method_name,omitempty"`
}

func (x *Hessian2MethodOptions) Reset() {
	*x = Hessian2MethodOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hessian2_extend_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hessian2MethodOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hessian2MethodOptions) ProtoMessage() {}

func (x *Hessian2MethodOptions) ProtoReflect() protoreflect.Message {
	mi := &file_hessian2_extend_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hessian2MethodOptions.ProtoReflect.Descriptor instead.
func (*Hessian2MethodOptions) Descriptor() ([]byte, []int) {
	return file_hessian2_extend_proto_rawDescGZIP(), []int{1}
}

func (x *Hessian2MethodOptions) GetMethodName() string {
	if x != nil {
		return x.MethodName
	}
	return ""
}

type Hessian2ServiceOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InterfaceName string `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
}

func (x *Hessian2ServiceOptions) Reset() {
	*x = Hessian2ServiceOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hessian2_extend_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hessian2ServiceOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hessian2ServiceOptions) ProtoMessage() {}

func (x *Hessian2ServiceOptions) ProtoReflect() protoreflect.Message {
	mi := &file_hessian2_extend_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hessian2ServiceOptions.ProtoReflect.Descriptor instead.
func (*Hessian2ServiceOptions) Descriptor() ([]byte, []int) {
	return file_hessian2_extend_proto_rawDescGZIP(), []int{2}
}

func (x *Hessian2ServiceOptions) GetInterfaceName() string {
	if x != nil {
		return x.InterfaceName
	}
	return ""
}

type Hessian2EnumOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JavaClassName string `protobuf:"bytes,1,opt,name=java_class_name,json=javaClassName,proto3" json:"java_class_name,omitempty"`
}

func (x *Hessian2EnumOptions) Reset() {
	*x = Hessian2EnumOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hessian2_extend_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hessian2EnumOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hessian2EnumOptions) ProtoMessage() {}

func (x *Hessian2EnumOptions) ProtoReflect() protoreflect.Message {
	mi := &file_hessian2_extend_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hessian2EnumOptions.ProtoReflect.Descriptor instead.
func (*Hessian2EnumOptions) Descriptor() ([]byte, []int) {
	return file_hessian2_extend_proto_rawDescGZIP(), []int{3}
}

func (x *Hessian2EnumOptions) GetJavaClassName() string {
	if x != nil {
		return x.JavaClassName
	}
	return ""
}

type Hessian2FieldOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsWrapper bool `protobuf:"varint,1,opt,name=is_wrapper,json=isWrapper,proto3" json:"is_wrapper,omitempty"`
}

func (x *Hessian2FieldOptions) Reset() {
	*x = Hessian2FieldOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hessian2_extend_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hessian2FieldOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hessian2FieldOptions) ProtoMessage() {}

func (x *Hessian2FieldOptions) ProtoReflect() protoreflect.Message {
	mi := &file_hessian2_extend_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hessian2FieldOptions.ProtoReflect.Descriptor instead.
func (*Hessian2FieldOptions) Descriptor() ([]byte, []int) {
	return file_hessian2_extend_proto_rawDescGZIP(), []int{4}
}

func (x *Hessian2FieldOptions) GetIsWrapper() bool {
	if x != nil {
		return x.IsWrapper
	}
	return false
}

var file_hessian2_extend_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*Hessian2MessageOptions)(nil),
		Field:         12345,
		Name:          "hessian2_extend.message_extend",
		Tag:           "bytes,12345,opt,name=message_extend",
		Filename:      "hessian2_extend.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*Hessian2MethodOptions)(nil),
		Field:         12345,
		Name:          "hessian2_extend.method_extend",
		Tag:           "bytes,12345,opt,name=method_extend",
		Filename:      "hessian2_extend.proto",
	},
	{
		ExtendedType:  (*descriptorpb.ServiceOptions)(nil),
		ExtensionType: (*Hessian2ServiceOptions)(nil),
		Field:         12345,
		Name:          "hessian2_extend.service_extend",
		Tag:           "bytes,12345,opt,name=service_extend",
		Filename:      "hessian2_extend.proto",
	},
	{
		ExtendedType:  (*descriptorpb.EnumOptions)(nil),
		ExtensionType: (*Hessian2EnumOptions)(nil),
		Field:         12345,
		Name:          "hessian2_extend.enum_extend",
		Tag:           "bytes,12345,opt,name=enum_extend",
		Filename:      "hessian2_extend.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*Hessian2FieldOptions)(nil),
		Field:         12345,
		Name:          "hessian2_extend.field_extend",
		Tag:           "bytes,12345,opt,name=field_extend",
		Filename:      "hessian2_extend.proto",
	},
}

// Extension fields to descriptorpb.MessageOptions.
var (
	// optional hessian2_extend.Hessian2MessageOptions message_extend = 12345;
	E_MessageExtend = &file_hessian2_extend_proto_extTypes[0]
)

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional hessian2_extend.Hessian2MethodOptions method_extend = 12345;
	E_MethodExtend = &file_hessian2_extend_proto_extTypes[1]
)

// Extension fields to descriptorpb.ServiceOptions.
var (
	// optional hessian2_extend.Hessian2ServiceOptions service_extend = 12345;
	E_ServiceExtend = &file_hessian2_extend_proto_extTypes[2]
)

// Extension fields to descriptorpb.EnumOptions.
var (
	// optional hessian2_extend.Hessian2EnumOptions enum_extend = 12345;
	E_EnumExtend = &file_hessian2_extend_proto_extTypes[3]
)

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional hessian2_extend.Hessian2FieldOptions field_extend = 12345;
	E_FieldExtend = &file_hessian2_extend_proto_extTypes[4]
)

var File_hessian2_extend_proto protoreflect.FileDescriptor

var file_hessian2_extend_proto_rawDesc = []byte{
	0x0a, 0x15, 0x68, 0x65, 0x73, 0x73, 0x69, 0x61, 0x6e, 0x32, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x68, 0x65, 0x73, 0x73, 0x69, 0x61, 0x6e,
	0x32, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaf, 0x01, 0x0a, 0x16, 0x48,
	0x65, 0x73, 0x73, 0x69, 0x61, 0x6e, 0x32, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6a, 0x61, 0x76, 0x61, 0x5f, 0x63, 0x6c,
	0x61, 0x73, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x6a, 0x61, 0x76, 0x61, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a,
	0x0e, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x50, 0x61, 0x74, 0x68, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x73, 0x5f, 0x69, 0x6e, 0x68, 0x65, 0x72,
	0x69, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x69, 0x73,
	0x49, 0x6e, 0x68, 0x65, 0x72, 0x69, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x72, 0x67, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x41, 0x72, 0x67, 0x73, 0x22, 0x38, 0x0a, 0x15,
	0x48, 0x65, 0x73, 0x73, 0x69, 0x61, 0x6e, 0x32, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x3f, 0x0a, 0x16, 0x48, 0x65, 0x73, 0x73, 0x69, 0x61,
	0x6e, 0x32, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x3d, 0x0a, 0x13, 0x48, 0x65, 0x73, 0x73, 0x69,
	0x61, 0x6e, 0x32, 0x45, 0x6e, 0x75, 0x6d, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x26,
	0x0a, 0x0f, 0x6a, 0x61, 0x76, 0x61, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6a, 0x61, 0x76, 0x61, 0x43, 0x6c, 0x61,
	0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x35, 0x0a, 0x14, 0x48, 0x65, 0x73, 0x73, 0x69, 0x61,
	0x6e, 0x32, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1d,
	0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x3a, 0x73, 0x0a,
	0x0e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x12,
	0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0xb9, 0x60, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x68, 0x65, 0x73, 0x73, 0x69, 0x61,
	0x6e, 0x32, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x48, 0x65, 0x73, 0x73, 0x69, 0x61,
	0x6e, 0x32, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x88,
	0x01, 0x01, 0x3a, 0x6f, 0x0a, 0x0d, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x64, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0xb9, 0x60, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x68, 0x65, 0x73,
	0x73, 0x69, 0x61, 0x6e, 0x32, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x48, 0x65, 0x73,
	0x73, 0x69, 0x61, 0x6e, 0x32, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x0c, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64,
	0x88, 0x01, 0x01, 0x3a, 0x73, 0x0a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x64, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xb9, 0x60, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e,
	0x68, 0x65, 0x73, 0x73, 0x69, 0x61, 0x6e, 0x32, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e,
	0x48, 0x65, 0x73, 0x73, 0x69, 0x61, 0x6e, 0x32, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x45,
	0x78, 0x74, 0x65, 0x6e, 0x64, 0x88, 0x01, 0x01, 0x3a, 0x67, 0x0a, 0x0b, 0x65, 0x6e, 0x75, 0x6d,
	0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xb9, 0x60, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x68,
	0x65, 0x73, 0x73, 0x69, 0x61, 0x6e, 0x32, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x48,
	0x65, 0x73, 0x73, 0x69, 0x61, 0x6e, 0x32, 0x45, 0x6e, 0x75, 0x6d, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x0a, 0x65, 0x6e, 0x75, 0x6d, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x88, 0x01,
	0x01, 0x3a, 0x6b, 0x0a, 0x0c, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x64, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0xb9, 0x60, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x68, 0x65, 0x73, 0x73, 0x69, 0x61,
	0x6e, 0x32, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x48, 0x65, 0x73, 0x73, 0x69, 0x61,
	0x6e, 0x32, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0b,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x88, 0x01, 0x01, 0x42, 0x44,
	0x5a, 0x42, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x6f,
	0x72, 0x67, 0x2f, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x33, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x65, 0x73, 0x73, 0x69, 0x61, 0x6e, 0x32, 0x5f, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x64, 0x3b, 0x68, 0x65, 0x73, 0x73, 0x69, 0x61, 0x6e, 0x32, 0x5f, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hessian2_extend_proto_rawDescOnce sync.Once
	file_hessian2_extend_proto_rawDescData = file_hessian2_extend_proto_rawDesc
)

func file_hessian2_extend_proto_rawDescGZIP() []byte {
	file_hessian2_extend_proto_rawDescOnce.Do(func() {
		file_hessian2_extend_proto_rawDescData = protoimpl.X.CompressGZIP(file_hessian2_extend_proto_rawDescData)
	})
	return file_hessian2_extend_proto_rawDescData
}

var file_hessian2_extend_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_hessian2_extend_proto_goTypes = []interface{}{
	(*Hessian2MessageOptions)(nil),      // 0: hessian2_extend.Hessian2MessageOptions
	(*Hessian2MethodOptions)(nil),       // 1: hessian2_extend.Hessian2MethodOptions
	(*Hessian2ServiceOptions)(nil),      // 2: hessian2_extend.Hessian2ServiceOptions
	(*Hessian2EnumOptions)(nil),         // 3: hessian2_extend.Hessian2EnumOptions
	(*Hessian2FieldOptions)(nil),        // 4: hessian2_extend.Hessian2FieldOptions
	(*descriptorpb.MessageOptions)(nil), // 5: google.protobuf.MessageOptions
	(*descriptorpb.MethodOptions)(nil),  // 6: google.protobuf.MethodOptions
	(*descriptorpb.ServiceOptions)(nil), // 7: google.protobuf.ServiceOptions
	(*descriptorpb.EnumOptions)(nil),    // 8: google.protobuf.EnumOptions
	(*descriptorpb.FieldOptions)(nil),   // 9: google.protobuf.FieldOptions
}
var file_hessian2_extend_proto_depIdxs = []int32{
	5,  // 0: hessian2_extend.message_extend:extendee -> google.protobuf.MessageOptions
	6,  // 1: hessian2_extend.method_extend:extendee -> google.protobuf.MethodOptions
	7,  // 2: hessian2_extend.service_extend:extendee -> google.protobuf.ServiceOptions
	8,  // 3: hessian2_extend.enum_extend:extendee -> google.protobuf.EnumOptions
	9,  // 4: hessian2_extend.field_extend:extendee -> google.protobuf.FieldOptions
	0,  // 5: hessian2_extend.message_extend:type_name -> hessian2_extend.Hessian2MessageOptions
	1,  // 6: hessian2_extend.method_extend:type_name -> hessian2_extend.Hessian2MethodOptions
	2,  // 7: hessian2_extend.service_extend:type_name -> hessian2_extend.Hessian2ServiceOptions
	3,  // 8: hessian2_extend.enum_extend:type_name -> hessian2_extend.Hessian2EnumOptions
	4,  // 9: hessian2_extend.field_extend:type_name -> hessian2_extend.Hessian2FieldOptions
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	5,  // [5:10] is the sub-list for extension type_name
	0,  // [0:5] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_hessian2_extend_proto_init() }
func file_hessian2_extend_proto_init() {
	if File_hessian2_extend_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hessian2_extend_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hessian2MessageOptions); i {
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
		file_hessian2_extend_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hessian2MethodOptions); i {
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
		file_hessian2_extend_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hessian2ServiceOptions); i {
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
		file_hessian2_extend_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hessian2EnumOptions); i {
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
		file_hessian2_extend_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hessian2FieldOptions); i {
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
			RawDescriptor: file_hessian2_extend_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 5,
			NumServices:   0,
		},
		GoTypes:           file_hessian2_extend_proto_goTypes,
		DependencyIndexes: file_hessian2_extend_proto_depIdxs,
		MessageInfos:      file_hessian2_extend_proto_msgTypes,
		ExtensionInfos:    file_hessian2_extend_proto_extTypes,
	}.Build()
	File_hessian2_extend_proto = out.File
	file_hessian2_extend_proto_rawDesc = nil
	file_hessian2_extend_proto_goTypes = nil
	file_hessian2_extend_proto_depIdxs = nil
}
