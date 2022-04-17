// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: skuInfoService.proto

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

type SkuInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int64            `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	SpuId          int64            `protobuf:"varint,2,opt,name=spu_id,json=spuId,proto3" json:"spu_id,omitempty"`
	SkuSn          string           `protobuf:"bytes,3,opt,name=sku_sn,json=skuSn,proto3" json:"sku_sn,omitempty"`
	Title          string           `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Type           string           `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	UnitId         int32            `protobuf:"varint,6,opt,name=unit_id,json=unitId,proto3" json:"unit_id,omitempty"`
	ShopId         int64            `protobuf:"varint,7,opt,name=shop_id,json=shopId,proto3" json:"shop_id,omitempty"`
	ImageId        int64            `protobuf:"varint,8,opt,name=image_id,json=imageId,proto3" json:"image_id,omitempty"`
	ImageUrl       string           `protobuf:"bytes,9,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	Price          float32          `protobuf:"fixed32,10,opt,name=price,proto3" json:"price,omitempty"`
	MarketPrice    float32          `protobuf:"fixed32,11,opt,name=market_price,json=marketPrice,proto3" json:"market_price,omitempty"`
	CostPrice      float32          `protobuf:"fixed32,12,opt,name=cost_price,json=costPrice,proto3" json:"cost_price,omitempty"`
	PromotionPrice float32          `protobuf:"fixed32,13,opt,name=promotion_price,json=promotionPrice,proto3" json:"promotion_price,omitempty"`
	Barcode        string           `protobuf:"bytes,14,opt,name=barcode,proto3" json:"barcode,omitempty"`
	Weight         float32          `protobuf:"fixed32,15,opt,name=weight,proto3" json:"weight,omitempty"`
	Volume         int32            `protobuf:"varint,16,opt,name=volume,proto3" json:"volume,omitempty"`
	Stock          int32            `protobuf:"varint,17,opt,name=stock,proto3" json:"stock,omitempty"`
	FrozenStock    int32            `protobuf:"varint,18,opt,name=frozen_stock,json=frozenStock,proto3" json:"frozen_stock,omitempty"`
	OutSkuNo       string           `protobuf:"bytes,19,opt,name=out_sku_no,json=outSkuNo,proto3" json:"out_sku_no,omitempty"`
	IsSpec         bool             `protobuf:"varint,20,opt,name=is_spec,json=isSpec,proto3" json:"is_spec,omitempty"`
	IsVirtual      bool             `protobuf:"varint,21,opt,name=is_virtual,json=isVirtual,proto3" json:"is_virtual,omitempty"`
	IsPresale      bool             `protobuf:"varint,22,opt,name=is_presale,json=isPresale,proto3" json:"is_presale,omitempty"`
	IsMemberRight  bool             `protobuf:"varint,23,opt,name=is_member_right,json=isMemberRight,proto3" json:"is_member_right,omitempty"`
	IsMemberPrice  bool             `protobuf:"varint,24,opt,name=is_member_price,json=isMemberPrice,proto3" json:"is_member_price,omitempty"`
	IsHideStock    bool             `protobuf:"varint,25,opt,name=is_hide_stock,json=isHideStock,proto3" json:"is_hide_stock,omitempty"`
	IsNew          bool             `protobuf:"varint,27,opt,name=is_new,json=isNew,proto3" json:"is_new,omitempty"`
	IsHot          bool             `protobuf:"varint,28,opt,name=is_hot,json=isHot,proto3" json:"is_hot,omitempty"`
	IsRecommend    bool             `protobuf:"varint,29,opt,name=is_recommend,json=isRecommend,proto3" json:"is_recommend,omitempty"`
	IsSpecial      bool             `protobuf:"varint,30,opt,name=is_special,json=isSpecial,proto3" json:"is_special,omitempty"`
	IsMaster       bool             `protobuf:"varint,31,opt,name=is_master,json=isMaster,proto3" json:"is_master,omitempty"`
	Indexes        string           `protobuf:"bytes,32,opt,name=indexes,proto3" json:"indexes,omitempty"`
	Status         string           `protobuf:"bytes,33,opt,name=status,proto3" json:"status,omitempty"`
	ListedAt       string           `protobuf:"bytes,34,opt,name=listed_at,json=listedAt,proto3" json:"listed_at,omitempty"`
	CreatedAt      string           `protobuf:"bytes,35,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt      string           `protobuf:"bytes,36,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	IndexesData    []int32          `protobuf:"varint,37,rep,packed,name=indexes_data,json=indexesData,proto3" json:"indexes_data,omitempty"`
	OwnDescData    map[int32]string `protobuf:"bytes,38,rep,name=own_desc_data,json=ownDescData,proto3" json:"own_desc_data,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	UnitName       string           `protobuf:"bytes,39,opt,name=unit_name,json=unitName,proto3" json:"unit_name,omitempty"`
	Vip            *SkuVip          `protobuf:"bytes,41,opt,name=vip,proto3" json:"vip,omitempty"`
	Coupon         *SkuCoupon       `protobuf:"bytes,42,opt,name=coupon,proto3" json:"coupon,omitempty"`
	Presale        *SkuPresale      `protobuf:"bytes,43,opt,name=presale,proto3" json:"presale,omitempty"`
}

func (x *SkuInfo) Reset() {
	*x = SkuInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_skuInfoService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SkuInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SkuInfo) ProtoMessage() {}

func (x *SkuInfo) ProtoReflect() protoreflect.Message {
	mi := &file_skuInfoService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SkuInfo.ProtoReflect.Descriptor instead.
func (*SkuInfo) Descriptor() ([]byte, []int) {
	return file_skuInfoService_proto_rawDescGZIP(), []int{0}
}

func (x *SkuInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SkuInfo) GetSpuId() int64 {
	if x != nil {
		return x.SpuId
	}
	return 0
}

func (x *SkuInfo) GetSkuSn() string {
	if x != nil {
		return x.SkuSn
	}
	return ""
}

func (x *SkuInfo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *SkuInfo) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *SkuInfo) GetUnitId() int32 {
	if x != nil {
		return x.UnitId
	}
	return 0
}

func (x *SkuInfo) GetShopId() int64 {
	if x != nil {
		return x.ShopId
	}
	return 0
}

func (x *SkuInfo) GetImageId() int64 {
	if x != nil {
		return x.ImageId
	}
	return 0
}

func (x *SkuInfo) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *SkuInfo) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *SkuInfo) GetMarketPrice() float32 {
	if x != nil {
		return x.MarketPrice
	}
	return 0
}

func (x *SkuInfo) GetCostPrice() float32 {
	if x != nil {
		return x.CostPrice
	}
	return 0
}

func (x *SkuInfo) GetPromotionPrice() float32 {
	if x != nil {
		return x.PromotionPrice
	}
	return 0
}

func (x *SkuInfo) GetBarcode() string {
	if x != nil {
		return x.Barcode
	}
	return ""
}

func (x *SkuInfo) GetWeight() float32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *SkuInfo) GetVolume() int32 {
	if x != nil {
		return x.Volume
	}
	return 0
}

func (x *SkuInfo) GetStock() int32 {
	if x != nil {
		return x.Stock
	}
	return 0
}

func (x *SkuInfo) GetFrozenStock() int32 {
	if x != nil {
		return x.FrozenStock
	}
	return 0
}

func (x *SkuInfo) GetOutSkuNo() string {
	if x != nil {
		return x.OutSkuNo
	}
	return ""
}

func (x *SkuInfo) GetIsSpec() bool {
	if x != nil {
		return x.IsSpec
	}
	return false
}

func (x *SkuInfo) GetIsVirtual() bool {
	if x != nil {
		return x.IsVirtual
	}
	return false
}

func (x *SkuInfo) GetIsPresale() bool {
	if x != nil {
		return x.IsPresale
	}
	return false
}

func (x *SkuInfo) GetIsMemberRight() bool {
	if x != nil {
		return x.IsMemberRight
	}
	return false
}

func (x *SkuInfo) GetIsMemberPrice() bool {
	if x != nil {
		return x.IsMemberPrice
	}
	return false
}

func (x *SkuInfo) GetIsHideStock() bool {
	if x != nil {
		return x.IsHideStock
	}
	return false
}

func (x *SkuInfo) GetIsNew() bool {
	if x != nil {
		return x.IsNew
	}
	return false
}

func (x *SkuInfo) GetIsHot() bool {
	if x != nil {
		return x.IsHot
	}
	return false
}

func (x *SkuInfo) GetIsRecommend() bool {
	if x != nil {
		return x.IsRecommend
	}
	return false
}

func (x *SkuInfo) GetIsSpecial() bool {
	if x != nil {
		return x.IsSpecial
	}
	return false
}

func (x *SkuInfo) GetIsMaster() bool {
	if x != nil {
		return x.IsMaster
	}
	return false
}

func (x *SkuInfo) GetIndexes() string {
	if x != nil {
		return x.Indexes
	}
	return ""
}

func (x *SkuInfo) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *SkuInfo) GetListedAt() string {
	if x != nil {
		return x.ListedAt
	}
	return ""
}

func (x *SkuInfo) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *SkuInfo) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *SkuInfo) GetIndexesData() []int32 {
	if x != nil {
		return x.IndexesData
	}
	return nil
}

func (x *SkuInfo) GetOwnDescData() map[int32]string {
	if x != nil {
		return x.OwnDescData
	}
	return nil
}

func (x *SkuInfo) GetUnitName() string {
	if x != nil {
		return x.UnitName
	}
	return ""
}

func (x *SkuInfo) GetVip() *SkuVip {
	if x != nil {
		return x.Vip
	}
	return nil
}

func (x *SkuInfo) GetCoupon() *SkuCoupon {
	if x != nil {
		return x.Coupon
	}
	return nil
}

func (x *SkuInfo) GetPresale() *SkuPresale {
	if x != nil {
		return x.Presale
	}
	return nil
}

type SkuPresale struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	SkuId            int64   `protobuf:"varint,2,opt,name=sku_id,json=skuId,proto3" json:"sku_id,omitempty"`
	PresaleType      int32   `protobuf:"varint,3,opt,name=presale_type,json=presaleType,proto3" json:"presale_type,omitempty"`
	PayedDate        string  `protobuf:"bytes,4,opt,name=payed_date,json=payedDate,proto3" json:"payed_date,omitempty"`
	AdvanceStartDate string  `protobuf:"bytes,5,opt,name=advance_start_date,json=advanceStartDate,proto3" json:"advance_start_date,omitempty"`
	AdvanceEndDate   string  `protobuf:"bytes,6,opt,name=advance_end_date,json=advanceEndDate,proto3" json:"advance_end_date,omitempty"`
	BalanceStartDate string  `protobuf:"bytes,7,opt,name=balance_start_date,json=balanceStartDate,proto3" json:"balance_start_date,omitempty"`
	BalanceEndDate   string  `protobuf:"bytes,8,opt,name=balance_end_date,json=balanceEndDate,proto3" json:"balance_end_date,omitempty"`
	AdvanceType      string  `protobuf:"bytes,9,opt,name=advance_type,json=advanceType,proto3" json:"advance_type,omitempty"`
	AdvanceRate      int32   `protobuf:"varint,10,opt,name=advance_rate,json=advanceRate,proto3" json:"advance_rate,omitempty"`
	AdvanceMoney     float32 `protobuf:"fixed32,11,opt,name=advance_money,json=advanceMoney,proto3" json:"advance_money,omitempty"`
	BalanceMoney     float32 `protobuf:"fixed32,12,opt,name=balance_money,json=balanceMoney,proto3" json:"balance_money,omitempty"`
	DeliveryType     string  `protobuf:"bytes,13,opt,name=delivery_type,json=deliveryType,proto3" json:"delivery_type,omitempty"`
	DeliveryDate     string  `protobuf:"bytes,14,opt,name=delivery_date,json=deliveryDate,proto3" json:"delivery_date,omitempty"`
	DeliveryDays     int32   `protobuf:"varint,15,opt,name=delivery_days,json=deliveryDays,proto3" json:"delivery_days,omitempty"`
	CreatedAt        string  `protobuf:"bytes,16,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt        string  `protobuf:"bytes,17,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *SkuPresale) Reset() {
	*x = SkuPresale{}
	if protoimpl.UnsafeEnabled {
		mi := &file_skuInfoService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SkuPresale) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SkuPresale) ProtoMessage() {}

func (x *SkuPresale) ProtoReflect() protoreflect.Message {
	mi := &file_skuInfoService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SkuPresale.ProtoReflect.Descriptor instead.
func (*SkuPresale) Descriptor() ([]byte, []int) {
	return file_skuInfoService_proto_rawDescGZIP(), []int{1}
}

func (x *SkuPresale) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SkuPresale) GetSkuId() int64 {
	if x != nil {
		return x.SkuId
	}
	return 0
}

func (x *SkuPresale) GetPresaleType() int32 {
	if x != nil {
		return x.PresaleType
	}
	return 0
}

func (x *SkuPresale) GetPayedDate() string {
	if x != nil {
		return x.PayedDate
	}
	return ""
}

func (x *SkuPresale) GetAdvanceStartDate() string {
	if x != nil {
		return x.AdvanceStartDate
	}
	return ""
}

func (x *SkuPresale) GetAdvanceEndDate() string {
	if x != nil {
		return x.AdvanceEndDate
	}
	return ""
}

func (x *SkuPresale) GetBalanceStartDate() string {
	if x != nil {
		return x.BalanceStartDate
	}
	return ""
}

func (x *SkuPresale) GetBalanceEndDate() string {
	if x != nil {
		return x.BalanceEndDate
	}
	return ""
}

func (x *SkuPresale) GetAdvanceType() string {
	if x != nil {
		return x.AdvanceType
	}
	return ""
}

func (x *SkuPresale) GetAdvanceRate() int32 {
	if x != nil {
		return x.AdvanceRate
	}
	return 0
}

func (x *SkuPresale) GetAdvanceMoney() float32 {
	if x != nil {
		return x.AdvanceMoney
	}
	return 0
}

func (x *SkuPresale) GetBalanceMoney() float32 {
	if x != nil {
		return x.BalanceMoney
	}
	return 0
}

func (x *SkuPresale) GetDeliveryType() string {
	if x != nil {
		return x.DeliveryType
	}
	return ""
}

func (x *SkuPresale) GetDeliveryDate() string {
	if x != nil {
		return x.DeliveryDate
	}
	return ""
}

func (x *SkuPresale) GetDeliveryDays() int32 {
	if x != nil {
		return x.DeliveryDays
	}
	return 0
}

func (x *SkuPresale) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *SkuPresale) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type SkuCoupon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	SkuId       int64  `protobuf:"varint,2,opt,name=sku_id,json=skuId,proto3" json:"sku_id,omitempty"`
	CouponId    int64  `protobuf:"varint,3,opt,name=coupon_id,json=couponId,proto3" json:"coupon_id,omitempty"`
	CouponTitle string `protobuf:"bytes,4,opt,name=coupon_title,json=couponTitle,proto3" json:"coupon_title,omitempty"`
	CreatedAt   string `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   string `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *SkuCoupon) Reset() {
	*x = SkuCoupon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_skuInfoService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SkuCoupon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SkuCoupon) ProtoMessage() {}

func (x *SkuCoupon) ProtoReflect() protoreflect.Message {
	mi := &file_skuInfoService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SkuCoupon.ProtoReflect.Descriptor instead.
func (*SkuCoupon) Descriptor() ([]byte, []int) {
	return file_skuInfoService_proto_rawDescGZIP(), []int{2}
}

func (x *SkuCoupon) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SkuCoupon) GetSkuId() int64 {
	if x != nil {
		return x.SkuId
	}
	return 0
}

func (x *SkuCoupon) GetCouponId() int64 {
	if x != nil {
		return x.CouponId
	}
	return 0
}

func (x *SkuCoupon) GetCouponTitle() string {
	if x != nil {
		return x.CouponTitle
	}
	return ""
}

func (x *SkuCoupon) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *SkuCoupon) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type SkuVip struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	SkuId        int64  `protobuf:"varint,2,opt,name=sku_id,json=skuId,proto3" json:"sku_id,omitempty"`
	MemberRankId int32  `protobuf:"varint,3,opt,name=member_rank_id,json=memberRankId,proto3" json:"member_rank_id,omitempty"`
	KeepType     string `protobuf:"bytes,4,opt,name=keep_type,json=keepType,proto3" json:"keep_type,omitempty"`
	KeepValue    int32  `protobuf:"varint,5,opt,name=keep_value,json=keepValue,proto3" json:"keep_value,omitempty"`
	CreatedAt    string `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt    string `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *SkuVip) Reset() {
	*x = SkuVip{}
	if protoimpl.UnsafeEnabled {
		mi := &file_skuInfoService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SkuVip) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SkuVip) ProtoMessage() {}

func (x *SkuVip) ProtoReflect() protoreflect.Message {
	mi := &file_skuInfoService_proto_msgTypes[3]
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
	return file_skuInfoService_proto_rawDescGZIP(), []int{3}
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

var File_skuInfoService_proto protoreflect.FileDescriptor

var file_skuInfoService_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x6b, 0x75, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x22, 0xab, 0x0a, 0x0a, 0x07, 0x53, 0x6b, 0x75, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x15, 0x0a, 0x06,
	0x73, 0x70, 0x75, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x70,
	0x75, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x6b, 0x75, 0x5f, 0x73, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x6b, 0x75, 0x53, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x6e, 0x69, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x6e, 0x69, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x73, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x5f, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x6d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x73, 0x74, 0x5f,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x63, 0x6f, 0x73,
	0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x0e, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x6f,
	0x63, 0x6b, 0x18, 0x11, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x12,
	0x21, 0x0a, 0x0c, 0x66, 0x72, 0x6f, 0x7a, 0x65, 0x6e, 0x5f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x18,
	0x12, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x66, 0x72, 0x6f, 0x7a, 0x65, 0x6e, 0x53, 0x74, 0x6f,
	0x63, 0x6b, 0x12, 0x1c, 0x0a, 0x0a, 0x6f, 0x75, 0x74, 0x5f, 0x73, 0x6b, 0x75, 0x5f, 0x6e, 0x6f,
	0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x75, 0x74, 0x53, 0x6b, 0x75, 0x4e, 0x6f,
	0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x14, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x69, 0x73, 0x53, 0x70, 0x65, 0x63, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f,
	0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x18, 0x15, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69,
	0x73, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x70,
	0x72, 0x65, 0x73, 0x61, 0x6c, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73,
	0x50, 0x72, 0x65, 0x73, 0x61, 0x6c, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x69, 0x73, 0x5f, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x5f, 0x72, 0x69, 0x67, 0x68, 0x74, 0x18, 0x17, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0d, 0x69, 0x73, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x69, 0x67, 0x68, 0x74, 0x12,
	0x26, 0x0a, 0x0f, 0x69, 0x73, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x18, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x69, 0x73, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x68, 0x69,
	0x64, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x18, 0x19, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b,
	0x69, 0x73, 0x48, 0x69, 0x64, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x15, 0x0a, 0x06, 0x69,
	0x73, 0x5f, 0x6e, 0x65, 0x77, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x4e,
	0x65, 0x77, 0x12, 0x15, 0x0a, 0x06, 0x69, 0x73, 0x5f, 0x68, 0x6f, 0x74, 0x18, 0x1c, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x48, 0x6f, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x73, 0x5f,
	0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0b, 0x69, 0x73, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x69, 0x73, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x09, 0x69, 0x73, 0x53, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x69,
	0x73, 0x5f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x69, 0x73, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x65, 0x73, 0x18, 0x20, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x21, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x69,
	0x73, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x22, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c,
	0x69, 0x73, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x23, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x24, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x25, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0b, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x65, 0x73, 0x44, 0x61, 0x74, 0x61, 0x12, 0x46, 0x0a, 0x0d, 0x6f, 0x77, 0x6e, 0x5f,
	0x64, 0x65, 0x73, 0x63, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x26, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x6b, 0x75, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x4f, 0x77, 0x6e, 0x44, 0x65, 0x73, 0x63, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x0b, 0x6f, 0x77, 0x6e, 0x44, 0x65, 0x73, 0x63, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x1b, 0x0a, 0x09, 0x75, 0x6e, 0x69, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x27, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x6e, 0x69, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a,
	0x03, 0x76, 0x69, 0x70, 0x18, 0x29, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x6b, 0x75, 0x56, 0x69, 0x70, 0x52, 0x03, 0x76, 0x69,
	0x70, 0x12, 0x2b, 0x0a, 0x06, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x18, 0x2a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x6b, 0x75,
	0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x52, 0x06, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x12, 0x2e,
	0x0a, 0x07, 0x70, 0x72, 0x65, 0x73, 0x61, 0x6c, 0x65, 0x18, 0x2b, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x6b, 0x75, 0x50, 0x72,
	0x65, 0x73, 0x61, 0x6c, 0x65, 0x52, 0x07, 0x70, 0x72, 0x65, 0x73, 0x61, 0x6c, 0x65, 0x1a, 0x3e,
	0x0a, 0x10, 0x4f, 0x77, 0x6e, 0x44, 0x65, 0x73, 0x63, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xe2,
	0x04, 0x0a, 0x0a, 0x53, 0x6b, 0x75, 0x50, 0x72, 0x65, 0x73, 0x61, 0x6c, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x15, 0x0a,
	0x06, 0x73, 0x6b, 0x75, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73,
	0x6b, 0x75, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x65, 0x73, 0x61, 0x6c, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x70, 0x72, 0x65, 0x73,
	0x61, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x79, 0x65, 0x64,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x79,
	0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63,
	0x65, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x10, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x5f,
	0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x45, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x2c,
	0x0a, 0x12, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x62, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x28, 0x0a, 0x10,
	0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x45,
	0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63,
	0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x64,
	0x76, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x64, 0x76,
	0x61, 0x6e, 0x63, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0b, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x61, 0x74, 0x65, 0x12, 0x23, 0x0a, 0x0d,
	0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x0c, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x4d, 0x6f, 0x6e, 0x65,
	0x79, 0x12, 0x23, 0x0a, 0x0d, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x6d, 0x6f, 0x6e,
	0x65, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65,
	0x72, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x64,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x44, 0x61, 0x74, 0x65,
	0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x64, 0x61, 0x79,
	0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x79, 0x44, 0x61, 0x79, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0xb0, 0x01, 0x0a, 0x09, 0x53, 0x6b, 0x75, 0x43, 0x6f, 0x75, 0x70, 0x6f,
	0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x6b, 0x75, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x73, 0x6b, 0x75, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x70,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x6f, 0x75,
	0x70, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x5f,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x75,
	0x70, 0x6f, 0x6e, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xcf, 0x01, 0x0a, 0x06, 0x53, 0x6b, 0x75, 0x56, 0x69,
	0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x6b, 0x75, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x73, 0x6b, 0x75, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x5f, 0x72, 0x61, 0x6e, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0c, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x61, 0x6e, 0x6b, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x6b, 0x65, 0x65, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6b, 0x65, 0x65, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6b,
	0x65, 0x65, 0x70, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x6b, 0x65, 0x65, 0x70, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_skuInfoService_proto_rawDescOnce sync.Once
	file_skuInfoService_proto_rawDescData = file_skuInfoService_proto_rawDesc
)

func file_skuInfoService_proto_rawDescGZIP() []byte {
	file_skuInfoService_proto_rawDescOnce.Do(func() {
		file_skuInfoService_proto_rawDescData = protoimpl.X.CompressGZIP(file_skuInfoService_proto_rawDescData)
	})
	return file_skuInfoService_proto_rawDescData
}

var file_skuInfoService_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_skuInfoService_proto_goTypes = []interface{}{
	(*SkuInfo)(nil),    // 0: services.SkuInfo
	(*SkuPresale)(nil), // 1: services.SkuPresale
	(*SkuCoupon)(nil),  // 2: services.SkuCoupon
	(*SkuVip)(nil),     // 3: services.SkuVip
	nil,                // 4: services.SkuInfo.OwnDescDataEntry
}
var file_skuInfoService_proto_depIdxs = []int32{
	4, // 0: services.SkuInfo.own_desc_data:type_name -> services.SkuInfo.OwnDescDataEntry
	3, // 1: services.SkuInfo.vip:type_name -> services.SkuVip
	2, // 2: services.SkuInfo.coupon:type_name -> services.SkuCoupon
	1, // 3: services.SkuInfo.presale:type_name -> services.SkuPresale
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_skuInfoService_proto_init() }
func file_skuInfoService_proto_init() {
	if File_skuInfoService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_skuInfoService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SkuInfo); i {
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
		file_skuInfoService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SkuPresale); i {
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
		file_skuInfoService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SkuCoupon); i {
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
		file_skuInfoService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_skuInfoService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_skuInfoService_proto_goTypes,
		DependencyIndexes: file_skuInfoService_proto_depIdxs,
		MessageInfos:      file_skuInfoService_proto_msgTypes,
	}.Build()
	File_skuInfoService_proto = out.File
	file_skuInfoService_proto_rawDesc = nil
	file_skuInfoService_proto_goTypes = nil
	file_skuInfoService_proto_depIdxs = nil
}
