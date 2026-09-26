package main

import (
	"fmt"
	"sync"
)

func main() {
	var status string
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(iter int) {
			status = fmt.Sprint(iter)
			wg.Done()
		}(i)
	}
	// time.Sleep(3 * time.Second)
	wg.Wait()
	fmt.Printf("status: %v\n", status)
}
