// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: brandService.proto

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

// 商品品牌
type Brand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32      `protobuf:"varint,1,opt,name=id,proto3" json:"id"`                                        //ID
	Name        string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`                                     //品牌名称
	ShopId      int64      `protobuf:"varint,3,opt,name=shop_id,json=shopId,proto3" json:"shop_id"`                  //分店ID
	BrandTypeId int64      `protobuf:"varint,4,opt,name=brand_type_id,json=brandTypeId,proto3" json:"brand_type_id"` //品牌类型
	LogoUrl     string     `protobuf:"bytes,5,opt,name=logo_url,json=logoUrl,proto3" json:"logo_url"`                //LOGO
	Letter      string     `protobuf:"bytes,6,opt,name=letter,proto3" json:"letter"`                                 //检索首字母
	Website     string     `protobuf:"bytes,7,opt,name=website,proto3" json:"website"`                               //品牌网站
	Desc        string     `protobuf:"bytes,8,opt,name=desc,proto3" json:"desc"`                                     //描述
	Sort        int32      `protobuf:"varint,9,opt,name=sort,proto3" json:"sort"`                                    //排序
	SpuCount    int32      `protobuf:"varint,10,opt,name=spu_count,json=spuCount,proto3" json:"spu_count"`           //商品数量
	CreatedAt   int64      `protobuf:"varint,11,opt,name=created_at,json=createdAt,proto3" json:"created_at"`        //创建时间
	UpdatedAt   int64      `protobuf:"varint,12,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`        //修改时间
	BrandType   *BrandType `protobuf:"bytes,13,opt,name=brand_type,json=brandType,proto3" json:"brand_type"`
}

func (x *Brand) Reset() {
	*x = Brand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_brandService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Brand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Brand) ProtoMessage() {}

func (x *Brand) ProtoReflect() protoreflect.Message {
	mi := &file_brandService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Brand.ProtoReflect.Descriptor instead.
func (*Brand) Descriptor() ([]byte, []int) {
	return file_brandService_proto_rawDescGZIP(), []int{0}
}

func (x *Brand) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Brand) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Brand) GetShopId() int64 {
	if x != nil {
		return x.ShopId
	}
	return 0
}

func (x *Brand) GetBrandTypeId() int64 {
	if x != nil {
		return x.BrandTypeId
	}
	return 0
}

func (x *Brand) GetLogoUrl() string {
	if x != nil {
		return x.LogoUrl
	}
	return ""
}

func (x *Brand) GetLetter() string {
	if x != nil {
		return x.Letter
	}
	return ""
}

func (x *Brand) GetWebsite() string {
	if x != nil {
		return x.Website
	}
	return ""
}

func (x *Brand) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *Brand) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *Brand) GetSpuCount() int32 {
	if x != nil {
		return x.SpuCount
	}
	return 0
}

func (x *Brand) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Brand) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *Brand) GetBrandType() *BrandType {
	if x != nil {
		return x.BrandType
	}
	return nil
}

// 商品品牌请求参数
type BrandRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Top       int32    `protobuf:"varint,1,opt,name=top,proto3" json:"top"`
	Paged     int64    `protobuf:"varint,2,opt,name=paged,proto3" json:"paged"`
	PageSize  int64    `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Keywords  string   `protobuf:"bytes,4,opt,name=keywords,proto3" json:"keywords"`
	Sorts     []string `protobuf:"bytes,5,rep,name=sorts,proto3" json:"sorts"`
	DateRange []string `protobuf:"bytes,6,rep,name=date_range,json=dateRange,proto3" json:"date_range"`
	Ids       []int32  `protobuf:"varint,7,rep,packed,name=ids,proto3" json:"ids"`
	Id        int32    `protobuf:"varint,8,opt,name=id,proto3" json:"id"`
	//以下为自定义参数
	Name        string `protobuf:"bytes,11,opt,name=name,proto3" json:"name"`                                     //品牌名称
	ShopId      int64  `protobuf:"varint,12,opt,name=shop_id,json=shopId,proto3" json:"shop_id"`                  //分店ID
	BrandTypeId int64  `protobuf:"varint,13,opt,name=brand_type_id,json=brandTypeId,proto3" json:"brand_type_id"` //品牌类型
}

func (x *BrandRequest) Reset() {
	*x = BrandRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_brandService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BrandRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrandRequest) ProtoMessage() {}

func (x *BrandRequest) ProtoReflect() protoreflect.Message {
	mi := &file_brandService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BrandRequest.ProtoReflect.Descriptor instead.
func (*BrandRequest) Descriptor() ([]byte, []int) {
	return file_brandService_proto_rawDescGZIP(), []int{1}
}

func (x *BrandRequest) GetTop() int32 {
	if x != nil {
		return x.Top
	}
	return 0
}

func (x *BrandRequest) GetPaged() int64 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *BrandRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *BrandRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *BrandRequest) GetSorts() []string {
	if x != nil {
		return x.Sorts
	}
	return nil
}

func (x *BrandRequest) GetDateRange() []string {
	if x != nil {
		return x.DateRange
	}
	return nil
}

func (x *BrandRequest) GetIds() []int32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *BrandRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BrandRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BrandRequest) GetShopId() int64 {
	if x != nil {
		return x.ShopId
	}
	return 0
}

func (x *BrandRequest) GetBrandTypeId() int64 {
	if x != nil {
		return x.BrandTypeId
	}
	return 0
}

// 商品品牌响应数据
type BrandResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg    string        `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Entity *Brand        `protobuf:"bytes,3,opt,name=entity,proto3" json:"entity"`
	Items  []*Brand      `protobuf:"bytes,4,rep,name=items,proto3" json:"items"`
}

func (x *BrandResponse) Reset() {
	*x = BrandResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_brandService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BrandResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrandResponse) ProtoMessage() {}

func (x *BrandResponse) ProtoReflect() protoreflect.Message {
	mi := &file_brandService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BrandResponse.ProtoReflect.Descriptor instead.
func (*BrandResponse) Descriptor() ([]byte, []int) {
	return file_brandService_proto_rawDescGZIP(), []int{2}
}

func (x *BrandResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *BrandResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *BrandResponse) GetEntity() *Brand {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *BrandResponse) GetItems() []*Brand {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_brandService_proto protoreflect.FileDescriptor

var file_brandService_proto_rawDesc = []byte{
	0x0a, 0x12, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x16, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xec, 0x02, 0x0a, 0x05, 0x42, 0x72,
	0x61, 0x6e, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x68, 0x6f, 0x70, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x68, 0x6f, 0x70, 0x49, 0x64,
	0x12, 0x22, 0x0a, 0x0d, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x6f, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f, 0x67, 0x6f, 0x55, 0x72, 0x6c, 0x12,
	0x16, 0x0a, 0x06, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x65, 0x62, 0x73, 0x69,
	0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x70, 0x75,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x70,
	0x75, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x32, 0x0a, 0x0a, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x62,
	0x72, 0x61, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x22, 0x97, 0x02, 0x0a, 0x0c, 0x42, 0x72, 0x61,
	0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x6f, 0x70,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x74, 0x6f, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x61, 0x67, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x6f,
	0x72, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x73, 0x6f, 0x72, 0x74, 0x73,
	0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x64, 0x61, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x05, 0x52, 0x03, 0x69, 0x64,
	0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x69, 0x64,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x12, 0x22,
	0x0a, 0x0d, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65,
	0x49, 0x64, 0x22, 0x96, 0x01, 0x0a, 0x0d, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50,
	0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x06, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x52, 0x06, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x12, 0x25, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x42,
	0x72, 0x61, 0x6e, 0x64, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x32, 0xdb, 0x02, 0x0a, 0x0c,
	0x42, 0x72, 0x61, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x06,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x34, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x1a, 0x17, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x72,
	0x61, 0x6e, 0x64, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x42,
	0x72, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x31,
	0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x3b, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x16, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x42,
	0x72, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39,
	0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_brandService_proto_rawDescOnce sync.Once
	file_brandService_proto_rawDescData = file_brandService_proto_rawDesc
)

func file_brandService_proto_rawDescGZIP() []byte {
	file_brandService_proto_rawDescOnce.Do(func() {
		file_brandService_proto_rawDescData = protoimpl.X.CompressGZIP(file_brandService_proto_rawDescData)
	})
	return file_brandService_proto_rawDescData
}

var file_brandService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_brandService_proto_goTypes = []interface{}{
	(*Brand)(nil),         // 0: services.Brand
	(*BrandRequest)(nil),  // 1: services.BrandRequest
	(*BrandResponse)(nil), // 2: services.BrandResponse
	(*BrandType)(nil),     // 3: services.BrandType
	(*common.Pager)(nil),  // 4: common.Pager
}
var file_brandService_proto_depIdxs = []int32{
	3,  // 0: services.Brand.brand_type:type_name -> services.BrandType
	4,  // 1: services.BrandResponse.pager:type_name -> common.Pager
	0,  // 2: services.BrandResponse.entity:type_name -> services.Brand
	0,  // 3: services.BrandResponse.items:type_name -> services.Brand
	0,  // 4: services.BrandService.Create:input_type -> services.Brand
	0,  // 5: services.BrandService.Update:input_type -> services.Brand
	0,  // 6: services.BrandService.Delete:input_type -> services.Brand
	0,  // 7: services.BrandService.Get:input_type -> services.Brand
	1,  // 8: services.BrandService.Search:input_type -> services.BrandRequest
	1,  // 9: services.BrandService.List:input_type -> services.BrandRequest
	2,  // 10: services.BrandService.Create:output_type -> services.BrandResponse
	2,  // 11: services.BrandService.Update:output_type -> services.BrandResponse
	2,  // 12: services.BrandService.Delete:output_type -> services.BrandResponse
	2,  // 13: services.BrandService.Get:output_type -> services.BrandResponse
	2,  // 14: services.BrandService.Search:output_type -> services.BrandResponse
	2,  // 15: services.BrandService.List:output_type -> services.BrandResponse
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_brandService_proto_init() }
func file_brandService_proto_init() {
	if File_brandService_proto != nil {
		return
	}
	file_brandTypeService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_brandService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Brand); i {
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
		file_brandService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BrandRequest); i {
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
		file_brandService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BrandResponse); i {
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
			RawDescriptor: file_brandService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_brandService_proto_goTypes,
		DependencyIndexes: file_brandService_proto_depIdxs,
		MessageInfos:      file_brandService_proto_msgTypes,
	}.Build()
	File_brandService_proto = out.File
	file_brandService_proto_rawDesc = nil
	file_brandService_proto_goTypes = nil
	file_brandService_proto_depIdxs = nil
}
