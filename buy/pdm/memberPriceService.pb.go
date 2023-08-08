// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: memberPriceService.proto

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

type MemberPrice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	OwnerType     string  `protobuf:"bytes,2,opt,name=owner_type,json=ownerType,proto3" json:"owner_type"`
	OwnerId       int64   `protobuf:"varint,3,opt,name=owner_id,json=ownerId,proto3" json:"owner_id"`
	SkuId         int64   `protobuf:"varint,4,opt,name=sku_id,json=skuId,proto3" json:"sku_id"`
	Method        string  `protobuf:"bytes,5,opt,name=method,proto3" json:"method"`
	DiscountRate  float32 `protobuf:"fixed32,6,opt,name=discount_rate,json=discountRate,proto3" json:"discount_rate"`
	DiscountPrice int64   `protobuf:"varint,7,opt,name=discount_price,json=discountPrice,proto3" json:"discount_price"`
	CreatedAt     string  `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt     string  `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
}

func (x *MemberPrice) Reset() {
	*x = MemberPrice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memberPriceService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberPrice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberPrice) ProtoMessage() {}

func (x *MemberPrice) ProtoReflect() protoreflect.Message {
	mi := &file_memberPriceService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberPrice.ProtoReflect.Descriptor instead.
func (*MemberPrice) Descriptor() ([]byte, []int) {
	return file_memberPriceService_proto_rawDescGZIP(), []int{0}
}

func (x *MemberPrice) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MemberPrice) GetOwnerType() string {
	if x != nil {
		return x.OwnerType
	}
	return ""
}

func (x *MemberPrice) GetOwnerId() int64 {
	if x != nil {
		return x.OwnerId
	}
	return 0
}

func (x *MemberPrice) GetSkuId() int64 {
	if x != nil {
		return x.SkuId
	}
	return 0
}

func (x *MemberPrice) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *MemberPrice) GetDiscountRate() float32 {
	if x != nil {
		return x.DiscountRate
	}
	return 0
}

func (x *MemberPrice) GetDiscountPrice() int64 {
	if x != nil {
		return x.DiscountPrice
	}
	return 0
}

func (x *MemberPrice) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *MemberPrice) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type MemberPriceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SpuId   int64          `protobuf:"varint,1,opt,name=spu_id,json=spuId,proto3" json:"spu_id"`
	Method  string         `protobuf:"bytes,2,opt,name=method,proto3" json:"method"`
	Details []*MemberPrice `protobuf:"bytes,3,rep,name=details,proto3" json:"details"`
}

func (x *MemberPriceRequest) Reset() {
	*x = MemberPriceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memberPriceService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberPriceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberPriceRequest) ProtoMessage() {}

func (x *MemberPriceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_memberPriceService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberPriceRequest.ProtoReflect.Descriptor instead.
func (*MemberPriceRequest) Descriptor() ([]byte, []int) {
	return file_memberPriceService_proto_rawDescGZIP(), []int{1}
}

func (x *MemberPriceRequest) GetSpuId() int64 {
	if x != nil {
		return x.SpuId
	}
	return 0
}

func (x *MemberPriceRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *MemberPriceRequest) GetDetails() []*MemberPrice {
	if x != nil {
		return x.Details
	}
	return nil
}

type MemberPriceData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *MemberPrice   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager  `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*MemberPrice `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Info   *common.Info   `protobuf:"bytes,4,opt,name=info,proto3" json:"info"`
	Spu    *Spu           `protobuf:"bytes,5,opt,name=spu,proto3" json:"spu"`
}

func (x *MemberPriceData) Reset() {
	*x = MemberPriceData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memberPriceService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberPriceData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberPriceData) ProtoMessage() {}

func (x *MemberPriceData) ProtoReflect() protoreflect.Message {
	mi := &file_memberPriceService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberPriceData.ProtoReflect.Descriptor instead.
func (*MemberPriceData) Descriptor() ([]byte, []int) {
	return file_memberPriceService_proto_rawDescGZIP(), []int{2}
}

func (x *MemberPriceData) GetEntity() *MemberPrice {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *MemberPriceData) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *MemberPriceData) GetItems() []*MemberPrice {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *MemberPriceData) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *MemberPriceData) GetSpu() *Spu {
	if x != nil {
		return x.Spu
	}
	return nil
}

type MemberPriceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *MemberPriceData `protobuf:"bytes,1,opt,name=data,proto3" json:"data"`
	Error *common.Error    `protobuf:"bytes,2,opt,name=error,proto3" json:"error"`
}

func (x *MemberPriceResponse) Reset() {
	*x = MemberPriceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memberPriceService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberPriceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberPriceResponse) ProtoMessage() {}

func (x *MemberPriceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_memberPriceService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberPriceResponse.ProtoReflect.Descriptor instead.
func (*MemberPriceResponse) Descriptor() ([]byte, []int) {
	return file_memberPriceService_proto_rawDescGZIP(), []int{3}
}

func (x *MemberPriceResponse) GetData() *MemberPriceData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *MemberPriceResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_memberPriceService_proto protoreflect.FileDescriptor

var file_memberPriceService_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x72, 0x69, 0x63, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x73, 0x70, 0x75, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x90, 0x02, 0x0a, 0x0b, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x6b, 0x75, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x6b, 0x75, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x72,
	0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x61, 0x74, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0d, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x74, 0x0a, 0x12,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x70, 0x75, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x73, 0x70, 0x75, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x12, 0x2f, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x22, 0xd5, 0x01, 0x0a, 0x0f, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2d, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x06, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61,
	0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x03, 0x73, 0x70, 0x75,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x53, 0x70, 0x75, 0x52, 0x03, 0x73, 0x70, 0x75, 0x22, 0x69, 0x0a, 0x13, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x50, 0x72, 0x69, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xa0, 0x01, 0x0a, 0x12, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x03,
	0x53, 0x65, 0x74, 0x12, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x44, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_memberPriceService_proto_rawDescOnce sync.Once
	file_memberPriceService_proto_rawDescData = file_memberPriceService_proto_rawDesc
)

func file_memberPriceService_proto_rawDescGZIP() []byte {
	file_memberPriceService_proto_rawDescOnce.Do(func() {
		file_memberPriceService_proto_rawDescData = protoimpl.X.CompressGZIP(file_memberPriceService_proto_rawDescData)
	})
	return file_memberPriceService_proto_rawDescData
}

var file_memberPriceService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_memberPriceService_proto_goTypes = []interface{}{
	(*MemberPrice)(nil),         // 0: services.MemberPrice
	(*MemberPriceRequest)(nil),  // 1: services.MemberPriceRequest
	(*MemberPriceData)(nil),     // 2: services.MemberPriceData
	(*MemberPriceResponse)(nil), // 3: services.MemberPriceResponse
	(*common.Pager)(nil),        // 4: common.Pager
	(*common.Info)(nil),         // 5: common.Info
	(*Spu)(nil),                 // 6: services.Spu
	(*common.Error)(nil),        // 7: common.Error
}
var file_memberPriceService_proto_depIdxs = []int32{
	0,  // 0: services.MemberPriceRequest.details:type_name -> services.MemberPrice
	0,  // 1: services.MemberPriceData.entity:type_name -> services.MemberPrice
	4,  // 2: services.MemberPriceData.pager:type_name -> common.Pager
	0,  // 3: services.MemberPriceData.items:type_name -> services.MemberPrice
	5,  // 4: services.MemberPriceData.info:type_name -> common.Info
	6,  // 5: services.MemberPriceData.spu:type_name -> services.Spu
	2,  // 6: services.MemberPriceResponse.data:type_name -> services.MemberPriceData
	7,  // 7: services.MemberPriceResponse.error:type_name -> common.Error
	1,  // 8: services.MemberPriceService.Set:input_type -> services.MemberPriceRequest
	1,  // 9: services.MemberPriceService.Get:input_type -> services.MemberPriceRequest
	3,  // 10: services.MemberPriceService.Set:output_type -> services.MemberPriceResponse
	3,  // 11: services.MemberPriceService.Get:output_type -> services.MemberPriceResponse
	10, // [10:12] is the sub-list for method output_type
	8,  // [8:10] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_memberPriceService_proto_init() }
func file_memberPriceService_proto_init() {
	if File_memberPriceService_proto != nil {
		return
	}
	file_spuService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_memberPriceService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberPrice); i {
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
		file_memberPriceService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberPriceRequest); i {
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
		file_memberPriceService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberPriceData); i {
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
		file_memberPriceService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberPriceResponse); i {
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
			RawDescriptor: file_memberPriceService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_memberPriceService_proto_goTypes,
		DependencyIndexes: file_memberPriceService_proto_depIdxs,
		MessageInfos:      file_memberPriceService_proto_msgTypes,
	}.Build()
	File_memberPriceService_proto = out.File
	file_memberPriceService_proto_rawDesc = nil
	file_memberPriceService_proto_goTypes = nil
	file_memberPriceService_proto_depIdxs = nil
}
