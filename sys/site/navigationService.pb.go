// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: navigationService.proto

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

type NavigationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged           int64   `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize        int64   `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Sorting         string  `protobuf:"bytes,3,opt,name=sorting,proto3" json:"sorting"`
	Keywords        string  `protobuf:"bytes,4,opt,name=keywords,proto3" json:"keywords"`
	Id              int64   `protobuf:"varint,5,opt,name=id,proto3" json:"id"`
	ParentId        int64   `protobuf:"varint,6,opt,name=parent_id,json=parentId,proto3" json:"parent_id"`
	Code            string  `protobuf:"bytes,7,opt,name=code,proto3" json:"code"`
	Name            string  `protobuf:"bytes,8,opt,name=name,proto3" json:"name"`
	DataType        string  `protobuf:"bytes,9,opt,name=data_type,json=dataType,proto3" json:"data_type"`
	Type            string  `protobuf:"bytes,10,opt,name=type,proto3" json:"type"`
	Tag             string  `protobuf:"bytes,11,opt,name=tag,proto3" json:"tag"`
	DataValue       string  `protobuf:"bytes,12,opt,name=data_value,json=dataValue,proto3" json:"data_value"`
	IsShow          string  `protobuf:"bytes,13,opt,name=is_show,json=isShow,proto3" json:"is_show"`
	IsNewWindowOpen string  `protobuf:"bytes,14,opt,name=is_new_window_open,json=isNewWindowOpen,proto3" json:"is_new_window_open"`
	Ids             []int64 `protobuf:"varint,15,rep,packed,name=ids,proto3" json:"ids"`
}

func (x *NavigationRequest) Reset() {
	*x = NavigationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_navigationService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NavigationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NavigationRequest) ProtoMessage() {}

func (x *NavigationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_navigationService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NavigationRequest.ProtoReflect.Descriptor instead.
func (*NavigationRequest) Descriptor() ([]byte, []int) {
	return file_navigationService_proto_rawDescGZIP(), []int{0}
}

func (x *NavigationRequest) GetPaged() int64 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *NavigationRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *NavigationRequest) GetSorting() string {
	if x != nil {
		return x.Sorting
	}
	return ""
}

func (x *NavigationRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *NavigationRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *NavigationRequest) GetParentId() int64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *NavigationRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *NavigationRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NavigationRequest) GetDataType() string {
	if x != nil {
		return x.DataType
	}
	return ""
}

func (x *NavigationRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *NavigationRequest) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *NavigationRequest) GetDataValue() string {
	if x != nil {
		return x.DataValue
	}
	return ""
}

func (x *NavigationRequest) GetIsShow() string {
	if x != nil {
		return x.IsShow
	}
	return ""
}

func (x *NavigationRequest) GetIsNewWindowOpen() string {
	if x != nil {
		return x.IsNewWindowOpen
	}
	return ""
}

func (x *NavigationRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type Navigation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                  int64         `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Type                string        `protobuf:"bytes,2,opt,name=type,proto3" json:"type"`
	Name                string        `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`
	Code                string        `protobuf:"bytes,4,opt,name=code,proto3" json:"code"`
	ParentId            int64         `protobuf:"varint,5,opt,name=parent_id,json=parentId,proto3" json:"parent_id"`
	DeptPath            string        `protobuf:"bytes,6,opt,name=dept_path,json=deptPath,proto3" json:"dept_path"`
	RouteType           string        `protobuf:"bytes,7,opt,name=route_type,json=routeType,proto3" json:"route_type"`
	DataType            string        `protobuf:"bytes,8,opt,name=data_type,json=dataType,proto3" json:"data_type"`
	DataValue           string        `protobuf:"bytes,9,opt,name=data_value,json=dataValue,proto3" json:"data_value"`
	ImageUrl            string        `protobuf:"bytes,10,opt,name=image_url,json=imageUrl,proto3" json:"image_url"`
	IsShow              string        `protobuf:"bytes,11,opt,name=is_show,json=isShow,proto3" json:"is_show"`
	IsNewWindowOpen     string        `protobuf:"bytes,12,opt,name=is_new_window_open,json=isNewWindowOpen,proto3" json:"is_new_window_open"`
	Sort                int32         `protobuf:"varint,13,opt,name=sort,proto3" json:"sort"`
	CreatedAt           int64         `protobuf:"varint,14,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt           int64         `protobuf:"varint,15,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	Parent              *Navigation   `protobuf:"bytes,16,opt,name=parent,proto3" json:"parent"`
	Tags                []string      `protobuf:"bytes,17,rep,name=tags,proto3" json:"tags"`
	Children            []*Navigation `protobuf:"bytes,18,rep,name=children,proto3" json:"children"`
	RouteTypeName       string        `protobuf:"bytes,19,opt,name=route_type_name,json=routeTypeName,proto3" json:"route_type_name"`
	DataTypeName        string        `protobuf:"bytes,20,opt,name=data_type_name,json=dataTypeName,proto3" json:"data_type_name"`
	IsShowName          string        `protobuf:"bytes,21,opt,name=is_show_name,json=isShowName,proto3" json:"is_show_name"`
	IsNewWindowOpenName string        `protobuf:"bytes,22,opt,name=is_new_window_open_name,json=isNewWindowOpenName,proto3" json:"is_new_window_open_name"`
}

func (x *Navigation) Reset() {
	*x = Navigation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_navigationService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Navigation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Navigation) ProtoMessage() {}

func (x *Navigation) ProtoReflect() protoreflect.Message {
	mi := &file_navigationService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Navigation.ProtoReflect.Descriptor instead.
func (*Navigation) Descriptor() ([]byte, []int) {
	return file_navigationService_proto_rawDescGZIP(), []int{1}
}

func (x *Navigation) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Navigation) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Navigation) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Navigation) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Navigation) GetParentId() int64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *Navigation) GetDeptPath() string {
	if x != nil {
		return x.DeptPath
	}
	return ""
}

func (x *Navigation) GetRouteType() string {
	if x != nil {
		return x.RouteType
	}
	return ""
}

func (x *Navigation) GetDataType() string {
	if x != nil {
		return x.DataType
	}
	return ""
}

func (x *Navigation) GetDataValue() string {
	if x != nil {
		return x.DataValue
	}
	return ""
}

func (x *Navigation) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *Navigation) GetIsShow() string {
	if x != nil {
		return x.IsShow
	}
	return ""
}

func (x *Navigation) GetIsNewWindowOpen() string {
	if x != nil {
		return x.IsNewWindowOpen
	}
	return ""
}

func (x *Navigation) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *Navigation) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Navigation) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *Navigation) GetParent() *Navigation {
	if x != nil {
		return x.Parent
	}
	return nil
}

func (x *Navigation) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Navigation) GetChildren() []*Navigation {
	if x != nil {
		return x.Children
	}
	return nil
}

func (x *Navigation) GetRouteTypeName() string {
	if x != nil {
		return x.RouteTypeName
	}
	return ""
}

func (x *Navigation) GetDataTypeName() string {
	if x != nil {
		return x.DataTypeName
	}
	return ""
}

func (x *Navigation) GetIsShowName() string {
	if x != nil {
		return x.IsShowName
	}
	return ""
}

func (x *Navigation) GetIsNewWindowOpenName() string {
	if x != nil {
		return x.IsNewWindowOpenName
	}
	return ""
}

type NavigationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Navigation   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*Navigation `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Info   string        `protobuf:"bytes,4,opt,name=info,proto3" json:"info"`
}

func (x *NavigationResponse) Reset() {
	*x = NavigationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_navigationService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NavigationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NavigationResponse) ProtoMessage() {}

func (x *NavigationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_navigationService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NavigationResponse.ProtoReflect.Descriptor instead.
func (*NavigationResponse) Descriptor() ([]byte, []int) {
	return file_navigationService_proto_rawDescGZIP(), []int{2}
}

func (x *NavigationResponse) GetEntity() *Navigation {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *NavigationResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *NavigationResponse) GetItems() []*Navigation {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *NavigationResponse) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

var File_navigationService_proto protoreflect.FileDescriptor

var file_navigationService_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x03, 0x0a, 0x11, 0x4e, 0x61, 0x76, 0x69, 0x67,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x61, 0x67,
	0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79,
	0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79,
	0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x61,
	0x74, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64,
	0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74,
	0x61, 0x67, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x1d, 0x0a,
	0x0a, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x64, 0x61, 0x74, 0x61, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x17, 0x0a, 0x07,
	0x69, 0x73, 0x5f, 0x73, 0x68, 0x6f, 0x77, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69,
	0x73, 0x53, 0x68, 0x6f, 0x77, 0x12, 0x2b, 0x0a, 0x12, 0x69, 0x73, 0x5f, 0x6e, 0x65, 0x77, 0x5f,
	0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x69, 0x73, 0x4e, 0x65, 0x77, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x4f, 0x70,
	0x65, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x03, 0x52,
	0x03, 0x69, 0x64, 0x73, 0x22, 0xbc, 0x05, 0x0a, 0x0a, 0x4e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x64, 0x65, 0x70, 0x74, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x64, 0x65, 0x70, 0x74, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x61,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74,
	0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x61, 0x74, 0x61, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72,
	0x6c, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x73, 0x68, 0x6f, 0x77, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x69, 0x73, 0x53, 0x68, 0x6f, 0x77, 0x12, 0x2b, 0x0a, 0x12, 0x69, 0x73,
	0x5f, 0x6e, 0x65, 0x77, 0x5f, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x5f, 0x6f, 0x70, 0x65, 0x6e,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x69, 0x73, 0x4e, 0x65, 0x77, 0x57, 0x69, 0x6e,
	0x64, 0x6f, 0x77, 0x4f, 0x70, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x2c, 0x0a, 0x06, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x4e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18,
	0x11, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x30, 0x0a, 0x08, 0x63,
	0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18, 0x12, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x12, 0x26, 0x0a,
	0x0f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64,
	0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0c, 0x69,
	0x73, 0x5f, 0x73, 0x68, 0x6f, 0x77, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x69, 0x73, 0x53, 0x68, 0x6f, 0x77, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x34, 0x0a,
	0x17, 0x69, 0x73, 0x5f, 0x6e, 0x65, 0x77, 0x5f, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x5f, 0x6f,
	0x70, 0x65, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13,
	0x69, 0x73, 0x4e, 0x65, 0x77, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x4f, 0x70, 0x65, 0x6e, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0xa7, 0x01, 0x0a, 0x12, 0x4e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x2a, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0xe1, 0x03,
	0x0a, 0x11, 0x4e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4e,
	0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4e,
	0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4e, 0x61, 0x76, 0x69,
	0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x3e, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4e, 0x61, 0x76, 0x69,
	0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x45, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1b, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x4e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4e, 0x61, 0x76, 0x69, 0x67,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x04,
	0x54, 0x72, 0x65, 0x65, 0x12, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x4e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4e, 0x61, 0x76,
	0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_navigationService_proto_rawDescOnce sync.Once
	file_navigationService_proto_rawDescData = file_navigationService_proto_rawDesc
)

func file_navigationService_proto_rawDescGZIP() []byte {
	file_navigationService_proto_rawDescOnce.Do(func() {
		file_navigationService_proto_rawDescData = protoimpl.X.CompressGZIP(file_navigationService_proto_rawDescData)
	})
	return file_navigationService_proto_rawDescData
}

var file_navigationService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_navigationService_proto_goTypes = []interface{}{
	(*NavigationRequest)(nil),  // 0: services.NavigationRequest
	(*Navigation)(nil),         // 1: services.Navigation
	(*NavigationResponse)(nil), // 2: services.NavigationResponse
	(*common.Pager)(nil),       // 3: common.Pager
}
var file_navigationService_proto_depIdxs = []int32{
	1,  // 0: services.Navigation.parent:type_name -> services.Navigation
	1,  // 1: services.Navigation.children:type_name -> services.Navigation
	1,  // 2: services.NavigationResponse.entity:type_name -> services.Navigation
	3,  // 3: services.NavigationResponse.pager:type_name -> common.Pager
	1,  // 4: services.NavigationResponse.items:type_name -> services.Navigation
	1,  // 5: services.NavigationService.Create:input_type -> services.Navigation
	1,  // 6: services.NavigationService.Update:input_type -> services.Navigation
	1,  // 7: services.NavigationService.Get:input_type -> services.Navigation
	1,  // 8: services.NavigationService.Delete:input_type -> services.Navigation
	0,  // 9: services.NavigationService.Search:input_type -> services.NavigationRequest
	0,  // 10: services.NavigationService.List:input_type -> services.NavigationRequest
	0,  // 11: services.NavigationService.Tree:input_type -> services.NavigationRequest
	2,  // 12: services.NavigationService.Create:output_type -> services.NavigationResponse
	2,  // 13: services.NavigationService.Update:output_type -> services.NavigationResponse
	2,  // 14: services.NavigationService.Get:output_type -> services.NavigationResponse
	2,  // 15: services.NavigationService.Delete:output_type -> services.NavigationResponse
	2,  // 16: services.NavigationService.Search:output_type -> services.NavigationResponse
	2,  // 17: services.NavigationService.List:output_type -> services.NavigationResponse
	2,  // 18: services.NavigationService.Tree:output_type -> services.NavigationResponse
	12, // [12:19] is the sub-list for method output_type
	5,  // [5:12] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_navigationService_proto_init() }
func file_navigationService_proto_init() {
	if File_navigationService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_navigationService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NavigationRequest); i {
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
		file_navigationService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Navigation); i {
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
		file_navigationService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NavigationResponse); i {
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
			RawDescriptor: file_navigationService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_navigationService_proto_goTypes,
		DependencyIndexes: file_navigationService_proto_depIdxs,
		MessageInfos:      file_navigationService_proto_msgTypes,
	}.Build()
	File_navigationService_proto = out.File
	file_navigationService_proto_rawDesc = nil
	file_navigationService_proto_goTypes = nil
	file_navigationService_proto_depIdxs = nil
}
