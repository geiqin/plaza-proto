// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: filesystemService.proto

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

//文件系统
type Filesystem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`                                //id
	Code      string `protobuf:"bytes,2,opt,name=code,proto3" json:"code"`                             //唯一标识
	Name      string `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`                             //名称
	Driver    string `protobuf:"bytes,4,opt,name=driver,proto3" json:"driver"`                         //驱动器
	AccessKey string `protobuf:"bytes,5,opt,name=access_key,json=accessKey,proto3" json:"access_key"`  //云端AccessKey
	SecretKey string `protobuf:"bytes,6,opt,name=secret_key,json=secretKey,proto3" json:"secret_key"`  //云端SecretKey
	Bucket    string `protobuf:"bytes,7,opt,name=bucket,proto3" json:"bucket"`                         //云端Bucket
	Transport string `protobuf:"bytes,8,opt,name=transport,proto3" json:"transport"`                   //云端Transport
	Domain    string `protobuf:"bytes,9,opt,name=domain,proto3" json:"domain"`                         //访问域名
	IsPrivate string `protobuf:"bytes,10,opt,name=is_private,json=isPrivate,proto3" json:"is_private"` //是否为私有资源（0否，1是）
	IsLocal   string `protobuf:"bytes,11,opt,name=is_local,json=isLocal,proto3" json:"is_local"`       //是否为本地储存（0否，1是）
	LocalPath string `protobuf:"bytes,12,opt,name=local_path,json=localPath,proto3" json:"local_path"` //本地根路径
	Status    string `protobuf:"bytes,13,opt,name=status,proto3" json:"status"`                        //状态: 1-正常,0-停用
	CreatedAt string `protobuf:"bytes,14,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt string `protobuf:"bytes,15,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
}

func (x *Filesystem) Reset() {
	*x = Filesystem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filesystemService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filesystem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filesystem) ProtoMessage() {}

func (x *Filesystem) ProtoReflect() protoreflect.Message {
	mi := &file_filesystemService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Filesystem.ProtoReflect.Descriptor instead.
func (*Filesystem) Descriptor() ([]byte, []int) {
	return file_filesystemService_proto_rawDescGZIP(), []int{0}
}

func (x *Filesystem) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Filesystem) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Filesystem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Filesystem) GetDriver() string {
	if x != nil {
		return x.Driver
	}
	return ""
}

func (x *Filesystem) GetAccessKey() string {
	if x != nil {
		return x.AccessKey
	}
	return ""
}

func (x *Filesystem) GetSecretKey() string {
	if x != nil {
		return x.SecretKey
	}
	return ""
}

func (x *Filesystem) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *Filesystem) GetTransport() string {
	if x != nil {
		return x.Transport
	}
	return ""
}

func (x *Filesystem) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *Filesystem) GetIsPrivate() string {
	if x != nil {
		return x.IsPrivate
	}
	return ""
}

func (x *Filesystem) GetIsLocal() string {
	if x != nil {
		return x.IsLocal
	}
	return ""
}

func (x *Filesystem) GetLocalPath() string {
	if x != nil {
		return x.LocalPath
	}
	return ""
}

func (x *Filesystem) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Filesystem) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Filesystem) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type FilesystemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged     int32   `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize  int32   `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Id        int32   `protobuf:"varint,3,opt,name=id,proto3" json:"id"`
	Keywords  string  `protobuf:"bytes,4,opt,name=keywords,proto3" json:"keywords"`
	Code      string  `protobuf:"bytes,5,opt,name=code,proto3" json:"code"`
	Name      string  `protobuf:"bytes,6,opt,name=name,proto3" json:"name"`
	Driver    string  `protobuf:"bytes,7,opt,name=driver,proto3" json:"driver"`                        //驱动器
	IsPrivate string  `protobuf:"bytes,8,opt,name=is_private,json=isPrivate,proto3" json:"is_private"` //是否为私有资源
	IsLocal   string  `protobuf:"bytes,9,opt,name=is_local,json=isLocal,proto3" json:"is_local"`       //是否为本地储存
	Status    string  `protobuf:"bytes,10,opt,name=status,proto3" json:"status"`                       //状态: 1-正常,0-停用
	Ids       []int32 `protobuf:"varint,11,rep,packed,name=ids,proto3" json:"ids"`
}

func (x *FilesystemRequest) Reset() {
	*x = FilesystemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filesystemService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilesystemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilesystemRequest) ProtoMessage() {}

func (x *FilesystemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_filesystemService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilesystemRequest.ProtoReflect.Descriptor instead.
func (*FilesystemRequest) Descriptor() ([]byte, []int) {
	return file_filesystemService_proto_rawDescGZIP(), []int{1}
}

func (x *FilesystemRequest) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *FilesystemRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *FilesystemRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FilesystemRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *FilesystemRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *FilesystemRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FilesystemRequest) GetDriver() string {
	if x != nil {
		return x.Driver
	}
	return ""
}

func (x *FilesystemRequest) GetIsPrivate() string {
	if x != nil {
		return x.IsPrivate
	}
	return ""
}

func (x *FilesystemRequest) GetIsLocal() string {
	if x != nil {
		return x.IsLocal
	}
	return ""
}

func (x *FilesystemRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *FilesystemRequest) GetIds() []int32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type FilesystemData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Filesystem   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*Filesystem `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Info   *common.Info  `protobuf:"bytes,4,opt,name=info,proto3" json:"info"`
}

func (x *FilesystemData) Reset() {
	*x = FilesystemData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filesystemService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilesystemData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilesystemData) ProtoMessage() {}

func (x *FilesystemData) ProtoReflect() protoreflect.Message {
	mi := &file_filesystemService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilesystemData.ProtoReflect.Descriptor instead.
func (*FilesystemData) Descriptor() ([]byte, []int) {
	return file_filesystemService_proto_rawDescGZIP(), []int{2}
}

func (x *FilesystemData) GetEntity() *Filesystem {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *FilesystemData) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *FilesystemData) GetItems() []*Filesystem {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *FilesystemData) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type FilesystemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *FilesystemData `protobuf:"bytes,1,opt,name=data,proto3" json:"data"`
	Error *common.Error   `protobuf:"bytes,2,opt,name=error,proto3" json:"error"`
}

func (x *FilesystemResponse) Reset() {
	*x = FilesystemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filesystemService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilesystemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilesystemResponse) ProtoMessage() {}

func (x *FilesystemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_filesystemService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilesystemResponse.ProtoReflect.Descriptor instead.
func (*FilesystemResponse) Descriptor() ([]byte, []int) {
	return file_filesystemService_proto_rawDescGZIP(), []int{3}
}

func (x *FilesystemResponse) GetData() *FilesystemData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *FilesystemResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_filesystemService_proto protoreflect.FileDescriptor

var file_filesystemService_proto_rawDesc = []byte{
	0x0a, 0x17, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x03, 0x0a, 0x0a, 0x46, 0x69, 0x6c, 0x65, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x4b, 0x65, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x4b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x73, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x69, 0x73, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x50, 0x61, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x22, 0x96, 0x02, 0x0a, 0x11, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79,
	0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79,
	0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x73, 0x50, 0x72, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x73, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x0b,
	0x20, 0x03, 0x28, 0x05, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0xb1, 0x01, 0x0a, 0x0e, 0x46, 0x69,
	0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2c, 0x0a, 0x06,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61,
	0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12,
	0x2a, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x20, 0x0a, 0x04, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x67, 0x0a,
	0x12, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x69, 0x6c,
	0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x9c, 0x03, 0x0a, 0x11, 0x46, 0x69, 0x6c, 0x65, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x06,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x1a, 0x1c, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x06,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x1a, 0x1c, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x06,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x1a, 0x1c, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x03,
	0x47, 0x65, 0x74, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46,
	0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x06, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x12, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46,
	0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x69, 0x6c, 0x65,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x43, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_filesystemService_proto_rawDescOnce sync.Once
	file_filesystemService_proto_rawDescData = file_filesystemService_proto_rawDesc
)

func file_filesystemService_proto_rawDescGZIP() []byte {
	file_filesystemService_proto_rawDescOnce.Do(func() {
		file_filesystemService_proto_rawDescData = protoimpl.X.CompressGZIP(file_filesystemService_proto_rawDescData)
	})
	return file_filesystemService_proto_rawDescData
}

var file_filesystemService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_filesystemService_proto_goTypes = []interface{}{
	(*Filesystem)(nil),         // 0: services.Filesystem
	(*FilesystemRequest)(nil),  // 1: services.FilesystemRequest
	(*FilesystemData)(nil),     // 2: services.FilesystemData
	(*FilesystemResponse)(nil), // 3: services.FilesystemResponse
	(*common.Pager)(nil),       // 4: common.Pager
	(*common.Info)(nil),        // 5: common.Info
	(*common.Error)(nil),       // 6: common.Error
}
var file_filesystemService_proto_depIdxs = []int32{
	0,  // 0: services.FilesystemData.entity:type_name -> services.Filesystem
	4,  // 1: services.FilesystemData.pager:type_name -> common.Pager
	0,  // 2: services.FilesystemData.items:type_name -> services.Filesystem
	5,  // 3: services.FilesystemData.info:type_name -> common.Info
	2,  // 4: services.FilesystemResponse.data:type_name -> services.FilesystemData
	6,  // 5: services.FilesystemResponse.error:type_name -> common.Error
	0,  // 6: services.FilesystemService.Create:input_type -> services.Filesystem
	0,  // 7: services.FilesystemService.Update:input_type -> services.Filesystem
	0,  // 8: services.FilesystemService.Delete:input_type -> services.Filesystem
	0,  // 9: services.FilesystemService.Get:input_type -> services.Filesystem
	1,  // 10: services.FilesystemService.Search:input_type -> services.FilesystemRequest
	1,  // 11: services.FilesystemService.List:input_type -> services.FilesystemRequest
	3,  // 12: services.FilesystemService.Create:output_type -> services.FilesystemResponse
	3,  // 13: services.FilesystemService.Update:output_type -> services.FilesystemResponse
	3,  // 14: services.FilesystemService.Delete:output_type -> services.FilesystemResponse
	3,  // 15: services.FilesystemService.Get:output_type -> services.FilesystemResponse
	3,  // 16: services.FilesystemService.Search:output_type -> services.FilesystemResponse
	3,  // 17: services.FilesystemService.List:output_type -> services.FilesystemResponse
	12, // [12:18] is the sub-list for method output_type
	6,  // [6:12] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_filesystemService_proto_init() }
func file_filesystemService_proto_init() {
	if File_filesystemService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_filesystemService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Filesystem); i {
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
		file_filesystemService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilesystemRequest); i {
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
		file_filesystemService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilesystemData); i {
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
		file_filesystemService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilesystemResponse); i {
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
			RawDescriptor: file_filesystemService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_filesystemService_proto_goTypes,
		DependencyIndexes: file_filesystemService_proto_depIdxs,
		MessageInfos:      file_filesystemService_proto_msgTypes,
	}.Build()
	File_filesystemService_proto = out.File
	file_filesystemService_proto_rawDesc = nil
	file_filesystemService_proto_goTypes = nil
	file_filesystemService_proto_depIdxs = nil
}
