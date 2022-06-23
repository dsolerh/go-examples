package main

import (
	"fmt"
	"net"
)

func main() {
	for i := 1; i <= 1024; i++ {
		conn, err := net.Dial("tcp", fmt.Sprintf("scanme.nmap.org:%d", i))
		if err != nil {
			fmt.Printf("port: %d [closed]\n", i)
			continue
		}
		conn.Close()
		fmt.Printf("port: %d [open]\n", i)
	}
}
