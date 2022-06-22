package main

import (
	"fmt"
)

func main() {
	stringStream := make(chan string)
	doneStream := make(chan interface{})
	go example1(stringStream, doneStream)
	for s := range stringStream {
		fmt.Printf("Outside: %s\n", s)
		if s == "b" {
			close(doneStream)
			fmt.Println("the end..")
		}
	}
}

func example1(stringStream chan<- string, done chan interface{}) {
	for _, s := range []string{"a", "b", "c"} {
		select {
		case <-done:
			fmt.Println("Job Done!")
			close(stringStream)
			return
		case stringStream <- s:
			fmt.Printf("Inside: %s\n", s)
		}
	}
}
