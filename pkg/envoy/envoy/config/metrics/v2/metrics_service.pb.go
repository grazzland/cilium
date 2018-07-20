// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/metrics/v2/metrics_service.proto

/*
Package v2 is a generated protocol buffer package.

It is generated from these files:
	envoy/config/metrics/v2/metrics_service.proto
	envoy/config/metrics/v2/stats.proto

It has these top-level messages:
	MetricsServiceConfig
	StatsSink
	StatsConfig
	TagSpecifier
	StatsdSink
	DogStatsdSink
	HystrixSink
*/
package v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import envoy_api_v2_core1 "github.com/cilium/cilium/pkg/envoy/envoy/api/v2/core"
import _ "github.com/lyft/protoc-gen-validate/validate"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Metrics Service is configured as a built-in *envoy.metrics_service* :ref:`StatsSink
// <envoy_api_msg_config.metrics.v2.StatsSink>`. This opaque configuration will be used to create
// Metrics Service.
type MetricsServiceConfig struct {
	// The upstream gRPC cluster that hosts the metrics service.
	GrpcService *envoy_api_v2_core1.GrpcService `protobuf:"bytes,1,opt,name=grpc_service,json=grpcService" json:"grpc_service,omitempty"`
}

func (m *MetricsServiceConfig) Reset()                    { *m = MetricsServiceConfig{} }
func (m *MetricsServiceConfig) String() string            { return proto.CompactTextString(m) }
func (*MetricsServiceConfig) ProtoMessage()               {}
func (*MetricsServiceConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *MetricsServiceConfig) GetGrpcService() *envoy_api_v2_core1.GrpcService {
	if m != nil {
		return m.GrpcService
	}
	return nil
}

func init() {
	proto.RegisterType((*MetricsServiceConfig)(nil), "envoy.config.metrics.v2.MetricsServiceConfig")
}

func init() { proto.RegisterFile("envoy/config/metrics/v2/metrics_service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 182 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4d, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0xcf, 0x4d, 0x2d, 0x29, 0xca, 0x4c, 0x2e,
	0xd6, 0x2f, 0x33, 0x82, 0x31, 0xe3, 0x8b, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0xf5, 0x0a, 0x8a,
	0xf2, 0x4b, 0xf2, 0x85, 0xc4, 0xc1, 0xca, 0xf5, 0x20, 0xca, 0xf5, 0xa0, 0x6a, 0xf4, 0xca, 0x8c,
	0xa4, 0x54, 0x20, 0xe6, 0x24, 0x16, 0x64, 0x82, 0x34, 0x27, 0xe7, 0x17, 0xa5, 0xea, 0xa7, 0x17,
	0x15, 0x24, 0xa3, 0x6a, 0x97, 0x12, 0x2f, 0x4b, 0xcc, 0xc9, 0x4c, 0x49, 0x2c, 0x49, 0xd5, 0x87,
	0x31, 0x20, 0x12, 0x4a, 0xa9, 0x5c, 0x22, 0xbe, 0x10, 0xc3, 0x82, 0x21, 0x1a, 0x9c, 0xc1, 0x36,
	0x08, 0xf9, 0x72, 0xf1, 0x20, 0x1b, 0x23, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x6d, 0x24, 0xa7, 0x07,
	0x71, 0x46, 0x62, 0x41, 0xa6, 0x5e, 0x99, 0x91, 0x1e, 0xc8, 0x36, 0x3d, 0xf7, 0xa2, 0x82, 0x64,
	0xa8, 0x5e, 0x27, 0xae, 0x5d, 0x2f, 0x0f, 0x30, 0xb3, 0x76, 0x31, 0x32, 0x09, 0x30, 0x06, 0x71,
	0xa7, 0x23, 0x24, 0x92, 0xd8, 0xc0, 0xb6, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xbd, 0xe4,
	0xe1, 0x00, 0xf6, 0x00, 0x00, 0x00,
}
