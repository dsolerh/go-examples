package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
	s1 := make([]int, 0)
	s1 = append(s1, 1)
	fmt.Printf("s1: %v\n", s1)
}
