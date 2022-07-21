// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: depositRecordService.proto

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

type DepositRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	MemberId  int64   `protobuf:"varint,2,opt,name=member_id,json=memberId,proto3" json:"member_id"`
	Direction int32   `protobuf:"varint,3,opt,name=direction,proto3" json:"direction"`
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

func (x *DepositRecord) Reset() {
	*x = DepositRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_depositRecordService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DepositRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DepositRecord) ProtoMessage() {}

func (x *DepositRecord) ProtoReflect() protoreflect.Message {
	mi := &file_depositRecordService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DepositRecord.ProtoReflect.Descriptor instead.
func (*DepositRecord) Descriptor() ([]byte, []int) {
	return file_depositRecordService_proto_rawDescGZIP(), []int{0}
}

func (x *DepositRecord) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DepositRecord) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

func (x *DepositRecord) GetDirection() int32 {
	if x != nil {
		return x.Direction
	}
	return 0
}

func (x *DepositRecord) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *DepositRecord) GetBalance() int64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

func (x *DepositRecord) GetSource() int32 {
	if x != nil {
		return x.Source
	}
	return 0
}

func (x *DepositRecord) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *DepositRecord) GetOrderSn() string {
	if x != nil {
		return x.OrderSn
	}
	return ""
}

func (x *DepositRecord) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *DepositRecord) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *DepositRecord) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *DepositRecord) GetMember() *Member {
	if x != nil {
		return x.Member
	}
	return nil
}

//查询参数
type DepositRecordRequest struct {
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

func (x *DepositRecordRequest) Reset() {
	*x = DepositRecordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_depositRecordService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DepositRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DepositRecordRequest) ProtoMessage() {}

func (x *DepositRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_depositRecordService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DepositRecordRequest.ProtoReflect.Descriptor instead.
func (*DepositRecordRequest) Descriptor() ([]byte, []int) {
	return file_depositRecordService_proto_rawDescGZIP(), []int{1}
}

func (x *DepositRecordRequest) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *DepositRecordRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *DepositRecordRequest) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

func (x *DepositRecordRequest) GetDirection() int32 {
	if x != nil {
		return x.Direction
	}
	return 0
}

func (x *DepositRecordRequest) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *DepositRecordRequest) GetBalance() int64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

func (x *DepositRecordRequest) GetSource() int32 {
	if x != nil {
		return x.Source
	}
	return 0
}

func (x *DepositRecordRequest) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *DepositRecordRequest) GetOrderSn() string {
	if x != nil {
		return x.OrderSn
	}
	return ""
}

func (x *DepositRecordRequest) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *DepositRecordRequest) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *DepositRecordRequest) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

//
type DepositRecordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *DepositRecord   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager    `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*DepositRecord `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Error  *common.Error    `protobuf:"bytes,4,opt,name=error,proto3" json:"error"`
	Info   *common.Info     `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
}

func (x *DepositRecordResponse) Reset() {
	*x = DepositRecordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_depositRecordService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DepositRecordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DepositRecordResponse) ProtoMessage() {}

func (x *DepositRecordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_depositRecordService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DepositRecordResponse.ProtoReflect.Descriptor instead.
func (*DepositRecordResponse) Descriptor() ([]byte, []int) {
	return file_depositRecordService_proto_rawDescGZIP(), []int{2}
}

func (x *DepositRecordResponse) GetEntity() *DepositRecord {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *DepositRecordResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *DepositRecordResponse) GetItems() []*DepositRecord {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *DepositRecordResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *DepositRecordResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_depositRecordService_proto protoreflect.FileDescriptor

var file_depositRecordService_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd6,
	0x02, 0x0a, 0x0d, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
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
	0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x22, 0xd2, 0x02, 0x0a, 0x14, 0x44, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
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
	0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x22, 0xe3, 0x01, 0x0a,
	0x15, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52,
	0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x23, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x32, 0xb2, 0x02, 0x0a, 0x14, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x06, 0x49,
	0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x1a, 0x1f,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x44, 0x0a, 0x06, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x64, 0x12, 0x17, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x1a, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x17,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x1a, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x06, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x12, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_depositRecordService_proto_rawDescOnce sync.Once
	file_depositRecordService_proto_rawDescData = file_depositRecordService_proto_rawDesc
)

func file_depositRecordService_proto_rawDescGZIP() []byte {
	file_depositRecordService_proto_rawDescOnce.Do(func() {
		file_depositRecordService_proto_rawDescData = protoimpl.X.CompressGZIP(file_depositRecordService_proto_rawDescData)
	})
	return file_depositRecordService_proto_rawDescData
}

var file_depositRecordService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_depositRecordService_proto_goTypes = []interface{}{
	(*DepositRecord)(nil),         // 0: services.DepositRecord
	(*DepositRecordRequest)(nil),  // 1: services.DepositRecordRequest
	(*DepositRecordResponse)(nil), // 2: services.DepositRecordResponse
	(*Member)(nil),                // 3: services.Member
	(*common.Pager)(nil),          // 4: common.Pager
	(*common.Error)(nil),          // 5: common.Error
	(*common.Info)(nil),           // 6: common.Info
}
var file_depositRecordService_proto_depIdxs = []int32{
	3,  // 0: services.DepositRecord.member:type_name -> services.Member
	0,  // 1: services.DepositRecordResponse.entity:type_name -> services.DepositRecord
	4,  // 2: services.DepositRecordResponse.pager:type_name -> common.Pager
	0,  // 3: services.DepositRecordResponse.items:type_name -> services.DepositRecord
	5,  // 4: services.DepositRecordResponse.error:type_name -> common.Error
	6,  // 5: services.DepositRecordResponse.info:type_name -> common.Info
	0,  // 6: services.DepositRecordService.Income:input_type -> services.DepositRecord
	0,  // 7: services.DepositRecordService.Expend:input_type -> services.DepositRecord
	0,  // 8: services.DepositRecordService.Get:input_type -> services.DepositRecord
	1,  // 9: services.DepositRecordService.Search:input_type -> services.DepositRecordRequest
	2,  // 10: services.DepositRecordService.Income:output_type -> services.DepositRecordResponse
	2,  // 11: services.DepositRecordService.Expend:output_type -> services.DepositRecordResponse
	2,  // 12: services.DepositRecordService.Get:output_type -> services.DepositRecordResponse
	2,  // 13: services.DepositRecordService.Search:output_type -> services.DepositRecordResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_depositRecordService_proto_init() }
func file_depositRecordService_proto_init() {
	if File_depositRecordService_proto != nil {
		return
	}
	file_memberService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_depositRecordService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DepositRecord); i {
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
		file_depositRecordService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DepositRecordRequest); i {
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
		file_depositRecordService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DepositRecordResponse); i {
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
			RawDescriptor: file_depositRecordService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_depositRecordService_proto_goTypes,
		DependencyIndexes: file_depositRecordService_proto_depIdxs,
		MessageInfos:      file_depositRecordService_proto_msgTypes,
	}.Build()
	File_depositRecordService_proto = out.File
	file_depositRecordService_proto_rawDesc = nil
	file_depositRecordService_proto_goTypes = nil
	file_depositRecordService_proto_depIdxs = nil
}
