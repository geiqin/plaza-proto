// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: warehouseGoodsService.proto

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

type WarehouseGoods struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`                                      //ID
	WarehouseId int64                 `protobuf:"varint,2,opt,name=warehouse_id,json=warehouseId,proto3" json:"warehouse_id"` //仓库id
	SpuId       int64                 `protobuf:"varint,3,opt,name=spu_id,json=spuId,proto3" json:"spu_id"`                   //商品id
	Inventory   int64                 `protobuf:"varint,4,opt,name=inventory,proto3" json:"inventory"`                        //总库存
	IsEnable    string                `protobuf:"bytes,5,opt,name=is_enable,json=isEnable,proto3" json:"is_enable"`           //是否启用（0否，1是）
	CreatedAt   string                `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt   string                `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	Warehouse   *Warehouse            `protobuf:"bytes,9,opt,name=warehouse,proto3" json:"warehouse"`
	Spu         *Spu                  `protobuf:"bytes,10,opt,name=spu,proto3" json:"spu"`
	SpecList    []*WarehouseGoodsSpec `protobuf:"bytes,11,rep,name=spec_list,json=specList,proto3" json:"spec_list"`
}

func (x *WarehouseGoods) Reset() {
	*x = WarehouseGoods{}
	if protoimpl.UnsafeEnabled {
		mi := &file_warehouseGoodsService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WarehouseGoods) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WarehouseGoods) ProtoMessage() {}

func (x *WarehouseGoods) ProtoReflect() protoreflect.Message {
	mi := &file_warehouseGoodsService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WarehouseGoods.ProtoReflect.Descriptor instead.
func (*WarehouseGoods) Descriptor() ([]byte, []int) {
	return file_warehouseGoodsService_proto_rawDescGZIP(), []int{0}
}

func (x *WarehouseGoods) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *WarehouseGoods) GetWarehouseId() int64 {
	if x != nil {
		return x.WarehouseId
	}
	return 0
}

func (x *WarehouseGoods) GetSpuId() int64 {
	if x != nil {
		return x.SpuId
	}
	return 0
}

func (x *WarehouseGoods) GetInventory() int64 {
	if x != nil {
		return x.Inventory
	}
	return 0
}

func (x *WarehouseGoods) GetIsEnable() string {
	if x != nil {
		return x.IsEnable
	}
	return ""
}

func (x *WarehouseGoods) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *WarehouseGoods) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *WarehouseGoods) GetWarehouse() *Warehouse {
	if x != nil {
		return x.Warehouse
	}
	return nil
}

func (x *WarehouseGoods) GetSpu() *Spu {
	if x != nil {
		return x.Spu
	}
	return nil
}

func (x *WarehouseGoods) GetSpecList() []*WarehouseGoodsSpec {
	if x != nil {
		return x.SpecList
	}
	return nil
}

type WarehouseGoodsSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               int64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`                                                       //ID
	WarehouseGoodsId int64                 `protobuf:"varint,2,opt,name=warehouse_goods_id,json=warehouseGoodsId,proto3" json:"warehouse_goods_id"` //仓库商品id
	WarehouseId      int64                 `protobuf:"varint,3,opt,name=warehouse_id,json=warehouseId,proto3" json:"warehouse_id"`                  //仓库id
	SpuId            int64                 `protobuf:"varint,4,opt,name=spu_id,json=spuId,proto3" json:"spu_id"`                                    //商品id
	Md5Key           string                `protobuf:"bytes,5,opt,name=md5_key,json=md5Key,proto3" json:"md5_key"`                                  //MD5Key
	Inventory        int64                 `protobuf:"varint,6,opt,name=inventory,proto3" json:"inventory"`                                         //库存
	Spec             []*GoodsSpecValueItem `protobuf:"bytes,7,rep,name=spec,proto3" json:"spec"`                                                    //规格值
	CreatedAt        string                `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt        string                `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	SpecText         string                `protobuf:"bytes,10,opt,name=spec_text,json=specText,proto3" json:"spec_text"`
}

func (x *WarehouseGoodsSpec) Reset() {
	*x = WarehouseGoodsSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_warehouseGoodsService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WarehouseGoodsSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WarehouseGoodsSpec) ProtoMessage() {}

func (x *WarehouseGoodsSpec) ProtoReflect() protoreflect.Message {
	mi := &file_warehouseGoodsService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WarehouseGoodsSpec.ProtoReflect.Descriptor instead.
func (*WarehouseGoodsSpec) Descriptor() ([]byte, []int) {
	return file_warehouseGoodsService_proto_rawDescGZIP(), []int{1}
}

func (x *WarehouseGoodsSpec) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *WarehouseGoodsSpec) GetWarehouseGoodsId() int64 {
	if x != nil {
		return x.WarehouseGoodsId
	}
	return 0
}

func (x *WarehouseGoodsSpec) GetWarehouseId() int64 {
	if x != nil {
		return x.WarehouseId
	}
	return 0
}

func (x *WarehouseGoodsSpec) GetSpuId() int64 {
	if x != nil {
		return x.SpuId
	}
	return 0
}

func (x *WarehouseGoodsSpec) GetMd5Key() string {
	if x != nil {
		return x.Md5Key
	}
	return ""
}

func (x *WarehouseGoodsSpec) GetInventory() int64 {
	if x != nil {
		return x.Inventory
	}
	return 0
}

func (x *WarehouseGoodsSpec) GetSpec() []*GoodsSpecValueItem {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *WarehouseGoodsSpec) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *WarehouseGoodsSpec) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *WarehouseGoodsSpec) GetSpecText() string {
	if x != nil {
		return x.SpecText
	}
	return ""
}

type GoodsSpecValueItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  string `protobuf:"bytes,1,opt,name=type,proto3" json:"type"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value"`
}

func (x *GoodsSpecValueItem) Reset() {
	*x = GoodsSpecValueItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_warehouseGoodsService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoodsSpecValueItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodsSpecValueItem) ProtoMessage() {}

func (x *GoodsSpecValueItem) ProtoReflect() protoreflect.Message {
	mi := &file_warehouseGoodsService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodsSpecValueItem.ProtoReflect.Descriptor instead.
func (*GoodsSpecValueItem) Descriptor() ([]byte, []int) {
	return file_warehouseGoodsService_proto_rawDescGZIP(), []int{2}
}

func (x *GoodsSpecValueItem) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *GoodsSpecValueItem) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type WarehouseGoodsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged       int64   `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize    int64   `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Sorting     string  `protobuf:"bytes,3,opt,name=sorting,proto3" json:"sorting"`
	Id          int64   `protobuf:"varint,4,opt,name=id,proto3" json:"id"`
	WarehouseId int64   `protobuf:"varint,5,opt,name=warehouse_id,json=warehouseId,proto3" json:"warehouse_id"` //仓库id
	SpuId       int64   `protobuf:"varint,6,opt,name=spu_id,json=spuId,proto3" json:"spu_id"`                   //商品id
	Inventory   int64   `protobuf:"varint,7,opt,name=inventory,proto3" json:"inventory"`                        //总库存
	IsEnable    string  `protobuf:"bytes,8,opt,name=is_enable,json=isEnable,proto3" json:"is_enable"`           //是否启用（0否，1是）
	CategoryId  int64   `protobuf:"varint,9,opt,name=category_id,json=categoryId,proto3" json:"category_id"`    //商品分类ID
	BrandId     int32   `protobuf:"varint,10,opt,name=brand_id,json=brandId,proto3" json:"brand_id"`            //品牌ID
	ShopId      int64   `protobuf:"varint,11,opt,name=shop_id,json=shopId,proto3" json:"shop_id"`               //分店ID
	TermId      int64   `protobuf:"varint,12,opt,name=term_id,json=termId,proto3" json:"term_id"`
	Title       string  `protobuf:"bytes,13,opt,name=title,proto3" json:"title"`       //商品标题
	Keywords    string  `protobuf:"bytes,14,opt,name=keywords,proto3" json:"keywords"` //商品标题
	Ids         []int64 `protobuf:"varint,15,rep,packed,name=ids,proto3" json:"ids"`
	SpuIds      []int64 `protobuf:"varint,16,rep,packed,name=spu_ids,json=spuIds,proto3" json:"spu_ids"`
	TermIds     []int64 `protobuf:"varint,17,rep,packed,name=term_ids,json=termIds,proto3" json:"term_ids"`
}

func (x *WarehouseGoodsRequest) Reset() {
	*x = WarehouseGoodsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_warehouseGoodsService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WarehouseGoodsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WarehouseGoodsRequest) ProtoMessage() {}

func (x *WarehouseGoodsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_warehouseGoodsService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WarehouseGoodsRequest.ProtoReflect.Descriptor instead.
func (*WarehouseGoodsRequest) Descriptor() ([]byte, []int) {
	return file_warehouseGoodsService_proto_rawDescGZIP(), []int{3}
}

func (x *WarehouseGoodsRequest) GetPaged() int64 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *WarehouseGoodsRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *WarehouseGoodsRequest) GetSorting() string {
	if x != nil {
		return x.Sorting
	}
	return ""
}

func (x *WarehouseGoodsRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *WarehouseGoodsRequest) GetWarehouseId() int64 {
	if x != nil {
		return x.WarehouseId
	}
	return 0
}

func (x *WarehouseGoodsRequest) GetSpuId() int64 {
	if x != nil {
		return x.SpuId
	}
	return 0
}

func (x *WarehouseGoodsRequest) GetInventory() int64 {
	if x != nil {
		return x.Inventory
	}
	return 0
}

func (x *WarehouseGoodsRequest) GetIsEnable() string {
	if x != nil {
		return x.IsEnable
	}
	return ""
}

func (x *WarehouseGoodsRequest) GetCategoryId() int64 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

func (x *WarehouseGoodsRequest) GetBrandId() int32 {
	if x != nil {
		return x.BrandId
	}
	return 0
}

func (x *WarehouseGoodsRequest) GetShopId() int64 {
	if x != nil {
		return x.ShopId
	}
	return 0
}

func (x *WarehouseGoodsRequest) GetTermId() int64 {
	if x != nil {
		return x.TermId
	}
	return 0
}

func (x *WarehouseGoodsRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *WarehouseGoodsRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *WarehouseGoodsRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *WarehouseGoodsRequest) GetSpuIds() []int64 {
	if x != nil {
		return x.SpuIds
	}
	return nil
}

func (x *WarehouseGoodsRequest) GetTermIds() []int64 {
	if x != nil {
		return x.TermIds
	}
	return nil
}

type WarehouseGoodsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *WarehouseGoods   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager     `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*WarehouseGoods `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Info   string            `protobuf:"bytes,4,opt,name=info,proto3" json:"info"`
}

func (x *WarehouseGoodsResponse) Reset() {
	*x = WarehouseGoodsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_warehouseGoodsService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WarehouseGoodsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WarehouseGoodsResponse) ProtoMessage() {}

func (x *WarehouseGoodsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_warehouseGoodsService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WarehouseGoodsResponse.ProtoReflect.Descriptor instead.
func (*WarehouseGoodsResponse) Descriptor() ([]byte, []int) {
	return file_warehouseGoodsService_proto_rawDescGZIP(), []int{4}
}

func (x *WarehouseGoodsResponse) GetEntity() *WarehouseGoods {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *WarehouseGoodsResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *WarehouseGoodsResponse) GetItems() []*WarehouseGoods {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *WarehouseGoodsResponse) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

type WarehouseGoodsSelectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pager *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items []*Spu        `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Info  string        `protobuf:"bytes,4,opt,name=info,proto3" json:"info"`
}

func (x *WarehouseGoodsSelectResponse) Reset() {
	*x = WarehouseGoodsSelectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_warehouseGoodsService_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WarehouseGoodsSelectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WarehouseGoodsSelectResponse) ProtoMessage() {}

func (x *WarehouseGoodsSelectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_warehouseGoodsService_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WarehouseGoodsSelectResponse.ProtoReflect.Descriptor instead.
func (*WarehouseGoodsSelectResponse) Descriptor() ([]byte, []int) {
	return file_warehouseGoodsService_proto_rawDescGZIP(), []int{5}
}

func (x *WarehouseGoodsSelectResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *WarehouseGoodsSelectResponse) GetItems() []*Spu {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *WarehouseGoodsSelectResponse) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

var File_warehouseGoodsService_proto protoreflect.FileDescriptor

var file_warehouseGoodsService_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x73, 0x70, 0x75, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x77, 0x61,
	0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe2, 0x02, 0x0a, 0x0e, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75,
	0x73, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x77, 0x61, 0x72, 0x65, 0x68,
	0x6f, 0x75, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x77,
	0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x70,
	0x75, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x70, 0x75, 0x49,
	0x64, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x12,
	0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x69, 0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x31, 0x0a, 0x09, 0x77, 0x61,
	0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75,
	0x73, 0x65, 0x52, 0x09, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x12, 0x1f, 0x0a,
	0x03, 0x73, 0x70, 0x75, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x70, 0x75, 0x52, 0x03, 0x73, 0x70, 0x75, 0x12, 0x39,
	0x0a, 0x09, 0x73, 0x70, 0x65, 0x63, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0b, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x57, 0x61, 0x72,
	0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x53, 0x70, 0x65, 0x63, 0x52,
	0x08, 0x73, 0x70, 0x65, 0x63, 0x4c, 0x69, 0x73, 0x74, 0x22, 0xd0, 0x02, 0x0a, 0x12, 0x57, 0x61,
	0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x53, 0x70, 0x65, 0x63,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x2c, 0x0a, 0x12, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x5f, 0x67, 0x6f,
	0x6f, 0x64, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x77, 0x61,
	0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x64, 0x12, 0x21,
	0x0a, 0x0c, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x49,
	0x64, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x70, 0x75, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x73, 0x70, 0x75, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x6d, 0x64, 0x35, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x64, 0x35, 0x4b, 0x65,
	0x79, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x12,
	0x30, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x53, 0x70,
	0x65, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x73, 0x70, 0x65,
	0x63, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x73, 0x70, 0x65, 0x63, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x73, 0x70, 0x65, 0x63, 0x54, 0x65, 0x78, 0x74, 0x22, 0x3e, 0x0a, 0x12,
	0x47, 0x6f, 0x6f, 0x64, 0x73, 0x53, 0x70, 0x65, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xcf, 0x03, 0x0a,
	0x15, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6f, 0x72,
	0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6f, 0x72, 0x74,
	0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x77, 0x61, 0x72, 0x65, 0x68,
	0x6f, 0x75, 0x73, 0x65, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x70, 0x75, 0x5f, 0x69, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x70, 0x75, 0x49, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x69,
	0x73, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x69, 0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x72, 0x61,
	0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x62, 0x72, 0x61,
	0x6e, 0x64, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x74, 0x65, 0x72, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x74, 0x65, 0x72, 0x6d, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18,
	0x0f, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x70,
	0x75, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x10, 0x20, 0x03, 0x28, 0x03, 0x52, 0x06, 0x73, 0x70, 0x75,
	0x49, 0x64, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x65, 0x72, 0x6d, 0x5f, 0x69, 0x64, 0x73, 0x18,
	0x11, 0x20, 0x03, 0x28, 0x03, 0x52, 0x07, 0x74, 0x65, 0x72, 0x6d, 0x49, 0x64, 0x73, 0x22, 0xb3,
	0x01, 0x0a, 0x16, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x47, 0x6f, 0x6f, 0x64,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x47, 0x6f,
	0x6f, 0x64, 0x73, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70,
	0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72,
	0x12, 0x2e, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x57, 0x61, 0x72, 0x65, 0x68,
	0x6f, 0x75, 0x73, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x69, 0x6e, 0x66, 0x6f, 0x22, 0x7c, 0x0a, 0x1c, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73,
	0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67,
	0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x53, 0x70, 0x75, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x32, 0xb6, 0x04, 0x0a, 0x15, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65,
	0x47, 0x6f, 0x6f, 0x64, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x06,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73,
	0x1a, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x57, 0x61, 0x72, 0x65,
	0x68, 0x6f, 0x75, 0x73, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e,
	0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x47, 0x6f, 0x6f, 0x64,
	0x73, 0x1a, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x57, 0x61, 0x72,
	0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12,
	0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x57, 0x61, 0x72, 0x65, 0x68,
	0x6f, 0x75, 0x73, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x1a, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x47, 0x6f,
	0x6f, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a,
	0x06, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x47, 0x6f, 0x6f, 0x64,
	0x73, 0x1a, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x57, 0x61, 0x72,
	0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1f, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75,
	0x73, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f,
	0x75, 0x73, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x4d, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1f, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73,
	0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75,
	0x73, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x58, 0x0a, 0x0b, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73,
	0x12, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x57, 0x61, 0x72, 0x65,
	0x68, 0x6f, 0x75, 0x73, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x26, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x57, 0x61, 0x72,
	0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x53, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2f,
	0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_warehouseGoodsService_proto_rawDescOnce sync.Once
	file_warehouseGoodsService_proto_rawDescData = file_warehouseGoodsService_proto_rawDesc
)

func file_warehouseGoodsService_proto_rawDescGZIP() []byte {
	file_warehouseGoodsService_proto_rawDescOnce.Do(func() {
		file_warehouseGoodsService_proto_rawDescData = protoimpl.X.CompressGZIP(file_warehouseGoodsService_proto_rawDescData)
	})
	return file_warehouseGoodsService_proto_rawDescData
}

var file_warehouseGoodsService_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_warehouseGoodsService_proto_goTypes = []interface{}{
	(*WarehouseGoods)(nil),               // 0: services.WarehouseGoods
	(*WarehouseGoodsSpec)(nil),           // 1: services.WarehouseGoodsSpec
	(*GoodsSpecValueItem)(nil),           // 2: services.GoodsSpecValueItem
	(*WarehouseGoodsRequest)(nil),        // 3: services.WarehouseGoodsRequest
	(*WarehouseGoodsResponse)(nil),       // 4: services.WarehouseGoodsResponse
	(*WarehouseGoodsSelectResponse)(nil), // 5: services.WarehouseGoodsSelectResponse
	(*Warehouse)(nil),                    // 6: services.Warehouse
	(*Spu)(nil),                          // 7: services.Spu
	(*common.Pager)(nil),                 // 8: common.Pager
}
var file_warehouseGoodsService_proto_depIdxs = []int32{
	6,  // 0: services.WarehouseGoods.warehouse:type_name -> services.Warehouse
	7,  // 1: services.WarehouseGoods.spu:type_name -> services.Spu
	1,  // 2: services.WarehouseGoods.spec_list:type_name -> services.WarehouseGoodsSpec
	2,  // 3: services.WarehouseGoodsSpec.spec:type_name -> services.GoodsSpecValueItem
	0,  // 4: services.WarehouseGoodsResponse.entity:type_name -> services.WarehouseGoods
	8,  // 5: services.WarehouseGoodsResponse.pager:type_name -> common.Pager
	0,  // 6: services.WarehouseGoodsResponse.items:type_name -> services.WarehouseGoods
	8,  // 7: services.WarehouseGoodsSelectResponse.pager:type_name -> common.Pager
	7,  // 8: services.WarehouseGoodsSelectResponse.items:type_name -> services.Spu
	0,  // 9: services.WarehouseGoodsService.Create:input_type -> services.WarehouseGoods
	0,  // 10: services.WarehouseGoodsService.UpdateInventory:input_type -> services.WarehouseGoods
	0,  // 11: services.WarehouseGoodsService.Delete:input_type -> services.WarehouseGoods
	0,  // 12: services.WarehouseGoodsService.Detail:input_type -> services.WarehouseGoods
	3,  // 13: services.WarehouseGoodsService.List:input_type -> services.WarehouseGoodsRequest
	3,  // 14: services.WarehouseGoodsService.Search:input_type -> services.WarehouseGoodsRequest
	3,  // 15: services.WarehouseGoodsService.SelectGoods:input_type -> services.WarehouseGoodsRequest
	4,  // 16: services.WarehouseGoodsService.Create:output_type -> services.WarehouseGoodsResponse
	4,  // 17: services.WarehouseGoodsService.UpdateInventory:output_type -> services.WarehouseGoodsResponse
	4,  // 18: services.WarehouseGoodsService.Delete:output_type -> services.WarehouseGoodsResponse
	4,  // 19: services.WarehouseGoodsService.Detail:output_type -> services.WarehouseGoodsResponse
	4,  // 20: services.WarehouseGoodsService.List:output_type -> services.WarehouseGoodsResponse
	4,  // 21: services.WarehouseGoodsService.Search:output_type -> services.WarehouseGoodsResponse
	5,  // 22: services.WarehouseGoodsService.SelectGoods:output_type -> services.WarehouseGoodsSelectResponse
	16, // [16:23] is the sub-list for method output_type
	9,  // [9:16] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_warehouseGoodsService_proto_init() }
func file_warehouseGoodsService_proto_init() {
	if File_warehouseGoodsService_proto != nil {
		return
	}
	file_spuService_proto_init()
	file_warehouseService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_warehouseGoodsService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WarehouseGoods); i {
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
		file_warehouseGoodsService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WarehouseGoodsSpec); i {
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
		file_warehouseGoodsService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoodsSpecValueItem); i {
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
		file_warehouseGoodsService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WarehouseGoodsRequest); i {
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
		file_warehouseGoodsService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WarehouseGoodsResponse); i {
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
		file_warehouseGoodsService_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WarehouseGoodsSelectResponse); i {
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
			RawDescriptor: file_warehouseGoodsService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_warehouseGoodsService_proto_goTypes,
		DependencyIndexes: file_warehouseGoodsService_proto_depIdxs,
		MessageInfos:      file_warehouseGoodsService_proto_msgTypes,
	}.Build()
	File_warehouseGoodsService_proto = out.File
	file_warehouseGoodsService_proto_rawDesc = nil
	file_warehouseGoodsService_proto_goTypes = nil
	file_warehouseGoodsService_proto_depIdxs = nil
}
