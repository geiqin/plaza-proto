// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: storeConfigService.proto

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

// 店铺配置
type StoreConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                  int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id"`                                                                   //ID
	StoreId             int64   `protobuf:"varint,2,opt,name=store_id,json=storeId,proto3" json:"store_id"`                                          //店铺ID
	CustomDomain        string  `protobuf:"bytes,3,opt,name=custom_domain,json=customDomain,proto3" json:"custom_domain"`                            //自定义域名
	StoragePublic       string  `protobuf:"bytes,4,opt,name=storage_public,json=storagePublic,proto3" json:"storage_public"`                         //公共储存
	StoragePrivate      string  `protobuf:"bytes,5,opt,name=storage_private,json=storagePrivate,proto3" json:"storage_private"`                      //私有储存
	SmsChannel          string  `protobuf:"bytes,6,opt,name=sms_channel,json=smsChannel,proto3" json:"sms_channel"`                                  //短信通道
	EmailChannel        string  `protobuf:"bytes,7,opt,name=email_channel,json=emailChannel,proto3" json:"email_channel"`                            //邮件通道
	ShopLimitCount      int32   `protobuf:"varint,8,opt,name=shop_limit_count,json=shopLimitCount,proto3" json:"shop_limit_count"`                   //商户限制数量
	RealstoreLimitCount int32   `protobuf:"varint,9,opt,name=realstore_limit_count,json=realstoreLimitCount,proto3" json:"realstore_limit_count"`    //门店限制数量
	WarehouseLimitCount int32   `protobuf:"varint,10,opt,name=warehouse_limit_count,json=warehouseLimitCount,proto3" json:"warehouse_limit_count"`   //仓库限制数量
	ManagerLimitCount   int32   `protobuf:"varint,11,opt,name=manager_limit_count,json=managerLimitCount,proto3" json:"manager_limit_count"`         //管理员限制数量
	MemberLimitCount    int32   `protobuf:"varint,12,opt,name=member_limit_count,json=memberLimitCount,proto3" json:"member_limit_count"`            //会员限制数量
	SpuLimitCount       int32   `protobuf:"varint,13,opt,name=spu_limit_count,json=spuLimitCount,proto3" json:"spu_limit_count"`                     //商品限制数量
	OrderFreeType       string  `protobuf:"bytes,14,opt,name=order_free_type,json=orderFreeType,proto3" json:"order_free_type"`                      //订单成交免费类型
	OrderFreeTotalLimit int32   `protobuf:"varint,15,opt,name=order_free_total_limit,json=orderFreeTotalLimit,proto3" json:"order_free_total_limit"` //订单成交免费总限制
	OrderFreeYearLimit  int32   `protobuf:"varint,16,opt,name=order_free_year_limit,json=orderFreeYearLimit,proto3" json:"order_free_year_limit"`    //订单成交免费年度限制
	OrderFreeMonthLimit int32   `protobuf:"varint,17,opt,name=order_free_month_limit,json=orderFreeMonthLimit,proto3" json:"order_free_month_limit"` //订单成交免费月度限制
	OrderCommissionRate float32 `protobuf:"fixed32,18,opt,name=order_commission_rate,json=orderCommissionRate,proto3" json:"order_commission_rate"`  //订单成交抽成比例
	MediaLimitCount     int32   `protobuf:"varint,19,opt,name=media_limit_count,json=mediaLimitCount,proto3" json:"media_limit_count"`               //媒体限制数量
	MediaLimitVolume    int32   `protobuf:"varint,20,opt,name=media_limit_volume,json=mediaLimitVolume,proto3" json:"media_limit_volume"`            //媒体限制容量
	ArticleLimitCount   int32   `protobuf:"varint,21,opt,name=article_limit_count,json=articleLimitCount,proto3" json:"article_limit_count"`         //文章限制数量
	CreatedAt           int64   `protobuf:"varint,22,opt,name=created_at,json=createdAt,proto3" json:"created_at"`                                   //
	UpdatedAt           int64   `protobuf:"varint,23,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`                                   //
}

func (x *StoreConfig) Reset() {
	*x = StoreConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storeConfigService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreConfig) ProtoMessage() {}

func (x *StoreConfig) ProtoReflect() protoreflect.Message {
	mi := &file_storeConfigService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreConfig.ProtoReflect.Descriptor instead.
func (*StoreConfig) Descriptor() ([]byte, []int) {
	return file_storeConfigService_proto_rawDescGZIP(), []int{0}
}

func (x *StoreConfig) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *StoreConfig) GetStoreId() int64 {
	if x != nil {
		return x.StoreId
	}
	return 0
}

func (x *StoreConfig) GetCustomDomain() string {
	if x != nil {
		return x.CustomDomain
	}
	return ""
}

func (x *StoreConfig) GetStoragePublic() string {
	if x != nil {
		return x.StoragePublic
	}
	return ""
}

func (x *StoreConfig) GetStoragePrivate() string {
	if x != nil {
		return x.StoragePrivate
	}
	return ""
}

func (x *StoreConfig) GetSmsChannel() string {
	if x != nil {
		return x.SmsChannel
	}
	return ""
}

func (x *StoreConfig) GetEmailChannel() string {
	if x != nil {
		return x.EmailChannel
	}
	return ""
}

func (x *StoreConfig) GetShopLimitCount() int32 {
	if x != nil {
		return x.ShopLimitCount
	}
	return 0
}

func (x *StoreConfig) GetRealstoreLimitCount() int32 {
	if x != nil {
		return x.RealstoreLimitCount
	}
	return 0
}

func (x *StoreConfig) GetWarehouseLimitCount() int32 {
	if x != nil {
		return x.WarehouseLimitCount
	}
	return 0
}

func (x *StoreConfig) GetManagerLimitCount() int32 {
	if x != nil {
		return x.ManagerLimitCount
	}
	return 0
}

func (x *StoreConfig) GetMemberLimitCount() int32 {
	if x != nil {
		return x.MemberLimitCount
	}
	return 0
}

func (x *StoreConfig) GetSpuLimitCount() int32 {
	if x != nil {
		return x.SpuLimitCount
	}
	return 0
}

func (x *StoreConfig) GetOrderFreeType() string {
	if x != nil {
		return x.OrderFreeType
	}
	return ""
}

func (x *StoreConfig) GetOrderFreeTotalLimit() int32 {
	if x != nil {
		return x.OrderFreeTotalLimit
	}
	return 0
}

func (x *StoreConfig) GetOrderFreeYearLimit() int32 {
	if x != nil {
		return x.OrderFreeYearLimit
	}
	return 0
}

func (x *StoreConfig) GetOrderFreeMonthLimit() int32 {
	if x != nil {
		return x.OrderFreeMonthLimit
	}
	return 0
}

func (x *StoreConfig) GetOrderCommissionRate() float32 {
	if x != nil {
		return x.OrderCommissionRate
	}
	return 0
}

func (x *StoreConfig) GetMediaLimitCount() int32 {
	if x != nil {
		return x.MediaLimitCount
	}
	return 0
}

func (x *StoreConfig) GetMediaLimitVolume() int32 {
	if x != nil {
		return x.MediaLimitVolume
	}
	return 0
}

func (x *StoreConfig) GetArticleLimitCount() int32 {
	if x != nil {
		return x.ArticleLimitCount
	}
	return 0
}

func (x *StoreConfig) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *StoreConfig) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

// 店铺配置请求参数
type StoreConfigRequest struct {
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
	StoreId int64 `protobuf:"varint,11,opt,name=store_id,json=storeId,proto3" json:"store_id"` //店铺ID
}

func (x *StoreConfigRequest) Reset() {
	*x = StoreConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storeConfigService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreConfigRequest) ProtoMessage() {}

func (x *StoreConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_storeConfigService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreConfigRequest.ProtoReflect.Descriptor instead.
func (*StoreConfigRequest) Descriptor() ([]byte, []int) {
	return file_storeConfigService_proto_rawDescGZIP(), []int{1}
}

func (x *StoreConfigRequest) GetTop() int32 {
	if x != nil {
		return x.Top
	}
	return 0
}

func (x *StoreConfigRequest) GetPaged() int64 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *StoreConfigRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *StoreConfigRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *StoreConfigRequest) GetSort() []string {
	if x != nil {
		return x.Sort
	}
	return nil
}

func (x *StoreConfigRequest) GetDateRange() []string {
	if x != nil {
		return x.DateRange
	}
	return nil
}

func (x *StoreConfigRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *StoreConfigRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *StoreConfigRequest) GetStoreId() int64 {
	if x != nil {
		return x.StoreId
	}
	return 0
}

// 店铺配置响应数据
type StoreConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *StoreConfig   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager  `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*StoreConfig `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Msg    string         `protobuf:"bytes,4,opt,name=msg,proto3" json:"msg"`
}

func (x *StoreConfigResponse) Reset() {
	*x = StoreConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storeConfigService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreConfigResponse) ProtoMessage() {}

func (x *StoreConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_storeConfigService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreConfigResponse.ProtoReflect.Descriptor instead.
func (*StoreConfigResponse) Descriptor() ([]byte, []int) {
	return file_storeConfigService_proto_rawDescGZIP(), []int{2}
}

func (x *StoreConfigResponse) GetEntity() *StoreConfig {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *StoreConfigResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *StoreConfigResponse) GetItems() []*StoreConfig {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *StoreConfigResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_storeConfigService_proto protoreflect.FileDescriptor

var file_storeConfigService_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcc, 0x07, 0x0a, 0x0b, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x12, 0x27,
	0x0a, 0x0f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6d, 0x73, 0x5f, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6d,
	0x73, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x28, 0x0a,
	0x10, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x73, 0x68, 0x6f, 0x70, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x32, 0x0a, 0x15, 0x72, 0x65, 0x61, 0x6c, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x13, 0x72, 0x65, 0x61, 0x6c, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x32, 0x0a, 0x15, 0x77,
	0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x13, 0x77, 0x61, 0x72, 0x65,
	0x68, 0x6f, 0x75, 0x73, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x2e, 0x0a, 0x13, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x2c, 0x0a, 0x12, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x26, 0x0a,
	0x0f, 0x73, 0x70, 0x75, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x73, 0x70, 0x75, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x66,
	0x72, 0x65, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x46, 0x72, 0x65, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x33, 0x0a,
	0x16, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x66, 0x72, 0x65, 0x65, 0x5f, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x13, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x46, 0x72, 0x65, 0x65, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x12, 0x31, 0x0a, 0x15, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x66, 0x72, 0x65, 0x65,
	0x5f, 0x79, 0x65, 0x61, 0x72, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x12, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x46, 0x72, 0x65, 0x65, 0x59, 0x65, 0x61, 0x72,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x33, 0x0a, 0x16, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x66,
	0x72, 0x65, 0x65, 0x5f, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x11, 0x20, 0x01, 0x28, 0x05, 0x52, 0x13, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x46, 0x72, 0x65, 0x65,
	0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x32, 0x0a, 0x15, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x72,
	0x61, 0x74, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x02, 0x52, 0x13, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x65, 0x12, 0x2a,
	0x0a, 0x11, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x13, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x6d, 0x65, 0x64, 0x69, 0x61,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x6d, 0x65,
	0x64, 0x69, 0x61, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x61, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x15, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x16, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x17, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xe5, 0x01, 0x0a, 0x12, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x74, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x74, 0x6f, 0x70, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6f,
	0x72, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x64, 0x61, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x67,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03,
	0x69, 0x64, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x22, 0xa8,
	0x01, 0x0a, 0x13, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61,
	0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x32, 0xa9, 0x03, 0x0a, 0x12, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x40, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x1a, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x40, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x1a, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x15,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x15, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x1a, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12,
	0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45,
	0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_storeConfigService_proto_rawDescOnce sync.Once
	file_storeConfigService_proto_rawDescData = file_storeConfigService_proto_rawDesc
)

func file_storeConfigService_proto_rawDescGZIP() []byte {
	file_storeConfigService_proto_rawDescOnce.Do(func() {
		file_storeConfigService_proto_rawDescData = protoimpl.X.CompressGZIP(file_storeConfigService_proto_rawDescData)
	})
	return file_storeConfigService_proto_rawDescData
}

var file_storeConfigService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_storeConfigService_proto_goTypes = []interface{}{
	(*StoreConfig)(nil),         // 0: services.StoreConfig
	(*StoreConfigRequest)(nil),  // 1: services.StoreConfigRequest
	(*StoreConfigResponse)(nil), // 2: services.StoreConfigResponse
	(*common.Pager)(nil),        // 3: common.Pager
}
var file_storeConfigService_proto_depIdxs = []int32{
	0, // 0: services.StoreConfigResponse.entity:type_name -> services.StoreConfig
	3, // 1: services.StoreConfigResponse.pager:type_name -> common.Pager
	0, // 2: services.StoreConfigResponse.items:type_name -> services.StoreConfig
	0, // 3: services.StoreConfigService.Create:input_type -> services.StoreConfig
	0, // 4: services.StoreConfigService.Update:input_type -> services.StoreConfig
	0, // 5: services.StoreConfigService.Delete:input_type -> services.StoreConfig
	0, // 6: services.StoreConfigService.Get:input_type -> services.StoreConfig
	1, // 7: services.StoreConfigService.Search:input_type -> services.StoreConfigRequest
	1, // 8: services.StoreConfigService.List:input_type -> services.StoreConfigRequest
	2, // 9: services.StoreConfigService.Create:output_type -> services.StoreConfigResponse
	2, // 10: services.StoreConfigService.Update:output_type -> services.StoreConfigResponse
	2, // 11: services.StoreConfigService.Delete:output_type -> services.StoreConfigResponse
	2, // 12: services.StoreConfigService.Get:output_type -> services.StoreConfigResponse
	2, // 13: services.StoreConfigService.Search:output_type -> services.StoreConfigResponse
	2, // 14: services.StoreConfigService.List:output_type -> services.StoreConfigResponse
	9, // [9:15] is the sub-list for method output_type
	3, // [3:9] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_storeConfigService_proto_init() }
func file_storeConfigService_proto_init() {
	if File_storeConfigService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_storeConfigService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreConfig); i {
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
		file_storeConfigService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreConfigRequest); i {
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
		file_storeConfigService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreConfigResponse); i {
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
			RawDescriptor: file_storeConfigService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_storeConfigService_proto_goTypes,
		DependencyIndexes: file_storeConfigService_proto_depIdxs,
		MessageInfos:      file_storeConfigService_proto_msgTypes,
	}.Build()
	File_storeConfigService_proto = out.File
	file_storeConfigService_proto_rawDesc = nil
	file_storeConfigService_proto_goTypes = nil
	file_storeConfigService_proto_depIdxs = nil
}
