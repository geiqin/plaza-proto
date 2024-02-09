// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: integralService.proto

package services

import (
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

//锁定积分信息
type LockingIntegralInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid             string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid"`                                                   //积分唯一标识
	OrderId          int64  `protobuf:"varint,2,opt,name=order_id,json=orderId,proto3" json:"order_id"`                             //订单id
	MemberId         int64  `protobuf:"varint,3,opt,name=member_id,json=memberId,proto3" json:"member_id"`                          //用户id
	BusinessId       int64  `protobuf:"varint,4,opt,name=business_id,json=businessId,proto3" json:"business_id"`                    //业务ID
	BusinessType     string `protobuf:"bytes,5,opt,name=business_type,json=businessType,proto3" json:"business_type"`               //业务类型
	BusinessTypeName string `protobuf:"bytes,6,opt,name=business_type_name,json=businessTypeName,proto3" json:"business_type_name"` //业务类型名称
	RefundIntegral   int64  `protobuf:"varint,7,opt,name=refund_integral,json=refundIntegral,proto3" json:"refund_integral"`        //退还积分
	Integral         int64  `protobuf:"varint,8,opt,name=integral,proto3" json:"integral"`                                          //积分
	Desc             string `protobuf:"bytes,9,opt,name=desc,proto3" json:"desc"`                                                   //描述
}

func (x *LockingIntegralInfo) Reset() {
	*x = LockingIntegralInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integralService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LockingIntegralInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LockingIntegralInfo) ProtoMessage() {}

func (x *LockingIntegralInfo) ProtoReflect() protoreflect.Message {
	mi := &file_integralService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LockingIntegralInfo.ProtoReflect.Descriptor instead.
func (*LockingIntegralInfo) Descriptor() ([]byte, []int) {
	return file_integralService_proto_rawDescGZIP(), []int{0}
}

func (x *LockingIntegralInfo) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *LockingIntegralInfo) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *LockingIntegralInfo) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

func (x *LockingIntegralInfo) GetBusinessId() int64 {
	if x != nil {
		return x.BusinessId
	}
	return 0
}

func (x *LockingIntegralInfo) GetBusinessType() string {
	if x != nil {
		return x.BusinessType
	}
	return ""
}

func (x *LockingIntegralInfo) GetBusinessTypeName() string {
	if x != nil {
		return x.BusinessTypeName
	}
	return ""
}

func (x *LockingIntegralInfo) GetRefundIntegral() int64 {
	if x != nil {
		return x.RefundIntegral
	}
	return 0
}

func (x *LockingIntegralInfo) GetIntegral() int64 {
	if x != nil {
		return x.Integral
	}
	return 0
}

func (x *LockingIntegralInfo) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

//锁定积分退还
type LockingIntegralRollbackInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid     string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid"`          //积分唯一标识
	Integral int64  `protobuf:"varint,2,opt,name=integral,proto3" json:"integral"` //退还积分
	Reason   string `protobuf:"bytes,3,opt,name=reason,proto3" json:"reason"`      //退还原因
}

func (x *LockingIntegralRollbackInfo) Reset() {
	*x = LockingIntegralRollbackInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integralService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LockingIntegralRollbackInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LockingIntegralRollbackInfo) ProtoMessage() {}

func (x *LockingIntegralRollbackInfo) ProtoReflect() protoreflect.Message {
	mi := &file_integralService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LockingIntegralRollbackInfo.ProtoReflect.Descriptor instead.
func (*LockingIntegralRollbackInfo) Descriptor() ([]byte, []int) {
	return file_integralService_proto_rawDescGZIP(), []int{1}
}

func (x *LockingIntegralRollbackInfo) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *LockingIntegralRollbackInfo) GetIntegral() int64 {
	if x != nil {
		return x.Integral
	}
	return 0
}

func (x *LockingIntegralRollbackInfo) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

//锁定积分生效
type LockingIntegralEffectInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid"` //积分唯一标识
}

func (x *LockingIntegralEffectInfo) Reset() {
	*x = LockingIntegralEffectInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integralService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LockingIntegralEffectInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LockingIntegralEffectInfo) ProtoMessage() {}

func (x *LockingIntegralEffectInfo) ProtoReflect() protoreflect.Message {
	mi := &file_integralService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LockingIntegralEffectInfo.ProtoReflect.Descriptor instead.
func (*LockingIntegralEffectInfo) Descriptor() ([]byte, []int) {
	return file_integralService_proto_rawDescGZIP(), []int{2}
}

func (x *LockingIntegralEffectInfo) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

//积分信息
type IntegralInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MemberId     int64  `protobuf:"varint,1,opt,name=member_id,json=memberId,proto3" json:"member_id"`            //用户ID
	Integral     int64  `protobuf:"varint,2,opt,name=integral,proto3" json:"integral"`                            //积分
	Msg          string `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg"`                                       //操作原因
	OperatorId   int64  `protobuf:"varint,4,opt,name=operator_id,json=operatorId,proto3" json:"operator_id"`      //操作员ID
	OperatorName string `protobuf:"bytes,5,opt,name=operator_name,json=operatorName,proto3" json:"operator_name"` //操作员名称
}

func (x *IntegralInfo) Reset() {
	*x = IntegralInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integralService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IntegralInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntegralInfo) ProtoMessage() {}

func (x *IntegralInfo) ProtoReflect() protoreflect.Message {
	mi := &file_integralService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntegralInfo.ProtoReflect.Descriptor instead.
func (*IntegralInfo) Descriptor() ([]byte, []int) {
	return file_integralService_proto_rawDescGZIP(), []int{3}
}

func (x *IntegralInfo) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

func (x *IntegralInfo) GetIntegral() int64 {
	if x != nil {
		return x.Integral
	}
	return 0
}

func (x *IntegralInfo) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *IntegralInfo) GetOperatorId() int64 {
	if x != nil {
		return x.OperatorId
	}
	return 0
}

func (x *IntegralInfo) GetOperatorName() string {
	if x != nil {
		return x.OperatorName
	}
	return ""
}

type IntegralRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`                                        //ID
	MemberId     int64  `protobuf:"varint,2,opt,name=member_id,json=memberId,proto3" json:"member_id"`            //用户ID
	Integral     int64  `protobuf:"varint,4,opt,name=integral,proto3" json:"integral"`                            //积分
	Msg          string `protobuf:"bytes,7,opt,name=msg,proto3" json:"msg"`                                       //操作原因
	OperatorId   int64  `protobuf:"varint,8,opt,name=operator_id,json=operatorId,proto3" json:"operator_id"`      //操作员ID
	OperatorName string `protobuf:"bytes,9,opt,name=operator_name,json=operatorName,proto3" json:"operator_name"` //操作员名称
}

func (x *IntegralRequest) Reset() {
	*x = IntegralRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integralService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IntegralRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntegralRequest) ProtoMessage() {}

func (x *IntegralRequest) ProtoReflect() protoreflect.Message {
	mi := &file_integralService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntegralRequest.ProtoReflect.Descriptor instead.
func (*IntegralRequest) Descriptor() ([]byte, []int) {
	return file_integralService_proto_rawDescGZIP(), []int{4}
}

func (x *IntegralRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *IntegralRequest) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

func (x *IntegralRequest) GetIntegral() int64 {
	if x != nil {
		return x.Integral
	}
	return 0
}

func (x *IntegralRequest) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *IntegralRequest) GetOperatorId() int64 {
	if x != nil {
		return x.OperatorId
	}
	return 0
}

func (x *IntegralRequest) GetOperatorName() string {
	if x != nil {
		return x.OperatorName
	}
	return ""
}

type IntegralResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *MemberIntegralLog `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Info   string             `protobuf:"bytes,4,opt,name=info,proto3" json:"info"`
}

func (x *IntegralResponse) Reset() {
	*x = IntegralResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integralService_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IntegralResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntegralResponse) ProtoMessage() {}

func (x *IntegralResponse) ProtoReflect() protoreflect.Message {
	mi := &file_integralService_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntegralResponse.ProtoReflect.Descriptor instead.
func (*IntegralResponse) Descriptor() ([]byte, []int) {
	return file_integralService_proto_rawDescGZIP(), []int{5}
}

func (x *IntegralResponse) GetEntity() *MemberIntegralLog {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *IntegralResponse) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

type LockingIntegralResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *IntegralLockChangeLog `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Info   string                 `protobuf:"bytes,4,opt,name=info,proto3" json:"info"`
}

func (x *LockingIntegralResponse) Reset() {
	*x = LockingIntegralResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integralService_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LockingIntegralResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LockingIntegralResponse) ProtoMessage() {}

func (x *LockingIntegralResponse) ProtoReflect() protoreflect.Message {
	mi := &file_integralService_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LockingIntegralResponse.ProtoReflect.Descriptor instead.
func (*LockingIntegralResponse) Descriptor() ([]byte, []int) {
	return file_integralService_proto_rawDescGZIP(), []int{6}
}

func (x *LockingIntegralResponse) GetEntity() *IntegralLockChangeLog {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *LockingIntegralResponse) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

var File_integralService_proto protoreflect.FileDescriptor

var file_integralService_proto_rawDesc = []byte{
	0x0a, 0x15, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x1a, 0x1e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61,
	0x6c, 0x4c, 0x6f, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x22, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x4c, 0x6f, 0x63, 0x6b, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x4c, 0x6f, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xae, 0x02, 0x0a, 0x13, 0x4c, 0x6f, 0x63, 0x6b, 0x69, 0x6e,
	0x67, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a,
	0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x62, 0x75,
	0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x2c, 0x0a, 0x12, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x62, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x54, 0x79, 0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a,
	0x0f, 0x72, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x72, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x49, 0x6e,
	0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72,
	0x61, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72,
	0x61, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x22, 0x65, 0x0a, 0x1b, 0x4c, 0x6f, 0x63, 0x6b, 0x69, 0x6e,
	0x67, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x52, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63,
	0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x69, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22, 0x2f, 0x0a,
	0x19, 0x4c, 0x6f, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c,
	0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x9f,
	0x01, 0x0a, 0x0c, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0xb2, 0x01, 0x0a, 0x0f, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12,
	0x1f, 0x0a, 0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64,
	0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x5b, 0x0a, 0x10, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x06, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x67,
	0x72, 0x61, 0x6c, 0x4c, 0x6f, 0x67, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x22, 0x66, 0x0a, 0x17, 0x4c, 0x6f, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a,
	0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61,
	0x6c, 0x4c, 0x6f, 0x63, 0x6b, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4c, 0x6f, 0x67, 0x52, 0x06,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0xc2, 0x03, 0x0a, 0x0f, 0x49,
	0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43,
	0x0a, 0x0b, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x41, 0x64, 0x64, 0x12, 0x16, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61,
	0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0b, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x55,
	0x73, 0x65, 0x12, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e,
	0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x1a, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5b, 0x0a, 0x15, 0x4c, 0x6f, 0x63, 0x6b,
	0x69, 0x6e, 0x67, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x49, 0x6e, 0x73, 0x65, 0x72,
	0x74, 0x12, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x6f, 0x63,
	0x6b, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f,
	0x1a, 0x21, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x6f, 0x63, 0x6b,
	0x69, 0x6e, 0x67, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x65, 0x0a, 0x17, 0x4c, 0x6f, 0x63, 0x6b, 0x69, 0x6e, 0x67,
	0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x52, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b,
	0x12, 0x25, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x6f, 0x63, 0x6b,
	0x69, 0x6e, 0x67, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x52, 0x6f, 0x6c, 0x6c, 0x62,
	0x61, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x21, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x4c, 0x6f, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72,
	0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x61, 0x0a, 0x15,
	0x4c, 0x6f, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x45,
	0x66, 0x66, 0x65, 0x63, 0x74, 0x12, 0x23, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x4c, 0x6f, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c,
	0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x21, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x6f, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_integralService_proto_rawDescOnce sync.Once
	file_integralService_proto_rawDescData = file_integralService_proto_rawDesc
)

func file_integralService_proto_rawDescGZIP() []byte {
	file_integralService_proto_rawDescOnce.Do(func() {
		file_integralService_proto_rawDescData = protoimpl.X.CompressGZIP(file_integralService_proto_rawDescData)
	})
	return file_integralService_proto_rawDescData
}

var file_integralService_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_integralService_proto_goTypes = []interface{}{
	(*LockingIntegralInfo)(nil),         // 0: services.LockingIntegralInfo
	(*LockingIntegralRollbackInfo)(nil), // 1: services.LockingIntegralRollbackInfo
	(*LockingIntegralEffectInfo)(nil),   // 2: services.LockingIntegralEffectInfo
	(*IntegralInfo)(nil),                // 3: services.IntegralInfo
	(*IntegralRequest)(nil),             // 4: services.IntegralRequest
	(*IntegralResponse)(nil),            // 5: services.IntegralResponse
	(*LockingIntegralResponse)(nil),     // 6: services.LockingIntegralResponse
	(*MemberIntegralLog)(nil),           // 7: services.MemberIntegralLog
	(*IntegralLockChangeLog)(nil),       // 8: services.IntegralLockChangeLog
}
var file_integralService_proto_depIdxs = []int32{
	7, // 0: services.IntegralResponse.entity:type_name -> services.MemberIntegralLog
	8, // 1: services.LockingIntegralResponse.entity:type_name -> services.IntegralLockChangeLog
	3, // 2: services.IntegralService.IntegralAdd:input_type -> services.IntegralInfo
	3, // 3: services.IntegralService.IntegralUse:input_type -> services.IntegralInfo
	0, // 4: services.IntegralService.LockingIntegralInsert:input_type -> services.LockingIntegralInfo
	1, // 5: services.IntegralService.LockingIntegralRollback:input_type -> services.LockingIntegralRollbackInfo
	2, // 6: services.IntegralService.LockingIntegralEffect:input_type -> services.LockingIntegralEffectInfo
	5, // 7: services.IntegralService.IntegralAdd:output_type -> services.IntegralResponse
	5, // 8: services.IntegralService.IntegralUse:output_type -> services.IntegralResponse
	6, // 9: services.IntegralService.LockingIntegralInsert:output_type -> services.LockingIntegralResponse
	6, // 10: services.IntegralService.LockingIntegralRollback:output_type -> services.LockingIntegralResponse
	6, // 11: services.IntegralService.LockingIntegralEffect:output_type -> services.LockingIntegralResponse
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_integralService_proto_init() }
func file_integralService_proto_init() {
	if File_integralService_proto != nil {
		return
	}
	file_memberIntegralLogService_proto_init()
	file_integralLockChangeLogService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_integralService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LockingIntegralInfo); i {
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
		file_integralService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LockingIntegralRollbackInfo); i {
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
		file_integralService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LockingIntegralEffectInfo); i {
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
		file_integralService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IntegralInfo); i {
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
		file_integralService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IntegralRequest); i {
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
		file_integralService_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IntegralResponse); i {
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
		file_integralService_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LockingIntegralResponse); i {
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
			RawDescriptor: file_integralService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_integralService_proto_goTypes,
		DependencyIndexes: file_integralService_proto_depIdxs,
		MessageInfos:      file_integralService_proto_msgTypes,
	}.Build()
	File_integralService_proto = out.File
	file_integralService_proto_rawDesc = nil
	file_integralService_proto_goTypes = nil
	file_integralService_proto_depIdxs = nil
}
