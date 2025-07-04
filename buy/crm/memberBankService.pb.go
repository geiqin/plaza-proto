// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: memberBankService.proto

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

type MemberBank struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id"`                                        //ID
	MemberId     int64   `protobuf:"varint,2,opt,name=member_id,json=memberId,proto3" json:"member_id"`            //用户ID
	BankName     string  `protobuf:"bytes,3,opt,name=bank_name,json=bankName,proto3" json:"bank_name"`             //开户平台
	BankAccount  string  `protobuf:"bytes,4,opt,name=bank_account,json=bankAccount,proto3" json:"bank_account"`    //开户账号
	BankUsername string  `protobuf:"bytes,5,opt,name=bank_username,json=bankUsername,proto3" json:"bank_username"` //开户姓名
	IsDefault    string  `protobuf:"bytes,6,opt,name=is_default,json=isDefault,proto3" json:"is_default"`          //是否为默认
	Status       string  `protobuf:"bytes,7,opt,name=status,proto3" json:"status"`                                 //状态：0 无效，1 正常
	CreatedAt    int64   `protobuf:"varint,8,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt    int64   `protobuf:"varint,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	Member       *Member `protobuf:"bytes,10,opt,name=member,proto3" json:"member"`
}

func (x *MemberBank) Reset() {
	*x = MemberBank{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memberBankService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberBank) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberBank) ProtoMessage() {}

func (x *MemberBank) ProtoReflect() protoreflect.Message {
	mi := &file_memberBankService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberBank.ProtoReflect.Descriptor instead.
func (*MemberBank) Descriptor() ([]byte, []int) {
	return file_memberBankService_proto_rawDescGZIP(), []int{0}
}

func (x *MemberBank) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MemberBank) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

func (x *MemberBank) GetBankName() string {
	if x != nil {
		return x.BankName
	}
	return ""
}

func (x *MemberBank) GetBankAccount() string {
	if x != nil {
		return x.BankAccount
	}
	return ""
}

func (x *MemberBank) GetBankUsername() string {
	if x != nil {
		return x.BankUsername
	}
	return ""
}

func (x *MemberBank) GetIsDefault() string {
	if x != nil {
		return x.IsDefault
	}
	return ""
}

func (x *MemberBank) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *MemberBank) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *MemberBank) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *MemberBank) GetMember() *Member {
	if x != nil {
		return x.Member
	}
	return nil
}

type MemberBankRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Top          int32    `protobuf:"varint,1,opt,name=top,proto3" json:"top"`
	Paged        int64    `protobuf:"varint,2,opt,name=paged,proto3" json:"paged"`
	PageSize     int64    `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Keywords     string   `protobuf:"bytes,4,opt,name=keywords,proto3" json:"keywords"`
	Sorts        []string `protobuf:"bytes,5,rep,name=sorts,proto3" json:"sorts"`
	DateRange    []string `protobuf:"bytes,6,rep,name=date_range,json=dateRange,proto3" json:"date_range"`
	Ids          []int64  `protobuf:"varint,7,rep,packed,name=ids,proto3" json:"ids"`
	Id           int64    `protobuf:"varint,8,opt,name=id,proto3" json:"id"`
	MemberId     int64    `protobuf:"varint,10,opt,name=member_id,json=memberId,proto3" json:"member_id"`
	BankUsername string   `protobuf:"bytes,11,opt,name=bank_username,json=bankUsername,proto3" json:"bank_username"`
	BankAccount  string   `protobuf:"bytes,12,opt,name=bank_account,json=bankAccount,proto3" json:"bank_account"`
	BankName     string   `protobuf:"bytes,13,opt,name=bank_name,json=bankName,proto3" json:"bank_name"`
	Status       string   `protobuf:"bytes,14,opt,name=status,proto3" json:"status"`
}

func (x *MemberBankRequest) Reset() {
	*x = MemberBankRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memberBankService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberBankRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberBankRequest) ProtoMessage() {}

func (x *MemberBankRequest) ProtoReflect() protoreflect.Message {
	mi := &file_memberBankService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberBankRequest.ProtoReflect.Descriptor instead.
func (*MemberBankRequest) Descriptor() ([]byte, []int) {
	return file_memberBankService_proto_rawDescGZIP(), []int{1}
}

func (x *MemberBankRequest) GetTop() int32 {
	if x != nil {
		return x.Top
	}
	return 0
}

func (x *MemberBankRequest) GetPaged() int64 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *MemberBankRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *MemberBankRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *MemberBankRequest) GetSorts() []string {
	if x != nil {
		return x.Sorts
	}
	return nil
}

func (x *MemberBankRequest) GetDateRange() []string {
	if x != nil {
		return x.DateRange
	}
	return nil
}

func (x *MemberBankRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *MemberBankRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MemberBankRequest) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

func (x *MemberBankRequest) GetBankUsername() string {
	if x != nil {
		return x.BankUsername
	}
	return ""
}

func (x *MemberBankRequest) GetBankAccount() string {
	if x != nil {
		return x.BankAccount
	}
	return ""
}

func (x *MemberBankRequest) GetBankName() string {
	if x != nil {
		return x.BankName
	}
	return ""
}

func (x *MemberBankRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type MemberBankResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *MemberBank   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*MemberBank `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Msg    string        `protobuf:"bytes,4,opt,name=msg,proto3" json:"msg"`
}

func (x *MemberBankResponse) Reset() {
	*x = MemberBankResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memberBankService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberBankResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberBankResponse) ProtoMessage() {}

func (x *MemberBankResponse) ProtoReflect() protoreflect.Message {
	mi := &file_memberBankService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberBankResponse.ProtoReflect.Descriptor instead.
func (*MemberBankResponse) Descriptor() ([]byte, []int) {
	return file_memberBankService_proto_rawDescGZIP(), []int{2}
}

func (x *MemberBankResponse) GetEntity() *MemberBank {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *MemberBankResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *MemberBankResponse) GetItems() []*MemberBank {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *MemberBankResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_memberBankService_proto protoreflect.FileDescriptor

var file_memberBankService_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x61, 0x6e, 0x6b, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbd, 0x02, 0x0a, 0x0a,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x61, 0x6e, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x61, 0x6e, 0x6b, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x61, 0x6e, 0x6b,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x61, 0x6e, 0x6b,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x62, 0x61, 0x6e, 0x6b, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x62, 0x61, 0x6e, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x69, 0x73, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x69, 0x73, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x28, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x22, 0xe5, 0x02, 0x0a, 0x11,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
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
	0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62,
	0x61, 0x6e, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x62,
	0x61, 0x6e, 0x6b, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x62, 0x61, 0x6e, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x62, 0x61, 0x6e, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0xa5, 0x01, 0x0a, 0x12, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x61,
	0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x61, 0x6e, 0x6b,
	0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x2a, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x61,
	0x6e, 0x6b, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x32, 0x97, 0x02, 0x0a, 0x11,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x61, 0x6e, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x3e, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x61, 0x6e,
	0x6b, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x42, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x3e, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x61, 0x6e,
	0x6b, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x42, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x3b, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x61, 0x6e, 0x6b, 0x1a, 0x1c,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x42, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45,
	0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x61, 0x6e, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_memberBankService_proto_rawDescOnce sync.Once
	file_memberBankService_proto_rawDescData = file_memberBankService_proto_rawDesc
)

func file_memberBankService_proto_rawDescGZIP() []byte {
	file_memberBankService_proto_rawDescOnce.Do(func() {
		file_memberBankService_proto_rawDescData = protoimpl.X.CompressGZIP(file_memberBankService_proto_rawDescData)
	})
	return file_memberBankService_proto_rawDescData
}

var file_memberBankService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_memberBankService_proto_goTypes = []interface{}{
	(*MemberBank)(nil),         // 0: services.MemberBank
	(*MemberBankRequest)(nil),  // 1: services.MemberBankRequest
	(*MemberBankResponse)(nil), // 2: services.MemberBankResponse
	(*Member)(nil),             // 3: services.Member
	(*common.Pager)(nil),       // 4: common.Pager
}
var file_memberBankService_proto_depIdxs = []int32{
	3, // 0: services.MemberBank.member:type_name -> services.Member
	0, // 1: services.MemberBankResponse.entity:type_name -> services.MemberBank
	4, // 2: services.MemberBankResponse.pager:type_name -> common.Pager
	0, // 3: services.MemberBankResponse.items:type_name -> services.MemberBank
	0, // 4: services.MemberBankService.Create:input_type -> services.MemberBank
	0, // 5: services.MemberBankService.Delete:input_type -> services.MemberBank
	0, // 6: services.MemberBankService.Get:input_type -> services.MemberBank
	1, // 7: services.MemberBankService.Search:input_type -> services.MemberBankRequest
	2, // 8: services.MemberBankService.Create:output_type -> services.MemberBankResponse
	2, // 9: services.MemberBankService.Delete:output_type -> services.MemberBankResponse
	2, // 10: services.MemberBankService.Get:output_type -> services.MemberBankResponse
	2, // 11: services.MemberBankService.Search:output_type -> services.MemberBankResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_memberBankService_proto_init() }
func file_memberBankService_proto_init() {
	if File_memberBankService_proto != nil {
		return
	}
	file_memberService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_memberBankService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberBank); i {
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
		file_memberBankService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberBankRequest); i {
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
		file_memberBankService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberBankResponse); i {
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
			RawDescriptor: file_memberBankService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_memberBankService_proto_goTypes,
		DependencyIndexes: file_memberBankService_proto_depIdxs,
		MessageInfos:      file_memberBankService_proto_msgTypes,
	}.Build()
	File_memberBankService_proto = out.File
	file_memberBankService_proto_rawDesc = nil
	file_memberBankService_proto_goTypes = nil
	file_memberBankService_proto_depIdxs = nil
}
