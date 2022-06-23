package main

import (
	"fmt"
	"sync"
)

// worker read from a int channel of ports
// and get the work done
func worker(ports chan int, wg *sync.WaitGroup) {
	for p := range ports {
		fmt.Println("Scanning port:", p)
		wg.Done()
	}
}
