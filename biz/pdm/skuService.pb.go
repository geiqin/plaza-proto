// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: skuService.proto

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

type Sku struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ItemId      int64   `protobuf:"varint,2,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	SkuSn       string  `protobuf:"bytes,3,opt,name=sku_sn,json=skuSn,proto3" json:"sku_sn,omitempty"`
	Name        string  `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Barcode     string  `protobuf:"bytes,5,opt,name=barcode,proto3" json:"barcode,omitempty"`
	Price       float32 `protobuf:"fixed32,6,opt,name=price,proto3" json:"price,omitempty"`
	OriginPrice float32 `protobuf:"fixed32,7,opt,name=origin_price,json=originPrice,proto3" json:"origin_price,omitempty"`
	CostPrice   float32 `protobuf:"fixed32,8,opt,name=cost_price,json=costPrice,proto3" json:"cost_price,omitempty"`
	Weight      float32 `protobuf:"fixed32,9,opt,name=weight,proto3" json:"weight,omitempty"`
	Quantity    int32   `protobuf:"varint,10,opt,name=quantity,proto3" json:"quantity,omitempty"`
	SoldNum     int32   `protobuf:"varint,11,opt,name=sold_num,json=soldNum,proto3" json:"sold_num,omitempty"`
	InitSoldNum int32   `protobuf:"varint,12,opt,name=init_sold_num,json=initSoldNum,proto3" json:"init_sold_num,omitempty"`
	FrozenNum   int32   `protobuf:"varint,13,opt,name=frozen_num,json=frozenNum,proto3" json:"frozen_num,omitempty"`
	Listed      bool    `protobuf:"varint,14,opt,name=listed,proto3" json:"listed,omitempty"`
	OutSkuNo    string  `protobuf:"bytes,15,opt,name=out_sku_no,json=outSkuNo,proto3" json:"out_sku_no,omitempty"`
	SpecDesc    string  `protobuf:"bytes,16,opt,name=spec_desc,json=specDesc,proto3" json:"spec_desc,omitempty"`
	Defaulted   bool    `protobuf:"varint,17,opt,name=defaulted,proto3" json:"defaulted,omitempty"`
	BuyQuota    int32   `protobuf:"varint,18,opt,name=buy_quota,json=buyQuota,proto3" json:"buy_quota,omitempty"`
	CreatedAt   string  `protobuf:"bytes,19,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   string  `protobuf:"bytes,20,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// @inject_tag: gorm:"many2many:spec_item_indices;"
	SpecValues []*SpecValue `protobuf:"bytes,21,rep,name=spec_values,json=specValues,proto3" json:"spec_values,omitempty"`
}

func (x *Sku) Reset() {
	*x = Sku{}
	if protoimpl.UnsafeEnabled {
		mi := &file_skuService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sku) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sku) ProtoMessage() {}

func (x *Sku) ProtoReflect() protoreflect.Message {
	mi := &file_skuService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sku.ProtoReflect.Descriptor instead.
func (*Sku) Descriptor() ([]byte, []int) {
	return file_skuService_proto_rawDescGZIP(), []int{0}
}

func (x *Sku) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Sku) GetItemId() int64 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *Sku) GetSkuSn() string {
	if x != nil {
		return x.SkuSn
	}
	return ""
}

func (x *Sku) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Sku) GetBarcode() string {
	if x != nil {
		return x.Barcode
	}
	return ""
}

func (x *Sku) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Sku) GetOriginPrice() float32 {
	if x != nil {
		return x.OriginPrice
	}
	return 0
}

func (x *Sku) GetCostPrice() float32 {
	if x != nil {
		return x.CostPrice
	}
	return 0
}

func (x *Sku) GetWeight() float32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *Sku) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *Sku) GetSoldNum() int32 {
	if x != nil {
		return x.SoldNum
	}
	return 0
}

func (x *Sku) GetInitSoldNum() int32 {
	if x != nil {
		return x.InitSoldNum
	}
	return 0
}

func (x *Sku) GetFrozenNum() int32 {
	if x != nil {
		return x.FrozenNum
	}
	return 0
}

func (x *Sku) GetListed() bool {
	if x != nil {
		return x.Listed
	}
	return false
}

func (x *Sku) GetOutSkuNo() string {
	if x != nil {
		return x.OutSkuNo
	}
	return ""
}

func (x *Sku) GetSpecDesc() string {
	if x != nil {
		return x.SpecDesc
	}
	return ""
}

func (x *Sku) GetDefaulted() bool {
	if x != nil {
		return x.Defaulted
	}
	return false
}

func (x *Sku) GetBuyQuota() int32 {
	if x != nil {
		return x.BuyQuota
	}
	return 0
}

func (x *Sku) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Sku) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Sku) GetSpecValues() []*SpecValue {
	if x != nil {
		return x.SpecValues
	}
	return nil
}

type SpecItemIndex struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ItemId      int64 `protobuf:"varint,2,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	SkuId       int64 `protobuf:"varint,3,opt,name=sku_id,json=skuId,proto3" json:"sku_id,omitempty"`
	SpecId      int32 `protobuf:"varint,4,opt,name=spec_id,json=specId,proto3" json:"spec_id,omitempty"`
	SpecValueId int32 `protobuf:"varint,5,opt,name=spec_value_id,json=specValueId,proto3" json:"spec_value_id,omitempty"`
}

func (x *SpecItemIndex) Reset() {
	*x = SpecItemIndex{}
	if protoimpl.UnsafeEnabled {
		mi := &file_skuService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpecItemIndex) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpecItemIndex) ProtoMessage() {}

func (x *SpecItemIndex) ProtoReflect() protoreflect.Message {
	mi := &file_skuService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpecItemIndex.ProtoReflect.Descriptor instead.
func (*SpecItemIndex) Descriptor() ([]byte, []int) {
	return file_skuService_proto_rawDescGZIP(), []int{1}
}

func (x *SpecItemIndex) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SpecItemIndex) GetItemId() int64 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *SpecItemIndex) GetSkuId() int64 {
	if x != nil {
		return x.SkuId
	}
	return 0
}

func (x *SpecItemIndex) GetSpecId() int32 {
	if x != nil {
		return x.SpecId
	}
	return 0
}

func (x *SpecItemIndex) GetSpecValueId() int32 {
	if x != nil {
		return x.SpecValueId
	}
	return 0
}

type SkuResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Sku          `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items  []*Sku        `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error  *common.Error `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info   *common.Info  `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *SkuResponse) Reset() {
	*x = SkuResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_skuService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SkuResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SkuResponse) ProtoMessage() {}

func (x *SkuResponse) ProtoReflect() protoreflect.Message {
	mi := &file_skuService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SkuResponse.ProtoReflect.Descriptor instead.
func (*SkuResponse) Descriptor() ([]byte, []int) {
	return file_skuService_proto_rawDescGZIP(), []int{2}
}

func (x *SkuResponse) GetEntity() *Sku {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *SkuResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *SkuResponse) GetItems() []*Sku {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *SkuResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *SkuResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_skuService_proto protoreflect.FileDescriptor

var file_skuService_proto_rawDesc = []byte{
	0x0a, 0x10, 0x73, 0x6b, 0x75, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x16, 0x73, 0x70, 0x65, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdf, 0x04, 0x0a, 0x03, 0x53, 0x6b, 0x75, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x6b, 0x75, 0x5f,
	0x73, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x6b, 0x75, 0x53, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x5f, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x6f, 0x72, 0x69, 0x67, 0x69,
	0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x63, 0x6f, 0x73, 0x74,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x6f, 0x6c,
	0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x6f, 0x6c,
	0x64, 0x4e, 0x75, 0x6d, 0x12, 0x22, 0x0a, 0x0d, 0x69, 0x6e, 0x69, 0x74, 0x5f, 0x73, 0x6f, 0x6c,
	0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x69, 0x6e, 0x69,
	0x74, 0x53, 0x6f, 0x6c, 0x64, 0x4e, 0x75, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x72, 0x6f, 0x7a,
	0x65, 0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x66, 0x72,
	0x6f, 0x7a, 0x65, 0x6e, 0x4e, 0x75, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x69, 0x73, 0x74, 0x65,
	0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x64, 0x12,
	0x1c, 0x0a, 0x0a, 0x6f, 0x75, 0x74, 0x5f, 0x73, 0x6b, 0x75, 0x5f, 0x6e, 0x6f, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x75, 0x74, 0x53, 0x6b, 0x75, 0x4e, 0x6f, 0x12, 0x1b, 0x0a,
	0x09, 0x73, 0x70, 0x65, 0x63, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x73, 0x70, 0x65, 0x63, 0x44, 0x65, 0x73, 0x63, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x65, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x64,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x75, 0x79, 0x5f,
	0x71, 0x75, 0x6f, 0x74, 0x61, 0x18, 0x12, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x62, 0x75, 0x79,
	0x51, 0x75, 0x6f, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x34, 0x0a, 0x0b, 0x73, 0x70, 0x65, 0x63, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x18, 0x15, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x53, 0x70, 0x65, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x73,
	0x70, 0x65, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x8c, 0x01, 0x0a, 0x0d, 0x53, 0x70,
	0x65, 0x63, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x69,
	0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x69, 0x74,
	0x65, 0x6d, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x6b, 0x75, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x6b, 0x75, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x73,
	0x70, 0x65, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x70,
	0x65, 0x63, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x73, 0x70, 0x65, 0x63, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x73, 0x70, 0x65,
	0x63, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x49, 0x64, 0x22, 0xc5, 0x01, 0x0a, 0x0b, 0x53, 0x6b, 0x75,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x53, 0x6b, 0x75, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12,
	0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70,
	0x61, 0x67, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53,
	0x6b, 0x75, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20,
	0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f,
	0x32, 0x7b, 0x0a, 0x0a, 0x53, 0x6b, 0x75, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d,
	0x0a, 0x13, 0x42, 0x61, 0x74, 0x63, 0x68, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x70, 0x65,
	0x63, 0x44, 0x65, 0x73, 0x63, 0x12, 0x0d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x53, 0x6b, 0x75, 0x1a, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x53, 0x6b, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2e, 0x0a,
	0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x53, 0x6b, 0x75, 0x1a, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x53, 0x6b, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a,
	0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_skuService_proto_rawDescOnce sync.Once
	file_skuService_proto_rawDescData = file_skuService_proto_rawDesc
)

func file_skuService_proto_rawDescGZIP() []byte {
	file_skuService_proto_rawDescOnce.Do(func() {
		file_skuService_proto_rawDescData = protoimpl.X.CompressGZIP(file_skuService_proto_rawDescData)
	})
	return file_skuService_proto_rawDescData
}

var file_skuService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_skuService_proto_goTypes = []interface{}{
	(*Sku)(nil),           // 0: services.Sku
	(*SpecItemIndex)(nil), // 1: services.SpecItemIndex
	(*SkuResponse)(nil),   // 2: services.SkuResponse
	(*SpecValue)(nil),     // 3: services.SpecValue
	(*common.Pager)(nil),  // 4: common.Pager
	(*common.Error)(nil),  // 5: common.Error
	(*common.Info)(nil),   // 6: common.Info
}
var file_skuService_proto_depIdxs = []int32{
	3, // 0: services.Sku.spec_values:type_name -> services.SpecValue
	0, // 1: services.SkuResponse.entity:type_name -> services.Sku
	4, // 2: services.SkuResponse.pager:type_name -> common.Pager
	0, // 3: services.SkuResponse.items:type_name -> services.Sku
	5, // 4: services.SkuResponse.error:type_name -> common.Error
	6, // 5: services.SkuResponse.info:type_name -> common.Info
	0, // 6: services.SkuService.BatchUpdateSpecDesc:input_type -> services.Sku
	0, // 7: services.SkuService.List:input_type -> services.Sku
	2, // 8: services.SkuService.BatchUpdateSpecDesc:output_type -> services.SkuResponse
	2, // 9: services.SkuService.List:output_type -> services.SkuResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_skuService_proto_init() }
func file_skuService_proto_init() {
	if File_skuService_proto != nil {
		return
	}
	file_specValueService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_skuService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sku); i {
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
		file_skuService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpecItemIndex); i {
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
		file_skuService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SkuResponse); i {
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
			RawDescriptor: file_skuService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_skuService_proto_goTypes,
		DependencyIndexes: file_skuService_proto_depIdxs,
		MessageInfos:      file_skuService_proto_msgTypes,
	}.Build()
	File_skuService_proto = out.File
	file_skuService_proto_rawDesc = nil
	file_skuService_proto_goTypes = nil
	file_skuService_proto_depIdxs = nil
}