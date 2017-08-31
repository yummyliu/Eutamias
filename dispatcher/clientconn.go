package main

import (
	"encoding/gob"
	"net"
	pb "github.com/yummyliu/Eutamias/rpc"
	"github.com/golang/protobuf/proto"
	"fmt"
)

func sendOut(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func handleNserver(con net.Conn) {
	log.Infof("get one Ns, addr=%s",con.RemoteAddr().String())
	defer func(con net.Conn) {
		err := con.Close()
		if err != nil {
			log.Error(err)
			return
		}
	}(con)

	for {
		var msg pb.Message
		conn_dec := gob.NewDecoder(con)
		err := conn_dec.Decode(&msg)
		if err != nil {
			log.Fatal(err)
			return
		}

		switch t := msg.Cmd; t {
		case pb.MsgCmd_C_NINFOREQ:
			handleNserverInfoReq(msg.Msg)
		default:
			log.Error("wrong cmd id")
		}
	}
}

func handleNserverInfoReq(msg []byte) {
	ninfoq := &pb.NinfoReq{}
	if err := proto.Unmarshal(msg, ninfoq); err != nil {
		log.Fatalf("failed to parse NinfoReq: ", err)
		return
	}

	log.Infof("req ninfo from client=%d", ninfoq.Id)
}
