// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: vipService.proto

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

// 会员体系
type Vip struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               int32       `protobuf:"varint,1,opt,name=id,proto3" json:"id"`                                                        //ID
	Name             string      `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`                                                     //名称
	Type             string      `protobuf:"bytes,3,opt,name=type,proto3" json:"type"`                                                     //类型
	ExpireType       string      `protobuf:"bytes,4,opt,name=expire_type,json=expireType,proto3" json:"expire_type"`                       //有效期类型
	ExpireMonths     int32       `protobuf:"varint,5,opt,name=expire_months,json=expireMonths,proto3" json:"expire_months"`                //有效期月数
	OpenMethod       string      `protobuf:"bytes,6,opt,name=open_method,json=openMethod,proto3" json:"open_method"`                       //开通方式
	UpMethod         string      `protobuf:"bytes,7,opt,name=up_method,json=upMethod,proto3" json:"up_method"`                             //升级方式
	GiftMethod       string      `protobuf:"bytes,8,opt,name=gift_method,json=giftMethod,proto3" json:"gift_method"`                       //礼包发放方式
	LevelCalculate   string      `protobuf:"bytes,9,opt,name=level_calculate,json=levelCalculate,proto3" json:"level_calculate"`           //等级计算时间
	ValidOrderAmount int64       `protobuf:"varint,10,opt,name=valid_order_amount,json=validOrderAmount,proto3" json:"valid_order_amount"` //有效订单金额
	IsNeedSync       string      `protobuf:"bytes,11,opt,name=is_need_sync,json=isNeedSync,proto3" json:"is_need_sync"`                    //需要同步更新
	GrowthRules      string      `protobuf:"bytes,12,opt,name=growth_rules,json=growthRules,proto3" json:"growth_rules"`                   //成长规则
	Status           string      `protobuf:"bytes,13,opt,name=status,proto3" json:"status"`                                                //状态
	CreatedAt        int64       `protobuf:"varint,14,opt,name=created_at,json=createdAt,proto3" json:"created_at"`                        //创建时间
	UpdatedAt        int64       `protobuf:"varint,15,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`                        //修改时间
	Levels           []*VipLevel `protobuf:"bytes,16,rep,name=levels,proto3" json:"levels"`
}

func (x *Vip) Reset() {
	*x = Vip{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vipService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vip) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vip) ProtoMessage() {}

func (x *Vip) ProtoReflect() protoreflect.Message {
	mi := &file_vipService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vip.ProtoReflect.Descriptor instead.
func (*Vip) Descriptor() ([]byte, []int) {
	return file_vipService_proto_rawDescGZIP(), []int{0}
}

func (x *Vip) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Vip) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Vip) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Vip) GetExpireType() string {
	if x != nil {
		return x.ExpireType
	}
	return ""
}

func (x *Vip) GetExpireMonths() int32 {
	if x != nil {
		return x.ExpireMonths
	}
	return 0
}

func (x *Vip) GetOpenMethod() string {
	if x != nil {
		return x.OpenMethod
	}
	return ""
}

func (x *Vip) GetUpMethod() string {
	if x != nil {
		return x.UpMethod
	}
	return ""
}

func (x *Vip) GetGiftMethod() string {
	if x != nil {
		return x.GiftMethod
	}
	return ""
}

func (x *Vip) GetLevelCalculate() string {
	if x != nil {
		return x.LevelCalculate
	}
	return ""
}

func (x *Vip) GetValidOrderAmount() int64 {
	if x != nil {
		return x.ValidOrderAmount
	}
	return 0
}

func (x *Vip) GetIsNeedSync() string {
	if x != nil {
		return x.IsNeedSync
	}
	return ""
}

func (x *Vip) GetGrowthRules() string {
	if x != nil {
		return x.GrowthRules
	}
	return ""
}

func (x *Vip) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Vip) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Vip) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *Vip) GetLevels() []*VipLevel {
	if x != nil {
		return x.Levels
	}
	return nil
}

// 会员体系请求参数
type VipRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Top       int32    `protobuf:"varint,1,opt,name=top,proto3" json:"top"`
	Paged     int64    `protobuf:"varint,2,opt,name=paged,proto3" json:"paged"`
	PageSize  int64    `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Keywords  string   `protobuf:"bytes,4,opt,name=keywords,proto3" json:"keywords"`
	Sort      []string `protobuf:"bytes,5,rep,name=sort,proto3" json:"sort"`
	DateRange []string `protobuf:"bytes,6,rep,name=date_range,json=dateRange,proto3" json:"date_range"`
	Ids       []int32  `protobuf:"varint,7,rep,packed,name=ids,proto3" json:"ids"`
	Id        int32    `protobuf:"varint,8,opt,name=id,proto3" json:"id"`
	//以下为自定义参数
	Name       string `protobuf:"bytes,11,opt,name=name,proto3" json:"name"`                                 //名称
	Type       string `protobuf:"bytes,12,opt,name=type,proto3" json:"type"`                                 //类型
	ExpireType string `protobuf:"bytes,13,opt,name=expire_type,json=expireType,proto3" json:"expire_type"`   //有效期类型
	IsNeedSync string `protobuf:"bytes,14,opt,name=is_need_sync,json=isNeedSync,proto3" json:"is_need_sync"` //需要同步更新
	Status     string `protobuf:"bytes,15,opt,name=status,proto3" json:"status"`                             //状态
}

func (x *VipRequest) Reset() {
	*x = VipRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vipService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VipRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VipRequest) ProtoMessage() {}

func (x *VipRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vipService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VipRequest.ProtoReflect.Descriptor instead.
func (*VipRequest) Descriptor() ([]byte, []int) {
	return file_vipService_proto_rawDescGZIP(), []int{1}
}

func (x *VipRequest) GetTop() int32 {
	if x != nil {
		return x.Top
	}
	return 0
}

func (x *VipRequest) GetPaged() int64 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *VipRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *VipRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *VipRequest) GetSort() []string {
	if x != nil {
		return x.Sort
	}
	return nil
}

func (x *VipRequest) GetDateRange() []string {
	if x != nil {
		return x.DateRange
	}
	return nil
}

func (x *VipRequest) GetIds() []int32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *VipRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *VipRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *VipRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *VipRequest) GetExpireType() string {
	if x != nil {
		return x.ExpireType
	}
	return ""
}

func (x *VipRequest) GetIsNeedSync() string {
	if x != nil {
		return x.IsNeedSync
	}
	return ""
}

func (x *VipRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

// 会员体系响应数据
type VipResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg    string        `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Entity *Vip          `protobuf:"bytes,3,opt,name=entity,proto3" json:"entity"`
	Items  []*Vip        `protobuf:"bytes,4,rep,name=items,proto3" json:"items"`
}

func (x *VipResponse) Reset() {
	*x = VipResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vipService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VipResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VipResponse) ProtoMessage() {}

func (x *VipResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vipService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VipResponse.ProtoReflect.Descriptor instead.
func (*VipResponse) Descriptor() ([]byte, []int) {
	return file_vipService_proto_rawDescGZIP(), []int{2}
}

func (x *VipResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *VipResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *VipResponse) GetEntity() *Vip {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *VipResponse) GetItems() []*Vip {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_vipService_proto protoreflect.FileDescriptor

var file_vipService_proto_rawDesc = []byte{
	0x0a, 0x10, 0x76, 0x69, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x15, 0x76, 0x69, 0x70, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x80, 0x04, 0x0a, 0x03, 0x56, 0x69, 0x70, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x5f, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x73, 0x12, 0x1f, 0x0a, 0x0b,
	0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x75, 0x70, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x70, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x67, 0x69,
	0x66, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x67, 0x69, 0x66, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x5f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x43, 0x61, 0x6c, 0x63, 0x75,
	0x6c, 0x61, 0x74, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x10, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x6e, 0x65, 0x65, 0x64, 0x5f, 0x73, 0x79,
	0x6e, 0x63, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x73, 0x4e, 0x65, 0x65, 0x64,
	0x53, 0x79, 0x6e, 0x63, 0x12, 0x21, 0x0a, 0x0c, 0x67, 0x72, 0x6f, 0x77, 0x74, 0x68, 0x5f, 0x72,
	0x75, 0x6c, 0x65, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x67, 0x72, 0x6f, 0x77,
	0x74, 0x68, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x2a, 0x0a,
	0x06, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x73, 0x18, 0x10, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x56, 0x69, 0x70, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x52, 0x06, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x73, 0x22, 0xc5, 0x02, 0x0a, 0x0a, 0x56, 0x69,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x6f, 0x70, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x74, 0x6f, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61,
	0x67, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72,
	0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x09, 0x64, 0x61, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x69, 0x64, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x05, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x6e, 0x65,
	0x65, 0x64, 0x5f, 0x73, 0x79, 0x6e, 0x63, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69,
	0x73, 0x4e, 0x65, 0x65, 0x64, 0x53, 0x79, 0x6e, 0x63, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x90, 0x01, 0x0a, 0x0b, 0x56, 0x69, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65,
	0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x56, 0x69, 0x70, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12,
	0x23, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x56, 0x69, 0x70, 0x52, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x32, 0xd6, 0x01, 0x0a, 0x0a, 0x56, 0x69, 0x70, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x12, 0x0d, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x56, 0x69, 0x70, 0x1a, 0x15, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x56, 0x69, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x0d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x56, 0x69, 0x70, 0x1a, 0x15,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x56, 0x69, 0x70, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0d,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x56, 0x69, 0x70, 0x1a, 0x15, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x56, 0x69, 0x70, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x56, 0x69, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x56, 0x69, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a,
	0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vipService_proto_rawDescOnce sync.Once
	file_vipService_proto_rawDescData = file_vipService_proto_rawDesc
)

func file_vipService_proto_rawDescGZIP() []byte {
	file_vipService_proto_rawDescOnce.Do(func() {
		file_vipService_proto_rawDescData = protoimpl.X.CompressGZIP(file_vipService_proto_rawDescData)
	})
	return file_vipService_proto_rawDescData
}

var file_vipService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_vipService_proto_goTypes = []interface{}{
	(*Vip)(nil),          // 0: services.Vip
	(*VipRequest)(nil),   // 1: services.VipRequest
	(*VipResponse)(nil),  // 2: services.VipResponse
	(*VipLevel)(nil),     // 3: services.VipLevel
	(*common.Pager)(nil), // 4: common.Pager
}
var file_vipService_proto_depIdxs = []int32{
	3, // 0: services.Vip.levels:type_name -> services.VipLevel
	4, // 1: services.VipResponse.pager:type_name -> common.Pager
	0, // 2: services.VipResponse.entity:type_name -> services.Vip
	0, // 3: services.VipResponse.items:type_name -> services.Vip
	0, // 4: services.VipService.Switch:input_type -> services.Vip
	0, // 5: services.VipService.Update:input_type -> services.Vip
	0, // 6: services.VipService.Get:input_type -> services.Vip
	1, // 7: services.VipService.List:input_type -> services.VipRequest
	2, // 8: services.VipService.Switch:output_type -> services.VipResponse
	2, // 9: services.VipService.Update:output_type -> services.VipResponse
	2, // 10: services.VipService.Get:output_type -> services.VipResponse
	2, // 11: services.VipService.List:output_type -> services.VipResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_vipService_proto_init() }
func file_vipService_proto_init() {
	if File_vipService_proto != nil {
		return
	}
	file_vipLevelService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_vipService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vip); i {
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
		file_vipService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VipRequest); i {
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
		file_vipService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VipResponse); i {
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
			RawDescriptor: file_vipService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_vipService_proto_goTypes,
		DependencyIndexes: file_vipService_proto_depIdxs,
		MessageInfos:      file_vipService_proto_msgTypes,
	}.Build()
	File_vipService_proto = out.File
	file_vipService_proto_rawDesc = nil
	file_vipService_proto_goTypes = nil
	file_vipService_proto_depIdxs = nil
}
