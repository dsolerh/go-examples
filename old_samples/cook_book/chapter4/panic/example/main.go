package main

import (
	"fmt"

	"github.com/dsolerh/examples/cook_book/chapter4/panic"
)

func main() {
	fmt.Println("before panic")
	panic.Catcher()
	fmt.Println("after panic")
}
