// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: designService.proto

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

type Design struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Name         string `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`
	LogoId       int64  `protobuf:"varint,4,opt,name=logo_id,json=logoId,proto3" json:"logo_id"`
	LogoUrl      string `protobuf:"bytes,5,opt,name=logo_url,json=logoUrl,proto3" json:"logo_url"`
	Content      string `protobuf:"bytes,6,opt,name=content,proto3" json:"content"`
	IsFooter     string `protobuf:"bytes,7,opt,name=is_footer,json=isFooter,proto3" json:"is_footer"`
	IsHeader     string `protobuf:"bytes,8,opt,name=is_header,json=isHeader,proto3" json:"is_header"`
	IsFullScreen string `protobuf:"bytes,9,opt,name=is_full_screen,json=isFullScreen,proto3" json:"is_full_screen"`
	SeoTitle     string `protobuf:"bytes,10,opt,name=seo_title,json=seoTitle,proto3" json:"seo_title"`
	SeoKeywords  string `protobuf:"bytes,11,opt,name=seo_keywords,json=seoKeywords,proto3" json:"seo_keywords"`
	SeoDesc      string `protobuf:"bytes,12,opt,name=seo_desc,json=seoDesc,proto3" json:"seo_desc"`
	AccessCount  int64  `protobuf:"varint,13,opt,name=access_count,json=accessCount,proto3" json:"access_count"`
	Status       string `protobuf:"bytes,14,opt,name=status,proto3" json:"status"`
	CreatedAt    string `protobuf:"bytes,15,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt    string `protobuf:"bytes,16,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
}

func (x *Design) Reset() {
	*x = Design{}
	if protoimpl.UnsafeEnabled {
		mi := &file_designService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Design) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Design) ProtoMessage() {}

func (x *Design) ProtoReflect() protoreflect.Message {
	mi := &file_designService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Design.ProtoReflect.Descriptor instead.
func (*Design) Descriptor() ([]byte, []int) {
	return file_designService_proto_rawDescGZIP(), []int{0}
}

func (x *Design) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Design) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Design) GetLogoId() int64 {
	if x != nil {
		return x.LogoId
	}
	return 0
}

func (x *Design) GetLogoUrl() string {
	if x != nil {
		return x.LogoUrl
	}
	return ""
}

func (x *Design) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Design) GetIsFooter() string {
	if x != nil {
		return x.IsFooter
	}
	return ""
}

func (x *Design) GetIsHeader() string {
	if x != nil {
		return x.IsHeader
	}
	return ""
}

func (x *Design) GetIsFullScreen() string {
	if x != nil {
		return x.IsFullScreen
	}
	return ""
}

func (x *Design) GetSeoTitle() string {
	if x != nil {
		return x.SeoTitle
	}
	return ""
}

func (x *Design) GetSeoKeywords() string {
	if x != nil {
		return x.SeoKeywords
	}
	return ""
}

func (x *Design) GetSeoDesc() string {
	if x != nil {
		return x.SeoDesc
	}
	return ""
}

func (x *Design) GetAccessCount() int64 {
	if x != nil {
		return x.AccessCount
	}
	return 0
}

func (x *Design) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Design) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Design) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type DesignRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int32 `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Sorting  int32 `protobuf:"varint,3,opt,name=sorting,proto3" json:"sorting"`
	//base params
	Id       int64   `protobuf:"varint,4,opt,name=id,proto3" json:"id"`
	Name     string  `protobuf:"bytes,5,opt,name=name,proto3" json:"name"`
	Status   string  `protobuf:"bytes,6,opt,name=status,proto3" json:"status"`
	ParentId int64   `protobuf:"varint,7,opt,name=parent_id,json=parentId,proto3" json:"parent_id"`
	Keywords string  `protobuf:"bytes,8,opt,name=keywords,proto3" json:"keywords"`
	Ids      []int64 `protobuf:"varint,10,rep,packed,name=ids,proto3" json:"ids"`
}

func (x *DesignRequest) Reset() {
	*x = DesignRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_designService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DesignRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DesignRequest) ProtoMessage() {}

func (x *DesignRequest) ProtoReflect() protoreflect.Message {
	mi := &file_designService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DesignRequest.ProtoReflect.Descriptor instead.
func (*DesignRequest) Descriptor() ([]byte, []int) {
	return file_designService_proto_rawDescGZIP(), []int{1}
}

func (x *DesignRequest) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *DesignRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *DesignRequest) GetSorting() int32 {
	if x != nil {
		return x.Sorting
	}
	return 0
}

func (x *DesignRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DesignRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DesignRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *DesignRequest) GetParentId() int64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *DesignRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *DesignRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

//
type DesignData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Design       `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*Design     `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Info   *common.Info  `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
}

func (x *DesignData) Reset() {
	*x = DesignData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_designService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DesignData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DesignData) ProtoMessage() {}

func (x *DesignData) ProtoReflect() protoreflect.Message {
	mi := &file_designService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DesignData.ProtoReflect.Descriptor instead.
func (*DesignData) Descriptor() ([]byte, []int) {
	return file_designService_proto_rawDescGZIP(), []int{2}
}

func (x *DesignData) GetEntity() *Design {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *DesignData) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *DesignData) GetItems() []*Design {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *DesignData) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

//
type DesignResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *DesignData   `protobuf:"bytes,1,opt,name=data,proto3" json:"data"`
	Error *common.Error `protobuf:"bytes,2,opt,name=error,proto3" json:"error"`
}

func (x *DesignResponse) Reset() {
	*x = DesignResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_designService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DesignResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DesignResponse) ProtoMessage() {}

func (x *DesignResponse) ProtoReflect() protoreflect.Message {
	mi := &file_designService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DesignResponse.ProtoReflect.Descriptor instead.
func (*DesignResponse) Descriptor() ([]byte, []int) {
	return file_designService_proto_rawDescGZIP(), []int{3}
}

func (x *DesignResponse) GetData() *DesignData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *DesignResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_designService_proto protoreflect.FileDescriptor

var file_designService_proto_rawDesc = []byte{
	0x0a, 0x13, 0x64, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a,
	0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xae, 0x03, 0x0a, 0x06, 0x44, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x6f, 0x67, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x6c, 0x6f, 0x67, 0x6f, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f,
	0x67, 0x6f, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f,
	0x67, 0x6f, 0x55, 0x72, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x66, 0x6f, 0x6f, 0x74, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x69, 0x73, 0x46, 0x6f, 0x6f, 0x74, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09,
	0x69, 0x73, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x69, 0x73, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x0e, 0x69, 0x73, 0x5f,
	0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x69, 0x73, 0x46, 0x75, 0x6c, 0x6c, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x12,
	0x1b, 0x0a, 0x09, 0x73, 0x65, 0x6f, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6f, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x73, 0x65, 0x6f, 0x5f, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x6f, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12,
	0x19, 0x0a, 0x08, 0x73, 0x65, 0x6f, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x73, 0x65, 0x6f, 0x44, 0x65, 0x73, 0x63, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0xe3, 0x01, 0x0a, 0x0d, 0x44, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74,
	0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69,
	0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6b,
	0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b,
	0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x0a,
	0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0xa5, 0x01, 0x0a, 0x0a, 0x44, 0x65,
	0x73, 0x69, 0x67, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x28, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72,
	0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x44, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12,
	0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66,
	0x6f, 0x22, 0x5f, 0x0a, 0x0e, 0x44, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x73,
	0x69, 0x67, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x23, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x32, 0xe8, 0x02, 0x0a, 0x0d, 0x44, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x10,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x73, 0x69, 0x67, 0x6e,
	0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x73, 0x69,
	0x67, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x06,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x44, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x10,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x73, 0x69, 0x67, 0x6e,
	0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x73, 0x69,
	0x67, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x03,
	0x47, 0x65, 0x74, 0x12, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44,
	0x65, 0x73, 0x69, 0x67, 0x6e, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x44, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x3d, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x17, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x44, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x3b, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x73,
	0x69, 0x67, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a,
	0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_designService_proto_rawDescOnce sync.Once
	file_designService_proto_rawDescData = file_designService_proto_rawDesc
)

func file_designService_proto_rawDescGZIP() []byte {
	file_designService_proto_rawDescOnce.Do(func() {
		file_designService_proto_rawDescData = protoimpl.X.CompressGZIP(file_designService_proto_rawDescData)
	})
	return file_designService_proto_rawDescData
}

var file_designService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_designService_proto_goTypes = []interface{}{
	(*Design)(nil),         // 0: services.Design
	(*DesignRequest)(nil),  // 1: services.DesignRequest
	(*DesignData)(nil),     // 2: services.DesignData
	(*DesignResponse)(nil), // 3: services.DesignResponse
	(*common.Pager)(nil),   // 4: common.Pager
	(*common.Info)(nil),    // 5: common.Info
	(*common.Error)(nil),   // 6: common.Error
}
var file_designService_proto_depIdxs = []int32{
	0,  // 0: services.DesignData.entity:type_name -> services.Design
	4,  // 1: services.DesignData.pager:type_name -> common.Pager
	0,  // 2: services.DesignData.items:type_name -> services.Design
	5,  // 3: services.DesignData.info:type_name -> common.Info
	2,  // 4: services.DesignResponse.data:type_name -> services.DesignData
	6,  // 5: services.DesignResponse.error:type_name -> common.Error
	0,  // 6: services.DesignService.Create:input_type -> services.Design
	0,  // 7: services.DesignService.Update:input_type -> services.Design
	0,  // 8: services.DesignService.Delete:input_type -> services.Design
	0,  // 9: services.DesignService.Get:input_type -> services.Design
	1,  // 10: services.DesignService.Search:input_type -> services.DesignRequest
	1,  // 11: services.DesignService.List:input_type -> services.DesignRequest
	3,  // 12: services.DesignService.Create:output_type -> services.DesignResponse
	3,  // 13: services.DesignService.Update:output_type -> services.DesignResponse
	3,  // 14: services.DesignService.Delete:output_type -> services.DesignResponse
	3,  // 15: services.DesignService.Get:output_type -> services.DesignResponse
	3,  // 16: services.DesignService.Search:output_type -> services.DesignResponse
	3,  // 17: services.DesignService.List:output_type -> services.DesignResponse
	12, // [12:18] is the sub-list for method output_type
	6,  // [6:12] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_designService_proto_init() }
func file_designService_proto_init() {
	if File_designService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_designService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Design); i {
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
		file_designService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DesignRequest); i {
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
		file_designService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DesignData); i {
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
		file_designService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DesignResponse); i {
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
			RawDescriptor: file_designService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_designService_proto_goTypes,
		DependencyIndexes: file_designService_proto_depIdxs,
		MessageInfos:      file_designService_proto_msgTypes,
	}.Build()
	File_designService_proto = out.File
	file_designService_proto_rawDesc = nil
	file_designService_proto_goTypes = nil
	file_designService_proto_depIdxs = nil
}
