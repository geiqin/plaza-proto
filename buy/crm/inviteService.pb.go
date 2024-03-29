// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: inviteService.proto

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

type Invite struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	FanId    int64  `protobuf:"varint,2,opt,name=fan_id,json=fanId,proto3" json:"fan_id"`
	MemberId int64  `protobuf:"varint,3,opt,name=member_id,json=memberId,proto3" json:"member_id"`
	LinkUrl  string `protobuf:"bytes,4,opt,name=link_url,json=linkUrl,proto3" json:"link_url"`
	Qrcode   string `protobuf:"bytes,5,opt,name=qrcode,proto3" json:"qrcode"`
	Size     int32  `protobuf:"varint,6,opt,name=size,proto3" json:"size"`
	Type     string `protobuf:"bytes,7,opt,name=type,proto3" json:"type"`
}

func (x *Invite) Reset() {
	*x = Invite{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inviteService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Invite) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Invite) ProtoMessage() {}

func (x *Invite) ProtoReflect() protoreflect.Message {
	mi := &file_inviteService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Invite.ProtoReflect.Descriptor instead.
func (*Invite) Descriptor() ([]byte, []int) {
	return file_inviteService_proto_rawDescGZIP(), []int{0}
}

func (x *Invite) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Invite) GetFanId() int64 {
	if x != nil {
		return x.FanId
	}
	return 0
}

func (x *Invite) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

func (x *Invite) GetLinkUrl() string {
	if x != nil {
		return x.LinkUrl
	}
	return ""
}

func (x *Invite) GetQrcode() string {
	if x != nil {
		return x.Qrcode
	}
	return ""
}

func (x *Invite) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Invite) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type InviteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int64  `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize int64  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Sorting  string `protobuf:"bytes,3,opt,name=sorting,proto3" json:"sorting"`
	//以下为自定义参数
	Id       int64  `protobuf:"varint,4,opt,name=id,proto3" json:"id"`
	FanId    int64  `protobuf:"varint,5,opt,name=fan_id,json=fanId,proto3" json:"fan_id"`
	MemberId int64  `protobuf:"varint,6,opt,name=member_id,json=memberId,proto3" json:"member_id"`
	Code     string `protobuf:"bytes,7,opt,name=code,proto3" json:"code"`
	Type     string `protobuf:"bytes,8,opt,name=type,proto3" json:"type"`
	Mode     int32  `protobuf:"varint,9,opt,name=mode,proto3" json:"mode"`
}

func (x *InviteRequest) Reset() {
	*x = InviteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inviteService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InviteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InviteRequest) ProtoMessage() {}

func (x *InviteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_inviteService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InviteRequest.ProtoReflect.Descriptor instead.
func (*InviteRequest) Descriptor() ([]byte, []int) {
	return file_inviteService_proto_rawDescGZIP(), []int{1}
}

func (x *InviteRequest) GetPaged() int64 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *InviteRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *InviteRequest) GetSorting() string {
	if x != nil {
		return x.Sorting
	}
	return ""
}

func (x *InviteRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *InviteRequest) GetFanId() int64 {
	if x != nil {
		return x.FanId
	}
	return 0
}

func (x *InviteRequest) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

func (x *InviteRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *InviteRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *InviteRequest) GetMode() int32 {
	if x != nil {
		return x.Mode
	}
	return 0
}

type InviteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Invite           `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager     `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*Invite         `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Info   string            `protobuf:"bytes,4,opt,name=info,proto3" json:"info"`
	Params map[string]string `protobuf:"bytes,5,rep,name=params,proto3" json:"params" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *InviteResponse) Reset() {
	*x = InviteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inviteService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InviteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InviteResponse) ProtoMessage() {}

func (x *InviteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_inviteService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InviteResponse.ProtoReflect.Descriptor instead.
func (*InviteResponse) Descriptor() ([]byte, []int) {
	return file_inviteService_proto_rawDescGZIP(), []int{2}
}

func (x *InviteResponse) GetEntity() *Invite {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *InviteResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *InviteResponse) GetItems() []*Invite {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *InviteResponse) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

func (x *InviteResponse) GetParams() map[string]string {
	if x != nil {
		return x.Params
	}
	return nil
}

var File_inviteService_proto protoreflect.FileDescriptor

var file_inviteService_proto_rawDesc = []byte{
	0x0a, 0x13, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a,
	0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xa7, 0x01, 0x0a, 0x06, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x15, 0x0a,
	0x06, 0x66, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x66,
	0x61, 0x6e, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x69, 0x6e, 0x6b, 0x55, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06,
	0x71, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x71, 0x72,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0xdc, 0x01, 0x0a,
	0x0d, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70,
	0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x66,
	0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x66, 0x61, 0x6e,
	0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x22, 0x94, 0x02, 0x0a, 0x0e,
	0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28,
	0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65,
	0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x26, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x3c, 0x0a, 0x06, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x32, 0xd1, 0x01, 0x0a, 0x0d, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x4d, 0x69, 0x6e, 0x69, 0x51, 0x72, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e,
	0x76, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x09, 0x41, 0x70, 0x70, 0x51, 0x72,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x04, 0x4c, 0x69, 0x6e,
	0x6b, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x76,
	0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_inviteService_proto_rawDescOnce sync.Once
	file_inviteService_proto_rawDescData = file_inviteService_proto_rawDesc
)

func file_inviteService_proto_rawDescGZIP() []byte {
	file_inviteService_proto_rawDescOnce.Do(func() {
		file_inviteService_proto_rawDescData = protoimpl.X.CompressGZIP(file_inviteService_proto_rawDescData)
	})
	return file_inviteService_proto_rawDescData
}

var file_inviteService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_inviteService_proto_goTypes = []interface{}{
	(*Invite)(nil),         // 0: services.Invite
	(*InviteRequest)(nil),  // 1: services.InviteRequest
	(*InviteResponse)(nil), // 2: services.InviteResponse
	nil,                    // 3: services.InviteResponse.ParamsEntry
	(*common.Pager)(nil),   // 4: common.Pager
}
var file_inviteService_proto_depIdxs = []int32{
	0, // 0: services.InviteResponse.entity:type_name -> services.Invite
	4, // 1: services.InviteResponse.pager:type_name -> common.Pager
	0, // 2: services.InviteResponse.items:type_name -> services.Invite
	3, // 3: services.InviteResponse.params:type_name -> services.InviteResponse.ParamsEntry
	1, // 4: services.InviteService.MiniQrcode:input_type -> services.InviteRequest
	1, // 5: services.InviteService.AppQrcode:input_type -> services.InviteRequest
	1, // 6: services.InviteService.Link:input_type -> services.InviteRequest
	2, // 7: services.InviteService.MiniQrcode:output_type -> services.InviteResponse
	2, // 8: services.InviteService.AppQrcode:output_type -> services.InviteResponse
	2, // 9: services.InviteService.Link:output_type -> services.InviteResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_inviteService_proto_init() }
func file_inviteService_proto_init() {
	if File_inviteService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_inviteService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Invite); i {
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
		file_inviteService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InviteRequest); i {
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
		file_inviteService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InviteResponse); i {
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
			RawDescriptor: file_inviteService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_inviteService_proto_goTypes,
		DependencyIndexes: file_inviteService_proto_depIdxs,
		MessageInfos:      file_inviteService_proto_msgTypes,
	}.Build()
	File_inviteService_proto = out.File
	file_inviteService_proto_rawDesc = nil
	file_inviteService_proto_goTypes = nil
	file_inviteService_proto_depIdxs = nil
}
