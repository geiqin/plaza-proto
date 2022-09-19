// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: refundService.proto

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

//退款
type Refund struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                  int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	RefundSn            string `protobuf:"bytes,2,opt,name=refund_sn,json=refundSn,proto3" json:"refund_sn"`
	TotalFee            int64  `protobuf:"varint,3,opt,name=total_fee,json=totalFee,proto3" json:"total_fee"`
	RefundFee           int64  `protobuf:"varint,4,opt,name=refund_fee,json=refundFee,proto3" json:"refund_fee"`
	Currency            string `protobuf:"bytes,5,opt,name=currency,proto3" json:"currency"`
	ChargeId            int64  `protobuf:"varint,8,opt,name=charge_id,json=chargeId,proto3" json:"charge_id"`
	ChargeSn            string `protobuf:"bytes,9,opt,name=charge_sn,json=chargeSn,proto3" json:"charge_sn"`
	ChargeTransactionNo string `protobuf:"bytes,10,opt,name=charge_transaction_no,json=chargeTransactionNo,proto3" json:"charge_transaction_no"`
	TargetUserType      string `protobuf:"bytes,11,opt,name=target_user_type,json=targetUserType,proto3" json:"target_user_type"`
	TargetUserId        int64  `protobuf:"varint,12,opt,name=target_user_id,json=targetUserId,proto3" json:"target_user_id"`
	SourceOrderId       int64  `protobuf:"varint,13,opt,name=source_order_id,json=sourceOrderId,proto3" json:"source_order_id"`
	RefundDesc          string `protobuf:"bytes,14,opt,name=refund_desc,json=refundDesc,proto3" json:"refund_desc"`
	RefundSource        string `protobuf:"bytes,15,opt,name=refund_source,json=refundSource,proto3" json:"refund_source"`
	Metadata            string `protobuf:"bytes,16,opt,name=metadata,proto3" json:"metadata"`
	ReturnExtra         string `protobuf:"bytes,17,opt,name=return_extra,json=returnExtra,proto3" json:"return_extra"`
	TransactionNo       string `protobuf:"bytes,18,opt,name=transaction_no,json=transactionNo,proto3" json:"transaction_no"`
	SafeguardId         int64  `protobuf:"varint,19,opt,name=safeguard_id,json=safeguardId,proto3" json:"safeguard_id"`
	SafeguardSn         string `protobuf:"bytes,20,opt,name=safeguard_sn,json=safeguardSn,proto3" json:"safeguard_sn"`
	Status              string `protobuf:"bytes,21,opt,name=status,proto3" json:"status"`
	FinishedAt          string `protobuf:"bytes,22,opt,name=finished_at,json=finishedAt,proto3" json:"finished_at"`
	CreatedAt           string `protobuf:"bytes,23,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt           string `protobuf:"bytes,24,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
}

func (x *Refund) Reset() {
	*x = Refund{}
	if protoimpl.UnsafeEnabled {
		mi := &file_refundService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Refund) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Refund) ProtoMessage() {}

func (x *Refund) ProtoReflect() protoreflect.Message {
	mi := &file_refundService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Refund.ProtoReflect.Descriptor instead.
func (*Refund) Descriptor() ([]byte, []int) {
	return file_refundService_proto_rawDescGZIP(), []int{0}
}

func (x *Refund) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Refund) GetRefundSn() string {
	if x != nil {
		return x.RefundSn
	}
	return ""
}

func (x *Refund) GetTotalFee() int64 {
	if x != nil {
		return x.TotalFee
	}
	return 0
}

func (x *Refund) GetRefundFee() int64 {
	if x != nil {
		return x.RefundFee
	}
	return 0
}

func (x *Refund) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *Refund) GetChargeId() int64 {
	if x != nil {
		return x.ChargeId
	}
	return 0
}

func (x *Refund) GetChargeSn() string {
	if x != nil {
		return x.ChargeSn
	}
	return ""
}

func (x *Refund) GetChargeTransactionNo() string {
	if x != nil {
		return x.ChargeTransactionNo
	}
	return ""
}

func (x *Refund) GetTargetUserType() string {
	if x != nil {
		return x.TargetUserType
	}
	return ""
}

func (x *Refund) GetTargetUserId() int64 {
	if x != nil {
		return x.TargetUserId
	}
	return 0
}

func (x *Refund) GetSourceOrderId() int64 {
	if x != nil {
		return x.SourceOrderId
	}
	return 0
}

func (x *Refund) GetRefundDesc() string {
	if x != nil {
		return x.RefundDesc
	}
	return ""
}

func (x *Refund) GetRefundSource() string {
	if x != nil {
		return x.RefundSource
	}
	return ""
}

func (x *Refund) GetMetadata() string {
	if x != nil {
		return x.Metadata
	}
	return ""
}

func (x *Refund) GetReturnExtra() string {
	if x != nil {
		return x.ReturnExtra
	}
	return ""
}

func (x *Refund) GetTransactionNo() string {
	if x != nil {
		return x.TransactionNo
	}
	return ""
}

func (x *Refund) GetSafeguardId() int64 {
	if x != nil {
		return x.SafeguardId
	}
	return 0
}

func (x *Refund) GetSafeguardSn() string {
	if x != nil {
		return x.SafeguardSn
	}
	return ""
}

func (x *Refund) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Refund) GetFinishedAt() string {
	if x != nil {
		return x.FinishedAt
	}
	return ""
}

func (x *Refund) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Refund) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type RefundRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int32 `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Top      int32 `protobuf:"varint,3,opt,name=top,proto3" json:"top"`
	//base params
	SafeguardId    int64  `protobuf:"varint,4,opt,name=safeguard_id,json=safeguardId,proto3" json:"safeguard_id"`
	SafeguardSn    string `protobuf:"bytes,5,opt,name=safeguard_sn,json=safeguardSn,proto3" json:"safeguard_sn"`
	SourceOrderId  int64  `protobuf:"varint,6,opt,name=source_order_id,json=sourceOrderId,proto3" json:"source_order_id"`
	Status         string `protobuf:"bytes,7,opt,name=status,proto3" json:"status"`
	OrderId        int64  `protobuf:"varint,8,opt,name=order_id,json=orderId,proto3" json:"order_id"`
	RefundSn       string `protobuf:"bytes,9,opt,name=refund_sn,json=refundSn,proto3" json:"refund_sn"`
	Metadata       string `protobuf:"bytes,10,opt,name=metadata,proto3" json:"metadata"`
	TargetUserType string `protobuf:"bytes,11,opt,name=target_user_type,json=targetUserType,proto3" json:"target_user_type"`
	TargetUserId   int64  `protobuf:"varint,12,opt,name=target_user_id,json=targetUserId,proto3" json:"target_user_id"`
}

func (x *RefundRequest) Reset() {
	*x = RefundRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_refundService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefundRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefundRequest) ProtoMessage() {}

func (x *RefundRequest) ProtoReflect() protoreflect.Message {
	mi := &file_refundService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefundRequest.ProtoReflect.Descriptor instead.
func (*RefundRequest) Descriptor() ([]byte, []int) {
	return file_refundService_proto_rawDescGZIP(), []int{1}
}

func (x *RefundRequest) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *RefundRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *RefundRequest) GetTop() int32 {
	if x != nil {
		return x.Top
	}
	return 0
}

func (x *RefundRequest) GetSafeguardId() int64 {
	if x != nil {
		return x.SafeguardId
	}
	return 0
}

func (x *RefundRequest) GetSafeguardSn() string {
	if x != nil {
		return x.SafeguardSn
	}
	return ""
}

func (x *RefundRequest) GetSourceOrderId() int64 {
	if x != nil {
		return x.SourceOrderId
	}
	return 0
}

func (x *RefundRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *RefundRequest) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *RefundRequest) GetRefundSn() string {
	if x != nil {
		return x.RefundSn
	}
	return ""
}

func (x *RefundRequest) GetMetadata() string {
	if x != nil {
		return x.Metadata
	}
	return ""
}

func (x *RefundRequest) GetTargetUserType() string {
	if x != nil {
		return x.TargetUserType
	}
	return ""
}

func (x *RefundRequest) GetTargetUserId() int64 {
	if x != nil {
		return x.TargetUserId
	}
	return 0
}

type RefundData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Refund       `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*Refund     `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Info   *common.Info  `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
}

func (x *RefundData) Reset() {
	*x = RefundData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_refundService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefundData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefundData) ProtoMessage() {}

func (x *RefundData) ProtoReflect() protoreflect.Message {
	mi := &file_refundService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefundData.ProtoReflect.Descriptor instead.
func (*RefundData) Descriptor() ([]byte, []int) {
	return file_refundService_proto_rawDescGZIP(), []int{2}
}

func (x *RefundData) GetEntity() *Refund {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *RefundData) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *RefundData) GetItems() []*Refund {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *RefundData) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type RefundResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *RefundData   `protobuf:"bytes,1,opt,name=data,proto3" json:"data"`
	Error *common.Error `protobuf:"bytes,2,opt,name=error,proto3" json:"error"`
}

func (x *RefundResponse) Reset() {
	*x = RefundResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_refundService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefundResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefundResponse) ProtoMessage() {}

func (x *RefundResponse) ProtoReflect() protoreflect.Message {
	mi := &file_refundService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefundResponse.ProtoReflect.Descriptor instead.
func (*RefundResponse) Descriptor() ([]byte, []int) {
	return file_refundService_proto_rawDescGZIP(), []int{3}
}

func (x *RefundResponse) GetData() *RefundData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *RefundResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_refundService_proto protoreflect.FileDescriptor

var file_refundService_proto_rawDesc = []byte{
	0x0a, 0x13, 0x72, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a,
	0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xdc, 0x05, 0x0a, 0x06, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x72, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x5f, 0x73, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x72, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x53, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x5f, 0x66, 0x65, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x46, 0x65, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x66, 0x75, 0x6e,
	0x64, 0x5f, 0x66, 0x65, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x72, 0x65, 0x66,
	0x75, 0x6e, 0x64, 0x46, 0x65, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x49, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x5f, 0x73, 0x6e, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x53, 0x6e, 0x12, 0x32, 0x0a, 0x15,
	0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x6e, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x63, 0x68, 0x61,
	0x72, 0x67, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x6f,
	0x12, 0x28, 0x0a, 0x10, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0c, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x26, 0x0a, 0x0f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x66, 0x75,
	0x6e, 0x64, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72,
	0x65, 0x66, 0x75, 0x6e, 0x64, 0x44, 0x65, 0x73, 0x63, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66,
	0x75, 0x6e, 0x64, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x72, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65,
	0x74, 0x75, 0x72, 0x6e, 0x5f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x12, 0x25, 0x0a,
	0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x6f, 0x18,
	0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x4e, 0x6f, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x61, 0x66, 0x65, 0x67, 0x75, 0x61, 0x72,
	0x64, 0x5f, 0x69, 0x64, 0x18, 0x13, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x73, 0x61, 0x66, 0x65,
	0x67, 0x75, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x61, 0x66, 0x65, 0x67,
	0x75, 0x61, 0x72, 0x64, 0x5f, 0x73, 0x6e, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73,
	0x61, 0x66, 0x65, 0x67, 0x75, 0x61, 0x72, 0x64, 0x53, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0xfe, 0x02, 0x0a, 0x0d, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x6f, 0x70, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x74, 0x6f, 0x70, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x61, 0x66, 0x65,
	0x67, 0x75, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b,
	0x73, 0x61, 0x66, 0x65, 0x67, 0x75, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x73,
	0x61, 0x66, 0x65, 0x67, 0x75, 0x61, 0x72, 0x64, 0x5f, 0x73, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x73, 0x61, 0x66, 0x65, 0x67, 0x75, 0x61, 0x72, 0x64, 0x53, 0x6e, 0x12, 0x26,
	0x0a, 0x0f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x19,
	0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x66,
	0x75, 0x6e, 0x64, 0x5f, 0x73, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65,
	0x66, 0x75, 0x6e, 0x64, 0x53, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x28, 0x0a, 0x10, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x0e,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x22, 0xa5, 0x01, 0x0a, 0x0a, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x28, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x66,
	0x75, 0x6e, 0x64, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70,
	0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72,
	0x12, 0x26, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x66, 0x75, 0x6e,
	0x64, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x5f, 0x0a, 0x0e, 0x52, 0x65,
	0x66, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xb3, 0x02, 0x0a, 0x0d,
	0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a,
	0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12,
	0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x66, 0x75, 0x6e,
	0x64, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x66,
	0x75, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x33, 0x0a,
	0x03, 0x47, 0x65, 0x74, 0x12, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x3d, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x17, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x3e, 0x0a, 0x0e, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x69, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52,
	0x65, 0x66, 0x75, 0x6e, 0x64, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_refundService_proto_rawDescOnce sync.Once
	file_refundService_proto_rawDescData = file_refundService_proto_rawDesc
)

func file_refundService_proto_rawDescGZIP() []byte {
	file_refundService_proto_rawDescOnce.Do(func() {
		file_refundService_proto_rawDescData = protoimpl.X.CompressGZIP(file_refundService_proto_rawDescData)
	})
	return file_refundService_proto_rawDescData
}

var file_refundService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_refundService_proto_goTypes = []interface{}{
	(*Refund)(nil),         // 0: services.Refund
	(*RefundRequest)(nil),  // 1: services.RefundRequest
	(*RefundData)(nil),     // 2: services.RefundData
	(*RefundResponse)(nil), // 3: services.RefundResponse
	(*common.Pager)(nil),   // 4: common.Pager
	(*common.Info)(nil),    // 5: common.Info
	(*common.Error)(nil),   // 6: common.Error
}
var file_refundService_proto_depIdxs = []int32{
	0,  // 0: services.RefundData.entity:type_name -> services.Refund
	4,  // 1: services.RefundData.pager:type_name -> common.Pager
	0,  // 2: services.RefundData.items:type_name -> services.Refund
	5,  // 3: services.RefundData.info:type_name -> common.Info
	2,  // 4: services.RefundResponse.data:type_name -> services.RefundData
	6,  // 5: services.RefundResponse.error:type_name -> common.Error
	0,  // 6: services.RefundService.Create:input_type -> services.Refund
	0,  // 7: services.RefundService.Delete:input_type -> services.Refund
	0,  // 8: services.RefundService.Get:input_type -> services.Refund
	1,  // 9: services.RefundService.Search:input_type -> services.RefundRequest
	0,  // 10: services.RefundService.Reconciliation:input_type -> services.Refund
	3,  // 11: services.RefundService.Create:output_type -> services.RefundResponse
	3,  // 12: services.RefundService.Delete:output_type -> services.RefundResponse
	3,  // 13: services.RefundService.Get:output_type -> services.RefundResponse
	3,  // 14: services.RefundService.Search:output_type -> services.RefundResponse
	3,  // 15: services.RefundService.Reconciliation:output_type -> services.RefundResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_refundService_proto_init() }
func file_refundService_proto_init() {
	if File_refundService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_refundService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Refund); i {
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
		file_refundService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefundRequest); i {
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
		file_refundService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefundData); i {
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
		file_refundService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefundResponse); i {
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
			RawDescriptor: file_refundService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_refundService_proto_goTypes,
		DependencyIndexes: file_refundService_proto_depIdxs,
		MessageInfos:      file_refundService_proto_msgTypes,
	}.Build()
	File_refundService_proto = out.File
	file_refundService_proto_rawDesc = nil
	file_refundService_proto_goTypes = nil
	file_refundService_proto_depIdxs = nil
}
