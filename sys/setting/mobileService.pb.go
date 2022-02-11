// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: mobileService.proto

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

// APP设置
type Mobile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name             string `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
	PackageName      string `protobuf:"bytes,2,opt,name=package_name,json=packageName,proto3" json:"package_name"`
	Sign             string `protobuf:"bytes,3,opt,name=sign,proto3" json:"sign"`
	AppId            string `protobuf:"bytes,4,opt,name=app_id,json=appId,proto3" json:"app_id"`
	AppSecret        string `protobuf:"bytes,5,opt,name=app_secret,json=appSecret,proto3" json:"app_secret"`
	ShowCoverAd      bool   `protobuf:"varint,6,opt,name=show_cover_ad,json=showCoverAd,proto3" json:"show_cover_ad"`
	ShowOtherAd      bool   `protobuf:"varint,7,opt,name=show_other_ad,json=showOtherAd,proto3" json:"show_other_ad"`
	OpenWxPay        bool   `protobuf:"varint,8,opt,name=open_wx_pay,json=openWxPay,proto3" json:"open_wx_pay"`
	OpenAliPay       bool   `protobuf:"varint,9,opt,name=open_ali_pay,json=openAliPay,proto3" json:"open_ali_pay"`
	PushAppKey       string `protobuf:"bytes,10,opt,name=push_app_key,json=pushAppKey,proto3" json:"push_app_key"`
	PushMasterSecret string `protobuf:"bytes,11,opt,name=push_master_secret,json=pushMasterSecret,proto3" json:"push_master_secret"`
}

func (x *Mobile) Reset() {
	*x = Mobile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mobileService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Mobile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mobile) ProtoMessage() {}

func (x *Mobile) ProtoReflect() protoreflect.Message {
	mi := &file_mobileService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mobile.ProtoReflect.Descriptor instead.
func (*Mobile) Descriptor() ([]byte, []int) {
	return file_mobileService_proto_rawDescGZIP(), []int{0}
}

func (x *Mobile) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Mobile) GetPackageName() string {
	if x != nil {
		return x.PackageName
	}
	return ""
}

func (x *Mobile) GetSign() string {
	if x != nil {
		return x.Sign
	}
	return ""
}

func (x *Mobile) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *Mobile) GetAppSecret() string {
	if x != nil {
		return x.AppSecret
	}
	return ""
}

func (x *Mobile) GetShowCoverAd() bool {
	if x != nil {
		return x.ShowCoverAd
	}
	return false
}

func (x *Mobile) GetShowOtherAd() bool {
	if x != nil {
		return x.ShowOtherAd
	}
	return false
}

func (x *Mobile) GetOpenWxPay() bool {
	if x != nil {
		return x.OpenWxPay
	}
	return false
}

func (x *Mobile) GetOpenAliPay() bool {
	if x != nil {
		return x.OpenAliPay
	}
	return false
}

func (x *Mobile) GetPushAppKey() string {
	if x != nil {
		return x.PushAppKey
	}
	return ""
}

func (x *Mobile) GetPushMasterSecret() string {
	if x != nil {
		return x.PushMasterSecret
	}
	return ""
}

//
type MobileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Mobile       `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Error  *common.Error `protobuf:"bytes,4,opt,name=error,proto3" json:"error"`
	Info   *common.Info  `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
}

func (x *MobileResponse) Reset() {
	*x = MobileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mobileService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MobileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MobileResponse) ProtoMessage() {}

func (x *MobileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mobileService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MobileResponse.ProtoReflect.Descriptor instead.
func (*MobileResponse) Descriptor() ([]byte, []int) {
	return file_mobileService_proto_rawDescGZIP(), []int{1}
}

func (x *MobileResponse) GetEntity() *Mobile {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *MobileResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *MobileResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_mobileService_proto protoreflect.FileDescriptor

var file_mobileService_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a,
	0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xe3, 0x02, 0x0a, 0x06, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x67, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x67, 0x6e, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x61, 0x70, 0x70, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x70, 0x70, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x22,
	0x0a, 0x0d, 0x73, 0x68, 0x6f, 0x77, 0x5f, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x73, 0x68, 0x6f, 0x77, 0x43, 0x6f, 0x76, 0x65, 0x72,
	0x41, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x73, 0x68, 0x6f, 0x77, 0x5f, 0x6f, 0x74, 0x68, 0x65, 0x72,
	0x5f, 0x61, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x73, 0x68, 0x6f, 0x77, 0x4f,
	0x74, 0x68, 0x65, 0x72, 0x41, 0x64, 0x12, 0x1e, 0x0a, 0x0b, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x77,
	0x78, 0x5f, 0x70, 0x61, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x6f, 0x70, 0x65,
	0x6e, 0x57, 0x78, 0x50, 0x61, 0x79, 0x12, 0x20, 0x0a, 0x0c, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x61,
	0x6c, 0x69, 0x5f, 0x70, 0x61, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x6f, 0x70,
	0x65, 0x6e, 0x41, 0x6c, 0x69, 0x50, 0x61, 0x79, 0x12, 0x20, 0x0a, 0x0c, 0x70, 0x75, 0x73, 0x68,
	0x5f, 0x61, 0x70, 0x70, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x70, 0x75, 0x73, 0x68, 0x41, 0x70, 0x70, 0x4b, 0x65, 0x79, 0x12, 0x2c, 0x0a, 0x12, 0x70, 0x75,
	0x73, 0x68, 0x5f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x70, 0x75, 0x73, 0x68, 0x4d, 0x61, 0x73, 0x74,
	0x65, 0x72, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x22, 0x81, 0x01, 0x0a, 0x0e, 0x4d, 0x6f, 0x62,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x06, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x52, 0x06, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0x76, 0x0a, 0x0d,
	0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x33, 0x0a,
	0x03, 0x53, 0x65, 0x74, 0x12, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x30, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mobileService_proto_rawDescOnce sync.Once
	file_mobileService_proto_rawDescData = file_mobileService_proto_rawDesc
)

func file_mobileService_proto_rawDescGZIP() []byte {
	file_mobileService_proto_rawDescOnce.Do(func() {
		file_mobileService_proto_rawDescData = protoimpl.X.CompressGZIP(file_mobileService_proto_rawDescData)
	})
	return file_mobileService_proto_rawDescData
}

var file_mobileService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_mobileService_proto_goTypes = []interface{}{
	(*Mobile)(nil),         // 0: services.Mobile
	(*MobileResponse)(nil), // 1: services.MobileResponse
	(*common.Error)(nil),   // 2: common.Error
	(*common.Info)(nil),    // 3: common.Info
	(*common.Empty)(nil),   // 4: common.Empty
}
var file_mobileService_proto_depIdxs = []int32{
	0, // 0: services.MobileResponse.entity:type_name -> services.Mobile
	2, // 1: services.MobileResponse.error:type_name -> common.Error
	3, // 2: services.MobileResponse.info:type_name -> common.Info
	0, // 3: services.MobileService.Set:input_type -> services.Mobile
	4, // 4: services.MobileService.Get:input_type -> common.Empty
	1, // 5: services.MobileService.Set:output_type -> services.MobileResponse
	1, // 6: services.MobileService.Get:output_type -> services.MobileResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_mobileService_proto_init() }
func file_mobileService_proto_init() {
	if File_mobileService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mobileService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Mobile); i {
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
		file_mobileService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MobileResponse); i {
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
			RawDescriptor: file_mobileService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mobileService_proto_goTypes,
		DependencyIndexes: file_mobileService_proto_depIdxs,
		MessageInfos:      file_mobileService_proto_msgTypes,
	}.Build()
	File_mobileService_proto = out.File
	file_mobileService_proto_rawDesc = nil
	file_mobileService_proto_goTypes = nil
	file_mobileService_proto_depIdxs = nil
}
