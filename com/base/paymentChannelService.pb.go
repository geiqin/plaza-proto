// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: paymentChannelService.proto

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

// 支付通道
type PaymentChannel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`                                              //ID
	ChannelCode    string `protobuf:"bytes,2,opt,name=channel_code,json=channelCode,proto3" json:"channel_code"`          //唯一标记
	ChannelName    string `protobuf:"bytes,3,opt,name=channel_name,json=channelName,proto3" json:"channel_name"`          //方式名称
	PaymentType    string `protobuf:"bytes,4,opt,name=payment_type,json=paymentType,proto3" json:"payment_type"`          //支付类型
	IconUrl        string `protobuf:"bytes,5,opt,name=icon_url,json=iconUrl,proto3" json:"icon_url"`                      //图标URL
	UniappProvider string `protobuf:"bytes,6,opt,name=uniapp_provider,json=uniappProvider,proto3" json:"uniapp_provider"` //UniApp标识
	Desc           string `protobuf:"bytes,7,opt,name=desc,proto3" json:"desc"`                                           //支付描述
	Sort           int32  `protobuf:"varint,8,opt,name=sort,proto3" json:"sort"`                                          //顺序
	CreatedAt      int64  `protobuf:"varint,9,opt,name=created_at,json=createdAt,proto3" json:"created_at"`               //创建时间
	UpdatedAt      int64  `protobuf:"varint,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`              //修改时间
}

func (x *PaymentChannel) Reset() {
	*x = PaymentChannel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paymentChannelService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentChannel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentChannel) ProtoMessage() {}

func (x *PaymentChannel) ProtoReflect() protoreflect.Message {
	mi := &file_paymentChannelService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentChannel.ProtoReflect.Descriptor instead.
func (*PaymentChannel) Descriptor() ([]byte, []int) {
	return file_paymentChannelService_proto_rawDescGZIP(), []int{0}
}

func (x *PaymentChannel) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PaymentChannel) GetChannelCode() string {
	if x != nil {
		return x.ChannelCode
	}
	return ""
}

func (x *PaymentChannel) GetChannelName() string {
	if x != nil {
		return x.ChannelName
	}
	return ""
}

func (x *PaymentChannel) GetPaymentType() string {
	if x != nil {
		return x.PaymentType
	}
	return ""
}

func (x *PaymentChannel) GetIconUrl() string {
	if x != nil {
		return x.IconUrl
	}
	return ""
}

func (x *PaymentChannel) GetUniappProvider() string {
	if x != nil {
		return x.UniappProvider
	}
	return ""
}

func (x *PaymentChannel) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *PaymentChannel) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *PaymentChannel) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *PaymentChannel) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

// 支付通道请求参数
type PaymentChannelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Top       int32    `protobuf:"varint,1,opt,name=top,proto3" json:"top"`
	Paged     int64    `protobuf:"varint,2,opt,name=paged,proto3" json:"paged"`
	PageSize  int64    `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Keywords  string   `protobuf:"bytes,4,opt,name=keywords,proto3" json:"keywords"`
	Sort      []string `protobuf:"bytes,5,rep,name=sort,proto3" json:"sort"`
	DateRange []string `protobuf:"bytes,6,rep,name=date_range,json=dateRange,proto3" json:"date_range"`
	Ids       []int32  `protobuf:"varint,7,rep,packed,name=ids,proto3" json:"ids"`
	Id        int32    `protobuf:"varint,8,opt,name=id,proto3" json:"id"`
	//以下为自定义参数
	ChannelCode string `protobuf:"bytes,11,opt,name=channel_code,json=channelCode,proto3" json:"channel_code"` //唯一标记
	ChannelName string `protobuf:"bytes,12,opt,name=channel_name,json=channelName,proto3" json:"channel_name"` //方式名称
	PaymentType string `protobuf:"bytes,13,opt,name=payment_type,json=paymentType,proto3" json:"payment_type"` //支付类型
}

func (x *PaymentChannelRequest) Reset() {
	*x = PaymentChannelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paymentChannelService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentChannelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentChannelRequest) ProtoMessage() {}

func (x *PaymentChannelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_paymentChannelService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentChannelRequest.ProtoReflect.Descriptor instead.
func (*PaymentChannelRequest) Descriptor() ([]byte, []int) {
	return file_paymentChannelService_proto_rawDescGZIP(), []int{1}
}

func (x *PaymentChannelRequest) GetTop() int32 {
	if x != nil {
		return x.Top
	}
	return 0
}

func (x *PaymentChannelRequest) GetPaged() int64 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *PaymentChannelRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *PaymentChannelRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *PaymentChannelRequest) GetSort() []string {
	if x != nil {
		return x.Sort
	}
	return nil
}

func (x *PaymentChannelRequest) GetDateRange() []string {
	if x != nil {
		return x.DateRange
	}
	return nil
}

func (x *PaymentChannelRequest) GetIds() []int32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *PaymentChannelRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PaymentChannelRequest) GetChannelCode() string {
	if x != nil {
		return x.ChannelCode
	}
	return ""
}

func (x *PaymentChannelRequest) GetChannelName() string {
	if x != nil {
		return x.ChannelName
	}
	return ""
}

func (x *PaymentChannelRequest) GetPaymentType() string {
	if x != nil {
		return x.PaymentType
	}
	return ""
}

// 支付通道响应数据
type PaymentChannelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *PaymentChannel   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager     `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*PaymentChannel `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Msg    string            `protobuf:"bytes,4,opt,name=msg,proto3" json:"msg"`
}

func (x *PaymentChannelResponse) Reset() {
	*x = PaymentChannelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paymentChannelService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentChannelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentChannelResponse) ProtoMessage() {}

func (x *PaymentChannelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_paymentChannelService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentChannelResponse.ProtoReflect.Descriptor instead.
func (*PaymentChannelResponse) Descriptor() ([]byte, []int) {
	return file_paymentChannelService_proto_rawDescGZIP(), []int{2}
}

func (x *PaymentChannelResponse) GetEntity() *PaymentChannel {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *PaymentChannelResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *PaymentChannelResponse) GetItems() []*PaymentChannel {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *PaymentChannelResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_paymentChannelService_proto protoreflect.FileDescriptor

var file_paymentChannelService_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb3, 0x02, 0x0a, 0x0e, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a,
	0x0c, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x63, 0x6f, 0x6e, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x63, 0x6f, 0x6e, 0x55, 0x72,
	0x6c, 0x12, 0x27, 0x0a, 0x0f, 0x75, 0x6e, 0x69, 0x61, 0x70, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x75, 0x6e, 0x69, 0x61,
	0x70, 0x70, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65,
	0x73, 0x63, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x6f,
	0x72, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x22, 0xb6, 0x02, 0x0a, 0x15, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x6f,
	0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x74, 0x6f, 0x70, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x61, 0x67,
	0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x09, 0x64, 0x61, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x05, 0x52, 0x03, 0x69, 0x64, 0x73,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0xb1, 0x01, 0x0a, 0x16, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x06,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50,
	0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x32, 0xd0, 0x03,
	0x0a, 0x15, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x1a, 0x20, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x46, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x1a, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x1a, 0x20, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x43, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x1a, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1f,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1f, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_paymentChannelService_proto_rawDescOnce sync.Once
	file_paymentChannelService_proto_rawDescData = file_paymentChannelService_proto_rawDesc
)

func file_paymentChannelService_proto_rawDescGZIP() []byte {
	file_paymentChannelService_proto_rawDescOnce.Do(func() {
		file_paymentChannelService_proto_rawDescData = protoimpl.X.CompressGZIP(file_paymentChannelService_proto_rawDescData)
	})
	return file_paymentChannelService_proto_rawDescData
}

var file_paymentChannelService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_paymentChannelService_proto_goTypes = []interface{}{
	(*PaymentChannel)(nil),         // 0: services.PaymentChannel
	(*PaymentChannelRequest)(nil),  // 1: services.PaymentChannelRequest
	(*PaymentChannelResponse)(nil), // 2: services.PaymentChannelResponse
	(*common.Pager)(nil),           // 3: common.Pager
}
var file_paymentChannelService_proto_depIdxs = []int32{
	0, // 0: services.PaymentChannelResponse.entity:type_name -> services.PaymentChannel
	3, // 1: services.PaymentChannelResponse.pager:type_name -> common.Pager
	0, // 2: services.PaymentChannelResponse.items:type_name -> services.PaymentChannel
	0, // 3: services.PaymentChannelService.Create:input_type -> services.PaymentChannel
	0, // 4: services.PaymentChannelService.Update:input_type -> services.PaymentChannel
	0, // 5: services.PaymentChannelService.Delete:input_type -> services.PaymentChannel
	0, // 6: services.PaymentChannelService.Get:input_type -> services.PaymentChannel
	1, // 7: services.PaymentChannelService.Search:input_type -> services.PaymentChannelRequest
	1, // 8: services.PaymentChannelService.List:input_type -> services.PaymentChannelRequest
	2, // 9: services.PaymentChannelService.Create:output_type -> services.PaymentChannelResponse
	2, // 10: services.PaymentChannelService.Update:output_type -> services.PaymentChannelResponse
	2, // 11: services.PaymentChannelService.Delete:output_type -> services.PaymentChannelResponse
	2, // 12: services.PaymentChannelService.Get:output_type -> services.PaymentChannelResponse
	2, // 13: services.PaymentChannelService.Search:output_type -> services.PaymentChannelResponse
	2, // 14: services.PaymentChannelService.List:output_type -> services.PaymentChannelResponse
	9, // [9:15] is the sub-list for method output_type
	3, // [3:9] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_paymentChannelService_proto_init() }
func file_paymentChannelService_proto_init() {
	if File_paymentChannelService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_paymentChannelService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentChannel); i {
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
		file_paymentChannelService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentChannelRequest); i {
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
		file_paymentChannelService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentChannelResponse); i {
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
			RawDescriptor: file_paymentChannelService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_paymentChannelService_proto_goTypes,
		DependencyIndexes: file_paymentChannelService_proto_depIdxs,
		MessageInfos:      file_paymentChannelService_proto_msgTypes,
	}.Build()
	File_paymentChannelService_proto = out.File
	file_paymentChannelService_proto_rawDesc = nil
	file_paymentChannelService_proto_goTypes = nil
	file_paymentChannelService_proto_depIdxs = nil
}
