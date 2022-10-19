package main

import "github.com/dsolerh/examples/cook.book/chapter6/storage"

func main() {
	if err := storage.Exec(); err != nil {
		panic(err)
	}
}
