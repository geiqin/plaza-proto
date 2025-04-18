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

// 积分管理
type Integral struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                     int32               `protobuf:"varint,1,opt,name=id,proto3" json:"id"`                                                                          //ID
	Name                   string              `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`                                                                       //积分名称
	IsProtected            string              `protobuf:"bytes,3,opt,name=is_protected,json=isProtected,proto3" json:"is_protected"`                                      //开启积分保护
	ProtectedDays          int32               `protobuf:"varint,4,opt,name=protected_days,json=protectedDays,proto3" json:"protected_days"`                               //积分保护期天数
	DeductAmount           int64               `protobuf:"varint,5,opt,name=deduct_amount,json=deductAmount,proto3" json:"deduct_amount"`                                  //抵现比例
	GivingType             string              `protobuf:"bytes,6,opt,name=giving_type,json=givingType,proto3" json:"giving_type"`                                         //发放时机
	IsDailyMaxLimit        string              `protobuf:"bytes,7,opt,name=is_daily_max_limit,json=isDailyMaxLimit,proto3" json:"is_daily_max_limit"`                      //开启获取上限
	DailyMaxLimit          int64               `protobuf:"varint,8,opt,name=daily_max_limit,json=dailyMaxLimit,proto3" json:"daily_max_limit"`                             //消费积分获取上限
	IsDeductRule           string              `protobuf:"bytes,9,opt,name=is_deduct_rule,json=isDeductRule,proto3" json:"is_deduct_rule"`                                 //开启抵现规则
	DeductRule             *IntegralDeductRule `protobuf:"bytes,10,opt,name=deduct_rule,json=deductRule,proto3" json:"deduct_rule"`                                        //成长规则
	CleanType              string              `protobuf:"bytes,11,opt,name=clean_type,json=cleanType,proto3" json:"clean_type"`                                           //清零类型
	CleanYear              int32               `protobuf:"varint,12,opt,name=clean_year,json=cleanYear,proto3" json:"clean_year"`                                          //清零每年
	CleanMonth             int32               `protobuf:"varint,13,opt,name=clean_month,json=cleanMonth,proto3" json:"clean_month"`                                       //清零每月
	CleanDay               int32               `protobuf:"varint,14,opt,name=clean_day,json=cleanDay,proto3" json:"clean_day"`                                             //清零每天
	TotalSpentIntegral     int64               `protobuf:"varint,15,opt,name=total_spent_integral,json=totalSpentIntegral,proto3" json:"total_spent_integral"`             //累计消耗积分
	TotalEarnedIntegral    int64               `protobuf:"varint,16,opt,name=total_earned_integral,json=totalEarnedIntegral,proto3" json:"total_earned_integral"`          //累计发放积分
	TotalProtectedIntegral int64               `protobuf:"varint,17,opt,name=total_protected_integral,json=totalProtectedIntegral,proto3" json:"total_protected_integral"` //保护期内积分
	Status                 string              `protobuf:"bytes,18,opt,name=status,proto3" json:"status"`                                                                  //启用状态
	CreatedAt              int64               `protobuf:"varint,19,opt,name=created_at,json=createdAt,proto3" json:"created_at"`                                          //创建时间
	UpdatedAt              int64               `protobuf:"varint,20,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`                                          //修改时间
}

func (x *Integral) Reset() {
	*x = Integral{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integralService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Integral) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Integral) ProtoMessage() {}

func (x *Integral) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Integral.ProtoReflect.Descriptor instead.
func (*Integral) Descriptor() ([]byte, []int) {
	return file_integralService_proto_rawDescGZIP(), []int{0}
}

func (x *Integral) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Integral) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Integral) GetIsProtected() string {
	if x != nil {
		return x.IsProtected
	}
	return ""
}

func (x *Integral) GetProtectedDays() int32 {
	if x != nil {
		return x.ProtectedDays
	}
	return 0
}

func (x *Integral) GetDeductAmount() int64 {
	if x != nil {
		return x.DeductAmount
	}
	return 0
}

func (x *Integral) GetGivingType() string {
	if x != nil {
		return x.GivingType
	}
	return ""
}

func (x *Integral) GetIsDailyMaxLimit() string {
	if x != nil {
		return x.IsDailyMaxLimit
	}
	return ""
}

func (x *Integral) GetDailyMaxLimit() int64 {
	if x != nil {
		return x.DailyMaxLimit
	}
	return 0
}

func (x *Integral) GetIsDeductRule() string {
	if x != nil {
		return x.IsDeductRule
	}
	return ""
}

func (x *Integral) GetDeductRule() *IntegralDeductRule {
	if x != nil {
		return x.DeductRule
	}
	return nil
}

func (x *Integral) GetCleanType() string {
	if x != nil {
		return x.CleanType
	}
	return ""
}

func (x *Integral) GetCleanYear() int32 {
	if x != nil {
		return x.CleanYear
	}
	return 0
}

func (x *Integral) GetCleanMonth() int32 {
	if x != nil {
		return x.CleanMonth
	}
	return 0
}

func (x *Integral) GetCleanDay() int32 {
	if x != nil {
		return x.CleanDay
	}
	return 0
}

func (x *Integral) GetTotalSpentIntegral() int64 {
	if x != nil {
		return x.TotalSpentIntegral
	}
	return 0
}

func (x *Integral) GetTotalEarnedIntegral() int64 {
	if x != nil {
		return x.TotalEarnedIntegral
	}
	return 0
}

func (x *Integral) GetTotalProtectedIntegral() int64 {
	if x != nil {
		return x.TotalProtectedIntegral
	}
	return 0
}

func (x *Integral) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Integral) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Integral) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

// 抵现规则
type IntegralDeductRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeductAmountLimit    string     `protobuf:"bytes,1,opt,name=deduct_amount_limit,json=deductAmountLimit,proto3" json:"deduct_amount_limit"`             //抵现金额上限
	DeductAmountLimitVal int64      `protobuf:"varint,2,opt,name=deduct_amount_limit_val,json=deductAmountLimitVal,proto3" json:"deduct_amount_limit_val"` //抵现金额上限值
	OrderAmountThreshold string     `protobuf:"bytes,3,opt,name=order_amount_threshold,json=orderAmountThreshold,proto3" json:"order_amount_threshold"`    //订单金额门槛
	MinOrderAmount       int64      `protobuf:"varint,4,opt,name=min_order_amount,json=minOrderAmount,proto3" json:"min_order_amount"`                     //订单金额最低为X元时可抵现
	SpuRangeType         string     `protobuf:"bytes,5,opt,name=spu_range_type,json=spuRangeType,proto3" json:"spu_range_type"`                            //适用商品类型
	SpuRangeIds          []int64    `protobuf:"varint,6,rep,packed,name=spu_range_ids,json=spuRangeIds,proto3" json:"spu_range_ids"`                       //适用商品IDS
	SpuRanges            []*SpuInfo `protobuf:"bytes,7,rep,name=spu_ranges,json=spuRanges,proto3" json:"spu_ranges"`
}

func (x *IntegralDeductRule) Reset() {
	*x = IntegralDeductRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integralService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IntegralDeductRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntegralDeductRule) ProtoMessage() {}

func (x *IntegralDeductRule) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use IntegralDeductRule.ProtoReflect.Descriptor instead.
func (*IntegralDeductRule) Descriptor() ([]byte, []int) {
	return file_integralService_proto_rawDescGZIP(), []int{1}
}

func (x *IntegralDeductRule) GetDeductAmountLimit() string {
	if x != nil {
		return x.DeductAmountLimit
	}
	return ""
}

func (x *IntegralDeductRule) GetDeductAmountLimitVal() int64 {
	if x != nil {
		return x.DeductAmountLimitVal
	}
	return 0
}

func (x *IntegralDeductRule) GetOrderAmountThreshold() string {
	if x != nil {
		return x.OrderAmountThreshold
	}
	return ""
}

func (x *IntegralDeductRule) GetMinOrderAmount() int64 {
	if x != nil {
		return x.MinOrderAmount
	}
	return 0
}

func (x *IntegralDeductRule) GetSpuRangeType() string {
	if x != nil {
		return x.SpuRangeType
	}
	return ""
}

func (x *IntegralDeductRule) GetSpuRangeIds() []int64 {
	if x != nil {
		return x.SpuRangeIds
	}
	return nil
}

func (x *IntegralDeductRule) GetSpuRanges() []*SpuInfo {
	if x != nil {
		return x.SpuRanges
	}
	return nil
}

// 积分信息
type IntegralValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MemberId         int64  `protobuf:"varint,1,opt,name=member_id,json=memberId,proto3" json:"member_id"`                          //用户ID
	Integral         int64  `protobuf:"varint,2,opt,name=integral,proto3" json:"integral"`                                          //积分
	Msg              string `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg"`                                                     //操作原因
	OperatorId       int64  `protobuf:"varint,4,opt,name=operator_id,json=operatorId,proto3" json:"operator_id"`                    //操作员ID
	OperatorName     string `protobuf:"bytes,5,opt,name=operator_name,json=operatorName,proto3" json:"operator_name"`               //操作员名称
	BusinessId       int64  `protobuf:"varint,6,opt,name=business_id,json=businessId,proto3" json:"business_id"`                    //业务ID
	BusinessType     string `protobuf:"bytes,7,opt,name=business_type,json=businessType,proto3" json:"business_type"`               //业务类型
	BusinessTypeName string `protobuf:"bytes,8,opt,name=business_type_name,json=businessTypeName,proto3" json:"business_type_name"` //业务类型名称
}

func (x *IntegralValue) Reset() {
	*x = IntegralValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integralService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IntegralValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntegralValue) ProtoMessage() {}

func (x *IntegralValue) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use IntegralValue.ProtoReflect.Descriptor instead.
func (*IntegralValue) Descriptor() ([]byte, []int) {
	return file_integralService_proto_rawDescGZIP(), []int{2}
}

func (x *IntegralValue) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

func (x *IntegralValue) GetIntegral() int64 {
	if x != nil {
		return x.Integral
	}
	return 0
}

func (x *IntegralValue) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *IntegralValue) GetOperatorId() int64 {
	if x != nil {
		return x.OperatorId
	}
	return 0
}

func (x *IntegralValue) GetOperatorName() string {
	if x != nil {
		return x.OperatorName
	}
	return ""
}

func (x *IntegralValue) GetBusinessId() int64 {
	if x != nil {
		return x.BusinessId
	}
	return 0
}

func (x *IntegralValue) GetBusinessType() string {
	if x != nil {
		return x.BusinessType
	}
	return ""
}

func (x *IntegralValue) GetBusinessTypeName() string {
	if x != nil {
		return x.BusinessTypeName
	}
	return ""
}

// 积分管理请求参数
type IntegralRequest struct {
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
	Name         string `protobuf:"bytes,11,opt,name=name,proto3" json:"name"`                                       //积分名称
	IsProtected  string `protobuf:"bytes,12,opt,name=is_protected,json=isProtected,proto3" json:"is_protected"`      //开启积分保护
	IsDeductRule string `protobuf:"bytes,13,opt,name=is_deduct_rule,json=isDeductRule,proto3" json:"is_deduct_rule"` //开启抵现规则
	CleanType    string `protobuf:"bytes,14,opt,name=clean_type,json=cleanType,proto3" json:"clean_type"`            //清零类型
	Status       string `protobuf:"bytes,15,opt,name=status,proto3" json:"status"`                                   //启用状态
}

func (x *IntegralRequest) Reset() {
	*x = IntegralRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integralService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IntegralRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntegralRequest) ProtoMessage() {}

func (x *IntegralRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use IntegralRequest.ProtoReflect.Descriptor instead.
func (*IntegralRequest) Descriptor() ([]byte, []int) {
	return file_integralService_proto_rawDescGZIP(), []int{3}
}

func (x *IntegralRequest) GetTop() int32 {
	if x != nil {
		return x.Top
	}
	return 0
}

func (x *IntegralRequest) GetPaged() int64 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *IntegralRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *IntegralRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *IntegralRequest) GetSort() []string {
	if x != nil {
		return x.Sort
	}
	return nil
}

func (x *IntegralRequest) GetDateRange() []string {
	if x != nil {
		return x.DateRange
	}
	return nil
}

func (x *IntegralRequest) GetIds() []int32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *IntegralRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *IntegralRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *IntegralRequest) GetIsProtected() string {
	if x != nil {
		return x.IsProtected
	}
	return ""
}

func (x *IntegralRequest) GetIsDeductRule() string {
	if x != nil {
		return x.IsDeductRule
	}
	return ""
}

func (x *IntegralRequest) GetCleanType() string {
	if x != nil {
		return x.CleanType
	}
	return ""
}

func (x *IntegralRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

// 积分管理响应数据
type IntegralResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg    string    `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg"`
	Entity *Integral `protobuf:"bytes,2,opt,name=entity,proto3" json:"entity"`
}

func (x *IntegralResponse) Reset() {
	*x = IntegralResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integralService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IntegralResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntegralResponse) ProtoMessage() {}

func (x *IntegralResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use IntegralResponse.ProtoReflect.Descriptor instead.
func (*IntegralResponse) Descriptor() ([]byte, []int) {
	return file_integralService_proto_rawDescGZIP(), []int{4}
}

func (x *IntegralResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *IntegralResponse) GetEntity() *Integral {
	if x != nil {
		return x.Entity
	}
	return nil
}

var File_integralService_proto protoreflect.FileDescriptor

var file_integralService_proto_rawDesc = []byte{
	0x0a, 0x15, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x1a, 0x15, 0x62, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xea, 0x05, 0x0a, 0x08, 0x49, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x73, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x69, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x25, 0x0a, 0x0e,
	0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x79, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x44,
	0x61, 0x79, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x64, 0x65, 0x64, 0x75,
	0x63, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x67, 0x69, 0x76, 0x69,
	0x6e, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x67,
	0x69, 0x76, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2b, 0x0a, 0x12, 0x69, 0x73, 0x5f,
	0x64, 0x61, 0x69, 0x6c, 0x79, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x69, 0x73, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x4d, 0x61,
	0x78, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x5f,
	0x6d, 0x61, 0x78, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0d, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x4d, 0x61, 0x78, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x24,
	0x0a, 0x0e, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x72, 0x75, 0x6c, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x73, 0x44, 0x65, 0x64, 0x75, 0x63, 0x74,
	0x52, 0x75, 0x6c, 0x65, 0x12, 0x3d, 0x0a, 0x0b, 0x64, 0x65, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x72,
	0x75, 0x6c, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x44, 0x65, 0x64,
	0x75, 0x63, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x0a, 0x64, 0x65, 0x64, 0x75, 0x63, 0x74, 0x52,
	0x75, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x65, 0x61, 0x6e, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x65, 0x61, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x65, 0x61, 0x6e, 0x5f, 0x79, 0x65, 0x61, 0x72,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x6c, 0x65, 0x61, 0x6e, 0x59, 0x65, 0x61,
	0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6c, 0x65, 0x61, 0x6e, 0x5f, 0x6d, 0x6f, 0x6e, 0x74, 0x68,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x6c, 0x65, 0x61, 0x6e, 0x4d, 0x6f, 0x6e,
	0x74, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x65, 0x61, 0x6e, 0x5f, 0x64, 0x61, 0x79, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x6c, 0x65, 0x61, 0x6e, 0x44, 0x61, 0x79, 0x12,
	0x30, 0x0a, 0x14, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x70, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x53, 0x70, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61,
	0x6c, 0x12, 0x32, 0x0a, 0x15, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x65, 0x61, 0x72, 0x6e, 0x65,
	0x64, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x18, 0x10, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x13, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x45, 0x61, 0x72, 0x6e, 0x65, 0x64, 0x49, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x6c, 0x12, 0x38, 0x0a, 0x18, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61,
	0x6c, 0x18, 0x11, 0x20, 0x01, 0x28, 0x03, 0x52, 0x16, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72,
	0x6f, 0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x13, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xd7, 0x02, 0x0a, 0x12, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72,
	0x61, 0x6c, 0x44, 0x65, 0x64, 0x75, 0x63, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x2e, 0x0a, 0x13,
	0x64, 0x65, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x64, 0x65, 0x64, 0x75, 0x63,
	0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x35, 0x0a, 0x17,
	0x64, 0x65, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x14, 0x64,
	0x65, 0x64, 0x75, 0x63, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x56, 0x61, 0x6c, 0x12, 0x34, 0x0a, 0x16, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x14, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x6d, 0x69, 0x6e,
	0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0e, 0x6d, 0x69, 0x6e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x73, 0x70, 0x75, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x70, 0x75,
	0x52, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x73, 0x70, 0x75,
	0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x03,
	0x52, 0x0b, 0x73, 0x70, 0x75, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x49, 0x64, 0x73, 0x12, 0x30, 0x0a,
	0x0a, 0x73, 0x70, 0x75, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x70, 0x75,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x73, 0x70, 0x75, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x22,
	0x94, 0x02, 0x0a, 0x0d, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73,
	0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x1f, 0x0a, 0x0b,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x23, 0x0a,
	0x0d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x69,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x62, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x54, 0x79,
	0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xdb, 0x02, 0x0a, 0x0f, 0x49, 0x6e, 0x74, 0x65, 0x67,
	0x72, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x6f,
	0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x74, 0x6f, 0x70, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x61, 0x67,
	0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x09, 0x64, 0x61, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x05, 0x52, 0x03, 0x69, 0x64, 0x73,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x65,
	0x63, 0x74, 0x65, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x73, 0x50, 0x72,
	0x6f, 0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x69, 0x73, 0x5f, 0x64, 0x65,
	0x64, 0x75, 0x63, 0x74, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x69, 0x73, 0x44, 0x65, 0x64, 0x75, 0x63, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x6c, 0x65, 0x61, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x6c, 0x65, 0x61, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x50, 0x0a, 0x10, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x2a, 0x0a, 0x06, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x52, 0x06,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x32, 0xea, 0x02, 0x0a, 0x0f, 0x49, 0x6e, 0x74, 0x65, 0x67,
	0x72, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x04, 0x53, 0x61,
	0x76, 0x65, 0x12, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e,
	0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x1a, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x12, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x1a,
	0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67,
	0x72, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a,
	0x10, 0x41, 0x64, 0x64, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x1a, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x49,
	0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x17, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x15, 0x52, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x49,
	0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x17, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_integralService_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_integralService_proto_goTypes = []interface{}{
	(*Integral)(nil),           // 0: services.Integral
	(*IntegralDeductRule)(nil), // 1: services.IntegralDeductRule
	(*IntegralValue)(nil),      // 2: services.IntegralValue
	(*IntegralRequest)(nil),    // 3: services.IntegralRequest
	(*IntegralResponse)(nil),   // 4: services.IntegralResponse
	(*SpuInfo)(nil),            // 5: services.SpuInfo
}
var file_integralService_proto_depIdxs = []int32{
	1, // 0: services.Integral.deduct_rule:type_name -> services.IntegralDeductRule
	5, // 1: services.IntegralDeductRule.spu_ranges:type_name -> services.SpuInfo
	0, // 2: services.IntegralResponse.entity:type_name -> services.Integral
	0, // 3: services.IntegralService.Save:input_type -> services.Integral
	0, // 4: services.IntegralService.Get:input_type -> services.Integral
	2, // 5: services.IntegralService.AddIntegralValue:input_type -> services.IntegralValue
	2, // 6: services.IntegralService.UseIntegralValue:input_type -> services.IntegralValue
	2, // 7: services.IntegralService.RollbackIntegralValue:input_type -> services.IntegralValue
	4, // 8: services.IntegralService.Save:output_type -> services.IntegralResponse
	4, // 9: services.IntegralService.Get:output_type -> services.IntegralResponse
	4, // 10: services.IntegralService.AddIntegralValue:output_type -> services.IntegralResponse
	4, // 11: services.IntegralService.UseIntegralValue:output_type -> services.IntegralResponse
	4, // 12: services.IntegralService.RollbackIntegralValue:output_type -> services.IntegralResponse
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_integralService_proto_init() }
func file_integralService_proto_init() {
	if File_integralService_proto != nil {
		return
	}
	file_baseInfoService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_integralService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Integral); i {
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
			switch v := v.(*IntegralDeductRule); i {
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
			switch v := v.(*IntegralValue); i {
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
		file_integralService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_integralService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
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
