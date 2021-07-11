// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: chargeService.proto

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

//支付凭证
type Charge struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ChargeSn       string  `protobuf:"bytes,2,opt,name=charge_sn,json=chargeSn,proto3" json:"charge_sn,omitempty"`
	Lived          bool    `protobuf:"varint,3,opt,name=lived,proto3" json:"lived,omitempty"`
	Paid           bool    `protobuf:"varint,4,opt,name=paid,proto3" json:"paid,omitempty"`
	Refunded       bool    `protobuf:"varint,5,opt,name=refunded,proto3" json:"refunded,omitempty"`
	Reversed       bool    `protobuf:"varint,6,opt,name=reversed,proto3" json:"reversed,omitempty"`
	PayGatewayId   int64   `protobuf:"varint,7,opt,name=pay_gateway_id,json=payGatewayId,proto3" json:"pay_gateway_id,omitempty"`
	OpenId         string  `protobuf:"bytes,8,opt,name=open_id,json=openId,proto3" json:"open_id,omitempty"`
	TargetUserType string  `protobuf:"bytes,9,opt,name=target_user_type,json=targetUserType,proto3" json:"target_user_type,omitempty"`
	TargetUserId   int64   `protobuf:"varint,10,opt,name=target_user_id,json=targetUserId,proto3" json:"target_user_id,omitempty"`
	Type           string  `protobuf:"bytes,11,opt,name=type,proto3" json:"type,omitempty"`
	OrderId        int64   `protobuf:"varint,12,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	OrderSn        string  `protobuf:"bytes,13,opt,name=order_sn,json=orderSn,proto3" json:"order_sn,omitempty"`
	Amount         float32 `protobuf:"fixed32,14,opt,name=amount,proto3" json:"amount,omitempty"`
	Currency       string  `protobuf:"bytes,15,opt,name=currency,proto3" json:"currency,omitempty"`
	Subject        string  `protobuf:"bytes,16,opt,name=subject,proto3" json:"subject,omitempty"`
	Body           string  `protobuf:"bytes,17,opt,name=body,proto3" json:"body,omitempty"`
	Extra          string  `protobuf:"bytes,18,opt,name=extra,proto3" json:"extra,omitempty"`
	Channel        string  `protobuf:"bytes,19,opt,name=channel,proto3" json:"channel,omitempty"`
	TradeType      string  `protobuf:"bytes,20,opt,name=trade_type,json=tradeType,proto3" json:"trade_type,omitempty"`
	TransactionNo  string  `protobuf:"bytes,21,opt,name=transaction_no,json=transactionNo,proto3" json:"transaction_no,omitempty"`
	RefundedAmount float32 `protobuf:"fixed32,22,opt,name=refunded_amount,json=refundedAmount,proto3" json:"refunded_amount,omitempty"`
	FailureCode    string  `protobuf:"bytes,23,opt,name=failure_code,json=failureCode,proto3" json:"failure_code,omitempty"`
	FailureMsg     string  `protobuf:"bytes,24,opt,name=failure_msg,json=failureMsg,proto3" json:"failure_msg,omitempty"`
	Credential     string  `protobuf:"bytes,25,opt,name=credential,proto3" json:"credential,omitempty"`
	ClientIp       string  `protobuf:"bytes,26,opt,name=client_ip,json=clientIp,proto3" json:"client_ip,omitempty"`
	Memo           string  `protobuf:"bytes,27,opt,name=memo,proto3" json:"memo,omitempty"`
	Metadata       string  `protobuf:"bytes,28,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Status         string  `protobuf:"bytes,29,opt,name=status,proto3" json:"status,omitempty"`
	PaidAt         string  `protobuf:"bytes,30,opt,name=paid_at,json=paidAt,proto3" json:"paid_at,omitempty"`
	ExpiredAt      string  `protobuf:"bytes,31,opt,name=expired_at,json=expiredAt,proto3" json:"expired_at,omitempty"`
	CreatedAt      string  `protobuf:"bytes,32,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt      string  `protobuf:"bytes,33,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Charge) Reset() {
	*x = Charge{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chargeService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Charge) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Charge) ProtoMessage() {}

func (x *Charge) ProtoReflect() protoreflect.Message {
	mi := &file_chargeService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Charge.ProtoReflect.Descriptor instead.
func (*Charge) Descriptor() ([]byte, []int) {
	return file_chargeService_proto_rawDescGZIP(), []int{0}
}

func (x *Charge) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Charge) GetChargeSn() string {
	if x != nil {
		return x.ChargeSn
	}
	return ""
}

func (x *Charge) GetLived() bool {
	if x != nil {
		return x.Lived
	}
	return false
}

func (x *Charge) GetPaid() bool {
	if x != nil {
		return x.Paid
	}
	return false
}

func (x *Charge) GetRefunded() bool {
	if x != nil {
		return x.Refunded
	}
	return false
}

func (x *Charge) GetReversed() bool {
	if x != nil {
		return x.Reversed
	}
	return false
}

func (x *Charge) GetPayGatewayId() int64 {
	if x != nil {
		return x.PayGatewayId
	}
	return 0
}

func (x *Charge) GetOpenId() string {
	if x != nil {
		return x.OpenId
	}
	return ""
}

func (x *Charge) GetTargetUserType() string {
	if x != nil {
		return x.TargetUserType
	}
	return ""
}

func (x *Charge) GetTargetUserId() int64 {
	if x != nil {
		return x.TargetUserId
	}
	return 0
}

func (x *Charge) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Charge) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *Charge) GetOrderSn() string {
	if x != nil {
		return x.OrderSn
	}
	return ""
}

func (x *Charge) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Charge) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *Charge) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *Charge) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *Charge) GetExtra() string {
	if x != nil {
		return x.Extra
	}
	return ""
}

func (x *Charge) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *Charge) GetTradeType() string {
	if x != nil {
		return x.TradeType
	}
	return ""
}

func (x *Charge) GetTransactionNo() string {
	if x != nil {
		return x.TransactionNo
	}
	return ""
}

func (x *Charge) GetRefundedAmount() float32 {
	if x != nil {
		return x.RefundedAmount
	}
	return 0
}

func (x *Charge) GetFailureCode() string {
	if x != nil {
		return x.FailureCode
	}
	return ""
}

func (x *Charge) GetFailureMsg() string {
	if x != nil {
		return x.FailureMsg
	}
	return ""
}

func (x *Charge) GetCredential() string {
	if x != nil {
		return x.Credential
	}
	return ""
}

func (x *Charge) GetClientIp() string {
	if x != nil {
		return x.ClientIp
	}
	return ""
}

func (x *Charge) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *Charge) GetMetadata() string {
	if x != nil {
		return x.Metadata
	}
	return ""
}

func (x *Charge) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Charge) GetPaidAt() string {
	if x != nil {
		return x.PaidAt
	}
	return ""
}

func (x *Charge) GetExpiredAt() string {
	if x != nil {
		return x.ExpiredAt
	}
	return ""
}

func (x *Charge) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Charge) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

//
type ChargeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Charge       `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items  []*Charge     `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error  *common.Error `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info   *common.Info  `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *ChargeResponse) Reset() {
	*x = ChargeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chargeService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChargeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChargeResponse) ProtoMessage() {}

func (x *ChargeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chargeService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChargeResponse.ProtoReflect.Descriptor instead.
func (*ChargeResponse) Descriptor() ([]byte, []int) {
	return file_chargeService_proto_rawDescGZIP(), []int{1}
}

func (x *ChargeResponse) GetEntity() *Charge {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *ChargeResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *ChargeResponse) GetItems() []*Charge {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ChargeResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *ChargeResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_chargeService_proto protoreflect.FileDescriptor

var file_chargeService_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a,
	0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xb0, 0x07, 0x0a, 0x06, 0x43, 0x68, 0x61, 0x72, 0x67, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x5f, 0x73, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x53, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69,
	0x76, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x6c, 0x69, 0x76, 0x65, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04,
	0x70, 0x61, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x65, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x65, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x64, 0x12, 0x24, 0x0a, 0x0e,
	0x70, 0x61, 0x79, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x70, 0x61, 0x79, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x70, 0x65, 0x6e, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x73, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x53, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x11, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61,
	0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x72, 0x61, 0x64, 0x65,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x72, 0x61,
	0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x6f, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x12, 0x27, 0x0a,
	0x0f, 0x72, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x65, 0x64, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x16, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0e, 0x72, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x65, 0x64,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72,
	0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x61,
	0x69, 0x6c, 0x75, 0x72, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x61, 0x69,
	0x6c, 0x75, 0x72, 0x65, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x4d, 0x73, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x18, 0x19, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x70, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18,
	0x1b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x17, 0x0a, 0x07, 0x70, 0x61, 0x69, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x70, 0x61, 0x69, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x20, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x21, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xce, 0x01, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x72, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x67, 0x65, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72,
	0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x67, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12,
	0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0xf5, 0x01, 0x0a, 0x0d, 0x43, 0x68, 0x61, 0x72, 0x67,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x68,
	0x61, 0x72, 0x67, 0x65, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x43, 0x68, 0x61, 0x72, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x3e, 0x0a, 0x0e, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x69, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x68,
	0x61, 0x72, 0x67, 0x65, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x43, 0x68, 0x61, 0x72, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x33, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x67, 0x65, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12,
	0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x57, 0x68, 0x65,
	0x72, 0x65, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x68,
	0x61, 0x72, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c,
	0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chargeService_proto_rawDescOnce sync.Once
	file_chargeService_proto_rawDescData = file_chargeService_proto_rawDesc
)

func file_chargeService_proto_rawDescGZIP() []byte {
	file_chargeService_proto_rawDescOnce.Do(func() {
		file_chargeService_proto_rawDescData = protoimpl.X.CompressGZIP(file_chargeService_proto_rawDescData)
	})
	return file_chargeService_proto_rawDescData
}

var file_chargeService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_chargeService_proto_goTypes = []interface{}{
	(*Charge)(nil),           // 0: services.Charge
	(*ChargeResponse)(nil),   // 1: services.ChargeResponse
	(*common.Pager)(nil),     // 2: common.Pager
	(*common.Error)(nil),     // 3: common.Error
	(*common.Info)(nil),      // 4: common.Info
	(*common.BaseWhere)(nil), // 5: common.BaseWhere
}
var file_chargeService_proto_depIdxs = []int32{
	0, // 0: services.ChargeResponse.entity:type_name -> services.Charge
	2, // 1: services.ChargeResponse.pager:type_name -> common.Pager
	0, // 2: services.ChargeResponse.items:type_name -> services.Charge
	3, // 3: services.ChargeResponse.error:type_name -> common.Error
	4, // 4: services.ChargeResponse.info:type_name -> common.Info
	0, // 5: services.ChargeService.Delete:input_type -> services.Charge
	0, // 6: services.ChargeService.Reconciliation:input_type -> services.Charge
	0, // 7: services.ChargeService.Get:input_type -> services.Charge
	5, // 8: services.ChargeService.Search:input_type -> common.BaseWhere
	1, // 9: services.ChargeService.Delete:output_type -> services.ChargeResponse
	1, // 10: services.ChargeService.Reconciliation:output_type -> services.ChargeResponse
	1, // 11: services.ChargeService.Get:output_type -> services.ChargeResponse
	1, // 12: services.ChargeService.Search:output_type -> services.ChargeResponse
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_chargeService_proto_init() }
func file_chargeService_proto_init() {
	if File_chargeService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chargeService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Charge); i {
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
		file_chargeService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChargeResponse); i {
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
			RawDescriptor: file_chargeService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chargeService_proto_goTypes,
		DependencyIndexes: file_chargeService_proto_depIdxs,
		MessageInfos:      file_chargeService_proto_msgTypes,
	}.Build()
	File_chargeService_proto = out.File
	file_chargeService_proto_rawDesc = nil
	file_chargeService_proto_goTypes = nil
	file_chargeService_proto_depIdxs = nil
}
