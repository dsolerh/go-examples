package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Printf("PopCount(108): %v\n", PopCount(999999999999))
	fmt.Printf("time.Since(start).Seconds(): %v\n", time.Since(start).Seconds())

	start = time.Now()
	fmt.Printf("PopCountShift(108): %v\n", PopCountShift(999999999999))
	fmt.Printf("time.Since(start).Seconds(): %v\n", time.Since(start).Seconds())
}
