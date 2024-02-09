// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: slideService.proto

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

type Slide struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Type       string `protobuf:"bytes,2,opt,name=type,proto3" json:"type"`
	Platform   string `protobuf:"bytes,3,opt,name=platform,proto3" json:"platform"`
	Name       string `protobuf:"bytes,4,opt,name=name,proto3" json:"name"`
	EventType  string `protobuf:"bytes,5,opt,name=event_type,json=eventType,proto3" json:"event_type"`
	EventValue string `protobuf:"bytes,6,opt,name=event_value,json=eventValue,proto3" json:"event_value"`
	ImageUrl   string `protobuf:"bytes,7,opt,name=image_url,json=imageUrl,proto3" json:"image_url"`
	Desc       string `protobuf:"bytes,8,opt,name=desc,proto3" json:"desc"`
	BgColor    string `protobuf:"bytes,9,opt,name=bg_color,json=bgColor,proto3" json:"bg_color"`
	Sort       int32  `protobuf:"varint,12,opt,name=sort,proto3" json:"sort"`
	Status     string `protobuf:"bytes,14,opt,name=status,proto3" json:"status"`
	CreatedAt  string `protobuf:"bytes,15,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt  string `protobuf:"bytes,16,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	//extends
	PlatformName  string `protobuf:"bytes,17,opt,name=platform_name,json=platformName,proto3" json:"platform_name"`
	EventTypeName string `protobuf:"bytes,18,opt,name=event_type_name,json=eventTypeName,proto3" json:"event_type_name"`
	StatusName    string `protobuf:"bytes,19,opt,name=status_name,json=statusName,proto3" json:"status_name"`
}

func (x *Slide) Reset() {
	*x = Slide{}
	if protoimpl.UnsafeEnabled {
		mi := &file_slideService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Slide) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Slide) ProtoMessage() {}

func (x *Slide) ProtoReflect() protoreflect.Message {
	mi := &file_slideService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Slide.ProtoReflect.Descriptor instead.
func (*Slide) Descriptor() ([]byte, []int) {
	return file_slideService_proto_rawDescGZIP(), []int{0}
}

func (x *Slide) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Slide) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Slide) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *Slide) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Slide) GetEventType() string {
	if x != nil {
		return x.EventType
	}
	return ""
}

func (x *Slide) GetEventValue() string {
	if x != nil {
		return x.EventValue
	}
	return ""
}

func (x *Slide) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *Slide) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *Slide) GetBgColor() string {
	if x != nil {
		return x.BgColor
	}
	return ""
}

func (x *Slide) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *Slide) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Slide) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Slide) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Slide) GetPlatformName() string {
	if x != nil {
		return x.PlatformName
	}
	return ""
}

func (x *Slide) GetEventTypeName() string {
	if x != nil {
		return x.EventTypeName
	}
	return ""
}

func (x *Slide) GetStatusName() string {
	if x != nil {
		return x.StatusName
	}
	return ""
}

type SlideRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int64  `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize int64  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Sorting  string `protobuf:"bytes,3,opt,name=sorting,proto3" json:"sorting"`
	//base params
	Id        int64   `protobuf:"varint,4,opt,name=id,proto3" json:"id"`
	Name      string  `protobuf:"bytes,5,opt,name=name,proto3" json:"name"`
	Type      string  `protobuf:"bytes,6,opt,name=type,proto3" json:"type"`
	Status    string  `protobuf:"bytes,7,opt,name=status,proto3" json:"status"`
	Platform  string  `protobuf:"bytes,8,opt,name=platform,proto3" json:"platform"`
	EventType string  `protobuf:"bytes,9,opt,name=event_type,json=eventType,proto3" json:"event_type"`
	Keywords  string  `protobuf:"bytes,10,opt,name=keywords,proto3" json:"keywords"`
	Ids       []int64 `protobuf:"varint,11,rep,packed,name=ids,proto3" json:"ids"`
}

func (x *SlideRequest) Reset() {
	*x = SlideRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_slideService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlideRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlideRequest) ProtoMessage() {}

func (x *SlideRequest) ProtoReflect() protoreflect.Message {
	mi := &file_slideService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlideRequest.ProtoReflect.Descriptor instead.
func (*SlideRequest) Descriptor() ([]byte, []int) {
	return file_slideService_proto_rawDescGZIP(), []int{1}
}

func (x *SlideRequest) GetPaged() int64 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *SlideRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *SlideRequest) GetSorting() string {
	if x != nil {
		return x.Sorting
	}
	return ""
}

func (x *SlideRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SlideRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SlideRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *SlideRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *SlideRequest) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *SlideRequest) GetEventType() string {
	if x != nil {
		return x.EventType
	}
	return ""
}

func (x *SlideRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *SlideRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type SlideResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Slide        `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*Slide      `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Info   string        `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
}

func (x *SlideResponse) Reset() {
	*x = SlideResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_slideService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlideResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlideResponse) ProtoMessage() {}

func (x *SlideResponse) ProtoReflect() protoreflect.Message {
	mi := &file_slideService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlideResponse.ProtoReflect.Descriptor instead.
func (*SlideResponse) Descriptor() ([]byte, []int) {
	return file_slideService_proto_rawDescGZIP(), []int{2}
}

func (x *SlideResponse) GetEntity() *Slide {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *SlideResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *SlideResponse) GetItems() []*Slide {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *SlideResponse) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

var File_slideService_proto protoreflect.FileDescriptor

var file_slideService_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x6c, 0x69, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xbf, 0x03, 0x0a, 0x05, 0x53, 0x6c, 0x69, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x65, 0x73, 0x63, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63,
	0x12, 0x19, 0x0a, 0x08, 0x62, 0x67, 0x5f, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x62, 0x67, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x6f, 0x72, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x12, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x94, 0x02, 0x0a, 0x0c, 0x53, 0x6c, 0x69, 0x64, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70,
	0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69,
	0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e,
	0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x1d, 0x0a,
	0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18,
	0x0b, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x98, 0x01, 0x0a, 0x0d, 0x53,
	0x6c, 0x69, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x06,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x6c, 0x69, 0x64, 0x65, 0x52, 0x06, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61,
	0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x6c, 0x69, 0x64, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0xdb, 0x02, 0x0a, 0x0c, 0x53, 0x6c, 0x69, 0x64, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x6c, 0x69, 0x64,
	0x65, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x6c, 0x69,
	0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x06,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x53, 0x6c, 0x69, 0x64, 0x65, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x53, 0x6c, 0x69, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x34, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x6c, 0x69, 0x64, 0x65, 0x1a, 0x17, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x6c, 0x69, 0x64, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12,
	0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x6c, 0x69, 0x64, 0x65,
	0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x6c, 0x69, 0x64,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x06, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x53, 0x6c, 0x69, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x6c, 0x69, 0x64, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x6c, 0x69, 0x64,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x53, 0x6c, 0x69, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_slideService_proto_rawDescOnce sync.Once
	file_slideService_proto_rawDescData = file_slideService_proto_rawDesc
)

func file_slideService_proto_rawDescGZIP() []byte {
	file_slideService_proto_rawDescOnce.Do(func() {
		file_slideService_proto_rawDescData = protoimpl.X.CompressGZIP(file_slideService_proto_rawDescData)
	})
	return file_slideService_proto_rawDescData
}

var file_slideService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_slideService_proto_goTypes = []interface{}{
	(*Slide)(nil),         // 0: services.Slide
	(*SlideRequest)(nil),  // 1: services.SlideRequest
	(*SlideResponse)(nil), // 2: services.SlideResponse
	(*common.Pager)(nil),  // 3: common.Pager
}
var file_slideService_proto_depIdxs = []int32{
	0, // 0: services.SlideResponse.entity:type_name -> services.Slide
	3, // 1: services.SlideResponse.pager:type_name -> common.Pager
	0, // 2: services.SlideResponse.items:type_name -> services.Slide
	0, // 3: services.SlideService.Create:input_type -> services.Slide
	0, // 4: services.SlideService.Update:input_type -> services.Slide
	0, // 5: services.SlideService.Delete:input_type -> services.Slide
	0, // 6: services.SlideService.Get:input_type -> services.Slide
	1, // 7: services.SlideService.Search:input_type -> services.SlideRequest
	1, // 8: services.SlideService.List:input_type -> services.SlideRequest
	2, // 9: services.SlideService.Create:output_type -> services.SlideResponse
	2, // 10: services.SlideService.Update:output_type -> services.SlideResponse
	2, // 11: services.SlideService.Delete:output_type -> services.SlideResponse
	2, // 12: services.SlideService.Get:output_type -> services.SlideResponse
	2, // 13: services.SlideService.Search:output_type -> services.SlideResponse
	2, // 14: services.SlideService.List:output_type -> services.SlideResponse
	9, // [9:15] is the sub-list for method output_type
	3, // [3:9] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_slideService_proto_init() }
func file_slideService_proto_init() {
	if File_slideService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_slideService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Slide); i {
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
		file_slideService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SlideRequest); i {
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
		file_slideService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SlideResponse); i {
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
			RawDescriptor: file_slideService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_slideService_proto_goTypes,
		DependencyIndexes: file_slideService_proto_depIdxs,
		MessageInfos:      file_slideService_proto_msgTypes,
	}.Build()
	File_slideService_proto = out.File
	file_slideService_proto_rawDesc = nil
	file_slideService_proto_goTypes = nil
	file_slideService_proto_depIdxs = nil
}
