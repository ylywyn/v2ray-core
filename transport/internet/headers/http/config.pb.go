package http

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

type Header struct {
	// "Accept", "Cookie", etc
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Each entry must be valid in one piece. Random entry will be chosen if multiple entries present.
	Value []string `protobuf:"bytes,2,rep,name=value" json:"value,omitempty"`
}

func (m *Header) Reset()                    { *m = Header{} }
func (m *Header) String() string            { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()               {}
func (*Header) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Header) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Header) GetValue() []string {
	if m != nil {
		return m.Value
	}
	return nil
}

// HTTP version. Default value "1.1".
type Version struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *Version) Reset()                    { *m = Version{} }
func (m *Version) String() string            { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()               {}
func (*Version) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Version) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// HTTP method. Default value "GET".
type Method struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *Method) Reset()                    { *m = Method{} }
func (m *Method) String() string            { return proto.CompactTextString(m) }
func (*Method) ProtoMessage()               {}
func (*Method) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Method) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type RequestConfig struct {
	// Full HTTP version like "1.1".
	Version *Version `protobuf:"bytes,1,opt,name=version" json:"version,omitempty"`
	// GET, POST, CONNECT etc
	Method *Method `protobuf:"bytes,2,opt,name=method" json:"method,omitempty"`
	// URI like "/login.php"
	Uri    []string  `protobuf:"bytes,3,rep,name=uri" json:"uri,omitempty"`
	Header []*Header `protobuf:"bytes,4,rep,name=header" json:"header,omitempty"`
}

func (m *RequestConfig) Reset()                    { *m = RequestConfig{} }
func (m *RequestConfig) String() string            { return proto.CompactTextString(m) }
func (*RequestConfig) ProtoMessage()               {}
func (*RequestConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RequestConfig) GetVersion() *Version {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *RequestConfig) GetMethod() *Method {
	if m != nil {
		return m.Method
	}
	return nil
}

func (m *RequestConfig) GetUri() []string {
	if m != nil {
		return m.Uri
	}
	return nil
}

func (m *RequestConfig) GetHeader() []*Header {
	if m != nil {
		return m.Header
	}
	return nil
}

type Status struct {
	// Status code. Default "200".
	Code string `protobuf:"bytes,1,opt,name=code" json:"code,omitempty"`
	// Statue reason. Default "OK".
	Reason string `protobuf:"bytes,2,opt,name=reason" json:"reason,omitempty"`
}

func (m *Status) Reset()                    { *m = Status{} }
func (m *Status) String() string            { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()               {}
func (*Status) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Status) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Status) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

type ResponseConfig struct {
	Version *Version  `protobuf:"bytes,1,opt,name=version" json:"version,omitempty"`
	Status  *Status   `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
	Header  []*Header `protobuf:"bytes,3,rep,name=header" json:"header,omitempty"`
}

func (m *ResponseConfig) Reset()                    { *m = ResponseConfig{} }
func (m *ResponseConfig) String() string            { return proto.CompactTextString(m) }
func (*ResponseConfig) ProtoMessage()               {}
func (*ResponseConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ResponseConfig) GetVersion() *Version {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *ResponseConfig) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ResponseConfig) GetHeader() []*Header {
	if m != nil {
		return m.Header
	}
	return nil
}

type Config struct {
	// Settings for authenticating requests. If not set, client side will not send authenication header, and server side will bypass authentication.
	Request *RequestConfig `protobuf:"bytes,1,opt,name=request" json:"request,omitempty"`
	// Settings for authenticating responses. If not set, client side will bypass authentication, and server side will not send authentication header.
	Response *ResponseConfig `protobuf:"bytes,2,opt,name=response" json:"response,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Config) GetRequest() *RequestConfig {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *Config) GetResponse() *ResponseConfig {
	if m != nil {
		return m.Response
	}
	return nil
}

func init() {
	proto.RegisterType((*Header)(nil), "v2ray.core.transport.internet.headers.http.Header")
	proto.RegisterType((*Version)(nil), "v2ray.core.transport.internet.headers.http.Version")
	proto.RegisterType((*Method)(nil), "v2ray.core.transport.internet.headers.http.Method")
	proto.RegisterType((*RequestConfig)(nil), "v2ray.core.transport.internet.headers.http.RequestConfig")
	proto.RegisterType((*Status)(nil), "v2ray.core.transport.internet.headers.http.Status")
	proto.RegisterType((*ResponseConfig)(nil), "v2ray.core.transport.internet.headers.http.ResponseConfig")
	proto.RegisterType((*Config)(nil), "v2ray.core.transport.internet.headers.http.Config")
}

func init() {
	proto.RegisterFile("v2ray.com/core/transport/internet/headers/http/config.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 399 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x94, 0x4f, 0x6b, 0xdb, 0x30,
	0x18, 0xc6, 0xb1, 0x9d, 0x39, 0xcb, 0x1b, 0x36, 0x86, 0x18, 0xc3, 0xa7, 0x2d, 0xf8, 0x14, 0x72,
	0x90, 0xc1, 0xd9, 0x65, 0xdb, 0x2d, 0xb9, 0x64, 0x85, 0x40, 0x51, 0x4a, 0x0e, 0xbd, 0xa9, 0x8e,
	0xda, 0x18, 0x1a, 0xc9, 0x95, 0xe4, 0x40, 0xbe, 0x41, 0x3f, 0x4b, 0xef, 0xfd, 0x6c, 0xbd, 0x16,
	0xfd, 0xb1, 0x9b, 0x1e, 0x0a, 0x75, 0x4b, 0x6f, 0xef, 0x6b, 0xbf, 0xcf, 0x8f, 0xf7, 0x79, 0x2c,
	0x19, 0xfe, 0xed, 0x73, 0x49, 0x0f, 0xb8, 0x10, 0xbb, 0xac, 0x10, 0x92, 0x65, 0x5a, 0x52, 0xae,
	0x2a, 0x21, 0x75, 0x56, 0x72, 0xcd, 0x24, 0x67, 0x3a, 0xdb, 0x32, 0xba, 0x61, 0x52, 0x65, 0x5b,
	0xad, 0xab, 0xac, 0x10, 0xfc, 0xb2, 0xbc, 0xc2, 0x95, 0x14, 0x5a, 0xa0, 0x49, 0x23, 0x96, 0x0c,
	0xb7, 0x42, 0xdc, 0x08, 0xb1, 0x17, 0x62, 0x23, 0x4c, 0x73, 0x88, 0x17, 0xb6, 0x47, 0x08, 0x7a,
	0x9c, 0xee, 0x58, 0x12, 0x8c, 0x82, 0xf1, 0x80, 0xd8, 0x1a, 0x7d, 0x87, 0x4f, 0x7b, 0x7a, 0x5d,
	0xb3, 0x24, 0x1c, 0x45, 0xe3, 0x01, 0x71, 0x4d, 0xfa, 0x0b, 0xfa, 0x6b, 0x26, 0x55, 0x29, 0xf8,
	0xd3, 0x80, 0x53, 0xf9, 0x81, 0x9f, 0x10, 0x2f, 0x99, 0xde, 0x8a, 0xcd, 0x0b, 0xef, 0x6f, 0x43,
	0xf8, 0x42, 0xd8, 0x4d, 0xcd, 0x94, 0x9e, 0xdb, 0xc5, 0xd1, 0x12, 0xfa, 0x7b, 0x87, 0xb4, 0x93,
	0xc3, 0x7c, 0x8a, 0x5f, 0x6f, 0x02, 0xfb, 0x6d, 0x48, 0xc3, 0x40, 0x27, 0x10, 0xef, 0xec, 0x02,
	0x49, 0x68, 0x69, 0x79, 0x17, 0x9a, 0x5b, 0x9d, 0x78, 0x02, 0xfa, 0x06, 0x51, 0x2d, 0xcb, 0x24,
	0xb2, 0x09, 0x98, 0xd2, 0xd0, 0x9d, 0x20, 0xe9, 0x8d, 0xa2, 0xae, 0x74, 0x97, 0x36, 0xf1, 0x84,
	0xf4, 0x37, 0xc4, 0x2b, 0x4d, 0x75, 0xad, 0x4c, 0xfe, 0x85, 0xd8, 0xb4, 0xf9, 0x9b, 0x1a, 0xfd,
	0x80, 0x58, 0x32, 0xaa, 0x04, 0xb7, 0x3e, 0x06, 0xc4, 0x77, 0xe9, 0x43, 0x00, 0x5f, 0x09, 0x53,
	0x95, 0xe0, 0x8a, 0x7d, 0x58, 0x82, 0xca, 0xee, 0xf5, 0x96, 0x04, 0x9d, 0x23, 0xe2, 0x09, 0x47,
	0x79, 0x45, 0xef, 0xce, 0xeb, 0x3e, 0x80, 0xd8, 0x3b, 0x5e, 0x41, 0x5f, 0xba, 0x43, 0xe4, 0x1d,
	0xff, 0xe9, 0xc2, 0x7d, 0x76, 0xfe, 0x48, 0x43, 0x42, 0x6b, 0xf8, 0x2c, 0x7d, 0xb0, 0xde, 0xf9,
	0xdf, 0x6e, 0xd4, 0xe3, 0x8f, 0x42, 0x5a, 0xd6, 0xac, 0x02, 0x73, 0x99, 0x3b, 0xa0, 0x66, 0x43,
	0xc7, 0x38, 0x35, 0x57, 0xfa, 0xbc, 0x67, 0x1e, 0xdd, 0x85, 0x93, 0x75, 0x4e, 0xe8, 0x01, 0xcf,
	0x8d, 0xfe, 0xac, 0xd5, 0xff, 0x6f, 0xf4, 0x0b, 0xaf, 0x5f, 0x68, 0x5d, 0x5d, 0xc4, 0xf6, 0x67,
	0x30, 0x7d, 0x0c, 0x00, 0x00, 0xff, 0xff, 0xd2, 0x30, 0x13, 0xcf, 0x4b, 0x04, 0x00, 0x00,
}
