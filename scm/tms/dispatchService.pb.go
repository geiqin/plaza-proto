// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: dispatchService.proto

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

type Dispatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int64             `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Code           string            `protobuf:"bytes,2,opt,name=code,proto3" json:"code"`
	OrderId        int64             `protobuf:"varint,3,opt,name=order_id,json=orderId,proto3" json:"order_id"`
	ShipmentCode   string            `protobuf:"bytes,4,opt,name=shipment_code,json=shipmentCode,proto3" json:"shipment_code"`
	Type           string            `protobuf:"bytes,5,opt,name=type,proto3" json:"type"`
	Freight        int64             `protobuf:"varint,6,opt,name=freight,proto3" json:"freight"`
	Protected      bool              `protobuf:"varint,7,opt,name=protected,proto3" json:"protected"`
	IsDelivery     bool              `protobuf:"varint,8,opt,name=is_delivery,json=isDelivery,proto3" json:"is_delivery"`
	IsAutoSurface  bool              `protobuf:"varint,9,opt,name=is_auto_surface,json=isAutoSurface,proto3" json:"is_auto_surface"`
	ShipperId      int32             `protobuf:"varint,10,opt,name=shipper_id,json=shipperId,proto3" json:"shipper_id"`
	LogisticsNo    string            `protobuf:"bytes,11,opt,name=logistics_no,json=logisticsNo,proto3" json:"logistics_no"`
	FetchCode      string            `protobuf:"bytes,12,opt,name=fetch_code,json=fetchCode,proto3" json:"fetch_code"`
	LocationId     int64             `protobuf:"varint,13,opt,name=location_id,json=locationId,proto3" json:"location_id"`
	FetchAt        string            `protobuf:"bytes,14,opt,name=fetch_at,json=fetchAt,proto3" json:"fetch_at"`
	DeliveryAt     string            `protobuf:"bytes,15,opt,name=delivery_at,json=deliveryAt,proto3" json:"delivery_at"`
	DeliveryType   string            `protobuf:"bytes,16,opt,name=delivery_type,json=deliveryType,proto3" json:"delivery_type"`
	MemberId       int64             `protobuf:"varint,17,opt,name=member_id,json=memberId,proto3" json:"member_id"`
	ReceiverName   string            `protobuf:"bytes,18,opt,name=receiver_name,json=receiverName,proto3" json:"receiver_name"`
	ReceiverAreaId int64             `protobuf:"varint,19,opt,name=receiver_area_id,json=receiverAreaId,proto3" json:"receiver_area_id"`
	ReceiverAddr   string            `protobuf:"bytes,20,opt,name=receiver_addr,json=receiverAddr,proto3" json:"receiver_addr"`
	ReceiverZip    string            `protobuf:"bytes,21,opt,name=receiver_zip,json=receiverZip,proto3" json:"receiver_zip"`
	ReceiverTel    string            `protobuf:"bytes,22,opt,name=receiver_tel,json=receiverTel,proto3" json:"receiver_tel"`
	ReceiverMobile string            `protobuf:"bytes,23,opt,name=receiver_mobile,json=receiverMobile,proto3" json:"receiver_mobile"`
	ReceiverEmail  string            `protobuf:"bytes,24,opt,name=receiver_email,json=receiverEmail,proto3" json:"receiver_email"`
	CreatorId      int64             `protobuf:"varint,25,opt,name=creator_id,json=creatorId,proto3" json:"creator_id"`
	CreatorName    string            `protobuf:"bytes,26,opt,name=creator_name,json=creatorName,proto3" json:"creator_name"`
	Status         string            `protobuf:"bytes,27,opt,name=status,proto3" json:"status"`
	Memo           string            `protobuf:"bytes,28,opt,name=memo,proto3" json:"memo"`
	ArrivedAt      string            `protobuf:"bytes,29,opt,name=arrived_at,json=arrivedAt,proto3" json:"arrived_at"`
	CreatedAt      string            `protobuf:"bytes,30,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt      string            `protobuf:"bytes,31,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	Details        []*DispatchDetail `protobuf:"bytes,32,rep,name=details,proto3" json:"details"`
	StartTime      string            `protobuf:"bytes,33,opt,name=start_time,json=startTime,proto3" json:"start_time"`
	EndTime        string            `protobuf:"bytes,34,opt,name=end_time,json=endTime,proto3" json:"end_time"`
	ShippedAt      string            `protobuf:"bytes,35,opt,name=shipped_at,json=shippedAt,proto3" json:"shipped_at"`
	Shipper        *Shipper          `protobuf:"bytes,36,opt,name=shipper,proto3" json:"shipper"`
	Location       *Fetch            `protobuf:"bytes,37,opt,name=location,proto3" json:"location"`
}

func (x *Dispatch) Reset() {
	*x = Dispatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dispatchService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dispatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dispatch) ProtoMessage() {}

func (x *Dispatch) ProtoReflect() protoreflect.Message {
	mi := &file_dispatchService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dispatch.ProtoReflect.Descriptor instead.
func (*Dispatch) Descriptor() ([]byte, []int) {
	return file_dispatchService_proto_rawDescGZIP(), []int{0}
}

func (x *Dispatch) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Dispatch) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Dispatch) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *Dispatch) GetShipmentCode() string {
	if x != nil {
		return x.ShipmentCode
	}
	return ""
}

func (x *Dispatch) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Dispatch) GetFreight() int64 {
	if x != nil {
		return x.Freight
	}
	return 0
}

func (x *Dispatch) GetProtected() bool {
	if x != nil {
		return x.Protected
	}
	return false
}

func (x *Dispatch) GetIsDelivery() bool {
	if x != nil {
		return x.IsDelivery
	}
	return false
}

func (x *Dispatch) GetIsAutoSurface() bool {
	if x != nil {
		return x.IsAutoSurface
	}
	return false
}

func (x *Dispatch) GetShipperId() int32 {
	if x != nil {
		return x.ShipperId
	}
	return 0
}

func (x *Dispatch) GetLogisticsNo() string {
	if x != nil {
		return x.LogisticsNo
	}
	return ""
}

func (x *Dispatch) GetFetchCode() string {
	if x != nil {
		return x.FetchCode
	}
	return ""
}

func (x *Dispatch) GetLocationId() int64 {
	if x != nil {
		return x.LocationId
	}
	return 0
}

func (x *Dispatch) GetFetchAt() string {
	if x != nil {
		return x.FetchAt
	}
	return ""
}

func (x *Dispatch) GetDeliveryAt() string {
	if x != nil {
		return x.DeliveryAt
	}
	return ""
}

func (x *Dispatch) GetDeliveryType() string {
	if x != nil {
		return x.DeliveryType
	}
	return ""
}

func (x *Dispatch) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

func (x *Dispatch) GetReceiverName() string {
	if x != nil {
		return x.ReceiverName
	}
	return ""
}

func (x *Dispatch) GetReceiverAreaId() int64 {
	if x != nil {
		return x.ReceiverAreaId
	}
	return 0
}

func (x *Dispatch) GetReceiverAddr() string {
	if x != nil {
		return x.ReceiverAddr
	}
	return ""
}

func (x *Dispatch) GetReceiverZip() string {
	if x != nil {
		return x.ReceiverZip
	}
	return ""
}

func (x *Dispatch) GetReceiverTel() string {
	if x != nil {
		return x.ReceiverTel
	}
	return ""
}

func (x *Dispatch) GetReceiverMobile() string {
	if x != nil {
		return x.ReceiverMobile
	}
	return ""
}

func (x *Dispatch) GetReceiverEmail() string {
	if x != nil {
		return x.ReceiverEmail
	}
	return ""
}

func (x *Dispatch) GetCreatorId() int64 {
	if x != nil {
		return x.CreatorId
	}
	return 0
}

func (x *Dispatch) GetCreatorName() string {
	if x != nil {
		return x.CreatorName
	}
	return ""
}

func (x *Dispatch) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Dispatch) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *Dispatch) GetArrivedAt() string {
	if x != nil {
		return x.ArrivedAt
	}
	return ""
}

func (x *Dispatch) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Dispatch) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Dispatch) GetDetails() []*DispatchDetail {
	if x != nil {
		return x.Details
	}
	return nil
}

func (x *Dispatch) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *Dispatch) GetEndTime() string {
	if x != nil {
		return x.EndTime
	}
	return ""
}

func (x *Dispatch) GetShippedAt() string {
	if x != nil {
		return x.ShippedAt
	}
	return ""
}

func (x *Dispatch) GetShipper() *Shipper {
	if x != nil {
		return x.Shipper
	}
	return nil
}

func (x *Dispatch) GetLocation() *Fetch {
	if x != nil {
		return x.Location
	}
	return nil
}

type DispatchDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	ShipOrderId   int64    `protobuf:"varint,2,opt,name=ship_order_id,json=shipOrderId,proto3" json:"ship_order_id"`
	OrderDetailId int64    `protobuf:"varint,3,opt,name=order_detail_id,json=orderDetailId,proto3" json:"order_detail_id"`
	ItemId        int64    `protobuf:"varint,4,opt,name=item_id,json=itemId,proto3" json:"item_id"`
	SkuId         int64    `protobuf:"varint,5,opt,name=sku_id,json=skuId,proto3" json:"sku_id"`
	Num           int32    `protobuf:"varint,6,opt,name=num,proto3" json:"num"`
	CreatedAt     string   `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt     string   `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	Sku           *SkuInfo `protobuf:"bytes,9,opt,name=sku,proto3" json:"sku"`
}

func (x *DispatchDetail) Reset() {
	*x = DispatchDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dispatchService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DispatchDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DispatchDetail) ProtoMessage() {}

func (x *DispatchDetail) ProtoReflect() protoreflect.Message {
	mi := &file_dispatchService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DispatchDetail.ProtoReflect.Descriptor instead.
func (*DispatchDetail) Descriptor() ([]byte, []int) {
	return file_dispatchService_proto_rawDescGZIP(), []int{1}
}

func (x *DispatchDetail) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DispatchDetail) GetShipOrderId() int64 {
	if x != nil {
		return x.ShipOrderId
	}
	return 0
}

func (x *DispatchDetail) GetOrderDetailId() int64 {
	if x != nil {
		return x.OrderDetailId
	}
	return 0
}

func (x *DispatchDetail) GetItemId() int64 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *DispatchDetail) GetSkuId() int64 {
	if x != nil {
		return x.SkuId
	}
	return 0
}

func (x *DispatchDetail) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *DispatchDetail) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *DispatchDetail) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *DispatchDetail) GetSku() *SkuInfo {
	if x != nil {
		return x.Sku
	}
	return nil
}

type DispatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged        int32   `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize     int32   `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Sorting      string  `protobuf:"bytes,3,opt,name=sorting,proto3" json:"sorting"`
	Keywords     string  `protobuf:"bytes,4,opt,name=keywords,proto3" json:"keywords"`
	Id           int64   `protobuf:"varint,5,opt,name=id,proto3" json:"id"`
	Type         string  `protobuf:"bytes,7,opt,name=type,proto3" json:"type"`
	Method       string  `protobuf:"bytes,8,opt,name=method,proto3" json:"method"`
	OrderId      int64   `protobuf:"varint,9,opt,name=order_id,json=orderId,proto3" json:"order_id"`
	DeliveryType string  `protobuf:"bytes,10,opt,name=delivery_type,json=deliveryType,proto3" json:"delivery_type"`
	Ids          []int64 `protobuf:"varint,6,rep,packed,name=ids,proto3" json:"ids"`
}

func (x *DispatchRequest) Reset() {
	*x = DispatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dispatchService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DispatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DispatchRequest) ProtoMessage() {}

func (x *DispatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dispatchService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DispatchRequest.ProtoReflect.Descriptor instead.
func (*DispatchRequest) Descriptor() ([]byte, []int) {
	return file_dispatchService_proto_rawDescGZIP(), []int{2}
}

func (x *DispatchRequest) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *DispatchRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *DispatchRequest) GetSorting() string {
	if x != nil {
		return x.Sorting
	}
	return ""
}

func (x *DispatchRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *DispatchRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DispatchRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *DispatchRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *DispatchRequest) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *DispatchRequest) GetDeliveryType() string {
	if x != nil {
		return x.DeliveryType
	}
	return ""
}

func (x *DispatchRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type DispatchData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Dispatch     `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*Dispatch   `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Info   *common.Info  `protobuf:"bytes,4,opt,name=info,proto3" json:"info"`
}

func (x *DispatchData) Reset() {
	*x = DispatchData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dispatchService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DispatchData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DispatchData) ProtoMessage() {}

func (x *DispatchData) ProtoReflect() protoreflect.Message {
	mi := &file_dispatchService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DispatchData.ProtoReflect.Descriptor instead.
func (*DispatchData) Descriptor() ([]byte, []int) {
	return file_dispatchService_proto_rawDescGZIP(), []int{3}
}

func (x *DispatchData) GetEntity() *Dispatch {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *DispatchData) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *DispatchData) GetItems() []*Dispatch {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *DispatchData) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type DispatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *DispatchData `protobuf:"bytes,1,opt,name=data,proto3" json:"data"`
	Error *common.Error `protobuf:"bytes,2,opt,name=error,proto3" json:"error"`
}

func (x *DispatchResponse) Reset() {
	*x = DispatchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dispatchService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DispatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DispatchResponse) ProtoMessage() {}

func (x *DispatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dispatchService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DispatchResponse.ProtoReflect.Descriptor instead.
func (*DispatchResponse) Descriptor() ([]byte, []int) {
	return file_dispatchService_proto_rawDescGZIP(), []int{4}
}

func (x *DispatchResponse) GetData() *DispatchData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *DispatchResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_dispatchService_proto protoreflect.FileDescriptor

var file_dispatchService_proto_rawDesc = []byte{
	0x0a, 0x15, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x73, 0x6b, 0x75, 0x49,
	0x6e, 0x66, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x12, 0x66, 0x65, 0x74, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbf, 0x09, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63,
	0x68, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x23, 0x0a, 0x0d, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e,
	0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x72, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x66, 0x72, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x65, 0x64,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65,
	0x72, 0x79, 0x12, 0x26, 0x0a, 0x0f, 0x69, 0x73, 0x5f, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x73, 0x75,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x69, 0x73, 0x41,
	0x75, 0x74, 0x6f, 0x53, 0x75, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x68,
	0x69, 0x70, 0x70, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x6f, 0x67,
	0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x6e, 0x6f, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x4e, 0x6f, 0x12, 0x1d, 0x0a, 0x0a,
	0x66, 0x65, 0x74, 0x63, 0x68, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x66, 0x65, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08,
	0x66, 0x65, 0x74, 0x63, 0x68, 0x5f, 0x61, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x66, 0x65, 0x74, 0x63, 0x68, 0x41, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x5f, 0x61, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x41, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x28, 0x0a, 0x10, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x72, 0x65, 0x61,
	0x5f, 0x69, 0x64, 0x18, 0x13, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x72, 0x65, 0x63, 0x65, 0x69,
	0x76, 0x65, 0x72, 0x41, 0x72, 0x65, 0x61, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x63,
	0x65, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x12, 0x21,
	0x0a, 0x0c, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x7a, 0x69, 0x70, 0x18, 0x15,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x5a, 0x69,
	0x70, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x74, 0x65,
	0x6c, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x72, 0x54, 0x65, 0x6c, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72,
	0x5f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x25, 0x0a,
	0x0e, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x19, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f,
	0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x1b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65,
	0x6d, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x72, 0x72, 0x69, 0x76, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x1d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x72, 0x72, 0x69, 0x76, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x1f,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x32, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x20, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x69, 0x73, 0x70,
	0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x21, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x22,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x23, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x64, 0x41, 0x74, 0x12, 0x2b, 0x0a, 0x07,
	0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x18, 0x24, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72,
	0x52, 0x07, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x08, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x25, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x91, 0x02, 0x0a, 0x0e, 0x44, 0x69, 0x73, 0x70, 0x61,
	0x74, 0x63, 0x68, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x73, 0x68, 0x69,
	0x70, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0b, 0x73, 0x68, 0x69, 0x70, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x26, 0x0a,
	0x0f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x15,
	0x0a, 0x06, 0x73, 0x6b, 0x75, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x73, 0x6b, 0x75, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x23, 0x0a, 0x03, 0x73, 0x6b, 0x75, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x6b,
	0x75, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x03, 0x73, 0x6b, 0x75, 0x22, 0x88, 0x02, 0x0a, 0x0f, 0x44,
	0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70,
	0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x6b,
	0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b,
	0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x23,
	0x0a, 0x0d, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x03,
	0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0xab, 0x01, 0x0a, 0x0c, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74,
	0x63, 0x68, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2a, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72,
	0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69,
	0x6e, 0x66, 0x6f, 0x22, 0x63, 0x0a, 0x10, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x8a, 0x02, 0x0a, 0x0f, 0x44, 0x69, 0x73,
	0x70, 0x61, 0x74, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x06,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x1a, 0x1a, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x69,
	0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x1a, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x19, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x19, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dispatchService_proto_rawDescOnce sync.Once
	file_dispatchService_proto_rawDescData = file_dispatchService_proto_rawDesc
)

func file_dispatchService_proto_rawDescGZIP() []byte {
	file_dispatchService_proto_rawDescOnce.Do(func() {
		file_dispatchService_proto_rawDescData = protoimpl.X.CompressGZIP(file_dispatchService_proto_rawDescData)
	})
	return file_dispatchService_proto_rawDescData
}

var file_dispatchService_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_dispatchService_proto_goTypes = []interface{}{
	(*Dispatch)(nil),         // 0: services.Dispatch
	(*DispatchDetail)(nil),   // 1: services.DispatchDetail
	(*DispatchRequest)(nil),  // 2: services.DispatchRequest
	(*DispatchData)(nil),     // 3: services.DispatchData
	(*DispatchResponse)(nil), // 4: services.DispatchResponse
	(*Shipper)(nil),          // 5: services.Shipper
	(*Fetch)(nil),            // 6: services.Fetch
	(*SkuInfo)(nil),          // 7: services.SkuInfo
	(*common.Pager)(nil),     // 8: common.Pager
	(*common.Info)(nil),      // 9: common.Info
	(*common.Error)(nil),     // 10: common.Error
}
var file_dispatchService_proto_depIdxs = []int32{
	1,  // 0: services.Dispatch.details:type_name -> services.DispatchDetail
	5,  // 1: services.Dispatch.shipper:type_name -> services.Shipper
	6,  // 2: services.Dispatch.location:type_name -> services.Fetch
	7,  // 3: services.DispatchDetail.sku:type_name -> services.SkuInfo
	0,  // 4: services.DispatchData.entity:type_name -> services.Dispatch
	8,  // 5: services.DispatchData.pager:type_name -> common.Pager
	0,  // 6: services.DispatchData.items:type_name -> services.Dispatch
	9,  // 7: services.DispatchData.info:type_name -> common.Info
	3,  // 8: services.DispatchResponse.data:type_name -> services.DispatchData
	10, // 9: services.DispatchResponse.error:type_name -> common.Error
	0,  // 10: services.DispatchService.Create:input_type -> services.Dispatch
	0,  // 11: services.DispatchService.Update:input_type -> services.Dispatch
	2,  // 12: services.DispatchService.Get:input_type -> services.DispatchRequest
	2,  // 13: services.DispatchService.List:input_type -> services.DispatchRequest
	4,  // 14: services.DispatchService.Create:output_type -> services.DispatchResponse
	4,  // 15: services.DispatchService.Update:output_type -> services.DispatchResponse
	4,  // 16: services.DispatchService.Get:output_type -> services.DispatchResponse
	4,  // 17: services.DispatchService.List:output_type -> services.DispatchResponse
	14, // [14:18] is the sub-list for method output_type
	10, // [10:14] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_dispatchService_proto_init() }
func file_dispatchService_proto_init() {
	if File_dispatchService_proto != nil {
		return
	}
	file_shipperService_proto_init()
	file_skuInfoService_proto_init()
	file_fetchService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_dispatchService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dispatch); i {
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
		file_dispatchService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DispatchDetail); i {
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
		file_dispatchService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DispatchRequest); i {
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
		file_dispatchService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DispatchData); i {
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
		file_dispatchService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DispatchResponse); i {
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
			RawDescriptor: file_dispatchService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dispatchService_proto_goTypes,
		DependencyIndexes: file_dispatchService_proto_depIdxs,
		MessageInfos:      file_dispatchService_proto_msgTypes,
	}.Build()
	File_dispatchService_proto = out.File
	file_dispatchService_proto_rawDesc = nil
	file_dispatchService_proto_goTypes = nil
	file_dispatchService_proto_depIdxs = nil
}
