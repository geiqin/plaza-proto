// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: personService.proto

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

//人员
type Person struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`                            //ID
	RealName          string `protobuf:"bytes,2,opt,name=real_name,json=realName,proto3" json:"real_name"` //真实姓名
	IdCard            string `protobuf:"bytes,3,opt,name=id_card,json=idCard,proto3" json:"id_card"`       //身份证号
	Birthday          string `protobuf:"bytes,4,opt,name=birthday,proto3" json:"birthday"`                 //出生日期
	Gender            string `protobuf:"bytes,5,opt,name=gender,proto3" json:"gender"`                     //性别：1-男，2-女
	Nation            string `protobuf:"bytes,6,opt,name=nation,proto3" json:"nation"`                     //民族
	Political         string `protobuf:"bytes,7,opt,name=political,proto3" json:"political"`               //政治面貌
	Nationality       string `protobuf:"bytes,8,opt,name=nationality,proto3" json:"nationality"`           //国籍
	Passport          string `protobuf:"bytes,9,opt,name=passport,proto3" json:"passport"`                 //护照编号
	FamilyId          int64  `protobuf:"varint,10,opt,name=family_id,json=familyId,proto3" json:"family_id"`
	FamilyRelation    string `protobuf:"bytes,11,opt,name=family_relation,json=familyRelation,proto3" json:"family_relation"`
	AreaId            int64  `protobuf:"varint,12,opt,name=area_id,json=areaId,proto3" json:"area_id"`                                   //区域ID
	Addr              string `protobuf:"bytes,13,opt,name=addr,proto3" json:"addr"`                                                      //住址
	Mobile            string `protobuf:"bytes,14,opt,name=mobile,proto3" json:"mobile"`                                                  //联系电话
	AvatarId          int64  `protobuf:"varint,15,opt,name=avatar_id,json=avatarId,proto3" json:"avatar_id"`                             //头像ID
	AvatarUrl         string `protobuf:"bytes,16,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url"`                           //头像URL
	StudentState      string `protobuf:"bytes,17,opt,name=student_state,json=studentState,proto3" json:"student_state"`                  //学生状况
	SoldierState      string `protobuf:"bytes,18,opt,name=soldier_state,json=soldierState,proto3" json:"soldier_state"`                  //军人状况
	CivilServantState string `protobuf:"bytes,19,opt,name=civil_servant_state,json=civilServantState,proto3" json:"civil_servant_state"` //公务员状况
	HandicappedState  string `protobuf:"bytes,20,opt,name=handicapped_state,json=handicappedState,proto3" json:"handicapped_state"`      //残疾状况
	Memo              string `protobuf:"bytes,21,opt,name=memo,proto3" json:"memo"`                                                      //备注
	Status            string `protobuf:"bytes,22,opt,name=status,proto3" json:"status"`                                                  //状态
	CreatedAt         string `protobuf:"bytes,23,opt,name=created_at,json=createdAt,proto3" json:"created_at"`                           //创建时间
	UpdatedAt         string `protobuf:"bytes,24,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`                           //修改时间
}

func (x *Person) Reset() {
	*x = Person{}
	if protoimpl.UnsafeEnabled {
		mi := &file_personService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person) ProtoMessage() {}

func (x *Person) ProtoReflect() protoreflect.Message {
	mi := &file_personService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person.ProtoReflect.Descriptor instead.
func (*Person) Descriptor() ([]byte, []int) {
	return file_personService_proto_rawDescGZIP(), []int{0}
}

func (x *Person) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Person) GetRealName() string {
	if x != nil {
		return x.RealName
	}
	return ""
}

func (x *Person) GetIdCard() string {
	if x != nil {
		return x.IdCard
	}
	return ""
}

func (x *Person) GetBirthday() string {
	if x != nil {
		return x.Birthday
	}
	return ""
}

func (x *Person) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *Person) GetNation() string {
	if x != nil {
		return x.Nation
	}
	return ""
}

func (x *Person) GetPolitical() string {
	if x != nil {
		return x.Political
	}
	return ""
}

func (x *Person) GetNationality() string {
	if x != nil {
		return x.Nationality
	}
	return ""
}

func (x *Person) GetPassport() string {
	if x != nil {
		return x.Passport
	}
	return ""
}

func (x *Person) GetFamilyId() int64 {
	if x != nil {
		return x.FamilyId
	}
	return 0
}

func (x *Person) GetFamilyRelation() string {
	if x != nil {
		return x.FamilyRelation
	}
	return ""
}

func (x *Person) GetAreaId() int64 {
	if x != nil {
		return x.AreaId
	}
	return 0
}

func (x *Person) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Person) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *Person) GetAvatarId() int64 {
	if x != nil {
		return x.AvatarId
	}
	return 0
}

func (x *Person) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

func (x *Person) GetStudentState() string {
	if x != nil {
		return x.StudentState
	}
	return ""
}

func (x *Person) GetSoldierState() string {
	if x != nil {
		return x.SoldierState
	}
	return ""
}

func (x *Person) GetCivilServantState() string {
	if x != nil {
		return x.CivilServantState
	}
	return ""
}

func (x *Person) GetHandicappedState() string {
	if x != nil {
		return x.HandicappedState
	}
	return ""
}

func (x *Person) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *Person) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Person) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Person) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type PersonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int32  `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize int32  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Sorting  string `protobuf:"bytes,3,opt,name=sorting,proto3" json:"sorting"`
	//----
	Id                int64   `protobuf:"varint,4,opt,name=id,proto3" json:"id"`
	RealName          string  `protobuf:"bytes,5,opt,name=real_name,json=realName,proto3" json:"real_name"`
	IdCard            string  `protobuf:"bytes,6,opt,name=id_card,json=idCard,proto3" json:"id_card"`
	Nation            string  `protobuf:"bytes,7,opt,name=nation,proto3" json:"nation"`
	Political         string  `protobuf:"bytes,8,opt,name=political,proto3" json:"political"`
	Nationality       string  `protobuf:"bytes,9,opt,name=nationality,proto3" json:"nationality"`
	Passport          string  `protobuf:"bytes,10,opt,name=passport,proto3" json:"passport"`
	AreaId            int64   `protobuf:"varint,11,opt,name=area_id,json=areaId,proto3" json:"area_id"`
	StudentState      string  `protobuf:"bytes,12,opt,name=student_state,json=studentState,proto3" json:"student_state"`
	SoldierState      string  `protobuf:"bytes,13,opt,name=soldier_state,json=soldierState,proto3" json:"soldier_state"`
	CivilServantState string  `protobuf:"bytes,14,opt,name=civil_servant_state,json=civilServantState,proto3" json:"civil_servant_state"`
	HandicappedState  string  `protobuf:"bytes,15,opt,name=handicapped_state,json=handicappedState,proto3" json:"handicapped_state"`
	StartAt           string  `protobuf:"bytes,16,opt,name=start_at,json=startAt,proto3" json:"start_at"`
	EndAt             string  `protobuf:"bytes,17,opt,name=end_at,json=endAt,proto3" json:"end_at"`
	Status            string  `protobuf:"bytes,18,opt,name=status,proto3" json:"status"`
	Keywords          string  `protobuf:"bytes,19,opt,name=keywords,proto3" json:"keywords"`
	Ids               []int64 `protobuf:"varint,20,rep,packed,name=ids,proto3" json:"ids"`
}

func (x *PersonRequest) Reset() {
	*x = PersonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_personService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonRequest) ProtoMessage() {}

func (x *PersonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_personService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonRequest.ProtoReflect.Descriptor instead.
func (*PersonRequest) Descriptor() ([]byte, []int) {
	return file_personService_proto_rawDescGZIP(), []int{1}
}

func (x *PersonRequest) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *PersonRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *PersonRequest) GetSorting() string {
	if x != nil {
		return x.Sorting
	}
	return ""
}

func (x *PersonRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PersonRequest) GetRealName() string {
	if x != nil {
		return x.RealName
	}
	return ""
}

func (x *PersonRequest) GetIdCard() string {
	if x != nil {
		return x.IdCard
	}
	return ""
}

func (x *PersonRequest) GetNation() string {
	if x != nil {
		return x.Nation
	}
	return ""
}

func (x *PersonRequest) GetPolitical() string {
	if x != nil {
		return x.Political
	}
	return ""
}

func (x *PersonRequest) GetNationality() string {
	if x != nil {
		return x.Nationality
	}
	return ""
}

func (x *PersonRequest) GetPassport() string {
	if x != nil {
		return x.Passport
	}
	return ""
}

func (x *PersonRequest) GetAreaId() int64 {
	if x != nil {
		return x.AreaId
	}
	return 0
}

func (x *PersonRequest) GetStudentState() string {
	if x != nil {
		return x.StudentState
	}
	return ""
}

func (x *PersonRequest) GetSoldierState() string {
	if x != nil {
		return x.SoldierState
	}
	return ""
}

func (x *PersonRequest) GetCivilServantState() string {
	if x != nil {
		return x.CivilServantState
	}
	return ""
}

func (x *PersonRequest) GetHandicappedState() string {
	if x != nil {
		return x.HandicappedState
	}
	return ""
}

func (x *PersonRequest) GetStartAt() string {
	if x != nil {
		return x.StartAt
	}
	return ""
}

func (x *PersonRequest) GetEndAt() string {
	if x != nil {
		return x.EndAt
	}
	return ""
}

func (x *PersonRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *PersonRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *PersonRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type PersonData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Person       `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*Person     `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Info   *common.Info  `protobuf:"bytes,4,opt,name=info,proto3" json:"info"`
}

func (x *PersonData) Reset() {
	*x = PersonData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_personService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersonData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonData) ProtoMessage() {}

func (x *PersonData) ProtoReflect() protoreflect.Message {
	mi := &file_personService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonData.ProtoReflect.Descriptor instead.
func (*PersonData) Descriptor() ([]byte, []int) {
	return file_personService_proto_rawDescGZIP(), []int{2}
}

func (x *PersonData) GetEntity() *Person {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *PersonData) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *PersonData) GetItems() []*Person {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *PersonData) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type PersonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *PersonData   `protobuf:"bytes,1,opt,name=data,proto3" json:"data"`
	Error *common.Error `protobuf:"bytes,2,opt,name=error,proto3" json:"error"`
}

func (x *PersonResponse) Reset() {
	*x = PersonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_personService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonResponse) ProtoMessage() {}

func (x *PersonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_personService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonResponse.ProtoReflect.Descriptor instead.
func (*PersonResponse) Descriptor() ([]byte, []int) {
	return file_personService_proto_rawDescGZIP(), []int{3}
}

func (x *PersonResponse) GetData() *PersonData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *PersonResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_personService_proto protoreflect.FileDescriptor

var file_personService_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a,
	0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xce, 0x05, 0x0a, 0x06, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x72, 0x65, 0x61, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x72, 0x65, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x64,
	0x5f, 0x63, 0x61, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x64, 0x43,
	0x61, 0x72, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x12,
	0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1c, 0x0a, 0x09, 0x70, 0x6f, 0x6c, 0x69, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x6f, 0x6c, 0x69, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x12, 0x20, 0x0a,
	0x0b, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x66,
	0x61, 0x6d, 0x69, 0x6c, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x66, 0x61, 0x6d, 0x69,
	0x6c, 0x79, 0x5f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x61, 0x72, 0x65, 0x61, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64,
	0x64, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x16,
	0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55,
	0x72, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x6f, 0x6c, 0x64, 0x69,
	0x65, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x73, 0x6f, 0x6c, 0x64, 0x69, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x2e, 0x0a, 0x13,
	0x63, 0x69, 0x76, 0x69, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x61, 0x6e, 0x74, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x63, 0x69, 0x76, 0x69, 0x6c,
	0x53, 0x65, 0x72, 0x76, 0x61, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x0a, 0x11,
	0x68, 0x61, 0x6e, 0x64, 0x69, 0x63, 0x61, 0x70, 0x70, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x68, 0x61, 0x6e, 0x64, 0x69, 0x63, 0x61,
	0x70, 0x70, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x6d,
	0x6f, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0xce, 0x04, 0x0a, 0x0d, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74,
	0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69,
	0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x61, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x17, 0x0a, 0x07, 0x69, 0x64, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x69, 0x64, 0x43, 0x61, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1c, 0x0a, 0x09, 0x70, 0x6f, 0x6c, 0x69, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x6f, 0x6c, 0x69, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x12, 0x20,
	0x0a, 0x0b, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x61, 0x72, 0x65, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61,
	0x72, 0x65, 0x61, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x74,
	0x75, 0x64, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x6f,
	0x6c, 0x64, 0x69, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x73, 0x6f, 0x6c, 0x64, 0x69, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x2e, 0x0a, 0x13, 0x63, 0x69, 0x76, 0x69, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x61, 0x6e, 0x74,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x63, 0x69,
	0x76, 0x69, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x61, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x2b, 0x0a, 0x11, 0x68, 0x61, 0x6e, 0x64, 0x69, 0x63, 0x61, 0x70, 0x70, 0x65, 0x64, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x68, 0x61, 0x6e, 0x64,
	0x69, 0x63, 0x61, 0x70, 0x70, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x61, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x41, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x65, 0x6e, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6e, 0x64, 0x41, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72,
	0x64, 0x73, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72,
	0x64, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x14, 0x20, 0x03, 0x28, 0x03, 0x52,
	0x03, 0x69, 0x64, 0x73, 0x22, 0xa5, 0x01, 0x0a, 0x0a, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x28, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a,
	0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67,
	0x65, 0x72, 0x12, 0x26, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x5f, 0x0a, 0x0e,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xe8, 0x02,
	0x0a, 0x0d, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x36, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x1a, 0x18, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x36, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x1a, 0x18, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x10,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x06,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x04, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_personService_proto_rawDescOnce sync.Once
	file_personService_proto_rawDescData = file_personService_proto_rawDesc
)

func file_personService_proto_rawDescGZIP() []byte {
	file_personService_proto_rawDescOnce.Do(func() {
		file_personService_proto_rawDescData = protoimpl.X.CompressGZIP(file_personService_proto_rawDescData)
	})
	return file_personService_proto_rawDescData
}

var file_personService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_personService_proto_goTypes = []interface{}{
	(*Person)(nil),         // 0: services.Person
	(*PersonRequest)(nil),  // 1: services.PersonRequest
	(*PersonData)(nil),     // 2: services.PersonData
	(*PersonResponse)(nil), // 3: services.PersonResponse
	(*common.Pager)(nil),   // 4: common.Pager
	(*common.Info)(nil),    // 5: common.Info
	(*common.Error)(nil),   // 6: common.Error
}
var file_personService_proto_depIdxs = []int32{
	0,  // 0: services.PersonData.entity:type_name -> services.Person
	4,  // 1: services.PersonData.pager:type_name -> common.Pager
	0,  // 2: services.PersonData.items:type_name -> services.Person
	5,  // 3: services.PersonData.info:type_name -> common.Info
	2,  // 4: services.PersonResponse.data:type_name -> services.PersonData
	6,  // 5: services.PersonResponse.error:type_name -> common.Error
	0,  // 6: services.PersonService.Create:input_type -> services.Person
	0,  // 7: services.PersonService.Update:input_type -> services.Person
	0,  // 8: services.PersonService.Delete:input_type -> services.Person
	0,  // 9: services.PersonService.Get:input_type -> services.Person
	1,  // 10: services.PersonService.Search:input_type -> services.PersonRequest
	1,  // 11: services.PersonService.List:input_type -> services.PersonRequest
	3,  // 12: services.PersonService.Create:output_type -> services.PersonResponse
	3,  // 13: services.PersonService.Update:output_type -> services.PersonResponse
	3,  // 14: services.PersonService.Delete:output_type -> services.PersonResponse
	3,  // 15: services.PersonService.Get:output_type -> services.PersonResponse
	3,  // 16: services.PersonService.Search:output_type -> services.PersonResponse
	3,  // 17: services.PersonService.List:output_type -> services.PersonResponse
	12, // [12:18] is the sub-list for method output_type
	6,  // [6:12] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_personService_proto_init() }
func file_personService_proto_init() {
	if File_personService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_personService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Person); i {
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
		file_personService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PersonRequest); i {
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
		file_personService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PersonData); i {
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
		file_personService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PersonResponse); i {
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
			RawDescriptor: file_personService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_personService_proto_goTypes,
		DependencyIndexes: file_personService_proto_depIdxs,
		MessageInfos:      file_personService_proto_msgTypes,
	}.Build()
	File_personService_proto = out.File
	file_personService_proto_rawDesc = nil
	file_personService_proto_goTypes = nil
	file_personService_proto_depIdxs = nil
}