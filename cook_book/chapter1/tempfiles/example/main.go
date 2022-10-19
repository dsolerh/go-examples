package main

import (
	"fmt"

	"github.com/dsolerh/examples/cook_book/chapter1/tempfiles"
)

func main() {
	if err := tempfiles.WorkWithTemp(); err != nil {
		fmt.Printf("err: %v\n", err)
	}
}
