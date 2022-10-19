package main

import "github.com/dsolerh/examples/cook_book/chapter4/global"

func main() {
	if err := global.UseLog(); err != nil {
		panic(err)
	}
}
