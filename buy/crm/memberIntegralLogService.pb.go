// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: memberIntegralLogService.proto

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

type MemberIntegralLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id"`                                                        //ID
	MemberId          int64   `protobuf:"varint,2,opt,name=member_id,json=memberId,proto3" json:"member_id"`                            //用户ID
	Type              string  `protobuf:"bytes,3,opt,name=type,proto3" json:"type"`                                                     //操作类型（0减少, 1增加）
	OriginalIntegral  int64   `protobuf:"varint,4,opt,name=original_integral,json=originalIntegral,proto3" json:"original_integral"`    //原始积分
	NewIntegral       int64   `protobuf:"varint,5,opt,name=new_integral,json=newIntegral,proto3" json:"new_integral"`                   //最新积分
	OperationIntegral int64   `protobuf:"varint,6,opt,name=operation_integral,json=operationIntegral,proto3" json:"operation_integral"` //操作积分
	Msg               string  `protobuf:"bytes,7,opt,name=msg,proto3" json:"msg"`                                                       //操作原因
	OperatorId        int64   `protobuf:"varint,8,opt,name=operator_id,json=operatorId,proto3" json:"operator_id"`                      //操作员ID
	OperatorName      string  `protobuf:"bytes,9,opt,name=operator_name,json=operatorName,proto3" json:"operator_name"`                 //操作员名称
	CreatedAt         string  `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt         string  `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	Member            *Member `protobuf:"bytes,12,opt,name=member,proto3" json:"member"`
}

func (x *MemberIntegralLog) Reset() {
	*x = MemberIntegralLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memberIntegralLogService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberIntegralLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberIntegralLog) ProtoMessage() {}

func (x *MemberIntegralLog) ProtoReflect() protoreflect.Message {
	mi := &file_memberIntegralLogService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberIntegralLog.ProtoReflect.Descriptor instead.
func (*MemberIntegralLog) Descriptor() ([]byte, []int) {
	return file_memberIntegralLogService_proto_rawDescGZIP(), []int{0}
}

func (x *MemberIntegralLog) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MemberIntegralLog) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

func (x *MemberIntegralLog) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *MemberIntegralLog) GetOriginalIntegral() int64 {
	if x != nil {
		return x.OriginalIntegral
	}
	return 0
}

func (x *MemberIntegralLog) GetNewIntegral() int64 {
	if x != nil {
		return x.NewIntegral
	}
	return 0
}

func (x *MemberIntegralLog) GetOperationIntegral() int64 {
	if x != nil {
		return x.OperationIntegral
	}
	return 0
}

func (x *MemberIntegralLog) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *MemberIntegralLog) GetOperatorId() int64 {
	if x != nil {
		return x.OperatorId
	}
	return 0
}

func (x *MemberIntegralLog) GetOperatorName() string {
	if x != nil {
		return x.OperatorName
	}
	return ""
}

func (x *MemberIntegralLog) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *MemberIntegralLog) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *MemberIntegralLog) GetMember() *Member {
	if x != nil {
		return x.Member
	}
	return nil
}

type MemberIntegralLogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int64  `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize int64  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Sorting  string `protobuf:"bytes,3,opt,name=sorting,proto3" json:"sorting"`
	//以下为自定义参数
	Type     string `protobuf:"bytes,4,opt,name=type,proto3" json:"type"`
	MemberId int64  `protobuf:"varint,5,opt,name=member_id,json=memberId,proto3" json:"member_id"`
}

func (x *MemberIntegralLogRequest) Reset() {
	*x = MemberIntegralLogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memberIntegralLogService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberIntegralLogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberIntegralLogRequest) ProtoMessage() {}

func (x *MemberIntegralLogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_memberIntegralLogService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberIntegralLogRequest.ProtoReflect.Descriptor instead.
func (*MemberIntegralLogRequest) Descriptor() ([]byte, []int) {
	return file_memberIntegralLogService_proto_rawDescGZIP(), []int{1}
}

func (x *MemberIntegralLogRequest) GetPaged() int64 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *MemberIntegralLogRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *MemberIntegralLogRequest) GetSorting() string {
	if x != nil {
		return x.Sorting
	}
	return ""
}

func (x *MemberIntegralLogRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *MemberIntegralLogRequest) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

type MemberIntegralLogData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *MemberIntegralLog   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager        `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*MemberIntegralLog `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Info   *common.Info         `protobuf:"bytes,4,opt,name=info,proto3" json:"info"`
}

func (x *MemberIntegralLogData) Reset() {
	*x = MemberIntegralLogData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memberIntegralLogService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberIntegralLogData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberIntegralLogData) ProtoMessage() {}

func (x *MemberIntegralLogData) ProtoReflect() protoreflect.Message {
	mi := &file_memberIntegralLogService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberIntegralLogData.ProtoReflect.Descriptor instead.
func (*MemberIntegralLogData) Descriptor() ([]byte, []int) {
	return file_memberIntegralLogService_proto_rawDescGZIP(), []int{2}
}

func (x *MemberIntegralLogData) GetEntity() *MemberIntegralLog {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *MemberIntegralLogData) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *MemberIntegralLogData) GetItems() []*MemberIntegralLog {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *MemberIntegralLogData) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type MemberIntegralLogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *MemberIntegralLogData `protobuf:"bytes,1,opt,name=data,proto3" json:"data"`
	Error *common.Error          `protobuf:"bytes,2,opt,name=error,proto3" json:"error"`
}

func (x *MemberIntegralLogResponse) Reset() {
	*x = MemberIntegralLogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memberIntegralLogService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberIntegralLogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberIntegralLogResponse) ProtoMessage() {}

func (x *MemberIntegralLogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_memberIntegralLogService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberIntegralLogResponse.ProtoReflect.Descriptor instead.
func (*MemberIntegralLogResponse) Descriptor() ([]byte, []int) {
	return file_memberIntegralLogService_proto_rawDescGZIP(), []int{3}
}

func (x *MemberIntegralLogResponse) GetData() *MemberIntegralLogData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *MemberIntegralLogResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_memberIntegralLogService_proto protoreflect.FileDescriptor

var file_memberIntegralLogService_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c,
	0x4c, 0x6f, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x93, 0x03, 0x0a, 0x11, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x6c, 0x4c, 0x6f, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x6f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x49, 0x6e,
	0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x65, 0x77, 0x5f, 0x69, 0x6e,
	0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6e, 0x65,
	0x77, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x12, 0x2d, 0x0a, 0x12, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x28,
	0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x52, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x98, 0x01, 0x0a, 0x18, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x4c, 0x6f, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74,
	0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69,
	0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x49, 0x64, 0x22, 0xc6, 0x01, 0x0a, 0x15, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x6e,
	0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x4c, 0x6f, 0x67, 0x44, 0x61, 0x74, 0x61, 0x12, 0x33, 0x0a,
	0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49,
	0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x4c, 0x6f, 0x67, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72,
	0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x31, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c,
	0x4c, 0x6f, 0x67, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x75, 0x0a, 0x19,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x4c, 0x6f,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61,
	0x6c, 0x4c, 0x6f, 0x67, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x23,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x32, 0x90, 0x02, 0x0a, 0x18, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x6e,
	0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x4c, 0x6f, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x4c, 0x0a, 0x06, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x1b, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65,
	0x67, 0x72, 0x61, 0x6c, 0x4c, 0x6f, 0x67, 0x1a, 0x23, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61,
	0x6c, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x51,
	0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x22, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c,
	0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65,
	0x67, 0x72, 0x61, 0x6c, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x53, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x22, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x6c, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x23, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_memberIntegralLogService_proto_rawDescOnce sync.Once
	file_memberIntegralLogService_proto_rawDescData = file_memberIntegralLogService_proto_rawDesc
)

func file_memberIntegralLogService_proto_rawDescGZIP() []byte {
	file_memberIntegralLogService_proto_rawDescOnce.Do(func() {
		file_memberIntegralLogService_proto_rawDescData = protoimpl.X.CompressGZIP(file_memberIntegralLogService_proto_rawDescData)
	})
	return file_memberIntegralLogService_proto_rawDescData
}

var file_memberIntegralLogService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_memberIntegralLogService_proto_goTypes = []interface{}{
	(*MemberIntegralLog)(nil),         // 0: services.MemberIntegralLog
	(*MemberIntegralLogRequest)(nil),  // 1: services.MemberIntegralLogRequest
	(*MemberIntegralLogData)(nil),     // 2: services.MemberIntegralLogData
	(*MemberIntegralLogResponse)(nil), // 3: services.MemberIntegralLogResponse
	(*Member)(nil),                    // 4: services.Member
	(*common.Pager)(nil),              // 5: common.Pager
	(*common.Info)(nil),               // 6: common.Info
	(*common.Error)(nil),              // 7: common.Error
}
var file_memberIntegralLogService_proto_depIdxs = []int32{
	4,  // 0: services.MemberIntegralLog.member:type_name -> services.Member
	0,  // 1: services.MemberIntegralLogData.entity:type_name -> services.MemberIntegralLog
	5,  // 2: services.MemberIntegralLogData.pager:type_name -> common.Pager
	0,  // 3: services.MemberIntegralLogData.items:type_name -> services.MemberIntegralLog
	6,  // 4: services.MemberIntegralLogData.info:type_name -> common.Info
	2,  // 5: services.MemberIntegralLogResponse.data:type_name -> services.MemberIntegralLogData
	7,  // 6: services.MemberIntegralLogResponse.error:type_name -> common.Error
	0,  // 7: services.MemberIntegralLogService.Detail:input_type -> services.MemberIntegralLog
	1,  // 8: services.MemberIntegralLogService.List:input_type -> services.MemberIntegralLogRequest
	1,  // 9: services.MemberIntegralLogService.Search:input_type -> services.MemberIntegralLogRequest
	3,  // 10: services.MemberIntegralLogService.Detail:output_type -> services.MemberIntegralLogResponse
	3,  // 11: services.MemberIntegralLogService.List:output_type -> services.MemberIntegralLogResponse
	3,  // 12: services.MemberIntegralLogService.Search:output_type -> services.MemberIntegralLogResponse
	10, // [10:13] is the sub-list for method output_type
	7,  // [7:10] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_memberIntegralLogService_proto_init() }
func file_memberIntegralLogService_proto_init() {
	if File_memberIntegralLogService_proto != nil {
		return
	}
	file_memberService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_memberIntegralLogService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberIntegralLog); i {
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
		file_memberIntegralLogService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberIntegralLogRequest); i {
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
		file_memberIntegralLogService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberIntegralLogData); i {
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
		file_memberIntegralLogService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberIntegralLogResponse); i {
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
			RawDescriptor: file_memberIntegralLogService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_memberIntegralLogService_proto_goTypes,
		DependencyIndexes: file_memberIntegralLogService_proto_depIdxs,
		MessageInfos:      file_memberIntegralLogService_proto_msgTypes,
	}.Build()
	File_memberIntegralLogService_proto = out.File
	file_memberIntegralLogService_proto_rawDesc = nil
	file_memberIntegralLogService_proto_goTypes = nil
	file_memberIntegralLogService_proto_depIdxs = nil
}
