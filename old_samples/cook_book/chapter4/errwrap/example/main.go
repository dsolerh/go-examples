package main

import (
	"fmt"

	"github.com/dsolerh/examples/cook_book/chapter4/errwrap"
)

func main() {
	errwrap.Wrap()
	fmt.Println()
	errwrap.Unwrap()
	fmt.Println()
	errwrap.StackTrace()
}
