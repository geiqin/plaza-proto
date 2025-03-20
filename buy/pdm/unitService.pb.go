// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: unitService.proto

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

// 单位库
type Unit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id"`                                           //ID
	Name         string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`                                        //单位名称
	Type         string  `protobuf:"bytes,3,opt,name=type,proto3" json:"type"`                                        //单位类型
	Symbol       string  `protobuf:"bytes,4,opt,name=symbol,proto3" json:"symbol"`                                    //单位符号
	ExchangeRate float32 `protobuf:"fixed32,5,opt,name=exchange_rate,json=exchangeRate,proto3" json:"exchange_rate"`  //换算率
	IsSystem     string  `protobuf:"bytes,6,opt,name=is_system,json=isSystem,proto3" json:"is_system"`                //创建来源
	Sort         int32   `protobuf:"varint,7,opt,name=sort,proto3" json:"sort"`                                       //排序
	Status       string  `protobuf:"bytes,8,opt,name=status,proto3" json:"status"`                                    //状态
	CreatedAt    int64   `protobuf:"varint,9,opt,name=created_at,json=createdAt,proto3" json:"created_at"`            //创建时间
	UpdatedAt    int64   `protobuf:"varint,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`           //修改时间
	TypeName     string  `protobuf:"bytes,11,opt,name=type_name,json=typeName,proto3" json:"type_name"`               //单位类型名
	IsSystemName string  `protobuf:"bytes,12,opt,name=is_system_name,json=isSystemName,proto3" json:"is_system_name"` //创建来源名
}

func (x *Unit) Reset() {
	*x = Unit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_unitService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unit) ProtoMessage() {}

func (x *Unit) ProtoReflect() protoreflect.Message {
	mi := &file_unitService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unit.ProtoReflect.Descriptor instead.
func (*Unit) Descriptor() ([]byte, []int) {
	return file_unitService_proto_rawDescGZIP(), []int{0}
}

func (x *Unit) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Unit) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Unit) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Unit) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *Unit) GetExchangeRate() float32 {
	if x != nil {
		return x.ExchangeRate
	}
	return 0
}

func (x *Unit) GetIsSystem() string {
	if x != nil {
		return x.IsSystem
	}
	return ""
}

func (x *Unit) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *Unit) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Unit) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Unit) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *Unit) GetTypeName() string {
	if x != nil {
		return x.TypeName
	}
	return ""
}

func (x *Unit) GetIsSystemName() string {
	if x != nil {
		return x.IsSystemName
	}
	return ""
}

// 单位库请求参数
type UnitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Top       int32    `protobuf:"varint,1,opt,name=top,proto3" json:"top"`
	Paged     int64    `protobuf:"varint,2,opt,name=paged,proto3" json:"paged"`
	PageSize  int64    `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Keywords  string   `protobuf:"bytes,4,opt,name=keywords,proto3" json:"keywords"`
	Sorts     []string `protobuf:"bytes,5,rep,name=sorts,proto3" json:"sorts"`
	DateRange []string `protobuf:"bytes,6,rep,name=date_range,json=dateRange,proto3" json:"date_range"`
	Ids       []int32  `protobuf:"varint,7,rep,packed,name=ids,proto3" json:"ids"`
	Id        int32    `protobuf:"varint,8,opt,name=id,proto3" json:"id"`
	//以下为自定义参数
	Name     string `protobuf:"bytes,11,opt,name=name,proto3" json:"name"`                         //单位名称
	Type     string `protobuf:"bytes,12,opt,name=type,proto3" json:"type"`                         //单位类型
	IsSystem string `protobuf:"bytes,13,opt,name=is_system,json=isSystem,proto3" json:"is_system"` //创建来源
	Status   string `protobuf:"bytes,14,opt,name=status,proto3" json:"status"`                     //状态
}

func (x *UnitRequest) Reset() {
	*x = UnitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_unitService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnitRequest) ProtoMessage() {}

func (x *UnitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_unitService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnitRequest.ProtoReflect.Descriptor instead.
func (*UnitRequest) Descriptor() ([]byte, []int) {
	return file_unitService_proto_rawDescGZIP(), []int{1}
}

func (x *UnitRequest) GetTop() int32 {
	if x != nil {
		return x.Top
	}
	return 0
}

func (x *UnitRequest) GetPaged() int64 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *UnitRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *UnitRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *UnitRequest) GetSorts() []string {
	if x != nil {
		return x.Sorts
	}
	return nil
}

func (x *UnitRequest) GetDateRange() []string {
	if x != nil {
		return x.DateRange
	}
	return nil
}

func (x *UnitRequest) GetIds() []int32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *UnitRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UnitRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UnitRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *UnitRequest) GetIsSystem() string {
	if x != nil {
		return x.IsSystem
	}
	return ""
}

func (x *UnitRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

// 单位库响应数据
type UnitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg    string        `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Entity *Unit         `protobuf:"bytes,3,opt,name=entity,proto3" json:"entity"`
	Items  []*Unit       `protobuf:"bytes,4,rep,name=items,proto3" json:"items"`
}

func (x *UnitResponse) Reset() {
	*x = UnitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_unitService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnitResponse) ProtoMessage() {}

func (x *UnitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_unitService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnitResponse.ProtoReflect.Descriptor instead.
func (*UnitResponse) Descriptor() ([]byte, []int) {
	return file_unitService_proto_rawDescGZIP(), []int{2}
}

func (x *UnitResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *UnitResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *UnitResponse) GetEntity() *Unit {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *UnitResponse) GetItems() []*Unit {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_unitService_proto protoreflect.FileDescriptor

var file_unitService_proto_rawDesc = []byte{
	0x0a, 0x11, 0x75, 0x6e, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xc5, 0x02, 0x0a, 0x04, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x78, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x0c, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x69, 0x73, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x69, 0x73, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x6f, 0x72, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x79, 0x70, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x69, 0x73, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x73, 0x53, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xa2, 0x02, 0x0a, 0x0b, 0x55, 0x6e, 0x69,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x6f, 0x70, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x74, 0x6f, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61,
	0x67, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x6f, 0x72,
	0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x73, 0x6f, 0x72, 0x74, 0x73, 0x12,
	0x1d, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x09, 0x64, 0x61, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x05, 0x52, 0x03, 0x69, 0x64, 0x73,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x73, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x93, 0x01,
	0x0a, 0x0c, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67,
	0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05,
	0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x24, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x32, 0x83, 0x03, 0x0a, 0x0b, 0x55, 0x6e, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x07, 0x50, 0x75, 0x6c, 0x6c, 0x4e, 0x65, 0x77, 0x12, 0x0e,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x1a, 0x16,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x55, 0x6e,
	0x69, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x55, 0x6e,
	0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x06,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x32, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0e, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0e, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12,
	0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x37, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_unitService_proto_rawDescOnce sync.Once
	file_unitService_proto_rawDescData = file_unitService_proto_rawDesc
)

func file_unitService_proto_rawDescGZIP() []byte {
	file_unitService_proto_rawDescOnce.Do(func() {
		file_unitService_proto_rawDescData = protoimpl.X.CompressGZIP(file_unitService_proto_rawDescData)
	})
	return file_unitService_proto_rawDescData
}

var file_unitService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_unitService_proto_goTypes = []interface{}{
	(*Unit)(nil),         // 0: services.Unit
	(*UnitRequest)(nil),  // 1: services.UnitRequest
	(*UnitResponse)(nil), // 2: services.UnitResponse
	(*common.Pager)(nil), // 3: common.Pager
}
var file_unitService_proto_depIdxs = []int32{
	3,  // 0: services.UnitResponse.pager:type_name -> common.Pager
	0,  // 1: services.UnitResponse.entity:type_name -> services.Unit
	0,  // 2: services.UnitResponse.items:type_name -> services.Unit
	0,  // 3: services.UnitService.PullNew:input_type -> services.Unit
	0,  // 4: services.UnitService.Create:input_type -> services.Unit
	0,  // 5: services.UnitService.Update:input_type -> services.Unit
	0,  // 6: services.UnitService.Delete:input_type -> services.Unit
	0,  // 7: services.UnitService.Get:input_type -> services.Unit
	1,  // 8: services.UnitService.Search:input_type -> services.UnitRequest
	1,  // 9: services.UnitService.List:input_type -> services.UnitRequest
	2,  // 10: services.UnitService.PullNew:output_type -> services.UnitResponse
	2,  // 11: services.UnitService.Create:output_type -> services.UnitResponse
	2,  // 12: services.UnitService.Update:output_type -> services.UnitResponse
	2,  // 13: services.UnitService.Delete:output_type -> services.UnitResponse
	2,  // 14: services.UnitService.Get:output_type -> services.UnitResponse
	2,  // 15: services.UnitService.Search:output_type -> services.UnitResponse
	2,  // 16: services.UnitService.List:output_type -> services.UnitResponse
	10, // [10:17] is the sub-list for method output_type
	3,  // [3:10] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_unitService_proto_init() }
func file_unitService_proto_init() {
	if File_unitService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_unitService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Unit); i {
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
		file_unitService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnitRequest); i {
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
		file_unitService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnitResponse); i {
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
			RawDescriptor: file_unitService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_unitService_proto_goTypes,
		DependencyIndexes: file_unitService_proto_depIdxs,
		MessageInfos:      file_unitService_proto_msgTypes,
	}.Build()
	File_unitService_proto = out.File
	file_unitService_proto_rawDesc = nil
	file_unitService_proto_goTypes = nil
	file_unitService_proto_depIdxs = nil
}
