// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: screeningPriceService.proto

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

// 价格筛选
type ScreeningPrice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`                               //ID
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`                            //名称
	MinPrice  int64  `protobuf:"varint,3,opt,name=min_price,json=minPrice,proto3" json:"min_price"`   //最小价格
	MaxPrice  int64  `protobuf:"varint,4,opt,name=max_price,json=maxPrice,proto3" json:"max_price"`   //最大价格
	Sort      int32  `protobuf:"varint,5,opt,name=sort,proto3" json:"sort"`                           //顺序
	Desc      string `protobuf:"bytes,6,opt,name=desc,proto3" json:"desc"`                            //描述
	IsEnable  string `protobuf:"bytes,7,opt,name=is_enable,json=isEnable,proto3" json:"is_enable"`    //是否启用
	CreatedAt string `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at"` //创建时间
	UpdatedAt string `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"` //修改时间
}

func (x *ScreeningPrice) Reset() {
	*x = ScreeningPrice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_screeningPriceService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScreeningPrice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScreeningPrice) ProtoMessage() {}

func (x *ScreeningPrice) ProtoReflect() protoreflect.Message {
	mi := &file_screeningPriceService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScreeningPrice.ProtoReflect.Descriptor instead.
func (*ScreeningPrice) Descriptor() ([]byte, []int) {
	return file_screeningPriceService_proto_rawDescGZIP(), []int{0}
}

func (x *ScreeningPrice) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ScreeningPrice) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ScreeningPrice) GetMinPrice() int64 {
	if x != nil {
		return x.MinPrice
	}
	return 0
}

func (x *ScreeningPrice) GetMaxPrice() int64 {
	if x != nil {
		return x.MaxPrice
	}
	return 0
}

func (x *ScreeningPrice) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *ScreeningPrice) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *ScreeningPrice) GetIsEnable() string {
	if x != nil {
		return x.IsEnable
	}
	return ""
}

func (x *ScreeningPrice) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *ScreeningPrice) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

// 价格筛选请求参数
type ScreeningPriceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Top       int32    `protobuf:"varint,1,opt,name=top,proto3" json:"top"`
	Paged     int64    `protobuf:"varint,2,opt,name=paged,proto3" json:"paged"`
	PageSize  int64    `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Keywords  string   `protobuf:"bytes,4,opt,name=keywords,proto3" json:"keywords"`
	Sorts     []string `protobuf:"bytes,5,rep,name=sorts,proto3" json:"sorts"`
	DateRange []string `protobuf:"bytes,6,rep,name=date_range,json=dateRange,proto3" json:"date_range"`
	Ids       []int64  `protobuf:"varint,7,rep,packed,name=ids,proto3" json:"ids"`
	Id        int64    `protobuf:"varint,8,opt,name=id,proto3" json:"id"`
	//以下为自定义参数
	Name     string `protobuf:"bytes,11,opt,name=name,proto3" json:"name"`                         //名称
	IsEnable string `protobuf:"bytes,12,opt,name=is_enable,json=isEnable,proto3" json:"is_enable"` //是否启用
}

func (x *ScreeningPriceRequest) Reset() {
	*x = ScreeningPriceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_screeningPriceService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScreeningPriceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScreeningPriceRequest) ProtoMessage() {}

func (x *ScreeningPriceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_screeningPriceService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScreeningPriceRequest.ProtoReflect.Descriptor instead.
func (*ScreeningPriceRequest) Descriptor() ([]byte, []int) {
	return file_screeningPriceService_proto_rawDescGZIP(), []int{1}
}

func (x *ScreeningPriceRequest) GetTop() int32 {
	if x != nil {
		return x.Top
	}
	return 0
}

func (x *ScreeningPriceRequest) GetPaged() int64 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *ScreeningPriceRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ScreeningPriceRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *ScreeningPriceRequest) GetSorts() []string {
	if x != nil {
		return x.Sorts
	}
	return nil
}

func (x *ScreeningPriceRequest) GetDateRange() []string {
	if x != nil {
		return x.DateRange
	}
	return nil
}

func (x *ScreeningPriceRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *ScreeningPriceRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ScreeningPriceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ScreeningPriceRequest) GetIsEnable() string {
	if x != nil {
		return x.IsEnable
	}
	return ""
}

// 价格筛选响应数据
type ScreeningPriceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg    string            `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg"`
	Pager  *common.Pager     `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Entity *ScreeningPrice   `protobuf:"bytes,3,opt,name=entity,proto3" json:"entity"`
	Items  []*ScreeningPrice `protobuf:"bytes,4,rep,name=items,proto3" json:"items"`
}

func (x *ScreeningPriceResponse) Reset() {
	*x = ScreeningPriceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_screeningPriceService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScreeningPriceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScreeningPriceResponse) ProtoMessage() {}

func (x *ScreeningPriceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_screeningPriceService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScreeningPriceResponse.ProtoReflect.Descriptor instead.
func (*ScreeningPriceResponse) Descriptor() ([]byte, []int) {
	return file_screeningPriceService_proto_rawDescGZIP(), []int{2}
}

func (x *ScreeningPriceResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *ScreeningPriceResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *ScreeningPriceResponse) GetEntity() *ScreeningPrice {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *ScreeningPriceResponse) GetItems() []*ScreeningPrice {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_screeningPriceService_proto protoreflect.FileDescriptor

var file_screeningPriceService_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf1, 0x01, 0x0a, 0x0e, 0x53,
	0x63, 0x72, 0x65, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x69, 0x6e, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x69, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64,
	0x65, 0x73, 0x63, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x80,
	0x02, 0x0a, 0x15, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x6f, 0x70, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x74, 0x6f, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61,
	0x67, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x6f, 0x72,
	0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x73, 0x6f, 0x72, 0x74, 0x73, 0x12,
	0x1d, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x09, 0x64, 0x61, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x22, 0xb1, 0x01, 0x0a, 0x16, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x23,
	0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61,
	0x67, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53,
	0x63, 0x72, 0x65, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x06, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x2e, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x32, 0xd0, 0x03, 0x0a, 0x15, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e,
	0x69, 0x6e, 0x67, 0x50, 0x72, 0x69, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x46, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x1a, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53,
	0x63, 0x72, 0x65, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x63, 0x72,
	0x65, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x69, 0x63, 0x65, 0x1a, 0x20, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x69, 0x6e, 0x67,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x46, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x1a, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53,
	0x63, 0x72, 0x65, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x18,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e,
	0x69, 0x6e, 0x67, 0x50, 0x72, 0x69, 0x63, 0x65, 0x1a, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x06,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x04, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53,
	0x63, 0x72, 0x65, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_screeningPriceService_proto_rawDescOnce sync.Once
	file_screeningPriceService_proto_rawDescData = file_screeningPriceService_proto_rawDesc
)

func file_screeningPriceService_proto_rawDescGZIP() []byte {
	file_screeningPriceService_proto_rawDescOnce.Do(func() {
		file_screeningPriceService_proto_rawDescData = protoimpl.X.CompressGZIP(file_screeningPriceService_proto_rawDescData)
	})
	return file_screeningPriceService_proto_rawDescData
}

var file_screeningPriceService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_screeningPriceService_proto_goTypes = []interface{}{
	(*ScreeningPrice)(nil),         // 0: services.ScreeningPrice
	(*ScreeningPriceRequest)(nil),  // 1: services.ScreeningPriceRequest
	(*ScreeningPriceResponse)(nil), // 2: services.ScreeningPriceResponse
	(*common.Pager)(nil),           // 3: common.Pager
}
var file_screeningPriceService_proto_depIdxs = []int32{
	3, // 0: services.ScreeningPriceResponse.pager:type_name -> common.Pager
	0, // 1: services.ScreeningPriceResponse.entity:type_name -> services.ScreeningPrice
	0, // 2: services.ScreeningPriceResponse.items:type_name -> services.ScreeningPrice
	0, // 3: services.ScreeningPriceService.Create:input_type -> services.ScreeningPrice
	0, // 4: services.ScreeningPriceService.Update:input_type -> services.ScreeningPrice
	0, // 5: services.ScreeningPriceService.Delete:input_type -> services.ScreeningPrice
	0, // 6: services.ScreeningPriceService.Get:input_type -> services.ScreeningPrice
	1, // 7: services.ScreeningPriceService.Search:input_type -> services.ScreeningPriceRequest
	1, // 8: services.ScreeningPriceService.List:input_type -> services.ScreeningPriceRequest
	2, // 9: services.ScreeningPriceService.Create:output_type -> services.ScreeningPriceResponse
	2, // 10: services.ScreeningPriceService.Update:output_type -> services.ScreeningPriceResponse
	2, // 11: services.ScreeningPriceService.Delete:output_type -> services.ScreeningPriceResponse
	2, // 12: services.ScreeningPriceService.Get:output_type -> services.ScreeningPriceResponse
	2, // 13: services.ScreeningPriceService.Search:output_type -> services.ScreeningPriceResponse
	2, // 14: services.ScreeningPriceService.List:output_type -> services.ScreeningPriceResponse
	9, // [9:15] is the sub-list for method output_type
	3, // [3:9] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_screeningPriceService_proto_init() }
func file_screeningPriceService_proto_init() {
	if File_screeningPriceService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_screeningPriceService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScreeningPrice); i {
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
		file_screeningPriceService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScreeningPriceRequest); i {
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
		file_screeningPriceService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScreeningPriceResponse); i {
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
			RawDescriptor: file_screeningPriceService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_screeningPriceService_proto_goTypes,
		DependencyIndexes: file_screeningPriceService_proto_depIdxs,
		MessageInfos:      file_screeningPriceService_proto_msgTypes,
	}.Build()
	File_screeningPriceService_proto = out.File
	file_screeningPriceService_proto_rawDesc = nil
	file_screeningPriceService_proto_goTypes = nil
	file_screeningPriceService_proto_depIdxs = nil
}
