// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: orderAddressService.proto

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

type OrderAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Name         string `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`
	AreaId       int64  `protobuf:"varint,4,opt,name=area_id,json=areaId,proto3" json:"area_id"`
	Addr         string `protobuf:"bytes,5,opt,name=addr,proto3" json:"addr"`
	Zip          string `protobuf:"bytes,6,opt,name=zip,proto3" json:"zip"`
	Phone        string `protobuf:"bytes,7,opt,name=phone,proto3" json:"phone"`
	Mobile       string `protobuf:"bytes,8,opt,name=mobile,proto3" json:"mobile"`
	Email        string `protobuf:"bytes,9,opt,name=email,proto3" json:"email"`
	DeliveryTime string `protobuf:"bytes,10,opt,name=delivery_time,json=deliveryTime,proto3" json:"delivery_time"`
	Confirmed    bool   `protobuf:"varint,11,opt,name=confirmed,proto3" json:"confirmed"`
	OrderId      int64  `protobuf:"varint,12,opt,name=order_id,json=orderId,proto3" json:"order_id"`
	AddressId    int64  `protobuf:"varint,13,opt,name=address_id,json=addressId,proto3" json:"address_id"`
	CreatedAt    string `protobuf:"bytes,14,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt    string `protobuf:"bytes,15,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	// @inject_tag: gorm:"-"
	Area *AreaInfo `protobuf:"bytes,16,opt,name=area,proto3" json:"area" gorm:"-"`
	Lng  string    `protobuf:"bytes,17,opt,name=lng,proto3" json:"lng"`
	Lat  string    `protobuf:"bytes,18,opt,name=lat,proto3" json:"lat"`
}

func (x *OrderAddress) Reset() {
	*x = OrderAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderAddressService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderAddress) ProtoMessage() {}

func (x *OrderAddress) ProtoReflect() protoreflect.Message {
	mi := &file_orderAddressService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderAddress.ProtoReflect.Descriptor instead.
func (*OrderAddress) Descriptor() ([]byte, []int) {
	return file_orderAddressService_proto_rawDescGZIP(), []int{0}
}

func (x *OrderAddress) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OrderAddress) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OrderAddress) GetAreaId() int64 {
	if x != nil {
		return x.AreaId
	}
	return 0
}

func (x *OrderAddress) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *OrderAddress) GetZip() string {
	if x != nil {
		return x.Zip
	}
	return ""
}

func (x *OrderAddress) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *OrderAddress) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *OrderAddress) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *OrderAddress) GetDeliveryTime() string {
	if x != nil {
		return x.DeliveryTime
	}
	return ""
}

func (x *OrderAddress) GetConfirmed() bool {
	if x != nil {
		return x.Confirmed
	}
	return false
}

func (x *OrderAddress) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *OrderAddress) GetAddressId() int64 {
	if x != nil {
		return x.AddressId
	}
	return 0
}

func (x *OrderAddress) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *OrderAddress) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *OrderAddress) GetArea() *AreaInfo {
	if x != nil {
		return x.Area
	}
	return nil
}

func (x *OrderAddress) GetLng() string {
	if x != nil {
		return x.Lng
	}
	return ""
}

func (x *OrderAddress) GetLat() string {
	if x != nil {
		return x.Lat
	}
	return ""
}

type OrderAddressResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *OrderAddress   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager   `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*OrderAddress `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Error  *common.Error   `protobuf:"bytes,4,opt,name=error,proto3" json:"error"`
	Info   *common.Info    `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
}

func (x *OrderAddressResponse) Reset() {
	*x = OrderAddressResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderAddressService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderAddressResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderAddressResponse) ProtoMessage() {}

func (x *OrderAddressResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orderAddressService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderAddressResponse.ProtoReflect.Descriptor instead.
func (*OrderAddressResponse) Descriptor() ([]byte, []int) {
	return file_orderAddressService_proto_rawDescGZIP(), []int{1}
}

func (x *OrderAddressResponse) GetEntity() *OrderAddress {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *OrderAddressResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *OrderAddressResponse) GetItems() []*OrderAddress {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *OrderAddressResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *OrderAddressResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_orderAddressService_proto protoreflect.FileDescriptor

var file_orderAddressService_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x61, 0x72, 0x65, 0x61, 0x49, 0x6e,
	0x66, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xbc, 0x03, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x72, 0x65, 0x61, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64,
	0x72, 0x12, 0x10, 0x0a, 0x03, 0x7a, 0x69, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x7a, 0x69, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62,
	0x69, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x26, 0x0a, 0x04, 0x61, 0x72, 0x65, 0x61, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x72, 0x65, 0x61,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x61, 0x72, 0x65, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6e,
	0x67, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x6e, 0x67, 0x12, 0x10, 0x0a, 0x03,
	0x6c, 0x61, 0x74, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x22, 0xe0,
	0x01, 0x0a, 0x14, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52,
	0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12,
	0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66,
	0x6f, 0x32, 0xe1, 0x02, 0x0a, 0x13, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x06, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x1a, 0x1e, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a,
	0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x1a,
	0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x42, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x1a, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x16, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x1a, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x12, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x57, 0x68,
	0x65, 0x72, 0x65, 0x1a, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_orderAddressService_proto_rawDescOnce sync.Once
	file_orderAddressService_proto_rawDescData = file_orderAddressService_proto_rawDesc
)

func file_orderAddressService_proto_rawDescGZIP() []byte {
	file_orderAddressService_proto_rawDescOnce.Do(func() {
		file_orderAddressService_proto_rawDescData = protoimpl.X.CompressGZIP(file_orderAddressService_proto_rawDescData)
	})
	return file_orderAddressService_proto_rawDescData
}

var file_orderAddressService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_orderAddressService_proto_goTypes = []interface{}{
	(*OrderAddress)(nil),         // 0: services.OrderAddress
	(*OrderAddressResponse)(nil), // 1: services.OrderAddressResponse
	(*AreaInfo)(nil),             // 2: services.AreaInfo
	(*common.Pager)(nil),         // 3: common.Pager
	(*common.Error)(nil),         // 4: common.Error
	(*common.Info)(nil),          // 5: common.Info
	(*common.BaseWhere)(nil),     // 6: common.BaseWhere
}
var file_orderAddressService_proto_depIdxs = []int32{
	2,  // 0: services.OrderAddress.area:type_name -> services.AreaInfo
	0,  // 1: services.OrderAddressResponse.entity:type_name -> services.OrderAddress
	3,  // 2: services.OrderAddressResponse.pager:type_name -> common.Pager
	0,  // 3: services.OrderAddressResponse.items:type_name -> services.OrderAddress
	4,  // 4: services.OrderAddressResponse.error:type_name -> common.Error
	5,  // 5: services.OrderAddressResponse.info:type_name -> common.Info
	0,  // 6: services.OrderAddressService.Create:input_type -> services.OrderAddress
	0,  // 7: services.OrderAddressService.Update:input_type -> services.OrderAddress
	0,  // 8: services.OrderAddressService.Delete:input_type -> services.OrderAddress
	0,  // 9: services.OrderAddressService.Get:input_type -> services.OrderAddress
	6,  // 10: services.OrderAddressService.Search:input_type -> common.BaseWhere
	1,  // 11: services.OrderAddressService.Create:output_type -> services.OrderAddressResponse
	1,  // 12: services.OrderAddressService.Update:output_type -> services.OrderAddressResponse
	1,  // 13: services.OrderAddressService.Delete:output_type -> services.OrderAddressResponse
	1,  // 14: services.OrderAddressService.Get:output_type -> services.OrderAddressResponse
	1,  // 15: services.OrderAddressService.Search:output_type -> services.OrderAddressResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_orderAddressService_proto_init() }
func file_orderAddressService_proto_init() {
	if File_orderAddressService_proto != nil {
		return
	}
	file_areaInfoService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_orderAddressService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderAddress); i {
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
		file_orderAddressService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderAddressResponse); i {
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
			RawDescriptor: file_orderAddressService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_orderAddressService_proto_goTypes,
		DependencyIndexes: file_orderAddressService_proto_depIdxs,
		MessageInfos:      file_orderAddressService_proto_msgTypes,
	}.Build()
	File_orderAddressService_proto = out.File
	file_orderAddressService_proto_rawDesc = nil
	file_orderAddressService_proto_goTypes = nil
	file_orderAddressService_proto_depIdxs = nil
}
