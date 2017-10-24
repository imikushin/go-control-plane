// Code generated by protoc-gen-go.
// source: api/bootstrap.proto
// DO NOT EDIT!

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/duration"
import google_protobuf1 "github.com/golang/protobuf/ptypes/struct"
import google_protobuf2 "github.com/golang/protobuf/ptypes/wrappers"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type LightstepConfig struct {
	CollectorCluster string `protobuf:"bytes,1,opt,name=collector_cluster,json=collectorCluster" json:"collector_cluster,omitempty"`
	AccessTokenFile  string `protobuf:"bytes,2,opt,name=access_token_file,json=accessTokenFile" json:"access_token_file,omitempty"`
}

func (m *LightstepConfig) Reset()                    { *m = LightstepConfig{} }
func (m *LightstepConfig) String() string            { return proto.CompactTextString(m) }
func (*LightstepConfig) ProtoMessage()               {}
func (*LightstepConfig) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *LightstepConfig) GetCollectorCluster() string {
	if m != nil {
		return m.CollectorCluster
	}
	return ""
}

func (m *LightstepConfig) GetAccessTokenFile() string {
	if m != nil {
		return m.AccessTokenFile
	}
	return ""
}

type ZipkinConfig struct {
	CollectorCluster  string `protobuf:"bytes,1,opt,name=collector_cluster,json=collectorCluster" json:"collector_cluster,omitempty"`
	CollectorEndpoint string `protobuf:"bytes,2,opt,name=collector_endpoint,json=collectorEndpoint" json:"collector_endpoint,omitempty"`
}

func (m *ZipkinConfig) Reset()                    { *m = ZipkinConfig{} }
func (m *ZipkinConfig) String() string            { return proto.CompactTextString(m) }
func (*ZipkinConfig) ProtoMessage()               {}
func (*ZipkinConfig) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *ZipkinConfig) GetCollectorCluster() string {
	if m != nil {
		return m.CollectorCluster
	}
	return ""
}

func (m *ZipkinConfig) GetCollectorEndpoint() string {
	if m != nil {
		return m.CollectorEndpoint
	}
	return ""
}

type Tracing struct {
	Http *Tracing_Http `protobuf:"bytes,1,opt,name=http" json:"http,omitempty"`
}

func (m *Tracing) Reset()                    { *m = Tracing{} }
func (m *Tracing) String() string            { return proto.CompactTextString(m) }
func (*Tracing) ProtoMessage()               {}
func (*Tracing) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *Tracing) GetHttp() *Tracing_Http {
	if m != nil {
		return m.Http
	}
	return nil
}

type Tracing_Http struct {
	Name   string                   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Config *google_protobuf1.Struct `protobuf:"bytes,2,opt,name=config" json:"config,omitempty"`
}

func (m *Tracing_Http) Reset()                    { *m = Tracing_Http{} }
func (m *Tracing_Http) String() string            { return proto.CompactTextString(m) }
func (*Tracing_Http) ProtoMessage()               {}
func (*Tracing_Http) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2, 0} }

func (m *Tracing_Http) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Tracing_Http) GetConfig() *google_protobuf1.Struct {
	if m != nil {
		return m.Config
	}
	return nil
}

type Admin struct {
	AccessLogPath string   `protobuf:"bytes,1,opt,name=access_log_path,json=accessLogPath" json:"access_log_path,omitempty"`
	ProfilePath   string   `protobuf:"bytes,2,opt,name=profile_path,json=profilePath" json:"profile_path,omitempty"`
	Address       *Address `protobuf:"bytes,3,opt,name=address" json:"address,omitempty"`
}

func (m *Admin) Reset()                    { *m = Admin{} }
func (m *Admin) String() string            { return proto.CompactTextString(m) }
func (*Admin) ProtoMessage()               {}
func (*Admin) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *Admin) GetAccessLogPath() string {
	if m != nil {
		return m.AccessLogPath
	}
	return ""
}

func (m *Admin) GetProfilePath() string {
	if m != nil {
		return m.ProfilePath
	}
	return ""
}

func (m *Admin) GetAddress() *Address {
	if m != nil {
		return m.Address
	}
	return nil
}

type ClusterManager struct {
	LocalClusterName   string                           `protobuf:"bytes,1,opt,name=local_cluster_name,json=localClusterName" json:"local_cluster_name,omitempty"`
	OutlierDetection   *ClusterManager_OutlierDetection `protobuf:"bytes,2,opt,name=outlier_detection,json=outlierDetection" json:"outlier_detection,omitempty"`
	UpstreamBindConfig *BindConfig                      `protobuf:"bytes,3,opt,name=upstream_bind_config,json=upstreamBindConfig" json:"upstream_bind_config,omitempty"`
	LoadStatsConfig    *ApiConfigSource                 `protobuf:"bytes,4,opt,name=load_stats_config,json=loadStatsConfig" json:"load_stats_config,omitempty"`
}

func (m *ClusterManager) Reset()                    { *m = ClusterManager{} }
func (m *ClusterManager) String() string            { return proto.CompactTextString(m) }
func (*ClusterManager) ProtoMessage()               {}
func (*ClusterManager) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

func (m *ClusterManager) GetLocalClusterName() string {
	if m != nil {
		return m.LocalClusterName
	}
	return ""
}

func (m *ClusterManager) GetOutlierDetection() *ClusterManager_OutlierDetection {
	if m != nil {
		return m.OutlierDetection
	}
	return nil
}

func (m *ClusterManager) GetUpstreamBindConfig() *BindConfig {
	if m != nil {
		return m.UpstreamBindConfig
	}
	return nil
}

func (m *ClusterManager) GetLoadStatsConfig() *ApiConfigSource {
	if m != nil {
		return m.LoadStatsConfig
	}
	return nil
}

type ClusterManager_OutlierDetection struct {
	EventLogPath string `protobuf:"bytes,1,opt,name=event_log_path,json=eventLogPath" json:"event_log_path,omitempty"`
}

func (m *ClusterManager_OutlierDetection) Reset()         { *m = ClusterManager_OutlierDetection{} }
func (m *ClusterManager_OutlierDetection) String() string { return proto.CompactTextString(m) }
func (*ClusterManager_OutlierDetection) ProtoMessage()    {}
func (*ClusterManager_OutlierDetection) Descriptor() ([]byte, []int) {
	return fileDescriptor3, []int{4, 0}
}

func (m *ClusterManager_OutlierDetection) GetEventLogPath() string {
	if m != nil {
		return m.EventLogPath
	}
	return ""
}

type StatsdSink struct {
	// Types that are valid to be assigned to StatsdSpecifier:
	//	*StatsdSink_Address
	//	*StatsdSink_TcpClusterName
	StatsdSpecifier isStatsdSink_StatsdSpecifier `protobuf_oneof:"statsd_specifier"`
}

func (m *StatsdSink) Reset()                    { *m = StatsdSink{} }
func (m *StatsdSink) String() string            { return proto.CompactTextString(m) }
func (*StatsdSink) ProtoMessage()               {}
func (*StatsdSink) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{5} }

type isStatsdSink_StatsdSpecifier interface {
	isStatsdSink_StatsdSpecifier()
}

type StatsdSink_Address struct {
	Address *Address `protobuf:"bytes,1,opt,name=address,oneof"`
}
type StatsdSink_TcpClusterName struct {
	TcpClusterName string `protobuf:"bytes,2,opt,name=tcp_cluster_name,json=tcpClusterName,oneof"`
}

func (*StatsdSink_Address) isStatsdSink_StatsdSpecifier()        {}
func (*StatsdSink_TcpClusterName) isStatsdSink_StatsdSpecifier() {}

func (m *StatsdSink) GetStatsdSpecifier() isStatsdSink_StatsdSpecifier {
	if m != nil {
		return m.StatsdSpecifier
	}
	return nil
}

func (m *StatsdSink) GetAddress() *Address {
	if x, ok := m.GetStatsdSpecifier().(*StatsdSink_Address); ok {
		return x.Address
	}
	return nil
}

func (m *StatsdSink) GetTcpClusterName() string {
	if x, ok := m.GetStatsdSpecifier().(*StatsdSink_TcpClusterName); ok {
		return x.TcpClusterName
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*StatsdSink) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _StatsdSink_OneofMarshaler, _StatsdSink_OneofUnmarshaler, _StatsdSink_OneofSizer, []interface{}{
		(*StatsdSink_Address)(nil),
		(*StatsdSink_TcpClusterName)(nil),
	}
}

func _StatsdSink_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*StatsdSink)
	// statsd_specifier
	switch x := m.StatsdSpecifier.(type) {
	case *StatsdSink_Address:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Address); err != nil {
			return err
		}
	case *StatsdSink_TcpClusterName:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.TcpClusterName)
	case nil:
	default:
		return fmt.Errorf("StatsdSink.StatsdSpecifier has unexpected type %T", x)
	}
	return nil
}

func _StatsdSink_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*StatsdSink)
	switch tag {
	case 1: // statsd_specifier.address
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Address)
		err := b.DecodeMessage(msg)
		m.StatsdSpecifier = &StatsdSink_Address{msg}
		return true, err
	case 2: // statsd_specifier.tcp_cluster_name
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.StatsdSpecifier = &StatsdSink_TcpClusterName{x}
		return true, err
	default:
		return false, nil
	}
}

func _StatsdSink_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*StatsdSink)
	// statsd_specifier
	switch x := m.StatsdSpecifier.(type) {
	case *StatsdSink_Address:
		s := proto.Size(x.Address)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *StatsdSink_TcpClusterName:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.TcpClusterName)))
		n += len(x.TcpClusterName)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type StatsSink struct {
	Name   string                   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Config *google_protobuf1.Struct `protobuf:"bytes,2,opt,name=config" json:"config,omitempty"`
}

func (m *StatsSink) Reset()                    { *m = StatsSink{} }
func (m *StatsSink) String() string            { return proto.CompactTextString(m) }
func (*StatsSink) ProtoMessage()               {}
func (*StatsSink) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{6} }

func (m *StatsSink) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *StatsSink) GetConfig() *google_protobuf1.Struct {
	if m != nil {
		return m.Config
	}
	return nil
}

type TagSpecifier struct {
	TagName string `protobuf:"bytes,1,opt,name=tag_name,json=tagName" json:"tag_name,omitempty"`
	Regex   string `protobuf:"bytes,2,opt,name=regex" json:"regex,omitempty"`
}

func (m *TagSpecifier) Reset()                    { *m = TagSpecifier{} }
func (m *TagSpecifier) String() string            { return proto.CompactTextString(m) }
func (*TagSpecifier) ProtoMessage()               {}
func (*TagSpecifier) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{7} }

func (m *TagSpecifier) GetTagName() string {
	if m != nil {
		return m.TagName
	}
	return ""
}

func (m *TagSpecifier) GetRegex() string {
	if m != nil {
		return m.Regex
	}
	return ""
}

type StatsConfig struct {
	StatsTags         []*TagSpecifier             `protobuf:"bytes,1,rep,name=stats_tags,json=statsTags" json:"stats_tags,omitempty"`
	UseAllDefaultTags *google_protobuf2.BoolValue `protobuf:"bytes,2,opt,name=use_all_default_tags,json=useAllDefaultTags" json:"use_all_default_tags,omitempty"`
}

func (m *StatsConfig) Reset()                    { *m = StatsConfig{} }
func (m *StatsConfig) String() string            { return proto.CompactTextString(m) }
func (*StatsConfig) ProtoMessage()               {}
func (*StatsConfig) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{8} }

func (m *StatsConfig) GetStatsTags() []*TagSpecifier {
	if m != nil {
		return m.StatsTags
	}
	return nil
}

func (m *StatsConfig) GetUseAllDefaultTags() *google_protobuf2.BoolValue {
	if m != nil {
		return m.UseAllDefaultTags
	}
	return nil
}

type Watchdog struct {
	MissTimeout      *google_protobuf.Duration `protobuf:"bytes,1,opt,name=miss_timeout,json=missTimeout" json:"miss_timeout,omitempty"`
	MegamissTimeout  *google_protobuf.Duration `protobuf:"bytes,2,opt,name=megamiss_timeout,json=megamissTimeout" json:"megamiss_timeout,omitempty"`
	KillTimeout      *google_protobuf.Duration `protobuf:"bytes,3,opt,name=kill_timeout,json=killTimeout" json:"kill_timeout,omitempty"`
	MultikillTimeout *google_protobuf.Duration `protobuf:"bytes,4,opt,name=multikill_timeout,json=multikillTimeout" json:"multikill_timeout,omitempty"`
}

func (m *Watchdog) Reset()                    { *m = Watchdog{} }
func (m *Watchdog) String() string            { return proto.CompactTextString(m) }
func (*Watchdog) ProtoMessage()               {}
func (*Watchdog) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{9} }

func (m *Watchdog) GetMissTimeout() *google_protobuf.Duration {
	if m != nil {
		return m.MissTimeout
	}
	return nil
}

func (m *Watchdog) GetMegamissTimeout() *google_protobuf.Duration {
	if m != nil {
		return m.MegamissTimeout
	}
	return nil
}

func (m *Watchdog) GetKillTimeout() *google_protobuf.Duration {
	if m != nil {
		return m.KillTimeout
	}
	return nil
}

func (m *Watchdog) GetMultikillTimeout() *google_protobuf.Duration {
	if m != nil {
		return m.MultikillTimeout
	}
	return nil
}

type Runtime struct {
	SymlinkRoot          string `protobuf:"bytes,1,opt,name=symlink_root,json=symlinkRoot" json:"symlink_root,omitempty"`
	Subdirectory         string `protobuf:"bytes,2,opt,name=subdirectory" json:"subdirectory,omitempty"`
	OverrideSubdirectory string `protobuf:"bytes,3,opt,name=override_subdirectory,json=overrideSubdirectory" json:"override_subdirectory,omitempty"`
}

func (m *Runtime) Reset()                    { *m = Runtime{} }
func (m *Runtime) String() string            { return proto.CompactTextString(m) }
func (*Runtime) ProtoMessage()               {}
func (*Runtime) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{10} }

func (m *Runtime) GetSymlinkRoot() string {
	if m != nil {
		return m.SymlinkRoot
	}
	return ""
}

func (m *Runtime) GetSubdirectory() string {
	if m != nil {
		return m.Subdirectory
	}
	return ""
}

func (m *Runtime) GetOverrideSubdirectory() string {
	if m != nil {
		return m.OverrideSubdirectory
	}
	return ""
}

type RateLimitServiceConfig struct {
	ClusterName string `protobuf:"bytes,1,opt,name=cluster_name,json=clusterName" json:"cluster_name,omitempty"`
}

func (m *RateLimitServiceConfig) Reset()                    { *m = RateLimitServiceConfig{} }
func (m *RateLimitServiceConfig) String() string            { return proto.CompactTextString(m) }
func (*RateLimitServiceConfig) ProtoMessage()               {}
func (*RateLimitServiceConfig) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{11} }

func (m *RateLimitServiceConfig) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

type Bootstrap struct {
	Node               *Node                       `protobuf:"bytes,1,opt,name=node" json:"node,omitempty"`
	StaticResources    *Bootstrap_StaticResources  `protobuf:"bytes,2,opt,name=static_resources,json=staticResources" json:"static_resources,omitempty"`
	DynamicResources   *Bootstrap_DynamicResources `protobuf:"bytes,3,opt,name=dynamic_resources,json=dynamicResources" json:"dynamic_resources,omitempty"`
	ClusterManager     *ClusterManager             `protobuf:"bytes,4,opt,name=cluster_manager,json=clusterManager" json:"cluster_manager,omitempty"`
	FlagsPath          string                      `protobuf:"bytes,5,opt,name=flags_path,json=flagsPath" json:"flags_path,omitempty"`
	StatsSinks         []*StatsSink                `protobuf:"bytes,6,rep,name=stats_sinks,json=statsSinks" json:"stats_sinks,omitempty"`
	StatsConfig        *StatsConfig                `protobuf:"bytes,13,opt,name=stats_config,json=statsConfig" json:"stats_config,omitempty"`
	StatsFlushInterval *google_protobuf.Duration   `protobuf:"bytes,7,opt,name=stats_flush_interval,json=statsFlushInterval" json:"stats_flush_interval,omitempty"`
	Watchdog           *Watchdog                   `protobuf:"bytes,8,opt,name=watchdog" json:"watchdog,omitempty"`
	Tracing            *Tracing                    `protobuf:"bytes,9,opt,name=tracing" json:"tracing,omitempty"`
	RateLimitService   *RateLimitServiceConfig     `protobuf:"bytes,10,opt,name=rate_limit_service,json=rateLimitService" json:"rate_limit_service,omitempty"`
	Runtime            *Runtime                    `protobuf:"bytes,11,opt,name=runtime" json:"runtime,omitempty"`
	Admin              *Admin                      `protobuf:"bytes,12,opt,name=admin" json:"admin,omitempty"`
}

func (m *Bootstrap) Reset()                    { *m = Bootstrap{} }
func (m *Bootstrap) String() string            { return proto.CompactTextString(m) }
func (*Bootstrap) ProtoMessage()               {}
func (*Bootstrap) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{12} }

func (m *Bootstrap) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *Bootstrap) GetStaticResources() *Bootstrap_StaticResources {
	if m != nil {
		return m.StaticResources
	}
	return nil
}

func (m *Bootstrap) GetDynamicResources() *Bootstrap_DynamicResources {
	if m != nil {
		return m.DynamicResources
	}
	return nil
}

func (m *Bootstrap) GetClusterManager() *ClusterManager {
	if m != nil {
		return m.ClusterManager
	}
	return nil
}

func (m *Bootstrap) GetFlagsPath() string {
	if m != nil {
		return m.FlagsPath
	}
	return ""
}

func (m *Bootstrap) GetStatsSinks() []*StatsSink {
	if m != nil {
		return m.StatsSinks
	}
	return nil
}

func (m *Bootstrap) GetStatsConfig() *StatsConfig {
	if m != nil {
		return m.StatsConfig
	}
	return nil
}

func (m *Bootstrap) GetStatsFlushInterval() *google_protobuf.Duration {
	if m != nil {
		return m.StatsFlushInterval
	}
	return nil
}

func (m *Bootstrap) GetWatchdog() *Watchdog {
	if m != nil {
		return m.Watchdog
	}
	return nil
}

func (m *Bootstrap) GetTracing() *Tracing {
	if m != nil {
		return m.Tracing
	}
	return nil
}

func (m *Bootstrap) GetRateLimitService() *RateLimitServiceConfig {
	if m != nil {
		return m.RateLimitService
	}
	return nil
}

func (m *Bootstrap) GetRuntime() *Runtime {
	if m != nil {
		return m.Runtime
	}
	return nil
}

func (m *Bootstrap) GetAdmin() *Admin {
	if m != nil {
		return m.Admin
	}
	return nil
}

type Bootstrap_StaticResources struct {
	Listeners []*Listener `protobuf:"bytes,1,rep,name=listeners" json:"listeners,omitempty"`
	Clusters  []*Cluster  `protobuf:"bytes,2,rep,name=clusters" json:"clusters,omitempty"`
	Secrets   []*Secret   `protobuf:"bytes,3,rep,name=secrets" json:"secrets,omitempty"`
}

func (m *Bootstrap_StaticResources) Reset()                    { *m = Bootstrap_StaticResources{} }
func (m *Bootstrap_StaticResources) String() string            { return proto.CompactTextString(m) }
func (*Bootstrap_StaticResources) ProtoMessage()               {}
func (*Bootstrap_StaticResources) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{12, 0} }

func (m *Bootstrap_StaticResources) GetListeners() []*Listener {
	if m != nil {
		return m.Listeners
	}
	return nil
}

func (m *Bootstrap_StaticResources) GetClusters() []*Cluster {
	if m != nil {
		return m.Clusters
	}
	return nil
}

func (m *Bootstrap_StaticResources) GetSecrets() []*Secret {
	if m != nil {
		return m.Secrets
	}
	return nil
}

type Bootstrap_DynamicResources struct {
	LdsConfig    *ConfigSource                            `protobuf:"bytes,1,opt,name=lds_config,json=ldsConfig" json:"lds_config,omitempty"`
	CdsConfig    *ConfigSource                            `protobuf:"bytes,2,opt,name=cds_config,json=cdsConfig" json:"cds_config,omitempty"`
	AdsConfig    *ApiConfigSource                         `protobuf:"bytes,3,opt,name=ads_config,json=adsConfig" json:"ads_config,omitempty"`
	DeprecatedV1 *Bootstrap_DynamicResources_DeprecatedV1 `protobuf:"bytes,4,opt,name=deprecated_v1,json=deprecatedV1" json:"deprecated_v1,omitempty"`
}

func (m *Bootstrap_DynamicResources) Reset()                    { *m = Bootstrap_DynamicResources{} }
func (m *Bootstrap_DynamicResources) String() string            { return proto.CompactTextString(m) }
func (*Bootstrap_DynamicResources) ProtoMessage()               {}
func (*Bootstrap_DynamicResources) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{12, 1} }

func (m *Bootstrap_DynamicResources) GetLdsConfig() *ConfigSource {
	if m != nil {
		return m.LdsConfig
	}
	return nil
}

func (m *Bootstrap_DynamicResources) GetCdsConfig() *ConfigSource {
	if m != nil {
		return m.CdsConfig
	}
	return nil
}

func (m *Bootstrap_DynamicResources) GetAdsConfig() *ApiConfigSource {
	if m != nil {
		return m.AdsConfig
	}
	return nil
}

func (m *Bootstrap_DynamicResources) GetDeprecatedV1() *Bootstrap_DynamicResources_DeprecatedV1 {
	if m != nil {
		return m.DeprecatedV1
	}
	return nil
}

type Bootstrap_DynamicResources_DeprecatedV1 struct {
	SdsConfig *ConfigSource `protobuf:"bytes,1,opt,name=sds_config,json=sdsConfig" json:"sds_config,omitempty"`
}

func (m *Bootstrap_DynamicResources_DeprecatedV1) Reset() {
	*m = Bootstrap_DynamicResources_DeprecatedV1{}
}
func (m *Bootstrap_DynamicResources_DeprecatedV1) String() string { return proto.CompactTextString(m) }
func (*Bootstrap_DynamicResources_DeprecatedV1) ProtoMessage()    {}
func (*Bootstrap_DynamicResources_DeprecatedV1) Descriptor() ([]byte, []int) {
	return fileDescriptor3, []int{12, 1, 0}
}

func (m *Bootstrap_DynamicResources_DeprecatedV1) GetSdsConfig() *ConfigSource {
	if m != nil {
		return m.SdsConfig
	}
	return nil
}

func init() {
	proto.RegisterType((*LightstepConfig)(nil), "envoy.api.v2.LightstepConfig")
	proto.RegisterType((*ZipkinConfig)(nil), "envoy.api.v2.ZipkinConfig")
	proto.RegisterType((*Tracing)(nil), "envoy.api.v2.Tracing")
	proto.RegisterType((*Tracing_Http)(nil), "envoy.api.v2.Tracing.Http")
	proto.RegisterType((*Admin)(nil), "envoy.api.v2.Admin")
	proto.RegisterType((*ClusterManager)(nil), "envoy.api.v2.ClusterManager")
	proto.RegisterType((*ClusterManager_OutlierDetection)(nil), "envoy.api.v2.ClusterManager.OutlierDetection")
	proto.RegisterType((*StatsdSink)(nil), "envoy.api.v2.StatsdSink")
	proto.RegisterType((*StatsSink)(nil), "envoy.api.v2.StatsSink")
	proto.RegisterType((*TagSpecifier)(nil), "envoy.api.v2.TagSpecifier")
	proto.RegisterType((*StatsConfig)(nil), "envoy.api.v2.StatsConfig")
	proto.RegisterType((*Watchdog)(nil), "envoy.api.v2.Watchdog")
	proto.RegisterType((*Runtime)(nil), "envoy.api.v2.Runtime")
	proto.RegisterType((*RateLimitServiceConfig)(nil), "envoy.api.v2.RateLimitServiceConfig")
	proto.RegisterType((*Bootstrap)(nil), "envoy.api.v2.Bootstrap")
	proto.RegisterType((*Bootstrap_StaticResources)(nil), "envoy.api.v2.Bootstrap.StaticResources")
	proto.RegisterType((*Bootstrap_DynamicResources)(nil), "envoy.api.v2.Bootstrap.DynamicResources")
	proto.RegisterType((*Bootstrap_DynamicResources_DeprecatedV1)(nil), "envoy.api.v2.Bootstrap.DynamicResources.DeprecatedV1")
}

func init() { proto.RegisterFile("api/bootstrap.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 1293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xdf, 0x6e, 0x13, 0xc7,
	0x17, 0xc6, 0x4e, 0x42, 0xe2, 0x63, 0x27, 0xb6, 0x87, 0x00, 0xc6, 0x82, 0x9f, 0x7e, 0x58, 0x88,
	0x52, 0x5a, 0x36, 0x4a, 0x68, 0x25, 0xaa, 0x22, 0x55, 0x84, 0x80, 0xa0, 0xa4, 0x14, 0x8d, 0x53,
	0x2a, 0x71, 0xb3, 0x9a, 0xec, 0x8e, 0x37, 0x43, 0xc6, 0x3b, 0xab, 0x99, 0x59, 0xd3, 0x5c, 0xf5,
	0xa2, 0x6a, 0x9f, 0xa0, 0x7d, 0x89, 0xbe, 0x42, 0x1f, 0xa9, 0xea, 0x3b, 0x54, 0xf3, 0x67, 0xd7,
	0xde, 0x25, 0x10, 0x5a, 0xf5, 0x6e, 0xe7, 0x9c, 0xef, 0xfb, 0x66, 0xe6, 0xfc, 0x9b, 0x85, 0x0b,
	0x24, 0x63, 0x5b, 0x87, 0x42, 0x68, 0xa5, 0x25, 0xc9, 0x82, 0x4c, 0x0a, 0x2d, 0x50, 0x87, 0xa6,
	0x33, 0x71, 0x12, 0x90, 0x8c, 0x05, 0xb3, 0x9d, 0x61, 0xdf, 0x40, 0x48, 0x1c, 0x4b, 0xaa, 0x94,
	0x03, 0x0c, 0x37, 0x2c, 0x8b, 0x28, 0xea, 0xd7, 0xeb, 0x66, 0x1d, 0xc5, 0x6a, 0x71, 0xc9, 0xab,
	0x4b, 0x55, 0x2e, 0xff, 0x97, 0x08, 0x91, 0x70, 0xba, 0x65, 0x57, 0x87, 0xf9, 0x64, 0x2b, 0xce,
	0x25, 0xd1, 0x4c, 0xa4, 0xde, 0x7f, 0xb5, 0xee, 0x57, 0x5a, 0xe6, 0x91, 0x7e, 0x17, 0xfb, 0x8d,
	0x24, 0x59, 0x46, 0xa5, 0x57, 0x1f, 0xbd, 0x86, 0xee, 0x3e, 0x4b, 0x8e, 0xb4, 0xd2, 0x34, 0x7b,
	0x28, 0xd2, 0x09, 0x4b, 0xd0, 0x27, 0xd0, 0x8f, 0x04, 0xe7, 0x34, 0xd2, 0x42, 0x86, 0x11, 0xcf,
	0x95, 0xa6, 0x72, 0xd0, 0xf8, 0x7f, 0xe3, 0x56, 0x0b, 0xf7, 0x4a, 0xc7, 0x43, 0x67, 0x47, 0xb7,
	0xa1, 0x4f, 0xa2, 0x88, 0x2a, 0x15, 0x6a, 0x71, 0x4c, 0xd3, 0x70, 0xc2, 0x38, 0x1d, 0x34, 0x2d,
	0xb8, 0xeb, 0x1c, 0x07, 0xc6, 0xfe, 0x98, 0x71, 0x3a, 0x7a, 0x0d, 0x9d, 0x57, 0x2c, 0x3b, 0x66,
	0xe9, 0xbf, 0xd9, 0xe8, 0x0e, 0xa0, 0x39, 0x98, 0xa6, 0x71, 0x26, 0x58, 0xaa, 0xfd, 0x4e, 0x73,
	0x99, 0x47, 0xde, 0x31, 0xfa, 0xa5, 0x01, 0xab, 0x07, 0x92, 0x44, 0x2c, 0x4d, 0x50, 0x00, 0xcb,
	0x47, 0x5a, 0x67, 0x56, 0xba, 0xbd, 0x33, 0x0c, 0x16, 0xd3, 0x15, 0x78, 0x50, 0xf0, 0x44, 0xeb,
	0x0c, 0x5b, 0xdc, 0xf0, 0x19, 0x2c, 0x9b, 0x15, 0x42, 0xb0, 0x9c, 0x92, 0x29, 0xf5, 0x47, 0xb2,
	0xdf, 0x68, 0x0b, 0xce, 0x47, 0xf6, 0xf4, 0x76, 0xeb, 0xf6, 0xce, 0xe5, 0xc0, 0x05, 0x38, 0x28,
	0x02, 0x1c, 0x8c, 0x6d, 0xf8, 0xb1, 0x87, 0x8d, 0x7e, 0x6a, 0xc0, 0xca, 0x83, 0x78, 0xca, 0x52,
	0x74, 0x13, 0x7c, 0x44, 0x42, 0x2e, 0x92, 0x30, 0x23, 0xfa, 0xc8, 0x2b, 0xaf, 0x3b, 0xf3, 0xbe,
	0x48, 0x5e, 0x10, 0x7d, 0x84, 0xae, 0x43, 0x27, 0x93, 0xc2, 0x04, 0xd2, 0x81, 0xdc, 0x1d, 0xdb,
	0xde, 0x66, 0x21, 0x5b, 0xb0, 0xea, 0x2b, 0x6c, 0xb0, 0x64, 0x8f, 0x71, 0xb1, 0x7a, 0xa9, 0x07,
	0xce, 0x89, 0x0b, 0xd4, 0xe8, 0xcf, 0x26, 0x6c, 0xf8, 0x48, 0x7e, 0x43, 0x52, 0x92, 0x50, 0x89,
	0x3e, 0x05, 0xc4, 0x45, 0x44, 0x78, 0x11, 0xf9, 0x70, 0xe1, 0xae, 0x3d, 0xeb, 0xf1, 0x84, 0xe7,
	0xe6, 0xde, 0xaf, 0xa0, 0x2f, 0x72, 0xcd, 0x19, 0x95, 0x61, 0x4c, 0x35, 0x8d, 0x4c, 0x01, 0xfa,
	0x10, 0xdc, 0xa9, 0xee, 0x5d, 0xdd, 0x26, 0xf8, 0xd6, 0xb1, 0xf6, 0x0a, 0x12, 0xee, 0x89, 0x9a,
	0x05, 0x7d, 0x0d, 0x9b, 0x79, 0xa6, 0xb4, 0xa4, 0x64, 0x1a, 0x1e, 0xb2, 0x34, 0x0e, 0x7d, 0x84,
	0xdd, 0xd5, 0x06, 0x55, 0xf9, 0x5d, 0x96, 0xc6, 0xae, 0x7e, 0x30, 0x2a, 0x58, 0x73, 0x1b, 0x7a,
	0x0a, 0x7d, 0x2e, 0x48, 0x1c, 0x2a, 0x4d, 0xb4, 0x2a, 0x84, 0x96, 0xad, 0xd0, 0xb5, 0x5a, 0x8c,
	0x32, 0xe6, 0x38, 0x63, 0x91, 0xcb, 0x88, 0xe2, 0xae, 0xe1, 0x8d, 0x0d, 0xcd, 0x99, 0x87, 0xf7,
	0xa0, 0x57, 0x3f, 0x3c, 0xba, 0x01, 0x1b, 0x74, 0x46, 0x53, 0x5d, 0x4f, 0x61, 0xc7, 0x5a, 0x7d,
	0x06, 0x47, 0x3f, 0x02, 0x58, 0xa1, 0x78, 0xcc, 0xd2, 0x63, 0xb4, 0x3d, 0x4f, 0x56, 0xe3, 0x3d,
	0xc9, 0x7a, 0x72, 0xae, 0x4c, 0x17, 0xba, 0x0d, 0x3d, 0x1d, 0x65, 0xd5, 0xcc, 0xd8, 0x32, 0x78,
	0x72, 0x0e, 0x6f, 0xe8, 0x28, 0x5b, 0xc8, 0xcc, 0x2e, 0x82, 0x9e, 0xbd, 0x6c, 0x1c, 0xaa, 0x8c,
	0x46, 0x6c, 0xc2, 0xa8, 0x1c, 0xbd, 0x80, 0x96, 0x3d, 0x80, 0xdd, 0xff, 0x3f, 0x29, 0xe3, 0xaf,
	0xa0, 0x73, 0x40, 0x92, 0x71, 0xb1, 0x03, 0xba, 0x02, 0x6b, 0x9a, 0x24, 0x8b, 0x35, 0xb3, 0xaa,
	0x49, 0x62, 0x4b, 0x65, 0x13, 0x56, 0x24, 0x4d, 0xe8, 0x0f, 0xbe, 0x70, 0xdd, 0x62, 0xf4, 0x5b,
	0x03, 0xda, 0x0b, 0xd1, 0x45, 0x5f, 0x00, 0xb8, 0x1c, 0x69, 0x92, 0x98, 0xc0, 0x2c, 0x9d, 0xd2,
	0x9a, 0x0b, 0x1b, 0xe2, 0x96, 0x45, 0x1f, 0x90, 0x44, 0xa1, 0x67, 0xb0, 0x99, 0x2b, 0x1a, 0x12,
	0xce, 0xc3, 0x98, 0x4e, 0x48, 0xce, 0xb5, 0x13, 0x69, 0xfa, 0xfe, 0xae, 0x5f, 0x65, 0x57, 0x08,
	0xfe, 0x92, 0xf0, 0x9c, 0xe2, 0x7e, 0xae, 0xe8, 0x03, 0xce, 0xf7, 0x1c, 0xcb, 0x88, 0x8d, 0x7e,
	0x6d, 0xc2, 0xda, 0xf7, 0x44, 0x47, 0x47, 0xb1, 0x48, 0xd0, 0x7d, 0xe8, 0x4c, 0x99, 0x99, 0x65,
	0x6c, 0x4a, 0x45, 0xae, 0x7d, 0xbe, 0xae, 0xbc, 0xa5, 0xb8, 0xe7, 0x47, 0x30, 0x6e, 0x1b, 0xf8,
	0x81, 0x43, 0xa3, 0x3d, 0xe8, 0x4d, 0x69, 0x42, 0x2a, 0x0a, 0xcd, 0xb3, 0x14, 0xba, 0x05, 0xa5,
	0x50, 0xb9, 0x0f, 0x9d, 0x63, 0xc6, 0x79, 0xa9, 0xb0, 0x74, 0xe6, 0x19, 0x0c, 0xbc, 0x60, 0x3f,
	0x86, 0xfe, 0x34, 0xe7, 0x9a, 0x55, 0x24, 0x96, 0xcf, 0x92, 0xe8, 0x95, 0x1c, 0xaf, 0x33, 0xfa,
	0xb9, 0x01, 0xab, 0x38, 0x4f, 0x8d, 0x82, 0x19, 0x48, 0xea, 0x64, 0xca, 0x59, 0x7a, 0x1c, 0x4a,
	0x21, 0xb4, 0xcf, 0x77, 0xdb, 0xdb, 0xb0, 0x10, 0x1a, 0x8d, 0xa0, 0xa3, 0xf2, 0xc3, 0x98, 0x49,
	0x3b, 0x85, 0x4f, 0x7c, 0xea, 0x2b, 0x36, 0x74, 0x17, 0x2e, 0x8a, 0x19, 0x95, 0x92, 0xc5, 0x34,
	0xac, 0x80, 0x97, 0x2c, 0x78, 0xb3, 0x70, 0x8e, 0x17, 0x7c, 0xa3, 0x2f, 0xe1, 0x12, 0x26, 0x9a,
	0xee, 0xb3, 0x29, 0xd3, 0x63, 0x2a, 0x67, 0x2c, 0xa2, 0xbe, 0x80, 0xae, 0x43, 0xe7, 0x94, 0xc9,
	0xd5, 0x8e, 0xe6, 0xad, 0x31, 0xfa, 0x03, 0xa0, 0xb5, 0x5b, 0x3c, 0xd6, 0xe8, 0x26, 0x2c, 0xa7,
	0x22, 0xa6, 0x3e, 0xa9, 0xa8, 0x5a, 0x6b, 0xcf, 0x45, 0x4c, 0xb1, 0xf5, 0x23, 0xec, 0x1a, 0x8a,
	0x45, 0xa1, 0xa4, 0xca, 0xce, 0x86, 0xa2, 0xb4, 0x3e, 0xaa, 0x8d, 0xa2, 0xf2, 0x3f, 0x60, 0x6c,
	0xf1, 0xb8, 0x80, 0xe3, 0xae, 0xaa, 0x1a, 0xd0, 0x77, 0xd0, 0x8f, 0x4f, 0x52, 0x32, 0xad, 0x88,
	0xba, 0xcc, 0xde, 0x7a, 0x97, 0xe8, 0x9e, 0x23, 0xcc, 0x55, 0x7b, 0x71, 0xcd, 0x82, 0x1e, 0x41,
	0xb7, 0x88, 0xc1, 0xd4, 0xcd, 0x5b, 0x9f, 0xeb, 0xab, 0xef, 0x9b, 0xc9, 0x78, 0x23, 0xaa, 0x3e,
	0x05, 0xd7, 0x00, 0x26, 0x9c, 0x24, 0xca, 0x4d, 0xb4, 0x15, 0x1b, 0xc8, 0x96, 0xb5, 0xd8, 0xd7,
	0xe6, 0x1e, 0xb4, 0x5d, 0xab, 0x2a, 0x96, 0x1e, 0xab, 0xc1, 0x79, 0xdb, 0xab, 0x97, 0xab, 0x3b,
	0x94, 0xe3, 0x06, 0xbb, 0xb6, 0x36, 0x9f, 0xca, 0xd4, 0x72, 0x65, 0x10, 0xaf, 0xfb, 0x42, 0x7c,
	0x9b, 0xea, 0x47, 0xba, 0xdb, 0xc8, 0x67, 0xf8, 0x19, 0x6c, 0x3a, 0xf6, 0x84, 0xe7, 0xea, 0x28,
	0x64, 0xa9, 0xa6, 0x72, 0x46, 0xf8, 0x60, 0xf5, 0xac, 0x72, 0x46, 0x96, 0xf6, 0xd8, 0xb0, 0x9e,
	0x7a, 0x12, 0xda, 0x81, 0xb5, 0x37, 0xbe, 0xcd, 0x07, 0x6b, 0x56, 0xe0, 0x52, 0xf5, 0x18, 0xc5,
	0x10, 0xc0, 0x25, 0xce, 0x3c, 0xb3, 0xda, 0xfd, 0x1e, 0x0c, 0x5a, 0xa7, 0x4d, 0x6e, 0xff, 0xef,
	0x80, 0x0b, 0x14, 0xc2, 0x80, 0x24, 0xd1, 0x34, 0xe4, 0xa6, 0x5c, 0x43, 0xe5, 0xea, 0x75, 0x00,
	0x96, 0x7b, 0xa3, 0xca, 0x3d, 0xbd, 0xaa, 0x71, 0x4f, 0xd6, 0xec, 0xe6, 0x10, 0xd2, 0x35, 0xe2,
	0xa0, 0x7d, 0xda, 0x21, 0x7c, 0x97, 0xe2, 0x02, 0x85, 0x3e, 0x86, 0x15, 0x62, 0x7e, 0x38, 0x06,
	0x1d, 0x0b, 0xbf, 0x50, 0x7f, 0x6d, 0xa6, 0x2c, 0xc5, 0x0e, 0x31, 0xfc, 0xbd, 0x01, 0xdd, 0x5a,
	0xed, 0xa2, 0xcf, 0xa0, 0xc5, 0x99, 0xd2, 0x34, 0xa5, 0xb2, 0x98, 0xcb, 0xb5, 0x48, 0xed, 0x7b,
	0x37, 0x9e, 0x03, 0xd1, 0x36, 0xac, 0xf9, 0xa2, 0x32, 0xcd, 0xb2, 0xf4, 0xf6, 0x31, 0x7d, 0x09,
	0xe2, 0x12, 0x86, 0x02, 0x58, 0x55, 0x34, 0x92, 0x54, 0x9b, 0x4e, 0x30, 0x8c, 0xcd, 0x5a, 0x5d,
	0x58, 0x27, 0x2e, 0x40, 0xc3, 0xbf, 0x9a, 0xd0, 0xab, 0xf7, 0x84, 0x79, 0x46, 0x78, 0x5c, 0xd6,
	0xd7, 0xa9, 0x7f, 0x78, 0x95, 0x57, 0xbe, 0xc5, 0xe3, 0x85, 0x17, 0x28, 0x9a, 0x53, 0x9b, 0x67,
	0x53, 0xa3, 0x92, 0x7a, 0x1f, 0x80, 0xcc, 0xa9, 0x4b, 0x1f, 0xf2, 0x7b, 0xd1, 0x22, 0x25, 0xfb,
	0x15, 0xac, 0xc7, 0x34, 0x93, 0x34, 0x22, 0x9a, 0xc6, 0xe1, 0x6c, 0xdb, 0xf7, 0xec, 0xe7, 0x1f,
	0x3a, 0x08, 0x82, 0xbd, 0x92, 0xfd, 0x72, 0x1b, 0x77, 0xe2, 0x85, 0xd5, 0xf0, 0x29, 0x74, 0x16,
	0xbd, 0xf6, 0x99, 0xfd, 0x47, 0xf1, 0x51, 0xc5, 0x31, 0x0f, 0xcf, 0xdb, 0xc6, 0xba, 0xfb, 0x77,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x0c, 0xbc, 0x98, 0xaa, 0xf4, 0x0c, 0x00, 0x00,
}