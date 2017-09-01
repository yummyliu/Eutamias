package main

import (
	"encoding/gob"
	"github.com/golang/protobuf/proto"
	pb "github.com/yummyliu/Eutamias/rpc"
	"net"
	"strconv"
)

func handleDserver(con net.Conn) {
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
			handleHeartBeat(msg.Msg)
		case pb.MsgCmd_C_NINFOUPD:
			handleNinfoUpt()
		default:
			log.Error("wrong cmd id")
		}
	}
}

func handleDHeartBeat(msg []byte){
	log.Notice("get hb")
}

func handleNinfoUpt() {
	log.Info("handleNinfoUpt")
	err := updateInfo()
	if err != nil {
		log.Error(err)
		return
	}
}

func updateInfo() error{
	addr := config.dIp+strconv.Itoa(int(config.dPort))
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
		return err
	}

	ninfo := &pb.Ninfo{
		CurConn : 0,
		MaxConn : uint64(MaxUser),
		Ip : config.ip,
		Port : config.port,
	}
	ninfo.CurConn = 0
	ninfo.MaxConn = MaxUser
	ninfo.Ip	= config.ip
	ninfo.Port = config.port

	outmsg, err := proto.Marshal(ninfo);
	if err != nil {
		log.Fatalf("failed to encode Ninfo: %s", err)
		return err
	}
	_, err = conn.Write(outmsg)
	if err != nil {
		log.Info(err)
		return err
	}

	return nil
}
