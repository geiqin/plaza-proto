// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: fypConfigService.proto

package services

import (
	common "github.com/geiqin/micro-kit/protobuf/common"
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

type FypConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	AppCode     string `protobuf:"bytes,2,opt,name=app_code,json=appCode,proto3" json:"app_code"`
	AppKey      string `protobuf:"bytes,3,opt,name=app_key,json=appKey,proto3" json:"app_key"`
	MonthlyCust string `protobuf:"bytes,4,opt,name=monthly_cust,json=monthlyCust,proto3" json:"monthly_cust"`
	ExpressType int32  `protobuf:"varint,5,opt,name=express_type,json=expressType,proto3" json:"express_type"`
	PayMethod   int32  `protobuf:"varint,6,opt,name=pay_method,json=payMethod,proto3" json:"pay_method"`
	AccessToke  string `protobuf:"bytes,7,opt,name=access_toke,json=accessToke,proto3" json:"access_toke"`
	ExpiredAt   string `protobuf:"bytes,8,opt,name=expired_at,json=expiredAt,proto3" json:"expired_at"`
	AddressId   int64  `protobuf:"varint,9,opt,name=address_id,json=addressId,proto3" json:"address_id"`
	CreatedAt   string `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt   string `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
}

func (x *FypConfig) Reset() {
	*x = FypConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fypConfigService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FypConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FypConfig) ProtoMessage() {}

func (x *FypConfig) ProtoReflect() protoreflect.Message {
	mi := &file_fypConfigService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FypConfig.ProtoReflect.Descriptor instead.
func (*FypConfig) Descriptor() ([]byte, []int) {
	return file_fypConfigService_proto_rawDescGZIP(), []int{0}
}

func (x *FypConfig) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FypConfig) GetAppCode() string {
	if x != nil {
		return x.AppCode
	}
	return ""
}

func (x *FypConfig) GetAppKey() string {
	if x != nil {
		return x.AppKey
	}
	return ""
}

func (x *FypConfig) GetMonthlyCust() string {
	if x != nil {
		return x.MonthlyCust
	}
	return ""
}

func (x *FypConfig) GetExpressType() int32 {
	if x != nil {
		return x.ExpressType
	}
	return 0
}

func (x *FypConfig) GetPayMethod() int32 {
	if x != nil {
		return x.PayMethod
	}
	return 0
}

func (x *FypConfig) GetAccessToke() string {
	if x != nil {
		return x.AccessToke
	}
	return ""
}

func (x *FypConfig) GetExpiredAt() string {
	if x != nil {
		return x.ExpiredAt
	}
	return ""
}

func (x *FypConfig) GetAddressId() int64 {
	if x != nil {
		return x.AddressId
	}
	return 0
}

func (x *FypConfig) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *FypConfig) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type FypConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info   *common.Info  `protobuf:"bytes,1,opt,name=info,proto3" json:"info"`
	Error  *common.Error `protobuf:"bytes,2,opt,name=error,proto3" json:"error"`
	Entity *FypConfig    `protobuf:"bytes,3,opt,name=entity,proto3" json:"entity"`
}

func (x *FypConfigResponse) Reset() {
	*x = FypConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fypConfigService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FypConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FypConfigResponse) ProtoMessage() {}

func (x *FypConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_fypConfigService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FypConfigResponse.ProtoReflect.Descriptor instead.
func (*FypConfigResponse) Descriptor() ([]byte, []int) {
	return file_fypConfigService_proto_rawDescGZIP(), []int{1}
}

func (x *FypConfigResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *FypConfigResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *FypConfigResponse) GetEntity() *FypConfig {
	if x != nil {
		return x.Entity
	}
	return nil
}

var File_fypConfigService_proto protoreflect.FileDescriptor

var file_fypConfigService_proto_rawDesc = []byte{
	0x0a, 0x16, 0x66, 0x79, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd1, 0x02, 0x0a, 0x09, 0x46, 0x79, 0x70, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x70, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x17,
	0x0a, 0x07, 0x61, 0x70, 0x70, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x70, 0x70, 0x4b, 0x65, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x6f, 0x6e, 0x74, 0x68,
	0x6c, 0x79, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d,
	0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x43, 0x75, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x78,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0b, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x61, 0x79, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x70, 0x61, 0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x87, 0x01, 0x0a, 0x11, 0x46, 0x79,
	0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66,
	0x6f, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x2b, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x46, 0x79, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x32, 0x89, 0x01, 0x0a, 0x10, 0x46, 0x79, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12,
	0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x79, 0x70, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x46, 0x79, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x04, 0x53, 0x61, 0x76, 0x65, 0x12, 0x13, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x79, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x79, 0x70, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fypConfigService_proto_rawDescOnce sync.Once
	file_fypConfigService_proto_rawDescData = file_fypConfigService_proto_rawDesc
)

func file_fypConfigService_proto_rawDescGZIP() []byte {
	file_fypConfigService_proto_rawDescOnce.Do(func() {
		file_fypConfigService_proto_rawDescData = protoimpl.X.CompressGZIP(file_fypConfigService_proto_rawDescData)
	})
	return file_fypConfigService_proto_rawDescData
}

var file_fypConfigService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_fypConfigService_proto_goTypes = []interface{}{
	(*FypConfig)(nil),         // 0: services.FypConfig
	(*FypConfigResponse)(nil), // 1: services.FypConfigResponse
	(*common.Info)(nil),       // 2: common.Info
	(*common.Error)(nil),      // 3: common.Error
}
var file_fypConfigService_proto_depIdxs = []int32{
	2, // 0: services.FypConfigResponse.info:type_name -> common.Info
	3, // 1: services.FypConfigResponse.error:type_name -> common.Error
	0, // 2: services.FypConfigResponse.entity:type_name -> services.FypConfig
	0, // 3: services.FypConfigService.Get:input_type -> services.FypConfig
	0, // 4: services.FypConfigService.Save:input_type -> services.FypConfig
	1, // 5: services.FypConfigService.Get:output_type -> services.FypConfigResponse
	1, // 6: services.FypConfigService.Save:output_type -> services.FypConfigResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_fypConfigService_proto_init() }
func file_fypConfigService_proto_init() {
	if File_fypConfigService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fypConfigService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FypConfig); i {
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
		file_fypConfigService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FypConfigResponse); i {
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
			RawDescriptor: file_fypConfigService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_fypConfigService_proto_goTypes,
		DependencyIndexes: file_fypConfigService_proto_depIdxs,
		MessageInfos:      file_fypConfigService_proto_msgTypes,
	}.Build()
	File_fypConfigService_proto = out.File
	file_fypConfigService_proto_rawDesc = nil
	file_fypConfigService_proto_goTypes = nil
	file_fypConfigService_proto_depIdxs = nil
}
