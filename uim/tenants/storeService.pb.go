// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: storeService.proto

package services

import (
	common "github.com/geiqin/microkit/protobuf/common"
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

type Store struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string       `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	LogoId        int64        `protobuf:"varint,3,opt,name=logo_id,json=logoId,proto3" json:"logo_id,omitempty"`
	LogoUrl       string       `protobuf:"bytes,4,opt,name=logo_url,json=logoUrl,proto3" json:"logo_url,omitempty"`
	CreatorId     int64        `protobuf:"varint,6,opt,name=creator_id,json=creatorId,proto3" json:"creator_id,omitempty"`
	ApplicationId int32        `protobuf:"varint,7,opt,name=application_id,json=applicationId,proto3" json:"application_id,omitempty"`
	Version       string       `protobuf:"bytes,8,opt,name=version,proto3" json:"version,omitempty"`
	Industry      int64        `protobuf:"varint,9,opt,name=industry,proto3" json:"industry,omitempty"`
	AreaId        int64        `protobuf:"varint,10,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	Address       string       `protobuf:"bytes,11,opt,name=address,proto3" json:"address,omitempty"`
	Lng           string       `protobuf:"bytes,12,opt,name=lng,proto3" json:"lng,omitempty"`
	Lat           string       `protobuf:"bytes,13,opt,name=lat,proto3" json:"lat,omitempty"`
	CustomDomain  string       `protobuf:"bytes,14,opt,name=custom_domain,json=customDomain,proto3" json:"custom_domain,omitempty"`
	Status        string       `protobuf:"bytes,15,opt,name=status,proto3" json:"status,omitempty"`
	Enabled       bool         `protobuf:"varint,16,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Trusted       bool         `protobuf:"varint,17,opt,name=trusted,proto3" json:"trusted,omitempty"`
	StartedAt     string       `protobuf:"bytes,18,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	ExpiredAt     string       `protobuf:"bytes,19,opt,name=expired_at,json=expiredAt,proto3" json:"expired_at,omitempty"`
	CreatedAt     string       `protobuf:"bytes,20,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string       `protobuf:"bytes,21,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Application   *Application `protobuf:"bytes,22,opt,name=application,proto3" json:"application,omitempty"`
	Ids           []int64      `protobuf:"varint,23,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *Store) Reset() {
	*x = Store{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storeService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Store) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Store) ProtoMessage() {}

func (x *Store) ProtoReflect() protoreflect.Message {
	mi := &file_storeService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Store.ProtoReflect.Descriptor instead.
func (*Store) Descriptor() ([]byte, []int) {
	return file_storeService_proto_rawDescGZIP(), []int{0}
}

func (x *Store) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Store) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Store) GetLogoId() int64 {
	if x != nil {
		return x.LogoId
	}
	return 0
}

func (x *Store) GetLogoUrl() string {
	if x != nil {
		return x.LogoUrl
	}
	return ""
}

func (x *Store) GetCreatorId() int64 {
	if x != nil {
		return x.CreatorId
	}
	return 0
}

func (x *Store) GetApplicationId() int32 {
	if x != nil {
		return x.ApplicationId
	}
	return 0
}

func (x *Store) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Store) GetIndustry() int64 {
	if x != nil {
		return x.Industry
	}
	return 0
}

func (x *Store) GetAreaId() int64 {
	if x != nil {
		return x.AreaId
	}
	return 0
}

func (x *Store) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Store) GetLng() string {
	if x != nil {
		return x.Lng
	}
	return ""
}

func (x *Store) GetLat() string {
	if x != nil {
		return x.Lat
	}
	return ""
}

func (x *Store) GetCustomDomain() string {
	if x != nil {
		return x.CustomDomain
	}
	return ""
}

func (x *Store) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Store) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *Store) GetTrusted() bool {
	if x != nil {
		return x.Trusted
	}
	return false
}

func (x *Store) GetStartedAt() string {
	if x != nil {
		return x.StartedAt
	}
	return ""
}

func (x *Store) GetExpiredAt() string {
	if x != nil {
		return x.ExpiredAt
	}
	return ""
}

func (x *Store) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Store) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Store) GetApplication() *Application {
	if x != nil {
		return x.Application
	}
	return nil
}

func (x *Store) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type StoreSecret struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StoreKey    string `protobuf:"bytes,1,opt,name=store_key,json=storeKey,proto3" json:"store_key,omitempty"`
	StoreSecret string `protobuf:"bytes,2,opt,name=store_secret,json=storeSecret,proto3" json:"store_secret,omitempty"`
}

func (x *StoreSecret) Reset() {
	*x = StoreSecret{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storeService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreSecret) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreSecret) ProtoMessage() {}

func (x *StoreSecret) ProtoReflect() protoreflect.Message {
	mi := &file_storeService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreSecret.ProtoReflect.Descriptor instead.
func (*StoreSecret) Descriptor() ([]byte, []int) {
	return file_storeService_proto_rawDescGZIP(), []int{1}
}

func (x *StoreSecret) GetStoreKey() string {
	if x != nil {
		return x.StoreKey
	}
	return ""
}

func (x *StoreSecret) GetStoreSecret() string {
	if x != nil {
		return x.StoreSecret
	}
	return ""
}

//
type StoreResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Store        `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items  []*Store      `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error  *common.Error `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info   *common.Info  `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *StoreResponse) Reset() {
	*x = StoreResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storeService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreResponse) ProtoMessage() {}

func (x *StoreResponse) ProtoReflect() protoreflect.Message {
	mi := &file_storeService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreResponse.ProtoReflect.Descriptor instead.
func (*StoreResponse) Descriptor() ([]byte, []int) {
	return file_storeService_proto_rawDescGZIP(), []int{2}
}

func (x *StoreResponse) GetEntity() *Store {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *StoreResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *StoreResponse) GetItems() []*Store {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *StoreResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *StoreResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_storeService_proto protoreflect.FileDescriptor

var file_storeService_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x18, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xea, 0x04, 0x0a, 0x05,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x6f, 0x67,
	0x6f, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6c, 0x6f, 0x67, 0x6f,
	0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x6f, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f, 0x67, 0x6f, 0x55, 0x72, 0x6c, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x69, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x69, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x72, 0x65,
	0x61, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x72, 0x65, 0x61,
	0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x10, 0x0a, 0x03,
	0x6c, 0x6e, 0x67, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x6e, 0x67, 0x12, 0x10,
	0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x61, 0x74,
	0x12, 0x23, 0x0a, 0x0d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x72, 0x75, 0x73, 0x74,
	0x65, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x74, 0x72, 0x75, 0x73, 0x74, 0x65,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x13,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x15, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x37, 0x0a,
	0x0b, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x16, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x17, 0x20,
	0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x4d, 0x0a, 0x0b, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x4b, 0x65, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x22, 0xcb, 0x01, 0x0a, 0x0d, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72,
	0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x23,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0xea, 0x04, 0x0a, 0x0c, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x06,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x38, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x6f,
	0x12, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x0a,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x0f, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x1a, 0x17, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x0a, 0x53, 0x65, 0x74, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x34, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x1a, 0x17, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x07, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64,
	0x65, 0x12, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a,
	0x03, 0x47, 0x65, 0x74, 0x12, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x32, 0x0a, 0x04, 0x4c, 0x6f, 0x63, 0x6b, 0x12, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x06, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x0f,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x1a,
	0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x06, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x12, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x42, 0x61,
	0x73, 0x65, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_storeService_proto_rawDescOnce sync.Once
	file_storeService_proto_rawDescData = file_storeService_proto_rawDesc
)

func file_storeService_proto_rawDescGZIP() []byte {
	file_storeService_proto_rawDescOnce.Do(func() {
		file_storeService_proto_rawDescData = protoimpl.X.CompressGZIP(file_storeService_proto_rawDescData)
	})
	return file_storeService_proto_rawDescData
}

var file_storeService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_storeService_proto_goTypes = []interface{}{
	(*Store)(nil),            // 0: services.Store
	(*StoreSecret)(nil),      // 1: services.StoreSecret
	(*StoreResponse)(nil),    // 2: services.StoreResponse
	(*Application)(nil),      // 3: services.Application
	(*common.Pager)(nil),     // 4: common.Pager
	(*common.Error)(nil),     // 5: common.Error
	(*common.Info)(nil),      // 6: common.Info
	(*common.BaseWhere)(nil), // 7: common.BaseWhere
}
var file_storeService_proto_depIdxs = []int32{
	3,  // 0: services.Store.application:type_name -> services.Application
	0,  // 1: services.StoreResponse.entity:type_name -> services.Store
	4,  // 2: services.StoreResponse.pager:type_name -> common.Pager
	0,  // 3: services.StoreResponse.items:type_name -> services.Store
	5,  // 4: services.StoreResponse.error:type_name -> common.Error
	6,  // 5: services.StoreResponse.info:type_name -> common.Info
	0,  // 6: services.StoreService.Create:input_type -> services.Store
	0,  // 7: services.StoreService.Update:input_type -> services.Store
	0,  // 8: services.StoreService.UpdateLogo:input_type -> services.Store
	0,  // 9: services.StoreService.UpdateName:input_type -> services.Store
	0,  // 10: services.StoreService.SetAddress:input_type -> services.Store
	0,  // 11: services.StoreService.Delete:input_type -> services.Store
	0,  // 12: services.StoreService.Upgrade:input_type -> services.Store
	0,  // 13: services.StoreService.Get:input_type -> services.Store
	0,  // 14: services.StoreService.Lock:input_type -> services.Store
	0,  // 15: services.StoreService.Unlock:input_type -> services.Store
	7,  // 16: services.StoreService.Search:input_type -> common.BaseWhere
	2,  // 17: services.StoreService.Create:output_type -> services.StoreResponse
	2,  // 18: services.StoreService.Update:output_type -> services.StoreResponse
	2,  // 19: services.StoreService.UpdateLogo:output_type -> services.StoreResponse
	2,  // 20: services.StoreService.UpdateName:output_type -> services.StoreResponse
	2,  // 21: services.StoreService.SetAddress:output_type -> services.StoreResponse
	2,  // 22: services.StoreService.Delete:output_type -> services.StoreResponse
	2,  // 23: services.StoreService.Upgrade:output_type -> services.StoreResponse
	2,  // 24: services.StoreService.Get:output_type -> services.StoreResponse
	2,  // 25: services.StoreService.Lock:output_type -> services.StoreResponse
	2,  // 26: services.StoreService.Unlock:output_type -> services.StoreResponse
	2,  // 27: services.StoreService.Search:output_type -> services.StoreResponse
	17, // [17:28] is the sub-list for method output_type
	6,  // [6:17] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_storeService_proto_init() }
func file_storeService_proto_init() {
	if File_storeService_proto != nil {
		return
	}
	file_applicationService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_storeService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Store); i {
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
		file_storeService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreSecret); i {
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
		file_storeService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreResponse); i {
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
			RawDescriptor: file_storeService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_storeService_proto_goTypes,
		DependencyIndexes: file_storeService_proto_depIdxs,
		MessageInfos:      file_storeService_proto_msgTypes,
	}.Build()
	File_storeService_proto = out.File
	file_storeService_proto_rawDesc = nil
	file_storeService_proto_goTypes = nil
	file_storeService_proto_depIdxs = nil
}
