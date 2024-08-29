package main

import "fmt"

func main() {
	val := 123
	shifted := val >> 2
	fmt.Printf("val    : %10b\n", val)
	fmt.Printf("shifted: %10b\n", shifted)
	and2 := val & 2
	fmt.Printf("and2   : %10b | (%10b & %10b) = %10b\n", and2, val, 2, and2)
	and1 := val & 1
	fmt.Printf("and1   : %10b | (%10b & %10b) = %10b\n", and1, val, 1, and1)
	val = (shifted << 2) | (and2 | and1)
	fmt.Printf("val    : %10b\n", val)
}
