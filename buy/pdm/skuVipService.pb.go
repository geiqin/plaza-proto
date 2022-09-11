// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: skuVipService.proto

package services

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

type SkuVip struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	SkuId        int64  `protobuf:"varint,2,opt,name=sku_id,json=skuId,proto3" json:"sku_id"`
	MemberRankId int32  `protobuf:"varint,3,opt,name=member_rank_id,json=memberRankId,proto3" json:"member_rank_id"`
	KeepType     string `protobuf:"bytes,4,opt,name=keep_type,json=keepType,proto3" json:"keep_type"`
	KeepValue    int32  `protobuf:"varint,5,opt,name=keep_value,json=keepValue,proto3" json:"keep_value"`
	CreatedAt    string `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt    string `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
}

func (x *SkuVip) Reset() {
	*x = SkuVip{}
	if protoimpl.UnsafeEnabled {
		mi := &file_skuVipService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SkuVip) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SkuVip) ProtoMessage() {}

func (x *SkuVip) ProtoReflect() protoreflect.Message {
	mi := &file_skuVipService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SkuVip.ProtoReflect.Descriptor instead.
func (*SkuVip) Descriptor() ([]byte, []int) {
	return file_skuVipService_proto_rawDescGZIP(), []int{0}
}

func (x *SkuVip) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SkuVip) GetSkuId() int64 {
	if x != nil {
		return x.SkuId
	}
	return 0
}

func (x *SkuVip) GetMemberRankId() int32 {
	if x != nil {
		return x.MemberRankId
	}
	return 0
}

func (x *SkuVip) GetKeepType() string {
	if x != nil {
		return x.KeepType
	}
	return ""
}

func (x *SkuVip) GetKeepValue() int32 {
	if x != nil {
		return x.KeepValue
	}
	return 0
}

func (x *SkuVip) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *SkuVip) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

var File_skuVipService_proto protoreflect.FileDescriptor

var file_skuVipService_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x6b, 0x75, 0x56, 0x69, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22,
	0xcf, 0x01, 0x0a, 0x06, 0x53, 0x6b, 0x75, 0x56, 0x69, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x6b,
	0x75, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x6b, 0x75, 0x49,
	0x64, 0x12, 0x24, 0x0a, 0x0e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x72, 0x61, 0x6e, 0x6b,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x52, 0x61, 0x6e, 0x6b, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6b, 0x65, 0x65, 0x70, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x65, 0x70,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6b, 0x65, 0x65, 0x70, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6b, 0x65, 0x65, 0x70, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_skuVipService_proto_rawDescOnce sync.Once
	file_skuVipService_proto_rawDescData = file_skuVipService_proto_rawDesc
)

func file_skuVipService_proto_rawDescGZIP() []byte {
	file_skuVipService_proto_rawDescOnce.Do(func() {
		file_skuVipService_proto_rawDescData = protoimpl.X.CompressGZIP(file_skuVipService_proto_rawDescData)
	})
	return file_skuVipService_proto_rawDescData
}

var file_skuVipService_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_skuVipService_proto_goTypes = []interface{}{
	(*SkuVip)(nil), // 0: services.SkuVip
}
var file_skuVipService_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_skuVipService_proto_init() }
func file_skuVipService_proto_init() {
	if File_skuVipService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_skuVipService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SkuVip); i {
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
			RawDescriptor: file_skuVipService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_skuVipService_proto_goTypes,
		DependencyIndexes: file_skuVipService_proto_depIdxs,
		MessageInfos:      file_skuVipService_proto_msgTypes,
	}.Build()
	File_skuVipService_proto = out.File
	file_skuVipService_proto_rawDesc = nil
	file_skuVipService_proto_goTypes = nil
	file_skuVipService_proto_depIdxs = nil
}
