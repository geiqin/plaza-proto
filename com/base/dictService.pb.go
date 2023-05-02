// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: dictService.proto

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

type DictRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged      int64   `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize   int64   `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Id         int64   `protobuf:"varint,3,opt,name=id,proto3" json:"id"`
	Keywords   string  `protobuf:"bytes,5,opt,name=keywords,proto3" json:"keywords"`
	DictTypeId int32   `protobuf:"varint,6,opt,name=dict_type_id,json=dictTypeId,proto3" json:"dict_type_id"`
	Status     int32   `protobuf:"varint,7,opt,name=status,proto3" json:"status"`
	Items      []*Dict `protobuf:"bytes,8,rep,name=items,proto3" json:"items"`
	Ids        []int64 `protobuf:"varint,4,rep,packed,name=ids,proto3" json:"ids"`
}

func (x *DictRequest) Reset() {
	*x = DictRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dictService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DictRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DictRequest) ProtoMessage() {}

func (x *DictRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dictService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DictRequest.ProtoReflect.Descriptor instead.
func (*DictRequest) Descriptor() ([]byte, []int) {
	return file_dictService_proto_rawDescGZIP(), []int{0}
}

func (x *DictRequest) GetPaged() int64 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *DictRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *DictRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DictRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *DictRequest) GetDictTypeId() int32 {
	if x != nil {
		return x.DictTypeId
	}
	return 0
}

func (x *DictRequest) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *DictRequest) GetItems() []*Dict {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *DictRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

// 字典信息
type Dict struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	DictTypeId int32  `protobuf:"varint,2,opt,name=dict_type_id,json=dictTypeId,proto3" json:"dict_type_id"`
	Value      string `protobuf:"bytes,3,opt,name=value,proto3" json:"value"`
	Text       string `protobuf:"bytes,4,opt,name=text,proto3" json:"text"`
	Term       string `protobuf:"bytes,5,opt,name=term,proto3" json:"term"`
	Memo       string `protobuf:"bytes,6,opt,name=memo,proto3" json:"memo"`
	Sort       int32  `protobuf:"varint,7,opt,name=sort,proto3" json:"sort"`
	Status     int32  `protobuf:"varint,8,opt,name=status,proto3" json:"status"`
	CreatedAt  string `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt  string `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
}

func (x *Dict) Reset() {
	*x = Dict{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dictService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dict) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dict) ProtoMessage() {}

func (x *Dict) ProtoReflect() protoreflect.Message {
	mi := &file_dictService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dict.ProtoReflect.Descriptor instead.
func (*Dict) Descriptor() ([]byte, []int) {
	return file_dictService_proto_rawDescGZIP(), []int{1}
}

func (x *Dict) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Dict) GetDictTypeId() int32 {
	if x != nil {
		return x.DictTypeId
	}
	return 0
}

func (x *Dict) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Dict) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Dict) GetTerm() string {
	if x != nil {
		return x.Term
	}
	return ""
}

func (x *Dict) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *Dict) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *Dict) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Dict) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Dict) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

//
type DictResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Dict            `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager    `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*Dict          `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Error  *common.Error    `protobuf:"bytes,4,opt,name=error,proto3" json:"error"`
	Info   *common.Info     `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
	Maps   map[string]*Dict `protobuf:"bytes,6,rep,name=maps,proto3" json:"maps" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *DictResponse) Reset() {
	*x = DictResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dictService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DictResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DictResponse) ProtoMessage() {}

func (x *DictResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dictService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DictResponse.ProtoReflect.Descriptor instead.
func (*DictResponse) Descriptor() ([]byte, []int) {
	return file_dictService_proto_rawDescGZIP(), []int{2}
}

func (x *DictResponse) GetEntity() *Dict {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *DictResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *DictResponse) GetItems() []*Dict {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *DictResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *DictResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *DictResponse) GetMaps() map[string]*Dict {
	if x != nil {
		return x.Maps
	}
	return nil
}

var File_dictService_proto protoreflect.FileDescriptor

var file_dictService_proto_rawDesc = []byte{
	0x0a, 0x11, 0x64, 0x69, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xde, 0x01, 0x0a, 0x0b, 0x44, 0x69, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12,
	0x20, 0x0a, 0x0c, 0x64, 0x69, 0x63, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x64, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x24, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x44, 0x69, 0x63, 0x74, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12,
	0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64,
	0x73, 0x22, 0xf4, 0x01, 0x0a, 0x04, 0x44, 0x69, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0c, 0x64, 0x69,
	0x63, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x64, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65,
	0x6d, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x6f,
	0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xc7, 0x02, 0x0a, 0x0c, 0x44, 0x69, 0x63,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x69, 0x63, 0x74, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52,
	0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x44, 0x69, 0x63, 0x74, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x23, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69,
	0x6e, 0x66, 0x6f, 0x12, 0x34, 0x0a, 0x04, 0x6d, 0x61, 0x70, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x69, 0x63,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4d, 0x61, 0x70, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x04, 0x6d, 0x61, 0x70, 0x73, 0x1a, 0x47, 0x0a, 0x09, 0x4d, 0x61, 0x70,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x24, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x44, 0x69, 0x63, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x32, 0xd6, 0x03, 0x0a, 0x0b, 0x44, 0x69, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x69, 0x63, 0x74, 0x1a, 0x16, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x69, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x69, 0x63, 0x74,
	0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x69, 0x63, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x06, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x44, 0x69, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x69, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x15, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x69, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44,
	0x69, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a,
	0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x44, 0x69, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x69, 0x63, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x69, 0x63, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x44, 0x69, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x3a, 0x0a, 0x07, 0x53, 0x65, 0x74, 0x53, 0x6f, 0x72, 0x74, 0x12, 0x15, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x69, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44,
	0x69, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a,
	0x09, 0x53, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x69, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x69, 0x63,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2f,
	0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_dictService_proto_rawDescOnce sync.Once
	file_dictService_proto_rawDescData = file_dictService_proto_rawDesc
)

func file_dictService_proto_rawDescGZIP() []byte {
	file_dictService_proto_rawDescOnce.Do(func() {
		file_dictService_proto_rawDescData = protoimpl.X.CompressGZIP(file_dictService_proto_rawDescData)
	})
	return file_dictService_proto_rawDescData
}

var file_dictService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_dictService_proto_goTypes = []interface{}{
	(*DictRequest)(nil),  // 0: services.DictRequest
	(*Dict)(nil),         // 1: services.Dict
	(*DictResponse)(nil), // 2: services.DictResponse
	nil,                  // 3: services.DictResponse.MapsEntry
	(*common.Pager)(nil), // 4: common.Pager
	(*common.Error)(nil), // 5: common.Error
	(*common.Info)(nil),  // 6: common.Info
}
var file_dictService_proto_depIdxs = []int32{
	1,  // 0: services.DictRequest.items:type_name -> services.Dict
	1,  // 1: services.DictResponse.entity:type_name -> services.Dict
	4,  // 2: services.DictResponse.pager:type_name -> common.Pager
	1,  // 3: services.DictResponse.items:type_name -> services.Dict
	5,  // 4: services.DictResponse.error:type_name -> common.Error
	6,  // 5: services.DictResponse.info:type_name -> common.Info
	3,  // 6: services.DictResponse.maps:type_name -> services.DictResponse.MapsEntry
	1,  // 7: services.DictResponse.MapsEntry.value:type_name -> services.Dict
	1,  // 8: services.DictService.Create:input_type -> services.Dict
	1,  // 9: services.DictService.Update:input_type -> services.Dict
	0,  // 10: services.DictService.Delete:input_type -> services.DictRequest
	0,  // 11: services.DictService.Get:input_type -> services.DictRequest
	0,  // 12: services.DictService.Search:input_type -> services.DictRequest
	0,  // 13: services.DictService.List:input_type -> services.DictRequest
	0,  // 14: services.DictService.SetSort:input_type -> services.DictRequest
	0,  // 15: services.DictService.SetStatus:input_type -> services.DictRequest
	2,  // 16: services.DictService.Create:output_type -> services.DictResponse
	2,  // 17: services.DictService.Update:output_type -> services.DictResponse
	2,  // 18: services.DictService.Delete:output_type -> services.DictResponse
	2,  // 19: services.DictService.Get:output_type -> services.DictResponse
	2,  // 20: services.DictService.Search:output_type -> services.DictResponse
	2,  // 21: services.DictService.List:output_type -> services.DictResponse
	2,  // 22: services.DictService.SetSort:output_type -> services.DictResponse
	2,  // 23: services.DictService.SetStatus:output_type -> services.DictResponse
	16, // [16:24] is the sub-list for method output_type
	8,  // [8:16] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_dictService_proto_init() }
func file_dictService_proto_init() {
	if File_dictService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dictService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DictRequest); i {
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
		file_dictService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dict); i {
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
		file_dictService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DictResponse); i {
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
			RawDescriptor: file_dictService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dictService_proto_goTypes,
		DependencyIndexes: file_dictService_proto_depIdxs,
		MessageInfos:      file_dictService_proto_msgTypes,
	}.Build()
	File_dictService_proto = out.File
	file_dictService_proto_rawDesc = nil
	file_dictService_proto_goTypes = nil
	file_dictService_proto_depIdxs = nil
}
