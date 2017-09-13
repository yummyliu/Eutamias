// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Clientmsg.proto

package rpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ServerTimeReq struct {
}

func (m *ServerTimeReq) Reset()                    { *m = ServerTimeReq{} }
func (m *ServerTimeReq) String() string            { return proto.CompactTextString(m) }
func (*ServerTimeReq) ProtoMessage()               {}
func (*ServerTimeReq) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type ServerTimeRsp struct {
	ServerTime uint64 `protobuf:"varint,1,opt,name=ServerTime" json:"ServerTime,omitempty"`
}

func (m *ServerTimeRsp) Reset()                    { *m = ServerTimeRsp{} }
func (m *ServerTimeRsp) String() string            { return proto.CompactTextString(m) }
func (*ServerTimeRsp) ProtoMessage()               {}
func (*ServerTimeRsp) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *ServerTimeRsp) GetServerTime() uint64 {
	if m != nil {
		return m.ServerTime
	}
	return 0
}

type LoginReq struct {
	Id           uint64       `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Passwd       string       `protobuf:"bytes,2,opt,name=passwd" json:"passwd,omitempty"`
	OnlineStatue OnlineStatus `protobuf:"varint,3,opt,name=online_statue,json=onlineStatue,enum=rpc.OnlineStatus" json:"online_statue,omitempty"`
}

func (m *LoginReq) Reset()                    { *m = LoginReq{} }
func (m *LoginReq) String() string            { return proto.CompactTextString(m) }
func (*LoginReq) ProtoMessage()               {}
func (*LoginReq) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *LoginReq) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *LoginReq) GetPasswd() string {
	if m != nil {
		return m.Passwd
	}
	return ""
}

func (m *LoginReq) GetOnlineStatue() OnlineStatus {
	if m != nil {
		return m.OnlineStatue
	}
	return OnlineStatus_OS_ONLINE
}

type LoginRsp struct {
	Id         uint64     `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	ResultCode ResultCode `protobuf:"varint,2,opt,name=result_code,json=resultCode,enum=rpc.ResultCode" json:"result_code,omitempty"`
	ResultStr  string     `protobuf:"bytes,3,opt,name=result_str,json=resultStr" json:"result_str,omitempty"`
}

func (m *LoginRsp) Reset()                    { *m = LoginRsp{} }
func (m *LoginRsp) String() string            { return proto.CompactTextString(m) }
func (*LoginRsp) ProtoMessage()               {}
func (*LoginRsp) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *LoginRsp) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *LoginRsp) GetResultCode() ResultCode {
	if m != nil {
		return m.ResultCode
	}
	return ResultCode_RC_DEFAULT
}

func (m *LoginRsp) GetResultStr() string {
	if m != nil {
		return m.ResultStr
	}
	return ""
}

type LogoutReq struct {
	Id uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *LogoutReq) Reset()                    { *m = LogoutReq{} }
func (m *LogoutReq) String() string            { return proto.CompactTextString(m) }
func (*LogoutReq) ProtoMessage()               {}
func (*LogoutReq) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *LogoutReq) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type LogoutRsp struct {
	Id         uint64     `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	ResultCode ResultCode `protobuf:"varint,2,opt,name=result_code,json=resultCode,enum=rpc.ResultCode" json:"result_code,omitempty"`
	ResultStr  string     `protobuf:"bytes,3,opt,name=result_str,json=resultStr" json:"result_str,omitempty"`
}

func (m *LogoutRsp) Reset()                    { *m = LogoutRsp{} }
func (m *LogoutRsp) String() string            { return proto.CompactTextString(m) }
func (*LogoutRsp) ProtoMessage()               {}
func (*LogoutRsp) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *LogoutRsp) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *LogoutRsp) GetResultCode() ResultCode {
	if m != nil {
		return m.ResultCode
	}
	return ResultCode_RC_DEFAULT
}

func (m *LogoutRsp) GetResultStr() string {
	if m != nil {
		return m.ResultStr
	}
	return ""
}

type SendMsgReq struct {
	Id     uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Peerid uint64 `protobuf:"varint,2,opt,name=peerid" json:"peerid,omitempty"`
}

func (m *SendMsgReq) Reset()                    { *m = SendMsgReq{} }
func (m *SendMsgReq) String() string            { return proto.CompactTextString(m) }
func (*SendMsgReq) ProtoMessage()               {}
func (*SendMsgReq) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *SendMsgReq) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SendMsgReq) GetPeerid() uint64 {
	if m != nil {
		return m.Peerid
	}
	return 0
}

type SendMsgRsp struct {
	Id     uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Peerid uint64 `protobuf:"varint,2,opt,name=peerid" json:"peerid,omitempty"`
}

func (m *SendMsgRsp) Reset()                    { *m = SendMsgRsp{} }
func (m *SendMsgRsp) String() string            { return proto.CompactTextString(m) }
func (*SendMsgRsp) ProtoMessage()               {}
func (*SendMsgRsp) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *SendMsgRsp) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SendMsgRsp) GetPeerid() uint64 {
	if m != nil {
		return m.Peerid
	}
	return 0
}

type MsgData struct {
	Id        uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	SessionId uint64 `protobuf:"varint,2,opt,name=session_id,json=sessionId" json:"session_id,omitempty"`
	Content   string `protobuf:"bytes,3,opt,name=content" json:"content,omitempty"`
}

func (m *MsgData) Reset()                    { *m = MsgData{} }
func (m *MsgData) String() string            { return proto.CompactTextString(m) }
func (*MsgData) ProtoMessage()               {}
func (*MsgData) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

func (m *MsgData) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *MsgData) GetSessionId() uint64 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

func (m *MsgData) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type MsgDataReadAck struct {
	ResultCode ResultCode `protobuf:"varint,1,opt,name=result_code,json=resultCode,enum=rpc.ResultCode" json:"result_code,omitempty"`
	ResultStr  string     `protobuf:"bytes,2,opt,name=result_str,json=resultStr" json:"result_str,omitempty"`
	Sip        string     `protobuf:"bytes,3,opt,name=Sip" json:"Sip,omitempty"`
	Sport      uint64     `protobuf:"varint,4,opt,name=Sport" json:"Sport,omitempty"`
}

func (m *MsgDataReadAck) Reset()                    { *m = MsgDataReadAck{} }
func (m *MsgDataReadAck) String() string            { return proto.CompactTextString(m) }
func (*MsgDataReadAck) ProtoMessage()               {}
func (*MsgDataReadAck) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{9} }

func (m *MsgDataReadAck) GetResultCode() ResultCode {
	if m != nil {
		return m.ResultCode
	}
	return ResultCode_RC_DEFAULT
}

func (m *MsgDataReadAck) GetResultStr() string {
	if m != nil {
		return m.ResultStr
	}
	return ""
}

func (m *MsgDataReadAck) GetSip() string {
	if m != nil {
		return m.Sip
	}
	return ""
}

func (m *MsgDataReadAck) GetSport() uint64 {
	if m != nil {
		return m.Sport
	}
	return 0
}

type UnreadMsgReq struct {
	Id uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *UnreadMsgReq) Reset()                    { *m = UnreadMsgReq{} }
func (m *UnreadMsgReq) String() string            { return proto.CompactTextString(m) }
func (*UnreadMsgReq) ProtoMessage()               {}
func (*UnreadMsgReq) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{10} }

func (m *UnreadMsgReq) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type UnreadMsgInfo struct {
	Fromid    uint64 `protobuf:"varint,1,opt,name=fromid" json:"fromid,omitempty"`
	SessionId uint64 `protobuf:"varint,2,opt,name=session_id,json=sessionId" json:"session_id,omitempty"`
	Content   string `protobuf:"bytes,3,opt,name=content" json:"content,omitempty"`
}

func (m *UnreadMsgInfo) Reset()                    { *m = UnreadMsgInfo{} }
func (m *UnreadMsgInfo) String() string            { return proto.CompactTextString(m) }
func (*UnreadMsgInfo) ProtoMessage()               {}
func (*UnreadMsgInfo) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{11} }

func (m *UnreadMsgInfo) GetFromid() uint64 {
	if m != nil {
		return m.Fromid
	}
	return 0
}

func (m *UnreadMsgInfo) GetSessionId() uint64 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

func (m *UnreadMsgInfo) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type UnreadMsgRsp struct {
	ResultCode    ResultCode       `protobuf:"varint,1,opt,name=result_code,json=resultCode,enum=rpc.ResultCode" json:"result_code,omitempty"`
	ResultStr     string           `protobuf:"bytes,2,opt,name=result_str,json=resultStr" json:"result_str,omitempty"`
	UnreadMsgList []*UnreadMsgInfo `protobuf:"bytes,3,rep,name=UnreadMsg_list,json=UnreadMsgList" json:"UnreadMsg_list,omitempty"`
	MsgCount      uint64           `protobuf:"varint,4,opt,name=msg_count,json=msgCount" json:"msg_count,omitempty"`
}

func (m *UnreadMsgRsp) Reset()                    { *m = UnreadMsgRsp{} }
func (m *UnreadMsgRsp) String() string            { return proto.CompactTextString(m) }
func (*UnreadMsgRsp) ProtoMessage()               {}
func (*UnreadMsgRsp) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{12} }

func (m *UnreadMsgRsp) GetResultCode() ResultCode {
	if m != nil {
		return m.ResultCode
	}
	return ResultCode_RC_DEFAULT
}

func (m *UnreadMsgRsp) GetResultStr() string {
	if m != nil {
		return m.ResultStr
	}
	return ""
}

func (m *UnreadMsgRsp) GetUnreadMsgList() []*UnreadMsgInfo {
	if m != nil {
		return m.UnreadMsgList
	}
	return nil
}

func (m *UnreadMsgRsp) GetMsgCount() uint64 {
	if m != nil {
		return m.MsgCount
	}
	return 0
}

type CreateSessionReq struct {
	Fromid uint64 `protobuf:"varint,1,opt,name=fromid" json:"fromid,omitempty"`
	Peerid uint64 `protobuf:"varint,2,opt,name=peerid" json:"peerid,omitempty"`
}

func (m *CreateSessionReq) Reset()                    { *m = CreateSessionReq{} }
func (m *CreateSessionReq) String() string            { return proto.CompactTextString(m) }
func (*CreateSessionReq) ProtoMessage()               {}
func (*CreateSessionReq) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{13} }

func (m *CreateSessionReq) GetFromid() uint64 {
	if m != nil {
		return m.Fromid
	}
	return 0
}

func (m *CreateSessionReq) GetPeerid() uint64 {
	if m != nil {
		return m.Peerid
	}
	return 0
}

type CreateSessionRsp struct {
	Fromid     uint64     `protobuf:"varint,1,opt,name=fromid" json:"fromid,omitempty"`
	Peerid     uint64     `protobuf:"varint,2,opt,name=peerid" json:"peerid,omitempty"`
	SessionId  uint64     `protobuf:"varint,3,opt,name=session_id,json=sessionId" json:"session_id,omitempty"`
	ResultCode ResultCode `protobuf:"varint,4,opt,name=result_code,json=resultCode,enum=rpc.ResultCode" json:"result_code,omitempty"`
	ResultStr  string     `protobuf:"bytes,5,opt,name=result_str,json=resultStr" json:"result_str,omitempty"`
	Sip        string     `protobuf:"bytes,6,opt,name=Sip" json:"Sip,omitempty"`
	Sport      uint64     `protobuf:"varint,7,opt,name=Sport" json:"Sport,omitempty"`
}

func (m *CreateSessionRsp) Reset()                    { *m = CreateSessionRsp{} }
func (m *CreateSessionRsp) String() string            { return proto.CompactTextString(m) }
func (*CreateSessionRsp) ProtoMessage()               {}
func (*CreateSessionRsp) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{14} }

func (m *CreateSessionRsp) GetFromid() uint64 {
	if m != nil {
		return m.Fromid
	}
	return 0
}

func (m *CreateSessionRsp) GetPeerid() uint64 {
	if m != nil {
		return m.Peerid
	}
	return 0
}

func (m *CreateSessionRsp) GetSessionId() uint64 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

func (m *CreateSessionRsp) GetResultCode() ResultCode {
	if m != nil {
		return m.ResultCode
	}
	return ResultCode_RC_DEFAULT
}

func (m *CreateSessionRsp) GetResultStr() string {
	if m != nil {
		return m.ResultStr
	}
	return ""
}

func (m *CreateSessionRsp) GetSip() string {
	if m != nil {
		return m.Sip
	}
	return ""
}

func (m *CreateSessionRsp) GetSport() uint64 {
	if m != nil {
		return m.Sport
	}
	return 0
}

type CloseSessionReq struct {
	Fromid    uint64 `protobuf:"varint,1,opt,name=fromid" json:"fromid,omitempty"`
	Peerid    uint64 `protobuf:"varint,2,opt,name=peerid" json:"peerid,omitempty"`
	SessionId uint64 `protobuf:"varint,3,opt,name=session_id,json=sessionId" json:"session_id,omitempty"`
}

func (m *CloseSessionReq) Reset()                    { *m = CloseSessionReq{} }
func (m *CloseSessionReq) String() string            { return proto.CompactTextString(m) }
func (*CloseSessionReq) ProtoMessage()               {}
func (*CloseSessionReq) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{15} }

func (m *CloseSessionReq) GetFromid() uint64 {
	if m != nil {
		return m.Fromid
	}
	return 0
}

func (m *CloseSessionReq) GetPeerid() uint64 {
	if m != nil {
		return m.Peerid
	}
	return 0
}

func (m *CloseSessionReq) GetSessionId() uint64 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

type CloseSessionRsq struct {
	ResultCode ResultCode `protobuf:"varint,1,opt,name=result_code,json=resultCode,enum=rpc.ResultCode" json:"result_code,omitempty"`
	ResultStr  string     `protobuf:"bytes,2,opt,name=result_str,json=resultStr" json:"result_str,omitempty"`
}

func (m *CloseSessionRsq) Reset()                    { *m = CloseSessionRsq{} }
func (m *CloseSessionRsq) String() string            { return proto.CompactTextString(m) }
func (*CloseSessionRsq) ProtoMessage()               {}
func (*CloseSessionRsq) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{16} }

func (m *CloseSessionRsq) GetResultCode() ResultCode {
	if m != nil {
		return m.ResultCode
	}
	return ResultCode_RC_DEFAULT
}

func (m *CloseSessionRsq) GetResultStr() string {
	if m != nil {
		return m.ResultStr
	}
	return ""
}

func init() {
	proto.RegisterType((*ServerTimeReq)(nil), "rpc.ServerTimeReq")
	proto.RegisterType((*ServerTimeRsp)(nil), "rpc.ServerTimeRsp")
	proto.RegisterType((*LoginReq)(nil), "rpc.LoginReq")
	proto.RegisterType((*LoginRsp)(nil), "rpc.LoginRsp")
	proto.RegisterType((*LogoutReq)(nil), "rpc.LogoutReq")
	proto.RegisterType((*LogoutRsp)(nil), "rpc.LogoutRsp")
	proto.RegisterType((*SendMsgReq)(nil), "rpc.SendMsgReq")
	proto.RegisterType((*SendMsgRsp)(nil), "rpc.SendMsgRsp")
	proto.RegisterType((*MsgData)(nil), "rpc.MsgData")
	proto.RegisterType((*MsgDataReadAck)(nil), "rpc.MsgDataReadAck")
	proto.RegisterType((*UnreadMsgReq)(nil), "rpc.UnreadMsgReq")
	proto.RegisterType((*UnreadMsgInfo)(nil), "rpc.UnreadMsgInfo")
	proto.RegisterType((*UnreadMsgRsp)(nil), "rpc.UnreadMsgRsp")
	proto.RegisterType((*CreateSessionReq)(nil), "rpc.CreateSessionReq")
	proto.RegisterType((*CreateSessionRsp)(nil), "rpc.CreateSessionRsp")
	proto.RegisterType((*CloseSessionReq)(nil), "rpc.CloseSessionReq")
	proto.RegisterType((*CloseSessionRsq)(nil), "rpc.CloseSessionRsq")
}

func init() { proto.RegisterFile("Clientmsg.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 510 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x95, 0xcd, 0x6e, 0xda, 0x40,
	0x10, 0xc7, 0x65, 0x20, 0x10, 0x4f, 0xc2, 0x47, 0x57, 0x55, 0x85, 0x1a, 0x25, 0x42, 0x7b, 0xe2,
	0x44, 0x2b, 0x5a, 0x55, 0xea, 0xb1, 0x21, 0x97, 0x48, 0x44, 0x95, 0xd6, 0xed, 0xd9, 0x71, 0xec,
	0xc1, 0xda, 0x06, 0x76, 0xcd, 0xce, 0xd2, 0x3e, 0x43, 0x9f, 0xa8, 0xef, 0xd2, 0xa7, 0xa9, 0x6c,
	0x8c, 0x21, 0x0e, 0x56, 0x45, 0x1b, 0xe5, 0xe6, 0xff, 0x7f, 0x3c, 0x5f, 0x3f, 0x46, 0x18, 0xba,
	0x93, 0xb9, 0x44, 0x65, 0x17, 0x14, 0x8f, 0x12, 0xa3, 0xad, 0x66, 0x75, 0x93, 0x84, 0xaf, 0x7b,
	0x97, 0x01, 0xe1, 0x15, 0xce, 0xa4, 0xc2, 0xb5, 0xcd, 0xbb, 0xd0, 0xf6, 0xd0, 0x7c, 0x47, 0xf3,
	0x45, 0x2e, 0x50, 0xe0, 0x92, 0xbf, 0x79, 0x60, 0x50, 0xc2, 0x2e, 0x00, 0xb6, 0x46, 0xdf, 0x19,
	0x38, 0xc3, 0x86, 0xd8, 0x71, 0xf8, 0x37, 0x38, 0x9e, 0xea, 0x58, 0x2a, 0x81, 0x4b, 0xd6, 0x81,
	0x9a, 0x8c, 0xf2, 0x77, 0x6a, 0x32, 0x62, 0xaf, 0xa0, 0x99, 0x04, 0x44, 0x3f, 0xa2, 0x7e, 0x6d,
	0xe0, 0x0c, 0x5d, 0x91, 0x2b, 0xf6, 0x01, 0xda, 0x5a, 0xcd, 0xa5, 0x42, 0x9f, 0x6c, 0x60, 0x57,
	0xd8, 0xaf, 0x0f, 0x9c, 0x61, 0x67, 0xfc, 0x62, 0x64, 0x92, 0x70, 0xf4, 0x39, 0x8b, 0x78, 0x69,
	0x80, 0xc4, 0xa9, 0xde, 0x2a, 0xe4, 0xf7, 0x9b, 0x5e, 0x94, 0x3c, 0xea, 0xf5, 0x16, 0x4e, 0x0c,
	0xd2, 0x6a, 0x6e, 0xfd, 0x50, 0x47, 0x98, 0x35, 0xec, 0x8c, 0xbb, 0x59, 0x45, 0x91, 0xf9, 0x13,
	0x1d, 0xa1, 0x00, 0x53, 0x3c, 0xb3, 0x73, 0xc8, 0x95, 0x4f, 0xd6, 0x64, 0x23, 0xb8, 0xc2, 0x5d,
	0x3b, 0x9e, 0x35, 0xfc, 0x0c, 0xdc, 0xa9, 0x8e, 0xf5, 0xca, 0xee, 0xd9, 0x8c, 0xcf, 0x8b, 0xe0,
	0x73, 0x8c, 0xf2, 0x3e, 0xfd, 0x0d, 0x54, 0x74, 0x43, 0x71, 0x15, 0x65, 0x44, 0x23, 0xd7, 0x94,
	0x1b, 0x22, 0x57, 0xbb, 0x59, 0x7b, 0x86, 0xac, 0xca, 0x12, 0xd0, 0xba, 0xa1, 0xf8, 0x2a, 0xb0,
	0xc1, 0xa3, 0x94, 0x73, 0x00, 0x42, 0x22, 0xa9, 0x95, 0x5f, 0xa4, 0xb9, 0xb9, 0x73, 0x1d, 0xb1,
	0x3e, 0xb4, 0x42, 0xad, 0x2c, 0x2a, 0x9b, 0x6f, 0xb0, 0x91, 0xfc, 0xa7, 0x03, 0x9d, 0xbc, 0xa8,
	0xc0, 0x20, 0xfa, 0x14, 0xde, 0x97, 0x19, 0x39, 0x87, 0x32, 0xaa, 0x95, 0x18, 0xb1, 0x1e, 0xd4,
	0x3d, 0x99, 0xe4, 0x9d, 0xd3, 0x47, 0xf6, 0x12, 0x8e, 0xbc, 0x44, 0x1b, 0xdb, 0x6f, 0x64, 0x93,
	0xae, 0x05, 0xbf, 0x80, 0xd3, 0xaf, 0xca, 0x60, 0x50, 0x41, 0x93, 0xdf, 0x42, 0xbb, 0x88, 0x5f,
	0xab, 0x99, 0x4e, 0x41, 0xcd, 0x8c, 0x5e, 0x14, 0x2f, 0xe5, 0xea, 0xdf, 0x69, 0xfc, 0x72, 0x76,
	0x47, 0xa0, 0xe4, 0xe9, 0x59, 0x7c, 0x84, 0x4e, 0xd1, 0xc0, 0x9f, 0x4b, 0x4a, 0x47, 0xa8, 0x0f,
	0x4f, 0xc6, 0x2c, 0xab, 0xf9, 0x60, 0x3d, 0xb1, 0xdd, 0x76, 0x2a, 0xc9, 0xb2, 0x33, 0x70, 0x17,
	0x14, 0xfb, 0xa1, 0x5e, 0xa9, 0x0d, 0xb8, 0xe3, 0x05, 0xc5, 0x93, 0x54, 0xf3, 0x4b, 0xe8, 0x4d,
	0x0c, 0x06, 0x16, 0xbd, 0xf5, 0x9a, 0x29, 0xbf, 0x2a, 0x3c, 0x55, 0xf7, 0xf5, 0xdb, 0x29, 0x17,
	0xa1, 0xe4, 0xd0, 0x22, 0x25, 0xf6, 0xf5, 0x32, 0xfb, 0x12, 0xd0, 0xc6, 0xa1, 0x40, 0x8f, 0x2a,
	0x8e, 0xab, 0xb9, 0xe7, 0xb8, 0x5a, 0xbb, 0xc7, 0x75, 0x9b, 0xfe, 0xf1, 0x6a, 0xfa, 0x0f, 0x3e,
	0x7f, 0x59, 0x8d, 0xdf, 0x95, 0x3a, 0xd0, 0xf2, 0xc9, 0xcf, 0xe7, 0xae, 0x99, 0x7d, 0x1b, 0xde,
	0xfd, 0x09, 0x00, 0x00, 0xff, 0xff, 0x08, 0x21, 0x1f, 0x80, 0x45, 0x06, 0x00, 0x00,
}
