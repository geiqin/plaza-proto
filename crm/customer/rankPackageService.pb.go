// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: rankPackageService.proto

package services

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

type RankPackage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int32                    `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	RankId       int32                    `protobuf:"varint,2,opt,name=rank_id,json=rankId,proto3" json:"rank_id"`
	IsSendPoint  bool                     `protobuf:"varint,3,opt,name=is_send_point,json=isSendPoint,proto3" json:"is_send_point"`
	IsSendCoupon bool                     `protobuf:"varint,4,opt,name=is_send_coupon,json=isSendCoupon,proto3" json:"is_send_coupon"`
	IsSendGift   bool                     `protobuf:"varint,5,opt,name=is_send_gift,json=isSendGift,proto3" json:"is_send_gift"`
	SendPoints   int32                    `protobuf:"varint,6,opt,name=send_points,json=sendPoints,proto3" json:"send_points"`
	SendCoupons  []*RankPackageSendCoupon `protobuf:"bytes,7,rep,name=send_coupons,json=sendCoupons,proto3" json:"send_coupons"`
	SendGifts    []*RankPackageSendGift   `protobuf:"bytes,8,rep,name=send_gifts,json=sendGifts,proto3" json:"send_gifts"`
}

func (x *RankPackage) Reset() {
	*x = RankPackage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rankPackageService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RankPackage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RankPackage) ProtoMessage() {}

func (x *RankPackage) ProtoReflect() protoreflect.Message {
	mi := &file_rankPackageService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RankPackage.ProtoReflect.Descriptor instead.
func (*RankPackage) Descriptor() ([]byte, []int) {
	return file_rankPackageService_proto_rawDescGZIP(), []int{0}
}

func (x *RankPackage) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RankPackage) GetRankId() int32 {
	if x != nil {
		return x.RankId
	}
	return 0
}

func (x *RankPackage) GetIsSendPoint() bool {
	if x != nil {
		return x.IsSendPoint
	}
	return false
}

func (x *RankPackage) GetIsSendCoupon() bool {
	if x != nil {
		return x.IsSendCoupon
	}
	return false
}

func (x *RankPackage) GetIsSendGift() bool {
	if x != nil {
		return x.IsSendGift
	}
	return false
}

func (x *RankPackage) GetSendPoints() int32 {
	if x != nil {
		return x.SendPoints
	}
	return 0
}

func (x *RankPackage) GetSendCoupons() []*RankPackageSendCoupon {
	if x != nil {
		return x.SendCoupons
	}
	return nil
}

func (x *RankPackage) GetSendGifts() []*RankPackageSendGift {
	if x != nil {
		return x.SendGifts
	}
	return nil
}

type RankPackageSendCoupon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title"`
	Desc  string `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc"`
	Num   int32  `protobuf:"varint,4,opt,name=num,proto3" json:"num"`
}

func (x *RankPackageSendCoupon) Reset() {
	*x = RankPackageSendCoupon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rankPackageService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RankPackageSendCoupon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RankPackageSendCoupon) ProtoMessage() {}

func (x *RankPackageSendCoupon) ProtoReflect() protoreflect.Message {
	mi := &file_rankPackageService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RankPackageSendCoupon.ProtoReflect.Descriptor instead.
func (*RankPackageSendCoupon) Descriptor() ([]byte, []int) {
	return file_rankPackageService_proto_rawDescGZIP(), []int{0, 0}
}

func (x *RankPackageSendCoupon) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RankPackageSendCoupon) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *RankPackageSendCoupon) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *RankPackageSendCoupon) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

type RankPackageSendGift struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Title    string `protobuf:"bytes,2,opt,name=title,proto3" json:"title"`
	ThumbUrl string `protobuf:"bytes,3,opt,name=thumb_url,json=thumbUrl,proto3" json:"thumb_url"`
	Num      int32  `protobuf:"varint,4,opt,name=num,proto3" json:"num"`
}

func (x *RankPackageSendGift) Reset() {
	*x = RankPackageSendGift{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rankPackageService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RankPackageSendGift) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RankPackageSendGift) ProtoMessage() {}

func (x *RankPackageSendGift) ProtoReflect() protoreflect.Message {
	mi := &file_rankPackageService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RankPackageSendGift.ProtoReflect.Descriptor instead.
func (*RankPackageSendGift) Descriptor() ([]byte, []int) {
	return file_rankPackageService_proto_rawDescGZIP(), []int{0, 1}
}

func (x *RankPackageSendGift) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RankPackageSendGift) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *RankPackageSendGift) GetThumbUrl() string {
	if x != nil {
		return x.ThumbUrl
	}
	return ""
}

func (x *RankPackageSendGift) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

var File_rankPackageService_proto protoreflect.FileDescriptor

var file_rankPackageService_proto_rawDesc = []byte{
	0x0a, 0x18, 0x72, 0x61, 0x6e, 0x6b, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x22, 0x86, 0x04, 0x0a, 0x0b, 0x52, 0x61, 0x6e, 0x6b, 0x50, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x61, 0x6e, 0x6b, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x61, 0x6e, 0x6b, 0x49, 0x64, 0x12, 0x22, 0x0a,
	0x0d, 0x69, 0x73, 0x5f, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x53, 0x65, 0x6e, 0x64, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x12, 0x24, 0x0a, 0x0e, 0x69, 0x73, 0x5f, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x63, 0x6f, 0x75,
	0x70, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x53, 0x65, 0x6e,
	0x64, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x73, 0x65,
	0x6e, 0x64, 0x5f, 0x67, 0x69, 0x66, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69,
	0x73, 0x53, 0x65, 0x6e, 0x64, 0x47, 0x69, 0x66, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x6e,
	0x64, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x73, 0x65, 0x6e, 0x64, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x44, 0x0a, 0x0c, 0x73, 0x65,
	0x6e, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x61, 0x6e, 0x6b,
	0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x63, 0x6f, 0x75,
	0x70, 0x6f, 0x6e, 0x52, 0x0b, 0x73, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x73,
	0x12, 0x3e, 0x0a, 0x0a, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x67, 0x69, 0x66, 0x74, 0x73, 0x18, 0x08,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x52, 0x61, 0x6e, 0x6b, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x73, 0x65, 0x6e, 0x64,
	0x5f, 0x67, 0x69, 0x66, 0x74, 0x52, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x47, 0x69, 0x66, 0x74, 0x73,
	0x1a, 0x59, 0x0a, 0x0b, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x1a, 0x60, 0x0a, 0x09, 0x73,
	0x65, 0x6e, 0x64, 0x5f, 0x67, 0x69, 0x66, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x55, 0x72, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x6e,
	0x75, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x42, 0x0c, 0x5a,
	0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_rankPackageService_proto_rawDescOnce sync.Once
	file_rankPackageService_proto_rawDescData = file_rankPackageService_proto_rawDesc
)

func file_rankPackageService_proto_rawDescGZIP() []byte {
	file_rankPackageService_proto_rawDescOnce.Do(func() {
		file_rankPackageService_proto_rawDescData = protoimpl.X.CompressGZIP(file_rankPackageService_proto_rawDescData)
	})
	return file_rankPackageService_proto_rawDescData
}

var file_rankPackageService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_rankPackageService_proto_goTypes = []interface{}{
	(*RankPackage)(nil),           // 0: services.RankPackage
	(*RankPackageSendCoupon)(nil), // 1: services.RankPackage.send_coupon
	(*RankPackageSendGift)(nil),   // 2: services.RankPackage.send_gift
}
var file_rankPackageService_proto_depIdxs = []int32{
	1, // 0: services.RankPackage.send_coupons:type_name -> services.RankPackage.send_coupon
	2, // 1: services.RankPackage.send_gifts:type_name -> services.RankPackage.send_gift
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_rankPackageService_proto_init() }
func file_rankPackageService_proto_init() {
	if File_rankPackageService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rankPackageService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RankPackage); i {
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
		file_rankPackageService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RankPackageSendCoupon); i {
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
		file_rankPackageService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RankPackageSendGift); i {
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
			RawDescriptor: file_rankPackageService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rankPackageService_proto_goTypes,
		DependencyIndexes: file_rankPackageService_proto_depIdxs,
		MessageInfos:      file_rankPackageService_proto_msgTypes,
	}.Build()
	File_rankPackageService_proto = out.File
	file_rankPackageService_proto_rawDesc = nil
	file_rankPackageService_proto_goTypes = nil
	file_rankPackageService_proto_depIdxs = nil
}
