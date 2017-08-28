package main

import (
	"github.com/op/go-logging"
	"fmt"
	pb "github.com/yummyliu/Eutamias/rpc"
	"time"
)

type ImClient struct {
	id uint64
}

func (c *ImClient) GetTime() time {
	timereq := &pb.ServerTimeReq{}

	return nil
}
