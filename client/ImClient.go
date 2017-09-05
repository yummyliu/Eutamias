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
	"strconv"
)

const (
	HEARTBEAT_DURATION = 3000

	CREATESESSION = 0
	SENDMSG = 1
	LOGOUT = 2

)

type ImClient struct {

	Nconn net.Conn
	SconnList map[uint64]net.Conn

	id uint64
	online_status pb.OnlineStatus
}

func (c *ImClient) writeMsgToN(cmd pb.MsgCmd, seq uint64, outmsg []byte) {
	msg := pb.Message{
		Cmd : cmd,
		Seq : seq,
		Msg : outmsg,
	}
	conn_enc := gob.NewEncoder(c.Nconn)
	err := conn_enc.Encode(msg)
	if err != nil {
		log.Print(err)
	}
	_,err = c.Nconn.Write(outmsg)
	if err != nil {
		log.Print(err)
	}
}

func (c *ImClient) login() error {
	conn, err := net.Dial("tcp", "127.0.0.1:54321")
	if err != nil {
		log.Fatal(err)
		return err
	}

	c.Nconn = conn

	return nil
}

func (c *ImClient) handleRevFromN() {
	log.Println("handleRev---")
	for {
		conbytes := make([]byte,100)
		if _,err := c.Nconn.Read(conbytes); err != nil {
			log.Fatal(err)
			return
		}
		var msg pb.Message
		conn_dec := gob.NewDecoder(c.Nconn)
		err := conn_dec.Decode(&msg)
		if err != nil {
			log.Fatal(err)
			return
		}

		switch t := msg.Cmd; t {
		case pb.MsgCmd_C_CREATESESSION:
			c.handleCreateSession(msg.Msg)
		default:
			log.Fatal("wrong cmd id")
		}

	}
}
func (c *ImClient) handleCreateSession(msg []byte) {
	createSessionRsq := &pb.CreateSessionRsp{}
	if err := proto.Unmarshal(msg, createSessionRsq); err != nil {
		log.Fatalf("failed to parse createSessionRsq: ", err)
		return
	}
	sconn, err := net.Dial("tcp", createSessionRsq.Sip+strconv.Itoa(int(createSessionRsq.Sport)))
	if err != nil {
		log.Fatal(err)
		return
	}
	c.SconnList[createSessionRsq.Peerid] = sconn
}
func (c *ImClient) sendhbtoN(delay time.Duration) {
	for {
		log.Println("send hb")
		hb := &pb.HeartBeat{}
		outmsg, err := proto.Marshal(hb);
		if err != nil {
			log.Fatalf("failed to encode HeartBeat: %s", err)
			return
		}
		c.writeMsgToN(pb.MsgCmd_C_HEARTBEAT, 0, outmsg)
	}
}

func (c *ImClient) createSession(peerid uint64) {
	log.Println("create session")
	cs := &pb.CreateSessionReq{
		Fromid : c.id,
		Peerid : peerid,
	}
	outmsg, err := proto.Marshal(cs)
	if err != nil {
		log.Fatalf("failed to encode HeartBeat: %s", err)
		return
	}
	c.writeMsgToN(pb.MsgCmd_C_CREATESESSION, 0, outmsg)
}
func (c *ImClient) sendmsg(peerid uint64, msgdata string) {
	log.Println("send msg")
	md := &pb.MsgData{
		Id : c.id,
		SessionId : peerid,
		Content : msgdata,
	}
	outmsg, err := proto.Marshal(md)
	if err != nil {
		log.Fatalf("failed to encode MsgData: %s", err)
		return
	}
	c.writeMsgToN(pb.MsgCmd_C_CREATESESSION, 0, outmsg)
}

func (c *ImClient) logout() {
	log.Println("logout")
}

func (c *ImClient) RunCmd(conn net.Conn){
	err := c.login()
	if err != nil {
		fmt.Print(err)
		return
	}
	go c.handleRevFromN()
	go c.sendhbtoN(HEARTBEAT_DURATION)
	for {
		var cmd []string
		fmt.Scanln(&cmd)
		cc,_:= strconv.Atoi(cmd[0])
		switch cc {
		case CREATESESSION:
			peerid,_ := strconv.Atoi(cmd[1])
			c.createSession(uint64(peerid))
		case SENDMSG:
			peerid,_ := strconv.Atoi(cmd[1])
			c.sendmsg(uint64(peerid), cmd[2])
		case LOGOUT:
			c.logout()
		}
	}
}
