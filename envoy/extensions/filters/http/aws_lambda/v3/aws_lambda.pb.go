// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/filters/http/aws_lambda/v3/aws_lambda.proto

package envoy_extensions_filters_http_aws_lambda_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
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

type Config struct {
	Arn                  string   `protobuf:"bytes,1,opt,name=arn,proto3" json:"arn,omitempty"`
	PayloadPassthrough   bool     `protobuf:"varint,2,opt,name=payload_passthrough,json=payloadPassthrough,proto3" json:"payload_passthrough,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_a22f80ea45c04127, []int{0}
}

func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetArn() string {
	if m != nil {
		return m.Arn
	}
	return ""
}

func (m *Config) GetPayloadPassthrough() bool {
	if m != nil {
		return m.PayloadPassthrough
	}
	return false
}

type PerRouteConfig struct {
	InvokeConfig         *Config  `protobuf:"bytes,1,opt,name=invoke_config,json=invokeConfig,proto3" json:"invoke_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PerRouteConfig) Reset()         { *m = PerRouteConfig{} }
func (m *PerRouteConfig) String() string { return proto.CompactTextString(m) }
func (*PerRouteConfig) ProtoMessage()    {}
func (*PerRouteConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_a22f80ea45c04127, []int{1}
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

func (m *PerRouteConfig) GetInvokeConfig() *Config {
	if m != nil {
		return m.InvokeConfig
	}
	return nil
}

func init() {
	proto.RegisterType((*Config)(nil), "envoy.extensions.filters.http.aws_lambda.v3.Config")
	proto.RegisterType((*PerRouteConfig)(nil), "envoy.extensions.filters.http.aws_lambda.v3.PerRouteConfig")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/http/aws_lambda/v3/aws_lambda.proto", fileDescriptor_a22f80ea45c04127)
}

var fileDescriptor_a22f80ea45c04127 = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xb1, 0x4e, 0xfb, 0x30,
	0x10, 0xc6, 0xe5, 0xfe, 0xff, 0x2a, 0x60, 0xa0, 0x42, 0x61, 0xa0, 0x74, 0x40, 0xa5, 0x53, 0x25,
	0x24, 0x5b, 0x6a, 0x58, 0x5a, 0x75, 0x69, 0x58, 0x41, 0x8a, 0x32, 0xb1, 0x55, 0x57, 0xe2, 0x36,
	0x16, 0xc1, 0x67, 0xd9, 0x6e, 0xda, 0xbe, 0x01, 0x1b, 0x3b, 0xef, 0xc0, 0x63, 0xf0, 0x52, 0x4c,
	0x28, 0x71, 0x50, 0x8b, 0x98, 0xba, 0xf9, 0x7c, 0xfe, 0xbe, 0xfb, 0x7d, 0x3e, 0x3a, 0x16, 0xaa,
	0xc0, 0x0d, 0x17, 0x6b, 0x27, 0x94, 0x95, 0xa8, 0x2c, 0x9f, 0xcb, 0xdc, 0x09, 0x63, 0x79, 0xe6,
	0x9c, 0xe6, 0xb0, 0xb2, 0xd3, 0x1c, 0x5e, 0x66, 0x29, 0xf0, 0x22, 0xdc, 0xa9, 0x98, 0x36, 0xe8,
	0x30, 0xb8, 0xa9, 0xd4, 0x6c, 0xab, 0x66, 0xb5, 0x9a, 0x95, 0x6a, 0xb6, 0xf3, 0xbe, 0x08, 0x3b,
	0xd7, 0xcb, 0x54, 0x03, 0x07, 0xa5, 0xd0, 0x81, 0xab, 0x46, 0x15, 0xc2, 0x94, 0x2a, 0xa9, 0x16,
	0xde, 0xaf, 0x73, 0x51, 0x40, 0x2e, 0x53, 0x70, 0x82, 0xff, 0x1c, 0x7c, 0xa3, 0xf7, 0x46, 0x68,
	0xf3, 0x0e, 0xd5, 0x5c, 0x2e, 0x82, 0x4b, 0xfa, 0x0f, 0x8c, 0x6a, 0x93, 0x2e, 0xe9, 0x1f, 0x45,
	0x07, 0x5f, 0xd1, 0x7f, 0xd3, 0x38, 0x23, 0x49, 0x79, 0x17, 0x70, 0x7a, 0xae, 0x61, 0x93, 0x23,
	0xa4, 0x53, 0x0d, 0xd6, 0xba, 0xcc, 0xe0, 0x72, 0x91, 0xb5, 0x1b, 0x5d, 0xd2, 0x3f, 0x4c, 0x82,
	0xba, 0x15, 0x6f, 0x3b, 0xa3, 0xe1, 0xfb, 0xe7, 0xeb, 0xd5, 0x2d, 0x1d, 0xf8, 0x18, 0x4f, 0xd5,
	0x80, 0x3a, 0xc2, 0xdf, 0x04, 0x03, 0xc8, 0x75, 0x06, 0xcc, 0x63, 0xf4, 0x3e, 0x08, 0x6d, 0xc5,
	0xc2, 0x24, 0xb8, 0x74, 0xa2, 0x26, 0x7b, 0xa4, 0xa7, 0x52, 0x15, 0xf8, 0x2c, 0xa6, 0xde, 0xa9,
	0x62, 0x3c, 0x1e, 0x84, 0x6c, 0x8f, 0x5f, 0xaa, 0xed, 0x93, 0x13, 0xef, 0xe4, 0xab, 0xd1, 0xa4,
	0xe4, 0x1c, 0xd3, 0xd1, 0x3e, 0x9c, 0xbf, 0xe1, 0xa2, 0x07, 0x3a, 0x94, 0xe8, 0x49, 0xb4, 0xc1,
	0xf5, 0x66, 0x1f, 0xa8, 0xa8, 0x35, 0x59, 0xd9, 0xfb, 0xaa, 0x8a, 0xcb, 0x75, 0xc4, 0x64, 0xd6,
	0xac, 0xf6, 0x12, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0xc0, 0x8e, 0x74, 0xd1, 0x40, 0x02, 0x00,
	0x00,
}
