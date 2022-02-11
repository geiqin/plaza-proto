// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: shoppingService.proto

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

// 购物商城参数
type Shopping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OpenStockFunc bool `protobuf:"varint,1,opt,name=open_stock_func,json=openStockFunc,proto3" json:"open_stock_func"`
}

func (x *Shopping) Reset() {
	*x = Shopping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shoppingService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Shopping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Shopping) ProtoMessage() {}

func (x *Shopping) ProtoReflect() protoreflect.Message {
	mi := &file_shoppingService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Shopping.ProtoReflect.Descriptor instead.
func (*Shopping) Descriptor() ([]byte, []int) {
	return file_shoppingService_proto_rawDescGZIP(), []int{0}
}

func (x *Shopping) GetOpenStockFunc() bool {
	if x != nil {
		return x.OpenStockFunc
	}
	return false
}

//
type ShoppingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Shopping     `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Error  *common.Error `protobuf:"bytes,4,opt,name=error,proto3" json:"error"`
	Info   *common.Info  `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
}

func (x *ShoppingResponse) Reset() {
	*x = ShoppingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shoppingService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShoppingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShoppingResponse) ProtoMessage() {}

func (x *ShoppingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shoppingService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShoppingResponse.ProtoReflect.Descriptor instead.
func (*ShoppingResponse) Descriptor() ([]byte, []int) {
	return file_shoppingService_proto_rawDescGZIP(), []int{1}
}

func (x *ShoppingResponse) GetEntity() *Shopping {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *ShoppingResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *ShoppingResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_shoppingService_proto protoreflect.FileDescriptor

var file_shoppingService_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x32, 0x0a, 0x08, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x12, 0x26, 0x0a, 0x0f, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x66,
	0x75, 0x6e, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x6f, 0x70, 0x65, 0x6e, 0x53,
	0x74, 0x6f, 0x63, 0x6b, 0x46, 0x75, 0x6e, 0x63, 0x22, 0x85, 0x01, 0x0a, 0x10, 0x53, 0x68, 0x6f,
	0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a,
	0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20,
	0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f,
	0x32, 0x7e, 0x0a, 0x0f, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x03, 0x53, 0x65, 0x74, 0x12, 0x12, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x1a, 0x1a,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x03,
	0x47, 0x65, 0x74, 0x12, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x68,
	0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shoppingService_proto_rawDescOnce sync.Once
	file_shoppingService_proto_rawDescData = file_shoppingService_proto_rawDesc
)

func file_shoppingService_proto_rawDescGZIP() []byte {
	file_shoppingService_proto_rawDescOnce.Do(func() {
		file_shoppingService_proto_rawDescData = protoimpl.X.CompressGZIP(file_shoppingService_proto_rawDescData)
	})
	return file_shoppingService_proto_rawDescData
}

var file_shoppingService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_shoppingService_proto_goTypes = []interface{}{
	(*Shopping)(nil),         // 0: services.Shopping
	(*ShoppingResponse)(nil), // 1: services.ShoppingResponse
	(*common.Error)(nil),     // 2: common.Error
	(*common.Info)(nil),      // 3: common.Info
	(*common.Empty)(nil),     // 4: common.Empty
}
var file_shoppingService_proto_depIdxs = []int32{
	0, // 0: services.ShoppingResponse.entity:type_name -> services.Shopping
	2, // 1: services.ShoppingResponse.error:type_name -> common.Error
	3, // 2: services.ShoppingResponse.info:type_name -> common.Info
	0, // 3: services.ShoppingService.Set:input_type -> services.Shopping
	4, // 4: services.ShoppingService.Get:input_type -> common.Empty
	1, // 5: services.ShoppingService.Set:output_type -> services.ShoppingResponse
	1, // 6: services.ShoppingService.Get:output_type -> services.ShoppingResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_shoppingService_proto_init() }
func file_shoppingService_proto_init() {
	if File_shoppingService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shoppingService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Shopping); i {
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
		file_shoppingService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShoppingResponse); i {
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
			RawDescriptor: file_shoppingService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shoppingService_proto_goTypes,
		DependencyIndexes: file_shoppingService_proto_depIdxs,
		MessageInfos:      file_shoppingService_proto_msgTypes,
	}.Build()
	File_shoppingService_proto = out.File
	file_shoppingService_proto_rawDesc = nil
	file_shoppingService_proto_goTypes = nil
	file_shoppingService_proto_depIdxs = nil
}
