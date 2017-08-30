package main

import (
	"log"
	"fmt"
	pb "github.com/yummyliu/Eutamias/rpc"
	"net"
	"time"
	//_ "bufio"
	"github.com/golang/protobuf/proto"
	"bytes"
	"encoding/binary"
)

type ImClient struct {
	id uint64
}

func (c *ImClient) GetTime() uint64 {
	timereq := &pb.ServerTimeReq{}
	conn, err := net.Dial("tcp", "127.0.0.1:54321")
	if err != nil {
		log.Fatal(err)
	}
	defer func (c *net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Print("err")
		}
	}(&conn)

	fmt.Fprintf(conn, timereq.String())

	return 0
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

		msg := new(pb.Message)
		msg.Cmd = 1
		msg.Seq = 0
		msg.Version = 1
//		msg.Msg = "hello"

		buf := new(bytes.Buffer)
		err := binary.Write(buf, binary.BigEndian, msg)
		if err != nil {
			log.Println(err.Error())
			continue
		}
		outmsg := buf.Bytes()

		length,err := conn.Write(outmsg)
		if err != nil {
			log.Print(err)
		}
		log.Printf("%s,w %d",conn.RemoteAddr().String(),length)
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
