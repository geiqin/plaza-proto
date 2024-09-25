// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: storageService.proto

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

// 存储
type Storage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`                                //ID
	Code        string `protobuf:"bytes,2,opt,name=code,proto3" json:"code"`                             //编码
	Name        string `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`                             //名称
	Type        string `protobuf:"bytes,4,opt,name=type,proto3" json:"type"`                             //类型
	AccessKey   string `protobuf:"bytes,5,opt,name=access_key,json=accessKey,proto3" json:"access_key"`  //访问密钥
	SecretKey   string `protobuf:"bytes,6,opt,name=secret_key,json=secretKey,proto3" json:"secret_key"`  //私有密钥
	Bucket      string `protobuf:"bytes,7,opt,name=bucket,proto3" json:"bucket"`                         //桶名称
	Endpoint    string `protobuf:"bytes,8,opt,name=endpoint,proto3" json:"endpoint"`                     //终端节点
	Zone        string `protobuf:"bytes,9,opt,name=zone,proto3" json:"zone"`                             //储存区
	Domain      string `protobuf:"bytes,10,opt,name=domain,proto3" json:"domain"`                        //访问域名
	HostUrl     string `protobuf:"bytes,11,opt,name=host_url,json=hostUrl,proto3" json:"host_url"`       //主机地址
	IsDefault   string `protobuf:"bytes,12,opt,name=is_default,json=isDefault,proto3" json:"is_default"` //是否默认
	IsPrivate   string `protobuf:"bytes,13,opt,name=is_private,json=isPrivate,proto3" json:"is_private"` //是否私有
	Description string `protobuf:"bytes,14,opt,name=description,proto3" json:"description"`              //描述
	Sort        int32  `protobuf:"varint,15,opt,name=sort,proto3" json:"sort"`                           //排序
	Status      string `protobuf:"bytes,16,opt,name=status,proto3" json:"status"`                        //状态
	CreatedAt   string `protobuf:"bytes,17,opt,name=created_at,json=createdAt,proto3" json:"created_at"` //创建时间
	UpdatedAt   string `protobuf:"bytes,18,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"` //修改时间
}

func (x *Storage) Reset() {
	*x = Storage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storageService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Storage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Storage) ProtoMessage() {}

func (x *Storage) ProtoReflect() protoreflect.Message {
	mi := &file_storageService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Storage.ProtoReflect.Descriptor instead.
func (*Storage) Descriptor() ([]byte, []int) {
	return file_storageService_proto_rawDescGZIP(), []int{0}
}

func (x *Storage) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Storage) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Storage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Storage) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Storage) GetAccessKey() string {
	if x != nil {
		return x.AccessKey
	}
	return ""
}

func (x *Storage) GetSecretKey() string {
	if x != nil {
		return x.SecretKey
	}
	return ""
}

func (x *Storage) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *Storage) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *Storage) GetZone() string {
	if x != nil {
		return x.Zone
	}
	return ""
}

func (x *Storage) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *Storage) GetHostUrl() string {
	if x != nil {
		return x.HostUrl
	}
	return ""
}

func (x *Storage) GetIsDefault() string {
	if x != nil {
		return x.IsDefault
	}
	return ""
}

func (x *Storage) GetIsPrivate() string {
	if x != nil {
		return x.IsPrivate
	}
	return ""
}

func (x *Storage) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Storage) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *Storage) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Storage) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Storage) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

// 存储请求参数
type StorageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Top       int32    `protobuf:"varint,1,opt,name=top,proto3" json:"top"`
	Paged     int64    `protobuf:"varint,2,opt,name=paged,proto3" json:"paged"`
	PageSize  int64    `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Keywords  string   `protobuf:"bytes,4,opt,name=keywords,proto3" json:"keywords"`
	Sort      []string `protobuf:"bytes,5,rep,name=sort,proto3" json:"sort"`
	DateRange []string `protobuf:"bytes,6,rep,name=date_range,json=dateRange,proto3" json:"date_range"`
	Ids       []int64  `protobuf:"varint,7,rep,packed,name=ids,proto3" json:"ids"`
	Id        int64    `protobuf:"varint,8,opt,name=id,proto3" json:"id"`
	//以下为自定义参数
	Code      string `protobuf:"bytes,11,opt,name=code,proto3" json:"code"`                            //编码
	Name      string `protobuf:"bytes,12,opt,name=name,proto3" json:"name"`                            //名称
	Type      string `protobuf:"bytes,13,opt,name=type,proto3" json:"type"`                            //类型
	Bucket    string `protobuf:"bytes,14,opt,name=bucket,proto3" json:"bucket"`                        //桶名称
	IsDefault string `protobuf:"bytes,15,opt,name=is_default,json=isDefault,proto3" json:"is_default"` //是否默认
	IsPrivate string `protobuf:"bytes,16,opt,name=is_private,json=isPrivate,proto3" json:"is_private"` //是否私有
	Status    string `protobuf:"bytes,17,opt,name=status,proto3" json:"status"`                        //状态
}

func (x *StorageRequest) Reset() {
	*x = StorageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storageService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageRequest) ProtoMessage() {}

func (x *StorageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_storageService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageRequest.ProtoReflect.Descriptor instead.
func (*StorageRequest) Descriptor() ([]byte, []int) {
	return file_storageService_proto_rawDescGZIP(), []int{1}
}

func (x *StorageRequest) GetTop() int32 {
	if x != nil {
		return x.Top
	}
	return 0
}

func (x *StorageRequest) GetPaged() int64 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *StorageRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *StorageRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *StorageRequest) GetSort() []string {
	if x != nil {
		return x.Sort
	}
	return nil
}

func (x *StorageRequest) GetDateRange() []string {
	if x != nil {
		return x.DateRange
	}
	return nil
}

func (x *StorageRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *StorageRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *StorageRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *StorageRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StorageRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *StorageRequest) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *StorageRequest) GetIsDefault() string {
	if x != nil {
		return x.IsDefault
	}
	return ""
}

func (x *StorageRequest) GetIsPrivate() string {
	if x != nil {
		return x.IsPrivate
	}
	return ""
}

func (x *StorageRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

// 上传凭证请求参数
type UploadTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Storage      string `protobuf:"bytes,1,opt,name=storage,proto3" json:"storage"`                               //存储编码
	StoreId      int64  `protobuf:"varint,2,opt,name=store_id,json=storeId,proto3" json:"store_id"`               //店铺ID
	UserType     string `protobuf:"bytes,3,opt,name=user_type,json=userType,proto3" json:"user_type"`             //用户类型
	UserId       int64  `protobuf:"varint,4,opt,name=user_id,json=userId,proto3" json:"user_id"`                  //用户ID
	OverwriteKey string `protobuf:"bytes,5,opt,name=overwrite_key,json=overwriteKey,proto3" json:"overwrite_key"` //需要覆盖的文件名
	KeyPrefix    string `protobuf:"bytes,6,opt,name=key_prefix,json=keyPrefix,proto3" json:"key_prefix"`          //为指定的空间和对象前缀生成上传策略
	CallbackUrl  string `protobuf:"bytes,7,opt,name=callback_url,json=callbackUrl,proto3" json:"callback_url"`    //回调URL
	FileType     string `protobuf:"bytes,8,opt,name=file_type,json=fileType,proto3" json:"file_type"`             //文件类型：image/video/audio/doc/other
}

func (x *UploadTokenRequest) Reset() {
	*x = UploadTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storageService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadTokenRequest) ProtoMessage() {}

func (x *UploadTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_storageService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadTokenRequest.ProtoReflect.Descriptor instead.
func (*UploadTokenRequest) Descriptor() ([]byte, []int) {
	return file_storageService_proto_rawDescGZIP(), []int{2}
}

func (x *UploadTokenRequest) GetStorage() string {
	if x != nil {
		return x.Storage
	}
	return ""
}

func (x *UploadTokenRequest) GetStoreId() int64 {
	if x != nil {
		return x.StoreId
	}
	return 0
}

func (x *UploadTokenRequest) GetUserType() string {
	if x != nil {
		return x.UserType
	}
	return ""
}

func (x *UploadTokenRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UploadTokenRequest) GetOverwriteKey() string {
	if x != nil {
		return x.OverwriteKey
	}
	return ""
}

func (x *UploadTokenRequest) GetKeyPrefix() string {
	if x != nil {
		return x.KeyPrefix
	}
	return ""
}

func (x *UploadTokenRequest) GetCallbackUrl() string {
	if x != nil {
		return x.CallbackUrl
	}
	return ""
}

func (x *UploadTokenRequest) GetFileType() string {
	if x != nil {
		return x.FileType
	}
	return ""
}

// 删除请求参数
type RemoveKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Storage string `protobuf:"bytes,1,opt,name=storage,proto3" json:"storage"` //存储编码
	Key     string `protobuf:"bytes,2,opt,name=key,proto3" json:"key"`         //文件KEY
}

func (x *RemoveKeyRequest) Reset() {
	*x = RemoveKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storageService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveKeyRequest) ProtoMessage() {}

func (x *RemoveKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_storageService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveKeyRequest.ProtoReflect.Descriptor instead.
func (*RemoveKeyRequest) Descriptor() ([]byte, []int) {
	return file_storageService_proto_rawDescGZIP(), []int{3}
}

func (x *RemoveKeyRequest) GetStorage() string {
	if x != nil {
		return x.Storage
	}
	return ""
}

func (x *RemoveKeyRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

// 上传凭证响应数据
type UploadTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token       string `protobuf:"bytes,1,opt,name=token,proto3" json:"token"`
	ExpiredTime int64  `protobuf:"varint,2,opt,name=expired_time,json=expiredTime,proto3" json:"expired_time"` //过期时间
}

func (x *UploadTokenResponse) Reset() {
	*x = UploadTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storageService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadTokenResponse) ProtoMessage() {}

func (x *UploadTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_storageService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadTokenResponse.ProtoReflect.Descriptor instead.
func (*UploadTokenResponse) Descriptor() ([]byte, []int) {
	return file_storageService_proto_rawDescGZIP(), []int{4}
}

func (x *UploadTokenResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *UploadTokenResponse) GetExpiredTime() int64 {
	if x != nil {
		return x.ExpiredTime
	}
	return 0
}

// 存储响应数据
type StorageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Storage      `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*Storage    `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Msg    string        `protobuf:"bytes,4,opt,name=msg,proto3" json:"msg"`
}

func (x *StorageResponse) Reset() {
	*x = StorageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storageService_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageResponse) ProtoMessage() {}

func (x *StorageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_storageService_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageResponse.ProtoReflect.Descriptor instead.
func (*StorageResponse) Descriptor() ([]byte, []int) {
	return file_storageService_proto_rawDescGZIP(), []int{5}
}

func (x *StorageResponse) GetEntity() *Storage {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *StorageResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *StorageResponse) GetItems() []*Storage {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *StorageResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_storageService_proto protoreflect.FileDescriptor

var file_storageService_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xd8, 0x03, 0x0a, 0x07, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x7a, 0x6f, 0x6e,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x6f, 0x73,
	0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x68, 0x6f, 0x73,
	0x74, 0x55, 0x72, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x73, 0x44, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x73, 0x50, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x11,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x12, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xf0,
	0x02, 0x0a, 0x0e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x74, 0x6f, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72,
	0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72,
	0x64, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x72,
	0x61, 0x6e, 0x67, 0x65, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x07, 0x20, 0x03,
	0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69,
	0x73, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x69, 0x73, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73,
	0x5f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x69, 0x73, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x83, 0x02, 0x0a, 0x12, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x76, 0x65, 0x72, 0x77, 0x72, 0x69, 0x74, 0x65,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x76, 0x65, 0x72,
	0x77, 0x72, 0x69, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x6b, 0x65, 0x79, 0x5f,
	0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6b, 0x65,
	0x79, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x61, 0x6c, 0x6c, 0x62,
	0x61, 0x63, 0x6b, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63,
	0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x55, 0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x3e, 0x0a, 0x10, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x4e, 0x0a, 0x13, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x9c, 0x01, 0x0a, 0x0f, 0x53, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x06, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x06,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50,
	0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x32, 0x8c, 0x04, 0x0a, 0x0e, 0x53, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x11, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a,
	0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x11,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f,
	0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x3d, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x44, 0x0a, 0x09, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x1a, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4b, 0x65,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_storageService_proto_rawDescOnce sync.Once
	file_storageService_proto_rawDescData = file_storageService_proto_rawDesc
)

func file_storageService_proto_rawDescGZIP() []byte {
	file_storageService_proto_rawDescOnce.Do(func() {
		file_storageService_proto_rawDescData = protoimpl.X.CompressGZIP(file_storageService_proto_rawDescData)
	})
	return file_storageService_proto_rawDescData
}

var file_storageService_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_storageService_proto_goTypes = []interface{}{
	(*Storage)(nil),             // 0: services.Storage
	(*StorageRequest)(nil),      // 1: services.StorageRequest
	(*UploadTokenRequest)(nil),  // 2: services.UploadTokenRequest
	(*RemoveKeyRequest)(nil),    // 3: services.RemoveKeyRequest
	(*UploadTokenResponse)(nil), // 4: services.UploadTokenResponse
	(*StorageResponse)(nil),     // 5: services.StorageResponse
	(*common.Pager)(nil),        // 6: common.Pager
}
var file_storageService_proto_depIdxs = []int32{
	0,  // 0: services.StorageResponse.entity:type_name -> services.Storage
	6,  // 1: services.StorageResponse.pager:type_name -> common.Pager
	0,  // 2: services.StorageResponse.items:type_name -> services.Storage
	0,  // 3: services.StorageService.Create:input_type -> services.Storage
	0,  // 4: services.StorageService.Update:input_type -> services.Storage
	0,  // 5: services.StorageService.Delete:input_type -> services.Storage
	0,  // 6: services.StorageService.Get:input_type -> services.Storage
	1,  // 7: services.StorageService.Search:input_type -> services.StorageRequest
	1,  // 8: services.StorageService.List:input_type -> services.StorageRequest
	2,  // 9: services.StorageService.GetUploadToken:input_type -> services.UploadTokenRequest
	3,  // 10: services.StorageService.RemoveKey:input_type -> services.RemoveKeyRequest
	5,  // 11: services.StorageService.Create:output_type -> services.StorageResponse
	5,  // 12: services.StorageService.Update:output_type -> services.StorageResponse
	5,  // 13: services.StorageService.Delete:output_type -> services.StorageResponse
	5,  // 14: services.StorageService.Get:output_type -> services.StorageResponse
	5,  // 15: services.StorageService.Search:output_type -> services.StorageResponse
	5,  // 16: services.StorageService.List:output_type -> services.StorageResponse
	4,  // 17: services.StorageService.GetUploadToken:output_type -> services.UploadTokenResponse
	5,  // 18: services.StorageService.RemoveKey:output_type -> services.StorageResponse
	11, // [11:19] is the sub-list for method output_type
	3,  // [3:11] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_storageService_proto_init() }
func file_storageService_proto_init() {
	if File_storageService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_storageService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Storage); i {
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
		file_storageService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageRequest); i {
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
		file_storageService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadTokenRequest); i {
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
		file_storageService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveKeyRequest); i {
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
		file_storageService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadTokenResponse); i {
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
		file_storageService_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageResponse); i {
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
			RawDescriptor: file_storageService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_storageService_proto_goTypes,
		DependencyIndexes: file_storageService_proto_depIdxs,
		MessageInfos:      file_storageService_proto_msgTypes,
	}.Build()
	File_storageService_proto = out.File
	file_storageService_proto_rawDesc = nil
	file_storageService_proto_goTypes = nil
	file_storageService_proto_depIdxs = nil
}
