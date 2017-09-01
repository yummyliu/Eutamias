package main

import (
	"encoding/gob"
	"github.com/golang/protobuf/proto"
	pb "github.com/yummyliu/Eutamias/rpc"
	"net"
)

func handleClient(con net.Conn) {
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
		case pb.MsgCmd_C_LOGIN:
			handleLogin(msg.Msg)
		case pb.MsgCmd_C_LOGOUT:
			handleLogout(msg.Msg)
		case pb.MsgCmd_C_SENDMSG:
			handleSendmsg(msg.Msg)
		default:
			log.Error("wrong cmd id")
		}
	}
}

func handleHeartBeat(msg []byte){
	log.Notice("get hb")
}

func handleLogin(msg []byte) {
	login := &pb.LoginReq{}
	if err := proto.Unmarshal(msg, login); err != nil {
		log.Fatalf("failed to parse login: ", err)
		return
	}

	log.Infof("login id=%d",login.GetId())
}

func handleLogout(msg []byte) {
	logout := &pb.LogoutReq{}
	if err := proto.Unmarshal(msg, logout); err != nil {
		log.Fatalf("failed to parse logout: ", err)
		return
	}

	log.Infof("logout id=%d",logout.GetId())
}

func handleSendmsg(msg []byte){
	sendmsg := &pb.SendMsgReq{}
	if err := proto.Unmarshal(msg, sendmsg); err != nil {
		log.Fatalf("failed to parse sendmsg: ", err)
		return
	}

	log.Infof("sendmsg id=%d",sendmsg.GetId())
}

