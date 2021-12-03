// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: knowledgeService.proto

package services

import (
	_ "github.com/geiqin/micro-kit/protobuf/common"
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

type KnowledgeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int32 `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Top      int32 `protobuf:"varint,3,opt,name=top,proto3" json:"top"`
	//base params
	Id         int64   `protobuf:"varint,4,opt,name=id,proto3" json:"id"`
	Name       string  `protobuf:"bytes,5,opt,name=name,proto3" json:"name"`
	ItemSn     string  `protobuf:"bytes,6,opt,name=item_sn,json=itemSn,proto3" json:"item_sn"`
	BrandId    int32   `protobuf:"varint,7,opt,name=brand_id,json=brandId,proto3" json:"brand_id"`
	TagId      int32   `protobuf:"varint,8,opt,name=tag_id,json=tagId,proto3" json:"tag_id"`
	CatId      int32   `protobuf:"varint,9,opt,name=cat_id,json=catId,proto3" json:"cat_id"`
	MediaType  int32   `protobuf:"varint,10,opt,name=media_type,json=mediaType,proto3" json:"media_type"`
	Author     string  `protobuf:"bytes,11,opt,name=author,proto3" json:"author"`
	CourseId   int64   `protobuf:"varint,12,opt,name=course_id,json=courseId,proto3" json:"course_id"`
	ColumnId   int64   `protobuf:"varint,13,opt,name=column_id,json=columnId,proto3" json:"column_id"`
	SaleMode   int32   `protobuf:"varint,14,opt,name=sale_mode,json=saleMode,proto3" json:"sale_mode"`
	StageId    int32   `protobuf:"varint,15,opt,name=stage_id,json=stageId,proto3" json:"stage_id"`
	SubjectId  int32   `protobuf:"varint,16,opt,name=subject_id,json=subjectId,proto3" json:"subject_id"`
	GradeId    int32   `protobuf:"varint,17,opt,name=grade_id,json=gradeId,proto3" json:"grade_id"`
	TextbookId int32   `protobuf:"varint,18,opt,name=textbook_id,json=textbookId,proto3" json:"textbook_id"`
	BookType   int32   `protobuf:"varint,19,opt,name=book_type,json=bookType,proto3" json:"book_type"`
	CourseIds  []int64 `protobuf:"varint,20,rep,packed,name=course_ids,json=courseIds,proto3" json:"course_ids"`
}

func (x *KnowledgeRequest) Reset() {
	*x = KnowledgeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_knowledgeService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KnowledgeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KnowledgeRequest) ProtoMessage() {}

func (x *KnowledgeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_knowledgeService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KnowledgeRequest.ProtoReflect.Descriptor instead.
func (*KnowledgeRequest) Descriptor() ([]byte, []int) {
	return file_knowledgeService_proto_rawDescGZIP(), []int{0}
}

func (x *KnowledgeRequest) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *KnowledgeRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *KnowledgeRequest) GetTop() int32 {
	if x != nil {
		return x.Top
	}
	return 0
}

func (x *KnowledgeRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *KnowledgeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *KnowledgeRequest) GetItemSn() string {
	if x != nil {
		return x.ItemSn
	}
	return ""
}

func (x *KnowledgeRequest) GetBrandId() int32 {
	if x != nil {
		return x.BrandId
	}
	return 0
}

func (x *KnowledgeRequest) GetTagId() int32 {
	if x != nil {
		return x.TagId
	}
	return 0
}

func (x *KnowledgeRequest) GetCatId() int32 {
	if x != nil {
		return x.CatId
	}
	return 0
}

func (x *KnowledgeRequest) GetMediaType() int32 {
	if x != nil {
		return x.MediaType
	}
	return 0
}

func (x *KnowledgeRequest) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *KnowledgeRequest) GetCourseId() int64 {
	if x != nil {
		return x.CourseId
	}
	return 0
}

func (x *KnowledgeRequest) GetColumnId() int64 {
	if x != nil {
		return x.ColumnId
	}
	return 0
}

func (x *KnowledgeRequest) GetSaleMode() int32 {
	if x != nil {
		return x.SaleMode
	}
	return 0
}

func (x *KnowledgeRequest) GetStageId() int32 {
	if x != nil {
		return x.StageId
	}
	return 0
}

func (x *KnowledgeRequest) GetSubjectId() int32 {
	if x != nil {
		return x.SubjectId
	}
	return 0
}

func (x *KnowledgeRequest) GetGradeId() int32 {
	if x != nil {
		return x.GradeId
	}
	return 0
}

func (x *KnowledgeRequest) GetTextbookId() int32 {
	if x != nil {
		return x.TextbookId
	}
	return 0
}

func (x *KnowledgeRequest) GetBookType() int32 {
	if x != nil {
		return x.BookType
	}
	return 0
}

func (x *KnowledgeRequest) GetCourseIds() []int64 {
	if x != nil {
		return x.CourseIds
	}
	return nil
}

var File_knowledgeService_proto protoreflect.FileDescriptor

var file_knowledgeService_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x69, 0x74, 0x65, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9d, 0x04, 0x0a, 0x10, 0x4b, 0x6e, 0x6f,
	0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61,
	0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x74, 0x6f, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x74,
	0x6f, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x73,
	0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x53, 0x6e, 0x12,
	0x19, 0x0a, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x74, 0x61,
	0x67, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x61, 0x67, 0x49,
	0x64, 0x12, 0x15, 0x0a, 0x06, 0x63, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x63, 0x61, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x64, 0x69,
	0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6d, 0x65,
	0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12,
	0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x61, 0x6c,
	0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x61,
	0x6c, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x67, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x74, 0x61, 0x67, 0x65, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x10, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64,
	0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x11, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x67, 0x72, 0x61, 0x64, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x74,
	0x65, 0x78, 0x74, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x12, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x74, 0x65, 0x78, 0x74, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x14, 0x20, 0x03, 0x28, 0x03, 0x52, 0x09, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x73, 0x32, 0x96, 0x03, 0x0a, 0x0d, 0x43, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x49, 0x74, 0x65, 0x6d, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x32,
	0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x32, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0e, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x1a, 0x16, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0e, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x1a, 0x16, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x07, 0x44, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x12, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x74,
	0x65, 0x6d, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x06,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x4b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0b,
	0x46, 0x72, 0x6f, 0x6e, 0x74, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1a, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x32, 0x9b, 0x04, 0x0a, 0x0d, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x1a, 0x16, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x74, 0x65,
	0x6d, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x06, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x49, 0x74, 0x65, 0x6d, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x3f, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1a, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x42, 0x0a, 0x0a, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1a,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4b, 0x6e, 0x6f, 0x77, 0x6c, 0x65,
	0x64, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0e, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x1a, 0x16, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x07, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x12, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d,
	0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x06, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x12, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x4b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0b, 0x46, 0x72,
	0x6f, 0x6e, 0x74, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x4b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_knowledgeService_proto_rawDescOnce sync.Once
	file_knowledgeService_proto_rawDescData = file_knowledgeService_proto_rawDesc
)

func file_knowledgeService_proto_rawDescGZIP() []byte {
	file_knowledgeService_proto_rawDescOnce.Do(func() {
		file_knowledgeService_proto_rawDescData = protoimpl.X.CompressGZIP(file_knowledgeService_proto_rawDescData)
	})
	return file_knowledgeService_proto_rawDescData
}

var file_knowledgeService_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_knowledgeService_proto_goTypes = []interface{}{
	(*KnowledgeRequest)(nil), // 0: services.KnowledgeRequest
	(*Item)(nil),             // 1: services.Item
	(*ItemResponse)(nil),     // 2: services.ItemResponse
}
var file_knowledgeService_proto_depIdxs = []int32{
	1,  // 0: services.CourseService.Create:input_type -> services.Item
	1,  // 1: services.CourseService.Update:input_type -> services.Item
	1,  // 2: services.CourseService.Delete:input_type -> services.Item
	1,  // 3: services.CourseService.Get:input_type -> services.Item
	1,  // 4: services.CourseService.Display:input_type -> services.Item
	0,  // 5: services.CourseService.Search:input_type -> services.KnowledgeRequest
	0,  // 6: services.CourseService.FrontSearch:input_type -> services.KnowledgeRequest
	1,  // 7: services.ColumnService.Create:input_type -> services.Item
	1,  // 8: services.ColumnService.Update:input_type -> services.Item
	1,  // 9: services.ColumnService.Delete:input_type -> services.Item
	0,  // 10: services.ColumnService.AddList:input_type -> services.KnowledgeRequest
	0,  // 11: services.ColumnService.RemoveList:input_type -> services.KnowledgeRequest
	1,  // 12: services.ColumnService.Get:input_type -> services.Item
	1,  // 13: services.ColumnService.Display:input_type -> services.Item
	0,  // 14: services.ColumnService.Search:input_type -> services.KnowledgeRequest
	0,  // 15: services.ColumnService.FrontSearch:input_type -> services.KnowledgeRequest
	2,  // 16: services.CourseService.Create:output_type -> services.ItemResponse
	2,  // 17: services.CourseService.Update:output_type -> services.ItemResponse
	2,  // 18: services.CourseService.Delete:output_type -> services.ItemResponse
	2,  // 19: services.CourseService.Get:output_type -> services.ItemResponse
	2,  // 20: services.CourseService.Display:output_type -> services.ItemResponse
	2,  // 21: services.CourseService.Search:output_type -> services.ItemResponse
	2,  // 22: services.CourseService.FrontSearch:output_type -> services.ItemResponse
	2,  // 23: services.ColumnService.Create:output_type -> services.ItemResponse
	2,  // 24: services.ColumnService.Update:output_type -> services.ItemResponse
	2,  // 25: services.ColumnService.Delete:output_type -> services.ItemResponse
	2,  // 26: services.ColumnService.AddList:output_type -> services.ItemResponse
	2,  // 27: services.ColumnService.RemoveList:output_type -> services.ItemResponse
	2,  // 28: services.ColumnService.Get:output_type -> services.ItemResponse
	2,  // 29: services.ColumnService.Display:output_type -> services.ItemResponse
	2,  // 30: services.ColumnService.Search:output_type -> services.ItemResponse
	2,  // 31: services.ColumnService.FrontSearch:output_type -> services.ItemResponse
	16, // [16:32] is the sub-list for method output_type
	0,  // [0:16] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_knowledgeService_proto_init() }
func file_knowledgeService_proto_init() {
	if File_knowledgeService_proto != nil {
		return
	}
	file_itemService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_knowledgeService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KnowledgeRequest); i {
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
			RawDescriptor: file_knowledgeService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_knowledgeService_proto_goTypes,
		DependencyIndexes: file_knowledgeService_proto_depIdxs,
		MessageInfos:      file_knowledgeService_proto_msgTypes,
	}.Build()
	File_knowledgeService_proto = out.File
	file_knowledgeService_proto_rawDesc = nil
	file_knowledgeService_proto_goTypes = nil
	file_knowledgeService_proto_depIdxs = nil
}