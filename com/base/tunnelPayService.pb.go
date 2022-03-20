// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: tunnelPayService.proto

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

type TunnelPay struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	Title      string `protobuf:"bytes,3,opt,name=title,proto3" json:"title"`
	Configs    string `protobuf:"bytes,4,opt,name=configs,proto3" json:"configs"`
	Version    int32  `protobuf:"varint,5,opt,name=version,proto3" json:"version"`
	PrivateKey string `protobuf:"bytes,6,opt,name=private_key,json=privateKey,proto3" json:"private_key"`
	PublicKey  string `protobuf:"bytes,7,opt,name=public_key,json=publicKey,proto3" json:"public_key"`
	Disabled   bool   `protobuf:"varint,8,opt,name=disabled,proto3" json:"disabled"`
}

func (x *TunnelPay) Reset() {
	*x = TunnelPay{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tunnelPayService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TunnelPay) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TunnelPay) ProtoMessage() {}

func (x *TunnelPay) ProtoReflect() protoreflect.Message {
	mi := &file_tunnelPayService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TunnelPay.ProtoReflect.Descriptor instead.
func (*TunnelPay) Descriptor() ([]byte, []int) {
	return file_tunnelPayService_proto_rawDescGZIP(), []int{0}
}

func (x *TunnelPay) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TunnelPay) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TunnelPay) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *TunnelPay) GetConfigs() string {
	if x != nil {
		return x.Configs
	}
	return ""
}

func (x *TunnelPay) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *TunnelPay) GetPrivateKey() string {
	if x != nil {
		return x.PrivateKey
	}
	return ""
}

func (x *TunnelPay) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

func (x *TunnelPay) GetDisabled() bool {
	if x != nil {
		return x.Disabled
	}
	return false
}

type TunnelPayRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int32  `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize int32  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Id       int32  `protobuf:"varint,3,opt,name=id,proto3" json:"id"`
	Name     string `protobuf:"bytes,4,opt,name=name,proto3" json:"name"`
	Version  int32  `protobuf:"varint,5,opt,name=version,proto3" json:"version"`
}

func (x *TunnelPayRequest) Reset() {
	*x = TunnelPayRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tunnelPayService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TunnelPayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TunnelPayRequest) ProtoMessage() {}

func (x *TunnelPayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tunnelPayService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TunnelPayRequest.ProtoReflect.Descriptor instead.
func (*TunnelPayRequest) Descriptor() ([]byte, []int) {
	return file_tunnelPayService_proto_rawDescGZIP(), []int{1}
}

func (x *TunnelPayRequest) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *TunnelPayRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *TunnelPayRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TunnelPayRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TunnelPayRequest) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

//
type TunnelPayResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *TunnelPay    `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*TunnelPay  `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Error  *common.Error `protobuf:"bytes,4,opt,name=error,proto3" json:"error"`
	Info   *common.Info  `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
}

func (x *TunnelPayResponse) Reset() {
	*x = TunnelPayResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tunnelPayService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TunnelPayResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TunnelPayResponse) ProtoMessage() {}

func (x *TunnelPayResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tunnelPayService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TunnelPayResponse.ProtoReflect.Descriptor instead.
func (*TunnelPayResponse) Descriptor() ([]byte, []int) {
	return file_tunnelPayService_proto_rawDescGZIP(), []int{2}
}

func (x *TunnelPayResponse) GetEntity() *TunnelPay {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *TunnelPayResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *TunnelPayResponse) GetItems() []*TunnelPay {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *TunnelPayResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *TunnelPayResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_tunnelPayService_proto protoreflect.FileDescriptor

var file_tunnelPayService_proto_rawDesc = []byte{
	0x0a, 0x16, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x50, 0x61, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd5, 0x01, 0x0a, 0x09, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c,
	0x50, 0x61, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b,
	0x65, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65,
	0x79, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x22, 0x83, 0x01,
	0x0a, 0x10, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x50, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x22, 0xd7, 0x01, 0x0a, 0x11, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x50, 0x61,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x50, 0x61, 0x79, 0x52, 0x06,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50,
	0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x50, 0x61, 0x79, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0xd3, 0x01,
	0x0a, 0x10, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x50, 0x61, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x39, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x50, 0x61, 0x79, 0x1a, 0x1b,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c,
	0x50, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a,
	0x09, 0x47, 0x65, 0x74, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x50, 0x61, 0x79, 0x1a,
	0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x75, 0x6e, 0x6e, 0x65,
	0x6c, 0x50, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43,
	0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x50, 0x61, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x50, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tunnelPayService_proto_rawDescOnce sync.Once
	file_tunnelPayService_proto_rawDescData = file_tunnelPayService_proto_rawDesc
)

func file_tunnelPayService_proto_rawDescGZIP() []byte {
	file_tunnelPayService_proto_rawDescOnce.Do(func() {
		file_tunnelPayService_proto_rawDescData = protoimpl.X.CompressGZIP(file_tunnelPayService_proto_rawDescData)
	})
	return file_tunnelPayService_proto_rawDescData
}

var file_tunnelPayService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_tunnelPayService_proto_goTypes = []interface{}{
	(*TunnelPay)(nil),         // 0: services.TunnelPay
	(*TunnelPayRequest)(nil),  // 1: services.TunnelPayRequest
	(*TunnelPayResponse)(nil), // 2: services.TunnelPayResponse
	(*common.Pager)(nil),      // 3: common.Pager
	(*common.Error)(nil),      // 4: common.Error
	(*common.Info)(nil),       // 5: common.Info
}
var file_tunnelPayService_proto_depIdxs = []int32{
	0, // 0: services.TunnelPayResponse.entity:type_name -> services.TunnelPay
	3, // 1: services.TunnelPayResponse.pager:type_name -> common.Pager
	0, // 2: services.TunnelPayResponse.items:type_name -> services.TunnelPay
	4, // 3: services.TunnelPayResponse.error:type_name -> common.Error
	5, // 4: services.TunnelPayResponse.info:type_name -> common.Info
	0, // 5: services.TunnelPayService.Get:input_type -> services.TunnelPay
	0, // 6: services.TunnelPayService.GetByName:input_type -> services.TunnelPay
	1, // 7: services.TunnelPayService.Search:input_type -> services.TunnelPayRequest
	2, // 8: services.TunnelPayService.Get:output_type -> services.TunnelPayResponse
	2, // 9: services.TunnelPayService.GetByName:output_type -> services.TunnelPayResponse
	2, // 10: services.TunnelPayService.Search:output_type -> services.TunnelPayResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_tunnelPayService_proto_init() }
func file_tunnelPayService_proto_init() {
	if File_tunnelPayService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tunnelPayService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TunnelPay); i {
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
		file_tunnelPayService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TunnelPayRequest); i {
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
		file_tunnelPayService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TunnelPayResponse); i {
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
			RawDescriptor: file_tunnelPayService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tunnelPayService_proto_goTypes,
		DependencyIndexes: file_tunnelPayService_proto_depIdxs,
		MessageInfos:      file_tunnelPayService_proto_msgTypes,
	}.Build()
	File_tunnelPayService_proto = out.File
	file_tunnelPayService_proto_rawDesc = nil
	file_tunnelPayService_proto_goTypes = nil
	file_tunnelPayService_proto_depIdxs = nil
}
