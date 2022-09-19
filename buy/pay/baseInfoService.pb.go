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

//支付信息
type ChargeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderType      string `protobuf:"bytes,1,opt,name=order_type,json=orderType,proto3" json:"order_type"`                  //订单类型
	OrderId        int64  `protobuf:"varint,2,opt,name=order_id,json=orderId,proto3" json:"order_id"`                       //订单id
	OrderSn        string `protobuf:"bytes,3,opt,name=order_sn,json=orderSn,proto3" json:"order_sn"`                        //订单编号
	Amount         int64  `protobuf:"varint,4,opt,name=amount,proto3" json:"amount"`                                        //支付金额
	Currency       string `protobuf:"bytes,5,opt,name=currency,proto3" json:"currency"`                                     //货币类型
	Subject        string `protobuf:"bytes,6,opt,name=subject,proto3" json:"subject"`                                       //支付标题
	Body           string `protobuf:"bytes,7,opt,name=body,proto3" json:"body"`                                             //主体描述
	Attach         string `protobuf:"bytes,8,opt,name=attach,proto3" json:"attach"`                                         //附加信息
	TargetUserType string `protobuf:"bytes,9,opt,name=target_user_type,json=targetUserType,proto3" json:"target_user_type"` //目标用户类型
	TargetUserId   int64  `protobuf:"varint,10,opt,name=target_user_id,json=targetUserId,proto3" json:"target_user_id"`     //目标用户ID
}

func (x *ChargeInfo) Reset() {
	*x = ChargeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_baseInfoService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChargeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChargeInfo) ProtoMessage() {}

func (x *ChargeInfo) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ChargeInfo.ProtoReflect.Descriptor instead.
func (*ChargeInfo) Descriptor() ([]byte, []int) {
	return file_baseInfoService_proto_rawDescGZIP(), []int{0}
}

func (x *ChargeInfo) GetOrderType() string {
	if x != nil {
		return x.OrderType
	}
	return ""
}

func (x *ChargeInfo) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *ChargeInfo) GetOrderSn() string {
	if x != nil {
		return x.OrderSn
	}
	return ""
}

func (x *ChargeInfo) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *ChargeInfo) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *ChargeInfo) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *ChargeInfo) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *ChargeInfo) GetAttach() string {
	if x != nil {
		return x.Attach
	}
	return ""
}

func (x *ChargeInfo) GetTargetUserType() string {
	if x != nil {
		return x.TargetUserType
	}
	return ""
}

func (x *ChargeInfo) GetTargetUserId() int64 {
	if x != nil {
		return x.TargetUserId
	}
	return 0
}

var File_baseInfoService_proto protoreflect.FileDescriptor

var file_baseInfoService_proto_rawDesc = []byte{
	0x0a, 0x15, 0x62, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x22, 0xab, 0x02, 0x0a, 0x0a, 0x43, 0x68, 0x61, 0x72, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x73, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x53, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x74, 0x74, 0x61, 0x63,
	0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x12,
	0x28, 0x0a, 0x10, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0c, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x42,
	0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_baseInfoService_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_baseInfoService_proto_goTypes = []interface{}{
	(*ChargeInfo)(nil), // 0: services.ChargeInfo
}
var file_baseInfoService_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_baseInfoService_proto_init() }
func file_baseInfoService_proto_init() {
	if File_baseInfoService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_baseInfoService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChargeInfo); i {
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
			NumMessages:   1,
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
