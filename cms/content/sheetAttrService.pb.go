// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: sheetAttrService.proto

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

type SheetAttrWhere struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int32 `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Id       int32 `protobuf:"varint,3,opt,name=id,proto3" json:"id"`
	// @inject_tag: gorm:"-"
	Ids      []int32 `protobuf:"varint,4,rep,packed,name=ids,proto3" json:"ids" gorm:"-"`
	Name     string  `protobuf:"bytes,5,opt,name=name,proto3" json:"name"`
	Keywords string  `protobuf:"bytes,6,opt,name=keywords,proto3" json:"keywords"`
	SheetId  int32   `protobuf:"varint,7,opt,name=sheet_id,json=sheetId,proto3" json:"sheet_id"`
}

func (x *SheetAttrWhere) Reset() {
	*x = SheetAttrWhere{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sheetAttrService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SheetAttrWhere) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SheetAttrWhere) ProtoMessage() {}

func (x *SheetAttrWhere) ProtoReflect() protoreflect.Message {
	mi := &file_sheetAttrService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SheetAttrWhere.ProtoReflect.Descriptor instead.
func (*SheetAttrWhere) Descriptor() ([]byte, []int) {
	return file_sheetAttrService_proto_rawDescGZIP(), []int{0}
}

func (x *SheetAttrWhere) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *SheetAttrWhere) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *SheetAttrWhere) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SheetAttrWhere) GetIds() []int32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *SheetAttrWhere) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SheetAttrWhere) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *SheetAttrWhere) GetSheetId() int32 {
	if x != nil {
		return x.SheetId
	}
	return 0
}

type SheetAttr struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	SheetId       int32  `protobuf:"varint,2,opt,name=sheet_id,json=sheetId,proto3" json:"sheet_id"`
	Name          string `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`
	DisplayName   string `protobuf:"bytes,4,opt,name=display_name,json=displayName,proto3" json:"display_name"`
	Type          string `protobuf:"bytes,5,opt,name=type,proto3" json:"type"`
	Required      bool   `protobuf:"varint,6,opt,name=required,proto3" json:"required"`
	Rule          string `protobuf:"bytes,7,opt,name=rule,proto3" json:"rule"`
	Tip           string `protobuf:"bytes,8,opt,name=tip,proto3" json:"tip"`
	Sorting       int32  `protobuf:"varint,9,opt,name=sorting,proto3" json:"sorting"`
	Data          string `protobuf:"bytes,10,opt,name=data,proto3" json:"data"`
	CreatedAt     string `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt     string `protobuf:"bytes,12,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	IsDisplayList bool   `protobuf:"varint,13,opt,name=is_display_list,json=isDisplayList,proto3" json:"is_display_list"`
	IsQuery       bool   `protobuf:"varint,14,opt,name=is_query,json=isQuery,proto3" json:"is_query"`
	QueryMethod   string `protobuf:"bytes,15,opt,name=query_method,json=queryMethod,proto3" json:"query_method"`
}

func (x *SheetAttr) Reset() {
	*x = SheetAttr{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sheetAttrService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SheetAttr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SheetAttr) ProtoMessage() {}

func (x *SheetAttr) ProtoReflect() protoreflect.Message {
	mi := &file_sheetAttrService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SheetAttr.ProtoReflect.Descriptor instead.
func (*SheetAttr) Descriptor() ([]byte, []int) {
	return file_sheetAttrService_proto_rawDescGZIP(), []int{1}
}

func (x *SheetAttr) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SheetAttr) GetSheetId() int32 {
	if x != nil {
		return x.SheetId
	}
	return 0
}

func (x *SheetAttr) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SheetAttr) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *SheetAttr) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *SheetAttr) GetRequired() bool {
	if x != nil {
		return x.Required
	}
	return false
}

func (x *SheetAttr) GetRule() string {
	if x != nil {
		return x.Rule
	}
	return ""
}

func (x *SheetAttr) GetTip() string {
	if x != nil {
		return x.Tip
	}
	return ""
}

func (x *SheetAttr) GetSorting() int32 {
	if x != nil {
		return x.Sorting
	}
	return 0
}

func (x *SheetAttr) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *SheetAttr) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *SheetAttr) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *SheetAttr) GetIsDisplayList() bool {
	if x != nil {
		return x.IsDisplayList
	}
	return false
}

func (x *SheetAttr) GetIsQuery() bool {
	if x != nil {
		return x.IsQuery
	}
	return false
}

func (x *SheetAttr) GetQueryMethod() string {
	if x != nil {
		return x.QueryMethod
	}
	return ""
}

type SheetAttrResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error  *common.Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error"`
	Info   *common.Info  `protobuf:"bytes,2,opt,name=info,proto3" json:"info"`
	Pager  *common.Pager `protobuf:"bytes,3,opt,name=pager,proto3" json:"pager"`
	Entity *SheetAttr    `protobuf:"bytes,4,opt,name=entity,proto3" json:"entity"`
	Items  []*SheetAttr  `protobuf:"bytes,5,rep,name=items,proto3" json:"items"`
}

func (x *SheetAttrResponse) Reset() {
	*x = SheetAttrResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sheetAttrService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SheetAttrResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SheetAttrResponse) ProtoMessage() {}

func (x *SheetAttrResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sheetAttrService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SheetAttrResponse.ProtoReflect.Descriptor instead.
func (*SheetAttrResponse) Descriptor() ([]byte, []int) {
	return file_sheetAttrService_proto_rawDescGZIP(), []int{2}
}

func (x *SheetAttrResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *SheetAttrResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *SheetAttrResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *SheetAttrResponse) GetEntity() *SheetAttr {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *SheetAttrResponse) GetItems() []*SheetAttr {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_sheetAttrService_proto protoreflect.FileDescriptor

var file_sheetAttrService_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x68, 0x65, 0x65, 0x74, 0x41, 0x74, 0x74, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb0, 0x01, 0x0a, 0x0e, 0x53, 0x68, 0x65, 0x65, 0x74, 0x41,
	0x74, 0x74, 0x72, 0x57, 0x68, 0x65, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x69,
	0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x05, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x19, 0x0a,
	0x08, 0x73, 0x68, 0x65, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x73, 0x68, 0x65, 0x65, 0x74, 0x49, 0x64, 0x22, 0x95, 0x03, 0x0a, 0x09, 0x53, 0x68, 0x65,
	0x65, 0x74, 0x41, 0x74, 0x74, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x68, 0x65, 0x65, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x68, 0x65, 0x65, 0x74, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x75, 0x6c, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x75, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x74, 0x69, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x69, 0x70, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x69, 0x73,
	0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0d, 0x69, 0x73, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x21, 0x0a,
	0x0c, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x71, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x22, 0xd7, 0x01, 0x0a, 0x11, 0x53, 0x68, 0x65, 0x65, 0x74, 0x41, 0x74, 0x74, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x23, 0x0a,
	0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67,
	0x65, 0x72, 0x12, 0x2b, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x68,
	0x65, 0x65, 0x74, 0x41, 0x74, 0x74, 0x72, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12,
	0x29, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x68, 0x65, 0x65, 0x74, 0x41,
	0x74, 0x74, 0x72, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x32, 0xd2, 0x03, 0x0a, 0x10, 0x53,
	0x68, 0x65, 0x65, 0x74, 0x41, 0x74, 0x74, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x3c, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x68, 0x65, 0x65, 0x74, 0x41, 0x74, 0x74, 0x72, 0x1a, 0x1b,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x68, 0x65, 0x65, 0x74, 0x41,
	0x74, 0x74, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a,
	0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x53, 0x68, 0x65, 0x65, 0x74, 0x41, 0x74, 0x74, 0x72, 0x1a, 0x1b, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x68, 0x65, 0x65, 0x74, 0x41, 0x74, 0x74,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x06, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x53, 0x68, 0x65, 0x65, 0x74, 0x41, 0x74, 0x74, 0x72, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a,
	0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x68, 0x65, 0x65, 0x74,
	0x41, 0x74, 0x74, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39,
	0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x53, 0x68, 0x65, 0x65, 0x74, 0x41, 0x74, 0x74, 0x72, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x68, 0x65, 0x65, 0x74, 0x41, 0x74, 0x74, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x06, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53,
	0x68, 0x65, 0x65, 0x74, 0x41, 0x74, 0x74, 0x72, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x1b, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x68, 0x65, 0x65, 0x74, 0x41, 0x74,
	0x74, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x04,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x53, 0x68, 0x65, 0x65, 0x74, 0x41, 0x74, 0x74, 0x72, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x1b,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x68, 0x65, 0x65, 0x74, 0x41,
	0x74, 0x74, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a,
	0x0a, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x6f, 0x72, 0x74, 0x12, 0x13, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x68, 0x65, 0x65, 0x74, 0x41, 0x74, 0x74, 0x72,
	0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x68, 0x65, 0x65,
	0x74, 0x41, 0x74, 0x74, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sheetAttrService_proto_rawDescOnce sync.Once
	file_sheetAttrService_proto_rawDescData = file_sheetAttrService_proto_rawDesc
)

func file_sheetAttrService_proto_rawDescGZIP() []byte {
	file_sheetAttrService_proto_rawDescOnce.Do(func() {
		file_sheetAttrService_proto_rawDescData = protoimpl.X.CompressGZIP(file_sheetAttrService_proto_rawDescData)
	})
	return file_sheetAttrService_proto_rawDescData
}

var file_sheetAttrService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_sheetAttrService_proto_goTypes = []interface{}{
	(*SheetAttrWhere)(nil),    // 0: services.SheetAttrWhere
	(*SheetAttr)(nil),         // 1: services.SheetAttr
	(*SheetAttrResponse)(nil), // 2: services.SheetAttrResponse
	(*common.Error)(nil),      // 3: common.Error
	(*common.Info)(nil),       // 4: common.Info
	(*common.Pager)(nil),      // 5: common.Pager
}
var file_sheetAttrService_proto_depIdxs = []int32{
	3,  // 0: services.SheetAttrResponse.error:type_name -> common.Error
	4,  // 1: services.SheetAttrResponse.info:type_name -> common.Info
	5,  // 2: services.SheetAttrResponse.pager:type_name -> common.Pager
	1,  // 3: services.SheetAttrResponse.entity:type_name -> services.SheetAttr
	1,  // 4: services.SheetAttrResponse.items:type_name -> services.SheetAttr
	1,  // 5: services.SheetAttrService.Create:input_type -> services.SheetAttr
	1,  // 6: services.SheetAttrService.Update:input_type -> services.SheetAttr
	0,  // 7: services.SheetAttrService.Delete:input_type -> services.SheetAttrWhere
	1,  // 8: services.SheetAttrService.Get:input_type -> services.SheetAttr
	0,  // 9: services.SheetAttrService.Search:input_type -> services.SheetAttrWhere
	0,  // 10: services.SheetAttrService.List:input_type -> services.SheetAttrWhere
	1,  // 11: services.SheetAttrService.ChangeSort:input_type -> services.SheetAttr
	2,  // 12: services.SheetAttrService.Create:output_type -> services.SheetAttrResponse
	2,  // 13: services.SheetAttrService.Update:output_type -> services.SheetAttrResponse
	2,  // 14: services.SheetAttrService.Delete:output_type -> services.SheetAttrResponse
	2,  // 15: services.SheetAttrService.Get:output_type -> services.SheetAttrResponse
	2,  // 16: services.SheetAttrService.Search:output_type -> services.SheetAttrResponse
	2,  // 17: services.SheetAttrService.List:output_type -> services.SheetAttrResponse
	2,  // 18: services.SheetAttrService.ChangeSort:output_type -> services.SheetAttrResponse
	12, // [12:19] is the sub-list for method output_type
	5,  // [5:12] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_sheetAttrService_proto_init() }
func file_sheetAttrService_proto_init() {
	if File_sheetAttrService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sheetAttrService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SheetAttrWhere); i {
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
		file_sheetAttrService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SheetAttr); i {
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
		file_sheetAttrService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SheetAttrResponse); i {
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
			RawDescriptor: file_sheetAttrService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sheetAttrService_proto_goTypes,
		DependencyIndexes: file_sheetAttrService_proto_depIdxs,
		MessageInfos:      file_sheetAttrService_proto_msgTypes,
	}.Build()
	File_sheetAttrService_proto = out.File
	file_sheetAttrService_proto_rawDesc = nil
	file_sheetAttrService_proto_goTypes = nil
	file_sheetAttrService_proto_depIdxs = nil
}
