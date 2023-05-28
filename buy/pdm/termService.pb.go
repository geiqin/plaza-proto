// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: termService.proto

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

type Term struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Coding            string   `protobuf:"bytes,2,opt,name=coding,proto3" json:"coding"`
	Name              string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`
	ViceName          string   `protobuf:"bytes,4,opt,name=vice_name,json=viceName,proto3" json:"vice_name"`
	ParentId          int64    `protobuf:"varint,5,opt,name=parent_id,json=parentId,proto3" json:"parent_id"`
	Describe          string   `protobuf:"bytes,6,opt,name=describe,proto3" json:"describe"`
	IconUrl           string   `protobuf:"bytes,7,opt,name=icon_url,json=iconUrl,proto3" json:"icon_url"`
	ImageUrl          string   `protobuf:"bytes,8,opt,name=image_url,json=imageUrl,proto3" json:"image_url"`
	BgColor           string   `protobuf:"bytes,9,opt,name=bg_color,json=bgColor,proto3" json:"bg_color"`
	IsHomeRecommended string   `protobuf:"bytes,10,opt,name=is_home_recommended,json=isHomeRecommended,proto3" json:"is_home_recommended"`
	StyleType         string   `protobuf:"bytes,11,opt,name=style_type,json=styleType,proto3" json:"style_type"`
	ShowLimit         int32    `protobuf:"varint,12,opt,name=show_limit,json=showLimit,proto3" json:"show_limit"`
	SeoTitle          string   `protobuf:"bytes,13,opt,name=seo_title,json=seoTitle,proto3" json:"seo_title"`
	SeoKeywords       string   `protobuf:"bytes,14,opt,name=seo_keywords,json=seoKeywords,proto3" json:"seo_keywords"`
	SeoDesc           string   `protobuf:"bytes,15,opt,name=seo_desc,json=seoDesc,proto3" json:"seo_desc"`
	DeptPath          string   `protobuf:"bytes,16,opt,name=dept_path,json=deptPath,proto3" json:"dept_path"`
	Sort              int32    `protobuf:"varint,17,opt,name=sort,proto3" json:"sort"`
	Status            string   `protobuf:"bytes,18,opt,name=status,proto3" json:"status"`
	CreatedAt         string   `protobuf:"bytes,19,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt         string   `protobuf:"bytes,20,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	Parent            *Term    `protobuf:"bytes,21,opt,name=parent,proto3" json:"parent"`
	Children          []*Term  `protobuf:"bytes,22,rep,name=children,proto3" json:"children"`
	GoodsList         []*Goods `protobuf:"bytes,23,rep,name=goods_list,json=goodsList,proto3" json:"goods_list"`
}

func (x *Term) Reset() {
	*x = Term{}
	if protoimpl.UnsafeEnabled {
		mi := &file_termService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Term) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Term) ProtoMessage() {}

func (x *Term) ProtoReflect() protoreflect.Message {
	mi := &file_termService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Term.ProtoReflect.Descriptor instead.
func (*Term) Descriptor() ([]byte, []int) {
	return file_termService_proto_rawDescGZIP(), []int{0}
}

func (x *Term) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Term) GetCoding() string {
	if x != nil {
		return x.Coding
	}
	return ""
}

func (x *Term) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Term) GetViceName() string {
	if x != nil {
		return x.ViceName
	}
	return ""
}

func (x *Term) GetParentId() int64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *Term) GetDescribe() string {
	if x != nil {
		return x.Describe
	}
	return ""
}

func (x *Term) GetIconUrl() string {
	if x != nil {
		return x.IconUrl
	}
	return ""
}

func (x *Term) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *Term) GetBgColor() string {
	if x != nil {
		return x.BgColor
	}
	return ""
}

func (x *Term) GetIsHomeRecommended() string {
	if x != nil {
		return x.IsHomeRecommended
	}
	return ""
}

func (x *Term) GetStyleType() string {
	if x != nil {
		return x.StyleType
	}
	return ""
}

func (x *Term) GetShowLimit() int32 {
	if x != nil {
		return x.ShowLimit
	}
	return 0
}

func (x *Term) GetSeoTitle() string {
	if x != nil {
		return x.SeoTitle
	}
	return ""
}

func (x *Term) GetSeoKeywords() string {
	if x != nil {
		return x.SeoKeywords
	}
	return ""
}

func (x *Term) GetSeoDesc() string {
	if x != nil {
		return x.SeoDesc
	}
	return ""
}

func (x *Term) GetDeptPath() string {
	if x != nil {
		return x.DeptPath
	}
	return ""
}

func (x *Term) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *Term) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Term) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Term) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Term) GetParent() *Term {
	if x != nil {
		return x.Parent
	}
	return nil
}

func (x *Term) GetChildren() []*Term {
	if x != nil {
		return x.Children
	}
	return nil
}

func (x *Term) GetGoodsList() []*Goods {
	if x != nil {
		return x.GoodsList
	}
	return nil
}

type TermRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged             int64   `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize          int64   `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Keywords          string  `protobuf:"bytes,3,opt,name=keywords,proto3" json:"keywords"`
	Id                int64   `protobuf:"varint,4,opt,name=id,proto3" json:"id"`
	ParentId          int64   `protobuf:"varint,5,opt,name=parent_id,json=parentId,proto3" json:"parent_id"`
	Coding            string  `protobuf:"bytes,6,opt,name=coding,proto3" json:"coding"`
	Name              string  `protobuf:"bytes,7,opt,name=name,proto3" json:"name"`
	IsHomeRecommended string  `protobuf:"bytes,8,opt,name=is_home_recommended,json=isHomeRecommended,proto3" json:"is_home_recommended"`
	Status            string  `protobuf:"bytes,9,opt,name=status,proto3" json:"status"`
	Top               int32   `protobuf:"varint,10,opt,name=top,proto3" json:"top"`
	Ids               []int64 `protobuf:"varint,11,rep,packed,name=ids,proto3" json:"ids"`
}

func (x *TermRequest) Reset() {
	*x = TermRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_termService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TermRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TermRequest) ProtoMessage() {}

func (x *TermRequest) ProtoReflect() protoreflect.Message {
	mi := &file_termService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TermRequest.ProtoReflect.Descriptor instead.
func (*TermRequest) Descriptor() ([]byte, []int) {
	return file_termService_proto_rawDescGZIP(), []int{1}
}

func (x *TermRequest) GetPaged() int64 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *TermRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *TermRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *TermRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TermRequest) GetParentId() int64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *TermRequest) GetCoding() string {
	if x != nil {
		return x.Coding
	}
	return ""
}

func (x *TermRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TermRequest) GetIsHomeRecommended() string {
	if x != nil {
		return x.IsHomeRecommended
	}
	return ""
}

func (x *TermRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *TermRequest) GetTop() int32 {
	if x != nil {
		return x.Top
	}
	return 0
}

func (x *TermRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type TermData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Term         `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*Term       `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Info   *common.Info  `protobuf:"bytes,4,opt,name=info,proto3" json:"info"`
}

func (x *TermData) Reset() {
	*x = TermData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_termService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TermData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TermData) ProtoMessage() {}

func (x *TermData) ProtoReflect() protoreflect.Message {
	mi := &file_termService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TermData.ProtoReflect.Descriptor instead.
func (*TermData) Descriptor() ([]byte, []int) {
	return file_termService_proto_rawDescGZIP(), []int{2}
}

func (x *TermData) GetEntity() *Term {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *TermData) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *TermData) GetItems() []*Term {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *TermData) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type TermResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *TermData     `protobuf:"bytes,1,opt,name=data,proto3" json:"data"`
	Error *common.Error `protobuf:"bytes,2,opt,name=error,proto3" json:"error"`
}

func (x *TermResponse) Reset() {
	*x = TermResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_termService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TermResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TermResponse) ProtoMessage() {}

func (x *TermResponse) ProtoReflect() protoreflect.Message {
	mi := &file_termService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TermResponse.ProtoReflect.Descriptor instead.
func (*TermResponse) Descriptor() ([]byte, []int) {
	return file_termService_proto_rawDescGZIP(), []int{3}
}

func (x *TermResponse) GetData() *TermData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *TermResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_termService_proto protoreflect.FileDescriptor

var file_termService_proto_rawDesc = []byte{
	0x0a, 0x11, 0x74, 0x65, 0x72, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x12, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbf, 0x05, 0x0a, 0x04, 0x54, 0x65, 0x72, 0x6d, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63,
	0x6f, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x76, 0x69,
	0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x69, 0x63, 0x6f, 0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x69, 0x63, 0x6f, 0x6e, 0x55, 0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x67, 0x5f, 0x63, 0x6f,
	0x6c, 0x6f, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x67, 0x43, 0x6f, 0x6c,
	0x6f, 0x72, 0x12, 0x2e, 0x0a, 0x13, 0x69, 0x73, 0x5f, 0x68, 0x6f, 0x6d, 0x65, 0x5f, 0x72, 0x65,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x11, 0x69, 0x73, 0x48, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64,
	0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x68, 0x6f, 0x77, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x68, 0x6f, 0x77, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x6f, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6f, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x73, 0x65, 0x6f, 0x5f, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x6f, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73,
	0x12, 0x19, 0x0a, 0x08, 0x73, 0x65, 0x6f, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x6f, 0x44, 0x65, 0x73, 0x63, 0x12, 0x1b, 0x0a, 0x09, 0x64,
	0x65, 0x70, 0x74, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x64, 0x65, 0x70, 0x74, 0x50, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74,
	0x18, 0x11, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x26, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x15, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x65,
	0x72, 0x6d, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x08, 0x63, 0x68,
	0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18, 0x16, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x65, 0x72, 0x6d, 0x52, 0x08, 0x63, 0x68,
	0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x12, 0x2e, 0x0a, 0x0a, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x17, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x09, 0x67, 0x6f, 0x6f,
	0x64, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x22, 0xa1, 0x02, 0x0a, 0x0b, 0x54, 0x65, 0x72, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79,
	0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79,
	0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2e,
	0x0a, 0x13, 0x69, 0x73, 0x5f, 0x68, 0x6f, 0x6d, 0x65, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x64, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x69, 0x73, 0x48,
	0x6f, 0x6d, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x6f, 0x70, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x74, 0x6f, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18,
	0x0b, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x9f, 0x01, 0x0a, 0x08, 0x54,
	0x65, 0x72, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x12, 0x26, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x54, 0x65, 0x72, 0x6d, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12,
	0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70,
	0x61, 0x67, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54,
	0x65, 0x72, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x5b, 0x0a, 0x0c,
	0x54, 0x65, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x65, 0x72, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x82, 0x03, 0x0a, 0x0b, 0x54, 0x65,
	0x72, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54,
	0x65, 0x72, 0x6d, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54,
	0x65, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x54, 0x65, 0x72, 0x6d, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x54, 0x65, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a,
	0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x54, 0x65, 0x72, 0x6d, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x54, 0x65, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2d, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x54, 0x65, 0x72, 0x6d, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x54, 0x65, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35,
	0x0a, 0x04, 0x54, 0x72, 0x65, 0x65, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x54, 0x65, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x65, 0x72, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12,
	0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x65, 0x72, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x54, 0x65, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e,
	0x0a, 0x0d, 0x48, 0x6f, 0x6d, 0x65, 0x46, 0x6c, 0x6f, 0x6f, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x65, 0x72, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x54, 0x65, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0d,
	0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_termService_proto_rawDescOnce sync.Once
	file_termService_proto_rawDescData = file_termService_proto_rawDesc
)

func file_termService_proto_rawDescGZIP() []byte {
	file_termService_proto_rawDescOnce.Do(func() {
		file_termService_proto_rawDescData = protoimpl.X.CompressGZIP(file_termService_proto_rawDescData)
	})
	return file_termService_proto_rawDescData
}

var file_termService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_termService_proto_goTypes = []interface{}{
	(*Term)(nil),         // 0: services.Term
	(*TermRequest)(nil),  // 1: services.TermRequest
	(*TermData)(nil),     // 2: services.TermData
	(*TermResponse)(nil), // 3: services.TermResponse
	(*Goods)(nil),        // 4: services.Goods
	(*common.Pager)(nil), // 5: common.Pager
	(*common.Info)(nil),  // 6: common.Info
	(*common.Error)(nil), // 7: common.Error
}
var file_termService_proto_depIdxs = []int32{
	0,  // 0: services.Term.parent:type_name -> services.Term
	0,  // 1: services.Term.children:type_name -> services.Term
	4,  // 2: services.Term.goods_list:type_name -> services.Goods
	0,  // 3: services.TermData.entity:type_name -> services.Term
	5,  // 4: services.TermData.pager:type_name -> common.Pager
	0,  // 5: services.TermData.items:type_name -> services.Term
	6,  // 6: services.TermData.info:type_name -> common.Info
	2,  // 7: services.TermResponse.data:type_name -> services.TermData
	7,  // 8: services.TermResponse.error:type_name -> common.Error
	0,  // 9: services.TermService.Create:input_type -> services.Term
	0,  // 10: services.TermService.Update:input_type -> services.Term
	0,  // 11: services.TermService.Delete:input_type -> services.Term
	0,  // 12: services.TermService.Get:input_type -> services.Term
	1,  // 13: services.TermService.Tree:input_type -> services.TermRequest
	1,  // 14: services.TermService.Search:input_type -> services.TermRequest
	1,  // 15: services.TermService.HomeFloorData:input_type -> services.TermRequest
	3,  // 16: services.TermService.Create:output_type -> services.TermResponse
	3,  // 17: services.TermService.Update:output_type -> services.TermResponse
	3,  // 18: services.TermService.Delete:output_type -> services.TermResponse
	3,  // 19: services.TermService.Get:output_type -> services.TermResponse
	3,  // 20: services.TermService.Tree:output_type -> services.TermResponse
	3,  // 21: services.TermService.Search:output_type -> services.TermResponse
	3,  // 22: services.TermService.HomeFloorData:output_type -> services.TermResponse
	16, // [16:23] is the sub-list for method output_type
	9,  // [9:16] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_termService_proto_init() }
func file_termService_proto_init() {
	if File_termService_proto != nil {
		return
	}
	file_goodsService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_termService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Term); i {
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
		file_termService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TermRequest); i {
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
		file_termService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TermData); i {
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
		file_termService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TermResponse); i {
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
			RawDescriptor: file_termService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_termService_proto_goTypes,
		DependencyIndexes: file_termService_proto_depIdxs,
		MessageInfos:      file_termService_proto_msgTypes,
	}.Build()
	File_termService_proto = out.File
	file_termService_proto_rawDesc = nil
	file_termService_proto_goTypes = nil
	file_termService_proto_depIdxs = nil
}
