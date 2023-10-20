// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: memberPayLevelService.proto

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

type MemberPayLevel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`                                                          //ID
	MemberId          int64  `protobuf:"varint,2,opt,name=member_id,json=memberId,proto3" json:"member_id"`                              //用户ID
	LevelId           int32  `protobuf:"varint,3,opt,name=level_id,json=levelId,proto3" json:"level_id"`                                 //等级ID
	BeginTime         string `protobuf:"bytes,4,opt,name=begin_time,json=beginTime,proto3" json:"begin_time"`                            //开始时间
	EndTime           string `protobuf:"bytes,5,opt,name=end_time,json=endTime,proto3" json:"end_time"`                                  //结束时间
	OrderTradeNum     int32  `protobuf:"varint,6,opt,name=order_trade_num,json=orderTradeNum,proto3" json:"order_trade_num"`             //成交订单笔数
	OrderTradeMoney   int64  `protobuf:"varint,7,opt,name=order_trade_money,json=orderTradeMoney,proto3" json:"order_trade_money"`       //成交订单总额
	RenewNum          int32  `protobuf:"varint,8,opt,name=renew_num,json=renewNum,proto3" json:"renew_num"`                              //续展次数
	SurplusTimeNumber int32  `protobuf:"varint,9,opt,name=surplus_time_number,json=surplusTimeNumber,proto3" json:"surplus_time_number"` //剩余时间值
	SurplusTimeUnit   string `protobuf:"bytes,10,opt,name=surplus_time_unit,json=surplusTimeUnit,proto3" json:"surplus_time_unit"`       //剩余时间单位
	IsForever         string `protobuf:"bytes,11,opt,name=is_forever,json=isForever,proto3" json:"is_forever"`                           //是否永久会员(0否，1是)
	Status            string `protobuf:"bytes,12,opt,name=status,proto3" json:"status"`                                                  //状态：0失效，1正常
	CreatedAt         string `protobuf:"bytes,13,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt         string `protobuf:"bytes,14,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	Level             *Level `protobuf:"bytes,15,opt,name=level,proto3" json:"level"`
}

func (x *MemberPayLevel) Reset() {
	*x = MemberPayLevel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memberPayLevelService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberPayLevel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberPayLevel) ProtoMessage() {}

func (x *MemberPayLevel) ProtoReflect() protoreflect.Message {
	mi := &file_memberPayLevelService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberPayLevel.ProtoReflect.Descriptor instead.
func (*MemberPayLevel) Descriptor() ([]byte, []int) {
	return file_memberPayLevelService_proto_rawDescGZIP(), []int{0}
}

func (x *MemberPayLevel) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MemberPayLevel) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

func (x *MemberPayLevel) GetLevelId() int32 {
	if x != nil {
		return x.LevelId
	}
	return 0
}

func (x *MemberPayLevel) GetBeginTime() string {
	if x != nil {
		return x.BeginTime
	}
	return ""
}

func (x *MemberPayLevel) GetEndTime() string {
	if x != nil {
		return x.EndTime
	}
	return ""
}

func (x *MemberPayLevel) GetOrderTradeNum() int32 {
	if x != nil {
		return x.OrderTradeNum
	}
	return 0
}

func (x *MemberPayLevel) GetOrderTradeMoney() int64 {
	if x != nil {
		return x.OrderTradeMoney
	}
	return 0
}

func (x *MemberPayLevel) GetRenewNum() int32 {
	if x != nil {
		return x.RenewNum
	}
	return 0
}

func (x *MemberPayLevel) GetSurplusTimeNumber() int32 {
	if x != nil {
		return x.SurplusTimeNumber
	}
	return 0
}

func (x *MemberPayLevel) GetSurplusTimeUnit() string {
	if x != nil {
		return x.SurplusTimeUnit
	}
	return ""
}

func (x *MemberPayLevel) GetIsForever() string {
	if x != nil {
		return x.IsForever
	}
	return ""
}

func (x *MemberPayLevel) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *MemberPayLevel) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *MemberPayLevel) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *MemberPayLevel) GetLevel() *Level {
	if x != nil {
		return x.Level
	}
	return nil
}

type MemberPayLevelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int64 `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	//以下为自定义参数
	Id       int32  `protobuf:"varint,3,opt,name=id,proto3" json:"id"`
	LevelId  int32  `protobuf:"varint,4,opt,name=level_id,json=levelId,proto3" json:"level_id"`
	MemberId int64  `protobuf:"varint,5,opt,name=member_id,json=memberId,proto3" json:"member_id"`
	Status   string `protobuf:"bytes,6,opt,name=status,proto3" json:"status"`
}

func (x *MemberPayLevelRequest) Reset() {
	*x = MemberPayLevelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memberPayLevelService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberPayLevelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberPayLevelRequest) ProtoMessage() {}

func (x *MemberPayLevelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_memberPayLevelService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberPayLevelRequest.ProtoReflect.Descriptor instead.
func (*MemberPayLevelRequest) Descriptor() ([]byte, []int) {
	return file_memberPayLevelService_proto_rawDescGZIP(), []int{1}
}

func (x *MemberPayLevelRequest) GetPaged() int64 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *MemberPayLevelRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *MemberPayLevelRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MemberPayLevelRequest) GetLevelId() int32 {
	if x != nil {
		return x.LevelId
	}
	return 0
}

func (x *MemberPayLevelRequest) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

func (x *MemberPayLevelRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type MemberPayLevelData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *MemberPayLevel   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager     `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*MemberPayLevel `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Info   *common.Info      `protobuf:"bytes,4,opt,name=info,proto3" json:"info"`
}

func (x *MemberPayLevelData) Reset() {
	*x = MemberPayLevelData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memberPayLevelService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberPayLevelData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberPayLevelData) ProtoMessage() {}

func (x *MemberPayLevelData) ProtoReflect() protoreflect.Message {
	mi := &file_memberPayLevelService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberPayLevelData.ProtoReflect.Descriptor instead.
func (*MemberPayLevelData) Descriptor() ([]byte, []int) {
	return file_memberPayLevelService_proto_rawDescGZIP(), []int{2}
}

func (x *MemberPayLevelData) GetEntity() *MemberPayLevel {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *MemberPayLevelData) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *MemberPayLevelData) GetItems() []*MemberPayLevel {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *MemberPayLevelData) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type MemberPayLevelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *MemberPayLevelData `protobuf:"bytes,1,opt,name=data,proto3" json:"data"`
	Error *common.Error       `protobuf:"bytes,2,opt,name=error,proto3" json:"error"`
}

func (x *MemberPayLevelResponse) Reset() {
	*x = MemberPayLevelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memberPayLevelService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberPayLevelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberPayLevelResponse) ProtoMessage() {}

func (x *MemberPayLevelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_memberPayLevelService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberPayLevelResponse.ProtoReflect.Descriptor instead.
func (*MemberPayLevelResponse) Descriptor() ([]byte, []int) {
	return file_memberPayLevelService_proto_rawDescGZIP(), []int{3}
}

func (x *MemberPayLevelResponse) GetData() *MemberPayLevelData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *MemberPayLevelResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type MemberPayLevelPosterData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserShareQrcode string `protobuf:"bytes,1,opt,name=user_share_qrcode,json=userShareQrcode,proto3" json:"user_share_qrcode"`
	UserShareUrl    string `protobuf:"bytes,2,opt,name=user_share_url,json=userShareUrl,proto3" json:"user_share_url"`
}

func (x *MemberPayLevelPosterData) Reset() {
	*x = MemberPayLevelPosterData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memberPayLevelService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberPayLevelPosterData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberPayLevelPosterData) ProtoMessage() {}

func (x *MemberPayLevelPosterData) ProtoReflect() protoreflect.Message {
	mi := &file_memberPayLevelService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberPayLevelPosterData.ProtoReflect.Descriptor instead.
func (*MemberPayLevelPosterData) Descriptor() ([]byte, []int) {
	return file_memberPayLevelService_proto_rawDescGZIP(), []int{4}
}

func (x *MemberPayLevelPosterData) GetUserShareQrcode() string {
	if x != nil {
		return x.UserShareQrcode
	}
	return ""
}

func (x *MemberPayLevelPosterData) GetUserShareUrl() string {
	if x != nil {
		return x.UserShareUrl
	}
	return ""
}

type MemberPayLevelPosterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *MemberPayLevelPosterData `protobuf:"bytes,1,opt,name=data,proto3" json:"data"`
	Error *common.Error             `protobuf:"bytes,2,opt,name=error,proto3" json:"error"`
}

func (x *MemberPayLevelPosterResponse) Reset() {
	*x = MemberPayLevelPosterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memberPayLevelService_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberPayLevelPosterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberPayLevelPosterResponse) ProtoMessage() {}

func (x *MemberPayLevelPosterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_memberPayLevelService_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberPayLevelPosterResponse.ProtoReflect.Descriptor instead.
func (*MemberPayLevelPosterResponse) Descriptor() ([]byte, []int) {
	return file_memberPayLevelService_proto_rawDescGZIP(), []int{5}
}

func (x *MemberPayLevelPosterResponse) GetData() *MemberPayLevelPosterData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *MemberPayLevelPosterResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_memberPayLevelService_proto protoreflect.FileDescriptor

var file_memberPayLevelService_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x61, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfb,
	0x03, 0x0a, 0x0e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x61, 0x79, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19,
	0x0a, 0x08, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x65, 0x67,
	0x69, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62,
	0x65, 0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x74, 0x72, 0x61,
	0x64, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x54, 0x72, 0x61, 0x64, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x2a, 0x0a, 0x11, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x6d, 0x6f, 0x6e, 0x65, 0x79,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x72, 0x61,
	0x64, 0x65, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x6e, 0x65, 0x77,
	0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x72, 0x65, 0x6e, 0x65,
	0x77, 0x4e, 0x75, 0x6d, 0x12, 0x2e, 0x0a, 0x13, 0x73, 0x75, 0x72, 0x70, 0x6c, 0x75, 0x73, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x11, 0x73, 0x75, 0x72, 0x70, 0x6c, 0x75, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x75, 0x72, 0x70, 0x6c, 0x75, 0x73, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x73, 0x75, 0x72, 0x70, 0x6c, 0x75, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x55, 0x6e, 0x69, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x66, 0x6f, 0x72, 0x65, 0x76, 0x65, 0x72, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x73, 0x46, 0x6f, 0x72, 0x65, 0x76, 0x65, 0x72, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x25, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0xaa, 0x01, 0x0a,
	0x15, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x61, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xbd, 0x01, 0x0a, 0x12, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x50, 0x61, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x30, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x50, 0x61, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72,
	0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x61, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x6f, 0x0a, 0x16, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x50, 0x61, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x50, 0x61, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x6c, 0x0a, 0x18, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x50, 0x61, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x50, 0x6f, 0x73, 0x74,
	0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2a, 0x0a, 0x11, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x5f, 0x71, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x53, 0x68, 0x61, 0x72, 0x65, 0x51, 0x72, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x75, 0x73, 0x65, 0x72,
	0x53, 0x68, 0x61, 0x72, 0x65, 0x55, 0x72, 0x6c, 0x22, 0x7b, 0x0a, 0x1c, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x50, 0x61, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x61, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xc6, 0x02, 0x0a, 0x15, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x50, 0x61, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x4c, 0x0a, 0x06, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x61, 0x79, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x1a, 0x26, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x61, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x50, 0x6f, 0x73,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a,
	0x03, 0x47, 0x65, 0x74, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x61, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x1a, 0x20,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x50, 0x61, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x4d, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1f, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x61,
	0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50,
	0x61, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x4b, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x61, 0x79, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x61, 0x79, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d,
	0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_memberPayLevelService_proto_rawDescOnce sync.Once
	file_memberPayLevelService_proto_rawDescData = file_memberPayLevelService_proto_rawDesc
)

func file_memberPayLevelService_proto_rawDescGZIP() []byte {
	file_memberPayLevelService_proto_rawDescOnce.Do(func() {
		file_memberPayLevelService_proto_rawDescData = protoimpl.X.CompressGZIP(file_memberPayLevelService_proto_rawDescData)
	})
	return file_memberPayLevelService_proto_rawDescData
}

var file_memberPayLevelService_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_memberPayLevelService_proto_goTypes = []interface{}{
	(*MemberPayLevel)(nil),               // 0: services.MemberPayLevel
	(*MemberPayLevelRequest)(nil),        // 1: services.MemberPayLevelRequest
	(*MemberPayLevelData)(nil),           // 2: services.MemberPayLevelData
	(*MemberPayLevelResponse)(nil),       // 3: services.MemberPayLevelResponse
	(*MemberPayLevelPosterData)(nil),     // 4: services.MemberPayLevelPosterData
	(*MemberPayLevelPosterResponse)(nil), // 5: services.MemberPayLevelPosterResponse
	(*Level)(nil),                        // 6: services.Level
	(*common.Pager)(nil),                 // 7: common.Pager
	(*common.Info)(nil),                  // 8: common.Info
	(*common.Error)(nil),                 // 9: common.Error
}
var file_memberPayLevelService_proto_depIdxs = []int32{
	6,  // 0: services.MemberPayLevel.level:type_name -> services.Level
	0,  // 1: services.MemberPayLevelData.entity:type_name -> services.MemberPayLevel
	7,  // 2: services.MemberPayLevelData.pager:type_name -> common.Pager
	0,  // 3: services.MemberPayLevelData.items:type_name -> services.MemberPayLevel
	8,  // 4: services.MemberPayLevelData.info:type_name -> common.Info
	2,  // 5: services.MemberPayLevelResponse.data:type_name -> services.MemberPayLevelData
	9,  // 6: services.MemberPayLevelResponse.error:type_name -> common.Error
	4,  // 7: services.MemberPayLevelPosterResponse.data:type_name -> services.MemberPayLevelPosterData
	9,  // 8: services.MemberPayLevelPosterResponse.error:type_name -> common.Error
	0,  // 9: services.MemberPayLevelService.Poster:input_type -> services.MemberPayLevel
	0,  // 10: services.MemberPayLevelService.Get:input_type -> services.MemberPayLevel
	1,  // 11: services.MemberPayLevelService.Search:input_type -> services.MemberPayLevelRequest
	1,  // 12: services.MemberPayLevelService.List:input_type -> services.MemberPayLevelRequest
	5,  // 13: services.MemberPayLevelService.Poster:output_type -> services.MemberPayLevelPosterResponse
	3,  // 14: services.MemberPayLevelService.Get:output_type -> services.MemberPayLevelResponse
	3,  // 15: services.MemberPayLevelService.Search:output_type -> services.MemberPayLevelResponse
	3,  // 16: services.MemberPayLevelService.List:output_type -> services.MemberPayLevelResponse
	13, // [13:17] is the sub-list for method output_type
	9,  // [9:13] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_memberPayLevelService_proto_init() }
func file_memberPayLevelService_proto_init() {
	if File_memberPayLevelService_proto != nil {
		return
	}
	file_levelService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_memberPayLevelService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberPayLevel); i {
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
		file_memberPayLevelService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberPayLevelRequest); i {
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
		file_memberPayLevelService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberPayLevelData); i {
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
		file_memberPayLevelService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberPayLevelResponse); i {
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
		file_memberPayLevelService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberPayLevelPosterData); i {
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
		file_memberPayLevelService_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberPayLevelPosterResponse); i {
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
			RawDescriptor: file_memberPayLevelService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_memberPayLevelService_proto_goTypes,
		DependencyIndexes: file_memberPayLevelService_proto_depIdxs,
		MessageInfos:      file_memberPayLevelService_proto_msgTypes,
	}.Build()
	File_memberPayLevelService_proto = out.File
	file_memberPayLevelService_proto_rawDesc = nil
	file_memberPayLevelService_proto_goTypes = nil
	file_memberPayLevelService_proto_depIdxs = nil
}