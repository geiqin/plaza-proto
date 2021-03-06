// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: appNavPersonalService.proto

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

type AppNavPersonal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	StoreId   int64   `protobuf:"varint,2,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`
	UserId    int64   `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	AppNavId  int32   `protobuf:"varint,4,opt,name=app_nav_id,json=appNavId,proto3" json:"app_nav_id,omitempty"`
	Sorting   int32   `protobuf:"varint,5,opt,name=sorting,proto3" json:"sorting,omitempty"`
	CreatedAt string  `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string  `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Nav       *AppNav `protobuf:"bytes,8,opt,name=nav,proto3" json:"nav,omitempty"`
}

func (x *AppNavPersonal) Reset() {
	*x = AppNavPersonal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appNavPersonalService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppNavPersonal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppNavPersonal) ProtoMessage() {}

func (x *AppNavPersonal) ProtoReflect() protoreflect.Message {
	mi := &file_appNavPersonalService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppNavPersonal.ProtoReflect.Descriptor instead.
func (*AppNavPersonal) Descriptor() ([]byte, []int) {
	return file_appNavPersonalService_proto_rawDescGZIP(), []int{0}
}

func (x *AppNavPersonal) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AppNavPersonal) GetStoreId() int64 {
	if x != nil {
		return x.StoreId
	}
	return 0
}

func (x *AppNavPersonal) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AppNavPersonal) GetAppNavId() int32 {
	if x != nil {
		return x.AppNavId
	}
	return 0
}

func (x *AppNavPersonal) GetSorting() int32 {
	if x != nil {
		return x.Sorting
	}
	return 0
}

func (x *AppNavPersonal) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *AppNavPersonal) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *AppNavPersonal) GetNav() *AppNav {
	if x != nil {
		return x.Nav
	}
	return nil
}

type AppNavPersonalWhere struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged     int32   `protobuf:"varint,1,opt,name=paged,proto3" json:"paged,omitempty"`
	PageSize  int32   `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Keywords  string  `protobuf:"bytes,3,opt,name=keywords,proto3" json:"keywords,omitempty"`
	Id        int64   `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
	Ids       []int64 `protobuf:"varint,5,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	StoreId   int64   `protobuf:"varint,6,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`
	UserId    int64   `protobuf:"varint,7,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	AppNavId  int32   `protobuf:"varint,8,opt,name=app_nav_id,json=appNavId,proto3" json:"app_nav_id,omitempty"`
	AppNavIds []int32 `protobuf:"varint,9,rep,packed,name=app_nav_ids,json=appNavIds,proto3" json:"app_nav_ids,omitempty"`
}

func (x *AppNavPersonalWhere) Reset() {
	*x = AppNavPersonalWhere{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appNavPersonalService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppNavPersonalWhere) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppNavPersonalWhere) ProtoMessage() {}

func (x *AppNavPersonalWhere) ProtoReflect() protoreflect.Message {
	mi := &file_appNavPersonalService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppNavPersonalWhere.ProtoReflect.Descriptor instead.
func (*AppNavPersonalWhere) Descriptor() ([]byte, []int) {
	return file_appNavPersonalService_proto_rawDescGZIP(), []int{1}
}

func (x *AppNavPersonalWhere) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *AppNavPersonalWhere) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *AppNavPersonalWhere) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *AppNavPersonalWhere) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AppNavPersonalWhere) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *AppNavPersonalWhere) GetStoreId() int64 {
	if x != nil {
		return x.StoreId
	}
	return 0
}

func (x *AppNavPersonalWhere) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AppNavPersonalWhere) GetAppNavId() int32 {
	if x != nil {
		return x.AppNavId
	}
	return 0
}

func (x *AppNavPersonalWhere) GetAppNavIds() []int32 {
	if x != nil {
		return x.AppNavIds
	}
	return nil
}

type AppNavPersonalResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *AppNavPersonal   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager  *common.Pager     `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items  []*AppNavPersonal `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error  *common.Error     `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info   *common.Info      `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *AppNavPersonalResponse) Reset() {
	*x = AppNavPersonalResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appNavPersonalService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppNavPersonalResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppNavPersonalResponse) ProtoMessage() {}

func (x *AppNavPersonalResponse) ProtoReflect() protoreflect.Message {
	mi := &file_appNavPersonalService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppNavPersonalResponse.ProtoReflect.Descriptor instead.
func (*AppNavPersonalResponse) Descriptor() ([]byte, []int) {
	return file_appNavPersonalService_proto_rawDescGZIP(), []int{2}
}

func (x *AppNavPersonalResponse) GetEntity() *AppNavPersonal {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *AppNavPersonalResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *AppNavPersonalResponse) GetItems() []*AppNavPersonal {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *AppNavPersonalResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *AppNavPersonalResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_appNavPersonalService_proto protoreflect.FileDescriptor

var file_appNavPersonalService_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x70, 0x4e, 0x61, 0x76, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x61, 0x70, 0x70, 0x4e,
	0x61, 0x76, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xee, 0x01, 0x0a, 0x0e, 0x41, 0x70, 0x70, 0x4e, 0x61, 0x76, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x61, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x0a, 0x61, 0x70, 0x70, 0x5f, 0x6e, 0x61,
	0x76, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x61, 0x70, 0x70, 0x4e,
	0x61, 0x76, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x22, 0x0a, 0x03,
	0x6e, 0x61, 0x76, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x4e, 0x61, 0x76, 0x52, 0x03, 0x6e, 0x61, 0x76,
	0x22, 0xf8, 0x01, 0x0a, 0x13, 0x41, 0x70, 0x70, 0x4e, 0x61, 0x76, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x61, 0x6c, 0x57, 0x68, 0x65, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6b,
	0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b,
	0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a,
	0x0a, 0x61, 0x70, 0x70, 0x5f, 0x6e, 0x61, 0x76, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x61, 0x70, 0x70, 0x4e, 0x61, 0x76, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0b, 0x61,
	0x70, 0x70, 0x5f, 0x6e, 0x61, 0x76, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x05,
	0x52, 0x09, 0x61, 0x70, 0x70, 0x4e, 0x61, 0x76, 0x49, 0x64, 0x73, 0x22, 0xe6, 0x01, 0x0a, 0x16,
	0x41, 0x70, 0x70, 0x4e, 0x61, 0x76, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x41, 0x70, 0x70, 0x4e, 0x61, 0x76, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c,
	0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x2e, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x4e, 0x61, 0x76, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x23, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04,
	0x69, 0x6e, 0x66, 0x6f, 0x32, 0xad, 0x01, 0x0a, 0x15, 0x41, 0x70, 0x70, 0x4e, 0x61, 0x76, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x49,
	0x0a, 0x04, 0x53, 0x61, 0x76, 0x65, 0x12, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x41, 0x70, 0x70, 0x4e, 0x61, 0x76, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c,
	0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x41, 0x70, 0x70, 0x4e, 0x61, 0x76, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x04, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x70, 0x70,
	0x4e, 0x61, 0x76, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x57, 0x68, 0x65, 0x72, 0x65,
	0x1a, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x4e,
	0x61, 0x76, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_appNavPersonalService_proto_rawDescOnce sync.Once
	file_appNavPersonalService_proto_rawDescData = file_appNavPersonalService_proto_rawDesc
)

func file_appNavPersonalService_proto_rawDescGZIP() []byte {
	file_appNavPersonalService_proto_rawDescOnce.Do(func() {
		file_appNavPersonalService_proto_rawDescData = protoimpl.X.CompressGZIP(file_appNavPersonalService_proto_rawDescData)
	})
	return file_appNavPersonalService_proto_rawDescData
}

var file_appNavPersonalService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_appNavPersonalService_proto_goTypes = []interface{}{
	(*AppNavPersonal)(nil),         // 0: services.AppNavPersonal
	(*AppNavPersonalWhere)(nil),    // 1: services.AppNavPersonalWhere
	(*AppNavPersonalResponse)(nil), // 2: services.AppNavPersonalResponse
	(*AppNav)(nil),                 // 3: services.AppNav
	(*common.Pager)(nil),           // 4: common.Pager
	(*common.Error)(nil),           // 5: common.Error
	(*common.Info)(nil),            // 6: common.Info
}
var file_appNavPersonalService_proto_depIdxs = []int32{
	3, // 0: services.AppNavPersonal.nav:type_name -> services.AppNav
	0, // 1: services.AppNavPersonalResponse.entity:type_name -> services.AppNavPersonal
	4, // 2: services.AppNavPersonalResponse.pager:type_name -> common.Pager
	0, // 3: services.AppNavPersonalResponse.items:type_name -> services.AppNavPersonal
	5, // 4: services.AppNavPersonalResponse.error:type_name -> common.Error
	6, // 5: services.AppNavPersonalResponse.info:type_name -> common.Info
	1, // 6: services.AppNavPersonalService.Save:input_type -> services.AppNavPersonalWhere
	1, // 7: services.AppNavPersonalService.List:input_type -> services.AppNavPersonalWhere
	2, // 8: services.AppNavPersonalService.Save:output_type -> services.AppNavPersonalResponse
	2, // 9: services.AppNavPersonalService.List:output_type -> services.AppNavPersonalResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_appNavPersonalService_proto_init() }
func file_appNavPersonalService_proto_init() {
	if File_appNavPersonalService_proto != nil {
		return
	}
	file_appNavService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_appNavPersonalService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppNavPersonal); i {
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
		file_appNavPersonalService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppNavPersonalWhere); i {
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
		file_appNavPersonalService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppNavPersonalResponse); i {
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
			RawDescriptor: file_appNavPersonalService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_appNavPersonalService_proto_goTypes,
		DependencyIndexes: file_appNavPersonalService_proto_depIdxs,
		MessageInfos:      file_appNavPersonalService_proto_msgTypes,
	}.Build()
	File_appNavPersonalService_proto = out.File
	file_appNavPersonalService_proto_rawDesc = nil
	file_appNavPersonalService_proto_goTypes = nil
	file_appNavPersonalService_proto_depIdxs = nil
}