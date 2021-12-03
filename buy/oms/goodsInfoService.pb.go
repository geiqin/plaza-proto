// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: goodsInfoService.proto

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

type GoodsInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemId       int64             `protobuf:"varint,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	ItemSn       string            `protobuf:"bytes,2,opt,name=item_sn,json=itemSn,proto3" json:"item_sn,omitempty"`
	Type         int32             `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	Name         string            `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Unit         string            `protobuf:"bytes,5,opt,name=unit,proto3" json:"unit,omitempty"`
	BrandId      int32             `protobuf:"varint,6,opt,name=brand_id,json=brandId,proto3" json:"brand_id,omitempty"`
	TaxonomyId   int64             `protobuf:"varint,7,opt,name=taxonomy_id,json=taxonomyId,proto3" json:"taxonomy_id,omitempty"`
	Quantity     int32             `protobuf:"varint,8,opt,name=quantity,proto3" json:"quantity,omitempty"`
	ThumbId      int64             `protobuf:"varint,9,opt,name=thumb_id,json=thumbId,proto3" json:"thumb_id,omitempty"`
	ThumbUrl     string            `protobuf:"bytes,10,opt,name=thumb_url,json=thumbUrl,proto3" json:"thumb_url,omitempty"`
	Barcode      string            `protobuf:"bytes,11,opt,name=barcode,proto3" json:"barcode,omitempty"`
	Price        float32           `protobuf:"fixed32,12,opt,name=price,proto3" json:"price,omitempty"`
	OriginPrice  float32           `protobuf:"fixed32,13,opt,name=origin_price,json=originPrice,proto3" json:"origin_price,omitempty"`
	CostPrice    float32           `protobuf:"fixed32,14,opt,name=cost_price,json=costPrice,proto3" json:"cost_price,omitempty"`
	Weight       float32           `protobuf:"fixed32,15,opt,name=weight,proto3" json:"weight,omitempty"`
	SkuId        int64             `protobuf:"varint,16,opt,name=sku_id,json=skuId,proto3" json:"sku_id,omitempty"` // 规格商品时使用
	SkuSn        string            `protobuf:"bytes,17,opt,name=sku_sn,json=skuSn,proto3" json:"sku_sn,omitempty"`
	SkuName      string            `protobuf:"bytes,18,opt,name=sku_name,json=skuName,proto3" json:"sku_name,omitempty"`
	SkuSpecDesc  string            `protobuf:"bytes,19,opt,name=sku_spec_desc,json=skuSpecDesc,proto3" json:"sku_spec_desc,omitempty"`
	MinBuy       int32             `protobuf:"varint,20,opt,name=min_buy,json=minBuy,proto3" json:"min_buy,omitempty"`
	BuyQuota     int32             `protobuf:"varint,21,opt,name=buy_quota,json=buyQuota,proto3" json:"buy_quota,omitempty"`
	IsSku        bool              `protobuf:"varint,22,opt,name=is_sku,json=isSku,proto3" json:"is_sku,omitempty"`
	IsVirtual    bool              `protobuf:"varint,23,opt,name=is_virtual,json=isVirtual,proto3" json:"is_virtual,omitempty"`
	IsRight      bool              `protobuf:"varint,24,opt,name=is_right,json=isRight,proto3" json:"is_right,omitempty"`
	IsPresale    bool              `protobuf:"varint,25,opt,name=is_presale,json=isPresale,proto3" json:"is_presale,omitempty"`
	VipLevelId   int32             `protobuf:"varint,26,opt,name=vip_level_id,json=vipLevelId,proto3" json:"vip_level_id,omitempty"`
	VipKeepType  int32             `protobuf:"varint,27,opt,name=vip_keep_type,json=vipKeepType,proto3" json:"vip_keep_type,omitempty"`
	VipKeepValue int32             `protobuf:"varint,28,opt,name=vip_keep_value,json=vipKeepValue,proto3" json:"vip_keep_value,omitempty"`
	ChildCount   int32             `protobuf:"varint,29,opt,name=child_count,json=childCount,proto3" json:"child_count,omitempty"`
	Ext11        int64             `protobuf:"varint,30,opt,name=ext11,proto3" json:"ext11,omitempty"`
	Ext21        int32             `protobuf:"varint,31,opt,name=ext21,proto3" json:"ext21,omitempty"`
	Ext22        int32             `protobuf:"varint,32,opt,name=ext22,proto3" json:"ext22,omitempty"`
	Ext31        string            `protobuf:"bytes,33,opt,name=ext31,proto3" json:"ext31,omitempty"`
	Ext32        string            `protobuf:"bytes,34,opt,name=ext32,proto3" json:"ext32,omitempty"`
	Ext33        string            `protobuf:"bytes,35,opt,name=ext33,proto3" json:"ext33,omitempty"`
	Metas        map[string]string `protobuf:"bytes,36,rep,name=metas,proto3" json:"metas,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GoodsInfo) Reset() {
	*x = GoodsInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goodsInfoService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoodsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodsInfo) ProtoMessage() {}

func (x *GoodsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_goodsInfoService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodsInfo.ProtoReflect.Descriptor instead.
func (*GoodsInfo) Descriptor() ([]byte, []int) {
	return file_goodsInfoService_proto_rawDescGZIP(), []int{0}
}

func (x *GoodsInfo) GetItemId() int64 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *GoodsInfo) GetItemSn() string {
	if x != nil {
		return x.ItemSn
	}
	return ""
}

func (x *GoodsInfo) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *GoodsInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GoodsInfo) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *GoodsInfo) GetBrandId() int32 {
	if x != nil {
		return x.BrandId
	}
	return 0
}

func (x *GoodsInfo) GetTaxonomyId() int64 {
	if x != nil {
		return x.TaxonomyId
	}
	return 0
}

func (x *GoodsInfo) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *GoodsInfo) GetThumbId() int64 {
	if x != nil {
		return x.ThumbId
	}
	return 0
}

func (x *GoodsInfo) GetThumbUrl() string {
	if x != nil {
		return x.ThumbUrl
	}
	return ""
}

func (x *GoodsInfo) GetBarcode() string {
	if x != nil {
		return x.Barcode
	}
	return ""
}

func (x *GoodsInfo) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *GoodsInfo) GetOriginPrice() float32 {
	if x != nil {
		return x.OriginPrice
	}
	return 0
}

func (x *GoodsInfo) GetCostPrice() float32 {
	if x != nil {
		return x.CostPrice
	}
	return 0
}

func (x *GoodsInfo) GetWeight() float32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *GoodsInfo) GetSkuId() int64 {
	if x != nil {
		return x.SkuId
	}
	return 0
}

func (x *GoodsInfo) GetSkuSn() string {
	if x != nil {
		return x.SkuSn
	}
	return ""
}

func (x *GoodsInfo) GetSkuName() string {
	if x != nil {
		return x.SkuName
	}
	return ""
}

func (x *GoodsInfo) GetSkuSpecDesc() string {
	if x != nil {
		return x.SkuSpecDesc
	}
	return ""
}

func (x *GoodsInfo) GetMinBuy() int32 {
	if x != nil {
		return x.MinBuy
	}
	return 0
}

func (x *GoodsInfo) GetBuyQuota() int32 {
	if x != nil {
		return x.BuyQuota
	}
	return 0
}

func (x *GoodsInfo) GetIsSku() bool {
	if x != nil {
		return x.IsSku
	}
	return false
}

func (x *GoodsInfo) GetIsVirtual() bool {
	if x != nil {
		return x.IsVirtual
	}
	return false
}

func (x *GoodsInfo) GetIsRight() bool {
	if x != nil {
		return x.IsRight
	}
	return false
}

func (x *GoodsInfo) GetIsPresale() bool {
	if x != nil {
		return x.IsPresale
	}
	return false
}

func (x *GoodsInfo) GetVipLevelId() int32 {
	if x != nil {
		return x.VipLevelId
	}
	return 0
}

func (x *GoodsInfo) GetVipKeepType() int32 {
	if x != nil {
		return x.VipKeepType
	}
	return 0
}

func (x *GoodsInfo) GetVipKeepValue() int32 {
	if x != nil {
		return x.VipKeepValue
	}
	return 0
}

func (x *GoodsInfo) GetChildCount() int32 {
	if x != nil {
		return x.ChildCount
	}
	return 0
}

func (x *GoodsInfo) GetExt11() int64 {
	if x != nil {
		return x.Ext11
	}
	return 0
}

func (x *GoodsInfo) GetExt21() int32 {
	if x != nil {
		return x.Ext21
	}
	return 0
}

func (x *GoodsInfo) GetExt22() int32 {
	if x != nil {
		return x.Ext22
	}
	return 0
}

func (x *GoodsInfo) GetExt31() string {
	if x != nil {
		return x.Ext31
	}
	return ""
}

func (x *GoodsInfo) GetExt32() string {
	if x != nil {
		return x.Ext32
	}
	return ""
}

func (x *GoodsInfo) GetExt33() string {
	if x != nil {
		return x.Ext33
	}
	return ""
}

func (x *GoodsInfo) GetMetas() map[string]string {
	if x != nil {
		return x.Metas
	}
	return nil
}

var File_goodsInfoService_proto protoreflect.FileDescriptor

var file_goodsInfoService_proto_rawDesc = []byte{
	0x0a, 0x16, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x22, 0xa7, 0x08, 0x0a, 0x09, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65,
	0x6d, 0x5f, 0x73, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d,
	0x53, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x6e,
	0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x12, 0x19,
	0x0a, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x61, 0x78,
	0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x74, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75,
	0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75,
	0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x5f,
	0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x55, 0x72, 0x6c, 0x12, 0x18,
	0x0a, 0x07, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x63, 0x6f, 0x73, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x6b, 0x75, 0x5f,
	0x69, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x6b, 0x75, 0x49, 0x64, 0x12,
	0x15, 0x0a, 0x06, 0x73, 0x6b, 0x75, 0x5f, 0x73, 0x6e, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x73, 0x6b, 0x75, 0x53, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x6b, 0x75, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6b, 0x75, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x22, 0x0a, 0x0d, 0x73, 0x6b, 0x75, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x5f, 0x64, 0x65,
	0x73, 0x63, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x6b, 0x75, 0x53, 0x70, 0x65,
	0x63, 0x44, 0x65, 0x73, 0x63, 0x12, 0x17, 0x0a, 0x07, 0x6d, 0x69, 0x6e, 0x5f, 0x62, 0x75, 0x79,
	0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6d, 0x69, 0x6e, 0x42, 0x75, 0x79, 0x12, 0x1b,
	0x0a, 0x09, 0x62, 0x75, 0x79, 0x5f, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x18, 0x15, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x62, 0x75, 0x79, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x12, 0x15, 0x0a, 0x06, 0x69,
	0x73, 0x5f, 0x73, 0x6b, 0x75, 0x18, 0x16, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x53,
	0x6b, 0x75, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c,
	0x18, 0x17, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61,
	0x6c, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x72, 0x69, 0x67, 0x68, 0x74, 0x18, 0x18, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x52, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x69, 0x73, 0x5f, 0x70, 0x72, 0x65, 0x73, 0x61, 0x6c, 0x65, 0x18, 0x19, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x09, 0x69, 0x73, 0x50, 0x72, 0x65, 0x73, 0x61, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0c, 0x76,
	0x69, 0x70, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x1a, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x76, 0x69, 0x70, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x22, 0x0a,
	0x0d, 0x76, 0x69, 0x70, 0x5f, 0x6b, 0x65, 0x65, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x1b,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x76, 0x69, 0x70, 0x4b, 0x65, 0x65, 0x70, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x24, 0x0a, 0x0e, 0x76, 0x69, 0x70, 0x5f, 0x6b, 0x65, 0x65, 0x70, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x76, 0x69, 0x70, 0x4b, 0x65,
	0x65, 0x70, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x68, 0x69, 0x6c, 0x64,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x68,
	0x69, 0x6c, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x31,
	0x31, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x65, 0x78, 0x74, 0x31, 0x31, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x78, 0x74, 0x32, 0x31, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x65,
	0x78, 0x74, 0x32, 0x31, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x32, 0x32, 0x18, 0x20, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x65, 0x78, 0x74, 0x32, 0x32, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x78,
	0x74, 0x33, 0x31, 0x18, 0x21, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x78, 0x74, 0x33, 0x31,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x33, 0x32, 0x18, 0x22, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x78, 0x74, 0x33, 0x32, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x33, 0x33, 0x18,
	0x23, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x78, 0x74, 0x33, 0x33, 0x12, 0x34, 0x0a, 0x05,
	0x6d, 0x65, 0x74, 0x61, 0x73, 0x18, 0x24, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x4d, 0x65, 0x74, 0x61, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x6d, 0x65, 0x74,
	0x61, 0x73, 0x1a, 0x38, 0x0a, 0x0a, 0x4d, 0x65, 0x74, 0x61, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0c, 0x5a, 0x0a,
	0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_goodsInfoService_proto_rawDescOnce sync.Once
	file_goodsInfoService_proto_rawDescData = file_goodsInfoService_proto_rawDesc
)

func file_goodsInfoService_proto_rawDescGZIP() []byte {
	file_goodsInfoService_proto_rawDescOnce.Do(func() {
		file_goodsInfoService_proto_rawDescData = protoimpl.X.CompressGZIP(file_goodsInfoService_proto_rawDescData)
	})
	return file_goodsInfoService_proto_rawDescData
}

var file_goodsInfoService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_goodsInfoService_proto_goTypes = []interface{}{
	(*GoodsInfo)(nil), // 0: services.GoodsInfo
	nil,               // 1: services.GoodsInfo.MetasEntry
}
var file_goodsInfoService_proto_depIdxs = []int32{
	1, // 0: services.GoodsInfo.metas:type_name -> services.GoodsInfo.MetasEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_goodsInfoService_proto_init() }
func file_goodsInfoService_proto_init() {
	if File_goodsInfoService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_goodsInfoService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoodsInfo); i {
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
			RawDescriptor: file_goodsInfoService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_goodsInfoService_proto_goTypes,
		DependencyIndexes: file_goodsInfoService_proto_depIdxs,
		MessageInfos:      file_goodsInfoService_proto_msgTypes,
	}.Build()
	File_goodsInfoService_proto = out.File
	file_goodsInfoService_proto_rawDesc = nil
	file_goodsInfoService_proto_goTypes = nil
	file_goodsInfoService_proto_depIdxs = nil
}
