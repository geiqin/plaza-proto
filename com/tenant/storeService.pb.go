// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: storeService.proto

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

type Store struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64        `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Name         string       `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	Title        string       `protobuf:"bytes,3,opt,name=title,proto3" json:"title"`
	LogoUrl      string       `protobuf:"bytes,5,opt,name=logo_url,json=logoUrl,proto3" json:"logo_url"`
	VersionId    int32        `protobuf:"varint,6,opt,name=version_id,json=versionId,proto3" json:"version_id"`
	CompanyId    int64        `protobuf:"varint,7,opt,name=company_id,json=companyId,proto3" json:"company_id"`
	CompanyName  string       `protobuf:"bytes,8,opt,name=company_name,json=companyName,proto3" json:"company_name"`
	Industry     string       `protobuf:"bytes,9,opt,name=industry,proto3" json:"industry"`
	ProvinceId   int64        `protobuf:"varint,10,opt,name=province_id,json=provinceId,proto3" json:"province_id"`
	CityId       int64        `protobuf:"varint,11,opt,name=city_id,json=cityId,proto3" json:"city_id"`
	CountyId     int64        `protobuf:"varint,12,opt,name=county_id,json=countyId,proto3" json:"county_id"`
	Address      string       `protobuf:"bytes,13,opt,name=address,proto3" json:"address"`
	Lng          float32      `protobuf:"fixed32,14,opt,name=lng,proto3" json:"lng"`
	Lat          float32      `protobuf:"fixed32,15,opt,name=lat,proto3" json:"lat"`
	Region       string       `protobuf:"bytes,16,opt,name=region,proto3" json:"region"`
	Config       *StoreConfig `protobuf:"bytes,17,opt,name=config,proto3" json:"config"`
	IsTrusted    string       `protobuf:"bytes,18,opt,name=is_trusted,json=isTrusted,proto3" json:"is_trusted"`
	Status       string       `protobuf:"bytes,19,opt,name=status,proto3" json:"status"`
	CreatorId    int64        `protobuf:"varint,20,opt,name=creator_id,json=creatorId,proto3" json:"creator_id"`
	ExpiredAt    string       `protobuf:"bytes,21,opt,name=expired_at,json=expiredAt,proto3" json:"expired_at"`
	CreatedAt    string       `protobuf:"bytes,22,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt    string       `protobuf:"bytes,23,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	Version      *Version     `protobuf:"bytes,24,opt,name=version,proto3" json:"version"`
	StatusName   string       `protobuf:"bytes,25,opt,name=status_name,json=statusName,proto3" json:"status_name"`
	ValidDays    int32        `protobuf:"varint,26,opt,name=valid_days,json=validDays,proto3" json:"valid_days"`
	ProvinceName string       `protobuf:"bytes,27,opt,name=province_name,json=provinceName,proto3" json:"province_name"`
	CityName     string       `protobuf:"bytes,28,opt,name=city_name,json=cityName,proto3" json:"city_name"`
	CountyName   string       `protobuf:"bytes,29,opt,name=county_name,json=countyName,proto3" json:"county_name"`
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

func (x *Store) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Store) GetLogoUrl() string {
	if x != nil {
		return x.LogoUrl
	}
	return ""
}

func (x *Store) GetVersionId() int32 {
	if x != nil {
		return x.VersionId
	}
	return 0
}

func (x *Store) GetCompanyId() int64 {
	if x != nil {
		return x.CompanyId
	}
	return 0
}

func (x *Store) GetCompanyName() string {
	if x != nil {
		return x.CompanyName
	}
	return ""
}

func (x *Store) GetIndustry() string {
	if x != nil {
		return x.Industry
	}
	return ""
}

func (x *Store) GetProvinceId() int64 {
	if x != nil {
		return x.ProvinceId
	}
	return 0
}

func (x *Store) GetCityId() int64 {
	if x != nil {
		return x.CityId
	}
	return 0
}

func (x *Store) GetCountyId() int64 {
	if x != nil {
		return x.CountyId
	}
	return 0
}

func (x *Store) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Store) GetLng() float32 {
	if x != nil {
		return x.Lng
	}
	return 0
}

func (x *Store) GetLat() float32 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *Store) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *Store) GetConfig() *StoreConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *Store) GetIsTrusted() string {
	if x != nil {
		return x.IsTrusted
	}
	return ""
}

func (x *Store) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Store) GetCreatorId() int64 {
	if x != nil {
		return x.CreatorId
	}
	return 0
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

func (x *Store) GetVersion() *Version {
	if x != nil {
		return x.Version
	}
	return nil
}

func (x *Store) GetStatusName() string {
	if x != nil {
		return x.StatusName
	}
	return ""
}

func (x *Store) GetValidDays() int32 {
	if x != nil {
		return x.ValidDays
	}
	return 0
}

func (x *Store) GetProvinceName() string {
	if x != nil {
		return x.ProvinceName
	}
	return ""
}

func (x *Store) GetCityName() string {
	if x != nil {
		return x.CityName
	}
	return ""
}

func (x *Store) GetCountyName() string {
	if x != nil {
		return x.CountyName
	}
	return ""
}

type StoreConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomDomain string `protobuf:"bytes,1,opt,name=custom_domain,json=customDomain,proto3" json:"custom_domain"`
}

func (x *StoreConfig) Reset() {
	*x = StoreConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storeService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreConfig) ProtoMessage() {}

func (x *StoreConfig) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use StoreConfig.ProtoReflect.Descriptor instead.
func (*StoreConfig) Descriptor() ([]byte, []int) {
	return file_storeService_proto_rawDescGZIP(), []int{1}
}

func (x *StoreConfig) GetCustomDomain() string {
	if x != nil {
		return x.CustomDomain
	}
	return ""
}

type StoreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged         int64   `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize      int64   `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Sorting       string  `protobuf:"bytes,3,opt,name=sorting,proto3" json:"sorting"`
	Id            int64   `protobuf:"varint,4,opt,name=id,proto3" json:"id"`
	Name          string  `protobuf:"bytes,5,opt,name=name,proto3" json:"name"`
	StoreKey      string  `protobuf:"bytes,6,opt,name=store_key,json=storeKey,proto3" json:"store_key"`
	StoreSecret   string  `protobuf:"bytes,7,opt,name=store_secret,json=storeSecret,proto3" json:"store_secret"`
	Industry      string  `protobuf:"bytes,8,opt,name=industry,proto3" json:"industry"`
	Status        string  `protobuf:"bytes,9,opt,name=status,proto3" json:"status"`
	StartDate     string  `protobuf:"bytes,10,opt,name=start_date,json=startDate,proto3" json:"start_date"`
	EndDate       string  `protobuf:"bytes,11,opt,name=end_date,json=endDate,proto3" json:"end_date"`
	Keywords      string  `protobuf:"bytes,12,opt,name=keywords,proto3" json:"keywords"`
	Region        string  `protobuf:"bytes,13,opt,name=region,proto3" json:"region"`
	CreatorId     int64   `protobuf:"varint,14,opt,name=creator_id,json=creatorId,proto3" json:"creator_id"`
	ApplicationId int32   `protobuf:"varint,15,opt,name=application_id,json=applicationId,proto3" json:"application_id"`
	AreaId        int64   `protobuf:"varint,16,opt,name=area_id,json=areaId,proto3" json:"area_id"`
	Ids           []int64 `protobuf:"varint,17,rep,packed,name=ids,proto3" json:"ids"`
}

func (x *StoreRequest) Reset() {
	*x = StoreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storeService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreRequest) ProtoMessage() {}

func (x *StoreRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use StoreRequest.ProtoReflect.Descriptor instead.
func (*StoreRequest) Descriptor() ([]byte, []int) {
	return file_storeService_proto_rawDescGZIP(), []int{2}
}

func (x *StoreRequest) GetPaged() int64 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *StoreRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *StoreRequest) GetSorting() string {
	if x != nil {
		return x.Sorting
	}
	return ""
}

func (x *StoreRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *StoreRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StoreRequest) GetStoreKey() string {
	if x != nil {
		return x.StoreKey
	}
	return ""
}

func (x *StoreRequest) GetStoreSecret() string {
	if x != nil {
		return x.StoreSecret
	}
	return ""
}

func (x *StoreRequest) GetIndustry() string {
	if x != nil {
		return x.Industry
	}
	return ""
}

func (x *StoreRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *StoreRequest) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *StoreRequest) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

func (x *StoreRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *StoreRequest) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *StoreRequest) GetCreatorId() int64 {
	if x != nil {
		return x.CreatorId
	}
	return 0
}

func (x *StoreRequest) GetApplicationId() int32 {
	if x != nil {
		return x.ApplicationId
	}
	return 0
}

func (x *StoreRequest) GetAreaId() int64 {
	if x != nil {
		return x.AreaId
	}
	return 0
}

func (x *StoreRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type StoreResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity  *Store        `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager   *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items   []*Store      `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Info    string        `protobuf:"bytes,4,opt,name=info,proto3" json:"info"`
	Modules []string      `protobuf:"bytes,5,rep,name=Modules,proto3" json:"Modules"`
}

func (x *StoreResponse) Reset() {
	*x = StoreResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storeService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreResponse) ProtoMessage() {}

func (x *StoreResponse) ProtoReflect() protoreflect.Message {
	mi := &file_storeService_proto_msgTypes[3]
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
	return file_storeService_proto_rawDescGZIP(), []int{3}
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

func (x *StoreResponse) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

func (x *StoreResponse) GetModules() []string {
	if x != nil {
		return x.Modules
	}
	return nil
}

var File_storeService_proto protoreflect.FileDescriptor

var file_storeService_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x14, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb8, 0x06, 0x0a, 0x05, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6c,
	0x6f, 0x67, 0x6f, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c,
	0x6f, 0x67, 0x6f, 0x55, 0x72, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x64, 0x75, 0x73,
	0x74, 0x72, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6e, 0x64, 0x75, 0x73,
	0x74, 0x72, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e,
	0x63, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x79, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6e, 0x67, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x03, 0x6c, 0x6e, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x12, 0x2d, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x74, 0x72, 0x75, 0x73, 0x74, 0x65, 0x64, 0x18, 0x12, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x73, 0x54, 0x72, 0x75, 0x73, 0x74, 0x65, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x14, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x2b, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x18, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x19,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x64, 0x61, 0x79, 0x73, 0x18, 0x1a,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x44, 0x61, 0x79, 0x73, 0x12,
	0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x1b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x69, 0x74, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x1d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x79, 0x4e, 0x61,
	0x6d, 0x65, 0x22, 0x32, 0x0a, 0x0b, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0xd2, 0x03, 0x0a, 0x0c, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6f,
	0x72, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6f, 0x72,
	0x74, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x64, 0x75,
	0x73, 0x74, 0x72, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6e, 0x64, 0x75,
	0x73, 0x74, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65,
	0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65,
	0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72,
	0x64, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72,
	0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0d, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x61, 0x72, 0x65, 0x61, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73,
	0x18, 0x11, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0xb2, 0x01, 0x0a, 0x0d,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a,
	0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x06,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50,
	0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73,
	0x32, 0xdc, 0x04, 0x0a, 0x0c, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x34, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x1a, 0x17, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x3a, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a,
	0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x03, 0x47, 0x65,
	0x74, 0x12, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a,
	0x09, 0x47, 0x65, 0x74, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x0f, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x1a, 0x17, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x12, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_storeService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_storeService_proto_goTypes = []interface{}{
	(*Store)(nil),         // 0: services.Store
	(*StoreConfig)(nil),   // 1: services.StoreConfig
	(*StoreRequest)(nil),  // 2: services.StoreRequest
	(*StoreResponse)(nil), // 3: services.StoreResponse
	(*Version)(nil),       // 4: services.Version
	(*common.Pager)(nil),  // 5: common.Pager
}
var file_storeService_proto_depIdxs = []int32{
	1,  // 0: services.Store.config:type_name -> services.StoreConfig
	4,  // 1: services.Store.version:type_name -> services.Version
	0,  // 2: services.StoreResponse.entity:type_name -> services.Store
	5,  // 3: services.StoreResponse.pager:type_name -> common.Pager
	0,  // 4: services.StoreResponse.items:type_name -> services.Store
	0,  // 5: services.StoreService.Create:input_type -> services.Store
	0,  // 6: services.StoreService.UpdateVersion:input_type -> services.Store
	0,  // 7: services.StoreService.UpdateProfile:input_type -> services.Store
	0,  // 8: services.StoreService.UpdateStatus:input_type -> services.Store
	2,  // 9: services.StoreService.Delete:input_type -> services.StoreRequest
	0,  // 10: services.StoreService.Get:input_type -> services.Store
	0,  // 11: services.StoreService.GetByName:input_type -> services.Store
	2,  // 12: services.StoreService.Search:input_type -> services.StoreRequest
	2,  // 13: services.StoreService.List:input_type -> services.StoreRequest
	2,  // 14: services.StoreService.GetModules:input_type -> services.StoreRequest
	3,  // 15: services.StoreService.Create:output_type -> services.StoreResponse
	3,  // 16: services.StoreService.UpdateVersion:output_type -> services.StoreResponse
	3,  // 17: services.StoreService.UpdateProfile:output_type -> services.StoreResponse
	3,  // 18: services.StoreService.UpdateStatus:output_type -> services.StoreResponse
	3,  // 19: services.StoreService.Delete:output_type -> services.StoreResponse
	3,  // 20: services.StoreService.Get:output_type -> services.StoreResponse
	3,  // 21: services.StoreService.GetByName:output_type -> services.StoreResponse
	3,  // 22: services.StoreService.Search:output_type -> services.StoreResponse
	3,  // 23: services.StoreService.List:output_type -> services.StoreResponse
	3,  // 24: services.StoreService.GetModules:output_type -> services.StoreResponse
	15, // [15:25] is the sub-list for method output_type
	5,  // [5:15] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_storeService_proto_init() }
func file_storeService_proto_init() {
	if File_storeService_proto != nil {
		return
	}
	file_versionService_proto_init()
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
		file_storeService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreRequest); i {
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
		file_storeService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
			NumMessages:   4,
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
