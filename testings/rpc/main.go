package main

import (
	"context"
	"log"
	"net"
	"net/rpc"
	"sync"
	"sync/atomic"
	"time"
)

type Payload struct {
	Data string
}

type service struct {
	wait chan struct{}
}

func (service) name() string { return "service" }

func (s *service) EchoWaitRpc(args Payload, reply *Payload) error {
	log.Printf("(EchoWaitRpc) args.Data: %v\n", args.Data)
	<-s.wait
	reply.Data = args.Data
	log.Printf("(EchoWaitRpc) reply.Data: %v\n", reply.Data)
	return nil
}

func (s *service) EchoRpc(args Payload, reply *Payload) error {
	log.Printf("(EchoRpc) args.Data: %v\n", args.Data)
	reply.Data = args.Data
	log.Printf("(EchoRpc) reply.Data: %v\n", reply.Data)
	return nil
}

func main() {
	// create
	li, err := net.Listen("tcp", ":12098")
	if err != nil {
		log.Fatal("Error starting listener:", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	service := &service{make(chan struct{})}
	server := rpc.NewServer()
	err = server.RegisterName(service.name(), service)
	if err != nil {
		log.Fatal("Error registering service:", err)
	}

	go func() {
		for {
			conn, err := li.Accept()
			if err != nil {
				// If we're shutting down, exit cleanly; otherwise log and continue.
				select {
				case <-ctx.Done():
					return
				default:
					log.Printf(service.name()+": err: %v\n", err)
					return
				}
			}
			log.Println(service.name() + ": Connection accepted")
			go server.ServeConn(conn)
		}
	}()

	client, err := rpc.Dial("tcp", "localhost:12098")
	if err != nil {
		log.Fatal("Error connecting to server:", err)
	}

	var wg sync.WaitGroup
	var total atomic.Uint32
	var completed atomic.Uint32
	var timeout atomic.Uint32
	toFail := 100_000
	wg.Add(toFail)
	total.Add(uint32(toFail))
	for _ = range toFail {
		go func() {
			defer wg.Done()
			success := callWaitTimeout(client, 5*time.Millisecond)
			if success {
				completed.Add(1)
			} else {
				timeout.Add(1)
			}
		}()
	}

	toSucceed := 100_000
	wg.Add(toSucceed)
	total.Add(uint32(toSucceed))
	for _ = range toSucceed {
		go func() {
			defer wg.Done()
			success := callTimeout(client, 50*time.Millisecond)
			if success {
				completed.Add(1)
			} else {
				timeout.Add(1)
			}
		}()
	}

	wg.Wait()
	log.Printf("total: %v\n", total.Load())
	log.Printf("completed: %v\n", completed.Load())
	log.Printf("timeout: %v\n", timeout.Load())
	time.Sleep(time.Second)
}

func callWaitTimeout(client *rpc.Client, timeout time.Duration) bool {
	args := Payload{Data: "Daniel"}
	var reply Payload

	done := make(chan *rpc.Call, 1)
	// start := time.Now()
	call := client.Go("service.EchoWaitRpc", args, &reply, done)
	select {
	case <-call.Done:
		// log.Printf("call.Error: %v\n", call.Error)
		// log.Printf("call.Reply: %v\n", call.Reply)
		return true
	case <-time.After(timeout):
		// log.Println("service.EchoWaitRpc timeout: %s", time.Since(start))
		return false
	}
}

func callTimeout(client *rpc.Client, timeout time.Duration) bool {
	args := Payload{Data: "Daniel"}
	var reply Payload

	// start := time.Now()
	call := client.Go("service.EchoRpc", args, &reply, make(chan *rpc.Call, 1))
	select {
	case <-call.Done:
		// log.Printf("call.Error: %v\n", call.Error)
		// log.Printf("call.Reply: %v\n", call.Reply)
		return true
	case <-time.After(timeout):
		// log.Println("service.EchoRpc timeout: %s", time.Since(start))
		return false
	}
}
