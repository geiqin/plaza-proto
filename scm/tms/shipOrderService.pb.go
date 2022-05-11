// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: shipOrderService.proto

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

type ShipOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int64              `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Code           string             `protobuf:"bytes,2,opt,name=code,proto3" json:"code"`
	OrderId        int64              `protobuf:"varint,3,opt,name=order_id,json=orderId,proto3" json:"order_id"`
	Method         string             `protobuf:"bytes,4,opt,name=method,proto3" json:"method"`
	Type           string             `protobuf:"bytes,5,opt,name=type,proto3" json:"type"`
	Freight        float32            `protobuf:"fixed32,6,opt,name=freight,proto3" json:"freight"`
	Protected      bool               `protobuf:"varint,7,opt,name=protected,proto3" json:"protected"`
	IsDelivery     bool               `protobuf:"varint,8,opt,name=is_delivery,json=isDelivery,proto3" json:"is_delivery"`
	ShipperId      int32              `protobuf:"varint,9,opt,name=shipper_id,json=shipperId,proto3" json:"shipper_id"`
	LogisticsNo    string             `protobuf:"bytes,10,opt,name=logistics_no,json=logisticsNo,proto3" json:"logistics_no"`
	FetchCode      string             `protobuf:"bytes,11,opt,name=fetch_code,json=fetchCode,proto3" json:"fetch_code"`
	LocationId     int64              `protobuf:"varint,12,opt,name=location_id,json=locationId,proto3" json:"location_id"`
	FetchAt        string             `protobuf:"bytes,13,opt,name=fetch_at,json=fetchAt,proto3" json:"fetch_at"`
	DeliveryAt     string             `protobuf:"bytes,14,opt,name=delivery_at,json=deliveryAt,proto3" json:"delivery_at"`
	DeliveryType   int32              `protobuf:"varint,33,opt,name=delivery_type,json=deliveryType,proto3" json:"delivery_type"`
	CustomerId     int64              `protobuf:"varint,15,opt,name=customer_id,json=customerId,proto3" json:"customer_id"`
	ReceiverName   string             `protobuf:"bytes,16,opt,name=receiver_name,json=receiverName,proto3" json:"receiver_name"`
	ReceiverAreaId int64              `protobuf:"varint,17,opt,name=receiver_area_id,json=receiverAreaId,proto3" json:"receiver_area_id"`
	ReceiverAddr   string             `protobuf:"bytes,18,opt,name=receiver_addr,json=receiverAddr,proto3" json:"receiver_addr"`
	ReceiverZip    string             `protobuf:"bytes,19,opt,name=receiver_zip,json=receiverZip,proto3" json:"receiver_zip"`
	ReceiverTel    string             `protobuf:"bytes,20,opt,name=receiver_tel,json=receiverTel,proto3" json:"receiver_tel"`
	ReceiverMobile string             `protobuf:"bytes,21,opt,name=receiver_mobile,json=receiverMobile,proto3" json:"receiver_mobile"`
	ReceiverEmail  string             `protobuf:"bytes,22,opt,name=receiver_email,json=receiverEmail,proto3" json:"receiver_email"`
	OpId           int64              `protobuf:"varint,23,opt,name=op_id,json=opId,proto3" json:"op_id"`
	Status         string             `protobuf:"bytes,24,opt,name=status,proto3" json:"status"`
	Memo           string             `protobuf:"bytes,25,opt,name=memo,proto3" json:"memo"`
	ArrivedAt      string             `protobuf:"bytes,26,opt,name=arrived_at,json=arrivedAt,proto3" json:"arrived_at"`
	CreatedAt      string             `protobuf:"bytes,27,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt      string             `protobuf:"bytes,28,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	Details        []*ShipOrderDetail `protobuf:"bytes,29,rep,name=details,proto3" json:"details"`
	StartTime      string             `protobuf:"bytes,30,opt,name=start_time,json=startTime,proto3" json:"start_time"`
	EndTime        string             `protobuf:"bytes,31,opt,name=end_time,json=endTime,proto3" json:"end_time"`
	ShippedAt      string             `protobuf:"bytes,32,opt,name=shipped_at,json=shippedAt,proto3" json:"shipped_at"`
	Shipper        *Shipper           `protobuf:"bytes,34,opt,name=shipper,proto3" json:"shipper"`
	Location       *Fetch             `protobuf:"bytes,35,opt,name=location,proto3" json:"location"`
}

func (x *ShipOrder) Reset() {
	*x = ShipOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shipOrderService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShipOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShipOrder) ProtoMessage() {}

func (x *ShipOrder) ProtoReflect() protoreflect.Message {
	mi := &file_shipOrderService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShipOrder.ProtoReflect.Descriptor instead.
func (*ShipOrder) Descriptor() ([]byte, []int) {
	return file_shipOrderService_proto_rawDescGZIP(), []int{0}
}

func (x *ShipOrder) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ShipOrder) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *ShipOrder) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *ShipOrder) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *ShipOrder) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ShipOrder) GetFreight() float32 {
	if x != nil {
		return x.Freight
	}
	return 0
}

func (x *ShipOrder) GetProtected() bool {
	if x != nil {
		return x.Protected
	}
	return false
}

func (x *ShipOrder) GetIsDelivery() bool {
	if x != nil {
		return x.IsDelivery
	}
	return false
}

func (x *ShipOrder) GetShipperId() int32 {
	if x != nil {
		return x.ShipperId
	}
	return 0
}

func (x *ShipOrder) GetLogisticsNo() string {
	if x != nil {
		return x.LogisticsNo
	}
	return ""
}

func (x *ShipOrder) GetFetchCode() string {
	if x != nil {
		return x.FetchCode
	}
	return ""
}

func (x *ShipOrder) GetLocationId() int64 {
	if x != nil {
		return x.LocationId
	}
	return 0
}

func (x *ShipOrder) GetFetchAt() string {
	if x != nil {
		return x.FetchAt
	}
	return ""
}

func (x *ShipOrder) GetDeliveryAt() string {
	if x != nil {
		return x.DeliveryAt
	}
	return ""
}

func (x *ShipOrder) GetDeliveryType() int32 {
	if x != nil {
		return x.DeliveryType
	}
	return 0
}

func (x *ShipOrder) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *ShipOrder) GetReceiverName() string {
	if x != nil {
		return x.ReceiverName
	}
	return ""
}

func (x *ShipOrder) GetReceiverAreaId() int64 {
	if x != nil {
		return x.ReceiverAreaId
	}
	return 0
}

func (x *ShipOrder) GetReceiverAddr() string {
	if x != nil {
		return x.ReceiverAddr
	}
	return ""
}

func (x *ShipOrder) GetReceiverZip() string {
	if x != nil {
		return x.ReceiverZip
	}
	return ""
}

func (x *ShipOrder) GetReceiverTel() string {
	if x != nil {
		return x.ReceiverTel
	}
	return ""
}

func (x *ShipOrder) GetReceiverMobile() string {
	if x != nil {
		return x.ReceiverMobile
	}
	return ""
}

func (x *ShipOrder) GetReceiverEmail() string {
	if x != nil {
		return x.ReceiverEmail
	}
	return ""
}

func (x *ShipOrder) GetOpId() int64 {
	if x != nil {
		return x.OpId
	}
	return 0
}

func (x *ShipOrder) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ShipOrder) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *ShipOrder) GetArrivedAt() string {
	if x != nil {
		return x.ArrivedAt
	}
	return ""
}

func (x *ShipOrder) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *ShipOrder) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *ShipOrder) GetDetails() []*ShipOrderDetail {
	if x != nil {
		return x.Details
	}
	return nil
}

func (x *ShipOrder) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *ShipOrder) GetEndTime() string {
	if x != nil {
		return x.EndTime
	}
	return ""
}

func (x *ShipOrder) GetShippedAt() string {
	if x != nil {
		return x.ShippedAt
	}
	return ""
}

func (x *ShipOrder) GetShipper() *Shipper {
	if x != nil {
		return x.Shipper
	}
	return nil
}

func (x *ShipOrder) GetLocation() *Fetch {
	if x != nil {
		return x.Location
	}
	return nil
}

type ShipOrderDetail struct {
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

func (x *ShipOrderDetail) Reset() {
	*x = ShipOrderDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shipOrderService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShipOrderDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShipOrderDetail) ProtoMessage() {}

func (x *ShipOrderDetail) ProtoReflect() protoreflect.Message {
	mi := &file_shipOrderService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShipOrderDetail.ProtoReflect.Descriptor instead.
func (*ShipOrderDetail) Descriptor() ([]byte, []int) {
	return file_shipOrderService_proto_rawDescGZIP(), []int{1}
}

func (x *ShipOrderDetail) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ShipOrderDetail) GetShipOrderId() int64 {
	if x != nil {
		return x.ShipOrderId
	}
	return 0
}

func (x *ShipOrderDetail) GetOrderDetailId() int64 {
	if x != nil {
		return x.OrderDetailId
	}
	return 0
}

func (x *ShipOrderDetail) GetItemId() int64 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *ShipOrderDetail) GetSkuId() int64 {
	if x != nil {
		return x.SkuId
	}
	return 0
}

func (x *ShipOrderDetail) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *ShipOrderDetail) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *ShipOrderDetail) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *ShipOrderDetail) GetSku() *SkuInfo {
	if x != nil {
		return x.Sku
	}
	return nil
}

type ShipOrderRequest struct {
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
	DeliveryType int32   `protobuf:"varint,10,opt,name=delivery_type,json=deliveryType,proto3" json:"delivery_type"`
	Ids          []int64 `protobuf:"varint,6,rep,packed,name=ids,proto3" json:"ids"`
}

func (x *ShipOrderRequest) Reset() {
	*x = ShipOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shipOrderService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShipOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShipOrderRequest) ProtoMessage() {}

func (x *ShipOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shipOrderService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShipOrderRequest.ProtoReflect.Descriptor instead.
func (*ShipOrderRequest) Descriptor() ([]byte, []int) {
	return file_shipOrderService_proto_rawDescGZIP(), []int{2}
}

func (x *ShipOrderRequest) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *ShipOrderRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ShipOrderRequest) GetSorting() string {
	if x != nil {
		return x.Sorting
	}
	return ""
}

func (x *ShipOrderRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *ShipOrderRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ShipOrderRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ShipOrderRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *ShipOrderRequest) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *ShipOrderRequest) GetDeliveryType() int32 {
	if x != nil {
		return x.DeliveryType
	}
	return 0
}

func (x *ShipOrderRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type ShipOrderData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *ShipOrder    `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*ShipOrder  `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Info   *common.Info  `protobuf:"bytes,4,opt,name=info,proto3" json:"info"`
}

func (x *ShipOrderData) Reset() {
	*x = ShipOrderData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shipOrderService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShipOrderData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShipOrderData) ProtoMessage() {}

func (x *ShipOrderData) ProtoReflect() protoreflect.Message {
	mi := &file_shipOrderService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShipOrderData.ProtoReflect.Descriptor instead.
func (*ShipOrderData) Descriptor() ([]byte, []int) {
	return file_shipOrderService_proto_rawDescGZIP(), []int{3}
}

func (x *ShipOrderData) GetEntity() *ShipOrder {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *ShipOrderData) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *ShipOrderData) GetItems() []*ShipOrder {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ShipOrderData) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type ShipOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *ShipOrderData `protobuf:"bytes,1,opt,name=data,proto3" json:"data"`
	Error *common.Error  `protobuf:"bytes,2,opt,name=error,proto3" json:"error"`
}

func (x *ShipOrderResponse) Reset() {
	*x = ShipOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shipOrderService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShipOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShipOrderResponse) ProtoMessage() {}

func (x *ShipOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shipOrderService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShipOrderResponse.ProtoReflect.Descriptor instead.
func (*ShipOrderResponse) Descriptor() ([]byte, []int) {
	return file_shipOrderService_proto_rawDescGZIP(), []int{4}
}

func (x *ShipOrderResponse) GetData() *ShipOrderData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ShipOrderResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_shipOrderService_proto protoreflect.FileDescriptor

var file_shipOrderService_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x68, 0x69, 0x70, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x73, 0x6b, 0x75,
	0x49, 0x6e, 0x66, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x12, 0x66, 0x65, 0x74, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe3, 0x08, 0x0a, 0x09, 0x53, 0x68, 0x69, 0x70, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x66, 0x72, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x07, 0x66, 0x72, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x74,
	0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x70, 0x72, 0x6f,
	0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x44,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x68, 0x69, 0x70, 0x70,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x68, 0x69,
	0x70, 0x70, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74,
	0x69, 0x63, 0x73, 0x5f, 0x6e, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x6f,
	0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x4e, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x65, 0x74,
	0x63, 0x68, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66,
	0x65, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x65, 0x74,
	0x63, 0x68, 0x5f, 0x61, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x65, 0x74,
	0x63, 0x68, 0x41, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79,
	0x5f, 0x61, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x41, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x21, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x64, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x72,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x10, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x28, 0x0a, 0x10, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x72, 0x65,
	0x61, 0x5f, 0x69, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x72, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x72, 0x41, 0x72, 0x65, 0x61, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x12, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x12,
	0x21, 0x0a, 0x0c, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x7a, 0x69, 0x70, 0x18,
	0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x5a,
	0x69, 0x70, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x74,
	0x65, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x72, 0x54, 0x65, 0x6c, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x72, 0x5f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x25,
	0x0a, 0x0e, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x13, 0x0a, 0x05, 0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x17,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x6f, 0x70, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x19, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x72, 0x72, 0x69, 0x76, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x72, 0x72, 0x69,
	0x76, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x33, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x1d,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x53, 0x68, 0x69, 0x70, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52,
	0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x20, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x2b, 0x0a, 0x07, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x18, 0x22, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x68,
	0x69, 0x70, 0x70, 0x65, 0x72, 0x52, 0x07, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x12, 0x2b,
	0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x23, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x65, 0x74, 0x63,
	0x68, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x92, 0x02, 0x0a, 0x0f,
	0x53, 0x68, 0x69, 0x70, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x22, 0x0a, 0x0d, 0x73, 0x68, 0x69, 0x70, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x73, 0x68, 0x69, 0x70, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x69,
	0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x69, 0x74,
	0x65, 0x6d, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x6b, 0x75, 0x5f, 0x69, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x6b, 0x75, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6e,
	0x75, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x23, 0x0a, 0x03, 0x73,
	0x6b, 0x75, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x53, 0x6b, 0x75, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x03, 0x73, 0x6b, 0x75,
	0x22, 0x89, 0x02, 0x0a, 0x10, 0x53, 0x68, 0x69, 0x70, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74,
	0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69,
	0x6e, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x64, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64,
	0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0xae, 0x01, 0x0a,
	0x0d, 0x53, 0x68, 0x69, 0x70, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2b,
	0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70,
	0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72,
	0x12, 0x29, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x20, 0x0a, 0x04, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x65, 0x0a,
	0x11, 0x53, 0x68, 0x69, 0x70, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x68, 0x69, 0x70,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x32, 0x93, 0x02, 0x0a, 0x10, 0x53, 0x68, 0x69, 0x70, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x06, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53,
	0x68, 0x69, 0x70, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x68, 0x69,
	0x70, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1a, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shipOrderService_proto_rawDescOnce sync.Once
	file_shipOrderService_proto_rawDescData = file_shipOrderService_proto_rawDesc
)

func file_shipOrderService_proto_rawDescGZIP() []byte {
	file_shipOrderService_proto_rawDescOnce.Do(func() {
		file_shipOrderService_proto_rawDescData = protoimpl.X.CompressGZIP(file_shipOrderService_proto_rawDescData)
	})
	return file_shipOrderService_proto_rawDescData
}

var file_shipOrderService_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_shipOrderService_proto_goTypes = []interface{}{
	(*ShipOrder)(nil),         // 0: services.ShipOrder
	(*ShipOrderDetail)(nil),   // 1: services.ShipOrderDetail
	(*ShipOrderRequest)(nil),  // 2: services.ShipOrderRequest
	(*ShipOrderData)(nil),     // 3: services.ShipOrderData
	(*ShipOrderResponse)(nil), // 4: services.ShipOrderResponse
	(*Shipper)(nil),           // 5: services.Shipper
	(*Fetch)(nil),             // 6: services.Fetch
	(*SkuInfo)(nil),           // 7: services.SkuInfo
	(*common.Pager)(nil),      // 8: common.Pager
	(*common.Info)(nil),       // 9: common.Info
	(*common.Error)(nil),      // 10: common.Error
}
var file_shipOrderService_proto_depIdxs = []int32{
	1,  // 0: services.ShipOrder.details:type_name -> services.ShipOrderDetail
	5,  // 1: services.ShipOrder.shipper:type_name -> services.Shipper
	6,  // 2: services.ShipOrder.location:type_name -> services.Fetch
	7,  // 3: services.ShipOrderDetail.sku:type_name -> services.SkuInfo
	0,  // 4: services.ShipOrderData.entity:type_name -> services.ShipOrder
	8,  // 5: services.ShipOrderData.pager:type_name -> common.Pager
	0,  // 6: services.ShipOrderData.items:type_name -> services.ShipOrder
	9,  // 7: services.ShipOrderData.info:type_name -> common.Info
	3,  // 8: services.ShipOrderResponse.data:type_name -> services.ShipOrderData
	10, // 9: services.ShipOrderResponse.error:type_name -> common.Error
	0,  // 10: services.ShipOrderService.Create:input_type -> services.ShipOrder
	0,  // 11: services.ShipOrderService.Update:input_type -> services.ShipOrder
	2,  // 12: services.ShipOrderService.Get:input_type -> services.ShipOrderRequest
	2,  // 13: services.ShipOrderService.List:input_type -> services.ShipOrderRequest
	4,  // 14: services.ShipOrderService.Create:output_type -> services.ShipOrderResponse
	4,  // 15: services.ShipOrderService.Update:output_type -> services.ShipOrderResponse
	4,  // 16: services.ShipOrderService.Get:output_type -> services.ShipOrderResponse
	4,  // 17: services.ShipOrderService.List:output_type -> services.ShipOrderResponse
	14, // [14:18] is the sub-list for method output_type
	10, // [10:14] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_shipOrderService_proto_init() }
func file_shipOrderService_proto_init() {
	if File_shipOrderService_proto != nil {
		return
	}
	file_shipperService_proto_init()
	file_skuInfoService_proto_init()
	file_fetchService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_shipOrderService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShipOrder); i {
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
		file_shipOrderService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShipOrderDetail); i {
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
		file_shipOrderService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShipOrderRequest); i {
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
		file_shipOrderService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShipOrderData); i {
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
		file_shipOrderService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShipOrderResponse); i {
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
			RawDescriptor: file_shipOrderService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shipOrderService_proto_goTypes,
		DependencyIndexes: file_shipOrderService_proto_depIdxs,
		MessageInfos:      file_shipOrderService_proto_msgTypes,
	}.Build()
	File_shipOrderService_proto = out.File
	file_shipOrderService_proto_rawDesc = nil
	file_shipOrderService_proto_goTypes = nil
	file_shipOrderService_proto_depIdxs = nil
}
