// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: applicationVersionService.proto

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

// 应用版本
type ApplicationVersion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int32          `protobuf:"varint,1,opt,name=id,proto3" json:"id"`                                            //ID
	Code          string         `protobuf:"bytes,2,opt,name=code,proto3" json:"code"`                                         //版本编码
	Name          string         `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`                                         //版本名称
	ApplicationId int32          `protobuf:"varint,4,opt,name=application_id,json=applicationId,proto3" json:"application_id"` //所属应用
	MarketPrice   int64          `protobuf:"varint,5,opt,name=market_price,json=marketPrice,proto3" json:"market_price"`       //市场价
	Price         int64          `protobuf:"varint,6,opt,name=price,proto3" json:"price"`                                      //优惠价
	Description   string         `protobuf:"bytes,7,opt,name=description,proto3" json:"description"`                           //描述
	IsDefault     string         `protobuf:"bytes,8,opt,name=is_default,json=isDefault,proto3" json:"is_default"`              //是否默认
	Config        *VersionConfig `protobuf:"bytes,9,opt,name=config,proto3" json:"config"`                                     //使用配置
	Status        string         `protobuf:"bytes,10,opt,name=status,proto3" json:"status"`                                    //状态
	CreatedAt     int64          `protobuf:"varint,11,opt,name=created_at,json=createdAt,proto3" json:"created_at"`            //
	UpdatedAt     int64          `protobuf:"varint,12,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`            //
	Application   *Application   `protobuf:"bytes,13,opt,name=application,proto3" json:"application"`
}

func (x *ApplicationVersion) Reset() {
	*x = ApplicationVersion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_applicationVersionService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplicationVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplicationVersion) ProtoMessage() {}

func (x *ApplicationVersion) ProtoReflect() protoreflect.Message {
	mi := &file_applicationVersionService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplicationVersion.ProtoReflect.Descriptor instead.
func (*ApplicationVersion) Descriptor() ([]byte, []int) {
	return file_applicationVersionService_proto_rawDescGZIP(), []int{0}
}

func (x *ApplicationVersion) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ApplicationVersion) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *ApplicationVersion) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ApplicationVersion) GetApplicationId() int32 {
	if x != nil {
		return x.ApplicationId
	}
	return 0
}

func (x *ApplicationVersion) GetMarketPrice() int64 {
	if x != nil {
		return x.MarketPrice
	}
	return 0
}

func (x *ApplicationVersion) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *ApplicationVersion) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ApplicationVersion) GetIsDefault() string {
	if x != nil {
		return x.IsDefault
	}
	return ""
}

func (x *ApplicationVersion) GetConfig() *VersionConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *ApplicationVersion) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ApplicationVersion) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *ApplicationVersion) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *ApplicationVersion) GetApplication() *Application {
	if x != nil {
		return x.Application
	}
	return nil
}

// 版本配置
type VersionConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShopLimitCount      int32   `protobuf:"varint,1,opt,name=shop_limit_count,json=shopLimitCount,proto3" json:"shop_limit_count"`                   //商户限制数量
	RealstoreLimitCount int32   `protobuf:"varint,2,opt,name=realstore_limit_count,json=realstoreLimitCount,proto3" json:"realstore_limit_count"`    //门店限制数量
	WarehouseLimitCount int32   `protobuf:"varint,3,opt,name=warehouse_limit_count,json=warehouseLimitCount,proto3" json:"warehouse_limit_count"`    //仓库限制数量
	ManagerLimitCount   int32   `protobuf:"varint,4,opt,name=manager_limit_count,json=managerLimitCount,proto3" json:"manager_limit_count"`          //管理员限制数量
	MemberLimitCount    int32   `protobuf:"varint,5,opt,name=member_limit_count,json=memberLimitCount,proto3" json:"member_limit_count"`             //会员限制数量
	SpuLimitCount       int32   `protobuf:"varint,6,opt,name=spu_limit_count,json=spuLimitCount,proto3" json:"spu_limit_count"`                      //商品限制数量
	OrderFreeType       string  `protobuf:"bytes,7,opt,name=order_free_type,json=orderFreeType,proto3" json:"order_free_type"`                       //订单成交免费类型
	OrderFreeTotalLimit int32   `protobuf:"varint,8,opt,name=order_free_total_limit,json=orderFreeTotalLimit,proto3" json:"order_free_total_limit"`  //订单成交免费总限制
	OrderFreeYearLimit  int32   `protobuf:"varint,9,opt,name=order_free_year_limit,json=orderFreeYearLimit,proto3" json:"order_free_year_limit"`     //订单成交免费年度限制
	OrderFreeMonthLimit int32   `protobuf:"varint,10,opt,name=order_free_month_limit,json=orderFreeMonthLimit,proto3" json:"order_free_month_limit"` //订单成交免费月度限制
	OrderCommissionRate float32 `protobuf:"fixed32,11,opt,name=order_commission_rate,json=orderCommissionRate,proto3" json:"order_commission_rate"`  //订单成交抽成比例
	MediaLimitCount     int32   `protobuf:"varint,12,opt,name=media_limit_count,json=mediaLimitCount,proto3" json:"media_limit_count"`               //媒体限制数量
	MediaLimitVolume    int32   `protobuf:"varint,13,opt,name=media_limit_volume,json=mediaLimitVolume,proto3" json:"media_limit_volume"`            //媒体限制容量
	ArticleLimitCount   int32   `protobuf:"varint,14,opt,name=article_limit_count,json=articleLimitCount,proto3" json:"article_limit_count"`         //文章限制数量
}

func (x *VersionConfig) Reset() {
	*x = VersionConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_applicationVersionService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionConfig) ProtoMessage() {}

func (x *VersionConfig) ProtoReflect() protoreflect.Message {
	mi := &file_applicationVersionService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionConfig.ProtoReflect.Descriptor instead.
func (*VersionConfig) Descriptor() ([]byte, []int) {
	return file_applicationVersionService_proto_rawDescGZIP(), []int{1}
}

func (x *VersionConfig) GetShopLimitCount() int32 {
	if x != nil {
		return x.ShopLimitCount
	}
	return 0
}

func (x *VersionConfig) GetRealstoreLimitCount() int32 {
	if x != nil {
		return x.RealstoreLimitCount
	}
	return 0
}

func (x *VersionConfig) GetWarehouseLimitCount() int32 {
	if x != nil {
		return x.WarehouseLimitCount
	}
	return 0
}

func (x *VersionConfig) GetManagerLimitCount() int32 {
	if x != nil {
		return x.ManagerLimitCount
	}
	return 0
}

func (x *VersionConfig) GetMemberLimitCount() int32 {
	if x != nil {
		return x.MemberLimitCount
	}
	return 0
}

func (x *VersionConfig) GetSpuLimitCount() int32 {
	if x != nil {
		return x.SpuLimitCount
	}
	return 0
}

func (x *VersionConfig) GetOrderFreeType() string {
	if x != nil {
		return x.OrderFreeType
	}
	return ""
}

func (x *VersionConfig) GetOrderFreeTotalLimit() int32 {
	if x != nil {
		return x.OrderFreeTotalLimit
	}
	return 0
}

func (x *VersionConfig) GetOrderFreeYearLimit() int32 {
	if x != nil {
		return x.OrderFreeYearLimit
	}
	return 0
}

func (x *VersionConfig) GetOrderFreeMonthLimit() int32 {
	if x != nil {
		return x.OrderFreeMonthLimit
	}
	return 0
}

func (x *VersionConfig) GetOrderCommissionRate() float32 {
	if x != nil {
		return x.OrderCommissionRate
	}
	return 0
}

func (x *VersionConfig) GetMediaLimitCount() int32 {
	if x != nil {
		return x.MediaLimitCount
	}
	return 0
}

func (x *VersionConfig) GetMediaLimitVolume() int32 {
	if x != nil {
		return x.MediaLimitVolume
	}
	return 0
}

func (x *VersionConfig) GetArticleLimitCount() int32 {
	if x != nil {
		return x.ArticleLimitCount
	}
	return 0
}

// 应用版本请求参数
type ApplicationVersionRequest struct {
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
	Code          string `protobuf:"bytes,11,opt,name=code,proto3" json:"code"`                                         //版本编码
	Name          string `protobuf:"bytes,12,opt,name=name,proto3" json:"name"`                                         //版本名称
	ApplicationId int32  `protobuf:"varint,13,opt,name=application_id,json=applicationId,proto3" json:"application_id"` //所属应用
	IsDefault     string `protobuf:"bytes,14,opt,name=is_default,json=isDefault,proto3" json:"is_default"`              //是否默认
	Status        string `protobuf:"bytes,15,opt,name=status,proto3" json:"status"`                                     //状态
}

func (x *ApplicationVersionRequest) Reset() {
	*x = ApplicationVersionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_applicationVersionService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplicationVersionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplicationVersionRequest) ProtoMessage() {}

func (x *ApplicationVersionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_applicationVersionService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplicationVersionRequest.ProtoReflect.Descriptor instead.
func (*ApplicationVersionRequest) Descriptor() ([]byte, []int) {
	return file_applicationVersionService_proto_rawDescGZIP(), []int{2}
}

func (x *ApplicationVersionRequest) GetTop() int32 {
	if x != nil {
		return x.Top
	}
	return 0
}

func (x *ApplicationVersionRequest) GetPaged() int64 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *ApplicationVersionRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ApplicationVersionRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *ApplicationVersionRequest) GetSort() []string {
	if x != nil {
		return x.Sort
	}
	return nil
}

func (x *ApplicationVersionRequest) GetDateRange() []string {
	if x != nil {
		return x.DateRange
	}
	return nil
}

func (x *ApplicationVersionRequest) GetIds() []int32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *ApplicationVersionRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ApplicationVersionRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *ApplicationVersionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ApplicationVersionRequest) GetApplicationId() int32 {
	if x != nil {
		return x.ApplicationId
	}
	return 0
}

func (x *ApplicationVersionRequest) GetIsDefault() string {
	if x != nil {
		return x.IsDefault
	}
	return ""
}

func (x *ApplicationVersionRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

// 应用版本响应数据
type ApplicationVersionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *ApplicationVersion   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager         `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*ApplicationVersion `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Msg    string                `protobuf:"bytes,4,opt,name=msg,proto3" json:"msg"`
}

func (x *ApplicationVersionResponse) Reset() {
	*x = ApplicationVersionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_applicationVersionService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplicationVersionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplicationVersionResponse) ProtoMessage() {}

func (x *ApplicationVersionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_applicationVersionService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplicationVersionResponse.ProtoReflect.Descriptor instead.
func (*ApplicationVersionResponse) Descriptor() ([]byte, []int) {
	return file_applicationVersionService_proto_rawDescGZIP(), []int{3}
}

func (x *ApplicationVersionResponse) GetEntity() *ApplicationVersion {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *ApplicationVersionResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *ApplicationVersionResponse) GetItems() []*ApplicationVersion {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ApplicationVersionResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_applicationVersionService_proto protoreflect.FileDescriptor

var file_applicationVersionService_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xad, 0x03, 0x0a, 0x12, 0x41, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0d, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x21,
	0x0a, 0x0c, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f,
	0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69,
	0x73, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x2f, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x37, 0x0a, 0x0b, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xaa, 0x05, 0x0a, 0x0d, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x28, 0x0a, 0x10, 0x73, 0x68,
	0x6f, 0x70, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x73, 0x68, 0x6f, 0x70, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x32, 0x0a, 0x15, 0x72, 0x65, 0x61, 0x6c, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x13, 0x72, 0x65, 0x61, 0x6c, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x32, 0x0a, 0x15, 0x77, 0x61, 0x72, 0x65,
	0x68, 0x6f, 0x75, 0x73, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x13, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75,
	0x73, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x13,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x12,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x73, 0x70,
	0x75, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0d, 0x73, 0x70, 0x75, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x66, 0x72, 0x65, 0x65,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x46, 0x72, 0x65, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x33, 0x0a, 0x16, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x66, 0x72, 0x65, 0x65, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x13, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x46, 0x72, 0x65, 0x65, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x31, 0x0a, 0x15, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x66, 0x72, 0x65, 0x65, 0x5f, 0x79, 0x65,
	0x61, 0x72, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x46, 0x72, 0x65, 0x65, 0x59, 0x65, 0x61, 0x72, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x12, 0x33, 0x0a, 0x16, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x66, 0x72, 0x65, 0x65,
	0x5f, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x13, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x46, 0x72, 0x65, 0x65, 0x4d, 0x6f, 0x6e,
	0x74, 0x68, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x32, 0x0a, 0x15, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x61, 0x74, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x02, 0x52, 0x13, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6d,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x6d,
	0x65, 0x64, 0x69, 0x61, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x6d, 0x65, 0x64, 0x69, 0x61,
	0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x10, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x56,
	0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x11, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xd7, 0x02, 0x0a, 0x19, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x03, 0x74, 0x6f, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77,
	0x6f, 0x72, 0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77,
	0x6f, 0x72, 0x64, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x07,
	0x20, 0x03, 0x28, 0x05, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x64,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x73,
	0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0xbd, 0x01, 0x0a, 0x1a, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34,
	0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67,
	0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x32,
	0x84, 0x04, 0x0a, 0x19, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a,
	0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x24, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a,
	0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x24, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a,
	0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x24, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a,
	0x03, 0x47, 0x65, 0x74, 0x12, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x1a, 0x24, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x06, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x12, 0x23, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x53, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x23, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_applicationVersionService_proto_rawDescOnce sync.Once
	file_applicationVersionService_proto_rawDescData = file_applicationVersionService_proto_rawDesc
)

func file_applicationVersionService_proto_rawDescGZIP() []byte {
	file_applicationVersionService_proto_rawDescOnce.Do(func() {
		file_applicationVersionService_proto_rawDescData = protoimpl.X.CompressGZIP(file_applicationVersionService_proto_rawDescData)
	})
	return file_applicationVersionService_proto_rawDescData
}

var file_applicationVersionService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_applicationVersionService_proto_goTypes = []interface{}{
	(*ApplicationVersion)(nil),         // 0: services.ApplicationVersion
	(*VersionConfig)(nil),              // 1: services.VersionConfig
	(*ApplicationVersionRequest)(nil),  // 2: services.ApplicationVersionRequest
	(*ApplicationVersionResponse)(nil), // 3: services.ApplicationVersionResponse
	(*Application)(nil),                // 4: services.Application
	(*common.Pager)(nil),               // 5: common.Pager
}
var file_applicationVersionService_proto_depIdxs = []int32{
	1,  // 0: services.ApplicationVersion.config:type_name -> services.VersionConfig
	4,  // 1: services.ApplicationVersion.application:type_name -> services.Application
	0,  // 2: services.ApplicationVersionResponse.entity:type_name -> services.ApplicationVersion
	5,  // 3: services.ApplicationVersionResponse.pager:type_name -> common.Pager
	0,  // 4: services.ApplicationVersionResponse.items:type_name -> services.ApplicationVersion
	0,  // 5: services.ApplicationVersionService.Create:input_type -> services.ApplicationVersion
	0,  // 6: services.ApplicationVersionService.Update:input_type -> services.ApplicationVersion
	0,  // 7: services.ApplicationVersionService.Delete:input_type -> services.ApplicationVersion
	0,  // 8: services.ApplicationVersionService.Get:input_type -> services.ApplicationVersion
	2,  // 9: services.ApplicationVersionService.Search:input_type -> services.ApplicationVersionRequest
	2,  // 10: services.ApplicationVersionService.List:input_type -> services.ApplicationVersionRequest
	3,  // 11: services.ApplicationVersionService.Create:output_type -> services.ApplicationVersionResponse
	3,  // 12: services.ApplicationVersionService.Update:output_type -> services.ApplicationVersionResponse
	3,  // 13: services.ApplicationVersionService.Delete:output_type -> services.ApplicationVersionResponse
	3,  // 14: services.ApplicationVersionService.Get:output_type -> services.ApplicationVersionResponse
	3,  // 15: services.ApplicationVersionService.Search:output_type -> services.ApplicationVersionResponse
	3,  // 16: services.ApplicationVersionService.List:output_type -> services.ApplicationVersionResponse
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_applicationVersionService_proto_init() }
func file_applicationVersionService_proto_init() {
	if File_applicationVersionService_proto != nil {
		return
	}
	file_applicationService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_applicationVersionService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplicationVersion); i {
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
		file_applicationVersionService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VersionConfig); i {
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
		file_applicationVersionService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplicationVersionRequest); i {
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
		file_applicationVersionService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplicationVersionResponse); i {
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
			RawDescriptor: file_applicationVersionService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_applicationVersionService_proto_goTypes,
		DependencyIndexes: file_applicationVersionService_proto_depIdxs,
		MessageInfos:      file_applicationVersionService_proto_msgTypes,
	}.Build()
	File_applicationVersionService_proto = out.File
	file_applicationVersionService_proto_rawDesc = nil
	file_applicationVersionService_proto_goTypes = nil
	file_applicationVersionService_proto_depIdxs = nil
}
