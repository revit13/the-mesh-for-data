// Copyright 2020 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.7.1
// source: data_credential_response.proto

package protobuf

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

type Credentials struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessKey string `protobuf:"bytes,1,opt,name=access_key,json=accessKey,proto3" json:"access_key,omitempty"`
	SecretKey string `protobuf:"bytes,2,opt,name=secret_key,json=secretKey,proto3" json:"secret_key,omitempty"`
	Username  string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Password  string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	ApiKey    string `protobuf:"bytes,5,opt,name=api_key,json=apiKey,proto3" json:"api_key,omitempty"`
}

func (x *Credentials) Reset() {
	*x = Credentials{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_credential_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Credentials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Credentials) ProtoMessage() {}

func (x *Credentials) ProtoReflect() protoreflect.Message {
	mi := &file_data_credential_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Credentials.ProtoReflect.Descriptor instead.
func (*Credentials) Descriptor() ([]byte, []int) {
	return file_data_credential_response_proto_rawDescGZIP(), []int{0}
}

func (x *Credentials) GetAccessKey() string {
	if x != nil {
		return x.AccessKey
	}
	return ""
}

func (x *Credentials) GetSecretKey() string {
	if x != nil {
		return x.SecretKey
	}
	return ""
}

func (x *Credentials) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Credentials) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Credentials) GetApiKey() string {
	if x != nil {
		return x.ApiKey
	}
	return ""
}

type DatasetCredentials struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DatasetId string       `protobuf:"bytes,1,opt,name=dataset_id,json=datasetId,proto3" json:"dataset_id,omitempty"` // identifier of asset - always needed. JSON expected. Interpreted by the Connector, can contain any additional information as part of JSON
	Creds     *Credentials `protobuf:"bytes,2,opt,name=creds,proto3" json:"creds,omitempty"`
}

func (x *DatasetCredentials) Reset() {
	*x = DatasetCredentials{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_credential_response_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatasetCredentials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatasetCredentials) ProtoMessage() {}

func (x *DatasetCredentials) ProtoReflect() protoreflect.Message {
	mi := &file_data_credential_response_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatasetCredentials.ProtoReflect.Descriptor instead.
func (*DatasetCredentials) Descriptor() ([]byte, []int) {
	return file_data_credential_response_proto_rawDescGZIP(), []int{1}
}

func (x *DatasetCredentials) GetDatasetId() string {
	if x != nil {
		return x.DatasetId
	}
	return ""
}

func (x *DatasetCredentials) GetCreds() *Credentials {
	if x != nil {
		return x.Creds
	}
	return nil
}

var File_data_credential_response_proto protoreflect.FileDescriptor

var file_data_credential_response_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0a, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x22, 0x9c, 0x01, 0x0a,
	0x0b, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x1d, 0x0a, 0x0a,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x22, 0x62, 0x0a, 0x12, 0x44,
	0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x49, 0x64,
	0x12, 0x2d, 0x0a, 0x05, 0x63, 0x72, 0x65, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x43, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x52, 0x05, 0x63, 0x72, 0x65, 0x64, 0x73, 0x42,
	0x47, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x74, 0x6d, 0x65, 0x73, 0x68, 0x5a, 0x38,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x62, 0x6d, 0x2f, 0x74,
	0x68, 0x65, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x66, 0x6f, 0x72, 0x2d, 0x64, 0x61, 0x74, 0x61,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_data_credential_response_proto_rawDescOnce sync.Once
	file_data_credential_response_proto_rawDescData = file_data_credential_response_proto_rawDesc
)

func file_data_credential_response_proto_rawDescGZIP() []byte {
	file_data_credential_response_proto_rawDescOnce.Do(func() {
		file_data_credential_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_data_credential_response_proto_rawDescData)
	})
	return file_data_credential_response_proto_rawDescData
}

var file_data_credential_response_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_data_credential_response_proto_goTypes = []interface{}{
	(*Credentials)(nil),        // 0: connectors.Credentials
	(*DatasetCredentials)(nil), // 1: connectors.DatasetCredentials
}
var file_data_credential_response_proto_depIdxs = []int32{
	0, // 0: connectors.DatasetCredentials.creds:type_name -> connectors.Credentials
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_data_credential_response_proto_init() }
func file_data_credential_response_proto_init() {
	if File_data_credential_response_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_data_credential_response_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Credentials); i {
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
		file_data_credential_response_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatasetCredentials); i {
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
			RawDescriptor: file_data_credential_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_data_credential_response_proto_goTypes,
		DependencyIndexes: file_data_credential_response_proto_depIdxs,
		MessageInfos:      file_data_credential_response_proto_msgTypes,
	}.Build()
	File_data_credential_response_proto = out.File
	file_data_credential_response_proto_rawDesc = nil
	file_data_credential_response_proto_goTypes = nil
	file_data_credential_response_proto_depIdxs = nil
}
