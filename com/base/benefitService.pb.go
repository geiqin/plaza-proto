// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: benefitService.proto

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

type Benefit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Code       string `protobuf:"bytes,2,opt,name=code,proto3" json:"code"`
	Name       string `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`
	Title      string `protobuf:"bytes,4,opt,name=title,proto3" json:"title"`
	Type       string `protobuf:"bytes,5,opt,name=type,proto3" json:"type"`
	Industry   string `protobuf:"bytes,6,opt,name=industry,proto3" json:"industry"`
	IconId     int64  `protobuf:"varint,7,opt,name=icon_id,json=iconId,proto3" json:"icon_id"`
	IconUrl    string `protobuf:"bytes,8,opt,name=icon_url,json=iconUrl,proto3" json:"icon_url"`
	Method     string `protobuf:"bytes,9,opt,name=method,proto3" json:"method"`
	PriceValue int32  `protobuf:"varint,10,opt,name=price_value,json=priceValue,proto3" json:"price_value"`
	Selected   bool   `protobuf:"varint,11,opt,name=selected,proto3" json:"selected"`
	Memo       string `protobuf:"bytes,12,opt,name=memo,proto3" json:"memo"`
	Status     string `protobuf:"bytes,13,opt,name=status,proto3" json:"status"`
	CreatedAt  string `protobuf:"bytes,14,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt  string `protobuf:"bytes,15,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
}

func (x *Benefit) Reset() {
	*x = Benefit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_benefitService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Benefit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Benefit) ProtoMessage() {}

func (x *Benefit) ProtoReflect() protoreflect.Message {
	mi := &file_benefitService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Benefit.ProtoReflect.Descriptor instead.
func (*Benefit) Descriptor() ([]byte, []int) {
	return file_benefitService_proto_rawDescGZIP(), []int{0}
}

func (x *Benefit) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Benefit) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Benefit) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Benefit) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Benefit) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Benefit) GetIndustry() string {
	if x != nil {
		return x.Industry
	}
	return ""
}

func (x *Benefit) GetIconId() int64 {
	if x != nil {
		return x.IconId
	}
	return 0
}

func (x *Benefit) GetIconUrl() string {
	if x != nil {
		return x.IconUrl
	}
	return ""
}

func (x *Benefit) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *Benefit) GetPriceValue() int32 {
	if x != nil {
		return x.PriceValue
	}
	return 0
}

func (x *Benefit) GetSelected() bool {
	if x != nil {
		return x.Selected
	}
	return false
}

func (x *Benefit) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *Benefit) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Benefit) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Benefit) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

// 查询参数
type BenefitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int64 `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	//以下为自定义参数
	Name   string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`
	Title  string   `protobuf:"bytes,4,opt,name=title,proto3" json:"title"`
	Type   string   `protobuf:"bytes,5,opt,name=type,proto3" json:"type"`
	Status string   `protobuf:"bytes,6,opt,name=status,proto3" json:"status"`
	Ids    []int32  `protobuf:"varint,7,rep,packed,name=ids,proto3" json:"ids"`
	Codes  []string `protobuf:"bytes,8,rep,name=codes,proto3" json:"codes"`
}

func (x *BenefitRequest) Reset() {
	*x = BenefitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_benefitService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BenefitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BenefitRequest) ProtoMessage() {}

func (x *BenefitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_benefitService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BenefitRequest.ProtoReflect.Descriptor instead.
func (*BenefitRequest) Descriptor() ([]byte, []int) {
	return file_benefitService_proto_rawDescGZIP(), []int{1}
}

func (x *BenefitRequest) GetPaged() int64 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *BenefitRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *BenefitRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BenefitRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *BenefitRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *BenefitRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *BenefitRequest) GetIds() []int32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *BenefitRequest) GetCodes() []string {
	if x != nil {
		return x.Codes
	}
	return nil
}

type BenefitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Benefit      `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*Benefit    `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Info   string        `protobuf:"bytes,4,opt,name=info,proto3" json:"info"`
}

func (x *BenefitResponse) Reset() {
	*x = BenefitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_benefitService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BenefitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BenefitResponse) ProtoMessage() {}

func (x *BenefitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_benefitService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BenefitResponse.ProtoReflect.Descriptor instead.
func (*BenefitResponse) Descriptor() ([]byte, []int) {
	return file_benefitService_proto_rawDescGZIP(), []int{2}
}

func (x *BenefitResponse) GetEntity() *Benefit {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *BenefitResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *BenefitResponse) GetItems() []*Benefit {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *BenefitResponse) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

var File_benefitService_proto protoreflect.FileDescriptor

var file_benefitService_proto_rawDesc = []byte{
	0x0a, 0x14, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xfa, 0x02, 0x0a, 0x07, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x12, 0x17, 0x0a,
	0x07, 0x69, 0x63, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x69, 0x63, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x63, 0x6f, 0x6e, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x63, 0x6f, 0x6e, 0x55, 0x72,
	0x6c, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x73, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x22, 0xc1, 0x01, 0x0a, 0x0e, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03,
	0x69, 0x64, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x05, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x63,
	0x6f, 0x64, 0x65, 0x73, 0x22, 0x9e, 0x01, 0x0a, 0x0f, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x52, 0x06, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65,
	0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0xf5, 0x02, 0x0a, 0x0e, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x65,
	0x6e, 0x65, 0x66, 0x69, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x38, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x1a,
	0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x65, 0x6e, 0x65, 0x66,
	0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x06,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x11, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74,
	0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x65, 0x6e, 0x65,
	0x66, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a,
	0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x65, 0x6e,
	0x65, 0x66, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d,
	0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x65, 0x6e, 0x65,
	0x66, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a,
	0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_benefitService_proto_rawDescOnce sync.Once
	file_benefitService_proto_rawDescData = file_benefitService_proto_rawDesc
)

func file_benefitService_proto_rawDescGZIP() []byte {
	file_benefitService_proto_rawDescOnce.Do(func() {
		file_benefitService_proto_rawDescData = protoimpl.X.CompressGZIP(file_benefitService_proto_rawDescData)
	})
	return file_benefitService_proto_rawDescData
}

var file_benefitService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_benefitService_proto_goTypes = []interface{}{
	(*Benefit)(nil),         // 0: services.Benefit
	(*BenefitRequest)(nil),  // 1: services.BenefitRequest
	(*BenefitResponse)(nil), // 2: services.BenefitResponse
	(*common.Pager)(nil),    // 3: common.Pager
}
var file_benefitService_proto_depIdxs = []int32{
	0, // 0: services.BenefitResponse.entity:type_name -> services.Benefit
	3, // 1: services.BenefitResponse.pager:type_name -> common.Pager
	0, // 2: services.BenefitResponse.items:type_name -> services.Benefit
	0, // 3: services.BenefitService.Create:input_type -> services.Benefit
	0, // 4: services.BenefitService.Update:input_type -> services.Benefit
	0, // 5: services.BenefitService.Delete:input_type -> services.Benefit
	0, // 6: services.BenefitService.Get:input_type -> services.Benefit
	1, // 7: services.BenefitService.Search:input_type -> services.BenefitRequest
	1, // 8: services.BenefitService.List:input_type -> services.BenefitRequest
	2, // 9: services.BenefitService.Create:output_type -> services.BenefitResponse
	2, // 10: services.BenefitService.Update:output_type -> services.BenefitResponse
	2, // 11: services.BenefitService.Delete:output_type -> services.BenefitResponse
	2, // 12: services.BenefitService.Get:output_type -> services.BenefitResponse
	2, // 13: services.BenefitService.Search:output_type -> services.BenefitResponse
	2, // 14: services.BenefitService.List:output_type -> services.BenefitResponse
	9, // [9:15] is the sub-list for method output_type
	3, // [3:9] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_benefitService_proto_init() }
func file_benefitService_proto_init() {
	if File_benefitService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_benefitService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Benefit); i {
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
		file_benefitService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BenefitRequest); i {
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
		file_benefitService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BenefitResponse); i {
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
			RawDescriptor: file_benefitService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_benefitService_proto_goTypes,
		DependencyIndexes: file_benefitService_proto_depIdxs,
		MessageInfos:      file_benefitService_proto_msgTypes,
	}.Build()
	File_benefitService_proto = out.File
	file_benefitService_proto_rawDesc = nil
	file_benefitService_proto_goTypes = nil
	file_benefitService_proto_depIdxs = nil
}
