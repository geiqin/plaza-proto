// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: commentService.proto

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

type Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64           `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	SkuId         int64           `protobuf:"varint,2,opt,name=sku_id,json=skuId,proto3" json:"sku_id,omitempty"`
	SpuId         int64           `protobuf:"varint,3,opt,name=spu_id,json=spuId,proto3" json:"spu_id,omitempty"`
	SpuName       string          `protobuf:"bytes,4,opt,name=spu_name,json=spuName,proto3" json:"spu_name,omitempty"`
	SpuAttributes string          `protobuf:"bytes,5,opt,name=spu_attributes,json=spuAttributes,proto3" json:"spu_attributes,omitempty"`
	ParentId      int32           `protobuf:"varint,6,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	MemberId      int64           `protobuf:"varint,7,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
	MemberName    string          `protobuf:"bytes,8,opt,name=member_name,json=memberName,proto3" json:"member_name,omitempty"`
	MemberAvatar  string          `protobuf:"bytes,9,opt,name=member_avatar,json=memberAvatar,proto3" json:"member_avatar,omitempty"`
	MemberIp      string          `protobuf:"bytes,10,opt,name=member_ip,json=memberIp,proto3" json:"member_ip,omitempty"`
	IsDisplay     bool            `protobuf:"varint,11,opt,name=is_display,json=isDisplay,proto3" json:"is_display,omitempty"`
	Content       string          `protobuf:"bytes,12,opt,name=content,proto3" json:"content,omitempty"`
	LikesCount    int32           `protobuf:"varint,13,opt,name=likes_count,json=likesCount,proto3" json:"likes_count,omitempty"`
	ReplyCount    int32           `protobuf:"varint,14,opt,name=reply_count,json=replyCount,proto3" json:"reply_count,omitempty"`
	CreatorId     int64           `protobuf:"varint,15,opt,name=creator_id,json=creatorId,proto3" json:"creator_id,omitempty"`
	CreatorName   string          `protobuf:"bytes,16,opt,name=creator_name,json=creatorName,proto3" json:"creator_name,omitempty"`
	CreatedAt     string          `protobuf:"bytes,17,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string          `protobuf:"bytes,18,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Attachments   []*Attachment   `protobuf:"bytes,19,rep,name=attachments,proto3" json:"attachments,omitempty"`
	Replies       []*CommentReply `protobuf:"bytes,20,rep,name=replies,proto3" json:"replies,omitempty"`
}

func (x *Comment) Reset() {
	*x = Comment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commentService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comment) ProtoMessage() {}

func (x *Comment) ProtoReflect() protoreflect.Message {
	mi := &file_commentService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comment.ProtoReflect.Descriptor instead.
func (*Comment) Descriptor() ([]byte, []int) {
	return file_commentService_proto_rawDescGZIP(), []int{0}
}

func (x *Comment) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Comment) GetSkuId() int64 {
	if x != nil {
		return x.SkuId
	}
	return 0
}

func (x *Comment) GetSpuId() int64 {
	if x != nil {
		return x.SpuId
	}
	return 0
}

func (x *Comment) GetSpuName() string {
	if x != nil {
		return x.SpuName
	}
	return ""
}

func (x *Comment) GetSpuAttributes() string {
	if x != nil {
		return x.SpuAttributes
	}
	return ""
}

func (x *Comment) GetParentId() int32 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *Comment) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

func (x *Comment) GetMemberName() string {
	if x != nil {
		return x.MemberName
	}
	return ""
}

func (x *Comment) GetMemberAvatar() string {
	if x != nil {
		return x.MemberAvatar
	}
	return ""
}

func (x *Comment) GetMemberIp() string {
	if x != nil {
		return x.MemberIp
	}
	return ""
}

func (x *Comment) GetIsDisplay() bool {
	if x != nil {
		return x.IsDisplay
	}
	return false
}

func (x *Comment) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Comment) GetLikesCount() int32 {
	if x != nil {
		return x.LikesCount
	}
	return 0
}

func (x *Comment) GetReplyCount() int32 {
	if x != nil {
		return x.ReplyCount
	}
	return 0
}

func (x *Comment) GetCreatorId() int64 {
	if x != nil {
		return x.CreatorId
	}
	return 0
}

func (x *Comment) GetCreatorName() string {
	if x != nil {
		return x.CreatorName
	}
	return ""
}

func (x *Comment) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Comment) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Comment) GetAttachments() []*Attachment {
	if x != nil {
		return x.Attachments
	}
	return nil
}

func (x *Comment) GetReplies() []*CommentReply {
	if x != nil {
		return x.Replies
	}
	return nil
}

type CommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged     int64   `protobuf:"varint,1,opt,name=paged,proto3" json:"paged,omitempty"`
	PageSize  int64   `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Keywords  string  `protobuf:"bytes,3,opt,name=keywords,proto3" json:"keywords,omitempty"`
	Id        int64   `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
	SpuName   string  `protobuf:"bytes,5,opt,name=spu_name,json=spuName,proto3" json:"spu_name,omitempty"`
	StartDate string  `protobuf:"bytes,6,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate   string  `protobuf:"bytes,7,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	Ids       []int64 `protobuf:"varint,8,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *CommentRequest) Reset() {
	*x = CommentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commentService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentRequest) ProtoMessage() {}

func (x *CommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commentService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentRequest.ProtoReflect.Descriptor instead.
func (*CommentRequest) Descriptor() ([]byte, []int) {
	return file_commentService_proto_rawDescGZIP(), []int{1}
}

func (x *CommentRequest) GetPaged() int64 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *CommentRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *CommentRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *CommentRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CommentRequest) GetSpuName() string {
	if x != nil {
		return x.SpuName
	}
	return ""
}

func (x *CommentRequest) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *CommentRequest) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

func (x *CommentRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type CommentData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Comment      `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items  []*Comment    `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Info   *common.Info  `protobuf:"bytes,4,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *CommentData) Reset() {
	*x = CommentData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commentService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentData) ProtoMessage() {}

func (x *CommentData) ProtoReflect() protoreflect.Message {
	mi := &file_commentService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentData.ProtoReflect.Descriptor instead.
func (*CommentData) Descriptor() ([]byte, []int) {
	return file_commentService_proto_rawDescGZIP(), []int{2}
}

func (x *CommentData) GetEntity() *Comment {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *CommentData) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *CommentData) GetItems() []*Comment {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *CommentData) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type CommentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *CommentData  `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Error *common.Error `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *CommentResponse) Reset() {
	*x = CommentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commentService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentResponse) ProtoMessage() {}

func (x *CommentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commentService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentResponse.ProtoReflect.Descriptor instead.
func (*CommentResponse) Descriptor() ([]byte, []int) {
	return file_commentService_proto_rawDescGZIP(), []int{3}
}

func (x *CommentResponse) GetData() *CommentData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *CommentResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_commentService_proto protoreflect.FileDescriptor

var file_commentService_proto_rawDesc = []byte{
	0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x05, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x6b, 0x75, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x6b, 0x75, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x70,
	0x75, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x70, 0x75, 0x49,
	0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x70, 0x75, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x70, 0x75, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e,
	0x73, 0x70, 0x75, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x70, 0x75, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23,
	0x0a, 0x0d, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x41, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x70,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x70,
	0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x69, 0x6b,
	0x65, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x6c, 0x69, 0x6b, 0x65, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65,
	0x70, 0x6c, 0x79, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x36, 0x0a, 0x0b, 0x61,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x13, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x74, 0x74, 0x61,
	0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x12, 0x30, 0x0a, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x65, 0x73, 0x18, 0x14,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x07, 0x72, 0x65,
	0x70, 0x6c, 0x69, 0x65, 0x73, 0x22, 0xd6, 0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6b,
	0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b,
	0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x70, 0x75, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x70, 0x75, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x69, 0x64, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0xa8,
	0x01, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x29,
	0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x27,
	0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x61, 0x0a, 0x0f, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xb1, 0x02, 0x0a,
	0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x36, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x19, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x33, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3d, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x18, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_commentService_proto_rawDescOnce sync.Once
	file_commentService_proto_rawDescData = file_commentService_proto_rawDesc
)

func file_commentService_proto_rawDescGZIP() []byte {
	file_commentService_proto_rawDescOnce.Do(func() {
		file_commentService_proto_rawDescData = protoimpl.X.CompressGZIP(file_commentService_proto_rawDescData)
	})
	return file_commentService_proto_rawDescData
}

var file_commentService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_commentService_proto_goTypes = []interface{}{
	(*Comment)(nil),         // 0: services.Comment
	(*CommentRequest)(nil),  // 1: services.CommentRequest
	(*CommentData)(nil),     // 2: services.CommentData
	(*CommentResponse)(nil), // 3: services.CommentResponse
	(*Attachment)(nil),      // 4: services.Attachment
	(*CommentReply)(nil),    // 5: services.CommentReply
	(*common.Pager)(nil),    // 6: common.Pager
	(*common.Info)(nil),     // 7: common.Info
	(*common.Error)(nil),    // 8: common.Error
}
var file_commentService_proto_depIdxs = []int32{
	4,  // 0: services.Comment.attachments:type_name -> services.Attachment
	5,  // 1: services.Comment.replies:type_name -> services.CommentReply
	0,  // 2: services.CommentData.entity:type_name -> services.Comment
	6,  // 3: services.CommentData.pager:type_name -> common.Pager
	0,  // 4: services.CommentData.items:type_name -> services.Comment
	7,  // 5: services.CommentData.info:type_name -> common.Info
	2,  // 6: services.CommentResponse.data:type_name -> services.CommentData
	8,  // 7: services.CommentResponse.error:type_name -> common.Error
	0,  // 8: services.CommentService.Create:input_type -> services.Comment
	0,  // 9: services.CommentService.Delete:input_type -> services.Comment
	0,  // 10: services.CommentService.Get:input_type -> services.Comment
	1,  // 11: services.CommentService.List:input_type -> services.CommentRequest
	1,  // 12: services.CommentService.Search:input_type -> services.CommentRequest
	3,  // 13: services.CommentService.Create:output_type -> services.CommentResponse
	3,  // 14: services.CommentService.Delete:output_type -> services.CommentResponse
	3,  // 15: services.CommentService.Get:output_type -> services.CommentResponse
	3,  // 16: services.CommentService.List:output_type -> services.CommentResponse
	3,  // 17: services.CommentService.Search:output_type -> services.CommentResponse
	13, // [13:18] is the sub-list for method output_type
	8,  // [8:13] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_commentService_proto_init() }
func file_commentService_proto_init() {
	if File_commentService_proto != nil {
		return
	}
	file_attachmentService_proto_init()
	file_commentReplyService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_commentService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Comment); i {
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
		file_commentService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentRequest); i {
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
		file_commentService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentData); i {
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
		file_commentService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentResponse); i {
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
			RawDescriptor: file_commentService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_commentService_proto_goTypes,
		DependencyIndexes: file_commentService_proto_depIdxs,
		MessageInfos:      file_commentService_proto_msgTypes,
	}.Build()
	File_commentService_proto = out.File
	file_commentService_proto_rawDesc = nil
	file_commentService_proto_goTypes = nil
	file_commentService_proto_depIdxs = nil
}
