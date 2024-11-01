package main

import (
	"fmt"

	"github.com/dsolerh/go-test-mono/first"
	"github.com/dsolerh/go-test-mono/second"
	third "github.com/dsolerh/go-test-mono/third/src"
)

func main() {
	fmt.Printf("first.Version(): %v\n", first.Version())
	fmt.Printf("second.Version(): %v\n", second.Version())
	fmt.Printf("third.Version(): %v\n", third.Version())
}
