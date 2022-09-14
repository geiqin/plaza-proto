// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: cartService.proto

package services

import (
	common "github.com/geiqin/micro-kit/protobuf/common"
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

type Cart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code          string           `protobuf:"bytes,1,opt,name=code,proto3" json:"code"`
	Currency      string           `protobuf:"bytes,2,opt,name=currency,proto3" json:"currency"`
	TotalNum      int32            `protobuf:"varint,3,opt,name=total_num,json=totalNum,proto3" json:"total_num"`
	TotalWeight   int64            `protobuf:"varint,4,opt,name=total_weight,json=totalWeight,proto3" json:"total_weight"`
	TotalPrice    int64            `protobuf:"varint,5,opt,name=total_price,json=totalPrice,proto3" json:"total_price"`
	TotalDiscount int64            `protobuf:"varint,6,opt,name=total_discount,json=totalDiscount,proto3" json:"total_discount"`
	TotalAmount   int64            `protobuf:"varint,7,opt,name=total_amount,json=totalAmount,proto3" json:"total_amount"`
	FreightFee    int64            `protobuf:"varint,8,opt,name=freight_fee,json=freightFee,proto3" json:"freight_fee"`
	InsureFee     int64            `protobuf:"varint,9,opt,name=insure_fee,json=insureFee,proto3" json:"insure_fee"`
	PackFee       int64            `protobuf:"varint,10,opt,name=pack_fee,json=packFee,proto3" json:"pack_fee"`
	MemberId      int64            `protobuf:"varint,11,opt,name=member_id,json=memberId,proto3" json:"member_id"`
	AddressId     int64            `protobuf:"varint,12,opt,name=address_id,json=addressId,proto3" json:"address_id"`
	CheckedStatus int32            `protobuf:"varint,13,opt,name=checked_status,json=checkedStatus,proto3" json:"checked_status"`
	Skus          []*CartItem      `protobuf:"bytes,14,rep,name=skus,proto3" json:"skus"`
	Promotions    []*PromotionInfo `protobuf:"bytes,15,rep,name=promotions,proto3" json:"promotions"`   //店铺营销活动
	Adjustments   []*Adjustment    `protobuf:"bytes,16,rep,name=adjustments,proto3" json:"adjustments"` //变动记录
	Presents      []*SkuInfo       `protobuf:"bytes,17,rep,name=presents,proto3" json:"presents"`       //可获得的赠品
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

func (x *Cart) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Cart) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *Cart) GetTotalNum() int32 {
	if x != nil {
		return x.TotalNum
	}
	return 0
}

func (x *Cart) GetTotalWeight() int64 {
	if x != nil {
		return x.TotalWeight
	}
	return 0
}

func (x *Cart) GetTotalPrice() int64 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

func (x *Cart) GetTotalDiscount() int64 {
	if x != nil {
		return x.TotalDiscount
	}
	return 0
}

func (x *Cart) GetTotalAmount() int64 {
	if x != nil {
		return x.TotalAmount
	}
	return 0
}

func (x *Cart) GetFreightFee() int64 {
	if x != nil {
		return x.FreightFee
	}
	return 0
}

func (x *Cart) GetInsureFee() int64 {
	if x != nil {
		return x.InsureFee
	}
	return 0
}

func (x *Cart) GetPackFee() int64 {
	if x != nil {
		return x.PackFee
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

func (x *Cart) GetCheckedStatus() int32 {
	if x != nil {
		return x.CheckedStatus
	}
	return 0
}

func (x *Cart) GetSkus() []*CartItem {
	if x != nil {
		return x.Skus
	}
	return nil
}

func (x *Cart) GetPromotions() []*PromotionInfo {
	if x != nil {
		return x.Promotions
	}
	return nil
}

func (x *Cart) GetAdjustments() []*Adjustment {
	if x != nil {
		return x.Adjustments
	}
	return nil
}

func (x *Cart) GetPresents() []*SkuInfo {
	if x != nil {
		return x.Presents
	}
	return nil
}

type CartItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SkuId            int64            `protobuf:"varint,1,opt,name=sku_id,json=skuId,proto3" json:"sku_id"`
	Num              int32            `protobuf:"varint,2,opt,name=num,proto3" json:"num"`
	Price            int64            `protobuf:"varint,3,opt,name=price,proto3" json:"price"`
	PromotionPrice   int64            `protobuf:"varint,4,opt,name=promotion_price,json=promotionPrice,proto3" json:"promotion_price"`
	SubtotalAmount   int64            `protobuf:"varint,5,opt,name=subtotal_amount,json=subtotalAmount,proto3" json:"subtotal_amount"`
	SubtotalDiscount int64            `protobuf:"varint,6,opt,name=subtotal_discount,json=subtotalDiscount,proto3" json:"subtotal_discount"`
	Checked          bool             `protobuf:"varint,7,opt,name=checked,proto3" json:"checked"`
	Sku              *SkuInfo         `protobuf:"bytes,8,opt,name=sku,proto3" json:"sku"`
	Promotions       []*PromotionInfo `protobuf:"bytes,9,rep,name=promotions,proto3" json:"promotions"`
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

func (x *CartItem) GetSkuId() int64 {
	if x != nil {
		return x.SkuId
	}
	return 0
}

func (x *CartItem) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *CartItem) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CartItem) GetPromotionPrice() int64 {
	if x != nil {
		return x.PromotionPrice
	}
	return 0
}

func (x *CartItem) GetSubtotalAmount() int64 {
	if x != nil {
		return x.SubtotalAmount
	}
	return 0
}

func (x *CartItem) GetSubtotalDiscount() int64 {
	if x != nil {
		return x.SubtotalDiscount
	}
	return 0
}

func (x *CartItem) GetChecked() bool {
	if x != nil {
		return x.Checked
	}
	return false
}

func (x *CartItem) GetSku() *SkuInfo {
	if x != nil {
		return x.Sku
	}
	return nil
}

func (x *CartItem) GetPromotions() []*PromotionInfo {
	if x != nil {
		return x.Promotions
	}
	return nil
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
	SkuIds    []int64         `protobuf:"varint,8,rep,packed,name=sku_ids,json=skuIds,proto3" json:"sku_ids"`
	SkuList   map[int64]int32 `protobuf:"bytes,9,rep,name=sku_list,json=skuList,proto3" json:"sku_list" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
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

func (x *CartRequest) GetSkuIds() []int64 {
	if x != nil {
		return x.SkuIds
	}
	return nil
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
	0x1a, 0x14, 0x73, 0x6b, 0x75, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x61, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe7, 0x04, 0x0a, 0x04,
	0x43, 0x61, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x6e, 0x75,
	0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4e, 0x75,
	0x6d, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x77, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x57, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x66, 0x72, 0x65, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x66, 0x65, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x66, 0x72, 0x65, 0x69, 0x67, 0x68, 0x74, 0x46, 0x65, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x73, 0x75, 0x72, 0x65, 0x5f, 0x66, 0x65, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x69, 0x6e, 0x73, 0x75, 0x72, 0x65, 0x46, 0x65, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x70, 0x61, 0x63, 0x6b, 0x5f, 0x66, 0x65, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x46, 0x65, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65,
	0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x26, 0x0a,
	0x04, 0x73, 0x6b, 0x75, 0x73, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x04, 0x73, 0x6b, 0x75, 0x73, 0x12, 0x37, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x36,
	0x0a, 0x0b, 0x61, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x10, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41,
	0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x61, 0x64, 0x6a, 0x75, 0x73,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x2d, 0x0a, 0x08, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e,
	0x74, 0x73, 0x18, 0x11, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x53, 0x6b, 0x75, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x70, 0x72, 0x65,
	0x73, 0x65, 0x6e, 0x74, 0x73, 0x22, 0xc0, 0x02, 0x0a, 0x08, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x6b, 0x75, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x73, 0x6b, 0x75, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x6d,
	0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x75,
	0x62, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0e, 0x73, 0x75, 0x62, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x2b, 0x0a, 0x11, 0x73, 0x75, 0x62, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10,
	0x73, 0x75, 0x62, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x64, 0x12, 0x23, 0x0a, 0x03, 0x73, 0x6b,
	0x75, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x53, 0x6b, 0x75, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x03, 0x73, 0x6b, 0x75, 0x12,
	0x37, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x09, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50,
	0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x70, 0x72,
	0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xce, 0x02, 0x0a, 0x0b, 0x43, 0x61, 0x72,
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
	0x52, 0x03, 0x6e, 0x75, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x6b, 0x75, 0x5f, 0x69, 0x64, 0x73,
	0x18, 0x08, 0x20, 0x03, 0x28, 0x03, 0x52, 0x06, 0x73, 0x6b, 0x75, 0x49, 0x64, 0x73, 0x12, 0x3d,
	0x0a, 0x08, 0x73, 0x6b, 0x75, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x6b, 0x75, 0x4c, 0x69, 0x73, 0x74, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x73, 0x6b, 0x75, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x3a, 0x0a,
	0x0c, 0x53, 0x6b, 0x75, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x9f, 0x01, 0x0a, 0x08, 0x43, 0x61,
	0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x26, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23,
	0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61,
	0x67, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61,
	0x72, 0x74, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x5b, 0x0a, 0x0c, 0x43,
	0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xe6, 0x02, 0x0a, 0x0b, 0x43, 0x61, 0x72,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12,
	0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x36, 0x0a, 0x03, 0x53, 0x65, 0x74, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x06, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61,
	0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x05, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x12, 0x15, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43,
	0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a,
	0x03, 0x47, 0x65, 0x74, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x07, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x64,
	0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_cartService_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_cartService_proto_goTypes = []interface{}{
	(*Cart)(nil),          // 0: services.Cart
	(*CartItem)(nil),      // 1: services.CartItem
	(*CartRequest)(nil),   // 2: services.CartRequest
	(*CartData)(nil),      // 3: services.CartData
	(*CartResponse)(nil),  // 4: services.CartResponse
	nil,                   // 5: services.CartRequest.SkuListEntry
	(*PromotionInfo)(nil), // 6: services.PromotionInfo
	(*Adjustment)(nil),    // 7: services.Adjustment
	(*SkuInfo)(nil),       // 8: services.SkuInfo
	(*common.Pager)(nil),  // 9: common.Pager
	(*common.Info)(nil),   // 10: common.Info
	(*common.Error)(nil),  // 11: common.Error
}
var file_cartService_proto_depIdxs = []int32{
	1,  // 0: services.Cart.skus:type_name -> services.CartItem
	6,  // 1: services.Cart.promotions:type_name -> services.PromotionInfo
	7,  // 2: services.Cart.adjustments:type_name -> services.Adjustment
	8,  // 3: services.Cart.presents:type_name -> services.SkuInfo
	8,  // 4: services.CartItem.sku:type_name -> services.SkuInfo
	6,  // 5: services.CartItem.promotions:type_name -> services.PromotionInfo
	5,  // 6: services.CartRequest.sku_list:type_name -> services.CartRequest.SkuListEntry
	0,  // 7: services.CartData.entity:type_name -> services.Cart
	9,  // 8: services.CartData.pager:type_name -> common.Pager
	0,  // 9: services.CartData.items:type_name -> services.Cart
	10, // 10: services.CartData.info:type_name -> common.Info
	3,  // 11: services.CartResponse.data:type_name -> services.CartData
	11, // 12: services.CartResponse.error:type_name -> common.Error
	2,  // 13: services.CartService.Add:input_type -> services.CartRequest
	2,  // 14: services.CartService.Set:input_type -> services.CartRequest
	2,  // 15: services.CartService.Remove:input_type -> services.CartRequest
	2,  // 16: services.CartService.Clear:input_type -> services.CartRequest
	2,  // 17: services.CartService.Get:input_type -> services.CartRequest
	2,  // 18: services.CartService.Checked:input_type -> services.CartRequest
	4,  // 19: services.CartService.Add:output_type -> services.CartResponse
	4,  // 20: services.CartService.Set:output_type -> services.CartResponse
	4,  // 21: services.CartService.Remove:output_type -> services.CartResponse
	4,  // 22: services.CartService.Clear:output_type -> services.CartResponse
	4,  // 23: services.CartService.Get:output_type -> services.CartResponse
	4,  // 24: services.CartService.Checked:output_type -> services.CartResponse
	19, // [19:25] is the sub-list for method output_type
	13, // [13:19] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_cartService_proto_init() }
func file_cartService_proto_init() {
	if File_cartService_proto != nil {
		return
	}
	file_skuInfoService_proto_init()
	file_promotionInfoService_proto_init()
	file_adjustmentService_proto_init()
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
			NumMessages:   6,
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
