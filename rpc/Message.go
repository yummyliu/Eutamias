package rpc

import (
)

const (
	C_HEARTBEAT = 1
	C_MSGDATA = 2
	C_MSGACK = 3
)

type Message struct {
	Cmd uint64
	Seq uint64
	Version uint64
	Msg  []byte
}
