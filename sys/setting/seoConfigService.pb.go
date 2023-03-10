// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: seoConfigService.proto

package services

import (
	common "github.com/geiqin/micro-kit/protobuf/common"
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

type SeoConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HomeSeoSiteTitle       string `protobuf:"bytes,1,opt,name=home_seo_site_title,json=homeSeoSiteTitle,proto3" json:"home_seo_site_title"`                   //站点标题，一般不超过80个字符
	HomeSeoSiteKeywords    string `protobuf:"bytes,2,opt,name=home_seo_site_keywords,json=homeSeoSiteKeywords,proto3" json:"home_seo_site_keywords"`          //站点关键字，一般不超过100个字符，多个关键字以半圆角逗号 [ , ] 隔开
	HomeSeoSiteDescription string `protobuf:"bytes,3,opt,name=home_seo_site_description,json=homeSeoSiteDescription,proto3" json:"home_seo_site_description"` //站点描述，一般不超过200个字符
}

func (x *SeoConfig) Reset() {
	*x = SeoConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_seoConfigService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeoConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeoConfig) ProtoMessage() {}

func (x *SeoConfig) ProtoReflect() protoreflect.Message {
	mi := &file_seoConfigService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeoConfig.ProtoReflect.Descriptor instead.
func (*SeoConfig) Descriptor() ([]byte, []int) {
	return file_seoConfigService_proto_rawDescGZIP(), []int{0}
}

func (x *SeoConfig) GetHomeSeoSiteTitle() string {
	if x != nil {
		return x.HomeSeoSiteTitle
	}
	return ""
}

func (x *SeoConfig) GetHomeSeoSiteKeywords() string {
	if x != nil {
		return x.HomeSeoSiteKeywords
	}
	return ""
}

func (x *SeoConfig) GetHomeSeoSiteDescription() string {
	if x != nil {
		return x.HomeSeoSiteDescription
	}
	return ""
}

type SeoConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *SeoConfig    `protobuf:"bytes,1,opt,name=data,proto3" json:"data"`
	Error *common.Error `protobuf:"bytes,2,opt,name=error,proto3" json:"error"`
}

func (x *SeoConfigResponse) Reset() {
	*x = SeoConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_seoConfigService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeoConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeoConfigResponse) ProtoMessage() {}

func (x *SeoConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_seoConfigService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeoConfigResponse.ProtoReflect.Descriptor instead.
func (*SeoConfigResponse) Descriptor() ([]byte, []int) {
	return file_seoConfigService_proto_rawDescGZIP(), []int{1}
}

func (x *SeoConfigResponse) GetData() *SeoConfig {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *SeoConfigResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_seoConfigService_proto protoreflect.FileDescriptor

var file_seoConfigService_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x65, 0x6f, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaa, 0x01, 0x0a, 0x09, 0x53, 0x65, 0x6f, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x2d, 0x0a, 0x13, 0x68, 0x6f, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x6f, 0x5f,
	0x73, 0x69, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x68, 0x6f, 0x6d, 0x65, 0x53, 0x65, 0x6f, 0x53, 0x69, 0x74, 0x65, 0x54, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x33, 0x0a, 0x16, 0x68, 0x6f, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x6f, 0x5f, 0x73,
	0x69, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x13, 0x68, 0x6f, 0x6d, 0x65, 0x53, 0x65, 0x6f, 0x53, 0x69, 0x74, 0x65, 0x4b,
	0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x39, 0x0a, 0x19, 0x68, 0x6f, 0x6d, 0x65, 0x5f,
	0x73, 0x65, 0x6f, 0x5f, 0x73, 0x69, 0x74, 0x65, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x68, 0x6f, 0x6d, 0x65,
	0x53, 0x65, 0x6f, 0x53, 0x69, 0x74, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x61, 0x0a, 0x11, 0x53, 0x65, 0x6f, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x53, 0x65, 0x6f, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x88, 0x01, 0x0a, 0x10, 0x53, 0x65, 0x6f, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x03, 0x53, 0x65,
	0x74, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x65, 0x6f,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x53, 0x65, 0x6f, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x13, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x65, 0x6f, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x65, 0x6f,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_seoConfigService_proto_rawDescOnce sync.Once
	file_seoConfigService_proto_rawDescData = file_seoConfigService_proto_rawDesc
)

func file_seoConfigService_proto_rawDescGZIP() []byte {
	file_seoConfigService_proto_rawDescOnce.Do(func() {
		file_seoConfigService_proto_rawDescData = protoimpl.X.CompressGZIP(file_seoConfigService_proto_rawDescData)
	})
	return file_seoConfigService_proto_rawDescData
}

var file_seoConfigService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_seoConfigService_proto_goTypes = []interface{}{
	(*SeoConfig)(nil),         // 0: services.SeoConfig
	(*SeoConfigResponse)(nil), // 1: services.SeoConfigResponse
	(*common.Error)(nil),      // 2: common.Error
}
var file_seoConfigService_proto_depIdxs = []int32{
	0, // 0: services.SeoConfigResponse.data:type_name -> services.SeoConfig
	2, // 1: services.SeoConfigResponse.error:type_name -> common.Error
	0, // 2: services.SeoConfigService.Set:input_type -> services.SeoConfig
	0, // 3: services.SeoConfigService.Get:input_type -> services.SeoConfig
	1, // 4: services.SeoConfigService.Set:output_type -> services.SeoConfigResponse
	1, // 5: services.SeoConfigService.Get:output_type -> services.SeoConfigResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_seoConfigService_proto_init() }
func file_seoConfigService_proto_init() {
	if File_seoConfigService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_seoConfigService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeoConfig); i {
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
		file_seoConfigService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeoConfigResponse); i {
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
			RawDescriptor: file_seoConfigService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_seoConfigService_proto_goTypes,
		DependencyIndexes: file_seoConfigService_proto_depIdxs,
		MessageInfos:      file_seoConfigService_proto_msgTypes,
	}.Build()
	File_seoConfigService_proto = out.File
	file_seoConfigService_proto_rawDesc = nil
	file_seoConfigService_proto_goTypes = nil
	file_seoConfigService_proto_depIdxs = nil
}