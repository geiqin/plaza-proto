// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: walletService.proto

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

type Wallet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id"`                                      //ID
	MemberId    int64   `protobuf:"varint,2,opt,name=member_id,json=memberId,proto3" json:"member_id"`          //用户ID
	NormalMoney int64   `protobuf:"varint,3,opt,name=normal_money,json=normalMoney,proto3" json:"normal_money"` //有效金额（可用余额=现金余额+赠送金+信用额度-冻结金额）
	FrozenMoney int64   `protobuf:"varint,4,opt,name=frozen_money,json=frozenMoney,proto3" json:"frozen_money"` //冻结金额
	GiveMoney   int64   `protobuf:"varint,5,opt,name=give_money,json=giveMoney,proto3" json:"give_money"`       //赠送金额
	CreditMoney int64   `protobuf:"varint,6,opt,name=credit_money,json=creditMoney,proto3" json:"credit_money"` //信用额度
	Status      string  `protobuf:"bytes,7,opt,name=status,proto3" json:"status"`                               //状态：0停用 1正常
	CreatedAt   string  `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt   string  `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	Member      *Member `protobuf:"bytes,10,opt,name=member,proto3" json:"member"`
}

func (x *Wallet) Reset() {
	*x = Wallet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_walletService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Wallet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Wallet) ProtoMessage() {}

func (x *Wallet) ProtoReflect() protoreflect.Message {
	mi := &file_walletService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Wallet.ProtoReflect.Descriptor instead.
func (*Wallet) Descriptor() ([]byte, []int) {
	return file_walletService_proto_rawDescGZIP(), []int{0}
}

func (x *Wallet) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Wallet) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

func (x *Wallet) GetNormalMoney() int64 {
	if x != nil {
		return x.NormalMoney
	}
	return 0
}

func (x *Wallet) GetFrozenMoney() int64 {
	if x != nil {
		return x.FrozenMoney
	}
	return 0
}

func (x *Wallet) GetGiveMoney() int64 {
	if x != nil {
		return x.GiveMoney
	}
	return 0
}

func (x *Wallet) GetCreditMoney() int64 {
	if x != nil {
		return x.CreditMoney
	}
	return 0
}

func (x *Wallet) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Wallet) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Wallet) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Wallet) GetMember() *Member {
	if x != nil {
		return x.Member
	}
	return nil
}

//查询参数
type WalletRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int32 `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	//以下为自定义参数
	Name     string `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`
	Id       int64  `protobuf:"varint,4,opt,name=id,proto3" json:"id"`
	MemberId int64  `protobuf:"varint,5,opt,name=member_id,json=memberId,proto3" json:"member_id"`
	Status   string `protobuf:"bytes,6,opt,name=status,proto3" json:"status"`
	Type     int32  `protobuf:"varint,7,opt,name=type,proto3" json:"type"`
}

func (x *WalletRequest) Reset() {
	*x = WalletRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_walletService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WalletRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WalletRequest) ProtoMessage() {}

func (x *WalletRequest) ProtoReflect() protoreflect.Message {
	mi := &file_walletService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WalletRequest.ProtoReflect.Descriptor instead.
func (*WalletRequest) Descriptor() ([]byte, []int) {
	return file_walletService_proto_rawDescGZIP(), []int{1}
}

func (x *WalletRequest) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *WalletRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *WalletRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *WalletRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *WalletRequest) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

func (x *WalletRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *WalletRequest) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

type WalletData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Wallet       `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*Wallet     `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Info   *common.Info  `protobuf:"bytes,4,opt,name=info,proto3" json:"info"`
}

func (x *WalletData) Reset() {
	*x = WalletData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_walletService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WalletData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WalletData) ProtoMessage() {}

func (x *WalletData) ProtoReflect() protoreflect.Message {
	mi := &file_walletService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WalletData.ProtoReflect.Descriptor instead.
func (*WalletData) Descriptor() ([]byte, []int) {
	return file_walletService_proto_rawDescGZIP(), []int{2}
}

func (x *WalletData) GetEntity() *Wallet {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *WalletData) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *WalletData) GetItems() []*Wallet {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *WalletData) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type WalletResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *WalletData   `protobuf:"bytes,1,opt,name=data,proto3" json:"data"`
	Error *common.Error `protobuf:"bytes,2,opt,name=error,proto3" json:"error"`
}

func (x *WalletResponse) Reset() {
	*x = WalletResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_walletService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WalletResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WalletResponse) ProtoMessage() {}

func (x *WalletResponse) ProtoReflect() protoreflect.Message {
	mi := &file_walletService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WalletResponse.ProtoReflect.Descriptor instead.
func (*WalletResponse) Descriptor() ([]byte, []int) {
	return file_walletService_proto_rawDescGZIP(), []int{3}
}

func (x *WalletResponse) GetData() *WalletData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *WalletResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_walletService_proto protoreflect.FileDescriptor

var file_walletService_proto_rawDesc = []byte{
	0x0a, 0x13, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a,
	0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x13, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbd, 0x02, 0x0a, 0x06, 0x57, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x21, 0x0a, 0x0c, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x5f, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x4d, 0x6f, 0x6e,
	0x65, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x72, 0x6f, 0x7a, 0x65, 0x6e, 0x5f, 0x6d, 0x6f, 0x6e,
	0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x66, 0x72, 0x6f, 0x7a, 0x65, 0x6e,
	0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x69, 0x76, 0x65, 0x5f, 0x6d, 0x6f,
	0x6e, 0x65, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x67, 0x69, 0x76, 0x65, 0x4d,
	0x6f, 0x6e, 0x65, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5f, 0x6d,
	0x6f, 0x6e, 0x65, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x64,
	0x69, 0x74, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x28, 0x0a,
	0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52,
	0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x22, 0xaf, 0x01, 0x0a, 0x0d, 0x57, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67,
	0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0xa5, 0x01, 0x0a, 0x0a, 0x57, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x28, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72,
	0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12,
	0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66,
	0x6f, 0x22, 0x5f, 0x0a, 0x0e, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x57, 0x61, 0x6c,
	0x6c, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x23, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x32, 0xcf, 0x03, 0x0a, 0x0d, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x17, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x3d, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x3f, 0x0a, 0x08, 0x43, 0x61, 0x73, 0x68, 0x49, 0x6e, 0x69, 0x74, 0x12, 0x17, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x41, 0x0a, 0x0a, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x53, 0x65, 0x6e, 0x64,
	0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x57, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x0b, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x03, 0x47, 0x65, 0x74,
	0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x57, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12,
	0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x57, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_walletService_proto_rawDescOnce sync.Once
	file_walletService_proto_rawDescData = file_walletService_proto_rawDesc
)

func file_walletService_proto_rawDescGZIP() []byte {
	file_walletService_proto_rawDescOnce.Do(func() {
		file_walletService_proto_rawDescData = protoimpl.X.CompressGZIP(file_walletService_proto_rawDescData)
	})
	return file_walletService_proto_rawDescData
}

var file_walletService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_walletService_proto_goTypes = []interface{}{
	(*Wallet)(nil),         // 0: services.Wallet
	(*WalletRequest)(nil),  // 1: services.WalletRequest
	(*WalletData)(nil),     // 2: services.WalletData
	(*WalletResponse)(nil), // 3: services.WalletResponse
	(*Member)(nil),         // 4: services.Member
	(*common.Pager)(nil),   // 5: common.Pager
	(*common.Info)(nil),    // 6: common.Info
	(*common.Error)(nil),   // 7: common.Error
}
var file_walletService_proto_depIdxs = []int32{
	4,  // 0: services.Wallet.member:type_name -> services.Member
	0,  // 1: services.WalletData.entity:type_name -> services.Wallet
	5,  // 2: services.WalletData.pager:type_name -> common.Pager
	0,  // 3: services.WalletData.items:type_name -> services.Wallet
	6,  // 4: services.WalletData.info:type_name -> common.Info
	2,  // 5: services.WalletResponse.data:type_name -> services.WalletData
	7,  // 6: services.WalletResponse.error:type_name -> common.Error
	1,  // 7: services.WalletService.Index:input_type -> services.WalletRequest
	1,  // 8: services.WalletService.Create:input_type -> services.WalletRequest
	1,  // 9: services.WalletService.CashInit:input_type -> services.WalletRequest
	1,  // 10: services.WalletService.VerifySend:input_type -> services.WalletRequest
	1,  // 11: services.WalletService.VerifyCheck:input_type -> services.WalletRequest
	1,  // 12: services.WalletService.Get:input_type -> services.WalletRequest
	1,  // 13: services.WalletService.Search:input_type -> services.WalletRequest
	3,  // 14: services.WalletService.Index:output_type -> services.WalletResponse
	3,  // 15: services.WalletService.Create:output_type -> services.WalletResponse
	3,  // 16: services.WalletService.CashInit:output_type -> services.WalletResponse
	3,  // 17: services.WalletService.VerifySend:output_type -> services.WalletResponse
	3,  // 18: services.WalletService.VerifyCheck:output_type -> services.WalletResponse
	3,  // 19: services.WalletService.Get:output_type -> services.WalletResponse
	3,  // 20: services.WalletService.Search:output_type -> services.WalletResponse
	14, // [14:21] is the sub-list for method output_type
	7,  // [7:14] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_walletService_proto_init() }
func file_walletService_proto_init() {
	if File_walletService_proto != nil {
		return
	}
	file_memberService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_walletService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Wallet); i {
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
		file_walletService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WalletRequest); i {
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
		file_walletService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WalletData); i {
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
		file_walletService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WalletResponse); i {
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
			RawDescriptor: file_walletService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_walletService_proto_goTypes,
		DependencyIndexes: file_walletService_proto_depIdxs,
		MessageInfos:      file_walletService_proto_msgTypes,
	}.Build()
	File_walletService_proto = out.File
	file_walletService_proto_rawDesc = nil
	file_walletService_proto_goTypes = nil
	file_walletService_proto_depIdxs = nil
}
