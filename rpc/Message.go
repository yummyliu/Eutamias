package rpc

import (
)

type Message struct {
	Cmd MsgCmd
	Seq uint64
	Msg  []byte
}
