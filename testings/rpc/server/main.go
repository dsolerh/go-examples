package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"sync"
	"time"
)

func main() {

	listener, err := net.Listen("tcp", ":12098")
	if err != nil {
		log.Fatal("Error starting listener:", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var wg sync.WaitGroup

	wg.Add(1)
	go serve(ctx, &wg, listener, echoService{"service", 5*time.Millisecond})

	fmt.Println("Press enter to stop...")
	var done string
	_, _ = fmt.Scanln(&done)

	cancel()
	// Closing the listener unblocks any goroutine waiting in Accept.
	listener.Close()
	wg.Wait()
	log.Println("Done")
}

func serve(ctx context.Context, wg *sync.WaitGroup, l net.Listener, service echoService) {
	defer wg.Done()
	server := rpc.NewServer()
	server.RegisterName(service.name, service)
	for {
		conn, err := l.Accept()
		if err != nil {
			// If we're shutting down, exit cleanly; otherwise log and continue.
			select {
			case <-ctx.Done():
				return
			default:
				log.Printf(service.name+": err: %v\n", err)
				return
			}
		}
		log.Println(service.name + ": Connection accepted")
		go server.ServeConn(conn)
	}
}

type (
	Payload struct {
		Data string
	}
	echoService struct {
		name string
		delay time.Duration
	}
)

func (s echoService) EchoDelay(args Payload, reply *Payload) error {
	time.Sleep(s.delay)
	reply.Data = s.name + args.Data
	return nil
}

func (s echoService) EchoOne(args Payload, reply *Payload) error {
	reply.Data = s.name + args.Data
	return nil
}

func (s echoService) EchoTwo(args Payload, reply *Payload) error {
	reply.Data = s.name + args.Data
	return nil
}

func (s echoService) EchoThree(args Payload, reply *Payload) error {
	reply.Data = s.name + args.Data
	return nil
}
