// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: complaintProgressService.proto

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

// 售后处理进度
type ComplaintProgress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`                                           //ID
	ComplaintId   int64  `protobuf:"varint,2,opt,name=complaint_id,json=complaintId,proto3" json:"complaint_id"`      //维权申请id
	CurrentStatus string `protobuf:"bytes,3,opt,name=current_status,json=currentStatus,proto3" json:"current_status"` //当前维权状态
	Note          string `protobuf:"bytes,4,opt,name=note,proto3" json:"note"`                                        //进度备注
	CreatedAt     string `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at"`             //创建时间
	UpdatedAt     string `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`             //修改时间
}

func (x *ComplaintProgress) Reset() {
	*x = ComplaintProgress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_complaintProgressService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComplaintProgress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComplaintProgress) ProtoMessage() {}

func (x *ComplaintProgress) ProtoReflect() protoreflect.Message {
	mi := &file_complaintProgressService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComplaintProgress.ProtoReflect.Descriptor instead.
func (*ComplaintProgress) Descriptor() ([]byte, []int) {
	return file_complaintProgressService_proto_rawDescGZIP(), []int{0}
}

func (x *ComplaintProgress) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ComplaintProgress) GetComplaintId() int64 {
	if x != nil {
		return x.ComplaintId
	}
	return 0
}

func (x *ComplaintProgress) GetCurrentStatus() string {
	if x != nil {
		return x.CurrentStatus
	}
	return ""
}

func (x *ComplaintProgress) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

func (x *ComplaintProgress) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *ComplaintProgress) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

// 售后处理进度请求参数
type ComplaintProgressRequest struct {
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
	ComplaintId   int64  `protobuf:"varint,11,opt,name=complaint_id,json=complaintId,proto3" json:"complaint_id"`      //维权申请id
	CurrentStatus string `protobuf:"bytes,12,opt,name=current_status,json=currentStatus,proto3" json:"current_status"` //当前维权状态
}

func (x *ComplaintProgressRequest) Reset() {
	*x = ComplaintProgressRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_complaintProgressService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComplaintProgressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComplaintProgressRequest) ProtoMessage() {}

func (x *ComplaintProgressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_complaintProgressService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComplaintProgressRequest.ProtoReflect.Descriptor instead.
func (*ComplaintProgressRequest) Descriptor() ([]byte, []int) {
	return file_complaintProgressService_proto_rawDescGZIP(), []int{1}
}

func (x *ComplaintProgressRequest) GetTop() int32 {
	if x != nil {
		return x.Top
	}
	return 0
}

func (x *ComplaintProgressRequest) GetPaged() int64 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *ComplaintProgressRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ComplaintProgressRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *ComplaintProgressRequest) GetSort() []string {
	if x != nil {
		return x.Sort
	}
	return nil
}

func (x *ComplaintProgressRequest) GetDateRange() []string {
	if x != nil {
		return x.DateRange
	}
	return nil
}

func (x *ComplaintProgressRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *ComplaintProgressRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ComplaintProgressRequest) GetComplaintId() int64 {
	if x != nil {
		return x.ComplaintId
	}
	return 0
}

func (x *ComplaintProgressRequest) GetCurrentStatus() string {
	if x != nil {
		return x.CurrentStatus
	}
	return ""
}

// 售后处理进度响应数据
type ComplaintProgressResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg    string               `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg"`
	Pager  *common.Pager        `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Entity *ComplaintProgress   `protobuf:"bytes,3,opt,name=entity,proto3" json:"entity"`
	Items  []*ComplaintProgress `protobuf:"bytes,4,rep,name=items,proto3" json:"items"`
}

func (x *ComplaintProgressResponse) Reset() {
	*x = ComplaintProgressResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_complaintProgressService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComplaintProgressResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComplaintProgressResponse) ProtoMessage() {}

func (x *ComplaintProgressResponse) ProtoReflect() protoreflect.Message {
	mi := &file_complaintProgressService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComplaintProgressResponse.ProtoReflect.Descriptor instead.
func (*ComplaintProgressResponse) Descriptor() ([]byte, []int) {
	return file_complaintProgressService_proto_rawDescGZIP(), []int{2}
}

func (x *ComplaintProgressResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *ComplaintProgressResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *ComplaintProgressResponse) GetEntity() *ComplaintProgress {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *ComplaintProgressResponse) GetItems() []*ComplaintProgress {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_complaintProgressService_proto protoreflect.FileDescriptor

var file_complaintProgressService_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbf, 0x01,
	0x0a, 0x11, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x6c,
	0x61, 0x69, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x74,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22,
	0x9a, 0x02, 0x0a, 0x18, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x74, 0x50, 0x72, 0x6f,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x74, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x74, 0x6f, 0x70, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70,
	0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6f, 0x72,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x64, 0x61, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69,
	0x64, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x61,
	0x69, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xba, 0x01, 0x0a,
	0x19, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x23, 0x0a, 0x05,
	0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65,
	0x72, 0x12, 0x33, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x61, 0x69, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x06,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x31, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x32, 0x8d, 0x02, 0x0a, 0x18, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1b, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x61, 0x69,
	0x6e, 0x74, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x1a, 0x23, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x74, 0x50,
	0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x53, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x22, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x74,
	0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x23, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c,
	0x61, 0x69, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x22,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x61,
	0x69, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x23, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_complaintProgressService_proto_rawDescOnce sync.Once
	file_complaintProgressService_proto_rawDescData = file_complaintProgressService_proto_rawDesc
)

func file_complaintProgressService_proto_rawDescGZIP() []byte {
	file_complaintProgressService_proto_rawDescOnce.Do(func() {
		file_complaintProgressService_proto_rawDescData = protoimpl.X.CompressGZIP(file_complaintProgressService_proto_rawDescData)
	})
	return file_complaintProgressService_proto_rawDescData
}

var file_complaintProgressService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_complaintProgressService_proto_goTypes = []interface{}{
	(*ComplaintProgress)(nil),         // 0: services.ComplaintProgress
	(*ComplaintProgressRequest)(nil),  // 1: services.ComplaintProgressRequest
	(*ComplaintProgressResponse)(nil), // 2: services.ComplaintProgressResponse
	(*common.Pager)(nil),              // 3: common.Pager
}
var file_complaintProgressService_proto_depIdxs = []int32{
	3, // 0: services.ComplaintProgressResponse.pager:type_name -> common.Pager
	0, // 1: services.ComplaintProgressResponse.entity:type_name -> services.ComplaintProgress
	0, // 2: services.ComplaintProgressResponse.items:type_name -> services.ComplaintProgress
	0, // 3: services.ComplaintProgressService.Get:input_type -> services.ComplaintProgress
	1, // 4: services.ComplaintProgressService.Search:input_type -> services.ComplaintProgressRequest
	1, // 5: services.ComplaintProgressService.List:input_type -> services.ComplaintProgressRequest
	2, // 6: services.ComplaintProgressService.Get:output_type -> services.ComplaintProgressResponse
	2, // 7: services.ComplaintProgressService.Search:output_type -> services.ComplaintProgressResponse
	2, // 8: services.ComplaintProgressService.List:output_type -> services.ComplaintProgressResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_complaintProgressService_proto_init() }
func file_complaintProgressService_proto_init() {
	if File_complaintProgressService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_complaintProgressService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComplaintProgress); i {
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
		file_complaintProgressService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComplaintProgressRequest); i {
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
		file_complaintProgressService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComplaintProgressResponse); i {
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
			RawDescriptor: file_complaintProgressService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_complaintProgressService_proto_goTypes,
		DependencyIndexes: file_complaintProgressService_proto_depIdxs,
		MessageInfos:      file_complaintProgressService_proto_msgTypes,
	}.Build()
	File_complaintProgressService_proto = out.File
	file_complaintProgressService_proto_rawDesc = nil
	file_complaintProgressService_proto_goTypes = nil
	file_complaintProgressService_proto_depIdxs = nil
}