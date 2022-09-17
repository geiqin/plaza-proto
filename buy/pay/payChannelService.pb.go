// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: payChannelService.proto

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

type PayChannel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int32       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	PayAccountId int32       `protobuf:"varint,2,opt,name=pay_account_id,json=payAccountId,proto3" json:"pay_account_id,omitempty"`
	ChannelId    int32       `protobuf:"varint,3,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	ChannelType  string      `protobuf:"bytes,4,opt,name=channel_type,json=channelType,proto3" json:"channel_type,omitempty"`
	AppId        string      `protobuf:"bytes,5,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	ConfigData   string      `protobuf:"bytes,6,opt,name=config_data,json=configData,proto3" json:"config_data,omitempty"`
	Status       string      `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
	CreatedAt    string      `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt    string      `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	PayAccount   *PayAccount `protobuf:"bytes,10,opt,name=pay_account,json=payAccount,proto3" json:"pay_account,omitempty"`
}

func (x *PayChannel) Reset() {
	*x = PayChannel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payChannelService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayChannel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayChannel) ProtoMessage() {}

func (x *PayChannel) ProtoReflect() protoreflect.Message {
	mi := &file_payChannelService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayChannel.ProtoReflect.Descriptor instead.
func (*PayChannel) Descriptor() ([]byte, []int) {
	return file_payChannelService_proto_rawDescGZIP(), []int{0}
}

func (x *PayChannel) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PayChannel) GetPayAccountId() int32 {
	if x != nil {
		return x.PayAccountId
	}
	return 0
}

func (x *PayChannel) GetChannelId() int32 {
	if x != nil {
		return x.ChannelId
	}
	return 0
}

func (x *PayChannel) GetChannelType() string {
	if x != nil {
		return x.ChannelType
	}
	return ""
}

func (x *PayChannel) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *PayChannel) GetConfigData() string {
	if x != nil {
		return x.ConfigData
	}
	return ""
}

func (x *PayChannel) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *PayChannel) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *PayChannel) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *PayChannel) GetPayAccount() *PayAccount {
	if x != nil {
		return x.PayAccount
	}
	return nil
}

type PayChannelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged        int32  `protobuf:"varint,1,opt,name=paged,proto3" json:"paged,omitempty"`
	PageSize     int32  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Keywords     string `protobuf:"bytes,3,opt,name=keywords,proto3" json:"keywords,omitempty"`
	Id           int32  `protobuf:"varint,5,opt,name=id,proto3" json:"id,omitempty"`
	PayAccountId int32  `protobuf:"varint,6,opt,name=pay_account_id,json=payAccountId,proto3" json:"pay_account_id,omitempty"`
	ChannelId    int32  `protobuf:"varint,7,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	ChannelType  string `protobuf:"bytes,8,opt,name=channel_type,json=channelType,proto3" json:"channel_type,omitempty"`
	AppId        string `protobuf:"bytes,9,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	Status       string `protobuf:"bytes,10,opt,name=status,proto3" json:"status,omitempty"`
	Reason       string `protobuf:"bytes,11,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (x *PayChannelRequest) Reset() {
	*x = PayChannelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payChannelService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayChannelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayChannelRequest) ProtoMessage() {}

func (x *PayChannelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payChannelService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayChannelRequest.ProtoReflect.Descriptor instead.
func (*PayChannelRequest) Descriptor() ([]byte, []int) {
	return file_payChannelService_proto_rawDescGZIP(), []int{1}
}

func (x *PayChannelRequest) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *PayChannelRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *PayChannelRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *PayChannelRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PayChannelRequest) GetPayAccountId() int32 {
	if x != nil {
		return x.PayAccountId
	}
	return 0
}

func (x *PayChannelRequest) GetChannelId() int32 {
	if x != nil {
		return x.ChannelId
	}
	return 0
}

func (x *PayChannelRequest) GetChannelType() string {
	if x != nil {
		return x.ChannelType
	}
	return ""
}

func (x *PayChannelRequest) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *PayChannelRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *PayChannelRequest) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

type PayChannelData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *PayChannel   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items  []*PayChannel `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Info   *common.Info  `protobuf:"bytes,4,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *PayChannelData) Reset() {
	*x = PayChannelData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payChannelService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayChannelData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayChannelData) ProtoMessage() {}

func (x *PayChannelData) ProtoReflect() protoreflect.Message {
	mi := &file_payChannelService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayChannelData.ProtoReflect.Descriptor instead.
func (*PayChannelData) Descriptor() ([]byte, []int) {
	return file_payChannelService_proto_rawDescGZIP(), []int{2}
}

func (x *PayChannelData) GetEntity() *PayChannel {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *PayChannelData) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *PayChannelData) GetItems() []*PayChannel {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *PayChannelData) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type PayChannelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *PayChannelData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Error *common.Error   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *PayChannelResponse) Reset() {
	*x = PayChannelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payChannelService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayChannelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayChannelResponse) ProtoMessage() {}

func (x *PayChannelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payChannelService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayChannelResponse.ProtoReflect.Descriptor instead.
func (*PayChannelResponse) Descriptor() ([]byte, []int) {
	return file_payChannelService_proto_rawDescGZIP(), []int{3}
}

func (x *PayChannelResponse) GetData() *PayChannelData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *PayChannelResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_payChannelService_proto protoreflect.FileDescriptor

var file_payChannelService_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x61, 0x79, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x70, 0x61, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xc9, 0x02, 0x0a, 0x0a, 0x50, 0x61, 0x79, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x24,
	0x0a, 0x0e, 0x70, 0x61, 0x79, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x70, 0x61, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x44, 0x61, 0x74, 0x61, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x35, 0x0a, 0x0b, 0x70, 0x61, 0x79, 0x5f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x0a, 0x70, 0x61, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xa1, 0x02, 0x0a, 0x11,
	0x50, 0x61, 0x79, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x24, 0x0a, 0x0e, 0x70, 0x61, 0x79, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x70, 0x61, 0x79, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f,
	0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22,
	0xb1, 0x01, 0x0a, 0x0e, 0x50, 0x61, 0x79, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x2c, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x61,
	0x79, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05,
	0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x50, 0x61, 0x79, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69,
	0x6e, 0x66, 0x6f, 0x22, 0x67, 0x0a, 0x12, 0x50, 0x61, 0x79, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x9d, 0x03, 0x0a,
	0x11, 0x50, 0x61, 0x79, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x3c, 0x0a, 0x04, 0x42, 0x69, 0x6e, 0x64, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x3e, 0x0a, 0x06, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x41, 0x0a, 0x09, 0x53, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50,
	0x61, 0x79, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x43, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x50, 0x61, 0x79, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12,
	0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x43, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b,
	0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_payChannelService_proto_rawDescOnce sync.Once
	file_payChannelService_proto_rawDescData = file_payChannelService_proto_rawDesc
)

func file_payChannelService_proto_rawDescGZIP() []byte {
	file_payChannelService_proto_rawDescOnce.Do(func() {
		file_payChannelService_proto_rawDescData = protoimpl.X.CompressGZIP(file_payChannelService_proto_rawDescData)
	})
	return file_payChannelService_proto_rawDescData
}

var file_payChannelService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_payChannelService_proto_goTypes = []interface{}{
	(*PayChannel)(nil),         // 0: services.PayChannel
	(*PayChannelRequest)(nil),  // 1: services.PayChannelRequest
	(*PayChannelData)(nil),     // 2: services.PayChannelData
	(*PayChannelResponse)(nil), // 3: services.PayChannelResponse
	(*PayAccount)(nil),         // 4: services.PayAccount
	(*common.Pager)(nil),       // 5: common.Pager
	(*common.Info)(nil),        // 6: common.Info
	(*common.Error)(nil),       // 7: common.Error
}
var file_payChannelService_proto_depIdxs = []int32{
	4,  // 0: services.PayChannel.pay_account:type_name -> services.PayAccount
	0,  // 1: services.PayChannelData.entity:type_name -> services.PayChannel
	5,  // 2: services.PayChannelData.pager:type_name -> common.Pager
	0,  // 3: services.PayChannelData.items:type_name -> services.PayChannel
	6,  // 4: services.PayChannelData.info:type_name -> common.Info
	2,  // 5: services.PayChannelResponse.data:type_name -> services.PayChannelData
	7,  // 6: services.PayChannelResponse.error:type_name -> common.Error
	0,  // 7: services.PayChannelService.Bind:input_type -> services.PayChannel
	0,  // 8: services.PayChannelService.Cancel:input_type -> services.PayChannel
	0,  // 9: services.PayChannelService.SetStatus:input_type -> services.PayChannel
	0,  // 10: services.PayChannelService.Get:input_type -> services.PayChannel
	1,  // 11: services.PayChannelService.List:input_type -> services.PayChannelRequest
	1,  // 12: services.PayChannelService.Search:input_type -> services.PayChannelRequest
	3,  // 13: services.PayChannelService.Bind:output_type -> services.PayChannelResponse
	3,  // 14: services.PayChannelService.Cancel:output_type -> services.PayChannelResponse
	3,  // 15: services.PayChannelService.SetStatus:output_type -> services.PayChannelResponse
	3,  // 16: services.PayChannelService.Get:output_type -> services.PayChannelResponse
	3,  // 17: services.PayChannelService.List:output_type -> services.PayChannelResponse
	3,  // 18: services.PayChannelService.Search:output_type -> services.PayChannelResponse
	13, // [13:19] is the sub-list for method output_type
	7,  // [7:13] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_payChannelService_proto_init() }
func file_payChannelService_proto_init() {
	if File_payChannelService_proto != nil {
		return
	}
	file_payAccountService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_payChannelService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayChannel); i {
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
		file_payChannelService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayChannelRequest); i {
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
		file_payChannelService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayChannelData); i {
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
		file_payChannelService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayChannelResponse); i {
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
			RawDescriptor: file_payChannelService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_payChannelService_proto_goTypes,
		DependencyIndexes: file_payChannelService_proto_depIdxs,
		MessageInfos:      file_payChannelService_proto_msgTypes,
	}.Build()
	File_payChannelService_proto = out.File
	file_payChannelService_proto_rawDesc = nil
	file_payChannelService_proto_goTypes = nil
	file_payChannelService_proto_depIdxs = nil
}
