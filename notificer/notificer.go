package main

import (
	gob "encoding/gob"
	"fmt"
	"github.com/op/go-logging"
	"net"
	"strconv"
	proto "github.com/golang/protobuf/proto"
	pb "github.com/yummyliu/Eutamias/rpc"
	"time"

)

var (
	config          Config
	onlineUser		uint64
	log             *logging.Logger
	dconn			net.Conn
)
const (
	MAX_ONLINE_USER			uint64			= 1000
	UPDATENINFO_DELAY_TIME	time.Duration	= 15000
)


func listen(addr string, f func(net.Conn)) {
	log.Infof("Listen on: %s", addr)
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
	// read configuration
	err := config.Read("notificer.ini")
	if err != nil {
		fmt.Println("config read error: ", err)
	}

	// init logger
	//err := init_log(config.LogFilePath)
	err = init_log("")
	if err != nil {
		log.Info("%s",err)
		return
	}

	// connect to dispatcher
	addr := config.dIp+":"+strconv.Itoa(int(config.dPort))
	log.Info(addr)
	dconn, err = net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
		return
	}

	// register in dispatcher
	err = registerInfoUntilSuccess(UPDATENINFO_DELAY_TIME)
	if err != nil {
		log.Info(err)
		return
	}

	// Listening Client
	listen("0.0.0.0:"+strconv.Itoa(int(config.port)), handleClient) }

func registerInfoUntilSuccess(delay time.Duration) error{
	log.Info("delay: %d", delay)
	ninfo := &pb.Ninfo{
		CurConn : 0,
		MaxConn : MAX_ONLINE_USER,
		Ip : config.ip,
		Port : config.port,
	}
	ninfo.CurConn = 0
	ninfo.MaxConn = MAX_ONLINE_USER
	ninfo.Ip	= config.ip
	ninfo.Port = config.port
	outmsg, err := proto.Marshal(ninfo);
	if err != nil {
		log.Fatalf("failed to encode Ninfo: %s", err)
		return err
	}

	msg := pb.Message{
		Cmd : pb.MsgCmd_C_NINFOUPD,
		Seq : 0,
		Msg : outmsg,
	}
	conn_enc := gob.NewEncoder(dconn)
	err = conn_enc.Encode(msg)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
