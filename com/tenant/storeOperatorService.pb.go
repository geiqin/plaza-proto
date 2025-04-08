// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: storeOperatorService.proto

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

// 运营人员
type StoreOperator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`                                //ID
	StoreId   int64  `protobuf:"varint,2,opt,name=store_id,json=storeId,proto3" json:"store_id"`       //店铺ID
	RealName  string `protobuf:"bytes,3,opt,name=real_name,json=realName,proto3" json:"real_name"`     //姓名
	Mobile    string `protobuf:"bytes,4,opt,name=mobile,proto3" json:"mobile"`                         //手机号码
	Wechat    string `protobuf:"bytes,5,opt,name=wechat,proto3" json:"wechat"`                         //微信账号
	Email     string `protobuf:"bytes,6,opt,name=email,proto3" json:"email"`                           //邮件
	GeiqinId  int64  `protobuf:"varint,7,opt,name=geiqin_id,json=geiqinId,proto3" json:"geiqin_id"`    //给亲账号
	UpdatedAt int64  `protobuf:"varint,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"` //创建时间
	CreatedAt int64  `protobuf:"varint,9,opt,name=created_at,json=createdAt,proto3" json:"created_at"` //修改时间
}

func (x *StoreOperator) Reset() {
	*x = StoreOperator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storeOperatorService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreOperator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreOperator) ProtoMessage() {}

func (x *StoreOperator) ProtoReflect() protoreflect.Message {
	mi := &file_storeOperatorService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreOperator.ProtoReflect.Descriptor instead.
func (*StoreOperator) Descriptor() ([]byte, []int) {
	return file_storeOperatorService_proto_rawDescGZIP(), []int{0}
}

func (x *StoreOperator) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *StoreOperator) GetStoreId() int64 {
	if x != nil {
		return x.StoreId
	}
	return 0
}

func (x *StoreOperator) GetRealName() string {
	if x != nil {
		return x.RealName
	}
	return ""
}

func (x *StoreOperator) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *StoreOperator) GetWechat() string {
	if x != nil {
		return x.Wechat
	}
	return ""
}

func (x *StoreOperator) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *StoreOperator) GetGeiqinId() int64 {
	if x != nil {
		return x.GeiqinId
	}
	return 0
}

func (x *StoreOperator) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *StoreOperator) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

// 运营人员请求参数
type StoreOperatorRequest struct {
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
	StoreId  int64  `protobuf:"varint,11,opt,name=store_id,json=storeId,proto3" json:"store_id"`    //店铺ID
	RealName string `protobuf:"bytes,12,opt,name=real_name,json=realName,proto3" json:"real_name"`  //姓名
	Mobile   string `protobuf:"bytes,13,opt,name=mobile,proto3" json:"mobile"`                      //手机号码
	Wechat   string `protobuf:"bytes,14,opt,name=wechat,proto3" json:"wechat"`                      //微信账号
	GeiqinId int64  `protobuf:"varint,15,opt,name=geiqin_id,json=geiqinId,proto3" json:"geiqin_id"` //给亲账号
}

func (x *StoreOperatorRequest) Reset() {
	*x = StoreOperatorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storeOperatorService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreOperatorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreOperatorRequest) ProtoMessage() {}

func (x *StoreOperatorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_storeOperatorService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreOperatorRequest.ProtoReflect.Descriptor instead.
func (*StoreOperatorRequest) Descriptor() ([]byte, []int) {
	return file_storeOperatorService_proto_rawDescGZIP(), []int{1}
}

func (x *StoreOperatorRequest) GetTop() int32 {
	if x != nil {
		return x.Top
	}
	return 0
}

func (x *StoreOperatorRequest) GetPaged() int64 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *StoreOperatorRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *StoreOperatorRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *StoreOperatorRequest) GetSorts() []string {
	if x != nil {
		return x.Sorts
	}
	return nil
}

func (x *StoreOperatorRequest) GetDateRange() []string {
	if x != nil {
		return x.DateRange
	}
	return nil
}

func (x *StoreOperatorRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *StoreOperatorRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *StoreOperatorRequest) GetStoreId() int64 {
	if x != nil {
		return x.StoreId
	}
	return 0
}

func (x *StoreOperatorRequest) GetRealName() string {
	if x != nil {
		return x.RealName
	}
	return ""
}

func (x *StoreOperatorRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *StoreOperatorRequest) GetWechat() string {
	if x != nil {
		return x.Wechat
	}
	return ""
}

func (x *StoreOperatorRequest) GetGeiqinId() int64 {
	if x != nil {
		return x.GeiqinId
	}
	return 0
}

// 运营人员响应数据
type StoreOperatorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg    string           `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg"`
	Pager  *common.Pager    `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Entity *StoreOperator   `protobuf:"bytes,3,opt,name=entity,proto3" json:"entity"`
	Items  []*StoreOperator `protobuf:"bytes,4,rep,name=items,proto3" json:"items"`
}

func (x *StoreOperatorResponse) Reset() {
	*x = StoreOperatorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storeOperatorService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreOperatorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreOperatorResponse) ProtoMessage() {}

func (x *StoreOperatorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_storeOperatorService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreOperatorResponse.ProtoReflect.Descriptor instead.
func (*StoreOperatorResponse) Descriptor() ([]byte, []int) {
	return file_storeOperatorService_proto_rawDescGZIP(), []int{2}
}

func (x *StoreOperatorResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *StoreOperatorResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *StoreOperatorResponse) GetEntity() *StoreOperator {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *StoreOperatorResponse) GetItems() []*StoreOperator {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_storeOperatorService_proto protoreflect.FileDescriptor

var file_storeOperatorService_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf8, 0x01, 0x0a, 0x0d, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x61, 0x6c, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x61, 0x6c, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x77,
	0x65, 0x63, 0x68, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x77, 0x65, 0x63,
	0x68, 0x61, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x65, 0x69,
	0x71, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x67, 0x65,
	0x69, 0x71, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0xd3, 0x02, 0x0a, 0x14, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
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
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x61, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x67, 0x65, 0x69, 0x71, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x67, 0x65, 0x69, 0x71, 0x69, 0x6e, 0x49, 0x64, 0x22, 0xae, 0x01, 0x0a, 0x15, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50,
	0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x2f, 0x0a, 0x06, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x2d, 0x0a, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x32, 0xfb, 0x02, 0x0a, 0x14,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x04, 0x53, 0x61, 0x76, 0x65, 0x12, 0x17, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x1a, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x1a, 0x1f, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41,
	0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x1a, 0x1f,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x4b, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1e, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49,
	0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_storeOperatorService_proto_rawDescOnce sync.Once
	file_storeOperatorService_proto_rawDescData = file_storeOperatorService_proto_rawDesc
)

func file_storeOperatorService_proto_rawDescGZIP() []byte {
	file_storeOperatorService_proto_rawDescOnce.Do(func() {
		file_storeOperatorService_proto_rawDescData = protoimpl.X.CompressGZIP(file_storeOperatorService_proto_rawDescData)
	})
	return file_storeOperatorService_proto_rawDescData
}

var file_storeOperatorService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_storeOperatorService_proto_goTypes = []interface{}{
	(*StoreOperator)(nil),         // 0: services.StoreOperator
	(*StoreOperatorRequest)(nil),  // 1: services.StoreOperatorRequest
	(*StoreOperatorResponse)(nil), // 2: services.StoreOperatorResponse
	(*common.Pager)(nil),          // 3: common.Pager
}
var file_storeOperatorService_proto_depIdxs = []int32{
	3, // 0: services.StoreOperatorResponse.pager:type_name -> common.Pager
	0, // 1: services.StoreOperatorResponse.entity:type_name -> services.StoreOperator
	0, // 2: services.StoreOperatorResponse.items:type_name -> services.StoreOperator
	0, // 3: services.StoreOperatorService.Save:input_type -> services.StoreOperator
	0, // 4: services.StoreOperatorService.Delete:input_type -> services.StoreOperator
	0, // 5: services.StoreOperatorService.Get:input_type -> services.StoreOperator
	1, // 6: services.StoreOperatorService.Search:input_type -> services.StoreOperatorRequest
	1, // 7: services.StoreOperatorService.List:input_type -> services.StoreOperatorRequest
	2, // 8: services.StoreOperatorService.Save:output_type -> services.StoreOperatorResponse
	2, // 9: services.StoreOperatorService.Delete:output_type -> services.StoreOperatorResponse
	2, // 10: services.StoreOperatorService.Get:output_type -> services.StoreOperatorResponse
	2, // 11: services.StoreOperatorService.Search:output_type -> services.StoreOperatorResponse
	2, // 12: services.StoreOperatorService.List:output_type -> services.StoreOperatorResponse
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_storeOperatorService_proto_init() }
func file_storeOperatorService_proto_init() {
	if File_storeOperatorService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_storeOperatorService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreOperator); i {
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
		file_storeOperatorService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreOperatorRequest); i {
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
		file_storeOperatorService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreOperatorResponse); i {
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
			RawDescriptor: file_storeOperatorService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_storeOperatorService_proto_goTypes,
		DependencyIndexes: file_storeOperatorService_proto_depIdxs,
		MessageInfos:      file_storeOperatorService_proto_msgTypes,
	}.Build()
	File_storeOperatorService_proto = out.File
	file_storeOperatorService_proto_rawDesc = nil
	file_storeOperatorService_proto_goTypes = nil
	file_storeOperatorService_proto_depIdxs = nil
}
