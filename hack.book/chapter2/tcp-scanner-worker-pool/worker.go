package main

import (
	"fmt"
	"net"
	"time"
)

// type counter struct {
// 	i  int
// 	mu sync.RWMutex
// }

// func (c *counter) Add() {
// 	c.mu.Lock()
// 	fmt.Printf("c.i: %v\n", c.i)
// 	c.i++
// 	fmt.Printf("c.i: %v\n", c.i)
// 	c.mu.Unlock()
// }

// func (c *counter) String() string {
// 	c.mu.RLock()
// 	defer c.mu.RUnlock()
// 	return fmt.Sprint(c.i)
// }

// var c *counter = &counter{}

// worker read from a int channel of ports
// and get the work done
func worker(ports chan int, r chan result) {
	for p := range ports {
		// fmt.Println("scanning port:", p)
		address := fmt.Sprintf("scanme.nmap.org:%d", p)
		conn, err := net.Dial("tcp", address)
		start := time.Now()
		if err != nil {
			// fmt.Printf("port: %d | took: %dns [closed]\n", p, time.Since(start).Nanoseconds())
			r <- result{port: p, time: time.Since(start), open: false}
			continue
		}
		conn.Close()
		// fmt.Printf("port: %d | took: %dns [open]\n", p, time.Since(start).Nanoseconds())
		r <- result{port: p, time: time.Since(start), open: true}
	}
}
