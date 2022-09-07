// Copyright 2020-2022 Buf Technologies, Inc.
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
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: buf/alpha/registry/v1alpha1/github.proto

package registryv1alpha1

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

type GithubAppConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId string `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
}

func (x *GithubAppConfig) Reset() {
	*x = GithubAppConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_alpha_registry_v1alpha1_github_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GithubAppConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GithubAppConfig) ProtoMessage() {}

func (x *GithubAppConfig) ProtoReflect() protoreflect.Message {
	mi := &file_buf_alpha_registry_v1alpha1_github_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GithubAppConfig.ProtoReflect.Descriptor instead.
func (*GithubAppConfig) Descriptor() ([]byte, []int) {
	return file_buf_alpha_registry_v1alpha1_github_proto_rawDescGZIP(), []int{0}
}

func (x *GithubAppConfig) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

type GetGithubAppConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetGithubAppConfigRequest) Reset() {
	*x = GetGithubAppConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_alpha_registry_v1alpha1_github_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGithubAppConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGithubAppConfigRequest) ProtoMessage() {}

func (x *GetGithubAppConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_buf_alpha_registry_v1alpha1_github_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGithubAppConfigRequest.ProtoReflect.Descriptor instead.
func (*GetGithubAppConfigRequest) Descriptor() ([]byte, []int) {
	return file_buf_alpha_registry_v1alpha1_github_proto_rawDescGZIP(), []int{1}
}

type GetGithubAppConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppConfig *GithubAppConfig `protobuf:"bytes,1,opt,name=app_config,json=appConfig,proto3" json:"app_config,omitempty"`
}

func (x *GetGithubAppConfigResponse) Reset() {
	*x = GetGithubAppConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_alpha_registry_v1alpha1_github_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGithubAppConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGithubAppConfigResponse) ProtoMessage() {}

func (x *GetGithubAppConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_buf_alpha_registry_v1alpha1_github_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGithubAppConfigResponse.ProtoReflect.Descriptor instead.
func (*GetGithubAppConfigResponse) Descriptor() ([]byte, []int) {
	return file_buf_alpha_registry_v1alpha1_github_proto_rawDescGZIP(), []int{2}
}

func (x *GetGithubAppConfigResponse) GetAppConfig() *GithubAppConfig {
	if x != nil {
		return x.AppConfig
	}
	return nil
}

var File_buf_alpha_registry_v1alpha1_github_proto protoreflect.FileDescriptor

var file_buf_alpha_registry_v1alpha1_github_proto_rawDesc = []byte{
	0x0a, 0x28, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x62, 0x75, 0x66, 0x2e,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x22, 0x28, 0x0a, 0x0f, 0x47, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70,
	0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49,
	0x64, 0x22, 0x1b, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x47, 0x69, 0x74, 0x68, 0x75, 0x62, 0x41, 0x70,
	0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x69,
	0x0a, 0x1a, 0x47, 0x65, 0x74, 0x47, 0x69, 0x74, 0x68, 0x75, 0x62, 0x41, 0x70, 0x70, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x0a,
	0x61, 0x70, 0x70, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2c, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x09,
	0x61, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x32, 0x97, 0x01, 0x0a, 0x0d, 0x47, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x85, 0x01, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x47, 0x69, 0x74, 0x68, 0x75, 0x62, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x36, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x47, 0x69, 0x74, 0x68, 0x75, 0x62, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x62, 0x75, 0x66,
	0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x98, 0x02, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x75, 0x66, 0x2e,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0b, 0x47, 0x69, 0x74, 0x68, 0x75, 0x62, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x59, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x62, 0x75, 0x66, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x62, 0x75, 0x66, 0x2f,
	0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x3b, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0xa2, 0x02, 0x03, 0x42, 0x41, 0x52, 0xaa, 0x02, 0x1b, 0x42, 0x75, 0x66, 0x2e, 0x41, 0x6c,
	0x70, 0x68, 0x61, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x56, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x1b, 0x42, 0x75, 0x66, 0x5c, 0x41, 0x6c, 0x70, 0x68,
	0x61, 0x5c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0xe2, 0x02, 0x27, 0x42, 0x75, 0x66, 0x5c, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x5c,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1e,
	0x42, 0x75, 0x66, 0x3a, 0x3a, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x3a, 0x3a, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_buf_alpha_registry_v1alpha1_github_proto_rawDescOnce sync.Once
	file_buf_alpha_registry_v1alpha1_github_proto_rawDescData = file_buf_alpha_registry_v1alpha1_github_proto_rawDesc
)

func file_buf_alpha_registry_v1alpha1_github_proto_rawDescGZIP() []byte {
	file_buf_alpha_registry_v1alpha1_github_proto_rawDescOnce.Do(func() {
		file_buf_alpha_registry_v1alpha1_github_proto_rawDescData = protoimpl.X.CompressGZIP(file_buf_alpha_registry_v1alpha1_github_proto_rawDescData)
	})
	return file_buf_alpha_registry_v1alpha1_github_proto_rawDescData
}

var file_buf_alpha_registry_v1alpha1_github_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_buf_alpha_registry_v1alpha1_github_proto_goTypes = []interface{}{
	(*GithubAppConfig)(nil),            // 0: buf.alpha.registry.v1alpha1.GithubAppConfig
	(*GetGithubAppConfigRequest)(nil),  // 1: buf.alpha.registry.v1alpha1.GetGithubAppConfigRequest
	(*GetGithubAppConfigResponse)(nil), // 2: buf.alpha.registry.v1alpha1.GetGithubAppConfigResponse
}
var file_buf_alpha_registry_v1alpha1_github_proto_depIdxs = []int32{
	0, // 0: buf.alpha.registry.v1alpha1.GetGithubAppConfigResponse.app_config:type_name -> buf.alpha.registry.v1alpha1.GithubAppConfig
	1, // 1: buf.alpha.registry.v1alpha1.GithubService.GetGithubAppConfig:input_type -> buf.alpha.registry.v1alpha1.GetGithubAppConfigRequest
	2, // 2: buf.alpha.registry.v1alpha1.GithubService.GetGithubAppConfig:output_type -> buf.alpha.registry.v1alpha1.GetGithubAppConfigResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_buf_alpha_registry_v1alpha1_github_proto_init() }
func file_buf_alpha_registry_v1alpha1_github_proto_init() {
	if File_buf_alpha_registry_v1alpha1_github_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_buf_alpha_registry_v1alpha1_github_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GithubAppConfig); i {
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
		file_buf_alpha_registry_v1alpha1_github_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGithubAppConfigRequest); i {
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
		file_buf_alpha_registry_v1alpha1_github_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGithubAppConfigResponse); i {
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
			RawDescriptor: file_buf_alpha_registry_v1alpha1_github_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_buf_alpha_registry_v1alpha1_github_proto_goTypes,
		DependencyIndexes: file_buf_alpha_registry_v1alpha1_github_proto_depIdxs,
		MessageInfos:      file_buf_alpha_registry_v1alpha1_github_proto_msgTypes,
	}.Build()
	File_buf_alpha_registry_v1alpha1_github_proto = out.File
	file_buf_alpha_registry_v1alpha1_github_proto_rawDesc = nil
	file_buf_alpha_registry_v1alpha1_github_proto_goTypes = nil
	file_buf_alpha_registry_v1alpha1_github_proto_depIdxs = nil
}
