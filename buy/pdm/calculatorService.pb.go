// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: calculatorService.proto

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

// 购买分组
type BuyGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Warehouse  *Warehouse      `protobuf:"bytes,1,opt,name=warehouse,proto3" json:"warehouse"`
	OrderBase  *BuyOrderBase   `protobuf:"bytes,2,opt,name=order_base,json=orderBase,proto3" json:"order_base"`    //订单基本信息
	OrderItems []*BuyOrderItem `protobuf:"bytes,3,rep,name=order_items,json=orderItems,proto3" json:"order_items"` //订单商品列表
}

func (x *BuyGroup) Reset() {
	*x = BuyGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculatorService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuyGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuyGroup) ProtoMessage() {}

func (x *BuyGroup) ProtoReflect() protoreflect.Message {
	mi := &file_calculatorService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuyGroup.ProtoReflect.Descriptor instead.
func (*BuyGroup) Descriptor() ([]byte, []int) {
	return file_calculatorService_proto_rawDescGZIP(), []int{0}
}

func (x *BuyGroup) GetWarehouse() *Warehouse {
	if x != nil {
		return x.Warehouse
	}
	return nil
}

func (x *BuyGroup) GetOrderBase() *BuyOrderBase {
	if x != nil {
		return x.OrderBase
	}
	return nil
}

func (x *BuyGroup) GetOrderItems() []*BuyOrderItem {
	if x != nil {
		return x.OrderItems
	}
	return nil
}

type BuyOrderBase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MemberId        int64   `protobuf:"varint,1,opt,name=member_id,json=memberId,proto3" json:"member_id"`                         //买家ID
	RealstoreId     int64   `protobuf:"varint,3,opt,name=realstore_id,json=realstoreId,proto3" json:"realstore_id"`                //多门店ID
	WarehouseId     int64   `protobuf:"varint,4,opt,name=warehouse_id,json=warehouseId,proto3" json:"warehouse_id"`                //仓库ID
	BuyCount        float32 `protobuf:"fixed32,5,opt,name=buy_count,json=buyCount,proto3" json:"buy_count"`                        //购买总数
	BuyCountNumber  int64   `protobuf:"varint,6,opt,name=buy_count_number,json=buyCountNumber,proto3" json:"buy_count_number"`     //购买总件数
	BuyCountWeighed int64   `protobuf:"varint,7,opt,name=buy_count_weighed,json=buyCountWeighed,proto3" json:"buy_count_weighed"`  //购买总称重
	TotalAmount     int64   `protobuf:"varint,8,opt,name=total_amount,json=totalAmount,proto3" json:"total_amount"`                //订单单价
	FinalAmount     int64   `protobuf:"varint,9,opt,name=final_amount,json=finalAmount,proto3" json:"final_amount"`                //订单实付
	IncreaseAmount  int64   `protobuf:"varint,10,opt,name=increase_amount,json=increaseAmount,proto3" json:"increase_amount"`      //增加金额
	DiscountAmount  int64   `protobuf:"varint,11,opt,name=discount_amount,json=discountAmount,proto3" json:"discount_amount"`      //优惠金额
	SpecTotalWeight int64   `protobuf:"varint,12,opt,name=spec_total_weight,json=specTotalWeight,proto3" json:"spec_total_weight"` //规格重量总计
	SpecTotalVolume int64   `protobuf:"varint,13,opt,name=spec_total_volume,json=specTotalVolume,proto3" json:"spec_total_volume"` //规格体积总计
}

func (x *BuyOrderBase) Reset() {
	*x = BuyOrderBase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculatorService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuyOrderBase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuyOrderBase) ProtoMessage() {}

func (x *BuyOrderBase) ProtoReflect() protoreflect.Message {
	mi := &file_calculatorService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuyOrderBase.ProtoReflect.Descriptor instead.
func (*BuyOrderBase) Descriptor() ([]byte, []int) {
	return file_calculatorService_proto_rawDescGZIP(), []int{1}
}

func (x *BuyOrderBase) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

func (x *BuyOrderBase) GetRealstoreId() int64 {
	if x != nil {
		return x.RealstoreId
	}
	return 0
}

func (x *BuyOrderBase) GetWarehouseId() int64 {
	if x != nil {
		return x.WarehouseId
	}
	return 0
}

func (x *BuyOrderBase) GetBuyCount() float32 {
	if x != nil {
		return x.BuyCount
	}
	return 0
}

func (x *BuyOrderBase) GetBuyCountNumber() int64 {
	if x != nil {
		return x.BuyCountNumber
	}
	return 0
}

func (x *BuyOrderBase) GetBuyCountWeighed() int64 {
	if x != nil {
		return x.BuyCountWeighed
	}
	return 0
}

func (x *BuyOrderBase) GetTotalAmount() int64 {
	if x != nil {
		return x.TotalAmount
	}
	return 0
}

func (x *BuyOrderBase) GetFinalAmount() int64 {
	if x != nil {
		return x.FinalAmount
	}
	return 0
}

func (x *BuyOrderBase) GetIncreaseAmount() int64 {
	if x != nil {
		return x.IncreaseAmount
	}
	return 0
}

func (x *BuyOrderBase) GetDiscountAmount() int64 {
	if x != nil {
		return x.DiscountAmount
	}
	return 0
}

func (x *BuyOrderBase) GetSpecTotalWeight() int64 {
	if x != nil {
		return x.SpecTotalWeight
	}
	return 0
}

func (x *BuyOrderBase) GetSpecTotalVolume() int64 {
	if x != nil {
		return x.SpecTotalVolume
	}
	return 0
}

type BuyOrderItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SpuId         int64    `protobuf:"varint,1,opt,name=spu_id,json=spuId,proto3" json:"spu_id"`                         //商品ID
	SkuKey        string   `protobuf:"bytes,2,opt,name=sku_key,json=skuKey,proto3" json:"sku_key"`                       //SKUKey
	SpecWeight    int64    `protobuf:"varint,3,opt,name=spec_weight,json=specWeight,proto3" json:"spec_weight"`          //规格重量
	SpecVolume    int64    `protobuf:"varint,4,opt,name=spec_volume,json=specVolume,proto3" json:"spec_volume"`          //规格体积
	Quantity      int64    `protobuf:"varint,5,opt,name=quantity,proto3" json:"quantity"`                                //购买数量
	UnitPrice     int64    `protobuf:"varint,6,opt,name=unit_price,json=unitPrice,proto3" json:"unit_price"`             //商品单价
	TotalPrice    int64    `protobuf:"varint,7,opt,name=total_price,json=totalPrice,proto3" json:"total_price"`          //合计价格
	DiscountShare int64    `protobuf:"varint,8,opt,name=discount_share,json=discountShare,proto3" json:"discount_share"` //优惠分摊
	FinalPrice    int64    `protobuf:"varint,9,opt,name=final_price,json=finalPrice,proto3" json:"final_price"`          //支付价格
	Product       *Product `protobuf:"bytes,10,opt,name=product,proto3" json:"product"`
}

func (x *BuyOrderItem) Reset() {
	*x = BuyOrderItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculatorService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuyOrderItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuyOrderItem) ProtoMessage() {}

func (x *BuyOrderItem) ProtoReflect() protoreflect.Message {
	mi := &file_calculatorService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuyOrderItem.ProtoReflect.Descriptor instead.
func (*BuyOrderItem) Descriptor() ([]byte, []int) {
	return file_calculatorService_proto_rawDescGZIP(), []int{2}
}

func (x *BuyOrderItem) GetSpuId() int64 {
	if x != nil {
		return x.SpuId
	}
	return 0
}

func (x *BuyOrderItem) GetSkuKey() string {
	if x != nil {
		return x.SkuKey
	}
	return ""
}

func (x *BuyOrderItem) GetSpecWeight() int64 {
	if x != nil {
		return x.SpecWeight
	}
	return 0
}

func (x *BuyOrderItem) GetSpecVolume() int64 {
	if x != nil {
		return x.SpecVolume
	}
	return 0
}

func (x *BuyOrderItem) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *BuyOrderItem) GetUnitPrice() int64 {
	if x != nil {
		return x.UnitPrice
	}
	return 0
}

func (x *BuyOrderItem) GetTotalPrice() int64 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

func (x *BuyOrderItem) GetDiscountShare() int64 {
	if x != nil {
		return x.DiscountShare
	}
	return 0
}

func (x *BuyOrderItem) GetFinalPrice() int64 {
	if x != nil {
		return x.FinalPrice
	}
	return 0
}

func (x *BuyOrderItem) GetProduct() *Product {
	if x != nil {
		return x.Product
	}
	return nil
}

// 购买商品信息
type PurchaseItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SpuId  int64  `protobuf:"varint,1,opt,name=spu_id,json=spuId,proto3" json:"spu_id"`   //商品ID
	SkuKey string `protobuf:"bytes,2,opt,name=sku_key,json=skuKey,proto3" json:"sku_key"` //规格sku_key
	Stock  int64  `protobuf:"varint,3,opt,name=stock,proto3" json:"stock"`                //数量
}

func (x *PurchaseItem) Reset() {
	*x = PurchaseItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculatorService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PurchaseItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PurchaseItem) ProtoMessage() {}

func (x *PurchaseItem) ProtoReflect() protoreflect.Message {
	mi := &file_calculatorService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PurchaseItem.ProtoReflect.Descriptor instead.
func (*PurchaseItem) Descriptor() ([]byte, []int) {
	return file_calculatorService_proto_rawDescGZIP(), []int{3}
}

func (x *PurchaseItem) GetSpuId() int64 {
	if x != nil {
		return x.SpuId
	}
	return 0
}

func (x *PurchaseItem) GetSkuKey() string {
	if x != nil {
		return x.SkuKey
	}
	return ""
}

func (x *PurchaseItem) GetStock() int64 {
	if x != nil {
		return x.Stock
	}
	return 0
}

// 计算请求
type CalculatorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShopId        int64           `protobuf:"varint,1,opt,name=shop_id,json=shopId,proto3" json:"shop_id"`
	RealstoreId   int64           `protobuf:"varint,2,opt,name=realstore_id,json=realstoreId,proto3" json:"realstore_id"`
	PurchaseItems []*PurchaseItem `protobuf:"bytes,3,rep,name=purchase_items,json=purchaseItems,proto3" json:"purchase_items"`
}

func (x *CalculatorRequest) Reset() {
	*x = CalculatorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculatorService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculatorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculatorRequest) ProtoMessage() {}

func (x *CalculatorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculatorService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculatorRequest.ProtoReflect.Descriptor instead.
func (*CalculatorRequest) Descriptor() ([]byte, []int) {
	return file_calculatorService_proto_rawDescGZIP(), []int{4}
}

func (x *CalculatorRequest) GetShopId() int64 {
	if x != nil {
		return x.ShopId
	}
	return 0
}

func (x *CalculatorRequest) GetRealstoreId() int64 {
	if x != nil {
		return x.RealstoreId
	}
	return 0
}

func (x *CalculatorRequest) GetPurchaseItems() []*PurchaseItem {
	if x != nil {
		return x.PurchaseItems
	}
	return nil
}

type CalculatorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg              string      `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg"`
	CalculateBuyData []*BuyGroup `protobuf:"bytes,2,rep,name=calculate_buy_data,json=calculateBuyData,proto3" json:"calculate_buy_data"`
}

func (x *CalculatorResponse) Reset() {
	*x = CalculatorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculatorService_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculatorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculatorResponse) ProtoMessage() {}

func (x *CalculatorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculatorService_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculatorResponse.ProtoReflect.Descriptor instead.
func (*CalculatorResponse) Descriptor() ([]byte, []int) {
	return file_calculatorService_proto_rawDescGZIP(), []int{5}
}

func (x *CalculatorResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *CalculatorResponse) GetCalculateBuyData() []*BuyGroup {
	if x != nil {
		return x.CalculateBuyData
	}
	return nil
}

var File_calculatorService_proto protoreflect.FileDescriptor

var file_calculatorService_proto_rawDesc = []byte{
	0x0a, 0x17, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x1a, 0x16, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xad, 0x01, 0x0a, 0x08, 0x42, 0x75, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x31,
	0x0a, 0x09, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x57, 0x61, 0x72,
	0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x09, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73,
	0x65, 0x12, 0x35, 0x0a, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x42, 0x75, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x61, 0x73, 0x65, 0x52, 0x09, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x42, 0x61, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x75, 0x79, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d,
	0x73, 0x22, 0xd4, 0x03, 0x0a, 0x0c, 0x42, 0x75, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x61,
	0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x21, 0x0a, 0x0c, 0x72, 0x65, 0x61, 0x6c, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x72, 0x65, 0x61, 0x6c, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f,
	0x75, 0x73, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x75, 0x79, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x62, 0x75, 0x79, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x62, 0x75, 0x79, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x62, 0x75,
	0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x11,
	0x62, 0x75, 0x79, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x77, 0x65, 0x69, 0x67, 0x68, 0x65,
	0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x62, 0x75, 0x79, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x57, 0x65, 0x69, 0x67, 0x68, 0x65, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x66,
	0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0b, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x27,
	0x0a, 0x0f, 0x69, 0x6e, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x69, 0x6e, 0x63, 0x72, 0x65, 0x61, 0x73,
	0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x2a, 0x0a, 0x11, 0x73, 0x70, 0x65, 0x63, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x77,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x73, 0x70, 0x65,
	0x63, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x2a, 0x0a, 0x11,
	0x73, 0x70, 0x65, 0x63, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x73, 0x70, 0x65, 0x63, 0x54, 0x6f, 0x74,
	0x61, 0x6c, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x22, 0xd1, 0x02, 0x0a, 0x0c, 0x42, 0x75, 0x79,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x70, 0x75,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x70, 0x75, 0x49, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x73, 0x6b, 0x75, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x6b, 0x75, 0x4b, 0x65, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x70, 0x65,
	0x63, 0x5f, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x73, 0x70, 0x65, 0x63, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x70,
	0x65, 0x63, 0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x73, 0x70, 0x65, 0x63, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x71,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x6e, 0x69, 0x74, 0x5f,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x6e, 0x69,
	0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0d, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x68, 0x61, 0x72, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x2b, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0x54, 0x0a, 0x0c,
	0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x15, 0x0a, 0x06,
	0x73, 0x70, 0x75, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x70,
	0x75, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x6b, 0x75, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6b, 0x75, 0x4b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x74, 0x6f, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x74, 0x6f,
	0x63, 0x6b, 0x22, 0x8e, 0x01, 0x0a, 0x11, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x68, 0x6f, 0x70,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x68, 0x6f, 0x70, 0x49,
	0x64, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x61, 0x6c, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x72, 0x65, 0x61, 0x6c, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x49, 0x64, 0x12, 0x3d, 0x0a, 0x0e, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65,
	0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x0d, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x49, 0x74,
	0x65, 0x6d, 0x73, 0x22, 0x68, 0x0a, 0x12, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x40, 0x0a, 0x12, 0x63,
	0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x62, 0x75, 0x79, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x42, 0x75, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x10, 0x63, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x42, 0x75, 0x79, 0x44, 0x61, 0x74, 0x61, 0x32, 0x64, 0x0a,
	0x11, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x4f, 0x0a, 0x10, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x42,
	0x75, 0x79, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43,
	0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_calculatorService_proto_rawDescOnce sync.Once
	file_calculatorService_proto_rawDescData = file_calculatorService_proto_rawDesc
)

func file_calculatorService_proto_rawDescGZIP() []byte {
	file_calculatorService_proto_rawDescOnce.Do(func() {
		file_calculatorService_proto_rawDescData = protoimpl.X.CompressGZIP(file_calculatorService_proto_rawDescData)
	})
	return file_calculatorService_proto_rawDescData
}

var file_calculatorService_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_calculatorService_proto_goTypes = []interface{}{
	(*BuyGroup)(nil),           // 0: services.BuyGroup
	(*BuyOrderBase)(nil),       // 1: services.BuyOrderBase
	(*BuyOrderItem)(nil),       // 2: services.BuyOrderItem
	(*PurchaseItem)(nil),       // 3: services.PurchaseItem
	(*CalculatorRequest)(nil),  // 4: services.CalculatorRequest
	(*CalculatorResponse)(nil), // 5: services.CalculatorResponse
	(*Warehouse)(nil),          // 6: services.Warehouse
	(*Product)(nil),            // 7: services.Product
}
var file_calculatorService_proto_depIdxs = []int32{
	6, // 0: services.BuyGroup.warehouse:type_name -> services.Warehouse
	1, // 1: services.BuyGroup.order_base:type_name -> services.BuyOrderBase
	2, // 2: services.BuyGroup.order_items:type_name -> services.BuyOrderItem
	7, // 3: services.BuyOrderItem.product:type_name -> services.Product
	3, // 4: services.CalculatorRequest.purchase_items:type_name -> services.PurchaseItem
	0, // 5: services.CalculatorResponse.calculate_buy_data:type_name -> services.BuyGroup
	4, // 6: services.CalculatorService.CalculateBuyData:input_type -> services.CalculatorRequest
	5, // 7: services.CalculatorService.CalculateBuyData:output_type -> services.CalculatorResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_calculatorService_proto_init() }
func file_calculatorService_proto_init() {
	if File_calculatorService_proto != nil {
		return
	}
	file_warehouseService_proto_init()
	file_productService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_calculatorService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuyGroup); i {
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
		file_calculatorService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuyOrderBase); i {
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
		file_calculatorService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuyOrderItem); i {
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
		file_calculatorService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PurchaseItem); i {
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
		file_calculatorService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculatorRequest); i {
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
		file_calculatorService_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculatorResponse); i {
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
			RawDescriptor: file_calculatorService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_calculatorService_proto_goTypes,
		DependencyIndexes: file_calculatorService_proto_depIdxs,
		MessageInfos:      file_calculatorService_proto_msgTypes,
	}.Build()
	File_calculatorService_proto = out.File
	file_calculatorService_proto_rawDesc = nil
	file_calculatorService_proto_goTypes = nil
	file_calculatorService_proto_depIdxs = nil
}
