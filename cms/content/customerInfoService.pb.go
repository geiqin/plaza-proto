// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: customerInfoService.proto

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

type Customer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	CustomerSn  string `protobuf:"bytes,2,opt,name=customer_sn,json=customerSn,proto3" json:"customer_sn"`
	Name        string `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`
	DisplayName string `protobuf:"bytes,4,opt,name=display_name,json=displayName,proto3" json:"display_name"`
	Realname    string `protobuf:"bytes,5,opt,name=realname,proto3" json:"realname"`
	Idcard      string `protobuf:"bytes,6,opt,name=idcard,proto3" json:"idcard"`
	Gender      int32  `protobuf:"varint,7,opt,name=gender,proto3" json:"gender"`
	Mobile      string `protobuf:"bytes,8,opt,name=mobile,proto3" json:"mobile"`
	Email       string `protobuf:"bytes,9,opt,name=email,proto3" json:"email"`
	AvatarId    int64  `protobuf:"varint,10,opt,name=avatar_id,json=avatarId,proto3" json:"avatar_id"`
	AvatarUrl   string `protobuf:"bytes,11,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url"`
}

func (x *Customer) Reset() {
	*x = Customer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customerInfoService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Customer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Customer) ProtoMessage() {}

func (x *Customer) ProtoReflect() protoreflect.Message {
	mi := &file_customerInfoService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Customer.ProtoReflect.Descriptor instead.
func (*Customer) Descriptor() ([]byte, []int) {
	return file_customerInfoService_proto_rawDescGZIP(), []int{0}
}

func (x *Customer) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Customer) GetCustomerSn() string {
	if x != nil {
		return x.CustomerSn
	}
	return ""
}

func (x *Customer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Customer) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Customer) GetRealname() string {
	if x != nil {
		return x.Realname
	}
	return ""
}

func (x *Customer) GetIdcard() string {
	if x != nil {
		return x.Idcard
	}
	return ""
}

func (x *Customer) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *Customer) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *Customer) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Customer) GetAvatarId() int64 {
	if x != nil {
		return x.AvatarId
	}
	return 0
}

func (x *Customer) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

var File_customerInfoService_proto protoreflect.FileDescriptor

var file_customerInfoService_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0xa8, 0x02, 0x0a, 0x08, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x73,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x53, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65,
	0x61, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65,
	0x61, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x64, 0x63, 0x61, 0x72, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x64, 0x63, 0x61, 0x72, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x6c,
	0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_customerInfoService_proto_rawDescOnce sync.Once
	file_customerInfoService_proto_rawDescData = file_customerInfoService_proto_rawDesc
)

func file_customerInfoService_proto_rawDescGZIP() []byte {
	file_customerInfoService_proto_rawDescOnce.Do(func() {
		file_customerInfoService_proto_rawDescData = protoimpl.X.CompressGZIP(file_customerInfoService_proto_rawDescData)
	})
	return file_customerInfoService_proto_rawDescData
}

var file_customerInfoService_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_customerInfoService_proto_goTypes = []interface{}{
	(*Customer)(nil), // 0: services.Customer
}
var file_customerInfoService_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_customerInfoService_proto_init() }
func file_customerInfoService_proto_init() {
	if File_customerInfoService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_customerInfoService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Customer); i {
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
			RawDescriptor: file_customerInfoService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_customerInfoService_proto_goTypes,
		DependencyIndexes: file_customerInfoService_proto_depIdxs,
		MessageInfos:      file_customerInfoService_proto_msgTypes,
	}.Build()
	File_customerInfoService_proto = out.File
	file_customerInfoService_proto_rawDesc = nil
	file_customerInfoService_proto_goTypes = nil
	file_customerInfoService_proto_depIdxs = nil
}
