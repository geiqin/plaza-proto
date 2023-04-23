// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: paymentService.proto

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

type Payment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id"`                                           //ID
	Code          string   `protobuf:"bytes,2,opt,name=code,proto3" json:"code"`                                        //唯一标记
	Name          string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`                                        //支付名称
	IconId        int64    `protobuf:"varint,4,opt,name=icon_id,json=iconId,proto3" json:"icon_id"`                     //图标ID
	IconUrl       string   `protobuf:"bytes,5,opt,name=icon_url,json=iconUrl,proto3" json:"icon_url"`                   //图标URL
	IsCod         string   `protobuf:"bytes,6,opt,name=is_cod,json=isCod,proto3" json:"is_cod"`                         //是否货到付款（0否，1是）
	IsUnderLine   string   `protobuf:"bytes,7,opt,name=is_under_line,json=isUnderLine,proto3" json:"is_under_line"`     //是否为线下支付（0否，1是）
	ApplyTerminal []string `protobuf:"bytes,8,rep,name=apply_terminal,json=applyTerminal,proto3" json:"apply_terminal"` //适用终端 一维数组json字符串存储（pc, h5, ios, android, alipay, weixin, baidu, toutiao, qq）
	AccountType   string   `protobuf:"bytes,9,opt,name=account_type,json=accountType,proto3" json:"account_type"`       //支付账户类型
	Desc          string   `protobuf:"bytes,10,opt,name=desc,proto3" json:"desc"`                                       //描述
	Sort          int32    `protobuf:"varint,11,opt,name=sort,proto3" json:"sort"`                                      //显示顺序
	Status        string   `protobuf:"bytes,12,opt,name=status,proto3" json:"status"`                                   //状态: 1启用，0 禁用
	CreatedAt     string   `protobuf:"bytes,13,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt     string   `protobuf:"bytes,14,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
}

func (x *Payment) Reset() {
	*x = Payment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paymentService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Payment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Payment) ProtoMessage() {}

func (x *Payment) ProtoReflect() protoreflect.Message {
	mi := &file_paymentService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Payment.ProtoReflect.Descriptor instead.
func (*Payment) Descriptor() ([]byte, []int) {
	return file_paymentService_proto_rawDescGZIP(), []int{0}
}

func (x *Payment) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Payment) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Payment) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Payment) GetIconId() int64 {
	if x != nil {
		return x.IconId
	}
	return 0
}

func (x *Payment) GetIconUrl() string {
	if x != nil {
		return x.IconUrl
	}
	return ""
}

func (x *Payment) GetIsCod() string {
	if x != nil {
		return x.IsCod
	}
	return ""
}

func (x *Payment) GetIsUnderLine() string {
	if x != nil {
		return x.IsUnderLine
	}
	return ""
}

func (x *Payment) GetApplyTerminal() []string {
	if x != nil {
		return x.ApplyTerminal
	}
	return nil
}

func (x *Payment) GetAccountType() string {
	if x != nil {
		return x.AccountType
	}
	return ""
}

func (x *Payment) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *Payment) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *Payment) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Payment) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Payment) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

//支付方式请求
type PaymentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int32 `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Top      int32 `protobuf:"varint,3,opt,name=top,proto3" json:"top"`
	//base params
	Id            int32  `protobuf:"varint,5,opt,name=id,proto3" json:"id"`
	Code          string `protobuf:"bytes,6,opt,name=code,proto3" json:"code"`
	Name          string `protobuf:"bytes,7,opt,name=name,proto3" json:"name"`
	ApplyTerminal string `protobuf:"bytes,8,opt,name=apply_terminal,json=applyTerminal,proto3" json:"apply_terminal"`
	Status        string `protobuf:"bytes,9,opt,name=status,proto3" json:"status"`
	IsCod         string `protobuf:"bytes,10,opt,name=is_cod,json=isCod,proto3" json:"is_cod"`
	IsUnderLine   string `protobuf:"bytes,11,opt,name=is_under_line,json=isUnderLine,proto3" json:"is_under_line"`
	AccountType   string `protobuf:"bytes,12,opt,name=account_type,json=accountType,proto3" json:"account_type"`
}

func (x *PaymentRequest) Reset() {
	*x = PaymentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paymentService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentRequest) ProtoMessage() {}

func (x *PaymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_paymentService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentRequest.ProtoReflect.Descriptor instead.
func (*PaymentRequest) Descriptor() ([]byte, []int) {
	return file_paymentService_proto_rawDescGZIP(), []int{1}
}

func (x *PaymentRequest) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *PaymentRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *PaymentRequest) GetTop() int32 {
	if x != nil {
		return x.Top
	}
	return 0
}

func (x *PaymentRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PaymentRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *PaymentRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PaymentRequest) GetApplyTerminal() string {
	if x != nil {
		return x.ApplyTerminal
	}
	return ""
}

func (x *PaymentRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *PaymentRequest) GetIsCod() string {
	if x != nil {
		return x.IsCod
	}
	return ""
}

func (x *PaymentRequest) GetIsUnderLine() string {
	if x != nil {
		return x.IsUnderLine
	}
	return ""
}

func (x *PaymentRequest) GetAccountType() string {
	if x != nil {
		return x.AccountType
	}
	return ""
}

type PaymentData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Payment      `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Items  []*Payment    `protobuf:"bytes,2,rep,name=items,proto3" json:"items"`
	Pager  *common.Pager `protobuf:"bytes,3,opt,name=pager,proto3" json:"pager"`
	Info   *common.Info  `protobuf:"bytes,4,opt,name=info,proto3" json:"info"`
}

func (x *PaymentData) Reset() {
	*x = PaymentData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paymentService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentData) ProtoMessage() {}

func (x *PaymentData) ProtoReflect() protoreflect.Message {
	mi := &file_paymentService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentData.ProtoReflect.Descriptor instead.
func (*PaymentData) Descriptor() ([]byte, []int) {
	return file_paymentService_proto_rawDescGZIP(), []int{2}
}

func (x *PaymentData) GetEntity() *Payment {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *PaymentData) GetItems() []*Payment {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *PaymentData) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *PaymentData) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type PaymentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *PaymentData  `protobuf:"bytes,1,opt,name=data,proto3" json:"data"`
	Error *common.Error `protobuf:"bytes,2,opt,name=error,proto3" json:"error"`
}

func (x *PaymentResponse) Reset() {
	*x = PaymentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paymentService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentResponse) ProtoMessage() {}

func (x *PaymentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_paymentService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentResponse.ProtoReflect.Descriptor instead.
func (*PaymentResponse) Descriptor() ([]byte, []int) {
	return file_paymentService_proto_rawDescGZIP(), []int{3}
}

func (x *PaymentResponse) GetData() *PaymentData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *PaymentResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_paymentService_proto protoreflect.FileDescriptor

var file_paymentService_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xf8, 0x02, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x63, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x69, 0x63, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x19, 0x0a, 0x08, 0x69, 0x63, 0x6f, 0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x69, 0x63, 0x6f, 0x6e, 0x55, 0x72, 0x6c, 0x12, 0x15, 0x0a, 0x06, 0x69,
	0x73, 0x5f, 0x63, 0x6f, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x73, 0x43,
	0x6f, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x75, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x6c,
	0x69, 0x6e, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x73, 0x55, 0x6e, 0x64,
	0x65, 0x72, 0x4c, 0x69, 0x6e, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f,
	0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d,
	0x61, 0x70, 0x70, 0x6c, 0x79, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x12, 0x21, 0x0a,
	0x0c, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x64, 0x65, 0x73, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xaa,
	0x02, 0x0a, 0x0e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x6f, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x03, 0x74, 0x6f, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25,
	0x0a, 0x0e, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x54, 0x65, 0x72,
	0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x15, 0x0a,
	0x06, 0x69, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69,
	0x73, 0x43, 0x6f, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x75, 0x6e, 0x64, 0x65, 0x72,
	0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x73, 0x55,
	0x6e, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x6e, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0xa8, 0x01, 0x0a, 0x0b,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x29, 0x0a, 0x06, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x06,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x27, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12,
	0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70,
	0x61, 0x67, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x61, 0x0a, 0x0f, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xd0, 0x02, 0x0a, 0x0e, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x06,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a,
	0x03, 0x47, 0x65, 0x74, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x04, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x06, 0x53, 0x77,
	0x69, 0x74, 0x63, 0x68, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x06, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b,
	0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_paymentService_proto_rawDescOnce sync.Once
	file_paymentService_proto_rawDescData = file_paymentService_proto_rawDesc
)

func file_paymentService_proto_rawDescGZIP() []byte {
	file_paymentService_proto_rawDescOnce.Do(func() {
		file_paymentService_proto_rawDescData = protoimpl.X.CompressGZIP(file_paymentService_proto_rawDescData)
	})
	return file_paymentService_proto_rawDescData
}

var file_paymentService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_paymentService_proto_goTypes = []interface{}{
	(*Payment)(nil),         // 0: services.Payment
	(*PaymentRequest)(nil),  // 1: services.PaymentRequest
	(*PaymentData)(nil),     // 2: services.PaymentData
	(*PaymentResponse)(nil), // 3: services.PaymentResponse
	(*common.Pager)(nil),    // 4: common.Pager
	(*common.Info)(nil),     // 5: common.Info
	(*common.Error)(nil),    // 6: common.Error
}
var file_paymentService_proto_depIdxs = []int32{
	0,  // 0: services.PaymentData.entity:type_name -> services.Payment
	0,  // 1: services.PaymentData.items:type_name -> services.Payment
	4,  // 2: services.PaymentData.pager:type_name -> common.Pager
	5,  // 3: services.PaymentData.info:type_name -> common.Info
	2,  // 4: services.PaymentResponse.data:type_name -> services.PaymentData
	6,  // 5: services.PaymentResponse.error:type_name -> common.Error
	1,  // 6: services.PaymentService.Update:input_type -> services.PaymentRequest
	1,  // 7: services.PaymentService.Get:input_type -> services.PaymentRequest
	1,  // 8: services.PaymentService.List:input_type -> services.PaymentRequest
	1,  // 9: services.PaymentService.Switch:input_type -> services.PaymentRequest
	1,  // 10: services.PaymentService.Search:input_type -> services.PaymentRequest
	3,  // 11: services.PaymentService.Update:output_type -> services.PaymentResponse
	3,  // 12: services.PaymentService.Get:output_type -> services.PaymentResponse
	3,  // 13: services.PaymentService.List:output_type -> services.PaymentResponse
	3,  // 14: services.PaymentService.Switch:output_type -> services.PaymentResponse
	3,  // 15: services.PaymentService.Search:output_type -> services.PaymentResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_paymentService_proto_init() }
func file_paymentService_proto_init() {
	if File_paymentService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_paymentService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Payment); i {
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
		file_paymentService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentRequest); i {
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
		file_paymentService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentData); i {
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
		file_paymentService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentResponse); i {
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
			RawDescriptor: file_paymentService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_paymentService_proto_goTypes,
		DependencyIndexes: file_paymentService_proto_depIdxs,
		MessageInfos:      file_paymentService_proto_msgTypes,
	}.Build()
	File_paymentService_proto = out.File
	file_paymentService_proto_rawDesc = nil
	file_paymentService_proto_goTypes = nil
	file_paymentService_proto_depIdxs = nil
}
