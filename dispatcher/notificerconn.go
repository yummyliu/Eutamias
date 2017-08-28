package main

import (
	"github.com/golang/protobuf/proto"
	pb "github.com/yummyliu/Eutamias/rpc"
	"net"
	"strconv"
)

func handleNserver(con net.Conn) {
	log.Infof("get one Ns, ip=%s",con.RemoteAddr().String())
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

	dninfo := &pb.Ninfo{}
	if err := proto.Unmarshal(conbytes, dninfo); err != nil {
		log.Fatalf("failed to parse dninfo: ", err)
	}
	log.Infof("ip= %s; port= %d; MaxConn=%d ", dninfo.Ip, dninfo.Port, dninfo.MaxConn)

	tmpns := new(Nserver)
	tmpns.Ip = dninfo.Ip
	tmpns.Port = dninfo.Port
	tmpns.MaxConn = dninfo.MaxConn
	tmpns.CurConn = dninfo.CurConn
	tmpns.Hostname = dninfo.Hostname
	NServerMap[tmpns.Ip+strconv.Itoa(int(tmpns.Port))] = *tmpns
}
