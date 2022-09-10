// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: memberInfoService.proto

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

type MemberInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	MemberSn    string `protobuf:"bytes,2,opt,name=member_sn,json=memberSn,proto3" json:"member_sn,omitempty"`
	Name        string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	DisplayName string `protobuf:"bytes,4,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	RealName    string `protobuf:"bytes,5,opt,name=real_name,json=realName,proto3" json:"real_name,omitempty"`
	IdCard      string `protobuf:"bytes,6,opt,name=id_card,json=idCard,proto3" json:"id_card,omitempty"`
	Gender      int32  `protobuf:"varint,7,opt,name=gender,proto3" json:"gender,omitempty"`
	Mobile      string `protobuf:"bytes,8,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Email       string `protobuf:"bytes,9,opt,name=email,proto3" json:"email,omitempty"`
	AvatarId    int64  `protobuf:"varint,10,opt,name=avatar_id,json=avatarId,proto3" json:"avatar_id,omitempty"`
	AvatarUrl   string `protobuf:"bytes,11,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
}

func (x *MemberInfo) Reset() {
	*x = MemberInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memberInfoService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberInfo) ProtoMessage() {}

func (x *MemberInfo) ProtoReflect() protoreflect.Message {
	mi := &file_memberInfoService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberInfo.ProtoReflect.Descriptor instead.
func (*MemberInfo) Descriptor() ([]byte, []int) {
	return file_memberInfoService_proto_rawDescGZIP(), []int{0}
}

func (x *MemberInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MemberInfo) GetMemberSn() string {
	if x != nil {
		return x.MemberSn
	}
	return ""
}

func (x *MemberInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MemberInfo) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *MemberInfo) GetRealName() string {
	if x != nil {
		return x.RealName
	}
	return ""
}

func (x *MemberInfo) GetIdCard() string {
	if x != nil {
		return x.IdCard
	}
	return ""
}

func (x *MemberInfo) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *MemberInfo) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *MemberInfo) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *MemberInfo) GetAvatarId() int64 {
	if x != nil {
		return x.AvatarId
	}
	return 0
}

func (x *MemberInfo) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

var File_memberInfoService_proto protoreflect.FileDescriptor

var file_memberInfoService_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x22, 0xa8, 0x02, 0x0a, 0x0a, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x73, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x53, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x61, 0x6c, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x61, 0x6c, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x64, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x64, 0x43, 0x61, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x67, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x42, 0x0d,
	0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_memberInfoService_proto_rawDescOnce sync.Once
	file_memberInfoService_proto_rawDescData = file_memberInfoService_proto_rawDesc
)

func file_memberInfoService_proto_rawDescGZIP() []byte {
	file_memberInfoService_proto_rawDescOnce.Do(func() {
		file_memberInfoService_proto_rawDescData = protoimpl.X.CompressGZIP(file_memberInfoService_proto_rawDescData)
	})
	return file_memberInfoService_proto_rawDescData
}

var file_memberInfoService_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_memberInfoService_proto_goTypes = []interface{}{
	(*MemberInfo)(nil), // 0: services.MemberInfo
}
var file_memberInfoService_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_memberInfoService_proto_init() }
func file_memberInfoService_proto_init() {
	if File_memberInfoService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_memberInfoService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberInfo); i {
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
			RawDescriptor: file_memberInfoService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_memberInfoService_proto_goTypes,
		DependencyIndexes: file_memberInfoService_proto_depIdxs,
		MessageInfos:      file_memberInfoService_proto_msgTypes,
	}.Build()
	File_memberInfoService_proto = out.File
	file_memberInfoService_proto_rawDesc = nil
	file_memberInfoService_proto_goTypes = nil
	file_memberInfoService_proto_depIdxs = nil
}
