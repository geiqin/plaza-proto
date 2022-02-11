// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: withdrawCfgService.proto

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

type WithdrawCfg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	MinMoney         float32 `protobuf:"fixed32,2,opt,name=min_money,json=minMoney,proto3" json:"min_money"`                           //单笔最小提现金额
	MaxMoney         float32 `protobuf:"fixed32,3,opt,name=max_money,json=maxMoney,proto3" json:"max_money"`                           //单笔最大提现金额
	ServiceMoneyRate float32 `protobuf:"fixed32,4,opt,name=service_money_rate,json=serviceMoneyRate,proto3" json:"service_money_rate"` //提现手续费比例
	MinServiceMoney  float32 `protobuf:"fixed32,5,opt,name=min_service_money,json=minServiceMoney,proto3" json:"min_service_money"`    //单笔最低收取的提现手续费
	CreatedAt        string  `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt        string  `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
}

func (x *WithdrawCfg) Reset() {
	*x = WithdrawCfg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_withdrawCfgService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WithdrawCfg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WithdrawCfg) ProtoMessage() {}

func (x *WithdrawCfg) ProtoReflect() protoreflect.Message {
	mi := &file_withdrawCfgService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WithdrawCfg.ProtoReflect.Descriptor instead.
func (*WithdrawCfg) Descriptor() ([]byte, []int) {
	return file_withdrawCfgService_proto_rawDescGZIP(), []int{0}
}

func (x *WithdrawCfg) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *WithdrawCfg) GetMinMoney() float32 {
	if x != nil {
		return x.MinMoney
	}
	return 0
}

func (x *WithdrawCfg) GetMaxMoney() float32 {
	if x != nil {
		return x.MaxMoney
	}
	return 0
}

func (x *WithdrawCfg) GetServiceMoneyRate() float32 {
	if x != nil {
		return x.ServiceMoneyRate
	}
	return 0
}

func (x *WithdrawCfg) GetMinServiceMoney() float32 {
	if x != nil {
		return x.MinServiceMoney
	}
	return 0
}

func (x *WithdrawCfg) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *WithdrawCfg) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type WithdrawCfgResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *WithdrawCfg   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager  `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*WithdrawCfg `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Error  *common.Error  `protobuf:"bytes,4,opt,name=error,proto3" json:"error"`
	Info   *common.Info   `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
}

func (x *WithdrawCfgResponse) Reset() {
	*x = WithdrawCfgResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_withdrawCfgService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WithdrawCfgResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WithdrawCfgResponse) ProtoMessage() {}

func (x *WithdrawCfgResponse) ProtoReflect() protoreflect.Message {
	mi := &file_withdrawCfgService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WithdrawCfgResponse.ProtoReflect.Descriptor instead.
func (*WithdrawCfgResponse) Descriptor() ([]byte, []int) {
	return file_withdrawCfgService_proto_rawDescGZIP(), []int{1}
}

func (x *WithdrawCfgResponse) GetEntity() *WithdrawCfg {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *WithdrawCfgResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *WithdrawCfgResponse) GetItems() []*WithdrawCfg {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *WithdrawCfgResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *WithdrawCfgResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_withdrawCfgService_proto protoreflect.FileDescriptor

var file_withdrawCfgService_proto_rawDesc = []byte{
	0x0a, 0x18, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x43, 0x66, 0x67, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xef, 0x01, 0x0a, 0x0b, 0x57, 0x69, 0x74, 0x68,
	0x64, 0x72, 0x61, 0x77, 0x43, 0x66, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x69, 0x6e, 0x5f, 0x6d,
	0x6f, 0x6e, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x6d, 0x69, 0x6e, 0x4d,
	0x6f, 0x6e, 0x65, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x6d, 0x6f, 0x6e, 0x65,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x4d, 0x6f, 0x6e, 0x65,
	0x79, 0x12, 0x2c, 0x0a, 0x12, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x6f, 0x6e,
	0x65, 0x79, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x10, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x52, 0x61, 0x74, 0x65, 0x12,
	0x2a, 0x0a, 0x11, 0x6d, 0x69, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6d,
	0x6f, 0x6e, 0x65, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0f, 0x6d, 0x69, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xdd, 0x01, 0x0a, 0x13, 0x57, 0x69,
	0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x43, 0x66, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2d, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x57, 0x69, 0x74,
	0x68, 0x64, 0x72, 0x61, 0x77, 0x43, 0x66, 0x67, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05,
	0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x43, 0x66, 0x67, 0x52, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0x8a, 0x01, 0x0a, 0x12, 0x57, 0x69,
	0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x43, 0x66, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x35, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x43, 0x66, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x03, 0x53, 0x65, 0x74, 0x12, 0x15,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72,
	0x61, 0x77, 0x43, 0x66, 0x67, 0x1a, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x43, 0x66, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_withdrawCfgService_proto_rawDescOnce sync.Once
	file_withdrawCfgService_proto_rawDescData = file_withdrawCfgService_proto_rawDesc
)

func file_withdrawCfgService_proto_rawDescGZIP() []byte {
	file_withdrawCfgService_proto_rawDescOnce.Do(func() {
		file_withdrawCfgService_proto_rawDescData = protoimpl.X.CompressGZIP(file_withdrawCfgService_proto_rawDescData)
	})
	return file_withdrawCfgService_proto_rawDescData
}

var file_withdrawCfgService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_withdrawCfgService_proto_goTypes = []interface{}{
	(*WithdrawCfg)(nil),         // 0: services.WithdrawCfg
	(*WithdrawCfgResponse)(nil), // 1: services.WithdrawCfgResponse
	(*common.Pager)(nil),        // 2: common.Pager
	(*common.Error)(nil),        // 3: common.Error
	(*common.Info)(nil),         // 4: common.Info
	(*common.Empty)(nil),        // 5: common.Empty
}
var file_withdrawCfgService_proto_depIdxs = []int32{
	0, // 0: services.WithdrawCfgResponse.entity:type_name -> services.WithdrawCfg
	2, // 1: services.WithdrawCfgResponse.pager:type_name -> common.Pager
	0, // 2: services.WithdrawCfgResponse.items:type_name -> services.WithdrawCfg
	3, // 3: services.WithdrawCfgResponse.error:type_name -> common.Error
	4, // 4: services.WithdrawCfgResponse.info:type_name -> common.Info
	5, // 5: services.WithdrawCfgService.Get:input_type -> common.Empty
	0, // 6: services.WithdrawCfgService.Set:input_type -> services.WithdrawCfg
	1, // 7: services.WithdrawCfgService.Get:output_type -> services.WithdrawCfgResponse
	1, // 8: services.WithdrawCfgService.Set:output_type -> services.WithdrawCfgResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_withdrawCfgService_proto_init() }
func file_withdrawCfgService_proto_init() {
	if File_withdrawCfgService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_withdrawCfgService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WithdrawCfg); i {
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
		file_withdrawCfgService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WithdrawCfgResponse); i {
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
			RawDescriptor: file_withdrawCfgService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_withdrawCfgService_proto_goTypes,
		DependencyIndexes: file_withdrawCfgService_proto_depIdxs,
		MessageInfos:      file_withdrawCfgService_proto_msgTypes,
	}.Build()
	File_withdrawCfgService_proto = out.File
	file_withdrawCfgService_proto_rawDesc = nil
	file_withdrawCfgService_proto_goTypes = nil
	file_withdrawCfgService_proto_depIdxs = nil
}