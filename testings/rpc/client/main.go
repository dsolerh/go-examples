package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Args struct {
	A, B int
}

type WhoArgs struct {
}

type WhoReply struct {
	Name string
}

type AppendArgs struct {
	N int
}

type AppendReply struct {
	Data []int
}

func main() {
	// Connect to the RPC server
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Error connecting to server:", err)
	}
	defer client.Close()

	// Prepare arguments
	args := &Args{A: 5, B: 3}
	var reply int

	// Call Multiply method
	err = client.Call("MathService.Multiply", args, &reply)
	if err != nil {
		log.Fatal("Error calling Multiply:", err)
	}
	fmt.Printf("5 * 3 = %d\n", reply)

	// Call Who method
	whoArgs := &WhoArgs{}
	var nameReply WhoReply
	err = client.Call("NameService.Who", whoArgs, &nameReply)
	if err != nil {
		log.Fatal("Error calling Who:", err)
	}
	fmt.Printf("Who = %s\n", nameReply.Name)

	var rr []AppendReply
	var appendReply AppendReply
	err = client.Call("MathService.Append", &AppendArgs{N: 1}, &appendReply)
	if err != nil {
		log.Fatal("Error calling Who:", err)
	}
	fmt.Printf("Data = %v\n", appendReply.Data)
	rr = append(rr, appendReply)
	err = client.Call("MathService.Append", &AppendArgs{N: 2}, &appendReply)
	if err != nil {
		log.Fatal("Error calling Who:", err)
	}
	fmt.Printf("Data = %v\n", appendReply.Data)
	rr = append(rr, appendReply)
	fmt.Printf("Data = %v\n", rr)
}
