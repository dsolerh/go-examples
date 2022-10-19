package main

import (
	"fmt"

	"github.com/dsolerh/examples/cook.book/chapter4/errwrap"
)

func main() {
	errwrap.Wrap()
	fmt.Println()
	errwrap.Unwrap()
	fmt.Println()
	errwrap.StackTrace()
}
