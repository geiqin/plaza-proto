// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: expressService.proto

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

type ExpressRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged      int32   `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize   int32   `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Sorting    string  `protobuf:"bytes,3,opt,name=sorting,proto3" json:"sorting"`
	Keywords   string  `protobuf:"bytes,4,opt,name=keywords,proto3" json:"keywords"`
	Id         int64   `protobuf:"varint,5,opt,name=id,proto3" json:"id"`
	ExpressId  int64   `protobuf:"varint,6,opt,name=express_id,json=expressId,proto3" json:"express_id"`
	Ids        []int64 `protobuf:"varint,7,rep,packed,name=ids,proto3" json:"ids"`
	ExpressIds []int64 `protobuf:"varint,8,rep,packed,name=express_ids,json=expressIds,proto3" json:"express_ids"`
}

func (x *ExpressRequest) Reset() {
	*x = ExpressRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_expressService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpressRequest) ProtoMessage() {}

func (x *ExpressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_expressService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpressRequest.ProtoReflect.Descriptor instead.
func (*ExpressRequest) Descriptor() ([]byte, []int) {
	return file_expressService_proto_rawDescGZIP(), []int{0}
}

func (x *ExpressRequest) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *ExpressRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ExpressRequest) GetSorting() string {
	if x != nil {
		return x.Sorting
	}
	return ""
}

func (x *ExpressRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *ExpressRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ExpressRequest) GetExpressId() int64 {
	if x != nil {
		return x.ExpressId
	}
	return 0
}

func (x *ExpressRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *ExpressRequest) GetExpressIds() []int64 {
	if x != nil {
		return x.ExpressIds
	}
	return nil
}

type Express struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int64            `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Name           string           `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	ChargingMethod int32            `protobuf:"varint,3,opt,name=charging_method,json=chargingMethod,proto3" json:"charging_method"`
	Defaulted      bool             `protobuf:"varint,4,opt,name=defaulted,proto3" json:"defaulted"`
	CreatedAt      string           `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt      string           `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	Details        []*ExpressDetail `protobuf:"bytes,7,rep,name=details,proto3" json:"details"`
}

func (x *Express) Reset() {
	*x = Express{}
	if protoimpl.UnsafeEnabled {
		mi := &file_expressService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Express) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Express) ProtoMessage() {}

func (x *Express) ProtoReflect() protoreflect.Message {
	mi := &file_expressService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Express.ProtoReflect.Descriptor instead.
func (*Express) Descriptor() ([]byte, []int) {
	return file_expressService_proto_rawDescGZIP(), []int{1}
}

func (x *Express) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Express) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Express) GetChargingMethod() int32 {
	if x != nil {
		return x.ChargingMethod
	}
	return 0
}

func (x *Express) GetDefaulted() bool {
	if x != nil {
		return x.Defaulted
	}
	return false
}

func (x *Express) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Express) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Express) GetDetails() []*ExpressDetail {
	if x != nil {
		return x.Details
	}
	return nil
}

type ExpressDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               int64          `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	ExpressId        int64          `protobuf:"varint,2,opt,name=express_id,json=expressId,proto3" json:"express_id"`
	FirstWeight      float32        `protobuf:"fixed32,3,opt,name=first_weight,json=firstWeight,proto3" json:"first_weight"`
	FirstNumber      int32          `protobuf:"varint,4,opt,name=first_number,json=firstNumber,proto3" json:"first_number"`
	ExpressFee       float32        `protobuf:"fixed32,5,opt,name=express_fee,json=expressFee,proto3" json:"express_fee"`
	AdditionalWeight float32        `protobuf:"fixed32,6,opt,name=additional_weight,json=additionalWeight,proto3" json:"additional_weight"`
	AdditionalNumber int32          `protobuf:"varint,7,opt,name=additional_number,json=additionalNumber,proto3" json:"additional_number"`
	AdditionalFee    float32        `protobuf:"fixed32,8,opt,name=additional_fee,json=additionalFee,proto3" json:"additional_fee"`
	CreatedAt        string         `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt        string         `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	AreaContent      string         `protobuf:"bytes,11,opt,name=area_content,json=areaContent,proto3" json:"area_content"`
	Areas            []*ExpressArea `protobuf:"bytes,12,rep,name=areas,proto3" json:"areas"`
}

func (x *ExpressDetail) Reset() {
	*x = ExpressDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_expressService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpressDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpressDetail) ProtoMessage() {}

func (x *ExpressDetail) ProtoReflect() protoreflect.Message {
	mi := &file_expressService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpressDetail.ProtoReflect.Descriptor instead.
func (*ExpressDetail) Descriptor() ([]byte, []int) {
	return file_expressService_proto_rawDescGZIP(), []int{2}
}

func (x *ExpressDetail) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ExpressDetail) GetExpressId() int64 {
	if x != nil {
		return x.ExpressId
	}
	return 0
}

func (x *ExpressDetail) GetFirstWeight() float32 {
	if x != nil {
		return x.FirstWeight
	}
	return 0
}

func (x *ExpressDetail) GetFirstNumber() int32 {
	if x != nil {
		return x.FirstNumber
	}
	return 0
}

func (x *ExpressDetail) GetExpressFee() float32 {
	if x != nil {
		return x.ExpressFee
	}
	return 0
}

func (x *ExpressDetail) GetAdditionalWeight() float32 {
	if x != nil {
		return x.AdditionalWeight
	}
	return 0
}

func (x *ExpressDetail) GetAdditionalNumber() int32 {
	if x != nil {
		return x.AdditionalNumber
	}
	return 0
}

func (x *ExpressDetail) GetAdditionalFee() float32 {
	if x != nil {
		return x.AdditionalFee
	}
	return 0
}

func (x *ExpressDetail) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *ExpressDetail) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *ExpressDetail) GetAreaContent() string {
	if x != nil {
		return x.AreaContent
	}
	return ""
}

func (x *ExpressDetail) GetAreas() []*ExpressArea {
	if x != nil {
		return x.Areas
	}
	return nil
}

type ExpressArea struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	ExpressDetailId int64  `protobuf:"varint,2,opt,name=express_detail_id,json=expressDetailId,proto3" json:"express_detail_id"`
	AreaId          int64  `protobuf:"varint,3,opt,name=area_id,json=areaId,proto3" json:"area_id"`
	CreatedAt       string `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt       string `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
}

func (x *ExpressArea) Reset() {
	*x = ExpressArea{}
	if protoimpl.UnsafeEnabled {
		mi := &file_expressService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpressArea) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpressArea) ProtoMessage() {}

func (x *ExpressArea) ProtoReflect() protoreflect.Message {
	mi := &file_expressService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpressArea.ProtoReflect.Descriptor instead.
func (*ExpressArea) Descriptor() ([]byte, []int) {
	return file_expressService_proto_rawDescGZIP(), []int{3}
}

func (x *ExpressArea) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ExpressArea) GetExpressDetailId() int64 {
	if x != nil {
		return x.ExpressDetailId
	}
	return 0
}

func (x *ExpressArea) GetAreaId() int64 {
	if x != nil {
		return x.AreaId
	}
	return 0
}

func (x *ExpressArea) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *ExpressArea) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type ExpressData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Express      `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*Express    `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Info   *common.Info  `protobuf:"bytes,4,opt,name=info,proto3" json:"info"`
}

func (x *ExpressData) Reset() {
	*x = ExpressData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_expressService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpressData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpressData) ProtoMessage() {}

func (x *ExpressData) ProtoReflect() protoreflect.Message {
	mi := &file_expressService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpressData.ProtoReflect.Descriptor instead.
func (*ExpressData) Descriptor() ([]byte, []int) {
	return file_expressService_proto_rawDescGZIP(), []int{4}
}

func (x *ExpressData) GetEntity() *Express {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *ExpressData) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *ExpressData) GetItems() []*Express {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ExpressData) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type ExpressResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *ExpressData  `protobuf:"bytes,1,opt,name=data,proto3" json:"data"`
	Error *common.Error `protobuf:"bytes,2,opt,name=error,proto3" json:"error"`
}

func (x *ExpressResponse) Reset() {
	*x = ExpressResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_expressService_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpressResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpressResponse) ProtoMessage() {}

func (x *ExpressResponse) ProtoReflect() protoreflect.Message {
	mi := &file_expressService_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpressResponse.ProtoReflect.Descriptor instead.
func (*ExpressResponse) Descriptor() ([]byte, []int) {
	return file_expressService_proto_rawDescGZIP(), []int{5}
}

func (x *ExpressResponse) GetData() *ExpressData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ExpressResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_expressService_proto protoreflect.FileDescriptor

var file_expressService_proto_rawDesc = []byte{
	0x0a, 0x14, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xdb, 0x01, 0x0a, 0x0e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6f, 0x72,
	0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6f, 0x72, 0x74,
	0x69, 0x6e, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73,
	0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x73, 0x18,
	0x08, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x49, 0x64,
	0x73, 0x22, 0xe5, 0x01, 0x0a, 0x07, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x68, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x63, 0x68, 0x61, 0x72,
	0x67, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x64,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x31, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0xb4, 0x03, 0x0a, 0x0d, 0x45, 0x78,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x65,
	0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x5f, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x0b, 0x66, 0x69, 0x72, 0x73, 0x74, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x21, 0x0a,
	0x0c, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0b, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x66, 0x65, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x46, 0x65,
	0x65, 0x12, 0x2b, 0x0a, 0x11, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f,
	0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x10, 0x61, 0x64,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x2b,
	0x0a, 0x11, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x61, 0x64, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x61,
	0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x66, 0x65, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x0d, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x46,
	0x65, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x21, 0x0a, 0x0c, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x72, 0x65, 0x61, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x12, 0x2b, 0x0a, 0x05, 0x61, 0x72, 0x65, 0x61, 0x73, 0x18, 0x0c, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x45, 0x78,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x41, 0x72, 0x65, 0x61, 0x52, 0x05, 0x61, 0x72, 0x65, 0x61, 0x73,
	0x22, 0xa0, 0x01, 0x0a, 0x0b, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x41, 0x72, 0x65, 0x61,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x2a, 0x0a, 0x11, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x65, 0x78, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x61, 0x72, 0x65, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61,
	0x72, 0x65, 0x61, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0xa8, 0x01, 0x0a, 0x0b, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x29, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x45,
	0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23,
	0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61,
	0x67, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x45, 0x78,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x20, 0x0a, 0x04,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x61,
	0x0a, 0x0f, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x29, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x32, 0xc4, 0x04, 0x0a, 0x0e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x11,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x45, 0x78, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38,
	0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x1a, 0x19, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x04, 0x43, 0x6f, 0x70, 0x79,
	0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x45, 0x78, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x45, 0x78, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x45, 0x78,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x04, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x45, 0x78, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12,
	0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x45, 0x78,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x42, 0x0a, 0x09, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x65, 0x64, 0x12, 0x18, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x45, 0x78, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_expressService_proto_rawDescOnce sync.Once
	file_expressService_proto_rawDescData = file_expressService_proto_rawDesc
)

func file_expressService_proto_rawDescGZIP() []byte {
	file_expressService_proto_rawDescOnce.Do(func() {
		file_expressService_proto_rawDescData = protoimpl.X.CompressGZIP(file_expressService_proto_rawDescData)
	})
	return file_expressService_proto_rawDescData
}

var file_expressService_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_expressService_proto_goTypes = []interface{}{
	(*ExpressRequest)(nil),  // 0: services.ExpressRequest
	(*Express)(nil),         // 1: services.Express
	(*ExpressDetail)(nil),   // 2: services.ExpressDetail
	(*ExpressArea)(nil),     // 3: services.ExpressArea
	(*ExpressData)(nil),     // 4: services.ExpressData
	(*ExpressResponse)(nil), // 5: services.ExpressResponse
	(*common.Pager)(nil),    // 6: common.Pager
	(*common.Info)(nil),     // 7: common.Info
	(*common.Error)(nil),    // 8: common.Error
}
var file_expressService_proto_depIdxs = []int32{
	2,  // 0: services.Express.details:type_name -> services.ExpressDetail
	3,  // 1: services.ExpressDetail.areas:type_name -> services.ExpressArea
	1,  // 2: services.ExpressData.entity:type_name -> services.Express
	6,  // 3: services.ExpressData.pager:type_name -> common.Pager
	1,  // 4: services.ExpressData.items:type_name -> services.Express
	7,  // 5: services.ExpressData.info:type_name -> common.Info
	4,  // 6: services.ExpressResponse.data:type_name -> services.ExpressData
	8,  // 7: services.ExpressResponse.error:type_name -> common.Error
	1,  // 8: services.ExpressService.Create:input_type -> services.Express
	1,  // 9: services.ExpressService.Update:input_type -> services.Express
	0,  // 10: services.ExpressService.Copy:input_type -> services.ExpressRequest
	0,  // 11: services.ExpressService.Delete:input_type -> services.ExpressRequest
	0,  // 12: services.ExpressService.Search:input_type -> services.ExpressRequest
	0,  // 13: services.ExpressService.List:input_type -> services.ExpressRequest
	1,  // 14: services.ExpressService.Get:input_type -> services.Express
	0,  // 15: services.ExpressService.Defaulted:input_type -> services.ExpressRequest
	0,  // 16: services.ExpressService.GetDefault:input_type -> services.ExpressRequest
	5,  // 17: services.ExpressService.Create:output_type -> services.ExpressResponse
	5,  // 18: services.ExpressService.Update:output_type -> services.ExpressResponse
	5,  // 19: services.ExpressService.Copy:output_type -> services.ExpressResponse
	5,  // 20: services.ExpressService.Delete:output_type -> services.ExpressResponse
	5,  // 21: services.ExpressService.Search:output_type -> services.ExpressResponse
	5,  // 22: services.ExpressService.List:output_type -> services.ExpressResponse
	5,  // 23: services.ExpressService.Get:output_type -> services.ExpressResponse
	5,  // 24: services.ExpressService.Defaulted:output_type -> services.ExpressResponse
	5,  // 25: services.ExpressService.GetDefault:output_type -> services.ExpressResponse
	17, // [17:26] is the sub-list for method output_type
	8,  // [8:17] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_expressService_proto_init() }
func file_expressService_proto_init() {
	if File_expressService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_expressService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpressRequest); i {
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
		file_expressService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Express); i {
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
		file_expressService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpressDetail); i {
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
		file_expressService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpressArea); i {
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
		file_expressService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpressData); i {
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
		file_expressService_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpressResponse); i {
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
			RawDescriptor: file_expressService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_expressService_proto_goTypes,
		DependencyIndexes: file_expressService_proto_depIdxs,
		MessageInfos:      file_expressService_proto_msgTypes,
	}.Build()
	File_expressService_proto = out.File
	file_expressService_proto_rawDesc = nil
	file_expressService_proto_goTypes = nil
	file_expressService_proto_depIdxs = nil
}