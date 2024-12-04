// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: baseInfoService.proto

package services

import (
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

type TreeItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key      int64       `protobuf:"varint,1,opt,name=key,proto3" json:"key"`           //ID
	ParentId int64       `protobuf:"varint,2,opt,name=parentId,proto3" json:"parentId"` //父ID
	Title    string      `protobuf:"bytes,3,opt,name=title,proto3" json:"title"`        //标题
	Sort     int32       `protobuf:"varint,4,opt,name=sort,proto3" json:"sort"`         //排序
	Children []*TreeItem `protobuf:"bytes,5,rep,name=children,proto3" json:"children"`
}

func (x *TreeItem) Reset() {
	*x = TreeItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_baseInfoService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TreeItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TreeItem) ProtoMessage() {}

func (x *TreeItem) ProtoReflect() protoreflect.Message {
	mi := &file_baseInfoService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TreeItem.ProtoReflect.Descriptor instead.
func (*TreeItem) Descriptor() ([]byte, []int) {
	return file_baseInfoService_proto_rawDescGZIP(), []int{0}
}

func (x *TreeItem) GetKey() int64 {
	if x != nil {
		return x.Key
	}
	return 0
}

func (x *TreeItem) GetParentId() int64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *TreeItem) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *TreeItem) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *TreeItem) GetChildren() []*TreeItem {
	if x != nil {
		return x.Children
	}
	return nil
}

// 上传凭证信息
type UploadTokenInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token       string `protobuf:"bytes,1,opt,name=token,proto3" json:"token"`                                 //上传凭证
	ExpiredTime int64  `protobuf:"varint,2,opt,name=expired_time,json=expiredTime,proto3" json:"expired_time"` //失效时间
	Key         string `protobuf:"bytes,3,opt,name=key,proto3" json:"key"`                                     //储存KEY
}

func (x *UploadTokenInfo) Reset() {
	*x = UploadTokenInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_baseInfoService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadTokenInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadTokenInfo) ProtoMessage() {}

func (x *UploadTokenInfo) ProtoReflect() protoreflect.Message {
	mi := &file_baseInfoService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadTokenInfo.ProtoReflect.Descriptor instead.
func (*UploadTokenInfo) Descriptor() ([]byte, []int) {
	return file_baseInfoService_proto_rawDescGZIP(), []int{1}
}

func (x *UploadTokenInfo) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *UploadTokenInfo) GetExpiredTime() int64 {
	if x != nil {
		return x.ExpiredTime
	}
	return 0
}

func (x *UploadTokenInfo) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

// 上传回调数据
type CallbackInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bucket       string     `protobuf:"bytes,1,opt,name=bucket,proto3" json:"bucket"`
	EndUser      string     `protobuf:"bytes,2,opt,name=endUser,proto3" json:"endUser"`
	Key          string     `protobuf:"bytes,3,opt,name=key,proto3" json:"key"`
	Hash         string     `protobuf:"bytes,4,opt,name=hash,proto3" json:"hash"`
	FileName     string     `protobuf:"bytes,5,opt,name=fileName,proto3" json:"fileName"`
	Size         int64      `protobuf:"varint,6,opt,name=size,proto3" json:"size"`
	Etag         string     `protobuf:"bytes,7,opt,name=etag,proto3" json:"etag"`
	Ext          string     `protobuf:"bytes,8,opt,name=ext,proto3" json:"ext"`
	MimeType     string     `protobuf:"bytes,9,opt,name=mimeType,proto3" json:"mimeType"`
	PersistentId string     `protobuf:"bytes,10,opt,name=persistentId,proto3" json:"persistentId"`
	Uuid         string     `protobuf:"bytes,12,opt,name=uuid,proto3" json:"uuid"`
	CatId        string     `protobuf:"bytes,14,opt,name=catId,proto3" json:"catId"`
	Channel      string     `protobuf:"bytes,15,opt,name=channel,proto3" json:"channel"`
	UserType     string     `protobuf:"bytes,16,opt,name=userType,proto3" json:"userType"`
	UserId       int64      `protobuf:"varint,17,opt,name=userId,proto3" json:"userId"`
	Storage      string     `protobuf:"bytes,18,opt,name=storage,proto3" json:"storage"`
	IsOverwrite  bool       `protobuf:"varint,19,opt,name=is_overwrite,json=isOverwrite,proto3" json:"is_overwrite"`
	ImageInfo    *ImageInfo `protobuf:"bytes,20,opt,name=imageInfo,proto3" json:"imageInfo"`
	AvInfo       *AvInfo    `protobuf:"bytes,21,opt,name=avInfo,proto3" json:"avInfo"`
}

func (x *CallbackInfo) Reset() {
	*x = CallbackInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_baseInfoService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallbackInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallbackInfo) ProtoMessage() {}

func (x *CallbackInfo) ProtoReflect() protoreflect.Message {
	mi := &file_baseInfoService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallbackInfo.ProtoReflect.Descriptor instead.
func (*CallbackInfo) Descriptor() ([]byte, []int) {
	return file_baseInfoService_proto_rawDescGZIP(), []int{2}
}

func (x *CallbackInfo) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *CallbackInfo) GetEndUser() string {
	if x != nil {
		return x.EndUser
	}
	return ""
}

func (x *CallbackInfo) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *CallbackInfo) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *CallbackInfo) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *CallbackInfo) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *CallbackInfo) GetEtag() string {
	if x != nil {
		return x.Etag
	}
	return ""
}

func (x *CallbackInfo) GetExt() string {
	if x != nil {
		return x.Ext
	}
	return ""
}

func (x *CallbackInfo) GetMimeType() string {
	if x != nil {
		return x.MimeType
	}
	return ""
}

func (x *CallbackInfo) GetPersistentId() string {
	if x != nil {
		return x.PersistentId
	}
	return ""
}

func (x *CallbackInfo) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *CallbackInfo) GetCatId() string {
	if x != nil {
		return x.CatId
	}
	return ""
}

func (x *CallbackInfo) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *CallbackInfo) GetUserType() string {
	if x != nil {
		return x.UserType
	}
	return ""
}

func (x *CallbackInfo) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CallbackInfo) GetStorage() string {
	if x != nil {
		return x.Storage
	}
	return ""
}

func (x *CallbackInfo) GetIsOverwrite() bool {
	if x != nil {
		return x.IsOverwrite
	}
	return false
}

func (x *CallbackInfo) GetImageInfo() *ImageInfo {
	if x != nil {
		return x.ImageInfo
	}
	return nil
}

func (x *CallbackInfo) GetAvInfo() *AvInfo {
	if x != nil {
		return x.AvInfo
	}
	return nil
}

// 媒体信息
type MediaInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image *ImageInfo `protobuf:"bytes,1,opt,name=image,proto3" json:"image"`
	Video *VideoInfo `protobuf:"bytes,2,opt,name=video,proto3" json:"video"`
	Audio *AudioInfo `protobuf:"bytes,3,opt,name=audio,proto3" json:"audio"`
}

func (x *MediaInfo) Reset() {
	*x = MediaInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_baseInfoService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MediaInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MediaInfo) ProtoMessage() {}

func (x *MediaInfo) ProtoReflect() protoreflect.Message {
	mi := &file_baseInfoService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MediaInfo.ProtoReflect.Descriptor instead.
func (*MediaInfo) Descriptor() ([]byte, []int) {
	return file_baseInfoService_proto_rawDescGZIP(), []int{3}
}

func (x *MediaInfo) GetImage() *ImageInfo {
	if x != nil {
		return x.Image
	}
	return nil
}

func (x *MediaInfo) GetVideo() *VideoInfo {
	if x != nil {
		return x.Video
	}
	return nil
}

func (x *MediaInfo) GetAudio() *AudioInfo {
	if x != nil {
		return x.Audio
	}
	return nil
}

// 回调图片信息
type ImageInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Format string `protobuf:"bytes,1,opt,name=format,proto3" json:"format"`
	Width  int64  `protobuf:"varint,2,opt,name=width,proto3" json:"width"`
	Height int64  `protobuf:"varint,3,opt,name=height,proto3" json:"height"`
}

func (x *ImageInfo) Reset() {
	*x = ImageInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_baseInfoService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageInfo) ProtoMessage() {}

func (x *ImageInfo) ProtoReflect() protoreflect.Message {
	mi := &file_baseInfoService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageInfo.ProtoReflect.Descriptor instead.
func (*ImageInfo) Descriptor() ([]byte, []int) {
	return file_baseInfoService_proto_rawDescGZIP(), []int{4}
}

func (x *ImageInfo) GetFormat() string {
	if x != nil {
		return x.Format
	}
	return ""
}

func (x *ImageInfo) GetWidth() int64 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *ImageInfo) GetHeight() int64 {
	if x != nil {
		return x.Height
	}
	return 0
}

// 回调音视频信息
type AvInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Audio  *AudioInfo  `protobuf:"bytes,1,opt,name=audio,proto3" json:"audio"`
	Video  *VideoInfo  `protobuf:"bytes,2,opt,name=video,proto3" json:"video"`
	Format *FormatInfo `protobuf:"bytes,3,opt,name=format,proto3" json:"format"`
}

func (x *AvInfo) Reset() {
	*x = AvInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_baseInfoService_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvInfo) ProtoMessage() {}

func (x *AvInfo) ProtoReflect() protoreflect.Message {
	mi := &file_baseInfoService_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvInfo.ProtoReflect.Descriptor instead.
func (*AvInfo) Descriptor() ([]byte, []int) {
	return file_baseInfoService_proto_rawDescGZIP(), []int{5}
}

func (x *AvInfo) GetAudio() *AudioInfo {
	if x != nil {
		return x.Audio
	}
	return nil
}

func (x *AvInfo) GetVideo() *VideoInfo {
	if x != nil {
		return x.Video
	}
	return nil
}

func (x *AvInfo) GetFormat() *FormatInfo {
	if x != nil {
		return x.Format
	}
	return nil
}

type AudioInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BitRate   string `protobuf:"bytes,1,opt,name=bit_rate,json=bitRate,proto3" json:"bit_rate"`       //BIT率
	CodecName string `protobuf:"bytes,2,opt,name=codec_name,json=codecName,proto3" json:"codec_name"` //编码名称
	CodecType string `protobuf:"bytes,3,opt,name=codec_type,json=codecType,proto3" json:"codec_type"` //编码类型
	Duration  string `protobuf:"bytes,4,opt,name=duration,proto3" json:"duration"`                    //时长
	StartTime string `protobuf:"bytes,8,opt,name=start_time,json=startTime,proto3" json:"start_time"` //视频开始时间
	Channels  int32  `protobuf:"varint,9,opt,name=channels,proto3" json:"channels"`                   //音频通道数
}

func (x *AudioInfo) Reset() {
	*x = AudioInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_baseInfoService_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AudioInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AudioInfo) ProtoMessage() {}

func (x *AudioInfo) ProtoReflect() protoreflect.Message {
	mi := &file_baseInfoService_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AudioInfo.ProtoReflect.Descriptor instead.
func (*AudioInfo) Descriptor() ([]byte, []int) {
	return file_baseInfoService_proto_rawDescGZIP(), []int{6}
}

func (x *AudioInfo) GetBitRate() string {
	if x != nil {
		return x.BitRate
	}
	return ""
}

func (x *AudioInfo) GetCodecName() string {
	if x != nil {
		return x.CodecName
	}
	return ""
}

func (x *AudioInfo) GetCodecType() string {
	if x != nil {
		return x.CodecType
	}
	return ""
}

func (x *AudioInfo) GetDuration() string {
	if x != nil {
		return x.Duration
	}
	return ""
}

func (x *AudioInfo) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *AudioInfo) GetChannels() int32 {
	if x != nil {
		return x.Channels
	}
	return 0
}

type VideoInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BitRate            string `protobuf:"bytes,1,opt,name=bit_rate,json=bitRate,proto3" json:"bit_rate"`       //BIT率
	CodecName          string `protobuf:"bytes,2,opt,name=codec_name,json=codecName,proto3" json:"codec_name"` //编码名称
	CodecType          string `protobuf:"bytes,3,opt,name=codec_type,json=codecType,proto3" json:"codec_type"` //编码类型
	Duration           string `protobuf:"bytes,4,opt,name=duration,proto3" json:"duration"`                    //时长
	DisplayAspectRatio string `protobuf:"bytes,5,opt,name=display_aspect_ratio,json=displayAspectRatio,proto3" json:"display_aspect_ratio"`
	Height             int64  `protobuf:"varint,6,opt,name=height,proto3" json:"height"`                       //视频高
	Width              int64  `protobuf:"varint,7,opt,name=width,proto3" json:"width"`                         //视频宽
	StartTime          string `protobuf:"bytes,8,opt,name=start_time,json=startTime,proto3" json:"start_time"` //视频开始时间
	Channels           int32  `protobuf:"varint,9,opt,name=channels,proto3" json:"channels"`                   //音频通道数
}

func (x *VideoInfo) Reset() {
	*x = VideoInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_baseInfoService_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoInfo) ProtoMessage() {}

func (x *VideoInfo) ProtoReflect() protoreflect.Message {
	mi := &file_baseInfoService_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoInfo.ProtoReflect.Descriptor instead.
func (*VideoInfo) Descriptor() ([]byte, []int) {
	return file_baseInfoService_proto_rawDescGZIP(), []int{7}
}

func (x *VideoInfo) GetBitRate() string {
	if x != nil {
		return x.BitRate
	}
	return ""
}

func (x *VideoInfo) GetCodecName() string {
	if x != nil {
		return x.CodecName
	}
	return ""
}

func (x *VideoInfo) GetCodecType() string {
	if x != nil {
		return x.CodecType
	}
	return ""
}

func (x *VideoInfo) GetDuration() string {
	if x != nil {
		return x.Duration
	}
	return ""
}

func (x *VideoInfo) GetDisplayAspectRatio() string {
	if x != nil {
		return x.DisplayAspectRatio
	}
	return ""
}

func (x *VideoInfo) GetHeight() int64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *VideoInfo) GetWidth() int64 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *VideoInfo) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *VideoInfo) GetChannels() int32 {
	if x != nil {
		return x.Channels
	}
	return 0
}

type FormatInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BitRate        string `protobuf:"bytes,1,opt,name=bit_rate,json=bitRate,proto3" json:"bit_rate"` //BIT率
	Duration       string `protobuf:"bytes,2,opt,name=duration,proto3" json:"duration"`              //时长
	FormatLongName string `protobuf:"bytes,3,opt,name=format_long_name,json=formatLongName,proto3" json:"format_long_name"`
	FormatName     string `protobuf:"bytes,4,opt,name=format_name,json=formatName,proto3" json:"format_name"`
	NbStreams      string `protobuf:"bytes,5,opt,name=nb_streams,json=nbStreams,proto3" json:"nb_streams"`
	Size           string `protobuf:"bytes,6,opt,name=size,proto3" json:"size"`
	StartTime      string `protobuf:"bytes,7,opt,name=start_time,json=startTime,proto3" json:"start_time"`
}

func (x *FormatInfo) Reset() {
	*x = FormatInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_baseInfoService_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FormatInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FormatInfo) ProtoMessage() {}

func (x *FormatInfo) ProtoReflect() protoreflect.Message {
	mi := &file_baseInfoService_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FormatInfo.ProtoReflect.Descriptor instead.
func (*FormatInfo) Descriptor() ([]byte, []int) {
	return file_baseInfoService_proto_rawDescGZIP(), []int{8}
}

func (x *FormatInfo) GetBitRate() string {
	if x != nil {
		return x.BitRate
	}
	return ""
}

func (x *FormatInfo) GetDuration() string {
	if x != nil {
		return x.Duration
	}
	return ""
}

func (x *FormatInfo) GetFormatLongName() string {
	if x != nil {
		return x.FormatLongName
	}
	return ""
}

func (x *FormatInfo) GetFormatName() string {
	if x != nil {
		return x.FormatName
	}
	return ""
}

func (x *FormatInfo) GetNbStreams() string {
	if x != nil {
		return x.NbStreams
	}
	return ""
}

func (x *FormatInfo) GetSize() string {
	if x != nil {
		return x.Size
	}
	return ""
}

func (x *FormatInfo) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

var File_baseInfoService_proto protoreflect.FileDescriptor

var file_baseInfoService_proto_rawDesc = []byte{
	0x0a, 0x15, 0x62, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x22, 0x92, 0x01, 0x0a, 0x08, 0x54, 0x72, 0x65, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x2e, 0x0a, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72,
	0x65, 0x6e, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x54, 0x72, 0x65, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x08, 0x63, 0x68,
	0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x22, 0x5c, 0x0a, 0x0f, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x21, 0x0a, 0x0c, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x22, 0x8e, 0x04, 0x0a, 0x0c, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63,
	0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x65, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x65, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73,
	0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x1a, 0x0a,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x65, 0x74, 0x61, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x65, 0x74, 0x61,
	0x67, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x78, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x65, 0x78, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x22, 0x0a, 0x0c, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x61, 0x74, 0x49, 0x64,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x61, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x11, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x6f, 0x76, 0x65, 0x72,
	0x77, 0x72, 0x69, 0x74, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x4f,
	0x76, 0x65, 0x72, 0x77, 0x72, 0x69, 0x74, 0x65, 0x12, 0x31, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x28, 0x0a, 0x06, 0x61,
	0x76, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x76, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x61,
	0x76, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x8c, 0x01, 0x0a, 0x09, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x29, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x29,
	0x0a, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x29, 0x0a, 0x05, 0x61, 0x75, 0x64,
	0x69, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x61,
	0x75, 0x64, 0x69, 0x6f, 0x22, 0x51, 0x0a, 0x09, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64,
	0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12,
	0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x8c, 0x01, 0x0a, 0x06, 0x41, 0x76, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x29, 0x0a, 0x05, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x75, 0x64,
	0x69, 0x6f, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x12, 0x29, 0x0a,
	0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x2c, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x22, 0xbb, 0x01, 0x0a, 0x09, 0x41, 0x75, 0x64, 0x69, 0x6f,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x69, 0x74, 0x5f, 0x72, 0x61, 0x74, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x69, 0x74, 0x52, 0x61, 0x74, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x73, 0x22, 0x9b, 0x02, 0x0a, 0x09, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x69, 0x74, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x69, 0x74, 0x52, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x6f, 0x64, 0x65, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x14, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x5f, 0x61, 0x73, 0x70, 0x65, 0x63, 0x74, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x41, 0x73,
	0x70, 0x65, 0x63, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x73, 0x22, 0xe0, 0x01, 0x0a, 0x0a, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x69, 0x74, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x69, 0x74, 0x52, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x10, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x5f, 0x6c, 0x6f, 0x6e, 0x67, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x4c, 0x6f, 0x6e, 0x67, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x62, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x62, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_baseInfoService_proto_rawDescOnce sync.Once
	file_baseInfoService_proto_rawDescData = file_baseInfoService_proto_rawDesc
)

func file_baseInfoService_proto_rawDescGZIP() []byte {
	file_baseInfoService_proto_rawDescOnce.Do(func() {
		file_baseInfoService_proto_rawDescData = protoimpl.X.CompressGZIP(file_baseInfoService_proto_rawDescData)
	})
	return file_baseInfoService_proto_rawDescData
}

var file_baseInfoService_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_baseInfoService_proto_goTypes = []interface{}{
	(*TreeItem)(nil),        // 0: services.TreeItem
	(*UploadTokenInfo)(nil), // 1: services.UploadTokenInfo
	(*CallbackInfo)(nil),    // 2: services.CallbackInfo
	(*MediaInfo)(nil),       // 3: services.MediaInfo
	(*ImageInfo)(nil),       // 4: services.ImageInfo
	(*AvInfo)(nil),          // 5: services.AvInfo
	(*AudioInfo)(nil),       // 6: services.AudioInfo
	(*VideoInfo)(nil),       // 7: services.VideoInfo
	(*FormatInfo)(nil),      // 8: services.FormatInfo
}
var file_baseInfoService_proto_depIdxs = []int32{
	0, // 0: services.TreeItem.children:type_name -> services.TreeItem
	4, // 1: services.CallbackInfo.imageInfo:type_name -> services.ImageInfo
	5, // 2: services.CallbackInfo.avInfo:type_name -> services.AvInfo
	4, // 3: services.MediaInfo.image:type_name -> services.ImageInfo
	7, // 4: services.MediaInfo.video:type_name -> services.VideoInfo
	6, // 5: services.MediaInfo.audio:type_name -> services.AudioInfo
	6, // 6: services.AvInfo.audio:type_name -> services.AudioInfo
	7, // 7: services.AvInfo.video:type_name -> services.VideoInfo
	8, // 8: services.AvInfo.format:type_name -> services.FormatInfo
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_baseInfoService_proto_init() }
func file_baseInfoService_proto_init() {
	if File_baseInfoService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_baseInfoService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TreeItem); i {
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
		file_baseInfoService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadTokenInfo); i {
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
		file_baseInfoService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallbackInfo); i {
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
		file_baseInfoService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MediaInfo); i {
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
		file_baseInfoService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageInfo); i {
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
		file_baseInfoService_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvInfo); i {
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
		file_baseInfoService_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AudioInfo); i {
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
		file_baseInfoService_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoInfo); i {
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
		file_baseInfoService_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FormatInfo); i {
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
			RawDescriptor: file_baseInfoService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_baseInfoService_proto_goTypes,
		DependencyIndexes: file_baseInfoService_proto_depIdxs,
		MessageInfos:      file_baseInfoService_proto_msgTypes,
	}.Build()
	File_baseInfoService_proto = out.File
	file_baseInfoService_proto_rawDesc = nil
	file_baseInfoService_proto_goTypes = nil
	file_baseInfoService_proto_depIdxs = nil
}
