package main

import (
	"fmt"

	"github.com/dsolerh/examples/cook_book/chapter3/tags"
)

func main() {
	if err := tags.EmptyStruct(); err != nil {
		panic(err)
	}
	fmt.Println()
	if err := tags.FullStruct(); err != nil {
		panic(err)
	}
}
