// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: itemStatsService.proto

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

type ItemStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	GoodsNum    int32  `protobuf:"varint,2,opt,name=goods_num,json=goodsNum,proto3" json:"goods_num"`
	ListedNum   int32  `protobuf:"varint,3,opt,name=listed_num,json=listedNum,proto3" json:"listed_num"`
	UnlistedNum int32  `protobuf:"varint,4,opt,name=unlisted_num,json=unlistedNum,proto3" json:"unlisted_num"`
	SelloutNum  int32  `protobuf:"varint,5,opt,name=sellout_num,json=selloutNum,proto3" json:"sellout_num"`
	InvalidNum  int32  `protobuf:"varint,6,opt,name=invalid_num,json=invalidNum,proto3" json:"invalid_num"`
	LastStatsAt string `protobuf:"bytes,7,opt,name=last_stats_at,json=lastStatsAt,proto3" json:"last_stats_at"`
	CreatedAt   string `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt   string `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
}

func (x *ItemStats) Reset() {
	*x = ItemStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_itemStatsService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemStats) ProtoMessage() {}

func (x *ItemStats) ProtoReflect() protoreflect.Message {
	mi := &file_itemStatsService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemStats.ProtoReflect.Descriptor instead.
func (*ItemStats) Descriptor() ([]byte, []int) {
	return file_itemStatsService_proto_rawDescGZIP(), []int{0}
}

func (x *ItemStats) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ItemStats) GetGoodsNum() int32 {
	if x != nil {
		return x.GoodsNum
	}
	return 0
}

func (x *ItemStats) GetListedNum() int32 {
	if x != nil {
		return x.ListedNum
	}
	return 0
}

func (x *ItemStats) GetUnlistedNum() int32 {
	if x != nil {
		return x.UnlistedNum
	}
	return 0
}

func (x *ItemStats) GetSelloutNum() int32 {
	if x != nil {
		return x.SelloutNum
	}
	return 0
}

func (x *ItemStats) GetInvalidNum() int32 {
	if x != nil {
		return x.InvalidNum
	}
	return 0
}

func (x *ItemStats) GetLastStatsAt() string {
	if x != nil {
		return x.LastStatsAt
	}
	return ""
}

func (x *ItemStats) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *ItemStats) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type ItemStatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *ItemStats    `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*ItemStats  `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Error  *common.Error `protobuf:"bytes,4,opt,name=error,proto3" json:"error"`
	Info   *common.Info  `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
}

func (x *ItemStatsResponse) Reset() {
	*x = ItemStatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_itemStatsService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemStatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemStatsResponse) ProtoMessage() {}

func (x *ItemStatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_itemStatsService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemStatsResponse.ProtoReflect.Descriptor instead.
func (*ItemStatsResponse) Descriptor() ([]byte, []int) {
	return file_itemStatsService_proto_rawDescGZIP(), []int{1}
}

func (x *ItemStatsResponse) GetEntity() *ItemStats {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *ItemStatsResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *ItemStatsResponse) GetItems() []*ItemStats {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ItemStatsResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *ItemStatsResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_itemStatsService_proto protoreflect.FileDescriptor

var file_itemStatsService_proto_rawDesc = []byte{
	0x0a, 0x16, 0x69, 0x74, 0x65, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9e, 0x02, 0x0a, 0x09, 0x49, 0x74, 0x65, 0x6d, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x5f, 0x6e, 0x75, 0x6d,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x4e, 0x75, 0x6d,
	0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x64, 0x4e, 0x75, 0x6d, 0x12,
	0x21, 0x0a, 0x0c, 0x75, 0x6e, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x75, 0x6e, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x64, 0x4e,
	0x75, 0x6d, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x6c, 0x6c, 0x6f, 0x75, 0x74, 0x5f, 0x6e, 0x75,
	0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x65, 0x6c, 0x6c, 0x6f, 0x75, 0x74,
	0x4e, 0x75, 0x6d, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x6e,
	0x75, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x4e, 0x75, 0x6d, 0x12, 0x22, 0x0a, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x61, 0x73,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xd7, 0x01, 0x0a, 0x11, 0x49, 0x74, 0x65, 0x6d, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x06,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x29,
	0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20,
	0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f,
	0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_itemStatsService_proto_rawDescOnce sync.Once
	file_itemStatsService_proto_rawDescData = file_itemStatsService_proto_rawDesc
)

func file_itemStatsService_proto_rawDescGZIP() []byte {
	file_itemStatsService_proto_rawDescOnce.Do(func() {
		file_itemStatsService_proto_rawDescData = protoimpl.X.CompressGZIP(file_itemStatsService_proto_rawDescData)
	})
	return file_itemStatsService_proto_rawDescData
}

var file_itemStatsService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_itemStatsService_proto_goTypes = []interface{}{
	(*ItemStats)(nil),         // 0: services.ItemStats
	(*ItemStatsResponse)(nil), // 1: services.ItemStatsResponse
	(*common.Pager)(nil),      // 2: common.Pager
	(*common.Error)(nil),      // 3: common.Error
	(*common.Info)(nil),       // 4: common.Info
}
var file_itemStatsService_proto_depIdxs = []int32{
	0, // 0: services.ItemStatsResponse.entity:type_name -> services.ItemStats
	2, // 1: services.ItemStatsResponse.pager:type_name -> common.Pager
	0, // 2: services.ItemStatsResponse.items:type_name -> services.ItemStats
	3, // 3: services.ItemStatsResponse.error:type_name -> common.Error
	4, // 4: services.ItemStatsResponse.info:type_name -> common.Info
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_itemStatsService_proto_init() }
func file_itemStatsService_proto_init() {
	if File_itemStatsService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_itemStatsService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemStats); i {
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
		file_itemStatsService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemStatsResponse); i {
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
			RawDescriptor: file_itemStatsService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_itemStatsService_proto_goTypes,
		DependencyIndexes: file_itemStatsService_proto_depIdxs,
		MessageInfos:      file_itemStatsService_proto_msgTypes,
	}.Build()
	File_itemStatsService_proto = out.File
	file_itemStatsService_proto_rawDesc = nil
	file_itemStatsService_proto_goTypes = nil
	file_itemStatsService_proto_depIdxs = nil
}
