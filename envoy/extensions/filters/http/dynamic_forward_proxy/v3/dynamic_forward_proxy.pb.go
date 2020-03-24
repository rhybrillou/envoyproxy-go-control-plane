// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/filters/http/dynamic_forward_proxy/v3/dynamic_forward_proxy.proto

package envoy_extensions_filters_http_dynamic_forward_proxy_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/common/dynamic_forward_proxy/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type FilterConfig struct {
	DnsCacheConfig       *v3.DnsCacheConfig `protobuf:"bytes,1,opt,name=dns_cache_config,json=dnsCacheConfig,proto3" json:"dns_cache_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *FilterConfig) Reset()         { *m = FilterConfig{} }
func (m *FilterConfig) String() string { return proto.CompactTextString(m) }
func (*FilterConfig) ProtoMessage()    {}
func (*FilterConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_919f718544b11cb2, []int{0}
}

func (m *FilterConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterConfig.Unmarshal(m, b)
}
func (m *FilterConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterConfig.Marshal(b, m, deterministic)
}
func (m *FilterConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterConfig.Merge(m, src)
}
func (m *FilterConfig) XXX_Size() int {
	return xxx_messageInfo_FilterConfig.Size(m)
}
func (m *FilterConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterConfig.DiscardUnknown(m)
}

var xxx_messageInfo_FilterConfig proto.InternalMessageInfo

func (m *FilterConfig) GetDnsCacheConfig() *v3.DnsCacheConfig {
	if m != nil {
		return m.DnsCacheConfig
	}
	return nil
}

type PerRouteConfig struct {
	// Types that are valid to be assigned to HostRewriteSpecifier:
	//	*PerRouteConfig_HostRewriteLiteral
	//	*PerRouteConfig_HostRewriteHeader
	HostRewriteSpecifier isPerRouteConfig_HostRewriteSpecifier `protobuf_oneof:"host_rewrite_specifier"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *PerRouteConfig) Reset()         { *m = PerRouteConfig{} }
func (m *PerRouteConfig) String() string { return proto.CompactTextString(m) }
func (*PerRouteConfig) ProtoMessage()    {}
func (*PerRouteConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_919f718544b11cb2, []int{1}
}

func (m *PerRouteConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PerRouteConfig.Unmarshal(m, b)
}
func (m *PerRouteConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PerRouteConfig.Marshal(b, m, deterministic)
}
func (m *PerRouteConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PerRouteConfig.Merge(m, src)
}
func (m *PerRouteConfig) XXX_Size() int {
	return xxx_messageInfo_PerRouteConfig.Size(m)
}
func (m *PerRouteConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_PerRouteConfig.DiscardUnknown(m)
}

var xxx_messageInfo_PerRouteConfig proto.InternalMessageInfo

type isPerRouteConfig_HostRewriteSpecifier interface {
	isPerRouteConfig_HostRewriteSpecifier()
}

type PerRouteConfig_HostRewriteLiteral struct {
	HostRewriteLiteral string `protobuf:"bytes,1,opt,name=host_rewrite_literal,json=hostRewriteLiteral,proto3,oneof"`
}

type PerRouteConfig_HostRewriteHeader struct {
	HostRewriteHeader string `protobuf:"bytes,2,opt,name=host_rewrite_header,json=hostRewriteHeader,proto3,oneof"`
}

func (*PerRouteConfig_HostRewriteLiteral) isPerRouteConfig_HostRewriteSpecifier() {}

func (*PerRouteConfig_HostRewriteHeader) isPerRouteConfig_HostRewriteSpecifier() {}

func (m *PerRouteConfig) GetHostRewriteSpecifier() isPerRouteConfig_HostRewriteSpecifier {
	if m != nil {
		return m.HostRewriteSpecifier
	}
	return nil
}

func (m *PerRouteConfig) GetHostRewriteLiteral() string {
	if x, ok := m.GetHostRewriteSpecifier().(*PerRouteConfig_HostRewriteLiteral); ok {
		return x.HostRewriteLiteral
	}
	return ""
}

func (m *PerRouteConfig) GetHostRewriteHeader() string {
	if x, ok := m.GetHostRewriteSpecifier().(*PerRouteConfig_HostRewriteHeader); ok {
		return x.HostRewriteHeader
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*PerRouteConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*PerRouteConfig_HostRewriteLiteral)(nil),
		(*PerRouteConfig_HostRewriteHeader)(nil),
	}
}

func init() {
	proto.RegisterType((*FilterConfig)(nil), "envoy.extensions.filters.http.dynamic_forward_proxy.v3.FilterConfig")
	proto.RegisterType((*PerRouteConfig)(nil), "envoy.extensions.filters.http.dynamic_forward_proxy.v3.PerRouteConfig")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/http/dynamic_forward_proxy/v3/dynamic_forward_proxy.proto", fileDescriptor_919f718544b11cb2)
}

var fileDescriptor_919f718544b11cb2 = []byte{
	// 405 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcf, 0x4a, 0xeb, 0x40,
	0x14, 0xc6, 0x6f, 0x02, 0xb7, 0xdc, 0x3b, 0x4a, 0xa9, 0x51, 0x34, 0x14, 0x14, 0xed, 0xca, 0xd5,
	0x44, 0x5a, 0x70, 0xd1, 0x55, 0x49, 0xff, 0x58, 0xa4, 0x8b, 0x90, 0x17, 0x08, 0x63, 0x32, 0x69,
	0x46, 0xd2, 0x99, 0x30, 0x33, 0x4d, 0xdb, 0x9d, 0x4b, 0xf1, 0x11, 0x7c, 0x14, 0xf7, 0x82, 0x4b,
	0x7d, 0x01, 0x1f, 0xc4, 0x95, 0x24, 0x13, 0xb5, 0xa1, 0xb5, 0x8b, 0xee, 0xc2, 0xf9, 0xe5, 0x7c,
	0xdf, 0xf9, 0xce, 0x1c, 0xe0, 0x62, 0x9a, 0xb2, 0x85, 0x85, 0xe7, 0x12, 0x53, 0x41, 0x18, 0x15,
	0x56, 0x48, 0x62, 0x89, 0xb9, 0xb0, 0x22, 0x29, 0x13, 0x2b, 0x58, 0x50, 0x34, 0x21, 0xbe, 0x17,
	0x32, 0x3e, 0x43, 0x3c, 0xf0, 0x12, 0xce, 0xe6, 0x0b, 0x2b, 0x6d, 0xad, 0x07, 0x30, 0xe1, 0x4c,
	0x32, 0xe3, 0x32, 0xd7, 0x84, 0x3f, 0x9a, 0xb0, 0xd0, 0x84, 0x99, 0x26, 0x5c, 0xdf, 0x9a, 0xb6,
	0xea, 0x9d, 0x95, 0x59, 0x7c, 0x36, 0x99, 0x30, 0xba, 0x61, 0x0a, 0x2a, 0x3c, 0x1f, 0xf9, 0x11,
	0x56, 0xce, 0xf5, 0xe3, 0x69, 0x90, 0x20, 0x0b, 0x51, 0xca, 0x24, 0x92, 0xb9, 0x82, 0x90, 0x48,
	0x4e, 0x45, 0x81, 0xcf, 0x56, 0x70, 0x8a, 0x79, 0xe6, 0x44, 0xe8, 0xb8, 0xf8, 0xe5, 0x28, 0x45,
	0x31, 0x09, 0x90, 0xc4, 0xd6, 0xd7, 0x87, 0x02, 0x8d, 0x57, 0x0d, 0xec, 0x0e, 0xf2, 0x18, 0x5d,
	0x46, 0x43, 0x32, 0x36, 0x24, 0xa8, 0x7d, 0xdb, 0x7b, 0x7e, 0x5e, 0x33, 0xb5, 0x53, 0xed, 0x7c,
	0xa7, 0xd9, 0x81, 0x2b, 0x0b, 0x50, 0x41, 0x7e, 0x8d, 0x0e, 0x7b, 0x54, 0x74, 0x33, 0x21, 0xa5,
	0x6d, 0xff, 0xfb, 0xb0, 0xff, 0x3e, 0x68, 0x7a, 0x4d, 0x73, 0xab, 0x41, 0x89, 0xb4, 0xaf, 0x1f,
	0x9f, 0xef, 0x4f, 0xfa, 0xa0, 0xab, 0x1c, 0x94, 0x6d, 0xb1, 0xde, 0x8d, 0xdb, 0x6d, 0xa2, 0x38,
	0x89, 0x10, 0x5c, 0x4e, 0xd0, 0x78, 0xd7, 0x40, 0xd5, 0xc1, 0xdc, 0x65, 0x53, 0x59, 0xc8, 0x1b,
	0x4d, 0x70, 0x10, 0x31, 0x21, 0x3d, 0x8e, 0x67, 0x9c, 0x48, 0xec, 0xc5, 0x44, 0x62, 0x8e, 0xe2,
	0x3c, 0xd8, 0xff, 0xe1, 0x1f, 0xd7, 0xc8, 0xa8, 0xab, 0xe0, 0x48, 0x31, 0xe3, 0x02, 0xec, 0x97,
	0x7a, 0x22, 0x8c, 0x02, 0xcc, 0x4d, 0xbd, 0x68, 0xd9, 0x5b, 0x6a, 0x19, 0xe6, 0xa8, 0x3d, 0xca,
	0x42, 0x5c, 0x81, 0xfe, 0x96, 0x21, 0xca, 0x33, 0xdb, 0x26, 0x38, 0x2c, 0xf9, 0x8b, 0x04, 0xfb,
	0x24, 0x24, 0x98, 0xdb, 0xb7, 0x4f, 0x77, 0x2f, 0x6f, 0x15, 0xbd, 0xa6, 0x83, 0x1e, 0x61, 0xea,
	0x51, 0x94, 0xd2, 0x76, 0x07, 0x6a, 0x9b, 0x3d, 0x45, 0x06, 0x0a, 0x38, 0x59, 0xdd, 0xc9, 0xae,
	0xc3, 0xd1, 0x6e, 0x2a, 0xf9, 0x99, 0xb4, 0x3e, 0x03, 0x00, 0x00, 0xff, 0xff, 0xc9, 0x25, 0x80,
	0x35, 0x51, 0x03, 0x00, 0x00,
}
