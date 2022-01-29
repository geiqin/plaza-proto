// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: companyService.proto

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

//往来单位
type Company struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                int64             `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Name              string            `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`
	ShortName         string            `protobuf:"bytes,4,opt,name=short_name,json=shortName,proto3" json:"short_name"`
	Pinyin            string            `protobuf:"bytes,5,opt,name=pinyin,proto3" json:"pinyin"`
	CompanyCatId      int32             `protobuf:"varint,6,opt,name=company_cat_id,json=companyCatId,proto3" json:"company_cat_id"`
	LegalMan          string            `protobuf:"bytes,7,opt,name=legal_man,json=legalMan,proto3" json:"legal_man"`
	RegisteredCapital float32           `protobuf:"fixed32,8,opt,name=registered_capital,json=registeredCapital,proto3" json:"registered_capital"`
	RegisteredDate    string            `protobuf:"bytes,9,opt,name=registered_date,json=registeredDate,proto3" json:"registered_date"`
	UnifiedSocialCode string            `protobuf:"bytes,10,opt,name=unified_social_code,json=unifiedSocialCode,proto3" json:"unified_social_code"`
	AreaId            int64             `protobuf:"varint,11,opt,name=area_id,json=areaId,proto3" json:"area_id"`
	Address           string            `protobuf:"bytes,12,opt,name=address,proto3" json:"address"`
	Contacts          string            `protobuf:"bytes,13,opt,name=contacts,proto3" json:"contacts"`
	ContactPost       string            `protobuf:"bytes,14,opt,name=contact_post,json=contactPost,proto3" json:"contact_post"`
	ContactTel        string            `protobuf:"bytes,15,opt,name=contact_tel,json=contactTel,proto3" json:"contact_tel"`
	ContactEmail      string            `protobuf:"bytes,16,opt,name=contact_email,json=contactEmail,proto3" json:"contact_email"`
	Telephone         string            `protobuf:"bytes,17,opt,name=telephone,proto3" json:"telephone"`
	InvoiceType       int32             `protobuf:"varint,18,opt,name=invoice_type,json=invoiceType,proto3" json:"invoice_type"`
	BankNo            string            `protobuf:"bytes,19,opt,name=bank_no,json=bankNo,proto3" json:"bank_no"`
	BankName          string            `protobuf:"bytes,20,opt,name=bank_name,json=bankName,proto3" json:"bank_name"`
	Nature            int32             `protobuf:"varint,21,opt,name=nature,proto3" json:"nature"`
	Memo              string            `protobuf:"bytes,22,opt,name=memo,proto3" json:"memo"`
	Status            int32             `protobuf:"varint,23,opt,name=status,proto3" json:"status"`
	CreatedAt         string            `protobuf:"bytes,24,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt         string            `protobuf:"bytes,25,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	Metas             map[string]string `protobuf:"bytes,26,rep,name=metas,proto3" json:"metas" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Cat               *CompanyCat       `protobuf:"bytes,27,opt,name=cat,proto3" json:"cat"`
}

func (x *Company) Reset() {
	*x = Company{}
	if protoimpl.UnsafeEnabled {
		mi := &file_companyService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Company) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Company) ProtoMessage() {}

func (x *Company) ProtoReflect() protoreflect.Message {
	mi := &file_companyService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Company.ProtoReflect.Descriptor instead.
func (*Company) Descriptor() ([]byte, []int) {
	return file_companyService_proto_rawDescGZIP(), []int{0}
}

func (x *Company) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Company) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Company) GetShortName() string {
	if x != nil {
		return x.ShortName
	}
	return ""
}

func (x *Company) GetPinyin() string {
	if x != nil {
		return x.Pinyin
	}
	return ""
}

func (x *Company) GetCompanyCatId() int32 {
	if x != nil {
		return x.CompanyCatId
	}
	return 0
}

func (x *Company) GetLegalMan() string {
	if x != nil {
		return x.LegalMan
	}
	return ""
}

func (x *Company) GetRegisteredCapital() float32 {
	if x != nil {
		return x.RegisteredCapital
	}
	return 0
}

func (x *Company) GetRegisteredDate() string {
	if x != nil {
		return x.RegisteredDate
	}
	return ""
}

func (x *Company) GetUnifiedSocialCode() string {
	if x != nil {
		return x.UnifiedSocialCode
	}
	return ""
}

func (x *Company) GetAreaId() int64 {
	if x != nil {
		return x.AreaId
	}
	return 0
}

func (x *Company) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Company) GetContacts() string {
	if x != nil {
		return x.Contacts
	}
	return ""
}

func (x *Company) GetContactPost() string {
	if x != nil {
		return x.ContactPost
	}
	return ""
}

func (x *Company) GetContactTel() string {
	if x != nil {
		return x.ContactTel
	}
	return ""
}

func (x *Company) GetContactEmail() string {
	if x != nil {
		return x.ContactEmail
	}
	return ""
}

func (x *Company) GetTelephone() string {
	if x != nil {
		return x.Telephone
	}
	return ""
}

func (x *Company) GetInvoiceType() int32 {
	if x != nil {
		return x.InvoiceType
	}
	return 0
}

func (x *Company) GetBankNo() string {
	if x != nil {
		return x.BankNo
	}
	return ""
}

func (x *Company) GetBankName() string {
	if x != nil {
		return x.BankName
	}
	return ""
}

func (x *Company) GetNature() int32 {
	if x != nil {
		return x.Nature
	}
	return 0
}

func (x *Company) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *Company) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Company) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Company) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Company) GetMetas() map[string]string {
	if x != nil {
		return x.Metas
	}
	return nil
}

func (x *Company) GetCat() *CompanyCat {
	if x != nil {
		return x.Cat
	}
	return nil
}

type CompanyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int32 `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Sorting  int32 `protobuf:"varint,3,opt,name=sorting,proto3" json:"sorting"`
	//base params
	Id           int64   `protobuf:"varint,4,opt,name=id,proto3" json:"id"`
	Name         string  `protobuf:"bytes,5,opt,name=name,proto3" json:"name"`
	ShortName    string  `protobuf:"bytes,6,opt,name=short_name,json=shortName,proto3" json:"short_name"`
	Status       int32   `protobuf:"varint,7,opt,name=status,proto3" json:"status"`
	AreaId       int64   `protobuf:"varint,8,opt,name=area_id,json=areaId,proto3" json:"area_id"`
	InvoiceType  int32   `protobuf:"varint,9,opt,name=invoice_type,json=invoiceType,proto3" json:"invoice_type"`
	Nature       int32   `protobuf:"varint,10,opt,name=nature,proto3" json:"nature"`
	CompanyCatId int32   `protobuf:"varint,11,opt,name=company_cat_id,json=companyCatId,proto3" json:"company_cat_id"`
	Ids          []int64 `protobuf:"varint,12,rep,packed,name=ids,proto3" json:"ids"`
}

func (x *CompanyRequest) Reset() {
	*x = CompanyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_companyService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompanyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompanyRequest) ProtoMessage() {}

func (x *CompanyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_companyService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompanyRequest.ProtoReflect.Descriptor instead.
func (*CompanyRequest) Descriptor() ([]byte, []int) {
	return file_companyService_proto_rawDescGZIP(), []int{1}
}

func (x *CompanyRequest) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *CompanyRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *CompanyRequest) GetSorting() int32 {
	if x != nil {
		return x.Sorting
	}
	return 0
}

func (x *CompanyRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CompanyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CompanyRequest) GetShortName() string {
	if x != nil {
		return x.ShortName
	}
	return ""
}

func (x *CompanyRequest) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *CompanyRequest) GetAreaId() int64 {
	if x != nil {
		return x.AreaId
	}
	return 0
}

func (x *CompanyRequest) GetInvoiceType() int32 {
	if x != nil {
		return x.InvoiceType
	}
	return 0
}

func (x *CompanyRequest) GetNature() int32 {
	if x != nil {
		return x.Nature
	}
	return 0
}

func (x *CompanyRequest) GetCompanyCatId() int32 {
	if x != nil {
		return x.CompanyCatId
	}
	return 0
}

func (x *CompanyRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

//
type CompanyData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Company      `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*Company    `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Info   *common.Info  `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
}

func (x *CompanyData) Reset() {
	*x = CompanyData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_companyService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompanyData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompanyData) ProtoMessage() {}

func (x *CompanyData) ProtoReflect() protoreflect.Message {
	mi := &file_companyService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompanyData.ProtoReflect.Descriptor instead.
func (*CompanyData) Descriptor() ([]byte, []int) {
	return file_companyService_proto_rawDescGZIP(), []int{2}
}

func (x *CompanyData) GetEntity() *Company {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *CompanyData) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *CompanyData) GetItems() []*Company {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *CompanyData) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

//
type CompanyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *CompanyData  `protobuf:"bytes,1,opt,name=data,proto3" json:"data"`
	Error *common.Error `protobuf:"bytes,2,opt,name=error,proto3" json:"error"`
}

func (x *CompanyResponse) Reset() {
	*x = CompanyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_companyService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompanyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompanyResponse) ProtoMessage() {}

func (x *CompanyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_companyService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompanyResponse.ProtoReflect.Descriptor instead.
func (*CompanyResponse) Descriptor() ([]byte, []int) {
	return file_companyService_proto_rawDescGZIP(), []int{3}
}

func (x *CompanyResponse) GetData() *CompanyData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *CompanyResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_companyService_proto protoreflect.FileDescriptor

var file_companyService_proto_rawDesc = []byte{
	0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x43, 0x61, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf6, 0x06, 0x0a,
	0x07, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70,
	0x69, 0x6e, 0x79, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x69, 0x6e,
	0x79, 0x69, 0x6e, 0x12, 0x24, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x63,
	0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x43, 0x61, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x65, 0x67,
	0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x65,
	0x67, 0x61, 0x6c, 0x4d, 0x61, 0x6e, 0x12, 0x2d, 0x0a, 0x12, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x65, 0x64, 0x5f, 0x63, 0x61, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x11, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x43, 0x61,
	0x70, 0x69, 0x74, 0x61, 0x6c, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x2e,
	0x0a, 0x13, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x75, 0x6e, 0x69,
	0x66, 0x69, 0x65, 0x64, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x17,
	0x0a, 0x07, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x61, 0x72, 0x65, 0x61, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x12, 0x21, 0x0a,
	0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x50, 0x6f, 0x73, 0x74,
	0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x74, 0x65, 0x6c, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x54, 0x65,
	0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x65, 0x6c, 0x65, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x69, 0x6e, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x61, 0x6e, 0x6b, 0x5f,
	0x6e, 0x6f, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x61, 0x6e, 0x6b, 0x4e, 0x6f,
	0x12, 0x1b, 0x0a, 0x09, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x61, 0x6e, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x16, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x17, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x19,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x32, 0x0a, 0x05, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x18, 0x1a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x6d, 0x65,
	0x74, 0x61, 0x73, 0x12, 0x26, 0x0a, 0x03, 0x63, 0x61, 0x74, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x43, 0x61, 0x74, 0x52, 0x03, 0x63, 0x61, 0x74, 0x1a, 0x38, 0x0a, 0x0a, 0x4d,
	0x65, 0x74, 0x61, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xc4, 0x02, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x6f,
	0x72, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x68, 0x6f, 0x72, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x17, 0x0a, 0x07, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x61, 0x72, 0x65, 0x61, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x6e, 0x76,
	0x6f, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0b, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f,
	0x63, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x43, 0x61, 0x74, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64,
	0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0xa8, 0x01, 0x0a,
	0x0b, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x44, 0x61, 0x74, 0x61, 0x12, 0x29, 0x0a, 0x06,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52,
	0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x61, 0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xb6, 0x02, 0x0a, 0x0e, 0x43,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a,
	0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x38, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x1a, 0x19,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x03, 0x47,
	0x65, 0x74, 0x12, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x3f, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x18, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_companyService_proto_rawDescOnce sync.Once
	file_companyService_proto_rawDescData = file_companyService_proto_rawDesc
)

func file_companyService_proto_rawDescGZIP() []byte {
	file_companyService_proto_rawDescOnce.Do(func() {
		file_companyService_proto_rawDescData = protoimpl.X.CompressGZIP(file_companyService_proto_rawDescData)
	})
	return file_companyService_proto_rawDescData
}

var file_companyService_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_companyService_proto_goTypes = []interface{}{
	(*Company)(nil),         // 0: services.Company
	(*CompanyRequest)(nil),  // 1: services.CompanyRequest
	(*CompanyData)(nil),     // 2: services.CompanyData
	(*CompanyResponse)(nil), // 3: services.CompanyResponse
	nil,                     // 4: services.Company.MetasEntry
	(*CompanyCat)(nil),      // 5: services.CompanyCat
	(*common.Pager)(nil),    // 6: common.Pager
	(*common.Info)(nil),     // 7: common.Info
	(*common.Error)(nil),    // 8: common.Error
}
var file_companyService_proto_depIdxs = []int32{
	4,  // 0: services.Company.metas:type_name -> services.Company.MetasEntry
	5,  // 1: services.Company.cat:type_name -> services.CompanyCat
	0,  // 2: services.CompanyData.entity:type_name -> services.Company
	6,  // 3: services.CompanyData.pager:type_name -> common.Pager
	0,  // 4: services.CompanyData.items:type_name -> services.Company
	7,  // 5: services.CompanyData.info:type_name -> common.Info
	2,  // 6: services.CompanyResponse.data:type_name -> services.CompanyData
	8,  // 7: services.CompanyResponse.error:type_name -> common.Error
	0,  // 8: services.CompanyService.Create:input_type -> services.Company
	0,  // 9: services.CompanyService.Update:input_type -> services.Company
	0,  // 10: services.CompanyService.Delete:input_type -> services.Company
	0,  // 11: services.CompanyService.Get:input_type -> services.Company
	1,  // 12: services.CompanyService.Search:input_type -> services.CompanyRequest
	3,  // 13: services.CompanyService.Create:output_type -> services.CompanyResponse
	3,  // 14: services.CompanyService.Update:output_type -> services.CompanyResponse
	3,  // 15: services.CompanyService.Delete:output_type -> services.CompanyResponse
	3,  // 16: services.CompanyService.Get:output_type -> services.CompanyResponse
	3,  // 17: services.CompanyService.Search:output_type -> services.CompanyResponse
	13, // [13:18] is the sub-list for method output_type
	8,  // [8:13] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_companyService_proto_init() }
func file_companyService_proto_init() {
	if File_companyService_proto != nil {
		return
	}
	file_companyCatService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_companyService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Company); i {
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
		file_companyService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompanyRequest); i {
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
		file_companyService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompanyData); i {
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
		file_companyService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompanyResponse); i {
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
			RawDescriptor: file_companyService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_companyService_proto_goTypes,
		DependencyIndexes: file_companyService_proto_depIdxs,
		MessageInfos:      file_companyService_proto_msgTypes,
	}.Build()
	File_companyService_proto = out.File
	file_companyService_proto_rawDesc = nil
	file_companyService_proto_goTypes = nil
	file_companyService_proto_depIdxs = nil
}
