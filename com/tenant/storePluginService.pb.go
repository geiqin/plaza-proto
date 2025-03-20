// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: storePluginService.proto

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

type StorePlugin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id"`                                       //ID
	StoreId       int64   `protobuf:"varint,2,opt,name=store_id,json=storeId,proto3" json:"store_id"`              //店铺ID
	PluginId      int32   `protobuf:"varint,3,opt,name=plugin_id,json=pluginId,proto3" json:"plugin_id"`           //插件ID
	PluginCode    string  `protobuf:"bytes,4,opt,name=plugin_code,json=pluginCode,proto3" json:"plugin_code"`      //插件code
	IsEnabled     string  `protobuf:"bytes,5,opt,name=is_enabled,json=isEnabled,proto3" json:"is_enabled"`         //是否已启用
	IsError       string  `protobuf:"bytes,6,opt,name=is_error,json=isError,proto3" json:"is_error"`               //是否有异常
	ErrorMsg      string  `protobuf:"bytes,7,opt,name=error_msg,json=errorMsg,proto3" json:"error_msg"`            //错误信息
	Config        string  `protobuf:"bytes,8,opt,name=config,proto3" json:"config"`                                //配置信息
	Sort          int32   `protobuf:"varint,9,opt,name=sort,proto3" json:"sort"`                                   //排序
	ExpiredTime   int64   `protobuf:"varint,10,opt,name=expired_time,json=expiredTime,proto3" json:"expired_time"` //到期时间
	Ver           int32   `protobuf:"varint,11,opt,name=ver,proto3" json:"ver"`                                    //版本号
	CreatedAt     int64   `protobuf:"varint,12,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt     int64   `protobuf:"varint,13,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	Plugin        *Plugin `protobuf:"bytes,14,opt,name=plugin,proto3" json:"plugin"`
	RemainingDays int64   `protobuf:"varint,15,opt,name=remaining_days,json=remainingDays,proto3" json:"remaining_days"` //剩余天数
	VerText       string  `protobuf:"bytes,16,opt,name=ver_text,json=verText,proto3" json:"ver_text"`                    //版本文本（版本号美化）
}

func (x *StorePlugin) Reset() {
	*x = StorePlugin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storePluginService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorePlugin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorePlugin) ProtoMessage() {}

func (x *StorePlugin) ProtoReflect() protoreflect.Message {
	mi := &file_storePluginService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorePlugin.ProtoReflect.Descriptor instead.
func (*StorePlugin) Descriptor() ([]byte, []int) {
	return file_storePluginService_proto_rawDescGZIP(), []int{0}
}

func (x *StorePlugin) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *StorePlugin) GetStoreId() int64 {
	if x != nil {
		return x.StoreId
	}
	return 0
}

func (x *StorePlugin) GetPluginId() int32 {
	if x != nil {
		return x.PluginId
	}
	return 0
}

func (x *StorePlugin) GetPluginCode() string {
	if x != nil {
		return x.PluginCode
	}
	return ""
}

func (x *StorePlugin) GetIsEnabled() string {
	if x != nil {
		return x.IsEnabled
	}
	return ""
}

func (x *StorePlugin) GetIsError() string {
	if x != nil {
		return x.IsError
	}
	return ""
}

func (x *StorePlugin) GetErrorMsg() string {
	if x != nil {
		return x.ErrorMsg
	}
	return ""
}

func (x *StorePlugin) GetConfig() string {
	if x != nil {
		return x.Config
	}
	return ""
}

func (x *StorePlugin) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *StorePlugin) GetExpiredTime() int64 {
	if x != nil {
		return x.ExpiredTime
	}
	return 0
}

func (x *StorePlugin) GetVer() int32 {
	if x != nil {
		return x.Ver
	}
	return 0
}

func (x *StorePlugin) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *StorePlugin) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *StorePlugin) GetPlugin() *Plugin {
	if x != nil {
		return x.Plugin
	}
	return nil
}

func (x *StorePlugin) GetRemainingDays() int64 {
	if x != nil {
		return x.RemainingDays
	}
	return 0
}

func (x *StorePlugin) GetVerText() string {
	if x != nil {
		return x.VerText
	}
	return ""
}

type PromotionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code   string `protobuf:"bytes,1,opt,name=code,proto3" json:"code"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	Status string `protobuf:"bytes,3,opt,name=status,proto3" json:"status"`
}

func (x *PromotionInfo) Reset() {
	*x = PromotionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storePluginService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PromotionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PromotionInfo) ProtoMessage() {}

func (x *PromotionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_storePluginService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PromotionInfo.ProtoReflect.Descriptor instead.
func (*PromotionInfo) Descriptor() ([]byte, []int) {
	return file_storePluginService_proto_rawDescGZIP(), []int{1}
}

func (x *PromotionInfo) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *PromotionInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PromotionInfo) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type StorePluginRequest struct {
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
	//base params
	StoreId     int64    `protobuf:"varint,10,opt,name=store_id,json=storeId,proto3" json:"store_id"`
	PluginId    int32    `protobuf:"varint,11,opt,name=plugin_id,json=pluginId,proto3" json:"plugin_id"`
	PluginCode  string   `protobuf:"bytes,12,opt,name=plugin_code,json=pluginCode,proto3" json:"plugin_code"`
	IsPromotion string   `protobuf:"bytes,13,opt,name=is_promotion,json=isPromotion,proto3" json:"is_promotion"`
	IsEnabled   string   `protobuf:"bytes,14,opt,name=is_enabled,json=isEnabled,proto3" json:"is_enabled"`
	IsError     string   `protobuf:"bytes,15,opt,name=is_error,json=isError,proto3" json:"is_error"`
	Mode        string   `protobuf:"bytes,16,opt,name=mode,proto3" json:"mode"` //ConfigList 数据模式: 1-plugin_base模式,2-plugin_data 模式，3-maps模式，无-对象列表
	StatusList  []string `protobuf:"bytes,17,rep,name=status_list,json=statusList,proto3" json:"status_list"`
	PluginIds   []int32  `protobuf:"varint,18,rep,packed,name=plugin_ids,json=pluginIds,proto3" json:"plugin_ids"`
	PluginCodes []string `protobuf:"bytes,19,rep,name=plugin_codes,json=pluginCodes,proto3" json:"plugin_codes"`
}

func (x *StorePluginRequest) Reset() {
	*x = StorePluginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storePluginService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorePluginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorePluginRequest) ProtoMessage() {}

func (x *StorePluginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_storePluginService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorePluginRequest.ProtoReflect.Descriptor instead.
func (*StorePluginRequest) Descriptor() ([]byte, []int) {
	return file_storePluginService_proto_rawDescGZIP(), []int{2}
}

func (x *StorePluginRequest) GetTop() int32 {
	if x != nil {
		return x.Top
	}
	return 0
}

func (x *StorePluginRequest) GetPaged() int64 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *StorePluginRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *StorePluginRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *StorePluginRequest) GetSort() []string {
	if x != nil {
		return x.Sort
	}
	return nil
}

func (x *StorePluginRequest) GetDateRange() []string {
	if x != nil {
		return x.DateRange
	}
	return nil
}

func (x *StorePluginRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *StorePluginRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *StorePluginRequest) GetStoreId() int64 {
	if x != nil {
		return x.StoreId
	}
	return 0
}

func (x *StorePluginRequest) GetPluginId() int32 {
	if x != nil {
		return x.PluginId
	}
	return 0
}

func (x *StorePluginRequest) GetPluginCode() string {
	if x != nil {
		return x.PluginCode
	}
	return ""
}

func (x *StorePluginRequest) GetIsPromotion() string {
	if x != nil {
		return x.IsPromotion
	}
	return ""
}

func (x *StorePluginRequest) GetIsEnabled() string {
	if x != nil {
		return x.IsEnabled
	}
	return ""
}

func (x *StorePluginRequest) GetIsError() string {
	if x != nil {
		return x.IsError
	}
	return ""
}

func (x *StorePluginRequest) GetMode() string {
	if x != nil {
		return x.Mode
	}
	return ""
}

func (x *StorePluginRequest) GetStatusList() []string {
	if x != nil {
		return x.StatusList
	}
	return nil
}

func (x *StorePluginRequest) GetPluginIds() []int32 {
	if x != nil {
		return x.PluginIds
	}
	return nil
}

func (x *StorePluginRequest) GetPluginCodes() []string {
	if x != nil {
		return x.PluginCodes
	}
	return nil
}

type StorePluginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg               string           `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg"`
	Entity            *StorePlugin     `protobuf:"bytes,2,opt,name=entity,proto3" json:"entity"`
	Pager             *common.Pager    `protobuf:"bytes,3,opt,name=pager,proto3" json:"pager"`
	Items             []*StorePlugin   `protobuf:"bytes,4,rep,name=items,proto3" json:"items"`
	PluginsValidCodes []string         `protobuf:"bytes,5,rep,name=plugins_valid_codes,json=pluginsValidCodes,proto3" json:"plugins_valid_codes"`
	PromotionList     []*PromotionInfo `protobuf:"bytes,6,rep,name=promotion_list,json=promotionList,proto3" json:"promotion_list"`
}

func (x *StorePluginResponse) Reset() {
	*x = StorePluginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storePluginService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorePluginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorePluginResponse) ProtoMessage() {}

func (x *StorePluginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_storePluginService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorePluginResponse.ProtoReflect.Descriptor instead.
func (*StorePluginResponse) Descriptor() ([]byte, []int) {
	return file_storePluginService_proto_rawDescGZIP(), []int{3}
}

func (x *StorePluginResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *StorePluginResponse) GetEntity() *StorePlugin {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *StorePluginResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *StorePluginResponse) GetItems() []*StorePlugin {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *StorePluginResponse) GetPluginsValidCodes() []string {
	if x != nil {
		return x.PluginsValidCodes
	}
	return nil
}

func (x *StorePluginResponse) GetPromotionList() []*PromotionInfo {
	if x != nil {
		return x.PromotionList
	}
	return nil
}

var File_storePluginService_proto protoreflect.FileDescriptor

var file_storePluginService_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd8, 0x03, 0x0a,
	0x0b, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x73, 0x45, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x73, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12,
	0x1b, 0x0a, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73, 0x67, 0x12, 0x16, 0x0a, 0x06,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x76,
	0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x76, 0x65, 0x72, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x28, 0x0a, 0x06, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x06, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69,
	0x6e, 0x67, 0x5f, 0x64, 0x61, 0x79, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x72,
	0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x79, 0x73, 0x12, 0x19, 0x0a, 0x08,
	0x76, 0x65, 0x72, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x54, 0x65, 0x78, 0x74, 0x22, 0x4f, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x6d, 0x6f,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xf7, 0x03, 0x0a, 0x12, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x74, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x74, 0x6f,
	0x70, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04,
	0x73, 0x6f, 0x72, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x61, 0x6e,
	0x67, 0x65, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x64, 0x61, 0x74, 0x65, 0x52, 0x61,
	0x6e, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x03,
	0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x73, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x69, 0x73, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6d,
	0x6f, 0x64, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x11,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x12,
	0x20, 0x03, 0x28, 0x05, 0x52, 0x09, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x64, 0x73, 0x12,
	0x21, 0x0a, 0x0c, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x18,
	0x13, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x43, 0x6f, 0x64,
	0x65, 0x73, 0x22, 0x98, 0x02, 0x0a, 0x13, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x2d, 0x0a, 0x06,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70,
	0x61, 0x67, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72,
	0x12, 0x2b, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x2e, 0x0a,
	0x13, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x11, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x3e, 0x0a,
	0x0e, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0d,
	0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x32, 0xdc, 0x06,
	0x0a, 0x12, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x07, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x12,
	0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x1a, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x1a, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x07, 0x55, 0x70, 0x67,
	0x72, 0x61, 0x64, 0x65, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x1a, 0x1d, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x06,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x1a, 0x1d, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40,
	0x0a, 0x06, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x1a,
	0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x40, 0x0a, 0x06, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x1a, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x3e, 0x0a, 0x04, 0x42, 0x61, 0x73, 0x65, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x1a, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x45, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1c, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x06, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x12, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x44, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x61, 0x76, 0x65,
	0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x1a, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x11, 0x50, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x1c, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x0d,
	0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1c, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b,
	0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_storePluginService_proto_rawDescOnce sync.Once
	file_storePluginService_proto_rawDescData = file_storePluginService_proto_rawDesc
)

func file_storePluginService_proto_rawDescGZIP() []byte {
	file_storePluginService_proto_rawDescOnce.Do(func() {
		file_storePluginService_proto_rawDescData = protoimpl.X.CompressGZIP(file_storePluginService_proto_rawDescData)
	})
	return file_storePluginService_proto_rawDescData
}

var file_storePluginService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_storePluginService_proto_goTypes = []interface{}{
	(*StorePlugin)(nil),         // 0: services.StorePlugin
	(*PromotionInfo)(nil),       // 1: services.PromotionInfo
	(*StorePluginRequest)(nil),  // 2: services.StorePluginRequest
	(*StorePluginResponse)(nil), // 3: services.StorePluginResponse
	(*Plugin)(nil),              // 4: services.Plugin
	(*common.Pager)(nil),        // 5: common.Pager
}
var file_storePluginService_proto_depIdxs = []int32{
	4,  // 0: services.StorePlugin.plugin:type_name -> services.Plugin
	0,  // 1: services.StorePluginResponse.entity:type_name -> services.StorePlugin
	5,  // 2: services.StorePluginResponse.pager:type_name -> common.Pager
	0,  // 3: services.StorePluginResponse.items:type_name -> services.StorePlugin
	1,  // 4: services.StorePluginResponse.promotion_list:type_name -> services.PromotionInfo
	0,  // 5: services.StorePluginService.Install:input_type -> services.StorePlugin
	0,  // 6: services.StorePluginService.Update:input_type -> services.StorePlugin
	0,  // 7: services.StorePluginService.Upgrade:input_type -> services.StorePlugin
	0,  // 8: services.StorePluginService.Remove:input_type -> services.StorePlugin
	0,  // 9: services.StorePluginService.Switch:input_type -> services.StorePlugin
	0,  // 10: services.StorePluginService.Detail:input_type -> services.StorePlugin
	0,  // 11: services.StorePluginService.Base:input_type -> services.StorePlugin
	2,  // 12: services.StorePluginService.List:input_type -> services.StorePluginRequest
	2,  // 13: services.StorePluginService.Search:input_type -> services.StorePluginRequest
	0,  // 14: services.StorePluginService.ConfigSave:input_type -> services.StorePlugin
	2,  // 15: services.StorePluginService.PluginsValidCodes:input_type -> services.StorePluginRequest
	2,  // 16: services.StorePluginService.PromotionList:input_type -> services.StorePluginRequest
	3,  // 17: services.StorePluginService.Install:output_type -> services.StorePluginResponse
	3,  // 18: services.StorePluginService.Update:output_type -> services.StorePluginResponse
	3,  // 19: services.StorePluginService.Upgrade:output_type -> services.StorePluginResponse
	3,  // 20: services.StorePluginService.Remove:output_type -> services.StorePluginResponse
	3,  // 21: services.StorePluginService.Switch:output_type -> services.StorePluginResponse
	3,  // 22: services.StorePluginService.Detail:output_type -> services.StorePluginResponse
	3,  // 23: services.StorePluginService.Base:output_type -> services.StorePluginResponse
	3,  // 24: services.StorePluginService.List:output_type -> services.StorePluginResponse
	3,  // 25: services.StorePluginService.Search:output_type -> services.StorePluginResponse
	3,  // 26: services.StorePluginService.ConfigSave:output_type -> services.StorePluginResponse
	3,  // 27: services.StorePluginService.PluginsValidCodes:output_type -> services.StorePluginResponse
	3,  // 28: services.StorePluginService.PromotionList:output_type -> services.StorePluginResponse
	17, // [17:29] is the sub-list for method output_type
	5,  // [5:17] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_storePluginService_proto_init() }
func file_storePluginService_proto_init() {
	if File_storePluginService_proto != nil {
		return
	}
	file_pluginService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_storePluginService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorePlugin); i {
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
		file_storePluginService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PromotionInfo); i {
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
		file_storePluginService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorePluginRequest); i {
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
		file_storePluginService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorePluginResponse); i {
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
			RawDescriptor: file_storePluginService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_storePluginService_proto_goTypes,
		DependencyIndexes: file_storePluginService_proto_depIdxs,
		MessageInfos:      file_storePluginService_proto_msgTypes,
	}.Build()
	File_storePluginService_proto = out.File
	file_storePluginService_proto_rawDesc = nil
	file_storePluginService_proto_goTypes = nil
	file_storePluginService_proto_depIdxs = nil
}
