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

func handleClient(con net.Conn) {
	log.Infof("get one client, ip=%s",con.RemoteAddr().String())
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
		err := conn_dec.Decode(&msg);
		if err != nil {
			log.Fatal(err)
			return
		}

		stime := &pb.ServerTimeRsp{}
		if err := proto.Unmarshal(msg.Msg, stime); err != nil {
			log.Fatalf("failed to parse servertime: ", err)
			return
		}

		log.Infof("get one %d %d %d %d\n",
		msg.Cmd, msg.Seq, msg.Version, stime.ServerTime)
	}
}
