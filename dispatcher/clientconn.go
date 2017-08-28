package main

import (
	"net"
	pb "github.com/yummyliu/Eutamias/rpc"
	"github.com/golang/protobuf/proto"
	"time"
	"fmt"
)

func sendOut(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func handleClient(con net.Conn) {
	log.Infof("get one client, ip=%s",con.RemoteAddr().String())
	defer func(con net.Conn) {
		err := con.Close()
		if err != nil {
			log.Error(err)
			return
		}
	}(con)

	conbytes := make([]byte, 100)
	_, err := con.Read(conbytes)
	if (err != nil) {
		log.Error(err);
		return
	}

	stime := &pb.ServerTimeReq{}
	if err := proto.Unmarshal(conbytes, stime); err != nil {
		log.Fatalf("failed to parse servertime: ", err)
		return
	}

	t := time.Now().Second()
	trsp := &pb.ServerTimeRsp{
		ServerTime : uint64(t),
	}
	outmsg, err := proto.Marshal(trsp);
	if err != nil {
		log.Fatalf("failed to encode ServerTimeRsp: %s", err)
		return
	}

	fmt.Fprintln(con, outmsg)
}
