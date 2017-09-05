package client

import (
	"log"
	pb "github.com/yummyliu/Eutamias/rpc"
	"net"
	"time"
	"github.com/golang/protobuf/proto"
	_ "encoding/binary"
	"encoding/gob"
	"fmt"
)

type ImClient struct {
	Conn net.Conn
	Sconn net.Conn

	id uint64
	online_status pb.OnlineStatus
}

const (
	HEARTBEAT_DURATION = 3000

	SENDMSG = 1
	LOGOUT = 2
)

func (c *ImClient) writeMsg(Cmd cmd, uint64 seq, uint64 outmsg []byte) {
	msg := new(pb.Message)
	msg.Cmd = cmd
	msg.Seq = seq
	msg.Msg = outmsg

	conn_enc := gob.NewEncoder(c.Conn)
	err = conn_enc.Encode(msg)
	_,err := c.Conn.Write(outmsg)
	if err != nil {
		log.Print(err)
	}
}
func (c *ImClient) login() (net.Conn,error) {
	conn, err := net.Dial("tcp", "127.0.0.1:54321")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return conn, nil
}

func (c *ImClient) handleRev(Conn net.Conn) {
	log.Println("handleRev")
	for {
		conbytes := make([]byte,100)
		if _,err := Conn.Read(conbytes); err != nil {
			log.Fatal(err)
			return
		}
		log.Println(conbytes)
	}
}
func (c *ImClient) sendhb(delay time.Duration) {
	for {
		log.Println("send hb")
		hb := &pb.HeartBeat{}
		outmsg, err := proto.Marshal(hb);
		if err != nil {
			log.Fatalf("failed to encode HeartBeat: %s", err)
			return
		}
		c.writeMsg(pb.MsgCmd_C_HEARTBEAT, 0, outmsg)
	}
}

func (c *ImClient) sendmsg(uint64 peerid) {
	log.Println("create session")
	cs := &pb.CreateSessionReq{
		Fromid : c.id
		Peerid : peerid
	}

	for {
		log.Println("sene hb")
		hb := &pb.HeartBeat{}
		outmsg, err := proto.Marshal(hb);
		if err != nil {
			log.Fatalf("failed to encode HeartBeat: %s", err)
			return
		}

		msg := new(pb.Message)
		msg.Cmd = pb.MsgCmd_C_HEARTBEAT
		msg.Seq = 0
		msg.Version = 1
		msg.Msg = outmsg

		conn_enc := gob.NewEncoder(conn)
		err = conn_enc.Encode(msg)
		length,err := conn.Write(outmsg)
		if err != nil {
			log.Print(err)
		}
	}
}

func (c *ImClient) logout(conn net.Conn) {
}

func (c *ImClient) RunCmd(conn net.Conn){
	fmt.Println("1: send msg; 2: logout")
	c.login()
	go handleRev(conn)
	go c.sendhb(HEARTBEAT_DURATION)
	for {
		var cmd []int
		fmt.Scanln(&cmd)
		switch cmd {
		case SENDMSG:
			sendmsg(cmd[1], cmd[2])
		case LOGOUT:
			logout()
		}
	}
}
