// Copyright 2020-2021 Buf Technologies, Inc.
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
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.17.1
// source: buf/alpha/api/v1alpha1/api.proto

package apiv1alpha1

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

// AccessType is the access type.
type AccessType int32

const (
	AccessType_ACCESS_TYPE_UNSPECIFIED AccessType = 0
	AccessType_ACCESS_TYPE_READ        AccessType = 1
	AccessType_ACCESS_TYPE_WRITE       AccessType = 2
)

// Enum value maps for AccessType.
var (
	AccessType_name = map[int32]string{
		0: "ACCESS_TYPE_UNSPECIFIED",
		1: "ACCESS_TYPE_READ",
		2: "ACCESS_TYPE_WRITE",
	}
	AccessType_value = map[string]int32{
		"ACCESS_TYPE_UNSPECIFIED": 0,
		"ACCESS_TYPE_READ":        1,
		"ACCESS_TYPE_WRITE":       2,
	}
)

func (x AccessType) Enum() *AccessType {
	p := new(AccessType)
	*p = x
	return p
}

func (x AccessType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AccessType) Descriptor() protoreflect.EnumDescriptor {
	return file_buf_alpha_api_v1alpha1_api_proto_enumTypes[0].Descriptor()
}

func (AccessType) Type() protoreflect.EnumType {
	return &file_buf_alpha_api_v1alpha1_api_proto_enumTypes[0]
}

func (x AccessType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *AccessType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = AccessType(num)
	return nil
}

// Deprecated: Use AccessType.Descriptor instead.
func (AccessType) EnumDescriptor() ([]byte, []int) {
	return file_buf_alpha_api_v1alpha1_api_proto_rawDescGZIP(), []int{0}
}

var file_buf_alpha_api_v1alpha1_api_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*AccessType)(nil),
		Field:         70001,
		Name:          "buf.alpha.api.v1alpha1.access_type",
		Tag:           "varint,70001,opt,name=access_type,enum=buf.alpha.api.v1alpha1.AccessType",
		Filename:      "buf/alpha/api/v1alpha1/api.proto",
	},
}

// Extension fields to descriptorpb.MethodOptions.
var (
	// access_type specifies the AccessType of the method.
	//
	// optional buf.alpha.api.v1alpha1.AccessType access_type = 70001;
	E_AccessType = &file_buf_alpha_api_v1alpha1_api_proto_extTypes[0]
)

var File_buf_alpha_api_v1alpha1_api_proto protoreflect.FileDescriptor

var file_buf_alpha_api_v1alpha1_api_proto_rawDesc = []byte{
	0x0a, 0x20, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x16, 0x62, 0x75, 0x66, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x56, 0x0a, 0x0a,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x41, 0x43,
	0x43, 0x45, 0x53, 0x53, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x41, 0x43, 0x43, 0x45, 0x53,
	0x53, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x10, 0x01, 0x12, 0x15, 0x0a,
	0x11, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x57, 0x52, 0x49,
	0x54, 0x45, 0x10, 0x02, 0x3a, 0x65, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0xf1, 0xa2, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x62, 0x75,
	0x66, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x0a, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x79, 0x70, 0x65, 0x42, 0x78, 0x0a, 0x1a, 0x63,
	0x6f, 0x6d, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x08, 0x41, 0x70, 0x69, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x62, 0x75, 0x66, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x62, 0x75, 0x66, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x2f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x61, 0x70, 0x69, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31,
}

var (
	file_buf_alpha_api_v1alpha1_api_proto_rawDescOnce sync.Once
	file_buf_alpha_api_v1alpha1_api_proto_rawDescData = file_buf_alpha_api_v1alpha1_api_proto_rawDesc
)

func file_buf_alpha_api_v1alpha1_api_proto_rawDescGZIP() []byte {
	file_buf_alpha_api_v1alpha1_api_proto_rawDescOnce.Do(func() {
		file_buf_alpha_api_v1alpha1_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_buf_alpha_api_v1alpha1_api_proto_rawDescData)
	})
	return file_buf_alpha_api_v1alpha1_api_proto_rawDescData
}

var file_buf_alpha_api_v1alpha1_api_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_buf_alpha_api_v1alpha1_api_proto_goTypes = []interface{}{
	(AccessType)(0),                    // 0: buf.alpha.api.v1alpha1.AccessType
	(*descriptorpb.MethodOptions)(nil), // 1: google.protobuf.MethodOptions
}
var file_buf_alpha_api_v1alpha1_api_proto_depIdxs = []int32{
	1, // 0: buf.alpha.api.v1alpha1.access_type:extendee -> google.protobuf.MethodOptions
	0, // 1: buf.alpha.api.v1alpha1.access_type:type_name -> buf.alpha.api.v1alpha1.AccessType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	1, // [1:2] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_buf_alpha_api_v1alpha1_api_proto_init() }
func file_buf_alpha_api_v1alpha1_api_proto_init() {
	if File_buf_alpha_api_v1alpha1_api_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_buf_alpha_api_v1alpha1_api_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_buf_alpha_api_v1alpha1_api_proto_goTypes,
		DependencyIndexes: file_buf_alpha_api_v1alpha1_api_proto_depIdxs,
		EnumInfos:         file_buf_alpha_api_v1alpha1_api_proto_enumTypes,
		ExtensionInfos:    file_buf_alpha_api_v1alpha1_api_proto_extTypes,
	}.Build()
	File_buf_alpha_api_v1alpha1_api_proto = out.File
	file_buf_alpha_api_v1alpha1_api_proto_rawDesc = nil
	file_buf_alpha_api_v1alpha1_api_proto_goTypes = nil
	file_buf_alpha_api_v1alpha1_api_proto_depIdxs = nil
}
