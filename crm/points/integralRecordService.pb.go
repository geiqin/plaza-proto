// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: integralRecordService.proto

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

type IntegralRecordWhere struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int32  `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize int32  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Sorting  string `protobuf:"bytes,3,opt,name=sorting,proto3" json:"sorting"`
	Keywords string `protobuf:"bytes,4,opt,name=keywords,proto3" json:"keywords"`
	Id       int64  `protobuf:"varint,5,opt,name=id,proto3" json:"id"`
	// @inject_tag: gorm:"-"
	Ids        []int64 `protobuf:"varint,6,rep,packed,name=ids,proto3" json:"ids" gorm:"-"`
	CustomerId int64   `protobuf:"varint,7,opt,name=customer_id,json=customerId,proto3" json:"customer_id"`
	Status     int32   `protobuf:"varint,8,opt,name=status,proto3" json:"status"`
	Points     int32   `protobuf:"varint,9,opt,name=points,proto3" json:"points"`
	// @inject_tag: gorm:"-"
	CustomerIds []int64 `protobuf:"varint,10,rep,packed,name=customer_ids,json=customerIds,proto3" json:"customer_ids" gorm:"-"`
	Direction   string  `protobuf:"bytes,11,opt,name=direction,proto3" json:"direction"`
	OrderId     int64   `protobuf:"varint,12,opt,name=order_id,json=orderId,proto3" json:"order_id"`
	Type        int32   `protobuf:"varint,13,opt,name=type,proto3" json:"type"`
	// @inject_tag: gorm:"-"
	Types []int32 `protobuf:"varint,14,rep,packed,name=types,proto3" json:"types" gorm:"-"`
	// @inject_tag: gorm:"-"
	Statuses []int32 `protobuf:"varint,15,rep,packed,name=statuses,proto3" json:"statuses" gorm:"-"`
}

func (x *IntegralRecordWhere) Reset() {
	*x = IntegralRecordWhere{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integralRecordService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IntegralRecordWhere) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntegralRecordWhere) ProtoMessage() {}

func (x *IntegralRecordWhere) ProtoReflect() protoreflect.Message {
	mi := &file_integralRecordService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntegralRecordWhere.ProtoReflect.Descriptor instead.
func (*IntegralRecordWhere) Descriptor() ([]byte, []int) {
	return file_integralRecordService_proto_rawDescGZIP(), []int{0}
}

func (x *IntegralRecordWhere) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *IntegralRecordWhere) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *IntegralRecordWhere) GetSorting() string {
	if x != nil {
		return x.Sorting
	}
	return ""
}

func (x *IntegralRecordWhere) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *IntegralRecordWhere) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *IntegralRecordWhere) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *IntegralRecordWhere) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *IntegralRecordWhere) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *IntegralRecordWhere) GetPoints() int32 {
	if x != nil {
		return x.Points
	}
	return 0
}

func (x *IntegralRecordWhere) GetCustomerIds() []int64 {
	if x != nil {
		return x.CustomerIds
	}
	return nil
}

func (x *IntegralRecordWhere) GetDirection() string {
	if x != nil {
		return x.Direction
	}
	return ""
}

func (x *IntegralRecordWhere) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *IntegralRecordWhere) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *IntegralRecordWhere) GetTypes() []int32 {
	if x != nil {
		return x.Types
	}
	return nil
}

func (x *IntegralRecordWhere) GetStatuses() []int32 {
	if x != nil {
		return x.Statuses
	}
	return nil
}

type IntegralRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	CustomerId   int64  `protobuf:"varint,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id"`
	IntegralId   int64  `protobuf:"varint,3,opt,name=integral_id,json=integralId,proto3" json:"integral_id"`
	Type         int32  `protobuf:"varint,4,opt,name=type,proto3" json:"type"`
	Direction    string `protobuf:"bytes,5,opt,name=direction,proto3" json:"direction"`
	Value        int32  `protobuf:"varint,6,opt,name=value,proto3" json:"value"`
	Balance      int32  `protobuf:"varint,7,opt,name=balance,proto3" json:"balance"`
	UsedValue    int32  `protobuf:"varint,8,opt,name=used_value,json=usedValue,proto3" json:"used_value"`
	EffectiveAt  string `protobuf:"bytes,9,opt,name=effective_at,json=effectiveAt,proto3" json:"effective_at"`
	ExpirationAt string `protobuf:"bytes,10,opt,name=expiration_at,json=expirationAt,proto3" json:"expiration_at"`
	Memo         string `protobuf:"bytes,11,opt,name=memo,proto3" json:"memo"`
	Status       int32  `protobuf:"varint,12,opt,name=status,proto3" json:"status"`
	CreatedAt    string `protobuf:"bytes,13,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt    string `protobuf:"bytes,14,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	OrderId      int64  `protobuf:"varint,15,opt,name=order_id,json=orderId,proto3" json:"order_id"`
}

func (x *IntegralRecord) Reset() {
	*x = IntegralRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integralRecordService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IntegralRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntegralRecord) ProtoMessage() {}

func (x *IntegralRecord) ProtoReflect() protoreflect.Message {
	mi := &file_integralRecordService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntegralRecord.ProtoReflect.Descriptor instead.
func (*IntegralRecord) Descriptor() ([]byte, []int) {
	return file_integralRecordService_proto_rawDescGZIP(), []int{1}
}

func (x *IntegralRecord) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *IntegralRecord) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *IntegralRecord) GetIntegralId() int64 {
	if x != nil {
		return x.IntegralId
	}
	return 0
}

func (x *IntegralRecord) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *IntegralRecord) GetDirection() string {
	if x != nil {
		return x.Direction
	}
	return ""
}

func (x *IntegralRecord) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *IntegralRecord) GetBalance() int32 {
	if x != nil {
		return x.Balance
	}
	return 0
}

func (x *IntegralRecord) GetUsedValue() int32 {
	if x != nil {
		return x.UsedValue
	}
	return 0
}

func (x *IntegralRecord) GetEffectiveAt() string {
	if x != nil {
		return x.EffectiveAt
	}
	return ""
}

func (x *IntegralRecord) GetExpirationAt() string {
	if x != nil {
		return x.ExpirationAt
	}
	return ""
}

func (x *IntegralRecord) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *IntegralRecord) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *IntegralRecord) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *IntegralRecord) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *IntegralRecord) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

type IntegralRecordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *IntegralRecord   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager     `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*IntegralRecord `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Error  *common.Error     `protobuf:"bytes,4,opt,name=error,proto3" json:"error"`
	Info   *common.Info      `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
}

func (x *IntegralRecordResponse) Reset() {
	*x = IntegralRecordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integralRecordService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IntegralRecordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntegralRecordResponse) ProtoMessage() {}

func (x *IntegralRecordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_integralRecordService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntegralRecordResponse.ProtoReflect.Descriptor instead.
func (*IntegralRecordResponse) Descriptor() ([]byte, []int) {
	return file_integralRecordService_proto_rawDescGZIP(), []int{2}
}

func (x *IntegralRecordResponse) GetEntity() *IntegralRecord {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *IntegralRecordResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *IntegralRecordResponse) GetItems() []*IntegralRecord {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *IntegralRecordResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *IntegralRecordResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_integralRecordService_proto protoreflect.FileDescriptor

var file_integralRecordService_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93, 0x03, 0x0a, 0x13, 0x49,
	0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x57, 0x68, 0x65,
	0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x12,
	0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x69,
	0x64, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x1f, 0x0a,
	0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x21,
	0x0a, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x0a,
	0x20, 0x03, 0x28, 0x03, 0x52, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64,
	0x73, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x05, 0x52, 0x05, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73,
	0x18, 0x0f, 0x20, 0x03, 0x28, 0x05, 0x52, 0x08, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73,
	0x22, 0xb0, 0x03, 0x0a, 0x0e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x69, 0x6e, 0x74, 0x65, 0x67,
	0x72, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x64, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x75, 0x73, 0x65,
	0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x66,
	0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x41, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65,
	0x6d, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x64, 0x22, 0xe6, 0x01, 0x0a, 0x16, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30,
	0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72,
	0x61, 0x6c, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05,
	0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0x64, 0x0a, 0x15,
	0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12,
	0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67,
	0x72, 0x61, 0x6c, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x20,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72,
	0x61, 0x6c, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x32, 0x66, 0x0a, 0x17, 0x4d, 0x79, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a,
	0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_integralRecordService_proto_rawDescOnce sync.Once
	file_integralRecordService_proto_rawDescData = file_integralRecordService_proto_rawDesc
)

func file_integralRecordService_proto_rawDescGZIP() []byte {
	file_integralRecordService_proto_rawDescOnce.Do(func() {
		file_integralRecordService_proto_rawDescData = protoimpl.X.CompressGZIP(file_integralRecordService_proto_rawDescData)
	})
	return file_integralRecordService_proto_rawDescData
}

var file_integralRecordService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_integralRecordService_proto_goTypes = []interface{}{
	(*IntegralRecordWhere)(nil),    // 0: services.IntegralRecordWhere
	(*IntegralRecord)(nil),         // 1: services.IntegralRecord
	(*IntegralRecordResponse)(nil), // 2: services.IntegralRecordResponse
	(*common.Pager)(nil),           // 3: common.Pager
	(*common.Error)(nil),           // 4: common.Error
	(*common.Info)(nil),            // 5: common.Info
}
var file_integralRecordService_proto_depIdxs = []int32{
	1, // 0: services.IntegralRecordResponse.entity:type_name -> services.IntegralRecord
	3, // 1: services.IntegralRecordResponse.pager:type_name -> common.Pager
	1, // 2: services.IntegralRecordResponse.items:type_name -> services.IntegralRecord
	4, // 3: services.IntegralRecordResponse.error:type_name -> common.Error
	5, // 4: services.IntegralRecordResponse.info:type_name -> common.Info
	0, // 5: services.IntegralRecordService.Search:input_type -> services.IntegralRecordWhere
	0, // 6: services.MyIntegralRecordService.Search:input_type -> services.IntegralRecordWhere
	2, // 7: services.IntegralRecordService.Search:output_type -> services.IntegralRecordResponse
	2, // 8: services.MyIntegralRecordService.Search:output_type -> services.IntegralRecordResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_integralRecordService_proto_init() }
func file_integralRecordService_proto_init() {
	if File_integralRecordService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_integralRecordService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IntegralRecordWhere); i {
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
		file_integralRecordService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IntegralRecord); i {
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
		file_integralRecordService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IntegralRecordResponse); i {
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
			RawDescriptor: file_integralRecordService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_integralRecordService_proto_goTypes,
		DependencyIndexes: file_integralRecordService_proto_depIdxs,
		MessageInfos:      file_integralRecordService_proto_msgTypes,
	}.Build()
	File_integralRecordService_proto = out.File
	file_integralRecordService_proto_rawDesc = nil
	file_integralRecordService_proto_goTypes = nil
	file_integralRecordService_proto_depIdxs = nil
}
