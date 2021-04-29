// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: itemPriceService.proto

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

type ItemPrice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Type        string  `protobuf:"bytes,2,opt,name=type,proto3" json:"type"`
	TypeValueId int64   `protobuf:"varint,3,opt,name=type_value_id,json=typeValueId,proto3" json:"type_value_id"`
	ItemId      int64   `protobuf:"varint,4,opt,name=item_id,json=itemId,proto3" json:"item_id"`
	SkuId       int64   `protobuf:"varint,5,opt,name=sku_id,json=skuId,proto3" json:"sku_id"`
	Method      string  `protobuf:"bytes,6,opt,name=method,proto3" json:"method"`
	Price       float32 `protobuf:"fixed32,7,opt,name=price,proto3" json:"price"`
	Discount    float32 `protobuf:"fixed32,8,opt,name=discount,proto3" json:"discount"`
	// @inject_tag: gorm:"-"
	Sku *Sku `protobuf:"bytes,9,opt,name=sku,proto3" json:"sku"  gorm:"-"`
}

func (x *ItemPrice) Reset() {
	*x = ItemPrice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_itemPriceService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemPrice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemPrice) ProtoMessage() {}

func (x *ItemPrice) ProtoReflect() protoreflect.Message {
	mi := &file_itemPriceService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemPrice.ProtoReflect.Descriptor instead.
func (*ItemPrice) Descriptor() ([]byte, []int) {
	return file_itemPriceService_proto_rawDescGZIP(), []int{0}
}

func (x *ItemPrice) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ItemPrice) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ItemPrice) GetTypeValueId() int64 {
	if x != nil {
		return x.TypeValueId
	}
	return 0
}

func (x *ItemPrice) GetItemId() int64 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *ItemPrice) GetSkuId() int64 {
	if x != nil {
		return x.SkuId
	}
	return 0
}

func (x *ItemPrice) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *ItemPrice) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *ItemPrice) GetDiscount() float32 {
	if x != nil {
		return x.Discount
	}
	return 0
}

func (x *ItemPrice) GetSku() *Sku {
	if x != nil {
		return x.Sku
	}
	return nil
}

type ItemPriceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *ItemPrice    `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*ItemPrice  `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Error  *common.Error `protobuf:"bytes,4,opt,name=error,proto3" json:"error"`
	Info   *common.Info  `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
}

func (x *ItemPriceResponse) Reset() {
	*x = ItemPriceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_itemPriceService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemPriceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemPriceResponse) ProtoMessage() {}

func (x *ItemPriceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_itemPriceService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemPriceResponse.ProtoReflect.Descriptor instead.
func (*ItemPriceResponse) Descriptor() ([]byte, []int) {
	return file_itemPriceService_proto_rawDescGZIP(), []int{1}
}

func (x *ItemPriceResponse) GetEntity() *ItemPrice {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *ItemPriceResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *ItemPriceResponse) GetItems() []*ItemPrice {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ItemPriceResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *ItemPriceResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

//价格方案
type PriceScheme struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemId          int64   `protobuf:"varint,1,opt,name=item_id,json=itemId,proto3" json:"item_id"`
	ItemSn          string  `protobuf:"bytes,2,opt,name=item_sn,json=itemSn,proto3" json:"item_sn"`
	ItemName        string  `protobuf:"bytes,3,opt,name=item_name,json=itemName,proto3" json:"item_name"`
	ItemPrice       float32 `protobuf:"fixed32,4,opt,name=item_price,json=itemPrice,proto3" json:"item_price"`
	ItemOriginPrice float32 `protobuf:"fixed32,5,opt,name=item_origin_price,json=itemOriginPrice,proto3" json:"item_origin_price"`
	ModelType       string  `protobuf:"bytes,7,opt,name=model_type,json=modelType,proto3" json:"model_type"`
	Type            string  `protobuf:"bytes,8,opt,name=type,proto3" json:"type"`
	Method          string  `protobuf:"bytes,9,opt,name=method,proto3" json:"method"`
	// @inject_tag: gorm:"-"
	Details []*ItemPrice `protobuf:"bytes,10,rep,name=details,proto3" json:"details"`
	// @inject_tag: gorm:"-"
	Skus []*Sku `protobuf:"bytes,11,rep,name=skus,proto3" json:"skus"`
	// @inject_tag: gorm:"-"
	Specs []*Spec `protobuf:"bytes,12,rep,name=specs,proto3" json:"specs"`
}

func (x *PriceScheme) Reset() {
	*x = PriceScheme{}
	if protoimpl.UnsafeEnabled {
		mi := &file_itemPriceService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PriceScheme) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PriceScheme) ProtoMessage() {}

func (x *PriceScheme) ProtoReflect() protoreflect.Message {
	mi := &file_itemPriceService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PriceScheme.ProtoReflect.Descriptor instead.
func (*PriceScheme) Descriptor() ([]byte, []int) {
	return file_itemPriceService_proto_rawDescGZIP(), []int{2}
}

func (x *PriceScheme) GetItemId() int64 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *PriceScheme) GetItemSn() string {
	if x != nil {
		return x.ItemSn
	}
	return ""
}

func (x *PriceScheme) GetItemName() string {
	if x != nil {
		return x.ItemName
	}
	return ""
}

func (x *PriceScheme) GetItemPrice() float32 {
	if x != nil {
		return x.ItemPrice
	}
	return 0
}

func (x *PriceScheme) GetItemOriginPrice() float32 {
	if x != nil {
		return x.ItemOriginPrice
	}
	return 0
}

func (x *PriceScheme) GetModelType() string {
	if x != nil {
		return x.ModelType
	}
	return ""
}

func (x *PriceScheme) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *PriceScheme) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *PriceScheme) GetDetails() []*ItemPrice {
	if x != nil {
		return x.Details
	}
	return nil
}

func (x *PriceScheme) GetSkus() []*Sku {
	if x != nil {
		return x.Skus
	}
	return nil
}

func (x *PriceScheme) GetSpecs() []*Spec {
	if x != nil {
		return x.Specs
	}
	return nil
}

type PriceSchemeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *PriceScheme   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager  `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*PriceScheme `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Error  *common.Error  `protobuf:"bytes,4,opt,name=error,proto3" json:"error"`
	Info   *common.Info   `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
}

func (x *PriceSchemeResponse) Reset() {
	*x = PriceSchemeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_itemPriceService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PriceSchemeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PriceSchemeResponse) ProtoMessage() {}

func (x *PriceSchemeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_itemPriceService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PriceSchemeResponse.ProtoReflect.Descriptor instead.
func (*PriceSchemeResponse) Descriptor() ([]byte, []int) {
	return file_itemPriceService_proto_rawDescGZIP(), []int{3}
}

func (x *PriceSchemeResponse) GetEntity() *PriceScheme {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *PriceSchemeResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *PriceSchemeResponse) GetItems() []*PriceScheme {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *PriceSchemeResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *PriceSchemeResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_itemPriceService_proto protoreflect.FileDescriptor

var file_itemPriceService_proto_rawDesc = []byte{
	0x0a, 0x16, 0x69, 0x74, 0x65, 0x6d, 0x50, 0x72, 0x69, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x73, 0x6b, 0x75, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x73, 0x70, 0x65, 0x63, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xee, 0x01, 0x0a, 0x09, 0x49,
	0x74, 0x65, 0x6d, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x0d,
	0x74, 0x79, 0x70, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0b, 0x74, 0x79, 0x70, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x49, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x6b, 0x75,
	0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x6b, 0x75, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x03, 0x73, 0x6b,
	0x75, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x53, 0x6b, 0x75, 0x52, 0x03, 0x73, 0x6b, 0x75, 0x22, 0xd7, 0x01, 0x0a, 0x11,
	0x49, 0x74, 0x65, 0x6d, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2b, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x74, 0x65,
	0x6d, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23,
	0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61,
	0x67, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x74,
	0x65, 0x6d, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x23,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0xea, 0x02, 0x0a, 0x0b, 0x50, 0x72, 0x69, 0x63, 0x65, 0x53,
	0x63, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x73, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x69, 0x74, 0x65, 0x6d, 0x53, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x74, 0x65, 0x6d,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0f,
	0x69, 0x74, 0x65, 0x6d, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x2d, 0x0a, 0x07, 0x64, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x21, 0x0a, 0x04, 0x73, 0x6b, 0x75,
	0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x53, 0x6b, 0x75, 0x52, 0x04, 0x73, 0x6b, 0x75, 0x73, 0x12, 0x24, 0x0a, 0x05,
	0x73, 0x70, 0x65, 0x63, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x70, 0x65, 0x63, 0x52, 0x05, 0x73, 0x70, 0x65,
	0x63, 0x73, 0x22, 0xdd, 0x01, 0x0a, 0x13, 0x50, 0x72, 0x69, 0x63, 0x65, 0x53, 0x63, 0x68, 0x65,
	0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d,
	0x65, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x2b,
	0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x53, 0x63,
	0x68, 0x65, 0x6d, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x23, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x32, 0xcc, 0x01, 0x0a, 0x10, 0x49, 0x74, 0x65, 0x6d, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x03, 0x53, 0x65, 0x74, 0x12, 0x15,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x53,
	0x63, 0x68, 0x65, 0x6d, 0x65, 0x1a, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x13, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x74,
	0x65, 0x6d, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x3e, 0x0a, 0x06, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x13, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x1a, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_itemPriceService_proto_rawDescOnce sync.Once
	file_itemPriceService_proto_rawDescData = file_itemPriceService_proto_rawDesc
)

func file_itemPriceService_proto_rawDescGZIP() []byte {
	file_itemPriceService_proto_rawDescOnce.Do(func() {
		file_itemPriceService_proto_rawDescData = protoimpl.X.CompressGZIP(file_itemPriceService_proto_rawDescData)
	})
	return file_itemPriceService_proto_rawDescData
}

var file_itemPriceService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_itemPriceService_proto_goTypes = []interface{}{
	(*ItemPrice)(nil),           // 0: services.ItemPrice
	(*ItemPriceResponse)(nil),   // 1: services.ItemPriceResponse
	(*PriceScheme)(nil),         // 2: services.PriceScheme
	(*PriceSchemeResponse)(nil), // 3: services.PriceSchemeResponse
	(*Sku)(nil),                 // 4: services.Sku
	(*common.Pager)(nil),        // 5: common.Pager
	(*common.Error)(nil),        // 6: common.Error
	(*common.Info)(nil),         // 7: common.Info
	(*Spec)(nil),                // 8: services.Spec
}
var file_itemPriceService_proto_depIdxs = []int32{
	4,  // 0: services.ItemPrice.sku:type_name -> services.Sku
	0,  // 1: services.ItemPriceResponse.entity:type_name -> services.ItemPrice
	5,  // 2: services.ItemPriceResponse.pager:type_name -> common.Pager
	0,  // 3: services.ItemPriceResponse.items:type_name -> services.ItemPrice
	6,  // 4: services.ItemPriceResponse.error:type_name -> common.Error
	7,  // 5: services.ItemPriceResponse.info:type_name -> common.Info
	0,  // 6: services.PriceScheme.details:type_name -> services.ItemPrice
	4,  // 7: services.PriceScheme.skus:type_name -> services.Sku
	8,  // 8: services.PriceScheme.specs:type_name -> services.Spec
	2,  // 9: services.PriceSchemeResponse.entity:type_name -> services.PriceScheme
	5,  // 10: services.PriceSchemeResponse.pager:type_name -> common.Pager
	2,  // 11: services.PriceSchemeResponse.items:type_name -> services.PriceScheme
	6,  // 12: services.PriceSchemeResponse.error:type_name -> common.Error
	7,  // 13: services.PriceSchemeResponse.info:type_name -> common.Info
	2,  // 14: services.ItemPriceService.Set:input_type -> services.PriceScheme
	0,  // 15: services.ItemPriceService.Get:input_type -> services.ItemPrice
	0,  // 16: services.ItemPriceService.Scheme:input_type -> services.ItemPrice
	3,  // 17: services.ItemPriceService.Set:output_type -> services.PriceSchemeResponse
	1,  // 18: services.ItemPriceService.Get:output_type -> services.ItemPriceResponse
	3,  // 19: services.ItemPriceService.Scheme:output_type -> services.PriceSchemeResponse
	17, // [17:20] is the sub-list for method output_type
	14, // [14:17] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_itemPriceService_proto_init() }
func file_itemPriceService_proto_init() {
	if File_itemPriceService_proto != nil {
		return
	}
	file_skuService_proto_init()
	file_specService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_itemPriceService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemPrice); i {
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
		file_itemPriceService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemPriceResponse); i {
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
		file_itemPriceService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PriceScheme); i {
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
		file_itemPriceService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PriceSchemeResponse); i {
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
			RawDescriptor: file_itemPriceService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_itemPriceService_proto_goTypes,
		DependencyIndexes: file_itemPriceService_proto_depIdxs,
		MessageInfos:      file_itemPriceService_proto_msgTypes,
	}.Build()
	File_itemPriceService_proto = out.File
	file_itemPriceService_proto_rawDesc = nil
	file_itemPriceService_proto_goTypes = nil
	file_itemPriceService_proto_depIdxs = nil
}
