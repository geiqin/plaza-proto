// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: itemCounterService.proto

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

type ItemCounter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ItemId       int64 `protobuf:"varint,2,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	InitSoldNum  int32 `protobuf:"varint,3,opt,name=init_sold_num,json=initSoldNum,proto3" json:"init_sold_num,omitempty"`
	ReviewNum    int32 `protobuf:"varint,4,opt,name=review_num,json=reviewNum,proto3" json:"review_num,omitempty"`
	ViewCount    int32 `protobuf:"varint,5,opt,name=view_count,json=viewCount,proto3" json:"view_count,omitempty"`
	BuyCount     int32 `protobuf:"varint,6,opt,name=buy_count,json=buyCount,proto3" json:"buy_count,omitempty"`
	WeekBuyCount int32 `protobuf:"varint,7,opt,name=week_buy_count,json=weekBuyCount,proto3" json:"week_buy_count,omitempty"`
}

func (x *ItemCounter) Reset() {
	*x = ItemCounter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_itemCounterService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemCounter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemCounter) ProtoMessage() {}

func (x *ItemCounter) ProtoReflect() protoreflect.Message {
	mi := &file_itemCounterService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemCounter.ProtoReflect.Descriptor instead.
func (*ItemCounter) Descriptor() ([]byte, []int) {
	return file_itemCounterService_proto_rawDescGZIP(), []int{0}
}

func (x *ItemCounter) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ItemCounter) GetItemId() int64 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *ItemCounter) GetInitSoldNum() int32 {
	if x != nil {
		return x.InitSoldNum
	}
	return 0
}

func (x *ItemCounter) GetReviewNum() int32 {
	if x != nil {
		return x.ReviewNum
	}
	return 0
}

func (x *ItemCounter) GetViewCount() int32 {
	if x != nil {
		return x.ViewCount
	}
	return 0
}

func (x *ItemCounter) GetBuyCount() int32 {
	if x != nil {
		return x.BuyCount
	}
	return 0
}

func (x *ItemCounter) GetWeekBuyCount() int32 {
	if x != nil {
		return x.WeekBuyCount
	}
	return 0
}

type ItemCounterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *ItemCounter   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager  *common.Pager  `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items  []*ItemCounter `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error  *common.Error  `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info   *common.Info   `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *ItemCounterResponse) Reset() {
	*x = ItemCounterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_itemCounterService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemCounterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemCounterResponse) ProtoMessage() {}

func (x *ItemCounterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_itemCounterService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemCounterResponse.ProtoReflect.Descriptor instead.
func (*ItemCounterResponse) Descriptor() ([]byte, []int) {
	return file_itemCounterService_proto_rawDescGZIP(), []int{1}
}

func (x *ItemCounterResponse) GetEntity() *ItemCounter {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *ItemCounterResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *ItemCounterResponse) GetItems() []*ItemCounter {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ItemCounterResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *ItemCounterResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_itemCounterService_proto protoreflect.FileDescriptor

var file_itemCounterService_proto_rawDesc = []byte{
	0x0a, 0x18, 0x69, 0x74, 0x65, 0x6d, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdb, 0x01, 0x0a, 0x0b, 0x49, 0x74, 0x65, 0x6d,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64,
	0x12, 0x22, 0x0a, 0x0d, 0x69, 0x6e, 0x69, 0x74, 0x5f, 0x73, 0x6f, 0x6c, 0x64, 0x5f, 0x6e, 0x75,
	0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x69, 0x6e, 0x69, 0x74, 0x53, 0x6f, 0x6c,
	0x64, 0x4e, 0x75, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x6e,
	0x75, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x4e, 0x75, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x75, 0x79, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x62, 0x75, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x24, 0x0a, 0x0e, 0x77, 0x65, 0x65, 0x6b, 0x5f, 0x62, 0x75, 0x79, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x77, 0x65, 0x65, 0x6b, 0x42, 0x75, 0x79,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xdd, 0x01, 0x0a, 0x13, 0x49, 0x74, 0x65, 0x6d, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a,
	0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x65, 0x72, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05,
	0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65,
	0x72, 0x12, 0x2b, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x23,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0x91, 0x01, 0x0a, 0x12, 0x49, 0x74, 0x65, 0x6d, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x03,
	0x47, 0x65, 0x74, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49,
	0x74, 0x65, 0x6d, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x1a, 0x1d, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x06, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x42,
	0x61, 0x73, 0x65, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_itemCounterService_proto_rawDescOnce sync.Once
	file_itemCounterService_proto_rawDescData = file_itemCounterService_proto_rawDesc
)

func file_itemCounterService_proto_rawDescGZIP() []byte {
	file_itemCounterService_proto_rawDescOnce.Do(func() {
		file_itemCounterService_proto_rawDescData = protoimpl.X.CompressGZIP(file_itemCounterService_proto_rawDescData)
	})
	return file_itemCounterService_proto_rawDescData
}

var file_itemCounterService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_itemCounterService_proto_goTypes = []interface{}{
	(*ItemCounter)(nil),         // 0: services.ItemCounter
	(*ItemCounterResponse)(nil), // 1: services.ItemCounterResponse
	(*common.Pager)(nil),        // 2: common.Pager
	(*common.Error)(nil),        // 3: common.Error
	(*common.Info)(nil),         // 4: common.Info
	(*common.BaseWhere)(nil),    // 5: common.BaseWhere
}
var file_itemCounterService_proto_depIdxs = []int32{
	0, // 0: services.ItemCounterResponse.entity:type_name -> services.ItemCounter
	2, // 1: services.ItemCounterResponse.pager:type_name -> common.Pager
	0, // 2: services.ItemCounterResponse.items:type_name -> services.ItemCounter
	3, // 3: services.ItemCounterResponse.error:type_name -> common.Error
	4, // 4: services.ItemCounterResponse.info:type_name -> common.Info
	0, // 5: services.ItemCounterService.Get:input_type -> services.ItemCounter
	5, // 6: services.ItemCounterService.Search:input_type -> common.BaseWhere
	1, // 7: services.ItemCounterService.Get:output_type -> services.ItemCounterResponse
	1, // 8: services.ItemCounterService.Search:output_type -> services.ItemCounterResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_itemCounterService_proto_init() }
func file_itemCounterService_proto_init() {
	if File_itemCounterService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_itemCounterService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemCounter); i {
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
		file_itemCounterService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemCounterResponse); i {
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
			RawDescriptor: file_itemCounterService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_itemCounterService_proto_goTypes,
		DependencyIndexes: file_itemCounterService_proto_depIdxs,
		MessageInfos:      file_itemCounterService_proto_msgTypes,
	}.Build()
	File_itemCounterService_proto = out.File
	file_itemCounterService_proto_rawDesc = nil
	file_itemCounterService_proto_goTypes = nil
	file_itemCounterService_proto_depIdxs = nil
}
