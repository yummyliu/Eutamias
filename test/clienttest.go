package main

import (
	client "github.com/yummyliu/Eutamias/client"
	"log"
	"net"
)

func main() {
	client := new(client.ImClient)
	con,err := net.Dial("tcp","127.0.0.1:54321")
	if err != nil {
		log.Print("connect faild")
		return
	}

	client.RunCmd(con)
}
