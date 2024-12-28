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

type CalculatorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AreaId         int64          `protobuf:"varint,1,opt,name=area_id,json=areaId,proto3" json:"area_id"`
	Longitude      string         `protobuf:"bytes,2,opt,name=longitude,proto3" json:"longitude"`
	Latitude       string         `protobuf:"bytes,3,opt,name=latitude,proto3" json:"latitude"`
	TotalNum       int32          `protobuf:"varint,4,opt,name=total_num,json=totalNum,proto3" json:"total_num"`
	TotalWeight    int64          `protobuf:"varint,5,opt,name=total_weight,json=totalWeight,proto3" json:"total_weight"`
	ChargingMethod int32          `protobuf:"varint,6,opt,name=charging_method,json=chargingMethod,proto3" json:"charging_method"`
	PurchaseSkus   []*PurchaseSku `protobuf:"bytes,7,rep,name=purchase_skus,json=purchaseSkus,proto3" json:"purchase_skus"`
}

func (x *CalculatorRequest) Reset() {
	*x = CalculatorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculatorService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculatorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculatorRequest) ProtoMessage() {}

func (x *CalculatorRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CalculatorRequest.ProtoReflect.Descriptor instead.
func (*CalculatorRequest) Descriptor() ([]byte, []int) {
	return file_calculatorService_proto_rawDescGZIP(), []int{0}
}

func (x *CalculatorRequest) GetAreaId() int64 {
	if x != nil {
		return x.AreaId
	}
	return 0
}

func (x *CalculatorRequest) GetLongitude() string {
	if x != nil {
		return x.Longitude
	}
	return ""
}

func (x *CalculatorRequest) GetLatitude() string {
	if x != nil {
		return x.Latitude
	}
	return ""
}

func (x *CalculatorRequest) GetTotalNum() int32 {
	if x != nil {
		return x.TotalNum
	}
	return 0
}

func (x *CalculatorRequest) GetTotalWeight() int64 {
	if x != nil {
		return x.TotalWeight
	}
	return 0
}

func (x *CalculatorRequest) GetChargingMethod() int32 {
	if x != nil {
		return x.ChargingMethod
	}
	return 0
}

func (x *CalculatorRequest) GetPurchaseSkus() []*PurchaseSku {
	if x != nil {
		return x.PurchaseSkus
	}
	return nil
}

type PurchaseSku struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SkuId             int64  `protobuf:"varint,1,opt,name=sku_id,json=skuId,proto3" json:"sku_id"`                                       //skuId
	Num               int32  `protobuf:"varint,2,opt,name=num,proto3" json:"num"`                                                        //数量
	Weight            int64  `protobuf:"varint,3,opt,name=weight,proto3" json:"weight"`                                                  //重量
	FreightType       string `protobuf:"bytes,4,opt,name=freight_type,json=freightType,proto3" json:"freight_type"`                      //运费类型
	FreightUniformFee int64  `protobuf:"varint,5,opt,name=freight_uniform_fee,json=freightUniformFee,proto3" json:"freight_uniform_fee"` //统一运费
	FreightExpressId  int64  `protobuf:"varint,6,opt,name=freight_express_id,json=freightExpressId,proto3" json:"freight_express_id"`    //运费模板ID
}

func (x *PurchaseSku) Reset() {
	*x = PurchaseSku{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculatorService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PurchaseSku) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PurchaseSku) ProtoMessage() {}

func (x *PurchaseSku) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use PurchaseSku.ProtoReflect.Descriptor instead.
func (*PurchaseSku) Descriptor() ([]byte, []int) {
	return file_calculatorService_proto_rawDescGZIP(), []int{1}
}

func (x *PurchaseSku) GetSkuId() int64 {
	if x != nil {
		return x.SkuId
	}
	return 0
}

func (x *PurchaseSku) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *PurchaseSku) GetWeight() int64 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *PurchaseSku) GetFreightType() string {
	if x != nil {
		return x.FreightType
	}
	return ""
}

func (x *PurchaseSku) GetFreightUniformFee() int64 {
	if x != nil {
		return x.FreightUniformFee
	}
	return 0
}

func (x *PurchaseSku) GetFreightExpressId() int64 {
	if x != nil {
		return x.FreightExpressId
	}
	return 0
}

type AdjustmentInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SkuId      int64  `protobuf:"varint,1,opt,name=sku_id,json=skuId,proto3" json:"sku_id"`
	Type       string `protobuf:"bytes,2,opt,name=type,proto3" json:"type"`
	Label      string `protobuf:"bytes,3,opt,name=label,proto3" json:"label"`
	OriginCode string `protobuf:"bytes,4,opt,name=origin_code,json=originCode,proto3" json:"origin_code"`
	Amount     int64  `protobuf:"varint,5,opt,name=amount,proto3" json:"amount"`
	Included   bool   `protobuf:"varint,6,opt,name=included,proto3" json:"included"`
}

func (x *AdjustmentInfo) Reset() {
	*x = AdjustmentInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculatorService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdjustmentInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdjustmentInfo) ProtoMessage() {}

func (x *AdjustmentInfo) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use AdjustmentInfo.ProtoReflect.Descriptor instead.
func (*AdjustmentInfo) Descriptor() ([]byte, []int) {
	return file_calculatorService_proto_rawDescGZIP(), []int{2}
}

func (x *AdjustmentInfo) GetSkuId() int64 {
	if x != nil {
		return x.SkuId
	}
	return 0
}

func (x *AdjustmentInfo) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *AdjustmentInfo) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *AdjustmentInfo) GetOriginCode() string {
	if x != nil {
		return x.OriginCode
	}
	return ""
}

func (x *AdjustmentInfo) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *AdjustmentInfo) GetIncluded() bool {
	if x != nil {
		return x.Included
	}
	return false
}

type ExpressResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExpressFee  int64             `protobuf:"varint,1,opt,name=express_fee,json=expressFee,proto3" json:"express_fee"`
	Adjustments []*AdjustmentInfo `protobuf:"bytes,2,rep,name=adjustments,proto3" json:"adjustments"`
}

func (x *ExpressResult) Reset() {
	*x = ExpressResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculatorService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpressResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpressResult) ProtoMessage() {}

func (x *ExpressResult) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ExpressResult.ProtoReflect.Descriptor instead.
func (*ExpressResult) Descriptor() ([]byte, []int) {
	return file_calculatorService_proto_rawDescGZIP(), []int{3}
}

func (x *ExpressResult) GetExpressFee() int64 {
	if x != nil {
		return x.ExpressFee
	}
	return 0
}

func (x *ExpressResult) GetAdjustments() []*AdjustmentInfo {
	if x != nil {
		return x.Adjustments
	}
	return nil
}

type DeliveryResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeliveryFee int64             `protobuf:"varint,1,opt,name=delivery_fee,json=deliveryFee,proto3" json:"delivery_fee"`
	Adjustments []*AdjustmentInfo `protobuf:"bytes,2,rep,name=adjustments,proto3" json:"adjustments"`
}

func (x *DeliveryResult) Reset() {
	*x = DeliveryResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculatorService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeliveryResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliveryResult) ProtoMessage() {}

func (x *DeliveryResult) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DeliveryResult.ProtoReflect.Descriptor instead.
func (*DeliveryResult) Descriptor() ([]byte, []int) {
	return file_calculatorService_proto_rawDescGZIP(), []int{4}
}

func (x *DeliveryResult) GetDeliveryFee() int64 {
	if x != nil {
		return x.DeliveryFee
	}
	return 0
}

func (x *DeliveryResult) GetAdjustments() []*AdjustmentInfo {
	if x != nil {
		return x.Adjustments
	}
	return nil
}

type CalculatorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExpressResult  *ExpressResult  `protobuf:"bytes,1,opt,name=express_result,json=expressResult,proto3" json:"express_result"`
	DeliveryResult *DeliveryResult `protobuf:"bytes,2,opt,name=delivery_result,json=deliveryResult,proto3" json:"delivery_result"`
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

func (x *CalculatorResponse) GetExpressResult() *ExpressResult {
	if x != nil {
		return x.ExpressResult
	}
	return nil
}

func (x *CalculatorResponse) GetDeliveryResult() *DeliveryResult {
	if x != nil {
		return x.DeliveryResult
	}
	return nil
}

var File_calculatorService_proto protoreflect.FileDescriptor

var file_calculatorService_proto_rawDesc = []byte{
	0x0a, 0x17, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x22, 0x8b, 0x02, 0x0a, 0x11, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x72, 0x65,
	0x61, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x72, 0x65, 0x61,
	0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x5f, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x27, 0x0a, 0x0f,
	0x63, 0x68, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x63, 0x68, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x67, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x3a, 0x0a, 0x0d, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73,
	0x65, 0x5f, 0x73, 0x6b, 0x75, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65,
	0x53, 0x6b, 0x75, 0x52, 0x0c, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x53, 0x6b, 0x75,
	0x73, 0x22, 0xcf, 0x01, 0x0a, 0x0b, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x53, 0x6b,
	0x75, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x6b, 0x75, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x73, 0x6b, 0x75, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x72, 0x65, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x72, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x66, 0x72, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x5f, 0x75, 0x6e, 0x69, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x66, 0x65, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x11, 0x66, 0x72, 0x65, 0x69, 0x67, 0x68, 0x74, 0x55, 0x6e, 0x69, 0x66, 0x6f,
	0x72, 0x6d, 0x46, 0x65, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x66, 0x72, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x5f, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x10, 0x66, 0x72, 0x65, 0x69, 0x67, 0x68, 0x74, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x49, 0x64, 0x22, 0xa6, 0x01, 0x0a, 0x0e, 0x41, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x6b, 0x75, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x6b, 0x75, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x72, 0x69, 0x67, 0x69,
	0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x22, 0x6c, 0x0a, 0x0d,
	0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1f, 0x0a,
	0x0b, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x66, 0x65, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x46, 0x65, 0x65, 0x12, 0x3a,
	0x0a, 0x0b, 0x61, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41,
	0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x61,
	0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x6f, 0x0a, 0x0e, 0x44, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x21, 0x0a, 0x0c,
	0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x66, 0x65, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0b, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x46, 0x65, 0x65, 0x12,
	0x3a, 0x0a, 0x0b, 0x61, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x41, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b,
	0x61, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x97, 0x01, 0x0a, 0x12,
	0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0e, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x52, 0x0d, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x41, 0x0a, 0x0f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x0e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0xaa, 0x01, 0x0a, 0x11, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c,
	0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x0a, 0x45,
	0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x46, 0x65, 0x65, 0x12, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65,
	0x72, 0x79, 0x46, 0x65, 0x65, 0x12, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61,
	0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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
	(*CalculatorRequest)(nil),  // 0: services.CalculatorRequest
	(*PurchaseSku)(nil),        // 1: services.PurchaseSku
	(*AdjustmentInfo)(nil),     // 2: services.AdjustmentInfo
	(*ExpressResult)(nil),      // 3: services.ExpressResult
	(*DeliveryResult)(nil),     // 4: services.DeliveryResult
	(*CalculatorResponse)(nil), // 5: services.CalculatorResponse
}
var file_calculatorService_proto_depIdxs = []int32{
	1, // 0: services.CalculatorRequest.purchase_skus:type_name -> services.PurchaseSku
	2, // 1: services.ExpressResult.adjustments:type_name -> services.AdjustmentInfo
	2, // 2: services.DeliveryResult.adjustments:type_name -> services.AdjustmentInfo
	3, // 3: services.CalculatorResponse.express_result:type_name -> services.ExpressResult
	4, // 4: services.CalculatorResponse.delivery_result:type_name -> services.DeliveryResult
	0, // 5: services.CalculatorService.ExpressFee:input_type -> services.CalculatorRequest
	0, // 6: services.CalculatorService.DeliveryFee:input_type -> services.CalculatorRequest
	5, // 7: services.CalculatorService.ExpressFee:output_type -> services.CalculatorResponse
	5, // 8: services.CalculatorService.DeliveryFee:output_type -> services.CalculatorResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_calculatorService_proto_init() }
func file_calculatorService_proto_init() {
	if File_calculatorService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_calculatorService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_calculatorService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PurchaseSku); i {
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
			switch v := v.(*AdjustmentInfo); i {
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
			switch v := v.(*ExpressResult); i {
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
			switch v := v.(*DeliveryResult); i {
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
