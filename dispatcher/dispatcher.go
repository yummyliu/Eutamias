package main

import (
	//	"flag"
	"fmt"
	"github.com/op/go-logging"
	"net"
	"os"
	"strconv"
)

var log *logging.Logger

func init_log(logFilePath string) error {
	log = logging.MustGetLogger("main")
	var format = logging.MustStringFormatter(
		`%{id:08x}--%{time}--%{level:.10s}--%{shortfile} %{message}`,
	)

	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	var backend = logging.NewLogBackend(logFile, "D:", 0)
	var backendFormatter = logging.NewBackendFormatter(backend, format)
	var backendLeveled = logging.AddModuleLevel(backendFormatter)
	backendLeveled.SetLevel(logging.DEBUG, "main")

	logging.SetBackend(backendFormatter)

	log.Info("asd")
	return nil
}

type Nserver struct {
	Ip       string
	Ip2      string // maybe has two ip, dianxin or wangtong
	port     int
	MaxConn  int
	CurConn  int
	hostname string
}

var nServerInfos []Nserver
var totalOnlineUser uint64

func SyncNinfo() {
	for {

	}
}

func handle_client(con net.Conn) {
	log.Info("get one client")
}

func Listen(f func(net.Conn), port int) {
	TCPService(fmt.Sprintf("0.0.0.0:%d", port), f)
}

var config *Config

func main() {
	//
	//	flag.Parse()
	//	if len(flag.Args()) == 0 {
	//		fmt.Println("usage: dispatcher config")
	//		return
	//	}
	//
	config.Read("dispatcher.ini")
	err := init_log(config.LogFilePath)
	if err != nil {
		fmt.Printf("%s", err.Error())
		return
	}

	nServerInfos = make([]Nserver, len(config.Nips))
	for i, n := range config.Nips {
		nServerInfos[i].Ip = n
		nServerInfos[i].port, err = strconv.Atoi(config.Nports[i])
		if err != nil {
			fmt.Println("port int ?")
			return
		}
		log.Infof("config: %s:%s:%s", config.Nips, config.Nports)
	}

	Listen(handle_client, 54321)
	Wait()
}
