// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: vipPriceService.proto

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

type VipPrice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`                                            //ID
	LevelId       int32  `protobuf:"varint,2,opt,name=level_id,json=levelId,proto3" json:"level_id"`                   //会员等级id
	SpuId         int64  `protobuf:"varint,3,opt,name=spu_id,json=spuId,proto3" json:"spu_id"`                         //SpuId
	Unique        string `protobuf:"bytes,4,opt,name=unique,proto3" json:"unique"`                                     //md5key
	Method        string `protobuf:"bytes,5,opt,name=method,proto3" json:"method"`                                     //优惠方式：1 打折,2 减价, 3 指定价
	DiscountValue int64  `protobuf:"varint,7,opt,name=discount_value,json=discountValue,proto3" json:"discount_value"` //优惠值
	CreatedAt     int64  `protobuf:"varint,9,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt     int64  `protobuf:"varint,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
}

func (x *VipPrice) Reset() {
	*x = VipPrice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vipPriceService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VipPrice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VipPrice) ProtoMessage() {}

func (x *VipPrice) ProtoReflect() protoreflect.Message {
	mi := &file_vipPriceService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VipPrice.ProtoReflect.Descriptor instead.
func (*VipPrice) Descriptor() ([]byte, []int) {
	return file_vipPriceService_proto_rawDescGZIP(), []int{0}
}

func (x *VipPrice) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *VipPrice) GetLevelId() int32 {
	if x != nil {
		return x.LevelId
	}
	return 0
}

func (x *VipPrice) GetSpuId() int64 {
	if x != nil {
		return x.SpuId
	}
	return 0
}

func (x *VipPrice) GetUnique() string {
	if x != nil {
		return x.Unique
	}
	return ""
}

func (x *VipPrice) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *VipPrice) GetDiscountValue() int64 {
	if x != nil {
		return x.DiscountValue
	}
	return 0
}

func (x *VipPrice) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *VipPrice) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

// 会员价格详情
type VipPriceDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SpuId              int64           `protobuf:"varint,1,opt,name=spu_id,json=spuId,proto3" json:"spu_id"`
	FreeDiscountMethod string          `protobuf:"bytes,2,opt,name=free_discount_method,json=freeDiscountMethod,proto3" json:"free_discount_method"`
	PayDiscountMethod  string          `protobuf:"bytes,3,opt,name=pay_discount_method,json=payDiscountMethod,proto3" json:"pay_discount_method"`
	FreeLevels         []*VipLevelInfo `protobuf:"bytes,4,rep,name=free_levels,json=freeLevels,proto3" json:"free_levels"`
	PayLevels          []*VipLevelInfo `protobuf:"bytes,5,rep,name=pay_levels,json=payLevels,proto3" json:"pay_levels"`
	FreeList           []*VipPrice     `protobuf:"bytes,6,rep,name=free_list,json=freeList,proto3" json:"free_list"`
	PayList            []*VipPrice     `protobuf:"bytes,7,rep,name=pay_list,json=payList,proto3" json:"pay_list"`
}

func (x *VipPriceDetail) Reset() {
	*x = VipPriceDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vipPriceService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VipPriceDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VipPriceDetail) ProtoMessage() {}

func (x *VipPriceDetail) ProtoReflect() protoreflect.Message {
	mi := &file_vipPriceService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VipPriceDetail.ProtoReflect.Descriptor instead.
func (*VipPriceDetail) Descriptor() ([]byte, []int) {
	return file_vipPriceService_proto_rawDescGZIP(), []int{1}
}

func (x *VipPriceDetail) GetSpuId() int64 {
	if x != nil {
		return x.SpuId
	}
	return 0
}

func (x *VipPriceDetail) GetFreeDiscountMethod() string {
	if x != nil {
		return x.FreeDiscountMethod
	}
	return ""
}

func (x *VipPriceDetail) GetPayDiscountMethod() string {
	if x != nil {
		return x.PayDiscountMethod
	}
	return ""
}

func (x *VipPriceDetail) GetFreeLevels() []*VipLevelInfo {
	if x != nil {
		return x.FreeLevels
	}
	return nil
}

func (x *VipPriceDetail) GetPayLevels() []*VipLevelInfo {
	if x != nil {
		return x.PayLevels
	}
	return nil
}

func (x *VipPriceDetail) GetFreeList() []*VipPrice {
	if x != nil {
		return x.FreeList
	}
	return nil
}

func (x *VipPriceDetail) GetPayList() []*VipPrice {
	if x != nil {
		return x.PayList
	}
	return nil
}

type VipPriceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64       `protobuf:"varint,1,opt,name=id,proto3" json:"id"`                          //ID
	LevelId int32       `protobuf:"varint,2,opt,name=level_id,json=levelId,proto3" json:"level_id"` //会员等级id
	SpuId   int64       `protobuf:"varint,3,opt,name=spu_id,json=spuId,proto3" json:"spu_id"`       //SpuId
	Unique  string      `protobuf:"bytes,4,opt,name=unique,proto3" json:"unique"`                   //md5key
	Method  string      `protobuf:"bytes,5,opt,name=method,proto3" json:"method"`                   //优惠方式：1 打折,2 减价, 3 指定价
	Ids     []int64     `protobuf:"varint,6,rep,packed,name=ids,proto3" json:"ids"`
	Details []*VipPrice `protobuf:"bytes,7,rep,name=details,proto3" json:"details"`
}

func (x *VipPriceRequest) Reset() {
	*x = VipPriceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vipPriceService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VipPriceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VipPriceRequest) ProtoMessage() {}

func (x *VipPriceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vipPriceService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VipPriceRequest.ProtoReflect.Descriptor instead.
func (*VipPriceRequest) Descriptor() ([]byte, []int) {
	return file_vipPriceService_proto_rawDescGZIP(), []int{2}
}

func (x *VipPriceRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *VipPriceRequest) GetLevelId() int32 {
	if x != nil {
		return x.LevelId
	}
	return 0
}

func (x *VipPriceRequest) GetSpuId() int64 {
	if x != nil {
		return x.SpuId
	}
	return 0
}

func (x *VipPriceRequest) GetUnique() string {
	if x != nil {
		return x.Unique
	}
	return ""
}

func (x *VipPriceRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *VipPriceRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *VipPriceRequest) GetDetails() []*VipPrice {
	if x != nil {
		return x.Details
	}
	return nil
}

type VipPriceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *VipPrice       `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager   `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*VipPrice     `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Msg    string          `protobuf:"bytes,4,opt,name=msg,proto3" json:"msg"`
	Detail *VipPriceDetail `protobuf:"bytes,5,opt,name=detail,proto3" json:"detail"`
}

func (x *VipPriceResponse) Reset() {
	*x = VipPriceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vipPriceService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VipPriceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VipPriceResponse) ProtoMessage() {}

func (x *VipPriceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vipPriceService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VipPriceResponse.ProtoReflect.Descriptor instead.
func (*VipPriceResponse) Descriptor() ([]byte, []int) {
	return file_vipPriceService_proto_rawDescGZIP(), []int{3}
}

func (x *VipPriceResponse) GetEntity() *VipPrice {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *VipPriceResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *VipPriceResponse) GetItems() []*VipPrice {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *VipPriceResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *VipPriceResponse) GetDetail() *VipPriceDetail {
	if x != nil {
		return x.Detail
	}
	return nil
}

var File_vipPriceService_proto protoreflect.FileDescriptor

var file_vipPriceService_proto_rawDesc = []byte{
	0x0a, 0x15, 0x76, 0x69, 0x70, 0x50, 0x72, 0x69, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x62, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe1, 0x01, 0x0a, 0x08,
	0x56, 0x69, 0x70, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x70, 0x75, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x70, 0x75, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x6e,
	0x69, 0x71, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x6e, 0x69, 0x71,
	0x75, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0d, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22,
	0xd9, 0x02, 0x0a, 0x0e, 0x56, 0x69, 0x70, 0x50, 0x72, 0x69, 0x63, 0x65, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x70, 0x75, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x73, 0x70, 0x75, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x14, 0x66, 0x72, 0x65,
	0x65, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x66, 0x72, 0x65, 0x65, 0x44, 0x69, 0x73,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x2e, 0x0a, 0x13, 0x70,
	0x61, 0x79, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x70, 0x61, 0x79, 0x44, 0x69, 0x73,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x37, 0x0a, 0x0b, 0x66,
	0x72, 0x65, 0x65, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x56, 0x69, 0x70, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x66, 0x72, 0x65, 0x65, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x73, 0x12, 0x35, 0x0a, 0x0a, 0x70, 0x61, 0x79, 0x5f, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x56, 0x69, 0x70, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x09, 0x70, 0x61, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x73, 0x12, 0x2f, 0x0a, 0x09, 0x66,
	0x72, 0x65, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x56, 0x69, 0x70, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x52, 0x08, 0x66, 0x72, 0x65, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x08,
	0x70, 0x61, 0x79, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x56, 0x69, 0x70, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x52, 0x07, 0x70, 0x61, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x22, 0xc3, 0x01, 0x0a, 0x0f,
	0x56, 0x69, 0x70, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x19, 0x0a, 0x08, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x70,
	0x75, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x70, 0x75, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03,
	0x69, 0x64, 0x73, 0x12, 0x2c, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x07,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x56, 0x69, 0x70, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x22, 0xd1, 0x01, 0x0a, 0x10, 0x56, 0x69, 0x70, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x56, 0x69, 0x70, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72,
	0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x56, 0x69, 0x70, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x12, 0x30, 0x0a, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x56,
	0x69, 0x70, 0x50, 0x72, 0x69, 0x63, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x06, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x32, 0x99, 0x02, 0x0a, 0x0f, 0x56, 0x69, 0x70, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x04, 0x53, 0x61, 0x76,
	0x65, 0x12, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x56, 0x69, 0x70,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x56, 0x69, 0x70, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x04, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x56, 0x69,
	0x70, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x56, 0x69, 0x70, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x06, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x56, 0x69, 0x70, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x56, 0x69, 0x70, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41,
	0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x56, 0x69, 0x70, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x56,
	0x69, 0x70, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vipPriceService_proto_rawDescOnce sync.Once
	file_vipPriceService_proto_rawDescData = file_vipPriceService_proto_rawDesc
)

func file_vipPriceService_proto_rawDescGZIP() []byte {
	file_vipPriceService_proto_rawDescOnce.Do(func() {
		file_vipPriceService_proto_rawDescData = protoimpl.X.CompressGZIP(file_vipPriceService_proto_rawDescData)
	})
	return file_vipPriceService_proto_rawDescData
}

var file_vipPriceService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_vipPriceService_proto_goTypes = []interface{}{
	(*VipPrice)(nil),         // 0: services.VipPrice
	(*VipPriceDetail)(nil),   // 1: services.VipPriceDetail
	(*VipPriceRequest)(nil),  // 2: services.VipPriceRequest
	(*VipPriceResponse)(nil), // 3: services.VipPriceResponse
	(*VipLevelInfo)(nil),     // 4: services.VipLevelInfo
	(*common.Pager)(nil),     // 5: common.Pager
}
var file_vipPriceService_proto_depIdxs = []int32{
	4,  // 0: services.VipPriceDetail.free_levels:type_name -> services.VipLevelInfo
	4,  // 1: services.VipPriceDetail.pay_levels:type_name -> services.VipLevelInfo
	0,  // 2: services.VipPriceDetail.free_list:type_name -> services.VipPrice
	0,  // 3: services.VipPriceDetail.pay_list:type_name -> services.VipPrice
	0,  // 4: services.VipPriceRequest.details:type_name -> services.VipPrice
	0,  // 5: services.VipPriceResponse.entity:type_name -> services.VipPrice
	5,  // 6: services.VipPriceResponse.pager:type_name -> common.Pager
	0,  // 7: services.VipPriceResponse.items:type_name -> services.VipPrice
	1,  // 8: services.VipPriceResponse.detail:type_name -> services.VipPriceDetail
	2,  // 9: services.VipPriceService.Save:input_type -> services.VipPriceRequest
	2,  // 10: services.VipPriceService.List:input_type -> services.VipPriceRequest
	2,  // 11: services.VipPriceService.Detail:input_type -> services.VipPriceRequest
	2,  // 12: services.VipPriceService.Delete:input_type -> services.VipPriceRequest
	3,  // 13: services.VipPriceService.Save:output_type -> services.VipPriceResponse
	3,  // 14: services.VipPriceService.List:output_type -> services.VipPriceResponse
	3,  // 15: services.VipPriceService.Detail:output_type -> services.VipPriceResponse
	3,  // 16: services.VipPriceService.Delete:output_type -> services.VipPriceResponse
	13, // [13:17] is the sub-list for method output_type
	9,  // [9:13] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_vipPriceService_proto_init() }
func file_vipPriceService_proto_init() {
	if File_vipPriceService_proto != nil {
		return
	}
	file_baseInfoService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_vipPriceService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VipPrice); i {
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
		file_vipPriceService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VipPriceDetail); i {
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
		file_vipPriceService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VipPriceRequest); i {
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
		file_vipPriceService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VipPriceResponse); i {
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
			RawDescriptor: file_vipPriceService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_vipPriceService_proto_goTypes,
		DependencyIndexes: file_vipPriceService_proto_depIdxs,
		MessageInfos:      file_vipPriceService_proto_msgTypes,
	}.Build()
	File_vipPriceService_proto = out.File
	file_vipPriceService_proto_rawDesc = nil
	file_vipPriceService_proto_goTypes = nil
	file_vipPriceService_proto_depIdxs = nil
}
