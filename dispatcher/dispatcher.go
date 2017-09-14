package main

import (
	"fmt"
	"github.com/op/go-logging"
	"net"
	"strconv"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "github.com/yummyliu/Eutamias/rpc"
	"google.golang.org/grpc/reflection"
	_ "github.com/golang/protobuf/proto"
)

var (
	config          Config
	NServerMap	= make(map[string]Nserver) // key:ip+":"+port
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

type dserver struct{}

func findBestNserver() (string, uint64){
	//  TODO:  <13-09-17, complete it > //
	for _,v := range NServerMap {
		return v.Ip, v.Port
	}
	return "",0
}

func (ds *dserver) GetMyNserver(ctx context.Context, ninfoq *pb.NinfoReq) (*pb.NinfoRsp, error) {
	// construct return response
	log.Info("get on one ninfo req %d", ninfoq.GetId())
	nip, nport := findBestNserver()
	ninfoRsp := &pb.NinfoRsp{
		Id : ninfoq.GetId(),
		Nip : nip,
		Nport : nport,
		Rc : pb.ResultCode_RC_OK,
	}

	return ninfoRsp, nil
}

func main() {
	// read conf
	err := config.Read("dispatcher.ini")
	if err != nil {
		fmt.Println("config read error: ", err)
	}

	//err := init_log(config.LogFilePath)
	err = init_log("")
	if err != nil {
		log.Info("%s",err)
		return
	}

	// Listening Notificer
	go listen("0.0.0.0:"+strconv.Itoa(config.nport), handleNserver)
	// Listening Client
	lis, err := net.Listen("tcp", "0.0.0.0:"+strconv.Itoa(config.cport))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	rpcserver := grpc.NewServer()
	pb.RegisterDServiceServer(rpcserver, &dserver{})
	reflection.Register(rpcserver)
	if err := rpcserver.Serve(lis); err != nil {
		log.Fatalf("failto serve: %v", err)
	}
}
