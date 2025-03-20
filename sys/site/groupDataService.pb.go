// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: groupDataService.proto

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

// 组合数据
type GroupData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32                     `protobuf:"varint,1,opt,name=id,proto3" json:"id"`                                                                                              //ID
	GroupId   int32                     `protobuf:"varint,2,opt,name=group_id,json=groupId,proto3" json:"group_id"`                                                                     //组合ID
	Value     map[string]*GroupDataItem `protobuf:"bytes,3,rep,name=value,proto3" json:"value" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` //组合数据
	Sort      int32                     `protobuf:"varint,4,opt,name=sort,proto3" json:"sort"`                                                                                          //排序
	Status    string                    `protobuf:"bytes,5,opt,name=status,proto3" json:"status"`                                                                                       //状态
	CreatedAt int64                     `protobuf:"varint,6,opt,name=created_at,json=createdAt,proto3" json:"created_at"`                                                               //创建时间
	UpdatedAt int64                     `protobuf:"varint,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`                                                               //修改时间
}

func (x *GroupData) Reset() {
	*x = GroupData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_groupDataService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupData) ProtoMessage() {}

func (x *GroupData) ProtoReflect() protoreflect.Message {
	mi := &file_groupDataService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupData.ProtoReflect.Descriptor instead.
func (*GroupData) Descriptor() ([]byte, []int) {
	return file_groupDataService_proto_rawDescGZIP(), []int{0}
}

func (x *GroupData) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GroupData) GetGroupId() int32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *GroupData) GetValue() map[string]*GroupDataItem {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *GroupData) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *GroupData) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *GroupData) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *GroupData) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

// 组合数据项
type GroupDataItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  string `protobuf:"bytes,1,opt,name=type,proto3" json:"type"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value"`
}

func (x *GroupDataItem) Reset() {
	*x = GroupDataItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_groupDataService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupDataItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupDataItem) ProtoMessage() {}

func (x *GroupDataItem) ProtoReflect() protoreflect.Message {
	mi := &file_groupDataService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupDataItem.ProtoReflect.Descriptor instead.
func (*GroupDataItem) Descriptor() ([]byte, []int) {
	return file_groupDataService_proto_rawDescGZIP(), []int{1}
}

func (x *GroupDataItem) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *GroupDataItem) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// 组合数据请求参数
type GroupDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Top       int32    `protobuf:"varint,1,opt,name=top,proto3" json:"top"`
	Paged     int64    `protobuf:"varint,2,opt,name=paged,proto3" json:"paged"`
	PageSize  int64    `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Keywords  string   `protobuf:"bytes,4,opt,name=keywords,proto3" json:"keywords"`
	Sorts     []string `protobuf:"bytes,5,rep,name=sorts,proto3" json:"sorts"`
	DateRange []string `protobuf:"bytes,6,rep,name=date_range,json=dateRange,proto3" json:"date_range"`
	Ids       []int32  `protobuf:"varint,7,rep,packed,name=ids,proto3" json:"ids"`
	Id        int32    `protobuf:"varint,8,opt,name=id,proto3" json:"id"`
	//以下为自定义参数
	GroupId int32  `protobuf:"varint,11,opt,name=group_id,json=groupId,proto3" json:"group_id"` //组合ID
	Status  string `protobuf:"bytes,12,opt,name=status,proto3" json:"status"`                   //状态
}

func (x *GroupDataRequest) Reset() {
	*x = GroupDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_groupDataService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupDataRequest) ProtoMessage() {}

func (x *GroupDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_groupDataService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupDataRequest.ProtoReflect.Descriptor instead.
func (*GroupDataRequest) Descriptor() ([]byte, []int) {
	return file_groupDataService_proto_rawDescGZIP(), []int{2}
}

func (x *GroupDataRequest) GetTop() int32 {
	if x != nil {
		return x.Top
	}
	return 0
}

func (x *GroupDataRequest) GetPaged() int64 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *GroupDataRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *GroupDataRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *GroupDataRequest) GetSorts() []string {
	if x != nil {
		return x.Sorts
	}
	return nil
}

func (x *GroupDataRequest) GetDateRange() []string {
	if x != nil {
		return x.DateRange
	}
	return nil
}

func (x *GroupDataRequest) GetIds() []int32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *GroupDataRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GroupDataRequest) GetGroupId() int32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *GroupDataRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

// 组合数据响应数据
type GroupDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg    string        `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Entity *GroupData    `protobuf:"bytes,3,opt,name=entity,proto3" json:"entity"`
	Items  []*GroupData  `protobuf:"bytes,4,rep,name=items,proto3" json:"items"`
}

func (x *GroupDataResponse) Reset() {
	*x = GroupDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_groupDataService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupDataResponse) ProtoMessage() {}

func (x *GroupDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_groupDataService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupDataResponse.ProtoReflect.Descriptor instead.
func (*GroupDataResponse) Descriptor() ([]byte, []int) {
	return file_groupDataService_proto_rawDescGZIP(), []int{3}
}

func (x *GroupDataResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GroupDataResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *GroupDataResponse) GetEntity() *GroupData {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *GroupDataResponse) GetItems() []*GroupData {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_groupDataService_proto protoreflect.FileDescriptor

var file_groupDataService_proto_rawDesc = []byte{
	0x0a, 0x16, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa9, 0x02, 0x0a, 0x09, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x34,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x61,
	0x74, 0x61, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x1a, 0x51,
	0x0a, 0x0a, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2d,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x61,
	0x74, 0x61, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x39, 0x0a, 0x0d, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x61, 0x74, 0x61, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xfd, 0x01, 0x0a,
	0x10, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x74, 0x6f, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72,
	0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72,
	0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x05, 0x73, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x07,
	0x20, 0x03, 0x28, 0x05, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xa2, 0x01, 0x0a,
	0x11, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6d, 0x73, 0x67, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67,
	0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x06, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x61, 0x74, 0x61, 0x52, 0x06,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x29, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x61, 0x74, 0x61, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x32, 0x8f, 0x03, 0x0a, 0x10, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x61, 0x74, 0x61, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x13,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x44,
	0x61, 0x74, 0x61, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x3c, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x61, 0x74,
	0x61, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x39, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x1b, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x06, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x41, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_groupDataService_proto_rawDescOnce sync.Once
	file_groupDataService_proto_rawDescData = file_groupDataService_proto_rawDesc
)

func file_groupDataService_proto_rawDescGZIP() []byte {
	file_groupDataService_proto_rawDescOnce.Do(func() {
		file_groupDataService_proto_rawDescData = protoimpl.X.CompressGZIP(file_groupDataService_proto_rawDescData)
	})
	return file_groupDataService_proto_rawDescData
}

var file_groupDataService_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_groupDataService_proto_goTypes = []interface{}{
	(*GroupData)(nil),         // 0: services.GroupData
	(*GroupDataItem)(nil),     // 1: services.GroupDataItem
	(*GroupDataRequest)(nil),  // 2: services.GroupDataRequest
	(*GroupDataResponse)(nil), // 3: services.GroupDataResponse
	nil,                       // 4: services.GroupData.ValueEntry
	(*common.Pager)(nil),      // 5: common.Pager
}
var file_groupDataService_proto_depIdxs = []int32{
	4,  // 0: services.GroupData.value:type_name -> services.GroupData.ValueEntry
	5,  // 1: services.GroupDataResponse.pager:type_name -> common.Pager
	0,  // 2: services.GroupDataResponse.entity:type_name -> services.GroupData
	0,  // 3: services.GroupDataResponse.items:type_name -> services.GroupData
	1,  // 4: services.GroupData.ValueEntry.value:type_name -> services.GroupDataItem
	0,  // 5: services.GroupDataService.Create:input_type -> services.GroupData
	0,  // 6: services.GroupDataService.Update:input_type -> services.GroupData
	0,  // 7: services.GroupDataService.Delete:input_type -> services.GroupData
	0,  // 8: services.GroupDataService.Get:input_type -> services.GroupData
	2,  // 9: services.GroupDataService.Search:input_type -> services.GroupDataRequest
	2,  // 10: services.GroupDataService.List:input_type -> services.GroupDataRequest
	3,  // 11: services.GroupDataService.Create:output_type -> services.GroupDataResponse
	3,  // 12: services.GroupDataService.Update:output_type -> services.GroupDataResponse
	3,  // 13: services.GroupDataService.Delete:output_type -> services.GroupDataResponse
	3,  // 14: services.GroupDataService.Get:output_type -> services.GroupDataResponse
	3,  // 15: services.GroupDataService.Search:output_type -> services.GroupDataResponse
	3,  // 16: services.GroupDataService.List:output_type -> services.GroupDataResponse
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_groupDataService_proto_init() }
func file_groupDataService_proto_init() {
	if File_groupDataService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_groupDataService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupData); i {
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
		file_groupDataService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupDataItem); i {
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
		file_groupDataService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupDataRequest); i {
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
		file_groupDataService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupDataResponse); i {
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
			RawDescriptor: file_groupDataService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_groupDataService_proto_goTypes,
		DependencyIndexes: file_groupDataService_proto_depIdxs,
		MessageInfos:      file_groupDataService_proto_msgTypes,
	}.Build()
	File_groupDataService_proto = out.File
	file_groupDataService_proto_rawDesc = nil
	file_groupDataService_proto_goTypes = nil
	file_groupDataService_proto_depIdxs = nil
}
