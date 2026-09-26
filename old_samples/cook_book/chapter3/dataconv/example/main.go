package main

import "github.com/dsolerh/examples/cook_book/chapter3/dataconv"

func main() {
	dataconv.ShowConv()
	if err := dataconv.Strconv(); err != nil {
		panic(err)
	}
	dataconv.Interfaces()
}
