package handler

import "github.com/dsolerh/examples/microservices_book/RPC-shared/contract"

type HelloWorldHandler struct{}

func (h *HelloWorldHandler) HelloWorld(args *contract.HelloWorldRequest, reply *contract.HelloWorldResponse) error {
	reply.Message = "Hello " + args.Name
	return nil
}
