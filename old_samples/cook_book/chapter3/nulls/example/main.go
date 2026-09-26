package main

import (
	"fmt"

	"github.com/dsolerh/examples/cook_book/chapter3/nulls"
)

func main() {
	if err := nulls.BaseEncoding(); err != nil {
		panic(err)
	}
	fmt.Println()
	if err := nulls.PointerEncoding(); err != nil {
		panic(err)
	}
	fmt.Println()
	if err := nulls.NullEncoding(); err != nil {
		panic(err)
	}
}
