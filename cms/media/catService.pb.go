// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: catService.proto

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

// 字典信息
type Cat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Type      string `protobuf:"bytes,2,opt,name=type,proto3" json:"type"`
	Slug      string `protobuf:"bytes,3,opt,name=slug,proto3" json:"slug"`
	Name      string `protobuf:"bytes,4,opt,name=name,proto3" json:"name"`
	Path      string `protobuf:"bytes,5,opt,name=path,proto3" json:"path"`
	ParentId  int64  `protobuf:"varint,6,opt,name=parent_id,json=parentId,proto3" json:"parent_id"`
	TotalNum  int32  `protobuf:"varint,7,opt,name=total_num,json=totalNum,proto3" json:"total_num"`
	TotalSize int64  `protobuf:"varint,8,opt,name=total_size,json=totalSize,proto3" json:"total_size"`
	CreatorId int64  `protobuf:"varint,9,opt,name=creator_id,json=creatorId,proto3" json:"creator_id"`
	Parent    *Cat   `protobuf:"bytes,10,opt,name=parent,proto3" json:"parent"`
	Children  []*Cat `protobuf:"bytes,11,rep,name=children,proto3" json:"children"`
}

func (x *Cat) Reset() {
	*x = Cat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cat) ProtoMessage() {}

func (x *Cat) ProtoReflect() protoreflect.Message {
	mi := &file_catService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cat.ProtoReflect.Descriptor instead.
func (*Cat) Descriptor() ([]byte, []int) {
	return file_catService_proto_rawDescGZIP(), []int{0}
}

func (x *Cat) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Cat) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Cat) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *Cat) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Cat) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Cat) GetParentId() int64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *Cat) GetTotalNum() int32 {
	if x != nil {
		return x.TotalNum
	}
	return 0
}

func (x *Cat) GetTotalSize() int64 {
	if x != nil {
		return x.TotalSize
	}
	return 0
}

func (x *Cat) GetCreatorId() int64 {
	if x != nil {
		return x.CreatorId
	}
	return 0
}

func (x *Cat) GetParent() *Cat {
	if x != nil {
		return x.Parent
	}
	return nil
}

func (x *Cat) GetChildren() []*Cat {
	if x != nil {
		return x.Children
	}
	return nil
}

type CatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int32 `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Top      int32 `protobuf:"varint,3,opt,name=top,proto3" json:"top"`
	//file params
	Id       int64   `protobuf:"varint,4,opt,name=id,proto3" json:"id"`
	Name     string  `protobuf:"bytes,5,opt,name=name,proto3" json:"name"`
	Slug     string  `protobuf:"bytes,6,opt,name=slug,proto3" json:"slug"`
	Type     string  `protobuf:"bytes,7,opt,name=type,proto3" json:"type"`
	ParentId int64   `protobuf:"varint,9,opt,name=parent_id,json=parentId,proto3" json:"parent_id"`
	Ids      []int64 `protobuf:"varint,11,rep,packed,name=ids,proto3" json:"ids"`
}

func (x *CatRequest) Reset() {
	*x = CatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CatRequest) ProtoMessage() {}

func (x *CatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_catService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CatRequest.ProtoReflect.Descriptor instead.
func (*CatRequest) Descriptor() ([]byte, []int) {
	return file_catService_proto_rawDescGZIP(), []int{1}
}

func (x *CatRequest) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *CatRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *CatRequest) GetTop() int32 {
	if x != nil {
		return x.Top
	}
	return 0
}

func (x *CatRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CatRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CatRequest) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *CatRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CatRequest) GetParentId() int64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *CatRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type CatData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Cat          `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*Cat        `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Info   *common.Info  `protobuf:"bytes,4,opt,name=info,proto3" json:"info"`
}

func (x *CatData) Reset() {
	*x = CatData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CatData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CatData) ProtoMessage() {}

func (x *CatData) ProtoReflect() protoreflect.Message {
	mi := &file_catService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CatData.ProtoReflect.Descriptor instead.
func (*CatData) Descriptor() ([]byte, []int) {
	return file_catService_proto_rawDescGZIP(), []int{2}
}

func (x *CatData) GetEntity() *Cat {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *CatData) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *CatData) GetItems() []*Cat {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *CatData) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type CatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *CatData      `protobuf:"bytes,1,opt,name=data,proto3" json:"data"`
	Error *common.Error `protobuf:"bytes,2,opt,name=error,proto3" json:"error"`
}

func (x *CatResponse) Reset() {
	*x = CatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CatResponse) ProtoMessage() {}

func (x *CatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_catService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CatResponse.ProtoReflect.Descriptor instead.
func (*CatResponse) Descriptor() ([]byte, []int) {
	return file_catService_proto_rawDescGZIP(), []int{3}
}

func (x *CatResponse) GetData() *CatData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *CatResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_catService_proto protoreflect.FileDescriptor

var file_catService_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xaf, 0x02, 0x0a, 0x03, 0x43, 0x61, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x6c, 0x75, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x6e, 0x75,
	0x6d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4e, 0x75,
	0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12,
	0x25, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x74, 0x52, 0x06,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72,
	0x65, 0x6e, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x74, 0x52, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65,
	0x6e, 0x22, 0xcc, 0x01, 0x0a, 0x0a, 0x43, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x6f, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x03, 0x74, 0x6f, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75,
	0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73,
	0x22, 0x9c, 0x01, 0x0a, 0x07, 0x43, 0x61, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x25, 0x0a, 0x06,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x74, 0x52, 0x06, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65,
	0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x43, 0x61, 0x74, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x20, 0x0a,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22,
	0x59, 0x0a, 0x0b, 0x43, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xea, 0x02, 0x0a, 0x0a, 0x43,
	0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x0d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43,
	0x61, 0x74, 0x1a, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x06, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x43, 0x61, 0x74, 0x1a, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x43, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x30, 0x0a,
	0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x43, 0x61, 0x74, 0x1a, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x43, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x2d, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x43, 0x61, 0x74, 0x1a, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x43, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2e,
	0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x43, 0x61, 0x74, 0x1a, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x43, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2e,
	0x0a, 0x04, 0x54, 0x72, 0x65, 0x65, 0x12, 0x0d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x43, 0x61, 0x74, 0x1a, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x43, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37,
	0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_catService_proto_rawDescOnce sync.Once
	file_catService_proto_rawDescData = file_catService_proto_rawDesc
)

func file_catService_proto_rawDescGZIP() []byte {
	file_catService_proto_rawDescOnce.Do(func() {
		file_catService_proto_rawDescData = protoimpl.X.CompressGZIP(file_catService_proto_rawDescData)
	})
	return file_catService_proto_rawDescData
}

var file_catService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_catService_proto_goTypes = []interface{}{
	(*Cat)(nil),          // 0: services.Cat
	(*CatRequest)(nil),   // 1: services.CatRequest
	(*CatData)(nil),      // 2: services.CatData
	(*CatResponse)(nil),  // 3: services.CatResponse
	(*common.Pager)(nil), // 4: common.Pager
	(*common.Info)(nil),  // 5: common.Info
	(*common.Error)(nil), // 6: common.Error
}
var file_catService_proto_depIdxs = []int32{
	0,  // 0: services.Cat.parent:type_name -> services.Cat
	0,  // 1: services.Cat.children:type_name -> services.Cat
	0,  // 2: services.CatData.entity:type_name -> services.Cat
	4,  // 3: services.CatData.pager:type_name -> common.Pager
	0,  // 4: services.CatData.items:type_name -> services.Cat
	5,  // 5: services.CatData.info:type_name -> common.Info
	2,  // 6: services.CatResponse.data:type_name -> services.CatData
	6,  // 7: services.CatResponse.error:type_name -> common.Error
	0,  // 8: services.CatService.Create:input_type -> services.Cat
	0,  // 9: services.CatService.Update:input_type -> services.Cat
	0,  // 10: services.CatService.Delete:input_type -> services.Cat
	0,  // 11: services.CatService.Get:input_type -> services.Cat
	0,  // 12: services.CatService.List:input_type -> services.Cat
	0,  // 13: services.CatService.Tree:input_type -> services.Cat
	1,  // 14: services.CatService.Search:input_type -> services.CatRequest
	3,  // 15: services.CatService.Create:output_type -> services.CatResponse
	3,  // 16: services.CatService.Update:output_type -> services.CatResponse
	3,  // 17: services.CatService.Delete:output_type -> services.CatResponse
	3,  // 18: services.CatService.Get:output_type -> services.CatResponse
	3,  // 19: services.CatService.List:output_type -> services.CatResponse
	3,  // 20: services.CatService.Tree:output_type -> services.CatResponse
	3,  // 21: services.CatService.Search:output_type -> services.CatResponse
	15, // [15:22] is the sub-list for method output_type
	8,  // [8:15] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_catService_proto_init() }
func file_catService_proto_init() {
	if File_catService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_catService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cat); i {
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
		file_catService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CatRequest); i {
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
		file_catService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CatData); i {
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
		file_catService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CatResponse); i {
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
			RawDescriptor: file_catService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_catService_proto_goTypes,
		DependencyIndexes: file_catService_proto_depIdxs,
		MessageInfos:      file_catService_proto_msgTypes,
	}.Build()
	File_catService_proto = out.File
	file_catService_proto_rawDesc = nil
	file_catService_proto_goTypes = nil
	file_catService_proto_depIdxs = nil
}
