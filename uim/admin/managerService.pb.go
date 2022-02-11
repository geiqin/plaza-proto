// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: managerService.proto

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

// 用户信息
type Manager struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	Type      string `protobuf:"bytes,3,opt,name=type,proto3" json:"type"`
	AvatarId  int64  `protobuf:"varint,4,opt,name=avatar_id,json=avatarId,proto3" json:"avatar_id"`
	AvatarUrl string `protobuf:"bytes,5,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url"`
	Gender    string `protobuf:"bytes,6,opt,name=gender,proto3" json:"gender"`
	RoleId    int32  `protobuf:"varint,7,opt,name=role_id,json=roleId,proto3" json:"role_id"`
	Status    string `protobuf:"bytes,8,opt,name=status,proto3" json:"status"`
	Memo      string `protobuf:"bytes,9,opt,name=memo,proto3" json:"memo"`
	CreatedAt string `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt string `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	Role      *Role  `protobuf:"bytes,12,opt,name=role,proto3" json:"role"`
}

func (x *Manager) Reset() {
	*x = Manager{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managerService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Manager) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Manager) ProtoMessage() {}

func (x *Manager) ProtoReflect() protoreflect.Message {
	mi := &file_managerService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Manager.ProtoReflect.Descriptor instead.
func (*Manager) Descriptor() ([]byte, []int) {
	return file_managerService_proto_rawDescGZIP(), []int{0}
}

func (x *Manager) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Manager) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Manager) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Manager) GetAvatarId() int64 {
	if x != nil {
		return x.AvatarId
	}
	return 0
}

func (x *Manager) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

func (x *Manager) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *Manager) GetRoleId() int32 {
	if x != nil {
		return x.RoleId
	}
	return 0
}

func (x *Manager) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Manager) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *Manager) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Manager) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Manager) GetRole() *Role {
	if x != nil {
		return x.Role
	}
	return nil
}

// 用户信息
type ManagerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int32  `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize int32  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Sorting  string `protobuf:"bytes,3,opt,name=sorting,proto3" json:"sorting"`
	//----
	Id        int64   `protobuf:"varint,4,opt,name=id,proto3" json:"id"`
	Name      string  `protobuf:"bytes,5,opt,name=name,proto3" json:"name"`
	Username  string  `protobuf:"bytes,6,opt,name=username,proto3" json:"username"`
	AvatarId  int64   `protobuf:"varint,7,opt,name=avatar_id,json=avatarId,proto3" json:"avatar_id"`
	AvatarUrl string  `protobuf:"bytes,8,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url"`
	Email     string  `protobuf:"bytes,9,opt,name=email,proto3" json:"email"`
	Mobile    string  `protobuf:"bytes,10,opt,name=mobile,proto3" json:"mobile"`
	Gender    string  `protobuf:"bytes,11,opt,name=gender,proto3" json:"gender"`
	Status    string  `protobuf:"bytes,12,opt,name=status,proto3" json:"status"`
	Password  string  `protobuf:"bytes,13,opt,name=password,proto3" json:"password"`
	Ids       []int64 `protobuf:"varint,14,rep,packed,name=ids,proto3" json:"ids"`
}

func (x *ManagerRequest) Reset() {
	*x = ManagerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managerService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManagerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManagerRequest) ProtoMessage() {}

func (x *ManagerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_managerService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManagerRequest.ProtoReflect.Descriptor instead.
func (*ManagerRequest) Descriptor() ([]byte, []int) {
	return file_managerService_proto_rawDescGZIP(), []int{1}
}

func (x *ManagerRequest) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *ManagerRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ManagerRequest) GetSorting() string {
	if x != nil {
		return x.Sorting
	}
	return ""
}

func (x *ManagerRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ManagerRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ManagerRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *ManagerRequest) GetAvatarId() int64 {
	if x != nil {
		return x.AvatarId
	}
	return 0
}

func (x *ManagerRequest) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

func (x *ManagerRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ManagerRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *ManagerRequest) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *ManagerRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ManagerRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *ManagerRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type ManagerData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Manager      `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*Manager    `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Info   *common.Info  `protobuf:"bytes,4,opt,name=info,proto3" json:"info"`
}

func (x *ManagerData) Reset() {
	*x = ManagerData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managerService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManagerData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManagerData) ProtoMessage() {}

func (x *ManagerData) ProtoReflect() protoreflect.Message {
	mi := &file_managerService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManagerData.ProtoReflect.Descriptor instead.
func (*ManagerData) Descriptor() ([]byte, []int) {
	return file_managerService_proto_rawDescGZIP(), []int{2}
}

func (x *ManagerData) GetEntity() *Manager {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *ManagerData) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *ManagerData) GetItems() []*Manager {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ManagerData) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type ManagerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *ManagerData  `protobuf:"bytes,1,opt,name=data,proto3" json:"data"`
	Error *common.Error `protobuf:"bytes,2,opt,name=error,proto3" json:"error"`
}

func (x *ManagerResponse) Reset() {
	*x = ManagerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managerService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManagerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManagerResponse) ProtoMessage() {}

func (x *ManagerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_managerService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManagerResponse.ProtoReflect.Descriptor instead.
func (*ManagerResponse) Descriptor() ([]byte, []int) {
	return file_managerService_proto_rawDescGZIP(), []int{3}
}

func (x *ManagerResponse) GetData() *ManagerData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ManagerResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_managerService_proto protoreflect.FileDescriptor

var file_managerService_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x72, 0x6f, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbc, 0x02, 0x0a, 0x07, 0x4d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x17,
	0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d,
	0x65, 0x6d, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x22, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52,
	0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0xe5, 0x02, 0x0a, 0x0e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6f,
	0x72, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x72,
	0x6c, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x69,
	0x64, 0x73, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0xa8, 0x01,
	0x0a, 0x0b, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x29, 0x0a,
	0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x27, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x61, 0x0a, 0x0f, 0x4d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xab, 0x04, 0x0a, 0x0e,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f,
	0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x3f, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x3c, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f,
	0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x3f, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x4c, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x41, 0x6e, 0x64, 0x50, 0x77, 0x64, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44,
	0x0a, 0x0b, 0x47, 0x65, 0x74, 0x42, 0x79, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x18, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x42, 0x79, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_managerService_proto_rawDescOnce sync.Once
	file_managerService_proto_rawDescData = file_managerService_proto_rawDesc
)

func file_managerService_proto_rawDescGZIP() []byte {
	file_managerService_proto_rawDescOnce.Do(func() {
		file_managerService_proto_rawDescData = protoimpl.X.CompressGZIP(file_managerService_proto_rawDescData)
	})
	return file_managerService_proto_rawDescData
}

var file_managerService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_managerService_proto_goTypes = []interface{}{
	(*Manager)(nil),         // 0: services.Manager
	(*ManagerRequest)(nil),  // 1: services.ManagerRequest
	(*ManagerData)(nil),     // 2: services.ManagerData
	(*ManagerResponse)(nil), // 3: services.ManagerResponse
	(*Role)(nil),            // 4: services.Role
	(*common.Pager)(nil),    // 5: common.Pager
	(*common.Info)(nil),     // 6: common.Info
	(*common.Error)(nil),    // 7: common.Error
}
var file_managerService_proto_depIdxs = []int32{
	4,  // 0: services.Manager.role:type_name -> services.Role
	0,  // 1: services.ManagerData.entity:type_name -> services.Manager
	5,  // 2: services.ManagerData.pager:type_name -> common.Pager
	0,  // 3: services.ManagerData.items:type_name -> services.Manager
	6,  // 4: services.ManagerData.info:type_name -> common.Info
	2,  // 5: services.ManagerResponse.data:type_name -> services.ManagerData
	7,  // 6: services.ManagerResponse.error:type_name -> common.Error
	1,  // 7: services.ManagerService.Create:input_type -> services.ManagerRequest
	1,  // 8: services.ManagerService.Update:input_type -> services.ManagerRequest
	1,  // 9: services.ManagerService.Get:input_type -> services.ManagerRequest
	1,  // 10: services.ManagerService.Delete:input_type -> services.ManagerRequest
	1,  // 11: services.ManagerService.Search:input_type -> services.ManagerRequest
	1,  // 12: services.ManagerService.GetByUsernameAndPwd:input_type -> services.ManagerRequest
	1,  // 13: services.ManagerService.GetByMobile:input_type -> services.ManagerRequest
	1,  // 14: services.ManagerService.GetByEmail:input_type -> services.ManagerRequest
	3,  // 15: services.ManagerService.Create:output_type -> services.ManagerResponse
	3,  // 16: services.ManagerService.Update:output_type -> services.ManagerResponse
	3,  // 17: services.ManagerService.Get:output_type -> services.ManagerResponse
	3,  // 18: services.ManagerService.Delete:output_type -> services.ManagerResponse
	3,  // 19: services.ManagerService.Search:output_type -> services.ManagerResponse
	3,  // 20: services.ManagerService.GetByUsernameAndPwd:output_type -> services.ManagerResponse
	3,  // 21: services.ManagerService.GetByMobile:output_type -> services.ManagerResponse
	3,  // 22: services.ManagerService.GetByEmail:output_type -> services.ManagerResponse
	15, // [15:23] is the sub-list for method output_type
	7,  // [7:15] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_managerService_proto_init() }
func file_managerService_proto_init() {
	if File_managerService_proto != nil {
		return
	}
	file_roleService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_managerService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Manager); i {
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
		file_managerService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManagerRequest); i {
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
		file_managerService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManagerData); i {
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
		file_managerService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManagerResponse); i {
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
			RawDescriptor: file_managerService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_managerService_proto_goTypes,
		DependencyIndexes: file_managerService_proto_depIdxs,
		MessageInfos:      file_managerService_proto_msgTypes,
	}.Build()
	File_managerService_proto = out.File
	file_managerService_proto_rawDesc = nil
	file_managerService_proto_goTypes = nil
	file_managerService_proto_depIdxs = nil
}