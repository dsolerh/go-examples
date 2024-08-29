package main

import "fmt"

func main() {
	BufferedExample()
}

func UnbufferedExample() {
	var ch chan []string
	ch = make(chan []string) // unbuffered channel

	go func() { ch <- []string{"a", "b"} }()
	a := <-ch

	fmt.Printf("ch: %v\n", ch)
	fmt.Printf("a: %v\n", a)
}

func BufferedExample() {
	var ch chan []string
	ch = make(chan []string, 2) // buffered channel

	// go func() {  }()
	ch <- []string{"a", "b"}
	ch <- []string{"a", "c"}
	a := <-ch

	fmt.Printf("ch: %v\n", ch)
	fmt.Printf("a: %v\n", a)
}
