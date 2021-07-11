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

type FetchWhere struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int32  `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize int32  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Sorting  string `protobuf:"bytes,3,opt,name=sorting,proto3" json:"sorting"`
	Keywords string `protobuf:"bytes,4,opt,name=keywords,proto3" json:"keywords"`
	Id       int64  `protobuf:"varint,5,opt,name=id,proto3" json:"id"`
	// @inject_tag: gorm:"-"
	Ids  []int64 `protobuf:"varint,6,rep,packed,name=ids,proto3" json:"ids" gorm:"-"`
	Lng  string  `protobuf:"bytes,7,opt,name=lng,proto3" json:"lng"`
	Lat  string  `protobuf:"bytes,8,opt,name=lat,proto3" json:"lat"`
	Name string  `protobuf:"bytes,9,opt,name=name,proto3" json:"name"`
}

func (x *FetchWhere) Reset() {
	*x = FetchWhere{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fetchService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchWhere) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchWhere) ProtoMessage() {}

func (x *FetchWhere) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use FetchWhere.ProtoReflect.Descriptor instead.
func (*FetchWhere) Descriptor() ([]byte, []int) {
	return file_fetchService_proto_rawDescGZIP(), []int{0}
}

func (x *FetchWhere) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *FetchWhere) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *FetchWhere) GetSorting() string {
	if x != nil {
		return x.Sorting
	}
	return ""
}

func (x *FetchWhere) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *FetchWhere) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FetchWhere) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *FetchWhere) GetLng() string {
	if x != nil {
		return x.Lng
	}
	return ""
}

func (x *FetchWhere) GetLat() string {
	if x != nil {
		return x.Lat
	}
	return ""
}

func (x *FetchWhere) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Fetch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Name                 string `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	AreaId               int64  `protobuf:"varint,3,opt,name=area_id,json=areaId,proto3" json:"area_id"`
	Addr                 string `protobuf:"bytes,4,opt,name=addr,proto3" json:"addr"`
	Lng                  string `protobuf:"bytes,5,opt,name=lng,proto3" json:"lng"`
	Lat                  string `protobuf:"bytes,6,opt,name=lat,proto3" json:"lat"`
	Tel                  string `protobuf:"bytes,7,opt,name=tel,proto3" json:"tel"`
	Mobile               string `protobuf:"bytes,8,opt,name=mobile,proto3" json:"mobile"`
	Reception            string `protobuf:"bytes,9,opt,name=reception,proto3" json:"reception"`
	ReceptionRepeatWeeks int32  `protobuf:"varint,10,opt,name=reception_repeat_weeks,json=receptionRepeatWeeks,proto3" json:"reception_repeat_weeks"`
	// @inject_tag: gorm:"-"
	ReceptionRepeatWeekArr []int32 `protobuf:"varint,11,rep,packed,name=reception_repeat_week_arr,json=receptionRepeatWeekArr,proto3" json:"reception_repeat_week_arr" gorm:"-"`
	IsFetchTime            bool    `protobuf:"varint,12,opt,name=is_fetch_time,json=isFetchTime,proto3" json:"is_fetch_time"`
	Fetch                  string  `protobuf:"bytes,13,opt,name=fetch,proto3" json:"fetch"`
	FetchRepeatWeeks       int32   `protobuf:"varint,14,opt,name=fetch_repeat_weeks,json=fetchRepeatWeeks,proto3" json:"fetch_repeat_weeks"`
	// @inject_tag: gorm:"-"
	FetchRepeatWeekArr []int32 `protobuf:"varint,15,rep,packed,name=fetch_repeat_week_arr,json=fetchRepeatWeekArr,proto3" json:"fetch_repeat_week_arr" gorm:"-"`
	SubFetchTime       string  `protobuf:"bytes,16,opt,name=sub_fetch_time,json=subFetchTime,proto3" json:"sub_fetch_time"`
	Appointment        string  `protobuf:"bytes,17,opt,name=appointment,proto3" json:"appointment"`
	AppointmentNum     int32   `protobuf:"varint,18,opt,name=appointment_num,json=appointmentNum,proto3" json:"appointment_num"`
	MaxAppointmentNum  int32   `protobuf:"varint,19,opt,name=max_appointment_num,json=maxAppointmentNum,proto3" json:"max_appointment_num"`
	Memo               string  `protobuf:"bytes,20,opt,name=memo,proto3" json:"memo"`
	CreatedAt          string  `protobuf:"bytes,21,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt          string  `protobuf:"bytes,22,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	// @inject_tag: gorm:"-"
	ReceptionTimes []*Times `protobuf:"bytes,23,rep,name=reception_times,json=receptionTimes,proto3" json:"reception_times" gorm:"-"`
	// @inject_tag: gorm:"-"
	FetchTimes []*Times        `protobuf:"bytes,24,rep,name=fetch_times,json=fetchTimes,proto3" json:"fetch_times" gorm:"-"`
	Galleries  []*FetchGallery `protobuf:"bytes,25,rep,name=galleries,proto3" json:"galleries"`
	// @inject_tag: gorm:"-"
	Area     *AreaInfo `protobuf:"bytes,26,opt,name=area,proto3" json:"area" gorm:"-"`
	Distance float32   `protobuf:"fixed32,27,opt,name=distance,proto3" json:"distance"`
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

type FetchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Fetch        `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*Fetch      `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Error  *common.Error `protobuf:"bytes,4,opt,name=error,proto3" json:"error"`
	Info   *common.Info  `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
}

func (x *FetchResponse) Reset() {
	*x = FetchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fetchService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchResponse) ProtoMessage() {}

func (x *FetchResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use FetchResponse.ProtoReflect.Descriptor instead.
func (*FetchResponse) Descriptor() ([]byte, []int) {
	return file_fetchService_proto_rawDescGZIP(), []int{2}
}

func (x *FetchResponse) GetEntity() *Fetch {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *FetchResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *FetchResponse) GetItems() []*Fetch {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *FetchResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *FetchResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
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
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcf, 0x01,
	0x0a, 0x0a, 0x46, 0x65, 0x74, 0x63, 0x68, 0x57, 0x68, 0x65, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x67,
	0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79,
	0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79,
	0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6e, 0x67, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x6e, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0xa9, 0x07, 0x0a, 0x05, 0x46, 0x65, 0x74, 0x63, 0x68, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a,
	0x07, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x61, 0x72, 0x65, 0x61, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6e,
	0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x6e, 0x67, 0x12, 0x10, 0x0a, 0x03,
	0x6c, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x74, 0x65, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x65, 0x6c,
	0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x65,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x63,
	0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x16, 0x72, 0x65, 0x63, 0x65, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x5f, 0x77, 0x65, 0x65, 0x6b, 0x73,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x14, 0x72, 0x65, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x57, 0x65, 0x65, 0x6b, 0x73, 0x12, 0x39, 0x0a, 0x19,
	0x72, 0x65, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74,
	0x5f, 0x77, 0x65, 0x65, 0x6b, 0x5f, 0x61, 0x72, 0x72, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x05, 0x52,
	0x16, 0x72, 0x65, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74,
	0x57, 0x65, 0x65, 0x6b, 0x41, 0x72, 0x72, 0x12, 0x22, 0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x66, 0x65,
	0x74, 0x63, 0x68, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b,
	0x69, 0x73, 0x46, 0x65, 0x74, 0x63, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66,
	0x65, 0x74, 0x63, 0x68, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x65, 0x74, 0x63,
	0x68, 0x12, 0x2c, 0x0a, 0x12, 0x66, 0x65, 0x74, 0x63, 0x68, 0x5f, 0x72, 0x65, 0x70, 0x65, 0x61,
	0x74, 0x5f, 0x77, 0x65, 0x65, 0x6b, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x66,
	0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x57, 0x65, 0x65, 0x6b, 0x73, 0x12,
	0x31, 0x0a, 0x15, 0x66, 0x65, 0x74, 0x63, 0x68, 0x5f, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x5f,
	0x77, 0x65, 0x65, 0x6b, 0x5f, 0x61, 0x72, 0x72, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x05, 0x52, 0x12,
	0x66, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x57, 0x65, 0x65, 0x6b, 0x41,
	0x72, 0x72, 0x12, 0x24, 0x0a, 0x0e, 0x73, 0x75, 0x62, 0x5f, 0x66, 0x65, 0x74, 0x63, 0x68, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x46,
	0x65, 0x74, 0x63, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x70, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61,
	0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x70,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x12, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0e, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x4e, 0x75, 0x6d, 0x12, 0x2e, 0x0a, 0x13, 0x6d, 0x61, 0x78, 0x5f, 0x61, 0x70, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x13, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x11, 0x6d, 0x61, 0x78, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x4e, 0x75, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x14, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x0f, 0x72, 0x65, 0x63, 0x65, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x17, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x52,
	0x0e, 0x72, 0x65, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x12,
	0x30, 0x0a, 0x0b, 0x66, 0x65, 0x74, 0x63, 0x68, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x18,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x52, 0x0a, 0x66, 0x65, 0x74, 0x63, 0x68, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x12, 0x34, 0x0a, 0x09, 0x67, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x69, 0x65, 0x73, 0x18, 0x19,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x46, 0x65, 0x74, 0x63, 0x68, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x52, 0x09, 0x67, 0x61,
	0x6c, 0x6c, 0x65, 0x72, 0x69, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x04, 0x61, 0x72, 0x65, 0x61, 0x18,
	0x1a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x41, 0x72, 0x65, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x61, 0x72, 0x65, 0x61, 0x12,
	0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x1b, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x22, 0xcb, 0x01, 0x0a, 0x0d,
	0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a,
	0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x06,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50,
	0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0xb0, 0x01, 0x0a, 0x0f, 0x46, 0x65,
	0x74, 0x63, 0x68, 0x43, 0x66, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a,
	0x04, 0x4f, 0x70, 0x65, 0x6e, 0x12, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x33, 0x0a, 0x05, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x12, 0x0d, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x06, 0x49, 0x73, 0x4f, 0x70, 0x65, 0x6e,
	0x12, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x32, 0xda, 0x03, 0x0a,
	0x0c, 0x46, 0x65, 0x74, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a,
	0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x1a, 0x17,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x06, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46,
	0x65, 0x74, 0x63, 0x68, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0f, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x1a, 0x17, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x65, 0x74,
	0x63, 0x68, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x37, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x57, 0x68, 0x65, 0x72, 0x65,
	0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x65, 0x74, 0x63,
	0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x08, 0x47,
	0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x1a, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x09, 0x47,
	0x65, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x17,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_fetchService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_fetchService_proto_goTypes = []interface{}{
	(*FetchWhere)(nil),       // 0: services.FetchWhere
	(*Fetch)(nil),            // 1: services.Fetch
	(*FetchResponse)(nil),    // 2: services.FetchResponse
	(*Times)(nil),            // 3: services.Times
	(*FetchGallery)(nil),     // 4: services.FetchGallery
	(*AreaInfo)(nil),         // 5: services.AreaInfo
	(*common.Pager)(nil),     // 6: common.Pager
	(*common.Error)(nil),     // 7: common.Error
	(*common.Info)(nil),      // 8: common.Info
	(*common.Empty)(nil),     // 9: common.Empty
	(*SettingResponse)(nil),  // 10: services.SettingResponse
	(*DateListResponse)(nil), // 11: services.DateListResponse
}
var file_fetchService_proto_depIdxs = []int32{
	3,  // 0: services.Fetch.reception_times:type_name -> services.Times
	3,  // 1: services.Fetch.fetch_times:type_name -> services.Times
	4,  // 2: services.Fetch.galleries:type_name -> services.FetchGallery
	5,  // 3: services.Fetch.area:type_name -> services.AreaInfo
	1,  // 4: services.FetchResponse.entity:type_name -> services.Fetch
	6,  // 5: services.FetchResponse.pager:type_name -> common.Pager
	1,  // 6: services.FetchResponse.items:type_name -> services.Fetch
	7,  // 7: services.FetchResponse.error:type_name -> common.Error
	8,  // 8: services.FetchResponse.info:type_name -> common.Info
	9,  // 9: services.FetchCfgService.Open:input_type -> common.Empty
	9,  // 10: services.FetchCfgService.Close:input_type -> common.Empty
	9,  // 11: services.FetchCfgService.IsOpen:input_type -> common.Empty
	1,  // 12: services.FetchService.Create:input_type -> services.Fetch
	1,  // 13: services.FetchService.Update:input_type -> services.Fetch
	0,  // 14: services.FetchService.Delete:input_type -> services.FetchWhere
	1,  // 15: services.FetchService.Get:input_type -> services.Fetch
	0,  // 16: services.FetchService.Search:input_type -> services.FetchWhere
	0,  // 17: services.FetchService.List:input_type -> services.FetchWhere
	0,  // 18: services.FetchService.GetTimes:input_type -> services.FetchWhere
	0,  // 19: services.FetchService.GetLatest:input_type -> services.FetchWhere
	10, // 20: services.FetchCfgService.Open:output_type -> services.SettingResponse
	10, // 21: services.FetchCfgService.Close:output_type -> services.SettingResponse
	10, // 22: services.FetchCfgService.IsOpen:output_type -> services.SettingResponse
	2,  // 23: services.FetchService.Create:output_type -> services.FetchResponse
	2,  // 24: services.FetchService.Update:output_type -> services.FetchResponse
	2,  // 25: services.FetchService.Delete:output_type -> services.FetchResponse
	2,  // 26: services.FetchService.Get:output_type -> services.FetchResponse
	2,  // 27: services.FetchService.Search:output_type -> services.FetchResponse
	2,  // 28: services.FetchService.List:output_type -> services.FetchResponse
	11, // 29: services.FetchService.GetTimes:output_type -> services.DateListResponse
	2,  // 30: services.FetchService.GetLatest:output_type -> services.FetchResponse
	20, // [20:31] is the sub-list for method output_type
	9,  // [9:20] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_fetchService_proto_init() }
func file_fetchService_proto_init() {
	if File_fetchService_proto != nil {
		return
	}
	file_timesService_proto_init()
	file_fetchGalleryService_proto_init()
	file_areaInfoService_proto_init()
	file_settingService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_fetchService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchWhere); i {
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
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   2,
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
