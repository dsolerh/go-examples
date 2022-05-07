package server

import (
	"fmt"
	"log"
	"net"
	"net/rpc"

	"github.com/dsolerh/examples/microservices.book/basic-RPC/contract"
)

type HelloWorldHandler struct{}

func (h *HelloWorldHandler) HelloWorld(args *contract.HelloWorldRequest, reply *contract.HelloWorldResponse) error {
	reply.Message = "Hello " + args.Name
	return nil
}

func StartServer() {
	port := 3021
	helloWorld := &HelloWorldHandler{}
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
