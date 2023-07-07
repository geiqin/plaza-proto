// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: smsService.proto

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

//验证码
type Sms struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id"` //ID
	CreatedAt string `protobuf:"bytes,14,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt string `protobuf:"bytes,15,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
}

func (x *Sms) Reset() {
	*x = Sms{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smsService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sms) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sms) ProtoMessage() {}

func (x *Sms) ProtoReflect() protoreflect.Message {
	mi := &file_smsService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sms.ProtoReflect.Descriptor instead.
func (*Sms) Descriptor() ([]byte, []int) {
	return file_smsService_proto_rawDescGZIP(), []int{0}
}

func (x *Sms) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Sms) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Sms) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type SmsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged      int64  `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize   int64  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Sorting    string `protobuf:"bytes,3,opt,name=sorting,proto3" json:"sorting"`
	Id         int64  `protobuf:"varint,4,opt,name=id,proto3" json:"id"`
	Type       string `protobuf:"bytes,5,opt,name=type,proto3" json:"type"`
	Mobile     string `protobuf:"bytes,6,opt,name=mobile,proto3" json:"mobile"`
	Code       string `protobuf:"bytes,7,opt,name=code,proto3" json:"code"`
	TemplateId int32  `protobuf:"varint,8,opt,name=template_id,json=templateId,proto3" json:"template_id"`
	IsCheck    string `protobuf:"bytes,9,opt,name=is_check,json=isCheck,proto3" json:"is_check"`
	Status     string `protobuf:"bytes,10,opt,name=status,proto3" json:"status"`
	UserId     int64  `protobuf:"varint,11,opt,name=user_id,json=userId,proto3" json:"user_id"`
	UserType   string `protobuf:"bytes,12,opt,name=user_type,json=userType,proto3" json:"user_type"`
}

func (x *SmsRequest) Reset() {
	*x = SmsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smsService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SmsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SmsRequest) ProtoMessage() {}

func (x *SmsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_smsService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SmsRequest.ProtoReflect.Descriptor instead.
func (*SmsRequest) Descriptor() ([]byte, []int) {
	return file_smsService_proto_rawDescGZIP(), []int{1}
}

func (x *SmsRequest) GetPaged() int64 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *SmsRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *SmsRequest) GetSorting() string {
	if x != nil {
		return x.Sorting
	}
	return ""
}

func (x *SmsRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SmsRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *SmsRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *SmsRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *SmsRequest) GetTemplateId() int32 {
	if x != nil {
		return x.TemplateId
	}
	return 0
}

func (x *SmsRequest) GetIsCheck() string {
	if x != nil {
		return x.IsCheck
	}
	return ""
}

func (x *SmsRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *SmsRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *SmsRequest) GetUserType() string {
	if x != nil {
		return x.UserType
	}
	return ""
}

type SmsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Sms          `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*Sms        `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Error  *common.Error `protobuf:"bytes,4,opt,name=error,proto3" json:"error"`
	Info   *common.Info  `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
}

func (x *SmsResponse) Reset() {
	*x = SmsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smsService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SmsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SmsResponse) ProtoMessage() {}

func (x *SmsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_smsService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SmsResponse.ProtoReflect.Descriptor instead.
func (*SmsResponse) Descriptor() ([]byte, []int) {
	return file_smsService_proto_rawDescGZIP(), []int{2}
}

func (x *SmsResponse) GetEntity() *Sms {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *SmsResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *SmsResponse) GetItems() []*Sms {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *SmsResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *SmsResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_smsService_proto protoreflect.FileDescriptor

var file_smsService_proto_rawDesc = []byte{
	0x0a, 0x10, 0x73, 0x6d, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x53, 0x0a, 0x03, 0x53, 0x6d, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0xb3, 0x02, 0x0a, 0x0a, 0x53, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e,
	0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49,
	0x64, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x73, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x22, 0xc5, 0x01, 0x0a, 0x0b, 0x53,
	0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x6d, 0x73, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52,
	0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x53, 0x6d, 0x73, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x23, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x32, 0xa0, 0x01, 0x0a, 0x0a, 0x53, 0x6d, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x31, 0x0a, 0x07, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67, 0x12, 0x0d, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x6d, 0x73, 0x1a, 0x15, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0d, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x6d, 0x73, 0x1a, 0x15, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x0d, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x6d, 0x73, 0x1a, 0x15, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_smsService_proto_rawDescOnce sync.Once
	file_smsService_proto_rawDescData = file_smsService_proto_rawDesc
)

func file_smsService_proto_rawDescGZIP() []byte {
	file_smsService_proto_rawDescOnce.Do(func() {
		file_smsService_proto_rawDescData = protoimpl.X.CompressGZIP(file_smsService_proto_rawDescData)
	})
	return file_smsService_proto_rawDescData
}

var file_smsService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_smsService_proto_goTypes = []interface{}{
	(*Sms)(nil),          // 0: services.Sms
	(*SmsRequest)(nil),   // 1: services.SmsRequest
	(*SmsResponse)(nil),  // 2: services.SmsResponse
	(*common.Pager)(nil), // 3: common.Pager
	(*common.Error)(nil), // 4: common.Error
	(*common.Info)(nil),  // 5: common.Info
}
var file_smsService_proto_depIdxs = []int32{
	0, // 0: services.SmsResponse.entity:type_name -> services.Sms
	3, // 1: services.SmsResponse.pager:type_name -> common.Pager
	0, // 2: services.SmsResponse.items:type_name -> services.Sms
	4, // 3: services.SmsResponse.error:type_name -> common.Error
	5, // 4: services.SmsResponse.info:type_name -> common.Info
	0, // 5: services.SmsService.SendMsg:input_type -> services.Sms
	0, // 6: services.SmsService.Get:input_type -> services.Sms
	0, // 7: services.SmsService.Search:input_type -> services.Sms
	2, // 8: services.SmsService.SendMsg:output_type -> services.SmsResponse
	2, // 9: services.SmsService.Get:output_type -> services.SmsResponse
	2, // 10: services.SmsService.Search:output_type -> services.SmsResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_smsService_proto_init() }
func file_smsService_proto_init() {
	if File_smsService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_smsService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sms); i {
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
		file_smsService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SmsRequest); i {
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
		file_smsService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SmsResponse); i {
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
			RawDescriptor: file_smsService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_smsService_proto_goTypes,
		DependencyIndexes: file_smsService_proto_depIdxs,
		MessageInfos:      file_smsService_proto_msgTypes,
	}.Build()
	File_smsService_proto = out.File
	file_smsService_proto_rawDesc = nil
	file_smsService_proto_goTypes = nil
	file_smsService_proto_depIdxs = nil
}
