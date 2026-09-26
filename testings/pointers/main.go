package main

import "fmt"

func main() {}

func Example1() {
	type data struct {
		Value int
	}

	v := new(data)

	fmt.Printf("v.Value: %v\n", v.Value)
}
