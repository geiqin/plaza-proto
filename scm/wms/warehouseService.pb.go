// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: warehouseService.proto

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

type Warehouse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64     `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Code       string    `protobuf:"bytes,2,opt,name=code,proto3" json:"code"`
	Name       string    `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`
	AreaId     int64     `protobuf:"varint,4,opt,name=area_id,json=areaId,proto3" json:"area_id"`
	Address    string    `protobuf:"bytes,5,opt,name=address,proto3" json:"address"`
	ContactMan string    `protobuf:"bytes,6,opt,name=contact_man,json=contactMan,proto3" json:"contact_man"`
	ContactTel string    `protobuf:"bytes,7,opt,name=contact_tel,json=contactTel,proto3" json:"contact_tel"`
	Memo       string    `protobuf:"bytes,8,opt,name=memo,proto3" json:"memo"`
	Status     string    `protobuf:"bytes,9,opt,name=status,proto3" json:"status"`
	CreatedAt  string    `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt  string    `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	Area       *AreaInfo `protobuf:"bytes,12,opt,name=area,proto3" json:"area"`
}

func (x *Warehouse) Reset() {
	*x = Warehouse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_warehouseService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Warehouse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Warehouse) ProtoMessage() {}

func (x *Warehouse) ProtoReflect() protoreflect.Message {
	mi := &file_warehouseService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Warehouse.ProtoReflect.Descriptor instead.
func (*Warehouse) Descriptor() ([]byte, []int) {
	return file_warehouseService_proto_rawDescGZIP(), []int{0}
}

func (x *Warehouse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Warehouse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Warehouse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Warehouse) GetAreaId() int64 {
	if x != nil {
		return x.AreaId
	}
	return 0
}

func (x *Warehouse) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Warehouse) GetContactMan() string {
	if x != nil {
		return x.ContactMan
	}
	return ""
}

func (x *Warehouse) GetContactTel() string {
	if x != nil {
		return x.ContactTel
	}
	return ""
}

func (x *Warehouse) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *Warehouse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Warehouse) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Warehouse) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Warehouse) GetArea() *AreaInfo {
	if x != nil {
		return x.Area
	}
	return nil
}

type WarehouseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int32   `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize int32   `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Sorting  string  `protobuf:"bytes,3,opt,name=sorting,proto3" json:"sorting"`
	Id       int64   `protobuf:"varint,4,opt,name=id,proto3" json:"id"`
	Code     string  `protobuf:"bytes,5,opt,name=code,proto3" json:"code"`
	Name     string  `protobuf:"bytes,6,opt,name=name,proto3" json:"name"`
	Status   string  `protobuf:"bytes,7,opt,name=status,proto3" json:"status"`
	Ids      []int64 `protobuf:"varint,8,rep,packed,name=ids,proto3" json:"ids"`
}

func (x *WarehouseRequest) Reset() {
	*x = WarehouseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_warehouseService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WarehouseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WarehouseRequest) ProtoMessage() {}

func (x *WarehouseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_warehouseService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WarehouseRequest.ProtoReflect.Descriptor instead.
func (*WarehouseRequest) Descriptor() ([]byte, []int) {
	return file_warehouseService_proto_rawDescGZIP(), []int{1}
}

func (x *WarehouseRequest) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *WarehouseRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *WarehouseRequest) GetSorting() string {
	if x != nil {
		return x.Sorting
	}
	return ""
}

func (x *WarehouseRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *WarehouseRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *WarehouseRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *WarehouseRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *WarehouseRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type WarehouseData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Warehouse    `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*Warehouse  `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Info   *common.Info  `protobuf:"bytes,4,opt,name=info,proto3" json:"info"`
}

func (x *WarehouseData) Reset() {
	*x = WarehouseData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_warehouseService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WarehouseData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WarehouseData) ProtoMessage() {}

func (x *WarehouseData) ProtoReflect() protoreflect.Message {
	mi := &file_warehouseService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WarehouseData.ProtoReflect.Descriptor instead.
func (*WarehouseData) Descriptor() ([]byte, []int) {
	return file_warehouseService_proto_rawDescGZIP(), []int{2}
}

func (x *WarehouseData) GetEntity() *Warehouse {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *WarehouseData) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *WarehouseData) GetItems() []*Warehouse {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *WarehouseData) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type WarehouseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *WarehouseData `protobuf:"bytes,1,opt,name=data,proto3" json:"data"`
	Error *common.Error  `protobuf:"bytes,2,opt,name=error,proto3" json:"error"`
}

func (x *WarehouseResponse) Reset() {
	*x = WarehouseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_warehouseService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WarehouseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WarehouseResponse) ProtoMessage() {}

func (x *WarehouseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_warehouseService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WarehouseResponse.ProtoReflect.Descriptor instead.
func (*WarehouseResponse) Descriptor() ([]byte, []int) {
	return file_warehouseService_proto_rawDescGZIP(), []int{3}
}

func (x *WarehouseResponse) GetData() *WarehouseData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *WarehouseResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_warehouseService_proto protoreflect.FileDescriptor

var file_warehouseService_proto_rawDesc = []byte{
	0x0a, 0x16, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x62, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xca, 0x02, 0x0a,
	0x09, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x72, 0x65, 0x61, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x5f, 0x6d, 0x61, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x4d, 0x61, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x5f, 0x74, 0x65, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x63, 0x74, 0x54, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x26, 0x0a, 0x04, 0x61, 0x72, 0x65, 0x61, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x72, 0x65, 0x61, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x04, 0x61, 0x72, 0x65, 0x61, 0x22, 0xc1, 0x01, 0x0a, 0x10, 0x57, 0x61,
	0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70,
	0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x69,
	0x64, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0xae, 0x01,
	0x0a, 0x0d, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x2b, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x57, 0x61, 0x72, 0x65, 0x68,
	0x6f, 0x75, 0x73, 0x65, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05,
	0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65,
	0x72, 0x12, 0x29, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x57, 0x61, 0x72, 0x65,
	0x68, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x20, 0x0a, 0x04,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x65,
	0x0a, 0x11, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x57, 0x61, 0x72,
	0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x8f, 0x03, 0x0a, 0x10, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f,
	0x75, 0x73, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x06, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x57, 0x61,
	0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x57, 0x61, 0x72, 0x65,
	0x68, 0x6f, 0x75, 0x73, 0x65, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x13, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65,
	0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x57, 0x61, 0x72, 0x65,
	0x68, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x41, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x57,
	0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x43, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1a, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_warehouseService_proto_rawDescOnce sync.Once
	file_warehouseService_proto_rawDescData = file_warehouseService_proto_rawDesc
)

func file_warehouseService_proto_rawDescGZIP() []byte {
	file_warehouseService_proto_rawDescOnce.Do(func() {
		file_warehouseService_proto_rawDescData = protoimpl.X.CompressGZIP(file_warehouseService_proto_rawDescData)
	})
	return file_warehouseService_proto_rawDescData
}

var file_warehouseService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_warehouseService_proto_goTypes = []interface{}{
	(*Warehouse)(nil),         // 0: services.Warehouse
	(*WarehouseRequest)(nil),  // 1: services.WarehouseRequest
	(*WarehouseData)(nil),     // 2: services.WarehouseData
	(*WarehouseResponse)(nil), // 3: services.WarehouseResponse
	(*AreaInfo)(nil),          // 4: services.AreaInfo
	(*common.Pager)(nil),      // 5: common.Pager
	(*common.Info)(nil),       // 6: common.Info
	(*common.Error)(nil),      // 7: common.Error
}
var file_warehouseService_proto_depIdxs = []int32{
	4,  // 0: services.Warehouse.area:type_name -> services.AreaInfo
	0,  // 1: services.WarehouseData.entity:type_name -> services.Warehouse
	5,  // 2: services.WarehouseData.pager:type_name -> common.Pager
	0,  // 3: services.WarehouseData.items:type_name -> services.Warehouse
	6,  // 4: services.WarehouseData.info:type_name -> common.Info
	2,  // 5: services.WarehouseResponse.data:type_name -> services.WarehouseData
	7,  // 6: services.WarehouseResponse.error:type_name -> common.Error
	0,  // 7: services.WarehouseService.Create:input_type -> services.Warehouse
	0,  // 8: services.WarehouseService.Update:input_type -> services.Warehouse
	0,  // 9: services.WarehouseService.Delete:input_type -> services.Warehouse
	0,  // 10: services.WarehouseService.Get:input_type -> services.Warehouse
	1,  // 11: services.WarehouseService.List:input_type -> services.WarehouseRequest
	1,  // 12: services.WarehouseService.Search:input_type -> services.WarehouseRequest
	3,  // 13: services.WarehouseService.Create:output_type -> services.WarehouseResponse
	3,  // 14: services.WarehouseService.Update:output_type -> services.WarehouseResponse
	3,  // 15: services.WarehouseService.Delete:output_type -> services.WarehouseResponse
	3,  // 16: services.WarehouseService.Get:output_type -> services.WarehouseResponse
	3,  // 17: services.WarehouseService.List:output_type -> services.WarehouseResponse
	3,  // 18: services.WarehouseService.Search:output_type -> services.WarehouseResponse
	13, // [13:19] is the sub-list for method output_type
	7,  // [7:13] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_warehouseService_proto_init() }
func file_warehouseService_proto_init() {
	if File_warehouseService_proto != nil {
		return
	}
	file_baseInfoService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_warehouseService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Warehouse); i {
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
		file_warehouseService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WarehouseRequest); i {
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
		file_warehouseService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WarehouseData); i {
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
		file_warehouseService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WarehouseResponse); i {
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
			RawDescriptor: file_warehouseService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_warehouseService_proto_goTypes,
		DependencyIndexes: file_warehouseService_proto_depIdxs,
		MessageInfos:      file_warehouseService_proto_msgTypes,
	}.Build()
	File_warehouseService_proto = out.File
	file_warehouseService_proto_rawDesc = nil
	file_warehouseService_proto_goTypes = nil
	file_warehouseService_proto_depIdxs = nil
}
