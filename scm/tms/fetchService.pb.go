// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: fetchService.proto

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

type FetchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int32  `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize int32  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Sorting  string `protobuf:"bytes,3,opt,name=sorting,proto3" json:"sorting"`
	Keywords string `protobuf:"bytes,4,opt,name=keywords,proto3" json:"keywords"`
	Id       int64  `protobuf:"varint,5,opt,name=id,proto3" json:"id"`
	// @inject_tag: gorm:"-"
	Ids  []int64 `protobuf:"varint,6,rep,packed,name=ids,proto3" json:"ids"`
	Lng  string  `protobuf:"bytes,7,opt,name=lng,proto3" json:"lng"`
	Lat  string  `protobuf:"bytes,8,opt,name=lat,proto3" json:"lat"`
	Name string  `protobuf:"bytes,9,opt,name=name,proto3" json:"name"`
}

func (x *FetchRequest) Reset() {
	*x = FetchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fetchService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchRequest) ProtoMessage() {}

func (x *FetchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fetchService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchRequest.ProtoReflect.Descriptor instead.
func (*FetchRequest) Descriptor() ([]byte, []int) {
	return file_fetchService_proto_rawDescGZIP(), []int{0}
}

func (x *FetchRequest) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *FetchRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *FetchRequest) GetSorting() string {
	if x != nil {
		return x.Sorting
	}
	return ""
}

func (x *FetchRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *FetchRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FetchRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *FetchRequest) GetLng() string {
	if x != nil {
		return x.Lng
	}
	return ""
}

func (x *FetchRequest) GetLat() string {
	if x != nil {
		return x.Lat
	}
	return ""
}

func (x *FetchRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Fetch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                     int64           `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Name                   string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	AreaId                 int64           `protobuf:"varint,3,opt,name=area_id,json=areaId,proto3" json:"area_id"`
	Addr                   string          `protobuf:"bytes,4,opt,name=addr,proto3" json:"addr"`
	Lng                    string          `protobuf:"bytes,5,opt,name=lng,proto3" json:"lng"`
	Lat                    string          `protobuf:"bytes,6,opt,name=lat,proto3" json:"lat"`
	Tel                    string          `protobuf:"bytes,7,opt,name=tel,proto3" json:"tel"`
	Mobile                 string          `protobuf:"bytes,8,opt,name=mobile,proto3" json:"mobile"`
	Reception              string          `protobuf:"bytes,9,opt,name=reception,proto3" json:"reception"`
	ReceptionRepeatWeeks   int32           `protobuf:"varint,10,opt,name=reception_repeat_weeks,json=receptionRepeatWeeks,proto3" json:"reception_repeat_weeks"`
	ReceptionRepeatWeekArr []int32         `protobuf:"varint,11,rep,packed,name=reception_repeat_week_arr,json=receptionRepeatWeekArr,proto3" json:"reception_repeat_week_arr"`
	IsFetchTime            bool            `protobuf:"varint,12,opt,name=is_fetch_time,json=isFetchTime,proto3" json:"is_fetch_time"`
	Fetch                  string          `protobuf:"bytes,13,opt,name=fetch,proto3" json:"fetch"`
	FetchRepeatWeeks       int32           `protobuf:"varint,14,opt,name=fetch_repeat_weeks,json=fetchRepeatWeeks,proto3" json:"fetch_repeat_weeks"`
	FetchRepeatWeekArr     []int32         `protobuf:"varint,15,rep,packed,name=fetch_repeat_week_arr,json=fetchRepeatWeekArr,proto3" json:"fetch_repeat_week_arr"`
	SubFetchTime           string          `protobuf:"bytes,16,opt,name=sub_fetch_time,json=subFetchTime,proto3" json:"sub_fetch_time"`
	Appointment            string          `protobuf:"bytes,17,opt,name=appointment,proto3" json:"appointment"`
	AppointmentNum         int32           `protobuf:"varint,18,opt,name=appointment_num,json=appointmentNum,proto3" json:"appointment_num"`
	MaxAppointmentNum      int32           `protobuf:"varint,19,opt,name=max_appointment_num,json=maxAppointmentNum,proto3" json:"max_appointment_num"`
	Memo                   string          `protobuf:"bytes,20,opt,name=memo,proto3" json:"memo"`
	CreatedAt              string          `protobuf:"bytes,21,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt              string          `protobuf:"bytes,22,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	ReceptionTimes         []*Times        `protobuf:"bytes,23,rep,name=reception_times,json=receptionTimes,proto3" json:"reception_times"`
	FetchTimes             []*Times        `protobuf:"bytes,24,rep,name=fetch_times,json=fetchTimes,proto3" json:"fetch_times"`
	Galleries              []*FetchGallery `protobuf:"bytes,25,rep,name=galleries,proto3" json:"galleries"`
	Area                   *AreaInfo       `protobuf:"bytes,26,opt,name=area,proto3" json:"area"`
	Distance               float32         `protobuf:"fixed32,27,opt,name=distance,proto3" json:"distance"`
}

func (x *Fetch) Reset() {
	*x = Fetch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fetchService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Fetch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Fetch) ProtoMessage() {}

func (x *Fetch) ProtoReflect() protoreflect.Message {
	mi := &file_fetchService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Fetch.ProtoReflect.Descriptor instead.
func (*Fetch) Descriptor() ([]byte, []int) {
	return file_fetchService_proto_rawDescGZIP(), []int{1}
}

func (x *Fetch) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Fetch) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Fetch) GetAreaId() int64 {
	if x != nil {
		return x.AreaId
	}
	return 0
}

func (x *Fetch) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Fetch) GetLng() string {
	if x != nil {
		return x.Lng
	}
	return ""
}

func (x *Fetch) GetLat() string {
	if x != nil {
		return x.Lat
	}
	return ""
}

func (x *Fetch) GetTel() string {
	if x != nil {
		return x.Tel
	}
	return ""
}

func (x *Fetch) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *Fetch) GetReception() string {
	if x != nil {
		return x.Reception
	}
	return ""
}

func (x *Fetch) GetReceptionRepeatWeeks() int32 {
	if x != nil {
		return x.ReceptionRepeatWeeks
	}
	return 0
}

func (x *Fetch) GetReceptionRepeatWeekArr() []int32 {
	if x != nil {
		return x.ReceptionRepeatWeekArr
	}
	return nil
}

func (x *Fetch) GetIsFetchTime() bool {
	if x != nil {
		return x.IsFetchTime
	}
	return false
}

func (x *Fetch) GetFetch() string {
	if x != nil {
		return x.Fetch
	}
	return ""
}

func (x *Fetch) GetFetchRepeatWeeks() int32 {
	if x != nil {
		return x.FetchRepeatWeeks
	}
	return 0
}

func (x *Fetch) GetFetchRepeatWeekArr() []int32 {
	if x != nil {
		return x.FetchRepeatWeekArr
	}
	return nil
}

func (x *Fetch) GetSubFetchTime() string {
	if x != nil {
		return x.SubFetchTime
	}
	return ""
}

func (x *Fetch) GetAppointment() string {
	if x != nil {
		return x.Appointment
	}
	return ""
}

func (x *Fetch) GetAppointmentNum() int32 {
	if x != nil {
		return x.AppointmentNum
	}
	return 0
}

func (x *Fetch) GetMaxAppointmentNum() int32 {
	if x != nil {
		return x.MaxAppointmentNum
	}
	return 0
}

func (x *Fetch) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *Fetch) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Fetch) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Fetch) GetReceptionTimes() []*Times {
	if x != nil {
		return x.ReceptionTimes
	}
	return nil
}

func (x *Fetch) GetFetchTimes() []*Times {
	if x != nil {
		return x.FetchTimes
	}
	return nil
}

func (x *Fetch) GetGalleries() []*FetchGallery {
	if x != nil {
		return x.Galleries
	}
	return nil
}

func (x *Fetch) GetArea() *AreaInfo {
	if x != nil {
		return x.Area
	}
	return nil
}

func (x *Fetch) GetDistance() float32 {
	if x != nil {
		return x.Distance
	}
	return 0
}

type FetchData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Fetch        `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*Fetch      `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Info   *common.Info  `protobuf:"bytes,4,opt,name=info,proto3" json:"info"`
}

func (x *FetchData) Reset() {
	*x = FetchData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fetchService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchData) ProtoMessage() {}

func (x *FetchData) ProtoReflect() protoreflect.Message {
	mi := &file_fetchService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchData.ProtoReflect.Descriptor instead.
func (*FetchData) Descriptor() ([]byte, []int) {
	return file_fetchService_proto_rawDescGZIP(), []int{2}
}

func (x *FetchData) GetEntity() *Fetch {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *FetchData) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *FetchData) GetItems() []*Fetch {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *FetchData) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type FetchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *FetchData    `protobuf:"bytes,1,opt,name=data,proto3" json:"data"`
	Error *common.Error `protobuf:"bytes,2,opt,name=error,proto3" json:"error"`
}

func (x *FetchResponse) Reset() {
	*x = FetchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fetchService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchResponse) ProtoMessage() {}

func (x *FetchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_fetchService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchResponse.ProtoReflect.Descriptor instead.
func (*FetchResponse) Descriptor() ([]byte, []int) {
	return file_fetchService_proto_rawDescGZIP(), []int{3}
}

func (x *FetchResponse) GetData() *FetchData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *FetchResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_fetchService_proto protoreflect.FileDescriptor

var file_fetchService_proto_rawDesc = []byte{
	0x0a, 0x12, 0x66, 0x65, 0x74, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x12, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x66, 0x65, 0x74, 0x63, 0x68, 0x47, 0x61, 0x6c, 0x6c,
	0x65, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x15, 0x61, 0x72, 0x65, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd1, 0x01, 0x0a, 0x0c, 0x46, 0x65, 0x74, 0x63,
	0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6f,
	0x72, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64,
	0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03,
	0x69, 0x64, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6e, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6c, 0x6e, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xa9, 0x07, 0x0a, 0x05,
	0x46, 0x65, 0x74, 0x63, 0x68, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x72, 0x65,
	0x61, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x72, 0x65, 0x61,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6e, 0x67, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x6e, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x65,
	0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x65, 0x6c, 0x12, 0x16, 0x0a, 0x06,
	0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f,
	0x62, 0x69, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x63, 0x65, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x16, 0x72, 0x65, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x5f, 0x77, 0x65, 0x65, 0x6b, 0x73, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x14, 0x72, 0x65, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70,
	0x65, 0x61, 0x74, 0x57, 0x65, 0x65, 0x6b, 0x73, 0x12, 0x39, 0x0a, 0x19, 0x72, 0x65, 0x63, 0x65,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x5f, 0x77, 0x65, 0x65,
	0x6b, 0x5f, 0x61, 0x72, 0x72, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x05, 0x52, 0x16, 0x72, 0x65, 0x63,
	0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x57, 0x65, 0x65, 0x6b,
	0x41, 0x72, 0x72, 0x12, 0x22, 0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x66, 0x65, 0x74, 0x63, 0x68, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x46, 0x65,
	0x74, 0x63, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x65, 0x74, 0x63, 0x68,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x65, 0x74, 0x63, 0x68, 0x12, 0x2c, 0x0a,
	0x12, 0x66, 0x65, 0x74, 0x63, 0x68, 0x5f, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x5f, 0x77, 0x65,
	0x65, 0x6b, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x66, 0x65, 0x74, 0x63, 0x68,
	0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x57, 0x65, 0x65, 0x6b, 0x73, 0x12, 0x31, 0x0a, 0x15, 0x66,
	0x65, 0x74, 0x63, 0x68, 0x5f, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x5f, 0x77, 0x65, 0x65, 0x6b,
	0x5f, 0x61, 0x72, 0x72, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x05, 0x52, 0x12, 0x66, 0x65, 0x74, 0x63,
	0x68, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x57, 0x65, 0x65, 0x6b, 0x41, 0x72, 0x72, 0x12, 0x24,
	0x0a, 0x0e, 0x73, 0x75, 0x62, 0x5f, 0x66, 0x65, 0x74, 0x63, 0x68, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x46, 0x65, 0x74, 0x63, 0x68,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x70, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x12, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0e, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x12,
	0x2e, 0x0a, 0x13, 0x6d, 0x61, 0x78, 0x5f, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x13, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x6d, 0x61,
	0x78, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x12,
	0x12, 0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d,
	0x65, 0x6d, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x38, 0x0a, 0x0f, 0x72, 0x65, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x18, 0x17, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x52, 0x0e, 0x72, 0x65, 0x63,
	0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x30, 0x0a, 0x0b, 0x66,
	0x65, 0x74, 0x63, 0x68, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x18, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x52, 0x0a, 0x66, 0x65, 0x74, 0x63, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x34, 0x0a,
	0x09, 0x67, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x69, 0x65, 0x73, 0x18, 0x19, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x65, 0x74, 0x63,
	0x68, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x52, 0x09, 0x67, 0x61, 0x6c, 0x6c, 0x65, 0x72,
	0x69, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x04, 0x61, 0x72, 0x65, 0x61, 0x18, 0x1a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x72, 0x65,
	0x61, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x61, 0x72, 0x65, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x64,
	0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x64,
	0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x22, 0xa2, 0x01, 0x0a, 0x09, 0x46, 0x65, 0x74, 0x63,
	0x68, 0x44, 0x61, 0x74, 0x61, 0x12, 0x27, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23,
	0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61,
	0x67, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x65,
	0x74, 0x63, 0x68, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x5d, 0x0a, 0x0d,
	0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xe4, 0x03, 0x0a, 0x0c,
	0x46, 0x65, 0x74, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x06,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x34, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x1a, 0x17, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x65,
	0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0f, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x1a, 0x17, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x12, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x65,
	0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x40, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x44, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x3e, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fetchService_proto_rawDescOnce sync.Once
	file_fetchService_proto_rawDescData = file_fetchService_proto_rawDesc
)

func file_fetchService_proto_rawDescGZIP() []byte {
	file_fetchService_proto_rawDescOnce.Do(func() {
		file_fetchService_proto_rawDescData = protoimpl.X.CompressGZIP(file_fetchService_proto_rawDescData)
	})
	return file_fetchService_proto_rawDescData
}

var file_fetchService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_fetchService_proto_goTypes = []interface{}{
	(*FetchRequest)(nil),     // 0: services.FetchRequest
	(*Fetch)(nil),            // 1: services.Fetch
	(*FetchData)(nil),        // 2: services.FetchData
	(*FetchResponse)(nil),    // 3: services.FetchResponse
	(*Times)(nil),            // 4: services.Times
	(*FetchGallery)(nil),     // 5: services.FetchGallery
	(*AreaInfo)(nil),         // 6: services.AreaInfo
	(*common.Pager)(nil),     // 7: common.Pager
	(*common.Info)(nil),      // 8: common.Info
	(*common.Error)(nil),     // 9: common.Error
	(*DateListResponse)(nil), // 10: services.DateListResponse
}
var file_fetchService_proto_depIdxs = []int32{
	4,  // 0: services.Fetch.reception_times:type_name -> services.Times
	4,  // 1: services.Fetch.fetch_times:type_name -> services.Times
	5,  // 2: services.Fetch.galleries:type_name -> services.FetchGallery
	6,  // 3: services.Fetch.area:type_name -> services.AreaInfo
	1,  // 4: services.FetchData.entity:type_name -> services.Fetch
	7,  // 5: services.FetchData.pager:type_name -> common.Pager
	1,  // 6: services.FetchData.items:type_name -> services.Fetch
	8,  // 7: services.FetchData.info:type_name -> common.Info
	2,  // 8: services.FetchResponse.data:type_name -> services.FetchData
	9,  // 9: services.FetchResponse.error:type_name -> common.Error
	1,  // 10: services.FetchService.Create:input_type -> services.Fetch
	1,  // 11: services.FetchService.Update:input_type -> services.Fetch
	0,  // 12: services.FetchService.Delete:input_type -> services.FetchRequest
	1,  // 13: services.FetchService.Get:input_type -> services.Fetch
	0,  // 14: services.FetchService.Search:input_type -> services.FetchRequest
	0,  // 15: services.FetchService.List:input_type -> services.FetchRequest
	0,  // 16: services.FetchService.GetTimes:input_type -> services.FetchRequest
	0,  // 17: services.FetchService.GetLatest:input_type -> services.FetchRequest
	3,  // 18: services.FetchService.Create:output_type -> services.FetchResponse
	3,  // 19: services.FetchService.Update:output_type -> services.FetchResponse
	3,  // 20: services.FetchService.Delete:output_type -> services.FetchResponse
	3,  // 21: services.FetchService.Get:output_type -> services.FetchResponse
	3,  // 22: services.FetchService.Search:output_type -> services.FetchResponse
	3,  // 23: services.FetchService.List:output_type -> services.FetchResponse
	10, // 24: services.FetchService.GetTimes:output_type -> services.DateListResponse
	3,  // 25: services.FetchService.GetLatest:output_type -> services.FetchResponse
	18, // [18:26] is the sub-list for method output_type
	10, // [10:18] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_fetchService_proto_init() }
func file_fetchService_proto_init() {
	if File_fetchService_proto != nil {
		return
	}
	file_timesService_proto_init()
	file_fetchGalleryService_proto_init()
	file_areaInfoService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_fetchService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchRequest); i {
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
		file_fetchService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Fetch); i {
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
		file_fetchService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchData); i {
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
		file_fetchService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchResponse); i {
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
			RawDescriptor: file_fetchService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_fetchService_proto_goTypes,
		DependencyIndexes: file_fetchService_proto_depIdxs,
		MessageInfos:      file_fetchService_proto_msgTypes,
	}.Build()
	File_fetchService_proto = out.File
	file_fetchService_proto_rawDesc = nil
	file_fetchService_proto_goTypes = nil
	file_fetchService_proto_depIdxs = nil
}
