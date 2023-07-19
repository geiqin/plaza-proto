// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: spuPropertyService.proto

package services

import (
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

//通用属性
type GenericProperty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Params []*GenericPropertyParam `protobuf:"bytes,1,rep,name=params,proto3" json:"params,omitempty"`
	Groups []*GenericPropertyGroup `protobuf:"bytes,2,rep,name=groups,proto3" json:"groups,omitempty"`
}

func (x *GenericProperty) Reset() {
	*x = GenericProperty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spuPropertyService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenericProperty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenericProperty) ProtoMessage() {}

func (x *GenericProperty) ProtoReflect() protoreflect.Message {
	mi := &file_spuPropertyService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenericProperty.ProtoReflect.Descriptor instead.
func (*GenericProperty) Descriptor() ([]byte, []int) {
	return file_spuPropertyService_proto_rawDescGZIP(), []int{0}
}

func (x *GenericProperty) GetParams() []*GenericPropertyParam {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *GenericProperty) GetGroups() []*GenericPropertyGroup {
	if x != nil {
		return x.Groups
	}
	return nil
}

//通用属性(分组)
type GenericPropertyGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId   int32  `protobuf:"varint,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	GroupName string `protobuf:"bytes,2,opt,name=group_name,json=groupName,proto3" json:"group_name,omitempty"`
}

func (x *GenericPropertyGroup) Reset() {
	*x = GenericPropertyGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spuPropertyService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenericPropertyGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenericPropertyGroup) ProtoMessage() {}

func (x *GenericPropertyGroup) ProtoReflect() protoreflect.Message {
	mi := &file_spuPropertyService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenericPropertyGroup.ProtoReflect.Descriptor instead.
func (*GenericPropertyGroup) Descriptor() ([]byte, []int) {
	return file_spuPropertyService_proto_rawDescGZIP(), []int{1}
}

func (x *GenericPropertyGroup) GetGroupId() int32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *GenericPropertyGroup) GetGroupName() string {
	if x != nil {
		return x.GroupName
	}
	return ""
}

//通用属性(参数)
type GenericPropertyParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId    int32  `protobuf:"varint,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	ParamId    int32  `protobuf:"varint,2,opt,name=param_id,json=paramId,proto3" json:"param_id,omitempty"`
	ParamName  string `protobuf:"bytes,3,opt,name=param_name,json=paramName,proto3" json:"param_name,omitempty"`
	ParamValue string `protobuf:"bytes,4,opt,name=param_value,json=paramValue,proto3" json:"param_value,omitempty"`
	ParamUnit  string `protobuf:"bytes,5,opt,name=param_unit,json=paramUnit,proto3" json:"param_unit,omitempty"`
}

func (x *GenericPropertyParam) Reset() {
	*x = GenericPropertyParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spuPropertyService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenericPropertyParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenericPropertyParam) ProtoMessage() {}

func (x *GenericPropertyParam) ProtoReflect() protoreflect.Message {
	mi := &file_spuPropertyService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenericPropertyParam.ProtoReflect.Descriptor instead.
func (*GenericPropertyParam) Descriptor() ([]byte, []int) {
	return file_spuPropertyService_proto_rawDescGZIP(), []int{2}
}

func (x *GenericPropertyParam) GetGroupId() int32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *GenericPropertyParam) GetParamId() int32 {
	if x != nil {
		return x.ParamId
	}
	return 0
}

func (x *GenericPropertyParam) GetParamName() string {
	if x != nil {
		return x.ParamName
	}
	return ""
}

func (x *GenericPropertyParam) GetParamValue() string {
	if x != nil {
		return x.ParamValue
	}
	return ""
}

func (x *GenericPropertyParam) GetParamUnit() string {
	if x != nil {
		return x.ParamUnit
	}
	return ""
}

//服务属性
type ServiceProperty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParamId   int32                       `protobuf:"varint,1,opt,name=param_id,json=paramId,proto3" json:"param_id,omitempty"`
	ParamName string                      `protobuf:"bytes,2,opt,name=param_name,json=paramName,proto3" json:"param_name,omitempty"`
	Values    []*ServicePropertyValueInfo `protobuf:"bytes,3,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *ServiceProperty) Reset() {
	*x = ServiceProperty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spuPropertyService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceProperty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceProperty) ProtoMessage() {}

func (x *ServiceProperty) ProtoReflect() protoreflect.Message {
	mi := &file_spuPropertyService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceProperty.ProtoReflect.Descriptor instead.
func (*ServiceProperty) Descriptor() ([]byte, []int) {
	return file_spuPropertyService_proto_rawDescGZIP(), []int{3}
}

func (x *ServiceProperty) GetParamId() int32 {
	if x != nil {
		return x.ParamId
	}
	return 0
}

func (x *ServiceProperty) GetParamName() string {
	if x != nil {
		return x.ParamName
	}
	return ""
}

func (x *ServiceProperty) GetValues() []*ServicePropertyValueInfo {
	if x != nil {
		return x.Values
	}
	return nil
}

//销售属性
type SaleProperty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SpecId   int32                    `protobuf:"varint,1,opt,name=spec_id,json=specId,proto3" json:"spec_id,omitempty"`
	SpecName string                   `protobuf:"bytes,2,opt,name=spec_name,json=specName,proto3" json:"spec_name,omitempty"`
	Values   []*SalePropertyValueInfo `protobuf:"bytes,3,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *SaleProperty) Reset() {
	*x = SaleProperty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spuPropertyService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaleProperty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaleProperty) ProtoMessage() {}

func (x *SaleProperty) ProtoReflect() protoreflect.Message {
	mi := &file_spuPropertyService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaleProperty.ProtoReflect.Descriptor instead.
func (*SaleProperty) Descriptor() ([]byte, []int) {
	return file_spuPropertyService_proto_rawDescGZIP(), []int{4}
}

func (x *SaleProperty) GetSpecId() int32 {
	if x != nil {
		return x.SpecId
	}
	return 0
}

func (x *SaleProperty) GetSpecName() string {
	if x != nil {
		return x.SpecName
	}
	return ""
}

func (x *SaleProperty) GetValues() []*SalePropertyValueInfo {
	if x != nil {
		return x.Values
	}
	return nil
}

//SKU自己的属性
type OwnProperty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SpecId    int32  `protobuf:"varint,1,opt,name=spec_id,json=specId,proto3" json:"spec_id,omitempty"`
	SpecName  string `protobuf:"bytes,2,opt,name=spec_name,json=specName,proto3" json:"spec_name,omitempty"`
	SpecValue string `protobuf:"bytes,3,opt,name=spec_value,json=specValue,proto3" json:"spec_value,omitempty"`
	ImageId   int64  `protobuf:"varint,4,opt,name=image_id,json=imageId,proto3" json:"image_id,omitempty"`
	ImageUrl  string `protobuf:"bytes,5,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
}

func (x *OwnProperty) Reset() {
	*x = OwnProperty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spuPropertyService_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OwnProperty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OwnProperty) ProtoMessage() {}

func (x *OwnProperty) ProtoReflect() protoreflect.Message {
	mi := &file_spuPropertyService_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OwnProperty.ProtoReflect.Descriptor instead.
func (*OwnProperty) Descriptor() ([]byte, []int) {
	return file_spuPropertyService_proto_rawDescGZIP(), []int{5}
}

func (x *OwnProperty) GetSpecId() int32 {
	if x != nil {
		return x.SpecId
	}
	return 0
}

func (x *OwnProperty) GetSpecName() string {
	if x != nil {
		return x.SpecName
	}
	return ""
}

func (x *OwnProperty) GetSpecValue() string {
	if x != nil {
		return x.SpecValue
	}
	return ""
}

func (x *OwnProperty) GetImageId() int64 {
	if x != nil {
		return x.ImageId
	}
	return 0
}

func (x *OwnProperty) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

type ServicePropertyValueInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ValueText string  `protobuf:"bytes,1,opt,name=value_text,json=valueText,proto3" json:"value_text,omitempty"`
	Price     float32 `protobuf:"fixed32,2,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *ServicePropertyValueInfo) Reset() {
	*x = ServicePropertyValueInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spuPropertyService_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServicePropertyValueInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServicePropertyValueInfo) ProtoMessage() {}

func (x *ServicePropertyValueInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spuPropertyService_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServicePropertyValueInfo.ProtoReflect.Descriptor instead.
func (*ServicePropertyValueInfo) Descriptor() ([]byte, []int) {
	return file_spuPropertyService_proto_rawDescGZIP(), []int{3, 0}
}

func (x *ServicePropertyValueInfo) GetValueText() string {
	if x != nil {
		return x.ValueText
	}
	return ""
}

func (x *ServicePropertyValueInfo) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type SalePropertyValueInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ValueText string `protobuf:"bytes,1,opt,name=value_text,json=valueText,proto3" json:"value_text,omitempty"`
	ImageId   int64  `protobuf:"varint,2,opt,name=image_id,json=imageId,proto3" json:"image_id,omitempty"`
	ImageUrl  string `protobuf:"bytes,3,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
}

func (x *SalePropertyValueInfo) Reset() {
	*x = SalePropertyValueInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spuPropertyService_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SalePropertyValueInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SalePropertyValueInfo) ProtoMessage() {}

func (x *SalePropertyValueInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spuPropertyService_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SalePropertyValueInfo.ProtoReflect.Descriptor instead.
func (*SalePropertyValueInfo) Descriptor() ([]byte, []int) {
	return file_spuPropertyService_proto_rawDescGZIP(), []int{4, 0}
}

func (x *SalePropertyValueInfo) GetValueText() string {
	if x != nil {
		return x.ValueText
	}
	return ""
}

func (x *SalePropertyValueInfo) GetImageId() int64 {
	if x != nil {
		return x.ImageId
	}
	return 0
}

func (x *SalePropertyValueInfo) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

var File_spuPropertyService_proto protoreflect.FileDescriptor

var file_spuPropertyService_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x70, 0x75, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x22, 0x81, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63,
	0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x36, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x50, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x12, 0x36, 0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x69, 0x63, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x22, 0x50, 0x0a, 0x14, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x69, 0x63, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xab, 0x01, 0x0a, 0x14, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x19,
	0x0a, 0x08, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x55, 0x6e, 0x69, 0x74, 0x22, 0xcc, 0x01, 0x0a, 0x0f, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x19, 0x0a, 0x08,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3c, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x79, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x1a, 0x41, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x74, 0x65, 0x78, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x54, 0x65, 0x78,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0xe4, 0x01, 0x0a, 0x0c, 0x53, 0x61, 0x6c, 0x65,
	0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x70, 0x65, 0x63,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x70, 0x65, 0x63, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x70, 0x65, 0x63, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x70, 0x65, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x39,
	0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x61, 0x6c, 0x65, 0x50, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x1a, 0x63, 0x0a, 0x0a, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x54, 0x65, 0x78, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x22, 0x9a,
	0x01, 0x0a, 0x0b, 0x4f, 0x77, 0x6e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x17,
	0x0a, 0x07, 0x73, 0x70, 0x65, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x73, 0x70, 0x65, 0x63, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x70, 0x65, 0x63, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x70, 0x65, 0x63,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x70, 0x65, 0x63, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x70, 0x65, 0x63, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x42, 0x0d, 0x5a, 0x0b, 0x2f,
	0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_spuPropertyService_proto_rawDescOnce sync.Once
	file_spuPropertyService_proto_rawDescData = file_spuPropertyService_proto_rawDesc
)

func file_spuPropertyService_proto_rawDescGZIP() []byte {
	file_spuPropertyService_proto_rawDescOnce.Do(func() {
		file_spuPropertyService_proto_rawDescData = protoimpl.X.CompressGZIP(file_spuPropertyService_proto_rawDescData)
	})
	return file_spuPropertyService_proto_rawDescData
}

var file_spuPropertyService_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_spuPropertyService_proto_goTypes = []interface{}{
	(*GenericProperty)(nil),          // 0: services.GenericProperty
	(*GenericPropertyGroup)(nil),     // 1: services.GenericPropertyGroup
	(*GenericPropertyParam)(nil),     // 2: services.GenericPropertyParam
	(*ServiceProperty)(nil),          // 3: services.ServiceProperty
	(*SaleProperty)(nil),             // 4: services.SaleProperty
	(*OwnProperty)(nil),              // 5: services.OwnProperty
	(*ServicePropertyValueInfo)(nil), // 6: services.ServiceProperty.value_info
	(*SalePropertyValueInfo)(nil),    // 7: services.SaleProperty.value_info
}
var file_spuPropertyService_proto_depIdxs = []int32{
	2, // 0: services.GenericProperty.params:type_name -> services.GenericPropertyParam
	1, // 1: services.GenericProperty.groups:type_name -> services.GenericPropertyGroup
	6, // 2: services.ServiceProperty.values:type_name -> services.ServiceProperty.value_info
	7, // 3: services.SaleProperty.values:type_name -> services.SaleProperty.value_info
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_spuPropertyService_proto_init() }
func file_spuPropertyService_proto_init() {
	if File_spuPropertyService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spuPropertyService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenericProperty); i {
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
		file_spuPropertyService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenericPropertyGroup); i {
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
		file_spuPropertyService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenericPropertyParam); i {
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
		file_spuPropertyService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceProperty); i {
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
		file_spuPropertyService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaleProperty); i {
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
		file_spuPropertyService_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OwnProperty); i {
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
		file_spuPropertyService_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServicePropertyValueInfo); i {
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
		file_spuPropertyService_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SalePropertyValueInfo); i {
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
			RawDescriptor: file_spuPropertyService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_spuPropertyService_proto_goTypes,
		DependencyIndexes: file_spuPropertyService_proto_depIdxs,
		MessageInfos:      file_spuPropertyService_proto_msgTypes,
	}.Build()
	File_spuPropertyService_proto = out.File
	file_spuPropertyService_proto_rawDesc = nil
	file_spuPropertyService_proto_goTypes = nil
	file_spuPropertyService_proto_depIdxs = nil
}
