// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: navService.proto

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

// 导航
type Nav struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string  `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Title         string  `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	TitleEn       string  `protobuf:"bytes,5,opt,name=title_en,json=titleEn,proto3" json:"title_en,omitempty"`
	Icon          string  `protobuf:"bytes,6,opt,name=icon,proto3" json:"icon,omitempty"`
	Nocache       bool    `protobuf:"varint,7,opt,name=nocache,proto3" json:"nocache,omitempty"`
	Meta          string  `protobuf:"bytes,8,opt,name=meta,proto3" json:"meta,omitempty"`
	NavId         int64   `protobuf:"varint,9,opt,name=nav_id,json=navId,proto3" json:"nav_id,omitempty"`
	RedirectUrl   string  `protobuf:"bytes,10,opt,name=redirect_url,json=redirectUrl,proto3" json:"redirect_url,omitempty"`
	ContainRoutes string  `protobuf:"bytes,11,opt,name=contain_routes,json=containRoutes,proto3" json:"contain_routes,omitempty"`
	PermissionId  int64   `protobuf:"varint,12,opt,name=permission_id,json=permissionId,proto3" json:"permission_id,omitempty"`
	HideSidebar   bool    `protobuf:"varint,13,opt,name=hide_sidebar,json=hideSidebar,proto3" json:"hide_sidebar,omitempty"`
	Disabled      bool    `protobuf:"varint,14,opt,name=disabled,proto3" json:"disabled,omitempty"`
	Sorting       string  `protobuf:"bytes,15,opt,name=sorting,proto3" json:"sorting,omitempty"`
	Memo          string  `protobuf:"bytes,16,opt,name=memo,proto3" json:"memo,omitempty"`
	CreatedAt     string  `protobuf:"bytes,17,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string  `protobuf:"bytes,18,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Children      []*Nav  `protobuf:"bytes,19,rep,name=children,proto3" json:"children,omitempty"`
	Ids           []int64 `protobuf:"varint,20,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *Nav) Reset() {
	*x = Nav{}
	if protoimpl.UnsafeEnabled {
		mi := &file_navService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Nav) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Nav) ProtoMessage() {}

func (x *Nav) ProtoReflect() protoreflect.Message {
	mi := &file_navService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Nav.ProtoReflect.Descriptor instead.
func (*Nav) Descriptor() ([]byte, []int) {
	return file_navService_proto_rawDescGZIP(), []int{0}
}

func (x *Nav) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Nav) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Nav) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Nav) GetTitleEn() string {
	if x != nil {
		return x.TitleEn
	}
	return ""
}

func (x *Nav) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *Nav) GetNocache() bool {
	if x != nil {
		return x.Nocache
	}
	return false
}

func (x *Nav) GetMeta() string {
	if x != nil {
		return x.Meta
	}
	return ""
}

func (x *Nav) GetNavId() int64 {
	if x != nil {
		return x.NavId
	}
	return 0
}

func (x *Nav) GetRedirectUrl() string {
	if x != nil {
		return x.RedirectUrl
	}
	return ""
}

func (x *Nav) GetContainRoutes() string {
	if x != nil {
		return x.ContainRoutes
	}
	return ""
}

func (x *Nav) GetPermissionId() int64 {
	if x != nil {
		return x.PermissionId
	}
	return 0
}

func (x *Nav) GetHideSidebar() bool {
	if x != nil {
		return x.HideSidebar
	}
	return false
}

func (x *Nav) GetDisabled() bool {
	if x != nil {
		return x.Disabled
	}
	return false
}

func (x *Nav) GetSorting() string {
	if x != nil {
		return x.Sorting
	}
	return ""
}

func (x *Nav) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *Nav) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Nav) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Nav) GetChildren() []*Nav {
	if x != nil {
		return x.Children
	}
	return nil
}

func (x *Nav) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type NavWhere struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int32  `protobuf:"varint,1,opt,name=paged,proto3" json:"paged,omitempty"`
	PageSize int32  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Sorting  string `protobuf:"bytes,3,opt,name=sorting,proto3" json:"sorting,omitempty"`
	//base params
	Id    int64   `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
	Name  string  `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	NavId int64   `protobuf:"varint,6,opt,name=nav_id,json=navId,proto3" json:"nav_id,omitempty"`
	Ids   []int64 `protobuf:"varint,8,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *NavWhere) Reset() {
	*x = NavWhere{}
	if protoimpl.UnsafeEnabled {
		mi := &file_navService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NavWhere) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NavWhere) ProtoMessage() {}

func (x *NavWhere) ProtoReflect() protoreflect.Message {
	mi := &file_navService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NavWhere.ProtoReflect.Descriptor instead.
func (*NavWhere) Descriptor() ([]byte, []int) {
	return file_navService_proto_rawDescGZIP(), []int{1}
}

func (x *NavWhere) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *NavWhere) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *NavWhere) GetSorting() string {
	if x != nil {
		return x.Sorting
	}
	return ""
}

func (x *NavWhere) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *NavWhere) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NavWhere) GetNavId() int64 {
	if x != nil {
		return x.NavId
	}
	return 0
}

func (x *NavWhere) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

//
type NavResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Nav          `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items  []*Nav        `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error  *common.Error `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info   *common.Info  `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *NavResponse) Reset() {
	*x = NavResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_navService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NavResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NavResponse) ProtoMessage() {}

func (x *NavResponse) ProtoReflect() protoreflect.Message {
	mi := &file_navService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NavResponse.ProtoReflect.Descriptor instead.
func (*NavResponse) Descriptor() ([]byte, []int) {
	return file_navService_proto_rawDescGZIP(), []int{2}
}

func (x *NavResponse) GetEntity() *Nav {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *NavResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *NavResponse) GetItems() []*Nav {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *NavResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *NavResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_navService_proto protoreflect.FileDescriptor

var file_navService_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6e, 0x61, 0x76, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x8a, 0x04, 0x0a, 0x03, 0x4e, 0x61, 0x76, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x5f, 0x65, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x45, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x69, 0x63, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e,
	0x12, 0x18, 0x0a, 0x07, 0x6e, 0x6f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x6e, 0x6f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65,
	0x74, 0x61, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x15,
	0x0a, 0x06, 0x6e, 0x61, 0x76, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x6e, 0x61, 0x76, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x12,
	0x23, 0x0a, 0x0d, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x68, 0x69, 0x64, 0x65, 0x5f, 0x73, 0x69, 0x64,
	0x65, 0x62, 0x61, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x68, 0x69, 0x64, 0x65,
	0x53, 0x69, 0x64, 0x65, 0x62, 0x61, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a,
	0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x6d,
	0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x12,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x29, 0x0a, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18, 0x13, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4e, 0x61, 0x76,
	0x52, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64,
	0x73, 0x18, 0x14, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0xa4, 0x01, 0x0a,
	0x08, 0x4e, 0x61, 0x76, 0x57, 0x68, 0x65, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67,
	0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73,
	0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x6e, 0x61,
	0x76, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6e, 0x61, 0x76, 0x49,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03,
	0x69, 0x64, 0x73, 0x22, 0xc5, 0x01, 0x0a, 0x0b, 0x4e, 0x61, 0x76, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4e,
	0x61, 0x76, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61,
	0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12,
	0x23, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4e, 0x61, 0x76, 0x52, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0xd1, 0x01, 0x0a, 0x0a,
	0x4e, 0x61, 0x76, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2d, 0x0a, 0x03, 0x47, 0x65,
	0x74, 0x12, 0x0d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4e, 0x61, 0x76,
	0x1a, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4e, 0x61, 0x76, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x06, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x12, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x42, 0x61, 0x73,
	0x65, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x4e, 0x61, 0x76, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x2e, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x4e, 0x61, 0x76, 0x1a, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x4e, 0x61, 0x76, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x2e, 0x0a, 0x04, 0x54, 0x72, 0x65, 0x65, 0x12, 0x0d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x4e, 0x61, 0x76, 0x1a, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x4e, 0x61, 0x76, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_navService_proto_rawDescOnce sync.Once
	file_navService_proto_rawDescData = file_navService_proto_rawDesc
)

func file_navService_proto_rawDescGZIP() []byte {
	file_navService_proto_rawDescOnce.Do(func() {
		file_navService_proto_rawDescData = protoimpl.X.CompressGZIP(file_navService_proto_rawDescData)
	})
	return file_navService_proto_rawDescData
}

var file_navService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_navService_proto_goTypes = []interface{}{
	(*Nav)(nil),              // 0: services.Nav
	(*NavWhere)(nil),         // 1: services.NavWhere
	(*NavResponse)(nil),      // 2: services.NavResponse
	(*common.Pager)(nil),     // 3: common.Pager
	(*common.Error)(nil),     // 4: common.Error
	(*common.Info)(nil),      // 5: common.Info
	(*common.BaseWhere)(nil), // 6: common.BaseWhere
}
var file_navService_proto_depIdxs = []int32{
	0,  // 0: services.Nav.children:type_name -> services.Nav
	0,  // 1: services.NavResponse.entity:type_name -> services.Nav
	3,  // 2: services.NavResponse.pager:type_name -> common.Pager
	0,  // 3: services.NavResponse.items:type_name -> services.Nav
	4,  // 4: services.NavResponse.error:type_name -> common.Error
	5,  // 5: services.NavResponse.info:type_name -> common.Info
	0,  // 6: services.NavService.Get:input_type -> services.Nav
	6,  // 7: services.NavService.Search:input_type -> common.BaseWhere
	0,  // 8: services.NavService.List:input_type -> services.Nav
	0,  // 9: services.NavService.Tree:input_type -> services.Nav
	2,  // 10: services.NavService.Get:output_type -> services.NavResponse
	2,  // 11: services.NavService.Search:output_type -> services.NavResponse
	2,  // 12: services.NavService.List:output_type -> services.NavResponse
	2,  // 13: services.NavService.Tree:output_type -> services.NavResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_navService_proto_init() }
func file_navService_proto_init() {
	if File_navService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_navService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Nav); i {
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
		file_navService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NavWhere); i {
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
		file_navService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NavResponse); i {
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
			RawDescriptor: file_navService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_navService_proto_goTypes,
		DependencyIndexes: file_navService_proto_depIdxs,
		MessageInfos:      file_navService_proto_msgTypes,
	}.Build()
	File_navService_proto = out.File
	file_navService_proto_rawDesc = nil
	file_navService_proto_goTypes = nil
	file_navService_proto_depIdxs = nil
}
