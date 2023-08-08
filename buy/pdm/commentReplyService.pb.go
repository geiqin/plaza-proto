// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: commentReplyService.proto

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

type CommentReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64         `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	FeedbackId  int64         `protobuf:"varint,2,opt,name=feedback_id,json=feedbackId,proto3" json:"feedback_id"`
	Content     string        `protobuf:"bytes,3,opt,name=content,proto3" json:"content"`
	CreatorId   int64         `protobuf:"varint,4,opt,name=creator_id,json=creatorId,proto3" json:"creator_id"`
	CreatorName string        `protobuf:"bytes,5,opt,name=creator_name,json=creatorName,proto3" json:"creator_name"`
	CreatorType string        `protobuf:"bytes,6,opt,name=creator_type,json=creatorType,proto3" json:"creator_type"`
	CreatedAt   string        `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt   string        `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	Attachments []*Attachment `protobuf:"bytes,9,rep,name=attachments,proto3" json:"attachments"`
}

func (x *CommentReply) Reset() {
	*x = CommentReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commentReplyService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentReply) ProtoMessage() {}

func (x *CommentReply) ProtoReflect() protoreflect.Message {
	mi := &file_commentReplyService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentReply.ProtoReflect.Descriptor instead.
func (*CommentReply) Descriptor() ([]byte, []int) {
	return file_commentReplyService_proto_rawDescGZIP(), []int{0}
}

func (x *CommentReply) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CommentReply) GetFeedbackId() int64 {
	if x != nil {
		return x.FeedbackId
	}
	return 0
}

func (x *CommentReply) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CommentReply) GetCreatorId() int64 {
	if x != nil {
		return x.CreatorId
	}
	return 0
}

func (x *CommentReply) GetCreatorName() string {
	if x != nil {
		return x.CreatorName
	}
	return ""
}

func (x *CommentReply) GetCreatorType() string {
	if x != nil {
		return x.CreatorType
	}
	return ""
}

func (x *CommentReply) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *CommentReply) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *CommentReply) GetAttachments() []*Attachment {
	if x != nil {
		return x.Attachments
	}
	return nil
}

type CommentReplyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged       int64   `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize    int64   `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Keywords    string  `protobuf:"bytes,3,opt,name=keywords,proto3" json:"keywords"`
	Id          int64   `protobuf:"varint,4,opt,name=id,proto3" json:"id"`
	FeedbackId  int64   `protobuf:"varint,5,opt,name=feedback_id,json=feedbackId,proto3" json:"feedback_id"`
	Content     string  `protobuf:"bytes,6,opt,name=content,proto3" json:"content"`
	Contact     string  `protobuf:"bytes,7,opt,name=contact,proto3" json:"contact"`
	CreatorId   int64   `protobuf:"varint,8,opt,name=creator_id,json=creatorId,proto3" json:"creator_id"`
	CreatorName string  `protobuf:"bytes,9,opt,name=creator_name,json=creatorName,proto3" json:"creator_name"`
	CreatorType string  `protobuf:"bytes,10,opt,name=creator_type,json=creatorType,proto3" json:"creator_type"`
	StartDate   string  `protobuf:"bytes,11,opt,name=start_date,json=startDate,proto3" json:"start_date"`
	EndDate     string  `protobuf:"bytes,12,opt,name=end_date,json=endDate,proto3" json:"end_date"`
	Ids         []int64 `protobuf:"varint,13,rep,packed,name=ids,proto3" json:"ids"`
}

func (x *CommentReplyRequest) Reset() {
	*x = CommentReplyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commentReplyService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentReplyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentReplyRequest) ProtoMessage() {}

func (x *CommentReplyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commentReplyService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentReplyRequest.ProtoReflect.Descriptor instead.
func (*CommentReplyRequest) Descriptor() ([]byte, []int) {
	return file_commentReplyService_proto_rawDescGZIP(), []int{1}
}

func (x *CommentReplyRequest) GetPaged() int64 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *CommentReplyRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *CommentReplyRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *CommentReplyRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CommentReplyRequest) GetFeedbackId() int64 {
	if x != nil {
		return x.FeedbackId
	}
	return 0
}

func (x *CommentReplyRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CommentReplyRequest) GetContact() string {
	if x != nil {
		return x.Contact
	}
	return ""
}

func (x *CommentReplyRequest) GetCreatorId() int64 {
	if x != nil {
		return x.CreatorId
	}
	return 0
}

func (x *CommentReplyRequest) GetCreatorName() string {
	if x != nil {
		return x.CreatorName
	}
	return ""
}

func (x *CommentReplyRequest) GetCreatorType() string {
	if x != nil {
		return x.CreatorType
	}
	return ""
}

func (x *CommentReplyRequest) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *CommentReplyRequest) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

func (x *CommentReplyRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type CommentReplyData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *CommentReply   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager   `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*CommentReply `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Info   *common.Info    `protobuf:"bytes,4,opt,name=info,proto3" json:"info"`
}

func (x *CommentReplyData) Reset() {
	*x = CommentReplyData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commentReplyService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentReplyData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentReplyData) ProtoMessage() {}

func (x *CommentReplyData) ProtoReflect() protoreflect.Message {
	mi := &file_commentReplyService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentReplyData.ProtoReflect.Descriptor instead.
func (*CommentReplyData) Descriptor() ([]byte, []int) {
	return file_commentReplyService_proto_rawDescGZIP(), []int{2}
}

func (x *CommentReplyData) GetEntity() *CommentReply {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *CommentReplyData) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *CommentReplyData) GetItems() []*CommentReply {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *CommentReplyData) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type CommentReplyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *CommentReplyData `protobuf:"bytes,1,opt,name=data,proto3" json:"data"`
	Error *common.Error     `protobuf:"bytes,2,opt,name=error,proto3" json:"error"`
}

func (x *CommentReplyResponse) Reset() {
	*x = CommentReplyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commentReplyService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentReplyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentReplyResponse) ProtoMessage() {}

func (x *CommentReplyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commentReplyService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentReplyResponse.ProtoReflect.Descriptor instead.
func (*CommentReplyResponse) Descriptor() ([]byte, []int) {
	return file_commentReplyService_proto_rawDescGZIP(), []int{3}
}

func (x *CommentReplyResponse) GetData() *CommentReplyData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *CommentReplyResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_commentReplyService_proto protoreflect.FileDescriptor

var file_commentReplyService_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xb4, 0x02, 0x0a, 0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63,
	0x6b, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x36, 0x0a, 0x0b, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x61, 0x74, 0x74,
	0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0xfa, 0x02, 0x0a, 0x13, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x63, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f,
	0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f,
	0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x03,
	0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0xb7, 0x01, 0x0a, 0x10, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2e, 0x0a, 0x06, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61,
	0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12,
	0x2c, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x20, 0x0a,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22,
	0x6b, 0x0a, 0x14, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xe8, 0x02, 0x0a,
	0x13, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x16,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x1a, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x1a, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12,
	0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x1a, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47,
	0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_commentReplyService_proto_rawDescOnce sync.Once
	file_commentReplyService_proto_rawDescData = file_commentReplyService_proto_rawDesc
)

func file_commentReplyService_proto_rawDescGZIP() []byte {
	file_commentReplyService_proto_rawDescOnce.Do(func() {
		file_commentReplyService_proto_rawDescData = protoimpl.X.CompressGZIP(file_commentReplyService_proto_rawDescData)
	})
	return file_commentReplyService_proto_rawDescData
}

var file_commentReplyService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_commentReplyService_proto_goTypes = []interface{}{
	(*CommentReply)(nil),         // 0: services.CommentReply
	(*CommentReplyRequest)(nil),  // 1: services.CommentReplyRequest
	(*CommentReplyData)(nil),     // 2: services.CommentReplyData
	(*CommentReplyResponse)(nil), // 3: services.CommentReplyResponse
	(*Attachment)(nil),           // 4: services.Attachment
	(*common.Pager)(nil),         // 5: common.Pager
	(*common.Info)(nil),          // 6: common.Info
	(*common.Error)(nil),         // 7: common.Error
}
var file_commentReplyService_proto_depIdxs = []int32{
	4,  // 0: services.CommentReply.attachments:type_name -> services.Attachment
	0,  // 1: services.CommentReplyData.entity:type_name -> services.CommentReply
	5,  // 2: services.CommentReplyData.pager:type_name -> common.Pager
	0,  // 3: services.CommentReplyData.items:type_name -> services.CommentReply
	6,  // 4: services.CommentReplyData.info:type_name -> common.Info
	2,  // 5: services.CommentReplyResponse.data:type_name -> services.CommentReplyData
	7,  // 6: services.CommentReplyResponse.error:type_name -> common.Error
	0,  // 7: services.CommentReplyService.Create:input_type -> services.CommentReply
	0,  // 8: services.CommentReplyService.Delete:input_type -> services.CommentReply
	0,  // 9: services.CommentReplyService.Get:input_type -> services.CommentReply
	1,  // 10: services.CommentReplyService.List:input_type -> services.CommentReplyRequest
	1,  // 11: services.CommentReplyService.Search:input_type -> services.CommentReplyRequest
	3,  // 12: services.CommentReplyService.Create:output_type -> services.CommentReplyResponse
	3,  // 13: services.CommentReplyService.Delete:output_type -> services.CommentReplyResponse
	3,  // 14: services.CommentReplyService.Get:output_type -> services.CommentReplyResponse
	3,  // 15: services.CommentReplyService.List:output_type -> services.CommentReplyResponse
	3,  // 16: services.CommentReplyService.Search:output_type -> services.CommentReplyResponse
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_commentReplyService_proto_init() }
func file_commentReplyService_proto_init() {
	if File_commentReplyService_proto != nil {
		return
	}
	file_attachmentService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_commentReplyService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentReply); i {
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
		file_commentReplyService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentReplyRequest); i {
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
		file_commentReplyService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentReplyData); i {
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
		file_commentReplyService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentReplyResponse); i {
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
			RawDescriptor: file_commentReplyService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_commentReplyService_proto_goTypes,
		DependencyIndexes: file_commentReplyService_proto_depIdxs,
		MessageInfos:      file_commentReplyService_proto_msgTypes,
	}.Build()
	File_commentReplyService_proto = out.File
	file_commentReplyService_proto_rawDesc = nil
	file_commentReplyService_proto_goTypes = nil
	file_commentReplyService_proto_depIdxs = nil
}
