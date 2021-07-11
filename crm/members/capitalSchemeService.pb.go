// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: capitalSchemeService.proto

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

type CapitalScheme struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Name           string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	StartAt        string  `protobuf:"bytes,3,opt,name=start_at,json=startAt,proto3" json:"start_at"`
	EndAt          string  `protobuf:"bytes,4,opt,name=end_at,json=endAt,proto3" json:"end_at"`
	ConditionType  int32   `protobuf:"varint,5,opt,name=condition_type,json=conditionType,proto3" json:"condition_type"`
	ConditionValue float32 `protobuf:"fixed32,6,opt,name=condition_value,json=conditionValue,proto3" json:"condition_value"`
	GiveMoney      float32 `protobuf:"fixed32,7,opt,name=give_money,json=giveMoney,proto3" json:"give_money"`
	GiveScore      int32   `protobuf:"varint,8,opt,name=give_score,json=giveScore,proto3" json:"give_score"`
	Stock          int32   `protobuf:"varint,9,opt,name=stock,proto3" json:"stock"`
	Memo           string  `protobuf:"bytes,10,opt,name=memo,proto3" json:"memo"`
	CreatedAt      string  `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt      string  `protobuf:"bytes,12,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
}

func (x *CapitalScheme) Reset() {
	*x = CapitalScheme{}
	if protoimpl.UnsafeEnabled {
		mi := &file_capitalSchemeService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CapitalScheme) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CapitalScheme) ProtoMessage() {}

func (x *CapitalScheme) ProtoReflect() protoreflect.Message {
	mi := &file_capitalSchemeService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CapitalScheme.ProtoReflect.Descriptor instead.
func (*CapitalScheme) Descriptor() ([]byte, []int) {
	return file_capitalSchemeService_proto_rawDescGZIP(), []int{0}
}

func (x *CapitalScheme) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CapitalScheme) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CapitalScheme) GetStartAt() string {
	if x != nil {
		return x.StartAt
	}
	return ""
}

func (x *CapitalScheme) GetEndAt() string {
	if x != nil {
		return x.EndAt
	}
	return ""
}

func (x *CapitalScheme) GetConditionType() int32 {
	if x != nil {
		return x.ConditionType
	}
	return 0
}

func (x *CapitalScheme) GetConditionValue() float32 {
	if x != nil {
		return x.ConditionValue
	}
	return 0
}

func (x *CapitalScheme) GetGiveMoney() float32 {
	if x != nil {
		return x.GiveMoney
	}
	return 0
}

func (x *CapitalScheme) GetGiveScore() int32 {
	if x != nil {
		return x.GiveScore
	}
	return 0
}

func (x *CapitalScheme) GetStock() int32 {
	if x != nil {
		return x.Stock
	}
	return 0
}

func (x *CapitalScheme) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *CapitalScheme) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *CapitalScheme) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

//查询参数
type CapitalSchemeWhere struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int32 `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	//以下为自定义参数
	StartAt       int64   `protobuf:"varint,3,opt,name=start_at,json=startAt,proto3" json:"start_at"`
	EndAt         string  `protobuf:"bytes,4,opt,name=end_at,json=endAt,proto3" json:"end_at"`
	ConditionType string  `protobuf:"bytes,5,opt,name=condition_type,json=conditionType,proto3" json:"condition_type"`
	Money         float32 `protobuf:"fixed32,6,opt,name=money,proto3" json:"money"`
	Name          string  `protobuf:"bytes,8,opt,name=name,proto3" json:"name"`
}

func (x *CapitalSchemeWhere) Reset() {
	*x = CapitalSchemeWhere{}
	if protoimpl.UnsafeEnabled {
		mi := &file_capitalSchemeService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CapitalSchemeWhere) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CapitalSchemeWhere) ProtoMessage() {}

func (x *CapitalSchemeWhere) ProtoReflect() protoreflect.Message {
	mi := &file_capitalSchemeService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CapitalSchemeWhere.ProtoReflect.Descriptor instead.
func (*CapitalSchemeWhere) Descriptor() ([]byte, []int) {
	return file_capitalSchemeService_proto_rawDescGZIP(), []int{1}
}

func (x *CapitalSchemeWhere) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *CapitalSchemeWhere) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *CapitalSchemeWhere) GetStartAt() int64 {
	if x != nil {
		return x.StartAt
	}
	return 0
}

func (x *CapitalSchemeWhere) GetEndAt() string {
	if x != nil {
		return x.EndAt
	}
	return ""
}

func (x *CapitalSchemeWhere) GetConditionType() string {
	if x != nil {
		return x.ConditionType
	}
	return ""
}

func (x *CapitalSchemeWhere) GetMoney() float32 {
	if x != nil {
		return x.Money
	}
	return 0
}

func (x *CapitalSchemeWhere) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

//
type CapitalSchemeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *CapitalScheme   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager    `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*CapitalScheme `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Error  *common.Error    `protobuf:"bytes,4,opt,name=error,proto3" json:"error"`
	Info   *common.Info     `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
}

func (x *CapitalSchemeResponse) Reset() {
	*x = CapitalSchemeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_capitalSchemeService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CapitalSchemeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CapitalSchemeResponse) ProtoMessage() {}

func (x *CapitalSchemeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_capitalSchemeService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CapitalSchemeResponse.ProtoReflect.Descriptor instead.
func (*CapitalSchemeResponse) Descriptor() ([]byte, []int) {
	return file_capitalSchemeService_proto_rawDescGZIP(), []int{2}
}

func (x *CapitalSchemeResponse) GetEntity() *CapitalScheme {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *CapitalSchemeResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *CapitalSchemeResponse) GetItems() []*CapitalScheme {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *CapitalSchemeResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *CapitalSchemeResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_capitalSchemeService_proto protoreflect.FileDescriptor

var file_capitalSchemeService_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x63, 0x61, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdb, 0x02, 0x0a, 0x0d, 0x43, 0x61,
	0x70, 0x69, 0x74, 0x61, 0x6c, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x41, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x65, 0x6e,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6e, 0x64, 0x41,
	0x74, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x69, 0x76, 0x65, 0x5f, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x67, 0x69, 0x76, 0x65, 0x4d, 0x6f, 0x6e, 0x65, 0x79,
	0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x69, 0x76, 0x65, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x67, 0x69, 0x76, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x73, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xca, 0x01, 0x0a, 0x12, 0x43, 0x61, 0x70, 0x69,
	0x74, 0x61, 0x6c, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x57, 0x68, 0x65, 0x72, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70,
	0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x41, 0x74, 0x12, 0x15, 0x0a, 0x06,
	0x65, 0x6e, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6e,
	0x64, 0x41, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f,
	0x6e, 0x65, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x65, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0xe3, 0x01, 0x0a, 0x15, 0x43, 0x61, 0x70, 0x69, 0x74, 0x61, 0x6c,
	0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f,
	0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x70, 0x69, 0x74, 0x61,
	0x6c, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12,
	0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70,
	0x61, 0x67, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43,
	0x61, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0xc6, 0x03, 0x0a, 0x14, 0x43,
	0x61, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x17, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x70, 0x69, 0x74, 0x61, 0x6c,
	0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x1a, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x43, 0x61, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x06, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43,
	0x61, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x1a, 0x1f, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x53,
	0x63, 0x68, 0x65, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x41, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x43, 0x61, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x1a,
	0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x70, 0x69, 0x74,
	0x61, 0x6c, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x49, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1c, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x53,
	0x63, 0x68, 0x65, 0x6d, 0x65, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x1f, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x53, 0x63, 0x68,
	0x65, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a,
	0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x43, 0x61, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x57, 0x68,
	0x65, 0x72, 0x65, 0x1a, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43,
	0x61, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x08, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x69,
	0x6e, 0x67, 0x12, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61,
	0x70, 0x69, 0x74, 0x61, 0x6c, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x57, 0x68, 0x65, 0x72, 0x65,
	0x1a, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x70, 0x69,
	0x74, 0x61, 0x6c, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_capitalSchemeService_proto_rawDescOnce sync.Once
	file_capitalSchemeService_proto_rawDescData = file_capitalSchemeService_proto_rawDesc
)

func file_capitalSchemeService_proto_rawDescGZIP() []byte {
	file_capitalSchemeService_proto_rawDescOnce.Do(func() {
		file_capitalSchemeService_proto_rawDescData = protoimpl.X.CompressGZIP(file_capitalSchemeService_proto_rawDescData)
	})
	return file_capitalSchemeService_proto_rawDescData
}

var file_capitalSchemeService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_capitalSchemeService_proto_goTypes = []interface{}{
	(*CapitalScheme)(nil),         // 0: services.CapitalScheme
	(*CapitalSchemeWhere)(nil),    // 1: services.CapitalSchemeWhere
	(*CapitalSchemeResponse)(nil), // 2: services.CapitalSchemeResponse
	(*common.Pager)(nil),          // 3: common.Pager
	(*common.Error)(nil),          // 4: common.Error
	(*common.Info)(nil),           // 5: common.Info
}
var file_capitalSchemeService_proto_depIdxs = []int32{
	0,  // 0: services.CapitalSchemeResponse.entity:type_name -> services.CapitalScheme
	3,  // 1: services.CapitalSchemeResponse.pager:type_name -> common.Pager
	0,  // 2: services.CapitalSchemeResponse.items:type_name -> services.CapitalScheme
	4,  // 3: services.CapitalSchemeResponse.error:type_name -> common.Error
	5,  // 4: services.CapitalSchemeResponse.info:type_name -> common.Info
	0,  // 5: services.CapitalSchemeService.Create:input_type -> services.CapitalScheme
	0,  // 6: services.CapitalSchemeService.Update:input_type -> services.CapitalScheme
	0,  // 7: services.CapitalSchemeService.Get:input_type -> services.CapitalScheme
	1,  // 8: services.CapitalSchemeService.Search:input_type -> services.CapitalSchemeWhere
	1,  // 9: services.CapitalSchemeService.List:input_type -> services.CapitalSchemeWhere
	1,  // 10: services.CapitalSchemeService.Matching:input_type -> services.CapitalSchemeWhere
	2,  // 11: services.CapitalSchemeService.Create:output_type -> services.CapitalSchemeResponse
	2,  // 12: services.CapitalSchemeService.Update:output_type -> services.CapitalSchemeResponse
	2,  // 13: services.CapitalSchemeService.Get:output_type -> services.CapitalSchemeResponse
	2,  // 14: services.CapitalSchemeService.Search:output_type -> services.CapitalSchemeResponse
	2,  // 15: services.CapitalSchemeService.List:output_type -> services.CapitalSchemeResponse
	2,  // 16: services.CapitalSchemeService.Matching:output_type -> services.CapitalSchemeResponse
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_capitalSchemeService_proto_init() }
func file_capitalSchemeService_proto_init() {
	if File_capitalSchemeService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_capitalSchemeService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CapitalScheme); i {
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
		file_capitalSchemeService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CapitalSchemeWhere); i {
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
		file_capitalSchemeService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CapitalSchemeResponse); i {
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
			RawDescriptor: file_capitalSchemeService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_capitalSchemeService_proto_goTypes,
		DependencyIndexes: file_capitalSchemeService_proto_depIdxs,
		MessageInfos:      file_capitalSchemeService_proto_msgTypes,
	}.Build()
	File_capitalSchemeService_proto = out.File
	file_capitalSchemeService_proto_rawDesc = nil
	file_capitalSchemeService_proto_goTypes = nil
	file_capitalSchemeService_proto_depIdxs = nil
}
