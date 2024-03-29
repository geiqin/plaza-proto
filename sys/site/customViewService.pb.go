// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: customViewService.proto

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

type CustomView struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Title        string `protobuf:"bytes,2,opt,name=title,proto3" json:"title"`
	Content      string `protobuf:"bytes,3,opt,name=content,proto3" json:"content"`
	IsFooter     string `protobuf:"bytes,4,opt,name=is_footer,json=isFooter,proto3" json:"is_footer"`
	IsHeader     string `protobuf:"bytes,5,opt,name=is_header,json=isHeader,proto3" json:"is_header"`
	IsFullScreen string `protobuf:"bytes,6,opt,name=is_full_screen,json=isFullScreen,proto3" json:"is_full_screen"`
	Images       string `protobuf:"bytes,7,opt,name=images,proto3" json:"images"`
	ImagesCount  int32  `protobuf:"varint,8,opt,name=images_count,json=imagesCount,proto3" json:"images_count"`
	AccessCount  int32  `protobuf:"varint,9,opt,name=access_count,json=accessCount,proto3" json:"access_count"`
	Status       string `protobuf:"bytes,10,opt,name=status,proto3" json:"status"`
	CreatedAt    string `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt    string `protobuf:"bytes,12,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
}

func (x *CustomView) Reset() {
	*x = CustomView{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customViewService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomView) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomView) ProtoMessage() {}

func (x *CustomView) ProtoReflect() protoreflect.Message {
	mi := &file_customViewService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomView.ProtoReflect.Descriptor instead.
func (*CustomView) Descriptor() ([]byte, []int) {
	return file_customViewService_proto_rawDescGZIP(), []int{0}
}

func (x *CustomView) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CustomView) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CustomView) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CustomView) GetIsFooter() string {
	if x != nil {
		return x.IsFooter
	}
	return ""
}

func (x *CustomView) GetIsHeader() string {
	if x != nil {
		return x.IsHeader
	}
	return ""
}

func (x *CustomView) GetIsFullScreen() string {
	if x != nil {
		return x.IsFullScreen
	}
	return ""
}

func (x *CustomView) GetImages() string {
	if x != nil {
		return x.Images
	}
	return ""
}

func (x *CustomView) GetImagesCount() int32 {
	if x != nil {
		return x.ImagesCount
	}
	return 0
}

func (x *CustomView) GetAccessCount() int32 {
	if x != nil {
		return x.AccessCount
	}
	return 0
}

func (x *CustomView) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *CustomView) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *CustomView) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type CustomViewRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int64  `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize int64  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Sorting  string `protobuf:"bytes,3,opt,name=sorting,proto3" json:"sorting"`
	//base params
	Id       int64   `protobuf:"varint,4,opt,name=id,proto3" json:"id"`
	Name     string  `protobuf:"bytes,5,opt,name=name,proto3" json:"name"`
	Code     string  `protobuf:"bytes,6,opt,name=code,proto3" json:"code"`
	ParentId int64   `protobuf:"varint,7,opt,name=parent_id,json=parentId,proto3" json:"parent_id"`
	Keywords string  `protobuf:"bytes,8,opt,name=keywords,proto3" json:"keywords"`
	Ids      []int64 `protobuf:"varint,10,rep,packed,name=ids,proto3" json:"ids"`
}

func (x *CustomViewRequest) Reset() {
	*x = CustomViewRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customViewService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomViewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomViewRequest) ProtoMessage() {}

func (x *CustomViewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_customViewService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomViewRequest.ProtoReflect.Descriptor instead.
func (*CustomViewRequest) Descriptor() ([]byte, []int) {
	return file_customViewService_proto_rawDescGZIP(), []int{1}
}

func (x *CustomViewRequest) GetPaged() int64 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *CustomViewRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *CustomViewRequest) GetSorting() string {
	if x != nil {
		return x.Sorting
	}
	return ""
}

func (x *CustomViewRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CustomViewRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CustomViewRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *CustomViewRequest) GetParentId() int64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *CustomViewRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *CustomViewRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type CustomViewResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *CustomView   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*CustomView `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Info   string        `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
}

func (x *CustomViewResponse) Reset() {
	*x = CustomViewResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customViewService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomViewResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomViewResponse) ProtoMessage() {}

func (x *CustomViewResponse) ProtoReflect() protoreflect.Message {
	mi := &file_customViewService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomViewResponse.ProtoReflect.Descriptor instead.
func (*CustomViewResponse) Descriptor() ([]byte, []int) {
	return file_customViewService_proto_rawDescGZIP(), []int{2}
}

func (x *CustomViewResponse) GetEntity() *CustomView {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *CustomViewResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *CustomViewResponse) GetItems() []*CustomView {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *CustomViewResponse) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

var File_customViewService_proto protoreflect.FileDescriptor

var file_customViewService_proto_rawDesc = []byte{
	0x0a, 0x17, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x56, 0x69, 0x65, 0x77, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe0, 0x02, 0x0a, 0x0a, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x56, 0x69, 0x65, 0x77, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x66, 0x6f, 0x6f, 0x74,
	0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x73, 0x46, 0x6f, 0x6f, 0x74,
	0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x73, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12,
	0x24, 0x0a, 0x0e, 0x69, 0x73, 0x5f, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x73, 0x63, 0x72, 0x65, 0x65,
	0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x73, 0x46, 0x75, 0x6c, 0x6c, 0x53,
	0x63, 0x72, 0x65, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x21, 0x0a,
	0x0c, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0b, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xe3, 0x01, 0x0a, 0x11, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x56, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x10, 0x0a,
	0x03, 0x69, 0x64, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22,
	0xa7, 0x01, 0x0a, 0x12, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x56, 0x69, 0x65, 0x77, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x56, 0x69, 0x65, 0x77, 0x52, 0x06, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67,
	0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x56, 0x69, 0x65, 0x77, 0x52, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0x9c, 0x03, 0x0a, 0x11, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x56, 0x69, 0x65, 0x77, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x3e, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x56, 0x69, 0x65, 0x77, 0x1a,
	0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x56, 0x69, 0x65, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x3e, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x56, 0x69, 0x65, 0x77, 0x1a,
	0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x56, 0x69, 0x65, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x3e, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x56, 0x69, 0x65, 0x77, 0x1a,
	0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x56, 0x69, 0x65, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x3b, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x56, 0x69, 0x65, 0x77, 0x1a, 0x1c, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x56, 0x69,
	0x65, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x06,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x56, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x56, 0x69, 0x65, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x56, 0x69, 0x65,
	0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x56, 0x69, 0x65, 0x77, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_customViewService_proto_rawDescOnce sync.Once
	file_customViewService_proto_rawDescData = file_customViewService_proto_rawDesc
)

func file_customViewService_proto_rawDescGZIP() []byte {
	file_customViewService_proto_rawDescOnce.Do(func() {
		file_customViewService_proto_rawDescData = protoimpl.X.CompressGZIP(file_customViewService_proto_rawDescData)
	})
	return file_customViewService_proto_rawDescData
}

var file_customViewService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_customViewService_proto_goTypes = []interface{}{
	(*CustomView)(nil),         // 0: services.CustomView
	(*CustomViewRequest)(nil),  // 1: services.CustomViewRequest
	(*CustomViewResponse)(nil), // 2: services.CustomViewResponse
	(*common.Pager)(nil),       // 3: common.Pager
}
var file_customViewService_proto_depIdxs = []int32{
	0, // 0: services.CustomViewResponse.entity:type_name -> services.CustomView
	3, // 1: services.CustomViewResponse.pager:type_name -> common.Pager
	0, // 2: services.CustomViewResponse.items:type_name -> services.CustomView
	0, // 3: services.CustomViewService.Create:input_type -> services.CustomView
	0, // 4: services.CustomViewService.Update:input_type -> services.CustomView
	0, // 5: services.CustomViewService.Delete:input_type -> services.CustomView
	0, // 6: services.CustomViewService.Get:input_type -> services.CustomView
	1, // 7: services.CustomViewService.Search:input_type -> services.CustomViewRequest
	1, // 8: services.CustomViewService.List:input_type -> services.CustomViewRequest
	2, // 9: services.CustomViewService.Create:output_type -> services.CustomViewResponse
	2, // 10: services.CustomViewService.Update:output_type -> services.CustomViewResponse
	2, // 11: services.CustomViewService.Delete:output_type -> services.CustomViewResponse
	2, // 12: services.CustomViewService.Get:output_type -> services.CustomViewResponse
	2, // 13: services.CustomViewService.Search:output_type -> services.CustomViewResponse
	2, // 14: services.CustomViewService.List:output_type -> services.CustomViewResponse
	9, // [9:15] is the sub-list for method output_type
	3, // [3:9] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_customViewService_proto_init() }
func file_customViewService_proto_init() {
	if File_customViewService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_customViewService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomView); i {
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
		file_customViewService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomViewRequest); i {
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
		file_customViewService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomViewResponse); i {
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
			RawDescriptor: file_customViewService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_customViewService_proto_goTypes,
		DependencyIndexes: file_customViewService_proto_depIdxs,
		MessageInfos:      file_customViewService_proto_msgTypes,
	}.Build()
	File_customViewService_proto = out.File
	file_customViewService_proto_rawDesc = nil
	file_customViewService_proto_goTypes = nil
	file_customViewService_proto_depIdxs = nil
}
