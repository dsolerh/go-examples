package main

import "github.com/dsolerh/examples/cook.book/chapter6/pools"

func main() {
	if err := pools.ExecWithTimeout(); err != nil {
		panic(err)
	}
}
