// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: dictionayValueService.proto

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

// 字典值信息
type DictionaryValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	DictionaryId int64  `protobuf:"varint,2,opt,name=dictionary_id,json=dictionaryId,proto3" json:"dictionary_id"`
	Value        string `protobuf:"bytes,3,opt,name=value,proto3" json:"value"`
	Text         string `protobuf:"bytes,4,opt,name=text,proto3" json:"text"`
	Flag         string `protobuf:"bytes,5,opt,name=flag,proto3" json:"flag"`
	Color        string `protobuf:"bytes,7,opt,name=color,proto3" json:"color"`
	Source       string `protobuf:"bytes,8,opt,name=source,proto3" json:"source"`
	Sort         int32  `protobuf:"varint,9,opt,name=sort,proto3" json:"sort"`
	IsEnabled    string `protobuf:"bytes,10,opt,name=is_enabled,json=isEnabled,proto3" json:"is_enabled"`
	CreatedAt    string `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt    string `protobuf:"bytes,12,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
}

func (x *DictionaryValue) Reset() {
	*x = DictionaryValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dictionayValueService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DictionaryValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DictionaryValue) ProtoMessage() {}

func (x *DictionaryValue) ProtoReflect() protoreflect.Message {
	mi := &file_dictionayValueService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DictionaryValue.ProtoReflect.Descriptor instead.
func (*DictionaryValue) Descriptor() ([]byte, []int) {
	return file_dictionayValueService_proto_rawDescGZIP(), []int{0}
}

func (x *DictionaryValue) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DictionaryValue) GetDictionaryId() int64 {
	if x != nil {
		return x.DictionaryId
	}
	return 0
}

func (x *DictionaryValue) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *DictionaryValue) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *DictionaryValue) GetFlag() string {
	if x != nil {
		return x.Flag
	}
	return ""
}

func (x *DictionaryValue) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *DictionaryValue) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *DictionaryValue) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *DictionaryValue) GetIsEnabled() string {
	if x != nil {
		return x.IsEnabled
	}
	return ""
}

func (x *DictionaryValue) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *DictionaryValue) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type DictionaryValueResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *DictionaryValue   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager      `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*DictionaryValue `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Info   string             `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
}

func (x *DictionaryValueResponse) Reset() {
	*x = DictionaryValueResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dictionayValueService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DictionaryValueResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DictionaryValueResponse) ProtoMessage() {}

func (x *DictionaryValueResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dictionayValueService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DictionaryValueResponse.ProtoReflect.Descriptor instead.
func (*DictionaryValueResponse) Descriptor() ([]byte, []int) {
	return file_dictionayValueService_proto_rawDescGZIP(), []int{1}
}

func (x *DictionaryValueResponse) GetEntity() *DictionaryValue {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *DictionaryValueResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *DictionaryValueResponse) GetItems() []*DictionaryValue {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *DictionaryValueResponse) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

var File_dictionayValueService_proto protoreflect.FileDescriptor

var file_dictionayValueService_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa3, 0x02, 0x0a, 0x0f, 0x44,
	0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x23,
	0x0a, 0x0d, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72,
	0x79, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x6c, 0x61,
	0x67, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73,
	0x6f, 0x72, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x22, 0xb6, 0x01, 0x0a, 0x17, 0x44, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x06,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x72, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12,
	0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70,
	0x61, 0x67, 0x65, 0x72, 0x12, 0x2f, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44,
	0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0xdc, 0x03, 0x0a, 0x16, 0x44, 0x69,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x19,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x69, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x21, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x48,
	0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x44, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x1a, 0x21, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44,
	0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x69,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x21, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x45, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x1a, 0x21, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x44, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x0a, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x6f, 0x72, 0x74, 0x12, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x44, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x1a, 0x21, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x69,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x44, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x1a, 0x21, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44,
	0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dictionayValueService_proto_rawDescOnce sync.Once
	file_dictionayValueService_proto_rawDescData = file_dictionayValueService_proto_rawDesc
)

func file_dictionayValueService_proto_rawDescGZIP() []byte {
	file_dictionayValueService_proto_rawDescOnce.Do(func() {
		file_dictionayValueService_proto_rawDescData = protoimpl.X.CompressGZIP(file_dictionayValueService_proto_rawDescData)
	})
	return file_dictionayValueService_proto_rawDescData
}

var file_dictionayValueService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_dictionayValueService_proto_goTypes = []interface{}{
	(*DictionaryValue)(nil),         // 0: services.DictionaryValue
	(*DictionaryValueResponse)(nil), // 1: services.DictionaryValueResponse
	(*common.Pager)(nil),            // 2: common.Pager
}
var file_dictionayValueService_proto_depIdxs = []int32{
	0, // 0: services.DictionaryValueResponse.entity:type_name -> services.DictionaryValue
	2, // 1: services.DictionaryValueResponse.pager:type_name -> common.Pager
	0, // 2: services.DictionaryValueResponse.items:type_name -> services.DictionaryValue
	0, // 3: services.DictionaryValueService.Create:input_type -> services.DictionaryValue
	0, // 4: services.DictionaryValueService.Update:input_type -> services.DictionaryValue
	0, // 5: services.DictionaryValueService.Delete:input_type -> services.DictionaryValue
	0, // 6: services.DictionaryValueService.Get:input_type -> services.DictionaryValue
	0, // 7: services.DictionaryValueService.UpdateSort:input_type -> services.DictionaryValue
	0, // 8: services.DictionaryValueService.UpdateEnabled:input_type -> services.DictionaryValue
	1, // 9: services.DictionaryValueService.Create:output_type -> services.DictionaryValueResponse
	1, // 10: services.DictionaryValueService.Update:output_type -> services.DictionaryValueResponse
	1, // 11: services.DictionaryValueService.Delete:output_type -> services.DictionaryValueResponse
	1, // 12: services.DictionaryValueService.Get:output_type -> services.DictionaryValueResponse
	1, // 13: services.DictionaryValueService.UpdateSort:output_type -> services.DictionaryValueResponse
	1, // 14: services.DictionaryValueService.UpdateEnabled:output_type -> services.DictionaryValueResponse
	9, // [9:15] is the sub-list for method output_type
	3, // [3:9] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_dictionayValueService_proto_init() }
func file_dictionayValueService_proto_init() {
	if File_dictionayValueService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dictionayValueService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DictionaryValue); i {
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
		file_dictionayValueService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DictionaryValueResponse); i {
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
			RawDescriptor: file_dictionayValueService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dictionayValueService_proto_goTypes,
		DependencyIndexes: file_dictionayValueService_proto_depIdxs,
		MessageInfos:      file_dictionayValueService_proto_msgTypes,
	}.Build()
	File_dictionayValueService_proto = out.File
	file_dictionayValueService_proto_rawDesc = nil
	file_dictionayValueService_proto_goTypes = nil
	file_dictionayValueService_proto_depIdxs = nil
}