// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: fanTrackService.proto

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

type FanTrack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Source    string `protobuf:"bytes,2,opt,name=source,proto3" json:"source"`
	CreatedAt string `protobuf:"bytes,17,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt string `protobuf:"bytes,18,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	// @inject_tag: gorm:"-"
	Ids []int64 `protobuf:"varint,20,rep,packed,name=ids,proto3" json:"ids" gorm:"-"`
}

func (x *FanTrack) Reset() {
	*x = FanTrack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fanTrackService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FanTrack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FanTrack) ProtoMessage() {}

func (x *FanTrack) ProtoReflect() protoreflect.Message {
	mi := &file_fanTrackService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FanTrack.ProtoReflect.Descriptor instead.
func (*FanTrack) Descriptor() ([]byte, []int) {
	return file_fanTrackService_proto_rawDescGZIP(), []int{0}
}

func (x *FanTrack) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FanTrack) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *FanTrack) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *FanTrack) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *FanTrack) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type FanTrackResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *FanTrack     `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*FanTrack   `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Error  *common.Error `protobuf:"bytes,4,opt,name=error,proto3" json:"error"`
	Info   *common.Info  `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
}

func (x *FanTrackResponse) Reset() {
	*x = FanTrackResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fanTrackService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FanTrackResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FanTrackResponse) ProtoMessage() {}

func (x *FanTrackResponse) ProtoReflect() protoreflect.Message {
	mi := &file_fanTrackService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FanTrackResponse.ProtoReflect.Descriptor instead.
func (*FanTrackResponse) Descriptor() ([]byte, []int) {
	return file_fanTrackService_proto_rawDescGZIP(), []int{1}
}

func (x *FanTrackResponse) GetEntity() *FanTrack {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *FanTrackResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *FanTrackResponse) GetItems() []*FanTrack {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *FanTrackResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *FanTrackResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_fanTrackService_proto protoreflect.FileDescriptor

var file_fanTrackService_proto_rawDesc = []byte{
	0x0a, 0x15, 0x66, 0x61, 0x6e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82, 0x01, 0x0a, 0x08, 0x46, 0x61, 0x6e, 0x54, 0x72, 0x61, 0x63,
	0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x14,
	0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0xd4, 0x01, 0x0a, 0x10, 0x46, 0x61,
	0x6e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a,
	0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x61, 0x6e, 0x54, 0x72, 0x61,
	0x63, 0x6b, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61,
	0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12,
	0x28, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x61, 0x6e, 0x54, 0x72, 0x61,
	0x63, 0x6b, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20,
	0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f,
	0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fanTrackService_proto_rawDescOnce sync.Once
	file_fanTrackService_proto_rawDescData = file_fanTrackService_proto_rawDesc
)

func file_fanTrackService_proto_rawDescGZIP() []byte {
	file_fanTrackService_proto_rawDescOnce.Do(func() {
		file_fanTrackService_proto_rawDescData = protoimpl.X.CompressGZIP(file_fanTrackService_proto_rawDescData)
	})
	return file_fanTrackService_proto_rawDescData
}

var file_fanTrackService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_fanTrackService_proto_goTypes = []interface{}{
	(*FanTrack)(nil),         // 0: services.FanTrack
	(*FanTrackResponse)(nil), // 1: services.FanTrackResponse
	(*common.Pager)(nil),     // 2: common.Pager
	(*common.Error)(nil),     // 3: common.Error
	(*common.Info)(nil),      // 4: common.Info
}
var file_fanTrackService_proto_depIdxs = []int32{
	0, // 0: services.FanTrackResponse.entity:type_name -> services.FanTrack
	2, // 1: services.FanTrackResponse.pager:type_name -> common.Pager
	0, // 2: services.FanTrackResponse.items:type_name -> services.FanTrack
	3, // 3: services.FanTrackResponse.error:type_name -> common.Error
	4, // 4: services.FanTrackResponse.info:type_name -> common.Info
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_fanTrackService_proto_init() }
func file_fanTrackService_proto_init() {
	if File_fanTrackService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fanTrackService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FanTrack); i {
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
		file_fanTrackService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FanTrackResponse); i {
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
			RawDescriptor: file_fanTrackService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_fanTrackService_proto_goTypes,
		DependencyIndexes: file_fanTrackService_proto_depIdxs,
		MessageInfos:      file_fanTrackService_proto_msgTypes,
	}.Build()
	File_fanTrackService_proto = out.File
	file_fanTrackService_proto_rawDesc = nil
	file_fanTrackService_proto_goTypes = nil
	file_fanTrackService_proto_depIdxs = nil
}
