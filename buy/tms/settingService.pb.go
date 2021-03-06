// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: settingService.proto

package services

import (
	common "github.com/geiqin/microkit/protobuf/common"
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

type Setting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                    int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	IsDelivery            bool   `protobuf:"varint,2,opt,name=is_delivery,json=isDelivery,proto3" json:"is_delivery"`
	IsFetch               bool   `protobuf:"varint,3,opt,name=is_fetch,json=isFetch,proto3" json:"is_fetch"`
	IsExpress             bool   `protobuf:"varint,4,opt,name=is_express,json=isExpress,proto3" json:"is_express"`
	ExpressChargingMethod int32  `protobuf:"varint,5,opt,name=express_charging_method,json=expressChargingMethod,proto3" json:"express_charging_method"`
	ShipperId             int32  `protobuf:"varint,6,opt,name=shipper_id,json=shipperId,proto3" json:"shipper_id"`
	CreatedAt             string `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt             string `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
}

func (x *Setting) Reset() {
	*x = Setting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_settingService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Setting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Setting) ProtoMessage() {}

func (x *Setting) ProtoReflect() protoreflect.Message {
	mi := &file_settingService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Setting.ProtoReflect.Descriptor instead.
func (*Setting) Descriptor() ([]byte, []int) {
	return file_settingService_proto_rawDescGZIP(), []int{0}
}

func (x *Setting) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Setting) GetIsDelivery() bool {
	if x != nil {
		return x.IsDelivery
	}
	return false
}

func (x *Setting) GetIsFetch() bool {
	if x != nil {
		return x.IsFetch
	}
	return false
}

func (x *Setting) GetIsExpress() bool {
	if x != nil {
		return x.IsExpress
	}
	return false
}

func (x *Setting) GetExpressChargingMethod() int32 {
	if x != nil {
		return x.ExpressChargingMethod
	}
	return 0
}

func (x *Setting) GetShipperId() int32 {
	if x != nil {
		return x.ShipperId
	}
	return 0
}

func (x *Setting) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Setting) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type SettingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Setting      `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*Setting    `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Error  *common.Error `protobuf:"bytes,4,opt,name=error,proto3" json:"error"`
	Info   *common.Info  `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
}

func (x *SettingResponse) Reset() {
	*x = SettingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_settingService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SettingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SettingResponse) ProtoMessage() {}

func (x *SettingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_settingService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SettingResponse.ProtoReflect.Descriptor instead.
func (*SettingResponse) Descriptor() ([]byte, []int) {
	return file_settingService_proto_rawDescGZIP(), []int{1}
}

func (x *SettingResponse) GetEntity() *Setting {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *SettingResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *SettingResponse) GetItems() []*Setting {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *SettingResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *SettingResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_settingService_proto protoreflect.FileDescriptor

var file_settingService_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x89, 0x02, 0x0a, 0x07, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79,
	0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x66, 0x65, 0x74, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x46, 0x65, 0x74, 0x63, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x69,
	0x73, 0x5f, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x69, 0x73, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x12, 0x36, 0x0a, 0x17, 0x65, 0x78,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x15, 0x65, 0x78, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x43, 0x68, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22,
	0xd1, 0x01, 0x0a, 0x0f, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23,
	0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61,
	0x67, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x23, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69,
	0x6e, 0x66, 0x6f, 0x32, 0x47, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x11, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x1a,
	0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a,
	0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_settingService_proto_rawDescOnce sync.Once
	file_settingService_proto_rawDescData = file_settingService_proto_rawDesc
)

func file_settingService_proto_rawDescGZIP() []byte {
	file_settingService_proto_rawDescOnce.Do(func() {
		file_settingService_proto_rawDescData = protoimpl.X.CompressGZIP(file_settingService_proto_rawDescData)
	})
	return file_settingService_proto_rawDescData
}

var file_settingService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_settingService_proto_goTypes = []interface{}{
	(*Setting)(nil),         // 0: services.Setting
	(*SettingResponse)(nil), // 1: services.SettingResponse
	(*common.Pager)(nil),    // 2: common.Pager
	(*common.Error)(nil),    // 3: common.Error
	(*common.Info)(nil),     // 4: common.Info
}
var file_settingService_proto_depIdxs = []int32{
	0, // 0: services.SettingResponse.entity:type_name -> services.Setting
	2, // 1: services.SettingResponse.pager:type_name -> common.Pager
	0, // 2: services.SettingResponse.items:type_name -> services.Setting
	3, // 3: services.SettingResponse.error:type_name -> common.Error
	4, // 4: services.SettingResponse.info:type_name -> common.Info
	0, // 5: services.SettingService.Get:input_type -> services.Setting
	1, // 6: services.SettingService.Get:output_type -> services.SettingResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_settingService_proto_init() }
func file_settingService_proto_init() {
	if File_settingService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_settingService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Setting); i {
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
		file_settingService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SettingResponse); i {
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
			RawDescriptor: file_settingService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_settingService_proto_goTypes,
		DependencyIndexes: file_settingService_proto_depIdxs,
		MessageInfos:      file_settingService_proto_msgTypes,
	}.Build()
	File_settingService_proto = out.File
	file_settingService_proto_rawDesc = nil
	file_settingService_proto_goTypes = nil
	file_settingService_proto_depIdxs = nil
}