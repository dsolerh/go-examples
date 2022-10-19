package main

import (
	"bytes"
	"fmt"

	"github.com/dsolerh/examples/cook_book/chapter1/interfaces"
)

func main() {
	in := bytes.NewReader([]byte("example"))
	out := &bytes.Buffer{}
	fmt.Print("stdout on Copy = ")
	if err := interfaces.Copy(in, out); err != nil {
		fmt.Println(err)
	}

	fmt.Println("out bytes buffer =", out.String())
	fmt.Print("stdout on PipeExample = ")
	if err := interfaces.PipeExample(); err != nil {
		fmt.Println(err)
	}
}
