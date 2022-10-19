package main

import (
	"fmt"

	"github.com/dsolerh/examples/cook_book/chapter4/basicerrors"
)

func main() {
	basicerrors.BasicErrors()

	err := basicerrors.SomeFunc()
	fmt.Println("custom error: ", err)
}
