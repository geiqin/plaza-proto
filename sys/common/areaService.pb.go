// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: areaService.proto

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

// 区域信息
type Area struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AreaId            int64   `protobuf:"varint,2,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	LevelType         string  `protobuf:"bytes,3,opt,name=level_type,json=levelType,proto3" json:"level_type,omitempty"`
	Name              string  `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Shortname         string  `protobuf:"bytes,5,opt,name=shortname,proto3" json:"shortname,omitempty"`
	ParentPath        string  `protobuf:"bytes,6,opt,name=parent_path,json=parentPath,proto3" json:"parent_path,omitempty"`
	Province          string  `protobuf:"bytes,7,opt,name=province,proto3" json:"province,omitempty"`
	City              string  `protobuf:"bytes,8,opt,name=city,proto3" json:"city,omitempty"`
	District          string  `protobuf:"bytes,9,opt,name=district,proto3" json:"district,omitempty"`
	ProvinceShortname string  `protobuf:"bytes,10,opt,name=province_shortname,json=provinceShortname,proto3" json:"province_shortname,omitempty"`
	CityShortname     string  `protobuf:"bytes,11,opt,name=city_shortname,json=cityShortname,proto3" json:"city_shortname,omitempty"`
	DistrictShortname string  `protobuf:"bytes,12,opt,name=district_shortname,json=districtShortname,proto3" json:"district_shortname,omitempty"`
	ProvincePinyin    string  `protobuf:"bytes,13,opt,name=province_pinyin,json=provincePinyin,proto3" json:"province_pinyin,omitempty"`
	CityPinyin        string  `protobuf:"bytes,14,opt,name=city_pinyin,json=cityPinyin,proto3" json:"city_pinyin,omitempty"`
	DistrictPinyin    string  `protobuf:"bytes,15,opt,name=district_pinyin,json=districtPinyin,proto3" json:"district_pinyin,omitempty"`
	Pinyin            string  `protobuf:"bytes,16,opt,name=pinyin,proto3" json:"pinyin,omitempty"`
	Jianpin           string  `protobuf:"bytes,17,opt,name=jianpin,proto3" json:"jianpin,omitempty"`
	FirstChar         string  `protobuf:"bytes,18,opt,name=first_char,json=firstChar,proto3" json:"first_char,omitempty"`
	CityCode          string  `protobuf:"bytes,19,opt,name=city_code,json=cityCode,proto3" json:"city_code,omitempty"`
	ZipCode           string  `protobuf:"bytes,20,opt,name=zip_code,json=zipCode,proto3" json:"zip_code,omitempty"`
	Lng               string  `protobuf:"bytes,21,opt,name=lng,proto3" json:"lng,omitempty"`
	Lat               string  `protobuf:"bytes,22,opt,name=lat,proto3" json:"lat,omitempty"`
	Remark1           string  `protobuf:"bytes,23,opt,name=remark1,proto3" json:"remark1,omitempty"`
	Remark2           string  `protobuf:"bytes,24,opt,name=remark2,proto3" json:"remark2,omitempty"`
	Children          []*Area `protobuf:"bytes,25,rep,name=children,proto3" json:"children,omitempty"`
}

func (x *Area) Reset() {
	*x = Area{}
	if protoimpl.UnsafeEnabled {
		mi := &file_areaService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Area) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Area) ProtoMessage() {}

func (x *Area) ProtoReflect() protoreflect.Message {
	mi := &file_areaService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Area.ProtoReflect.Descriptor instead.
func (*Area) Descriptor() ([]byte, []int) {
	return file_areaService_proto_rawDescGZIP(), []int{0}
}

func (x *Area) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Area) GetAreaId() int64 {
	if x != nil {
		return x.AreaId
	}
	return 0
}

func (x *Area) GetLevelType() string {
	if x != nil {
		return x.LevelType
	}
	return ""
}

func (x *Area) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Area) GetShortname() string {
	if x != nil {
		return x.Shortname
	}
	return ""
}

func (x *Area) GetParentPath() string {
	if x != nil {
		return x.ParentPath
	}
	return ""
}

func (x *Area) GetProvince() string {
	if x != nil {
		return x.Province
	}
	return ""
}

func (x *Area) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Area) GetDistrict() string {
	if x != nil {
		return x.District
	}
	return ""
}

func (x *Area) GetProvinceShortname() string {
	if x != nil {
		return x.ProvinceShortname
	}
	return ""
}

func (x *Area) GetCityShortname() string {
	if x != nil {
		return x.CityShortname
	}
	return ""
}

func (x *Area) GetDistrictShortname() string {
	if x != nil {
		return x.DistrictShortname
	}
	return ""
}

func (x *Area) GetProvincePinyin() string {
	if x != nil {
		return x.ProvincePinyin
	}
	return ""
}

func (x *Area) GetCityPinyin() string {
	if x != nil {
		return x.CityPinyin
	}
	return ""
}

func (x *Area) GetDistrictPinyin() string {
	if x != nil {
		return x.DistrictPinyin
	}
	return ""
}

func (x *Area) GetPinyin() string {
	if x != nil {
		return x.Pinyin
	}
	return ""
}

func (x *Area) GetJianpin() string {
	if x != nil {
		return x.Jianpin
	}
	return ""
}

func (x *Area) GetFirstChar() string {
	if x != nil {
		return x.FirstChar
	}
	return ""
}

func (x *Area) GetCityCode() string {
	if x != nil {
		return x.CityCode
	}
	return ""
}

func (x *Area) GetZipCode() string {
	if x != nil {
		return x.ZipCode
	}
	return ""
}

func (x *Area) GetLng() string {
	if x != nil {
		return x.Lng
	}
	return ""
}

func (x *Area) GetLat() string {
	if x != nil {
		return x.Lat
	}
	return ""
}

func (x *Area) GetRemark1() string {
	if x != nil {
		return x.Remark1
	}
	return ""
}

func (x *Area) GetRemark2() string {
	if x != nil {
		return x.Remark2
	}
	return ""
}

func (x *Area) GetChildren() []*Area {
	if x != nil {
		return x.Children
	}
	return nil
}

type AreaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Area         `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items  []*Area       `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error  *common.Error `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *AreaResponse) Reset() {
	*x = AreaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_areaService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AreaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AreaResponse) ProtoMessage() {}

func (x *AreaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_areaService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AreaResponse.ProtoReflect.Descriptor instead.
func (*AreaResponse) Descriptor() ([]byte, []int) {
	return file_areaService_proto_rawDescGZIP(), []int{1}
}

func (x *AreaResponse) GetEntity() *Area {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *AreaResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *AreaResponse) GetItems() []*Area {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *AreaResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_areaService_proto protoreflect.FileDescriptor

var file_areaService_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x72, 0x65, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xf2, 0x05, 0x0a, 0x04, 0x41, 0x72, 0x65, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x72, 0x65,
	0x61, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x72, 0x65, 0x61,
	0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x50, 0x61, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x63, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74,
	0x12, 0x2d, 0x0a, 0x12, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x5f, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x25, 0x0a, 0x0e, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x69, 0x74, 0x79, 0x53, 0x68, 0x6f,
	0x72, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x12, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69,
	0x63, 0x74, 0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x11, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x53, 0x68, 0x6f, 0x72,
	0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63,
	0x65, 0x5f, 0x70, 0x69, 0x6e, 0x79, 0x69, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x50, 0x69, 0x6e, 0x79, 0x69, 0x6e, 0x12, 0x1f,
	0x0a, 0x0b, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x70, 0x69, 0x6e, 0x79, 0x69, 0x6e, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x69, 0x74, 0x79, 0x50, 0x69, 0x6e, 0x79, 0x69, 0x6e, 0x12,
	0x27, 0x0a, 0x0f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x5f, 0x70, 0x69, 0x6e, 0x79,
	0x69, 0x6e, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69,
	0x63, 0x74, 0x50, 0x69, 0x6e, 0x79, 0x69, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x69, 0x6e, 0x79,
	0x69, 0x6e, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x69, 0x6e, 0x79, 0x69, 0x6e,
	0x12, 0x18, 0x0a, 0x07, 0x6a, 0x69, 0x61, 0x6e, 0x70, 0x69, 0x6e, 0x18, 0x11, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6a, 0x69, 0x61, 0x6e, 0x70, 0x69, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x43, 0x68, 0x61, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x69, 0x74,
	0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x69,
	0x74, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x7a, 0x69, 0x70, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x7a, 0x69, 0x70, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6e, 0x67, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6c, 0x6e, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x31,
	0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x31, 0x12,
	0x18, 0x0a, 0x07, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x32, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x32, 0x12, 0x2a, 0x0a, 0x08, 0x63, 0x68, 0x69,
	0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18, 0x19, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x72, 0x65, 0x61, 0x52, 0x08, 0x63, 0x68, 0x69,
	0x6c, 0x64, 0x72, 0x65, 0x6e, 0x22, 0xa6, 0x01, 0x0a, 0x0c, 0x41, 0x72, 0x65, 0x61, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x41, 0x72, 0x65, 0x61, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23,
	0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61,
	0x67, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x72,
	0x65, 0x61, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xd9,
	0x01, 0x0a, 0x0b, 0x41, 0x72, 0x65, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2f,
	0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x41, 0x72, 0x65, 0x61, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x41, 0x72, 0x65, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x35, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x16, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x72, 0x65, 0x61, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x04, 0x54, 0x72, 0x65, 0x65, 0x12, 0x0e,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x72, 0x65, 0x61, 0x1a, 0x16,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x72, 0x65, 0x61, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x72, 0x65, 0x61,
	0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x72, 0x65, 0x61,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_areaService_proto_rawDescOnce sync.Once
	file_areaService_proto_rawDescData = file_areaService_proto_rawDesc
)

func file_areaService_proto_rawDescGZIP() []byte {
	file_areaService_proto_rawDescOnce.Do(func() {
		file_areaService_proto_rawDescData = protoimpl.X.CompressGZIP(file_areaService_proto_rawDescData)
	})
	return file_areaService_proto_rawDescData
}

var file_areaService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_areaService_proto_goTypes = []interface{}{
	(*Area)(nil),             // 0: services.Area
	(*AreaResponse)(nil),     // 1: services.AreaResponse
	(*common.Pager)(nil),     // 2: common.Pager
	(*common.Error)(nil),     // 3: common.Error
	(*common.BaseWhere)(nil), // 4: common.BaseWhere
}
var file_areaService_proto_depIdxs = []int32{
	0, // 0: services.Area.children:type_name -> services.Area
	0, // 1: services.AreaResponse.entity:type_name -> services.Area
	2, // 2: services.AreaResponse.pager:type_name -> common.Pager
	0, // 3: services.AreaResponse.items:type_name -> services.Area
	3, // 4: services.AreaResponse.error:type_name -> common.Error
	0, // 5: services.AreaService.Get:input_type -> services.Area
	4, // 6: services.AreaService.Search:input_type -> common.BaseWhere
	0, // 7: services.AreaService.Tree:input_type -> services.Area
	0, // 8: services.AreaService.List:input_type -> services.Area
	1, // 9: services.AreaService.Get:output_type -> services.AreaResponse
	1, // 10: services.AreaService.Search:output_type -> services.AreaResponse
	1, // 11: services.AreaService.Tree:output_type -> services.AreaResponse
	1, // 12: services.AreaService.List:output_type -> services.AreaResponse
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_areaService_proto_init() }
func file_areaService_proto_init() {
	if File_areaService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_areaService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Area); i {
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
		file_areaService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AreaResponse); i {
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
			RawDescriptor: file_areaService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_areaService_proto_goTypes,
		DependencyIndexes: file_areaService_proto_depIdxs,
		MessageInfos:      file_areaService_proto_msgTypes,
	}.Build()
	File_areaService_proto = out.File
	file_areaService_proto_rawDesc = nil
	file_areaService_proto_goTypes = nil
	file_areaService_proto_depIdxs = nil
}
