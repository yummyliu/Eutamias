package main

import (
	"bufio"
	"fmt"
	"github.com/op/go-logging"
	"net"
	"os"
	"strconv"
)

type Nserver struct {
	Ip       string
	Port     int
	MaxConn  int
	CurConn  int
	Hostname string
}

type ClientChan chan<- string

var (
	clients  = make(map[ClientChan]bool)
	entering = make(chan ClientChan)
	leaving  = make(chan ClientChan)
	messages = make(chan string)

	config          Config
	nServerInfos    []Nserver
	totalOnlineUser uint64
	log             *logging.Logger
)

func init_log(logFilePath string) error {
	log = logging.MustGetLogger("main")
	var format = logging.MustStringFormatter(
		`%{id:08x}--%{time}--%{level:.10s}--%{shortfile} %{message}`,
	)

	if logFilePath != "" {
		logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			return err
		}
		var backend = logging.NewLogBackend(logFile, "D:", 0)
		var backendFormatter = logging.NewBackendFormatter(backend, format)
		var backendLeveled = logging.AddModuleLevel(backendFormatter)
		backendLeveled.SetLevel(logging.DEBUG, "main")
		logging.SetBackend(backendFormatter)
	} else {
		var backend = logging.NewLogBackend(os.Stderr, "D:", 0)
		var backendFormatter = logging.NewBackendFormatter(backend, format)
		var backendLeveled = logging.AddModuleLevel(backendFormatter)
		backendLeveled.SetLevel(logging.DEBUG, "main")
		logging.SetBackend(backendFormatter)
	}
	return nil
}

func sendOut(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func handleClient(con net.Conn) {
	log.Info("get one client")
	out := make(chan string)
	go sendOut(con, out)

	who := con.RemoteAddr().String()
	out <- "You are " + who

	input := bufio.NewScanner(con)
	for input.Scan() {
		out <- "have got you msg: " + input.Text()
	}

	con.Close()
}
func handleNserver(con net.Conn) {
	log.Info("get one Nserver")
	out := make(chan string)
	go sendOut(con, out)

	who := con.RemoteAddr().String()
	out <- "You are " + who

	input := bufio.NewScanner(con)
	for input.Scan() {
		out <- "have got you msg: " + input.Text()
	}

	con.Close()
}
func Listen(addr string, f func(net.Conn)) {
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
	config.Read("dispatcher.ini")

	// Listening
	go Listen("0.0.0.0:"+strconv.Itoa(config.cport), handleClient)
	Listen("0.0.0.0:"+strconv.Itoa(config.nport), handleNserver)
}
