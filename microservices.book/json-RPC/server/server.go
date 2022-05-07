package server

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"

	"github.com/dsolerh/examples/microservices.book/RPC-shared/handler"
)

type HelloWorldHandler struct{}

func StartServer() {
	port := 3210
	helloWorld := new(handler.HelloWorldHandler)
	rpc.Register(helloWorld)

	l, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to listen on given port: %s", err))
	}
	log.Printf("Server starting on port: %v\n", port)
	http.Serve(l, http.HandlerFunc(httpHandler))
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	serverCodec := jsonrpc.NewServerCodec(&HttpConn{in: r.Body, out: w})
	err := rpc.ServeRequest(serverCodec)
	if err != nil {
		log.Printf("Error while serving JSON request: %v", err)
		http.Error(w, "Error while serving JSON request, details have been logged", http.StatusInternalServerError)
		return
	}
}

type HttpConn struct {
	in  io.Reader
	out io.Writer
}

func (c *HttpConn) Read(p []byte) (n int, err error)  { return c.in.Read(p) }
func (c *HttpConn) Write(d []byte) (n int, err error) { return c.out.Write(d) }
func (c *HttpConn) Close() error                      { return nil }
