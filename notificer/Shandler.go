package main

import (
	"encoding/gob"
	_ "github.com/golang/protobuf/proto"
	pb "github.com/yummyliu/Eutamias/rpc"
	"net"
)

func handleSserver(con net.Conn) {
	log.Infof("get one Ds, addr=%s",con.RemoteAddr().String())
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
		case pb.MsgCmd_C_HEARTBEAT:
			handleHeartBeat(msg, con)
//		case pb.MsgCmd_C_SINFOUPD:
//			handleNinfoUpt()
		default:
			log.Error("wrong cmd id")
		}
	}
}

