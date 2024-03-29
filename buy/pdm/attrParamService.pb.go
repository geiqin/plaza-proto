// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: attrParamService.proto

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

type AttrParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int32            `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Name          string           `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	CategoryId    int64            `protobuf:"varint,3,opt,name=category_id,json=categoryId,proto3" json:"category_id"`
	GroupId       int32            `protobuf:"varint,4,opt,name=group_id,json=groupId,proto3" json:"group_id"`
	Type          string           `protobuf:"bytes,5,opt,name=type,proto3" json:"type"`
	InputType     string           `protobuf:"bytes,6,opt,name=input_type,json=inputType,proto3" json:"input_type"`
	IsAllowAlias  string           `protobuf:"bytes,7,opt,name=is_allow_alias,json=isAllowAlias,proto3" json:"is_allow_alias"`
	IsSale        string           `protobuf:"bytes,8,opt,name=is_sale,json=isSale,proto3" json:"is_sale"`
	IsNumeric     string           `protobuf:"bytes,9,opt,name=is_numeric,json=isNumeric,proto3" json:"is_numeric"`
	IsKey         string           `protobuf:"bytes,10,opt,name=is_key,json=isKey,proto3" json:"is_key"`
	IsSearch      string           `protobuf:"bytes,11,opt,name=is_search,json=isSearch,proto3" json:"is_search"`
	IsDisplay     string           `protobuf:"bytes,12,opt,name=is_display,json=isDisplay,proto3" json:"is_display"`
	IsMust        string           `protobuf:"bytes,13,opt,name=is_must,json=isMust,proto3" json:"is_must"`
	IsMulti       string           `protobuf:"bytes,14,opt,name=is_multi,json=isMulti,proto3" json:"is_multi"`
	Unit          string           `protobuf:"bytes,15,opt,name=unit,proto3" json:"unit"`
	Sort          int32            `protobuf:"varint,16,opt,name=sort,proto3" json:"sort"`
	GoodsCount    int32            `protobuf:"varint,17,opt,name=goods_count,json=goodsCount,proto3" json:"goods_count"`
	Status        string           `protobuf:"bytes,18,opt,name=status,proto3" json:"status"`
	CreatedAt     string           `protobuf:"bytes,19,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt     string           `protobuf:"bytes,20,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	Group         *AttrGroup       `protobuf:"bytes,21,opt,name=group,proto3" json:"group"`
	Category      *Category        `protobuf:"bytes,22,opt,name=category,proto3" json:"category"`
	Values        []string         `protobuf:"bytes,23,rep,name=values,proto3" json:"values"`
	ServicePrices map[string]int64 `protobuf:"bytes,24,rep,name=service_prices,json=servicePrices,proto3" json:"service_prices" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"` //服务价清单
}

func (x *AttrParam) Reset() {
	*x = AttrParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attrParamService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttrParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttrParam) ProtoMessage() {}

func (x *AttrParam) ProtoReflect() protoreflect.Message {
	mi := &file_attrParamService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttrParam.ProtoReflect.Descriptor instead.
func (*AttrParam) Descriptor() ([]byte, []int) {
	return file_attrParamService_proto_rawDescGZIP(), []int{0}
}

func (x *AttrParam) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AttrParam) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AttrParam) GetCategoryId() int64 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

func (x *AttrParam) GetGroupId() int32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *AttrParam) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *AttrParam) GetInputType() string {
	if x != nil {
		return x.InputType
	}
	return ""
}

func (x *AttrParam) GetIsAllowAlias() string {
	if x != nil {
		return x.IsAllowAlias
	}
	return ""
}

func (x *AttrParam) GetIsSale() string {
	if x != nil {
		return x.IsSale
	}
	return ""
}

func (x *AttrParam) GetIsNumeric() string {
	if x != nil {
		return x.IsNumeric
	}
	return ""
}

func (x *AttrParam) GetIsKey() string {
	if x != nil {
		return x.IsKey
	}
	return ""
}

func (x *AttrParam) GetIsSearch() string {
	if x != nil {
		return x.IsSearch
	}
	return ""
}

func (x *AttrParam) GetIsDisplay() string {
	if x != nil {
		return x.IsDisplay
	}
	return ""
}

func (x *AttrParam) GetIsMust() string {
	if x != nil {
		return x.IsMust
	}
	return ""
}

func (x *AttrParam) GetIsMulti() string {
	if x != nil {
		return x.IsMulti
	}
	return ""
}

func (x *AttrParam) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *AttrParam) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *AttrParam) GetGoodsCount() int32 {
	if x != nil {
		return x.GoodsCount
	}
	return 0
}

func (x *AttrParam) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *AttrParam) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *AttrParam) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *AttrParam) GetGroup() *AttrGroup {
	if x != nil {
		return x.Group
	}
	return nil
}

func (x *AttrParam) GetCategory() *Category {
	if x != nil {
		return x.Category
	}
	return nil
}

func (x *AttrParam) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

func (x *AttrParam) GetServicePrices() map[string]int64 {
	if x != nil {
		return x.ServicePrices
	}
	return nil
}

type AttrParamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged      int64   `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize   int64   `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Keywords   string  `protobuf:"bytes,3,opt,name=keywords,proto3" json:"keywords"`
	Id         int32   `protobuf:"varint,4,opt,name=id,proto3" json:"id"`
	Name       string  `protobuf:"bytes,5,opt,name=name,proto3" json:"name"`
	CategoryId int32   `protobuf:"varint,6,opt,name=category_id,json=categoryId,proto3" json:"category_id"`
	GroupId    int32   `protobuf:"varint,7,opt,name=group_id,json=groupId,proto3" json:"group_id"`
	Type       string  `protobuf:"bytes,8,opt,name=type,proto3" json:"type"`
	ParamId    int32   `protobuf:"varint,9,opt,name=param_id,json=paramId,proto3" json:"param_id"`
	ParamValue string  `protobuf:"bytes,10,opt,name=param_value,json=paramValue,proto3" json:"param_value"`
	ValuePrice float32 `protobuf:"fixed32,11,opt,name=value_price,json=valuePrice,proto3" json:"value_price"`
	Ids        []int32 `protobuf:"varint,12,rep,packed,name=ids,proto3" json:"ids"`
}

func (x *AttrParamRequest) Reset() {
	*x = AttrParamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attrParamService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttrParamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttrParamRequest) ProtoMessage() {}

func (x *AttrParamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_attrParamService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttrParamRequest.ProtoReflect.Descriptor instead.
func (*AttrParamRequest) Descriptor() ([]byte, []int) {
	return file_attrParamService_proto_rawDescGZIP(), []int{1}
}

func (x *AttrParamRequest) GetPaged() int64 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *AttrParamRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *AttrParamRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *AttrParamRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AttrParamRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AttrParamRequest) GetCategoryId() int32 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

func (x *AttrParamRequest) GetGroupId() int32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *AttrParamRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *AttrParamRequest) GetParamId() int32 {
	if x != nil {
		return x.ParamId
	}
	return 0
}

func (x *AttrParamRequest) GetParamValue() string {
	if x != nil {
		return x.ParamValue
	}
	return ""
}

func (x *AttrParamRequest) GetValuePrice() float32 {
	if x != nil {
		return x.ValuePrice
	}
	return 0
}

func (x *AttrParamRequest) GetIds() []int32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type AttrParamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity     *AttrParam    `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	ParamValue string        `protobuf:"bytes,2,opt,name=param_value,json=paramValue,proto3" json:"param_value"`
	Pager      *common.Pager `protobuf:"bytes,3,opt,name=pager,proto3" json:"pager"`
	Items      []*AttrParam  `protobuf:"bytes,4,rep,name=items,proto3" json:"items"`
	Info       string        `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
}

func (x *AttrParamResponse) Reset() {
	*x = AttrParamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attrParamService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttrParamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttrParamResponse) ProtoMessage() {}

func (x *AttrParamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_attrParamService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttrParamResponse.ProtoReflect.Descriptor instead.
func (*AttrParamResponse) Descriptor() ([]byte, []int) {
	return file_attrParamService_proto_rawDescGZIP(), []int{2}
}

func (x *AttrParamResponse) GetEntity() *AttrParam {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *AttrParamResponse) GetParamValue() string {
	if x != nil {
		return x.ParamValue
	}
	return ""
}

func (x *AttrParamResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *AttrParamResponse) GetItems() []*AttrParam {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *AttrParamResponse) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

var File_attrParamService_proto protoreflect.FileDescriptor

var file_attrParamService_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x74, 0x74, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x61, 0x74, 0x74, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x06, 0x0a, 0x09, 0x41, 0x74, 0x74, 0x72, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x69, 0x73, 0x5f, 0x61, 0x6c, 0x6c, 0x6f,
	0x77, 0x5f, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69,
	0x73, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x41, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x69,
	0x73, 0x5f, 0x73, 0x61, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x73,
	0x53, 0x61, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x6e, 0x75, 0x6d, 0x65, 0x72,
	0x69, 0x63, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x73, 0x4e, 0x75, 0x6d, 0x65,
	0x72, 0x69, 0x63, 0x12, 0x15, 0x0a, 0x06, 0x69, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x73, 0x4b, 0x65, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73,
	0x5f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69,
	0x73, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x64, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x73, 0x44,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x6d, 0x75, 0x73,
	0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x73, 0x4d, 0x75, 0x73, 0x74, 0x12,
	0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x69, 0x73, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x6e,
	0x69, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x6f,
	0x72, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x11, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x12, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x29, 0x0a, 0x05, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x05, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x12, 0x2e, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x18, 0x16, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x17,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x4d, 0x0a, 0x0e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x73, 0x18, 0x18,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x41, 0x74, 0x74, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x40, 0x0a, 0x12, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xc4, 0x02,
	0x0a, 0x10, 0x41, 0x74, 0x74, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64,
	0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f,
	0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x5f, 0x69,
	0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x49, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x05, 0x52,
	0x03, 0x69, 0x64, 0x73, 0x22, 0xc5, 0x01, 0x0a, 0x11, 0x41, 0x74, 0x74, 0x72, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52,
	0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x29, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0xc8, 0x03, 0x0a,
	0x10, 0x41, 0x74, 0x74, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x3a, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x74, 0x74, 0x72,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a,
	0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x1b, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x06, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41,
	0x74, 0x74, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x13, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x74, 0x74,
	0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f,
	0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x74,
	0x74, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x41, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x41, 0x74, 0x74, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x43, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1a,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_attrParamService_proto_rawDescOnce sync.Once
	file_attrParamService_proto_rawDescData = file_attrParamService_proto_rawDesc
)

func file_attrParamService_proto_rawDescGZIP() []byte {
	file_attrParamService_proto_rawDescOnce.Do(func() {
		file_attrParamService_proto_rawDescData = protoimpl.X.CompressGZIP(file_attrParamService_proto_rawDescData)
	})
	return file_attrParamService_proto_rawDescData
}

var file_attrParamService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_attrParamService_proto_goTypes = []interface{}{
	(*AttrParam)(nil),         // 0: services.AttrParam
	(*AttrParamRequest)(nil),  // 1: services.AttrParamRequest
	(*AttrParamResponse)(nil), // 2: services.AttrParamResponse
	nil,                       // 3: services.AttrParam.ServicePricesEntry
	(*AttrGroup)(nil),         // 4: services.AttrGroup
	(*Category)(nil),          // 5: services.Category
	(*common.Pager)(nil),      // 6: common.Pager
}
var file_attrParamService_proto_depIdxs = []int32{
	4,  // 0: services.AttrParam.group:type_name -> services.AttrGroup
	5,  // 1: services.AttrParam.category:type_name -> services.Category
	3,  // 2: services.AttrParam.service_prices:type_name -> services.AttrParam.ServicePricesEntry
	0,  // 3: services.AttrParamResponse.entity:type_name -> services.AttrParam
	6,  // 4: services.AttrParamResponse.pager:type_name -> common.Pager
	0,  // 5: services.AttrParamResponse.items:type_name -> services.AttrParam
	0,  // 6: services.AttrParamService.Create:input_type -> services.AttrParam
	0,  // 7: services.AttrParamService.Update:input_type -> services.AttrParam
	0,  // 8: services.AttrParamService.Delete:input_type -> services.AttrParam
	0,  // 9: services.AttrParamService.Get:input_type -> services.AttrParam
	1,  // 10: services.AttrParamService.List:input_type -> services.AttrParamRequest
	1,  // 11: services.AttrParamService.Search:input_type -> services.AttrParamRequest
	1,  // 12: services.AttrParamService.AddValue:input_type -> services.AttrParamRequest
	2,  // 13: services.AttrParamService.Create:output_type -> services.AttrParamResponse
	2,  // 14: services.AttrParamService.Update:output_type -> services.AttrParamResponse
	2,  // 15: services.AttrParamService.Delete:output_type -> services.AttrParamResponse
	2,  // 16: services.AttrParamService.Get:output_type -> services.AttrParamResponse
	2,  // 17: services.AttrParamService.List:output_type -> services.AttrParamResponse
	2,  // 18: services.AttrParamService.Search:output_type -> services.AttrParamResponse
	2,  // 19: services.AttrParamService.AddValue:output_type -> services.AttrParamResponse
	13, // [13:20] is the sub-list for method output_type
	6,  // [6:13] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_attrParamService_proto_init() }
func file_attrParamService_proto_init() {
	if File_attrParamService_proto != nil {
		return
	}
	file_attrGroupService_proto_init()
	file_categoryService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_attrParamService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttrParam); i {
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
		file_attrParamService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttrParamRequest); i {
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
		file_attrParamService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttrParamResponse); i {
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
			RawDescriptor: file_attrParamService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_attrParamService_proto_goTypes,
		DependencyIndexes: file_attrParamService_proto_depIdxs,
		MessageInfos:      file_attrParamService_proto_msgTypes,
	}.Build()
	File_attrParamService_proto = out.File
	file_attrParamService_proto_rawDesc = nil
	file_attrParamService_proto_goTypes = nil
	file_attrParamService_proto_depIdxs = nil
}
