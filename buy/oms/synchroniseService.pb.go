// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: synchroniseService.proto

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

type SynchroniseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int32 `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	//----------
	OrderId  int64   `protobuf:"varint,3,opt,name=order_id,json=orderId,proto3" json:"order_id"`
	OrderNo  string  `protobuf:"bytes,4,opt,name=order_no,json=orderNo,proto3" json:"order_no"`
	OrderIds []int64 `protobuf:"varint,34,rep,packed,name=order_ids,json=orderIds,proto3" json:"order_ids"`
}

func (x *SynchroniseRequest) Reset() {
	*x = SynchroniseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_synchroniseService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SynchroniseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SynchroniseRequest) ProtoMessage() {}

func (x *SynchroniseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_synchroniseService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SynchroniseRequest.ProtoReflect.Descriptor instead.
func (*SynchroniseRequest) Descriptor() ([]byte, []int) {
	return file_synchroniseService_proto_rawDescGZIP(), []int{0}
}

func (x *SynchroniseRequest) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *SynchroniseRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *SynchroniseRequest) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *SynchroniseRequest) GetOrderNo() string {
	if x != nil {
		return x.OrderNo
	}
	return ""
}

func (x *SynchroniseRequest) GetOrderIds() []int64 {
	if x != nil {
		return x.OrderIds
	}
	return nil
}

type SynchroniseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info string `protobuf:"bytes,1,opt,name=info,proto3" json:"info"`
}

func (x *SynchroniseResponse) Reset() {
	*x = SynchroniseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_synchroniseService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SynchroniseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SynchroniseResponse) ProtoMessage() {}

func (x *SynchroniseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_synchroniseService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SynchroniseResponse.ProtoReflect.Descriptor instead.
func (*SynchroniseResponse) Descriptor() ([]byte, []int) {
	return file_synchroniseService_proto_rawDescGZIP(), []int{1}
}

func (x *SynchroniseResponse) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

var File_synchroniseService_proto protoreflect.FileDescriptor

var file_synchroniseService_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x73, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x22, 0x9a, 0x01, 0x0a, 0x12, 0x53, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f,
	0x6e, 0x69, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x5f, 0x6e, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x4e, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x73, 0x18, 0x22, 0x20, 0x03, 0x28, 0x03, 0x52, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x73, 0x22, 0x29, 0x0a, 0x13, 0x53, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x73, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0xab, 0x01, 0x0a,
	0x12, 0x53, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x73, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x08, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x61, 0x79, 0x12,
	0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x68,
	0x72, 0x6f, 0x6e, 0x69, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f,
	0x6e, 0x69, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a,
	0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x68, 0x69, 0x70, 0x12, 0x1c, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69,
	0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x73, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e,
	0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_synchroniseService_proto_rawDescOnce sync.Once
	file_synchroniseService_proto_rawDescData = file_synchroniseService_proto_rawDesc
)

func file_synchroniseService_proto_rawDescGZIP() []byte {
	file_synchroniseService_proto_rawDescOnce.Do(func() {
		file_synchroniseService_proto_rawDescData = protoimpl.X.CompressGZIP(file_synchroniseService_proto_rawDescData)
	})
	return file_synchroniseService_proto_rawDescData
}

var file_synchroniseService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_synchroniseService_proto_goTypes = []interface{}{
	(*SynchroniseRequest)(nil),  // 0: services.SynchroniseRequest
	(*SynchroniseResponse)(nil), // 1: services.SynchroniseResponse
}
var file_synchroniseService_proto_depIdxs = []int32{
	0, // 0: services.SynchroniseService.OrderPay:input_type -> services.SynchroniseRequest
	0, // 1: services.SynchroniseService.OrderShip:input_type -> services.SynchroniseRequest
	1, // 2: services.SynchroniseService.OrderPay:output_type -> services.SynchroniseResponse
	1, // 3: services.SynchroniseService.OrderShip:output_type -> services.SynchroniseResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_synchroniseService_proto_init() }
func file_synchroniseService_proto_init() {
	if File_synchroniseService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_synchroniseService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SynchroniseRequest); i {
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
		file_synchroniseService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SynchroniseResponse); i {
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
			RawDescriptor: file_synchroniseService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_synchroniseService_proto_goTypes,
		DependencyIndexes: file_synchroniseService_proto_depIdxs,
		MessageInfos:      file_synchroniseService_proto_msgTypes,
	}.Build()
	File_synchroniseService_proto = out.File
	file_synchroniseService_proto_rawDesc = nil
	file_synchroniseService_proto_goTypes = nil
	file_synchroniseService_proto_depIdxs = nil
}
