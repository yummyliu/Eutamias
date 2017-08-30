package rpc

import (
)

type Message struct {
	Cmd uint64
	Seq uint64
	Version uint64
	//Msg interface{}
}
