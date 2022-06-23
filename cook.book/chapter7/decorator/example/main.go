package main

import "github.com/dsolerh/examples/cook.book/chapter7/decorator"

func main() {
	if err := decorator.Exec(); err != nil {
		panic(err)
	}
}
