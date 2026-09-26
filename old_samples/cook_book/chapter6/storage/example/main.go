package main

import "github.com/dsolerh/examples/cook_book/chapter6/storage"

func main() {
	if err := storage.Exec(); err != nil {
		panic(err)
	}
}
