package main

import (
	"fmt"

	"github.com/dsolerh/go-test-mono/foo"
	"github.com/dsolerh/go-test-mono/publisher"
	"github.com/dsolerh/go-test-mono/utils"
)

func main() {
	fmt.Printf("utils.Version(): %v\n", utils.Version())
	fmt.Printf("publisher.Version(): %v\n", publisher.Version())
	fmt.Printf("foo.Version(): %v\n", foo.Version())
}
