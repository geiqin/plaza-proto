// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: skuInfoService.proto

package services

import (
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

type SkuInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SkuId       int64   `protobuf:"varint,1,opt,name=sku_id,json=skuId,proto3" json:"sku_id"`
	SkuSn       string  `protobuf:"bytes,2,opt,name=sku_sn,json=skuSn,proto3" json:"sku_sn"`
	ModelType   string  `protobuf:"bytes,3,opt,name=model_type,json=modelType,proto3" json:"model_type"`
	Name        string  `protobuf:"bytes,4,opt,name=name,proto3" json:"name"`
	Unit        string  `protobuf:"bytes,5,opt,name=unit,proto3" json:"unit"`
	BrandId     int32   `protobuf:"varint,6,opt,name=brand_id,json=brandId,proto3" json:"brand_id"`
	CategoryId  int32   `protobuf:"varint,7,opt,name=category_id,json=categoryId,proto3" json:"category_id"`
	Quantity    int32   `protobuf:"varint,8,opt,name=quantity,proto3" json:"quantity"`
	ThumbId     int64   `protobuf:"varint,9,opt,name=thumb_id,json=thumbId,proto3" json:"thumb_id"`
	ThumbUrl    string  `protobuf:"bytes,10,opt,name=thumb_url,json=thumbUrl,proto3" json:"thumb_url"`
	Barcode     string  `protobuf:"bytes,11,opt,name=barcode,proto3" json:"barcode"`
	Price       int64   `protobuf:"varint,12,opt,name=price,proto3" json:"price"`
	OriginPrice int64   `protobuf:"varint,13,opt,name=origin_price,json=originPrice,proto3" json:"origin_price"`
	CostPrice   int64   `protobuf:"varint,14,opt,name=cost_price,json=costPrice,proto3" json:"cost_price"`
	Weight      float32 `protobuf:"fixed32,15,opt,name=weight,proto3" json:"weight"`
}

func (x *SkuInfo) Reset() {
	*x = SkuInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_skuInfoService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SkuInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SkuInfo) ProtoMessage() {}

func (x *SkuInfo) ProtoReflect() protoreflect.Message {
	mi := &file_skuInfoService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SkuInfo.ProtoReflect.Descriptor instead.
func (*SkuInfo) Descriptor() ([]byte, []int) {
	return file_skuInfoService_proto_rawDescGZIP(), []int{0}
}

func (x *SkuInfo) GetSkuId() int64 {
	if x != nil {
		return x.SkuId
	}
	return 0
}

func (x *SkuInfo) GetSkuSn() string {
	if x != nil {
		return x.SkuSn
	}
	return ""
}

func (x *SkuInfo) GetModelType() string {
	if x != nil {
		return x.ModelType
	}
	return ""
}

func (x *SkuInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SkuInfo) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *SkuInfo) GetBrandId() int32 {
	if x != nil {
		return x.BrandId
	}
	return 0
}

func (x *SkuInfo) GetCategoryId() int32 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

func (x *SkuInfo) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *SkuInfo) GetThumbId() int64 {
	if x != nil {
		return x.ThumbId
	}
	return 0
}

func (x *SkuInfo) GetThumbUrl() string {
	if x != nil {
		return x.ThumbUrl
	}
	return ""
}

func (x *SkuInfo) GetBarcode() string {
	if x != nil {
		return x.Barcode
	}
	return ""
}

func (x *SkuInfo) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *SkuInfo) GetOriginPrice() int64 {
	if x != nil {
		return x.OriginPrice
	}
	return 0
}

func (x *SkuInfo) GetCostPrice() int64 {
	if x != nil {
		return x.CostPrice
	}
	return 0
}

func (x *SkuInfo) GetWeight() float32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

var File_skuInfoService_proto protoreflect.FileDescriptor

var file_skuInfoService_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x6b, 0x75, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x22, 0x98, 0x03, 0x0a, 0x07, 0x53, 0x6b, 0x75, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x15, 0x0a, 0x06,
	0x73, 0x6b, 0x75, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x6b,
	0x75, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x6b, 0x75, 0x5f, 0x73, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x6b, 0x75, 0x53, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x6e, 0x69,
	0x74, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x68, 0x75,
	0x6d, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x74, 0x68, 0x75,
	0x6d, 0x62, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x55, 0x72,
	0x6c, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x5f, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x73, 0x74, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x42, 0x0c, 0x5a, 0x0a, 0x2e,
	0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_skuInfoService_proto_rawDescOnce sync.Once
	file_skuInfoService_proto_rawDescData = file_skuInfoService_proto_rawDesc
)

func file_skuInfoService_proto_rawDescGZIP() []byte {
	file_skuInfoService_proto_rawDescOnce.Do(func() {
		file_skuInfoService_proto_rawDescData = protoimpl.X.CompressGZIP(file_skuInfoService_proto_rawDescData)
	})
	return file_skuInfoService_proto_rawDescData
}

var file_skuInfoService_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_skuInfoService_proto_goTypes = []interface{}{
	(*SkuInfo)(nil), // 0: services.SkuInfo
}
var file_skuInfoService_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_skuInfoService_proto_init() }
func file_skuInfoService_proto_init() {
	if File_skuInfoService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_skuInfoService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SkuInfo); i {
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
			RawDescriptor: file_skuInfoService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_skuInfoService_proto_goTypes,
		DependencyIndexes: file_skuInfoService_proto_depIdxs,
		MessageInfos:      file_skuInfoService_proto_msgTypes,
	}.Build()
	File_skuInfoService_proto = out.File
	file_skuInfoService_proto_rawDesc = nil
	file_skuInfoService_proto_goTypes = nil
	file_skuInfoService_proto_depIdxs = nil
}
