// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: navTypeService.proto

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

// 导航
type NavType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Name      string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`
	Title     string   `protobuf:"bytes,4,opt,name=title,proto3" json:"title"`
	VersionId int32    `protobuf:"varint,5,opt,name=version_id,json=versionId,proto3" json:"version_id"`
	Tag       string   `protobuf:"bytes,6,opt,name=tag,proto3" json:"tag"`
	Ver       int32    `protobuf:"varint,7,opt,name=ver,proto3" json:"ver"`
	Memo      string   `protobuf:"bytes,8,opt,name=memo,proto3" json:"memo"`
	Status    string   `protobuf:"bytes,9,opt,name=status,proto3" json:"status"`
	CreatedAt string   `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt string   `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	Navs      []*Nav   `protobuf:"bytes,12,rep,name=navs,proto3" json:"navs"`
	Version   *Version `protobuf:"bytes,13,opt,name=version,proto3" json:"version"`
}

func (x *NavType) Reset() {
	*x = NavType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_navTypeService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NavType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NavType) ProtoMessage() {}

func (x *NavType) ProtoReflect() protoreflect.Message {
	mi := &file_navTypeService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NavType.ProtoReflect.Descriptor instead.
func (*NavType) Descriptor() ([]byte, []int) {
	return file_navTypeService_proto_rawDescGZIP(), []int{0}
}

func (x *NavType) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *NavType) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NavType) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *NavType) GetVersionId() int32 {
	if x != nil {
		return x.VersionId
	}
	return 0
}

func (x *NavType) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *NavType) GetVer() int32 {
	if x != nil {
		return x.Ver
	}
	return 0
}

func (x *NavType) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *NavType) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *NavType) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *NavType) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *NavType) GetNavs() []*Nav {
	if x != nil {
		return x.Navs
	}
	return nil
}

func (x *NavType) GetVersion() *Version {
	if x != nil {
		return x.Version
	}
	return nil
}

type NavTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int64  `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize int64  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Sorting  string `protobuf:"bytes,3,opt,name=sorting,proto3" json:"sorting"`
	//base params
	Id        int64   `protobuf:"varint,4,opt,name=id,proto3" json:"id"`
	Name      string  `protobuf:"bytes,5,opt,name=name,proto3" json:"name"`
	Title     string  `protobuf:"bytes,6,opt,name=title,proto3" json:"title"`
	Status    string  `protobuf:"bytes,7,opt,name=status,proto3" json:"status"`
	VersionId int32   `protobuf:"varint,8,opt,name=version_id,json=versionId,proto3" json:"version_id"`
	StoreId   int64   `protobuf:"varint,9,opt,name=store_id,json=storeId,proto3" json:"store_id"`
	Tag       string  `protobuf:"bytes,10,opt,name=tag,proto3" json:"tag"`
	Ver       int32   `protobuf:"varint,11,opt,name=ver,proto3" json:"ver"`
	Ids       []int64 `protobuf:"varint,12,rep,packed,name=ids,proto3" json:"ids"`
}

func (x *NavTypeRequest) Reset() {
	*x = NavTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_navTypeService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NavTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NavTypeRequest) ProtoMessage() {}

func (x *NavTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_navTypeService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NavTypeRequest.ProtoReflect.Descriptor instead.
func (*NavTypeRequest) Descriptor() ([]byte, []int) {
	return file_navTypeService_proto_rawDescGZIP(), []int{1}
}

func (x *NavTypeRequest) GetPaged() int64 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *NavTypeRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *NavTypeRequest) GetSorting() string {
	if x != nil {
		return x.Sorting
	}
	return ""
}

func (x *NavTypeRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *NavTypeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NavTypeRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *NavTypeRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *NavTypeRequest) GetVersionId() int32 {
	if x != nil {
		return x.VersionId
	}
	return 0
}

func (x *NavTypeRequest) GetStoreId() int64 {
	if x != nil {
		return x.StoreId
	}
	return 0
}

func (x *NavTypeRequest) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *NavTypeRequest) GetVer() int32 {
	if x != nil {
		return x.Ver
	}
	return 0
}

func (x *NavTypeRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type NavTypeData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *NavType      `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*NavType    `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Info   *common.Info  `protobuf:"bytes,4,opt,name=info,proto3" json:"info"`
}

func (x *NavTypeData) Reset() {
	*x = NavTypeData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_navTypeService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NavTypeData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NavTypeData) ProtoMessage() {}

func (x *NavTypeData) ProtoReflect() protoreflect.Message {
	mi := &file_navTypeService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NavTypeData.ProtoReflect.Descriptor instead.
func (*NavTypeData) Descriptor() ([]byte, []int) {
	return file_navTypeService_proto_rawDescGZIP(), []int{2}
}

func (x *NavTypeData) GetEntity() *NavType {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *NavTypeData) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *NavTypeData) GetItems() []*NavType {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *NavTypeData) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type NavTypeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *NavTypeData  `protobuf:"bytes,1,opt,name=data,proto3" json:"data"`
	Error *common.Error `protobuf:"bytes,2,opt,name=error,proto3" json:"error"`
}

func (x *NavTypeResponse) Reset() {
	*x = NavTypeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_navTypeService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NavTypeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NavTypeResponse) ProtoMessage() {}

func (x *NavTypeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_navTypeService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NavTypeResponse.ProtoReflect.Descriptor instead.
func (*NavTypeResponse) Descriptor() ([]byte, []int) {
	return file_navTypeService_proto_rawDescGZIP(), []int{3}
}

func (x *NavTypeResponse) GetData() *NavTypeData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *NavTypeResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_navTypeService_proto protoreflect.FileDescriptor

var file_navTypeService_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6e, 0x61, 0x76, 0x54, 0x79, 0x70, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x6e, 0x61, 0x76, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc0, 0x02, 0x0a, 0x07,
	0x4e, 0x61, 0x76, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74,
	0x61, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x76, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x21,
	0x0a, 0x04, 0x6e, 0x61, 0x76, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4e, 0x61, 0x76, 0x52, 0x04, 0x6e, 0x61, 0x76,
	0x73, 0x12, 0x2b, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x9f,
	0x02, 0x0a, 0x0e, 0x4e, 0x61, 0x76, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x19, 0x0a, 0x08, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61,
	0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x10, 0x0a, 0x03,
	0x76, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x76, 0x65, 0x72, 0x12, 0x10,
	0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73,
	0x22, 0xa8, 0x01, 0x0a, 0x0b, 0x4e, 0x61, 0x76, 0x54, 0x79, 0x70, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x29, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4e, 0x61, 0x76, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70,
	0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72,
	0x12, 0x27, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4e, 0x61, 0x76, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x61, 0x0a, 0x0f, 0x4e,
	0x61, 0x76, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4e, 0x61, 0x76, 0x54, 0x79, 0x70, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xf5,
	0x02, 0x0a, 0x0e, 0x4e, 0x61, 0x76, 0x54, 0x79, 0x70, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x35, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x4e, 0x61, 0x76, 0x54, 0x79, 0x70, 0x65, 0x1a, 0x19, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4e, 0x61, 0x76, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4e, 0x61,
	0x76, 0x54, 0x79, 0x70, 0x65, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x4e, 0x61, 0x76, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x38, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4e, 0x61, 0x76, 0x54, 0x79, 0x70, 0x65, 0x1a,
	0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4e, 0x61, 0x76, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x06,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x4e, 0x61, 0x76, 0x54, 0x79, 0x70, 0x65, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x4e, 0x61, 0x76, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4e, 0x61, 0x76, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4e, 0x61, 0x76, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4e, 0x61, 0x76, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x4e, 0x61, 0x76, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_navTypeService_proto_rawDescOnce sync.Once
	file_navTypeService_proto_rawDescData = file_navTypeService_proto_rawDesc
)

func file_navTypeService_proto_rawDescGZIP() []byte {
	file_navTypeService_proto_rawDescOnce.Do(func() {
		file_navTypeService_proto_rawDescData = protoimpl.X.CompressGZIP(file_navTypeService_proto_rawDescData)
	})
	return file_navTypeService_proto_rawDescData
}

var file_navTypeService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_navTypeService_proto_goTypes = []interface{}{
	(*NavType)(nil),         // 0: services.NavType
	(*NavTypeRequest)(nil),  // 1: services.NavTypeRequest
	(*NavTypeData)(nil),     // 2: services.NavTypeData
	(*NavTypeResponse)(nil), // 3: services.NavTypeResponse
	(*Nav)(nil),             // 4: services.Nav
	(*Version)(nil),         // 5: services.Version
	(*common.Pager)(nil),    // 6: common.Pager
	(*common.Info)(nil),     // 7: common.Info
	(*common.Error)(nil),    // 8: common.Error
}
var file_navTypeService_proto_depIdxs = []int32{
	4,  // 0: services.NavType.navs:type_name -> services.Nav
	5,  // 1: services.NavType.version:type_name -> services.Version
	0,  // 2: services.NavTypeData.entity:type_name -> services.NavType
	6,  // 3: services.NavTypeData.pager:type_name -> common.Pager
	0,  // 4: services.NavTypeData.items:type_name -> services.NavType
	7,  // 5: services.NavTypeData.info:type_name -> common.Info
	2,  // 6: services.NavTypeResponse.data:type_name -> services.NavTypeData
	8,  // 7: services.NavTypeResponse.error:type_name -> common.Error
	0,  // 8: services.NavTypeService.Get:input_type -> services.NavType
	0,  // 9: services.NavTypeService.Create:input_type -> services.NavType
	0,  // 10: services.NavTypeService.Update:input_type -> services.NavType
	0,  // 11: services.NavTypeService.Delete:input_type -> services.NavType
	1,  // 12: services.NavTypeService.Search:input_type -> services.NavTypeRequest
	1,  // 13: services.NavTypeService.List:input_type -> services.NavTypeRequest
	3,  // 14: services.NavTypeService.Get:output_type -> services.NavTypeResponse
	3,  // 15: services.NavTypeService.Create:output_type -> services.NavTypeResponse
	3,  // 16: services.NavTypeService.Update:output_type -> services.NavTypeResponse
	3,  // 17: services.NavTypeService.Delete:output_type -> services.NavTypeResponse
	3,  // 18: services.NavTypeService.Search:output_type -> services.NavTypeResponse
	3,  // 19: services.NavTypeService.List:output_type -> services.NavTypeResponse
	14, // [14:20] is the sub-list for method output_type
	8,  // [8:14] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_navTypeService_proto_init() }
func file_navTypeService_proto_init() {
	if File_navTypeService_proto != nil {
		return
	}
	file_navService_proto_init()
	file_versionService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_navTypeService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NavType); i {
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
		file_navTypeService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NavTypeRequest); i {
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
		file_navTypeService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NavTypeData); i {
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
		file_navTypeService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NavTypeResponse); i {
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
			RawDescriptor: file_navTypeService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_navTypeService_proto_goTypes,
		DependencyIndexes: file_navTypeService_proto_depIdxs,
		MessageInfos:      file_navTypeService_proto_msgTypes,
	}.Build()
	File_navTypeService_proto = out.File
	file_navTypeService_proto_rawDesc = nil
	file_navTypeService_proto_goTypes = nil
	file_navTypeService_proto_depIdxs = nil
}
