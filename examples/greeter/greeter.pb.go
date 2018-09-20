// Code generated by protoc-gen-go. DO NOT EDIT.
// source: examples/greeter/greeter.proto

package greeter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Request struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_greeter_a33cad897a59fda9, []int{0}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "greeter.Request")
}

func init() {
	proto.RegisterFile("examples/greeter/greeter.proto", fileDescriptor_greeter_a33cad897a59fda9)
}

var fileDescriptor_greeter_a33cad897a59fda9 = []byte{
	// 132 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0x2d, 0xd6, 0x4f, 0x2f, 0x4a, 0x4d, 0x2d, 0x49, 0x2d, 0x82, 0xd1, 0x7a, 0x05,
	0x45, 0xf9, 0x25, 0xf9, 0x42, 0xec, 0x50, 0xae, 0x94, 0x3c, 0x86, 0xc2, 0xa2, 0xd4, 0xe2, 0x82,
	0xfc, 0xbc, 0xe2, 0x54, 0x88, 0x4a, 0x25, 0x59, 0x2e, 0xf6, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2,
	0x12, 0x21, 0x21, 0x2e, 0x96, 0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20,
	0x30, 0xdb, 0xc8, 0x92, 0x8b, 0xdd, 0x1d, 0xa2, 0x51, 0x48, 0x8f, 0x8b, 0xd5, 0x23, 0x35, 0x27,
	0x27, 0x5f, 0x48, 0x40, 0x0f, 0x66, 0x19, 0x54, 0xa7, 0x94, 0x20, 0x92, 0x08, 0xc4, 0x74, 0x25,
	0x86, 0x24, 0x36, 0xb0, 0x05, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x74, 0x5e, 0xe0, 0xeb,
	0xac, 0x00, 0x00, 0x00,
}
