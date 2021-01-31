// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: jmartin127/traffic/v1/traffic.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

type GetTravelTimeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OriginAddress      string `protobuf:"bytes,1,opt,name=originAddress,proto3" json:"originAddress,omitempty"`           // By default maps to URL query parameter `origin`.
	DestinationAddress string `protobuf:"bytes,2,opt,name=destinationAddress,proto3" json:"destinationAddress,omitempty"` // By default maps to URL query parameter `destination`.
}

func (x *GetTravelTimeRequest) Reset() {
	*x = GetTravelTimeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jmartin127_traffic_v1_traffic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTravelTimeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTravelTimeRequest) ProtoMessage() {}

func (x *GetTravelTimeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_jmartin127_traffic_v1_traffic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTravelTimeRequest.ProtoReflect.Descriptor instead.
func (*GetTravelTimeRequest) Descriptor() ([]byte, []int) {
	return file_jmartin127_traffic_v1_traffic_proto_rawDescGZIP(), []int{0}
}

func (x *GetTravelTimeRequest) GetOriginAddress() string {
	if x != nil {
		return x.OriginAddress
	}
	return ""
}

func (x *GetTravelTimeRequest) GetDestinationAddress() string {
	if x != nil {
		return x.DestinationAddress
	}
	return ""
}

type GetTravelTimeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TravelTime *duration.Duration `protobuf:"bytes,1,opt,name=travelTime,proto3" json:"travelTime,omitempty"`
}

func (x *GetTravelTimeReply) Reset() {
	*x = GetTravelTimeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jmartin127_traffic_v1_traffic_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTravelTimeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTravelTimeReply) ProtoMessage() {}

func (x *GetTravelTimeReply) ProtoReflect() protoreflect.Message {
	mi := &file_jmartin127_traffic_v1_traffic_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTravelTimeReply.ProtoReflect.Descriptor instead.
func (*GetTravelTimeReply) Descriptor() ([]byte, []int) {
	return file_jmartin127_traffic_v1_traffic_proto_rawDescGZIP(), []int{1}
}

func (x *GetTravelTimeReply) GetTravelTime() *duration.Duration {
	if x != nil {
		return x.TravelTime
	}
	return nil
}

var File_jmartin127_traffic_v1_traffic_proto protoreflect.FileDescriptor

var file_jmartin127_traffic_v1_traffic_proto_rawDesc = []byte{
	0x0a, 0x23, 0x6a, 0x6d, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x31, 0x32, 0x37, 0x2f, 0x74, 0x72, 0x61,
	0x66, 0x66, 0x69, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x6a, 0x6d, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x31, 0x32,
	0x37, 0x2e, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32,
	0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x7e, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x76, 0x65, 0x6c,
	0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x0d, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0d, 0x6f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x37, 0x0a, 0x12, 0x64, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52,
	0x12, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x22, 0x4f, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x76, 0x65, 0x6c,
	0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x39, 0x0a, 0x0a, 0x74, 0x72, 0x61,
	0x76, 0x65, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x74, 0x72, 0x61, 0x76, 0x65, 0x6c,
	0x54, 0x69, 0x6d, 0x65, 0x32, 0xd3, 0x02, 0x0a, 0x07, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63,
	0x12, 0xc7, 0x02, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x2b, 0x2e, 0x6a, 0x6d, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x31, 0x32, 0x37, 0x2e,
	0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72,
	0x61, 0x76, 0x65, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x29, 0x2e, 0x6a, 0x6d, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x31, 0x32, 0x37, 0x2e, 0x74, 0x72, 0x61,
	0x66, 0x66, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x76, 0x65,
	0x6c, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0xdd, 0x01, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x2f, 0x74, 0x72,
	0x61, 0x76, 0x65, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x92, 0x41, 0xbe, 0x01, 0x12, 0x3b, 0x52, 0x65,
	0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x20, 0x74, 0x68, 0x65, 0x20, 0x74, 0x72, 0x61, 0x76, 0x65,
	0x6c, 0x20, 0x74, 0x69, 0x6d, 0x65, 0x20, 0x62, 0x65, 0x74, 0x77, 0x65, 0x65, 0x6e, 0x20, 0x74,
	0x68, 0x65, 0x20, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x64, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x7f, 0x55, 0x73, 0x65, 0x20, 0x74,
	0x68, 0x69, 0x73, 0x20, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x20, 0x74, 0x6f, 0x20,
	0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x20, 0x74, 0x68, 0x65, 0x20, 0x74, 0x72,
	0x61, 0x76, 0x65, 0x6c, 0x20, 0x74, 0x69, 0x6d, 0x65, 0x20, 0x62, 0x65, 0x74, 0x77, 0x65, 0x65,
	0x6e, 0x20, 0x61, 0x6e, 0x20, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x20, 0x61, 0x6e, 0x64, 0x20,
	0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x20, 0x42, 0x6f, 0x74,
	0x68, 0x20, 0x74, 0x68, 0x65, 0x20, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x20, 0x61, 0x6e, 0x64,
	0x20, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x61, 0x72, 0x65,
	0x20, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x2e, 0x42, 0x44, 0x5a, 0x42, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x6d, 0x61, 0x72, 0x74, 0x69, 0x6e,
	0x31, 0x32, 0x37, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x6a, 0x6d, 0x61, 0x72, 0x74,
	0x69, 0x6e, 0x31, 0x32, 0x37, 0x2f, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x2f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_jmartin127_traffic_v1_traffic_proto_rawDescOnce sync.Once
	file_jmartin127_traffic_v1_traffic_proto_rawDescData = file_jmartin127_traffic_v1_traffic_proto_rawDesc
)

func file_jmartin127_traffic_v1_traffic_proto_rawDescGZIP() []byte {
	file_jmartin127_traffic_v1_traffic_proto_rawDescOnce.Do(func() {
		file_jmartin127_traffic_v1_traffic_proto_rawDescData = protoimpl.X.CompressGZIP(file_jmartin127_traffic_v1_traffic_proto_rawDescData)
	})
	return file_jmartin127_traffic_v1_traffic_proto_rawDescData
}

var file_jmartin127_traffic_v1_traffic_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_jmartin127_traffic_v1_traffic_proto_goTypes = []interface{}{
	(*GetTravelTimeRequest)(nil), // 0: jmartin127.traffic.v1.GetTravelTimeRequest
	(*GetTravelTimeReply)(nil),   // 1: jmartin127.traffic.v1.GetTravelTimeReply
	(*duration.Duration)(nil),    // 2: google.protobuf.Duration
}
var file_jmartin127_traffic_v1_traffic_proto_depIdxs = []int32{
	2, // 0: jmartin127.traffic.v1.GetTravelTimeReply.travelTime:type_name -> google.protobuf.Duration
	0, // 1: jmartin127.traffic.v1.Traffic.GetTravelTime:input_type -> jmartin127.traffic.v1.GetTravelTimeRequest
	1, // 2: jmartin127.traffic.v1.Traffic.GetTravelTime:output_type -> jmartin127.traffic.v1.GetTravelTimeReply
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_jmartin127_traffic_v1_traffic_proto_init() }
func file_jmartin127_traffic_v1_traffic_proto_init() {
	if File_jmartin127_traffic_v1_traffic_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_jmartin127_traffic_v1_traffic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTravelTimeRequest); i {
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
		file_jmartin127_traffic_v1_traffic_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTravelTimeReply); i {
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
			RawDescriptor: file_jmartin127_traffic_v1_traffic_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_jmartin127_traffic_v1_traffic_proto_goTypes,
		DependencyIndexes: file_jmartin127_traffic_v1_traffic_proto_depIdxs,
		MessageInfos:      file_jmartin127_traffic_v1_traffic_proto_msgTypes,
	}.Build()
	File_jmartin127_traffic_v1_traffic_proto = out.File
	file_jmartin127_traffic_v1_traffic_proto_rawDesc = nil
	file_jmartin127_traffic_v1_traffic_proto_goTypes = nil
	file_jmartin127_traffic_v1_traffic_proto_depIdxs = nil
}
