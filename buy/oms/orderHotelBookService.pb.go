// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: orderHotelBookService.proto

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

type HotelBookReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomPricePlanId int64  `protobuf:"varint,1,opt,name=room_price_plan_id,json=roomPricePlanId,proto3" json:"room_price_plan_id"`
	CheckinTime     string `protobuf:"bytes,2,opt,name=checkin_time,json=checkinTime,proto3" json:"checkin_time"`
	CheckoutTime    string `protobuf:"bytes,3,opt,name=checkout_time,json=checkoutTime,proto3" json:"checkout_time"`
	RoomNum         int32  `protobuf:"varint,4,opt,name=room_num,json=roomNum,proto3" json:"room_num"`
	IsExtraBed      bool   `protobuf:"varint,5,opt,name=is_extra_bed,json=isExtraBed,proto3" json:"is_extra_bed"`      // 是否加床
	BookRoomJson    string `protobuf:"bytes,6,opt,name=book_room_json,json=bookRoomJson,proto3" json:"book_room_json"` // 预留信息(json数据), 包括字段: name-字段名, value-值
	// @inject_tag: gorm:"-"
	BookCheckins []*HotelBookCheckin `protobuf:"bytes,7,rep,name=book_checkins,json=bookCheckins,proto3" json:"book_checkins" gorm:"-"`
}

func (x *HotelBookReq) Reset() {
	*x = HotelBookReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderHotelBookService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HotelBookReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HotelBookReq) ProtoMessage() {}

func (x *HotelBookReq) ProtoReflect() protoreflect.Message {
	mi := &file_orderHotelBookService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HotelBookReq.ProtoReflect.Descriptor instead.
func (*HotelBookReq) Descriptor() ([]byte, []int) {
	return file_orderHotelBookService_proto_rawDescGZIP(), []int{0}
}

func (x *HotelBookReq) GetRoomPricePlanId() int64 {
	if x != nil {
		return x.RoomPricePlanId
	}
	return 0
}

func (x *HotelBookReq) GetCheckinTime() string {
	if x != nil {
		return x.CheckinTime
	}
	return ""
}

func (x *HotelBookReq) GetCheckoutTime() string {
	if x != nil {
		return x.CheckoutTime
	}
	return ""
}

func (x *HotelBookReq) GetRoomNum() int32 {
	if x != nil {
		return x.RoomNum
	}
	return 0
}

func (x *HotelBookReq) GetIsExtraBed() bool {
	if x != nil {
		return x.IsExtraBed
	}
	return false
}

func (x *HotelBookReq) GetBookRoomJson() string {
	if x != nil {
		return x.BookRoomJson
	}
	return ""
}

func (x *HotelBookReq) GetBookCheckins() []*HotelBookCheckin {
	if x != nil {
		return x.BookCheckins
	}
	return nil
}

// 入住人信息
type HotelBookCheckin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GuestName         string `protobuf:"bytes,1,opt,name=guest_name,json=guestName,proto3" json:"guest_name"`
	GuestMobile       string `protobuf:"bytes,2,opt,name=guest_mobile,json=guestMobile,proto3" json:"guest_mobile"`
	GuestDocumentType int32  `protobuf:"varint,3,opt,name=guest_document_type,json=guestDocumentType,proto3" json:"guest_document_type"`
	GuestDocumentNo   string `protobuf:"bytes,4,opt,name=guest_document_no,json=guestDocumentNo,proto3" json:"guest_document_no"`
	// @inject_tag: gorm:"-"
	TogetherDetails []*HotelBookCheckin `protobuf:"bytes,5,rep,name=together_details,json=togetherDetails,proto3" json:"together_details" gorm:"-"`
}

func (x *HotelBookCheckin) Reset() {
	*x = HotelBookCheckin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderHotelBookService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HotelBookCheckin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HotelBookCheckin) ProtoMessage() {}

func (x *HotelBookCheckin) ProtoReflect() protoreflect.Message {
	mi := &file_orderHotelBookService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HotelBookCheckin.ProtoReflect.Descriptor instead.
func (*HotelBookCheckin) Descriptor() ([]byte, []int) {
	return file_orderHotelBookService_proto_rawDescGZIP(), []int{1}
}

func (x *HotelBookCheckin) GetGuestName() string {
	if x != nil {
		return x.GuestName
	}
	return ""
}

func (x *HotelBookCheckin) GetGuestMobile() string {
	if x != nil {
		return x.GuestMobile
	}
	return ""
}

func (x *HotelBookCheckin) GetGuestDocumentType() int32 {
	if x != nil {
		return x.GuestDocumentType
	}
	return 0
}

func (x *HotelBookCheckin) GetGuestDocumentNo() string {
	if x != nil {
		return x.GuestDocumentNo
	}
	return ""
}

func (x *HotelBookCheckin) GetTogetherDetails() []*HotelBookCheckin {
	if x != nil {
		return x.TogetherDetails
	}
	return nil
}

type HotelBook struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bookable             bool    `protobuf:"varint,1,opt,name=bookable,proto3" json:"bookable"`                                                         // 当前是否可预定
	MaxBookCount         int32   `protobuf:"varint,2,opt,name=max_book_count,json=maxBookCount,proto3" json:"max_book_count"`                           // 订单最多可预订房间数(0-不限)
	MinBookCount         int32   `protobuf:"varint,3,opt,name=min_book_count,json=minBookCount,proto3" json:"min_book_count"`                           // 订单最低预订房间数
	BookDay              int32   `protobuf:"varint,4,opt,name=book_day,json=bookDay,proto3" json:"book_day"`                                            // 最多可预定天数(0-不限)
	AheadDate            string  `protobuf:"bytes,5,opt,name=ahead_date,json=aheadDate,proto3" json:"ahead_date"`                                       // 可预定最晚的离店日期
	BookRoomJson         string  `protobuf:"bytes,6,opt,name=book_room_json,json=bookRoomJson,proto3" json:"book_room_json"`                            // 预留信息(json数据), 包括字段: name-字段名, tips_word-输入框内引导提示文案
	IsBookContentConfirm bool    `protobuf:"varint,7,opt,name=is_book_content_confirm,json=isBookContentConfirm,proto3" json:"is_book_content_confirm"` // 下单前是否需要进行预订须知确认
	BookContent          string  `protobuf:"bytes,8,opt,name=book_content,json=bookContent,proto3" json:"book_content"`                                 // 预定须知内容
	BookName             string  `protobuf:"bytes,9,opt,name=book_name,json=bookName,proto3" json:"book_name"`                                          // 房型名称
	IsOnlineDeposit      bool    `protobuf:"varint,10,opt,name=is_online_deposit,json=isOnlineDeposit,proto3" json:"is_online_deposit"`                 // 是否线上支付押金
	GuestNum             int32   `protobuf:"varint,11,opt,name=guest_num,json=guestNum,proto3" json:"guest_num"`                                        // 最多可住人数
	IsNeedHousemate      bool    `protobuf:"varint,12,opt,name=is_need_housemate,json=isNeedHousemate,proto3" json:"is_need_housemate"`                 // 是否填写同住人信息
	IsExtraBed           bool    `protobuf:"varint,13,opt,name=is_extra_bed,json=isExtraBed,proto3" json:"is_extra_bed"`                                // 可否加床
	ExtraBedName         string  `protobuf:"bytes,14,opt,name=extra_bed_name,json=extraBedName,proto3" json:"extra_bed_name"`                           // 加床自定义名称
	ExtraBedPrice        float32 `protobuf:"fixed32,15,opt,name=extra_bed_price,json=extraBedPrice,proto3" json:"extra_bed_price"`                      // 加床费
	BreakfastCount       int32   `protobuf:"varint,16,opt,name=breakfast_count,json=breakfastCount,proto3" json:"breakfast_count"`                      // 早餐份数
	GiftContent          string  `protobuf:"bytes,17,opt,name=gift_content,json=giftContent,proto3" json:"gift_content"`                                // 礼包内容
}

func (x *HotelBook) Reset() {
	*x = HotelBook{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderHotelBookService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HotelBook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HotelBook) ProtoMessage() {}

func (x *HotelBook) ProtoReflect() protoreflect.Message {
	mi := &file_orderHotelBookService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HotelBook.ProtoReflect.Descriptor instead.
func (*HotelBook) Descriptor() ([]byte, []int) {
	return file_orderHotelBookService_proto_rawDescGZIP(), []int{2}
}

func (x *HotelBook) GetBookable() bool {
	if x != nil {
		return x.Bookable
	}
	return false
}

func (x *HotelBook) GetMaxBookCount() int32 {
	if x != nil {
		return x.MaxBookCount
	}
	return 0
}

func (x *HotelBook) GetMinBookCount() int32 {
	if x != nil {
		return x.MinBookCount
	}
	return 0
}

func (x *HotelBook) GetBookDay() int32 {
	if x != nil {
		return x.BookDay
	}
	return 0
}

func (x *HotelBook) GetAheadDate() string {
	if x != nil {
		return x.AheadDate
	}
	return ""
}

func (x *HotelBook) GetBookRoomJson() string {
	if x != nil {
		return x.BookRoomJson
	}
	return ""
}

func (x *HotelBook) GetIsBookContentConfirm() bool {
	if x != nil {
		return x.IsBookContentConfirm
	}
	return false
}

func (x *HotelBook) GetBookContent() string {
	if x != nil {
		return x.BookContent
	}
	return ""
}

func (x *HotelBook) GetBookName() string {
	if x != nil {
		return x.BookName
	}
	return ""
}

func (x *HotelBook) GetIsOnlineDeposit() bool {
	if x != nil {
		return x.IsOnlineDeposit
	}
	return false
}

func (x *HotelBook) GetGuestNum() int32 {
	if x != nil {
		return x.GuestNum
	}
	return 0
}

func (x *HotelBook) GetIsNeedHousemate() bool {
	if x != nil {
		return x.IsNeedHousemate
	}
	return false
}

func (x *HotelBook) GetIsExtraBed() bool {
	if x != nil {
		return x.IsExtraBed
	}
	return false
}

func (x *HotelBook) GetExtraBedName() string {
	if x != nil {
		return x.ExtraBedName
	}
	return ""
}

func (x *HotelBook) GetExtraBedPrice() float32 {
	if x != nil {
		return x.ExtraBedPrice
	}
	return 0
}

func (x *HotelBook) GetBreakfastCount() int32 {
	if x != nil {
		return x.BreakfastCount
	}
	return 0
}

func (x *HotelBook) GetGiftContent() string {
	if x != nil {
		return x.GiftContent
	}
	return ""
}

var File_orderHotelBookService_proto protoreflect.FileDescriptor

var file_orderHotelBookService_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x42, 0x6f, 0x6f, 0x6b,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0xa7, 0x02, 0x0a, 0x0c, 0x48, 0x6f, 0x74, 0x65,
	0x6c, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x2b, 0x0a, 0x12, 0x72, 0x6f, 0x6f, 0x6d,
	0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x72, 0x6f, 0x6f, 0x6d, 0x50, 0x72, 0x69, 0x63, 0x65, 0x50,
	0x6c, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x6f, 0x75, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x4e, 0x75, 0x6d, 0x12, 0x20, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x65,
	0x78, 0x74, 0x72, 0x61, 0x5f, 0x62, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a,
	0x69, 0x73, 0x45, 0x78, 0x74, 0x72, 0x61, 0x42, 0x65, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x62, 0x6f,
	0x6f, 0x6b, 0x5f, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x6a, 0x73, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x62, 0x6f, 0x6f, 0x6b, 0x52, 0x6f, 0x6f, 0x6d, 0x4a, 0x73, 0x6f, 0x6e,
	0x12, 0x3f, 0x0a, 0x0d, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e,
	0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x42, 0x6f, 0x6f, 0x6b, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x69, 0x6e, 0x52, 0x0c, 0x62, 0x6f, 0x6f, 0x6b, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e,
	0x73, 0x22, 0xf7, 0x01, 0x0a, 0x10, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x42, 0x6f, 0x6f, 0x6b, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x75, 0x65, 0x73, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x75, 0x65, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x67, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x6d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x67, 0x75, 0x65,
	0x73, 0x74, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x67, 0x75, 0x65, 0x73,
	0x74, 0x5f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x67, 0x75, 0x65, 0x73, 0x74, 0x44, 0x6f, 0x63, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x67, 0x75, 0x65, 0x73,
	0x74, 0x5f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x6f, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x67, 0x75, 0x65, 0x73, 0x74, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x4e, 0x6f, 0x12, 0x45, 0x0a, 0x10, 0x74, 0x6f, 0x67, 0x65, 0x74, 0x68, 0x65, 0x72,
	0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x42,
	0x6f, 0x6f, 0x6b, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x52, 0x0f, 0x74, 0x6f, 0x67, 0x65,
	0x74, 0x68, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0xfb, 0x04, 0x0a, 0x09,
	0x48, 0x6f, 0x74, 0x65, 0x6c, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x6f, 0x6f,
	0x6b, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x62, 0x6f, 0x6f,
	0x6b, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x6d, 0x61, 0x78, 0x5f, 0x62, 0x6f, 0x6f,
	0x6b, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6d,
	0x61, 0x78, 0x42, 0x6f, 0x6f, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x6d,
	0x69, 0x6e, 0x5f, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0c, 0x6d, 0x69, 0x6e, 0x42, 0x6f, 0x6f, 0x6b, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x64, 0x61, 0x79, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x44, 0x61, 0x79, 0x12, 0x1d, 0x0a, 0x0a,
	0x61, 0x68, 0x65, 0x61, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x61, 0x68, 0x65, 0x61, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x62,
	0x6f, 0x6f, 0x6b, 0x5f, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x6a, 0x73, 0x6f, 0x6e, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x6f, 0x6f, 0x6b, 0x52, 0x6f, 0x6f, 0x6d, 0x4a, 0x73, 0x6f,
	0x6e, 0x12, 0x35, 0x0a, 0x17, 0x69, 0x73, 0x5f, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x14, 0x69, 0x73, 0x42, 0x6f, 0x6f, 0x6b, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6f, 0x6f, 0x6b,
	0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x62, 0x6f, 0x6f, 0x6b, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x62,
	0x6f, 0x6f, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x62, 0x6f, 0x6f, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x69, 0x73, 0x5f, 0x6f,
	0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0f, 0x69, 0x73, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x44, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x6e, 0x75,
	0x6d, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x67, 0x75, 0x65, 0x73, 0x74, 0x4e, 0x75,
	0x6d, 0x12, 0x2a, 0x0a, 0x11, 0x69, 0x73, 0x5f, 0x6e, 0x65, 0x65, 0x64, 0x5f, 0x68, 0x6f, 0x75,
	0x73, 0x65, 0x6d, 0x61, 0x74, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x69, 0x73,
	0x4e, 0x65, 0x65, 0x64, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x6d, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a,
	0x0c, 0x69, 0x73, 0x5f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x62, 0x65, 0x64, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x45, 0x78, 0x74, 0x72, 0x61, 0x42, 0x65, 0x64, 0x12,
	0x24, 0x0a, 0x0e, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x62, 0x65, 0x64, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x78, 0x74, 0x72, 0x61, 0x42, 0x65,
	0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x62,
	0x65, 0x64, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d,
	0x65, 0x78, 0x74, 0x72, 0x61, 0x42, 0x65, 0x64, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x27, 0x0a,
	0x0f, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x66, 0x61, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x10, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x66, 0x61, 0x73,
	0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x67, 0x69, 0x66, 0x74, 0x5f, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x67, 0x69,
	0x66, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_orderHotelBookService_proto_rawDescOnce sync.Once
	file_orderHotelBookService_proto_rawDescData = file_orderHotelBookService_proto_rawDesc
)

func file_orderHotelBookService_proto_rawDescGZIP() []byte {
	file_orderHotelBookService_proto_rawDescOnce.Do(func() {
		file_orderHotelBookService_proto_rawDescData = protoimpl.X.CompressGZIP(file_orderHotelBookService_proto_rawDescData)
	})
	return file_orderHotelBookService_proto_rawDescData
}

var file_orderHotelBookService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_orderHotelBookService_proto_goTypes = []interface{}{
	(*HotelBookReq)(nil),     // 0: services.HotelBookReq
	(*HotelBookCheckin)(nil), // 1: services.HotelBookCheckin
	(*HotelBook)(nil),        // 2: services.HotelBook
}
var file_orderHotelBookService_proto_depIdxs = []int32{
	1, // 0: services.HotelBookReq.book_checkins:type_name -> services.HotelBookCheckin
	1, // 1: services.HotelBookCheckin.together_details:type_name -> services.HotelBookCheckin
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_orderHotelBookService_proto_init() }
func file_orderHotelBookService_proto_init() {
	if File_orderHotelBookService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_orderHotelBookService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HotelBookReq); i {
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
		file_orderHotelBookService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HotelBookCheckin); i {
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
		file_orderHotelBookService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HotelBook); i {
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
			RawDescriptor: file_orderHotelBookService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_orderHotelBookService_proto_goTypes,
		DependencyIndexes: file_orderHotelBookService_proto_depIdxs,
		MessageInfos:      file_orderHotelBookService_proto_msgTypes,
	}.Build()
	File_orderHotelBookService_proto = out.File
	file_orderHotelBookService_proto_rawDesc = nil
	file_orderHotelBookService_proto_goTypes = nil
	file_orderHotelBookService_proto_depIdxs = nil
}
