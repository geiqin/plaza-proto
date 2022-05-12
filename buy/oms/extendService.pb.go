// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: extendService.proto

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

//订单详情步骤
type OrderFlow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Current *OrderFlowStep   `protobuf:"bytes,1,opt,name=current,proto3" json:"current,omitempty"`
	Steps   []*OrderFlowStep `protobuf:"bytes,2,rep,name=steps,proto3" json:"steps,omitempty"`
}

func (x *OrderFlow) Reset() {
	*x = OrderFlow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_extendService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderFlow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderFlow) ProtoMessage() {}

func (x *OrderFlow) ProtoReflect() protoreflect.Message {
	mi := &file_extendService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderFlow.ProtoReflect.Descriptor instead.
func (*OrderFlow) Descriptor() ([]byte, []int) {
	return file_extendService_proto_rawDescGZIP(), []int{0}
}

func (x *OrderFlow) GetCurrent() *OrderFlowStep {
	if x != nil {
		return x.Current
	}
	return nil
}

func (x *OrderFlow) GetSteps() []*OrderFlowStep {
	if x != nil {
		return x.Steps
	}
	return nil
}

type OrderFlowStep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code        string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ShortName   string `protobuf:"bytes,3,opt,name=short_name,json=shortName,proto3" json:"short_name,omitempty"`
	IconUrl     string `protobuf:"bytes,4,opt,name=icon_url,json=iconUrl,proto3" json:"icon_url,omitempty"`
	Status      string `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	HappenAt    string `protobuf:"bytes,6,opt,name=happen_at,json=happenAt,proto3" json:"happen_at,omitempty"`
	Description string `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *OrderFlowStep) Reset() {
	*x = OrderFlowStep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_extendService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderFlowStep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderFlowStep) ProtoMessage() {}

func (x *OrderFlowStep) ProtoReflect() protoreflect.Message {
	mi := &file_extendService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderFlowStep.ProtoReflect.Descriptor instead.
func (*OrderFlowStep) Descriptor() ([]byte, []int) {
	return file_extendService_proto_rawDescGZIP(), []int{0, 0}
}

func (x *OrderFlowStep) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *OrderFlowStep) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OrderFlowStep) GetShortName() string {
	if x != nil {
		return x.ShortName
	}
	return ""
}

func (x *OrderFlowStep) GetIconUrl() string {
	if x != nil {
		return x.IconUrl
	}
	return ""
}

func (x *OrderFlowStep) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *OrderFlowStep) GetHappenAt() string {
	if x != nil {
		return x.HappenAt
	}
	return ""
}

func (x *OrderFlowStep) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_extendService_proto protoreflect.FileDescriptor

var file_extendService_proto_rawDesc = []byte{
	0x0a, 0x13, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22,
	0xb1, 0x02, 0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x46, 0x6c, 0x6f, 0x77, 0x12, 0x32, 0x0a,
	0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x46,
	0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x74, 0x65, 0x70, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x12, 0x2e, 0x0a, 0x05, 0x73, 0x74, 0x65, 0x70, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x46, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x74, 0x65, 0x70, 0x52, 0x05, 0x73, 0x74, 0x65, 0x70,
	0x73, 0x1a, 0xbf, 0x01, 0x0a, 0x04, 0x73, 0x74, 0x65, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x63, 0x6f, 0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x63, 0x6f, 0x6e, 0x55, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x68, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x5f, 0x61,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x41,
	0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_extendService_proto_rawDescOnce sync.Once
	file_extendService_proto_rawDescData = file_extendService_proto_rawDesc
)

func file_extendService_proto_rawDescGZIP() []byte {
	file_extendService_proto_rawDescOnce.Do(func() {
		file_extendService_proto_rawDescData = protoimpl.X.CompressGZIP(file_extendService_proto_rawDescData)
	})
	return file_extendService_proto_rawDescData
}

var file_extendService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_extendService_proto_goTypes = []interface{}{
	(*OrderFlow)(nil),     // 0: services.OrderFlow
	(*OrderFlowStep)(nil), // 1: services.OrderFlow.step
}
var file_extendService_proto_depIdxs = []int32{
	1, // 0: services.OrderFlow.current:type_name -> services.OrderFlow.step
	1, // 1: services.OrderFlow.steps:type_name -> services.OrderFlow.step
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_extendService_proto_init() }
func file_extendService_proto_init() {
	if File_extendService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_extendService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderFlow); i {
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
		file_extendService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderFlowStep); i {
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
			RawDescriptor: file_extendService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_extendService_proto_goTypes,
		DependencyIndexes: file_extendService_proto_depIdxs,
		MessageInfos:      file_extendService_proto_msgTypes,
	}.Build()
	File_extendService_proto = out.File
	file_extendService_proto_rawDesc = nil
	file_extendService_proto_goTypes = nil
	file_extendService_proto_depIdxs = nil
}
