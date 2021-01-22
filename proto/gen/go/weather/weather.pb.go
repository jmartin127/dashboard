// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: weather/weather.proto

package weather

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetCurrentWeatherRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"` // By default maps to URL query parameter `address`.
}

func (x *GetCurrentWeatherRequest) Reset() {
	*x = GetCurrentWeatherRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weather_weather_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCurrentWeatherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurrentWeatherRequest) ProtoMessage() {}

func (x *GetCurrentWeatherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_weather_weather_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurrentWeatherRequest.ProtoReflect.Descriptor instead.
func (*GetCurrentWeatherRequest) Descriptor() ([]byte, []int) {
	return file_weather_weather_proto_rawDescGZIP(), []int{0}
}

func (x *GetCurrentWeatherRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type GetCurrentWeatherReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TempFahrenheit   int32 `protobuf:"varint,1,opt,name=tempFahrenheit,proto3" json:"tempFahrenheit,omitempty"`
	PrecipitationPct int32 `protobuf:"varint,2,opt,name=precipitationPct,proto3" json:"precipitationPct,omitempty"`
	HumidityPct      int32 `protobuf:"varint,3,opt,name=humidityPct,proto3" json:"humidityPct,omitempty"`
	WindMPH          int32 `protobuf:"varint,4,opt,name=windMPH,proto3" json:"windMPH,omitempty"`
}

func (x *GetCurrentWeatherReply) Reset() {
	*x = GetCurrentWeatherReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weather_weather_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCurrentWeatherReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurrentWeatherReply) ProtoMessage() {}

func (x *GetCurrentWeatherReply) ProtoReflect() protoreflect.Message {
	mi := &file_weather_weather_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurrentWeatherReply.ProtoReflect.Descriptor instead.
func (*GetCurrentWeatherReply) Descriptor() ([]byte, []int) {
	return file_weather_weather_proto_rawDescGZIP(), []int{1}
}

func (x *GetCurrentWeatherReply) GetTempFahrenheit() int32 {
	if x != nil {
		return x.TempFahrenheit
	}
	return 0
}

func (x *GetCurrentWeatherReply) GetPrecipitationPct() int32 {
	if x != nil {
		return x.PrecipitationPct
	}
	return 0
}

func (x *GetCurrentWeatherReply) GetHumidityPct() int32 {
	if x != nil {
		return x.HumidityPct
	}
	return 0
}

func (x *GetCurrentWeatherReply) GetWindMPH() int32 {
	if x != nil {
		return x.WindMPH
	}
	return 0
}

var File_weather_weather_proto protoreflect.FileDescriptor

var file_weather_weather_proto_rawDesc = []byte{
	0x0a, 0x15, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2f, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3d, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0xa8, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x26, 0x0a, 0x0e, 0x74, 0x65, 0x6d, 0x70, 0x46, 0x61, 0x68, 0x72, 0x65, 0x6e, 0x68,
	0x65, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x74, 0x65, 0x6d, 0x70, 0x46,
	0x61, 0x68, 0x72, 0x65, 0x6e, 0x68, 0x65, 0x69, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x70, 0x72, 0x65,
	0x63, 0x69, 0x70, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x63, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x10, 0x70, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x50, 0x63, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x68, 0x75, 0x6d, 0x69, 0x64, 0x69, 0x74,
	0x79, 0x50, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x68, 0x75, 0x6d, 0x69,
	0x64, 0x69, 0x74, 0x79, 0x50, 0x63, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x69, 0x6e, 0x64, 0x4d,
	0x50, 0x48, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x77, 0x69, 0x6e, 0x64, 0x4d, 0x50,
	0x48, 0x32, 0x7c, 0x0a, 0x07, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x12, 0x71, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65,
	0x72, 0x12, 0x21, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x47,
	0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f,
	0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x42,
	0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x6d,
	0x61, 0x72, 0x74, 0x69, 0x6e, 0x31, 0x32, 0x37, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f,
	0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_weather_weather_proto_rawDescOnce sync.Once
	file_weather_weather_proto_rawDescData = file_weather_weather_proto_rawDesc
)

func file_weather_weather_proto_rawDescGZIP() []byte {
	file_weather_weather_proto_rawDescOnce.Do(func() {
		file_weather_weather_proto_rawDescData = protoimpl.X.CompressGZIP(file_weather_weather_proto_rawDescData)
	})
	return file_weather_weather_proto_rawDescData
}

var file_weather_weather_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_weather_weather_proto_goTypes = []interface{}{
	(*GetCurrentWeatherRequest)(nil), // 0: weather.GetCurrentWeatherRequest
	(*GetCurrentWeatherReply)(nil),   // 1: weather.GetCurrentWeatherReply
}
var file_weather_weather_proto_depIdxs = []int32{
	0, // 0: weather.Weather.GetCurrentWeather:input_type -> weather.GetCurrentWeatherRequest
	1, // 1: weather.Weather.GetCurrentWeather:output_type -> weather.GetCurrentWeatherReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_weather_weather_proto_init() }
func file_weather_weather_proto_init() {
	if File_weather_weather_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_weather_weather_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCurrentWeatherRequest); i {
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
		file_weather_weather_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCurrentWeatherReply); i {
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
			RawDescriptor: file_weather_weather_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_weather_weather_proto_goTypes,
		DependencyIndexes: file_weather_weather_proto_depIdxs,
		MessageInfos:      file_weather_weather_proto_msgTypes,
	}.Build()
	File_weather_weather_proto = out.File
	file_weather_weather_proto_rawDesc = nil
	file_weather_weather_proto_goTypes = nil
	file_weather_weather_proto_depIdxs = nil
}
