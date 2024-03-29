// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: promotionRuleConfigService.proto

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

//活动常用规则
type PromotionRuleConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderTotalAmount      int64           `protobuf:"varint,1,opt,name=order_total_amount,json=orderTotalAmount,proto3" json:"order_total_amount,omitempty"`                //订单总金额
	OrderTotalCount       int32           `protobuf:"varint,2,opt,name=order_total_count,json=orderTotalCount,proto3" json:"order_total_count,omitempty"`                   //订单总数量
	PromotionItemsTotal   int64           `protobuf:"varint,3,opt,name=promotion_items_total,json=promotionItemsTotal,proto3" json:"promotion_items_total,omitempty"`       //订单中促销项目总额
	WhichOrder            int32           `protobuf:"varint,4,opt,name=which_order,json=whichOrder,proto3" json:"which_order,omitempty"`                                    //第N笔订单
	MemberRankIds         []int32         `protobuf:"varint,5,rep,packed,name=member_rank_ids,json=memberRankIds,proto3" json:"member_rank_ids,omitempty"`                  //消费者会员等级组
	OverduePayMinute      int32           `protobuf:"varint,6,opt,name=overdue_pay_minute,json=overduePayMinute,proto3" json:"overdue_pay_minute,omitempty"`                //超时多少分钟未支付订单取消
	IsNewMember           bool            `protobuf:"varint,7,opt,name=is_new_member,json=isNewMember,proto3" json:"is_new_member,omitempty"`                               //是否为新客户
	IsOldMember           bool            `protobuf:"varint,8,opt,name=is_old_member,json=isOldMember,proto3" json:"is_old_member,omitempty"`                               //是否为老客户
	IsMutex               bool            `protobuf:"varint,9,opt,name=is_mutex,json=isMutex,proto3" json:"is_mutex,omitempty"`                                             //是否互斥，true 表示互斥，false 表示不互斥
	TimeRepeatType        string          `protobuf:"bytes,10,opt,name=time_repeat_type,json=timeRepeatType,proto3" json:"time_repeat_type,omitempty"`                      //周期重复类型：0-不重复，1-每天，2-每周，3-每月
	UseRange              int32           `protobuf:"varint,11,opt,name=use_range,json=useRange,proto3" json:"use_range,omitempty"`                                         //使用范围，0—全场，1—分店，2—类别，3—商品
	LimitCatchType        string          `protobuf:"bytes,13,opt,name=limit_catch_type,json=limitCatchType,proto3" json:"limit_catch_type,omitempty"`                      //限购类型：0不限购 1每人每种商品限购 2每人每种商品前XX件享受折扣
	LimitCatchNum         int32           `protobuf:"varint,14,opt,name=limit_catch_num,json=limitCatchNum,proto3" json:"limit_catch_num,omitempty"`                        //限购数量/每个用户可以领取的数量
	AllowPaymentTerminals []string        `protobuf:"bytes,15,rep,name=allow_payment_terminals,json=allowPaymentTerminals,proto3" json:"allow_payment_terminals,omitempty"` //允许的支付终端
	AllowPaymentCodes     []string        `protobuf:"bytes,16,rep,name=allow_payment_codes,json=allowPaymentCodes,proto3" json:"allow_payment_codes,omitempty"`             //允许的支付方式
	TimeRepeatData        *TimeRepeatInfo `protobuf:"bytes,17,opt,name=time_repeat_data,json=timeRepeatData,proto3" json:"time_repeat_data,omitempty"`                      //周期重复数据设置
}

func (x *PromotionRuleConfig) Reset() {
	*x = PromotionRuleConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_promotionRuleConfigService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PromotionRuleConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PromotionRuleConfig) ProtoMessage() {}

func (x *PromotionRuleConfig) ProtoReflect() protoreflect.Message {
	mi := &file_promotionRuleConfigService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PromotionRuleConfig.ProtoReflect.Descriptor instead.
func (*PromotionRuleConfig) Descriptor() ([]byte, []int) {
	return file_promotionRuleConfigService_proto_rawDescGZIP(), []int{0}
}

func (x *PromotionRuleConfig) GetOrderTotalAmount() int64 {
	if x != nil {
		return x.OrderTotalAmount
	}
	return 0
}

func (x *PromotionRuleConfig) GetOrderTotalCount() int32 {
	if x != nil {
		return x.OrderTotalCount
	}
	return 0
}

func (x *PromotionRuleConfig) GetPromotionItemsTotal() int64 {
	if x != nil {
		return x.PromotionItemsTotal
	}
	return 0
}

func (x *PromotionRuleConfig) GetWhichOrder() int32 {
	if x != nil {
		return x.WhichOrder
	}
	return 0
}

func (x *PromotionRuleConfig) GetMemberRankIds() []int32 {
	if x != nil {
		return x.MemberRankIds
	}
	return nil
}

func (x *PromotionRuleConfig) GetOverduePayMinute() int32 {
	if x != nil {
		return x.OverduePayMinute
	}
	return 0
}

func (x *PromotionRuleConfig) GetIsNewMember() bool {
	if x != nil {
		return x.IsNewMember
	}
	return false
}

func (x *PromotionRuleConfig) GetIsOldMember() bool {
	if x != nil {
		return x.IsOldMember
	}
	return false
}

func (x *PromotionRuleConfig) GetIsMutex() bool {
	if x != nil {
		return x.IsMutex
	}
	return false
}

func (x *PromotionRuleConfig) GetTimeRepeatType() string {
	if x != nil {
		return x.TimeRepeatType
	}
	return ""
}

func (x *PromotionRuleConfig) GetUseRange() int32 {
	if x != nil {
		return x.UseRange
	}
	return 0
}

func (x *PromotionRuleConfig) GetLimitCatchType() string {
	if x != nil {
		return x.LimitCatchType
	}
	return ""
}

func (x *PromotionRuleConfig) GetLimitCatchNum() int32 {
	if x != nil {
		return x.LimitCatchNum
	}
	return 0
}

func (x *PromotionRuleConfig) GetAllowPaymentTerminals() []string {
	if x != nil {
		return x.AllowPaymentTerminals
	}
	return nil
}

func (x *PromotionRuleConfig) GetAllowPaymentCodes() []string {
	if x != nil {
		return x.AllowPaymentCodes
	}
	return nil
}

func (x *PromotionRuleConfig) GetTimeRepeatData() *TimeRepeatInfo {
	if x != nil {
		return x.TimeRepeatData
	}
	return nil
}

//按周期重复数据
type TimeRepeatInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TimeBegan  string  `protobuf:"bytes,1,opt,name=time_began,json=timeBegan,proto3" json:"time_began,omitempty"`            //每天开始时间：hh:mm
	TimeEnded  string  `protobuf:"bytes,2,opt,name=time_ended,json=timeEnded,proto3" json:"time_ended,omitempty"`            //每天结束时间：hh:mm
	MonthlyDay int32   `protobuf:"varint,3,opt,name=monthly_day,json=monthlyDay,proto3" json:"monthly_day,omitempty"`        //每月的几号
	WeeklyDays []int32 `protobuf:"varint,4,rep,packed,name=weekly_days,json=weeklyDays,proto3" json:"weekly_days,omitempty"` //每周的那几天
}

func (x *TimeRepeatInfo) Reset() {
	*x = TimeRepeatInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_promotionRuleConfigService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeRepeatInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeRepeatInfo) ProtoMessage() {}

func (x *TimeRepeatInfo) ProtoReflect() protoreflect.Message {
	mi := &file_promotionRuleConfigService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeRepeatInfo.ProtoReflect.Descriptor instead.
func (*TimeRepeatInfo) Descriptor() ([]byte, []int) {
	return file_promotionRuleConfigService_proto_rawDescGZIP(), []int{1}
}

func (x *TimeRepeatInfo) GetTimeBegan() string {
	if x != nil {
		return x.TimeBegan
	}
	return ""
}

func (x *TimeRepeatInfo) GetTimeEnded() string {
	if x != nil {
		return x.TimeEnded
	}
	return ""
}

func (x *TimeRepeatInfo) GetMonthlyDay() int32 {
	if x != nil {
		return x.MonthlyDay
	}
	return 0
}

func (x *TimeRepeatInfo) GetWeeklyDays() []int32 {
	if x != nil {
		return x.WeeklyDays
	}
	return nil
}

var File_promotionRuleConfigService_proto protoreflect.FileDescriptor

var file_promotionRuleConfigService_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x75, 0x6c, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0xc2, 0x05, 0x0a,
	0x13, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x75, 0x6c, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x2c, 0x0a, 0x12, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x10, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x32,
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x13, 0x70,
	0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x54, 0x6f, 0x74,
	0x61, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x68, 0x69, 0x63, 0x68, 0x5f, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x77, 0x68, 0x69, 0x63, 0x68, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x72, 0x61,
	0x6e, 0x6b, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0d, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x52, 0x61, 0x6e, 0x6b, 0x49, 0x64, 0x73, 0x12, 0x2c, 0x0a, 0x12, 0x6f,
	0x76, 0x65, 0x72, 0x64, 0x75, 0x65, 0x5f, 0x70, 0x61, 0x79, 0x5f, 0x6d, 0x69, 0x6e, 0x75, 0x74,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x6f, 0x76, 0x65, 0x72, 0x64, 0x75, 0x65,
	0x50, 0x61, 0x79, 0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x69, 0x73, 0x5f,
	0x6e, 0x65, 0x77, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0b, 0x69, 0x73, 0x4e, 0x65, 0x77, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x22, 0x0a,
	0x0d, 0x69, 0x73, 0x5f, 0x6f, 0x6c, 0x64, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x4f, 0x6c, 0x64, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x6d, 0x75, 0x74, 0x65, 0x78, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x4d, 0x75, 0x74, 0x65, 0x78, 0x12, 0x28, 0x0a, 0x10,
	0x74, 0x69, 0x6d, 0x65, 0x5f, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x70, 0x65,
	0x61, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x5f, 0x72, 0x61,
	0x6e, 0x67, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x75, 0x73, 0x65, 0x52, 0x61,
	0x6e, 0x67, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x63, 0x61, 0x74,
	0x63, 0x68, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x43, 0x61, 0x74, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x26, 0x0a,
	0x0f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x63, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x6e, 0x75, 0x6d,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x43, 0x61, 0x74,
	0x63, 0x68, 0x4e, 0x75, 0x6d, 0x12, 0x36, 0x0a, 0x17, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x73,
	0x18, 0x0f, 0x20, 0x03, 0x28, 0x09, 0x52, 0x15, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x73, 0x12, 0x2e, 0x0a,
	0x13, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x73, 0x18, 0x10, 0x20, 0x03, 0x28, 0x09, 0x52, 0x11, 0x61, 0x6c, 0x6c, 0x6f,
	0x77, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x42, 0x0a,
	0x10, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x0e, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x22, 0x90, 0x01, 0x0a, 0x0e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x62, 0x65, 0x67,
	0x61, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x65,
	0x67, 0x61, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x65, 0x6e, 0x64, 0x65,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x45, 0x6e, 0x64,
	0x65, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x5f, 0x64, 0x61,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79,
	0x44, 0x61, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x5f, 0x64, 0x61,
	0x79, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0a, 0x77, 0x65, 0x65, 0x6b, 0x6c, 0x79,
	0x44, 0x61, 0x79, 0x73, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_promotionRuleConfigService_proto_rawDescOnce sync.Once
	file_promotionRuleConfigService_proto_rawDescData = file_promotionRuleConfigService_proto_rawDesc
)

func file_promotionRuleConfigService_proto_rawDescGZIP() []byte {
	file_promotionRuleConfigService_proto_rawDescOnce.Do(func() {
		file_promotionRuleConfigService_proto_rawDescData = protoimpl.X.CompressGZIP(file_promotionRuleConfigService_proto_rawDescData)
	})
	return file_promotionRuleConfigService_proto_rawDescData
}

var file_promotionRuleConfigService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_promotionRuleConfigService_proto_goTypes = []interface{}{
	(*PromotionRuleConfig)(nil), // 0: services.PromotionRuleConfig
	(*TimeRepeatInfo)(nil),      // 1: services.TimeRepeatInfo
}
var file_promotionRuleConfigService_proto_depIdxs = []int32{
	1, // 0: services.PromotionRuleConfig.time_repeat_data:type_name -> services.TimeRepeatInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_promotionRuleConfigService_proto_init() }
func file_promotionRuleConfigService_proto_init() {
	if File_promotionRuleConfigService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_promotionRuleConfigService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PromotionRuleConfig); i {
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
		file_promotionRuleConfigService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeRepeatInfo); i {
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
			RawDescriptor: file_promotionRuleConfigService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_promotionRuleConfigService_proto_goTypes,
		DependencyIndexes: file_promotionRuleConfigService_proto_depIdxs,
		MessageInfos:      file_promotionRuleConfigService_proto_msgTypes,
	}.Build()
	File_promotionRuleConfigService_proto = out.File
	file_promotionRuleConfigService_proto_rawDesc = nil
	file_promotionRuleConfigService_proto_goTypes = nil
	file_promotionRuleConfigService_proto_depIdxs = nil
}
