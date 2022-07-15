// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: promotionService.proto

package services

import (
	common "github.com/geiqin/micro-kit/protobuf/common"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

//营销活动
type Promotion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                                                       //活动ID
	Type               string                 `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`                                                                    //活动类型:Coupon优惠劵[CPN]/满减送/限时折扣/秒杀/拼团/通用
	Code               string                 `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`                                                                    //活动编码:如：coupon_20210101212
	Title              string                 `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`                                                                  //活动标题
	BeganAt            string                 `protobuf:"bytes,5,opt,name=began_at,json=beganAt,proto3" json:"began_at,omitempty"`                                               //活动生效时间
	EndedAt            string                 `protobuf:"bytes,6,opt,name=ended_at,json=endedAt,proto3" json:"ended_at,omitempty"`                                               //活动结束时间
	CoverId            int64                  `protobuf:"varint,7,opt,name=cover_id,json=coverId,proto3" json:"cover_id,omitempty"`                                              //促销封面ID
	CoverUrl           string                 `protobuf:"bytes,8,opt,name=cover_url,json=coverUrl,proto3" json:"cover_url,omitempty"`                                            //促销封面url
	AssetUrl           string                 `protobuf:"bytes,9,opt,name=asset_url,json=assetUrl,proto3" json:"asset_url,omitempty"`                                            //促销详情链接
	Description        string                 `protobuf:"bytes,10,opt,name=description,proto3" json:"description,omitempty"`                                                     //活动描述
	Level              int32                  `protobuf:"varint,11,opt,name=level,proto3" json:"level,omitempty"`                                                                //促销级别：1 商品级，2订单级，3特殊级
	VariantRange       string                 `protobuf:"bytes,12,opt,name=variant_range,json=variantRange,proto3" json:"variant_range,omitempty"`                               //变体商品范围: 0-全部，1-部分参与，2-部分不参与
	ConfigData         string                 `protobuf:"bytes,13,opt,name=config_data,json=configData,proto3" json:"config_data,omitempty"`                                     //自定义配置数据json
	TotalCount         int32                  `protobuf:"varint,14,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`                                    //发放总量 【备选项】
	AssignCount        int32                  `protobuf:"varint,15,opt,name=assign_count,json=assignCount,proto3" json:"assign_count,omitempty"`                                 //已发放数量 【备选项】
	UsedCount          int32                  `protobuf:"varint,16,opt,name=used_count,json=usedCount,proto3" json:"used_count,omitempty"`                                       //已使用数量 【备选项】
	FailedCount        int32                  `protobuf:"varint,17,opt,name=failed_count,json=failedCount,proto3" json:"failed_count,omitempty"`                                 //失败，失效数量 【备选项】
	Status             string                 `protobuf:"bytes,18,opt,name=status,proto3" json:"status,omitempty"`                                                               //活动状态：0未开始 1进行中 2已结束 3已失效
	CreatedAt          string                 `protobuf:"bytes,19,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`                                        //创建时间
	UpdatedAt          string                 `protobuf:"bytes,20,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`                                        //修改时间
	VariantRangeSkuIds []int64                `protobuf:"varint,21,rep,packed,name=variant_range_sku_ids,json=variantRangeSkuIds,proto3" json:"variant_range_sku_ids,omitempty"` //变体商品IDS
	ActionConfig       *PromotionActionConfig `protobuf:"bytes,22,opt,name=action_config,json=actionConfig,proto3" json:"action_config,omitempty"`                               //限制规则配置
	RuleConfig         *PromotionRuleConfig   `protobuf:"bytes,23,opt,name=rule_config,json=ruleConfig,proto3" json:"rule_config,omitempty"`                                     //优惠策略配置
	Variants           *PromotionVariant      `protobuf:"bytes,24,opt,name=variants,proto3" json:"variants,omitempty"`                                                           //参与活动商品变体
}

func (x *Promotion) Reset() {
	*x = Promotion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_promotionService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Promotion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Promotion) ProtoMessage() {}

func (x *Promotion) ProtoReflect() protoreflect.Message {
	mi := &file_promotionService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Promotion.ProtoReflect.Descriptor instead.
func (*Promotion) Descriptor() ([]byte, []int) {
	return file_promotionService_proto_rawDescGZIP(), []int{0}
}

func (x *Promotion) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Promotion) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Promotion) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Promotion) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Promotion) GetBeganAt() string {
	if x != nil {
		return x.BeganAt
	}
	return ""
}

func (x *Promotion) GetEndedAt() string {
	if x != nil {
		return x.EndedAt
	}
	return ""
}

func (x *Promotion) GetCoverId() int64 {
	if x != nil {
		return x.CoverId
	}
	return 0
}

func (x *Promotion) GetCoverUrl() string {
	if x != nil {
		return x.CoverUrl
	}
	return ""
}

func (x *Promotion) GetAssetUrl() string {
	if x != nil {
		return x.AssetUrl
	}
	return ""
}

func (x *Promotion) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Promotion) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *Promotion) GetVariantRange() string {
	if x != nil {
		return x.VariantRange
	}
	return ""
}

func (x *Promotion) GetConfigData() string {
	if x != nil {
		return x.ConfigData
	}
	return ""
}

func (x *Promotion) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *Promotion) GetAssignCount() int32 {
	if x != nil {
		return x.AssignCount
	}
	return 0
}

func (x *Promotion) GetUsedCount() int32 {
	if x != nil {
		return x.UsedCount
	}
	return 0
}

func (x *Promotion) GetFailedCount() int32 {
	if x != nil {
		return x.FailedCount
	}
	return 0
}

func (x *Promotion) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Promotion) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Promotion) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Promotion) GetVariantRangeSkuIds() []int64 {
	if x != nil {
		return x.VariantRangeSkuIds
	}
	return nil
}

func (x *Promotion) GetActionConfig() *PromotionActionConfig {
	if x != nil {
		return x.ActionConfig
	}
	return nil
}

func (x *Promotion) GetRuleConfig() *PromotionRuleConfig {
	if x != nil {
		return x.RuleConfig
	}
	return nil
}

func (x *Promotion) GetVariants() *PromotionVariant {
	if x != nil {
		return x.Variants
	}
	return nil
}

type PromotionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int32  `protobuf:"varint,1,opt,name=paged,proto3" json:"paged,omitempty"`
	PageSize int32  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Keywords string `protobuf:"bytes,3,opt,name=keywords,proto3" json:"keywords,omitempty"`
	//以下为自定义参数
	Id        int32   `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
	Title     string  `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
	Status    string  `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	TypeId    int32   `protobuf:"varint,7,opt,name=type_id,json=typeId,proto3" json:"type_id,omitempty"`
	StartDate string  `protobuf:"bytes,8,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate   string  `protobuf:"bytes,9,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	Ids       []int32 `protobuf:"varint,10,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *PromotionRequest) Reset() {
	*x = PromotionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_promotionService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PromotionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PromotionRequest) ProtoMessage() {}

func (x *PromotionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_promotionService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PromotionRequest.ProtoReflect.Descriptor instead.
func (*PromotionRequest) Descriptor() ([]byte, []int) {
	return file_promotionService_proto_rawDescGZIP(), []int{1}
}

func (x *PromotionRequest) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *PromotionRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *PromotionRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *PromotionRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PromotionRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PromotionRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *PromotionRequest) GetTypeId() int32 {
	if x != nil {
		return x.TypeId
	}
	return 0
}

func (x *PromotionRequest) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *PromotionRequest) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

func (x *PromotionRequest) GetIds() []int32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type PromotionData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Promotion    `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items  []*Promotion  `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Info   *common.Info  `protobuf:"bytes,4,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *PromotionData) Reset() {
	*x = PromotionData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_promotionService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PromotionData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PromotionData) ProtoMessage() {}

func (x *PromotionData) ProtoReflect() protoreflect.Message {
	mi := &file_promotionService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PromotionData.ProtoReflect.Descriptor instead.
func (*PromotionData) Descriptor() ([]byte, []int) {
	return file_promotionService_proto_rawDescGZIP(), []int{2}
}

func (x *PromotionData) GetEntity() *Promotion {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *PromotionData) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *PromotionData) GetItems() []*Promotion {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *PromotionData) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type PromotionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *PromotionData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Error *common.Error  `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *PromotionResponse) Reset() {
	*x = PromotionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_promotionService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PromotionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PromotionResponse) ProtoMessage() {}

func (x *PromotionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_promotionService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PromotionResponse.ProtoReflect.Descriptor instead.
func (*PromotionResponse) Descriptor() ([]byte, []int) {
	return file_promotionService_proto_rawDescGZIP(), []int{3}
}

func (x *PromotionResponse) GetData() *PromotionData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *PromotionResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_promotionService_proto protoreflect.FileDescriptor

var file_promotionService_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x70, 0x72, 0x6f, 0x6d, 0x6f,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x75, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x70, 0x72, 0x6f,
	0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaf, 0x06, 0x0a, 0x09, 0x50,
	0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x65, 0x67, 0x61, 0x6e, 0x5f,
	0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x65, 0x67, 0x61, 0x6e, 0x41,
	0x74, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x41, 0x74, 0x12, 0x19, 0x0a, 0x08,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x55, 0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x73, 0x73, 0x65, 0x74, 0x55, 0x72,
	0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x76, 0x61, 0x72,
	0x69, 0x61, 0x6e, 0x74, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x21, 0x0a, 0x0c, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x75, 0x73, 0x65, 0x64, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x13, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x31, 0x0a, 0x15, 0x76,
	0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x73, 0x6b, 0x75,
	0x5f, 0x69, 0x64, 0x73, 0x18, 0x15, 0x20, 0x03, 0x28, 0x03, 0x52, 0x12, 0x76, 0x61, 0x72, 0x69,
	0x61, 0x6e, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x6b, 0x75, 0x49, 0x64, 0x73, 0x12, 0x44,
	0x0a, 0x0d, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18,
	0x16, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0c, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x3e, 0x0a, 0x0b, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x17, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x75,
	0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0a, 0x72, 0x75, 0x6c, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x36, 0x0a, 0x08, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x73,
	0x18, 0x18, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x72, 0x69, 0x61,
	0x6e, 0x74, 0x52, 0x08, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x73, 0x22, 0x84, 0x02, 0x0a,
	0x10, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x17,
	0x0a, 0x07, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x74, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x05, 0x52, 0x03,
	0x69, 0x64, 0x73, 0x22, 0xae, 0x01, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f,
	0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2b, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72,
	0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04,
	0x69, 0x6e, 0x66, 0x6f, 0x22, 0x65, 0x0a, 0x11, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x88, 0x04, 0x0a, 0x10,
	0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x3c, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x1a,
	0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x6f,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c,
	0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1b, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x06,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1b, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x05, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50,
	0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x04, 0x53, 0x74, 0x6f, 0x70, 0x12,
	0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x6f,
	0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x1a,
	0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x6f,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41,
	0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72,
	0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x43, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1a, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_promotionService_proto_rawDescOnce sync.Once
	file_promotionService_proto_rawDescData = file_promotionService_proto_rawDesc
)

func file_promotionService_proto_rawDescGZIP() []byte {
	file_promotionService_proto_rawDescOnce.Do(func() {
		file_promotionService_proto_rawDescData = protoimpl.X.CompressGZIP(file_promotionService_proto_rawDescData)
	})
	return file_promotionService_proto_rawDescData
}

var file_promotionService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_promotionService_proto_goTypes = []interface{}{
	(*Promotion)(nil),             // 0: services.Promotion
	(*PromotionRequest)(nil),      // 1: services.PromotionRequest
	(*PromotionData)(nil),         // 2: services.PromotionData
	(*PromotionResponse)(nil),     // 3: services.PromotionResponse
	(*PromotionActionConfig)(nil), // 4: services.PromotionActionConfig
	(*PromotionRuleConfig)(nil),   // 5: services.PromotionRuleConfig
	(*PromotionVariant)(nil),      // 6: services.PromotionVariant
	(*common.Pager)(nil),          // 7: common.Pager
	(*common.Info)(nil),           // 8: common.Info
	(*common.Error)(nil),          // 9: common.Error
}
var file_promotionService_proto_depIdxs = []int32{
	4,  // 0: services.Promotion.action_config:type_name -> services.PromotionActionConfig
	5,  // 1: services.Promotion.rule_config:type_name -> services.PromotionRuleConfig
	6,  // 2: services.Promotion.variants:type_name -> services.PromotionVariant
	0,  // 3: services.PromotionData.entity:type_name -> services.Promotion
	7,  // 4: services.PromotionData.pager:type_name -> common.Pager
	0,  // 5: services.PromotionData.items:type_name -> services.Promotion
	8,  // 6: services.PromotionData.info:type_name -> common.Info
	2,  // 7: services.PromotionResponse.data:type_name -> services.PromotionData
	9,  // 8: services.PromotionResponse.error:type_name -> common.Error
	0,  // 9: services.PromotionService.Create:input_type -> services.Promotion
	0,  // 10: services.PromotionService.Update:input_type -> services.Promotion
	0,  // 11: services.PromotionService.Delete:input_type -> services.Promotion
	0,  // 12: services.PromotionService.Start:input_type -> services.Promotion
	0,  // 13: services.PromotionService.Stop:input_type -> services.Promotion
	0,  // 14: services.PromotionService.Get:input_type -> services.Promotion
	1,  // 15: services.PromotionService.List:input_type -> services.PromotionRequest
	1,  // 16: services.PromotionService.Search:input_type -> services.PromotionRequest
	3,  // 17: services.PromotionService.Create:output_type -> services.PromotionResponse
	3,  // 18: services.PromotionService.Update:output_type -> services.PromotionResponse
	3,  // 19: services.PromotionService.Delete:output_type -> services.PromotionResponse
	3,  // 20: services.PromotionService.Start:output_type -> services.PromotionResponse
	3,  // 21: services.PromotionService.Stop:output_type -> services.PromotionResponse
	3,  // 22: services.PromotionService.Get:output_type -> services.PromotionResponse
	3,  // 23: services.PromotionService.List:output_type -> services.PromotionResponse
	3,  // 24: services.PromotionService.Search:output_type -> services.PromotionResponse
	17, // [17:25] is the sub-list for method output_type
	9,  // [9:17] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_promotionService_proto_init() }
func file_promotionService_proto_init() {
	if File_promotionService_proto != nil {
		return
	}
	file_promotionActionConfigService_proto_init()
	file_promotionRuleConfigService_proto_init()
	file_promotionVariantService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_promotionService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Promotion); i {
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
		file_promotionService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PromotionRequest); i {
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
		file_promotionService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PromotionData); i {
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
		file_promotionService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PromotionResponse); i {
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
			RawDescriptor: file_promotionService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_promotionService_proto_goTypes,
		DependencyIndexes: file_promotionService_proto_depIdxs,
		MessageInfos:      file_promotionService_proto_msgTypes,
	}.Build()
	File_promotionService_proto = out.File
	file_promotionService_proto_rawDesc = nil
	file_promotionService_proto_goTypes = nil
	file_promotionService_proto_depIdxs = nil
}
