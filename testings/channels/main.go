package main

import (
	"fmt"
)

type empty = struct{}

func waitNil() <-chan empty {
	return nil
}

func waitImmediate() <-chan empty {
	ch := make(chan empty, 1)
	ch <- empty{}
	return ch
}

func read(ch <-chan empty) int {
	counter := 0
	for range ch {
		counter++
	}
	return counter
}

func main() {
	select {
	case <-waitNil():
		fmt.Println("should never get here")
	case <-waitImmediate():
		fmt.Println("so perfect")
	}
}
