// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: wxService.proto

package services

import (
	common "github.com/geiqin/microkit/protobuf/common"
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

//微信支付参数(发起微信支付需要的参数)
type WxParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TimeStamp string `protobuf:"bytes,2,opt,name=time_stamp,json=timeStamp,proto3" json:"time_stamp,omitempty"`
	NonceStr  string `protobuf:"bytes,3,opt,name=nonce_str,json=nonceStr,proto3" json:"nonce_str,omitempty"`
	Package   string `protobuf:"bytes,4,opt,name=package,proto3" json:"package,omitempty"`
	SignType  string `protobuf:"bytes,5,opt,name=sign_type,json=signType,proto3" json:"sign_type,omitempty"`
	PaySign   string `protobuf:"bytes,6,opt,name=pay_sign,json=paySign,proto3" json:"pay_sign,omitempty"`
}

func (x *WxParam) Reset() {
	*x = WxParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wxService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WxParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WxParam) ProtoMessage() {}

func (x *WxParam) ProtoReflect() protoreflect.Message {
	mi := &file_wxService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WxParam.ProtoReflect.Descriptor instead.
func (*WxParam) Descriptor() ([]byte, []int) {
	return file_wxService_proto_rawDescGZIP(), []int{0}
}

func (x *WxParam) GetTimeStamp() string {
	if x != nil {
		return x.TimeStamp
	}
	return ""
}

func (x *WxParam) GetNonceStr() string {
	if x != nil {
		return x.NonceStr
	}
	return ""
}

func (x *WxParam) GetPackage() string {
	if x != nil {
		return x.Package
	}
	return ""
}

func (x *WxParam) GetSignType() string {
	if x != nil {
		return x.SignType
	}
	return ""
}

func (x *WxParam) GetPaySign() string {
	if x != nil {
		return x.PaySign
	}
	return ""
}

type WxResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *WxParam      `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Error  *common.Error `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info   *common.Info  `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *WxResponse) Reset() {
	*x = WxResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wxService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WxResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WxResponse) ProtoMessage() {}

func (x *WxResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wxService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WxResponse.ProtoReflect.Descriptor instead.
func (*WxResponse) Descriptor() ([]byte, []int) {
	return file_wxService_proto_rawDescGZIP(), []int{1}
}

func (x *WxResponse) GetEntity() *WxParam {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *WxResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *WxResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_wxService_proto protoreflect.FileDescriptor

var file_wxService_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x77, 0x78, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x01,
	0x0a, 0x07, 0x57, 0x78, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x69, 0x6d,
	0x65, 0x5f, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x6f, 0x6e, 0x63,
	0x65, 0x5f, 0x73, 0x74, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x6e,
	0x63, 0x65, 0x53, 0x74, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x73, 0x69, 0x67, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08,
	0x70, 0x61, 0x79, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x70, 0x61, 0x79, 0x53, 0x69, 0x67, 0x6e, 0x22, 0x7e, 0x0a, 0x0a, 0x57, 0x78, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x57, 0x78, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0xfd, 0x01, 0x0a, 0x09, 0x57, 0x78, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x07, 0x4d, 0x69, 0x6e, 0x69, 0x50, 0x61, 0x79,
	0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x57, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x39, 0x0a, 0x05, 0x48, 0x35, 0x50, 0x61, 0x79, 0x12, 0x18, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x57, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a,
	0x06, 0x41, 0x70, 0x70, 0x50, 0x61, 0x79, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x57, 0x78, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x08, 0x4d, 0x69, 0x63,
	0x72, 0x6f, 0x50, 0x61, 0x79, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x57, 0x78, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wxService_proto_rawDescOnce sync.Once
	file_wxService_proto_rawDescData = file_wxService_proto_rawDesc
)

func file_wxService_proto_rawDescGZIP() []byte {
	file_wxService_proto_rawDescOnce.Do(func() {
		file_wxService_proto_rawDescData = protoimpl.X.CompressGZIP(file_wxService_proto_rawDescData)
	})
	return file_wxService_proto_rawDescData
}

var file_wxService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_wxService_proto_goTypes = []interface{}{
	(*WxParam)(nil),        // 0: services.WxParam
	(*WxResponse)(nil),     // 1: services.WxResponse
	(*common.Error)(nil),   // 2: common.Error
	(*common.Info)(nil),    // 3: common.Info
	(*PaymentRequest)(nil), // 4: services.PaymentRequest
}
var file_wxService_proto_depIdxs = []int32{
	0, // 0: services.WxResponse.entity:type_name -> services.WxParam
	2, // 1: services.WxResponse.error:type_name -> common.Error
	3, // 2: services.WxResponse.info:type_name -> common.Info
	4, // 3: services.WxService.MiniPay:input_type -> services.PaymentRequest
	4, // 4: services.WxService.H5Pay:input_type -> services.PaymentRequest
	4, // 5: services.WxService.AppPay:input_type -> services.PaymentRequest
	4, // 6: services.WxService.MicroPay:input_type -> services.PaymentRequest
	1, // 7: services.WxService.MiniPay:output_type -> services.WxResponse
	1, // 8: services.WxService.H5Pay:output_type -> services.WxResponse
	1, // 9: services.WxService.AppPay:output_type -> services.WxResponse
	1, // 10: services.WxService.MicroPay:output_type -> services.WxResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_wxService_proto_init() }
func file_wxService_proto_init() {
	if File_wxService_proto != nil {
		return
	}
	file_payment_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_wxService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WxParam); i {
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
		file_wxService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WxResponse); i {
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
			RawDescriptor: file_wxService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_wxService_proto_goTypes,
		DependencyIndexes: file_wxService_proto_depIdxs,
		MessageInfos:      file_wxService_proto_msgTypes,
	}.Build()
	File_wxService_proto = out.File
	file_wxService_proto_rawDesc = nil
	file_wxService_proto_goTypes = nil
	file_wxService_proto_depIdxs = nil
}