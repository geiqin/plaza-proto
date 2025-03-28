// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: memberQrcodeService.proto

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

type MemberQrcode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`                                                            //ID
	MemberId           int64  `protobuf:"varint,2,opt,name=member_id,json=memberId,proto3" json:"member_id"`                                //客户ID
	MemberQrcode       string `protobuf:"bytes,3,opt,name=member_qrcode,json=memberQrcode,proto3" json:"member_qrcode"`                     //会员二维码
	WeixinMiniQrcode   string `protobuf:"bytes,4,opt,name=weixin_mini_qrcode,json=weixinMiniQrcode,proto3" json:"weixin_mini_qrcode"`       //微信小程序二维码
	AlipayMiniQrcode   string `protobuf:"bytes,5,opt,name=alipay_mini_qrcode,json=alipayMiniQrcode,proto3" json:"alipay_mini_qrcode"`       //支付宝小程序二维码
	DouyinMiniQrcode   string `protobuf:"bytes,6,opt,name=douyin_mini_qrcode,json=douyinMiniQrcode,proto3" json:"douyin_mini_qrcode"`       //抖音小程序二维码
	ToutiaoMiniQrcode  string `protobuf:"bytes,7,opt,name=toutiao_mini_qrcode,json=toutiaoMiniQrcode,proto3" json:"toutiao_mini_qrcode"`    //头条小程序二维码
	BaiduMiniQrcode    string `protobuf:"bytes,8,opt,name=baidu_mini_qrcode,json=baiduMiniQrcode,proto3" json:"baidu_mini_qrcode"`          //百度小程序二维码
	KuaishouMiniQrcode string `protobuf:"bytes,9,opt,name=kuaishou_mini_qrcode,json=kuaishouMiniQrcode,proto3" json:"kuaishou_mini_qrcode"` //快手小程序二维码
	CreatedAt          int64  `protobuf:"varint,10,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt          int64  `protobuf:"varint,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
}

func (x *MemberQrcode) Reset() {
	*x = MemberQrcode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memberQrcodeService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberQrcode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberQrcode) ProtoMessage() {}

func (x *MemberQrcode) ProtoReflect() protoreflect.Message {
	mi := &file_memberQrcodeService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberQrcode.ProtoReflect.Descriptor instead.
func (*MemberQrcode) Descriptor() ([]byte, []int) {
	return file_memberQrcodeService_proto_rawDescGZIP(), []int{0}
}

func (x *MemberQrcode) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MemberQrcode) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

func (x *MemberQrcode) GetMemberQrcode() string {
	if x != nil {
		return x.MemberQrcode
	}
	return ""
}

func (x *MemberQrcode) GetWeixinMiniQrcode() string {
	if x != nil {
		return x.WeixinMiniQrcode
	}
	return ""
}

func (x *MemberQrcode) GetAlipayMiniQrcode() string {
	if x != nil {
		return x.AlipayMiniQrcode
	}
	return ""
}

func (x *MemberQrcode) GetDouyinMiniQrcode() string {
	if x != nil {
		return x.DouyinMiniQrcode
	}
	return ""
}

func (x *MemberQrcode) GetToutiaoMiniQrcode() string {
	if x != nil {
		return x.ToutiaoMiniQrcode
	}
	return ""
}

func (x *MemberQrcode) GetBaiduMiniQrcode() string {
	if x != nil {
		return x.BaiduMiniQrcode
	}
	return ""
}

func (x *MemberQrcode) GetKuaishouMiniQrcode() string {
	if x != nil {
		return x.KuaishouMiniQrcode
	}
	return ""
}

func (x *MemberQrcode) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *MemberQrcode) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type MemberQrcodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int64  `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize int64  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Sorting  string `protobuf:"bytes,3,opt,name=sorting,proto3" json:"sorting"`
	//以下为自定义参数
	Id        int64   `protobuf:"varint,4,opt,name=id,proto3" json:"id"`
	Platform  string  `protobuf:"bytes,5,opt,name=platform,proto3" json:"platform"`
	MemberId  int64   `protobuf:"varint,6,opt,name=member_id,json=memberId,proto3" json:"member_id"`
	MemberIds []int64 `protobuf:"varint,7,rep,packed,name=member_ids,json=memberIds,proto3" json:"member_ids"`
	Ids       []int64 `protobuf:"varint,8,rep,packed,name=ids,proto3" json:"ids"`
}

func (x *MemberQrcodeRequest) Reset() {
	*x = MemberQrcodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memberQrcodeService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberQrcodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberQrcodeRequest) ProtoMessage() {}

func (x *MemberQrcodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_memberQrcodeService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberQrcodeRequest.ProtoReflect.Descriptor instead.
func (*MemberQrcodeRequest) Descriptor() ([]byte, []int) {
	return file_memberQrcodeService_proto_rawDescGZIP(), []int{1}
}

func (x *MemberQrcodeRequest) GetPaged() int64 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *MemberQrcodeRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *MemberQrcodeRequest) GetSorting() string {
	if x != nil {
		return x.Sorting
	}
	return ""
}

func (x *MemberQrcodeRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MemberQrcodeRequest) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *MemberQrcodeRequest) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

func (x *MemberQrcodeRequest) GetMemberIds() []int64 {
	if x != nil {
		return x.MemberIds
	}
	return nil
}

func (x *MemberQrcodeRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type MemberQrcodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *MemberQrcode   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager   `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*MemberQrcode `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Info   string          `protobuf:"bytes,4,opt,name=info,proto3" json:"info"`
}

func (x *MemberQrcodeResponse) Reset() {
	*x = MemberQrcodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memberQrcodeService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberQrcodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberQrcodeResponse) ProtoMessage() {}

func (x *MemberQrcodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_memberQrcodeService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberQrcodeResponse.ProtoReflect.Descriptor instead.
func (*MemberQrcodeResponse) Descriptor() ([]byte, []int) {
	return file_memberQrcodeService_proto_rawDescGZIP(), []int{2}
}

func (x *MemberQrcodeResponse) GetEntity() *MemberQrcode {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *MemberQrcodeResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *MemberQrcodeResponse) GetItems() []*MemberQrcode {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *MemberQrcodeResponse) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

var File_memberQrcodeService_proto protoreflect.FileDescriptor

var file_memberQrcodeService_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x51, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb6, 0x03, 0x0a, 0x0c, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x51, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x5f, 0x71, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x51, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x77,
	0x65, 0x69, 0x78, 0x69, 0x6e, 0x5f, 0x6d, 0x69, 0x6e, 0x69, 0x5f, 0x71, 0x72, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x77, 0x65, 0x69, 0x78, 0x69, 0x6e, 0x4d,
	0x69, 0x6e, 0x69, 0x51, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x61, 0x6c, 0x69,
	0x70, 0x61, 0x79, 0x5f, 0x6d, 0x69, 0x6e, 0x69, 0x5f, 0x71, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x61, 0x6c, 0x69, 0x70, 0x61, 0x79, 0x4d, 0x69, 0x6e,
	0x69, 0x51, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x64, 0x6f, 0x75, 0x79, 0x69,
	0x6e, 0x5f, 0x6d, 0x69, 0x6e, 0x69, 0x5f, 0x71, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x4d, 0x69, 0x6e, 0x69, 0x51,
	0x72, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x74, 0x6f, 0x75, 0x74, 0x69, 0x61, 0x6f,
	0x5f, 0x6d, 0x69, 0x6e, 0x69, 0x5f, 0x71, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x11, 0x74, 0x6f, 0x75, 0x74, 0x69, 0x61, 0x6f, 0x4d, 0x69, 0x6e, 0x69, 0x51,
	0x72, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x62, 0x61, 0x69, 0x64, 0x75, 0x5f, 0x6d,
	0x69, 0x6e, 0x69, 0x5f, 0x71, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x62, 0x61, 0x69, 0x64, 0x75, 0x4d, 0x69, 0x6e, 0x69, 0x51, 0x72, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x30, 0x0a, 0x14, 0x6b, 0x75, 0x61, 0x69, 0x73, 0x68, 0x6f, 0x75, 0x5f, 0x6d, 0x69,
	0x6e, 0x69, 0x5f, 0x71, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x12, 0x6b, 0x75, 0x61, 0x69, 0x73, 0x68, 0x6f, 0x75, 0x4d, 0x69, 0x6e, 0x69, 0x51, 0x72, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0xdc, 0x01, 0x0a, 0x13, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x51, 0x72, 0x63, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67,
	0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73,
	0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x07, 0x20,
	0x03, 0x28, 0x03, 0x52, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x73, 0x12, 0x10,
	0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73,
	0x22, 0xad, 0x01, 0x0a, 0x14, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x51, 0x72, 0x63, 0x6f, 0x64,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x51, 0x72, 0x63, 0x6f, 0x64,
	0x65, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x2c,
	0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x51,
	0x72, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f,
	0x32, 0x98, 0x01, 0x0a, 0x13, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x51, 0x72, 0x63, 0x6f, 0x64,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12,
	0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x51, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x1a, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x51, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x04, 0x53, 0x61, 0x76,
	0x65, 0x12, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x51, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x1a, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x51, 0x72, 0x63, 0x6f, 0x64,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2f,
	0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_memberQrcodeService_proto_rawDescOnce sync.Once
	file_memberQrcodeService_proto_rawDescData = file_memberQrcodeService_proto_rawDesc
)

func file_memberQrcodeService_proto_rawDescGZIP() []byte {
	file_memberQrcodeService_proto_rawDescOnce.Do(func() {
		file_memberQrcodeService_proto_rawDescData = protoimpl.X.CompressGZIP(file_memberQrcodeService_proto_rawDescData)
	})
	return file_memberQrcodeService_proto_rawDescData
}

var file_memberQrcodeService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_memberQrcodeService_proto_goTypes = []interface{}{
	(*MemberQrcode)(nil),         // 0: services.MemberQrcode
	(*MemberQrcodeRequest)(nil),  // 1: services.MemberQrcodeRequest
	(*MemberQrcodeResponse)(nil), // 2: services.MemberQrcodeResponse
	(*common.Pager)(nil),         // 3: common.Pager
}
var file_memberQrcodeService_proto_depIdxs = []int32{
	0, // 0: services.MemberQrcodeResponse.entity:type_name -> services.MemberQrcode
	3, // 1: services.MemberQrcodeResponse.pager:type_name -> common.Pager
	0, // 2: services.MemberQrcodeResponse.items:type_name -> services.MemberQrcode
	0, // 3: services.MemberQrcodeService.Get:input_type -> services.MemberQrcode
	0, // 4: services.MemberQrcodeService.Save:input_type -> services.MemberQrcode
	2, // 5: services.MemberQrcodeService.Get:output_type -> services.MemberQrcodeResponse
	2, // 6: services.MemberQrcodeService.Save:output_type -> services.MemberQrcodeResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_memberQrcodeService_proto_init() }
func file_memberQrcodeService_proto_init() {
	if File_memberQrcodeService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_memberQrcodeService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberQrcode); i {
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
		file_memberQrcodeService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberQrcodeRequest); i {
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
		file_memberQrcodeService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberQrcodeResponse); i {
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
			RawDescriptor: file_memberQrcodeService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_memberQrcodeService_proto_goTypes,
		DependencyIndexes: file_memberQrcodeService_proto_depIdxs,
		MessageInfos:      file_memberQrcodeService_proto_msgTypes,
	}.Build()
	File_memberQrcodeService_proto = out.File
	file_memberQrcodeService_proto_rawDesc = nil
	file_memberQrcodeService_proto_goTypes = nil
	file_memberQrcodeService_proto_depIdxs = nil
}
