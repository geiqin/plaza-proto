// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: familyMoveService.proto

package services

import (
	common "github.com/geiqin/micro-kit/protobuf/common"
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

//易地搬迁
type FamilyMove struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	FamilyId     int64   `protobuf:"varint,2,opt,name=family_id,json=familyId,proto3" json:"family_id"`
	RealName     string  `protobuf:"bytes,3,opt,name=real_name,json=realName,proto3" json:"real_name"`
	IdCard       int64   `protobuf:"varint,4,opt,name=id_card,json=idCard,proto3" json:"id_card"`
	MoveYear     int32   `protobuf:"varint,5,opt,name=move_year,json=moveYear,proto3" json:"move_year"`
	MoveMonth    int32   `protobuf:"varint,6,opt,name=move_month,json=moveMonth,proto3" json:"move_month"`
	MoveDay      int32   `protobuf:"varint,7,opt,name=move_day,json=moveDay,proto3" json:"move_day"`
	FromStreetId int64   `protobuf:"varint,8,opt,name=from_street_id,json=fromStreetId,proto3" json:"from_street_id"`
	FromAddr     string  `protobuf:"bytes,9,opt,name=from_addr,json=fromAddr,proto3" json:"from_addr"`
	ToStreetId   int64   `protobuf:"varint,10,opt,name=to_street_id,json=toStreetId,proto3" json:"to_street_id"`
	ToAddr       string  `protobuf:"bytes,11,opt,name=to_addr,json=toAddr,proto3" json:"to_addr"`
	Status       string  `protobuf:"bytes,12,opt,name=status,proto3" json:"status"`
	CreatedAt    string  `protobuf:"bytes,13,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt    string  `protobuf:"bytes,14,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	Family       *Family `protobuf:"bytes,15,opt,name=family,proto3" json:"family"`
	FromStreet   *Street `protobuf:"bytes,16,opt,name=from_street,json=fromStreet,proto3" json:"from_street"`
	ToStreet     *Street `protobuf:"bytes,17,opt,name=to_street,json=toStreet,proto3" json:"to_street"`
}

func (x *FamilyMove) Reset() {
	*x = FamilyMove{}
	if protoimpl.UnsafeEnabled {
		mi := &file_familyMoveService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FamilyMove) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FamilyMove) ProtoMessage() {}

func (x *FamilyMove) ProtoReflect() protoreflect.Message {
	mi := &file_familyMoveService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FamilyMove.ProtoReflect.Descriptor instead.
func (*FamilyMove) Descriptor() ([]byte, []int) {
	return file_familyMoveService_proto_rawDescGZIP(), []int{0}
}

func (x *FamilyMove) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FamilyMove) GetFamilyId() int64 {
	if x != nil {
		return x.FamilyId
	}
	return 0
}

func (x *FamilyMove) GetRealName() string {
	if x != nil {
		return x.RealName
	}
	return ""
}

func (x *FamilyMove) GetIdCard() int64 {
	if x != nil {
		return x.IdCard
	}
	return 0
}

func (x *FamilyMove) GetMoveYear() int32 {
	if x != nil {
		return x.MoveYear
	}
	return 0
}

func (x *FamilyMove) GetMoveMonth() int32 {
	if x != nil {
		return x.MoveMonth
	}
	return 0
}

func (x *FamilyMove) GetMoveDay() int32 {
	if x != nil {
		return x.MoveDay
	}
	return 0
}

func (x *FamilyMove) GetFromStreetId() int64 {
	if x != nil {
		return x.FromStreetId
	}
	return 0
}

func (x *FamilyMove) GetFromAddr() string {
	if x != nil {
		return x.FromAddr
	}
	return ""
}

func (x *FamilyMove) GetToStreetId() int64 {
	if x != nil {
		return x.ToStreetId
	}
	return 0
}

func (x *FamilyMove) GetToAddr() string {
	if x != nil {
		return x.ToAddr
	}
	return ""
}

func (x *FamilyMove) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *FamilyMove) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *FamilyMove) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *FamilyMove) GetFamily() *Family {
	if x != nil {
		return x.Family
	}
	return nil
}

func (x *FamilyMove) GetFromStreet() *Street {
	if x != nil {
		return x.FromStreet
	}
	return nil
}

func (x *FamilyMove) GetToStreet() *Street {
	if x != nil {
		return x.ToStreet
	}
	return nil
}

type FamilyMoveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int32  `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize int32  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Sorting  string `protobuf:"bytes,3,opt,name=sorting,proto3" json:"sorting"`
	//----
	FamilyId  int64   `protobuf:"varint,4,opt,name=family_id,json=familyId,proto3" json:"family_id"`
	MoveYear  int32   `protobuf:"varint,5,opt,name=move_year,json=moveYear,proto3" json:"move_year"`
	MoveMonth int32   `protobuf:"varint,6,opt,name=move_month,json=moveMonth,proto3" json:"move_month"`
	RealName  string  `protobuf:"bytes,7,opt,name=real_name,json=realName,proto3" json:"real_name"`
	IdCard    string  `protobuf:"bytes,8,opt,name=id_card,json=idCard,proto3" json:"id_card"`
	Status    string  `protobuf:"bytes,9,opt,name=status,proto3" json:"status"`
	Ids       []int64 `protobuf:"varint,10,rep,packed,name=ids,proto3" json:"ids"`
}

func (x *FamilyMoveRequest) Reset() {
	*x = FamilyMoveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_familyMoveService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FamilyMoveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FamilyMoveRequest) ProtoMessage() {}

func (x *FamilyMoveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_familyMoveService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FamilyMoveRequest.ProtoReflect.Descriptor instead.
func (*FamilyMoveRequest) Descriptor() ([]byte, []int) {
	return file_familyMoveService_proto_rawDescGZIP(), []int{1}
}

func (x *FamilyMoveRequest) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *FamilyMoveRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *FamilyMoveRequest) GetSorting() string {
	if x != nil {
		return x.Sorting
	}
	return ""
}

func (x *FamilyMoveRequest) GetFamilyId() int64 {
	if x != nil {
		return x.FamilyId
	}
	return 0
}

func (x *FamilyMoveRequest) GetMoveYear() int32 {
	if x != nil {
		return x.MoveYear
	}
	return 0
}

func (x *FamilyMoveRequest) GetMoveMonth() int32 {
	if x != nil {
		return x.MoveMonth
	}
	return 0
}

func (x *FamilyMoveRequest) GetRealName() string {
	if x != nil {
		return x.RealName
	}
	return ""
}

func (x *FamilyMoveRequest) GetIdCard() string {
	if x != nil {
		return x.IdCard
	}
	return ""
}

func (x *FamilyMoveRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *FamilyMoveRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type FamilyMoveData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *FamilyMove   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*FamilyMove `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Info   *common.Info  `protobuf:"bytes,4,opt,name=info,proto3" json:"info"`
}

func (x *FamilyMoveData) Reset() {
	*x = FamilyMoveData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_familyMoveService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FamilyMoveData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FamilyMoveData) ProtoMessage() {}

func (x *FamilyMoveData) ProtoReflect() protoreflect.Message {
	mi := &file_familyMoveService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FamilyMoveData.ProtoReflect.Descriptor instead.
func (*FamilyMoveData) Descriptor() ([]byte, []int) {
	return file_familyMoveService_proto_rawDescGZIP(), []int{2}
}

func (x *FamilyMoveData) GetEntity() *FamilyMove {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *FamilyMoveData) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *FamilyMoveData) GetItems() []*FamilyMove {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *FamilyMoveData) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type FamilyMoveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *FamilyMoveData `protobuf:"bytes,1,opt,name=data,proto3" json:"data"`
	Error *common.Error   `protobuf:"bytes,2,opt,name=error,proto3" json:"error"`
}

func (x *FamilyMoveResponse) Reset() {
	*x = FamilyMoveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_familyMoveService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FamilyMoveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FamilyMoveResponse) ProtoMessage() {}

func (x *FamilyMoveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_familyMoveService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FamilyMoveResponse.ProtoReflect.Descriptor instead.
func (*FamilyMoveResponse) Descriptor() ([]byte, []int) {
	return file_familyMoveService_proto_rawDescGZIP(), []int{3}
}

func (x *FamilyMoveResponse) GetData() *FamilyMoveData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *FamilyMoveResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_familyMoveService_proto protoreflect.FileDescriptor

var file_familyMoveService_proto_rawDesc = []byte{
	0x0a, 0x17, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x4d, 0x6f, 0x76, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x73, 0x74, 0x72,
	0x65, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xa6, 0x04, 0x0a, 0x0a, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x4d, 0x6f, 0x76, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x72, 0x65, 0x61, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x72, 0x65, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x64, 0x5f,
	0x63, 0x61, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x69, 0x64, 0x43, 0x61,
	0x72, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x79, 0x65, 0x61, 0x72, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6d, 0x6f, 0x76, 0x65, 0x59, 0x65, 0x61, 0x72, 0x12,
	0x1d, 0x0a, 0x0a, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x6d, 0x6f, 0x76, 0x65, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x19,
	0x0a, 0x08, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x64, 0x61, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x6d, 0x6f, 0x76, 0x65, 0x44, 0x61, 0x79, 0x12, 0x24, 0x0a, 0x0e, 0x66, 0x72, 0x6f,
	0x6d, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x53, 0x74, 0x72, 0x65, 0x65, 0x74, 0x49, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x66, 0x72, 0x6f, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x12, 0x20, 0x0a, 0x0c,
	0x74, 0x6f, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x74, 0x6f, 0x53, 0x74, 0x72, 0x65, 0x65, 0x74, 0x49, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x74, 0x6f, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x74, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x28, 0x0a,
	0x06, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x52,
	0x06, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x12, 0x31, 0x0a, 0x0b, 0x66, 0x72, 0x6f, 0x6d, 0x5f,
	0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x65, 0x74, 0x52, 0x0a,
	0x66, 0x72, 0x6f, 0x6d, 0x53, 0x74, 0x72, 0x65, 0x65, 0x74, 0x12, 0x2d, 0x0a, 0x09, 0x74, 0x6f,
	0x5f, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x65, 0x74, 0x52,
	0x08, 0x74, 0x6f, 0x53, 0x74, 0x72, 0x65, 0x65, 0x74, 0x22, 0x99, 0x02, 0x0a, 0x11, 0x46, 0x61,
	0x6d, 0x69, 0x6c, 0x79, 0x4d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1b, 0x0a, 0x09,
	0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x6f, 0x76,
	0x65, 0x5f, 0x79, 0x65, 0x61, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6d, 0x6f,
	0x76, 0x65, 0x59, 0x65, 0x61, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x6d,
	0x6f, 0x6e, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6d, 0x6f, 0x76, 0x65,
	0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x61, 0x6c, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x61, 0x6c, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x64, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x64, 0x43, 0x61, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x03,
	0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0xb1, 0x01, 0x0a, 0x0e, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79,
	0x4d, 0x6f, 0x76, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2c, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x4d, 0x6f, 0x76, 0x65, 0x52, 0x06,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50,
	0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x4d, 0x6f, 0x76, 0x65,
	0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x67, 0x0a, 0x12, 0x46, 0x61, 0x6d,
	0x69, 0x6c, 0x79, 0x4d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2c, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x4d,
	0x6f, 0x76, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x23, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x32, 0xd7, 0x02, 0x0a, 0x11, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x4d, 0x6f, 0x76,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x61,
	0x6d, 0x69, 0x6c, 0x79, 0x4d, 0x6f, 0x76, 0x65, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x4d, 0x6f, 0x76, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x61,
	0x6d, 0x69, 0x6c, 0x79, 0x4d, 0x6f, 0x76, 0x65, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x4d, 0x6f, 0x76, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x61,
	0x6d, 0x69, 0x6c, 0x79, 0x4d, 0x6f, 0x76, 0x65, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x4d, 0x6f, 0x76, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12,
	0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x61, 0x6d, 0x69, 0x6c,
	0x79, 0x4d, 0x6f, 0x76, 0x65, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x4d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12,
	0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x61, 0x6d, 0x69, 0x6c,
	0x79, 0x4d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x4d, 0x6f,
	0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a,
	0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_familyMoveService_proto_rawDescOnce sync.Once
	file_familyMoveService_proto_rawDescData = file_familyMoveService_proto_rawDesc
)

func file_familyMoveService_proto_rawDescGZIP() []byte {
	file_familyMoveService_proto_rawDescOnce.Do(func() {
		file_familyMoveService_proto_rawDescData = protoimpl.X.CompressGZIP(file_familyMoveService_proto_rawDescData)
	})
	return file_familyMoveService_proto_rawDescData
}

var file_familyMoveService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_familyMoveService_proto_goTypes = []interface{}{
	(*FamilyMove)(nil),         // 0: services.FamilyMove
	(*FamilyMoveRequest)(nil),  // 1: services.FamilyMoveRequest
	(*FamilyMoveData)(nil),     // 2: services.FamilyMoveData
	(*FamilyMoveResponse)(nil), // 3: services.FamilyMoveResponse
	(*Family)(nil),             // 4: services.Family
	(*Street)(nil),             // 5: services.Street
	(*common.Pager)(nil),       // 6: common.Pager
	(*common.Info)(nil),        // 7: common.Info
	(*common.Error)(nil),       // 8: common.Error
}
var file_familyMoveService_proto_depIdxs = []int32{
	4,  // 0: services.FamilyMove.family:type_name -> services.Family
	5,  // 1: services.FamilyMove.from_street:type_name -> services.Street
	5,  // 2: services.FamilyMove.to_street:type_name -> services.Street
	0,  // 3: services.FamilyMoveData.entity:type_name -> services.FamilyMove
	6,  // 4: services.FamilyMoveData.pager:type_name -> common.Pager
	0,  // 5: services.FamilyMoveData.items:type_name -> services.FamilyMove
	7,  // 6: services.FamilyMoveData.info:type_name -> common.Info
	2,  // 7: services.FamilyMoveResponse.data:type_name -> services.FamilyMoveData
	8,  // 8: services.FamilyMoveResponse.error:type_name -> common.Error
	0,  // 9: services.FamilyMoveService.Create:input_type -> services.FamilyMove
	0,  // 10: services.FamilyMoveService.Update:input_type -> services.FamilyMove
	0,  // 11: services.FamilyMoveService.Delete:input_type -> services.FamilyMove
	0,  // 12: services.FamilyMoveService.Get:input_type -> services.FamilyMove
	1,  // 13: services.FamilyMoveService.Search:input_type -> services.FamilyMoveRequest
	3,  // 14: services.FamilyMoveService.Create:output_type -> services.FamilyMoveResponse
	3,  // 15: services.FamilyMoveService.Update:output_type -> services.FamilyMoveResponse
	3,  // 16: services.FamilyMoveService.Delete:output_type -> services.FamilyMoveResponse
	3,  // 17: services.FamilyMoveService.Get:output_type -> services.FamilyMoveResponse
	3,  // 18: services.FamilyMoveService.Search:output_type -> services.FamilyMoveResponse
	14, // [14:19] is the sub-list for method output_type
	9,  // [9:14] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_familyMoveService_proto_init() }
func file_familyMoveService_proto_init() {
	if File_familyMoveService_proto != nil {
		return
	}
	file_familyService_proto_init()
	file_streetService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_familyMoveService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FamilyMove); i {
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
		file_familyMoveService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FamilyMoveRequest); i {
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
		file_familyMoveService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FamilyMoveData); i {
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
		file_familyMoveService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FamilyMoveResponse); i {
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
			RawDescriptor: file_familyMoveService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_familyMoveService_proto_goTypes,
		DependencyIndexes: file_familyMoveService_proto_depIdxs,
		MessageInfos:      file_familyMoveService_proto_msgTypes,
	}.Build()
	File_familyMoveService_proto = out.File
	file_familyMoveService_proto_rawDesc = nil
	file_familyMoveService_proto_goTypes = nil
	file_familyMoveService_proto_depIdxs = nil
}
