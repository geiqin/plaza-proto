// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: surplusRecordService.proto

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

type SurplusRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	MemberId  int64   `protobuf:"varint,2,opt,name=member_id,json=memberId,proto3" json:"member_id"`
	Direction string  `protobuf:"bytes,3,opt,name=direction,proto3" json:"direction"`
	Amount    int64   `protobuf:"varint,4,opt,name=amount,proto3" json:"amount"`
	Balance   int64   `protobuf:"varint,5,opt,name=balance,proto3" json:"balance"`
	Source    int32   `protobuf:"varint,6,opt,name=source,proto3" json:"source"`
	OrderId   int64   `protobuf:"varint,7,opt,name=order_id,json=orderId,proto3" json:"order_id"`
	OrderSn   string  `protobuf:"bytes,8,opt,name=order_sn,json=orderSn,proto3" json:"order_sn"`
	Memo      string  `protobuf:"bytes,9,opt,name=memo,proto3" json:"memo"`
	CreatedAt string  `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt string  `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	Member    *Member `protobuf:"bytes,12,opt,name=member,proto3" json:"member"`
}

func (x *SurplusRecord) Reset() {
	*x = SurplusRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_surplusRecordService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SurplusRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SurplusRecord) ProtoMessage() {}

func (x *SurplusRecord) ProtoReflect() protoreflect.Message {
	mi := &file_surplusRecordService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SurplusRecord.ProtoReflect.Descriptor instead.
func (*SurplusRecord) Descriptor() ([]byte, []int) {
	return file_surplusRecordService_proto_rawDescGZIP(), []int{0}
}

func (x *SurplusRecord) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SurplusRecord) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

func (x *SurplusRecord) GetDirection() string {
	if x != nil {
		return x.Direction
	}
	return ""
}

func (x *SurplusRecord) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *SurplusRecord) GetBalance() int64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

func (x *SurplusRecord) GetSource() int32 {
	if x != nil {
		return x.Source
	}
	return 0
}

func (x *SurplusRecord) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *SurplusRecord) GetOrderSn() string {
	if x != nil {
		return x.OrderSn
	}
	return ""
}

func (x *SurplusRecord) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *SurplusRecord) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *SurplusRecord) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *SurplusRecord) GetMember() *Member {
	if x != nil {
		return x.Member
	}
	return nil
}

//查询参数
type SurplusRecordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int32 `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	//以下为自定义参数
	MemberId  int64  `protobuf:"varint,3,opt,name=member_id,json=memberId,proto3" json:"member_id"`
	Direction int32  `protobuf:"varint,4,opt,name=direction,proto3" json:"direction"`
	Amount    int64  `protobuf:"varint,5,opt,name=amount,proto3" json:"amount"`
	Balance   int64  `protobuf:"varint,6,opt,name=balance,proto3" json:"balance"`
	Source    int32  `protobuf:"varint,7,opt,name=source,proto3" json:"source"`
	OrderId   int64  `protobuf:"varint,8,opt,name=order_id,json=orderId,proto3" json:"order_id"`
	OrderSn   string `protobuf:"bytes,9,opt,name=order_sn,json=orderSn,proto3" json:"order_sn"`
	Memo      string `protobuf:"bytes,10,opt,name=memo,proto3" json:"memo"`
	StartDate string `protobuf:"bytes,11,opt,name=start_date,json=startDate,proto3" json:"start_date"`
	EndDate   string `protobuf:"bytes,12,opt,name=end_date,json=endDate,proto3" json:"end_date"`
}

func (x *SurplusRecordRequest) Reset() {
	*x = SurplusRecordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_surplusRecordService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SurplusRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SurplusRecordRequest) ProtoMessage() {}

func (x *SurplusRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_surplusRecordService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SurplusRecordRequest.ProtoReflect.Descriptor instead.
func (*SurplusRecordRequest) Descriptor() ([]byte, []int) {
	return file_surplusRecordService_proto_rawDescGZIP(), []int{1}
}

func (x *SurplusRecordRequest) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *SurplusRecordRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *SurplusRecordRequest) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

func (x *SurplusRecordRequest) GetDirection() int32 {
	if x != nil {
		return x.Direction
	}
	return 0
}

func (x *SurplusRecordRequest) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *SurplusRecordRequest) GetBalance() int64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

func (x *SurplusRecordRequest) GetSource() int32 {
	if x != nil {
		return x.Source
	}
	return 0
}

func (x *SurplusRecordRequest) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *SurplusRecordRequest) GetOrderSn() string {
	if x != nil {
		return x.OrderSn
	}
	return ""
}

func (x *SurplusRecordRequest) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *SurplusRecordRequest) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *SurplusRecordRequest) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

type SurplusRecordData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *SurplusRecord   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager    `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*SurplusRecord `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Info   *common.Info     `protobuf:"bytes,4,opt,name=info,proto3" json:"info"`
}

func (x *SurplusRecordData) Reset() {
	*x = SurplusRecordData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_surplusRecordService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SurplusRecordData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SurplusRecordData) ProtoMessage() {}

func (x *SurplusRecordData) ProtoReflect() protoreflect.Message {
	mi := &file_surplusRecordService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SurplusRecordData.ProtoReflect.Descriptor instead.
func (*SurplusRecordData) Descriptor() ([]byte, []int) {
	return file_surplusRecordService_proto_rawDescGZIP(), []int{2}
}

func (x *SurplusRecordData) GetEntity() *SurplusRecord {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *SurplusRecordData) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *SurplusRecordData) GetItems() []*SurplusRecord {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *SurplusRecordData) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type SurplusRecordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *SurplusRecordData `protobuf:"bytes,1,opt,name=data,proto3" json:"data"`
	Error *common.Error      `protobuf:"bytes,2,opt,name=error,proto3" json:"error"`
}

func (x *SurplusRecordResponse) Reset() {
	*x = SurplusRecordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_surplusRecordService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SurplusRecordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SurplusRecordResponse) ProtoMessage() {}

func (x *SurplusRecordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_surplusRecordService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SurplusRecordResponse.ProtoReflect.Descriptor instead.
func (*SurplusRecordResponse) Descriptor() ([]byte, []int) {
	return file_surplusRecordService_proto_rawDescGZIP(), []int{3}
}

func (x *SurplusRecordResponse) GetData() *SurplusRecordData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *SurplusRecordResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_surplusRecordService_proto protoreflect.FileDescriptor

var file_surplusRecordService_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x73, 0x75, 0x72, 0x70, 0x6c, 0x75, 0x73, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd6,
	0x02, 0x0a, 0x0d, 0x53, 0x75, 0x72, 0x70, 0x6c, 0x75, 0x73, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x6e, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6d,
	0x65, 0x6d, 0x6f, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x28, 0x0a,
	0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52,
	0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x22, 0xd2, 0x02, 0x0a, 0x14, 0x53, 0x75, 0x72, 0x70,
	0x6c, 0x75, 0x73, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x6e, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65,
	0x6d, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x22, 0xba, 0x01, 0x0a,
	0x11, 0x53, 0x75, 0x72, 0x70, 0x6c, 0x75, 0x73, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x2f, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x75,
	0x72, 0x70, 0x6c, 0x75, 0x73, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x06, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65,
	0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x53, 0x75, 0x72, 0x70, 0x6c, 0x75, 0x73, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x6d, 0x0a, 0x15, 0x53, 0x75, 0x72,
	0x70, 0x6c, 0x75, 0x73, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x75, 0x72, 0x70,
	0x6c, 0x75, 0x73, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xb2, 0x02, 0x0a, 0x14, 0x53, 0x75, 0x72,
	0x70, 0x6c, 0x75, 0x73, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x44, 0x0a, 0x06, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x12, 0x17, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x75, 0x72, 0x70, 0x6c, 0x75, 0x73, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x1a, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x53, 0x75, 0x72, 0x70, 0x6c, 0x75, 0x73, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x06, 0x45, 0x78, 0x70, 0x65, 0x6e,
	0x64, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x75, 0x72,
	0x70, 0x6c, 0x75, 0x73, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x1a, 0x1f, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x75, 0x72, 0x70, 0x6c, 0x75, 0x73, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a,
	0x03, 0x47, 0x65, 0x74, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x53, 0x75, 0x72, 0x70, 0x6c, 0x75, 0x73, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x1a, 0x1f, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x75, 0x72, 0x70, 0x6c, 0x75, 0x73,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x4b, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1e, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x75, 0x72, 0x70, 0x6c, 0x75, 0x73, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x75, 0x72, 0x70, 0x6c, 0x75, 0x73, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a,
	0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_surplusRecordService_proto_rawDescOnce sync.Once
	file_surplusRecordService_proto_rawDescData = file_surplusRecordService_proto_rawDesc
)

func file_surplusRecordService_proto_rawDescGZIP() []byte {
	file_surplusRecordService_proto_rawDescOnce.Do(func() {
		file_surplusRecordService_proto_rawDescData = protoimpl.X.CompressGZIP(file_surplusRecordService_proto_rawDescData)
	})
	return file_surplusRecordService_proto_rawDescData
}

var file_surplusRecordService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_surplusRecordService_proto_goTypes = []interface{}{
	(*SurplusRecord)(nil),         // 0: services.SurplusRecord
	(*SurplusRecordRequest)(nil),  // 1: services.SurplusRecordRequest
	(*SurplusRecordData)(nil),     // 2: services.SurplusRecordData
	(*SurplusRecordResponse)(nil), // 3: services.SurplusRecordResponse
	(*Member)(nil),                // 4: services.Member
	(*common.Pager)(nil),          // 5: common.Pager
	(*common.Info)(nil),           // 6: common.Info
	(*common.Error)(nil),          // 7: common.Error
}
var file_surplusRecordService_proto_depIdxs = []int32{
	4,  // 0: services.SurplusRecord.member:type_name -> services.Member
	0,  // 1: services.SurplusRecordData.entity:type_name -> services.SurplusRecord
	5,  // 2: services.SurplusRecordData.pager:type_name -> common.Pager
	0,  // 3: services.SurplusRecordData.items:type_name -> services.SurplusRecord
	6,  // 4: services.SurplusRecordData.info:type_name -> common.Info
	2,  // 5: services.SurplusRecordResponse.data:type_name -> services.SurplusRecordData
	7,  // 6: services.SurplusRecordResponse.error:type_name -> common.Error
	0,  // 7: services.SurplusRecordService.Income:input_type -> services.SurplusRecord
	0,  // 8: services.SurplusRecordService.Expend:input_type -> services.SurplusRecord
	0,  // 9: services.SurplusRecordService.Get:input_type -> services.SurplusRecord
	1,  // 10: services.SurplusRecordService.Search:input_type -> services.SurplusRecordRequest
	3,  // 11: services.SurplusRecordService.Income:output_type -> services.SurplusRecordResponse
	3,  // 12: services.SurplusRecordService.Expend:output_type -> services.SurplusRecordResponse
	3,  // 13: services.SurplusRecordService.Get:output_type -> services.SurplusRecordResponse
	3,  // 14: services.SurplusRecordService.Search:output_type -> services.SurplusRecordResponse
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_surplusRecordService_proto_init() }
func file_surplusRecordService_proto_init() {
	if File_surplusRecordService_proto != nil {
		return
	}
	file_memberService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_surplusRecordService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SurplusRecord); i {
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
		file_surplusRecordService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SurplusRecordRequest); i {
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
		file_surplusRecordService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SurplusRecordData); i {
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
		file_surplusRecordService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SurplusRecordResponse); i {
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
			RawDescriptor: file_surplusRecordService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_surplusRecordService_proto_goTypes,
		DependencyIndexes: file_surplusRecordService_proto_depIdxs,
		MessageInfos:      file_surplusRecordService_proto_msgTypes,
	}.Build()
	File_surplusRecordService_proto = out.File
	file_surplusRecordService_proto_rawDesc = nil
	file_surplusRecordService_proto_goTypes = nil
	file_surplusRecordService_proto_depIdxs = nil
}
