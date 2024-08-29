package main

import (
	"fmt"
	"math"
)

func main() {
	big := math.MaxInt64

	fmt.Printf("%64b %d\n", int64(big), int64(big))
	fmt.Printf("%64b %d\n", int32(big), int32(big))
	fmt.Printf("%64b %d\n", int16(big), int16(big))
	fmt.Printf("%64b %d\n", int8(big), int8(big))
	fmt.Printf("%64b %d\n", int(big), int(big))
}
