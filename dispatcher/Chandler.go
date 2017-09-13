package main

import (
	"encoding/gob"
	"net"
	pb "github.com/yummyliu/Eutamias/rpc"
	"github.com/golang/protobuf/proto"
)

func handleClient(con net.Conn) {
	log.Infof("client connect, addr=%s",con.RemoteAddr().String())
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
			log.Debug(err.Error(),": ", msg.Cmd)
			return
		}

		switch t := msg.Cmd; t {
		case pb.MsgCmd_C_NINFOREQ:
			handleNserverInfoReq(msg.Msg, con)
		default:
			log.Error("wrong cmd id")
		}
	}
}

func findBestNserver() (string, uint64){
	//  TODO:  <13-09-17, complete it > //
	for _,v := range NServerMap {
		return v.Ip, v.Port
	}
	return "",0
}

func handleNserverInfoReq(msg []byte, con net.Conn) {
	ninfoq := &pb.NinfoReq{}
	if err := proto.Unmarshal(msg, ninfoq); err != nil {
		log.Fatalf(err.Error())
		return
	}
	log.Infof("req ninfo from client=%d", ninfoq.Id)

	// construct return response
	nip, nport := findBestNserver()
	ninfoRsp := &pb.NinfoRsp{
		Id : ninfoq.GetId(),
		Nip : nip,
		Nport : nport,
		Rc : pb.ResultCode_RC_OK,
	}

	// send back to C
	outmsg, err := proto.Marshal(ninfoRsp)
	if err != nil {
		log.Fatalf("failed to encode MsgData: %s", err)
		return
	}
	writeMsgToC(pb.MsgCmd_C_NINFOREQ, 0, outmsg, con)
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
