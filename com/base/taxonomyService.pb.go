// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: taxonomyService.proto

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

type TaxonomyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int64   `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize int64   `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Sorting  string  `protobuf:"bytes,3,opt,name=sorting,proto3" json:"sorting"`
	Keywords string  `protobuf:"bytes,4,opt,name=keywords,proto3" json:"keywords"`
	Id       int64   `protobuf:"varint,5,opt,name=id,proto3" json:"id"`
	ParentId int64   `protobuf:"varint,6,opt,name=parent_id,json=parentId,proto3" json:"parent_id"`
	SkuType  string  `protobuf:"bytes,7,opt,name=sku_type,json=skuType,proto3" json:"sku_type"`
	Name     bool    `protobuf:"varint,8,opt,name=name,proto3" json:"name"`
	Ids      []int64 `protobuf:"varint,9,rep,packed,name=ids,proto3" json:"ids"`
}

func (x *TaxonomyRequest) Reset() {
	*x = TaxonomyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_taxonomyService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaxonomyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaxonomyRequest) ProtoMessage() {}

func (x *TaxonomyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_taxonomyService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaxonomyRequest.ProtoReflect.Descriptor instead.
func (*TaxonomyRequest) Descriptor() ([]byte, []int) {
	return file_taxonomyService_proto_rawDescGZIP(), []int{0}
}

func (x *TaxonomyRequest) GetPaged() int64 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *TaxonomyRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *TaxonomyRequest) GetSorting() string {
	if x != nil {
		return x.Sorting
	}
	return ""
}

func (x *TaxonomyRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *TaxonomyRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TaxonomyRequest) GetParentId() int64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *TaxonomyRequest) GetSkuType() string {
	if x != nil {
		return x.SkuType
	}
	return ""
}

func (x *TaxonomyRequest) GetName() bool {
	if x != nil {
		return x.Name
	}
	return false
}

func (x *TaxonomyRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type NewTaxonomyIndex struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaxId    int64 `protobuf:"varint,1,opt,name=max_id,json=maxId,proto3" json:"max_id"`
	NewId    int64 `protobuf:"varint,2,opt,name=new_id,json=newId,proto3" json:"new_id"`
	NewIndex int64 `protobuf:"varint,3,opt,name=new_index,json=newIndex,proto3" json:"new_index"`
}

func (x *NewTaxonomyIndex) Reset() {
	*x = NewTaxonomyIndex{}
	if protoimpl.UnsafeEnabled {
		mi := &file_taxonomyService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewTaxonomyIndex) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewTaxonomyIndex) ProtoMessage() {}

func (x *NewTaxonomyIndex) ProtoReflect() protoreflect.Message {
	mi := &file_taxonomyService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewTaxonomyIndex.ProtoReflect.Descriptor instead.
func (*NewTaxonomyIndex) Descriptor() ([]byte, []int) {
	return file_taxonomyService_proto_rawDescGZIP(), []int{1}
}

func (x *NewTaxonomyIndex) GetMaxId() int64 {
	if x != nil {
		return x.MaxId
	}
	return 0
}

func (x *NewTaxonomyIndex) GetNewId() int64 {
	if x != nil {
		return x.NewId
	}
	return 0
}

func (x *NewTaxonomyIndex) GetNewIndex() int64 {
	if x != nil {
		return x.NewIndex
	}
	return 0
}

type Taxonomy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64       `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Name      string      `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	ParentId  int64       `protobuf:"varint,4,opt,name=parent_id,json=parentId,proto3" json:"parent_id"`
	Path      string      `protobuf:"bytes,5,opt,name=path,proto3" json:"path"`
	Depth     int32       `protobuf:"varint,6,opt,name=depth,proto3" json:"depth"`
	Letter    string      `protobuf:"bytes,7,opt,name=letter,proto3" json:"letter"`
	SkuType   string      `protobuf:"bytes,8,opt,name=sku_type,json=skuType,proto3" json:"sku_type"`
	IsVirtual bool        `protobuf:"varint,9,opt,name=is_virtual,json=isVirtual,proto3" json:"is_virtual"`
	Memo      string      `protobuf:"bytes,10,opt,name=memo,proto3" json:"memo"`
	CreatedAt string      `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt string      `protobuf:"bytes,12,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	Parent    *Taxonomy   `protobuf:"bytes,13,opt,name=parent,proto3" json:"parent"`
	Children  []*Taxonomy `protobuf:"bytes,14,rep,name=children,proto3" json:"children"`
}

func (x *Taxonomy) Reset() {
	*x = Taxonomy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_taxonomyService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Taxonomy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Taxonomy) ProtoMessage() {}

func (x *Taxonomy) ProtoReflect() protoreflect.Message {
	mi := &file_taxonomyService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Taxonomy.ProtoReflect.Descriptor instead.
func (*Taxonomy) Descriptor() ([]byte, []int) {
	return file_taxonomyService_proto_rawDescGZIP(), []int{2}
}

func (x *Taxonomy) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Taxonomy) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Taxonomy) GetParentId() int64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *Taxonomy) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Taxonomy) GetDepth() int32 {
	if x != nil {
		return x.Depth
	}
	return 0
}

func (x *Taxonomy) GetLetter() string {
	if x != nil {
		return x.Letter
	}
	return ""
}

func (x *Taxonomy) GetSkuType() string {
	if x != nil {
		return x.SkuType
	}
	return ""
}

func (x *Taxonomy) GetIsVirtual() bool {
	if x != nil {
		return x.IsVirtual
	}
	return false
}

func (x *Taxonomy) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *Taxonomy) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Taxonomy) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Taxonomy) GetParent() *Taxonomy {
	if x != nil {
		return x.Parent
	}
	return nil
}

func (x *Taxonomy) GetChildren() []*Taxonomy {
	if x != nil {
		return x.Children
	}
	return nil
}

type TaxonomyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity   *Taxonomy     `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager    *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items    []*Taxonomy   `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Info     string        `protobuf:"bytes,4,opt,name=info,proto3" json:"info"`
	NewIndex int32         `protobuf:"varint,5,opt,name=new_index,json=newIndex,proto3" json:"new_index"`
}

func (x *TaxonomyResponse) Reset() {
	*x = TaxonomyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_taxonomyService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaxonomyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaxonomyResponse) ProtoMessage() {}

func (x *TaxonomyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_taxonomyService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaxonomyResponse.ProtoReflect.Descriptor instead.
func (*TaxonomyResponse) Descriptor() ([]byte, []int) {
	return file_taxonomyService_proto_rawDescGZIP(), []int{3}
}

func (x *TaxonomyResponse) GetEntity() *Taxonomy {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *TaxonomyResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *TaxonomyResponse) GetItems() []*Taxonomy {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *TaxonomyResponse) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

func (x *TaxonomyResponse) GetNewIndex() int32 {
	if x != nil {
		return x.NewIndex
	}
	return 0
}

var File_taxonomyService_proto protoreflect.FileDescriptor

var file_taxonomyService_proto_rawDesc = []byte{
	0x0a, 0x15, 0x74, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe8, 0x01, 0x0a, 0x0f, 0x54, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6f,
	0x72, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64,
	0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x19,
	0x0a, 0x08, 0x73, 0x6b, 0x75, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x73, 0x6b, 0x75, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x69, 0x64, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22,
	0x5d, 0x0a, 0x10, 0x4e, 0x65, 0x77, 0x54, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x15, 0x0a, 0x06, 0x6d, 0x61, 0x78, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x6d, 0x61, 0x78, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x6e, 0x65,
	0x77, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6e, 0x65, 0x77, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6e, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0xf5,
	0x02, 0x0a, 0x08, 0x54, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x70, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x64, 0x65, 0x70, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x72,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x12, 0x19,
	0x0a, 0x08, 0x73, 0x6b, 0x75, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x73, 0x6b, 0x75, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f,
	0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69,
	0x73, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x2a, 0x0a, 0x06, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x52, 0x06,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72,
	0x65, 0x6e, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x54, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x52, 0x08, 0x63, 0x68,
	0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x22, 0xbe, 0x01, 0x0a, 0x10, 0x54, 0x61, 0x78, 0x6f, 0x6e,
	0x6f, 0x6d, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x52,
	0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x65,
	0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6e,
	0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x32, 0xc3, 0x03, 0x0a, 0x0f, 0x54, 0x61, 0x78, 0x6f,
	0x6e, 0x6f, 0x6d, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x06, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x54, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x1a, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x61, 0x78,
	0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x1a, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x54, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x12, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x1a, 0x1a,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x61, 0x78, 0x6f, 0x6e, 0x6f,
	0x6d, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x06,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x54, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x1a, 0x1a, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x12, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x61,
	0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x04, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54,
	0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x61, 0x78, 0x6f, 0x6e, 0x6f,
	0x6d, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x04,
	0x54, 0x72, 0x65, 0x65, 0x12, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x54, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x61, 0x78, 0x6f, 0x6e,
	0x6f, 0x6d, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a,
	0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_taxonomyService_proto_rawDescOnce sync.Once
	file_taxonomyService_proto_rawDescData = file_taxonomyService_proto_rawDesc
)

func file_taxonomyService_proto_rawDescGZIP() []byte {
	file_taxonomyService_proto_rawDescOnce.Do(func() {
		file_taxonomyService_proto_rawDescData = protoimpl.X.CompressGZIP(file_taxonomyService_proto_rawDescData)
	})
	return file_taxonomyService_proto_rawDescData
}

var file_taxonomyService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_taxonomyService_proto_goTypes = []interface{}{
	(*TaxonomyRequest)(nil),  // 0: services.TaxonomyRequest
	(*NewTaxonomyIndex)(nil), // 1: services.NewTaxonomyIndex
	(*Taxonomy)(nil),         // 2: services.Taxonomy
	(*TaxonomyResponse)(nil), // 3: services.TaxonomyResponse
	(*common.Pager)(nil),     // 4: common.Pager
}
var file_taxonomyService_proto_depIdxs = []int32{
	2,  // 0: services.Taxonomy.parent:type_name -> services.Taxonomy
	2,  // 1: services.Taxonomy.children:type_name -> services.Taxonomy
	2,  // 2: services.TaxonomyResponse.entity:type_name -> services.Taxonomy
	4,  // 3: services.TaxonomyResponse.pager:type_name -> common.Pager
	2,  // 4: services.TaxonomyResponse.items:type_name -> services.Taxonomy
	2,  // 5: services.TaxonomyService.Create:input_type -> services.Taxonomy
	2,  // 6: services.TaxonomyService.Update:input_type -> services.Taxonomy
	2,  // 7: services.TaxonomyService.Get:input_type -> services.Taxonomy
	2,  // 8: services.TaxonomyService.Delete:input_type -> services.Taxonomy
	0,  // 9: services.TaxonomyService.Search:input_type -> services.TaxonomyRequest
	0,  // 10: services.TaxonomyService.List:input_type -> services.TaxonomyRequest
	0,  // 11: services.TaxonomyService.Tree:input_type -> services.TaxonomyRequest
	3,  // 12: services.TaxonomyService.Create:output_type -> services.TaxonomyResponse
	3,  // 13: services.TaxonomyService.Update:output_type -> services.TaxonomyResponse
	3,  // 14: services.TaxonomyService.Get:output_type -> services.TaxonomyResponse
	3,  // 15: services.TaxonomyService.Delete:output_type -> services.TaxonomyResponse
	3,  // 16: services.TaxonomyService.Search:output_type -> services.TaxonomyResponse
	3,  // 17: services.TaxonomyService.List:output_type -> services.TaxonomyResponse
	3,  // 18: services.TaxonomyService.Tree:output_type -> services.TaxonomyResponse
	12, // [12:19] is the sub-list for method output_type
	5,  // [5:12] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_taxonomyService_proto_init() }
func file_taxonomyService_proto_init() {
	if File_taxonomyService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_taxonomyService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaxonomyRequest); i {
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
		file_taxonomyService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewTaxonomyIndex); i {
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
		file_taxonomyService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Taxonomy); i {
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
		file_taxonomyService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaxonomyResponse); i {
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
			RawDescriptor: file_taxonomyService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_taxonomyService_proto_goTypes,
		DependencyIndexes: file_taxonomyService_proto_depIdxs,
		MessageInfos:      file_taxonomyService_proto_msgTypes,
	}.Build()
	File_taxonomyService_proto = out.File
	file_taxonomyService_proto_rawDesc = nil
	file_taxonomyService_proto_goTypes = nil
	file_taxonomyService_proto_depIdxs = nil
}
