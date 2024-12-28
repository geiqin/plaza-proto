// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: localConfigService.proto

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

// 同城配送配置
type LocalConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsTimedArrival string             `protobuf:"bytes,1,opt,name=is_timed_arrival,json=isTimedArrival,proto3" json:"is_timed_arrival"` //是否开启定时达(0否、1是)
	TimedArrival   *LocalTimedArrival `protobuf:"bytes,2,opt,name=timed_arrival,json=timedArrival,proto3" json:"timed_arrival"`         //定时达规则
	Description    string             `protobuf:"bytes,3,opt,name=description,proto3" json:"description"`                               //配送说明
}

func (x *LocalConfig) Reset() {
	*x = LocalConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_localConfigService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalConfig) ProtoMessage() {}

func (x *LocalConfig) ProtoReflect() protoreflect.Message {
	mi := &file_localConfigService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalConfig.ProtoReflect.Descriptor instead.
func (*LocalConfig) Descriptor() ([]byte, []int) {
	return file_localConfigService_proto_rawDescGZIP(), []int{0}
}

func (x *LocalConfig) GetIsTimedArrival() string {
	if x != nil {
		return x.IsTimedArrival
	}
	return ""
}

func (x *LocalConfig) GetTimedArrival() *LocalTimedArrival {
	if x != nil {
		return x.TimedArrival
	}
	return nil
}

func (x *LocalConfig) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

// 定时送达数据
type LocalTimedArrival struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RepeatType         string   `protobuf:"bytes,1,opt,name=repeat_type,json=repeatType,proto3" json:"repeat_type"`                            //配送时间:all_day全天、repeat_day每天重复、repeat_week每周重复
	RepeatWeeks        int32    `protobuf:"varint,2,opt,name=repeat_weeks,json=repeatWeeks,proto3" json:"repeat_weeks"`                        //自提时间每周重复的星期数按|运算: 1周一，2周二，4周三，8周四，16周五，32周六，64周日
	TimeSection        string   `protobuf:"bytes,3,opt,name=time_section,json=timeSection,proto3" json:"time_section"`                         //配送时段细分: day按天，hour按小时，half_hour按半小时
	AppointmentType    string   `protobuf:"bytes,4,opt,name=appointment_type,json=appointmentType,proto3" json:"appointment_type"`             //预约下单：no无需提前、day提前X天、hour提前X小时、minute提前X分钟
	AppointmentValue   int32    `protobuf:"varint,5,opt,name=appointment_value,json=appointmentValue,proto3" json:"appointment_value"`         //预约下单的天数、小时、分钟
	AppointmentMaxDays int32    `protobuf:"varint,6,opt,name=appointment_max_days,json=appointmentMaxDays,proto3" json:"appointment_max_days"` //最长预约天数(0只能当天配送)
	Ranges             []string `protobuf:"bytes,7,rep,name=ranges,proto3" json:"ranges"`                                                      //范围 待定
	DeliveryTimes      []string `protobuf:"bytes,8,rep,name=delivery_times,json=deliveryTimes,proto3" json:"delivery_times"`                   // 配送时段 待定
}

func (x *LocalTimedArrival) Reset() {
	*x = LocalTimedArrival{}
	if protoimpl.UnsafeEnabled {
		mi := &file_localConfigService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalTimedArrival) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalTimedArrival) ProtoMessage() {}

func (x *LocalTimedArrival) ProtoReflect() protoreflect.Message {
	mi := &file_localConfigService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalTimedArrival.ProtoReflect.Descriptor instead.
func (*LocalTimedArrival) Descriptor() ([]byte, []int) {
	return file_localConfigService_proto_rawDescGZIP(), []int{1}
}

func (x *LocalTimedArrival) GetRepeatType() string {
	if x != nil {
		return x.RepeatType
	}
	return ""
}

func (x *LocalTimedArrival) GetRepeatWeeks() int32 {
	if x != nil {
		return x.RepeatWeeks
	}
	return 0
}

func (x *LocalTimedArrival) GetTimeSection() string {
	if x != nil {
		return x.TimeSection
	}
	return ""
}

func (x *LocalTimedArrival) GetAppointmentType() string {
	if x != nil {
		return x.AppointmentType
	}
	return ""
}

func (x *LocalTimedArrival) GetAppointmentValue() int32 {
	if x != nil {
		return x.AppointmentValue
	}
	return 0
}

func (x *LocalTimedArrival) GetAppointmentMaxDays() int32 {
	if x != nil {
		return x.AppointmentMaxDays
	}
	return 0
}

func (x *LocalTimedArrival) GetRanges() []string {
	if x != nil {
		return x.Ranges
	}
	return nil
}

func (x *LocalTimedArrival) GetDeliveryTimes() []string {
	if x != nil {
		return x.DeliveryTimes
	}
	return nil
}

var File_localConfigService_proto protoreflect.FileDescriptor

var file_localConfigService_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x22, 0x9b, 0x01, 0x0a, 0x0b, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x28, 0x0a, 0x10, 0x69, 0x73, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x64,
	0x5f, 0x61, 0x72, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x69, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x64, 0x41, 0x72, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x12, 0x40,
	0x0a, 0x0d, 0x74, 0x69, 0x6d, 0x65, 0x64, 0x5f, 0x61, 0x72, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x64, 0x41, 0x72, 0x72, 0x69, 0x76,
	0x61, 0x6c, 0x52, 0x0c, 0x74, 0x69, 0x6d, 0x65, 0x64, 0x41, 0x72, 0x72, 0x69, 0x76, 0x61, 0x6c,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0xc3, 0x02, 0x0a, 0x11, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65,
	0x64, 0x41, 0x72, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x70, 0x65,
	0x61, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72,
	0x65, 0x70, 0x65, 0x61, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x70,
	0x65, 0x61, 0x74, 0x5f, 0x77, 0x65, 0x65, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0b, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x57, 0x65, 0x65, 0x6b, 0x73, 0x12, 0x21, 0x0a, 0x0c,
	0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x29, 0x0a, 0x10, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x70, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x61, 0x70,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x61, 0x70, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x64, 0x61, 0x79, 0x73, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x4d, 0x61, 0x78, 0x44, 0x61, 0x79, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x6e,
	0x67, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x72, 0x61, 0x6e, 0x67, 0x65,
	0x73, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_localConfigService_proto_rawDescOnce sync.Once
	file_localConfigService_proto_rawDescData = file_localConfigService_proto_rawDesc
)

func file_localConfigService_proto_rawDescGZIP() []byte {
	file_localConfigService_proto_rawDescOnce.Do(func() {
		file_localConfigService_proto_rawDescData = protoimpl.X.CompressGZIP(file_localConfigService_proto_rawDescData)
	})
	return file_localConfigService_proto_rawDescData
}

var file_localConfigService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_localConfigService_proto_goTypes = []interface{}{
	(*LocalConfig)(nil),       // 0: services.LocalConfig
	(*LocalTimedArrival)(nil), // 1: services.LocalTimedArrival
}
var file_localConfigService_proto_depIdxs = []int32{
	1, // 0: services.LocalConfig.timed_arrival:type_name -> services.LocalTimedArrival
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_localConfigService_proto_init() }
func file_localConfigService_proto_init() {
	if File_localConfigService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_localConfigService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocalConfig); i {
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
		file_localConfigService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocalTimedArrival); i {
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
			RawDescriptor: file_localConfigService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_localConfigService_proto_goTypes,
		DependencyIndexes: file_localConfigService_proto_depIdxs,
		MessageInfos:      file_localConfigService_proto_msgTypes,
	}.Build()
	File_localConfigService_proto = out.File
	file_localConfigService_proto_rawDesc = nil
	file_localConfigService_proto_goTypes = nil
	file_localConfigService_proto_depIdxs = nil
}
