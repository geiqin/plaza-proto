// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: payAccountService.proto

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

type PayAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int32         `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Code         string        `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Title        string        `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Type         string        `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	ProviderMode string        `protobuf:"bytes,5,opt,name=provider_mode,json=providerMode,proto3" json:"provider_mode,omitempty"`
	Version      string        `protobuf:"bytes,6,opt,name=version,proto3" json:"version,omitempty"`
	ConfigData   string        `protobuf:"bytes,7,opt,name=config_data,json=configData,proto3" json:"config_data,omitempty"`
	Status       string        `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
	CreatedAt    string        `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt    string        `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	WxConfig     *WxConfig     `protobuf:"bytes,11,opt,name=wx_config,json=wxConfig,proto3" json:"wx_config,omitempty"`
	AlipayConfig *AlipayConfig `protobuf:"bytes,12,opt,name=alipay_config,json=alipayConfig,proto3" json:"alipay_config,omitempty"`
}

func (x *PayAccount) Reset() {
	*x = PayAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payAccountService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayAccount) ProtoMessage() {}

func (x *PayAccount) ProtoReflect() protoreflect.Message {
	mi := &file_payAccountService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayAccount.ProtoReflect.Descriptor instead.
func (*PayAccount) Descriptor() ([]byte, []int) {
	return file_payAccountService_proto_rawDescGZIP(), []int{0}
}

func (x *PayAccount) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PayAccount) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *PayAccount) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PayAccount) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *PayAccount) GetProviderMode() string {
	if x != nil {
		return x.ProviderMode
	}
	return ""
}

func (x *PayAccount) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *PayAccount) GetConfigData() string {
	if x != nil {
		return x.ConfigData
	}
	return ""
}

func (x *PayAccount) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *PayAccount) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *PayAccount) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *PayAccount) GetWxConfig() *WxConfig {
	if x != nil {
		return x.WxConfig
	}
	return nil
}

func (x *PayAccount) GetAlipayConfig() *AlipayConfig {
	if x != nil {
		return x.AlipayConfig
	}
	return nil
}

type WxConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MchId             string `protobuf:"bytes,1,opt,name=mch_id,json=mchId,proto3" json:"mch_id,omitempty"`
	SignType          string `protobuf:"bytes,2,opt,name=sign_type,json=signType,proto3" json:"sign_type,omitempty"`
	Md5Key            string `protobuf:"bytes,3,opt,name=md5_key,json=md5Key,proto3" json:"md5_key,omitempty"`
	CertSerialNo      string `protobuf:"bytes,4,opt,name=cert_serial_no,json=certSerialNo,proto3" json:"cert_serial_no,omitempty"`
	PrivateKey        string `protobuf:"bytes,5,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	PublicCert        string `protobuf:"bytes,6,opt,name=public_cert,json=publicCert,proto3" json:"public_cert,omitempty"`
	V3ApiKey          string `protobuf:"bytes,7,opt,name=v3_api_key,json=v3ApiKey,proto3" json:"v3_api_key,omitempty"`
	V3PrivateKey      string `protobuf:"bytes,8,opt,name=v3_private_key,json=v3PrivateKey,proto3" json:"v3_private_key,omitempty"`
	ApiClientCertFile string `protobuf:"bytes,9,opt,name=api_client_cert_file,json=apiClientCertFile,proto3" json:"api_client_cert_file,omitempty"`
	ApiClientKeyFile  string `protobuf:"bytes,10,opt,name=api_client_key_file,json=apiClientKeyFile,proto3" json:"api_client_key_file,omitempty"`
	ApiClientP12File  string `protobuf:"bytes,11,opt,name=api_client_p12_file,json=apiClientP12File,proto3" json:"api_client_p12_file,omitempty"`
}

func (x *WxConfig) Reset() {
	*x = WxConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payAccountService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WxConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WxConfig) ProtoMessage() {}

func (x *WxConfig) ProtoReflect() protoreflect.Message {
	mi := &file_payAccountService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WxConfig.ProtoReflect.Descriptor instead.
func (*WxConfig) Descriptor() ([]byte, []int) {
	return file_payAccountService_proto_rawDescGZIP(), []int{1}
}

func (x *WxConfig) GetMchId() string {
	if x != nil {
		return x.MchId
	}
	return ""
}

func (x *WxConfig) GetSignType() string {
	if x != nil {
		return x.SignType
	}
	return ""
}

func (x *WxConfig) GetMd5Key() string {
	if x != nil {
		return x.Md5Key
	}
	return ""
}

func (x *WxConfig) GetCertSerialNo() string {
	if x != nil {
		return x.CertSerialNo
	}
	return ""
}

func (x *WxConfig) GetPrivateKey() string {
	if x != nil {
		return x.PrivateKey
	}
	return ""
}

func (x *WxConfig) GetPublicCert() string {
	if x != nil {
		return x.PublicCert
	}
	return ""
}

func (x *WxConfig) GetV3ApiKey() string {
	if x != nil {
		return x.V3ApiKey
	}
	return ""
}

func (x *WxConfig) GetV3PrivateKey() string {
	if x != nil {
		return x.V3PrivateKey
	}
	return ""
}

func (x *WxConfig) GetApiClientCertFile() string {
	if x != nil {
		return x.ApiClientCertFile
	}
	return ""
}

func (x *WxConfig) GetApiClientKeyFile() string {
	if x != nil {
		return x.ApiClientKeyFile
	}
	return ""
}

func (x *WxConfig) GetApiClientP12File() string {
	if x != nil {
		return x.ApiClientP12File
	}
	return ""
}

type AlipayConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MchId string `protobuf:"bytes,1,opt,name=mch_id,json=mchId,proto3" json:"mch_id,omitempty"`
}

func (x *AlipayConfig) Reset() {
	*x = AlipayConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payAccountService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlipayConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlipayConfig) ProtoMessage() {}

func (x *AlipayConfig) ProtoReflect() protoreflect.Message {
	mi := &file_payAccountService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlipayConfig.ProtoReflect.Descriptor instead.
func (*AlipayConfig) Descriptor() ([]byte, []int) {
	return file_payAccountService_proto_rawDescGZIP(), []int{2}
}

func (x *AlipayConfig) GetMchId() string {
	if x != nil {
		return x.MchId
	}
	return ""
}

type PayAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged        int32  `protobuf:"varint,1,opt,name=paged,proto3" json:"paged,omitempty"`
	PageSize     int32  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Keywords     string `protobuf:"bytes,3,opt,name=keywords,proto3" json:"keywords,omitempty"`
	Id           int32  `protobuf:"varint,5,opt,name=id,proto3" json:"id,omitempty"`
	Code         string `protobuf:"bytes,6,opt,name=code,proto3" json:"code,omitempty"`
	Name         string `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	Type         string `protobuf:"bytes,8,opt,name=type,proto3" json:"type,omitempty"`
	ProviderMode string `protobuf:"bytes,9,opt,name=provider_mode,json=providerMode,proto3" json:"provider_mode,omitempty"`
	Status       string `protobuf:"bytes,10,opt,name=status,proto3" json:"status,omitempty"`
	Reason       string `protobuf:"bytes,11,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (x *PayAccountRequest) Reset() {
	*x = PayAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payAccountService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayAccountRequest) ProtoMessage() {}

func (x *PayAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payAccountService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayAccountRequest.ProtoReflect.Descriptor instead.
func (*PayAccountRequest) Descriptor() ([]byte, []int) {
	return file_payAccountService_proto_rawDescGZIP(), []int{3}
}

func (x *PayAccountRequest) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *PayAccountRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *PayAccountRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *PayAccountRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PayAccountRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *PayAccountRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PayAccountRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *PayAccountRequest) GetProviderMode() string {
	if x != nil {
		return x.ProviderMode
	}
	return ""
}

func (x *PayAccountRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *PayAccountRequest) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

type PayAccountData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *PayAccount   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items  []*PayAccount `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Info   *common.Info  `protobuf:"bytes,4,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *PayAccountData) Reset() {
	*x = PayAccountData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payAccountService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayAccountData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayAccountData) ProtoMessage() {}

func (x *PayAccountData) ProtoReflect() protoreflect.Message {
	mi := &file_payAccountService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayAccountData.ProtoReflect.Descriptor instead.
func (*PayAccountData) Descriptor() ([]byte, []int) {
	return file_payAccountService_proto_rawDescGZIP(), []int{4}
}

func (x *PayAccountData) GetEntity() *PayAccount {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *PayAccountData) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *PayAccountData) GetItems() []*PayAccount {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *PayAccountData) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type PayAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *PayAccountData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Error *common.Error   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *PayAccountResponse) Reset() {
	*x = PayAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payAccountService_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayAccountResponse) ProtoMessage() {}

func (x *PayAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payAccountService_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayAccountResponse.ProtoReflect.Descriptor instead.
func (*PayAccountResponse) Descriptor() ([]byte, []int) {
	return file_payAccountService_proto_rawDescGZIP(), []int{5}
}

func (x *PayAccountResponse) GetData() *PayAccountData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *PayAccountResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_payAccountService_proto protoreflect.FileDescriptor

var file_payAccountService_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x61, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfe, 0x02, 0x0a, 0x0a, 0x50, 0x61, 0x79, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f,
	0x6d, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x2f, 0x0a, 0x09, 0x77, 0x78, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x57, 0x78, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x08, 0x77, 0x78, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3b, 0x0a, 0x0d, 0x61, 0x6c,
	0x69, 0x70, 0x61, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x6c, 0x69,
	0x70, 0x61, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0c, 0x61, 0x6c, 0x69, 0x70, 0x61,
	0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x92, 0x03, 0x0a, 0x08, 0x57, 0x78, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x15, 0x0a, 0x06, 0x6d, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x63, 0x68, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73,
	0x69, 0x67, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x73, 0x69, 0x67, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x6d, 0x64, 0x35, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x64, 0x35, 0x4b, 0x65,
	0x79, 0x12, 0x24, 0x0a, 0x0e, 0x63, 0x65, 0x72, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x5f, 0x6e, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x65, 0x72, 0x74, 0x53,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x43, 0x65, 0x72, 0x74, 0x12, 0x1c, 0x0a, 0x0a, 0x76, 0x33, 0x5f,
	0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x76,
	0x33, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x12, 0x24, 0x0a, 0x0e, 0x76, 0x33, 0x5f, 0x70, 0x72,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x76, 0x33, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x2f, 0x0a,
	0x14, 0x61, 0x70, 0x69, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x65, 0x72, 0x74,
	0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x61, 0x70, 0x69,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x65, 0x72, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x2d,
	0x0a, 0x13, 0x61, 0x70, 0x69, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x6b, 0x65, 0x79,
	0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x61, 0x70, 0x69,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x2d, 0x0a,
	0x13, 0x61, 0x70, 0x69, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x31, 0x32, 0x5f,
	0x66, 0x69, 0x6c, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x61, 0x70, 0x69, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x31, 0x32, 0x46, 0x69, 0x6c, 0x65, 0x22, 0x25, 0x0a, 0x0c,
	0x41, 0x6c, 0x69, 0x70, 0x61, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x15, 0x0a, 0x06,
	0x6d, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x63,
	0x68, 0x49, 0x64, 0x22, 0x83, 0x02, 0x0a, 0x11, 0x50, 0x61, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67,
	0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22, 0xb1, 0x01, 0x0a, 0x0e, 0x50, 0x61,
	0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2c, 0x0a, 0x06,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61,
	0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12,
	0x2a, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x20, 0x0a, 0x04, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x67, 0x0a,
	0x12, 0x50, 0x61, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x79,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xd7, 0x02, 0x0a, 0x11, 0x50, 0x61, 0x79, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x06,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x50, 0x61, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x1c, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x06,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x50, 0x61, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x1c, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x06,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x50, 0x61, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x1c, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x03,
	0x47, 0x65, 0x74, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50,
	0x61, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x06, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x12, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50,
	0x61, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_payAccountService_proto_rawDescOnce sync.Once
	file_payAccountService_proto_rawDescData = file_payAccountService_proto_rawDesc
)

func file_payAccountService_proto_rawDescGZIP() []byte {
	file_payAccountService_proto_rawDescOnce.Do(func() {
		file_payAccountService_proto_rawDescData = protoimpl.X.CompressGZIP(file_payAccountService_proto_rawDescData)
	})
	return file_payAccountService_proto_rawDescData
}

var file_payAccountService_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_payAccountService_proto_goTypes = []interface{}{
	(*PayAccount)(nil),         // 0: services.PayAccount
	(*WxConfig)(nil),           // 1: services.WxConfig
	(*AlipayConfig)(nil),       // 2: services.AlipayConfig
	(*PayAccountRequest)(nil),  // 3: services.PayAccountRequest
	(*PayAccountData)(nil),     // 4: services.PayAccountData
	(*PayAccountResponse)(nil), // 5: services.PayAccountResponse
	(*common.Pager)(nil),       // 6: common.Pager
	(*common.Info)(nil),        // 7: common.Info
	(*common.Error)(nil),       // 8: common.Error
}
var file_payAccountService_proto_depIdxs = []int32{
	1,  // 0: services.PayAccount.wx_config:type_name -> services.WxConfig
	2,  // 1: services.PayAccount.alipay_config:type_name -> services.AlipayConfig
	0,  // 2: services.PayAccountData.entity:type_name -> services.PayAccount
	6,  // 3: services.PayAccountData.pager:type_name -> common.Pager
	0,  // 4: services.PayAccountData.items:type_name -> services.PayAccount
	7,  // 5: services.PayAccountData.info:type_name -> common.Info
	4,  // 6: services.PayAccountResponse.data:type_name -> services.PayAccountData
	8,  // 7: services.PayAccountResponse.error:type_name -> common.Error
	0,  // 8: services.PayAccountService.Create:input_type -> services.PayAccount
	0,  // 9: services.PayAccountService.Update:input_type -> services.PayAccount
	0,  // 10: services.PayAccountService.Delete:input_type -> services.PayAccount
	0,  // 11: services.PayAccountService.Get:input_type -> services.PayAccount
	3,  // 12: services.PayAccountService.Search:input_type -> services.PayAccountRequest
	5,  // 13: services.PayAccountService.Create:output_type -> services.PayAccountResponse
	5,  // 14: services.PayAccountService.Update:output_type -> services.PayAccountResponse
	5,  // 15: services.PayAccountService.Delete:output_type -> services.PayAccountResponse
	5,  // 16: services.PayAccountService.Get:output_type -> services.PayAccountResponse
	5,  // 17: services.PayAccountService.Search:output_type -> services.PayAccountResponse
	13, // [13:18] is the sub-list for method output_type
	8,  // [8:13] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_payAccountService_proto_init() }
func file_payAccountService_proto_init() {
	if File_payAccountService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_payAccountService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayAccount); i {
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
		file_payAccountService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WxConfig); i {
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
		file_payAccountService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlipayConfig); i {
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
		file_payAccountService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayAccountRequest); i {
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
		file_payAccountService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayAccountData); i {
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
		file_payAccountService_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayAccountResponse); i {
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
			RawDescriptor: file_payAccountService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_payAccountService_proto_goTypes,
		DependencyIndexes: file_payAccountService_proto_depIdxs,
		MessageInfos:      file_payAccountService_proto_msgTypes,
	}.Build()
	File_payAccountService_proto = out.File
	file_payAccountService_proto_rawDesc = nil
	file_payAccountService_proto_goTypes = nil
	file_payAccountService_proto_depIdxs = nil
}
