// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: deptService.proto

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

type Dept struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Name      string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	ParentId  int32   `protobuf:"varint,3,opt,name=parent_id,json=parentId,proto3" json:"parent_id"`
	Path      string  `protobuf:"bytes,4,opt,name=path,proto3" json:"path"`
	Sort      int32   `protobuf:"varint,5,opt,name=sort,proto3" json:"sort"`
	Leader    string  `protobuf:"bytes,6,opt,name=leader,proto3" json:"leader"`
	Phone     string  `protobuf:"bytes,7,opt,name=phone,proto3" json:"phone"`
	Email     string  `protobuf:"bytes,8,opt,name=email,proto3" json:"email"`
	Status    string  `protobuf:"bytes,9,opt,name=status,proto3" json:"status"`
	CreatedAt string  `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt string  `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	Children  []*Dept `protobuf:"bytes,12,rep,name=children,proto3" json:"children"`
}

func (x *Dept) Reset() {
	*x = Dept{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deptService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dept) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dept) ProtoMessage() {}

func (x *Dept) ProtoReflect() protoreflect.Message {
	mi := &file_deptService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dept.ProtoReflect.Descriptor instead.
func (*Dept) Descriptor() ([]byte, []int) {
	return file_deptService_proto_rawDescGZIP(), []int{0}
}

func (x *Dept) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Dept) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Dept) GetParentId() int32 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *Dept) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Dept) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *Dept) GetLeader() string {
	if x != nil {
		return x.Leader
	}
	return ""
}

func (x *Dept) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Dept) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Dept) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Dept) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Dept) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Dept) GetChildren() []*Dept {
	if x != nil {
		return x.Children
	}
	return nil
}

type DeptRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int64 `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Top      int64 `protobuf:"varint,3,opt,name=top,proto3" json:"top"`
	//my self
	Id       int32   `protobuf:"varint,4,opt,name=id,proto3" json:"id"`
	Name     string  `protobuf:"bytes,5,opt,name=name,proto3" json:"name"`
	ParentId int32   `protobuf:"varint,6,opt,name=parent_id,json=parentId,proto3" json:"parent_id"`
	Status   string  `protobuf:"bytes,7,opt,name=status,proto3" json:"status"`
	Ids      []int32 `protobuf:"varint,8,rep,packed,name=ids,proto3" json:"ids"`
}

func (x *DeptRequest) Reset() {
	*x = DeptRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deptService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeptRequest) ProtoMessage() {}

func (x *DeptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_deptService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeptRequest.ProtoReflect.Descriptor instead.
func (*DeptRequest) Descriptor() ([]byte, []int) {
	return file_deptService_proto_rawDescGZIP(), []int{1}
}

func (x *DeptRequest) GetPaged() int64 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *DeptRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *DeptRequest) GetTop() int64 {
	if x != nil {
		return x.Top
	}
	return 0
}

func (x *DeptRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DeptRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DeptRequest) GetParentId() int32 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *DeptRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *DeptRequest) GetIds() []int32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type DeptData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Dept             `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager     `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*Dept           `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Params map[string]string `protobuf:"bytes,4,rep,name=params,proto3" json:"params" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Info   *common.Info      `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
}

func (x *DeptData) Reset() {
	*x = DeptData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deptService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeptData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeptData) ProtoMessage() {}

func (x *DeptData) ProtoReflect() protoreflect.Message {
	mi := &file_deptService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeptData.ProtoReflect.Descriptor instead.
func (*DeptData) Descriptor() ([]byte, []int) {
	return file_deptService_proto_rawDescGZIP(), []int{2}
}

func (x *DeptData) GetEntity() *Dept {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *DeptData) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *DeptData) GetItems() []*Dept {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *DeptData) GetParams() map[string]string {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *DeptData) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type DeptResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error *common.Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error"`
	Data  *DeptData     `protobuf:"bytes,2,opt,name=data,proto3" json:"data"`
}

func (x *DeptResponse) Reset() {
	*x = DeptResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deptService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeptResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeptResponse) ProtoMessage() {}

func (x *DeptResponse) ProtoReflect() protoreflect.Message {
	mi := &file_deptService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeptResponse.ProtoReflect.Descriptor instead.
func (*DeptResponse) Descriptor() ([]byte, []int) {
	return file_deptService_proto_rawDescGZIP(), []int{3}
}

func (x *DeptResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *DeptResponse) GetData() *DeptData {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_deptService_proto protoreflect.FileDescriptor

var file_deptService_proto_rawDesc = []byte{
	0x0a, 0x11, 0x64, 0x65, 0x70, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xb5, 0x02, 0x0a, 0x04, 0x44, 0x65, 0x70, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x6f,
	0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x2a, 0x0a, 0x08,
	0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x70, 0x74, 0x52, 0x08,
	0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x22, 0xbd, 0x01, 0x0a, 0x0b, 0x44, 0x65, 0x70,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74,
	0x6f, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x74, 0x6f, 0x70, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x08, 0x20,
	0x03, 0x28, 0x05, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x92, 0x02, 0x0a, 0x08, 0x44, 0x65, 0x70,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x26, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x44, 0x65, 0x70, 0x74, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a,
	0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67,
	0x65, 0x72, 0x12, 0x24, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x70,
	0x74, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x36, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x70, 0x74, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x1a, 0x39, 0x0a, 0x0b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x5b, 0x0a,
	0x0c, 0x44, 0x65, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x26, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x70, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x87, 0x03, 0x0a, 0x0b, 0x44,
	0x65, 0x70, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x44, 0x65, 0x70, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x44, 0x65, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x32,
	0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x70, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x32, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0e, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x70, 0x74, 0x1a, 0x16, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0e, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x70, 0x74, 0x1a, 0x16, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x70, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x04, 0x54, 0x72, 0x65, 0x65, 0x12,
	0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x70, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x44, 0x65, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x37, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x70, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x06, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44,
	0x65, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_deptService_proto_rawDescOnce sync.Once
	file_deptService_proto_rawDescData = file_deptService_proto_rawDesc
)

func file_deptService_proto_rawDescGZIP() []byte {
	file_deptService_proto_rawDescOnce.Do(func() {
		file_deptService_proto_rawDescData = protoimpl.X.CompressGZIP(file_deptService_proto_rawDescData)
	})
	return file_deptService_proto_rawDescData
}

var file_deptService_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_deptService_proto_goTypes = []interface{}{
	(*Dept)(nil),         // 0: services.Dept
	(*DeptRequest)(nil),  // 1: services.DeptRequest
	(*DeptData)(nil),     // 2: services.DeptData
	(*DeptResponse)(nil), // 3: services.DeptResponse
	nil,                  // 4: services.DeptData.ParamsEntry
	(*common.Pager)(nil), // 5: common.Pager
	(*common.Info)(nil),  // 6: common.Info
	(*common.Error)(nil), // 7: common.Error
}
var file_deptService_proto_depIdxs = []int32{
	0,  // 0: services.Dept.children:type_name -> services.Dept
	0,  // 1: services.DeptData.entity:type_name -> services.Dept
	5,  // 2: services.DeptData.pager:type_name -> common.Pager
	0,  // 3: services.DeptData.items:type_name -> services.Dept
	4,  // 4: services.DeptData.params:type_name -> services.DeptData.ParamsEntry
	6,  // 5: services.DeptData.info:type_name -> common.Info
	7,  // 6: services.DeptResponse.error:type_name -> common.Error
	2,  // 7: services.DeptResponse.data:type_name -> services.DeptData
	0,  // 8: services.DeptService.Create:input_type -> services.Dept
	0,  // 9: services.DeptService.Update:input_type -> services.Dept
	0,  // 10: services.DeptService.Delete:input_type -> services.Dept
	0,  // 11: services.DeptService.Get:input_type -> services.Dept
	1,  // 12: services.DeptService.Tree:input_type -> services.DeptRequest
	1,  // 13: services.DeptService.List:input_type -> services.DeptRequest
	1,  // 14: services.DeptService.Search:input_type -> services.DeptRequest
	3,  // 15: services.DeptService.Create:output_type -> services.DeptResponse
	3,  // 16: services.DeptService.Update:output_type -> services.DeptResponse
	3,  // 17: services.DeptService.Delete:output_type -> services.DeptResponse
	3,  // 18: services.DeptService.Get:output_type -> services.DeptResponse
	3,  // 19: services.DeptService.Tree:output_type -> services.DeptResponse
	3,  // 20: services.DeptService.List:output_type -> services.DeptResponse
	3,  // 21: services.DeptService.Search:output_type -> services.DeptResponse
	15, // [15:22] is the sub-list for method output_type
	8,  // [8:15] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_deptService_proto_init() }
func file_deptService_proto_init() {
	if File_deptService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_deptService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dept); i {
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
		file_deptService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeptRequest); i {
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
		file_deptService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeptData); i {
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
		file_deptService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeptResponse); i {
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
			RawDescriptor: file_deptService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_deptService_proto_goTypes,
		DependencyIndexes: file_deptService_proto_depIdxs,
		MessageInfos:      file_deptService_proto_msgTypes,
	}.Build()
	File_deptService_proto = out.File
	file_deptService_proto_rawDesc = nil
	file_deptService_proto_goTypes = nil
	file_deptService_proto_depIdxs = nil
}
