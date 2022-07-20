// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: promotionVariantService.proto

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

//营销活动变体
type PromotionVariant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                              //活动ID
	PromotionId   int64    `protobuf:"varint,2,opt,name=promotion_id,json=promotionId,proto3" json:"promotion_id,omitempty"`         //促销ID
	PromotionType string   `protobuf:"bytes,3,opt,name=promotion_type,json=promotionType,proto3" json:"promotion_type,omitempty"`    //促销类型 [冗余]
	Type          string   `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`                                           //变体类型：spu,sku
	SpuId         int64    `protobuf:"varint,5,opt,name=spu_id,json=spuId,proto3" json:"spu_id,omitempty"`                           //商品ID [冗余]
	SkuId         int64    `protobuf:"varint,6,opt,name=sku_id,json=skuId,proto3" json:"sku_id,omitempty"`                           //货品ID
	Stock         int32    `protobuf:"varint,7,opt,name=stock,proto3" json:"stock,omitempty"`                                        //促销库存
	Method        int64    `protobuf:"varint,8,opt,name=method,proto3" json:"method,omitempty"`                                      //优惠方式：1 折扣 2减价,3固定金额
	DiscountRate  float32  `protobuf:"fixed32,9,opt,name=discount_rate,json=discountRate,proto3" json:"discount_rate,omitempty"`     //优惠折扣
	DiscountMoney float32  `protobuf:"fixed32,10,opt,name=discount_money,json=discountMoney,proto3" json:"discount_money,omitempty"` //优惠金额(立减金额/固定金额)
	QuantityLimit int32    `protobuf:"varint,11,opt,name=quantity_limit,json=quantityLimit,proto3" json:"quantity_limit,omitempty"`  //购买数量限制
	Enabled       bool     `protobuf:"varint,12,opt,name=enabled,proto3" json:"enabled,omitempty"`                                   //已启用
	CreatedAt     string   `protobuf:"bytes,13,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`               //创建时间
	UpdatedAt     string   `protobuf:"bytes,14,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`               //修改时间
	Spu           *SpuInfo `protobuf:"bytes,15,opt,name=spu,proto3" json:"spu,omitempty"`
	Sku           *SkuInfo `protobuf:"bytes,16,opt,name=sku,proto3" json:"sku,omitempty"`
}

func (x *PromotionVariant) Reset() {
	*x = PromotionVariant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_promotionVariantService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PromotionVariant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PromotionVariant) ProtoMessage() {}

func (x *PromotionVariant) ProtoReflect() protoreflect.Message {
	mi := &file_promotionVariantService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PromotionVariant.ProtoReflect.Descriptor instead.
func (*PromotionVariant) Descriptor() ([]byte, []int) {
	return file_promotionVariantService_proto_rawDescGZIP(), []int{0}
}

func (x *PromotionVariant) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PromotionVariant) GetPromotionId() int64 {
	if x != nil {
		return x.PromotionId
	}
	return 0
}

func (x *PromotionVariant) GetPromotionType() string {
	if x != nil {
		return x.PromotionType
	}
	return ""
}

func (x *PromotionVariant) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *PromotionVariant) GetSpuId() int64 {
	if x != nil {
		return x.SpuId
	}
	return 0
}

func (x *PromotionVariant) GetSkuId() int64 {
	if x != nil {
		return x.SkuId
	}
	return 0
}

func (x *PromotionVariant) GetStock() int32 {
	if x != nil {
		return x.Stock
	}
	return 0
}

func (x *PromotionVariant) GetMethod() int64 {
	if x != nil {
		return x.Method
	}
	return 0
}

func (x *PromotionVariant) GetDiscountRate() float32 {
	if x != nil {
		return x.DiscountRate
	}
	return 0
}

func (x *PromotionVariant) GetDiscountMoney() float32 {
	if x != nil {
		return x.DiscountMoney
	}
	return 0
}

func (x *PromotionVariant) GetQuantityLimit() int32 {
	if x != nil {
		return x.QuantityLimit
	}
	return 0
}

func (x *PromotionVariant) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *PromotionVariant) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *PromotionVariant) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *PromotionVariant) GetSpu() *SpuInfo {
	if x != nil {
		return x.Spu
	}
	return nil
}

func (x *PromotionVariant) GetSku() *SkuInfo {
	if x != nil {
		return x.Sku
	}
	return nil
}

var File_promotionVariantService_proto protoreflect.FileDescriptor

var file_promotionVariantService_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x72, 0x69, 0x61,
	0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x14, 0x73, 0x6b, 0x75, 0x49, 0x6e,
	0x66, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xf1, 0x03, 0x0a, 0x10, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x72,
	0x69, 0x61, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x6d,
	0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x6d, 0x6f,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x70, 0x75, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x73, 0x70, 0x75, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x6b, 0x75,
	0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x6b, 0x75, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x23,
	0x0a, 0x0d, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x61, 0x74, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f,
	0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x71, 0x75,
	0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0d, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x23, 0x0a, 0x03, 0x73, 0x70, 0x75,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x53, 0x70, 0x75, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x03, 0x73, 0x70, 0x75, 0x12, 0x23,
	0x0a, 0x03, 0x73, 0x6b, 0x75, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x6b, 0x75, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x03,
	0x73, 0x6b, 0x75, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_promotionVariantService_proto_rawDescOnce sync.Once
	file_promotionVariantService_proto_rawDescData = file_promotionVariantService_proto_rawDesc
)

func file_promotionVariantService_proto_rawDescGZIP() []byte {
	file_promotionVariantService_proto_rawDescOnce.Do(func() {
		file_promotionVariantService_proto_rawDescData = protoimpl.X.CompressGZIP(file_promotionVariantService_proto_rawDescData)
	})
	return file_promotionVariantService_proto_rawDescData
}

var file_promotionVariantService_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_promotionVariantService_proto_goTypes = []interface{}{
	(*PromotionVariant)(nil), // 0: services.PromotionVariant
	(*SpuInfo)(nil),          // 1: services.SpuInfo
	(*SkuInfo)(nil),          // 2: services.SkuInfo
}
var file_promotionVariantService_proto_depIdxs = []int32{
	1, // 0: services.PromotionVariant.spu:type_name -> services.SpuInfo
	2, // 1: services.PromotionVariant.sku:type_name -> services.SkuInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_promotionVariantService_proto_init() }
func file_promotionVariantService_proto_init() {
	if File_promotionVariantService_proto != nil {
		return
	}
	file_skuInfoService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_promotionVariantService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PromotionVariant); i {
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
			RawDescriptor: file_promotionVariantService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_promotionVariantService_proto_goTypes,
		DependencyIndexes: file_promotionVariantService_proto_depIdxs,
		MessageInfos:      file_promotionVariantService_proto_msgTypes,
	}.Build()
	File_promotionVariantService_proto = out.File
	file_promotionVariantService_proto_rawDesc = nil
	file_promotionVariantService_proto_goTypes = nil
	file_promotionVariantService_proto_depIdxs = nil
}
