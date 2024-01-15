// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: paymentWayService.proto

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

type PaymentWay struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`                                                 //ID
	WayCode         string `protobuf:"bytes,2,opt,name=way_code,json=wayCode,proto3" json:"way_code"`                         //唯一标记
	WayName         string `protobuf:"bytes,3,opt,name=way_name,json=wayName,proto3" json:"way_name"`                         //支付方式名称
	PaymentType     string `protobuf:"bytes,4,opt,name=payment_type,json=paymentType,proto3" json:"payment_type"`             //支付类型
	IconUrl         string `protobuf:"bytes,5,opt,name=icon_url,json=iconUrl,proto3" json:"icon_url"`                         //图标URL
	Desc            string `protobuf:"bytes,6,opt,name=desc,proto3" json:"desc"`                                              //支付描述
	UniAppProvider  int32  `protobuf:"varint,7,opt,name=uni_app_provider,json=uniAppProvider,proto3" json:"uni_app_provider"` //UniApp支付标识
	Sort            int32  `protobuf:"varint,8,opt,name=sort,proto3" json:"sort"`                                             //顺序
	CreatedAt       string `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt       string `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	PaymentTypeName string `protobuf:"bytes,11,opt,name=payment_type_name,json=paymentTypeName,proto3" json:"payment_type_name"`
}

func (x *PaymentWay) Reset() {
	*x = PaymentWay{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paymentWayService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentWay) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentWay) ProtoMessage() {}

func (x *PaymentWay) ProtoReflect() protoreflect.Message {
	mi := &file_paymentWayService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentWay.ProtoReflect.Descriptor instead.
func (*PaymentWay) Descriptor() ([]byte, []int) {
	return file_paymentWayService_proto_rawDescGZIP(), []int{0}
}

func (x *PaymentWay) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PaymentWay) GetWayCode() string {
	if x != nil {
		return x.WayCode
	}
	return ""
}

func (x *PaymentWay) GetWayName() string {
	if x != nil {
		return x.WayName
	}
	return ""
}

func (x *PaymentWay) GetPaymentType() string {
	if x != nil {
		return x.PaymentType
	}
	return ""
}

func (x *PaymentWay) GetIconUrl() string {
	if x != nil {
		return x.IconUrl
	}
	return ""
}

func (x *PaymentWay) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *PaymentWay) GetUniAppProvider() int32 {
	if x != nil {
		return x.UniAppProvider
	}
	return 0
}

func (x *PaymentWay) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *PaymentWay) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *PaymentWay) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *PaymentWay) GetPaymentTypeName() string {
	if x != nil {
		return x.PaymentTypeName
	}
	return ""
}

type PaymentWayRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged       int64  `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize    int64  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Keywords    string `protobuf:"bytes,3,opt,name=keywords,proto3" json:"keywords"`
	Id          int32  `protobuf:"varint,5,opt,name=id,proto3" json:"id"`
	WayCode     int32  `protobuf:"varint,6,opt,name=way_code,json=wayCode,proto3" json:"way_code"`
	WayName     string `protobuf:"bytes,7,opt,name=way_name,json=wayName,proto3" json:"way_name"`
	PaymentType string `protobuf:"bytes,8,opt,name=payment_type,json=paymentType,proto3" json:"payment_type"`
}

func (x *PaymentWayRequest) Reset() {
	*x = PaymentWayRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paymentWayService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentWayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentWayRequest) ProtoMessage() {}

func (x *PaymentWayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_paymentWayService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentWayRequest.ProtoReflect.Descriptor instead.
func (*PaymentWayRequest) Descriptor() ([]byte, []int) {
	return file_paymentWayService_proto_rawDescGZIP(), []int{1}
}

func (x *PaymentWayRequest) GetPaged() int64 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *PaymentWayRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *PaymentWayRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *PaymentWayRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PaymentWayRequest) GetWayCode() int32 {
	if x != nil {
		return x.WayCode
	}
	return 0
}

func (x *PaymentWayRequest) GetWayName() string {
	if x != nil {
		return x.WayName
	}
	return ""
}

func (x *PaymentWayRequest) GetPaymentType() string {
	if x != nil {
		return x.PaymentType
	}
	return ""
}

type PaymentWayData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *PaymentWay   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*PaymentWay `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Msg    string        `protobuf:"bytes,4,opt,name=msg,proto3" json:"msg"`
}

func (x *PaymentWayData) Reset() {
	*x = PaymentWayData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paymentWayService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentWayData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentWayData) ProtoMessage() {}

func (x *PaymentWayData) ProtoReflect() protoreflect.Message {
	mi := &file_paymentWayService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentWayData.ProtoReflect.Descriptor instead.
func (*PaymentWayData) Descriptor() ([]byte, []int) {
	return file_paymentWayService_proto_rawDescGZIP(), []int{2}
}

func (x *PaymentWayData) GetEntity() *PaymentWay {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *PaymentWayData) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *PaymentWayData) GetItems() []*PaymentWay {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *PaymentWayData) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type PaymentWayResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *PaymentWayData `protobuf:"bytes,1,opt,name=data,proto3" json:"data"`
	Error *common.Error   `protobuf:"bytes,2,opt,name=error,proto3" json:"error"`
}

func (x *PaymentWayResponse) Reset() {
	*x = PaymentWayResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paymentWayService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentWayResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentWayResponse) ProtoMessage() {}

func (x *PaymentWayResponse) ProtoReflect() protoreflect.Message {
	mi := &file_paymentWayService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentWayResponse.ProtoReflect.Descriptor instead.
func (*PaymentWayResponse) Descriptor() ([]byte, []int) {
	return file_paymentWayService_proto_rawDescGZIP(), []int{3}
}

func (x *PaymentWayResponse) GetData() *PaymentWayData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *PaymentWayResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_paymentWayService_proto protoreflect.FileDescriptor

var file_paymentWayService_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x57, 0x61, 0x79, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcc, 0x02, 0x0a, 0x0a, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x57, 0x61, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x77, 0x61, 0x79, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x77, 0x61, 0x79, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x77, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x77, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x69, 0x63, 0x6f, 0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x69, 0x63, 0x6f, 0x6e, 0x55, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73,
	0x63, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x28, 0x0a,
	0x10, 0x75, 0x6e, 0x69, 0x5f, 0x61, 0x70, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x75, 0x6e, 0x69, 0x41, 0x70, 0x70, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xcb, 0x01, 0x0a, 0x11, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x57, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x77, 0x61,
	0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x77, 0x61,
	0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x77, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x77, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x22, 0xa1, 0x01, 0x0a, 0x0e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x57,
	0x61, 0x79, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2c, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x57, 0x61, 0x79, 0x52, 0x06, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67,
	0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x57, 0x61, 0x79, 0x52, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x67, 0x0a, 0x12, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x57, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x57, 0x61,
	0x79, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x32, 0x97, 0x02, 0x0a, 0x11, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x57, 0x61, 0x79, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x57, 0x61, 0x79, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x57, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x57, 0x61, 0x79, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x57, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x14, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x57, 0x61, 0x79, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x57, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1b, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x57, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x57, 0x61, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e,
	0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_paymentWayService_proto_rawDescOnce sync.Once
	file_paymentWayService_proto_rawDescData = file_paymentWayService_proto_rawDesc
)

func file_paymentWayService_proto_rawDescGZIP() []byte {
	file_paymentWayService_proto_rawDescOnce.Do(func() {
		file_paymentWayService_proto_rawDescData = protoimpl.X.CompressGZIP(file_paymentWayService_proto_rawDescData)
	})
	return file_paymentWayService_proto_rawDescData
}

var file_paymentWayService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_paymentWayService_proto_goTypes = []interface{}{
	(*PaymentWay)(nil),         // 0: services.PaymentWay
	(*PaymentWayRequest)(nil),  // 1: services.PaymentWayRequest
	(*PaymentWayData)(nil),     // 2: services.PaymentWayData
	(*PaymentWayResponse)(nil), // 3: services.PaymentWayResponse
	(*common.Pager)(nil),       // 4: common.Pager
	(*common.Error)(nil),       // 5: common.Error
}
var file_paymentWayService_proto_depIdxs = []int32{
	0, // 0: services.PaymentWayData.entity:type_name -> services.PaymentWay
	4, // 1: services.PaymentWayData.pager:type_name -> common.Pager
	0, // 2: services.PaymentWayData.items:type_name -> services.PaymentWay
	2, // 3: services.PaymentWayResponse.data:type_name -> services.PaymentWayData
	5, // 4: services.PaymentWayResponse.error:type_name -> common.Error
	0, // 5: services.PaymentWayService.Create:input_type -> services.PaymentWay
	0, // 6: services.PaymentWayService.Delete:input_type -> services.PaymentWay
	0, // 7: services.PaymentWayService.Get:input_type -> services.PaymentWay
	1, // 8: services.PaymentWayService.Search:input_type -> services.PaymentWayRequest
	3, // 9: services.PaymentWayService.Create:output_type -> services.PaymentWayResponse
	3, // 10: services.PaymentWayService.Delete:output_type -> services.PaymentWayResponse
	3, // 11: services.PaymentWayService.Get:output_type -> services.PaymentWayResponse
	3, // 12: services.PaymentWayService.Search:output_type -> services.PaymentWayResponse
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_paymentWayService_proto_init() }
func file_paymentWayService_proto_init() {
	if File_paymentWayService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_paymentWayService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentWay); i {
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
		file_paymentWayService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentWayRequest); i {
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
		file_paymentWayService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentWayData); i {
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
		file_paymentWayService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentWayResponse); i {
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
			RawDescriptor: file_paymentWayService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_paymentWayService_proto_goTypes,
		DependencyIndexes: file_paymentWayService_proto_depIdxs,
		MessageInfos:      file_paymentWayService_proto_msgTypes,
	}.Build()
	File_paymentWayService_proto = out.File
	file_paymentWayService_proto_rawDesc = nil
	file_paymentWayService_proto_goTypes = nil
	file_paymentWayService_proto_depIdxs = nil
}
