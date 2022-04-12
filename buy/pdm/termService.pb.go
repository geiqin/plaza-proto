// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: termService.proto

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

type Term struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32       `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Name        string      `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	Title       string      `protobuf:"bytes,3,opt,name=title,proto3" json:"title"`
	ParentId    int32       `protobuf:"varint,4,opt,name=parent_id,json=parentId,proto3" json:"parent_id"`
	Path        string      `protobuf:"bytes,5,opt,name=path,proto3" json:"path"`
	MediaId     int64       `protobuf:"varint,6,opt,name=media_id,json=mediaId,proto3" json:"media_id"`
	MediaUrl    string      `protobuf:"bytes,7,opt,name=media_url,json=mediaUrl,proto3" json:"media_url"`
	CategoryIds string      `protobuf:"bytes,8,opt,name=category_ids,json=categoryIds,proto3" json:"category_ids"`
	Memo        string      `protobuf:"bytes,9,opt,name=memo,proto3" json:"memo"`
	Sort        int32       `protobuf:"varint,10,opt,name=sort,proto3" json:"sort"`
	CreatedAt   string      `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt   string      `protobuf:"bytes,12,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	Parent      *Term       `protobuf:"bytes,13,opt,name=parent,proto3" json:"parent"`
	Children    []*Term     `protobuf:"bytes,14,rep,name=children,proto3" json:"children"`
	Categories  []*Category `protobuf:"bytes,15,rep,name=categories,proto3" json:"categories"`
}

func (x *Term) Reset() {
	*x = Term{}
	if protoimpl.UnsafeEnabled {
		mi := &file_termService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Term) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Term) ProtoMessage() {}

func (x *Term) ProtoReflect() protoreflect.Message {
	mi := &file_termService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Term.ProtoReflect.Descriptor instead.
func (*Term) Descriptor() ([]byte, []int) {
	return file_termService_proto_rawDescGZIP(), []int{0}
}

func (x *Term) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Term) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Term) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Term) GetParentId() int32 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *Term) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Term) GetMediaId() int64 {
	if x != nil {
		return x.MediaId
	}
	return 0
}

func (x *Term) GetMediaUrl() string {
	if x != nil {
		return x.MediaUrl
	}
	return ""
}

func (x *Term) GetCategoryIds() string {
	if x != nil {
		return x.CategoryIds
	}
	return ""
}

func (x *Term) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *Term) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *Term) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Term) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Term) GetParent() *Term {
	if x != nil {
		return x.Parent
	}
	return nil
}

func (x *Term) GetChildren() []*Term {
	if x != nil {
		return x.Children
	}
	return nil
}

func (x *Term) GetCategories() []*Category {
	if x != nil {
		return x.Categories
	}
	return nil
}

type TermRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int32  `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize int32  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Keywords string `protobuf:"bytes,3,opt,name=keywords,proto3" json:"keywords"`
	Id       int32  `protobuf:"varint,4,opt,name=id,proto3" json:"id"`
	Name     string `protobuf:"bytes,5,opt,name=name,proto3" json:"name"`
}

func (x *TermRequest) Reset() {
	*x = TermRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_termService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TermRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TermRequest) ProtoMessage() {}

func (x *TermRequest) ProtoReflect() protoreflect.Message {
	mi := &file_termService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TermRequest.ProtoReflect.Descriptor instead.
func (*TermRequest) Descriptor() ([]byte, []int) {
	return file_termService_proto_rawDescGZIP(), []int{1}
}

func (x *TermRequest) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *TermRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *TermRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *TermRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TermRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type TermData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Term         `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*Term       `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Info   *common.Info  `protobuf:"bytes,4,opt,name=info,proto3" json:"info"`
}

func (x *TermData) Reset() {
	*x = TermData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_termService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TermData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TermData) ProtoMessage() {}

func (x *TermData) ProtoReflect() protoreflect.Message {
	mi := &file_termService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TermData.ProtoReflect.Descriptor instead.
func (*TermData) Descriptor() ([]byte, []int) {
	return file_termService_proto_rawDescGZIP(), []int{2}
}

func (x *TermData) GetEntity() *Term {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *TermData) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *TermData) GetItems() []*Term {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *TermData) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type TermResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *TermData     `protobuf:"bytes,1,opt,name=data,proto3" json:"data"`
	Error *common.Error `protobuf:"bytes,2,opt,name=error,proto3" json:"error"`
}

func (x *TermResponse) Reset() {
	*x = TermResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_termService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TermResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TermResponse) ProtoMessage() {}

func (x *TermResponse) ProtoReflect() protoreflect.Message {
	mi := &file_termService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TermResponse.ProtoReflect.Descriptor instead.
func (*TermResponse) Descriptor() ([]byte, []int) {
	return file_termService_proto_rawDescGZIP(), []int{3}
}

func (x *TermResponse) GetData() *TermData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *TermResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_termService_proto protoreflect.FileDescriptor

var file_termService_proto_rawDesc = []byte{
	0x0a, 0x11, 0x74, 0x65, 0x72, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x15, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xba, 0x03, 0x0a, 0x04, 0x54, 0x65, 0x72, 0x6d,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x19, 0x0a, 0x08, 0x6d,
	0x65, 0x64, 0x69, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6d,
	0x65, 0x64, 0x69, 0x61, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x64, 0x69, 0x61,
	0x55, 0x72, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f,
	0x69, 0x64, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x49, 0x64, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f,
	0x72, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x26, 0x0a, 0x06,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x65, 0x72, 0x6d, 0x52, 0x06, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e,
	0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x54, 0x65, 0x72, 0x6d, 0x52, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e,
	0x12, 0x32, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x0f,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x69, 0x65, 0x73, 0x22, 0x80, 0x01, 0x0a, 0x0b, 0x54, 0x65, 0x72, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70,
	0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f,
	0x72, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f,
	0x72, 0x64, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x9f, 0x01, 0x0a, 0x08, 0x54, 0x65, 0x72, 0x6d,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x26, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x54, 0x65, 0x72, 0x6d, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05,
	0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65,
	0x72, 0x12, 0x24, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x65, 0x72, 0x6d,
	0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x5b, 0x0a, 0x0c, 0x54, 0x65, 0x72,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x54, 0x65, 0x72, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xc2, 0x02, 0x0a, 0x0b, 0x54, 0x65, 0x72, 0x6d, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x65, 0x72, 0x6d,
	0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x65, 0x72, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x65,
	0x72, 0x6d, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x65,
	0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x54, 0x65, 0x72, 0x6d, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x54, 0x65, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x03,
	0x47, 0x65, 0x74, 0x12, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54,
	0x65, 0x72, 0x6d, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54,
	0x65, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x04, 0x54,
	0x72, 0x65, 0x65, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54,
	0x65, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x65, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x15, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x65, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54,
	0x65, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2e,
	0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_termService_proto_rawDescOnce sync.Once
	file_termService_proto_rawDescData = file_termService_proto_rawDesc
)

func file_termService_proto_rawDescGZIP() []byte {
	file_termService_proto_rawDescOnce.Do(func() {
		file_termService_proto_rawDescData = protoimpl.X.CompressGZIP(file_termService_proto_rawDescData)
	})
	return file_termService_proto_rawDescData
}

var file_termService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_termService_proto_goTypes = []interface{}{
	(*Term)(nil),         // 0: services.Term
	(*TermRequest)(nil),  // 1: services.TermRequest
	(*TermData)(nil),     // 2: services.TermData
	(*TermResponse)(nil), // 3: services.TermResponse
	(*Category)(nil),     // 4: services.Category
	(*common.Pager)(nil), // 5: common.Pager
	(*common.Info)(nil),  // 6: common.Info
	(*common.Error)(nil), // 7: common.Error
}
var file_termService_proto_depIdxs = []int32{
	0,  // 0: services.Term.parent:type_name -> services.Term
	0,  // 1: services.Term.children:type_name -> services.Term
	4,  // 2: services.Term.categories:type_name -> services.Category
	0,  // 3: services.TermData.entity:type_name -> services.Term
	5,  // 4: services.TermData.pager:type_name -> common.Pager
	0,  // 5: services.TermData.items:type_name -> services.Term
	6,  // 6: services.TermData.info:type_name -> common.Info
	2,  // 7: services.TermResponse.data:type_name -> services.TermData
	7,  // 8: services.TermResponse.error:type_name -> common.Error
	0,  // 9: services.TermService.Create:input_type -> services.Term
	0,  // 10: services.TermService.Update:input_type -> services.Term
	0,  // 11: services.TermService.Delete:input_type -> services.Term
	0,  // 12: services.TermService.Get:input_type -> services.Term
	1,  // 13: services.TermService.Tree:input_type -> services.TermRequest
	1,  // 14: services.TermService.Search:input_type -> services.TermRequest
	3,  // 15: services.TermService.Create:output_type -> services.TermResponse
	3,  // 16: services.TermService.Update:output_type -> services.TermResponse
	3,  // 17: services.TermService.Delete:output_type -> services.TermResponse
	3,  // 18: services.TermService.Get:output_type -> services.TermResponse
	3,  // 19: services.TermService.Tree:output_type -> services.TermResponse
	3,  // 20: services.TermService.Search:output_type -> services.TermResponse
	15, // [15:21] is the sub-list for method output_type
	9,  // [9:15] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_termService_proto_init() }
func file_termService_proto_init() {
	if File_termService_proto != nil {
		return
	}
	file_categoryService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_termService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Term); i {
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
		file_termService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TermRequest); i {
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
		file_termService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TermData); i {
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
		file_termService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TermResponse); i {
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
			RawDescriptor: file_termService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_termService_proto_goTypes,
		DependencyIndexes: file_termService_proto_depIdxs,
		MessageInfos:      file_termService_proto_msgTypes,
	}.Build()
	File_termService_proto = out.File
	file_termService_proto_rawDesc = nil
	file_termService_proto_goTypes = nil
	file_termService_proto_depIdxs = nil
}