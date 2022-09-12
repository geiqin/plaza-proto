// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: checkoutService.proto

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

type Checkout struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CheckoutId       string          `protobuf:"bytes,1,opt,name=checkout_id,json=checkoutId,proto3" json:"checkout_id"`
	Currency         string          `protobuf:"bytes,2,opt,name=currency,proto3" json:"currency"`
	TotalNum         int32           `protobuf:"varint,3,opt,name=total_num,json=totalNum,proto3" json:"total_num"`
	TotalWeight      int64           `protobuf:"varint,4,opt,name=total_weight,json=totalWeight,proto3" json:"total_weight"`
	TotalPrice       int64           `protobuf:"varint,5,opt,name=total_price,json=totalPrice,proto3" json:"total_price"`
	TotalDiscount    int64           `protobuf:"varint,6,opt,name=total_discount,json=totalDiscount,proto3" json:"total_discount"`
	TotalAmount      int64           `protobuf:"varint,7,opt,name=total_amount,json=totalAmount,proto3" json:"total_amount"`
	ShippingFee      int64           `protobuf:"varint,8,opt,name=shipping_fee,json=shippingFee,proto3" json:"shipping_fee"`
	InsureFee        int64           `protobuf:"varint,9,opt,name=insure_fee,json=insureFee,proto3" json:"insure_fee"`
	PackFee          int64           `protobuf:"varint,10,opt,name=pack_fee,json=packFee,proto3" json:"pack_fee"`
	MemberId         int64           `protobuf:"varint,11,opt,name=member_id,json=memberId,proto3" json:"member_id"`
	AddressId        int64           `protobuf:"varint,12,opt,name=address_id,json=addressId,proto3" json:"address_id"`
	ShippingAddress  *OrderAddress   `protobuf:"bytes,13,opt,name=shipping_address,json=shippingAddress,proto3" json:"shipping_address"`    //收货地址
	Details          []*CheckoutItem `protobuf:"bytes,14,rep,name=details,proto3" json:"details"`                                           //商品明细
	AvailableCoupons []*CouponInfo   `protobuf:"bytes,15,rep,name=available_coupons,json=availableCoupons,proto3" json:"available_coupons"` //可使用的优惠劵集合
	ShipmentList     []*Shipment     `protobuf:"bytes,16,rep,name=shipment_list,json=shipmentList,proto3" json:"shipment_list"`             //可选的配送方式
	PaymentList      []*Payment      `protobuf:"bytes,17,rep,name=payment_list,json=paymentList,proto3" json:"payment_list"`                //可选的支付方式
}

func (x *Checkout) Reset() {
	*x = Checkout{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkoutService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Checkout) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Checkout) ProtoMessage() {}

func (x *Checkout) ProtoReflect() protoreflect.Message {
	mi := &file_checkoutService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Checkout.ProtoReflect.Descriptor instead.
func (*Checkout) Descriptor() ([]byte, []int) {
	return file_checkoutService_proto_rawDescGZIP(), []int{0}
}

func (x *Checkout) GetCheckoutId() string {
	if x != nil {
		return x.CheckoutId
	}
	return ""
}

func (x *Checkout) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *Checkout) GetTotalNum() int32 {
	if x != nil {
		return x.TotalNum
	}
	return 0
}

func (x *Checkout) GetTotalWeight() int64 {
	if x != nil {
		return x.TotalWeight
	}
	return 0
}

func (x *Checkout) GetTotalPrice() int64 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

func (x *Checkout) GetTotalDiscount() int64 {
	if x != nil {
		return x.TotalDiscount
	}
	return 0
}

func (x *Checkout) GetTotalAmount() int64 {
	if x != nil {
		return x.TotalAmount
	}
	return 0
}

func (x *Checkout) GetShippingFee() int64 {
	if x != nil {
		return x.ShippingFee
	}
	return 0
}

func (x *Checkout) GetInsureFee() int64 {
	if x != nil {
		return x.InsureFee
	}
	return 0
}

func (x *Checkout) GetPackFee() int64 {
	if x != nil {
		return x.PackFee
	}
	return 0
}

func (x *Checkout) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

func (x *Checkout) GetAddressId() int64 {
	if x != nil {
		return x.AddressId
	}
	return 0
}

func (x *Checkout) GetShippingAddress() *OrderAddress {
	if x != nil {
		return x.ShippingAddress
	}
	return nil
}

func (x *Checkout) GetDetails() []*CheckoutItem {
	if x != nil {
		return x.Details
	}
	return nil
}

func (x *Checkout) GetAvailableCoupons() []*CouponInfo {
	if x != nil {
		return x.AvailableCoupons
	}
	return nil
}

func (x *Checkout) GetShipmentList() []*Shipment {
	if x != nil {
		return x.ShipmentList
	}
	return nil
}

func (x *Checkout) GetPaymentList() []*Payment {
	if x != nil {
		return x.PaymentList
	}
	return nil
}

type CheckoutItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SkuId            int64    `protobuf:"varint,1,opt,name=sku_id,json=skuId,proto3" json:"sku_id"`
	Num              int32    `protobuf:"varint,2,opt,name=num,proto3" json:"num"`
	Price            int64    `protobuf:"varint,3,opt,name=price,proto3" json:"price"`
	PromotionPrice   int64    `protobuf:"varint,4,opt,name=promotion_price,json=promotionPrice,proto3" json:"promotion_price"`
	SubtotalAmount   int64    `protobuf:"varint,5,opt,name=subtotal_amount,json=subtotalAmount,proto3" json:"subtotal_amount"`
	SubtotalDiscount int64    `protobuf:"varint,6,opt,name=subtotal_discount,json=subtotalDiscount,proto3" json:"subtotal_discount"`
	Sku              *SkuInfo `protobuf:"bytes,8,opt,name=sku,proto3" json:"sku"`
}

func (x *CheckoutItem) Reset() {
	*x = CheckoutItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkoutService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckoutItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckoutItem) ProtoMessage() {}

func (x *CheckoutItem) ProtoReflect() protoreflect.Message {
	mi := &file_checkoutService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckoutItem.ProtoReflect.Descriptor instead.
func (*CheckoutItem) Descriptor() ([]byte, []int) {
	return file_checkoutService_proto_rawDescGZIP(), []int{1}
}

func (x *CheckoutItem) GetSkuId() int64 {
	if x != nil {
		return x.SkuId
	}
	return 0
}

func (x *CheckoutItem) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *CheckoutItem) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CheckoutItem) GetPromotionPrice() int64 {
	if x != nil {
		return x.PromotionPrice
	}
	return 0
}

func (x *CheckoutItem) GetSubtotalAmount() int64 {
	if x != nil {
		return x.SubtotalAmount
	}
	return 0
}

func (x *CheckoutItem) GetSubtotalDiscount() int64 {
	if x != nil {
		return x.SubtotalDiscount
	}
	return 0
}

func (x *CheckoutItem) GetSku() *SkuInfo {
	if x != nil {
		return x.Sku
	}
	return nil
}

//结账台请求参数
type CheckoutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CheckoutId      string                 `protobuf:"bytes,1,opt,name=checkout_id,json=checkoutId,proto3" json:"checkout_id"`                 //结账台ID
	OrderType       string                 `protobuf:"bytes,2,opt,name=order_type,json=orderType,proto3" json:"order_type"`                    //订单类型（默认为普通订单）
	AddressId       int64                  `protobuf:"varint,3,opt,name=address_id,json=addressId,proto3" json:"address_id"`                   //收货地址ID
	Postscript      string                 `protobuf:"bytes,4,opt,name=postscript,proto3" json:"postscript"`                                   //买家留言(50字以内)
	ShipmentId      int32                  `protobuf:"varint,5,opt,name=shipment_id,json=shipmentId,proto3" json:"shipment_id"`                //选中的配送方式
	PaymentId       int32                  `protobuf:"varint,6,opt,name=payment_id,json=paymentId,proto3" json:"payment_id"`                   //选中的支付方式
	Terminal        string                 `protobuf:"bytes,7,opt,name=terminal,proto3" json:"terminal"`                                       //下单来源终端：site/app/wx_mini/ali_mini/h5/pos
	CouponId        int64                  `protobuf:"varint,8,opt,name=coupon_id,json=couponId,proto3" json:"coupon_id"`                      //选中的优惠劵凭证ID
	Integral        int32                  `protobuf:"varint,9,opt,name=integral,proto3" json:"integral"`                                      //使用的积分数
	Surplus         int64                  `protobuf:"varint,10,opt,name=surplus,proto3" json:"surplus"`                                       //使用的余额数
	AgentMemberId   int64                  `protobuf:"varint,11,opt,name=agent_member_id,json=agentMemberId,proto3" json:"agent_member_id"`    //代理的客户ID（操作员代客下单）
	RechargeMoney   int64                  `protobuf:"varint,12,opt,name=recharge_money,json=rechargeMoney,proto3" json:"recharge_money"`      //订单金额【充值订单】
	RechargeSubject string                 `protobuf:"bytes,13,opt,name=recharge_subject,json=rechargeSubject,proto3" json:"recharge_subject"` //订单标题【充值订单】
	DirectSkuId     int64                  `protobuf:"varint,14,opt,name=direct_sku_id,json=directSkuId,proto3" json:"direct_sku_id"`          //直接购买SkuId
	CartSkuIds      []int64                `protobuf:"varint,15,rep,packed,name=cart_sku_ids,json=cartSkuIds,proto3" json:"cart_sku_ids"`      //购物车选中的SkuIds
	Shops           []*CheckoutRequestShop `protobuf:"bytes,16,rep,name=shops,proto3" json:"shops"`                                            //多店铺数据提交
}

func (x *CheckoutRequest) Reset() {
	*x = CheckoutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkoutService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckoutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckoutRequest) ProtoMessage() {}

func (x *CheckoutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_checkoutService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckoutRequest.ProtoReflect.Descriptor instead.
func (*CheckoutRequest) Descriptor() ([]byte, []int) {
	return file_checkoutService_proto_rawDescGZIP(), []int{2}
}

func (x *CheckoutRequest) GetCheckoutId() string {
	if x != nil {
		return x.CheckoutId
	}
	return ""
}

func (x *CheckoutRequest) GetOrderType() string {
	if x != nil {
		return x.OrderType
	}
	return ""
}

func (x *CheckoutRequest) GetAddressId() int64 {
	if x != nil {
		return x.AddressId
	}
	return 0
}

func (x *CheckoutRequest) GetPostscript() string {
	if x != nil {
		return x.Postscript
	}
	return ""
}

func (x *CheckoutRequest) GetShipmentId() int32 {
	if x != nil {
		return x.ShipmentId
	}
	return 0
}

func (x *CheckoutRequest) GetPaymentId() int32 {
	if x != nil {
		return x.PaymentId
	}
	return 0
}

func (x *CheckoutRequest) GetTerminal() string {
	if x != nil {
		return x.Terminal
	}
	return ""
}

func (x *CheckoutRequest) GetCouponId() int64 {
	if x != nil {
		return x.CouponId
	}
	return 0
}

func (x *CheckoutRequest) GetIntegral() int32 {
	if x != nil {
		return x.Integral
	}
	return 0
}

func (x *CheckoutRequest) GetSurplus() int64 {
	if x != nil {
		return x.Surplus
	}
	return 0
}

func (x *CheckoutRequest) GetAgentMemberId() int64 {
	if x != nil {
		return x.AgentMemberId
	}
	return 0
}

func (x *CheckoutRequest) GetRechargeMoney() int64 {
	if x != nil {
		return x.RechargeMoney
	}
	return 0
}

func (x *CheckoutRequest) GetRechargeSubject() string {
	if x != nil {
		return x.RechargeSubject
	}
	return ""
}

func (x *CheckoutRequest) GetDirectSkuId() int64 {
	if x != nil {
		return x.DirectSkuId
	}
	return 0
}

func (x *CheckoutRequest) GetCartSkuIds() []int64 {
	if x != nil {
		return x.CartSkuIds
	}
	return nil
}

func (x *CheckoutRequest) GetShops() []*CheckoutRequestShop {
	if x != nil {
		return x.Shops
	}
	return nil
}

type CheckoutData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Checkout *Checkout         `protobuf:"bytes,1,opt,name=checkout,proto3" json:"checkout"`                                                                                     //结算信息
	Order    *Order            `protobuf:"bytes,2,opt,name=order,proto3" json:"order"`                                                                                           //订单信息
	Request  *CheckoutRequest  `protobuf:"bytes,3,opt,name=request,proto3" json:"request"`                                                                                       //请求数据原始返回
	Params   map[string]string `protobuf:"bytes,5,rep,name=params,proto3" json:"params" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` //扩展信息
}

func (x *CheckoutData) Reset() {
	*x = CheckoutData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkoutService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckoutData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckoutData) ProtoMessage() {}

func (x *CheckoutData) ProtoReflect() protoreflect.Message {
	mi := &file_checkoutService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckoutData.ProtoReflect.Descriptor instead.
func (*CheckoutData) Descriptor() ([]byte, []int) {
	return file_checkoutService_proto_rawDescGZIP(), []int{3}
}

func (x *CheckoutData) GetCheckout() *Checkout {
	if x != nil {
		return x.Checkout
	}
	return nil
}

func (x *CheckoutData) GetOrder() *Order {
	if x != nil {
		return x.Order
	}
	return nil
}

func (x *CheckoutData) GetRequest() *CheckoutRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *CheckoutData) GetParams() map[string]string {
	if x != nil {
		return x.Params
	}
	return nil
}

type CheckoutResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *CheckoutData `protobuf:"bytes,1,opt,name=data,proto3" json:"data"`
	Error *common.Error `protobuf:"bytes,2,opt,name=error,proto3" json:"error"`
}

func (x *CheckoutResponse) Reset() {
	*x = CheckoutResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkoutService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckoutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckoutResponse) ProtoMessage() {}

func (x *CheckoutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_checkoutService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckoutResponse.ProtoReflect.Descriptor instead.
func (*CheckoutResponse) Descriptor() ([]byte, []int) {
	return file_checkoutService_proto_rawDescGZIP(), []int{4}
}

func (x *CheckoutResponse) GetData() *CheckoutData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *CheckoutResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type CheckoutRequestShop struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShopId     int64  `protobuf:"varint,1,opt,name=shop_id,json=shopId,proto3" json:"shop_id"`             //分店ID
	ShipmentId int32  `protobuf:"varint,2,opt,name=shipment_id,json=shipmentId,proto3" json:"shipment_id"` //选中的配送方式
	Postscript string `protobuf:"bytes,3,opt,name=postscript,proto3" json:"postscript"`                    //商家的留言
}

func (x *CheckoutRequestShop) Reset() {
	*x = CheckoutRequestShop{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkoutService_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckoutRequestShop) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckoutRequestShop) ProtoMessage() {}

func (x *CheckoutRequestShop) ProtoReflect() protoreflect.Message {
	mi := &file_checkoutService_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckoutRequestShop.ProtoReflect.Descriptor instead.
func (*CheckoutRequestShop) Descriptor() ([]byte, []int) {
	return file_checkoutService_proto_rawDescGZIP(), []int{2, 0}
}

func (x *CheckoutRequestShop) GetShopId() int64 {
	if x != nil {
		return x.ShopId
	}
	return 0
}

func (x *CheckoutRequestShop) GetShipmentId() int32 {
	if x != nil {
		return x.ShipmentId
	}
	return 0
}

func (x *CheckoutRequestShop) GetPostscript() string {
	if x != nil {
		return x.Postscript
	}
	return ""
}

var File_checkoutService_proto protoreflect.FileDescriptor

var file_checkoutService_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x15, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x73, 0x6b, 0x75, 0x49, 0x6e,
	0x66, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xb2, 0x05, 0x0a, 0x08, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f,
	0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x5f, 0x66, 0x65, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x73, 0x68, 0x69, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x46, 0x65, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x73, 0x75, 0x72,
	0x65, 0x5f, 0x66, 0x65, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x69, 0x6e, 0x73,
	0x75, 0x72, 0x65, 0x46, 0x65, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x63, 0x6b, 0x5f, 0x66,
	0x65, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x46, 0x65,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x41, 0x0a,
	0x10, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52,
	0x0f, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x30, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x0e, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x6f, 0x75, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x12, 0x41, 0x0a, 0x11, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f,
	0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x10, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x6f,
	0x75, 0x70, 0x6f, 0x6e, 0x73, 0x12, 0x37, 0x0a, 0x0d, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x10, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x0c, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x34,
	0x0a, 0x0c, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x11,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x22, 0xf1, 0x01, 0x0a, 0x0c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75,
	0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x6b, 0x75, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x6b, 0x75, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x6e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x70,
	0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x27, 0x0a,
	0x0f, 0x73, 0x75, 0x62, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x73, 0x75, 0x62, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2b, 0x0a, 0x11, 0x73, 0x75, 0x62, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x10, 0x73, 0x75, 0x62, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x44, 0x69, 0x73, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x03, 0x73, 0x6b, 0x75, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x6b, 0x75, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x03, 0x73, 0x6b, 0x75, 0x22, 0x97, 0x05, 0x0a, 0x0f, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x70,
	0x6f, 0x73, 0x74, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73,
	0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x74,
	0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74,
	0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x70, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x6f, 0x75, 0x70,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x72, 0x70, 0x6c, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x73, 0x75, 0x72, 0x70, 0x6c, 0x75, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0d, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x5f, 0x6d,
	0x6f, 0x6e, 0x65, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x72, 0x65, 0x63, 0x68,
	0x61, 0x72, 0x67, 0x65, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x12, 0x29, 0x0a, 0x10, 0x72, 0x65, 0x63,
	0x68, 0x61, 0x72, 0x67, 0x65, 0x5f, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x53, 0x75, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x22, 0x0a, 0x0d, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x73,
	0x6b, 0x75, 0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x53, 0x6b, 0x75, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0c, 0x63, 0x61, 0x72, 0x74,
	0x5f, 0x73, 0x6b, 0x75, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0a,
	0x63, 0x61, 0x72, 0x74, 0x53, 0x6b, 0x75, 0x49, 0x64, 0x73, 0x12, 0x34, 0x0a, 0x05, 0x73, 0x68,
	0x6f, 0x70, 0x73, 0x18, 0x10, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x52, 0x05, 0x73, 0x68, 0x6f, 0x70, 0x73,
	0x1a, 0x60, 0x0a, 0x04, 0x73, 0x68, 0x6f, 0x70, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x68, 0x6f, 0x70,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x68, 0x6f, 0x70, 0x49,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x22, 0x91, 0x02, 0x0a, 0x0c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x2e, 0x0a, 0x08, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x08, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x6f, 0x75, 0x74, 0x12, 0x25, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x33, 0x0a, 0x07, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x3a, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x6f, 0x75, 0x74, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x63, 0x0a, 0x10, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f,
	0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xdb, 0x01, 0x0a, 0x0f,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x40, 0x0a, 0x05, 0x57, 0x72, 0x69, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x41, 0x0a, 0x06, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x12, 0x19, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x08, 0x52, 0x65, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65,
	0x12, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_checkoutService_proto_rawDescOnce sync.Once
	file_checkoutService_proto_rawDescData = file_checkoutService_proto_rawDesc
)

func file_checkoutService_proto_rawDescGZIP() []byte {
	file_checkoutService_proto_rawDescOnce.Do(func() {
		file_checkoutService_proto_rawDescData = protoimpl.X.CompressGZIP(file_checkoutService_proto_rawDescData)
	})
	return file_checkoutService_proto_rawDescData
}

var file_checkoutService_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_checkoutService_proto_goTypes = []interface{}{
	(*Checkout)(nil),            // 0: services.Checkout
	(*CheckoutItem)(nil),        // 1: services.CheckoutItem
	(*CheckoutRequest)(nil),     // 2: services.CheckoutRequest
	(*CheckoutData)(nil),        // 3: services.CheckoutData
	(*CheckoutResponse)(nil),    // 4: services.CheckoutResponse
	(*CheckoutRequestShop)(nil), // 5: services.CheckoutRequest.shop
	nil,                         // 6: services.CheckoutData.ParamsEntry
	(*OrderAddress)(nil),        // 7: services.OrderAddress
	(*CouponInfo)(nil),          // 8: services.CouponInfo
	(*Shipment)(nil),            // 9: services.Shipment
	(*Payment)(nil),             // 10: services.Payment
	(*SkuInfo)(nil),             // 11: services.SkuInfo
	(*Order)(nil),               // 12: services.Order
	(*common.Error)(nil),        // 13: common.Error
}
var file_checkoutService_proto_depIdxs = []int32{
	7,  // 0: services.Checkout.shipping_address:type_name -> services.OrderAddress
	1,  // 1: services.Checkout.details:type_name -> services.CheckoutItem
	8,  // 2: services.Checkout.available_coupons:type_name -> services.CouponInfo
	9,  // 3: services.Checkout.shipment_list:type_name -> services.Shipment
	10, // 4: services.Checkout.payment_list:type_name -> services.Payment
	11, // 5: services.CheckoutItem.sku:type_name -> services.SkuInfo
	5,  // 6: services.CheckoutRequest.shops:type_name -> services.CheckoutRequest.shop
	0,  // 7: services.CheckoutData.checkout:type_name -> services.Checkout
	12, // 8: services.CheckoutData.order:type_name -> services.Order
	2,  // 9: services.CheckoutData.request:type_name -> services.CheckoutRequest
	6,  // 10: services.CheckoutData.params:type_name -> services.CheckoutData.ParamsEntry
	3,  // 11: services.CheckoutResponse.data:type_name -> services.CheckoutData
	13, // 12: services.CheckoutResponse.error:type_name -> common.Error
	2,  // 13: services.CheckoutService.Write:input_type -> services.CheckoutRequest
	2,  // 14: services.CheckoutService.Submit:input_type -> services.CheckoutRequest
	2,  // 15: services.CheckoutService.Recharge:input_type -> services.CheckoutRequest
	4,  // 16: services.CheckoutService.Write:output_type -> services.CheckoutResponse
	4,  // 17: services.CheckoutService.Submit:output_type -> services.CheckoutResponse
	4,  // 18: services.CheckoutService.Recharge:output_type -> services.CheckoutResponse
	16, // [16:19] is the sub-list for method output_type
	13, // [13:16] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_checkoutService_proto_init() }
func file_checkoutService_proto_init() {
	if File_checkoutService_proto != nil {
		return
	}
	file_orderService_proto_init()
	file_orderAddressService_proto_init()
	file_couponInfoService_proto_init()
	file_paymentService_proto_init()
	file_shipmentService_proto_init()
	file_skuInfoService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_checkoutService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Checkout); i {
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
		file_checkoutService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckoutItem); i {
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
		file_checkoutService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckoutRequest); i {
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
		file_checkoutService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckoutData); i {
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
		file_checkoutService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckoutResponse); i {
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
		file_checkoutService_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckoutRequestShop); i {
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
			RawDescriptor: file_checkoutService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_checkoutService_proto_goTypes,
		DependencyIndexes: file_checkoutService_proto_depIdxs,
		MessageInfos:      file_checkoutService_proto_msgTypes,
	}.Build()
	File_checkoutService_proto = out.File
	file_checkoutService_proto_rawDesc = nil
	file_checkoutService_proto_goTypes = nil
	file_checkoutService_proto_depIdxs = nil
}
