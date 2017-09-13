package main

import (
	"encoding/gob"
	"github.com/golang/protobuf/proto"
	pb "github.com/yummyliu/Eutamias/rpc"
	"net"
)

type Nserver struct {
	Ip       string
	Port     uint64
	MaxConn  uint64
	CurConn  uint64
	Hostname string
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
		case pb.MsgCmd_C_HEARTBEAT:
			handleHeartBeat(msg.Msg)
		case pb.MsgCmd_C_NINFOUPD:
			handleNserverInfoUpd(msg.Msg)
		default:
			log.Error("wrong cmd id")
		}
	}
}

func handleHeartBeat(msg []byte){
	log.Notice("get hb")
}

func handleNserverInfoUpd(msg []byte) {
	ninfo := &pb.Ninfo{}
	if err := proto.Unmarshal(msg, ninfo); err != nil {
		log.Fatalf("failed to parse Ninfo: ", err)
		return
	}
	log.Infof("req ninfo, nip:%s, nport:%d", ninfo.Ip, ninfo.Port)
	NServerMap[ninfo.GetIp()+":"+strconv.Itoa(ninfo.GetPort())] = Nserver{
		Ip : ninfo.GetIp(),
		Port : ninfo.GetPort(),
		MaxConn : ninfo.GetMaxConn(),
		CurConn : ninfo.GetCurConn(),
		Hostname : ninfo.GetHostname(),
	}
}
