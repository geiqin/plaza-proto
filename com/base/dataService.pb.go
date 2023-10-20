// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: dataService.proto

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

// 区域信息
type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`                      //ID
	Type      string `protobuf:"bytes,2,opt,name=type,proto3" json:"type"`                   //类型
	Code      string `protobuf:"bytes,3,opt,name=code,proto3" json:"code"`                   //唯一编码
	Name      string `protobuf:"bytes,4,opt,name=name,proto3" json:"name"`                   //数据名称
	Content   string `protobuf:"bytes,5,opt,name=content,proto3" json:"content"`             //JSON内容
	IsInit    string `protobuf:"bytes,6,opt,name=is_init,json=isInit,proto3" json:"is_init"` //是否初始化
	Status    string `protobuf:"bytes,7,opt,name=status,proto3" json:"status"`               //状态：0禁用，1可用
	CreatedAt string `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt string `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dataService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_dataService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_dataService_proto_rawDescGZIP(), []int{0}
}

func (x *Data) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Data) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Data) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Data) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Data) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Data) GetIsInit() string {
	if x != nil {
		return x.IsInit
	}
	return ""
}

func (x *Data) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Data) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Data) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

//查询参数
type DataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int64  `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize int64  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Keywords string `protobuf:"bytes,3,opt,name=keywords,proto3" json:"keywords"`
	//以下为自定义参数
	Id       int64   `protobuf:"varint,4,opt,name=id,proto3" json:"id"`
	Name     string  `protobuf:"bytes,5,opt,name=name,proto3" json:"name"`
	Type     string  `protobuf:"bytes,6,opt,name=type,proto3" json:"type"`
	Status   string  `protobuf:"bytes,7,opt,name=status,proto3" json:"status"`
	Ids      []int64 `protobuf:"varint,8,rep,packed,name=ids,proto3" json:"ids"`
	ParentId int64   `protobuf:"varint,9,opt,name=parent_id,json=parentId,proto3" json:"parent_id"`
}

func (x *DataRequest) Reset() {
	*x = DataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dataService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataRequest) ProtoMessage() {}

func (x *DataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dataService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataRequest.ProtoReflect.Descriptor instead.
func (*DataRequest) Descriptor() ([]byte, []int) {
	return file_dataService_proto_rawDescGZIP(), []int{1}
}

func (x *DataRequest) GetPaged() int64 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *DataRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *DataRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *DataRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DataRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DataRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *DataRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *DataRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *DataRequest) GetParentId() int64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

type DataData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Data         `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*Data       `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Msg    string        `protobuf:"bytes,4,opt,name=msg,proto3" json:"msg"`
}

func (x *DataData) Reset() {
	*x = DataData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dataService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataData) ProtoMessage() {}

func (x *DataData) ProtoReflect() protoreflect.Message {
	mi := &file_dataService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataData.ProtoReflect.Descriptor instead.
func (*DataData) Descriptor() ([]byte, []int) {
	return file_dataService_proto_rawDescGZIP(), []int{2}
}

func (x *DataData) GetEntity() *Data {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *DataData) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *DataData) GetItems() []*Data {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *DataData) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type DataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *DataData     `protobuf:"bytes,1,opt,name=data,proto3" json:"data"`
	Error *common.Error `protobuf:"bytes,2,opt,name=error,proto3" json:"error"`
}

func (x *DataResponse) Reset() {
	*x = DataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dataService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataResponse) ProtoMessage() {}

func (x *DataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dataService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataResponse.ProtoReflect.Descriptor instead.
func (*DataResponse) Descriptor() ([]byte, []int) {
	return file_dataService_proto_rawDescGZIP(), []int{3}
}

func (x *DataResponse) GetData() *DataData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *DataResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_dataService_proto protoreflect.FileDescriptor

var file_dataService_proto_rawDesc = []byte{
	0x0a, 0x11, 0x64, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xdb, 0x01, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x69, 0x6e, 0x69, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x69, 0x73, 0x49, 0x6e, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xdb,
	0x01, 0x0a, 0x0b, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70,
	0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a,
	0x03, 0x69, 0x64, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12,
	0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x8f, 0x01, 0x0a,
	0x08, 0x44, 0x61, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x12, 0x26, 0x0a, 0x06, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52,
	0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x10, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x5b,
	0x0a, 0x0c, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xce, 0x02, 0x0a, 0x0b,
	0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x32, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0e, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x16, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0e,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x16,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x15, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b,
	0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_dataService_proto_rawDescOnce sync.Once
	file_dataService_proto_rawDescData = file_dataService_proto_rawDesc
)

func file_dataService_proto_rawDescGZIP() []byte {
	file_dataService_proto_rawDescOnce.Do(func() {
		file_dataService_proto_rawDescData = protoimpl.X.CompressGZIP(file_dataService_proto_rawDescData)
	})
	return file_dataService_proto_rawDescData
}

var file_dataService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_dataService_proto_goTypes = []interface{}{
	(*Data)(nil),         // 0: services.Data
	(*DataRequest)(nil),  // 1: services.DataRequest
	(*DataData)(nil),     // 2: services.DataData
	(*DataResponse)(nil), // 3: services.DataResponse
	(*common.Pager)(nil), // 4: common.Pager
	(*common.Error)(nil), // 5: common.Error
}
var file_dataService_proto_depIdxs = []int32{
	0,  // 0: services.DataData.entity:type_name -> services.Data
	4,  // 1: services.DataData.pager:type_name -> common.Pager
	0,  // 2: services.DataData.items:type_name -> services.Data
	2,  // 3: services.DataResponse.data:type_name -> services.DataData
	5,  // 4: services.DataResponse.error:type_name -> common.Error
	0,  // 5: services.DataService.Create:input_type -> services.Data
	0,  // 6: services.DataService.Update:input_type -> services.Data
	0,  // 7: services.DataService.Delete:input_type -> services.Data
	0,  // 8: services.DataService.Get:input_type -> services.Data
	1,  // 9: services.DataService.Search:input_type -> services.DataRequest
	1,  // 10: services.DataService.List:input_type -> services.DataRequest
	3,  // 11: services.DataService.Create:output_type -> services.DataResponse
	3,  // 12: services.DataService.Update:output_type -> services.DataResponse
	3,  // 13: services.DataService.Delete:output_type -> services.DataResponse
	3,  // 14: services.DataService.Get:output_type -> services.DataResponse
	3,  // 15: services.DataService.Search:output_type -> services.DataResponse
	3,  // 16: services.DataService.List:output_type -> services.DataResponse
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_dataService_proto_init() }
func file_dataService_proto_init() {
	if File_dataService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dataService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
		file_dataService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataRequest); i {
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
		file_dataService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataData); i {
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
		file_dataService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataResponse); i {
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
			RawDescriptor: file_dataService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dataService_proto_goTypes,
		DependencyIndexes: file_dataService_proto_depIdxs,
		MessageInfos:      file_dataService_proto_msgTypes,
	}.Build()
	File_dataService_proto = out.File
	file_dataService_proto_rawDesc = nil
	file_dataService_proto_goTypes = nil
	file_dataService_proto_depIdxs = nil
}
