package main

import (
	"fmt"
	"github.com/op/go-logging"
	"net"
	"strconv"
)

type Nserver struct {
	Ip       string
	Port     int64
	MaxConn  int64
	CurConn  int64
	Hostname string
}

var (
	config          Config
	NServerMap	= make(map[string]Nserver) // key:ip+port
	totalOnlineUser uint64
	log             *logging.Logger
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
	//	err := init_log(config.LogFilePath)
	err := init_log("")
	if err != nil {
		log.Info("%s",err)
		return
	}

	// read conf
	err = config.Read("dispatcher.ini")
	if err != nil {
		fmt.Println("config read error: ", err)
	}

	// Listening Notificer
	go listen("0.0.0.0:"+strconv.Itoa(config.nport), handleNserver)
	// Listening Client
	listen("0.0.0.0:"+strconv.Itoa(config.cport), handleClient)
}
