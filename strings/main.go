package main

import "fmt"

func main() {
	var s = []byte{84, 104, 101, 32, 97, 110, 115, 119, 101, 114, 32, 105, 115, 32, 53, 48, 33}
	fmt.Printf("%s", string(s))
}
