package main

import "github.com/dsolerh/examples/cook.book/chapter6/mongodb"

func main() {
	if err := mongodb.Exec("mongodb://localhost:27018"); err != nil {
		panic(err)
	}
}
