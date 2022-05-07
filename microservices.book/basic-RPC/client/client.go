package client

import (
	"fmt"
	"log"
	"net/rpc"

	"github.com/dsolerh/examples/microservices.book/basic-RPC/contract"
)

func CreateClient() *rpc.Client {
	port := 3021
	client, err := rpc.Dial("tcp", fmt.Sprintf("localhost:%v", port))
	if err != nil {
		log.Fatal("dialing: ", err)
	}
	return client
}

func PerformRequest(client *rpc.Client) contract.HelloWorldResponse {
	args := &contract.HelloWorldRequest{Name: "World"}
	var reply contract.HelloWorldResponse

	err := client.Call("HelloWorldHandler.HelloWorld", args, &reply)
	if err != nil {
		log.Fatal("error: ", err)
	}
	return reply
}
