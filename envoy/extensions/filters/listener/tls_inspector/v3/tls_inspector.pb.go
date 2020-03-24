// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/filters/listener/tls_inspector/v3/tls_inspector.proto

package envoy_extensions_filters_listener_tls_inspector_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
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

type TlsInspector struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TlsInspector) Reset()         { *m = TlsInspector{} }
func (m *TlsInspector) String() string { return proto.CompactTextString(m) }
func (*TlsInspector) ProtoMessage()    {}
func (*TlsInspector) Descriptor() ([]byte, []int) {
	return fileDescriptor_a303ec5271932b24, []int{0}
}

func (m *TlsInspector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TlsInspector.Unmarshal(m, b)
}
func (m *TlsInspector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TlsInspector.Marshal(b, m, deterministic)
}
func (m *TlsInspector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TlsInspector.Merge(m, src)
}
func (m *TlsInspector) XXX_Size() int {
	return xxx_messageInfo_TlsInspector.Size(m)
}
func (m *TlsInspector) XXX_DiscardUnknown() {
	xxx_messageInfo_TlsInspector.DiscardUnknown(m)
}

var xxx_messageInfo_TlsInspector proto.InternalMessageInfo

func init() {
	proto.RegisterType((*TlsInspector)(nil), "envoy.extensions.filters.listener.tls_inspector.v3.TlsInspector")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/listener/tls_inspector/v3/tls_inspector.proto", fileDescriptor_a303ec5271932b24)
}

var fileDescriptor_a303ec5271932b24 = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x4b, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xad, 0x28, 0x49, 0xcd, 0x2b, 0xce, 0xcc, 0xcf, 0x2b, 0xd6, 0x4f, 0xcb, 0xcc,
	0x29, 0x49, 0x2d, 0x2a, 0xd6, 0xcf, 0xc9, 0x2c, 0x2e, 0x49, 0xcd, 0x4b, 0x2d, 0xd2, 0x2f, 0xc9,
	0x29, 0x8e, 0xcf, 0xcc, 0x2b, 0x2e, 0x48, 0x4d, 0x2e, 0xc9, 0x2f, 0xd2, 0x2f, 0x33, 0x46, 0x15,
	0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x32, 0x02, 0x9b, 0xa3, 0x87, 0x30, 0x47, 0x0f, 0x6a,
	0x8e, 0x1e, 0xcc, 0x1c, 0x3d, 0x54, 0x6d, 0x65, 0xc6, 0x52, 0xb2, 0xa5, 0x29, 0x05, 0x89, 0xfa,
	0x89, 0x79, 0x79, 0xf9, 0x25, 0x89, 0x25, 0x60, 0xbb, 0x8b, 0x4b, 0x12, 0x4b, 0x4a, 0x8b, 0x21,
	0x46, 0x4a, 0x29, 0x62, 0x48, 0x97, 0xa5, 0x16, 0x81, 0xcc, 0xce, 0xcc, 0x4b, 0x87, 0x28, 0x51,
	0x0a, 0xe4, 0xe2, 0x09, 0xc9, 0x29, 0xf6, 0x84, 0x19, 0x6a, 0xe5, 0x38, 0xeb, 0x68, 0x87, 0x9c,
	0x0d, 0x97, 0x15, 0xc4, 0x31, 0xc9, 0xf9, 0x79, 0x69, 0x99, 0xe9, 0x50, 0x87, 0xe0, 0x74, 0x87,
	0x91, 0x1e, 0xb2, 0x11, 0x4e, 0xf1, 0xbb, 0x1a, 0x4e, 0x5c, 0x64, 0x63, 0x12, 0x60, 0xe2, 0x72,
	0xc8, 0xcc, 0xd7, 0x03, 0x1b, 0x54, 0x50, 0x94, 0x5f, 0x51, 0xa9, 0x47, 0xba, 0x07, 0x9d, 0x04,
	0x91, 0x4d, 0x0e, 0x00, 0xb9, 0x38, 0x80, 0x31, 0x89, 0x0d, 0xec, 0x74, 0x63, 0x40, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x6c, 0x95, 0xcb, 0x09, 0x7a, 0x01, 0x00, 0x00,
}
