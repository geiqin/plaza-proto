// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: userService.proto

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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id"`                                                            //ID
	Type              string  `protobuf:"bytes,2,opt,name=type,proto3" json:"type"`                                                         //类型：free免费等级，1付费等级
	Name              string  `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`                                                         //等级名称
	BgMode            string  `protobuf:"bytes,4,opt,name=bg_mode,json=bgMode,proto3" json:"bg_mode"`                                       //背景模式：0背景色，1背景图
	BgColor           string  `protobuf:"bytes,5,opt,name=bg_color,json=bgColor,proto3" json:"bg_color"`                                    //卡面背景颜色
	BgImageUrl        string  `protobuf:"bytes,6,opt,name=bg_image_url,json=bgImageUrl,proto3" json:"bg_image_url"`                         //背景图URL
	IconUrl           string  `protobuf:"bytes,7,opt,name=icon_url,json=iconUrl,proto3" json:"icon_url"`                                    //图标URL
	Grade             int32   `protobuf:"varint,8,opt,name=grade,proto3" json:"grade"`                                                      //会员等级级别
	GrowthValue       int32   `protobuf:"varint,9,opt,name=growth_value,json=growthValue,proto3" json:"growth_value"`                       //需要成长值【免费会员】
	Description       string  `protobuf:"bytes,10,opt,name=description,proto3" json:"description"`                                          //等级描述/使用须知
	IsFreeShipping    string  `protobuf:"bytes,11,opt,name=is_free_shipping,json=isFreeShipping,proto3" json:"is_free_shipping"`            //是否会员包邮(0否，1是)【会员权益】
	IsDiscount        string  `protobuf:"bytes,12,opt,name=is_discount,json=isDiscount,proto3" json:"is_discount"`                          //是否会员折扣(0否，1是)【会员权益】
	IsPointFeedback   string  `protobuf:"bytes,13,opt,name=is_point_feedback,json=isPointFeedback,proto3" json:"is_point_feedback"`         //是否积分回馈(0否，1是)【会员权益】
	DiscountRate      float32 `protobuf:"fixed32,14,opt,name=discount_rate,json=discountRate,proto3" json:"discount_rate"`                  //会员折扣率(折扣率取0到10之间)
	PointFeedbackRate float32 `protobuf:"fixed32,15,opt,name=point_feedback_rate,json=pointFeedbackRate,proto3" json:"point_feedback_rate"` //积分回馈倍率(积分倍率不能大于 50)
	IsAutoUpgrade     string  `protobuf:"bytes,16,opt,name=is_auto_upgrade,json=isAutoUpgrade,proto3" json:"is_auto_upgrade"`               //是否开启自动升级【免费会员】
	IsAutoRenewal     string  `protobuf:"bytes,17,opt,name=is_auto_renewal,json=isAutoRenewal,proto3" json:"is_auto_renewal"`               //是否开启自动续费【付费会员】
	Status            string  `protobuf:"bytes,18,opt,name=status,proto3" json:"status"`                                                    //状态：启用，禁用
	CreatedAt         int64   `protobuf:"varint,19,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt         int64   `protobuf:"varint,20,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_userService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_userService_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetBgMode() string {
	if x != nil {
		return x.BgMode
	}
	return ""
}

func (x *User) GetBgColor() string {
	if x != nil {
		return x.BgColor
	}
	return ""
}

func (x *User) GetBgImageUrl() string {
	if x != nil {
		return x.BgImageUrl
	}
	return ""
}

func (x *User) GetIconUrl() string {
	if x != nil {
		return x.IconUrl
	}
	return ""
}

func (x *User) GetGrade() int32 {
	if x != nil {
		return x.Grade
	}
	return 0
}

func (x *User) GetGrowthValue() int32 {
	if x != nil {
		return x.GrowthValue
	}
	return 0
}

func (x *User) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *User) GetIsFreeShipping() string {
	if x != nil {
		return x.IsFreeShipping
	}
	return ""
}

func (x *User) GetIsDiscount() string {
	if x != nil {
		return x.IsDiscount
	}
	return ""
}

func (x *User) GetIsPointFeedback() string {
	if x != nil {
		return x.IsPointFeedback
	}
	return ""
}

func (x *User) GetDiscountRate() float32 {
	if x != nil {
		return x.DiscountRate
	}
	return 0
}

func (x *User) GetPointFeedbackRate() float32 {
	if x != nil {
		return x.PointFeedbackRate
	}
	return 0
}

func (x *User) GetIsAutoUpgrade() string {
	if x != nil {
		return x.IsAutoUpgrade
	}
	return ""
}

func (x *User) GetIsAutoRenewal() string {
	if x != nil {
		return x.IsAutoRenewal
	}
	return ""
}

func (x *User) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *User) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *User) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type UserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId     string            `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id"`
	ClientSecret string            `protobuf:"bytes,2,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret"`
	Type         string            `protobuf:"bytes,3,opt,name=type,proto3" json:"type"` //类型： username/mobile/email
	Username     string            `protobuf:"bytes,4,opt,name=username,proto3" json:"username"`
	Password     string            `protobuf:"bytes,5,opt,name=password,proto3" json:"password"`
	Code         string            `protobuf:"bytes,6,opt,name=code,proto3" json:"code"`
	Mobile       string            `protobuf:"bytes,7,opt,name=mobile,proto3" json:"mobile"`
	Email        string            `protobuf:"bytes,8,opt,name=email,proto3" json:"email"`
	RedirectUri  string            `protobuf:"bytes,9,opt,name=redirect_uri,json=redirectUri,proto3" json:"redirect_uri"`
	RefreshToken string            `protobuf:"bytes,10,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token"`
	StoreId      string            `protobuf:"bytes,11,opt,name=store_id,json=storeId,proto3" json:"store_id"`
	StoreName    string            `protobuf:"bytes,12,opt,name=store_name,json=storeName,proto3" json:"store_name"`
	Scope        string            `protobuf:"bytes,13,opt,name=scope,proto3" json:"scope"`
	State        string            `protobuf:"bytes,14,opt,name=state,proto3" json:"state"`
	GrantType    string            `protobuf:"bytes,15,opt,name=grant_type,json=grantType,proto3" json:"grant_type"`
	ReferrerId   int64             `protobuf:"varint,16,opt,name=referrer_id,json=referrerId,proto3" json:"referrer_id"`
	Extends      map[string]string `protobuf:"bytes,17,rep,name=extends,proto3" json:"extends" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *UserRequest) Reset() {
	*x = UserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRequest) ProtoMessage() {}

func (x *UserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_userService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRequest.ProtoReflect.Descriptor instead.
func (*UserRequest) Descriptor() ([]byte, []int) {
	return file_userService_proto_rawDescGZIP(), []int{1}
}

func (x *UserRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *UserRequest) GetClientSecret() string {
	if x != nil {
		return x.ClientSecret
	}
	return ""
}

func (x *UserRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *UserRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *UserRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *UserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserRequest) GetRedirectUri() string {
	if x != nil {
		return x.RedirectUri
	}
	return ""
}

func (x *UserRequest) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *UserRequest) GetStoreId() string {
	if x != nil {
		return x.StoreId
	}
	return ""
}

func (x *UserRequest) GetStoreName() string {
	if x != nil {
		return x.StoreName
	}
	return ""
}

func (x *UserRequest) GetScope() string {
	if x != nil {
		return x.Scope
	}
	return ""
}

func (x *UserRequest) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *UserRequest) GetGrantType() string {
	if x != nil {
		return x.GrantType
	}
	return ""
}

func (x *UserRequest) GetReferrerId() int64 {
	if x != nil {
		return x.ReferrerId
	}
	return 0
}

func (x *UserRequest) GetExtends() map[string]string {
	if x != nil {
		return x.Extends
	}
	return nil
}

// 用户切换请求【必须是服务调用】
type SwitchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mode        string            `protobuf:"bytes,1,opt,name=mode,proto3" json:"mode"`                                                                                                //授权模式
	SessionKey  string            `protobuf:"bytes,2,opt,name=session_key,json=sessionKey,proto3" json:"session_key"`                                                                  //会话Key
	UserId      int64             `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id"`                                                                             //用户ID
	StoreId     int64             `protobuf:"varint,4,opt,name=store_id,json=storeId,proto3" json:"store_id"`                                                                          //平台ID
	RealstoreId int64             `protobuf:"varint,5,opt,name=realstore_id,json=realstoreId,proto3" json:"realstore_id"`                                                              //多门店ID
	ShopId      int64             `protobuf:"varint,6,opt,name=shop_id,json=shopId,proto3" json:"shop_id"`                                                                             //多商户ID
	PlatformId  int64             `protobuf:"varint,7,opt,name=platform_id,json=platformId,proto3" json:"platform_id"`                                                                 //平台ID
	ClientId    string            `protobuf:"bytes,8,opt,name=client_id,json=clientId,proto3" json:"client_id"`                                                                        //ClientID
	Nickname    string            `protobuf:"bytes,9,opt,name=nickname,proto3" json:"nickname"`                                                                                        //用户昵称
	Extends     map[string]string `protobuf:"bytes,10,rep,name=extends,proto3" json:"extends" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` //附加信息
}

func (x *SwitchRequest) Reset() {
	*x = SwitchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SwitchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SwitchRequest) ProtoMessage() {}

func (x *SwitchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_userService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SwitchRequest.ProtoReflect.Descriptor instead.
func (*SwitchRequest) Descriptor() ([]byte, []int) {
	return file_userService_proto_rawDescGZIP(), []int{2}
}

func (x *SwitchRequest) GetMode() string {
	if x != nil {
		return x.Mode
	}
	return ""
}

func (x *SwitchRequest) GetSessionKey() string {
	if x != nil {
		return x.SessionKey
	}
	return ""
}

func (x *SwitchRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *SwitchRequest) GetStoreId() int64 {
	if x != nil {
		return x.StoreId
	}
	return 0
}

func (x *SwitchRequest) GetRealstoreId() int64 {
	if x != nil {
		return x.RealstoreId
	}
	return 0
}

func (x *SwitchRequest) GetShopId() int64 {
	if x != nil {
		return x.ShopId
	}
	return 0
}

func (x *SwitchRequest) GetPlatformId() int64 {
	if x != nil {
		return x.PlatformId
	}
	return 0
}

func (x *SwitchRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *SwitchRequest) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *SwitchRequest) GetExtends() map[string]string {
	if x != nil {
		return x.Extends
	}
	return nil
}

type UserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity map[string]string `protobuf:"bytes,1,rep,name=entity,proto3" json:"entity" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Pager  *common.Pager     `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*User           `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Msg    string            `protobuf:"bytes,4,opt,name=msg,proto3" json:"msg"`
}

func (x *UserResponse) Reset() {
	*x = UserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserResponse) ProtoMessage() {}

func (x *UserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_userService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserResponse.ProtoReflect.Descriptor instead.
func (*UserResponse) Descriptor() ([]byte, []int) {
	return file_userService_proto_rawDescGZIP(), []int{3}
}

func (x *UserResponse) GetEntity() map[string]string {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *UserResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *UserResponse) GetItems() []*User {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *UserResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_userService_proto protoreflect.FileDescriptor

var file_userService_proto_rawDesc = []byte{
	0x0a, 0x11, 0x75, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x15, 0x62, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfc, 0x04, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x67, 0x5f, 0x6d,
	0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x67, 0x4d, 0x6f, 0x64,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x67, 0x5f, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x67, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x0c,
	0x62, 0x67, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x62, 0x67, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x19,
	0x0a, 0x08, 0x69, 0x63, 0x6f, 0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x69, 0x63, 0x6f, 0x6e, 0x55, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x61,
	0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x67, 0x72, 0x61, 0x64, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x67, 0x72, 0x6f, 0x77, 0x74, 0x68, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x67, 0x72, 0x6f, 0x77, 0x74, 0x68, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x10, 0x69, 0x73, 0x5f, 0x66, 0x72, 0x65, 0x65, 0x5f,
	0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x69, 0x73, 0x46, 0x72, 0x65, 0x65, 0x53, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x1f,
	0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x73, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x2a, 0x0a, 0x11, 0x69, 0x73, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x66, 0x65, 0x65, 0x64,
	0x62, 0x61, 0x63, 0x6b, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x69, 0x73, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x23, 0x0a, 0x0d, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x0c, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x61, 0x74, 0x65,
	0x12, 0x2e, 0x0a, 0x13, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61,
	0x63, 0x6b, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x02, 0x52, 0x11, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x61, 0x74, 0x65,
	0x12, 0x26, 0x0a, 0x0f, 0x69, 0x73, 0x5f, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x75, 0x70, 0x67, 0x72,
	0x61, 0x64, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x73, 0x41, 0x75, 0x74,
	0x6f, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x69, 0x73, 0x5f, 0x61,
	0x75, 0x74, 0x6f, 0x5f, 0x72, 0x65, 0x6e, 0x65, 0x77, 0x61, 0x6c, 0x18, 0x11, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x69, 0x73, 0x41, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x61, 0x6c,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x13, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xc5, 0x04, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x55, 0x72, 0x69, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x19,
	0x0a, 0x08, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70,
	0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x3c, 0x0a, 0x07, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x73, 0x18,
	0x11, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x45, 0x78, 0x74,
	0x65, 0x6e, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x64, 0x73, 0x1a, 0x3a, 0x0a, 0x0c, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x8a,
	0x03, 0x0a, 0x0d, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6d, 0x6f, 0x64, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19,
	0x0a, 0x08, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x61,
	0x6c, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0b, 0x72, 0x65, 0x61, 0x6c, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x73, 0x68, 0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73,
	0x68, 0x6f, 0x70, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x3e, 0x0a, 0x07, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x24, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x77, 0x69, 0x74,
	0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x73, 0x1a,
	0x3a, 0x0a, 0x0c, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xe2, 0x01, 0x0a, 0x0c,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x06,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x24, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6d, 0x73, 0x67, 0x1a, 0x39, 0x0a, 0x0b, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x32, 0x91, 0x04, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x38, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0c, 0x53, 0x77,
	0x69, 0x74, 0x63, 0x68, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a,
	0x0f, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x53, 0x65, 0x6e, 0x64,
	0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x3e, 0x0a, 0x0b, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x3c, 0x0a, 0x09, 0x46, 0x6f, 0x72, 0x67, 0x65, 0x74, 0x50, 0x77, 0x64, 0x12, 0x15,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x46, 0x0a, 0x13, 0x46, 0x6f, 0x72, 0x67, 0x65, 0x74, 0x50, 0x77, 0x64, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0d, 0x41, 0x70, 0x70, 0x4d, 0x6f,
	0x62, 0x69, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x06, 0x4c, 0x6f, 0x67,
	0x6f, 0x75, 0x74, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_userService_proto_rawDescOnce sync.Once
	file_userService_proto_rawDescData = file_userService_proto_rawDesc
)

func file_userService_proto_rawDescGZIP() []byte {
	file_userService_proto_rawDescOnce.Do(func() {
		file_userService_proto_rawDescData = protoimpl.X.CompressGZIP(file_userService_proto_rawDescData)
	})
	return file_userService_proto_rawDescData
}

var file_userService_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_userService_proto_goTypes = []interface{}{
	(*User)(nil),          // 0: services.User
	(*UserRequest)(nil),   // 1: services.UserRequest
	(*SwitchRequest)(nil), // 2: services.SwitchRequest
	(*UserResponse)(nil),  // 3: services.UserResponse
	nil,                   // 4: services.UserRequest.ExtendsEntry
	nil,                   // 5: services.SwitchRequest.ExtendsEntry
	nil,                   // 6: services.UserResponse.EntityEntry
	(*common.Pager)(nil),  // 7: common.Pager
	(*AuthRequest)(nil),   // 8: services.AuthRequest
	(*AuthResponse)(nil),  // 9: services.AuthResponse
}
var file_userService_proto_depIdxs = []int32{
	4,  // 0: services.UserRequest.extends:type_name -> services.UserRequest.ExtendsEntry
	5,  // 1: services.SwitchRequest.extends:type_name -> services.SwitchRequest.ExtendsEntry
	6,  // 2: services.UserResponse.entity:type_name -> services.UserResponse.EntityEntry
	7,  // 3: services.UserResponse.pager:type_name -> common.Pager
	0,  // 4: services.UserResponse.items:type_name -> services.User
	8,  // 5: services.UserService.Login:input_type -> services.AuthRequest
	2,  // 6: services.UserService.SwitchHandle:input_type -> services.SwitchRequest
	1,  // 7: services.UserService.LoginVerifySend:input_type -> services.UserRequest
	1,  // 8: services.UserService.VerifyEntry:input_type -> services.UserRequest
	1,  // 9: services.UserService.ForgetPwd:input_type -> services.UserRequest
	1,  // 10: services.UserService.ForgetPwdVerifySend:input_type -> services.UserRequest
	1,  // 11: services.UserService.AppMobileBind:input_type -> services.UserRequest
	1,  // 12: services.UserService.Logout:input_type -> services.UserRequest
	9,  // 13: services.UserService.Login:output_type -> services.AuthResponse
	9,  // 14: services.UserService.SwitchHandle:output_type -> services.AuthResponse
	3,  // 15: services.UserService.LoginVerifySend:output_type -> services.UserResponse
	3,  // 16: services.UserService.VerifyEntry:output_type -> services.UserResponse
	3,  // 17: services.UserService.ForgetPwd:output_type -> services.UserResponse
	3,  // 18: services.UserService.ForgetPwdVerifySend:output_type -> services.UserResponse
	3,  // 19: services.UserService.AppMobileBind:output_type -> services.UserResponse
	3,  // 20: services.UserService.Logout:output_type -> services.UserResponse
	13, // [13:21] is the sub-list for method output_type
	5,  // [5:13] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_userService_proto_init() }
func file_userService_proto_init() {
	if File_userService_proto != nil {
		return
	}
	file_baseInfoService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_userService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_userService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRequest); i {
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
		file_userService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SwitchRequest); i {
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
		file_userService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserResponse); i {
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
			RawDescriptor: file_userService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_userService_proto_goTypes,
		DependencyIndexes: file_userService_proto_depIdxs,
		MessageInfos:      file_userService_proto_msgTypes,
	}.Build()
	File_userService_proto = out.File
	file_userService_proto_rawDesc = nil
	file_userService_proto_goTypes = nil
	file_userService_proto_depIdxs = nil
}
