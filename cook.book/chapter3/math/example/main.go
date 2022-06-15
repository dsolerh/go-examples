package main

import (
	"fmt"

	"github.com/dsolerh/examples/cook.book/math"
)

func main() {
	math.Examples()
	for i := 50; i < 70; i++ {
		fmt.Printf("%v ", math.Fib(i))
	}
	fmt.Println()
}
