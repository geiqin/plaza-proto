// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: orderPayLogService.proto

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

type OrderPayLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	OrderId          int64  `protobuf:"varint,2,opt,name=order_id,json=orderId,proto3" json:"order_id"`
	Type             string `protobuf:"bytes,3,opt,name=type,proto3" json:"type"`
	Currency         int64  `protobuf:"varint,5,opt,name=currency,proto3" json:"currency"`
	Amount           int64  `protobuf:"varint,6,opt,name=amount,proto3" json:"amount"`
	TradeUserType    string `protobuf:"bytes,7,opt,name=trade_user_type,json=tradeUserType,proto3" json:"trade_user_type"`
	TradeUserId      int64  `protobuf:"varint,8,opt,name=trade_user_id,json=tradeUserId,proto3" json:"trade_user_id"`
	PayAccountId     int32  `protobuf:"varint,9,opt,name=pay_account_id,json=payAccountId,proto3" json:"pay_account_id"`
	TrxTransactionId string `protobuf:"bytes,10,opt,name=trx_transaction_id,json=trxTransactionId,proto3" json:"trx_transaction_id"`
	ChannelType      string `protobuf:"bytes,11,opt,name=channel_type,json=channelType,proto3" json:"channel_type"`
	ChargeId         int64  `protobuf:"varint,12,opt,name=charge_id,json=chargeId,proto3" json:"charge_id"`
	OriginCode       string `protobuf:"bytes,13,opt,name=origin_code,json=originCode,proto3" json:"origin_code"`
}

func (x *OrderPayLog) Reset() {
	*x = OrderPayLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderPayLogService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderPayLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderPayLog) ProtoMessage() {}

func (x *OrderPayLog) ProtoReflect() protoreflect.Message {
	mi := &file_orderPayLogService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderPayLog.ProtoReflect.Descriptor instead.
func (*OrderPayLog) Descriptor() ([]byte, []int) {
	return file_orderPayLogService_proto_rawDescGZIP(), []int{0}
}

func (x *OrderPayLog) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OrderPayLog) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *OrderPayLog) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *OrderPayLog) GetCurrency() int64 {
	if x != nil {
		return x.Currency
	}
	return 0
}

func (x *OrderPayLog) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *OrderPayLog) GetTradeUserType() string {
	if x != nil {
		return x.TradeUserType
	}
	return ""
}

func (x *OrderPayLog) GetTradeUserId() int64 {
	if x != nil {
		return x.TradeUserId
	}
	return 0
}

func (x *OrderPayLog) GetPayAccountId() int32 {
	if x != nil {
		return x.PayAccountId
	}
	return 0
}

func (x *OrderPayLog) GetTrxTransactionId() string {
	if x != nil {
		return x.TrxTransactionId
	}
	return ""
}

func (x *OrderPayLog) GetChannelType() string {
	if x != nil {
		return x.ChannelType
	}
	return ""
}

func (x *OrderPayLog) GetChargeId() int64 {
	if x != nil {
		return x.ChargeId
	}
	return 0
}

func (x *OrderPayLog) GetOriginCode() string {
	if x != nil {
		return x.OriginCode
	}
	return ""
}

//支付请求参数
type OrderPayLogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int32  `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize int32  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Sorting  string `protobuf:"bytes,3,opt,name=sorting,proto3" json:"sorting"`
	Keywords string `protobuf:"bytes,4,opt,name=keywords,proto3" json:"keywords"`
	//-----
	Id          int32   `protobuf:"varint,5,opt,name=id,proto3" json:"id"`
	OrderId     int64   `protobuf:"varint,6,opt,name=order_id,json=orderId,proto3" json:"order_id"`
	ChannelType string  `protobuf:"bytes,7,opt,name=channel_type,json=channelType,proto3" json:"channel_type"`
	Type        string  `protobuf:"bytes,8,opt,name=type,proto3" json:"type"`
	ChargeId    int64   `protobuf:"varint,9,opt,name=charge_id,json=chargeId,proto3" json:"charge_id"`
	OriginCode  string  `protobuf:"bytes,10,opt,name=origin_code,json=originCode,proto3" json:"origin_code"`
	Ids         []int64 `protobuf:"varint,11,rep,packed,name=ids,proto3" json:"ids"`
}

func (x *OrderPayLogRequest) Reset() {
	*x = OrderPayLogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderPayLogService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderPayLogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderPayLogRequest) ProtoMessage() {}

func (x *OrderPayLogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orderPayLogService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderPayLogRequest.ProtoReflect.Descriptor instead.
func (*OrderPayLogRequest) Descriptor() ([]byte, []int) {
	return file_orderPayLogService_proto_rawDescGZIP(), []int{1}
}

func (x *OrderPayLogRequest) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *OrderPayLogRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *OrderPayLogRequest) GetSorting() string {
	if x != nil {
		return x.Sorting
	}
	return ""
}

func (x *OrderPayLogRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *OrderPayLogRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OrderPayLogRequest) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *OrderPayLogRequest) GetChannelType() string {
	if x != nil {
		return x.ChannelType
	}
	return ""
}

func (x *OrderPayLogRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *OrderPayLogRequest) GetChargeId() int64 {
	if x != nil {
		return x.ChargeId
	}
	return 0
}

func (x *OrderPayLogRequest) GetOriginCode() string {
	if x != nil {
		return x.OriginCode
	}
	return ""
}

func (x *OrderPayLogRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type OrderPayLogData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *OrderPayLog   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Items  []*OrderPayLog `protobuf:"bytes,2,rep,name=items,proto3" json:"items"`
	Pager  *common.Pager  `protobuf:"bytes,3,opt,name=pager,proto3" json:"pager"`
	Info   *common.Info   `protobuf:"bytes,4,opt,name=info,proto3" json:"info"`
}

func (x *OrderPayLogData) Reset() {
	*x = OrderPayLogData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderPayLogService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderPayLogData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderPayLogData) ProtoMessage() {}

func (x *OrderPayLogData) ProtoReflect() protoreflect.Message {
	mi := &file_orderPayLogService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderPayLogData.ProtoReflect.Descriptor instead.
func (*OrderPayLogData) Descriptor() ([]byte, []int) {
	return file_orderPayLogService_proto_rawDescGZIP(), []int{2}
}

func (x *OrderPayLogData) GetEntity() *OrderPayLog {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *OrderPayLogData) GetItems() []*OrderPayLog {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *OrderPayLogData) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *OrderPayLogData) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type OrderPayLogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *OrderPayLogData `protobuf:"bytes,1,opt,name=data,proto3" json:"data"`
	Error *common.Error    `protobuf:"bytes,2,opt,name=error,proto3" json:"error"`
}

func (x *OrderPayLogResponse) Reset() {
	*x = OrderPayLogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderPayLogService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderPayLogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderPayLogResponse) ProtoMessage() {}

func (x *OrderPayLogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orderPayLogService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderPayLogResponse.ProtoReflect.Descriptor instead.
func (*OrderPayLogResponse) Descriptor() ([]byte, []int) {
	return file_orderPayLogService_proto_rawDescGZIP(), []int{3}
}

func (x *OrderPayLogResponse) GetData() *OrderPayLogData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *OrderPayLogResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_orderPayLogService_proto protoreflect.FileDescriptor

var file_orderPayLogService_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x61, 0x79, 0x4c, 0x6f, 0x67, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x81, 0x03, 0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x50, 0x61, 0x79, 0x4c, 0x6f, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x74, 0x72,
	0x61, 0x64, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x64, 0x65, 0x55, 0x73, 0x65, 0x72, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x64, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x70, 0x61, 0x79, 0x5f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c,
	0x70, 0x61, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x12,
	0x74, 0x72, 0x78, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x74, 0x72, 0x78, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x22, 0xaf, 0x02, 0x0a, 0x12,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x61, 0x79, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x12,
	0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x69,
	0x64, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0xb4, 0x01,
	0x0a, 0x0f, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x61, 0x79, 0x4c, 0x6f, 0x67, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x2d, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x50, 0x61, 0x79, 0x4c, 0x6f, 0x67, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x12, 0x2b, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x50, 0x61, 0x79, 0x4c, 0x6f, 0x67, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x23, 0x0a,
	0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67,
	0x65, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04,
	0x69, 0x6e, 0x66, 0x6f, 0x22, 0x69, 0x0a, 0x13, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x61, 0x79,
	0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x61, 0x79, 0x4c, 0x6f, 0x67,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32,
	0x9c, 0x01, 0x0a, 0x12, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x61, 0x79, 0x4c, 0x6f, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x15, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x61,
	0x79, 0x4c, 0x6f, 0x67, 0x1a, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x61, 0x79, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12,
	0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x50, 0x61, 0x79, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x61,
	0x79, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d,
	0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_orderPayLogService_proto_rawDescOnce sync.Once
	file_orderPayLogService_proto_rawDescData = file_orderPayLogService_proto_rawDesc
)

func file_orderPayLogService_proto_rawDescGZIP() []byte {
	file_orderPayLogService_proto_rawDescOnce.Do(func() {
		file_orderPayLogService_proto_rawDescData = protoimpl.X.CompressGZIP(file_orderPayLogService_proto_rawDescData)
	})
	return file_orderPayLogService_proto_rawDescData
}

var file_orderPayLogService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_orderPayLogService_proto_goTypes = []interface{}{
	(*OrderPayLog)(nil),         // 0: services.OrderPayLog
	(*OrderPayLogRequest)(nil),  // 1: services.OrderPayLogRequest
	(*OrderPayLogData)(nil),     // 2: services.OrderPayLogData
	(*OrderPayLogResponse)(nil), // 3: services.OrderPayLogResponse
	(*common.Pager)(nil),        // 4: common.Pager
	(*common.Info)(nil),         // 5: common.Info
	(*common.Error)(nil),        // 6: common.Error
}
var file_orderPayLogService_proto_depIdxs = []int32{
	0, // 0: services.OrderPayLogData.entity:type_name -> services.OrderPayLog
	0, // 1: services.OrderPayLogData.items:type_name -> services.OrderPayLog
	4, // 2: services.OrderPayLogData.pager:type_name -> common.Pager
	5, // 3: services.OrderPayLogData.info:type_name -> common.Info
	2, // 4: services.OrderPayLogResponse.data:type_name -> services.OrderPayLogData
	6, // 5: services.OrderPayLogResponse.error:type_name -> common.Error
	0, // 6: services.OrderPayLogService.Get:input_type -> services.OrderPayLog
	1, // 7: services.OrderPayLogService.Search:input_type -> services.OrderPayLogRequest
	3, // 8: services.OrderPayLogService.Get:output_type -> services.OrderPayLogResponse
	3, // 9: services.OrderPayLogService.Search:output_type -> services.OrderPayLogResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_orderPayLogService_proto_init() }
func file_orderPayLogService_proto_init() {
	if File_orderPayLogService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_orderPayLogService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderPayLog); i {
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
		file_orderPayLogService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderPayLogRequest); i {
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
		file_orderPayLogService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderPayLogData); i {
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
		file_orderPayLogService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderPayLogResponse); i {
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
			RawDescriptor: file_orderPayLogService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_orderPayLogService_proto_goTypes,
		DependencyIndexes: file_orderPayLogService_proto_depIdxs,
		MessageInfos:      file_orderPayLogService_proto_msgTypes,
	}.Build()
	File_orderPayLogService_proto = out.File
	file_orderPayLogService_proto_rawDesc = nil
	file_orderPayLogService_proto_goTypes = nil
	file_orderPayLogService_proto_depIdxs = nil
}
