// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: unitService.proto

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

type Unit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	Type        string `protobuf:"bytes,3,opt,name=type,proto3" json:"type"`
	Source      string `protobuf:"bytes,4,opt,name=source,proto3" json:"source"`
	Sort        int32  `protobuf:"varint,5,opt,name=sort,proto3" json:"sort"`
	CreatorId   int64  `protobuf:"varint,6,opt,name=creator_id,json=creatorId,proto3" json:"creator_id"`
	CreatorName string `protobuf:"bytes,7,opt,name=creator_name,json=creatorName,proto3" json:"creator_name"`
	Status      string `protobuf:"bytes,8,opt,name=status,proto3" json:"status"`
	CreatedAt   string `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt   string `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
}

func (x *Unit) Reset() {
	*x = Unit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_unitService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unit) ProtoMessage() {}

func (x *Unit) ProtoReflect() protoreflect.Message {
	mi := &file_unitService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unit.ProtoReflect.Descriptor instead.
func (*Unit) Descriptor() ([]byte, []int) {
	return file_unitService_proto_rawDescGZIP(), []int{0}
}

func (x *Unit) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Unit) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Unit) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Unit) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Unit) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *Unit) GetCreatorId() int64 {
	if x != nil {
		return x.CreatorId
	}
	return 0
}

func (x *Unit) GetCreatorName() string {
	if x != nil {
		return x.CreatorName
	}
	return ""
}

func (x *Unit) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Unit) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Unit) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type UnitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged     int32  `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize  int32  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Keywords  string `protobuf:"bytes,3,opt,name=keywords,proto3" json:"keywords"`
	Id        int32  `protobuf:"varint,4,opt,name=id,proto3" json:"id"`
	Name      string `protobuf:"bytes,5,opt,name=name,proto3" json:"name"`
	Type      string `protobuf:"bytes,6,opt,name=type,proto3" json:"type"`
	Source    string `protobuf:"bytes,7,opt,name=source,proto3" json:"source"`
	CreatorId string `protobuf:"bytes,8,opt,name=creator_id,json=creatorId,proto3" json:"creator_id"`
	Status    string `protobuf:"bytes,9,opt,name=status,proto3" json:"status"`
}

func (x *UnitRequest) Reset() {
	*x = UnitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_unitService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnitRequest) ProtoMessage() {}

func (x *UnitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_unitService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnitRequest.ProtoReflect.Descriptor instead.
func (*UnitRequest) Descriptor() ([]byte, []int) {
	return file_unitService_proto_rawDescGZIP(), []int{1}
}

func (x *UnitRequest) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *UnitRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *UnitRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *UnitRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UnitRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UnitRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *UnitRequest) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *UnitRequest) GetCreatorId() string {
	if x != nil {
		return x.CreatorId
	}
	return ""
}

func (x *UnitRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type UnitData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Unit         `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*Unit       `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Info   *common.Info  `protobuf:"bytes,4,opt,name=info,proto3" json:"info"`
}

func (x *UnitData) Reset() {
	*x = UnitData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_unitService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnitData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnitData) ProtoMessage() {}

func (x *UnitData) ProtoReflect() protoreflect.Message {
	mi := &file_unitService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnitData.ProtoReflect.Descriptor instead.
func (*UnitData) Descriptor() ([]byte, []int) {
	return file_unitService_proto_rawDescGZIP(), []int{2}
}

func (x *UnitData) GetEntity() *Unit {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *UnitData) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *UnitData) GetItems() []*Unit {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *UnitData) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type UnitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *UnitData     `protobuf:"bytes,1,opt,name=data,proto3" json:"data"`
	Error *common.Error `protobuf:"bytes,2,opt,name=error,proto3" json:"error"`
}

func (x *UnitResponse) Reset() {
	*x = UnitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_unitService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnitResponse) ProtoMessage() {}

func (x *UnitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_unitService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnitResponse.ProtoReflect.Descriptor instead.
func (*UnitResponse) Descriptor() ([]byte, []int) {
	return file_unitService_proto_rawDescGZIP(), []int{3}
}

func (x *UnitResponse) GetData() *UnitData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *UnitResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_unitService_proto protoreflect.FileDescriptor

var file_unitService_proto_rawDesc = []byte{
	0x0a, 0x11, 0x75, 0x6e, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x82, 0x02, 0x0a, 0x04, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xe3, 0x01, 0x0a, 0x0b, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77,
	0x6f, 0x72, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77,
	0x6f, 0x72, 0x64, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f,
	0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x9f, 0x01, 0x0a, 0x08,
	0x55, 0x6e, 0x69, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x26, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05,
	0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x55, 0x6e, 0x69, 0x74, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x20, 0x0a, 0x04, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x5b, 0x0a,
	0x0c, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xc2, 0x02, 0x0a, 0x0b, 0x55,
	0x6e, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x55, 0x6e, 0x69, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x55, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x06,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30,
	0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2d, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x35, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x55, 0x6e, 0x69, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_unitService_proto_rawDescOnce sync.Once
	file_unitService_proto_rawDescData = file_unitService_proto_rawDesc
)

func file_unitService_proto_rawDescGZIP() []byte {
	file_unitService_proto_rawDescOnce.Do(func() {
		file_unitService_proto_rawDescData = protoimpl.X.CompressGZIP(file_unitService_proto_rawDescData)
	})
	return file_unitService_proto_rawDescData
}

var file_unitService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_unitService_proto_goTypes = []interface{}{
	(*Unit)(nil),         // 0: services.Unit
	(*UnitRequest)(nil),  // 1: services.UnitRequest
	(*UnitData)(nil),     // 2: services.UnitData
	(*UnitResponse)(nil), // 3: services.UnitResponse
	(*common.Pager)(nil), // 4: common.Pager
	(*common.Info)(nil),  // 5: common.Info
	(*common.Error)(nil), // 6: common.Error
}
var file_unitService_proto_depIdxs = []int32{
	0,  // 0: services.UnitData.entity:type_name -> services.Unit
	4,  // 1: services.UnitData.pager:type_name -> common.Pager
	0,  // 2: services.UnitData.items:type_name -> services.Unit
	5,  // 3: services.UnitData.info:type_name -> common.Info
	2,  // 4: services.UnitResponse.data:type_name -> services.UnitData
	6,  // 5: services.UnitResponse.error:type_name -> common.Error
	0,  // 6: services.UnitService.Create:input_type -> services.Unit
	0,  // 7: services.UnitService.Update:input_type -> services.Unit
	0,  // 8: services.UnitService.Delete:input_type -> services.Unit
	0,  // 9: services.UnitService.Get:input_type -> services.Unit
	1,  // 10: services.UnitService.List:input_type -> services.UnitRequest
	1,  // 11: services.UnitService.Search:input_type -> services.UnitRequest
	3,  // 12: services.UnitService.Create:output_type -> services.UnitResponse
	3,  // 13: services.UnitService.Update:output_type -> services.UnitResponse
	3,  // 14: services.UnitService.Delete:output_type -> services.UnitResponse
	3,  // 15: services.UnitService.Get:output_type -> services.UnitResponse
	3,  // 16: services.UnitService.List:output_type -> services.UnitResponse
	3,  // 17: services.UnitService.Search:output_type -> services.UnitResponse
	12, // [12:18] is the sub-list for method output_type
	6,  // [6:12] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_unitService_proto_init() }
func file_unitService_proto_init() {
	if File_unitService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_unitService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Unit); i {
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
		file_unitService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnitRequest); i {
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
		file_unitService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnitData); i {
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
		file_unitService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnitResponse); i {
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
			RawDescriptor: file_unitService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_unitService_proto_goTypes,
		DependencyIndexes: file_unitService_proto_depIdxs,
		MessageInfos:      file_unitService_proto_msgTypes,
	}.Build()
	File_unitService_proto = out.File
	file_unitService_proto_rawDesc = nil
	file_unitService_proto_goTypes = nil
	file_unitService_proto_depIdxs = nil
}