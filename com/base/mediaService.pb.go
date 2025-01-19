// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: mediaService.proto

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

// 公共图标
type Media struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64      `protobuf:"varint,1,opt,name=id,proto3" json:"id"`                                //ID
	Hash        string     `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash"`                             //Hash值
	Code        string     `protobuf:"bytes,3,opt,name=code,proto3" json:"code"`                             //编码
	Name        string     `protobuf:"bytes,4,opt,name=name,proto3" json:"name"`                             //名称
	Storage     string     `protobuf:"bytes,5,opt,name=storage,proto3" json:"storage"`                       //储存名
	Type        string     `protobuf:"bytes,6,opt,name=type,proto3" json:"type"`                             //类型
	RawName     string     `protobuf:"bytes,8,opt,name=raw_name,json=rawName,proto3" json:"raw_name"`        //原始名称
	Extension   string     `protobuf:"bytes,9,opt,name=extension,proto3" json:"extension"`                   //文件扩展名
	SavePath    string     `protobuf:"bytes,10,opt,name=save_path,json=savePath,proto3" json:"save_path"`    //存相对路径
	Url         string     `protobuf:"bytes,11,opt,name=url,proto3" json:"url"`                              //完整路径
	MimeType    string     `protobuf:"bytes,12,opt,name=mime_type,json=mimeType,proto3" json:"mime_type"`    //文件类型
	Size        int64      `protobuf:"varint,13,opt,name=size,proto3" json:"size"`                           //文件大小
	Width       int64      `protobuf:"varint,14,opt,name=width,proto3" json:"width"`                         //宽
	Height      int64      `protobuf:"varint,15,opt,name=height,proto3" json:"height"`                       //高
	Description string     `protobuf:"bytes,16,opt,name=description,proto3" json:"description"`              //描述
	CatId       int64      `protobuf:"varint,17,opt,name=cat_id,json=catId,proto3" json:"cat_id"`            //分类
	Sort        int32      `protobuf:"varint,18,opt,name=sort,proto3" json:"sort"`                           //排序值
	CreatedAt   string     `protobuf:"bytes,19,opt,name=created_at,json=createdAt,proto3" json:"created_at"` //创建时间
	UpdatedAt   string     `protobuf:"bytes,20,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"` //修改时间
	Cat         *MediaCat  `protobuf:"bytes,21,opt,name=cat,proto3" json:"cat"`
	MediaInfo   *MediaInfo `protobuf:"bytes,22,opt,name=media_info,json=mediaInfo,proto3" json:"media_info"` //媒体信息
}

func (x *Media) Reset() {
	*x = Media{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mediaService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Media) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Media) ProtoMessage() {}

func (x *Media) ProtoReflect() protoreflect.Message {
	mi := &file_mediaService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Media.ProtoReflect.Descriptor instead.
func (*Media) Descriptor() ([]byte, []int) {
	return file_mediaService_proto_rawDescGZIP(), []int{0}
}

func (x *Media) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Media) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *Media) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Media) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Media) GetStorage() string {
	if x != nil {
		return x.Storage
	}
	return ""
}

func (x *Media) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Media) GetRawName() string {
	if x != nil {
		return x.RawName
	}
	return ""
}

func (x *Media) GetExtension() string {
	if x != nil {
		return x.Extension
	}
	return ""
}

func (x *Media) GetSavePath() string {
	if x != nil {
		return x.SavePath
	}
	return ""
}

func (x *Media) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Media) GetMimeType() string {
	if x != nil {
		return x.MimeType
	}
	return ""
}

func (x *Media) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Media) GetWidth() int64 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *Media) GetHeight() int64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Media) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Media) GetCatId() int64 {
	if x != nil {
		return x.CatId
	}
	return 0
}

func (x *Media) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *Media) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Media) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Media) GetCat() *MediaCat {
	if x != nil {
		return x.Cat
	}
	return nil
}

func (x *Media) GetMediaInfo() *MediaInfo {
	if x != nil {
		return x.MediaInfo
	}
	return nil
}

type MediaStatistics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type   string             `protobuf:"bytes,1,opt,name=type,proto3" json:"type"`      //类型
	Size   int64              `protobuf:"varint,2,opt,name=size,proto3" json:"size"`     //大小
	Number int64              `protobuf:"varint,3,opt,name=number,proto3" json:"number"` //数量
	Unit   string             `protobuf:"bytes,4,opt,name=unit,proto3" json:"unit"`      //单位
	Data   []*MediaStatistics `protobuf:"bytes,5,rep,name=data,proto3" json:"data"`      //子项统计
}

func (x *MediaStatistics) Reset() {
	*x = MediaStatistics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mediaService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MediaStatistics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MediaStatistics) ProtoMessage() {}

func (x *MediaStatistics) ProtoReflect() protoreflect.Message {
	mi := &file_mediaService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MediaStatistics.ProtoReflect.Descriptor instead.
func (*MediaStatistics) Descriptor() ([]byte, []int) {
	return file_mediaService_proto_rawDescGZIP(), []int{1}
}

func (x *MediaStatistics) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *MediaStatistics) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *MediaStatistics) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *MediaStatistics) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *MediaStatistics) GetData() []*MediaStatistics {
	if x != nil {
		return x.Data
	}
	return nil
}

// 公共图标请求参数
type MediaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Top       int32    `protobuf:"varint,1,opt,name=top,proto3" json:"top"`
	Paged     int64    `protobuf:"varint,2,opt,name=paged,proto3" json:"paged"`
	PageSize  int64    `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Keywords  string   `protobuf:"bytes,4,opt,name=keywords,proto3" json:"keywords"`
	Sort      []string `protobuf:"bytes,5,rep,name=sort,proto3" json:"sort"`
	DateRange []string `protobuf:"bytes,6,rep,name=date_range,json=dateRange,proto3" json:"date_range"`
	Ids       []int64  `protobuf:"varint,7,rep,packed,name=ids,proto3" json:"ids"`
	Id        int64    `protobuf:"varint,8,opt,name=id,proto3" json:"id"`
	//以下为自定义参数
	Code     string `protobuf:"bytes,11,opt,name=code,proto3" json:"code"`                         //编码
	Name     string `protobuf:"bytes,12,opt,name=name,proto3" json:"name"`                         //名称
	Storage  string `protobuf:"bytes,13,opt,name=storage,proto3" json:"storage"`                   //储存名
	Type     string `protobuf:"bytes,14,opt,name=type,proto3" json:"type"`                         //类型
	Filename string `protobuf:"bytes,15,opt,name=filename,proto3" json:"filename"`                 //文件名称
	MimeType string `protobuf:"bytes,16,opt,name=mime_type,json=mimeType,proto3" json:"mime_type"` //文件类型
	CatId    int64  `protobuf:"varint,17,opt,name=cat_id,json=catId,proto3" json:"cat_id"`
}

func (x *MediaRequest) Reset() {
	*x = MediaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mediaService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MediaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MediaRequest) ProtoMessage() {}

func (x *MediaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mediaService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MediaRequest.ProtoReflect.Descriptor instead.
func (*MediaRequest) Descriptor() ([]byte, []int) {
	return file_mediaService_proto_rawDescGZIP(), []int{2}
}

func (x *MediaRequest) GetTop() int32 {
	if x != nil {
		return x.Top
	}
	return 0
}

func (x *MediaRequest) GetPaged() int64 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *MediaRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *MediaRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *MediaRequest) GetSort() []string {
	if x != nil {
		return x.Sort
	}
	return nil
}

func (x *MediaRequest) GetDateRange() []string {
	if x != nil {
		return x.DateRange
	}
	return nil
}

func (x *MediaRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *MediaRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MediaRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *MediaRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MediaRequest) GetStorage() string {
	if x != nil {
		return x.Storage
	}
	return ""
}

func (x *MediaRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *MediaRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *MediaRequest) GetMimeType() string {
	if x != nil {
		return x.MimeType
	}
	return ""
}

func (x *MediaRequest) GetCatId() int64 {
	if x != nil {
		return x.CatId
	}
	return 0
}

// 公共图标响应数据
type MediaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg        string            `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg"`
	Pager      *common.Pager     `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Entity     *Media            `protobuf:"bytes,3,opt,name=entity,proto3" json:"entity"`
	Items      []*Media          `protobuf:"bytes,4,rep,name=items,proto3" json:"items"`
	Statistics *MediaStatistics  `protobuf:"bytes,5,opt,name=statistics,proto3" json:"statistics"`
	Token      *UploadTokenInfo  `protobuf:"bytes,6,opt,name=token,proto3" json:"token"`
	Callback   map[string]string `protobuf:"bytes,7,rep,name=callback,proto3" json:"callback" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *MediaResponse) Reset() {
	*x = MediaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mediaService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MediaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MediaResponse) ProtoMessage() {}

func (x *MediaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mediaService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MediaResponse.ProtoReflect.Descriptor instead.
func (*MediaResponse) Descriptor() ([]byte, []int) {
	return file_mediaService_proto_rawDescGZIP(), []int{3}
}

func (x *MediaResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *MediaResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *MediaResponse) GetEntity() *Media {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *MediaResponse) GetItems() []*Media {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *MediaResponse) GetStatistics() *MediaStatistics {
	if x != nil {
		return x.Statistics
	}
	return nil
}

func (x *MediaResponse) GetToken() *UploadTokenInfo {
	if x != nil {
		return x.Token
	}
	return nil
}

func (x *MediaResponse) GetCallback() map[string]string {
	if x != nil {
		return x.Callback
	}
	return nil
}

var File_mediaService_proto protoreflect.FileDescriptor

var file_mediaService_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x15, 0x62, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x43,
	0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xad, 0x04, 0x0a, 0x05, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73,
	0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x61, 0x77, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x61, 0x77, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09,
	0x73, 0x61, 0x76, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x73, 0x61, 0x76, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x6d,
	0x69, 0x6d, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6d, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x77, 0x69, 0x64,
	0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x15, 0x0a, 0x06,
	0x63, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x61,
	0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x12, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x24, 0x0a, 0x03, 0x63, 0x61, 0x74, 0x18, 0x15, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65,
	0x64, 0x69, 0x61, 0x43, 0x61, 0x74, 0x52, 0x03, 0x63, 0x61, 0x74, 0x12, 0x32, 0x0a, 0x0a, 0x6d,
	0x65, 0x64, 0x69, 0x61, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x22,
	0x94, 0x01, 0x0a, 0x0f, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74,
	0x69, 0x63, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x12, 0x2d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xea, 0x02, 0x0a, 0x0c, 0x4d, 0x65, 0x64, 0x69, 0x61,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x6f, 0x70, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x74, 0x6f, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67,
	0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x09, 0x64, 0x61, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x69,
	0x64, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x6d, 0x69, 0x6d, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x10, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6d, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x06,
	0x63, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x61,
	0x74, 0x49, 0x64, 0x22, 0x82, 0x03, 0x0a, 0x0d, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x06,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x52, 0x06, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x25, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x4d, 0x65, 0x64, 0x69, 0x61, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x39, 0x0a, 0x0a,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x64, 0x69,
	0x61, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x0a, 0x73, 0x74, 0x61,
	0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12, 0x2f, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x41, 0x0a, 0x08, 0x63, 0x61, 0x6c, 0x6c,
	0x62, 0x61, 0x63, 0x6b, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x08, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x1a, 0x3b, 0x0a, 0x0d, 0x43,
	0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0xa3, 0x04, 0x0a, 0x0c, 0x4d, 0x65, 0x64,
	0x69, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x06, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d,
	0x65, 0x64, 0x69, 0x61, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x4d, 0x65, 0x64, 0x69, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x34, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12,
	0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61,
	0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x64, 0x69,
	0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x03, 0x47,
	0x65, 0x74, 0x12, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65,
	0x64, 0x69, 0x61, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d,
	0x65, 0x64, 0x69, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b,
	0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x64, 0x69,
	0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x04, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d,
	0x65, 0x64, 0x69, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73,
	0x74, 0x69, 0x63, 0x73, 0x12, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x4d, 0x65, 0x64, 0x69, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0b, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0e, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x16, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x49,
	0x6e, 0x66, 0x6f, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d,
	0x65, 0x64, 0x69, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d,
	0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mediaService_proto_rawDescOnce sync.Once
	file_mediaService_proto_rawDescData = file_mediaService_proto_rawDesc
)

func file_mediaService_proto_rawDescGZIP() []byte {
	file_mediaService_proto_rawDescOnce.Do(func() {
		file_mediaService_proto_rawDescData = protoimpl.X.CompressGZIP(file_mediaService_proto_rawDescData)
	})
	return file_mediaService_proto_rawDescData
}

var file_mediaService_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_mediaService_proto_goTypes = []interface{}{
	(*Media)(nil),           // 0: services.Media
	(*MediaStatistics)(nil), // 1: services.MediaStatistics
	(*MediaRequest)(nil),    // 2: services.MediaRequest
	(*MediaResponse)(nil),   // 3: services.MediaResponse
	nil,                     // 4: services.MediaResponse.CallbackEntry
	(*MediaCat)(nil),        // 5: services.MediaCat
	(*MediaInfo)(nil),       // 6: services.MediaInfo
	(*common.Pager)(nil),    // 7: common.Pager
	(*UploadTokenInfo)(nil), // 8: services.UploadTokenInfo
	(*CallbackInfo)(nil),    // 9: services.CallbackInfo
}
var file_mediaService_proto_depIdxs = []int32{
	5,  // 0: services.Media.cat:type_name -> services.MediaCat
	6,  // 1: services.Media.media_info:type_name -> services.MediaInfo
	1,  // 2: services.MediaStatistics.data:type_name -> services.MediaStatistics
	7,  // 3: services.MediaResponse.pager:type_name -> common.Pager
	0,  // 4: services.MediaResponse.entity:type_name -> services.Media
	0,  // 5: services.MediaResponse.items:type_name -> services.Media
	1,  // 6: services.MediaResponse.statistics:type_name -> services.MediaStatistics
	8,  // 7: services.MediaResponse.token:type_name -> services.UploadTokenInfo
	4,  // 8: services.MediaResponse.callback:type_name -> services.MediaResponse.CallbackEntry
	0,  // 9: services.MediaService.Create:input_type -> services.Media
	0,  // 10: services.MediaService.Update:input_type -> services.Media
	0,  // 11: services.MediaService.Delete:input_type -> services.Media
	0,  // 12: services.MediaService.Get:input_type -> services.Media
	2,  // 13: services.MediaService.Search:input_type -> services.MediaRequest
	2,  // 14: services.MediaService.List:input_type -> services.MediaRequest
	2,  // 15: services.MediaService.Statistics:input_type -> services.MediaRequest
	2,  // 16: services.MediaService.UploadToken:input_type -> services.MediaRequest
	9,  // 17: services.MediaService.UploadCallback:input_type -> services.CallbackInfo
	3,  // 18: services.MediaService.Create:output_type -> services.MediaResponse
	3,  // 19: services.MediaService.Update:output_type -> services.MediaResponse
	3,  // 20: services.MediaService.Delete:output_type -> services.MediaResponse
	3,  // 21: services.MediaService.Get:output_type -> services.MediaResponse
	3,  // 22: services.MediaService.Search:output_type -> services.MediaResponse
	3,  // 23: services.MediaService.List:output_type -> services.MediaResponse
	3,  // 24: services.MediaService.Statistics:output_type -> services.MediaResponse
	3,  // 25: services.MediaService.UploadToken:output_type -> services.MediaResponse
	3,  // 26: services.MediaService.UploadCallback:output_type -> services.MediaResponse
	18, // [18:27] is the sub-list for method output_type
	9,  // [9:18] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_mediaService_proto_init() }
func file_mediaService_proto_init() {
	if File_mediaService_proto != nil {
		return
	}
	file_baseInfoService_proto_init()
	file_mediaCatService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_mediaService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Media); i {
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
		file_mediaService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MediaStatistics); i {
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
		file_mediaService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MediaRequest); i {
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
		file_mediaService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MediaResponse); i {
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
			RawDescriptor: file_mediaService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mediaService_proto_goTypes,
		DependencyIndexes: file_mediaService_proto_depIdxs,
		MessageInfos:      file_mediaService_proto_msgTypes,
	}.Build()
	File_mediaService_proto = out.File
	file_mediaService_proto_rawDesc = nil
	file_mediaService_proto_goTypes = nil
	file_mediaService_proto_depIdxs = nil
}