package main

import "github.com/dsolerh/examples/cook_book/chapter1/bytestring"

func main() {
	err := bytestring.WorkWithBuffer()
	if err != nil {
		panic(err)
	}
	// each of these print to stdout
	bytestring.SearchString()
	bytestring.ModifyString()
	bytestring.StringReader()
}
