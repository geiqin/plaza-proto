// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: orderTransactService.proto

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

type OrderTransact struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	OrderId       int64  `protobuf:"varint,2,opt,name=order_id,json=orderId,proto3" json:"order_id"`
	Type          string `protobuf:"bytes,3,opt,name=type,proto3" json:"type"`
	Currency      int64  `protobuf:"varint,5,opt,name=currency,proto3" json:"currency"`
	Amount        int64  `protobuf:"varint,6,opt,name=amount,proto3" json:"amount"`
	TradeUserType string `protobuf:"bytes,7,opt,name=trade_user_type,json=tradeUserType,proto3" json:"trade_user_type"`
	TradeUserId   int64  `protobuf:"varint,8,opt,name=trade_user_id,json=tradeUserId,proto3" json:"trade_user_id"`
	PayAccountId  int32  `protobuf:"varint,9,opt,name=pay_account_id,json=payAccountId,proto3" json:"pay_account_id"`
	TransactionNo string `protobuf:"bytes,10,opt,name=transaction_no,json=transactionNo,proto3" json:"transaction_no"`
	ChannelType   string `protobuf:"bytes,11,opt,name=channel_type,json=channelType,proto3" json:"channel_type"`
	ChargeId      int64  `protobuf:"varint,12,opt,name=charge_id,json=chargeId,proto3" json:"charge_id"`
	RefundId      int64  `protobuf:"varint,13,opt,name=refund_id,json=refundId,proto3" json:"refund_id"`
}

func (x *OrderTransact) Reset() {
	*x = OrderTransact{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderTransactService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderTransact) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderTransact) ProtoMessage() {}

func (x *OrderTransact) ProtoReflect() protoreflect.Message {
	mi := &file_orderTransactService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderTransact.ProtoReflect.Descriptor instead.
func (*OrderTransact) Descriptor() ([]byte, []int) {
	return file_orderTransactService_proto_rawDescGZIP(), []int{0}
}

func (x *OrderTransact) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OrderTransact) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *OrderTransact) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *OrderTransact) GetCurrency() int64 {
	if x != nil {
		return x.Currency
	}
	return 0
}

func (x *OrderTransact) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *OrderTransact) GetTradeUserType() string {
	if x != nil {
		return x.TradeUserType
	}
	return ""
}

func (x *OrderTransact) GetTradeUserId() int64 {
	if x != nil {
		return x.TradeUserId
	}
	return 0
}

func (x *OrderTransact) GetPayAccountId() int32 {
	if x != nil {
		return x.PayAccountId
	}
	return 0
}

func (x *OrderTransact) GetTransactionNo() string {
	if x != nil {
		return x.TransactionNo
	}
	return ""
}

func (x *OrderTransact) GetChannelType() string {
	if x != nil {
		return x.ChannelType
	}
	return ""
}

func (x *OrderTransact) GetChargeId() int64 {
	if x != nil {
		return x.ChargeId
	}
	return 0
}

func (x *OrderTransact) GetRefundId() int64 {
	if x != nil {
		return x.RefundId
	}
	return 0
}

//支付请求参数
type OrderTransactRequest struct {
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
	RefundId    int64   `protobuf:"varint,10,opt,name=refund_id,json=refundId,proto3" json:"refund_id"`
	Ids         []int64 `protobuf:"varint,11,rep,packed,name=ids,proto3" json:"ids"`
}

func (x *OrderTransactRequest) Reset() {
	*x = OrderTransactRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderTransactService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderTransactRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderTransactRequest) ProtoMessage() {}

func (x *OrderTransactRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orderTransactService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderTransactRequest.ProtoReflect.Descriptor instead.
func (*OrderTransactRequest) Descriptor() ([]byte, []int) {
	return file_orderTransactService_proto_rawDescGZIP(), []int{1}
}

func (x *OrderTransactRequest) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *OrderTransactRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *OrderTransactRequest) GetSorting() string {
	if x != nil {
		return x.Sorting
	}
	return ""
}

func (x *OrderTransactRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *OrderTransactRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OrderTransactRequest) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *OrderTransactRequest) GetChannelType() string {
	if x != nil {
		return x.ChannelType
	}
	return ""
}

func (x *OrderTransactRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *OrderTransactRequest) GetChargeId() int64 {
	if x != nil {
		return x.ChargeId
	}
	return 0
}

func (x *OrderTransactRequest) GetRefundId() int64 {
	if x != nil {
		return x.RefundId
	}
	return 0
}

func (x *OrderTransactRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type OrderTransactData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *OrderTransact   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Items  []*OrderTransact `protobuf:"bytes,2,rep,name=items,proto3" json:"items"`
	Info   *common.Info     `protobuf:"bytes,3,opt,name=info,proto3" json:"info"`
}

func (x *OrderTransactData) Reset() {
	*x = OrderTransactData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderTransactService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderTransactData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderTransactData) ProtoMessage() {}

func (x *OrderTransactData) ProtoReflect() protoreflect.Message {
	mi := &file_orderTransactService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderTransactData.ProtoReflect.Descriptor instead.
func (*OrderTransactData) Descriptor() ([]byte, []int) {
	return file_orderTransactService_proto_rawDescGZIP(), []int{2}
}

func (x *OrderTransactData) GetEntity() *OrderTransact {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *OrderTransactData) GetItems() []*OrderTransact {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *OrderTransactData) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type OrderTransactResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *OrderTransactData `protobuf:"bytes,1,opt,name=data,proto3" json:"data"`
	Error *common.Error      `protobuf:"bytes,2,opt,name=error,proto3" json:"error"`
}

func (x *OrderTransactResponse) Reset() {
	*x = OrderTransactResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderTransactService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderTransactResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderTransactResponse) ProtoMessage() {}

func (x *OrderTransactResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orderTransactService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderTransactResponse.ProtoReflect.Descriptor instead.
func (*OrderTransactResponse) Descriptor() ([]byte, []int) {
	return file_orderTransactService_proto_rawDescGZIP(), []int{3}
}

func (x *OrderTransactResponse) GetData() *OrderTransactData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *OrderTransactResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_orderTransactService_proto protoreflect.FileDescriptor

var file_orderTransactService_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf8, 0x02, 0x0a, 0x0d, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x26,
	0x0a, 0x0f, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x64, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x74,
	0x72, 0x61, 0x64, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x70, 0x61,
	0x79, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0c, 0x70, 0x61, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x6e, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x68,
	0x61, 0x72, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63,
	0x68, 0x61, 0x72, 0x67, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x66, 0x75, 0x6e,
	0x64, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x72, 0x65, 0x66, 0x75,
	0x6e, 0x64, 0x49, 0x64, 0x22, 0xad, 0x02, 0x0a, 0x14, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61,
	0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65,
	0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65,
	0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x72,
	0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x68, 0x61,
	0x72, 0x67, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x5f,
	0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x72, 0x65, 0x66, 0x75, 0x6e, 0x64,
	0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x03, 0x52,
	0x03, 0x69, 0x64, 0x73, 0x22, 0x95, 0x01, 0x0a, 0x11, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2f, 0x0a, 0x06, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x2d, 0x0a, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x6d, 0x0a, 0x15,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xa6, 0x01, 0x0a, 0x14,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x17, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x1a, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x12, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_orderTransactService_proto_rawDescOnce sync.Once
	file_orderTransactService_proto_rawDescData = file_orderTransactService_proto_rawDesc
)

func file_orderTransactService_proto_rawDescGZIP() []byte {
	file_orderTransactService_proto_rawDescOnce.Do(func() {
		file_orderTransactService_proto_rawDescData = protoimpl.X.CompressGZIP(file_orderTransactService_proto_rawDescData)
	})
	return file_orderTransactService_proto_rawDescData
}

var file_orderTransactService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_orderTransactService_proto_goTypes = []interface{}{
	(*OrderTransact)(nil),         // 0: services.OrderTransact
	(*OrderTransactRequest)(nil),  // 1: services.OrderTransactRequest
	(*OrderTransactData)(nil),     // 2: services.OrderTransactData
	(*OrderTransactResponse)(nil), // 3: services.OrderTransactResponse
	(*common.Info)(nil),           // 4: common.Info
	(*common.Error)(nil),          // 5: common.Error
}
var file_orderTransactService_proto_depIdxs = []int32{
	0, // 0: services.OrderTransactData.entity:type_name -> services.OrderTransact
	0, // 1: services.OrderTransactData.items:type_name -> services.OrderTransact
	4, // 2: services.OrderTransactData.info:type_name -> common.Info
	2, // 3: services.OrderTransactResponse.data:type_name -> services.OrderTransactData
	5, // 4: services.OrderTransactResponse.error:type_name -> common.Error
	0, // 5: services.OrderTransactService.Get:input_type -> services.OrderTransact
	1, // 6: services.OrderTransactService.Search:input_type -> services.OrderTransactRequest
	3, // 7: services.OrderTransactService.Get:output_type -> services.OrderTransactResponse
	3, // 8: services.OrderTransactService.Search:output_type -> services.OrderTransactResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_orderTransactService_proto_init() }
func file_orderTransactService_proto_init() {
	if File_orderTransactService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_orderTransactService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderTransact); i {
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
		file_orderTransactService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderTransactRequest); i {
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
		file_orderTransactService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderTransactData); i {
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
		file_orderTransactService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderTransactResponse); i {
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
			RawDescriptor: file_orderTransactService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_orderTransactService_proto_goTypes,
		DependencyIndexes: file_orderTransactService_proto_depIdxs,
		MessageInfos:      file_orderTransactService_proto_msgTypes,
	}.Build()
	File_orderTransactService_proto = out.File
	file_orderTransactService_proto_rawDesc = nil
	file_orderTransactService_proto_goTypes = nil
	file_orderTransactService_proto_depIdxs = nil
}
