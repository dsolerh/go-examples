package main

import "fmt"

func main() {
	var i uint32 = 126
	var r int32
	x := fmt.Sprintf("%b", i)
	for _, bit := range x {
		if bit == '1' {
			r++
		}
	}
	fmt.Printf("r: %v\n", r)

	fmt.Printf("(i2): %v\n", (i % 2))
}
