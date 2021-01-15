// Code generated by protoc-gen-go. DO NOT EDIT.
// source: traffic/traffic.proto

package traffic

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetTrafficRequest struct {
	Origin               string   `protobuf:"bytes,1,opt,name=origin,proto3" json:"origin,omitempty"`
	Destination          string   `protobuf:"bytes,2,opt,name=destination,proto3" json:"destination,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTrafficRequest) Reset()         { *m = GetTrafficRequest{} }
func (m *GetTrafficRequest) String() string { return proto.CompactTextString(m) }
func (*GetTrafficRequest) ProtoMessage()    {}
func (*GetTrafficRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b91e267bbb4bdf8, []int{0}
}

func (m *GetTrafficRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTrafficRequest.Unmarshal(m, b)
}
func (m *GetTrafficRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTrafficRequest.Marshal(b, m, deterministic)
}
func (m *GetTrafficRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTrafficRequest.Merge(m, src)
}
func (m *GetTrafficRequest) XXX_Size() int {
	return xxx_messageInfo_GetTrafficRequest.Size(m)
}
func (m *GetTrafficRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTrafficRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTrafficRequest proto.InternalMessageInfo

func (m *GetTrafficRequest) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *GetTrafficRequest) GetDestination() string {
	if m != nil {
		return m.Destination
	}
	return ""
}

type GetTrafficReply struct {
	DistanceMiles        float32  `protobuf:"fixed32,1,opt,name=distanceMiles,proto3" json:"distanceMiles,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTrafficReply) Reset()         { *m = GetTrafficReply{} }
func (m *GetTrafficReply) String() string { return proto.CompactTextString(m) }
func (*GetTrafficReply) ProtoMessage()    {}
func (*GetTrafficReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b91e267bbb4bdf8, []int{1}
}

func (m *GetTrafficReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTrafficReply.Unmarshal(m, b)
}
func (m *GetTrafficReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTrafficReply.Marshal(b, m, deterministic)
}
func (m *GetTrafficReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTrafficReply.Merge(m, src)
}
func (m *GetTrafficReply) XXX_Size() int {
	return xxx_messageInfo_GetTrafficReply.Size(m)
}
func (m *GetTrafficReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTrafficReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetTrafficReply proto.InternalMessageInfo

func (m *GetTrafficReply) GetDistanceMiles() float32 {
	if m != nil {
		return m.DistanceMiles
	}
	return 0
}

func init() {
	proto.RegisterType((*GetTrafficRequest)(nil), "traffic.GetTrafficRequest")
	proto.RegisterType((*GetTrafficReply)(nil), "traffic.GetTrafficReply")
}

func init() {
	proto.RegisterFile("traffic/traffic.proto", fileDescriptor_9b91e267bbb4bdf8)
}

var fileDescriptor_9b91e267bbb4bdf8 = []byte{
	// 211 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x29, 0x4a, 0x4c,
	0x4b, 0xcb, 0x4c, 0xd6, 0x87, 0xd2, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xec, 0x50, 0xae,
	0x92, 0x2f, 0x97, 0xa0, 0x7b, 0x6a, 0x49, 0x08, 0x84, 0x17, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c,
	0x22, 0x24, 0xc6, 0xc5, 0x96, 0x5f, 0x94, 0x99, 0x9e, 0x99, 0x27, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1,
	0x19, 0x04, 0xe5, 0x09, 0x29, 0x70, 0x71, 0xa7, 0xa4, 0x16, 0x97, 0x64, 0xe6, 0x25, 0x96, 0x64,
	0xe6, 0xe7, 0x49, 0x30, 0x81, 0x25, 0x91, 0x85, 0x94, 0xcc, 0xb9, 0xf8, 0x91, 0x8d, 0x2b, 0xc8,
	0xa9, 0x14, 0x52, 0xe1, 0xe2, 0x4d, 0xc9, 0x2c, 0x2e, 0x49, 0xcc, 0x4b, 0x4e, 0xf5, 0xcd, 0xcc,
	0x49, 0x2d, 0x06, 0x9b, 0xc9, 0x14, 0x84, 0x2a, 0x68, 0x14, 0xcc, 0xc5, 0x0e, 0xd5, 0x25, 0xe4,
	0xc1, 0xc5, 0x87, 0x30, 0x23, 0x24, 0x33, 0x37, 0x55, 0x48, 0x4a, 0x0f, 0xe6, 0x7a, 0x0c, 0xb7,
	0x4a, 0x49, 0x60, 0x95, 0x2b, 0xc8, 0xa9, 0x54, 0x62, 0x70, 0x32, 0x8c, 0xd2, 0x4f, 0xcf, 0x2c,
	0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0xca, 0x4d, 0x2c, 0x2a, 0xc9, 0xcc, 0x33,
	0x34, 0x32, 0xd7, 0x4f, 0x49, 0x2c, 0xce, 0x48, 0xca, 0x4f, 0x2c, 0x4a, 0xd1, 0x47, 0x0b, 0x9e,
	0x24, 0x36, 0x70, 0xf8, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x1c, 0x39, 0xb9, 0xbf, 0x38,
	0x01, 0x00, 0x00,
}
