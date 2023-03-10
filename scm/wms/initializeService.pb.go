// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: initializeService.proto

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

type Initialize struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64               `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	InitializeSn string              `protobuf:"bytes,2,opt,name=initialize_sn,json=initializeSn,proto3" json:"initialize_sn"`
	WarehouseId  int64               `protobuf:"varint,3,opt,name=warehouse_id,json=warehouseId,proto3" json:"warehouse_id"`
	TotalNum     int32               `protobuf:"varint,4,opt,name=total_num,json=totalNum,proto3" json:"total_num"`
	TotalCost    int64               `protobuf:"varint,5,opt,name=total_cost,json=totalCost,proto3" json:"total_cost"`
	FailReason   string              `protobuf:"bytes,8,opt,name=fail_reason,json=failReason,proto3" json:"fail_reason"`
	Memo         string              `protobuf:"bytes,9,opt,name=memo,proto3" json:"memo"`
	Status       string              `protobuf:"bytes,10,opt,name=status,proto3" json:"status"`
	OperatorId   int64               `protobuf:"varint,11,opt,name=operator_id,json=operatorId,proto3" json:"operator_id"`
	CreatedAt    string              `protobuf:"bytes,12,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt    string              `protobuf:"bytes,13,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	Details      []*InitializeDetail `protobuf:"bytes,14,rep,name=details,proto3" json:"details"`
	Warehouse    *Warehouse          `protobuf:"bytes,15,opt,name=warehouse,proto3" json:"warehouse"`
	Operator     *UserInfo           `protobuf:"bytes,16,opt,name=operator,proto3" json:"operator"`
}

func (x *Initialize) Reset() {
	*x = Initialize{}
	if protoimpl.UnsafeEnabled {
		mi := &file_initializeService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Initialize) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Initialize) ProtoMessage() {}

func (x *Initialize) ProtoReflect() protoreflect.Message {
	mi := &file_initializeService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Initialize.ProtoReflect.Descriptor instead.
func (*Initialize) Descriptor() ([]byte, []int) {
	return file_initializeService_proto_rawDescGZIP(), []int{0}
}

func (x *Initialize) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Initialize) GetInitializeSn() string {
	if x != nil {
		return x.InitializeSn
	}
	return ""
}

func (x *Initialize) GetWarehouseId() int64 {
	if x != nil {
		return x.WarehouseId
	}
	return 0
}

func (x *Initialize) GetTotalNum() int32 {
	if x != nil {
		return x.TotalNum
	}
	return 0
}

func (x *Initialize) GetTotalCost() int64 {
	if x != nil {
		return x.TotalCost
	}
	return 0
}

func (x *Initialize) GetFailReason() string {
	if x != nil {
		return x.FailReason
	}
	return ""
}

func (x *Initialize) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *Initialize) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Initialize) GetOperatorId() int64 {
	if x != nil {
		return x.OperatorId
	}
	return 0
}

func (x *Initialize) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Initialize) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Initialize) GetDetails() []*InitializeDetail {
	if x != nil {
		return x.Details
	}
	return nil
}

func (x *Initialize) GetWarehouse() *Warehouse {
	if x != nil {
		return x.Warehouse
	}
	return nil
}

func (x *Initialize) GetOperator() *UserInfo {
	if x != nil {
		return x.Operator
	}
	return nil
}

type InitializeDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	InitializeId    int64    `protobuf:"varint,2,opt,name=initialize_id,json=initializeId,proto3" json:"initialize_id"`
	SpuId           int64    `protobuf:"varint,3,opt,name=spu_id,json=spuId,proto3" json:"spu_id"`
	SkuId           int64    `protobuf:"varint,4,opt,name=sku_id,json=skuId,proto3" json:"sku_id"`
	Num             int32    `protobuf:"varint,5,opt,name=num,proto3" json:"num"`
	CostPrice       int64    `protobuf:"varint,6,opt,name=cost_price,json=costPrice,proto3" json:"cost_price"`
	SubtotalCost    int64    `protobuf:"varint,7,opt,name=subtotal_cost,json=subtotalCost,proto3" json:"subtotal_cost"`
	ProductionDate  string   `protobuf:"bytes,8,opt,name=production_date,json=productionDate,proto3" json:"production_date"`
	ProductionBatch int64    `protobuf:"varint,9,opt,name=production_batch,json=productionBatch,proto3" json:"production_batch"`
	Memo            string   `protobuf:"bytes,10,opt,name=memo,proto3" json:"memo"`
	CreatedAt       string   `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt       string   `protobuf:"bytes,12,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	Sku             *SkuInfo `protobuf:"bytes,13,opt,name=sku,proto3" json:"sku"`
}

func (x *InitializeDetail) Reset() {
	*x = InitializeDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_initializeService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitializeDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitializeDetail) ProtoMessage() {}

func (x *InitializeDetail) ProtoReflect() protoreflect.Message {
	mi := &file_initializeService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitializeDetail.ProtoReflect.Descriptor instead.
func (*InitializeDetail) Descriptor() ([]byte, []int) {
	return file_initializeService_proto_rawDescGZIP(), []int{1}
}

func (x *InitializeDetail) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *InitializeDetail) GetInitializeId() int64 {
	if x != nil {
		return x.InitializeId
	}
	return 0
}

func (x *InitializeDetail) GetSpuId() int64 {
	if x != nil {
		return x.SpuId
	}
	return 0
}

func (x *InitializeDetail) GetSkuId() int64 {
	if x != nil {
		return x.SkuId
	}
	return 0
}

func (x *InitializeDetail) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *InitializeDetail) GetCostPrice() int64 {
	if x != nil {
		return x.CostPrice
	}
	return 0
}

func (x *InitializeDetail) GetSubtotalCost() int64 {
	if x != nil {
		return x.SubtotalCost
	}
	return 0
}

func (x *InitializeDetail) GetProductionDate() string {
	if x != nil {
		return x.ProductionDate
	}
	return ""
}

func (x *InitializeDetail) GetProductionBatch() int64 {
	if x != nil {
		return x.ProductionBatch
	}
	return 0
}

func (x *InitializeDetail) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *InitializeDetail) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *InitializeDetail) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *InitializeDetail) GetSku() *SkuInfo {
	if x != nil {
		return x.Sku
	}
	return nil
}

type InitializeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged        int32   `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize     int32   `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Id           int64   `protobuf:"varint,3,opt,name=id,proto3" json:"id"`
	Type         string  `protobuf:"bytes,5,opt,name=type,proto3" json:"type"`
	SheetType    int32   `protobuf:"varint,6,opt,name=sheet_type,json=sheetType,proto3" json:"sheet_type"`
	Status       string  `protobuf:"bytes,7,opt,name=status,proto3" json:"status"`
	WarehouseId  int64   `protobuf:"varint,8,opt,name=warehouse_id,json=warehouseId,proto3" json:"warehouse_id"`
	Failure      string  `protobuf:"bytes,9,opt,name=failure,proto3" json:"failure"`
	InitializeId int64   `protobuf:"varint,10,opt,name=initialize_id,json=initializeId,proto3" json:"initialize_id"`
	Ids          []int64 `protobuf:"varint,4,rep,packed,name=ids,proto3" json:"ids"`
}

func (x *InitializeRequest) Reset() {
	*x = InitializeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_initializeService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitializeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitializeRequest) ProtoMessage() {}

func (x *InitializeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_initializeService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitializeRequest.ProtoReflect.Descriptor instead.
func (*InitializeRequest) Descriptor() ([]byte, []int) {
	return file_initializeService_proto_rawDescGZIP(), []int{2}
}

func (x *InitializeRequest) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *InitializeRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *InitializeRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *InitializeRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *InitializeRequest) GetSheetType() int32 {
	if x != nil {
		return x.SheetType
	}
	return 0
}

func (x *InitializeRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *InitializeRequest) GetWarehouseId() int64 {
	if x != nil {
		return x.WarehouseId
	}
	return 0
}

func (x *InitializeRequest) GetFailure() string {
	if x != nil {
		return x.Failure
	}
	return ""
}

func (x *InitializeRequest) GetInitializeId() int64 {
	if x != nil {
		return x.InitializeId
	}
	return 0
}

func (x *InitializeRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type InitializeData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity  *Initialize         `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager   *common.Pager       `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items   []*Initialize       `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Details []*InitializeDetail `protobuf:"bytes,4,rep,name=details,proto3" json:"details"`
	Info    *common.Info        `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
}

func (x *InitializeData) Reset() {
	*x = InitializeData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_initializeService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitializeData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitializeData) ProtoMessage() {}

func (x *InitializeData) ProtoReflect() protoreflect.Message {
	mi := &file_initializeService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitializeData.ProtoReflect.Descriptor instead.
func (*InitializeData) Descriptor() ([]byte, []int) {
	return file_initializeService_proto_rawDescGZIP(), []int{3}
}

func (x *InitializeData) GetEntity() *Initialize {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *InitializeData) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *InitializeData) GetItems() []*Initialize {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *InitializeData) GetDetails() []*InitializeDetail {
	if x != nil {
		return x.Details
	}
	return nil
}

func (x *InitializeData) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type InitializeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *InitializeData `protobuf:"bytes,1,opt,name=data,proto3" json:"data"`
	Error *common.Error   `protobuf:"bytes,2,opt,name=error,proto3" json:"error"`
}

func (x *InitializeResponse) Reset() {
	*x = InitializeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_initializeService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitializeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitializeResponse) ProtoMessage() {}

func (x *InitializeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_initializeService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitializeResponse.ProtoReflect.Descriptor instead.
func (*InitializeResponse) Descriptor() ([]byte, []int) {
	return file_initializeService_proto_rawDescGZIP(), []int{4}
}

func (x *InitializeResponse) GetData() *InitializeData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *InitializeResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_initializeService_proto protoreflect.FileDescriptor

var file_initializeService_proto_rawDesc = []byte{
	0x0a, 0x17, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x62, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x77,
	0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe5, 0x03, 0x0a, 0x0a, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61,
	0x6c, 0x69, 0x7a, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69,
	0x7a, 0x65, 0x5f, 0x73, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x53, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x77, 0x61, 0x72,
	0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0b, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x61, 0x69, 0x6c,
	0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x6d,
	0x6f, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x34, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18,
	0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x31, 0x0a, 0x09, 0x77, 0x61,
	0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75,
	0x73, 0x65, 0x52, 0x09, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x12, 0x2e, 0x0a,
	0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x96, 0x03,
	0x0a, 0x10, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x61, 0x6c, 0x69, 0x7a, 0x65, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x70, 0x75, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x70, 0x75, 0x49, 0x64, 0x12, 0x15,
	0x0a, 0x06, 0x73, 0x6b, 0x75, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x73, 0x6b, 0x75, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x73, 0x74, 0x5f,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x73,
	0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x73,
	0x75, 0x62, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12,
	0x12, 0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d,
	0x65, 0x6d, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x23, 0x0a, 0x03, 0x73, 0x6b, 0x75, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x6b, 0x75, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x03, 0x73, 0x6b, 0x75, 0x22, 0x95, 0x02, 0x0a, 0x11, 0x49, 0x6e, 0x69, 0x74, 0x69,
	0x61, 0x6c, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x67,
	0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x68, 0x65, 0x65, 0x74, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x68, 0x65, 0x65, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x77, 0x61,
	0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0b, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x61, 0x6c, 0x69, 0x7a, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x69, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0xe7,
	0x01, 0x0a, 0x0e, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x2c, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x69,
	0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12,
	0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70,
	0x61, 0x67, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49,
	0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x12, 0x34, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x69,
	0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x07, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x67, 0x0a, 0x12, 0x49, 0x6e, 0x69, 0x74,
	0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69,
	0x7a, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x32, 0xd7, 0x01, 0x0a, 0x11, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x69,
	0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x14,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61,
	0x6c, 0x69, 0x7a, 0x65, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1b,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61,
	0x6c, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2f,
	0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_initializeService_proto_rawDescOnce sync.Once
	file_initializeService_proto_rawDescData = file_initializeService_proto_rawDesc
)

func file_initializeService_proto_rawDescGZIP() []byte {
	file_initializeService_proto_rawDescOnce.Do(func() {
		file_initializeService_proto_rawDescData = protoimpl.X.CompressGZIP(file_initializeService_proto_rawDescData)
	})
	return file_initializeService_proto_rawDescData
}

var file_initializeService_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_initializeService_proto_goTypes = []interface{}{
	(*Initialize)(nil),         // 0: services.Initialize
	(*InitializeDetail)(nil),   // 1: services.InitializeDetail
	(*InitializeRequest)(nil),  // 2: services.InitializeRequest
	(*InitializeData)(nil),     // 3: services.InitializeData
	(*InitializeResponse)(nil), // 4: services.InitializeResponse
	(*Warehouse)(nil),          // 5: services.Warehouse
	(*UserInfo)(nil),           // 6: services.UserInfo
	(*SkuInfo)(nil),            // 7: services.SkuInfo
	(*common.Pager)(nil),       // 8: common.Pager
	(*common.Info)(nil),        // 9: common.Info
	(*common.Error)(nil),       // 10: common.Error
}
var file_initializeService_proto_depIdxs = []int32{
	1,  // 0: services.Initialize.details:type_name -> services.InitializeDetail
	5,  // 1: services.Initialize.warehouse:type_name -> services.Warehouse
	6,  // 2: services.Initialize.operator:type_name -> services.UserInfo
	7,  // 3: services.InitializeDetail.sku:type_name -> services.SkuInfo
	0,  // 4: services.InitializeData.entity:type_name -> services.Initialize
	8,  // 5: services.InitializeData.pager:type_name -> common.Pager
	0,  // 6: services.InitializeData.items:type_name -> services.Initialize
	1,  // 7: services.InitializeData.details:type_name -> services.InitializeDetail
	9,  // 8: services.InitializeData.info:type_name -> common.Info
	3,  // 9: services.InitializeResponse.data:type_name -> services.InitializeData
	10, // 10: services.InitializeResponse.error:type_name -> common.Error
	0,  // 11: services.InitializeService.Create:input_type -> services.Initialize
	0,  // 12: services.InitializeService.Get:input_type -> services.Initialize
	2,  // 13: services.InitializeService.Search:input_type -> services.InitializeRequest
	4,  // 14: services.InitializeService.Create:output_type -> services.InitializeResponse
	4,  // 15: services.InitializeService.Get:output_type -> services.InitializeResponse
	4,  // 16: services.InitializeService.Search:output_type -> services.InitializeResponse
	14, // [14:17] is the sub-list for method output_type
	11, // [11:14] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_initializeService_proto_init() }
func file_initializeService_proto_init() {
	if File_initializeService_proto != nil {
		return
	}
	file_baseInfoService_proto_init()
	file_warehouseService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_initializeService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Initialize); i {
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
		file_initializeService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitializeDetail); i {
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
		file_initializeService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitializeRequest); i {
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
		file_initializeService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitializeData); i {
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
		file_initializeService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitializeResponse); i {
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
			RawDescriptor: file_initializeService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_initializeService_proto_goTypes,
		DependencyIndexes: file_initializeService_proto_depIdxs,
		MessageInfos:      file_initializeService_proto_msgTypes,
	}.Build()
	File_initializeService_proto = out.File
	file_initializeService_proto_rawDesc = nil
	file_initializeService_proto_goTypes = nil
	file_initializeService_proto_depIdxs = nil
}