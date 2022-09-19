// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: alipayService.proto

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

//支付宝
type Alipay struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Charge *ChargeInfo `protobuf:"bytes,1,opt,name=charge,proto3" json:"charge"` //订单信息
}

func (x *Alipay) Reset() {
	*x = Alipay{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alipayService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Alipay) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Alipay) ProtoMessage() {}

func (x *Alipay) ProtoReflect() protoreflect.Message {
	mi := &file_alipayService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Alipay.ProtoReflect.Descriptor instead.
func (*Alipay) Descriptor() ([]byte, []int) {
	return file_alipayService_proto_rawDescGZIP(), []int{0}
}

func (x *Alipay) GetCharge() *ChargeInfo {
	if x != nil {
		return x.Charge
	}
	return nil
}

type AlipayData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Alipay           `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Params map[string]string `protobuf:"bytes,2,rep,name=params,proto3" json:"params" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *AlipayData) Reset() {
	*x = AlipayData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alipayService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlipayData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlipayData) ProtoMessage() {}

func (x *AlipayData) ProtoReflect() protoreflect.Message {
	mi := &file_alipayService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlipayData.ProtoReflect.Descriptor instead.
func (*AlipayData) Descriptor() ([]byte, []int) {
	return file_alipayService_proto_rawDescGZIP(), []int{1}
}

func (x *AlipayData) GetEntity() *Alipay {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *AlipayData) GetParams() map[string]string {
	if x != nil {
		return x.Params
	}
	return nil
}

type AlipayResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *AlipayData   `protobuf:"bytes,1,opt,name=data,proto3" json:"data"`
	Error *common.Error `protobuf:"bytes,2,opt,name=error,proto3" json:"error"`
}

func (x *AlipayResponse) Reset() {
	*x = AlipayResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alipayService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlipayResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlipayResponse) ProtoMessage() {}

func (x *AlipayResponse) ProtoReflect() protoreflect.Message {
	mi := &file_alipayService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlipayResponse.ProtoReflect.Descriptor instead.
func (*AlipayResponse) Descriptor() ([]byte, []int) {
	return file_alipayService_proto_rawDescGZIP(), []int{2}
}

func (x *AlipayResponse) GetData() *AlipayData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *AlipayResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_alipayService_proto protoreflect.FileDescriptor

var file_alipayService_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x6c, 0x69, 0x70, 0x61, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a,
	0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x15, 0x62, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x36, 0x0a, 0x06, 0x41, 0x6c, 0x69,
	0x70, 0x61, 0x79, 0x12, 0x2c, 0x0a, 0x06, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43,
	0x68, 0x61, 0x72, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x63, 0x68, 0x61, 0x72, 0x67,
	0x65, 0x22, 0xab, 0x01, 0x0a, 0x0a, 0x41, 0x6c, 0x69, 0x70, 0x61, 0x79, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x28, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x6c, 0x69, 0x70,
	0x61, 0x79, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x38, 0x0a, 0x06, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x6c, 0x69, 0x70, 0x61, 0x79, 0x44, 0x61, 0x74, 0x61,
	0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x5f, 0x0a, 0x0e, 0x41, 0x6c, 0x69, 0x70, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x6c, 0x69, 0x70, 0x61,
	0x79, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x32, 0xb8, 0x01, 0x0a, 0x0d, 0x41, 0x6c, 0x69, 0x70, 0x61, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x36, 0x0a, 0x06, 0x57, 0x61, 0x70, 0x50, 0x61, 0x79, 0x12, 0x10, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x6c, 0x69, 0x70, 0x61, 0x79, 0x1a, 0x18,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x6c, 0x69, 0x70, 0x61, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x07, 0x50, 0x61,
	0x67, 0x65, 0x50, 0x61, 0x79, 0x12, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x41, 0x6c, 0x69, 0x70, 0x61, 0x79, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x41, 0x6c, 0x69, 0x70, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x06, 0x41, 0x70, 0x70, 0x50, 0x61, 0x79, 0x12, 0x10, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x6c, 0x69, 0x70, 0x61, 0x79, 0x1a,
	0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x6c, 0x69, 0x70, 0x61,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2f,
	0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_alipayService_proto_rawDescOnce sync.Once
	file_alipayService_proto_rawDescData = file_alipayService_proto_rawDesc
)

func file_alipayService_proto_rawDescGZIP() []byte {
	file_alipayService_proto_rawDescOnce.Do(func() {
		file_alipayService_proto_rawDescData = protoimpl.X.CompressGZIP(file_alipayService_proto_rawDescData)
	})
	return file_alipayService_proto_rawDescData
}

var file_alipayService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_alipayService_proto_goTypes = []interface{}{
	(*Alipay)(nil),         // 0: services.Alipay
	(*AlipayData)(nil),     // 1: services.AlipayData
	(*AlipayResponse)(nil), // 2: services.AlipayResponse
	nil,                    // 3: services.AlipayData.ParamsEntry
	(*ChargeInfo)(nil),     // 4: services.ChargeInfo
	(*common.Error)(nil),   // 5: common.Error
}
var file_alipayService_proto_depIdxs = []int32{
	4, // 0: services.Alipay.charge:type_name -> services.ChargeInfo
	0, // 1: services.AlipayData.entity:type_name -> services.Alipay
	3, // 2: services.AlipayData.params:type_name -> services.AlipayData.ParamsEntry
	1, // 3: services.AlipayResponse.data:type_name -> services.AlipayData
	5, // 4: services.AlipayResponse.error:type_name -> common.Error
	0, // 5: services.AlipayService.WapPay:input_type -> services.Alipay
	0, // 6: services.AlipayService.PagePay:input_type -> services.Alipay
	0, // 7: services.AlipayService.AppPay:input_type -> services.Alipay
	2, // 8: services.AlipayService.WapPay:output_type -> services.AlipayResponse
	2, // 9: services.AlipayService.PagePay:output_type -> services.AlipayResponse
	2, // 10: services.AlipayService.AppPay:output_type -> services.AlipayResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_alipayService_proto_init() }
func file_alipayService_proto_init() {
	if File_alipayService_proto != nil {
		return
	}
	file_baseInfoService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_alipayService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Alipay); i {
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
		file_alipayService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlipayData); i {
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
		file_alipayService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlipayResponse); i {
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
			RawDescriptor: file_alipayService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_alipayService_proto_goTypes,
		DependencyIndexes: file_alipayService_proto_depIdxs,
		MessageInfos:      file_alipayService_proto_msgTypes,
	}.Build()
	File_alipayService_proto = out.File
	file_alipayService_proto_rawDesc = nil
	file_alipayService_proto_goTypes = nil
	file_alipayService_proto_depIdxs = nil
}
