// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: attrGroupService.proto

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

type AttrGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CategoryId int64     `protobuf:"varint,3,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	Sort       int32     `protobuf:"varint,4,opt,name=sort,proto3" json:"sort,omitempty"`
	ParamCount int32     `protobuf:"varint,5,opt,name=param_count,json=paramCount,proto3" json:"param_count,omitempty"`
	Status     string    `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	CreatedAt  string    `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt  string    `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Category   *Category `protobuf:"bytes,9,opt,name=category,proto3" json:"category,omitempty"`
}

func (x *AttrGroup) Reset() {
	*x = AttrGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attrGroupService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttrGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttrGroup) ProtoMessage() {}

func (x *AttrGroup) ProtoReflect() protoreflect.Message {
	mi := &file_attrGroupService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttrGroup.ProtoReflect.Descriptor instead.
func (*AttrGroup) Descriptor() ([]byte, []int) {
	return file_attrGroupService_proto_rawDescGZIP(), []int{0}
}

func (x *AttrGroup) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AttrGroup) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AttrGroup) GetCategoryId() int64 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

func (x *AttrGroup) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *AttrGroup) GetParamCount() int32 {
	if x != nil {
		return x.ParamCount
	}
	return 0
}

func (x *AttrGroup) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *AttrGroup) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *AttrGroup) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *AttrGroup) GetCategory() *Category {
	if x != nil {
		return x.Category
	}
	return nil
}

type AttrGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged      int64  `protobuf:"varint,1,opt,name=paged,proto3" json:"paged,omitempty"`
	PageSize   int64  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Id         int32  `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	Name       string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	CategoryId int64  `protobuf:"varint,5,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
}

func (x *AttrGroupRequest) Reset() {
	*x = AttrGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attrGroupService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttrGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttrGroupRequest) ProtoMessage() {}

func (x *AttrGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_attrGroupService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttrGroupRequest.ProtoReflect.Descriptor instead.
func (*AttrGroupRequest) Descriptor() ([]byte, []int) {
	return file_attrGroupService_proto_rawDescGZIP(), []int{1}
}

func (x *AttrGroupRequest) GetPaged() int64 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *AttrGroupRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *AttrGroupRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AttrGroupRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AttrGroupRequest) GetCategoryId() int64 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

type AttrGroupData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *AttrGroup    `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items  []*AttrGroup  `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Info   *common.Info  `protobuf:"bytes,4,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *AttrGroupData) Reset() {
	*x = AttrGroupData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attrGroupService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttrGroupData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttrGroupData) ProtoMessage() {}

func (x *AttrGroupData) ProtoReflect() protoreflect.Message {
	mi := &file_attrGroupService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttrGroupData.ProtoReflect.Descriptor instead.
func (*AttrGroupData) Descriptor() ([]byte, []int) {
	return file_attrGroupService_proto_rawDescGZIP(), []int{2}
}

func (x *AttrGroupData) GetEntity() *AttrGroup {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *AttrGroupData) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *AttrGroupData) GetItems() []*AttrGroup {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *AttrGroupData) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type AttrGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *AttrGroupData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Error *common.Error  `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *AttrGroupResponse) Reset() {
	*x = AttrGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attrGroupService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttrGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttrGroupResponse) ProtoMessage() {}

func (x *AttrGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_attrGroupService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttrGroupResponse.ProtoReflect.Descriptor instead.
func (*AttrGroupResponse) Descriptor() ([]byte, []int) {
	return file_attrGroupService_proto_rawDescGZIP(), []int{3}
}

func (x *AttrGroupResponse) GetData() *AttrGroupData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *AttrGroupResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_attrGroupService_proto protoreflect.FileDescriptor

var file_attrGroupService_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x74, 0x74, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x02, 0x0a,
	0x09, 0x41, 0x74, 0x74, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73,
	0x6f, 0x72, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x2e, 0x0a, 0x08, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x22, 0x8a, 0x01, 0x0a, 0x10, 0x41,
	0x74, 0x74, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x22, 0xae, 0x01, 0x0a, 0x0d, 0x41, 0x74, 0x74, 0x72,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2b, 0x0a, 0x06, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x06,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50,
	0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x65, 0x0a, 0x11, 0x41, 0x74, 0x74, 0x72,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32,
	0xc2, 0x02, 0x0a, 0x10, 0x41, 0x74, 0x74, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x13,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41,
	0x74, 0x74, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3a, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x1a,
	0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x06,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x1a, 0x1b, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12,
	0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x41, 0x74, 0x74, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x41, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1a, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_attrGroupService_proto_rawDescOnce sync.Once
	file_attrGroupService_proto_rawDescData = file_attrGroupService_proto_rawDesc
)

func file_attrGroupService_proto_rawDescGZIP() []byte {
	file_attrGroupService_proto_rawDescOnce.Do(func() {
		file_attrGroupService_proto_rawDescData = protoimpl.X.CompressGZIP(file_attrGroupService_proto_rawDescData)
	})
	return file_attrGroupService_proto_rawDescData
}

var file_attrGroupService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_attrGroupService_proto_goTypes = []interface{}{
	(*AttrGroup)(nil),         // 0: services.AttrGroup
	(*AttrGroupRequest)(nil),  // 1: services.AttrGroupRequest
	(*AttrGroupData)(nil),     // 2: services.AttrGroupData
	(*AttrGroupResponse)(nil), // 3: services.AttrGroupResponse
	(*Category)(nil),          // 4: services.Category
	(*common.Pager)(nil),      // 5: common.Pager
	(*common.Info)(nil),       // 6: common.Info
	(*common.Error)(nil),      // 7: common.Error
}
var file_attrGroupService_proto_depIdxs = []int32{
	4,  // 0: services.AttrGroup.category:type_name -> services.Category
	0,  // 1: services.AttrGroupData.entity:type_name -> services.AttrGroup
	5,  // 2: services.AttrGroupData.pager:type_name -> common.Pager
	0,  // 3: services.AttrGroupData.items:type_name -> services.AttrGroup
	6,  // 4: services.AttrGroupData.info:type_name -> common.Info
	2,  // 5: services.AttrGroupResponse.data:type_name -> services.AttrGroupData
	7,  // 6: services.AttrGroupResponse.error:type_name -> common.Error
	0,  // 7: services.AttrGroupService.Create:input_type -> services.AttrGroup
	0,  // 8: services.AttrGroupService.Update:input_type -> services.AttrGroup
	0,  // 9: services.AttrGroupService.Delete:input_type -> services.AttrGroup
	0,  // 10: services.AttrGroupService.Get:input_type -> services.AttrGroup
	1,  // 11: services.AttrGroupService.Search:input_type -> services.AttrGroupRequest
	3,  // 12: services.AttrGroupService.Create:output_type -> services.AttrGroupResponse
	3,  // 13: services.AttrGroupService.Update:output_type -> services.AttrGroupResponse
	3,  // 14: services.AttrGroupService.Delete:output_type -> services.AttrGroupResponse
	3,  // 15: services.AttrGroupService.Get:output_type -> services.AttrGroupResponse
	3,  // 16: services.AttrGroupService.Search:output_type -> services.AttrGroupResponse
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_attrGroupService_proto_init() }
func file_attrGroupService_proto_init() {
	if File_attrGroupService_proto != nil {
		return
	}
	file_categoryService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_attrGroupService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttrGroup); i {
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
		file_attrGroupService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttrGroupRequest); i {
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
		file_attrGroupService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttrGroupData); i {
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
		file_attrGroupService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttrGroupResponse); i {
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
			RawDescriptor: file_attrGroupService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_attrGroupService_proto_goTypes,
		DependencyIndexes: file_attrGroupService_proto_depIdxs,
		MessageInfos:      file_attrGroupService_proto_msgTypes,
	}.Build()
	File_attrGroupService_proto = out.File
	file_attrGroupService_proto_rawDesc = nil
	file_attrGroupService_proto_goTypes = nil
	file_attrGroupService_proto_depIdxs = nil
}
