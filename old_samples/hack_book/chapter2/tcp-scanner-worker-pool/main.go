package main

import (
	"fmt"
	"sort"
	"time"
)

const (
	MIN_PORT = 1
	MAX_PORT = 150
)

type result struct {
	port int
	open bool
	time time.Duration
}

type SortByStatus []result

func (a SortByStatus) Len() int           { return len(a) }
func (a SortByStatus) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByStatus) Less(i, j int) bool { return a[i].port < a[j].port }

func main() {
	ports := make(chan int, 100)
	results := make(chan result, 5)
	var openports []result
	var closeports []result
	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for i := MIN_PORT; i <= MAX_PORT; i++ {
			ports <- i
			// fmt.Printf("--i: %v\n", i)
		}
	}()

	for i := MIN_PORT; i <= MAX_PORT; i++ {
		// fmt.Printf("i: %v\n", i)
		r := <-results
		// fmt.Printf("index: %d | r: %v\n", i, r)
		if r.open {
			openports = append(openports, r)
		} else {
			closeports = append(closeports, r)
		}
	}
	fmt.Println("scann complete")

	close(ports)
	close(results)
	sort.Sort(SortByStatus(openports))
	fmt.Println("Ports open:")
	for _, port := range openports {
		fmt.Printf("\tport: %d  | took: %dms\n", port.port, port.time.Microseconds())
	}
	fmt.Println()
	fmt.Println("Ports close:")
	for _, port := range closeports {
		fmt.Printf("\tport: %d  | took: %dms\n", port.port, port.time.Milliseconds())
	}
	fmt.Println()
}
