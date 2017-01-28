package web

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

type FileServer struct {
	Entry []*FileServer_Entry `protobuf:"bytes,1,rep,name=entry" json:"entry,omitempty"`
}

func (m *FileServer) Reset()                    { *m = FileServer{} }
func (m *FileServer) String() string            { return proto.CompactTextString(m) }
func (*FileServer) ProtoMessage()               {}
func (*FileServer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *FileServer) GetEntry() []*FileServer_Entry {
	if m != nil {
		return m.Entry
	}
	return nil
}

type FileServer_Entry struct {
	// Types that are valid to be assigned to FileOrDir:
	//	*FileServer_Entry_File
	//	*FileServer_Entry_Directory
	FileOrDir isFileServer_Entry_FileOrDir `protobuf_oneof:"FileOrDir"`
	Path      string                       `protobuf:"bytes,3,opt,name=path" json:"path,omitempty"`
}

func (m *FileServer_Entry) Reset()                    { *m = FileServer_Entry{} }
func (m *FileServer_Entry) String() string            { return proto.CompactTextString(m) }
func (*FileServer_Entry) ProtoMessage()               {}
func (*FileServer_Entry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type isFileServer_Entry_FileOrDir interface {
	isFileServer_Entry_FileOrDir()
}

type FileServer_Entry_File struct {
	File string `protobuf:"bytes,1,opt,name=File,oneof"`
}
type FileServer_Entry_Directory struct {
	Directory string `protobuf:"bytes,2,opt,name=Directory,oneof"`
}

func (*FileServer_Entry_File) isFileServer_Entry_FileOrDir()      {}
func (*FileServer_Entry_Directory) isFileServer_Entry_FileOrDir() {}

func (m *FileServer_Entry) GetFileOrDir() isFileServer_Entry_FileOrDir {
	if m != nil {
		return m.FileOrDir
	}
	return nil
}

func (m *FileServer_Entry) GetFile() string {
	if x, ok := m.GetFileOrDir().(*FileServer_Entry_File); ok {
		return x.File
	}
	return ""
}

func (m *FileServer_Entry) GetDirectory() string {
	if x, ok := m.GetFileOrDir().(*FileServer_Entry_Directory); ok {
		return x.Directory
	}
	return ""
}

func (m *FileServer_Entry) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*FileServer_Entry) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _FileServer_Entry_OneofMarshaler, _FileServer_Entry_OneofUnmarshaler, _FileServer_Entry_OneofSizer, []interface{}{
		(*FileServer_Entry_File)(nil),
		(*FileServer_Entry_Directory)(nil),
	}
}

func _FileServer_Entry_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*FileServer_Entry)
	// FileOrDir
	switch x := m.FileOrDir.(type) {
	case *FileServer_Entry_File:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.File)
	case *FileServer_Entry_Directory:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Directory)
	case nil:
	default:
		return fmt.Errorf("FileServer_Entry.FileOrDir has unexpected type %T", x)
	}
	return nil
}

func _FileServer_Entry_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*FileServer_Entry)
	switch tag {
	case 1: // FileOrDir.File
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.FileOrDir = &FileServer_Entry_File{x}
		return true, err
	case 2: // FileOrDir.Directory
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.FileOrDir = &FileServer_Entry_Directory{x}
		return true, err
	default:
		return false, nil
	}
}

func _FileServer_Entry_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*FileServer_Entry)
	// FileOrDir
	switch x := m.FileOrDir.(type) {
	case *FileServer_Entry_File:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.File)))
		n += len(x.File)
	case *FileServer_Entry_Directory:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Directory)))
		n += len(x.Directory)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Server struct {
	Domain   []string                               `protobuf:"bytes,1,rep,name=domain" json:"domain,omitempty"`
	Settings *v2ray_core_common_serial.TypedMessage `protobuf:"bytes,2,opt,name=settings" json:"settings,omitempty"`
}

func (m *Server) Reset()                    { *m = Server{} }
func (m *Server) String() string            { return proto.CompactTextString(m) }
func (*Server) ProtoMessage()               {}
func (*Server) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Server) GetDomain() []string {
	if m != nil {
		return m.Domain
	}
	return nil
}

func (m *Server) GetSettings() *v2ray_core_common_serial.TypedMessage {
	if m != nil {
		return m.Settings
	}
	return nil
}

type Config struct {
	Server []*Server `protobuf:"bytes,1,rep,name=server" json:"server,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Config) GetServer() []*Server {
	if m != nil {
		return m.Server
	}
	return nil
}

func init() {
	proto.RegisterType((*FileServer)(nil), "v2ray.core.app.web.FileServer")
	proto.RegisterType((*FileServer_Entry)(nil), "v2ray.core.app.web.FileServer.Entry")
	proto.RegisterType((*Server)(nil), "v2ray.core.app.web.Server")
	proto.RegisterType((*Config)(nil), "v2ray.core.app.web.Config")
}

func init() { proto.RegisterFile("v2ray.com/core/app/web/config.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x91, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0x87, 0x4d, 0xff, 0x04, 0x33, 0xb9, 0x2d, 0x52, 0x42, 0x0f, 0x52, 0xaa, 0x48, 0x4f, 0x1b,
	0x89, 0x37, 0x11, 0xc4, 0xb4, 0x8a, 0x17, 0x51, 0xa2, 0x28, 0x78, 0x50, 0x36, 0xe9, 0x58, 0x17,
	0x9a, 0xec, 0xb2, 0x59, 0x5a, 0xf2, 0x46, 0xe2, 0x53, 0xca, 0xee, 0x46, 0x2b, 0xda, 0x5b, 0x26,
	0xf3, 0xfd, 0x66, 0x26, 0x5f, 0xe0, 0x60, 0x95, 0x28, 0xd6, 0xd0, 0x42, 0x94, 0x71, 0x21, 0x14,
	0xc6, 0x4c, 0xca, 0x78, 0x8d, 0x79, 0x5c, 0x88, 0xea, 0x8d, 0x2f, 0xa8, 0x54, 0x42, 0x0b, 0x42,
	0xbe, 0x21, 0x85, 0x94, 0x49, 0x49, 0xd7, 0x98, 0x0f, 0x8f, 0xff, 0x04, 0x0b, 0x51, 0x96, 0xa2,
	0x8a, 0x6b, 0x54, 0x9c, 0x2d, 0x63, 0xdd, 0x48, 0x9c, 0xbf, 0x96, 0x58, 0xd7, 0x6c, 0x81, 0x6e,
	0xca, 0xf8, 0xc3, 0x03, 0xb8, 0xe2, 0x4b, 0xbc, 0x47, 0xb5, 0x42, 0x45, 0x4e, 0xa1, 0x8f, 0x95,
	0x56, 0x4d, 0xe4, 0x8d, 0xba, 0x93, 0x30, 0x39, 0xa4, 0xff, 0x97, 0xd0, 0x0d, 0x4e, 0x2f, 0x0d,
	0x9b, 0xb9, 0xc8, 0xf0, 0x05, 0xfa, 0xb6, 0x26, 0x7b, 0xd0, 0x33, 0x4c, 0xe4, 0x8d, 0xbc, 0x49,
	0x70, 0xbd, 0x93, 0xd9, 0x8a, 0xec, 0x43, 0x30, 0xe3, 0x0a, 0x0b, 0x2d, 0x54, 0x13, 0x75, 0xda,
	0xd6, 0xe6, 0x15, 0x21, 0xd0, 0x93, 0x4c, 0xbf, 0x47, 0x5d, 0xd3, 0xca, 0xec, 0x73, 0x1a, 0x42,
	0x60, 0xb2, 0xb7, 0x6a, 0xc6, 0xd5, 0x78, 0x0e, 0x7e, 0x7b, 0xe5, 0x00, 0xfc, 0xb9, 0x28, 0x19,
	0xaf, 0xec, 0x99, 0x41, 0xd6, 0x56, 0x24, 0x85, 0xdd, 0x1a, 0xb5, 0xe6, 0xd5, 0xa2, 0xb6, 0x1b,
	0xc2, 0xe4, 0xe8, 0xf7, 0x07, 0x38, 0x1b, 0xd4, 0xd9, 0xa0, 0x0f, 0xc6, 0xc6, 0x8d, 0x93, 0x91,
	0xfd, 0xe4, 0xc6, 0x67, 0xe0, 0x4f, 0xad, 0x66, 0x92, 0x80, 0x5f, 0xdb, 0x7d, 0xad, 0x8c, 0xe1,
	0x36, 0x19, 0xee, 0xa2, 0xac, 0x25, 0xd3, 0x73, 0x18, 0x14, 0xa2, 0xdc, 0x02, 0xa6, 0xa1, 0x9b,
	0x7a, 0x67, 0xac, 0x3f, 0x77, 0xd7, 0x98, 0x7f, 0x76, 0xc8, 0x63, 0x92, 0xb1, 0x86, 0x4e, 0x0d,
	0x76, 0x21, 0x25, 0x7d, 0xc2, 0x3c, 0xf7, 0xed, 0x6f, 0x39, 0xf9, 0x0a, 0x00, 0x00, 0xff, 0xff,
	0xc9, 0x59, 0x48, 0x03, 0x03, 0x02, 0x00, 0x00,
}
