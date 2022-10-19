package server

import (
	"fmt"
	"log"
	"net"
	"net/rpc"

	"github.com/dsolerh/examples/microservices.book/RPC-shared/handler"
)

func StartServer() {
	port := 3021
	helloWorld := &handler.HelloWorldHandler{}
	rpc.Register(helloWorld)

	l, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to listen on given port: %s", err))
	}

	for {
		conn, _ := l.Accept()
		go rpc.ServeConn(conn)
	}
}

func main() {

}
