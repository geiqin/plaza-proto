// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: vipBenefitService.proto

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

// VIP权益
type VipBenefit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`                                        //ID
	Code         string `protobuf:"bytes,2,opt,name=code,proto3" json:"code"`                                     //权益标识
	Name         string `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`                                     //权益名称
	Type         string `protobuf:"bytes,4,opt,name=type,proto3" json:"type"`                                     //权益类型
	Industry     string `protobuf:"bytes,5,opt,name=industry,proto3" json:"industry"`                             //所属行业
	IconUrl      string `protobuf:"bytes,6,opt,name=icon_url,json=iconUrl,proto3" json:"icon_url"`                //权益图标
	PriceValue   int64  `protobuf:"varint,7,opt,name=price_value,json=priceValue,proto3" json:"price_value"`      //权益价值
	IsInit       string `protobuf:"bytes,8,opt,name=is_init,json=isInit,proto3" json:"is_init"`                   //初始化
	VerifyMethod string `protobuf:"bytes,9,opt,name=verify_method,json=verifyMethod,proto3" json:"verify_method"` //核销方式
	Desc         string `protobuf:"bytes,10,opt,name=desc,proto3" json:"desc"`                                    //权益描述
	Status       string `protobuf:"bytes,11,opt,name=status,proto3" json:"status"`                                //状态
	CreatedAt    string `protobuf:"bytes,12,opt,name=created_at,json=createdAt,proto3" json:"created_at"`         //创建日期
	UpdatedAt    string `protobuf:"bytes,13,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`         //修改日期
}

func (x *VipBenefit) Reset() {
	*x = VipBenefit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vipBenefitService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VipBenefit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VipBenefit) ProtoMessage() {}

func (x *VipBenefit) ProtoReflect() protoreflect.Message {
	mi := &file_vipBenefitService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VipBenefit.ProtoReflect.Descriptor instead.
func (*VipBenefit) Descriptor() ([]byte, []int) {
	return file_vipBenefitService_proto_rawDescGZIP(), []int{0}
}

func (x *VipBenefit) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *VipBenefit) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *VipBenefit) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *VipBenefit) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *VipBenefit) GetIndustry() string {
	if x != nil {
		return x.Industry
	}
	return ""
}

func (x *VipBenefit) GetIconUrl() string {
	if x != nil {
		return x.IconUrl
	}
	return ""
}

func (x *VipBenefit) GetPriceValue() int64 {
	if x != nil {
		return x.PriceValue
	}
	return 0
}

func (x *VipBenefit) GetIsInit() string {
	if x != nil {
		return x.IsInit
	}
	return ""
}

func (x *VipBenefit) GetVerifyMethod() string {
	if x != nil {
		return x.VerifyMethod
	}
	return ""
}

func (x *VipBenefit) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *VipBenefit) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *VipBenefit) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *VipBenefit) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

// VIP权益请求参数
type VipBenefitRequest struct {
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
	Code         string `protobuf:"bytes,11,opt,name=code,proto3" json:"code"`                                     //权益标识
	Name         string `protobuf:"bytes,12,opt,name=name,proto3" json:"name"`                                     //权益名称
	Type         string `protobuf:"bytes,13,opt,name=type,proto3" json:"type"`                                     //权益类型
	Industry     string `protobuf:"bytes,14,opt,name=industry,proto3" json:"industry"`                             //所属行业
	IsInit       string `protobuf:"bytes,15,opt,name=is_init,json=isInit,proto3" json:"is_init"`                   //初始化
	VerifyMethod string `protobuf:"bytes,16,opt,name=verify_method,json=verifyMethod,proto3" json:"verify_method"` //核销方式
	Status       string `protobuf:"bytes,17,opt,name=status,proto3" json:"status"`                                 //状态
}

func (x *VipBenefitRequest) Reset() {
	*x = VipBenefitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vipBenefitService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VipBenefitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VipBenefitRequest) ProtoMessage() {}

func (x *VipBenefitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vipBenefitService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VipBenefitRequest.ProtoReflect.Descriptor instead.
func (*VipBenefitRequest) Descriptor() ([]byte, []int) {
	return file_vipBenefitService_proto_rawDescGZIP(), []int{1}
}

func (x *VipBenefitRequest) GetTop() int32 {
	if x != nil {
		return x.Top
	}
	return 0
}

func (x *VipBenefitRequest) GetPaged() int64 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *VipBenefitRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *VipBenefitRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *VipBenefitRequest) GetSort() []string {
	if x != nil {
		return x.Sort
	}
	return nil
}

func (x *VipBenefitRequest) GetDateRange() []string {
	if x != nil {
		return x.DateRange
	}
	return nil
}

func (x *VipBenefitRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *VipBenefitRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *VipBenefitRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *VipBenefitRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *VipBenefitRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *VipBenefitRequest) GetIndustry() string {
	if x != nil {
		return x.Industry
	}
	return ""
}

func (x *VipBenefitRequest) GetIsInit() string {
	if x != nil {
		return x.IsInit
	}
	return ""
}

func (x *VipBenefitRequest) GetVerifyMethod() string {
	if x != nil {
		return x.VerifyMethod
	}
	return ""
}

func (x *VipBenefitRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

// VIP权益响应数据
type VipBenefitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *VipBenefit   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*VipBenefit `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Msg    string        `protobuf:"bytes,4,opt,name=msg,proto3" json:"msg"`
}

func (x *VipBenefitResponse) Reset() {
	*x = VipBenefitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vipBenefitService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VipBenefitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VipBenefitResponse) ProtoMessage() {}

func (x *VipBenefitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vipBenefitService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VipBenefitResponse.ProtoReflect.Descriptor instead.
func (*VipBenefitResponse) Descriptor() ([]byte, []int) {
	return file_vipBenefitService_proto_rawDescGZIP(), []int{2}
}

func (x *VipBenefitResponse) GetEntity() *VipBenefit {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *VipBenefitResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *VipBenefitResponse) GetItems() []*VipBenefit {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *VipBenefitResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_vipBenefitService_proto protoreflect.FileDescriptor

var file_vipBenefitService_proto_rawDesc = []byte{
	0x0a, 0x17, 0x76, 0x69, 0x70, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd8, 0x02, 0x0a, 0x0a, 0x56, 0x69, 0x70, 0x42, 0x65,
	0x6e, 0x65, 0x66, 0x69, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x12, 0x19, 0x0a,
	0x08, 0x69, 0x63, 0x6f, 0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x69, 0x63, 0x6f, 0x6e, 0x55, 0x72, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f,
	0x69, 0x6e, 0x69, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x73, 0x49, 0x6e,
	0x69, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x76, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0xf7, 0x02, 0x0a, 0x11, 0x56, 0x69, 0x70, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x6f, 0x70, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x74, 0x6f, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67,
	0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x09, 0x64, 0x61, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x69,
	0x64, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x64,
	0x75, 0x73, 0x74, 0x72, 0x79, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6e, 0x64,
	0x75, 0x73, 0x74, 0x72, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x69, 0x6e, 0x69, 0x74,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x73, 0x49, 0x6e, 0x69, 0x74, 0x12, 0x23,
	0x0a, 0x0d, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18,
	0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x11, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xa5, 0x01, 0x0a, 0x12,
	0x56, 0x69, 0x70, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x56, 0x69,
	0x70, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05,
	0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x56, 0x69, 0x70, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x32, 0x9c, 0x03, 0x0a, 0x11, 0x56, 0x69, 0x70, 0x42, 0x65, 0x6e, 0x65, 0x66,
	0x69, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x06, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x56,
	0x69, 0x70, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x56, 0x69, 0x70, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x06, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x56,
	0x69, 0x70, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x56, 0x69, 0x70, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x06, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x56,
	0x69, 0x70, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x56, 0x69, 0x70, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x03, 0x47, 0x65, 0x74,
	0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x56, 0x69, 0x70, 0x42,
	0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x56, 0x69, 0x70, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x12, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x56, 0x69, 0x70, 0x42,
	0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x56, 0x69, 0x70, 0x42, 0x65, 0x6e, 0x65,
	0x66, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a,
	0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x56, 0x69, 0x70, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x56, 0x69,
	0x70, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vipBenefitService_proto_rawDescOnce sync.Once
	file_vipBenefitService_proto_rawDescData = file_vipBenefitService_proto_rawDesc
)

func file_vipBenefitService_proto_rawDescGZIP() []byte {
	file_vipBenefitService_proto_rawDescOnce.Do(func() {
		file_vipBenefitService_proto_rawDescData = protoimpl.X.CompressGZIP(file_vipBenefitService_proto_rawDescData)
	})
	return file_vipBenefitService_proto_rawDescData
}

var file_vipBenefitService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_vipBenefitService_proto_goTypes = []interface{}{
	(*VipBenefit)(nil),         // 0: services.VipBenefit
	(*VipBenefitRequest)(nil),  // 1: services.VipBenefitRequest
	(*VipBenefitResponse)(nil), // 2: services.VipBenefitResponse
	(*common.Pager)(nil),       // 3: common.Pager
}
var file_vipBenefitService_proto_depIdxs = []int32{
	0, // 0: services.VipBenefitResponse.entity:type_name -> services.VipBenefit
	3, // 1: services.VipBenefitResponse.pager:type_name -> common.Pager
	0, // 2: services.VipBenefitResponse.items:type_name -> services.VipBenefit
	0, // 3: services.VipBenefitService.Create:input_type -> services.VipBenefit
	0, // 4: services.VipBenefitService.Update:input_type -> services.VipBenefit
	0, // 5: services.VipBenefitService.Delete:input_type -> services.VipBenefit
	0, // 6: services.VipBenefitService.Get:input_type -> services.VipBenefit
	1, // 7: services.VipBenefitService.Search:input_type -> services.VipBenefitRequest
	1, // 8: services.VipBenefitService.List:input_type -> services.VipBenefitRequest
	2, // 9: services.VipBenefitService.Create:output_type -> services.VipBenefitResponse
	2, // 10: services.VipBenefitService.Update:output_type -> services.VipBenefitResponse
	2, // 11: services.VipBenefitService.Delete:output_type -> services.VipBenefitResponse
	2, // 12: services.VipBenefitService.Get:output_type -> services.VipBenefitResponse
	2, // 13: services.VipBenefitService.Search:output_type -> services.VipBenefitResponse
	2, // 14: services.VipBenefitService.List:output_type -> services.VipBenefitResponse
	9, // [9:15] is the sub-list for method output_type
	3, // [3:9] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_vipBenefitService_proto_init() }
func file_vipBenefitService_proto_init() {
	if File_vipBenefitService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vipBenefitService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VipBenefit); i {
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
		file_vipBenefitService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VipBenefitRequest); i {
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
		file_vipBenefitService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VipBenefitResponse); i {
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
			RawDescriptor: file_vipBenefitService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_vipBenefitService_proto_goTypes,
		DependencyIndexes: file_vipBenefitService_proto_depIdxs,
		MessageInfos:      file_vipBenefitService_proto_msgTypes,
	}.Build()
	File_vipBenefitService_proto = out.File
	file_vipBenefitService_proto_rawDesc = nil
	file_vipBenefitService_proto_goTypes = nil
	file_vipBenefitService_proto_depIdxs = nil
}
