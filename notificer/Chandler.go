package main

import (
	"encoding/gob"
	"github.com/golang/protobuf/proto"
	pb "github.com/yummyliu/Eutamias/rpc"
	"net"
)

func handleClient(con net.Conn) {
	log.Infof("get one Client, addr=%s",con.RemoteAddr().String())
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
		case pb.MsgCmd_C_LOGIN:
			handleLogin(msg, con)
		case pb.MsgCmd_C_LOGOUT:
			handleLogout(msg, con)
		case pb.MsgCmd_C_SENDMSG:
			handleSendmsg(msg, con)
		default:
			log.Error("wrong cmd id")
		}
	}
}

func writeMsgToC(cmd pb.MsgCmd, seq uint64, outmsg []byte, con net.Conn) {
	msg := pb.Message{
		Cmd : cmd,
		Seq : seq,
		Msg : outmsg,
	}
	conn_enc := gob.NewEncoder(con)
	err := conn_enc.Encode(msg)
	if err != nil {
		log.Error(err.Error())
	}
}

func handleHeartBeat(msg pb.Message, con net.Conn){
	log.Notice("get hb")
}

func handleLogin(msg pb.Message, con net.Conn) {
	login := &pb.LoginReq{}
	if err := proto.Unmarshal(msg.Msg, login); err != nil {
		log.Fatalf("failed to parse login: ", err)
		return
	}
	log.Infof("login id=%d",login.GetId())

	loginRsp := &pb.LoginRsp{
		Id : login.GetId(),
		ResultCode : pb.ResultCode_RC_OK,
	}
	outbytes, err := proto.Marshal(loginRsp);
	if err != nil {
		log.Error(err)
		return
	}

	writeMsgToC(pb.MsgCmd_C_LOGIN, msg.Seq, outbytes, con)
}

func handleLogout(msg pb.Message, con net.Conn) {
	logout := &pb.LogoutReq{}
	if err := proto.Unmarshal(msg.Msg, logout); err != nil {
		log.Fatalf("failed to parse logout: ", err)
		return
	}

	log.Infof("logout id=%d",logout.GetId())
}

func handleSendmsg(msg pb.Message, con net.Conn){
	sendmsg := &pb.SendMsgReq{}
	if err := proto.Unmarshal(msg.Msg, sendmsg); err != nil {
		log.Fatalf("failed to parse sendmsg: ", err)
		return
	}

	log.Infof("sendmsg id=%d",sendmsg.GetId())
}

