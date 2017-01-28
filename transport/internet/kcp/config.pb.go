package kcp

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import v2ray_core_common_serial "v2ray.com/core/common/serial"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Maximum Transmission Unit, in bytes.
type MTU struct {
	Value uint32 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *MTU) Reset()                    { *m = MTU{} }
func (m *MTU) String() string            { return proto.CompactTextString(m) }
func (*MTU) ProtoMessage()               {}
func (*MTU) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *MTU) GetValue() uint32 {
	if m != nil {
		return m.Value
	}
	return 0
}

// Transmission Time Interview, in milli-sec.
type TTI struct {
	Value uint32 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *TTI) Reset()                    { *m = TTI{} }
func (m *TTI) String() string            { return proto.CompactTextString(m) }
func (*TTI) ProtoMessage()               {}
func (*TTI) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TTI) GetValue() uint32 {
	if m != nil {
		return m.Value
	}
	return 0
}

// Uplink capacity, in MB.
type UplinkCapacity struct {
	Value uint32 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *UplinkCapacity) Reset()                    { *m = UplinkCapacity{} }
func (m *UplinkCapacity) String() string            { return proto.CompactTextString(m) }
func (*UplinkCapacity) ProtoMessage()               {}
func (*UplinkCapacity) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *UplinkCapacity) GetValue() uint32 {
	if m != nil {
		return m.Value
	}
	return 0
}

// Downlink capacity, in MB.
type DownlinkCapacity struct {
	Value uint32 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *DownlinkCapacity) Reset()                    { *m = DownlinkCapacity{} }
func (m *DownlinkCapacity) String() string            { return proto.CompactTextString(m) }
func (*DownlinkCapacity) ProtoMessage()               {}
func (*DownlinkCapacity) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *DownlinkCapacity) GetValue() uint32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type WriteBuffer struct {
	// Buffer size in bytes.
	Size uint32 `protobuf:"varint,1,opt,name=size" json:"size,omitempty"`
}

func (m *WriteBuffer) Reset()                    { *m = WriteBuffer{} }
func (m *WriteBuffer) String() string            { return proto.CompactTextString(m) }
func (*WriteBuffer) ProtoMessage()               {}
func (*WriteBuffer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *WriteBuffer) GetSize() uint32 {
	if m != nil {
		return m.Size
	}
	return 0
}

type ReadBuffer struct {
	// Buffer size in bytes.
	Size uint32 `protobuf:"varint,1,opt,name=size" json:"size,omitempty"`
}

func (m *ReadBuffer) Reset()                    { *m = ReadBuffer{} }
func (m *ReadBuffer) String() string            { return proto.CompactTextString(m) }
func (*ReadBuffer) ProtoMessage()               {}
func (*ReadBuffer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ReadBuffer) GetSize() uint32 {
	if m != nil {
		return m.Size
	}
	return 0
}

type ConnectionReuse struct {
	Enable bool `protobuf:"varint,1,opt,name=enable" json:"enable,omitempty"`
}

func (m *ConnectionReuse) Reset()                    { *m = ConnectionReuse{} }
func (m *ConnectionReuse) String() string            { return proto.CompactTextString(m) }
func (*ConnectionReuse) ProtoMessage()               {}
func (*ConnectionReuse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ConnectionReuse) GetEnable() bool {
	if m != nil {
		return m.Enable
	}
	return false
}

type Config struct {
	Mtu              *MTU                                   `protobuf:"bytes,1,opt,name=mtu" json:"mtu,omitempty"`
	Tti              *TTI                                   `protobuf:"bytes,2,opt,name=tti" json:"tti,omitempty"`
	UplinkCapacity   *UplinkCapacity                        `protobuf:"bytes,3,opt,name=uplink_capacity,json=uplinkCapacity" json:"uplink_capacity,omitempty"`
	DownlinkCapacity *DownlinkCapacity                      `protobuf:"bytes,4,opt,name=downlink_capacity,json=downlinkCapacity" json:"downlink_capacity,omitempty"`
	Congestion       bool                                   `protobuf:"varint,5,opt,name=congestion" json:"congestion,omitempty"`
	WriteBuffer      *WriteBuffer                           `protobuf:"bytes,6,opt,name=write_buffer,json=writeBuffer" json:"write_buffer,omitempty"`
	ReadBuffer       *ReadBuffer                            `protobuf:"bytes,7,opt,name=read_buffer,json=readBuffer" json:"read_buffer,omitempty"`
	HeaderConfig     *v2ray_core_common_serial.TypedMessage `protobuf:"bytes,8,opt,name=header_config,json=headerConfig" json:"header_config,omitempty"`
	ConnectionReuse  *ConnectionReuse                       `protobuf:"bytes,9,opt,name=connection_reuse,json=connectionReuse" json:"connection_reuse,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Config) GetMtu() *MTU {
	if m != nil {
		return m.Mtu
	}
	return nil
}

func (m *Config) GetTti() *TTI {
	if m != nil {
		return m.Tti
	}
	return nil
}

func (m *Config) GetUplinkCapacity() *UplinkCapacity {
	if m != nil {
		return m.UplinkCapacity
	}
	return nil
}

func (m *Config) GetDownlinkCapacity() *DownlinkCapacity {
	if m != nil {
		return m.DownlinkCapacity
	}
	return nil
}

func (m *Config) GetCongestion() bool {
	if m != nil {
		return m.Congestion
	}
	return false
}

func (m *Config) GetWriteBuffer() *WriteBuffer {
	if m != nil {
		return m.WriteBuffer
	}
	return nil
}

func (m *Config) GetReadBuffer() *ReadBuffer {
	if m != nil {
		return m.ReadBuffer
	}
	return nil
}

func (m *Config) GetHeaderConfig() *v2ray_core_common_serial.TypedMessage {
	if m != nil {
		return m.HeaderConfig
	}
	return nil
}

func (m *Config) GetConnectionReuse() *ConnectionReuse {
	if m != nil {
		return m.ConnectionReuse
	}
	return nil
}

func init() {
	proto.RegisterType((*MTU)(nil), "v2ray.core.transport.internet.kcp.MTU")
	proto.RegisterType((*TTI)(nil), "v2ray.core.transport.internet.kcp.TTI")
	proto.RegisterType((*UplinkCapacity)(nil), "v2ray.core.transport.internet.kcp.UplinkCapacity")
	proto.RegisterType((*DownlinkCapacity)(nil), "v2ray.core.transport.internet.kcp.DownlinkCapacity")
	proto.RegisterType((*WriteBuffer)(nil), "v2ray.core.transport.internet.kcp.WriteBuffer")
	proto.RegisterType((*ReadBuffer)(nil), "v2ray.core.transport.internet.kcp.ReadBuffer")
	proto.RegisterType((*ConnectionReuse)(nil), "v2ray.core.transport.internet.kcp.ConnectionReuse")
	proto.RegisterType((*Config)(nil), "v2ray.core.transport.internet.kcp.Config")
}

func init() { proto.RegisterFile("v2ray.com/core/transport/internet/kcp/config.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 491 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x54, 0xdf, 0x6f, 0xd3, 0x30,
	0x18, 0xd4, 0xd6, 0xad, 0x8c, 0x2f, 0xdb, 0x5a, 0x22, 0x84, 0x22, 0x90, 0xd0, 0x5a, 0x89, 0x69,
	0x3c, 0xe0, 0x40, 0xf6, 0xc2, 0x73, 0xcb, 0x4b, 0x35, 0x15, 0x81, 0x95, 0x82, 0x34, 0x69, 0x0a,
	0xae, 0xf3, 0xb5, 0x44, 0x6d, 0xec, 0xc8, 0x71, 0x56, 0x95, 0xff, 0x08, 0xfe, 0x4a, 0x64, 0xa7,
	0xe9, 0x2f, 0x34, 0x96, 0xb7, 0xd8, 0xbe, 0x3b, 0x5b, 0xf7, 0xdd, 0x05, 0x82, 0xfb, 0x40, 0xb1,
	0x25, 0xe1, 0x32, 0xf5, 0xb9, 0x54, 0xe8, 0x6b, 0xc5, 0x44, 0x9e, 0x49, 0xa5, 0xfd, 0x44, 0x68,
	0x54, 0x02, 0xb5, 0x3f, 0xe3, 0x99, 0xcf, 0xa5, 0x98, 0x24, 0x53, 0x92, 0x29, 0xa9, 0xa5, 0xdb,
	0xa9, 0x38, 0x0a, 0xc9, 0x1a, 0x4f, 0x2a, 0x3c, 0x99, 0xf1, 0xec, 0xe5, 0xfb, 0x3d, 0x59, 0x2e,
	0xd3, 0x54, 0x0a, 0x3f, 0x47, 0x95, 0xb0, 0xb9, 0xaf, 0x97, 0x19, 0xc6, 0x51, 0x8a, 0x79, 0xce,
	0xa6, 0x58, 0x8a, 0x76, 0x5f, 0x41, 0x63, 0x18, 0x8e, 0xdc, 0xe7, 0x70, 0x7c, 0xcf, 0xe6, 0x05,
	0x7a, 0x07, 0x17, 0x07, 0x57, 0x67, 0xb4, 0x5c, 0x98, 0xc3, 0x30, 0x1c, 0x3c, 0x70, 0x78, 0x09,
	0xe7, 0xa3, 0x6c, 0x9e, 0x88, 0x59, 0x9f, 0x65, 0x8c, 0x27, 0x7a, 0xf9, 0x00, 0xee, 0x0a, 0xda,
	0x9f, 0xe4, 0x42, 0xd4, 0x40, 0x76, 0xc0, 0xf9, 0xae, 0x12, 0x8d, 0xbd, 0x62, 0x32, 0x41, 0xe5,
	0xba, 0x70, 0x94, 0x27, 0xbf, 0x2a, 0x8c, 0xfd, 0xee, 0x5e, 0x00, 0x50, 0x64, 0xf1, 0x7f, 0x10,
	0x6f, 0xa1, 0xd5, 0x97, 0x42, 0x20, 0xd7, 0x89, 0x14, 0x14, 0x8b, 0x1c, 0xdd, 0x17, 0xd0, 0x44,
	0xc1, 0xc6, 0xf3, 0x12, 0x78, 0x42, 0x57, 0xab, 0xee, 0xef, 0x63, 0x68, 0xf6, 0xad, 0xc3, 0xee,
	0x47, 0x68, 0xa4, 0xba, 0xb0, 0xe7, 0x4e, 0x70, 0x49, 0x1e, 0x75, 0x9a, 0x0c, 0xc3, 0x11, 0x35,
	0x14, 0xc3, 0xd4, 0x3a, 0xf1, 0x0e, 0x6b, 0x33, 0xc3, 0x70, 0x40, 0x0d, 0xc5, 0xbd, 0x85, 0x56,
	0x61, 0x0d, 0x8c, 0xf8, 0xca, 0x17, 0xaf, 0x61, 0x55, 0x3e, 0xd4, 0x50, 0xd9, 0xb5, 0x9e, 0x9e,
	0x17, 0xbb, 0xa3, 0xf8, 0x01, 0xcf, 0xe2, 0x95, 0xe9, 0x1b, 0xf5, 0x23, 0xab, 0x7e, 0x5d, 0x43,
	0x7d, 0x7f, 0x60, 0xb4, 0x1d, 0xef, 0x8f, 0xf0, 0x35, 0x00, 0x97, 0x62, 0x8a, 0xb9, 0xf1, 0xd9,
	0x3b, 0xb6, 0xc6, 0x6e, 0xed, 0xb8, 0x5f, 0xe1, 0x74, 0x61, 0x86, 0x19, 0x8d, 0xed, 0xac, 0xbc,
	0xa6, 0xbd, 0x9c, 0xd4, 0xb8, 0x7c, 0x2b, 0x03, 0xd4, 0x59, 0x6c, 0x05, 0xe2, 0x33, 0x38, 0x0a,
	0x59, 0x5c, 0x29, 0x3e, 0xb1, 0x8a, 0xef, 0x6a, 0x28, 0x6e, 0x22, 0x43, 0x41, 0x6d, 0xe2, 0x73,
	0x03, 0x67, 0x3f, 0x91, 0xc5, 0xa8, 0xa2, 0xb2, 0x67, 0xde, 0xc9, 0xbf, 0x43, 0x2c, 0x1b, 0x44,
	0xca, 0x06, 0x91, 0xd0, 0x34, 0x68, 0x58, 0x16, 0x88, 0x9e, 0x96, 0xe4, 0x55, 0x82, 0xee, 0xa0,
	0xcd, 0xd7, 0xb9, 0x8b, 0x94, 0x09, 0x9e, 0xf7, 0xd4, 0xea, 0x05, 0x35, 0x5e, 0xb8, 0x17, 0x59,
	0xda, 0xe2, 0xbb, 0x1b, 0xbd, 0x3b, 0x78, 0xc3, 0x65, 0xfa, 0xb8, 0x52, 0xcf, 0x29, 0xdf, 0xf3,
	0xc5, 0xb4, 0xfb, 0xb6, 0x31, 0xe3, 0xd9, 0x9f, 0xc3, 0xce, 0xb7, 0x80, 0xb2, 0x25, 0xe9, 0x1b,
	0x56, 0xb8, 0x66, 0x0d, 0x2a, 0xd6, 0x0d, 0xcf, 0xc6, 0x4d, 0xfb, 0x37, 0xb8, 0xfe, 0x1b, 0x00,
	0x00, 0xff, 0xff, 0x77, 0xae, 0x78, 0x3a, 0x98, 0x04, 0x00, 0x00,
}
