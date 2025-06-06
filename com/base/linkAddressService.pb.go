// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: linkAddressService.proto

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

// 页面链接
type LinkAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64         `protobuf:"varint,1,opt,name=id,proto3" json:"id"`                                   //ID
	CategoryId int64         `protobuf:"varint,2,opt,name=category_id,json=categoryId,proto3" json:"category_id"` //分类ID
	Type       string        `protobuf:"bytes,3,opt,name=type,proto3" json:"type"`                                //类型
	Name       string        `protobuf:"bytes,4,opt,name=name,proto3" json:"name"`                                //分类名称
	Url        string        `protobuf:"bytes,5,opt,name=url,proto3" json:"url"`                                  //页面链接
	Param      string        `protobuf:"bytes,6,opt,name=param,proto3" json:"param"`                              //参数
	Example    string        `protobuf:"bytes,7,opt,name=example,proto3" json:"example"`                          //事例
	Sort       int32         `protobuf:"varint,8,opt,name=sort,proto3" json:"sort"`                               //排序
	Status     string        `protobuf:"bytes,9,opt,name=status,proto3" json:"status"`                            //状态
	UpdatedAt  int64         `protobuf:"varint,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`   //创建时间
	CreatedAt  int64         `protobuf:"varint,11,opt,name=created_at,json=createdAt,proto3" json:"created_at"`   //修改时间
	Category   *LinkCategory `protobuf:"bytes,12,opt,name=category,proto3" json:"category"`
}

func (x *LinkAddress) Reset() {
	*x = LinkAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_linkAddressService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LinkAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LinkAddress) ProtoMessage() {}

func (x *LinkAddress) ProtoReflect() protoreflect.Message {
	mi := &file_linkAddressService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LinkAddress.ProtoReflect.Descriptor instead.
func (*LinkAddress) Descriptor() ([]byte, []int) {
	return file_linkAddressService_proto_rawDescGZIP(), []int{0}
}

func (x *LinkAddress) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *LinkAddress) GetCategoryId() int64 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

func (x *LinkAddress) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *LinkAddress) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LinkAddress) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *LinkAddress) GetParam() string {
	if x != nil {
		return x.Param
	}
	return ""
}

func (x *LinkAddress) GetExample() string {
	if x != nil {
		return x.Example
	}
	return ""
}

func (x *LinkAddress) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *LinkAddress) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *LinkAddress) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *LinkAddress) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *LinkAddress) GetCategory() *LinkCategory {
	if x != nil {
		return x.Category
	}
	return nil
}

// 页面链接请求参数
type LinkAddressRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Top       int32    `protobuf:"varint,1,opt,name=top,proto3" json:"top"`
	Paged     int64    `protobuf:"varint,2,opt,name=paged,proto3" json:"paged"`
	PageSize  int64    `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Keywords  string   `protobuf:"bytes,4,opt,name=keywords,proto3" json:"keywords"`
	Sorts     []string `protobuf:"bytes,5,rep,name=sorts,proto3" json:"sorts"`
	DateRange []string `protobuf:"bytes,6,rep,name=date_range,json=dateRange,proto3" json:"date_range"`
	Ids       []int64  `protobuf:"varint,7,rep,packed,name=ids,proto3" json:"ids"`
	Id        int64    `protobuf:"varint,8,opt,name=id,proto3" json:"id"`
	//以下为自定义参数
	CategoryId int64  `protobuf:"varint,11,opt,name=category_id,json=categoryId,proto3" json:"category_id"` //分类ID
	Type       string `protobuf:"bytes,12,opt,name=type,proto3" json:"type"`                                //类型
	Name       string `protobuf:"bytes,13,opt,name=name,proto3" json:"name"`                                //分类名称
	Status     string `protobuf:"bytes,14,opt,name=status,proto3" json:"status"`                            //状态
}

func (x *LinkAddressRequest) Reset() {
	*x = LinkAddressRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_linkAddressService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LinkAddressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LinkAddressRequest) ProtoMessage() {}

func (x *LinkAddressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_linkAddressService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LinkAddressRequest.ProtoReflect.Descriptor instead.
func (*LinkAddressRequest) Descriptor() ([]byte, []int) {
	return file_linkAddressService_proto_rawDescGZIP(), []int{1}
}

func (x *LinkAddressRequest) GetTop() int32 {
	if x != nil {
		return x.Top
	}
	return 0
}

func (x *LinkAddressRequest) GetPaged() int64 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *LinkAddressRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *LinkAddressRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *LinkAddressRequest) GetSorts() []string {
	if x != nil {
		return x.Sorts
	}
	return nil
}

func (x *LinkAddressRequest) GetDateRange() []string {
	if x != nil {
		return x.DateRange
	}
	return nil
}

func (x *LinkAddressRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *LinkAddressRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *LinkAddressRequest) GetCategoryId() int64 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

func (x *LinkAddressRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *LinkAddressRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LinkAddressRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

// 页面链接响应数据
type LinkAddressResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg    string         `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg"`
	Pager  *common.Pager  `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Entity *LinkAddress   `protobuf:"bytes,3,opt,name=entity,proto3" json:"entity"`
	Items  []*LinkAddress `protobuf:"bytes,4,rep,name=items,proto3" json:"items"`
}

func (x *LinkAddressResponse) Reset() {
	*x = LinkAddressResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_linkAddressService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LinkAddressResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LinkAddressResponse) ProtoMessage() {}

func (x *LinkAddressResponse) ProtoReflect() protoreflect.Message {
	mi := &file_linkAddressService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LinkAddressResponse.ProtoReflect.Descriptor instead.
func (*LinkAddressResponse) Descriptor() ([]byte, []int) {
	return file_linkAddressService_proto_rawDescGZIP(), []int{2}
}

func (x *LinkAddressResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *LinkAddressResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *LinkAddressResponse) GetEntity() *LinkAddress {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *LinkAddressResponse) GetItems() []*LinkAddress {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_linkAddressService_proto protoreflect.FileDescriptor

var file_linkAddressService_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6c, 0x69, 0x6e, 0x6b, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x6c, 0x69, 0x6e, 0x6b, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xc6, 0x02, 0x0a, 0x0b, 0x4c, 0x69, 0x6e, 0x6b, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x6f, 0x72,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x32, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x22, 0xad, 0x02, 0x0a, 0x12,
	0x4c, 0x69, 0x6e, 0x6b, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x74, 0x6f, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70,
	0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f,
	0x72, 0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f,
	0x72, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x05, 0x73, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18,
	0x07, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xa8, 0x01, 0x0a, 0x13,
	0x4c, 0x69, 0x6e, 0x6b, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61,
	0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x06, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x2b, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x32, 0xa9, 0x03, 0x0a, 0x12, 0x4c, 0x69, 0x6e, 0x6b, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a,
	0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x1a, 0x1d,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x40, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x1a, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x6e, 0x6b,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x40, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x1a, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x69,
	0x6e, 0x6b, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x1a, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x6e,
	0x6b, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x47, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1c, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x04, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c,
	0x69, 0x6e, 0x6b, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x6e,
	0x6b, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_linkAddressService_proto_rawDescOnce sync.Once
	file_linkAddressService_proto_rawDescData = file_linkAddressService_proto_rawDesc
)

func file_linkAddressService_proto_rawDescGZIP() []byte {
	file_linkAddressService_proto_rawDescOnce.Do(func() {
		file_linkAddressService_proto_rawDescData = protoimpl.X.CompressGZIP(file_linkAddressService_proto_rawDescData)
	})
	return file_linkAddressService_proto_rawDescData
}

var file_linkAddressService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_linkAddressService_proto_goTypes = []interface{}{
	(*LinkAddress)(nil),         // 0: services.LinkAddress
	(*LinkAddressRequest)(nil),  // 1: services.LinkAddressRequest
	(*LinkAddressResponse)(nil), // 2: services.LinkAddressResponse
	(*LinkCategory)(nil),        // 3: services.LinkCategory
	(*common.Pager)(nil),        // 4: common.Pager
}
var file_linkAddressService_proto_depIdxs = []int32{
	3,  // 0: services.LinkAddress.category:type_name -> services.LinkCategory
	4,  // 1: services.LinkAddressResponse.pager:type_name -> common.Pager
	0,  // 2: services.LinkAddressResponse.entity:type_name -> services.LinkAddress
	0,  // 3: services.LinkAddressResponse.items:type_name -> services.LinkAddress
	0,  // 4: services.LinkAddressService.Create:input_type -> services.LinkAddress
	0,  // 5: services.LinkAddressService.Update:input_type -> services.LinkAddress
	0,  // 6: services.LinkAddressService.Delete:input_type -> services.LinkAddress
	0,  // 7: services.LinkAddressService.Get:input_type -> services.LinkAddress
	1,  // 8: services.LinkAddressService.Search:input_type -> services.LinkAddressRequest
	1,  // 9: services.LinkAddressService.List:input_type -> services.LinkAddressRequest
	2,  // 10: services.LinkAddressService.Create:output_type -> services.LinkAddressResponse
	2,  // 11: services.LinkAddressService.Update:output_type -> services.LinkAddressResponse
	2,  // 12: services.LinkAddressService.Delete:output_type -> services.LinkAddressResponse
	2,  // 13: services.LinkAddressService.Get:output_type -> services.LinkAddressResponse
	2,  // 14: services.LinkAddressService.Search:output_type -> services.LinkAddressResponse
	2,  // 15: services.LinkAddressService.List:output_type -> services.LinkAddressResponse
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_linkAddressService_proto_init() }
func file_linkAddressService_proto_init() {
	if File_linkAddressService_proto != nil {
		return
	}
	file_linkCategoryService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_linkAddressService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LinkAddress); i {
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
		file_linkAddressService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LinkAddressRequest); i {
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
		file_linkAddressService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LinkAddressResponse); i {
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
			RawDescriptor: file_linkAddressService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_linkAddressService_proto_goTypes,
		DependencyIndexes: file_linkAddressService_proto_depIdxs,
		MessageInfos:      file_linkAddressService_proto_msgTypes,
	}.Build()
	File_linkAddressService_proto = out.File
	file_linkAddressService_proto_rawDesc = nil
	file_linkAddressService_proto_goTypes = nil
	file_linkAddressService_proto_depIdxs = nil
}
