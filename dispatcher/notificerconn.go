package main

import (
	"encoding/gob"
	"github.com/golang/protobuf/proto"
	pb "github.com/yummyliu/Eutamias/rpc"
	"net"
)

type Nserver struct {
	Ip       string
	Port     int64
	MaxConn  int64
	CurConn  int64
	Hostname string
}

func handleClient(con net.Conn) {
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
		case pb.MsgCmd_C_HEARTBEAT:
			handleHeartBeat(msg.Msg)
		case pb.MsgCmd_C_NINFOUPD:
			handleNserverInfo(msg.Msg)
		default:
			log.Error("wrong cmd id")
		}
	}
}

func handleHeartBeat(msg []byte){
	log.Notice("get hb")
}

func handleNserverInfo(msg []byte) {
	ninfo := &pb.Ninfo{}
	if err := proto.Unmarshal(msg, ninfo); err != nil {
		log.Fatalf("failed to parse Ninfo: ", err)
		return
	}

	log.Infof("get ninfo from ip=%s", ninfo.Ip)
}
