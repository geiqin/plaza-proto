// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: spuBrowseService.proto

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

// 商品浏览
type SpuBrowse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`                                //ID
	SpuId     int64 `protobuf:"varint,2,opt,name=spu_id,json=spuId,proto3" json:"spu_id"`             //SpuID
	MemberId  int64 `protobuf:"varint,3,opt,name=member_id,json=memberId,proto3" json:"member_id"`    //会员ID
	CreatedAt int64 `protobuf:"varint,4,opt,name=created_at,json=createdAt,proto3" json:"created_at"` //创建时间
	UpdatedAt int64 `protobuf:"varint,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"` //修改时间
	Spu       *Spu  `protobuf:"bytes,6,opt,name=spu,proto3" json:"spu"`
}

func (x *SpuBrowse) Reset() {
	*x = SpuBrowse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spuBrowseService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpuBrowse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpuBrowse) ProtoMessage() {}

func (x *SpuBrowse) ProtoReflect() protoreflect.Message {
	mi := &file_spuBrowseService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpuBrowse.ProtoReflect.Descriptor instead.
func (*SpuBrowse) Descriptor() ([]byte, []int) {
	return file_spuBrowseService_proto_rawDescGZIP(), []int{0}
}

func (x *SpuBrowse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SpuBrowse) GetSpuId() int64 {
	if x != nil {
		return x.SpuId
	}
	return 0
}

func (x *SpuBrowse) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

func (x *SpuBrowse) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *SpuBrowse) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *SpuBrowse) GetSpu() *Spu {
	if x != nil {
		return x.Spu
	}
	return nil
}

// 商品浏览请求参数
type SpuBrowseRequest struct {
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
	SpuId    int64 `protobuf:"varint,11,opt,name=spu_id,json=spuId,proto3" json:"spu_id"`          //SpuID
	MemberId int64 `protobuf:"varint,12,opt,name=member_id,json=memberId,proto3" json:"member_id"` //会员ID
}

func (x *SpuBrowseRequest) Reset() {
	*x = SpuBrowseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spuBrowseService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpuBrowseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpuBrowseRequest) ProtoMessage() {}

func (x *SpuBrowseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spuBrowseService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpuBrowseRequest.ProtoReflect.Descriptor instead.
func (*SpuBrowseRequest) Descriptor() ([]byte, []int) {
	return file_spuBrowseService_proto_rawDescGZIP(), []int{1}
}

func (x *SpuBrowseRequest) GetTop() int32 {
	if x != nil {
		return x.Top
	}
	return 0
}

func (x *SpuBrowseRequest) GetPaged() int64 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *SpuBrowseRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *SpuBrowseRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *SpuBrowseRequest) GetSorts() []string {
	if x != nil {
		return x.Sorts
	}
	return nil
}

func (x *SpuBrowseRequest) GetDateRange() []string {
	if x != nil {
		return x.DateRange
	}
	return nil
}

func (x *SpuBrowseRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *SpuBrowseRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SpuBrowseRequest) GetSpuId() int64 {
	if x != nil {
		return x.SpuId
	}
	return 0
}

func (x *SpuBrowseRequest) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

// 商品浏览响应数据
type SpuBrowseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg    string        `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Entity *SpuBrowse    `protobuf:"bytes,3,opt,name=entity,proto3" json:"entity"`
	Items  []*SpuBrowse  `protobuf:"bytes,4,rep,name=items,proto3" json:"items"`
	Count  int64         `protobuf:"varint,5,opt,name=count,proto3" json:"count"`
}

func (x *SpuBrowseResponse) Reset() {
	*x = SpuBrowseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spuBrowseService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpuBrowseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpuBrowseResponse) ProtoMessage() {}

func (x *SpuBrowseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spuBrowseService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpuBrowseResponse.ProtoReflect.Descriptor instead.
func (*SpuBrowseResponse) Descriptor() ([]byte, []int) {
	return file_spuBrowseService_proto_rawDescGZIP(), []int{2}
}

func (x *SpuBrowseResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *SpuBrowseResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *SpuBrowseResponse) GetEntity() *SpuBrowse {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *SpuBrowseResponse) GetItems() []*SpuBrowse {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *SpuBrowseResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_spuBrowseService_proto protoreflect.FileDescriptor

var file_spuBrowseService_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x70, 0x75, 0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x73, 0x70, 0x75, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xae, 0x01, 0x0a, 0x09, 0x53, 0x70, 0x75, 0x42,
	0x72, 0x6f, 0x77, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x70, 0x75, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x70, 0x75, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1f, 0x0a, 0x03, 0x73, 0x70, 0x75, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x53, 0x70, 0x75, 0x52, 0x03, 0x73, 0x70, 0x75, 0x22, 0xfe, 0x01, 0x0a, 0x10, 0x53, 0x70, 0x75,
	0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x74, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x74, 0x6f, 0x70, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x73,
	0x6f, 0x72, 0x74, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x61, 0x6e,
	0x67, 0x65, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x64, 0x61, 0x74, 0x65, 0x52, 0x61,
	0x6e, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x03,
	0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x70, 0x75, 0x5f, 0x69, 0x64, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x70, 0x75, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x22, 0xb8, 0x01, 0x0a, 0x11, 0x53, 0x70,
	0x75, 0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73,
	0x67, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52,
	0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x53, 0x70, 0x75, 0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x52, 0x06, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x12, 0x29, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x70,
	0x75, 0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x32, 0xd2, 0x02, 0x0a, 0x10, 0x53, 0x70, 0x75, 0x42, 0x72, 0x6f, 0x77,
	0x73, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x06, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53,
	0x70, 0x75, 0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x53, 0x70, 0x75, 0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x70, 0x75,
	0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x53, 0x70, 0x75, 0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x13, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x70, 0x75, 0x42, 0x72, 0x6f, 0x77, 0x73,
	0x65, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x70, 0x75,
	0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x43, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1a, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x70, 0x75, 0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x53, 0x70, 0x75, 0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x70, 0x75, 0x42, 0x72, 0x6f,
	0x77, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x70, 0x75, 0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spuBrowseService_proto_rawDescOnce sync.Once
	file_spuBrowseService_proto_rawDescData = file_spuBrowseService_proto_rawDesc
)

func file_spuBrowseService_proto_rawDescGZIP() []byte {
	file_spuBrowseService_proto_rawDescOnce.Do(func() {
		file_spuBrowseService_proto_rawDescData = protoimpl.X.CompressGZIP(file_spuBrowseService_proto_rawDescData)
	})
	return file_spuBrowseService_proto_rawDescData
}

var file_spuBrowseService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_spuBrowseService_proto_goTypes = []interface{}{
	(*SpuBrowse)(nil),         // 0: services.SpuBrowse
	(*SpuBrowseRequest)(nil),  // 1: services.SpuBrowseRequest
	(*SpuBrowseResponse)(nil), // 2: services.SpuBrowseResponse
	(*Spu)(nil),               // 3: services.Spu
	(*common.Pager)(nil),      // 4: common.Pager
}
var file_spuBrowseService_proto_depIdxs = []int32{
	3, // 0: services.SpuBrowse.spu:type_name -> services.Spu
	4, // 1: services.SpuBrowseResponse.pager:type_name -> common.Pager
	0, // 2: services.SpuBrowseResponse.entity:type_name -> services.SpuBrowse
	0, // 3: services.SpuBrowseResponse.items:type_name -> services.SpuBrowse
	0, // 4: services.SpuBrowseService.Create:input_type -> services.SpuBrowse
	0, // 5: services.SpuBrowseService.Delete:input_type -> services.SpuBrowse
	0, // 6: services.SpuBrowseService.Get:input_type -> services.SpuBrowse
	1, // 7: services.SpuBrowseService.Search:input_type -> services.SpuBrowseRequest
	1, // 8: services.SpuBrowseService.Count:input_type -> services.SpuBrowseRequest
	2, // 9: services.SpuBrowseService.Create:output_type -> services.SpuBrowseResponse
	2, // 10: services.SpuBrowseService.Delete:output_type -> services.SpuBrowseResponse
	2, // 11: services.SpuBrowseService.Get:output_type -> services.SpuBrowseResponse
	2, // 12: services.SpuBrowseService.Search:output_type -> services.SpuBrowseResponse
	2, // 13: services.SpuBrowseService.Count:output_type -> services.SpuBrowseResponse
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_spuBrowseService_proto_init() }
func file_spuBrowseService_proto_init() {
	if File_spuBrowseService_proto != nil {
		return
	}
	file_spuService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_spuBrowseService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpuBrowse); i {
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
		file_spuBrowseService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpuBrowseRequest); i {
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
		file_spuBrowseService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpuBrowseResponse); i {
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
			RawDescriptor: file_spuBrowseService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spuBrowseService_proto_goTypes,
		DependencyIndexes: file_spuBrowseService_proto_depIdxs,
		MessageInfos:      file_spuBrowseService_proto_msgTypes,
	}.Build()
	File_spuBrowseService_proto = out.File
	file_spuBrowseService_proto_rawDesc = nil
	file_spuBrowseService_proto_goTypes = nil
	file_spuBrowseService_proto_depIdxs = nil
}
