// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: levelOrderLogService.proto

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

type LevelOrderLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id"`                                               //ID
	LogNo          string  `protobuf:"bytes,2,opt,name=log_no,json=logNo,proto3" json:"log_no"`                             //支付订单编号
	MemberId       int64   `protobuf:"varint,3,opt,name=member_id,json=memberId,proto3" json:"member_id"`                   //用户ID
	LevelId        int64   `protobuf:"varint,4,opt,name=level_id,json=levelId,proto3" json:"level_id"`                      //LevelID
	Money          int64   `protobuf:"varint,5,opt,name=money,proto3" json:"money"`                                         //销售价
	PayMoney       int64   `protobuf:"varint,6,opt,name=pay_money,json=payMoney,proto3" json:"pay_money"`                   //支付金额
	Type           string  `protobuf:"bytes,7,opt,name=type,proto3" json:"type"`                                            //类型：buy正常购买,give赠送
	PeriodUnit     string  `protobuf:"bytes,8,opt,name=period_unit,json=periodUnit,proto3" json:"period_unit"`              //周期单位：年,季,月,天
	PeriodValue    int32   `protobuf:"varint,9,opt,name=period_value,json=periodValue,proto3" json:"period_value"`          //周期值
	IsContinuation string  `protobuf:"bytes,10,opt,name=is_continuation,json=isContinuation,proto3" json:"is_continuation"` //是否为连续性：如连续包月
	PayTime        string  `protobuf:"bytes,11,opt,name=pay_time,json=payTime,proto3" json:"pay_time"`                      //支付时间
	PaymentId      int32   `protobuf:"varint,12,opt,name=payment_id,json=paymentId,proto3" json:"payment_id"`               //支付方式Id
	Status         string  `protobuf:"bytes,13,opt,name=status,proto3" json:"status"`                                       //状态：0待支付，1已支付，2已失效
	SettleStatus   string  `protobuf:"bytes,14,opt,name=settle_status,json=settleStatus,proto3" json:"settle_status"`       //结算状态：0待支付，1已结算，2结算失败
	CreatedAt      string  `protobuf:"bytes,15,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt      string  `protobuf:"bytes,16,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	Member         *Member `protobuf:"bytes,17,opt,name=member,proto3" json:"member"`
}

func (x *LevelOrderLog) Reset() {
	*x = LevelOrderLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_levelOrderLogService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LevelOrderLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LevelOrderLog) ProtoMessage() {}

func (x *LevelOrderLog) ProtoReflect() protoreflect.Message {
	mi := &file_levelOrderLogService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LevelOrderLog.ProtoReflect.Descriptor instead.
func (*LevelOrderLog) Descriptor() ([]byte, []int) {
	return file_levelOrderLogService_proto_rawDescGZIP(), []int{0}
}

func (x *LevelOrderLog) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *LevelOrderLog) GetLogNo() string {
	if x != nil {
		return x.LogNo
	}
	return ""
}

func (x *LevelOrderLog) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

func (x *LevelOrderLog) GetLevelId() int64 {
	if x != nil {
		return x.LevelId
	}
	return 0
}

func (x *LevelOrderLog) GetMoney() int64 {
	if x != nil {
		return x.Money
	}
	return 0
}

func (x *LevelOrderLog) GetPayMoney() int64 {
	if x != nil {
		return x.PayMoney
	}
	return 0
}

func (x *LevelOrderLog) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *LevelOrderLog) GetPeriodUnit() string {
	if x != nil {
		return x.PeriodUnit
	}
	return ""
}

func (x *LevelOrderLog) GetPeriodValue() int32 {
	if x != nil {
		return x.PeriodValue
	}
	return 0
}

func (x *LevelOrderLog) GetIsContinuation() string {
	if x != nil {
		return x.IsContinuation
	}
	return ""
}

func (x *LevelOrderLog) GetPayTime() string {
	if x != nil {
		return x.PayTime
	}
	return ""
}

func (x *LevelOrderLog) GetPaymentId() int32 {
	if x != nil {
		return x.PaymentId
	}
	return 0
}

func (x *LevelOrderLog) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *LevelOrderLog) GetSettleStatus() string {
	if x != nil {
		return x.SettleStatus
	}
	return ""
}

func (x *LevelOrderLog) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *LevelOrderLog) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *LevelOrderLog) GetMember() *Member {
	if x != nil {
		return x.Member
	}
	return nil
}

//查询参数
type LevelOrderLogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int64 `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	//以下为自定义参数
	MemberId   int64  `protobuf:"varint,3,opt,name=member_id,json=memberId,proto3" json:"member_id"`
	Type       string `protobuf:"bytes,4,opt,name=type,proto3" json:"type"`
	LevelId    int32  `protobuf:"varint,5,opt,name=level_id,json=levelId,proto3" json:"level_id"`
	Money      int64  `protobuf:"varint,6,opt,name=money,proto3" json:"money"`
	PaymentId  int32  `protobuf:"varint,7,opt,name=payment_id,json=paymentId,proto3" json:"payment_id"`
	LogNo      int64  `protobuf:"varint,8,opt,name=log_no,json=logNo,proto3" json:"log_no"`
	PeriodUnit string `protobuf:"bytes,9,opt,name=period_unit,json=periodUnit,proto3" json:"period_unit"`
	Memo       string `protobuf:"bytes,10,opt,name=memo,proto3" json:"memo"`
	StartDate  string `protobuf:"bytes,11,opt,name=start_date,json=startDate,proto3" json:"start_date"`
	EndDate    string `protobuf:"bytes,12,opt,name=end_date,json=endDate,proto3" json:"end_date"`
}

func (x *LevelOrderLogRequest) Reset() {
	*x = LevelOrderLogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_levelOrderLogService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LevelOrderLogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LevelOrderLogRequest) ProtoMessage() {}

func (x *LevelOrderLogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_levelOrderLogService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LevelOrderLogRequest.ProtoReflect.Descriptor instead.
func (*LevelOrderLogRequest) Descriptor() ([]byte, []int) {
	return file_levelOrderLogService_proto_rawDescGZIP(), []int{1}
}

func (x *LevelOrderLogRequest) GetPaged() int64 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *LevelOrderLogRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *LevelOrderLogRequest) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

func (x *LevelOrderLogRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *LevelOrderLogRequest) GetLevelId() int32 {
	if x != nil {
		return x.LevelId
	}
	return 0
}

func (x *LevelOrderLogRequest) GetMoney() int64 {
	if x != nil {
		return x.Money
	}
	return 0
}

func (x *LevelOrderLogRequest) GetPaymentId() int32 {
	if x != nil {
		return x.PaymentId
	}
	return 0
}

func (x *LevelOrderLogRequest) GetLogNo() int64 {
	if x != nil {
		return x.LogNo
	}
	return 0
}

func (x *LevelOrderLogRequest) GetPeriodUnit() string {
	if x != nil {
		return x.PeriodUnit
	}
	return ""
}

func (x *LevelOrderLogRequest) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *LevelOrderLogRequest) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *LevelOrderLogRequest) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

type LevelOrderLogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *LevelOrderLog   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager    `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*LevelOrderLog `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Info   string           `protobuf:"bytes,4,opt,name=info,proto3" json:"info"`
}

func (x *LevelOrderLogResponse) Reset() {
	*x = LevelOrderLogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_levelOrderLogService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LevelOrderLogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LevelOrderLogResponse) ProtoMessage() {}

func (x *LevelOrderLogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_levelOrderLogService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LevelOrderLogResponse.ProtoReflect.Descriptor instead.
func (*LevelOrderLogResponse) Descriptor() ([]byte, []int) {
	return file_levelOrderLogService_proto_rawDescGZIP(), []int{2}
}

func (x *LevelOrderLogResponse) GetEntity() *LevelOrderLog {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *LevelOrderLogResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *LevelOrderLogResponse) GetItems() []*LevelOrderLog {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *LevelOrderLogResponse) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

var File_levelOrderLogService_proto protoreflect.FileDescriptor

var file_levelOrderLogService_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x81,
	0x04, 0x0a, 0x0d, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x6f, 0x67,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x15, 0x0a, 0x06, 0x6c, 0x6f, 0x67, 0x5f, 0x6e, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6c, 0x6f, 0x67, 0x4e, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x79, 0x5f, 0x6d, 0x6f, 0x6e,
	0x65, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x79, 0x4d, 0x6f, 0x6e,
	0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64,
	0x5f, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x65, 0x72,
	0x69, 0x6f, 0x64, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x65, 0x72, 0x69, 0x6f,
	0x64, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x70,
	0x65, 0x72, 0x69, 0x6f, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x69, 0x73,
	0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x69, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65,
	0x74, 0x74, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x28, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x22, 0xd0, 0x02, 0x0a, 0x14, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f,
	0x6e, 0x65, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x65, 0x79,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x15, 0x0a, 0x06, 0x6c, 0x6f, 0x67, 0x5f, 0x6e, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x6c, 0x6f, 0x67, 0x4e, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64,
	0x5f, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x65, 0x72,
	0x69, 0x6f, 0x64, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e,
	0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e,
	0x64, 0x44, 0x61, 0x74, 0x65, 0x22, 0xb0, 0x01, 0x0a, 0x15, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2f, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05,
	0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x52, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0xb2, 0x02, 0x0a, 0x14, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x44, 0x0a, 0x06, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x12, 0x17, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x4c, 0x6f, 0x67, 0x1a, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x06, 0x45, 0x78, 0x70, 0x65, 0x6e,
	0x64, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x1a, 0x1f, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a,
	0x03, 0x47, 0x65, 0x74, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x1a, 0x1f, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x4b, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1e, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a,
	0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_levelOrderLogService_proto_rawDescOnce sync.Once
	file_levelOrderLogService_proto_rawDescData = file_levelOrderLogService_proto_rawDesc
)

func file_levelOrderLogService_proto_rawDescGZIP() []byte {
	file_levelOrderLogService_proto_rawDescOnce.Do(func() {
		file_levelOrderLogService_proto_rawDescData = protoimpl.X.CompressGZIP(file_levelOrderLogService_proto_rawDescData)
	})
	return file_levelOrderLogService_proto_rawDescData
}

var file_levelOrderLogService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_levelOrderLogService_proto_goTypes = []interface{}{
	(*LevelOrderLog)(nil),         // 0: services.LevelOrderLog
	(*LevelOrderLogRequest)(nil),  // 1: services.LevelOrderLogRequest
	(*LevelOrderLogResponse)(nil), // 2: services.LevelOrderLogResponse
	(*Member)(nil),                // 3: services.Member
	(*common.Pager)(nil),          // 4: common.Pager
}
var file_levelOrderLogService_proto_depIdxs = []int32{
	3, // 0: services.LevelOrderLog.member:type_name -> services.Member
	0, // 1: services.LevelOrderLogResponse.entity:type_name -> services.LevelOrderLog
	4, // 2: services.LevelOrderLogResponse.pager:type_name -> common.Pager
	0, // 3: services.LevelOrderLogResponse.items:type_name -> services.LevelOrderLog
	0, // 4: services.LevelOrderLogService.Income:input_type -> services.LevelOrderLog
	0, // 5: services.LevelOrderLogService.Expend:input_type -> services.LevelOrderLog
	0, // 6: services.LevelOrderLogService.Get:input_type -> services.LevelOrderLog
	1, // 7: services.LevelOrderLogService.Search:input_type -> services.LevelOrderLogRequest
	2, // 8: services.LevelOrderLogService.Income:output_type -> services.LevelOrderLogResponse
	2, // 9: services.LevelOrderLogService.Expend:output_type -> services.LevelOrderLogResponse
	2, // 10: services.LevelOrderLogService.Get:output_type -> services.LevelOrderLogResponse
	2, // 11: services.LevelOrderLogService.Search:output_type -> services.LevelOrderLogResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_levelOrderLogService_proto_init() }
func file_levelOrderLogService_proto_init() {
	if File_levelOrderLogService_proto != nil {
		return
	}
	file_memberService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_levelOrderLogService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LevelOrderLog); i {
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
		file_levelOrderLogService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LevelOrderLogRequest); i {
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
		file_levelOrderLogService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LevelOrderLogResponse); i {
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
			RawDescriptor: file_levelOrderLogService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_levelOrderLogService_proto_goTypes,
		DependencyIndexes: file_levelOrderLogService_proto_depIdxs,
		MessageInfos:      file_levelOrderLogService_proto_msgTypes,
	}.Build()
	File_levelOrderLogService_proto = out.File
	file_levelOrderLogService_proto_rawDesc = nil
	file_levelOrderLogService_proto_goTypes = nil
	file_levelOrderLogService_proto_depIdxs = nil
}
