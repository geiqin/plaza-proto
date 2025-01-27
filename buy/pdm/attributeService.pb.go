// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: attributeService.proto

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

// 属性库
type Attribute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64             `protobuf:"varint,1,opt,name=id,proto3" json:"id"`                                          //ID
	Name         string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`                                       //属性项名
	IsBuyShow    string            `protobuf:"bytes,3,opt,name=is_buy_show,json=isBuyShow,proto3" json:"is_buy_show"`          //下单页展示
	IsMultiValue string            `protobuf:"bytes,4,opt,name=is_multi_value,json=isMultiValue,proto3" json:"is_multi_value"` //属性值多选
	IsRequired   string            `protobuf:"bytes,5,opt,name=is_required,json=isRequired,proto3" json:"is_required"`         //属性值必选
	Values       []*AttributeValue `protobuf:"bytes,6,rep,name=values,proto3" json:"values"`                                   //属性值列表
	SpuCount     int32             `protobuf:"varint,7,opt,name=spu_count,json=spuCount,proto3" json:"spu_count"`              //商品数量
	CreatedAt    string            `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at"`            //创建时间
	UpdatedAt    string            `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`            //修改时间
}

func (x *Attribute) Reset() {
	*x = Attribute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attributeService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Attribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attribute) ProtoMessage() {}

func (x *Attribute) ProtoReflect() protoreflect.Message {
	mi := &file_attributeService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attribute.ProtoReflect.Descriptor instead.
func (*Attribute) Descriptor() ([]byte, []int) {
	return file_attributeService_proto_rawDescGZIP(), []int{0}
}

func (x *Attribute) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Attribute) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Attribute) GetIsBuyShow() string {
	if x != nil {
		return x.IsBuyShow
	}
	return ""
}

func (x *Attribute) GetIsMultiValue() string {
	if x != nil {
		return x.IsMultiValue
	}
	return ""
}

func (x *Attribute) GetIsRequired() string {
	if x != nil {
		return x.IsRequired
	}
	return ""
}

func (x *Attribute) GetValues() []*AttributeValue {
	if x != nil {
		return x.Values
	}
	return nil
}

func (x *Attribute) GetSpuCount() int32 {
	if x != nil {
		return x.SpuCount
	}
	return 0
}

func (x *Attribute) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Attribute) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type AttributeValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
	AddPrice int64  `protobuf:"varint,2,opt,name=add_price,json=addPrice,proto3" json:"add_price"`
}

func (x *AttributeValue) Reset() {
	*x = AttributeValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attributeService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttributeValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttributeValue) ProtoMessage() {}

func (x *AttributeValue) ProtoReflect() protoreflect.Message {
	mi := &file_attributeService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttributeValue.ProtoReflect.Descriptor instead.
func (*AttributeValue) Descriptor() ([]byte, []int) {
	return file_attributeService_proto_rawDescGZIP(), []int{1}
}

func (x *AttributeValue) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AttributeValue) GetAddPrice() int64 {
	if x != nil {
		return x.AddPrice
	}
	return 0
}

// 属性库请求参数
type AttributeRequest struct {
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
	Name         string `protobuf:"bytes,11,opt,name=name,proto3" json:"name"`                                       //属性项名
	IsBuyShow    string `protobuf:"bytes,12,opt,name=is_buy_show,json=isBuyShow,proto3" json:"is_buy_show"`          //下单页展示
	IsMultiValue string `protobuf:"bytes,13,opt,name=is_multi_value,json=isMultiValue,proto3" json:"is_multi_value"` //属性值多选
	IsRequired   string `protobuf:"bytes,14,opt,name=is_required,json=isRequired,proto3" json:"is_required"`         //属性值必选
}

func (x *AttributeRequest) Reset() {
	*x = AttributeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attributeService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttributeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttributeRequest) ProtoMessage() {}

func (x *AttributeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_attributeService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttributeRequest.ProtoReflect.Descriptor instead.
func (*AttributeRequest) Descriptor() ([]byte, []int) {
	return file_attributeService_proto_rawDescGZIP(), []int{2}
}

func (x *AttributeRequest) GetTop() int32 {
	if x != nil {
		return x.Top
	}
	return 0
}

func (x *AttributeRequest) GetPaged() int64 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *AttributeRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *AttributeRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *AttributeRequest) GetSorts() []string {
	if x != nil {
		return x.Sorts
	}
	return nil
}

func (x *AttributeRequest) GetDateRange() []string {
	if x != nil {
		return x.DateRange
	}
	return nil
}

func (x *AttributeRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *AttributeRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AttributeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AttributeRequest) GetIsBuyShow() string {
	if x != nil {
		return x.IsBuyShow
	}
	return ""
}

func (x *AttributeRequest) GetIsMultiValue() string {
	if x != nil {
		return x.IsMultiValue
	}
	return ""
}

func (x *AttributeRequest) GetIsRequired() string {
	if x != nil {
		return x.IsRequired
	}
	return ""
}

// 属性库响应数据
type AttributeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg    string        `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Entity *Attribute    `protobuf:"bytes,3,opt,name=entity,proto3" json:"entity"`
	Items  []*Attribute  `protobuf:"bytes,4,rep,name=items,proto3" json:"items"`
}

func (x *AttributeResponse) Reset() {
	*x = AttributeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attributeService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttributeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttributeResponse) ProtoMessage() {}

func (x *AttributeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_attributeService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttributeResponse.ProtoReflect.Descriptor instead.
func (*AttributeResponse) Descriptor() ([]byte, []int) {
	return file_attributeService_proto_rawDescGZIP(), []int{3}
}

func (x *AttributeResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *AttributeResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *AttributeResponse) GetEntity() *Attribute {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *AttributeResponse) GetItems() []*Attribute {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_attributeService_proto protoreflect.FileDescriptor

var file_attributeService_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa3, 0x02, 0x0a, 0x09, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x62, 0x75,
	0x79, 0x5f, 0x73, 0x68, 0x6f, 0x77, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x73,
	0x42, 0x75, 0x79, 0x53, 0x68, 0x6f, 0x77, 0x12, 0x24, 0x0a, 0x0e, 0x69, 0x73, 0x5f, 0x6d, 0x75,
	0x6c, 0x74, 0x69, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x69, 0x73, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x69, 0x73, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x69, 0x73, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x30,
	0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x12, 0x1b, 0x0a, 0x09, 0x73, 0x70, 0x75, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x70, 0x75, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x41, 0x0a, 0x0e, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x64, 0x64, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x61, 0x64, 0x64, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0xc5,
	0x02, 0x0a, 0x10, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x03, 0x74, 0x6f, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77,
	0x6f, 0x72, 0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77,
	0x6f, 0x72, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x05, 0x73, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73,
	0x18, 0x07, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1e, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x62, 0x75, 0x79, 0x5f, 0x73, 0x68, 0x6f, 0x77, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x73, 0x42, 0x75, 0x79, 0x53, 0x68, 0x6f, 0x77, 0x12,
	0x24, 0x0a, 0x0e, 0x69, 0x73, 0x5f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x73, 0x4d, 0x75, 0x6c, 0x74, 0x69,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x72, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0xa2, 0x01, 0x0a, 0x11, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x23,
	0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61,
	0x67, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x12, 0x29, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x32, 0x8f, 0x03, 0x0a, 0x10,
	0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x3c, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x1a,
	0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c,
	0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x1a, 0x1b, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x06,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x1a, 0x1b, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x03, 0x47, 0x65,
	0x74, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12,
	0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x04, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a,
	0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_attributeService_proto_rawDescOnce sync.Once
	file_attributeService_proto_rawDescData = file_attributeService_proto_rawDesc
)

func file_attributeService_proto_rawDescGZIP() []byte {
	file_attributeService_proto_rawDescOnce.Do(func() {
		file_attributeService_proto_rawDescData = protoimpl.X.CompressGZIP(file_attributeService_proto_rawDescData)
	})
	return file_attributeService_proto_rawDescData
}

var file_attributeService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_attributeService_proto_goTypes = []interface{}{
	(*Attribute)(nil),         // 0: services.Attribute
	(*AttributeValue)(nil),    // 1: services.AttributeValue
	(*AttributeRequest)(nil),  // 2: services.AttributeRequest
	(*AttributeResponse)(nil), // 3: services.AttributeResponse
	(*common.Pager)(nil),      // 4: common.Pager
}
var file_attributeService_proto_depIdxs = []int32{
	1,  // 0: services.Attribute.values:type_name -> services.AttributeValue
	4,  // 1: services.AttributeResponse.pager:type_name -> common.Pager
	0,  // 2: services.AttributeResponse.entity:type_name -> services.Attribute
	0,  // 3: services.AttributeResponse.items:type_name -> services.Attribute
	0,  // 4: services.AttributeService.Create:input_type -> services.Attribute
	0,  // 5: services.AttributeService.Update:input_type -> services.Attribute
	0,  // 6: services.AttributeService.Delete:input_type -> services.Attribute
	0,  // 7: services.AttributeService.Get:input_type -> services.Attribute
	2,  // 8: services.AttributeService.Search:input_type -> services.AttributeRequest
	2,  // 9: services.AttributeService.List:input_type -> services.AttributeRequest
	3,  // 10: services.AttributeService.Create:output_type -> services.AttributeResponse
	3,  // 11: services.AttributeService.Update:output_type -> services.AttributeResponse
	3,  // 12: services.AttributeService.Delete:output_type -> services.AttributeResponse
	3,  // 13: services.AttributeService.Get:output_type -> services.AttributeResponse
	3,  // 14: services.AttributeService.Search:output_type -> services.AttributeResponse
	3,  // 15: services.AttributeService.List:output_type -> services.AttributeResponse
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_attributeService_proto_init() }
func file_attributeService_proto_init() {
	if File_attributeService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_attributeService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Attribute); i {
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
		file_attributeService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttributeValue); i {
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
		file_attributeService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttributeRequest); i {
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
		file_attributeService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttributeResponse); i {
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
			RawDescriptor: file_attributeService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_attributeService_proto_goTypes,
		DependencyIndexes: file_attributeService_proto_depIdxs,
		MessageInfos:      file_attributeService_proto_msgTypes,
	}.Build()
	File_attributeService_proto = out.File
	file_attributeService_proto_rawDesc = nil
	file_attributeService_proto_goTypes = nil
	file_attributeService_proto_depIdxs = nil
}
