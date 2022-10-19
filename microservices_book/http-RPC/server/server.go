package server

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/dsolerh/examples/microservices.book/RPC-shared/handler"
)

func StartServer() {
	port := 3210
	helloWorld := &handler.HelloWorldHandler{}
	rpc.Register(helloWorld)
	rpc.HandleHTTP()

	l, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to listen on given port: %s", err))
	}

	log.Printf("Server starting on port: %v\n", port)
	http.Serve(l, nil)
}
