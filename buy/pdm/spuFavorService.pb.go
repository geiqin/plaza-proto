// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: spuFavorService.proto

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

type SpuFavor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	MemberId  int64  `protobuf:"varint,2,opt,name=member_id,json=memberId,proto3" json:"member_id"`
	SpuId     int64  `protobuf:"varint,3,opt,name=spu_id,json=spuId,proto3" json:"spu_id"`
	CreatedAt string `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt string `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	Spu       *Spu   `protobuf:"bytes,6,opt,name=spu,proto3" json:"spu"`
}

func (x *SpuFavor) Reset() {
	*x = SpuFavor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spuFavorService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpuFavor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpuFavor) ProtoMessage() {}

func (x *SpuFavor) ProtoReflect() protoreflect.Message {
	mi := &file_spuFavorService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpuFavor.ProtoReflect.Descriptor instead.
func (*SpuFavor) Descriptor() ([]byte, []int) {
	return file_spuFavorService_proto_rawDescGZIP(), []int{0}
}

func (x *SpuFavor) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SpuFavor) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

func (x *SpuFavor) GetSpuId() int64 {
	if x != nil {
		return x.SpuId
	}
	return 0
}

func (x *SpuFavor) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *SpuFavor) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *SpuFavor) GetSpu() *Spu {
	if x != nil {
		return x.Spu
	}
	return nil
}

type SpuFavorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int32  `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize int32  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Keywords string `protobuf:"bytes,3,opt,name=keywords,proto3" json:"keywords"`
	Id       int64  `protobuf:"varint,4,opt,name=id,proto3" json:"id"`
	MemberId int64  `protobuf:"varint,5,opt,name=member_id,json=memberId,proto3" json:"member_id"`
}

func (x *SpuFavorRequest) Reset() {
	*x = SpuFavorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spuFavorService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpuFavorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpuFavorRequest) ProtoMessage() {}

func (x *SpuFavorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spuFavorService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpuFavorRequest.ProtoReflect.Descriptor instead.
func (*SpuFavorRequest) Descriptor() ([]byte, []int) {
	return file_spuFavorService_proto_rawDescGZIP(), []int{1}
}

func (x *SpuFavorRequest) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *SpuFavorRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *SpuFavorRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *SpuFavorRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SpuFavorRequest) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

type SpuFavorData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *SpuFavor     `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*SpuFavor   `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Info   *common.Info  `protobuf:"bytes,4,opt,name=info,proto3" json:"info"`
}

func (x *SpuFavorData) Reset() {
	*x = SpuFavorData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spuFavorService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpuFavorData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpuFavorData) ProtoMessage() {}

func (x *SpuFavorData) ProtoReflect() protoreflect.Message {
	mi := &file_spuFavorService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpuFavorData.ProtoReflect.Descriptor instead.
func (*SpuFavorData) Descriptor() ([]byte, []int) {
	return file_spuFavorService_proto_rawDescGZIP(), []int{2}
}

func (x *SpuFavorData) GetEntity() *SpuFavor {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *SpuFavorData) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *SpuFavorData) GetItems() []*SpuFavor {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *SpuFavorData) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type SpuFavorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *SpuFavorData `protobuf:"bytes,1,opt,name=data,proto3" json:"data"`
	Error *common.Error `protobuf:"bytes,2,opt,name=error,proto3" json:"error"`
}

func (x *SpuFavorResponse) Reset() {
	*x = SpuFavorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spuFavorService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpuFavorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpuFavorResponse) ProtoMessage() {}

func (x *SpuFavorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spuFavorService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpuFavorResponse.ProtoReflect.Descriptor instead.
func (*SpuFavorResponse) Descriptor() ([]byte, []int) {
	return file_spuFavorService_proto_rawDescGZIP(), []int{3}
}

func (x *SpuFavorResponse) GetData() *SpuFavorData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *SpuFavorResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_spuFavorService_proto protoreflect.FileDescriptor

var file_spuFavorService_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x70, 0x75, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x73, 0x70, 0x75, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xad, 0x01, 0x0a, 0x08, 0x53, 0x70, 0x75, 0x46, 0x61,
	0x76, 0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x15, 0x0a, 0x06, 0x73, 0x70, 0x75, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x73, 0x70, 0x75, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1f, 0x0a, 0x03, 0x73, 0x70, 0x75, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x70,
	0x75, 0x52, 0x03, 0x73, 0x70, 0x75, 0x22, 0x8d, 0x01, 0x0a, 0x0f, 0x53, 0x70, 0x75, 0x46, 0x61,
	0x76, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61,
	0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x22, 0xab, 0x01, 0x0a, 0x0c, 0x53, 0x70, 0x75, 0x46, 0x61,
	0x76, 0x6f, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2a, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x53, 0x70, 0x75, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x52, 0x06, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65,
	0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x53, 0x70, 0x75, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x52, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04,
	0x69, 0x6e, 0x66, 0x6f, 0x22, 0x63, 0x0a, 0x10, 0x53, 0x70, 0x75, 0x46, 0x61, 0x76, 0x6f, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x53, 0x70, 0x75, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xc6, 0x01, 0x0a, 0x0f, 0x53, 0x70,
	0x75, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a,
	0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x53, 0x70, 0x75, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x1a, 0x1a, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x70, 0x75, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x70, 0x75,
	0x46, 0x61, 0x76, 0x6f, 0x72, 0x1a, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x53, 0x70, 0x75, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3f, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x19, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x70, 0x75, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x53, 0x70, 0x75, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spuFavorService_proto_rawDescOnce sync.Once
	file_spuFavorService_proto_rawDescData = file_spuFavorService_proto_rawDesc
)

func file_spuFavorService_proto_rawDescGZIP() []byte {
	file_spuFavorService_proto_rawDescOnce.Do(func() {
		file_spuFavorService_proto_rawDescData = protoimpl.X.CompressGZIP(file_spuFavorService_proto_rawDescData)
	})
	return file_spuFavorService_proto_rawDescData
}

var file_spuFavorService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_spuFavorService_proto_goTypes = []interface{}{
	(*SpuFavor)(nil),         // 0: services.SpuFavor
	(*SpuFavorRequest)(nil),  // 1: services.SpuFavorRequest
	(*SpuFavorData)(nil),     // 2: services.SpuFavorData
	(*SpuFavorResponse)(nil), // 3: services.SpuFavorResponse
	(*Spu)(nil),              // 4: services.Spu
	(*common.Pager)(nil),     // 5: common.Pager
	(*common.Info)(nil),      // 6: common.Info
	(*common.Error)(nil),     // 7: common.Error
}
var file_spuFavorService_proto_depIdxs = []int32{
	4,  // 0: services.SpuFavor.spu:type_name -> services.Spu
	0,  // 1: services.SpuFavorData.entity:type_name -> services.SpuFavor
	5,  // 2: services.SpuFavorData.pager:type_name -> common.Pager
	0,  // 3: services.SpuFavorData.items:type_name -> services.SpuFavor
	6,  // 4: services.SpuFavorData.info:type_name -> common.Info
	2,  // 5: services.SpuFavorResponse.data:type_name -> services.SpuFavorData
	7,  // 6: services.SpuFavorResponse.error:type_name -> common.Error
	0,  // 7: services.SpuFavorService.Create:input_type -> services.SpuFavor
	0,  // 8: services.SpuFavorService.Delete:input_type -> services.SpuFavor
	1,  // 9: services.SpuFavorService.Search:input_type -> services.SpuFavorRequest
	3,  // 10: services.SpuFavorService.Create:output_type -> services.SpuFavorResponse
	3,  // 11: services.SpuFavorService.Delete:output_type -> services.SpuFavorResponse
	3,  // 12: services.SpuFavorService.Search:output_type -> services.SpuFavorResponse
	10, // [10:13] is the sub-list for method output_type
	7,  // [7:10] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_spuFavorService_proto_init() }
func file_spuFavorService_proto_init() {
	if File_spuFavorService_proto != nil {
		return
	}
	file_spuService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_spuFavorService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpuFavor); i {
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
		file_spuFavorService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpuFavorRequest); i {
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
		file_spuFavorService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpuFavorData); i {
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
		file_spuFavorService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpuFavorResponse); i {
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
			RawDescriptor: file_spuFavorService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spuFavorService_proto_goTypes,
		DependencyIndexes: file_spuFavorService_proto_depIdxs,
		MessageInfos:      file_spuFavorService_proto_msgTypes,
	}.Build()
	File_spuFavorService_proto = out.File
	file_spuFavorService_proto_rawDesc = nil
	file_spuFavorService_proto_goTypes = nil
	file_spuFavorService_proto_depIdxs = nil
}
