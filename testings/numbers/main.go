package main

import (
	"fmt"
	"math"
)

func main() {
	example2()
}

func example1() {
	big := math.MaxInt64

	fmt.Printf("%64b %d\n", int64(big), int64(big))
	fmt.Printf("%64b %d\n", int32(big), int32(big))
	fmt.Printf("%64b %d\n", int16(big), int16(big))
	fmt.Printf("%64b %d\n", int8(big), int8(big))
	fmt.Printf("%64b %d\n", int(big), int(big))
}

func example2() {
	const tickRate = 4
	for _, number := range []int{1500, 2500, 500, 750} {
		val := number * tickRate / 1000
		fmt.Printf("%d * %d / 1000 = %d\n", number, tickRate, val)
	}
}
