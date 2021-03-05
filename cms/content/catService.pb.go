// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: catService.proto

package services

import (
	common "github.com/geiqin/microkit/protobuf/common"
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

type CatWhere struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int32  `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize int32  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Top      int32  `protobuf:"varint,3,opt,name=top,proto3" json:"top"`
	Slug     string `protobuf:"bytes,4,opt,name=slug,proto3" json:"slug"`
	// @inject_tag: gorm:"-"
	Slugs []string `protobuf:"bytes,5,rep,name=slugs,proto3" json:"slugs" gorm:"-"`
	// @inject_tag: gorm:"-"
	Ids []int32 `protobuf:"varint,6,rep,packed,name=ids,proto3" json:"ids" gorm:"-"`
}

func (x *CatWhere) Reset() {
	*x = CatWhere{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CatWhere) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CatWhere) ProtoMessage() {}

func (x *CatWhere) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CatWhere.ProtoReflect.Descriptor instead.
func (*CatWhere) Descriptor() ([]byte, []int) {
	return file_catService_proto_rawDescGZIP(), []int{0}
}

func (x *CatWhere) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *CatWhere) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *CatWhere) GetTop() int32 {
	if x != nil {
		return x.Top
	}
	return 0
}

func (x *CatWhere) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *CatWhere) GetSlugs() []string {
	if x != nil {
		return x.Slugs
	}
	return nil
}

func (x *CatWhere) GetIds() []int32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

// 字典信息
type Cat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Name         string `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	Slug         string `protobuf:"bytes,3,opt,name=slug,proto3" json:"slug"`
	CatId        int32  `protobuf:"varint,4,opt,name=cat_id,json=catId,proto3" json:"cat_id"`
	ThumbId      int64  `protobuf:"varint,5,opt,name=thumb_id,json=thumbId,proto3" json:"thumb_id"`
	ThumbUrl     string `protobuf:"bytes,6,opt,name=thumb_url,json=thumbUrl,proto3" json:"thumb_url"`
	Depth        string `protobuf:"bytes,7,opt,name=depth,proto3" json:"depth"`
	Description  string `protobuf:"bytes,8,opt,name=description,proto3" json:"description"`
	System       bool   `protobuf:"varint,17,opt,name=system,proto3" json:"system"`
	Sorting      int32  `protobuf:"varint,9,opt,name=sorting,proto3" json:"sorting"`
	ArticleCount int32  `protobuf:"varint,10,opt,name=article_count,json=articleCount,proto3" json:"article_count"`
	CreatedAt    string `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt    string `protobuf:"bytes,12,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	Children     []*Cat `protobuf:"bytes,13,rep,name=children,proto3" json:"children"`
	// @inject_tag: gorm:"-"
	Ids []int32 `protobuf:"varint,14,rep,packed,name=ids,proto3" json:"ids" gorm:"-"`
}

func (x *Cat) Reset() {
	*x = Cat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cat) ProtoMessage() {}

func (x *Cat) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Cat.ProtoReflect.Descriptor instead.
func (*Cat) Descriptor() ([]byte, []int) {
	return file_catService_proto_rawDescGZIP(), []int{1}
}

func (x *Cat) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Cat) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Cat) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *Cat) GetCatId() int32 {
	if x != nil {
		return x.CatId
	}
	return 0
}

func (x *Cat) GetThumbId() int64 {
	if x != nil {
		return x.ThumbId
	}
	return 0
}

func (x *Cat) GetThumbUrl() string {
	if x != nil {
		return x.ThumbUrl
	}
	return ""
}

func (x *Cat) GetDepth() string {
	if x != nil {
		return x.Depth
	}
	return ""
}

func (x *Cat) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Cat) GetSystem() bool {
	if x != nil {
		return x.System
	}
	return false
}

func (x *Cat) GetSorting() int32 {
	if x != nil {
		return x.Sorting
	}
	return 0
}

func (x *Cat) GetArticleCount() int32 {
	if x != nil {
		return x.ArticleCount
	}
	return 0
}

func (x *Cat) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Cat) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Cat) GetChildren() []*Cat {
	if x != nil {
		return x.Children
	}
	return nil
}

func (x *Cat) GetIds() []int32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

//
type CatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Cat          `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*Cat        `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Error  *common.Error `protobuf:"bytes,4,opt,name=error,proto3" json:"error"`
	Info   *common.Info  `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
}

func (x *CatResponse) Reset() {
	*x = CatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CatResponse) ProtoMessage() {}

func (x *CatResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CatResponse.ProtoReflect.Descriptor instead.
func (*CatResponse) Descriptor() ([]byte, []int) {
	return file_catService_proto_rawDescGZIP(), []int{2}
}

func (x *CatResponse) GetEntity() *Cat {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *CatResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *CatResponse) GetItems() []*Cat {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *CatResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *CatResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_catService_proto protoreflect.FileDescriptor

var file_catService_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x8b, 0x01, 0x0a, 0x08, 0x43, 0x61, 0x74, 0x57, 0x68, 0x65, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x67,
	0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x74, 0x6f, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x74, 0x6f,
	0x70, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x73, 0x6c, 0x75, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x6c, 0x75, 0x67, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x73, 0x6c, 0x75, 0x67, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x69,
	0x64, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x05, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x96, 0x03,
	0x0a, 0x03, 0x43, 0x61, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75,
	0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x12, 0x15, 0x0a,
	0x06, 0x63, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63,
	0x61, 0x74, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x5f, 0x69, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x49, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x55, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05,
	0x64, 0x65, 0x70, 0x74, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x64, 0x65, 0x70,
	0x74, 0x68, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x18, 0x11,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73,
	0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x29, 0x0a, 0x08, 0x63, 0x68, 0x69,
	0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x74, 0x52, 0x08, 0x63, 0x68, 0x69, 0x6c,
	0x64, 0x72, 0x65, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x0e, 0x20, 0x03, 0x28,
	0x05, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0xc5, 0x01, 0x0a, 0x0b, 0x43, 0x61, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x43, 0x61, 0x74, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a,
	0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67,
	0x65, 0x72, 0x12, 0x23, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x74,
	0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x04,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0xeb,
	0x02, 0x0a, 0x0a, 0x43, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x30, 0x0a,
	0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x43, 0x61, 0x74, 0x1a, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x43, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x30, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0d, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x74, 0x1a, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x30, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0d, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x74, 0x1a, 0x15, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0d, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x74, 0x1a, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x34, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x11, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a,
	0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x57, 0x68,
	0x65, 0x72, 0x65, 0x1a, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43,
	0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x04,
	0x54, 0x72, 0x65, 0x65, 0x12, 0x0d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x43, 0x61, 0x74, 0x1a, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43,
	0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a,
	0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
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

var file_catService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_catService_proto_goTypes = []interface{}{
	(*CatWhere)(nil),         // 0: services.CatWhere
	(*Cat)(nil),              // 1: services.Cat
	(*CatResponse)(nil),      // 2: services.CatResponse
	(*common.Pager)(nil),     // 3: common.Pager
	(*common.Error)(nil),     // 4: common.Error
	(*common.Info)(nil),      // 5: common.Info
	(*common.BaseWhere)(nil), // 6: common.BaseWhere
}
var file_catService_proto_depIdxs = []int32{
	1,  // 0: services.Cat.children:type_name -> services.Cat
	1,  // 1: services.CatResponse.entity:type_name -> services.Cat
	3,  // 2: services.CatResponse.pager:type_name -> common.Pager
	1,  // 3: services.CatResponse.items:type_name -> services.Cat
	4,  // 4: services.CatResponse.error:type_name -> common.Error
	5,  // 5: services.CatResponse.info:type_name -> common.Info
	1,  // 6: services.CatService.Create:input_type -> services.Cat
	1,  // 7: services.CatService.Update:input_type -> services.Cat
	1,  // 8: services.CatService.Delete:input_type -> services.Cat
	1,  // 9: services.CatService.Get:input_type -> services.Cat
	6,  // 10: services.CatService.Search:input_type -> common.BaseWhere
	6,  // 11: services.CatService.List:input_type -> common.BaseWhere
	1,  // 12: services.CatService.Tree:input_type -> services.Cat
	2,  // 13: services.CatService.Create:output_type -> services.CatResponse
	2,  // 14: services.CatService.Update:output_type -> services.CatResponse
	2,  // 15: services.CatService.Delete:output_type -> services.CatResponse
	2,  // 16: services.CatService.Get:output_type -> services.CatResponse
	2,  // 17: services.CatService.Search:output_type -> services.CatResponse
	2,  // 18: services.CatService.List:output_type -> services.CatResponse
	2,  // 19: services.CatService.Tree:output_type -> services.CatResponse
	13, // [13:20] is the sub-list for method output_type
	6,  // [6:13] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_catService_proto_init() }
func file_catService_proto_init() {
	if File_catService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_catService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CatWhere); i {
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
		file_catService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
			NumMessages:   3,
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
