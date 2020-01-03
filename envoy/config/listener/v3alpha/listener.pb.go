// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/listener/v3alpha/listener.proto

package envoy_config_listener_v3alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3alpha "github.com/envoyproxy/go-control-plane/envoy/config/core/v3alpha"
	v2 "github.com/envoyproxy/go-control-plane/envoy/config/listener/v2"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Listener_DrainType int32

const (
	Listener_DEFAULT     Listener_DrainType = 0
	Listener_MODIFY_ONLY Listener_DrainType = 1
)

var Listener_DrainType_name = map[int32]string{
	0: "DEFAULT",
	1: "MODIFY_ONLY",
}

var Listener_DrainType_value = map[string]int32{
	"DEFAULT":     0,
	"MODIFY_ONLY": 1,
}

func (x Listener_DrainType) String() string {
	return proto.EnumName(Listener_DrainType_name, int32(x))
}

func (Listener_DrainType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_45aaa9e70cb8d8e4, []int{0, 0}
}

type Listener struct {
	Name                                string                            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Address                             *v3alpha.Address                  `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	FilterChains                        []*FilterChain                    `protobuf:"bytes,3,rep,name=filter_chains,json=filterChains,proto3" json:"filter_chains,omitempty"`
	HiddenEnvoyDeprecatedUseOriginalDst *wrappers.BoolValue               `protobuf:"bytes,4,opt,name=hidden_envoy_deprecated_use_original_dst,json=hiddenEnvoyDeprecatedUseOriginalDst,proto3" json:"hidden_envoy_deprecated_use_original_dst,omitempty"` // Deprecated: Do not use.
	PerConnectionBufferLimitBytes       *wrappers.UInt32Value             `protobuf:"bytes,5,opt,name=per_connection_buffer_limit_bytes,json=perConnectionBufferLimitBytes,proto3" json:"per_connection_buffer_limit_bytes,omitempty"`
	Metadata                            *v3alpha.Metadata                 `protobuf:"bytes,6,opt,name=metadata,proto3" json:"metadata,omitempty"`
	DeprecatedV1                        *Listener_DeprecatedV1            `protobuf:"bytes,7,opt,name=deprecated_v1,json=deprecatedV1,proto3" json:"deprecated_v1,omitempty"`
	DrainType                           Listener_DrainType                `protobuf:"varint,8,opt,name=drain_type,json=drainType,proto3,enum=envoy.config.listener.v3alpha.Listener_DrainType" json:"drain_type,omitempty"`
	ListenerFilters                     []*ListenerFilter                 `protobuf:"bytes,9,rep,name=listener_filters,json=listenerFilters,proto3" json:"listener_filters,omitempty"`
	ListenerFiltersTimeout              *duration.Duration                `protobuf:"bytes,15,opt,name=listener_filters_timeout,json=listenerFiltersTimeout,proto3" json:"listener_filters_timeout,omitempty"`
	ContinueOnListenerFiltersTimeout    bool                              `protobuf:"varint,17,opt,name=continue_on_listener_filters_timeout,json=continueOnListenerFiltersTimeout,proto3" json:"continue_on_listener_filters_timeout,omitempty"`
	Transparent                         *wrappers.BoolValue               `protobuf:"bytes,10,opt,name=transparent,proto3" json:"transparent,omitempty"`
	Freebind                            *wrappers.BoolValue               `protobuf:"bytes,11,opt,name=freebind,proto3" json:"freebind,omitempty"`
	SocketOptions                       []*v3alpha.SocketOption           `protobuf:"bytes,13,rep,name=socket_options,json=socketOptions,proto3" json:"socket_options,omitempty"`
	TcpFastOpenQueueLength              *wrappers.UInt32Value             `protobuf:"bytes,12,opt,name=tcp_fast_open_queue_length,json=tcpFastOpenQueueLength,proto3" json:"tcp_fast_open_queue_length,omitempty"`
	TrafficDirection                    v3alpha.TrafficDirection          `protobuf:"varint,16,opt,name=traffic_direction,json=trafficDirection,proto3,enum=envoy.config.core.v3alpha.TrafficDirection" json:"traffic_direction,omitempty"`
	UdpListenerConfig                   *UdpListenerConfig                `protobuf:"bytes,18,opt,name=udp_listener_config,json=udpListenerConfig,proto3" json:"udp_listener_config,omitempty"`
	ApiListener                         *v2.ApiListener                   `protobuf:"bytes,19,opt,name=api_listener,json=apiListener,proto3" json:"api_listener,omitempty"`
	ConnectionBalanceConfig             *Listener_ConnectionBalanceConfig `protobuf:"bytes,20,opt,name=connection_balance_config,json=connectionBalanceConfig,proto3" json:"connection_balance_config,omitempty"`
	ReusePort                           bool                              `protobuf:"varint,21,opt,name=reuse_port,json=reusePort,proto3" json:"reuse_port,omitempty"`
	XXX_NoUnkeyedLiteral                struct{}                          `json:"-"`
	XXX_unrecognized                    []byte                            `json:"-"`
	XXX_sizecache                       int32                             `json:"-"`
}

func (m *Listener) Reset()         { *m = Listener{} }
func (m *Listener) String() string { return proto.CompactTextString(m) }
func (*Listener) ProtoMessage()    {}
func (*Listener) Descriptor() ([]byte, []int) {
	return fileDescriptor_45aaa9e70cb8d8e4, []int{0}
}

func (m *Listener) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Listener.Unmarshal(m, b)
}
func (m *Listener) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Listener.Marshal(b, m, deterministic)
}
func (m *Listener) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Listener.Merge(m, src)
}
func (m *Listener) XXX_Size() int {
	return xxx_messageInfo_Listener.Size(m)
}
func (m *Listener) XXX_DiscardUnknown() {
	xxx_messageInfo_Listener.DiscardUnknown(m)
}

var xxx_messageInfo_Listener proto.InternalMessageInfo

func (m *Listener) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Listener) GetAddress() *v3alpha.Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Listener) GetFilterChains() []*FilterChain {
	if m != nil {
		return m.FilterChains
	}
	return nil
}

// Deprecated: Do not use.
func (m *Listener) GetHiddenEnvoyDeprecatedUseOriginalDst() *wrappers.BoolValue {
	if m != nil {
		return m.HiddenEnvoyDeprecatedUseOriginalDst
	}
	return nil
}

func (m *Listener) GetPerConnectionBufferLimitBytes() *wrappers.UInt32Value {
	if m != nil {
		return m.PerConnectionBufferLimitBytes
	}
	return nil
}

func (m *Listener) GetMetadata() *v3alpha.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Listener) GetDeprecatedV1() *Listener_DeprecatedV1 {
	if m != nil {
		return m.DeprecatedV1
	}
	return nil
}

func (m *Listener) GetDrainType() Listener_DrainType {
	if m != nil {
		return m.DrainType
	}
	return Listener_DEFAULT
}

func (m *Listener) GetListenerFilters() []*ListenerFilter {
	if m != nil {
		return m.ListenerFilters
	}
	return nil
}

func (m *Listener) GetListenerFiltersTimeout() *duration.Duration {
	if m != nil {
		return m.ListenerFiltersTimeout
	}
	return nil
}

func (m *Listener) GetContinueOnListenerFiltersTimeout() bool {
	if m != nil {
		return m.ContinueOnListenerFiltersTimeout
	}
	return false
}

func (m *Listener) GetTransparent() *wrappers.BoolValue {
	if m != nil {
		return m.Transparent
	}
	return nil
}

func (m *Listener) GetFreebind() *wrappers.BoolValue {
	if m != nil {
		return m.Freebind
	}
	return nil
}

func (m *Listener) GetSocketOptions() []*v3alpha.SocketOption {
	if m != nil {
		return m.SocketOptions
	}
	return nil
}

func (m *Listener) GetTcpFastOpenQueueLength() *wrappers.UInt32Value {
	if m != nil {
		return m.TcpFastOpenQueueLength
	}
	return nil
}

func (m *Listener) GetTrafficDirection() v3alpha.TrafficDirection {
	if m != nil {
		return m.TrafficDirection
	}
	return v3alpha.TrafficDirection_UNSPECIFIED
}

func (m *Listener) GetUdpListenerConfig() *UdpListenerConfig {
	if m != nil {
		return m.UdpListenerConfig
	}
	return nil
}

func (m *Listener) GetApiListener() *v2.ApiListener {
	if m != nil {
		return m.ApiListener
	}
	return nil
}

func (m *Listener) GetConnectionBalanceConfig() *Listener_ConnectionBalanceConfig {
	if m != nil {
		return m.ConnectionBalanceConfig
	}
	return nil
}

func (m *Listener) GetReusePort() bool {
	if m != nil {
		return m.ReusePort
	}
	return false
}

type Listener_DeprecatedV1 struct {
	BindToPort           *wrappers.BoolValue `protobuf:"bytes,1,opt,name=bind_to_port,json=bindToPort,proto3" json:"bind_to_port,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Listener_DeprecatedV1) Reset()         { *m = Listener_DeprecatedV1{} }
func (m *Listener_DeprecatedV1) String() string { return proto.CompactTextString(m) }
func (*Listener_DeprecatedV1) ProtoMessage()    {}
func (*Listener_DeprecatedV1) Descriptor() ([]byte, []int) {
	return fileDescriptor_45aaa9e70cb8d8e4, []int{0, 0}
}

func (m *Listener_DeprecatedV1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Listener_DeprecatedV1.Unmarshal(m, b)
}
func (m *Listener_DeprecatedV1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Listener_DeprecatedV1.Marshal(b, m, deterministic)
}
func (m *Listener_DeprecatedV1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Listener_DeprecatedV1.Merge(m, src)
}
func (m *Listener_DeprecatedV1) XXX_Size() int {
	return xxx_messageInfo_Listener_DeprecatedV1.Size(m)
}
func (m *Listener_DeprecatedV1) XXX_DiscardUnknown() {
	xxx_messageInfo_Listener_DeprecatedV1.DiscardUnknown(m)
}

var xxx_messageInfo_Listener_DeprecatedV1 proto.InternalMessageInfo

func (m *Listener_DeprecatedV1) GetBindToPort() *wrappers.BoolValue {
	if m != nil {
		return m.BindToPort
	}
	return nil
}

type Listener_ConnectionBalanceConfig struct {
	// Types that are valid to be assigned to BalanceType:
	//	*Listener_ConnectionBalanceConfig_ExactBalance_
	BalanceType          isListener_ConnectionBalanceConfig_BalanceType `protobuf_oneof:"balance_type"`
	XXX_NoUnkeyedLiteral struct{}                                       `json:"-"`
	XXX_unrecognized     []byte                                         `json:"-"`
	XXX_sizecache        int32                                          `json:"-"`
}

func (m *Listener_ConnectionBalanceConfig) Reset()         { *m = Listener_ConnectionBalanceConfig{} }
func (m *Listener_ConnectionBalanceConfig) String() string { return proto.CompactTextString(m) }
func (*Listener_ConnectionBalanceConfig) ProtoMessage()    {}
func (*Listener_ConnectionBalanceConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_45aaa9e70cb8d8e4, []int{0, 1}
}

func (m *Listener_ConnectionBalanceConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Listener_ConnectionBalanceConfig.Unmarshal(m, b)
}
func (m *Listener_ConnectionBalanceConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Listener_ConnectionBalanceConfig.Marshal(b, m, deterministic)
}
func (m *Listener_ConnectionBalanceConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Listener_ConnectionBalanceConfig.Merge(m, src)
}
func (m *Listener_ConnectionBalanceConfig) XXX_Size() int {
	return xxx_messageInfo_Listener_ConnectionBalanceConfig.Size(m)
}
func (m *Listener_ConnectionBalanceConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_Listener_ConnectionBalanceConfig.DiscardUnknown(m)
}

var xxx_messageInfo_Listener_ConnectionBalanceConfig proto.InternalMessageInfo

type isListener_ConnectionBalanceConfig_BalanceType interface {
	isListener_ConnectionBalanceConfig_BalanceType()
}

type Listener_ConnectionBalanceConfig_ExactBalance_ struct {
	ExactBalance *Listener_ConnectionBalanceConfig_ExactBalance `protobuf:"bytes,1,opt,name=exact_balance,json=exactBalance,proto3,oneof"`
}

func (*Listener_ConnectionBalanceConfig_ExactBalance_) isListener_ConnectionBalanceConfig_BalanceType() {
}

func (m *Listener_ConnectionBalanceConfig) GetBalanceType() isListener_ConnectionBalanceConfig_BalanceType {
	if m != nil {
		return m.BalanceType
	}
	return nil
}

func (m *Listener_ConnectionBalanceConfig) GetExactBalance() *Listener_ConnectionBalanceConfig_ExactBalance {
	if x, ok := m.GetBalanceType().(*Listener_ConnectionBalanceConfig_ExactBalance_); ok {
		return x.ExactBalance
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Listener_ConnectionBalanceConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Listener_ConnectionBalanceConfig_ExactBalance_)(nil),
	}
}

type Listener_ConnectionBalanceConfig_ExactBalance struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Listener_ConnectionBalanceConfig_ExactBalance) Reset() {
	*m = Listener_ConnectionBalanceConfig_ExactBalance{}
}
func (m *Listener_ConnectionBalanceConfig_ExactBalance) String() string {
	return proto.CompactTextString(m)
}
func (*Listener_ConnectionBalanceConfig_ExactBalance) ProtoMessage() {}
func (*Listener_ConnectionBalanceConfig_ExactBalance) Descriptor() ([]byte, []int) {
	return fileDescriptor_45aaa9e70cb8d8e4, []int{0, 1, 0}
}

func (m *Listener_ConnectionBalanceConfig_ExactBalance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Listener_ConnectionBalanceConfig_ExactBalance.Unmarshal(m, b)
}
func (m *Listener_ConnectionBalanceConfig_ExactBalance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Listener_ConnectionBalanceConfig_ExactBalance.Marshal(b, m, deterministic)
}
func (m *Listener_ConnectionBalanceConfig_ExactBalance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Listener_ConnectionBalanceConfig_ExactBalance.Merge(m, src)
}
func (m *Listener_ConnectionBalanceConfig_ExactBalance) XXX_Size() int {
	return xxx_messageInfo_Listener_ConnectionBalanceConfig_ExactBalance.Size(m)
}
func (m *Listener_ConnectionBalanceConfig_ExactBalance) XXX_DiscardUnknown() {
	xxx_messageInfo_Listener_ConnectionBalanceConfig_ExactBalance.DiscardUnknown(m)
}

var xxx_messageInfo_Listener_ConnectionBalanceConfig_ExactBalance proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("envoy.config.listener.v3alpha.Listener_DrainType", Listener_DrainType_name, Listener_DrainType_value)
	proto.RegisterType((*Listener)(nil), "envoy.config.listener.v3alpha.Listener")
	proto.RegisterType((*Listener_DeprecatedV1)(nil), "envoy.config.listener.v3alpha.Listener.DeprecatedV1")
	proto.RegisterType((*Listener_ConnectionBalanceConfig)(nil), "envoy.config.listener.v3alpha.Listener.ConnectionBalanceConfig")
	proto.RegisterType((*Listener_ConnectionBalanceConfig_ExactBalance)(nil), "envoy.config.listener.v3alpha.Listener.ConnectionBalanceConfig.ExactBalance")
}

func init() {
	proto.RegisterFile("envoy/config/listener/v3alpha/listener.proto", fileDescriptor_45aaa9e70cb8d8e4)
}

var fileDescriptor_45aaa9e70cb8d8e4 = []byte{
	// 1031 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x95, 0x6f, 0x4f, 0x1b, 0x37,
	0x1c, 0xc7, 0x7b, 0x94, 0x96, 0xe0, 0x24, 0x10, 0xcc, 0x5a, 0xae, 0x11, 0x54, 0x29, 0x65, 0x6a,
	0x3a, 0xd6, 0xcb, 0x08, 0xd5, 0x26, 0x21, 0xb4, 0x8a, 0x10, 0x50, 0x3b, 0xa5, 0x84, 0x5e, 0xa1,
	0x82, 0x47, 0x9e, 0x73, 0xe7, 0x80, 0xb5, 0xc3, 0xf6, 0x6c, 0x1f, 0x2d, 0xda, 0x1b, 0x98, 0xf6,
	0x12, 0xf6, 0x1a, 0xf6, 0x12, 0xf6, 0x74, 0xef, 0x69, 0xda, 0xa3, 0xc9, 0xbe, 0xbb, 0x70, 0xd0,
	0x04, 0x32, 0xed, 0xd9, 0xd9, 0xfe, 0x7d, 0x3f, 0xfe, 0xfd, 0xf3, 0xef, 0xc0, 0xd7, 0x84, 0x9d,
	0xf3, 0x8b, 0x46, 0xc0, 0x59, 0x9f, 0x9e, 0x34, 0x22, 0xaa, 0x34, 0x61, 0x44, 0x36, 0xce, 0xd7,
	0x71, 0x24, 0x4e, 0xf1, 0x60, 0xc3, 0x13, 0x92, 0x6b, 0x0e, 0x97, 0xac, 0xb5, 0x97, 0x58, 0x7b,
	0x83, 0xc3, 0xd4, 0xba, 0xfa, 0xec, 0x0a, 0x2c, 0xe0, 0x92, 0x0c, 0x40, 0x38, 0x0c, 0x25, 0x51,
	0x2a, 0xe1, 0x54, 0x57, 0x46, 0x1b, 0xf6, 0xb0, 0x22, 0xa9, 0xd5, 0xea, 0x08, 0xdf, 0x9a, 0x0d,
	0x2c, 0x28, 0xba, 0xea, 0x5a, 0xf5, 0xbb, 0xf1, 0x02, 0x41, 0x01, 0x3f, 0x13, 0x9c, 0x11, 0xa6,
	0xd5, 0x78, 0xc2, 0x38, 0x14, 0x28, 0x27, 0xb6, 0x81, 0x27, 0xc2, 0xc5, 0x13, 0xce, 0x4f, 0x22,
	0x62, 0x9c, 0x69, 0x60, 0xc6, 0xb8, 0xc6, 0x9a, 0x72, 0x96, 0x61, 0x1f, 0xa7, 0xa7, 0x76, 0xd5,
	0x8b, 0xfb, 0x8d, 0x30, 0x96, 0xd6, 0x60, 0xd4, 0xf9, 0x47, 0x89, 0x85, 0x20, 0x32, 0xd3, 0x3f,
	0x89, 0x43, 0x81, 0xf3, 0xdc, 0xc6, 0x39, 0x91, 0x8a, 0x72, 0x46, 0x59, 0xe6, 0xc0, 0xc2, 0x39,
	0x8e, 0x68, 0x88, 0x35, 0x69, 0x64, 0x1f, 0xc9, 0xc1, 0xf2, 0x9f, 0x15, 0x50, 0xe8, 0xa4, 0x3e,
	0x43, 0x08, 0x26, 0x19, 0x3e, 0x23, 0xae, 0x53, 0x73, 0xea, 0xd3, 0xbe, 0xfd, 0x86, 0xbb, 0x60,
	0x2a, 0x2d, 0x88, 0x3b, 0x51, 0x73, 0xea, 0xc5, 0xe6, 0xb2, 0x77, 0xa5, 0xb2, 0xa6, 0x22, 0x59,
	0x55, 0xbd, 0xad, 0xc4, 0xb2, 0x55, 0xf8, 0xa7, 0x75, 0xef, 0x37, 0x67, 0xa2, 0xe2, 0xf8, 0x99,
	0x18, 0x76, 0x41, 0xb9, 0x4f, 0x23, 0x6d, 0x32, 0x73, 0x8a, 0x29, 0x53, 0xee, 0xdd, 0xda, 0xdd,
	0x7a, 0xb1, 0xf9, 0x95, 0x77, 0x63, 0x9f, 0x78, 0xbb, 0x56, 0xb3, 0x6d, 0x24, 0x7e, 0xa9, 0x7f,
	0xb9, 0x50, 0x30, 0x06, 0xf5, 0x53, 0x1a, 0x86, 0x84, 0x21, 0x4b, 0x40, 0x21, 0x11, 0x92, 0x04,
	0x58, 0x93, 0x10, 0xc5, 0x8a, 0x20, 0x2e, 0xe9, 0x09, 0x65, 0x38, 0x42, 0xa1, 0xd2, 0xee, 0xa4,
	0xf5, 0xbc, 0xea, 0x25, 0x89, 0xf4, 0xb2, 0x44, 0x7a, 0x2d, 0xce, 0xa3, 0x0f, 0x38, 0x8a, 0x49,
	0x6b, 0xc2, 0x75, 0xfc, 0xa7, 0x09, 0x6f, 0xc7, 0xe0, 0xda, 0x03, 0xda, 0xa1, 0x22, 0xdd, 0x94,
	0xd5, 0x56, 0x1a, 0xf6, 0xc1, 0x13, 0x91, 0x94, 0x97, 0x91, 0xc0, 0x64, 0x1b, 0xf5, 0xe2, 0x7e,
	0x9f, 0x48, 0x14, 0xd1, 0x33, 0xaa, 0x51, 0xef, 0x42, 0x13, 0xe5, 0xde, 0xb3, 0xf7, 0x2d, 0x7e,
	0x76, 0xdf, 0xe1, 0x1b, 0xa6, 0xd7, 0x9b, 0xf6, 0x46, 0x7f, 0x49, 0x10, 0xb9, 0x3d, 0xa0, 0xb4,
	0x2c, 0xa4, 0x63, 0x18, 0x2d, 0x83, 0x80, 0xaf, 0x40, 0xe1, 0x8c, 0x68, 0x1c, 0x62, 0x8d, 0xdd,
	0xfb, 0x16, 0xf7, 0xf4, 0x86, 0xc4, 0xbf, 0x4d, 0x4d, 0xfd, 0x81, 0x08, 0x1e, 0x83, 0x72, 0x2e,
	0x25, 0xe7, 0x6b, 0xee, 0x94, 0xa5, 0xbc, 0xbc, 0x25, 0xe1, 0x59, 0x33, 0x78, 0x97, 0x19, 0xf8,
	0xb0, 0xe6, 0x97, 0xc2, 0xdc, 0x0a, 0xee, 0x03, 0x10, 0x4a, 0x4c, 0x19, 0xd2, 0x17, 0x82, 0xb8,
	0x85, 0x9a, 0x53, 0x9f, 0x69, 0xae, 0x8d, 0xcd, 0x35, 0xca, 0x83, 0x0b, 0x41, 0xfc, 0xe9, 0x30,
	0xfb, 0x84, 0x47, 0xa0, 0x32, 0x78, 0x39, 0x49, 0x95, 0x95, 0x3b, 0x6d, 0x1b, 0xe4, 0xc5, 0x98,
	0xdc, 0xa4, 0x51, 0xfc, 0xd9, 0xe8, 0xca, 0x5a, 0xc1, 0xf7, 0xc0, 0xbd, 0x4e, 0x46, 0x9a, 0x9e,
	0x11, 0x1e, 0x6b, 0x77, 0xd6, 0x66, 0xe4, 0xd1, 0x67, 0x65, 0x6a, 0xa7, 0xef, 0xcf, 0x7f, 0x78,
	0x8d, 0x76, 0x90, 0x08, 0xe1, 0x1e, 0x58, 0x09, 0x38, 0xd3, 0x94, 0xc5, 0x04, 0x71, 0x86, 0x46,
	0x5e, 0x30, 0x57, 0x73, 0xea, 0x05, 0xbf, 0x96, 0xd9, 0x76, 0x59, 0x67, 0x38, 0x6f, 0x13, 0x14,
	0xb5, 0xc4, 0x4c, 0x09, 0x2c, 0x09, 0xd3, 0x2e, 0xb8, 0xad, 0x5d, 0xfd, 0xbc, 0x39, 0xfc, 0x16,
	0x14, 0xfa, 0x92, 0x90, 0x1e, 0x65, 0xa1, 0x5b, 0xbc, 0x55, 0x3a, 0xb0, 0x85, 0x7b, 0x60, 0x46,
	0xf1, 0xe0, 0x27, 0xa2, 0x11, 0x17, 0x76, 0x6e, 0xb8, 0x65, 0x9b, 0xf2, 0x67, 0x37, 0x34, 0xda,
	0x7b, 0x2b, 0xe8, 0x5a, 0x7b, 0xbf, 0xac, 0x72, 0x2b, 0x05, 0x8f, 0x40, 0x55, 0x07, 0x02, 0xf5,
	0xb1, 0x32, 0x44, 0xc2, 0xd0, 0xcf, 0x31, 0x89, 0x09, 0x8a, 0x08, 0x3b, 0xd1, 0xa7, 0x6e, 0x69,
	0x8c, 0x37, 0xf1, 0x50, 0x07, 0x62, 0x17, 0x2b, 0xdd, 0x15, 0x84, 0xbd, 0x33, 0xe2, 0x8e, 0xd5,
	0xc2, 0x23, 0x30, 0xa7, 0x25, 0xee, 0xf7, 0x69, 0x80, 0x42, 0x2a, 0x93, 0x17, 0xe3, 0x56, 0x6c,
	0xdf, 0xad, 0xde, 0xe0, 0xec, 0x41, 0xa2, 0x69, 0x67, 0x12, 0xbf, 0xa2, 0xaf, 0xed, 0xc0, 0x1f,
	0xc1, 0xfc, 0x90, 0xb1, 0xed, 0x42, 0xeb, 0xec, 0x37, 0xb7, 0xf4, 0xde, 0x61, 0x28, 0xb2, 0x82,
	0x6e, 0x5b, 0x13, 0x7f, 0x2e, 0xbe, 0xbe, 0x05, 0x5f, 0x83, 0x52, 0xfe, 0x1f, 0xe4, 0xce, 0x5b,
	0xf4, 0x97, 0xa3, 0xd0, 0x4d, 0x6f, 0x4b, 0xd0, 0x0c, 0xe1, 0x17, 0xf1, 0xe5, 0x02, 0xfe, 0x02,
	0x1e, 0xe5, 0xc7, 0x0e, 0x8e, 0x30, 0x0b, 0x48, 0xe6, 0xf1, 0x17, 0x16, 0xfb, 0x6a, 0xdc, 0x57,
	0x98, 0x9b, 0x3c, 0x09, 0x27, 0x0d, 0x60, 0x21, 0x18, 0x7e, 0x00, 0x97, 0x00, 0x90, 0xc4, 0x0c,
	0x56, 0xc1, 0xa5, 0x76, 0x1f, 0xd8, 0xc6, 0x9e, 0xb6, 0x3b, 0xfb, 0x5c, 0xea, 0xea, 0x47, 0x50,
	0xca, 0x0f, 0x0c, 0xb8, 0x09, 0x4a, 0xa6, 0xc7, 0x90, 0xe6, 0x89, 0xc0, 0xb9, 0xb5, 0x2f, 0x81,
	0xb1, 0x3f, 0xe0, 0x86, 0xb6, 0xf1, 0xfc, 0xf7, 0xbf, 0x7e, 0x7d, 0xbc, 0x02, 0x96, 0x93, 0x60,
	0xb0, 0xa0, 0x26, 0x2f, 0x43, 0x27, 0x53, 0xf5, 0x8f, 0x09, 0xb0, 0x30, 0x22, 0x18, 0xa8, 0x40,
	0x99, 0x7c, 0xc2, 0x81, 0xce, 0x72, 0x95, 0x7a, 0xd1, 0xf9, 0x9f, 0x49, 0xf2, 0x76, 0x0c, 0x34,
	0xdd, 0x7a, 0x7d, 0xc7, 0x2f, 0x91, 0xdc, 0xba, 0xfa, 0x0e, 0x94, 0xf2, 0xe7, 0x1b, 0x5b, 0x26,
	0x96, 0x4d, 0xb0, 0x31, 0x3c, 0x96, 0x71, 0xae, 0xd8, 0x78, 0x69, 0x10, 0x0d, 0xf0, 0xe2, 0x3f,
	0x21, 0x5a, 0xf3, 0xa0, 0x94, 0xf5, 0x88, 0x99, 0xd3, 0xf0, 0xee, 0xdf, 0x2d, 0x67, 0xf9, 0x39,
	0x98, 0x1e, 0x0c, 0x60, 0x58, 0x04, 0x53, 0xed, 0x9d, 0xdd, 0xad, 0xc3, 0xce, 0x41, 0xe5, 0x0e,
	0x9c, 0x05, 0xc5, 0xb7, 0xdd, 0xf6, 0x9b, 0xdd, 0x63, 0xd4, 0xdd, 0xeb, 0x1c, 0x57, 0x9c, 0x8d,
	0x45, 0x73, 0xeb, 0x02, 0x78, 0x30, 0xf4, 0xd6, 0x1f, 0x26, 0x0b, 0x33, 0x95, 0xd9, 0xd6, 0xf7,
	0x60, 0x95, 0xf2, 0x24, 0x9d, 0x42, 0xf2, 0x4f, 0x17, 0x37, 0x67, 0xb6, 0x55, 0xce, 0xe4, 0xfb,
	0xa6, 0xfe, 0xfb, 0x4e, 0xef, 0xbe, 0x6d, 0x84, 0xf5, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xc0,
	0x9f, 0x17, 0x2d, 0x5c, 0x0a, 0x00, 0x00,
}
