// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: streetService.proto

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

type StreetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int64   `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize int64   `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Sorting  string  `protobuf:"bytes,3,opt,name=sorting,proto3" json:"sorting"`
	Keywords string  `protobuf:"bytes,4,opt,name=keywords,proto3" json:"keywords"`
	Id       int64   `protobuf:"varint,5,opt,name=id,proto3" json:"id"`
	ParentId int64   `protobuf:"varint,6,opt,name=parent_id,json=parentId,proto3" json:"parent_id"`
	AreaId   int64   `protobuf:"varint,7,opt,name=area_id,json=areaId,proto3" json:"area_id"`
	Name     bool    `protobuf:"varint,8,opt,name=name,proto3" json:"name"`
	Ids      []int64 `protobuf:"varint,9,rep,packed,name=ids,proto3" json:"ids"`
}

func (x *StreetRequest) Reset() {
	*x = StreetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_streetService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreetRequest) ProtoMessage() {}

func (x *StreetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_streetService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreetRequest.ProtoReflect.Descriptor instead.
func (*StreetRequest) Descriptor() ([]byte, []int) {
	return file_streetService_proto_rawDescGZIP(), []int{0}
}

func (x *StreetRequest) GetPaged() int64 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *StreetRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *StreetRequest) GetSorting() string {
	if x != nil {
		return x.Sorting
	}
	return ""
}

func (x *StreetRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *StreetRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *StreetRequest) GetParentId() int64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *StreetRequest) GetAreaId() int64 {
	if x != nil {
		return x.AreaId
	}
	return 0
}

func (x *StreetRequest) GetName() bool {
	if x != nil {
		return x.Name
	}
	return false
}

func (x *StreetRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type NewStreetIndex struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaxId    int64 `protobuf:"varint,1,opt,name=max_id,json=maxId,proto3" json:"max_id"`
	NewId    int64 `protobuf:"varint,2,opt,name=new_id,json=newId,proto3" json:"new_id"`
	NewIndex int64 `protobuf:"varint,3,opt,name=new_index,json=newIndex,proto3" json:"new_index"`
}

func (x *NewStreetIndex) Reset() {
	*x = NewStreetIndex{}
	if protoimpl.UnsafeEnabled {
		mi := &file_streetService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewStreetIndex) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewStreetIndex) ProtoMessage() {}

func (x *NewStreetIndex) ProtoReflect() protoreflect.Message {
	mi := &file_streetService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewStreetIndex.ProtoReflect.Descriptor instead.
func (*NewStreetIndex) Descriptor() ([]byte, []int) {
	return file_streetService_proto_rawDescGZIP(), []int{1}
}

func (x *NewStreetIndex) GetMaxId() int64 {
	if x != nil {
		return x.MaxId
	}
	return 0
}

func (x *NewStreetIndex) GetNewId() int64 {
	if x != nil {
		return x.NewId
	}
	return 0
}

func (x *NewStreetIndex) GetNewIndex() int64 {
	if x != nil {
		return x.NewIndex
	}
	return 0
}

type Street struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64     `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Name      string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	AreaId    int64     `protobuf:"varint,3,opt,name=area_id,json=areaId,proto3" json:"area_id"`
	ParentId  int64     `protobuf:"varint,4,opt,name=parent_id,json=parentId,proto3" json:"parent_id"`
	Path      string    `protobuf:"bytes,5,opt,name=path,proto3" json:"path"`
	Depth     int32     `protobuf:"varint,6,opt,name=depth,proto3" json:"depth"`
	Lng       string    `protobuf:"bytes,7,opt,name=lng,proto3" json:"lng"`
	Lat       string    `protobuf:"bytes,8,opt,name=lat,proto3" json:"lat"`
	Memo      string    `protobuf:"bytes,9,opt,name=memo,proto3" json:"memo"`
	CreatedAt string    `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt string    `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	Area      *AreaInfo `protobuf:"bytes,12,opt,name=area,proto3" json:"area"`
	Parent    *Street   `protobuf:"bytes,13,opt,name=parent,proto3" json:"parent"`
	Children  []*Street `protobuf:"bytes,14,rep,name=children,proto3" json:"children"`
}

func (x *Street) Reset() {
	*x = Street{}
	if protoimpl.UnsafeEnabled {
		mi := &file_streetService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Street) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Street) ProtoMessage() {}

func (x *Street) ProtoReflect() protoreflect.Message {
	mi := &file_streetService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Street.ProtoReflect.Descriptor instead.
func (*Street) Descriptor() ([]byte, []int) {
	return file_streetService_proto_rawDescGZIP(), []int{2}
}

func (x *Street) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Street) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Street) GetAreaId() int64 {
	if x != nil {
		return x.AreaId
	}
	return 0
}

func (x *Street) GetParentId() int64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *Street) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Street) GetDepth() int32 {
	if x != nil {
		return x.Depth
	}
	return 0
}

func (x *Street) GetLng() string {
	if x != nil {
		return x.Lng
	}
	return ""
}

func (x *Street) GetLat() string {
	if x != nil {
		return x.Lat
	}
	return ""
}

func (x *Street) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *Street) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Street) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Street) GetArea() *AreaInfo {
	if x != nil {
		return x.Area
	}
	return nil
}

func (x *Street) GetParent() *Street {
	if x != nil {
		return x.Parent
	}
	return nil
}

func (x *Street) GetChildren() []*Street {
	if x != nil {
		return x.Children
	}
	return nil
}

type StreetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity   *Street       `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager    *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items    []*Street     `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Info     string        `protobuf:"bytes,4,opt,name=info,proto3" json:"info"`
	NewIndex int32         `protobuf:"varint,5,opt,name=new_index,json=newIndex,proto3" json:"new_index"`
}

func (x *StreetResponse) Reset() {
	*x = StreetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_streetService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreetResponse) ProtoMessage() {}

func (x *StreetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_streetService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreetResponse.ProtoReflect.Descriptor instead.
func (*StreetResponse) Descriptor() ([]byte, []int) {
	return file_streetService_proto_rawDescGZIP(), []int{3}
}

func (x *StreetResponse) GetEntity() *Street {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *StreetResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *StreetResponse) GetItems() []*Street {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *StreetResponse) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

func (x *StreetResponse) GetNewIndex() int32 {
	if x != nil {
		return x.NewIndex
	}
	return 0
}

var File_streetService_proto protoreflect.FileDescriptor

var file_streetService_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a,
	0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x15, 0x62, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe4, 0x01, 0x0a, 0x0d, 0x53, 0x74,
	0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77,
	0x6f, 0x72, 0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77,
	0x6f, 0x72, 0x64, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x61, 0x72, 0x65, 0x61, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73,
	0x22, 0x5b, 0x0a, 0x0e, 0x4e, 0x65, 0x77, 0x53, 0x74, 0x72, 0x65, 0x65, 0x74, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x15, 0x0a, 0x06, 0x6d, 0x61, 0x78, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x6d, 0x61, 0x78, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x6e, 0x65, 0x77,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6e, 0x65, 0x77, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x6e, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x82, 0x03,
	0x0a, 0x06, 0x53, 0x74, 0x72, 0x65, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07,
	0x61, 0x72, 0x65, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61,
	0x72, 0x65, 0x61, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x70, 0x74, 0x68, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x64, 0x65, 0x70, 0x74, 0x68, 0x12, 0x10, 0x0a, 0x03,
	0x6c, 0x6e, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x6e, 0x67, 0x12, 0x10,
	0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x61, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6d, 0x65, 0x6d, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x26, 0x0a, 0x04, 0x61, 0x72, 0x65, 0x61, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x72, 0x65, 0x61,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x61, 0x72, 0x65, 0x61, 0x12, 0x28, 0x0a, 0x06, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x65, 0x74, 0x52, 0x06, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e,
	0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x65, 0x74, 0x52, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72,
	0x65, 0x6e, 0x22, 0xb8, 0x01, 0x0a, 0x0e, 0x53, 0x74, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x53, 0x74, 0x72, 0x65, 0x65, 0x74, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12,
	0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70,
	0x61, 0x67, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53,
	0x74, 0x72, 0x65, 0x65, 0x74, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f,
	0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x6e, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x32, 0xa5, 0x03,
	0x0a, 0x0d, 0x53, 0x74, 0x72, 0x65, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x36, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x65, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x72,
	0x65, 0x65, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53,
	0x74, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x33, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x65, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x10,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x65, 0x74,
	0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x72, 0x65,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x06,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x04, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53,
	0x74, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x04, 0x54, 0x72, 0x65, 0x65,
	0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x72, 0x65,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_streetService_proto_rawDescOnce sync.Once
	file_streetService_proto_rawDescData = file_streetService_proto_rawDesc
)

func file_streetService_proto_rawDescGZIP() []byte {
	file_streetService_proto_rawDescOnce.Do(func() {
		file_streetService_proto_rawDescData = protoimpl.X.CompressGZIP(file_streetService_proto_rawDescData)
	})
	return file_streetService_proto_rawDescData
}

var file_streetService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_streetService_proto_goTypes = []interface{}{
	(*StreetRequest)(nil),  // 0: services.StreetRequest
	(*NewStreetIndex)(nil), // 1: services.NewStreetIndex
	(*Street)(nil),         // 2: services.Street
	(*StreetResponse)(nil), // 3: services.StreetResponse
	(*AreaInfo)(nil),       // 4: services.AreaInfo
	(*common.Pager)(nil),   // 5: common.Pager
}
var file_streetService_proto_depIdxs = []int32{
	4,  // 0: services.Street.area:type_name -> services.AreaInfo
	2,  // 1: services.Street.parent:type_name -> services.Street
	2,  // 2: services.Street.children:type_name -> services.Street
	2,  // 3: services.StreetResponse.entity:type_name -> services.Street
	5,  // 4: services.StreetResponse.pager:type_name -> common.Pager
	2,  // 5: services.StreetResponse.items:type_name -> services.Street
	2,  // 6: services.StreetService.Create:input_type -> services.Street
	2,  // 7: services.StreetService.Update:input_type -> services.Street
	2,  // 8: services.StreetService.Get:input_type -> services.Street
	2,  // 9: services.StreetService.Delete:input_type -> services.Street
	0,  // 10: services.StreetService.Search:input_type -> services.StreetRequest
	0,  // 11: services.StreetService.List:input_type -> services.StreetRequest
	0,  // 12: services.StreetService.Tree:input_type -> services.StreetRequest
	3,  // 13: services.StreetService.Create:output_type -> services.StreetResponse
	3,  // 14: services.StreetService.Update:output_type -> services.StreetResponse
	3,  // 15: services.StreetService.Get:output_type -> services.StreetResponse
	3,  // 16: services.StreetService.Delete:output_type -> services.StreetResponse
	3,  // 17: services.StreetService.Search:output_type -> services.StreetResponse
	3,  // 18: services.StreetService.List:output_type -> services.StreetResponse
	3,  // 19: services.StreetService.Tree:output_type -> services.StreetResponse
	13, // [13:20] is the sub-list for method output_type
	6,  // [6:13] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_streetService_proto_init() }
func file_streetService_proto_init() {
	if File_streetService_proto != nil {
		return
	}
	file_baseInfoService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_streetService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreetRequest); i {
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
		file_streetService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewStreetIndex); i {
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
		file_streetService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Street); i {
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
		file_streetService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreetResponse); i {
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
			RawDescriptor: file_streetService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_streetService_proto_goTypes,
		DependencyIndexes: file_streetService_proto_depIdxs,
		MessageInfos:      file_streetService_proto_msgTypes,
	}.Build()
	File_streetService_proto = out.File
	file_streetService_proto_rawDesc = nil
	file_streetService_proto_goTypes = nil
	file_streetService_proto_depIdxs = nil
}
