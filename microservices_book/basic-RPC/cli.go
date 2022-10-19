package main

import (
	"fmt"

	"github.com/dsolerh/examples/microservices.book/basic-RPC/client"
	"github.com/dsolerh/examples/microservices.book/basic-RPC/server"
)

func main() {
	go server.StartServer()

	c := client.CreateClient()
	defer c.Close()

	reply := client.PerformRequest(c)
	fmt.Println(reply.Message)
}
