// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v2/cluster/circuit_breaker.proto

/*
Package cluster is a generated protocol buffer package.

It is generated from these files:
	envoy/api/v2/cluster/circuit_breaker.proto
	envoy/api/v2/cluster/outlier_detection.proto

It has these top-level messages:
	CircuitBreakers
	OutlierDetection
*/
package cluster

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import envoy_api_v2_core "github.com/cilium/cilium/pkg/envoy/envoy/api/v2/core"
import google_protobuf1 "github.com/golang/protobuf/ptypes/wrappers"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// :ref:`Circuit breaking<arch_overview_circuit_break>` settings can be
// specified individually for each defined priority.
type CircuitBreakers struct {
	// If multiple :ref:`Thresholds<envoy_api_msg_cluster.CircuitBreakers.Thresholds>`
	// are defined with the same :ref:`RoutingPriority<envoy_api_enum_core.RoutingPriority>`,
	// the first one in the list is used. If no Thresholds is defined for a given
	// :ref:`RoutingPriority<envoy_api_enum_core.RoutingPriority>`, the default values
	// are used.
	Thresholds []*CircuitBreakers_Thresholds `protobuf:"bytes,1,rep,name=thresholds" json:"thresholds,omitempty"`
}

func (m *CircuitBreakers) Reset()                    { *m = CircuitBreakers{} }
func (m *CircuitBreakers) String() string            { return proto.CompactTextString(m) }
func (*CircuitBreakers) ProtoMessage()               {}
func (*CircuitBreakers) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CircuitBreakers) GetThresholds() []*CircuitBreakers_Thresholds {
	if m != nil {
		return m.Thresholds
	}
	return nil
}

// A Thresholds defines CircuitBreaker settings for a
// :ref:`RoutingPriority<envoy_api_enum_core.RoutingPriority>`.
type CircuitBreakers_Thresholds struct {
	// The :ref:`RoutingPriority<envoy_api_enum_core.RoutingPriority>`
	// the specified CircuitBreaker settings apply to.
	// [#comment:TODO(htuch): add (validate.rules).enum.defined_only = true once
	// https://github.com/lyft/protoc-gen-validate/issues/42 is resolved.]
	Priority envoy_api_v2_core.RoutingPriority `protobuf:"varint,1,opt,name=priority,enum=envoy.api.v2.core.RoutingPriority" json:"priority,omitempty"`
	// The maximum number of connections that Envoy will make to the upstream
	// cluster. If not specified, the default is 1024.
	MaxConnections *google_protobuf1.UInt32Value `protobuf:"bytes,2,opt,name=max_connections,json=maxConnections" json:"max_connections,omitempty"`
	// The maximum number of pending requests that Envoy will allow to the
	// upstream cluster. If not specified, the default is 1024.
	MaxPendingRequests *google_protobuf1.UInt32Value `protobuf:"bytes,3,opt,name=max_pending_requests,json=maxPendingRequests" json:"max_pending_requests,omitempty"`
	// The maximum number of parallel requests that Envoy will make to the
	// upstream cluster. If not specified, the default is 1024.
	MaxRequests *google_protobuf1.UInt32Value `protobuf:"bytes,4,opt,name=max_requests,json=maxRequests" json:"max_requests,omitempty"`
	// The maximum number of parallel retries that Envoy will allow to the
	// upstream cluster. If not specified, the default is 3.
	MaxRetries *google_protobuf1.UInt32Value `protobuf:"bytes,5,opt,name=max_retries,json=maxRetries" json:"max_retries,omitempty"`
}

func (m *CircuitBreakers_Thresholds) Reset()                    { *m = CircuitBreakers_Thresholds{} }
func (m *CircuitBreakers_Thresholds) String() string            { return proto.CompactTextString(m) }
func (*CircuitBreakers_Thresholds) ProtoMessage()               {}
func (*CircuitBreakers_Thresholds) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *CircuitBreakers_Thresholds) GetPriority() envoy_api_v2_core.RoutingPriority {
	if m != nil {
		return m.Priority
	}
	return envoy_api_v2_core.RoutingPriority_DEFAULT
}

func (m *CircuitBreakers_Thresholds) GetMaxConnections() *google_protobuf1.UInt32Value {
	if m != nil {
		return m.MaxConnections
	}
	return nil
}

func (m *CircuitBreakers_Thresholds) GetMaxPendingRequests() *google_protobuf1.UInt32Value {
	if m != nil {
		return m.MaxPendingRequests
	}
	return nil
}

func (m *CircuitBreakers_Thresholds) GetMaxRequests() *google_protobuf1.UInt32Value {
	if m != nil {
		return m.MaxRequests
	}
	return nil
}

func (m *CircuitBreakers_Thresholds) GetMaxRetries() *google_protobuf1.UInt32Value {
	if m != nil {
		return m.MaxRetries
	}
	return nil
}

func init() {
	proto.RegisterType((*CircuitBreakers)(nil), "envoy.api.v2.cluster.CircuitBreakers")
	proto.RegisterType((*CircuitBreakers_Thresholds)(nil), "envoy.api.v2.cluster.CircuitBreakers.Thresholds")
}

func init() { proto.RegisterFile("envoy/api/v2/cluster/circuit_breaker.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 369 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xcf, 0x6e, 0xe2, 0x30,
	0x10, 0x87, 0x15, 0xd8, 0x7f, 0x72, 0x56, 0x20, 0x45, 0x68, 0x15, 0x21, 0x84, 0x10, 0x87, 0x15,
	0xda, 0x83, 0xbd, 0x0a, 0xe7, 0xdd, 0xd5, 0x82, 0x38, 0xf4, 0x82, 0x50, 0xda, 0x72, 0xe8, 0x05,
	0x39, 0x61, 0x1a, 0xac, 0x26, 0xb6, 0x6b, 0x3b, 0x34, 0xbc, 0x51, 0xd5, 0x37, 0xe9, 0x6b, 0xb4,
	0x2f, 0x52, 0x11, 0xa7, 0xa1, 0x45, 0x3d, 0x70, 0x1b, 0xcf, 0xfc, 0xbe, 0x4f, 0xa3, 0x31, 0xfa,
	0x05, 0x7c, 0x2b, 0x76, 0x84, 0x4a, 0x46, 0xb6, 0x01, 0x89, 0xd3, 0x5c, 0x1b, 0x50, 0x24, 0x66,
	0x2a, 0xce, 0x99, 0x59, 0x45, 0x0a, 0xe8, 0x0d, 0x28, 0x2c, 0x95, 0x30, 0xc2, 0xeb, 0x94, 0x59,
	0x4c, 0x25, 0xc3, 0xdb, 0x00, 0x57, 0xd9, 0x6e, 0xef, 0xbd, 0x41, 0x28, 0x20, 0x11, 0xd5, 0x60,
	0x99, 0x6e, 0x3f, 0x11, 0x22, 0x49, 0x81, 0x94, 0xaf, 0x28, 0xbf, 0x26, 0x77, 0x8a, 0x4a, 0x09,
	0x4a, 0x57, 0xf3, 0x4e, 0x22, 0x12, 0x51, 0x96, 0x64, 0x5f, 0xd9, 0xee, 0xf0, 0xb1, 0x89, 0xda,
	0x53, 0xbb, 0xc3, 0xc4, 0xae, 0xa0, 0xbd, 0x05, 0x42, 0x66, 0xa3, 0x40, 0x6f, 0x44, 0xba, 0xd6,
	0xbe, 0x33, 0x68, 0x8e, 0xdc, 0xe0, 0x37, 0xfe, 0x68, 0x25, 0x7c, 0x84, 0xe2, 0x8b, 0x9a, 0x0b,
	0xdf, 0x38, 0xba, 0xcf, 0x0d, 0x84, 0x0e, 0x23, 0xef, 0x2f, 0xfa, 0x26, 0x15, 0x13, 0x8a, 0x99,
	0x9d, 0xef, 0x0c, 0x9c, 0x51, 0x2b, 0x18, 0x1e, 0xe9, 0x85, 0x02, 0x1c, 0x8a, 0xdc, 0x30, 0x9e,
	0x2c, 0xaa, 0x64, 0x58, 0x33, 0xde, 0x0c, 0xb5, 0x33, 0x5a, 0xac, 0x62, 0xc1, 0x39, 0xc4, 0x86,
	0x09, 0xae, 0xfd, 0xc6, 0xc0, 0x19, 0xb9, 0x41, 0x0f, 0xdb, 0x23, 0xe0, 0xd7, 0x23, 0xe0, 0xcb,
	0x33, 0x6e, 0xc6, 0xc1, 0x92, 0xa6, 0x39, 0x84, 0xad, 0x8c, 0x16, 0xd3, 0x03, 0xe3, 0xcd, 0x51,
	0x67, 0xaf, 0x91, 0xc0, 0xd7, 0x8c, 0x27, 0x2b, 0x05, 0xb7, 0x39, 0x68, 0xa3, 0xfd, 0xe6, 0x09,
	0x2e, 0x2f, 0xa3, 0xc5, 0xc2, 0x82, 0x61, 0xc5, 0x79, 0xff, 0xd0, 0xf7, 0xbd, 0xaf, 0xf6, 0x7c,
	0x3a, 0xc1, 0xe3, 0x66, 0xb4, 0xa8, 0x05, 0x7f, 0x90, 0x6b, 0x05, 0x46, 0x31, 0xd0, 0xfe, 0xe7,
	0x13, 0x78, 0x54, 0xf2, 0x65, 0x7e, 0xf2, 0xf3, 0xfe, 0xa9, 0xef, 0x5c, 0x7d, 0xad, 0xfe, 0xe6,
	0xa1, 0xf1, 0x63, 0x56, 0xde, 0xf4, 0xbf, 0x64, 0x78, 0x19, 0xe0, 0xa9, 0x6d, 0xcf, 0xcf, 0xa3,
	0x2f, 0xa5, 0x69, 0xfc, 0x12, 0x00, 0x00, 0xff, 0xff, 0x41, 0xcf, 0x7d, 0x2c, 0x92, 0x02, 0x00,
	0x00,
}
