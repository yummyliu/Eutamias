package main

import (
	"fmt"
	"github.com/op/go-logging"
	"net"
	"strconv"
	_ "github.com/golang/protobuf/proto"
	_ "github.com/yummyliu/Eutamias/rpc"
)

var (
	config          Config
	OnlineUser		uint64
	log             *logging.Logger
)
const (
	MaxUser		uint64 = 1000
)


func listen(addr string, f func(net.Conn)) {
	log.Infof("Listen on; %s", addr)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Info(err)
			continue
		}
		go f(conn)
	}
}


func main() {
	// read conf
	err := config.Read("notificer.ini")
	if err != nil {
		fmt.Println("config read error: ", err)
	}

	//err := init_log(config.LogFilePath)
	err = init_log("")
	if err != nil {
		log.Info("%s",err)
		return
	}

	// Send My info to dispatcher
	err = updateInfo()
	if err != nil {
		log.Info(err)
		return
	}

	// Listening Client
	listen("0.0.0.0:"+strconv.Itoa(int(config.port)), handleClient)
}
