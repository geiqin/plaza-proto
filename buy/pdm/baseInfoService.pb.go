// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: baseInfoService.proto

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

type GoodsSpecBase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title  string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title"`
	Values []string `protobuf:"bytes,2,rep,name=values,proto3" json:"values"`
}

func (x *GoodsSpecBase) Reset() {
	*x = GoodsSpecBase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_baseInfoService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoodsSpecBase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodsSpecBase) ProtoMessage() {}

func (x *GoodsSpecBase) ProtoReflect() protoreflect.Message {
	mi := &file_baseInfoService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodsSpecBase.ProtoReflect.Descriptor instead.
func (*GoodsSpecBase) Descriptor() ([]byte, []int) {
	return file_baseInfoService_proto_rawDescGZIP(), []int{0}
}

func (x *GoodsSpecBase) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GoodsSpecBase) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

type GoodsSpecification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Choose []*GoodsSpecChoose `protobuf:"bytes,1,rep,name=choose,proto3" json:"choose"`
}

func (x *GoodsSpecification) Reset() {
	*x = GoodsSpecification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_baseInfoService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoodsSpecification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodsSpecification) ProtoMessage() {}

func (x *GoodsSpecification) ProtoReflect() protoreflect.Message {
	mi := &file_baseInfoService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodsSpecification.ProtoReflect.Descriptor instead.
func (*GoodsSpecification) Descriptor() ([]byte, []int) {
	return file_baseInfoService_proto_rawDescGZIP(), []int{1}
}

func (x *GoodsSpecification) GetChoose() []*GoodsSpecChoose {
	if x != nil {
		return x.Choose
	}
	return nil
}

type GoodsSpecChoose struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64                   `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Name    string                  `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	GoodsId int64                   `protobuf:"varint,3,opt,name=goods_id,json=goodsId,proto3" json:"goods_id"`
	Value   []*GoodsSpecChooseValue `protobuf:"bytes,4,rep,name=value,proto3" json:"value"`
}

func (x *GoodsSpecChoose) Reset() {
	*x = GoodsSpecChoose{}
	if protoimpl.UnsafeEnabled {
		mi := &file_baseInfoService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoodsSpecChoose) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodsSpecChoose) ProtoMessage() {}

func (x *GoodsSpecChoose) ProtoReflect() protoreflect.Message {
	mi := &file_baseInfoService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodsSpecChoose.ProtoReflect.Descriptor instead.
func (*GoodsSpecChoose) Descriptor() ([]byte, []int) {
	return file_baseInfoService_proto_rawDescGZIP(), []int{2}
}

func (x *GoodsSpecChoose) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GoodsSpecChoose) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GoodsSpecChoose) GetGoodsId() int64 {
	if x != nil {
		return x.GoodsId
	}
	return 0
}

func (x *GoodsSpecChoose) GetValue() []*GoodsSpecChooseValue {
	if x != nil {
		return x.Value
	}
	return nil
}

type GoodsSpecChooseValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
	ImageUrl string `protobuf:"bytes,2,opt,name=image_url,json=imageUrl,proto3" json:"image_url"`
}

func (x *GoodsSpecChooseValue) Reset() {
	*x = GoodsSpecChooseValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_baseInfoService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoodsSpecChooseValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodsSpecChooseValue) ProtoMessage() {}

func (x *GoodsSpecChooseValue) ProtoReflect() protoreflect.Message {
	mi := &file_baseInfoService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodsSpecChooseValue.ProtoReflect.Descriptor instead.
func (*GoodsSpecChooseValue) Descriptor() ([]byte, []int) {
	return file_baseInfoService_proto_rawDescGZIP(), []int{3}
}

func (x *GoodsSpecChooseValue) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GoodsSpecChooseValue) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

type GoodsParameter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title  string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title"`
	Values []string `protobuf:"bytes,2,rep,name=values,proto3" json:"values"`
}

func (x *GoodsParameter) Reset() {
	*x = GoodsParameter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_baseInfoService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoodsParameter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodsParameter) ProtoMessage() {}

func (x *GoodsParameter) ProtoReflect() protoreflect.Message {
	mi := &file_baseInfoService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodsParameter.ProtoReflect.Descriptor instead.
func (*GoodsParameter) Descriptor() ([]byte, []int) {
	return file_baseInfoService_proto_rawDescGZIP(), []int{4}
}

func (x *GoodsParameter) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GoodsParameter) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

type SpecTypeProperty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SpecId   int64                    `protobuf:"varint,1,opt,name=spec_id,json=specId,proto3" json:"spec_id"`
	SpecName string                   `protobuf:"bytes,2,opt,name=spec_name,json=specName,proto3" json:"spec_name"`
	Values   []*SpecTypePropertyValue `protobuf:"bytes,3,rep,name=values,proto3" json:"values"`
}

func (x *SpecTypeProperty) Reset() {
	*x = SpecTypeProperty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_baseInfoService_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpecTypeProperty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpecTypeProperty) ProtoMessage() {}

func (x *SpecTypeProperty) ProtoReflect() protoreflect.Message {
	mi := &file_baseInfoService_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpecTypeProperty.ProtoReflect.Descriptor instead.
func (*SpecTypeProperty) Descriptor() ([]byte, []int) {
	return file_baseInfoService_proto_rawDescGZIP(), []int{5}
}

func (x *SpecTypeProperty) GetSpecId() int64 {
	if x != nil {
		return x.SpecId
	}
	return 0
}

func (x *SpecTypeProperty) GetSpecName() string {
	if x != nil {
		return x.SpecName
	}
	return ""
}

func (x *SpecTypeProperty) GetValues() []*SpecTypePropertyValue {
	if x != nil {
		return x.Values
	}
	return nil
}

type SpecificationData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SpecId    int32  `protobuf:"varint,1,opt,name=spec_id,json=specId,proto3" json:"spec_id"`
	SpecName  string `protobuf:"bytes,2,opt,name=spec_name,json=specName,proto3" json:"spec_name"`
	SpecValue string `protobuf:"bytes,3,opt,name=spec_value,json=specValue,proto3" json:"spec_value"`
	ImageUrl  string `protobuf:"bytes,4,opt,name=image_url,json=imageUrl,proto3" json:"image_url"`
}

func (x *SpecificationData) Reset() {
	*x = SpecificationData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_baseInfoService_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpecificationData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpecificationData) ProtoMessage() {}

func (x *SpecificationData) ProtoReflect() protoreflect.Message {
	mi := &file_baseInfoService_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpecificationData.ProtoReflect.Descriptor instead.
func (*SpecificationData) Descriptor() ([]byte, []int) {
	return file_baseInfoService_proto_rawDescGZIP(), []int{6}
}

func (x *SpecificationData) GetSpecId() int32 {
	if x != nil {
		return x.SpecId
	}
	return 0
}

func (x *SpecificationData) GetSpecName() string {
	if x != nil {
		return x.SpecName
	}
	return ""
}

func (x *SpecificationData) GetSpecValue() string {
	if x != nil {
		return x.SpecValue
	}
	return ""
}

func (x *SpecificationData) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

type SpecTypePropertyValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
	ImageId  int64  `protobuf:"varint,2,opt,name=image_id,json=imageId,proto3" json:"image_id"`
	ImageUrl string `protobuf:"bytes,3,opt,name=image_url,json=imageUrl,proto3" json:"image_url"`
}

func (x *SpecTypePropertyValue) Reset() {
	*x = SpecTypePropertyValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_baseInfoService_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpecTypePropertyValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpecTypePropertyValue) ProtoMessage() {}

func (x *SpecTypePropertyValue) ProtoReflect() protoreflect.Message {
	mi := &file_baseInfoService_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpecTypePropertyValue.ProtoReflect.Descriptor instead.
func (*SpecTypePropertyValue) Descriptor() ([]byte, []int) {
	return file_baseInfoService_proto_rawDescGZIP(), []int{5, 0}
}

func (x *SpecTypePropertyValue) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SpecTypePropertyValue) GetImageId() int64 {
	if x != nil {
		return x.ImageId
	}
	return 0
}

func (x *SpecTypePropertyValue) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

var File_baseInfoService_proto protoreflect.FileDescriptor

var file_baseInfoService_proto_rawDesc = []byte{
	0x0a, 0x15, 0x62, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x22, 0x3d, 0x0a, 0x0d, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x53, 0x70, 0x65, 0x63, 0x42, 0x61,
	0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x22, 0x47, 0x0a, 0x12, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x31, 0x0a, 0x06, 0x63, 0x68, 0x6f, 0x6f, 0x73, 0x65,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x53, 0x70, 0x65, 0x63, 0x43, 0x68, 0x6f, 0x6f, 0x73,
	0x65, 0x52, 0x06, 0x63, 0x68, 0x6f, 0x6f, 0x73, 0x65, 0x22, 0x86, 0x01, 0x0a, 0x0f, 0x47, 0x6f,
	0x6f, 0x64, 0x73, 0x53, 0x70, 0x65, 0x63, 0x43, 0x68, 0x6f, 0x6f, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x53, 0x70, 0x65, 0x63,
	0x43, 0x68, 0x6f, 0x6f, 0x73, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0x47, 0x0a, 0x14, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x53, 0x70, 0x65, 0x63, 0x43,
	0x68, 0x6f, 0x6f, 0x73, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x22, 0x3e, 0x0a, 0x0e, 0x47,
	0x6f, 0x6f, 0x64, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0xd7, 0x01, 0x0a, 0x10,
	0x53, 0x70, 0x65, 0x63, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79,
	0x12, 0x17, 0x0a, 0x07, 0x73, 0x70, 0x65, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x73, 0x70, 0x65, 0x63, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x70, 0x65,
	0x63, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x70,
	0x65, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x53, 0x70, 0x65, 0x63, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x79, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x1a, 0x53, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x55, 0x72, 0x6c, 0x22, 0x85, 0x01, 0x0a, 0x11, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x17, 0x0a, 0x07, 0x73,
	0x70, 0x65, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x70,
	0x65, 0x63, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x70, 0x65, 0x63, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x70, 0x65, 0x63, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x70, 0x65, 0x63, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x70, 0x65, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x42, 0x0d, 0x5a,
	0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_baseInfoService_proto_rawDescOnce sync.Once
	file_baseInfoService_proto_rawDescData = file_baseInfoService_proto_rawDesc
)

func file_baseInfoService_proto_rawDescGZIP() []byte {
	file_baseInfoService_proto_rawDescOnce.Do(func() {
		file_baseInfoService_proto_rawDescData = protoimpl.X.CompressGZIP(file_baseInfoService_proto_rawDescData)
	})
	return file_baseInfoService_proto_rawDescData
}

var file_baseInfoService_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_baseInfoService_proto_goTypes = []interface{}{
	(*GoodsSpecBase)(nil),         // 0: services.GoodsSpecBase
	(*GoodsSpecification)(nil),    // 1: services.GoodsSpecification
	(*GoodsSpecChoose)(nil),       // 2: services.GoodsSpecChoose
	(*GoodsSpecChooseValue)(nil),  // 3: services.GoodsSpecChooseValue
	(*GoodsParameter)(nil),        // 4: services.GoodsParameter
	(*SpecTypeProperty)(nil),      // 5: services.SpecTypeProperty
	(*SpecificationData)(nil),     // 6: services.SpecificationData
	(*SpecTypePropertyValue)(nil), // 7: services.SpecTypeProperty.value
}
var file_baseInfoService_proto_depIdxs = []int32{
	2, // 0: services.GoodsSpecification.choose:type_name -> services.GoodsSpecChoose
	3, // 1: services.GoodsSpecChoose.value:type_name -> services.GoodsSpecChooseValue
	7, // 2: services.SpecTypeProperty.values:type_name -> services.SpecTypeProperty.value
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_baseInfoService_proto_init() }
func file_baseInfoService_proto_init() {
	if File_baseInfoService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_baseInfoService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoodsSpecBase); i {
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
		file_baseInfoService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoodsSpecification); i {
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
		file_baseInfoService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoodsSpecChoose); i {
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
		file_baseInfoService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoodsSpecChooseValue); i {
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
		file_baseInfoService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoodsParameter); i {
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
		file_baseInfoService_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpecTypeProperty); i {
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
		file_baseInfoService_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpecificationData); i {
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
		file_baseInfoService_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpecTypePropertyValue); i {
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
			RawDescriptor: file_baseInfoService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_baseInfoService_proto_goTypes,
		DependencyIndexes: file_baseInfoService_proto_depIdxs,
		MessageInfos:      file_baseInfoService_proto_msgTypes,
	}.Build()
	File_baseInfoService_proto = out.File
	file_baseInfoService_proto_rawDesc = nil
	file_baseInfoService_proto_goTypes = nil
	file_baseInfoService_proto_depIdxs = nil
}
