package main

import "github.com/dsolerh/examples/cook_book/chapter7/rest"

func main() {
	if err := rest.Exec(); err != nil {
		panic(err)
	}
}
