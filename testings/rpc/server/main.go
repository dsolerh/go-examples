package main

import (
	"log"
	"net"
	"net/rpc"
)

// Args represents the arguments for our RPC methods
type Args struct {
	A, B int
}

// MathService provides mathematical operations
type MathService struct {
	Param int
}

// Multiply multiplies two integers
func (m *MathService) Multiply(args *Args, reply *int) error {
	*reply = m.InternalMultiply(args.A, args.B)
	log.Printf("Multiply called: %d * %d = %d", args.A, args.B, *reply)
	return nil
}

type AppendArgs struct {
	N int
}
type AppendReply struct {
	Data []int
}

func (m *MathService) Append(args *AppendArgs, reply *AppendReply) error {
	reply.Data = append(reply.Data, args.N)
	return nil
}

func (m *MathService) InternalMultiply(a, b int) int { return a*b + m.Param }

type NameService struct {
	Name string
}

type WhoArgs struct {
}
type WhoReply struct {
	Name string
}

func (n *NameService) Who(_ *WhoArgs, reply *WhoReply) error {
	reply.Name = n.Name
	return nil
}

func main() {
	// Create a new instance of our service
	mathService := &MathService{Param: 10}

	// Register the service with RPC
	err := rpc.Register(mathService)
	if err != nil {
		log.Fatal("Error registering service:", err)
	}

	err = rpc.Register(&NameService{Name: "Daniel"})
	if err != nil {
		log.Fatal("Error registering service:", err)
	}

	// Listen on TCP port 1234
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Error starting listener:", err)
	}
	defer listener.Close()

	log.Println("RPC server listening on port 1234...")

	// Accept connections and serve RPC requests
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}
		log.Println("New client connected")
		go rpc.ServeConn(conn)
	}
}
