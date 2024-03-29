// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: tagService.proto

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

// 标签信息
type Tag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	TotalNum int32  `protobuf:"varint,5,opt,name=total_num,json=totalNum,proto3" json:"total_num"`
}

func (x *Tag) Reset() {
	*x = Tag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tagService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tag) ProtoMessage() {}

func (x *Tag) ProtoReflect() protoreflect.Message {
	mi := &file_tagService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tag.ProtoReflect.Descriptor instead.
func (*Tag) Descriptor() ([]byte, []int) {
	return file_tagService_proto_rawDescGZIP(), []int{0}
}

func (x *Tag) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Tag) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Tag) GetTotalNum() int32 {
	if x != nil {
		return x.TotalNum
	}
	return 0
}

type TatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int64 `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Top      int64 `protobuf:"varint,3,opt,name=top,proto3" json:"top"`
	//file params
	Id   int64   `protobuf:"varint,4,opt,name=id,proto3" json:"id"`
	Name string  `protobuf:"bytes,5,opt,name=name,proto3" json:"name"`
	Slug string  `protobuf:"bytes,6,opt,name=slug,proto3" json:"slug"`
	Ids  []int64 `protobuf:"varint,7,rep,packed,name=ids,proto3" json:"ids"`
}

func (x *TatRequest) Reset() {
	*x = TatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tagService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TatRequest) ProtoMessage() {}

func (x *TatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tagService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TatRequest.ProtoReflect.Descriptor instead.
func (*TatRequest) Descriptor() ([]byte, []int) {
	return file_tagService_proto_rawDescGZIP(), []int{1}
}

func (x *TatRequest) GetPaged() int64 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *TatRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *TatRequest) GetTop() int64 {
	if x != nil {
		return x.Top
	}
	return 0
}

func (x *TatRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TatRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TatRequest) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *TatRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type TagResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Tag          `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*Tag        `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Info   string        `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
}

func (x *TagResponse) Reset() {
	*x = TagResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tagService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagResponse) ProtoMessage() {}

func (x *TagResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tagService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagResponse.ProtoReflect.Descriptor instead.
func (*TagResponse) Descriptor() ([]byte, []int) {
	return file_tagService_proto_rawDescGZIP(), []int{2}
}

func (x *TagResponse) GetEntity() *Tag {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *TagResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *TagResponse) GetItems() []*Tag {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *TagResponse) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

var File_tagService_proto protoreflect.FileDescriptor

var file_tagService_proto_rawDesc = []byte{
	0x0a, 0x10, 0x74, 0x61, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x46, 0x0a, 0x03, 0x54, 0x61, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x22, 0x9b, 0x01, 0x0a, 0x0a, 0x54, 0x61, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x6f, 0x70,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x74, 0x6f, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73,
	0x6c, 0x75, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x03,
	0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x92, 0x01, 0x0a, 0x0b, 0x54, 0x61, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x54, 0x61, 0x67, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05,
	0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65,
	0x72, 0x12, 0x23, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x61, 0x67, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0xd8, 0x01, 0x0a, 0x0a, 0x54,
	0x61, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x0d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54,
	0x61, 0x67, 0x1a, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x61,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x06, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x54, 0x61, 0x67, 0x1a, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x54, 0x61, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2d, 0x0a,
	0x03, 0x47, 0x65, 0x74, 0x12, 0x0d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x54, 0x61, 0x67, 0x1a, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54,
	0x61, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x06,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x54, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tagService_proto_rawDescOnce sync.Once
	file_tagService_proto_rawDescData = file_tagService_proto_rawDesc
)

func file_tagService_proto_rawDescGZIP() []byte {
	file_tagService_proto_rawDescOnce.Do(func() {
		file_tagService_proto_rawDescData = protoimpl.X.CompressGZIP(file_tagService_proto_rawDescData)
	})
	return file_tagService_proto_rawDescData
}

var file_tagService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_tagService_proto_goTypes = []interface{}{
	(*Tag)(nil),          // 0: services.Tag
	(*TatRequest)(nil),   // 1: services.TatRequest
	(*TagResponse)(nil),  // 2: services.TagResponse
	(*common.Pager)(nil), // 3: common.Pager
}
var file_tagService_proto_depIdxs = []int32{
	0, // 0: services.TagResponse.entity:type_name -> services.Tag
	3, // 1: services.TagResponse.pager:type_name -> common.Pager
	0, // 2: services.TagResponse.items:type_name -> services.Tag
	0, // 3: services.TagService.Create:input_type -> services.Tag
	0, // 4: services.TagService.Delete:input_type -> services.Tag
	0, // 5: services.TagService.Get:input_type -> services.Tag
	1, // 6: services.TagService.Search:input_type -> services.TatRequest
	2, // 7: services.TagService.Create:output_type -> services.TagResponse
	2, // 8: services.TagService.Delete:output_type -> services.TagResponse
	2, // 9: services.TagService.Get:output_type -> services.TagResponse
	2, // 10: services.TagService.Search:output_type -> services.TagResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_tagService_proto_init() }
func file_tagService_proto_init() {
	if File_tagService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tagService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tag); i {
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
		file_tagService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TatRequest); i {
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
		file_tagService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagResponse); i {
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
			RawDescriptor: file_tagService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tagService_proto_goTypes,
		DependencyIndexes: file_tagService_proto_depIdxs,
		MessageInfos:      file_tagService_proto_msgTypes,
	}.Build()
	File_tagService_proto = out.File
	file_tagService_proto_rawDesc = nil
	file_tagService_proto_goTypes = nil
	file_tagService_proto_depIdxs = nil
}
