// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: tagGroupService.proto

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

// 客户标签组
type TagGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`                                 //ID
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`                              //标签组名称
	AddRule   string `protobuf:"bytes,3,opt,name=add_rule,json=addRule,proto3" json:"add_rule"`         //添加规则
	IsHandAdd string `protobuf:"bytes,4,opt,name=is_hand_add,json=isHandAdd,proto3" json:"is_hand_add"` //手动打标
	IsAutoAdd string `protobuf:"bytes,5,opt,name=is_auto_add,json=isAutoAdd,proto3" json:"is_auto_add"` //自动打标
	Status    string `protobuf:"bytes,6,opt,name=status,proto3" json:"status"`                          //状态
	CreatedAt int64  `protobuf:"varint,7,opt,name=created_at,json=createdAt,proto3" json:"created_at"`  //创建时间
	UpdatedAt int64  `protobuf:"varint,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`  //修改时间
	Tags      []*Tag `protobuf:"bytes,9,rep,name=tags,proto3" json:"tags"`
}

func (x *TagGroup) Reset() {
	*x = TagGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tagGroupService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagGroup) ProtoMessage() {}

func (x *TagGroup) ProtoReflect() protoreflect.Message {
	mi := &file_tagGroupService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagGroup.ProtoReflect.Descriptor instead.
func (*TagGroup) Descriptor() ([]byte, []int) {
	return file_tagGroupService_proto_rawDescGZIP(), []int{0}
}

func (x *TagGroup) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TagGroup) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TagGroup) GetAddRule() string {
	if x != nil {
		return x.AddRule
	}
	return ""
}

func (x *TagGroup) GetIsHandAdd() string {
	if x != nil {
		return x.IsHandAdd
	}
	return ""
}

func (x *TagGroup) GetIsAutoAdd() string {
	if x != nil {
		return x.IsAutoAdd
	}
	return ""
}

func (x *TagGroup) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *TagGroup) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *TagGroup) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *TagGroup) GetTags() []*Tag {
	if x != nil {
		return x.Tags
	}
	return nil
}

// 客户标签组请求参数
type TagGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Top       int32    `protobuf:"varint,1,opt,name=top,proto3" json:"top"`
	Paged     int64    `protobuf:"varint,2,opt,name=paged,proto3" json:"paged"`
	PageSize  int64    `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Keywords  string   `protobuf:"bytes,4,opt,name=keywords,proto3" json:"keywords"`
	Sort      []string `protobuf:"bytes,5,rep,name=sort,proto3" json:"sort"`
	DateRange []string `protobuf:"bytes,6,rep,name=date_range,json=dateRange,proto3" json:"date_range"`
	Ids       []int64  `protobuf:"varint,7,rep,packed,name=ids,proto3" json:"ids"`
	Id        int64    `protobuf:"varint,8,opt,name=id,proto3" json:"id"`
	//以下为自定义参数
	Name      string `protobuf:"bytes,11,opt,name=name,proto3" json:"name"`                              //标签组名称
	IsHandAdd string `protobuf:"bytes,12,opt,name=is_hand_add,json=isHandAdd,proto3" json:"is_hand_add"` //手动打标
	IsAutoAdd string `protobuf:"bytes,13,opt,name=is_auto_add,json=isAutoAdd,proto3" json:"is_auto_add"` //自动打标
	Status    string `protobuf:"bytes,14,opt,name=status,proto3" json:"status"`                          //状态
}

func (x *TagGroupRequest) Reset() {
	*x = TagGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tagGroupService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagGroupRequest) ProtoMessage() {}

func (x *TagGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tagGroupService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagGroupRequest.ProtoReflect.Descriptor instead.
func (*TagGroupRequest) Descriptor() ([]byte, []int) {
	return file_tagGroupService_proto_rawDescGZIP(), []int{1}
}

func (x *TagGroupRequest) GetTop() int32 {
	if x != nil {
		return x.Top
	}
	return 0
}

func (x *TagGroupRequest) GetPaged() int64 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *TagGroupRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *TagGroupRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *TagGroupRequest) GetSort() []string {
	if x != nil {
		return x.Sort
	}
	return nil
}

func (x *TagGroupRequest) GetDateRange() []string {
	if x != nil {
		return x.DateRange
	}
	return nil
}

func (x *TagGroupRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *TagGroupRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TagGroupRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TagGroupRequest) GetIsHandAdd() string {
	if x != nil {
		return x.IsHandAdd
	}
	return ""
}

func (x *TagGroupRequest) GetIsAutoAdd() string {
	if x != nil {
		return x.IsAutoAdd
	}
	return ""
}

func (x *TagGroupRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

// 客户标签组响应数据
type TagGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg    string        `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Entity *TagGroup     `protobuf:"bytes,3,opt,name=entity,proto3" json:"entity"`
	Items  []*TagGroup   `protobuf:"bytes,4,rep,name=items,proto3" json:"items"`
}

func (x *TagGroupResponse) Reset() {
	*x = TagGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tagGroupService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagGroupResponse) ProtoMessage() {}

func (x *TagGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tagGroupService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagGroupResponse.ProtoReflect.Descriptor instead.
func (*TagGroupResponse) Descriptor() ([]byte, []int) {
	return file_tagGroupService_proto_rawDescGZIP(), []int{2}
}

func (x *TagGroupResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *TagGroupResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *TagGroupResponse) GetEntity() *TagGroup {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *TagGroupResponse) GetItems() []*TagGroup {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_tagGroupService_proto protoreflect.FileDescriptor

var file_tagGroupService_proto_rawDesc = []byte{
	0x0a, 0x15, 0x74, 0x61, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x74, 0x61, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82, 0x02, 0x0a, 0x08, 0x54, 0x61, 0x67, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x64, 0x64, 0x5f, 0x72,
	0x75, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x52, 0x75,
	0x6c, 0x65, 0x12, 0x1e, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x5f, 0x61, 0x64,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x73, 0x48, 0x61, 0x6e, 0x64, 0x41,
	0x64, 0x64, 0x12, 0x1e, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x61, 0x64,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x73, 0x41, 0x75, 0x74, 0x6f, 0x41,
	0x64, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x21, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73,
	0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0xb3, 0x02, 0x0a, 0x0f,
	0x54, 0x61, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x74, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x74, 0x6f,
	0x70, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04,
	0x73, 0x6f, 0x72, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x61, 0x6e,
	0x67, 0x65, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x64, 0x61, 0x74, 0x65, 0x52, 0x61,
	0x6e, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x03,
	0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0b, 0x69, 0x73, 0x5f,
	0x68, 0x61, 0x6e, 0x64, 0x5f, 0x61, 0x64, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x69, 0x73, 0x48, 0x61, 0x6e, 0x64, 0x41, 0x64, 0x64, 0x12, 0x1e, 0x0a, 0x0b, 0x69, 0x73, 0x5f,
	0x61, 0x75, 0x74, 0x6f, 0x5f, 0x61, 0x64, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x69, 0x73, 0x41, 0x75, 0x74, 0x6f, 0x41, 0x64, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x9f, 0x01, 0x0a, 0x10, 0x54, 0x61, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x2a, 0x0a,
	0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x61, 0x67, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x28, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x54, 0x61, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x32, 0x82, 0x03, 0x0a, 0x0f, 0x54, 0x61, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x61, 0x67,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x1a, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x54, 0x61, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x12, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x61, 0x67, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x1a, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x61, 0x67,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x3a, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x61, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x1a, 0x1a, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x61, 0x67, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x03, 0x47,
	0x65, 0x74, 0x12, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x61,
	0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x1a, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x54, 0x61, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x19,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x61, 0x67, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x61, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x61, 0x67, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x61, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tagGroupService_proto_rawDescOnce sync.Once
	file_tagGroupService_proto_rawDescData = file_tagGroupService_proto_rawDesc
)

func file_tagGroupService_proto_rawDescGZIP() []byte {
	file_tagGroupService_proto_rawDescOnce.Do(func() {
		file_tagGroupService_proto_rawDescData = protoimpl.X.CompressGZIP(file_tagGroupService_proto_rawDescData)
	})
	return file_tagGroupService_proto_rawDescData
}

var file_tagGroupService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_tagGroupService_proto_goTypes = []interface{}{
	(*TagGroup)(nil),         // 0: services.TagGroup
	(*TagGroupRequest)(nil),  // 1: services.TagGroupRequest
	(*TagGroupResponse)(nil), // 2: services.TagGroupResponse
	(*Tag)(nil),              // 3: services.Tag
	(*common.Pager)(nil),     // 4: common.Pager
}
var file_tagGroupService_proto_depIdxs = []int32{
	3,  // 0: services.TagGroup.tags:type_name -> services.Tag
	4,  // 1: services.TagGroupResponse.pager:type_name -> common.Pager
	0,  // 2: services.TagGroupResponse.entity:type_name -> services.TagGroup
	0,  // 3: services.TagGroupResponse.items:type_name -> services.TagGroup
	0,  // 4: services.TagGroupService.Create:input_type -> services.TagGroup
	0,  // 5: services.TagGroupService.Update:input_type -> services.TagGroup
	0,  // 6: services.TagGroupService.Delete:input_type -> services.TagGroup
	0,  // 7: services.TagGroupService.Get:input_type -> services.TagGroup
	1,  // 8: services.TagGroupService.Search:input_type -> services.TagGroupRequest
	1,  // 9: services.TagGroupService.List:input_type -> services.TagGroupRequest
	2,  // 10: services.TagGroupService.Create:output_type -> services.TagGroupResponse
	2,  // 11: services.TagGroupService.Update:output_type -> services.TagGroupResponse
	2,  // 12: services.TagGroupService.Delete:output_type -> services.TagGroupResponse
	2,  // 13: services.TagGroupService.Get:output_type -> services.TagGroupResponse
	2,  // 14: services.TagGroupService.Search:output_type -> services.TagGroupResponse
	2,  // 15: services.TagGroupService.List:output_type -> services.TagGroupResponse
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_tagGroupService_proto_init() }
func file_tagGroupService_proto_init() {
	if File_tagGroupService_proto != nil {
		return
	}
	file_tagService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tagGroupService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagGroup); i {
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
		file_tagGroupService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagGroupRequest); i {
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
		file_tagGroupService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagGroupResponse); i {
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
			RawDescriptor: file_tagGroupService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tagGroupService_proto_goTypes,
		DependencyIndexes: file_tagGroupService_proto_depIdxs,
		MessageInfos:      file_tagGroupService_proto_msgTypes,
	}.Build()
	File_tagGroupService_proto = out.File
	file_tagGroupService_proto_rawDesc = nil
	file_tagGroupService_proto_goTypes = nil
	file_tagGroupService_proto_depIdxs = nil
}
