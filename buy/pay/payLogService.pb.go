// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: payLogService.proto

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

//支付日志
type PayLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64          `protobuf:"varint,1,opt,name=id,proto3" json:"id"`                                            //ID
	LogNo        string         `protobuf:"bytes,2,opt,name=log_no,json=logNo,proto3" json:"log_no"`                          //支付日志订单号
	BusinessType string         `protobuf:"bytes,3,opt,name=business_type,json=businessType,proto3" json:"business_type"`     //订单类型: order 订单, wallet 钱包 vip 会员等级
	MemberId     int64          `protobuf:"varint,4,opt,name=member_id,json=memberId,proto3" json:"member_id"`                //用户ID
	PaymentId    int32          `protobuf:"varint,5,opt,name=payment_id,json=paymentId,proto3" json:"payment_id"`             //支付方式ID
	PaymentCode  string         `protobuf:"bytes,6,opt,name=payment_code,json=paymentCode,proto3" json:"payment_code"`        //支付方式标识
	PaymentName  string         `protobuf:"bytes,7,opt,name=payment_name,json=paymentName,proto3" json:"payment_name"`        //支付方式名称
	Subject      string         `protobuf:"bytes,8,opt,name=subject,proto3" json:"subject"`                                   //订单名称
	TotalPrice   int64          `protobuf:"varint,9,opt,name=total_price,json=totalPrice,proto3" json:"total_price"`          //业务订单金额
	PayPrice     int64          `protobuf:"varint,10,opt,name=pay_price,json=payPrice,proto3" json:"pay_price"`               //支付金额
	Currency     string         `protobuf:"bytes,11,opt,name=currency,proto3" json:"currency"`                                //货币代码
	TradeNo      string         `protobuf:"bytes,12,opt,name=trade_no,json=tradeNo,proto3" json:"trade_no"`                   //支付平台交易号
	BuyerUser    string         `protobuf:"bytes,13,opt,name=buyer_user,json=buyerUser,proto3" json:"buyer_user"`             //支付平台用户帐号
	PayAccountId int32          `protobuf:"varint,14,opt,name=pay_account_id,json=payAccountId,proto3" json:"pay_account_id"` //支付结算账号ID
	Status       string         `protobuf:"bytes,15,opt,name=status,proto3" json:"status"`                                    //状态（0待支付, 1已支付, 2已关闭）正常30分钟内未支付将关闭
	PaidTime     string         `protobuf:"bytes,16,opt,name=paid_time,json=paidTime,proto3" json:"paid_time"`                //支付时间
	ClosedTime   string         `protobuf:"bytes,17,opt,name=closed_time,json=closedTime,proto3" json:"closed_time"`          //关闭时间
	CreatedAt    string         `protobuf:"bytes,18,opt,name=created_at,json=createdAt,proto3" json:"created_at"`             //创建时间
	UpdatedAt    string         `protobuf:"bytes,19,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`             //修改时间
	Values       []*PayLogValue `protobuf:"bytes,20,rep,name=values,proto3" json:"values"`                                    //支付日志关联业务数据
}

func (x *PayLog) Reset() {
	*x = PayLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payLogService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayLog) ProtoMessage() {}

func (x *PayLog) ProtoReflect() protoreflect.Message {
	mi := &file_payLogService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayLog.ProtoReflect.Descriptor instead.
func (*PayLog) Descriptor() ([]byte, []int) {
	return file_payLogService_proto_rawDescGZIP(), []int{0}
}

func (x *PayLog) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PayLog) GetLogNo() string {
	if x != nil {
		return x.LogNo
	}
	return ""
}

func (x *PayLog) GetBusinessType() string {
	if x != nil {
		return x.BusinessType
	}
	return ""
}

func (x *PayLog) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

func (x *PayLog) GetPaymentId() int32 {
	if x != nil {
		return x.PaymentId
	}
	return 0
}

func (x *PayLog) GetPaymentCode() string {
	if x != nil {
		return x.PaymentCode
	}
	return ""
}

func (x *PayLog) GetPaymentName() string {
	if x != nil {
		return x.PaymentName
	}
	return ""
}

func (x *PayLog) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *PayLog) GetTotalPrice() int64 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

func (x *PayLog) GetPayPrice() int64 {
	if x != nil {
		return x.PayPrice
	}
	return 0
}

func (x *PayLog) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *PayLog) GetTradeNo() string {
	if x != nil {
		return x.TradeNo
	}
	return ""
}

func (x *PayLog) GetBuyerUser() string {
	if x != nil {
		return x.BuyerUser
	}
	return ""
}

func (x *PayLog) GetPayAccountId() int32 {
	if x != nil {
		return x.PayAccountId
	}
	return 0
}

func (x *PayLog) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *PayLog) GetPaidTime() string {
	if x != nil {
		return x.PaidTime
	}
	return ""
}

func (x *PayLog) GetClosedTime() string {
	if x != nil {
		return x.ClosedTime
	}
	return ""
}

func (x *PayLog) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *PayLog) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *PayLog) GetValues() []*PayLogValue {
	if x != nil {
		return x.Values
	}
	return nil
}

//支付日志关联业务数据
type PayLogValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`                                   //ID
	PayLogId   int64  `protobuf:"varint,2,opt,name=pay_log_id,json=payLogId,proto3" json:"pay_log_id"`     //支付日志id
	BusinessId int64  `protobuf:"varint,3,opt,name=business_id,json=businessId,proto3" json:"business_id"` //业务订单id
	BusinessNo string `protobuf:"bytes,4,opt,name=business_no,json=businessNo,proto3" json:"business_no"`  //业务订单号
	CreatedAt  string `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at"`     //创建时间
	UpdatedAt  string `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`     //修改时间
}

func (x *PayLogValue) Reset() {
	*x = PayLogValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payLogService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayLogValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayLogValue) ProtoMessage() {}

func (x *PayLogValue) ProtoReflect() protoreflect.Message {
	mi := &file_payLogService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayLogValue.ProtoReflect.Descriptor instead.
func (*PayLogValue) Descriptor() ([]byte, []int) {
	return file_payLogService_proto_rawDescGZIP(), []int{1}
}

func (x *PayLogValue) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PayLogValue) GetPayLogId() int64 {
	if x != nil {
		return x.PayLogId
	}
	return 0
}

func (x *PayLogValue) GetBusinessId() int64 {
	if x != nil {
		return x.BusinessId
	}
	return 0
}

func (x *PayLogValue) GetBusinessNo() string {
	if x != nil {
		return x.BusinessNo
	}
	return ""
}

func (x *PayLogValue) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *PayLogValue) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type PayLogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int32 `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Top      int32 `protobuf:"varint,3,opt,name=top,proto3" json:"top"`
	//base params
	Id           int64  `protobuf:"varint,4,opt,name=id,proto3" json:"id"`                                        //ID
	LogNo        string `protobuf:"bytes,5,opt,name=log_no,json=logNo,proto3" json:"log_no"`                      //支付日志订单号
	BusinessType string `protobuf:"bytes,6,opt,name=business_type,json=businessType,proto3" json:"business_type"` //订单类型: order 订单, wallet 钱包 vip 会员等级
	MemberId     int64  `protobuf:"varint,7,opt,name=member_id,json=memberId,proto3" json:"member_id"`            //用户ID
	PaymentId    int32  `protobuf:"varint,8,opt,name=payment_id,json=paymentId,proto3" json:"payment_id"`         //支付方式ID
	PaymentCode  string `protobuf:"bytes,9,opt,name=payment_code,json=paymentCode,proto3" json:"payment_code"`    //支付方式标识
	PaymentName  string `protobuf:"bytes,10,opt,name=payment_name,json=paymentName,proto3" json:"payment_name"`   //支付方式名称
	Subject      string `protobuf:"bytes,11,opt,name=subject,proto3" json:"subject"`                              //订单名称
	TradeNo      string `protobuf:"bytes,12,opt,name=trade_no,json=tradeNo,proto3" json:"trade_no"`               //支付平台交易号
	BuyerUser    string `protobuf:"bytes,13,opt,name=buyer_user,json=buyerUser,proto3" json:"buyer_user"`         //支付平台用户帐号
	BusinessId   int64  `protobuf:"varint,14,opt,name=business_id,json=businessId,proto3" json:"business_id"`     //业务订单ID
	BusinessNo   string `protobuf:"bytes,15,opt,name=business_no,json=businessNo,proto3" json:"business_no"`      //业务订单编号
	Status       string `protobuf:"bytes,16,opt,name=status,proto3" json:"status"`                                //状态
	BeginTime    string `protobuf:"bytes,17,opt,name=begin_time,json=beginTime,proto3" json:"begin_time"`         //开始时间
	EndTime      string `protobuf:"bytes,18,opt,name=end_time,json=endTime,proto3" json:"end_time"`               //结束时间
}

func (x *PayLogRequest) Reset() {
	*x = PayLogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payLogService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayLogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayLogRequest) ProtoMessage() {}

func (x *PayLogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payLogService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayLogRequest.ProtoReflect.Descriptor instead.
func (*PayLogRequest) Descriptor() ([]byte, []int) {
	return file_payLogService_proto_rawDescGZIP(), []int{2}
}

func (x *PayLogRequest) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *PayLogRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *PayLogRequest) GetTop() int32 {
	if x != nil {
		return x.Top
	}
	return 0
}

func (x *PayLogRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PayLogRequest) GetLogNo() string {
	if x != nil {
		return x.LogNo
	}
	return ""
}

func (x *PayLogRequest) GetBusinessType() string {
	if x != nil {
		return x.BusinessType
	}
	return ""
}

func (x *PayLogRequest) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

func (x *PayLogRequest) GetPaymentId() int32 {
	if x != nil {
		return x.PaymentId
	}
	return 0
}

func (x *PayLogRequest) GetPaymentCode() string {
	if x != nil {
		return x.PaymentCode
	}
	return ""
}

func (x *PayLogRequest) GetPaymentName() string {
	if x != nil {
		return x.PaymentName
	}
	return ""
}

func (x *PayLogRequest) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *PayLogRequest) GetTradeNo() string {
	if x != nil {
		return x.TradeNo
	}
	return ""
}

func (x *PayLogRequest) GetBuyerUser() string {
	if x != nil {
		return x.BuyerUser
	}
	return ""
}

func (x *PayLogRequest) GetBusinessId() int64 {
	if x != nil {
		return x.BusinessId
	}
	return 0
}

func (x *PayLogRequest) GetBusinessNo() string {
	if x != nil {
		return x.BusinessNo
	}
	return ""
}

func (x *PayLogRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *PayLogRequest) GetBeginTime() string {
	if x != nil {
		return x.BeginTime
	}
	return ""
}

func (x *PayLogRequest) GetEndTime() string {
	if x != nil {
		return x.EndTime
	}
	return ""
}

type PayLogData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *PayLog       `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*PayLog     `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Info   *common.Info  `protobuf:"bytes,4,opt,name=info,proto3" json:"info"`
}

func (x *PayLogData) Reset() {
	*x = PayLogData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payLogService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayLogData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayLogData) ProtoMessage() {}

func (x *PayLogData) ProtoReflect() protoreflect.Message {
	mi := &file_payLogService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayLogData.ProtoReflect.Descriptor instead.
func (*PayLogData) Descriptor() ([]byte, []int) {
	return file_payLogService_proto_rawDescGZIP(), []int{3}
}

func (x *PayLogData) GetEntity() *PayLog {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *PayLogData) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *PayLogData) GetItems() []*PayLog {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *PayLogData) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type PayLogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *PayLogData   `protobuf:"bytes,1,opt,name=data,proto3" json:"data"`
	Error *common.Error `protobuf:"bytes,2,opt,name=error,proto3" json:"error"`
}

func (x *PayLogResponse) Reset() {
	*x = PayLogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payLogService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayLogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayLogResponse) ProtoMessage() {}

func (x *PayLogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payLogService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayLogResponse.ProtoReflect.Descriptor instead.
func (*PayLogResponse) Descriptor() ([]byte, []int) {
	return file_payLogService_proto_rawDescGZIP(), []int{4}
}

func (x *PayLogResponse) GetData() *PayLogData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *PayLogResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_payLogService_proto protoreflect.FileDescriptor

var file_payLogService_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x61, 0x79, 0x4c, 0x6f, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a,
	0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xed, 0x04, 0x0a, 0x06, 0x50, 0x61, 0x79, 0x4c, 0x6f, 0x67, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x15, 0x0a,
	0x06, 0x6c, 0x6f, 0x67, 0x5f, 0x6e, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c,
	0x6f, 0x67, 0x4e, 0x6f, 0x12, 0x23, 0x0a, 0x0d, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x79, 0x5f, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x79, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12,
	0x19, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x6e, 0x6f, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x74, 0x72, 0x61, 0x64, 0x65, 0x4e, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x75,
	0x79, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x62, 0x75, 0x79, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x0e, 0x70, 0x61, 0x79,
	0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0c, 0x70, 0x61, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x69, 0x64, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x69, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6c, 0x6f, 0x73, 0x65,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x2d, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x14, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50,
	0x61, 0x79, 0x4c, 0x6f, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x22, 0xbb, 0x01, 0x0a, 0x0b, 0x50, 0x61, 0x79, 0x4c, 0x6f, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1c, 0x0a, 0x0a, 0x70, 0x61, 0x79, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x79, 0x4c, 0x6f, 0x67, 0x49, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x49,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x6e, 0x6f,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x4e, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x22, 0x8a, 0x04, 0x0a, 0x0d, 0x50, 0x61, 0x79, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x6f, 0x70, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x03, 0x74, 0x6f, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x6c, 0x6f, 0x67, 0x5f, 0x6e,
	0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x4e, 0x6f, 0x12, 0x23,
	0x0a, 0x0d, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x21, 0x0a, 0x0c, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12,
	0x19, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x6e, 0x6f, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x74, 0x72, 0x61, 0x64, 0x65, 0x4e, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x75,
	0x79, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x62, 0x75, 0x79, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x75,
	0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x6e, 0x6f, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x4e, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x12,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xa5, 0x01,
	0x0a, 0x0a, 0x50, 0x61, 0x79, 0x4c, 0x6f, 0x67, 0x44, 0x61, 0x74, 0x61, 0x12, 0x28, 0x0a, 0x06,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x4c, 0x6f, 0x67, 0x52, 0x06,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50,
	0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x4c, 0x6f, 0x67, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x5f, 0x0a, 0x0e, 0x50, 0x61, 0x79, 0x4c, 0x6f, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x50, 0x61, 0x79, 0x4c, 0x6f, 0x67, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xac, 0x03, 0x0a, 0x0d, 0x50, 0x61, 0x79, 0x4c, 0x6f,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x06, 0x49, 0x6e, 0x73, 0x65,
	0x72, 0x74, 0x12, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x61,
	0x79, 0x4c, 0x6f, 0x67, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x50, 0x61, 0x79, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x37, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x10, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x4c, 0x6f, 0x67, 0x1a, 0x18, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x4c, 0x6f, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x05, 0x43, 0x6c, 0x6f,
	0x73, 0x65, 0x12, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x61,
	0x79, 0x4c, 0x6f, 0x67, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x50, 0x61, 0x79, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x38, 0x0a, 0x08, 0x54, 0x79, 0x70, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x10, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x4c, 0x6f, 0x67, 0x1a, 0x18,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x4c, 0x6f, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x06, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x50, 0x61, 0x79, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x4c, 0x6f, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x04, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x79,
	0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x4c,
	0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_payLogService_proto_rawDescOnce sync.Once
	file_payLogService_proto_rawDescData = file_payLogService_proto_rawDesc
)

func file_payLogService_proto_rawDescGZIP() []byte {
	file_payLogService_proto_rawDescOnce.Do(func() {
		file_payLogService_proto_rawDescData = protoimpl.X.CompressGZIP(file_payLogService_proto_rawDescData)
	})
	return file_payLogService_proto_rawDescData
}

var file_payLogService_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_payLogService_proto_goTypes = []interface{}{
	(*PayLog)(nil),         // 0: services.PayLog
	(*PayLogValue)(nil),    // 1: services.PayLogValue
	(*PayLogRequest)(nil),  // 2: services.PayLogRequest
	(*PayLogData)(nil),     // 3: services.PayLogData
	(*PayLogResponse)(nil), // 4: services.PayLogResponse
	(*common.Pager)(nil),   // 5: common.Pager
	(*common.Info)(nil),    // 6: common.Info
	(*common.Error)(nil),   // 7: common.Error
}
var file_payLogService_proto_depIdxs = []int32{
	1,  // 0: services.PayLog.values:type_name -> services.PayLogValue
	0,  // 1: services.PayLogData.entity:type_name -> services.PayLog
	5,  // 2: services.PayLogData.pager:type_name -> common.Pager
	0,  // 3: services.PayLogData.items:type_name -> services.PayLog
	6,  // 4: services.PayLogData.info:type_name -> common.Info
	3,  // 5: services.PayLogResponse.data:type_name -> services.PayLogData
	7,  // 6: services.PayLogResponse.error:type_name -> common.Error
	0,  // 7: services.PayLogService.Insert:input_type -> services.PayLog
	0,  // 8: services.PayLogService.Success:input_type -> services.PayLog
	0,  // 9: services.PayLogService.Close:input_type -> services.PayLog
	0,  // 10: services.PayLogService.TypeList:input_type -> services.PayLog
	2,  // 11: services.PayLogService.Detail:input_type -> services.PayLogRequest
	2,  // 12: services.PayLogService.List:input_type -> services.PayLogRequest
	2,  // 13: services.PayLogService.Search:input_type -> services.PayLogRequest
	4,  // 14: services.PayLogService.Insert:output_type -> services.PayLogResponse
	4,  // 15: services.PayLogService.Success:output_type -> services.PayLogResponse
	4,  // 16: services.PayLogService.Close:output_type -> services.PayLogResponse
	4,  // 17: services.PayLogService.TypeList:output_type -> services.PayLogResponse
	4,  // 18: services.PayLogService.Detail:output_type -> services.PayLogResponse
	4,  // 19: services.PayLogService.List:output_type -> services.PayLogResponse
	4,  // 20: services.PayLogService.Search:output_type -> services.PayLogResponse
	14, // [14:21] is the sub-list for method output_type
	7,  // [7:14] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_payLogService_proto_init() }
func file_payLogService_proto_init() {
	if File_payLogService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_payLogService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayLog); i {
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
		file_payLogService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayLogValue); i {
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
		file_payLogService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayLogRequest); i {
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
		file_payLogService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayLogData); i {
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
		file_payLogService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayLogResponse); i {
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
			RawDescriptor: file_payLogService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_payLogService_proto_goTypes,
		DependencyIndexes: file_payLogService_proto_depIdxs,
		MessageInfos:      file_payLogService_proto_msgTypes,
	}.Build()
	File_payLogService_proto = out.File
	file_payLogService_proto_rawDesc = nil
	file_payLogService_proto_goTypes = nil
	file_payLogService_proto_depIdxs = nil
}
