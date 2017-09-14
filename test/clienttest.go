package main

import (
	client "github.com/yummyliu/Eutamias/client"
)

func main() {
	client := new(client.ImClient)
	client.Id = 234
	client.RunCmd("127.0.0.1:54321")
}
