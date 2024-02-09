// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: cartService.proto

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

type CartGoods struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CartId           int64                `protobuf:"varint,1,opt,name=cart_id,json=cartId,proto3" json:"cart_id"`                                                                             //ID
	Title            string               `protobuf:"bytes,2,opt,name=title,proto3" json:"title"`                                                                                              //标题
	ImageUrl         string               `protobuf:"bytes,3,opt,name=image_url,json=imageUrl,proto3" json:"image_url"`                                                                        //封面图片
	SpuId            int64                `protobuf:"varint,4,opt,name=spu_id,json=spuId,proto3" json:"spu_id"`                                                                                //商品spuId
	Md5Key           string               `protobuf:"bytes,5,opt,name=md5_key,json=md5Key,proto3" json:"md5_key"`                                                                              //商品规格md5_key
	Inventory        int64                `protobuf:"varint,6,opt,name=inventory,proto3" json:"inventory"`                                                                                     //库存数量
	OriginalPrice    int64                `protobuf:"varint,7,opt,name=original_price,json=originalPrice,proto3" json:"original_price"`                                                        //原价
	Price            int64                `protobuf:"varint,8,opt,name=price,proto3" json:"price"`                                                                                             //销售价格
	SpecBuyMinNumber int64                `protobuf:"varint,9,opt,name=spec_buy_min_number,json=specBuyMinNumber,proto3" json:"spec_buy_min_number"`                                           //购买数量
	SpecBuyMaxNumber int64                `protobuf:"varint,10,opt,name=spec_buy_max_number,json=specBuyMaxNumber,proto3" json:"spec_buy_max_number"`                                          //购买数量
	SpecWeight       int64                `protobuf:"varint,11,opt,name=spec_weight,json=specWeight,proto3" json:"spec_weight"`                                                                //规格重量
	SpecVolume       int64                `protobuf:"varint,12,opt,name=spec_volume,json=specVolume,proto3" json:"spec_volume"`                                                                //规格体积
	SpecCoding       string               `protobuf:"bytes,13,opt,name=spec_coding,json=specCoding,proto3" json:"spec_coding"`                                                                 //规格编码
	SpecBarcode      string               `protobuf:"bytes,14,opt,name=spec_barcode,json=specBarcode,proto3" json:"spec_barcode"`                                                              //规格条码
	Unit             string               `protobuf:"bytes,15,opt,name=unit,proto3" json:"unit"`                                                                                               //库存单位
	Extends          map[string]string    `protobuf:"bytes,16,rep,name=extends,proto3" json:"extends" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` //扩展数据
	IsInvalid        string               `protobuf:"bytes,17,opt,name=is_invalid,json=isInvalid,proto3" json:"is_invalid"`                                                                    //是否无效（0否，1是）
	IsError          string               `protobuf:"bytes,18,opt,name=is_error,json=isError,proto3" json:"is_error"`                                                                          //是否错误（0否，1是）
	ErrorMsg         string               `protobuf:"bytes,19,opt,name=error_msg,json=errorMsg,proto3" json:"error_msg"`                                                                       //错误信息
	Stock            int64                `protobuf:"varint,20,opt,name=stock,proto3" json:"stock"`                                                                                            //购买数量
	SubtotalPrice    int64                `protobuf:"varint,21,opt,name=subtotal_price,json=subtotalPrice,proto3" json:"subtotal_price"`                                                       //商品小计
	Spec             []*SpecificationData `protobuf:"bytes,22,rep,name=spec,proto3" json:"spec"`                                                                                               //规格描述
	SpecDesc         string               `protobuf:"bytes,23,opt,name=spec_desc,json=specDesc,proto3" json:"spec_desc"`
	BrandName        string               `protobuf:"bytes,24,opt,name=brand_name,json=brandName,proto3" json:"brand_name"`
	GoodsUrl         string               `protobuf:"bytes,25,opt,name=goods_url,json=goodsUrl,proto3" json:"goods_url"`
}

func (x *CartGoods) Reset() {
	*x = CartGoods{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cartService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartGoods) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartGoods) ProtoMessage() {}

func (x *CartGoods) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CartGoods.ProtoReflect.Descriptor instead.
func (*CartGoods) Descriptor() ([]byte, []int) {
	return file_cartService_proto_rawDescGZIP(), []int{0}
}

func (x *CartGoods) GetCartId() int64 {
	if x != nil {
		return x.CartId
	}
	return 0
}

func (x *CartGoods) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CartGoods) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *CartGoods) GetSpuId() int64 {
	if x != nil {
		return x.SpuId
	}
	return 0
}

func (x *CartGoods) GetMd5Key() string {
	if x != nil {
		return x.Md5Key
	}
	return ""
}

func (x *CartGoods) GetInventory() int64 {
	if x != nil {
		return x.Inventory
	}
	return 0
}

func (x *CartGoods) GetOriginalPrice() int64 {
	if x != nil {
		return x.OriginalPrice
	}
	return 0
}

func (x *CartGoods) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CartGoods) GetSpecBuyMinNumber() int64 {
	if x != nil {
		return x.SpecBuyMinNumber
	}
	return 0
}

func (x *CartGoods) GetSpecBuyMaxNumber() int64 {
	if x != nil {
		return x.SpecBuyMaxNumber
	}
	return 0
}

func (x *CartGoods) GetSpecWeight() int64 {
	if x != nil {
		return x.SpecWeight
	}
	return 0
}

func (x *CartGoods) GetSpecVolume() int64 {
	if x != nil {
		return x.SpecVolume
	}
	return 0
}

func (x *CartGoods) GetSpecCoding() string {
	if x != nil {
		return x.SpecCoding
	}
	return ""
}

func (x *CartGoods) GetSpecBarcode() string {
	if x != nil {
		return x.SpecBarcode
	}
	return ""
}

func (x *CartGoods) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *CartGoods) GetExtends() map[string]string {
	if x != nil {
		return x.Extends
	}
	return nil
}

func (x *CartGoods) GetIsInvalid() string {
	if x != nil {
		return x.IsInvalid
	}
	return ""
}

func (x *CartGoods) GetIsError() string {
	if x != nil {
		return x.IsError
	}
	return ""
}

func (x *CartGoods) GetErrorMsg() string {
	if x != nil {
		return x.ErrorMsg
	}
	return ""
}

func (x *CartGoods) GetStock() int64 {
	if x != nil {
		return x.Stock
	}
	return 0
}

func (x *CartGoods) GetSubtotalPrice() int64 {
	if x != nil {
		return x.SubtotalPrice
	}
	return 0
}

func (x *CartGoods) GetSpec() []*SpecificationData {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *CartGoods) GetSpecDesc() string {
	if x != nil {
		return x.SpecDesc
	}
	return ""
}

func (x *CartGoods) GetBrandName() string {
	if x != nil {
		return x.BrandName
	}
	return ""
}

func (x *CartGoods) GetGoodsUrl() string {
	if x != nil {
		return x.GoodsUrl
	}
	return ""
}

// 购物车数据
type CartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BuyNumber      int64            `protobuf:"varint,1,opt,name=buy_number,json=buyNumber,proto3" json:"buy_number"`
	TotalPrice     int64            `protobuf:"varint,2,opt,name=total_price,json=totalPrice,proto3" json:"total_price"`
	AddressId      int64            `protobuf:"varint,3,opt,name=address_id,json=addressId,proto3" json:"address_id"`
	CartList       []*CartGoods     `protobuf:"bytes,4,rep,name=cart_list,json=cartList,proto3" json:"cart_list"`
	PromotionsData []*PromotionInfo `protobuf:"bytes,5,rep,name=promotions_data,json=promotionsData,proto3" json:"promotions_data"`
}

func (x *CartResponse) Reset() {
	*x = CartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cartService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartResponse) ProtoMessage() {}

func (x *CartResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CartResponse.ProtoReflect.Descriptor instead.
func (*CartResponse) Descriptor() ([]byte, []int) {
	return file_cartService_proto_rawDescGZIP(), []int{1}
}

func (x *CartResponse) GetBuyNumber() int64 {
	if x != nil {
		return x.BuyNumber
	}
	return 0
}

func (x *CartResponse) GetTotalPrice() int64 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

func (x *CartResponse) GetAddressId() int64 {
	if x != nil {
		return x.AddressId
	}
	return 0
}

func (x *CartResponse) GetCartList() []*CartGoods {
	if x != nil {
		return x.CartList
	}
	return nil
}

func (x *CartResponse) GetPromotionsData() []*PromotionInfo {
	if x != nil {
		return x.PromotionsData
	}
	return nil
}

type CartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type         string               `protobuf:"bytes,1,opt,name=type,proto3" json:"type"`
	Id           int64                `protobuf:"varint,2,opt,name=id,proto3" json:"id"`
	SpuId        int64                `protobuf:"varint,3,opt,name=spu_id,json=spuId,proto3" json:"spu_id"`
	Stock        int64                `protobuf:"varint,4,opt,name=stock,proto3" json:"stock"`
	BuyEventType string               `protobuf:"bytes,5,opt,name=buy_event_type,json=buyEventType,proto3" json:"buy_event_type"`
	Spec         []*SpecificationData `protobuf:"bytes,6,rep,name=spec,proto3" json:"spec"`
	Ids          []int64              `protobuf:"varint,9,rep,packed,name=ids,proto3" json:"ids"`
	MemberId     int64                `protobuf:"varint,10,opt,name=member_id,json=memberId,proto3" json:"member_id"`
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

func (x *CartRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CartRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CartRequest) GetSpuId() int64 {
	if x != nil {
		return x.SpuId
	}
	return 0
}

func (x *CartRequest) GetStock() int64 {
	if x != nil {
		return x.Stock
	}
	return 0
}

func (x *CartRequest) GetBuyEventType() string {
	if x != nil {
		return x.BuyEventType
	}
	return ""
}

func (x *CartRequest) GetSpec() []*SpecificationData {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *CartRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *CartRequest) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

type CartStockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *CartGoods `protobuf:"bytes,1,opt,name=data,proto3" json:"data"`
	Info string     `protobuf:"bytes,2,opt,name=info,proto3" json:"info"`
}

func (x *CartStockResponse) Reset() {
	*x = CartStockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cartService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartStockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartStockResponse) ProtoMessage() {}

func (x *CartStockResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CartStockResponse.ProtoReflect.Descriptor instead.
func (*CartStockResponse) Descriptor() ([]byte, []int) {
	return file_cartService_proto_rawDescGZIP(), []int{3}
}

func (x *CartStockResponse) GetData() *CartGoods {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *CartStockResponse) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

var File_cartService_proto protoreflect.FileDescriptor

var file_cartService_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x61, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x15, 0x62,
	0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf0, 0x06, 0x0a, 0x09, 0x43, 0x61, 0x72, 0x74, 0x47, 0x6f, 0x6f,
	0x64, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x61, 0x72, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x61, 0x72, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x15,
	0x0a, 0x06, 0x73, 0x70, 0x75, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x73, 0x70, 0x75, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x6d, 0x64, 0x35, 0x5f, 0x6b, 0x65, 0x79,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x64, 0x35, 0x4b, 0x65, 0x79, 0x12, 0x1c,
	0x0a, 0x09, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x25, 0x0a, 0x0e,
	0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x2d, 0x0a, 0x13, 0x73, 0x70, 0x65,
	0x63, 0x5f, 0x62, 0x75, 0x79, 0x5f, 0x6d, 0x69, 0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x73, 0x70, 0x65, 0x63, 0x42, 0x75, 0x79, 0x4d,
	0x69, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x13, 0x73, 0x70, 0x65, 0x63,
	0x5f, 0x62, 0x75, 0x79, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x73, 0x70, 0x65, 0x63, 0x42, 0x75, 0x79, 0x4d, 0x61,
	0x78, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x70, 0x65, 0x63, 0x5f,
	0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x70,
	0x65, 0x63, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x70, 0x65, 0x63,
	0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73,
	0x70, 0x65, 0x63, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x70, 0x65,
	0x63, 0x5f, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x73, 0x70, 0x65, 0x63, 0x43, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x70,
	0x65, 0x63, 0x5f, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x73, 0x70, 0x65, 0x63, 0x42, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x6e, 0x69,
	0x74, 0x12, 0x3a, 0x0a, 0x07, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x73, 0x18, 0x10, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61,
	0x72, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x73, 0x12, 0x1d, 0x0a,
	0x0a, 0x69, 0x73, 0x5f, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x69, 0x73, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08,
	0x69, 0x73, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x69, 0x73, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x5f, 0x6d, 0x73, 0x67, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x4d, 0x73, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x75,
	0x62, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x15, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0d, 0x73, 0x75, 0x62, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x2f, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x16, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x70, 0x65, 0x63, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x73, 0x70,
	0x65, 0x63, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x70, 0x65, 0x63, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x18,
	0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x70, 0x65, 0x63, 0x44, 0x65, 0x73, 0x63, 0x12,
	0x1d, 0x0a, 0x0a, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x18, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x19, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x55, 0x72, 0x6c, 0x1a, 0x3a, 0x0a, 0x0c, 0x45,
	0x78, 0x74, 0x65, 0x6e, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xe1, 0x01, 0x0a, 0x0c, 0x43, 0x61, 0x72, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x75, 0x79, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x62, 0x75,
	0x79, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x09, 0x63, 0x61, 0x72, 0x74, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52,
	0x08, 0x63, 0x61, 0x72, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x0f, 0x70, 0x72, 0x6f,
	0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72,
	0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0e, 0x70, 0x72, 0x6f,
	0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x44, 0x61, 0x74, 0x61, 0x22, 0xe4, 0x01, 0x0a, 0x0b,
	0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x15, 0x0a, 0x06, 0x73, 0x70, 0x75, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x73, 0x70, 0x75, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x24, 0x0a, 0x0e,
	0x62, 0x75, 0x79, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x75, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x70, 0x65, 0x63,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x73,
	0x70, 0x65, 0x63, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x03,
	0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x50, 0x0a, 0x11, 0x43, 0x61, 0x72, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x43, 0x61, 0x72, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x69, 0x6e, 0x66, 0x6f, 0x32, 0xa9, 0x03, 0x0a, 0x0b, 0x43, 0x61, 0x72, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x15, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x36,
	0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x06, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x38, 0x0a, 0x05, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x72,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x05, 0x53,
	0x74, 0x6f, 0x63, 0x6b, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x05, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43,
	0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x07, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x64, 0x12,
	0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_cartService_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_cartService_proto_goTypes = []interface{}{
	(*CartGoods)(nil),         // 0: services.CartGoods
	(*CartResponse)(nil),      // 1: services.CartResponse
	(*CartRequest)(nil),       // 2: services.CartRequest
	(*CartStockResponse)(nil), // 3: services.CartStockResponse
	nil,                       // 4: services.CartGoods.ExtendsEntry
	(*SpecificationData)(nil), // 5: services.SpecificationData
	(*PromotionInfo)(nil),     // 6: services.PromotionInfo
}
var file_cartService_proto_depIdxs = []int32{
	4,  // 0: services.CartGoods.extends:type_name -> services.CartGoods.ExtendsEntry
	5,  // 1: services.CartGoods.spec:type_name -> services.SpecificationData
	0,  // 2: services.CartResponse.cart_list:type_name -> services.CartGoods
	6,  // 3: services.CartResponse.promotions_data:type_name -> services.PromotionInfo
	5,  // 4: services.CartRequest.spec:type_name -> services.SpecificationData
	0,  // 5: services.CartStockResponse.data:type_name -> services.CartGoods
	2,  // 6: services.CartService.Index:input_type -> services.CartRequest
	2,  // 7: services.CartService.Add:input_type -> services.CartRequest
	2,  // 8: services.CartService.Remove:input_type -> services.CartRequest
	2,  // 9: services.CartService.Clear:input_type -> services.CartRequest
	2,  // 10: services.CartService.Stock:input_type -> services.CartRequest
	2,  // 11: services.CartService.Count:input_type -> services.CartRequest
	2,  // 12: services.CartService.Checked:input_type -> services.CartRequest
	1,  // 13: services.CartService.Index:output_type -> services.CartResponse
	1,  // 14: services.CartService.Add:output_type -> services.CartResponse
	1,  // 15: services.CartService.Remove:output_type -> services.CartResponse
	1,  // 16: services.CartService.Clear:output_type -> services.CartResponse
	3,  // 17: services.CartService.Stock:output_type -> services.CartStockResponse
	1,  // 18: services.CartService.Count:output_type -> services.CartResponse
	1,  // 19: services.CartService.Checked:output_type -> services.CartResponse
	13, // [13:20] is the sub-list for method output_type
	6,  // [6:13] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_cartService_proto_init() }
func file_cartService_proto_init() {
	if File_cartService_proto != nil {
		return
	}
	file_baseInfoService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cartService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartGoods); i {
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
			switch v := v.(*CartStockResponse); i {
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
			NumMessages:   5,
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
