// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: rankRecordService.proto

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

type RankRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	MemberId  int64   `protobuf:"varint,2,opt,name=member_id,json=memberId,proto3" json:"member_id"`
	RankId    int32   `protobuf:"varint,3,opt,name=rank_id,json=rankId,proto3" json:"rank_id"`
	OrderId   int64   `protobuf:"varint,4,opt,name=order_id,json=orderId,proto3" json:"order_id"`
	OrderSn   string  `protobuf:"bytes,5,opt,name=order_sn,json=orderSn,proto3" json:"order_sn"`
	KeepType  string  `protobuf:"bytes,6,opt,name=keep_type,json=keepType,proto3" json:"keep_type"`
	KeepValue int32   `protobuf:"varint,7,opt,name=keep_value,json=keepValue,proto3" json:"keep_value"`
	CreatedAt string  `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt string  `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	Member    *Member `protobuf:"bytes,10,opt,name=member,proto3" json:"member"`
}

func (x *RankRecord) Reset() {
	*x = RankRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rankRecordService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RankRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RankRecord) ProtoMessage() {}

func (x *RankRecord) ProtoReflect() protoreflect.Message {
	mi := &file_rankRecordService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RankRecord.ProtoReflect.Descriptor instead.
func (*RankRecord) Descriptor() ([]byte, []int) {
	return file_rankRecordService_proto_rawDescGZIP(), []int{0}
}

func (x *RankRecord) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RankRecord) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

func (x *RankRecord) GetRankId() int32 {
	if x != nil {
		return x.RankId
	}
	return 0
}

func (x *RankRecord) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *RankRecord) GetOrderSn() string {
	if x != nil {
		return x.OrderSn
	}
	return ""
}

func (x *RankRecord) GetKeepType() string {
	if x != nil {
		return x.KeepType
	}
	return ""
}

func (x *RankRecord) GetKeepValue() int32 {
	if x != nil {
		return x.KeepValue
	}
	return 0
}

func (x *RankRecord) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *RankRecord) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *RankRecord) GetMember() *Member {
	if x != nil {
		return x.Member
	}
	return nil
}

type RankRecordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int32   `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize int32   `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Sorting  string  `protobuf:"bytes,3,opt,name=sorting,proto3" json:"sorting"`
	MemberId int64   `protobuf:"varint,4,opt,name=member_id,json=memberId,proto3" json:"member_id"`
	RankId   int32   `protobuf:"varint,5,opt,name=rank_id,json=rankId,proto3" json:"rank_id"`
	OrderId  int64   `protobuf:"varint,6,opt,name=order_id,json=orderId,proto3" json:"order_id"`
	OrderSn  string  `protobuf:"bytes,7,opt,name=order_sn,json=orderSn,proto3" json:"order_sn"`
	Id       int64   `protobuf:"varint,8,opt,name=id,proto3" json:"id"`
	Ids      []int64 `protobuf:"varint,9,rep,packed,name=ids,proto3" json:"ids"`
}

func (x *RankRecordRequest) Reset() {
	*x = RankRecordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rankRecordService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RankRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RankRecordRequest) ProtoMessage() {}

func (x *RankRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rankRecordService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RankRecordRequest.ProtoReflect.Descriptor instead.
func (*RankRecordRequest) Descriptor() ([]byte, []int) {
	return file_rankRecordService_proto_rawDescGZIP(), []int{1}
}

func (x *RankRecordRequest) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *RankRecordRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *RankRecordRequest) GetSorting() string {
	if x != nil {
		return x.Sorting
	}
	return ""
}

func (x *RankRecordRequest) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

func (x *RankRecordRequest) GetRankId() int32 {
	if x != nil {
		return x.RankId
	}
	return 0
}

func (x *RankRecordRequest) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *RankRecordRequest) GetOrderSn() string {
	if x != nil {
		return x.OrderSn
	}
	return ""
}

func (x *RankRecordRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RankRecordRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type RankRecordData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *RankRecord   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*RankRecord `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Info   *common.Info  `protobuf:"bytes,4,opt,name=info,proto3" json:"info"`
}

func (x *RankRecordData) Reset() {
	*x = RankRecordData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rankRecordService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RankRecordData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RankRecordData) ProtoMessage() {}

func (x *RankRecordData) ProtoReflect() protoreflect.Message {
	mi := &file_rankRecordService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RankRecordData.ProtoReflect.Descriptor instead.
func (*RankRecordData) Descriptor() ([]byte, []int) {
	return file_rankRecordService_proto_rawDescGZIP(), []int{2}
}

func (x *RankRecordData) GetEntity() *RankRecord {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *RankRecordData) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *RankRecordData) GetItems() []*RankRecord {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *RankRecordData) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type RankRecordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *RankRecordData `protobuf:"bytes,1,opt,name=data,proto3" json:"data"`
	Error *common.Error   `protobuf:"bytes,2,opt,name=error,proto3" json:"error"`
}

func (x *RankRecordResponse) Reset() {
	*x = RankRecordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rankRecordService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RankRecordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RankRecordResponse) ProtoMessage() {}

func (x *RankRecordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rankRecordService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RankRecordResponse.ProtoReflect.Descriptor instead.
func (*RankRecordResponse) Descriptor() ([]byte, []int) {
	return file_rankRecordService_proto_rawDescGZIP(), []int{3}
}

func (x *RankRecordResponse) GetData() *RankRecordData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *RankRecordResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_rankRecordService_proto protoreflect.FileDescriptor

var file_rankRecordService_proto_rawDesc = []byte{
	0x0a, 0x17, 0x72, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xac, 0x02, 0x0a, 0x0a,
	0x52, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x61, 0x6e, 0x6b, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x61, 0x6e, 0x6b, 0x49, 0x64,
	0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x53, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x6b, 0x65, 0x65, 0x70, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x65, 0x70, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6b, 0x65, 0x65, 0x70, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6b, 0x65, 0x65, 0x70, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x28, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x22, 0xee, 0x01, 0x0a, 0x11, 0x52,
	0x61, 0x6e, 0x6b, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1b, 0x0a,
	0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x61,
	0x6e, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x61, 0x6e,
	0x6b, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19,
	0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73,
	0x18, 0x09, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0xb1, 0x01, 0x0a, 0x0e,
	0x52, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2c,
	0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05,
	0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65,
	0x72, 0x12, 0x2a, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x61, 0x6e, 0x6b,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x20, 0x0a,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22,
	0x67, 0x0a, 0x12, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52,
	0x61, 0x6e, 0x6b, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x5a, 0x0a, 0x11, 0x52, 0x61, 0x6e, 0x6b,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a,
	0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x52, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rankRecordService_proto_rawDescOnce sync.Once
	file_rankRecordService_proto_rawDescData = file_rankRecordService_proto_rawDesc
)

func file_rankRecordService_proto_rawDescGZIP() []byte {
	file_rankRecordService_proto_rawDescOnce.Do(func() {
		file_rankRecordService_proto_rawDescData = protoimpl.X.CompressGZIP(file_rankRecordService_proto_rawDescData)
	})
	return file_rankRecordService_proto_rawDescData
}

var file_rankRecordService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_rankRecordService_proto_goTypes = []interface{}{
	(*RankRecord)(nil),         // 0: services.RankRecord
	(*RankRecordRequest)(nil),  // 1: services.RankRecordRequest
	(*RankRecordData)(nil),     // 2: services.RankRecordData
	(*RankRecordResponse)(nil), // 3: services.RankRecordResponse
	(*Member)(nil),             // 4: services.Member
	(*common.Pager)(nil),       // 5: common.Pager
	(*common.Info)(nil),        // 6: common.Info
	(*common.Error)(nil),       // 7: common.Error
}
var file_rankRecordService_proto_depIdxs = []int32{
	4, // 0: services.RankRecord.member:type_name -> services.Member
	0, // 1: services.RankRecordData.entity:type_name -> services.RankRecord
	5, // 2: services.RankRecordData.pager:type_name -> common.Pager
	0, // 3: services.RankRecordData.items:type_name -> services.RankRecord
	6, // 4: services.RankRecordData.info:type_name -> common.Info
	2, // 5: services.RankRecordResponse.data:type_name -> services.RankRecordData
	7, // 6: services.RankRecordResponse.error:type_name -> common.Error
	1, // 7: services.RankRecordService.Search:input_type -> services.RankRecordRequest
	3, // 8: services.RankRecordService.Search:output_type -> services.RankRecordResponse
	8, // [8:9] is the sub-list for method output_type
	7, // [7:8] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_rankRecordService_proto_init() }
func file_rankRecordService_proto_init() {
	if File_rankRecordService_proto != nil {
		return
	}
	file_memberService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rankRecordService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RankRecord); i {
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
		file_rankRecordService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RankRecordRequest); i {
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
		file_rankRecordService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RankRecordData); i {
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
		file_rankRecordService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RankRecordResponse); i {
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
			RawDescriptor: file_rankRecordService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rankRecordService_proto_goTypes,
		DependencyIndexes: file_rankRecordService_proto_depIdxs,
		MessageInfos:      file_rankRecordService_proto_msgTypes,
	}.Build()
	File_rankRecordService_proto = out.File
	file_rankRecordService_proto_rawDesc = nil
	file_rankRecordService_proto_goTypes = nil
	file_rankRecordService_proto_depIdxs = nil
}
