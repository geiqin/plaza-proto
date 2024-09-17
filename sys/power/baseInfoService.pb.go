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

type TreeItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key      int64       `protobuf:"varint,1,opt,name=key,proto3" json:"key"`           //ID
	ParentId int64       `protobuf:"varint,2,opt,name=parentId,proto3" json:"parentId"` //父ID
	Title    string      `protobuf:"bytes,3,opt,name=title,proto3" json:"title"`        //标题
	Sort     int32       `protobuf:"varint,4,opt,name=sort,proto3" json:"sort"`         //排序
	Children []*TreeItem `protobuf:"bytes,5,rep,name=children,proto3" json:"children"`
}

func (x *TreeItem) Reset() {
	*x = TreeItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_baseInfoService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TreeItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TreeItem) ProtoMessage() {}

func (x *TreeItem) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use TreeItem.ProtoReflect.Descriptor instead.
func (*TreeItem) Descriptor() ([]byte, []int) {
	return file_baseInfoService_proto_rawDescGZIP(), []int{0}
}

func (x *TreeItem) GetKey() int64 {
	if x != nil {
		return x.Key
	}
	return 0
}

func (x *TreeItem) GetParentId() int64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *TreeItem) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *TreeItem) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *TreeItem) GetChildren() []*TreeItem {
	if x != nil {
		return x.Children
	}
	return nil
}

var File_baseInfoService_proto protoreflect.FileDescriptor

var file_baseInfoService_proto_rawDesc = []byte{
	0x0a, 0x15, 0x62, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x22, 0x92, 0x01, 0x0a, 0x08, 0x54, 0x72, 0x65, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x2e, 0x0a, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72,
	0x65, 0x6e, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x54, 0x72, 0x65, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x08, 0x63, 0x68,
	0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_baseInfoService_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_baseInfoService_proto_goTypes = []interface{}{
	(*TreeItem)(nil), // 0: services.TreeItem
}
var file_baseInfoService_proto_depIdxs = []int32{
	0, // 0: services.TreeItem.children:type_name -> services.TreeItem
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_baseInfoService_proto_init() }
func file_baseInfoService_proto_init() {
	if File_baseInfoService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_baseInfoService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TreeItem); i {
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
			NumMessages:   1,
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
