// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: mainNavService.proto

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

type MainNav struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`                                        //
	Type        string `protobuf:"bytes,2,opt,name=type,proto3" json:"type"`                                     //类型
	Platform    string `protobuf:"bytes,3,opt,name=platform,proto3" json:"platform"`                             //所属平台
	Name        string `protobuf:"bytes,4,opt,name=name,proto3" json:"name"`                                     //名称
	EventType   string `protobuf:"bytes,5,opt,name=event_type,json=eventType,proto3" json:"event_type"`          //事件类型
	EventValue  string `protobuf:"bytes,6,opt,name=event_value,json=eventValue,proto3" json:"event_value"`       //事件值
	ImageUrl    string `protobuf:"bytes,7,opt,name=image_url,json=imageUrl,proto3" json:"image_url"`             //图片地址
	Desc        string `protobuf:"bytes,8,opt,name=desc,proto3" json:"desc"`                                     //描述
	BgColor     string `protobuf:"bytes,9,opt,name=bg_color,json=bgColor,proto3" json:"bg_color"`                //背景色值
	IsNeedLogin string `protobuf:"bytes,10,opt,name=is_need_login,json=isNeedLogin,proto3" json:"is_need_login"` //是否需要登录
	Sort        int32  `protobuf:"varint,11,opt,name=sort,proto3" json:"sort"`                                   //排序
	Status      string `protobuf:"bytes,12,opt,name=status,proto3" json:"status"`                                //状态
	CreatedAt   int64  `protobuf:"varint,13,opt,name=created_at,json=createdAt,proto3" json:"created_at"`        //创建时间
	UpdatedAt   int64  `protobuf:"varint,14,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`        //修改时间
	//extends
	PlatformName  string `protobuf:"bytes,17,opt,name=platform_name,json=platformName,proto3" json:"platform_name"`
	EventTypeName string `protobuf:"bytes,18,opt,name=event_type_name,json=eventTypeName,proto3" json:"event_type_name"`
	StatusName    string `protobuf:"bytes,19,opt,name=status_name,json=statusName,proto3" json:"status_name"`
}

func (x *MainNav) Reset() {
	*x = MainNav{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mainNavService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MainNav) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MainNav) ProtoMessage() {}

func (x *MainNav) ProtoReflect() protoreflect.Message {
	mi := &file_mainNavService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MainNav.ProtoReflect.Descriptor instead.
func (*MainNav) Descriptor() ([]byte, []int) {
	return file_mainNavService_proto_rawDescGZIP(), []int{0}
}

func (x *MainNav) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MainNav) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *MainNav) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *MainNav) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MainNav) GetEventType() string {
	if x != nil {
		return x.EventType
	}
	return ""
}

func (x *MainNav) GetEventValue() string {
	if x != nil {
		return x.EventValue
	}
	return ""
}

func (x *MainNav) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *MainNav) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *MainNav) GetBgColor() string {
	if x != nil {
		return x.BgColor
	}
	return ""
}

func (x *MainNav) GetIsNeedLogin() string {
	if x != nil {
		return x.IsNeedLogin
	}
	return ""
}

func (x *MainNav) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *MainNav) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *MainNav) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *MainNav) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *MainNav) GetPlatformName() string {
	if x != nil {
		return x.PlatformName
	}
	return ""
}

func (x *MainNav) GetEventTypeName() string {
	if x != nil {
		return x.EventTypeName
	}
	return ""
}

func (x *MainNav) GetStatusName() string {
	if x != nil {
		return x.StatusName
	}
	return ""
}

type MainNavRequest struct {
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
	Type        string `protobuf:"bytes,11,opt,name=type,proto3" json:"type"`                                    //类型
	Platform    string `protobuf:"bytes,12,opt,name=platform,proto3" json:"platform"`                            //所属平台
	Name        string `protobuf:"bytes,13,opt,name=name,proto3" json:"name"`                                    //名称
	EventType   string `protobuf:"bytes,14,opt,name=event_type,json=eventType,proto3" json:"event_type"`         //事件类型
	IsNeedLogin string `protobuf:"bytes,15,opt,name=is_need_login,json=isNeedLogin,proto3" json:"is_need_login"` //是否需要登录
	Status      string `protobuf:"bytes,16,opt,name=status,proto3" json:"status"`                                //状态
}

func (x *MainNavRequest) Reset() {
	*x = MainNavRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mainNavService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MainNavRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MainNavRequest) ProtoMessage() {}

func (x *MainNavRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mainNavService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MainNavRequest.ProtoReflect.Descriptor instead.
func (*MainNavRequest) Descriptor() ([]byte, []int) {
	return file_mainNavService_proto_rawDescGZIP(), []int{1}
}

func (x *MainNavRequest) GetTop() int32 {
	if x != nil {
		return x.Top
	}
	return 0
}

func (x *MainNavRequest) GetPaged() int64 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *MainNavRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *MainNavRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *MainNavRequest) GetSort() []string {
	if x != nil {
		return x.Sort
	}
	return nil
}

func (x *MainNavRequest) GetDateRange() []string {
	if x != nil {
		return x.DateRange
	}
	return nil
}

func (x *MainNavRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *MainNavRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MainNavRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *MainNavRequest) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *MainNavRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MainNavRequest) GetEventType() string {
	if x != nil {
		return x.EventType
	}
	return ""
}

func (x *MainNavRequest) GetIsNeedLogin() string {
	if x != nil {
		return x.IsNeedLogin
	}
	return ""
}

func (x *MainNavRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type MainNavResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *MainNav      `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*MainNav    `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Msg    string        `protobuf:"bytes,5,opt,name=msg,proto3" json:"msg"`
}

func (x *MainNavResponse) Reset() {
	*x = MainNavResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mainNavService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MainNavResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MainNavResponse) ProtoMessage() {}

func (x *MainNavResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mainNavService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MainNavResponse.ProtoReflect.Descriptor instead.
func (*MainNavResponse) Descriptor() ([]byte, []int) {
	return file_mainNavService_proto_rawDescGZIP(), []int{2}
}

func (x *MainNavResponse) GetEntity() *MainNav {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *MainNavResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *MainNavResponse) GetItems() []*MainNav {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *MainNavResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_mainNavService_proto protoreflect.FileDescriptor

var file_mainNavService_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6d, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x76, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xe5, 0x03, 0x0a, 0x07, 0x4d, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x76, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x64, 0x65, 0x73, 0x63, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x67, 0x5f, 0x63, 0x6f, 0x6c, 0x6f, 0x72,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x67, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12,
	0x22, 0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x6e, 0x65, 0x65, 0x64, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x73, 0x4e, 0x65, 0x65, 0x64, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x23, 0x0a,
	0x0d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x11,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xe5, 0x02, 0x0a, 0x0e,
	0x4d, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x76, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x74, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x74, 0x6f, 0x70,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x73,
	0x6f, 0x72, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x61, 0x6e, 0x67,
	0x65, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x64, 0x61, 0x74, 0x65, 0x52, 0x61, 0x6e,
	0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x03, 0x52,
	0x03, 0x69, 0x64, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x6e, 0x65,
	0x65, 0x64, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x69, 0x73, 0x4e, 0x65, 0x65, 0x64, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x9c, 0x01, 0x0a, 0x0f, 0x4d, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x76, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x4d, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x76, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72,
	0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x4d, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x76, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d,
	0x73, 0x67, 0x32, 0xf5, 0x02, 0x0a, 0x0e, 0x4d, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x76, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x69, 0x6e, 0x4e,
	0x61, 0x76, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x61,
	0x69, 0x6e, 0x4e, 0x61, 0x76, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x38, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x76, 0x1a, 0x19, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x76, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x06, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d,
	0x61, 0x69, 0x6e, 0x4e, 0x61, 0x76, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x4d, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x76, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x11, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x76, 0x1a, 0x19, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x76,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x06, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x4d, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x76, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x69, 0x6e, 0x4e, 0x61,
	0x76, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x04, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d,
	0x61, 0x69, 0x6e, 0x4e, 0x61, 0x76, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x76,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e,
	0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_mainNavService_proto_rawDescOnce sync.Once
	file_mainNavService_proto_rawDescData = file_mainNavService_proto_rawDesc
)

func file_mainNavService_proto_rawDescGZIP() []byte {
	file_mainNavService_proto_rawDescOnce.Do(func() {
		file_mainNavService_proto_rawDescData = protoimpl.X.CompressGZIP(file_mainNavService_proto_rawDescData)
	})
	return file_mainNavService_proto_rawDescData
}

var file_mainNavService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_mainNavService_proto_goTypes = []interface{}{
	(*MainNav)(nil),         // 0: services.MainNav
	(*MainNavRequest)(nil),  // 1: services.MainNavRequest
	(*MainNavResponse)(nil), // 2: services.MainNavResponse
	(*common.Pager)(nil),    // 3: common.Pager
}
var file_mainNavService_proto_depIdxs = []int32{
	0, // 0: services.MainNavResponse.entity:type_name -> services.MainNav
	3, // 1: services.MainNavResponse.pager:type_name -> common.Pager
	0, // 2: services.MainNavResponse.items:type_name -> services.MainNav
	0, // 3: services.MainNavService.Create:input_type -> services.MainNav
	0, // 4: services.MainNavService.Update:input_type -> services.MainNav
	0, // 5: services.MainNavService.Delete:input_type -> services.MainNav
	0, // 6: services.MainNavService.Get:input_type -> services.MainNav
	1, // 7: services.MainNavService.Search:input_type -> services.MainNavRequest
	1, // 8: services.MainNavService.List:input_type -> services.MainNavRequest
	2, // 9: services.MainNavService.Create:output_type -> services.MainNavResponse
	2, // 10: services.MainNavService.Update:output_type -> services.MainNavResponse
	2, // 11: services.MainNavService.Delete:output_type -> services.MainNavResponse
	2, // 12: services.MainNavService.Get:output_type -> services.MainNavResponse
	2, // 13: services.MainNavService.Search:output_type -> services.MainNavResponse
	2, // 14: services.MainNavService.List:output_type -> services.MainNavResponse
	9, // [9:15] is the sub-list for method output_type
	3, // [3:9] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_mainNavService_proto_init() }
func file_mainNavService_proto_init() {
	if File_mainNavService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mainNavService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MainNav); i {
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
		file_mainNavService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MainNavRequest); i {
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
		file_mainNavService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MainNavResponse); i {
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
			RawDescriptor: file_mainNavService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mainNavService_proto_goTypes,
		DependencyIndexes: file_mainNavService_proto_depIdxs,
		MessageInfos:      file_mainNavService_proto_msgTypes,
	}.Build()
	File_mainNavService_proto = out.File
	file_mainNavService_proto_rawDesc = nil
	file_mainNavService_proto_goTypes = nil
	file_mainNavService_proto_depIdxs = nil
}
