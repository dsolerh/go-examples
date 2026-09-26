package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	fmt.Println(s)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	words := []string{
		"alpha",
		"beta",
		"delta",
		"gamma",
		"pi",
		"zeta",
		"eta",
		"theta",
		"epsilon",
	}

	wg.Add(len(words))

	for i, x := range words {
		go printSomething(fmt.Sprintf("%d: %s", i, x), &wg)
	}

	wg.Wait()
	// time.Sleep(1 * time.Second)

	wg.Add(1)
	printSomething("second something", &wg)
}
