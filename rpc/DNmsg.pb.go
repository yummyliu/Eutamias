// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc/DNmsg.proto

package rpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Ninfo struct {
	Ip       string `protobuf:"bytes,1,opt,name=ip" json:"ip,omitempty"`
	Port     int64  `protobuf:"varint,2,opt,name=port" json:"port,omitempty"`
	MaxConn  int64  `protobuf:"varint,3,opt,name=MaxConn" json:"MaxConn,omitempty"`
	CurConn  int64  `protobuf:"varint,4,opt,name=CurConn" json:"CurConn,omitempty"`
	Hostname string `protobuf:"bytes,5,opt,name=Hostname" json:"Hostname,omitempty"`
}

func (m *Ninfo) Reset()                    { *m = Ninfo{} }
func (m *Ninfo) String() string            { return proto.CompactTextString(m) }
func (*Ninfo) ProtoMessage()               {}
func (*Ninfo) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Ninfo) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *Ninfo) GetPort() int64 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Ninfo) GetMaxConn() int64 {
	if m != nil {
		return m.MaxConn
	}
	return 0
}

func (m *Ninfo) GetCurConn() int64 {
	if m != nil {
		return m.CurConn
	}
	return 0
}

func (m *Ninfo) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func init() {
	proto.RegisterType((*Ninfo)(nil), "rpc.Ninfo")
}

func init() { proto.RegisterFile("rpc/DNmsg.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 139 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x2a, 0x48, 0xd6,
	0x77, 0xf1, 0xcb, 0x2d, 0x4e, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2e, 0x2a, 0x48,
	0x56, 0xaa, 0xe6, 0x62, 0xf5, 0xcb, 0xcc, 0x4b, 0xcb, 0x17, 0xe2, 0xe3, 0x62, 0xca, 0x2c, 0x90,
	0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62, 0xca, 0x2c, 0x10, 0x12, 0xe2, 0x62, 0x29, 0xc8, 0x2f,
	0x2a, 0x91, 0x60, 0x52, 0x60, 0xd4, 0x60, 0x0e, 0x02, 0xb3, 0x85, 0x24, 0xb8, 0xd8, 0x7d, 0x13,
	0x2b, 0x9c, 0xf3, 0xf3, 0xf2, 0x24, 0x98, 0xc1, 0xc2, 0x30, 0x2e, 0x48, 0xc6, 0xb9, 0xb4, 0x08,
	0x2c, 0xc3, 0x02, 0x91, 0x81, 0x72, 0x85, 0xa4, 0xb8, 0x38, 0x3c, 0xf2, 0x8b, 0x4b, 0xf2, 0x12,
	0x73, 0x53, 0x25, 0x58, 0xc1, 0xa6, 0xc3, 0xf9, 0x49, 0x6c, 0x60, 0x87, 0x18, 0x03, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x42, 0xbd, 0x0b, 0x33, 0x9b, 0x00, 0x00, 0x00,
}
