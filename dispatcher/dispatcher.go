package main

import (
	"bufio"
	"fmt"
	"github.com/op/go-logging"
	"net"
	"os"
	"strconv"
)

var (
	clients = make(map[ClientChan]bool)
	entering = make(chan ClientChan)
	leaving = make(chan ClientChan)
	messages = make(chan string)

	config Config
	nServerInfos []Nserver
	totalOnlineUser uint64
	log *logging.Logger
)

type Nserver struct {
	Ip       string
	Ip2      string // maybe has two ip, dianxin or wangtong
	Port     int
	MaxConn  int
	CurConn  int
	Hostname string
}
type ClientChan chan<- string



func broadcaster() {
	for {
		select {
		case msg := <- messages:
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <- leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

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

func SyncNinfo() {
	for {

	}
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func handle_client(con net.Conn) {
	log.Info("get one client")
	ch := make(chan string)
	go clientWriter(con, ch)

	who := con.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- ch

	input := bufio.NewScanner(con)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}

	leaving <- ch
	messages <- who + " has left"
	con.Close()
}

func main() {
//	err := init_log(config.LogFilePath)
	err := init_log("")
	if err != nil {
		fmt.Printf("%s", err.Error())
		return
	}

	// read conf
	config.Read("dispatcher.ini")
	nServerInfos = make([]Nserver, len(config.Nips))
	for i, n := range config.Nips {
		nServerInfos[i].Ip = n
		nServerInfos[i].Port, err = strconv.Atoi(config.Nports[i])
		if err != nil {
			fmt.Println("port int ?")
			return
		}
		log.Infof("config: %s:%s:%s", config.Nips, config.Nports,config.cport)
	}

	// Listening
	listener, err := net.Listen("tcp", "0.0.0.0:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Info(err)
			continue
		}
		go handle_client(conn)
	}
}
