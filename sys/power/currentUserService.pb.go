// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: currentUserService.proto

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

// 授权请求信息
type CurrentUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	Username    string `protobuf:"bytes,3,opt,name=username,proto3" json:"username"`
	Password    string `protobuf:"bytes,4,opt,name=password,proto3" json:"password"`
	OldPassword string `protobuf:"bytes,5,opt,name=old_password,json=oldPassword,proto3" json:"old_password"`
	Email       string `protobuf:"bytes,6,opt,name=email,proto3" json:"email"`
	Mobile      string `protobuf:"bytes,7,opt,name=mobile,proto3" json:"mobile"`
	Tag         string `protobuf:"bytes,8,opt,name=tag,proto3" json:"tag"`
}

func (x *CurrentUserRequest) Reset() {
	*x = CurrentUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_currentUserService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrentUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrentUserRequest) ProtoMessage() {}

func (x *CurrentUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_currentUserService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrentUserRequest.ProtoReflect.Descriptor instead.
func (*CurrentUserRequest) Descriptor() ([]byte, []int) {
	return file_currentUserService_proto_rawDescGZIP(), []int{0}
}

func (x *CurrentUserRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CurrentUserRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CurrentUserRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CurrentUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CurrentUserRequest) GetOldPassword() string {
	if x != nil {
		return x.OldPassword
	}
	return ""
}

func (x *CurrentUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CurrentUserRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *CurrentUserRequest) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

type CurrentUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User        *User         `protobuf:"bytes,1,opt,name=user,proto3" json:"user"`
	Navs        []*Nav        `protobuf:"bytes,2,rep,name=navs,proto3" json:"navs"`
	Permissions []string      `protobuf:"bytes,3,rep,name=permissions,proto3" json:"permissions"`
	Pager       *common.Pager `protobuf:"bytes,4,opt,name=pager,proto3" json:"pager"`
	Items       []*User       `protobuf:"bytes,5,rep,name=items,proto3" json:"items"`
	Info        string        `protobuf:"bytes,6,opt,name=info,proto3" json:"info"`
}

func (x *CurrentUserResponse) Reset() {
	*x = CurrentUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_currentUserService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrentUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrentUserResponse) ProtoMessage() {}

func (x *CurrentUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_currentUserService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrentUserResponse.ProtoReflect.Descriptor instead.
func (*CurrentUserResponse) Descriptor() ([]byte, []int) {
	return file_currentUserService_proto_rawDescGZIP(), []int{1}
}

func (x *CurrentUserResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *CurrentUserResponse) GetNavs() []*Nav {
	if x != nil {
		return x.Navs
	}
	return nil
}

func (x *CurrentUserResponse) GetPermissions() []string {
	if x != nil {
		return x.Permissions
	}
	return nil
}

func (x *CurrentUserResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *CurrentUserResponse) GetItems() []*User {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *CurrentUserResponse) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

var File_currentUserService_proto protoreflect.FileDescriptor

var file_currentUserService_proto_rawDesc = []byte{
	0x0a, 0x18, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x75, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x6e, 0x61, 0x76, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd3, 0x01, 0x0a,
	0x12, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12,
	0x21, 0x0a, 0x0c, 0x6f, 0x6c, 0x64, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x6c, 0x64, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74,
	0x61, 0x67, 0x22, 0xdd, 0x01, 0x0a, 0x13, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x21,
	0x0a, 0x04, 0x6e, 0x61, 0x76, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4e, 0x61, 0x76, 0x52, 0x04, 0x6e, 0x61, 0x76,
	0x73, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65,
	0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x32, 0xe0, 0x02, 0x0a, 0x12, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x05, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x49, 0x0a, 0x08, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x0c,
	0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x0e, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x16, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x0c, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79,
	0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x43, 0x0a, 0x09, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x50, 0x77, 0x64, 0x12, 0x1c, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_currentUserService_proto_rawDescOnce sync.Once
	file_currentUserService_proto_rawDescData = file_currentUserService_proto_rawDesc
)

func file_currentUserService_proto_rawDescGZIP() []byte {
	file_currentUserService_proto_rawDescOnce.Do(func() {
		file_currentUserService_proto_rawDescData = protoimpl.X.CompressGZIP(file_currentUserService_proto_rawDescData)
	})
	return file_currentUserService_proto_rawDescData
}

var file_currentUserService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_currentUserService_proto_goTypes = []interface{}{
	(*CurrentUserRequest)(nil),  // 0: services.CurrentUserRequest
	(*CurrentUserResponse)(nil), // 1: services.CurrentUserResponse
	(*User)(nil),                // 2: services.User
	(*Nav)(nil),                 // 3: services.Nav
	(*common.Pager)(nil),        // 4: common.Pager
	(*UserResponse)(nil),        // 5: services.UserResponse
}
var file_currentUserService_proto_depIdxs = []int32{
	2, // 0: services.CurrentUserResponse.user:type_name -> services.User
	3, // 1: services.CurrentUserResponse.navs:type_name -> services.Nav
	4, // 2: services.CurrentUserResponse.pager:type_name -> common.Pager
	2, // 3: services.CurrentUserResponse.items:type_name -> services.User
	0, // 4: services.CurrentUserService.Index:input_type -> services.CurrentUserRequest
	0, // 5: services.CurrentUserService.BaseInfo:input_type -> services.CurrentUserRequest
	2, // 6: services.CurrentUserService.ModifyMobile:input_type -> services.User
	2, // 7: services.CurrentUserService.ModifyAvatar:input_type -> services.User
	0, // 8: services.CurrentUserService.ModifyPwd:input_type -> services.CurrentUserRequest
	1, // 9: services.CurrentUserService.Index:output_type -> services.CurrentUserResponse
	1, // 10: services.CurrentUserService.BaseInfo:output_type -> services.CurrentUserResponse
	5, // 11: services.CurrentUserService.ModifyMobile:output_type -> services.UserResponse
	5, // 12: services.CurrentUserService.ModifyAvatar:output_type -> services.UserResponse
	5, // 13: services.CurrentUserService.ModifyPwd:output_type -> services.UserResponse
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_currentUserService_proto_init() }
func file_currentUserService_proto_init() {
	if File_currentUserService_proto != nil {
		return
	}
	file_userService_proto_init()
	file_navService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_currentUserService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CurrentUserRequest); i {
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
		file_currentUserService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CurrentUserResponse); i {
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
			RawDescriptor: file_currentUserService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_currentUserService_proto_goTypes,
		DependencyIndexes: file_currentUserService_proto_depIdxs,
		MessageInfos:      file_currentUserService_proto_msgTypes,
	}.Build()
	File_currentUserService_proto = out.File
	file_currentUserService_proto_rawDesc = nil
	file_currentUserService_proto_goTypes = nil
	file_currentUserService_proto_depIdxs = nil
}
