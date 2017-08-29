package rpc

import (
)

type Message struct {
	cmd uint64
	seq uint64
	version uint64
	msg interface{}
}
