// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: orderDetailService.proto

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

type OrderDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                    int64          `protobuf:"varint,1,opt,name=id,proto3" json:"id"`                                                                      //ID
	OrderId               int64          `protobuf:"varint,2,opt,name=order_id,json=orderId,proto3" json:"order_id"`                                             //订单ID
	SpuId                 int64          `protobuf:"varint,3,opt,name=spu_id,json=spuId,proto3" json:"spu_id"`                                                   //商品ID
	SkuKey                string         `protobuf:"bytes,4,opt,name=sku_key,json=skuKey,proto3" json:"sku_key"`                                                 //SKUKey
	Name                  string         `protobuf:"bytes,5,opt,name=name,proto3" json:"name"`                                                                   //商品名称
	ImageUrl              string         `protobuf:"bytes,6,opt,name=image_url,json=imageUrl,proto3" json:"image_url"`                                           //图片路径
	SpecWeight            int64          `protobuf:"varint,7,opt,name=spec_weight,json=specWeight,proto3" json:"spec_weight"`                                    //规格重量
	SpecVolume            int64          `protobuf:"varint,8,opt,name=spec_volume,json=specVolume,proto3" json:"spec_volume"`                                    //规格体积
	SpecCoding            string         `protobuf:"bytes,9,opt,name=spec_coding,json=specCoding,proto3" json:"spec_coding"`                                     //规格编码
	SpecBarcode           string         `protobuf:"bytes,10,opt,name=spec_barcode,json=specBarcode,proto3" json:"spec_barcode"`                                 //规格条码
	UnitId                int64          `protobuf:"varint,11,opt,name=unit_id,json=unitId,proto3" json:"unit_id"`                                               //计量单位ID
	UnitName              string         `protobuf:"bytes,12,opt,name=unit_name,json=unitName,proto3" json:"unit_name"`                                          //计量单位名称
	UnitRate              float32        `protobuf:"fixed32,13,opt,name=unit_rate,json=unitRate,proto3" json:"unit_rate"`                                        //单位换算率
	IsWeighed             string         `protobuf:"bytes,14,opt,name=is_weighed,json=isWeighed,proto3" json:"is_weighed"`                                       //称重商品
	BuyNumber             int64          `protobuf:"varint,15,opt,name=buy_number,json=buyNumber,proto3" json:"buy_number"`                                      //购买数量
	Integral              int64          `protobuf:"varint,16,opt,name=integral,proto3" json:"integral"`                                                         //消耗积分
	SalePrice             int64          `protobuf:"varint,17,opt,name=sale_price,json=salePrice,proto3" json:"sale_price"`                                      //销售价
	Price                 int64          `protobuf:"varint,18,opt,name=price,proto3" json:"price"`                                                               //成交价
	DiscountAmount        int64          `protobuf:"varint,19,opt,name=discount_amount,json=discountAmount,proto3" json:"discount_amount"`                       //优惠金额
	SubtotalAmount        int64          `protobuf:"varint,20,opt,name=subtotal_amount,json=subtotalAmount,proto3" json:"subtotal_amount"`                       //合计金额
	RefundedAmount        int64          `protobuf:"varint,21,opt,name=refunded_amount,json=refundedAmount,proto3" json:"refunded_amount"`                       //已退款金额
	ReturnedNumber        int64          `protobuf:"varint,22,opt,name=returned_number,json=returnedNumber,proto3" json:"returned_number"`                       //已退货数量
	DeliveredNumber       int64          `protobuf:"varint,23,opt,name=delivered_number,json=deliveredNumber,proto3" json:"delivered_number"`                    //已发货数量
	IsAllDelivered        string         `protobuf:"bytes,24,opt,name=is_all_delivered,json=isAllDelivered,proto3" json:"is_all_delivered"`                      //已全部发货
	DiscountDetails       string         `protobuf:"bytes,25,opt,name=discount_details,json=discountDetails,proto3" json:"discount_details"`                     //优惠明细
	Spec                  []*SkuSpecInfo `protobuf:"bytes,26,rep,name=spec,proto3" json:"spec"`                                                                  //规格
	SpecDesc              string         `protobuf:"bytes,27,opt,name=spec_desc,json=specDesc,proto3" json:"spec_desc"`                                          //规格描述
	OutPrimaryId          string         `protobuf:"bytes,28,opt,name=out_primary_id,json=outPrimaryId,proto3" json:"out_primary_id"`                            //外部主表主键
	OutSecondaryId        string         `protobuf:"bytes,29,opt,name=out_secondary_id,json=outSecondaryId,proto3" json:"out_secondary_id"`                      //外部次表主键
	IsGift                string         `protobuf:"bytes,30,opt,name=is_gift,json=isGift,proto3" json:"is_gift"`                                                //是否为赠品
	CreatedAt             string         `protobuf:"bytes,31,opt,name=created_at,json=createdAt,proto3" json:"created_at"`                                       //创建时间
	UpdatedAt             string         `protobuf:"bytes,32,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`                                       //修改时间
	OrderaftersaleBtnText string         `protobuf:"bytes,33,opt,name=orderaftersale_btn_text,json=orderaftersaleBtnText,proto3" json:"orderaftersale_btn_text"` //订单售后服务按钮文本
	Orderaftersale        *LastAftersale `protobuf:"bytes,34,opt,name=orderaftersale,proto3" json:"orderaftersale"`                                              //最新一条售后单
}

func (x *OrderDetail) Reset() {
	*x = OrderDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderDetailService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderDetail) ProtoMessage() {}

func (x *OrderDetail) ProtoReflect() protoreflect.Message {
	mi := &file_orderDetailService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderDetail.ProtoReflect.Descriptor instead.
func (*OrderDetail) Descriptor() ([]byte, []int) {
	return file_orderDetailService_proto_rawDescGZIP(), []int{0}
}

func (x *OrderDetail) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OrderDetail) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *OrderDetail) GetSpuId() int64 {
	if x != nil {
		return x.SpuId
	}
	return 0
}

func (x *OrderDetail) GetSkuKey() string {
	if x != nil {
		return x.SkuKey
	}
	return ""
}

func (x *OrderDetail) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OrderDetail) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *OrderDetail) GetSpecWeight() int64 {
	if x != nil {
		return x.SpecWeight
	}
	return 0
}

func (x *OrderDetail) GetSpecVolume() int64 {
	if x != nil {
		return x.SpecVolume
	}
	return 0
}

func (x *OrderDetail) GetSpecCoding() string {
	if x != nil {
		return x.SpecCoding
	}
	return ""
}

func (x *OrderDetail) GetSpecBarcode() string {
	if x != nil {
		return x.SpecBarcode
	}
	return ""
}

func (x *OrderDetail) GetUnitId() int64 {
	if x != nil {
		return x.UnitId
	}
	return 0
}

func (x *OrderDetail) GetUnitName() string {
	if x != nil {
		return x.UnitName
	}
	return ""
}

func (x *OrderDetail) GetUnitRate() float32 {
	if x != nil {
		return x.UnitRate
	}
	return 0
}

func (x *OrderDetail) GetIsWeighed() string {
	if x != nil {
		return x.IsWeighed
	}
	return ""
}

func (x *OrderDetail) GetBuyNumber() int64 {
	if x != nil {
		return x.BuyNumber
	}
	return 0
}

func (x *OrderDetail) GetIntegral() int64 {
	if x != nil {
		return x.Integral
	}
	return 0
}

func (x *OrderDetail) GetSalePrice() int64 {
	if x != nil {
		return x.SalePrice
	}
	return 0
}

func (x *OrderDetail) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *OrderDetail) GetDiscountAmount() int64 {
	if x != nil {
		return x.DiscountAmount
	}
	return 0
}

func (x *OrderDetail) GetSubtotalAmount() int64 {
	if x != nil {
		return x.SubtotalAmount
	}
	return 0
}

func (x *OrderDetail) GetRefundedAmount() int64 {
	if x != nil {
		return x.RefundedAmount
	}
	return 0
}

func (x *OrderDetail) GetReturnedNumber() int64 {
	if x != nil {
		return x.ReturnedNumber
	}
	return 0
}

func (x *OrderDetail) GetDeliveredNumber() int64 {
	if x != nil {
		return x.DeliveredNumber
	}
	return 0
}

func (x *OrderDetail) GetIsAllDelivered() string {
	if x != nil {
		return x.IsAllDelivered
	}
	return ""
}

func (x *OrderDetail) GetDiscountDetails() string {
	if x != nil {
		return x.DiscountDetails
	}
	return ""
}

func (x *OrderDetail) GetSpec() []*SkuSpecInfo {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *OrderDetail) GetSpecDesc() string {
	if x != nil {
		return x.SpecDesc
	}
	return ""
}

func (x *OrderDetail) GetOutPrimaryId() string {
	if x != nil {
		return x.OutPrimaryId
	}
	return ""
}

func (x *OrderDetail) GetOutSecondaryId() string {
	if x != nil {
		return x.OutSecondaryId
	}
	return ""
}

func (x *OrderDetail) GetIsGift() string {
	if x != nil {
		return x.IsGift
	}
	return ""
}

func (x *OrderDetail) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *OrderDetail) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *OrderDetail) GetOrderaftersaleBtnText() string {
	if x != nil {
		return x.OrderaftersaleBtnText
	}
	return ""
}

func (x *OrderDetail) GetOrderaftersale() *LastAftersale {
	if x != nil {
		return x.Orderaftersale
	}
	return nil
}

// 最后条售后
type LastAftersale struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`                                                //ID
	OrderNo        string `protobuf:"bytes,2,opt,name=order_no,json=orderNo,proto3" json:"order_no"`                        //订单编号
	OrderId        int64  `protobuf:"varint,3,opt,name=order_id,json=orderId,proto3" json:"order_id"`                       //订单ID
	OrderDetailId  int64  `protobuf:"varint,4,opt,name=order_detail_id,json=orderDetailId,proto3" json:"order_detail_id"`   //订单详情id
	SpuId          int64  `protobuf:"varint,5,opt,name=spu_id,json=spuId,proto3" json:"spu_id"`                             //SpuId
	SkuKey         string `protobuf:"bytes,6,opt,name=sku_key,json=skuKey,proto3" json:"sku_key"`                           //SkuKey
	MemberId       int64  `protobuf:"varint,7,opt,name=member_id,json=memberId,proto3" json:"member_id"`                    //用户id
	Type           string `protobuf:"bytes,8,opt,name=type,proto3" json:"type"`                                             //业务类型（0仅退款, 1退货退款）
	RefundMethod   string `protobuf:"bytes,9,opt,name=refund_method,json=refundMethod,proto3" json:"refund_method"`         //退款类型（0原路退回, 1退至钱包, 2手动处理）
	Reason         string `protobuf:"bytes,10,opt,name=reason,proto3" json:"reason"`                                        //申请原因
	Number         int32  `protobuf:"varint,11,opt,name=number,proto3" json:"number"`                                       //退货数量
	Price          int64  `protobuf:"varint,12,opt,name=price,proto3" json:"price"`                                         //退款金额
	FreightSubsidy int64  `protobuf:"varint,13,opt,name=freight_subsidy,json=freightSubsidy,proto3" json:"freight_subsidy"` //运费补贴金
	Msg            string `protobuf:"bytes,14,opt,name=msg,proto3" json:"msg"`                                              //退款说明
	Status         string `protobuf:"bytes,15,opt,name=status,proto3" json:"status"`                                        //状态（0待确认, 1待退货, 2待审核, 3已完成, 4已拒绝, 5已取消）
	StatusText     string `protobuf:"bytes,28,opt,name=status_text,json=statusText,proto3" json:"status_text"`
}

func (x *LastAftersale) Reset() {
	*x = LastAftersale{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderDetailService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LastAftersale) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LastAftersale) ProtoMessage() {}

func (x *LastAftersale) ProtoReflect() protoreflect.Message {
	mi := &file_orderDetailService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LastAftersale.ProtoReflect.Descriptor instead.
func (*LastAftersale) Descriptor() ([]byte, []int) {
	return file_orderDetailService_proto_rawDescGZIP(), []int{1}
}

func (x *LastAftersale) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *LastAftersale) GetOrderNo() string {
	if x != nil {
		return x.OrderNo
	}
	return ""
}

func (x *LastAftersale) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *LastAftersale) GetOrderDetailId() int64 {
	if x != nil {
		return x.OrderDetailId
	}
	return 0
}

func (x *LastAftersale) GetSpuId() int64 {
	if x != nil {
		return x.SpuId
	}
	return 0
}

func (x *LastAftersale) GetSkuKey() string {
	if x != nil {
		return x.SkuKey
	}
	return ""
}

func (x *LastAftersale) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

func (x *LastAftersale) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *LastAftersale) GetRefundMethod() string {
	if x != nil {
		return x.RefundMethod
	}
	return ""
}

func (x *LastAftersale) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *LastAftersale) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *LastAftersale) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *LastAftersale) GetFreightSubsidy() int64 {
	if x != nil {
		return x.FreightSubsidy
	}
	return 0
}

func (x *LastAftersale) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *LastAftersale) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *LastAftersale) GetStatusText() string {
	if x != nil {
		return x.StatusText
	}
	return ""
}

var File_orderDetailService_proto protoreflect.FileDescriptor

var file_orderDetailService_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x1a, 0x15, 0x62, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x09, 0x0a, 0x0b,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x70, 0x75, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x70, 0x75, 0x49, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x73, 0x6b, 0x75, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x6b, 0x75, 0x4b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x70, 0x65, 0x63, 0x5f,
	0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x70,
	0x65, 0x63, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x70, 0x65, 0x63,
	0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73,
	0x70, 0x65, 0x63, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x70, 0x65,
	0x63, 0x5f, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x73, 0x70, 0x65, 0x63, 0x43, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x70,
	0x65, 0x63, 0x5f, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x73, 0x70, 0x65, 0x63, 0x42, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x6e, 0x69, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x75, 0x6e, 0x69, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x6e, 0x69, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x6e, 0x69, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x6e, 0x69, 0x74, 0x5f, 0x72, 0x61, 0x74, 0x65,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x75, 0x6e, 0x69, 0x74, 0x52, 0x61, 0x74, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x77, 0x65, 0x69, 0x67, 0x68, 0x65, 0x64, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x73, 0x57, 0x65, 0x69, 0x67, 0x68, 0x65, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x62, 0x75, 0x79, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x62, 0x75, 0x79, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1a,
	0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x18, 0x10, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x61,
	0x6c, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x73, 0x61, 0x6c, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x27, 0x0a, 0x0f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x13, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x75, 0x62, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0e, 0x73, 0x75, 0x62, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x65, 0x64, 0x5f, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x15, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x72, 0x65, 0x66, 0x75,
	0x6e, 0x64, 0x65, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65,
	0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x16, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0e, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x65, 0x64,
	0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x17, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x64,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x65, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x28,
	0x0a, 0x10, 0x69, 0x73, 0x5f, 0x61, 0x6c, 0x6c, 0x5f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x65, 0x64, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x69, 0x73, 0x41, 0x6c, 0x6c, 0x44,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x65, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x19, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x12, 0x29, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x1a, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x6b, 0x75,
	0x53, 0x70, 0x65, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x1b,
	0x0a, 0x09, 0x73, 0x70, 0x65, 0x63, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x18, 0x1b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x73, 0x70, 0x65, 0x63, 0x44, 0x65, 0x73, 0x63, 0x12, 0x24, 0x0a, 0x0e, 0x6f,
	0x75, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x1c, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x75, 0x74, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x49,
	0x64, 0x12, 0x28, 0x0a, 0x10, 0x6f, 0x75, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61,
	0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x75, 0x74,
	0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x69,
	0x73, 0x5f, 0x67, 0x69, 0x66, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x73,
	0x47, 0x69, 0x66, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x20, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x36, 0x0a, 0x17, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x61, 0x66, 0x74, 0x65, 0x72,
	0x73, 0x61, 0x6c, 0x65, 0x5f, 0x62, 0x74, 0x6e, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x21, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x15, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x61, 0x66, 0x74, 0x65, 0x72, 0x73,
	0x61, 0x6c, 0x65, 0x42, 0x74, 0x6e, 0x54, 0x65, 0x78, 0x74, 0x12, 0x3f, 0x0a, 0x0e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x61, 0x66, 0x74, 0x65, 0x72, 0x73, 0x61, 0x6c, 0x65, 0x18, 0x22, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x61,
	0x73, 0x74, 0x41, 0x66, 0x74, 0x65, 0x72, 0x73, 0x61, 0x6c, 0x65, 0x52, 0x0e, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x61, 0x66, 0x74, 0x65, 0x72, 0x73, 0x61, 0x6c, 0x65, 0x22, 0xbd, 0x03, 0x0a, 0x0d,
	0x4c, 0x61, 0x73, 0x74, 0x41, 0x66, 0x74, 0x65, 0x72, 0x73, 0x61, 0x6c, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a,
	0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x6e, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x73,
	0x70, 0x75, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x70, 0x75,
	0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x6b, 0x75, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6b, 0x75, 0x4b, 0x65, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d,
	0x72, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x66, 0x72, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x5f, 0x73, 0x75, 0x62, 0x73, 0x69, 0x64, 0x79, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0e, 0x66, 0x72, 0x65, 0x69, 0x67, 0x68, 0x74, 0x53, 0x75, 0x62, 0x73, 0x69, 0x64, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d,
	0x73, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x54, 0x65, 0x78, 0x74, 0x42, 0x0d, 0x5a, 0x0b, 0x2f,
	0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_orderDetailService_proto_rawDescOnce sync.Once
	file_orderDetailService_proto_rawDescData = file_orderDetailService_proto_rawDesc
)

func file_orderDetailService_proto_rawDescGZIP() []byte {
	file_orderDetailService_proto_rawDescOnce.Do(func() {
		file_orderDetailService_proto_rawDescData = protoimpl.X.CompressGZIP(file_orderDetailService_proto_rawDescData)
	})
	return file_orderDetailService_proto_rawDescData
}

var file_orderDetailService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_orderDetailService_proto_goTypes = []interface{}{
	(*OrderDetail)(nil),   // 0: services.OrderDetail
	(*LastAftersale)(nil), // 1: services.LastAftersale
	(*SkuSpecInfo)(nil),   // 2: services.SkuSpecInfo
}
var file_orderDetailService_proto_depIdxs = []int32{
	2, // 0: services.OrderDetail.spec:type_name -> services.SkuSpecInfo
	1, // 1: services.OrderDetail.orderaftersale:type_name -> services.LastAftersale
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_orderDetailService_proto_init() }
func file_orderDetailService_proto_init() {
	if File_orderDetailService_proto != nil {
		return
	}
	file_baseInfoService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_orderDetailService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderDetail); i {
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
		file_orderDetailService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LastAftersale); i {
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
			RawDescriptor: file_orderDetailService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_orderDetailService_proto_goTypes,
		DependencyIndexes: file_orderDetailService_proto_depIdxs,
		MessageInfos:      file_orderDetailService_proto_msgTypes,
	}.Build()
	File_orderDetailService_proto = out.File
	file_orderDetailService_proto_rawDesc = nil
	file_orderDetailService_proto_goTypes = nil
	file_orderDetailService_proto_depIdxs = nil
}
