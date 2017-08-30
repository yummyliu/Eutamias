package main

import (
	"log"
	pb "github.com/yummyliu/Eutamias/rpc"
	"net"
	"time"
	//_ "bufio"
	"github.com/golang/protobuf/proto"
	_ "encoding/binary"
	"encoding/gob"
)

type ImClient struct {
	id uint64
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
func (c *ImClient) sendhb(conn net.Conn, delay time.Duration) {
	// send hb
	for {
		log.Println("sene hb")
		t := time.Now().Unix()//Format("2017-4-5:1:02:02\n")
		trsp := &pb.ServerTimeRsp{
			ServerTime : uint64(t),
		}
		outmsg, err := proto.Marshal(trsp);
		if err != nil {
			log.Fatalf("failed to encode ServerTimeRsp: %s", err)
			return
		}
		length,err := conn.Write(outmsg)
		if err != nil {
			log.Print(err)
			return
		}
		log.Printf("%s,w %d, %d",conn.RemoteAddr().String(),length, t)
		time.Sleep(delay)
	}
}

func (c *ImClient) sendmsg(conn net.Conn, delay time.Duration) {
	// send msg
	for {
		log.Println("sene msg")
		//pb
		t := time.Now().Unix()//Format("2017-4-5:1:02:02\n")
		trsp := &pb.ServerTimeRsp{
			ServerTime : uint64(t),
		}
		outmsg, err := proto.Marshal(trsp);
		if err != nil {
			log.Fatalf("failed to encode ServerTimeRsp: %s", err)
			return
		}
		//msg
		msg := new(pb.Message)
		msg.Cmd = 1
		msg.Seq = 0
		msg.Version = 1
		msg.Msg = outmsg
		//write to conn
		conn_enc := gob.NewEncoder(conn)
		err = conn_enc.Encode(msg)
		if err != nil {
			log.Print(err)
		}
		log.Printf("%s,",conn.RemoteAddr().String())
		time.Sleep(delay)
	}
}

func main() {
	client := new(ImClient)

	con,err := client.login()
	if err != nil {
		return
	}
	go client.sendmsg(con, 1000 * time.Millisecond)
	client.handleRev(con)
}
