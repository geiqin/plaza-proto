// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: pluginService.proto

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

type Plugin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`                                  //ID
	Code       string `protobuf:"bytes,2,opt,name=code,proto3" json:"code"`                               //唯一标记
	Name       string `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`                               //插件名称
	Type       string `protobuf:"bytes,4,opt,name=type,proto3" json:"type"`                               //插件类型
	IconUrl    string `protobuf:"bytes,5,opt,name=icon_url,json=iconUrl,proto3" json:"icon_url"`          //插件图标
	Desc       string `protobuf:"bytes,6,opt,name=desc,proto3" json:"desc"`                               //插件描述
	Config     string `protobuf:"bytes,7,opt,name=config,proto3" json:"config"`                           //配置
	AuthorId   string `protobuf:"bytes,8,opt,name=author_id,json=authorId,proto3" json:"author_id"`       //作者ID
	AuthorName string `protobuf:"bytes,9,opt,name=author_name,json=authorName,proto3" json:"author_name"` //作者名称
	Version    string `protobuf:"bytes,10,opt,name=version,proto3" json:"version"`                        //版本
	Url        string `protobuf:"bytes,11,opt,name=url,proto3" json:"url"`                                //URL地址
	Weight     int32  `protobuf:"varint,12,opt,name=weight,proto3" json:"weight"`                         //权重
	IsEnable   string `protobuf:"bytes,13,opt,name=is_enable,json=isEnable,proto3" json:"is_enable"`      //是否启用（0否，1是）
	CreatedAt  string `protobuf:"bytes,14,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt  string `protobuf:"bytes,15,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
}

func (x *Plugin) Reset() {
	*x = Plugin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pluginService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Plugin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Plugin) ProtoMessage() {}

func (x *Plugin) ProtoReflect() protoreflect.Message {
	mi := &file_pluginService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Plugin.ProtoReflect.Descriptor instead.
func (*Plugin) Descriptor() ([]byte, []int) {
	return file_pluginService_proto_rawDescGZIP(), []int{0}
}

func (x *Plugin) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Plugin) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Plugin) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Plugin) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Plugin) GetIconUrl() string {
	if x != nil {
		return x.IconUrl
	}
	return ""
}

func (x *Plugin) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *Plugin) GetConfig() string {
	if x != nil {
		return x.Config
	}
	return ""
}

func (x *Plugin) GetAuthorId() string {
	if x != nil {
		return x.AuthorId
	}
	return ""
}

func (x *Plugin) GetAuthorName() string {
	if x != nil {
		return x.AuthorName
	}
	return ""
}

func (x *Plugin) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Plugin) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Plugin) GetWeight() int32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *Plugin) GetIsEnable() string {
	if x != nil {
		return x.IsEnable
	}
	return ""
}

func (x *Plugin) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Plugin) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type PluginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int64  `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize int64  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Sorting  string `protobuf:"bytes,3,opt,name=sorting,proto3" json:"sorting"`
	//base params
	Id       int32   `protobuf:"varint,4,opt,name=id,proto3" json:"id"`
	Name     string  `protobuf:"bytes,5,opt,name=name,proto3" json:"name"`
	Code     string  `protobuf:"bytes,6,opt,name=code,proto3" json:"code"`
	Type     string  `protobuf:"bytes,7,opt,name=type,proto3" json:"type"`
	IsEnable string  `protobuf:"bytes,8,opt,name=is_enable,json=isEnable,proto3" json:"is_enable"`
	AuthorId int64   `protobuf:"varint,9,opt,name=author_id,json=authorId,proto3" json:"author_id"`
	Keywords string  `protobuf:"bytes,10,opt,name=keywords,proto3" json:"keywords"`
	Ids      []int32 `protobuf:"varint,11,rep,packed,name=ids,proto3" json:"ids"`
}

func (x *PluginRequest) Reset() {
	*x = PluginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pluginService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginRequest) ProtoMessage() {}

func (x *PluginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pluginService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginRequest.ProtoReflect.Descriptor instead.
func (*PluginRequest) Descriptor() ([]byte, []int) {
	return file_pluginService_proto_rawDescGZIP(), []int{1}
}

func (x *PluginRequest) GetPaged() int64 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *PluginRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *PluginRequest) GetSorting() string {
	if x != nil {
		return x.Sorting
	}
	return ""
}

func (x *PluginRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PluginRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PluginRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *PluginRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *PluginRequest) GetIsEnable() string {
	if x != nil {
		return x.IsEnable
	}
	return ""
}

func (x *PluginRequest) GetAuthorId() int64 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

func (x *PluginRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *PluginRequest) GetIds() []int32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

//
type PluginData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Plugin       `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*Plugin     `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Info   *common.Info  `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
}

func (x *PluginData) Reset() {
	*x = PluginData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pluginService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginData) ProtoMessage() {}

func (x *PluginData) ProtoReflect() protoreflect.Message {
	mi := &file_pluginService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginData.ProtoReflect.Descriptor instead.
func (*PluginData) Descriptor() ([]byte, []int) {
	return file_pluginService_proto_rawDescGZIP(), []int{2}
}

func (x *PluginData) GetEntity() *Plugin {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *PluginData) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *PluginData) GetItems() []*Plugin {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *PluginData) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

//
type PluginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *PluginData   `protobuf:"bytes,1,opt,name=data,proto3" json:"data"`
	Error *common.Error `protobuf:"bytes,2,opt,name=error,proto3" json:"error"`
}

func (x *PluginResponse) Reset() {
	*x = PluginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pluginService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginResponse) ProtoMessage() {}

func (x *PluginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pluginService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginResponse.ProtoReflect.Descriptor instead.
func (*PluginResponse) Descriptor() ([]byte, []int) {
	return file_pluginService_proto_rawDescGZIP(), []int{3}
}

func (x *PluginResponse) GetData() *PluginData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *PluginResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_pluginService_proto protoreflect.FileDescriptor

var file_pluginService_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a,
	0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xf8, 0x02, 0x0a, 0x06, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x63, 0x6f,
	0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x63, 0x6f,
	0x6e, 0x55, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x90, 0x02,
	0x0a, 0x0d, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x73, 0x45,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x10,
	0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x05, 0x52, 0x03, 0x69, 0x64, 0x73,
	0x22, 0xa5, 0x01, 0x0a, 0x0a, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x28, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x26,
	0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x5f, 0x0a, 0x0e, 0x50, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xe8, 0x02, 0x0a, 0x0d, 0x50, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x06, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x10, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x1a,
	0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x06, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x10, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x1a, 0x18, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12,
	0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pluginService_proto_rawDescOnce sync.Once
	file_pluginService_proto_rawDescData = file_pluginService_proto_rawDesc
)

func file_pluginService_proto_rawDescGZIP() []byte {
	file_pluginService_proto_rawDescOnce.Do(func() {
		file_pluginService_proto_rawDescData = protoimpl.X.CompressGZIP(file_pluginService_proto_rawDescData)
	})
	return file_pluginService_proto_rawDescData
}

var file_pluginService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pluginService_proto_goTypes = []interface{}{
	(*Plugin)(nil),         // 0: services.Plugin
	(*PluginRequest)(nil),  // 1: services.PluginRequest
	(*PluginData)(nil),     // 2: services.PluginData
	(*PluginResponse)(nil), // 3: services.PluginResponse
	(*common.Pager)(nil),   // 4: common.Pager
	(*common.Info)(nil),    // 5: common.Info
	(*common.Error)(nil),   // 6: common.Error
}
var file_pluginService_proto_depIdxs = []int32{
	0,  // 0: services.PluginData.entity:type_name -> services.Plugin
	4,  // 1: services.PluginData.pager:type_name -> common.Pager
	0,  // 2: services.PluginData.items:type_name -> services.Plugin
	5,  // 3: services.PluginData.info:type_name -> common.Info
	2,  // 4: services.PluginResponse.data:type_name -> services.PluginData
	6,  // 5: services.PluginResponse.error:type_name -> common.Error
	0,  // 6: services.PluginService.Create:input_type -> services.Plugin
	0,  // 7: services.PluginService.Update:input_type -> services.Plugin
	0,  // 8: services.PluginService.Delete:input_type -> services.Plugin
	0,  // 9: services.PluginService.Get:input_type -> services.Plugin
	1,  // 10: services.PluginService.List:input_type -> services.PluginRequest
	1,  // 11: services.PluginService.Search:input_type -> services.PluginRequest
	3,  // 12: services.PluginService.Create:output_type -> services.PluginResponse
	3,  // 13: services.PluginService.Update:output_type -> services.PluginResponse
	3,  // 14: services.PluginService.Delete:output_type -> services.PluginResponse
	3,  // 15: services.PluginService.Get:output_type -> services.PluginResponse
	3,  // 16: services.PluginService.List:output_type -> services.PluginResponse
	3,  // 17: services.PluginService.Search:output_type -> services.PluginResponse
	12, // [12:18] is the sub-list for method output_type
	6,  // [6:12] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_pluginService_proto_init() }
func file_pluginService_proto_init() {
	if File_pluginService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pluginService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Plugin); i {
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
		file_pluginService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginRequest); i {
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
		file_pluginService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginData); i {
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
		file_pluginService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginResponse); i {
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
			RawDescriptor: file_pluginService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pluginService_proto_goTypes,
		DependencyIndexes: file_pluginService_proto_depIdxs,
		MessageInfos:      file_pluginService_proto_msgTypes,
	}.Build()
	File_pluginService_proto = out.File
	file_pluginService_proto_rawDesc = nil
	file_pluginService_proto_goTypes = nil
	file_pluginService_proto_depIdxs = nil
}
