// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: cartService.proto

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

type Cart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalNum       int32           `protobuf:"varint,1,opt,name=total_num,json=totalNum,proto3" json:"total_num"`
	TotalWeight    float32         `protobuf:"fixed32,2,opt,name=total_weight,json=totalWeight,proto3" json:"total_weight"`
	Currency       string          `protobuf:"bytes,3,opt,name=currency,proto3" json:"currency"`
	GoodsAmount    float32         `protobuf:"fixed32,4,opt,name=goods_amount,json=goodsAmount,proto3" json:"goods_amount"`
	DiscountAmount float32         `protobuf:"fixed32,5,opt,name=discount_amount,json=discountAmount,proto3" json:"discount_amount"`
	ShippingFee    float32         `protobuf:"fixed32,6,opt,name=shipping_fee,json=shippingFee,proto3" json:"shipping_fee"`
	InsureFee      float32         `protobuf:"fixed32,7,opt,name=insure_fee,json=insureFee,proto3" json:"insure_fee"`
	PackFee        float32         `protobuf:"fixed32,8,opt,name=pack_fee,json=packFee,proto3" json:"pack_fee"`
	Amount         float32         `protobuf:"fixed32,9,opt,name=amount,proto3" json:"amount"`
	MemberId       int64           `protobuf:"varint,10,opt,name=member_id,json=memberId,proto3" json:"member_id"`
	AddressId      int64           `protobuf:"varint,11,opt,name=address_id,json=addressId,proto3" json:"address_id"`
	Items          []*CartItem     `protobuf:"bytes,12,rep,name=items,proto3" json:"items"`
	SkuList        map[int64]int32 `protobuf:"bytes,13,rep,name=sku_list,json=skuList,proto3" json:"sku_list" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *Cart) Reset() {
	*x = Cart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cartService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cart) ProtoMessage() {}

func (x *Cart) ProtoReflect() protoreflect.Message {
	mi := &file_cartService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cart.ProtoReflect.Descriptor instead.
func (*Cart) Descriptor() ([]byte, []int) {
	return file_cartService_proto_rawDescGZIP(), []int{0}
}

func (x *Cart) GetTotalNum() int32 {
	if x != nil {
		return x.TotalNum
	}
	return 0
}

func (x *Cart) GetTotalWeight() float32 {
	if x != nil {
		return x.TotalWeight
	}
	return 0
}

func (x *Cart) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *Cart) GetGoodsAmount() float32 {
	if x != nil {
		return x.GoodsAmount
	}
	return 0
}

func (x *Cart) GetDiscountAmount() float32 {
	if x != nil {
		return x.DiscountAmount
	}
	return 0
}

func (x *Cart) GetShippingFee() float32 {
	if x != nil {
		return x.ShippingFee
	}
	return 0
}

func (x *Cart) GetInsureFee() float32 {
	if x != nil {
		return x.InsureFee
	}
	return 0
}

func (x *Cart) GetPackFee() float32 {
	if x != nil {
		return x.PackFee
	}
	return 0
}

func (x *Cart) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Cart) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

func (x *Cart) GetAddressId() int64 {
	if x != nil {
		return x.AddressId
	}
	return 0
}

func (x *Cart) GetItems() []*CartItem {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *Cart) GetSkuList() map[int64]int32 {
	if x != nil {
		return x.SkuList
	}
	return nil
}

type CartItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SpuId             int64   `protobuf:"varint,1,opt,name=spu_id,json=spuId,proto3" json:"spu_id"`
	SkuId             int64   `protobuf:"varint,2,opt,name=sku_id,json=skuId,proto3" json:"sku_id"`
	SkuSn             string  `protobuf:"bytes,3,opt,name=sku_sn,json=skuSn,proto3" json:"sku_sn"`
	ImageId           int64   `protobuf:"varint,4,opt,name=image_id,json=imageId,proto3" json:"image_id"`
	ImageUrl          string  `protobuf:"bytes,5,opt,name=image_url,json=imageUrl,proto3" json:"image_url"`
	Title             string  `protobuf:"bytes,6,opt,name=title,proto3" json:"title"`
	Weight            float32 `protobuf:"fixed32,7,opt,name=weight,proto3" json:"weight"`
	Num               int32   `protobuf:"varint,8,opt,name=num,proto3" json:"num"`
	Price             float32 `protobuf:"fixed32,9,opt,name=price,proto3" json:"price"`
	MarketPrice       float32 `protobuf:"fixed32,10,opt,name=market_price,json=marketPrice,proto3" json:"market_price"`
	PromotionPrice    float32 `protobuf:"fixed32,11,opt,name=promotion_price,json=promotionPrice,proto3" json:"promotion_price"`
	CostPrice         float32 `protobuf:"fixed32,12,opt,name=cost_price,json=costPrice,proto3" json:"cost_price"`
	RowGoodsAmount    float32 `protobuf:"fixed32,13,opt,name=row_goods_amount,json=rowGoodsAmount,proto3" json:"row_goods_amount"`
	RowDiscountAmount float32 `protobuf:"fixed32,14,opt,name=row_discount_amount,json=rowDiscountAmount,proto3" json:"row_discount_amount"`
	RowAmount         float32 `protobuf:"fixed32,15,opt,name=row_amount,json=rowAmount,proto3" json:"row_amount"`
	BrandId           int32   `protobuf:"varint,16,opt,name=brand_id,json=brandId,proto3" json:"brand_id"`
	BrandName         string  `protobuf:"bytes,17,opt,name=brand_name,json=brandName,proto3" json:"brand_name"`
	UnitId            int32   `protobuf:"varint,18,opt,name=unit_id,json=unitId,proto3" json:"unit_id"`
	UnitName          string  `protobuf:"bytes,19,opt,name=unit_name,json=unitName,proto3" json:"unit_name"`
	Indexes           string  `protobuf:"bytes,20,opt,name=indexes,proto3" json:"indexes"`
	OwnDesc           string  `protobuf:"bytes,21,opt,name=own_desc,json=ownDesc,proto3" json:"own_desc"`
	IsSpec            bool    `protobuf:"varint,22,opt,name=is_spec,json=isSpec,proto3" json:"is_spec"`
	IsVirtual         bool    `protobuf:"varint,23,opt,name=is_virtual,json=isVirtual,proto3" json:"is_virtual"`
	IsPresale         bool    `protobuf:"varint,24,opt,name=is_presale,json=isPresale,proto3" json:"is_presale"`
	Status            string  `protobuf:"bytes,25,opt,name=status,proto3" json:"status"`
}

func (x *CartItem) Reset() {
	*x = CartItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cartService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartItem) ProtoMessage() {}

func (x *CartItem) ProtoReflect() protoreflect.Message {
	mi := &file_cartService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartItem.ProtoReflect.Descriptor instead.
func (*CartItem) Descriptor() ([]byte, []int) {
	return file_cartService_proto_rawDescGZIP(), []int{1}
}

func (x *CartItem) GetSpuId() int64 {
	if x != nil {
		return x.SpuId
	}
	return 0
}

func (x *CartItem) GetSkuId() int64 {
	if x != nil {
		return x.SkuId
	}
	return 0
}

func (x *CartItem) GetSkuSn() string {
	if x != nil {
		return x.SkuSn
	}
	return ""
}

func (x *CartItem) GetImageId() int64 {
	if x != nil {
		return x.ImageId
	}
	return 0
}

func (x *CartItem) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *CartItem) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CartItem) GetWeight() float32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *CartItem) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *CartItem) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CartItem) GetMarketPrice() float32 {
	if x != nil {
		return x.MarketPrice
	}
	return 0
}

func (x *CartItem) GetPromotionPrice() float32 {
	if x != nil {
		return x.PromotionPrice
	}
	return 0
}

func (x *CartItem) GetCostPrice() float32 {
	if x != nil {
		return x.CostPrice
	}
	return 0
}

func (x *CartItem) GetRowGoodsAmount() float32 {
	if x != nil {
		return x.RowGoodsAmount
	}
	return 0
}

func (x *CartItem) GetRowDiscountAmount() float32 {
	if x != nil {
		return x.RowDiscountAmount
	}
	return 0
}

func (x *CartItem) GetRowAmount() float32 {
	if x != nil {
		return x.RowAmount
	}
	return 0
}

func (x *CartItem) GetBrandId() int32 {
	if x != nil {
		return x.BrandId
	}
	return 0
}

func (x *CartItem) GetBrandName() string {
	if x != nil {
		return x.BrandName
	}
	return ""
}

func (x *CartItem) GetUnitId() int32 {
	if x != nil {
		return x.UnitId
	}
	return 0
}

func (x *CartItem) GetUnitName() string {
	if x != nil {
		return x.UnitName
	}
	return ""
}

func (x *CartItem) GetIndexes() string {
	if x != nil {
		return x.Indexes
	}
	return ""
}

func (x *CartItem) GetOwnDesc() string {
	if x != nil {
		return x.OwnDesc
	}
	return ""
}

func (x *CartItem) GetIsSpec() bool {
	if x != nil {
		return x.IsSpec
	}
	return false
}

func (x *CartItem) GetIsVirtual() bool {
	if x != nil {
		return x.IsVirtual
	}
	return false
}

func (x *CartItem) GetIsPresale() bool {
	if x != nil {
		return x.IsPresale
	}
	return false
}

func (x *CartItem) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type CartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged     int32           `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize  int32           `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Keywords  string          `protobuf:"bytes,3,opt,name=keywords,proto3" json:"keywords"`
	RowId     string          `protobuf:"bytes,4,opt,name=row_id,json=rowId,proto3" json:"row_id"`
	SkuId     int64           `protobuf:"varint,5,opt,name=sku_id,json=skuId,proto3" json:"sku_id"`
	Direction string          `protobuf:"bytes,6,opt,name=direction,proto3" json:"direction"`
	Num       int32           `protobuf:"varint,7,opt,name=num,proto3" json:"num"`
	SkuList   map[int64]int32 `protobuf:"bytes,13,rep,name=sku_list,json=skuList,proto3" json:"sku_list" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *CartRequest) Reset() {
	*x = CartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cartService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartRequest) ProtoMessage() {}

func (x *CartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cartService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartRequest.ProtoReflect.Descriptor instead.
func (*CartRequest) Descriptor() ([]byte, []int) {
	return file_cartService_proto_rawDescGZIP(), []int{2}
}

func (x *CartRequest) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *CartRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *CartRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *CartRequest) GetRowId() string {
	if x != nil {
		return x.RowId
	}
	return ""
}

func (x *CartRequest) GetSkuId() int64 {
	if x != nil {
		return x.SkuId
	}
	return 0
}

func (x *CartRequest) GetDirection() string {
	if x != nil {
		return x.Direction
	}
	return ""
}

func (x *CartRequest) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *CartRequest) GetSkuList() map[int64]int32 {
	if x != nil {
		return x.SkuList
	}
	return nil
}

type CartData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Cart         `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*Cart       `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Info   *common.Info  `protobuf:"bytes,4,opt,name=info,proto3" json:"info"`
}

func (x *CartData) Reset() {
	*x = CartData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cartService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartData) ProtoMessage() {}

func (x *CartData) ProtoReflect() protoreflect.Message {
	mi := &file_cartService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartData.ProtoReflect.Descriptor instead.
func (*CartData) Descriptor() ([]byte, []int) {
	return file_cartService_proto_rawDescGZIP(), []int{3}
}

func (x *CartData) GetEntity() *Cart {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *CartData) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *CartData) GetItems() []*Cart {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *CartData) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type CartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *CartData     `protobuf:"bytes,1,opt,name=data,proto3" json:"data"`
	Error *common.Error `protobuf:"bytes,2,opt,name=error,proto3" json:"error"`
}

func (x *CartResponse) Reset() {
	*x = CartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cartService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartResponse) ProtoMessage() {}

func (x *CartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cartService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartResponse.ProtoReflect.Descriptor instead.
func (*CartResponse) Descriptor() ([]byte, []int) {
	return file_cartService_proto_rawDescGZIP(), []int{4}
}

func (x *CartResponse) GetData() *CartData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *CartResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_cartService_proto protoreflect.FileDescriptor

var file_cartService_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x61, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xfd, 0x03, 0x0a, 0x04, 0x43, 0x61, 0x72, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f,
	0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x5f, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x67, 0x6f, 0x6f,
	0x64, 0x73, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x0e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x65,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x46, 0x65, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x73, 0x75, 0x72, 0x65, 0x5f, 0x66,
	0x65, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x69, 0x6e, 0x73, 0x75, 0x72, 0x65,
	0x46, 0x65, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x63, 0x6b, 0x5f, 0x66, 0x65, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x46, 0x65, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x69,
	0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x49, 0x64, 0x12, 0x28, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x72,
	0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x36, 0x0a, 0x08,
	0x73, 0x6b, 0x75, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x2e, 0x53,
	0x6b, 0x75, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x73, 0x6b, 0x75,
	0x4c, 0x69, 0x73, 0x74, 0x1a, 0x3a, 0x0a, 0x0c, 0x53, 0x6b, 0x75, 0x4c, 0x69, 0x73, 0x74, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0xd5, 0x05, 0x0a, 0x08, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x15, 0x0a,
	0x06, 0x73, 0x70, 0x75, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73,
	0x70, 0x75, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x6b, 0x75, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x6b, 0x75, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x73,
	0x6b, 0x75, 0x5f, 0x73, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x6b, 0x75,
	0x53, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0e, 0x70, 0x72,
	0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x6f, 0x73, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x09, 0x63, 0x6f, 0x73, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x72,
	0x6f, 0x77, 0x5f, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0e, 0x72, 0x6f, 0x77, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x13, 0x72, 0x6f, 0x77, 0x5f, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x11, 0x72, 0x6f, 0x77, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x6f, 0x77, 0x5f, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x72, 0x6f, 0x77, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x5f, 0x69, 0x64,
	0x18, 0x10, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x49, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x11, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x6e, 0x69, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x12, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x75, 0x6e, 0x69, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x6e, 0x69, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x6e, 0x69, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x18,
	0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x12, 0x19,
	0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6f, 0x77, 0x6e, 0x44, 0x65, 0x73, 0x63, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f,
	0x73, 0x70, 0x65, 0x63, 0x18, 0x16, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x53, 0x70,
	0x65, 0x63, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c,
	0x18, 0x17, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61,
	0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x70, 0x72, 0x65, 0x73, 0x61, 0x6c, 0x65, 0x18,
	0x18, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x50, 0x72, 0x65, 0x73, 0x61, 0x6c, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x19, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xb5, 0x02, 0x0a, 0x0b, 0x43, 0x61, 0x72,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6b,
	0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b,
	0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x15, 0x0a, 0x06, 0x72, 0x6f, 0x77, 0x5f, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x6f, 0x77, 0x49, 0x64, 0x12, 0x15,
	0x0a, 0x06, 0x73, 0x6b, 0x75, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x73, 0x6b, 0x75, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x03, 0x6e, 0x75, 0x6d, 0x12, 0x3d, 0x0a, 0x08, 0x73, 0x6b, 0x75, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x53,
	0x6b, 0x75, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x73, 0x6b, 0x75,
	0x4c, 0x69, 0x73, 0x74, 0x1a, 0x3a, 0x0a, 0x0c, 0x53, 0x6b, 0x75, 0x4c, 0x69, 0x73, 0x74, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x9f, 0x01, 0x0a, 0x08, 0x43, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x26, 0x0a,
	0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x52, 0x06, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61,
	0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x22, 0x5b, 0x0a, 0x0c, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32,
	0xa7, 0x03, 0x0a, 0x0b, 0x43, 0x61, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x36, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x72,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x39, 0x0a, 0x06, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x12, 0x15, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43,
	0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a,
	0x08, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x08, 0x44, 0x65,
	0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x05, 0x43, 0x6c, 0x65, 0x61, 0x72,
	0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x36, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cartService_proto_rawDescOnce sync.Once
	file_cartService_proto_rawDescData = file_cartService_proto_rawDesc
)

func file_cartService_proto_rawDescGZIP() []byte {
	file_cartService_proto_rawDescOnce.Do(func() {
		file_cartService_proto_rawDescData = protoimpl.X.CompressGZIP(file_cartService_proto_rawDescData)
	})
	return file_cartService_proto_rawDescData
}

var file_cartService_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_cartService_proto_goTypes = []interface{}{
	(*Cart)(nil),         // 0: services.Cart
	(*CartItem)(nil),     // 1: services.CartItem
	(*CartRequest)(nil),  // 2: services.CartRequest
	(*CartData)(nil),     // 3: services.CartData
	(*CartResponse)(nil), // 4: services.CartResponse
	nil,                  // 5: services.Cart.SkuListEntry
	nil,                  // 6: services.CartRequest.SkuListEntry
	(*common.Pager)(nil), // 7: common.Pager
	(*common.Info)(nil),  // 8: common.Info
	(*common.Error)(nil), // 9: common.Error
}
var file_cartService_proto_depIdxs = []int32{
	1,  // 0: services.Cart.items:type_name -> services.CartItem
	5,  // 1: services.Cart.sku_list:type_name -> services.Cart.SkuListEntry
	6,  // 2: services.CartRequest.sku_list:type_name -> services.CartRequest.SkuListEntry
	0,  // 3: services.CartData.entity:type_name -> services.Cart
	7,  // 4: services.CartData.pager:type_name -> common.Pager
	0,  // 5: services.CartData.items:type_name -> services.Cart
	8,  // 6: services.CartData.info:type_name -> common.Info
	3,  // 7: services.CartResponse.data:type_name -> services.CartData
	9,  // 8: services.CartResponse.error:type_name -> common.Error
	2,  // 9: services.CartService.Add:input_type -> services.CartRequest
	2,  // 10: services.CartService.Update:input_type -> services.CartRequest
	2,  // 11: services.CartService.Remove:input_type -> services.CartRequest
	2,  // 12: services.CartService.Increase:input_type -> services.CartRequest
	2,  // 13: services.CartService.Decrease:input_type -> services.CartRequest
	2,  // 14: services.CartService.Clear:input_type -> services.CartRequest
	2,  // 15: services.CartService.Get:input_type -> services.CartRequest
	4,  // 16: services.CartService.Add:output_type -> services.CartResponse
	4,  // 17: services.CartService.Update:output_type -> services.CartResponse
	4,  // 18: services.CartService.Remove:output_type -> services.CartResponse
	4,  // 19: services.CartService.Increase:output_type -> services.CartResponse
	4,  // 20: services.CartService.Decrease:output_type -> services.CartResponse
	4,  // 21: services.CartService.Clear:output_type -> services.CartResponse
	4,  // 22: services.CartService.Get:output_type -> services.CartResponse
	16, // [16:23] is the sub-list for method output_type
	9,  // [9:16] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_cartService_proto_init() }
func file_cartService_proto_init() {
	if File_cartService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cartService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cart); i {
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
		file_cartService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartItem); i {
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
		file_cartService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartRequest); i {
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
		file_cartService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartData); i {
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
		file_cartService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartResponse); i {
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
			RawDescriptor: file_cartService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cartService_proto_goTypes,
		DependencyIndexes: file_cartService_proto_depIdxs,
		MessageInfos:      file_cartService_proto_msgTypes,
	}.Build()
	File_cartService_proto = out.File
	file_cartService_proto_rawDesc = nil
	file_cartService_proto_goTypes = nil
	file_cartService_proto_depIdxs = nil
}
