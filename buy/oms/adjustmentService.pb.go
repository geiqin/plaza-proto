// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: adjustmentService.proto

package services

import (
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

type Adjustment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderId         int64   `protobuf:"varint,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	OrderItemId     int64   `protobuf:"varint,3,opt,name=order_item_id,json=orderItemId,proto3" json:"order_item_id,omitempty"`
	OrderItemUnitId int64   `protobuf:"varint,4,opt,name=order_item_unit_id,json=orderItemUnitId,proto3" json:"order_item_unit_id,omitempty"`
	Type            string  `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	Label           string  `protobuf:"bytes,6,opt,name=label,proto3" json:"label,omitempty"`
	OriginCode      string  `protobuf:"bytes,7,opt,name=origin_code,json=originCode,proto3" json:"origin_code,omitempty"`
	Amount          float32 `protobuf:"fixed32,8,opt,name=amount,proto3" json:"amount,omitempty"`
	Included        bool    `protobuf:"varint,9,opt,name=included,proto3" json:"included,omitempty"`
	CreatedAt       string  `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt       string  `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Adjustment) Reset() {
	*x = Adjustment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adjustmentService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Adjustment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Adjustment) ProtoMessage() {}

func (x *Adjustment) ProtoReflect() protoreflect.Message {
	mi := &file_adjustmentService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Adjustment.ProtoReflect.Descriptor instead.
func (*Adjustment) Descriptor() ([]byte, []int) {
	return file_adjustmentService_proto_rawDescGZIP(), []int{0}
}

func (x *Adjustment) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Adjustment) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *Adjustment) GetOrderItemId() int64 {
	if x != nil {
		return x.OrderItemId
	}
	return 0
}

func (x *Adjustment) GetOrderItemUnitId() int64 {
	if x != nil {
		return x.OrderItemUnitId
	}
	return 0
}

func (x *Adjustment) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Adjustment) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *Adjustment) GetOriginCode() string {
	if x != nil {
		return x.OriginCode
	}
	return ""
}

func (x *Adjustment) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Adjustment) GetIncluded() bool {
	if x != nil {
		return x.Included
	}
	return false
}

func (x *Adjustment) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Adjustment) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

var File_adjustmentService_proto protoreflect.FileDescriptor

var file_adjustmentService_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x22, 0xc5, 0x02, 0x0a, 0x0a, 0x41, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x22, 0x0a,
	0x0d, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x49,
	0x64, 0x12, 0x2b, 0x0a, 0x12, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f,
	0x75, 0x6e, 0x69, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x55, 0x6e, 0x69, 0x74, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x0c, 0x5a, 0x0a, 0x2e,
	0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_adjustmentService_proto_rawDescOnce sync.Once
	file_adjustmentService_proto_rawDescData = file_adjustmentService_proto_rawDesc
)

func file_adjustmentService_proto_rawDescGZIP() []byte {
	file_adjustmentService_proto_rawDescOnce.Do(func() {
		file_adjustmentService_proto_rawDescData = protoimpl.X.CompressGZIP(file_adjustmentService_proto_rawDescData)
	})
	return file_adjustmentService_proto_rawDescData
}

var file_adjustmentService_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_adjustmentService_proto_goTypes = []interface{}{
	(*Adjustment)(nil), // 0: services.Adjustment
}
var file_adjustmentService_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_adjustmentService_proto_init() }
func file_adjustmentService_proto_init() {
	if File_adjustmentService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_adjustmentService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Adjustment); i {
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
			RawDescriptor: file_adjustmentService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_adjustmentService_proto_goTypes,
		DependencyIndexes: file_adjustmentService_proto_depIdxs,
		MessageInfos:      file_adjustmentService_proto_msgTypes,
	}.Build()
	File_adjustmentService_proto = out.File
	file_adjustmentService_proto_rawDesc = nil
	file_adjustmentService_proto_goTypes = nil
	file_adjustmentService_proto_depIdxs = nil
}