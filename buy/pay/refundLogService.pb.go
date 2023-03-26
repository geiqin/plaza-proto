// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: refundLogService.proto

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

//退款日志
type RefundLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`                                            //ID
	LogNo         string `protobuf:"bytes,2,opt,name=log_no,json=logNo,proto3" json:"log_no"`                          //退款单号
	PayLogId      int64  `protobuf:"varint,3,opt,name=pay_log_id,json=payLogId,proto3" json:"pay_log_id"`              //支付日志id
	MemberId      int64  `protobuf:"varint,4,opt,name=member_id,json=memberId,proto3" json:"member_id"`                //用户ID
	BusinessId    int64  `protobuf:"varint,5,opt,name=business_id,json=businessId,proto3" json:"business_id"`          // 业务订单id
	BusinessType  string `protobuf:"bytes,6,opt,name=business_type,json=businessType,proto3" json:"business_type"`     //订单类型: order 订单, wallet 钱包 vip 会员等级
	TradeNo       string `protobuf:"bytes,7,opt,name=trade_no,json=tradeNo,proto3" json:"trade_no"`                    //支付平台交易号
	BuyerUser     string `protobuf:"bytes,8,opt,name=buyer_user,json=buyerUser,proto3" json:"buyer_user"`              //支付平台用户帐号
	RefundPrice   int64  `protobuf:"varint,9,opt,name=refund_price,json=refundPrice,proto3" json:"refund_price"`       //退款金额
	PayPrice      int64  `protobuf:"varint,10,opt,name=pay_price,json=payPrice,proto3" json:"pay_price"`               //订单实际支付金额
	Reason        string `protobuf:"bytes,11,opt,name=reason,proto3" json:"reason"`                                    //原因描述
	PaymentId     int32  `protobuf:"varint,12,opt,name=payment_id,json=paymentId,proto3" json:"payment_id"`            //支付方式ID
	PaymentCode   string `protobuf:"bytes,13,opt,name=payment_code,json=paymentCode,proto3" json:"payment_code"`       //支付方式标识
	PaymentName   string `protobuf:"bytes,14,opt,name=payment_name,json=paymentName,proto3" json:"payment_name"`       //支付方式名称
	PayAccountId  int32  `protobuf:"varint,15,opt,name=pay_account_id,json=payAccountId,proto3" json:"pay_account_id"` //支付结算账号ID
	Refundment    string `protobuf:"bytes,16,opt,name=refundment,proto3" json:"refundment"`                            //退款类型（0原路退回, 1退至钱包, 2手动处理）
	RequestParams string `protobuf:"bytes,17,opt,name=request_params,json=requestParams,proto3" json:"request_params"` //支付平台返回参数（以json存储）
	Status        string `protobuf:"bytes,18,opt,name=status,proto3" json:"status"`                                    //退款状态：0待退款，1退款成功，2退款失败
	CreatedAt     string `protobuf:"bytes,19,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt     string `protobuf:"bytes,20,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
}

func (x *RefundLog) Reset() {
	*x = RefundLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_refundLogService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefundLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefundLog) ProtoMessage() {}

func (x *RefundLog) ProtoReflect() protoreflect.Message {
	mi := &file_refundLogService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefundLog.ProtoReflect.Descriptor instead.
func (*RefundLog) Descriptor() ([]byte, []int) {
	return file_refundLogService_proto_rawDescGZIP(), []int{0}
}

func (x *RefundLog) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RefundLog) GetLogNo() string {
	if x != nil {
		return x.LogNo
	}
	return ""
}

func (x *RefundLog) GetPayLogId() int64 {
	if x != nil {
		return x.PayLogId
	}
	return 0
}

func (x *RefundLog) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

func (x *RefundLog) GetBusinessId() int64 {
	if x != nil {
		return x.BusinessId
	}
	return 0
}

func (x *RefundLog) GetBusinessType() string {
	if x != nil {
		return x.BusinessType
	}
	return ""
}

func (x *RefundLog) GetTradeNo() string {
	if x != nil {
		return x.TradeNo
	}
	return ""
}

func (x *RefundLog) GetBuyerUser() string {
	if x != nil {
		return x.BuyerUser
	}
	return ""
}

func (x *RefundLog) GetRefundPrice() int64 {
	if x != nil {
		return x.RefundPrice
	}
	return 0
}

func (x *RefundLog) GetPayPrice() int64 {
	if x != nil {
		return x.PayPrice
	}
	return 0
}

func (x *RefundLog) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *RefundLog) GetPaymentId() int32 {
	if x != nil {
		return x.PaymentId
	}
	return 0
}

func (x *RefundLog) GetPaymentCode() string {
	if x != nil {
		return x.PaymentCode
	}
	return ""
}

func (x *RefundLog) GetPaymentName() string {
	if x != nil {
		return x.PaymentName
	}
	return ""
}

func (x *RefundLog) GetPayAccountId() int32 {
	if x != nil {
		return x.PayAccountId
	}
	return 0
}

func (x *RefundLog) GetRefundment() string {
	if x != nil {
		return x.Refundment
	}
	return ""
}

func (x *RefundLog) GetRequestParams() string {
	if x != nil {
		return x.RequestParams
	}
	return ""
}

func (x *RefundLog) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *RefundLog) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *RefundLog) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type RefundLogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int32 `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Top      int32 `protobuf:"varint,3,opt,name=top,proto3" json:"top"`
	//base params
	BusinessId   int64  `protobuf:"varint,4,opt,name=business_id,json=businessId,proto3" json:"business_id"`      // 业务订单id
	BusinessType string `protobuf:"bytes,5,opt,name=business_type,json=businessType,proto3" json:"business_type"` //订单类型: order 订单, wallet 钱包 vip 会员等级
	TradeNo      string `protobuf:"bytes,6,opt,name=trade_no,json=tradeNo,proto3" json:"trade_no"`                //支付平台交易号
	BuyerUser    string `protobuf:"bytes,7,opt,name=buyer_user,json=buyerUser,proto3" json:"buyer_user"`          //支付平台用户帐号
	PaymentId    int32  `protobuf:"varint,8,opt,name=payment_id,json=paymentId,proto3" json:"payment_id"`         //支付方式ID
	PaymentCode  string `protobuf:"bytes,9,opt,name=payment_code,json=paymentCode,proto3" json:"payment_code"`    //支付方式标识
	PaymentName  string `protobuf:"bytes,10,opt,name=payment_name,json=paymentName,proto3" json:"payment_name"`   //支付方式名称
	Refundment   string `protobuf:"bytes,11,opt,name=refundment,proto3" json:"refundment"`                        //退款类型（0原路退回, 1退至钱包, 2手动处理）
}

func (x *RefundLogRequest) Reset() {
	*x = RefundLogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_refundLogService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefundLogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefundLogRequest) ProtoMessage() {}

func (x *RefundLogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_refundLogService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefundLogRequest.ProtoReflect.Descriptor instead.
func (*RefundLogRequest) Descriptor() ([]byte, []int) {
	return file_refundLogService_proto_rawDescGZIP(), []int{1}
}

func (x *RefundLogRequest) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *RefundLogRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *RefundLogRequest) GetTop() int32 {
	if x != nil {
		return x.Top
	}
	return 0
}

func (x *RefundLogRequest) GetBusinessId() int64 {
	if x != nil {
		return x.BusinessId
	}
	return 0
}

func (x *RefundLogRequest) GetBusinessType() string {
	if x != nil {
		return x.BusinessType
	}
	return ""
}

func (x *RefundLogRequest) GetTradeNo() string {
	if x != nil {
		return x.TradeNo
	}
	return ""
}

func (x *RefundLogRequest) GetBuyerUser() string {
	if x != nil {
		return x.BuyerUser
	}
	return ""
}

func (x *RefundLogRequest) GetPaymentId() int32 {
	if x != nil {
		return x.PaymentId
	}
	return 0
}

func (x *RefundLogRequest) GetPaymentCode() string {
	if x != nil {
		return x.PaymentCode
	}
	return ""
}

func (x *RefundLogRequest) GetPaymentName() string {
	if x != nil {
		return x.PaymentName
	}
	return ""
}

func (x *RefundLogRequest) GetRefundment() string {
	if x != nil {
		return x.Refundment
	}
	return ""
}

type RefundLogData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *RefundLog    `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*RefundLog  `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Info   *common.Info  `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
}

func (x *RefundLogData) Reset() {
	*x = RefundLogData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_refundLogService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefundLogData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefundLogData) ProtoMessage() {}

func (x *RefundLogData) ProtoReflect() protoreflect.Message {
	mi := &file_refundLogService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefundLogData.ProtoReflect.Descriptor instead.
func (*RefundLogData) Descriptor() ([]byte, []int) {
	return file_refundLogService_proto_rawDescGZIP(), []int{2}
}

func (x *RefundLogData) GetEntity() *RefundLog {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *RefundLogData) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *RefundLogData) GetItems() []*RefundLog {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *RefundLogData) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type RefundLogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *RefundLogData `protobuf:"bytes,1,opt,name=data,proto3" json:"data"`
	Error *common.Error  `protobuf:"bytes,2,opt,name=error,proto3" json:"error"`
}

func (x *RefundLogResponse) Reset() {
	*x = RefundLogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_refundLogService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefundLogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefundLogResponse) ProtoMessage() {}

func (x *RefundLogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_refundLogService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefundLogResponse.ProtoReflect.Descriptor instead.
func (*RefundLogResponse) Descriptor() ([]byte, []int) {
	return file_refundLogService_proto_rawDescGZIP(), []int{3}
}

func (x *RefundLogResponse) GetData() *RefundLogData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *RefundLogResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_refundLogService_proto protoreflect.FileDescriptor

var file_refundLogService_proto_rawDesc = []byte{
	0x0a, 0x16, 0x72, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x4c, 0x6f, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xed, 0x04, 0x0a, 0x09, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64,
	0x4c, 0x6f, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x6c, 0x6f, 0x67, 0x5f, 0x6e, 0x6f, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x4e, 0x6f, 0x12, 0x1c, 0x0a, 0x0a, 0x70, 0x61,
	0x79, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x70, 0x61, 0x79, 0x4c, 0x6f, 0x67, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x62, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62,
	0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x74,
	0x72, 0x61, 0x64, 0x65, 0x5f, 0x6e, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74,
	0x72, 0x61, 0x64, 0x65, 0x4e, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x75, 0x79, 0x65, 0x72, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x75, 0x79, 0x65,
	0x72, 0x55, 0x73, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x5f,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x72, 0x65, 0x66,
	0x75, 0x6e, 0x64, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x79, 0x5f,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x79,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x70, 0x61, 0x79, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x70, 0x61, 0x79, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x66, 0x75,
	0x6e, 0x64, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65,
	0x66, 0x75, 0x6e, 0x64, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xdc, 0x02, 0x0a, 0x10, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64,
	0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61,
	0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x74, 0x6f, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x74, 0x6f, 0x70, 0x12,
	0x1f, 0x0a, 0x0b, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x49, 0x64,
	0x12, 0x23, 0x0a, 0x0d, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x6e,
	0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x72, 0x61, 0x64, 0x65, 0x4e, 0x6f,
	0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x75, 0x79, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x75, 0x79, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x21,
	0x0a, 0x0c, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x6d, 0x65,
	0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x66, 0x75, 0x6e, 0x64,
	0x6d, 0x65, 0x6e, 0x74, 0x22, 0xae, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x4c,
	0x6f, 0x67, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2b, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x4c, 0x6f, 0x67, 0x52, 0x06, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65,
	0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x4c, 0x6f, 0x67, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x65, 0x0a, 0x11, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x4c,
	0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x4c, 0x6f, 0x67, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x8e, 0x02, 0x0a,
	0x10, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x4c, 0x6f, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x3c, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x4c, 0x6f, 0x67,
	0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x66, 0x75,
	0x6e, 0x64, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x3c, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x4c, 0x6f, 0x67, 0x1a, 0x1b,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64,
	0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a,
	0x03, 0x47, 0x65, 0x74, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x4c, 0x6f, 0x67, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x4c, 0x6f, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x12, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65,
	0x66, 0x75, 0x6e, 0x64, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64,
	0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a,
	0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_refundLogService_proto_rawDescOnce sync.Once
	file_refundLogService_proto_rawDescData = file_refundLogService_proto_rawDesc
)

func file_refundLogService_proto_rawDescGZIP() []byte {
	file_refundLogService_proto_rawDescOnce.Do(func() {
		file_refundLogService_proto_rawDescData = protoimpl.X.CompressGZIP(file_refundLogService_proto_rawDescData)
	})
	return file_refundLogService_proto_rawDescData
}

var file_refundLogService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_refundLogService_proto_goTypes = []interface{}{
	(*RefundLog)(nil),         // 0: services.RefundLog
	(*RefundLogRequest)(nil),  // 1: services.RefundLogRequest
	(*RefundLogData)(nil),     // 2: services.RefundLogData
	(*RefundLogResponse)(nil), // 3: services.RefundLogResponse
	(*common.Pager)(nil),      // 4: common.Pager
	(*common.Info)(nil),       // 5: common.Info
	(*common.Error)(nil),      // 6: common.Error
}
var file_refundLogService_proto_depIdxs = []int32{
	0,  // 0: services.RefundLogData.entity:type_name -> services.RefundLog
	4,  // 1: services.RefundLogData.pager:type_name -> common.Pager
	0,  // 2: services.RefundLogData.items:type_name -> services.RefundLog
	5,  // 3: services.RefundLogData.info:type_name -> common.Info
	2,  // 4: services.RefundLogResponse.data:type_name -> services.RefundLogData
	6,  // 5: services.RefundLogResponse.error:type_name -> common.Error
	0,  // 6: services.RefundLogService.Create:input_type -> services.RefundLog
	0,  // 7: services.RefundLogService.Delete:input_type -> services.RefundLog
	0,  // 8: services.RefundLogService.Get:input_type -> services.RefundLog
	1,  // 9: services.RefundLogService.Search:input_type -> services.RefundLogRequest
	3,  // 10: services.RefundLogService.Create:output_type -> services.RefundLogResponse
	3,  // 11: services.RefundLogService.Delete:output_type -> services.RefundLogResponse
	3,  // 12: services.RefundLogService.Get:output_type -> services.RefundLogResponse
	3,  // 13: services.RefundLogService.Search:output_type -> services.RefundLogResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_refundLogService_proto_init() }
func file_refundLogService_proto_init() {
	if File_refundLogService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_refundLogService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefundLog); i {
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
		file_refundLogService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefundLogRequest); i {
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
		file_refundLogService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefundLogData); i {
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
		file_refundLogService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefundLogResponse); i {
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
			RawDescriptor: file_refundLogService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_refundLogService_proto_goTypes,
		DependencyIndexes: file_refundLogService_proto_depIdxs,
		MessageInfos:      file_refundLogService_proto_msgTypes,
	}.Build()
	File_refundLogService_proto = out.File
	file_refundLogService_proto_rawDesc = nil
	file_refundLogService_proto_goTypes = nil
	file_refundLogService_proto_depIdxs = nil
}