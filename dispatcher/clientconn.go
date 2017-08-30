package main

import (
	"bytes"
	"encoding/binary"
//	"bufio"
	"net"
	pb "github.com/yummyliu/Eutamias/rpc"
	_ "github.com/golang/protobuf/proto"
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
		conbytes := make([]byte,100)
		length,err := con.Read(conbytes);
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Infof("get %d byte",length)
		//stime := &pb.ServerTimeRsp{}
		//if err := proto.Unmarshal(conbytes[:length], stime); err != nil {
		//	log.Fatalf("failed to parse servertime: ", err)
		//	return
		//}
		buf := bytes.NewReader(conbytes[:length])
		var msg pb.Message
		err = binary.Read(buf, binary.BigEndian, &msg)
		if err != nil {
			fmt.Println("binary.read failed", err)
			return
		}
		log.Infof("get one %d %d %d %s\n", msg.Cmd, msg.Seq, msg.Version)
	}
}
